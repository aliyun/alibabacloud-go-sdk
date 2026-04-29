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
		"cn-qingdao":                  dara.String("polardb.aliyuncs.com"),
		"cn-beijing":                  dara.String("polardb.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("polardb.aliyuncs.com"),
		"cn-hangzhou":                 dara.String("polardb.aliyuncs.com"),
		"cn-shanghai":                 dara.String("polardb.aliyuncs.com"),
		"cn-heyuan":                   dara.String("polardb.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("polardb.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("polardb.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("polardb.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("polardb.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("polardb.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("polardb.aliyuncs.com"),
		"cn-edge-1":                   dara.String("polardb.aliyuncs.com"),
		"cn-fujian":                   dara.String("polardb.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("polardb.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("polardb.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("polardb.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("polardb.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("polardb.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("polardb.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("polardb.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("polardb.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("polardb.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("polardb.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("polardb.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("polardb.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("polardb.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("polardb.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("polardb.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("polardb.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("polardb.aliyuncs.com"),
		"cn-wuhan":                    dara.String("polardb.aliyuncs.com"),
		"cn-yushanfang":               dara.String("polardb.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("polardb.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("polardb.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("polardb.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("polardb.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("polardb.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("polardb.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("polardb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - AbortDBClusterMigrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbortDBClusterMigrationResponse
func (client *Client) AbortDBClusterMigrationWithOptions(request *AbortDBClusterMigrationRequest, runtime *dara.RuntimeOptions) (_result *AbortDBClusterMigrationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AbortDBClusterMigration"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AbortDBClusterMigrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AbortDBClusterMigrationRequest
//
// @return AbortDBClusterMigrationResponse
func (client *Client) AbortDBClusterMigration(request *AbortDBClusterMigrationRequest) (_result *AbortDBClusterMigrationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AbortDBClusterMigrationResponse{}
	_body, _err := client.AbortDBClusterMigrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加角色权限
//
// @param request - AddEncryptionDBRolePrivilegeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddEncryptionDBRolePrivilegeResponse
func (client *Client) AddEncryptionDBRolePrivilegeWithOptions(request *AddEncryptionDBRolePrivilegeRequest, runtime *dara.RuntimeOptions) (_result *AddEncryptionDBRolePrivilegeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.RolePrivilegeConfig) {
		query["RolePrivilegeConfig"] = request.RolePrivilegeConfig
	}

	if !dara.IsNil(request.RolePrivilegeName) {
		query["RolePrivilegeName"] = request.RolePrivilegeName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddEncryptionDBRolePrivilege"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddEncryptionDBRolePrivilegeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加角色权限
//
// @param request - AddEncryptionDBRolePrivilegeRequest
//
// @return AddEncryptionDBRolePrivilegeResponse
func (client *Client) AddEncryptionDBRolePrivilege(request *AddEncryptionDBRolePrivilegeRequest) (_result *AddEncryptionDBRolePrivilegeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddEncryptionDBRolePrivilegeResponse{}
	_body, _err := client.AddEncryptionDBRolePrivilegeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加sql防火墙配置
//
// @param request - AddFirewallRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFirewallRulesResponse
func (client *Client) AddFirewallRulesWithOptions(request *AddFirewallRulesRequest, runtime *dara.RuntimeOptions) (_result *AddFirewallRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.RuleConfig) {
		query["RuleConfig"] = request.RuleConfig
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFirewallRules"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFirewallRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加sql防火墙配置
//
// @param request - AddFirewallRulesRequest
//
// @return AddFirewallRulesResponse
func (client *Client) AddFirewallRules(request *AddFirewallRulesRequest) (_result *AddFirewallRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddFirewallRulesResponse{}
	_body, _err := client.AddFirewallRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加PolarClaw MCP Server
//
// @param tmpReq - AddPolarClawMCPServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPolarClawMCPServerResponse
func (client *Client) AddPolarClawMCPServerWithOptions(tmpReq *AddPolarClawMCPServerRequest, runtime *dara.RuntimeOptions) (_result *AddPolarClawMCPServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddPolarClawMCPServerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ServerConfig) {
		request.ServerConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServerConfig, dara.String("ServerConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ServerConfigShrink) {
		query["ServerConfig"] = request.ServerConfigShrink
	}

	if !dara.IsNil(request.ServerName) {
		query["ServerName"] = request.ServerName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddPolarClawMCPServer"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddPolarClawMCPServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加PolarClaw MCP Server
//
// @param request - AddPolarClawMCPServerRequest
//
// @return AddPolarClawMCPServerResponse
func (client *Client) AddPolarClawMCPServer(request *AddPolarClawMCPServerRequest) (_result *AddPolarClawMCPServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddPolarClawMCPServerResponse{}
	_body, _err := client.AddPolarClawMCPServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增polarfs bucket路径
//
// @param request - AddPolarFsPathMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPolarFsPathMappingResponse
func (client *Client) AddPolarFsPathMappingWithOptions(request *AddPolarFsPathMappingRequest, runtime *dara.RuntimeOptions) (_result *AddPolarFsPathMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomBucketPathList) {
		query["CustomBucketPathList"] = request.CustomBucketPathList
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.PolarFsInstanceId) {
		query["PolarFsInstanceId"] = request.PolarFsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddPolarFsPathMapping"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddPolarFsPathMappingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增polarfs bucket路径
//
// @param request - AddPolarFsPathMappingRequest
//
// @return AddPolarFsPathMappingResponse
func (client *Client) AddPolarFsPathMapping(request *AddPolarFsPathMappingRequest) (_result *AddPolarFsPathMappingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddPolarFsPathMappingResponse{}
	_body, _err := client.AddPolarFsPathMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增PolarFs Quota规则
//
// @param request - AddPolarFsQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPolarFsQuotaResponse
func (client *Client) AddPolarFsQuotaWithOptions(request *AddPolarFsQuotaRequest, runtime *dara.RuntimeOptions) (_result *AddPolarFsQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.PolarFsInstanceId) {
		query["PolarFsInstanceId"] = request.PolarFsInstanceId
	}

	if !dara.IsNil(request.Quotas) {
		query["Quotas"] = request.Quotas
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddPolarFsQuota"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddPolarFsQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增PolarFs Quota规则
//
// @param request - AddPolarFsQuotaRequest
//
// @return AddPolarFsQuotaResponse
func (client *Client) AddPolarFsQuota(request *AddPolarFsQuotaRequest) (_result *AddPolarFsQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddPolarFsQuotaResponse{}
	_body, _err := client.AddPolarFsQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加SQL限流规则
//
// @param request - AddSQLRateLimitingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSQLRateLimitingRulesResponse
func (client *Client) AddSQLRateLimitingRulesWithOptions(request *AddSQLRateLimitingRulesRequest, runtime *dara.RuntimeOptions) (_result *AddSQLRateLimitingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.RuleConfig) {
		query["RuleConfig"] = request.RuleConfig
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSQLRateLimitingRules"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSQLRateLimitingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加SQL限流规则
//
// @param request - AddSQLRateLimitingRulesRequest
//
// @return AddSQLRateLimitingRulesResponse
func (client *Client) AddSQLRateLimitingRules(request *AddSQLRateLimitingRulesRequest) (_result *AddSQLRateLimitingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddSQLRateLimitingRulesResponse{}
	_body, _err := client.AddSQLRateLimitingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 应用提示词策略到实例
//
// @param tmpReq - ApplyApplicationPromptsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyApplicationPromptsResponse
func (client *Client) ApplyApplicationPromptsWithOptions(tmpReq *ApplyApplicationPromptsRequest, runtime *dara.RuntimeOptions) (_result *ApplyApplicationPromptsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ApplyApplicationPromptsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DisabledPromptIds) {
		request.DisabledPromptIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DisabledPromptIds, dara.String("DisabledPromptIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.EnabledPromptIds) {
		request.EnabledPromptIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EnabledPromptIds, dara.String("EnabledPromptIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.DisabledPromptIdsShrink) {
		query["DisabledPromptIds"] = request.DisabledPromptIdsShrink
	}

	if !dara.IsNil(request.EnabledPromptIdsShrink) {
		query["EnabledPromptIds"] = request.EnabledPromptIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyApplicationPrompts"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyApplicationPromptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用提示词策略到实例
//
// @param request - ApplyApplicationPromptsRequest
//
// @return ApplyApplicationPromptsResponse
func (client *Client) ApplyApplicationPrompts(request *ApplyApplicationPromptsRequest) (_result *ApplyApplicationPromptsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyApplicationPromptsResponse{}
	_body, _err := client.ApplyApplicationPromptsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 挂载PolarFS到PolarDB应用
//
// @param request - AttachApplicationPolarFSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachApplicationPolarFSResponse
func (client *Client) AttachApplicationPolarFSWithOptions(request *AttachApplicationPolarFSRequest, runtime *dara.RuntimeOptions) (_result *AttachApplicationPolarFSResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.PolarFSAccessKeyId) {
		query["PolarFSAccessKeyId"] = request.PolarFSAccessKeyId
	}

	if !dara.IsNil(request.PolarFSAccessKeySecret) {
		query["PolarFSAccessKeySecret"] = request.PolarFSAccessKeySecret
	}

	if !dara.IsNil(request.PolarFSInstanceId) {
		query["PolarFSInstanceId"] = request.PolarFSInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachApplicationPolarFS"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachApplicationPolarFSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 挂载PolarFS到PolarDB应用
//
// @param request - AttachApplicationPolarFSRequest
//
// @return AttachApplicationPolarFSResponse
func (client *Client) AttachApplicationPolarFS(request *AttachApplicationPolarFSRequest) (_result *AttachApplicationPolarFSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachApplicationPolarFSResponse{}
	_body, _err := client.AttachApplicationPolarFSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 绑定PolarClaw Agent
//
// @param request - BindPolarClawAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindPolarClawAgentResponse
func (client *Client) BindPolarClawAgentWithOptions(request *BindPolarClawAgentRequest, runtime *dara.RuntimeOptions) (_result *BindPolarClawAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.ChannelAccountId) {
		query["ChannelAccountId"] = request.ChannelAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindPolarClawAgent"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindPolarClawAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定PolarClaw Agent
//
// @param request - BindPolarClawAgentRequest
//
// @return BindPolarClawAgentResponse
func (client *Client) BindPolarClawAgent(request *BindPolarClawAgentRequest) (_result *BindPolarClawAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindPolarClawAgentResponse{}
	_body, _err := client.BindPolarClawAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels O\\\\\\&M events at a time.
//
// @param request - CancelActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelActiveOperationTasksResponse
func (client *Client) CancelActiveOperationTasksWithOptions(request *CancelActiveOperationTasksRequest, runtime *dara.RuntimeOptions) (_result *CancelActiveOperationTasksResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TaskIds) {
		query["TaskIds"] = request.TaskIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelActiveOperationTasks"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelActiveOperationTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels O\\\\\\&M events at a time.
//
// @param request - CancelActiveOperationTasksRequest
//
// @return CancelActiveOperationTasksResponse
func (client *Client) CancelActiveOperationTasks(request *CancelActiveOperationTasksRequest) (_result *CancelActiveOperationTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelActiveOperationTasksResponse{}
	_body, _err := client.CancelActiveOperationTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消周期任务策略
//
// @param request - CancelCronJobPolicyServerlessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelCronJobPolicyServerlessResponse
func (client *Client) CancelCronJobPolicyServerlessWithOptions(request *CancelCronJobPolicyServerlessRequest, runtime *dara.RuntimeOptions) (_result *CancelCronJobPolicyServerlessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
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
		Action:      dara.String("CancelCronJobPolicyServerless"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelCronJobPolicyServerlessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消周期任务策略
//
// @param request - CancelCronJobPolicyServerlessRequest
//
// @return CancelCronJobPolicyServerlessResponse
func (client *Client) CancelCronJobPolicyServerless(request *CancelCronJobPolicyServerlessRequest) (_result *CancelCronJobPolicyServerlessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelCronJobPolicyServerlessResponse{}
	_body, _err := client.CancelCronJobPolicyServerlessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消目录的配额
//
// @param request - CancelPolarFsFileQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelPolarFsFileQuotaResponse
func (client *Client) CancelPolarFsFileQuotaWithOptions(request *CancelPolarFsFileQuotaRequest, runtime *dara.RuntimeOptions) (_result *CancelPolarFsFileQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.FilePathIds) {
		query["FilePathIds"] = request.FilePathIds
	}

	if !dara.IsNil(request.PolarFsInstanceId) {
		query["PolarFsInstanceId"] = request.PolarFsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelPolarFsFileQuota"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelPolarFsFileQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消目录的配额
//
// @param request - CancelPolarFsFileQuotaRequest
//
// @return CancelPolarFsFileQuotaResponse
func (client *Client) CancelPolarFsFileQuota(request *CancelPolarFsFileQuotaRequest) (_result *CancelPolarFsFileQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelPolarFsFileQuotaResponse{}
	_body, _err := client.CancelPolarFsFileQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels scheduled tasks that are not yet started.
//
// @param request - CancelScheduleTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelScheduleTasksResponse
func (client *Client) CancelScheduleTasksWithOptions(request *CancelScheduleTasksRequest, runtime *dara.RuntimeOptions) (_result *CancelScheduleTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelScheduleTasks"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelScheduleTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels scheduled tasks that are not yet started.
//
// @param request - CancelScheduleTasksRequest
//
// @return CancelScheduleTasksResponse
func (client *Client) CancelScheduleTasks(request *CancelScheduleTasksRequest) (_result *CancelScheduleTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelScheduleTasksResponse{}
	_body, _err := client.CancelScheduleTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether an account name is valid or unique in a cluster.
//
// @param request - CheckAccountNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAccountNameResponse
func (client *Client) CheckAccountNameWithOptions(request *CheckAccountNameRequest, runtime *dara.RuntimeOptions) (_result *CheckAccountNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckAccountName"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckAccountNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether an account name is valid or unique in a cluster.
//
// @param request - CheckAccountNameRequest
//
// @return CheckAccountNameResponse
func (client *Client) CheckAccountName(request *CheckAccountNameRequest) (_result *CheckAccountNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckAccountNameResponse{}
	_body, _err := client.CheckAccountNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检查PolarDB账号名称
//
// @param request - CheckAccountNameZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAccountNameZonalResponse
func (client *Client) CheckAccountNameZonalWithOptions(request *CheckAccountNameZonalRequest, runtime *dara.RuntimeOptions) (_result *CheckAccountNameZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckAccountNameZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckAccountNameZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查PolarDB账号名称
//
// @param request - CheckAccountNameZonalRequest
//
// @return CheckAccountNameZonalResponse
func (client *Client) CheckAccountNameZonal(request *CheckAccountNameZonalRequest) (_result *CheckAccountNameZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckAccountNameZonalResponse{}
	_body, _err := client.CheckAccountNameZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检查连接串
//
// @param request - CheckConnectionStringRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckConnectionStringResponse
func (client *Client) CheckConnectionStringWithOptions(request *CheckConnectionStringRequest, runtime *dara.RuntimeOptions) (_result *CheckConnectionStringResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckConnectionString"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckConnectionStringResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查连接串
//
// @param request - CheckConnectionStringRequest
//
// @return CheckConnectionStringResponse
func (client *Client) CheckConnectionString(request *CheckConnectionStringRequest) (_result *CheckConnectionStringResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckConnectionStringResponse{}
	_body, _err := client.CheckConnectionStringWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether a database name is valid or whether the name is already used by another database in the current cluster.
//
// @param request - CheckDBNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDBNameResponse
func (client *Client) CheckDBNameWithOptions(request *CheckDBNameRequest, runtime *dara.RuntimeOptions) (_result *CheckDBNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDBName"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDBNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a database name is valid or whether the name is already used by another database in the current cluster.
//
// @param request - CheckDBNameRequest
//
// @return CheckDBNameResponse
func (client *Client) CheckDBName(request *CheckDBNameRequest) (_result *CheckDBNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckDBNameResponse{}
	_body, _err := client.CheckDBNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检查PolarDB边缘云数据库名
//
// @param request - CheckDBNameZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDBNameZonalResponse
func (client *Client) CheckDBNameZonalWithOptions(request *CheckDBNameZonalRequest, runtime *dara.RuntimeOptions) (_result *CheckDBNameZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDBNameZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDBNameZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查PolarDB边缘云数据库名
//
// @param request - CheckDBNameZonalRequest
//
// @return CheckDBNameZonalResponse
func (client *Client) CheckDBNameZonal(request *CheckDBNameZonalRequest) (_result *CheckDBNameZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckDBNameZonalResponse{}
	_body, _err := client.CheckDBNameZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether the cluster is authorized to use Key Management Service (KMS).
//
// @param request - CheckKMSAuthorizedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckKMSAuthorizedResponse
func (client *Client) CheckKMSAuthorizedWithOptions(request *CheckKMSAuthorizedRequest, runtime *dara.RuntimeOptions) (_result *CheckKMSAuthorizedResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.TDERegion) {
		query["TDERegion"] = request.TDERegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckKMSAuthorized"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckKMSAuthorizedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the cluster is authorized to use Key Management Service (KMS).
//
// @param request - CheckKMSAuthorizedRequest
//
// @return CheckKMSAuthorizedResponse
func (client *Client) CheckKMSAuthorized(request *CheckKMSAuthorizedRequest) (_result *CheckKMSAuthorizedResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckKMSAuthorizedResponse{}
	_body, _err := client.CheckKMSAuthorizedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用于检查PolarFS实例中配额设置的一致性状态。
//
// Description:
//
// ## 请求说明
//
// 该API允许用户验证指定PolarFS实例内的配额配置是否一致，包括但不限于目录路径上的存储容量和inode限制。如果存在不一致的情况，将返回具体的不一致路径列表及可能的错误信息。
//
// ### 注意事项
//
// - 确保`PolarFsInstanceId`参数正确无误地指向了目标PolarFS实例。
//
// - 当系统检测到配额不一致时，除了返回`IsConsistent=false`外，还会提供`InconsistentPaths`数组来指示具体哪些路径存在问题。
//
// - 如果请求成功但没有发现任何不一致，则`InconsistentPaths`为空数组，并且`IsConsistent=true`。
//
// - 错误处理：若请求过程中遇到权限不足、资源不存在等问题，请参考提供的错误码定义部分以获取更详细的错误信息。
//
// @param request - CheckPolarFsQuotaConsistencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckPolarFsQuotaConsistencyResponse
func (client *Client) CheckPolarFsQuotaConsistencyWithOptions(request *CheckPolarFsQuotaConsistencyRequest, runtime *dara.RuntimeOptions) (_result *CheckPolarFsQuotaConsistencyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableRepair) {
		query["EnableRepair"] = request.EnableRepair
	}

	if !dara.IsNil(request.EnableStrictCalculate) {
		query["EnableStrictCalculate"] = request.EnableStrictCalculate
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.PolarFsInstanceId) {
		query["PolarFsInstanceId"] = request.PolarFsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckPolarFsQuotaConsistency"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckPolarFsQuotaConsistencyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于检查PolarFS实例中配额设置的一致性状态。
//
// Description:
//
// ## 请求说明
//
// 该API允许用户验证指定PolarFS实例内的配额配置是否一致，包括但不限于目录路径上的存储容量和inode限制。如果存在不一致的情况，将返回具体的不一致路径列表及可能的错误信息。
//
// ### 注意事项
//
// - 确保`PolarFsInstanceId`参数正确无误地指向了目标PolarFS实例。
//
// - 当系统检测到配额不一致时，除了返回`IsConsistent=false`外，还会提供`InconsistentPaths`数组来指示具体哪些路径存在问题。
//
// - 如果请求成功但没有发现任何不一致，则`InconsistentPaths`为空数组，并且`IsConsistent=true`。
//
// - 错误处理：若请求过程中遇到权限不足、资源不存在等问题，请参考提供的错误码定义部分以获取更详细的错误信息。
//
// @param request - CheckPolarFsQuotaConsistencyRequest
//
// @return CheckPolarFsQuotaConsistencyResponse
func (client *Client) CheckPolarFsQuotaConsistency(request *CheckPolarFsQuotaConsistencyRequest) (_result *CheckPolarFsQuotaConsistencyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckPolarFsQuotaConsistencyResponse{}
	_body, _err := client.CheckPolarFsQuotaConsistencyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether a service-linked role (SLR) is created.
//
// @param request - CheckServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckServiceLinkedRoleResponse
func (client *Client) CheckServiceLinkedRoleWithOptions(request *CheckServiceLinkedRoleRequest, runtime *dara.RuntimeOptions) (_result *CheckServiceLinkedRoleResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckServiceLinkedRole"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a service-linked role (SLR) is created.
//
// @param request - CheckServiceLinkedRoleRequest
//
// @return CheckServiceLinkedRoleResponse
func (client *Client) CheckServiceLinkedRole(request *CheckServiceLinkedRoleRequest) (_result *CheckServiceLinkedRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.CheckServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 支持基础版支持clone文件或目录快照
//
// @param request - ClonePolarFsBasicSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClonePolarFsBasicSnapshotResponse
func (client *Client) ClonePolarFsBasicSnapshotWithOptions(request *ClonePolarFsBasicSnapshotRequest, runtime *dara.RuntimeOptions) (_result *ClonePolarFsBasicSnapshotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.PolarFsInstanceId) {
		query["PolarFsInstanceId"] = request.PolarFsInstanceId
	}

	if !dara.IsNil(request.SourcePath) {
		query["SourcePath"] = request.SourcePath
	}

	if !dara.IsNil(request.TargetPath) {
		query["TargetPath"] = request.TargetPath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClonePolarFsBasicSnapshot"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClonePolarFsBasicSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 支持基础版支持clone文件或目录快照
//
// @param request - ClonePolarFsBasicSnapshotRequest
//
// @return ClonePolarFsBasicSnapshotResponse
func (client *Client) ClonePolarFsBasicSnapshot(request *ClonePolarFsBasicSnapshotRequest) (_result *ClonePolarFsBasicSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClonePolarFsBasicSnapshotResponse{}
	_body, _err := client.ClonePolarFsBasicSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the PolarDB for AI feature for a cluster.
//
// @param request - CloseAITaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseAITaskResponse
func (client *Client) CloseAITaskWithOptions(request *CloseAITaskRequest, runtime *dara.RuntimeOptions) (_result *CloseAITaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("CloseAITask"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseAITaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the PolarDB for AI feature for a cluster.
//
// @param request - CloseAITaskRequest
//
// @return CloseAITaskResponse
func (client *Client) CloseAITask(request *CloseAITaskRequest) (_result *CloseAITaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloseAITaskResponse{}
	_body, _err := client.CloseAITaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels or completes the migration task that upgrades an RDS cluster to a PolarDB cluster.
//
// Description:
//
//	  You can call this operation to cancel the migration task before data migration.
//
//		- You can call this operation to perform the migration task after data migration.
//
// > Before you call this operation, ensure that a one-click upgrade task has been created for the cluster. You can call the [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) operation to create an upgrade task. Set the **CreationOption*	- parameter to **MigrationFromRDS**. For more information, see [Create a PolarDB for MySQL cluster by using the Migration from RDS method](https://help.aliyun.com/document_detail/121582.html).
//
// @param request - CloseDBClusterMigrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseDBClusterMigrationResponse
func (client *Client) CloseDBClusterMigrationWithOptions(request *CloseDBClusterMigrationRequest, runtime *dara.RuntimeOptions) (_result *CloseDBClusterMigrationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContinueEnableBinlog) {
		query["ContinueEnableBinlog"] = request.ContinueEnableBinlog
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseDBClusterMigration"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseDBClusterMigrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels or completes the migration task that upgrades an RDS cluster to a PolarDB cluster.
//
// Description:
//
//	  You can call this operation to cancel the migration task before data migration.
//
//		- You can call this operation to perform the migration task after data migration.
//
// > Before you call this operation, ensure that a one-click upgrade task has been created for the cluster. You can call the [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) operation to create an upgrade task. Set the **CreationOption*	- parameter to **MigrationFromRDS**. For more information, see [Create a PolarDB for MySQL cluster by using the Migration from RDS method](https://help.aliyun.com/document_detail/121582.html).
//
// @param request - CloseDBClusterMigrationRequest
//
// @return CloseDBClusterMigrationResponse
func (client *Client) CloseDBClusterMigration(request *CloseDBClusterMigrationRequest) (_result *CloseDBClusterMigrationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloseDBClusterMigrationResponse{}
	_body, _err := client.CloseDBClusterMigrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ContinueDBClusterMigrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinueDBClusterMigrationResponse
func (client *Client) ContinueDBClusterMigrationWithOptions(request *ContinueDBClusterMigrationRequest, runtime *dara.RuntimeOptions) (_result *ContinueDBClusterMigrationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ForceSwitch) {
		query["ForceSwitch"] = request.ForceSwitch
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContinueDBClusterMigration"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ContinueDBClusterMigrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ContinueDBClusterMigrationRequest
//
// @return ContinueDBClusterMigrationResponse
func (client *Client) ContinueDBClusterMigration(request *ContinueDBClusterMigrationRequest) (_result *ContinueDBClusterMigrationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ContinueDBClusterMigrationResponse{}
	_body, _err := client.ContinueDBClusterMigrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据集
//
// @param request - CreateAIDBClusterDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAIDBClusterDatasetResponse
func (client *Client) CreateAIDBClusterDatasetWithOptions(request *CreateAIDBClusterDatasetRequest, runtime *dara.RuntimeOptions) (_result *CreateAIDBClusterDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.DatasetType) {
		query["DatasetType"] = request.DatasetType
	}

	if !dara.IsNil(request.ImportMode) {
		query["ImportMode"] = request.ImportMode
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

	if !dara.IsNil(request.TrainMode) {
		query["TrainMode"] = request.TrainMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAIDBClusterDataset"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAIDBClusterDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据集
//
// @param request - CreateAIDBClusterDatasetRequest
//
// @return CreateAIDBClusterDatasetResponse
func (client *Client) CreateAIDBClusterDataset(request *CreateAIDBClusterDatasetRequest) (_result *CreateAIDBClusterDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAIDBClusterDatasetResponse{}
	_body, _err := client.CreateAIDBClusterDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建模型评测任务
//
// @param request - CreateAIDBClusterTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAIDBClusterTaskResponse
func (client *Client) CreateAIDBClusterTaskWithOptions(request *CreateAIDBClusterTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateAIDBClusterTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !dara.IsNil(request.DatasetPath) {
		query["DatasetPath"] = request.DatasetPath
	}

	if !dara.IsNil(request.EvalDatasetPath) {
		query["EvalDatasetPath"] = request.EvalDatasetPath
	}

	if !dara.IsNil(request.KubeType) {
		query["KubeType"] = request.KubeType
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.ModelSource) {
		query["ModelSource"] = request.ModelSource
	}

	if !dara.IsNil(request.ModelType) {
		query["ModelType"] = request.ModelType
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

	if !dara.IsNil(request.RunningParameter) {
		query["RunningParameter"] = request.RunningParameter
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.VPCId) {
		query["VPCId"] = request.VPCId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAIDBClusterTask"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAIDBClusterTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建模型评测任务
//
// @param request - CreateAIDBClusterTaskRequest
//
// @return CreateAIDBClusterTaskResponse
func (client *Client) CreateAIDBClusterTask(request *CreateAIDBClusterTaskRequest) (_result *CreateAIDBClusterTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAIDBClusterTaskResponse{}
	_body, _err := client.CreateAIDBClusterTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建AI节点
//
// @param request - CreateAINodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAINodesResponse
func (client *Client) CreateAINodesWithOptions(request *CreateAINodesRequest, runtime *dara.RuntimeOptions) (_result *CreateAINodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNodes) {
		query["DBNodes"] = request.DBNodes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAINodes"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAINodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建AI节点
//
// @param request - CreateAINodesRequest
//
// @return CreateAINodesResponse
func (client *Client) CreateAINodes(request *CreateAINodesRequest) (_result *CreateAINodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAINodesResponse{}
	_body, _err := client.CreateAINodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a database account for a PolarDB cluster.
//
// @param request - CreateAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccountResponse
func (client *Client) CreateAccountWithOptions(request *CreateAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountDescription) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.AccountPrivilege) {
		query["AccountPrivilege"] = request.AccountPrivilege
	}

	if !dara.IsNil(request.AccountType) {
		query["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PrivForAllDB) {
		query["PrivForAllDB"] = request.PrivForAllDB
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
		Action:      dara.String("CreateAccount"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a database account for a PolarDB cluster.
//
// @param request - CreateAccountRequest
//
// @return CreateAccountResponse
func (client *Client) CreateAccount(request *CreateAccountRequest) (_result *CreateAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAccountResponse{}
	_body, _err := client.CreateAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建边缘云账号
//
// @param request - CreateAccountZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccountZonalResponse
func (client *Client) CreateAccountZonalWithOptions(request *CreateAccountZonalRequest, runtime *dara.RuntimeOptions) (_result *CreateAccountZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountDescription) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.AccountPrivilege) {
		query["AccountPrivilege"] = request.AccountPrivilege
	}

	if !dara.IsNil(request.AccountType) {
		query["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PrivForAllDB) {
		query["PrivForAllDB"] = request.PrivForAllDB
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
		Action:      dara.String("CreateAccountZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccountZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建边缘云账号
//
// @param request - CreateAccountZonalRequest
//
// @return CreateAccountZonalResponse
func (client *Client) CreateAccountZonal(request *CreateAccountZonalRequest) (_result *CreateAccountZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAccountZonalResponse{}
	_body, _err := client.CreateAccountZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a lightweight license activation code.
//
// @param request - CreateActivationCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateActivationCodeResponse
func (client *Client) CreateActivationCodeWithOptions(request *CreateActivationCodeRequest, runtime *dara.RuntimeOptions) (_result *CreateActivationCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunOrderId) {
		query["AliyunOrderId"] = request.AliyunOrderId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.MacAddress) {
		query["MacAddress"] = request.MacAddress
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SystemIdentifier) {
		query["SystemIdentifier"] = request.SystemIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateActivationCode"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateActivationCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a lightweight license activation code.
//
// @param request - CreateActivationCodeRequest
//
// @return CreateActivationCodeResponse
func (client *Client) CreateActivationCode(request *CreateActivationCodeRequest) (_result *CreateActivationCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateActivationCodeResponse{}
	_body, _err := client.CreateActivationCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建PolarDB应用
//
// @param tmpReq - CreateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationResponse
func (client *Client) CreateApplicationWithOptions(tmpReq *CreateApplicationRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Components) {
		request.ComponentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Components, dara.String("Components"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Endpoints) {
		request.EndpointsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Endpoints, dara.String("Endpoints"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.KnowledgeApplicationSpec) {
		request.KnowledgeApplicationSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KnowledgeApplicationSpec, dara.String("KnowledgeApplicationSpec"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MemApplicationSpec) {
		request.MemApplicationSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MemApplicationSpec, dara.String("MemApplicationSpec"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AIDBClusterId) {
		query["AIDBClusterId"] = request.AIDBClusterId
	}

	if !dara.IsNil(request.ApplicationType) {
		query["ApplicationType"] = request.ApplicationType
	}

	if !dara.IsNil(request.Architecture) {
		query["Architecture"] = request.Architecture
	}

	if !dara.IsNil(request.AuthProvider) {
		query["AuthProvider"] = request.AuthProvider
	}

	if !dara.IsNil(request.AuthProviderConfig) {
		query["AuthProviderConfig"] = request.AuthProviderConfig
	}

	if !dara.IsNil(request.AutoAllocatePublicEip) {
		query["AutoAllocatePublicEip"] = request.AutoAllocatePublicEip
	}

	if !dara.IsNil(request.AutoCreatePolarFs) {
		query["AutoCreatePolarFs"] = request.AutoCreatePolarFs
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.ComponentsShrink) {
		query["Components"] = request.ComponentsShrink
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EndpointsShrink) {
		query["Endpoints"] = request.EndpointsShrink
	}

	if !dara.IsNil(request.KnowledgeApplicationSpecShrink) {
		query["KnowledgeApplicationSpec"] = request.KnowledgeApplicationSpecShrink
	}

	if !dara.IsNil(request.MemApplicationSpecShrink) {
		query["MemApplicationSpec"] = request.MemApplicationSpecShrink
	}

	if !dara.IsNil(request.ModelApi) {
		query["ModelApi"] = request.ModelApi
	}

	if !dara.IsNil(request.ModelApiKey) {
		query["ModelApiKey"] = request.ModelApiKey
	}

	if !dara.IsNil(request.ModelBaseUrl) {
		query["ModelBaseUrl"] = request.ModelBaseUrl
	}

	if !dara.IsNil(request.ModelFrom) {
		query["ModelFrom"] = request.ModelFrom
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PolarFSInstanceId) {
		query["PolarFSInstanceId"] = request.PolarFSInstanceId
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SecurityIPArrayName) {
		query["SecurityIPArrayName"] = request.SecurityIPArrayName
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !dara.IsNil(request.SecurityIPType) {
		query["SecurityIPType"] = request.SecurityIPType
	}

	if !dara.IsNil(request.SkillTemplateId) {
		query["SkillTemplateId"] = request.SkillTemplateId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TargetVersion) {
		query["TargetVersion"] = request.TargetVersion
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplication"),
		Version:     dara.String("2017-08-01"),
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
// 创建PolarDB应用
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
// 创建PolarDB应用终端节点地址
//
// @param request - CreateApplicationEndpointAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationEndpointAddressResponse
func (client *Client) CreateApplicationEndpointAddressWithOptions(request *CreateApplicationEndpointAddressRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationEndpointAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.NetType) {
		query["NetType"] = request.NetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplicationEndpointAddress"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApplicationEndpointAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建PolarDB应用终端节点地址
//
// @param request - CreateApplicationEndpointAddressRequest
//
// @return CreateApplicationEndpointAddressResponse
func (client *Client) CreateApplicationEndpointAddress(request *CreateApplicationEndpointAddressRequest) (_result *CreateApplicationEndpointAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApplicationEndpointAddressResponse{}
	_body, _err := client.CreateApplicationEndpointAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建应用提示词策略
//
// @param request - CreateApplicationPromptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationPromptResponse
func (client *Client) CreateApplicationPromptWithOptions(request *CreateApplicationPromptRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationPromptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.PromptName) {
		query["PromptName"] = request.PromptName
	}

	if !dara.IsNil(request.PromptType) {
		query["PromptType"] = request.PromptType
	}

	if !dara.IsNil(request.PromptValue) {
		query["PromptValue"] = request.PromptValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplicationPrompt"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApplicationPromptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用提示词策略
//
// @param request - CreateApplicationPromptRequest
//
// @return CreateApplicationPromptResponse
func (client *Client) CreateApplicationPrompt(request *CreateApplicationPromptRequest) (_result *CreateApplicationPromptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApplicationPromptResponse{}
	_body, _err := client.CreateApplicationPromptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a full snapshot backup for a PolarDB cluster.
//
// Description:
//
// >
//
//   - You can manually create up to three backups for each cluster.
//
//   - The `Exceeding the daily backup times of this DB cluster` error message indicates that three manual backups already exist in your cluster. You must delete existing backups before you call this operation to manually create backups. For more information about how to delete backups, see [Delete backups](https://help.aliyun.com/document_detail/98101.html).
//
//   - After you call this operation, a backup task is created in the backend. The task may be time-consuming if you want to back up large amounts of data.
//
// @param request - CreateBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackupResponse
func (client *Client) CreateBackupWithOptions(request *CreateBackupRequest, runtime *dara.RuntimeOptions) (_result *CreateBackupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a full snapshot backup for a PolarDB cluster.
//
// Description:
//
// >
//
//   - You can manually create up to three backups for each cluster.
//
//   - The `Exceeding the daily backup times of this DB cluster` error message indicates that three manual backups already exist in your cluster. You must delete existing backups before you call this operation to manually create backups. For more information about how to delete backups, see [Delete backups](https://help.aliyun.com/document_detail/98101.html).
//
//   - After you call this operation, a backup task is created in the backend. The task may be time-consuming if you want to back up large amounts of data.
//
// @param request - CreateBackupRequest
//
// @return CreateBackupResponse
func (client *Client) CreateBackup(request *CreateBackupRequest) (_result *CreateBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBackupResponse{}
	_body, _err := client.CreateBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建预算策略
//
// @param request - CreateBudgetPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBudgetPolicyResponse
func (client *Client) CreateBudgetPolicyWithOptions(request *CreateBudgetPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateBudgetPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertThresholdPct) {
		query["AlertThresholdPct"] = request.AlertThresholdPct
	}

	if !dara.IsNil(request.BudgetDimensionRefId) {
		query["BudgetDimensionRefId"] = request.BudgetDimensionRefId
	}

	if !dara.IsNil(request.BudgetPoints) {
		query["BudgetPoints"] = request.BudgetPoints
	}

	if !dara.IsNil(request.BudgetType) {
		query["BudgetType"] = request.BudgetType
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResetDayOfMonth) {
		query["ResetDayOfMonth"] = request.ResetDayOfMonth
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBudgetPolicy"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBudgetPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建预算策略
//
// @param request - CreateBudgetPolicyRequest
//
// @return CreateBudgetPolicyResponse
func (client *Client) CreateBudgetPolicy(request *CreateBudgetPolicyRequest) (_result *CreateBudgetPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBudgetPolicyResponse{}
	_body, _err := client.CreateBudgetPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a cluster that is used to store cold data.
//
// @param request - CreateColdStorageInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateColdStorageInstanceResponse
func (client *Client) CreateColdStorageInstanceWithOptions(request *CreateColdStorageInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateColdStorageInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ColdStorageInstanceDescription) {
		query["ColdStorageInstanceDescription"] = request.ColdStorageInstanceDescription
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateColdStorageInstance"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateColdStorageInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cluster that is used to store cold data.
//
// @param request - CreateColdStorageInstanceRequest
//
// @return CreateColdStorageInstanceResponse
func (client *Client) CreateColdStorageInstance(request *CreateColdStorageInstanceRequest) (_result *CreateColdStorageInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateColdStorageInstanceResponse{}
	_body, _err := client.CreateColdStorageInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建消费者
//
// @param request - CreateConsumerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerResponse
func (client *Client) CreateConsumerWithOptions(request *CreateConsumerRequest, runtime *dara.RuntimeOptions) (_result *CreateConsumerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerGroupName) {
		query["ConsumerGroupName"] = request.ConsumerGroupName
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.KeyType) {
		query["KeyType"] = request.KeyType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NickName) {
		query["NickName"] = request.NickName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConsumer"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConsumerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建消费者
//
// @param request - CreateConsumerRequest
//
// @return CreateConsumerResponse
func (client *Client) CreateConsumer(request *CreateConsumerRequest) (_result *CreateConsumerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateConsumerResponse{}
	_body, _err := client.CreateConsumerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建消费者组
//
// @param request - CreateConsumerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerGroupResponse
func (client *Client) CreateConsumerGroupWithOptions(request *CreateConsumerGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateConsumerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerGroupName) {
		query["ConsumerGroupName"] = request.ConsumerGroupName
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.IsDefault) {
		query["IsDefault"] = request.IsDefault
	}

	if !dara.IsNil(request.NickName) {
		query["NickName"] = request.NickName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConsumerGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建消费者组
//
// @param request - CreateConsumerGroupRequest
//
// @return CreateConsumerGroupResponse
func (client *Client) CreateConsumerGroup(request *CreateConsumerGroupRequest) (_result *CreateConsumerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.CreateConsumerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建限流策略
//
// @param request - CreateCostRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCostRuleResponse
func (client *Client) CreateCostRuleWithOptions(request *CreateCostRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateCostRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CacheCostPointsPerMillion) {
		query["CacheCostPointsPerMillion"] = request.CacheCostPointsPerMillion
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.InputCostPointsPerMillion) {
		query["InputCostPointsPerMillion"] = request.InputCostPointsPerMillion
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.ModelServiceId) {
		query["ModelServiceId"] = request.ModelServiceId
	}

	if !dara.IsNil(request.OutputCostPointsPerMillion) {
		query["OutputCostPointsPerMillion"] = request.OutputCostPointsPerMillion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCostRule"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCostRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建限流策略
//
// @param request - CreateCostRuleRequest
//
// @return CreateCostRuleResponse
func (client *Client) CreateCostRule(request *CreateCostRuleRequest) (_result *CreateCostRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCostRuleResponse{}
	_body, _err := client.CreateCostRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// serverless创建周期任务策略
//
// @param request - CreateCronJobPolicyServerlessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCronJobPolicyServerlessResponse
func (client *Client) CreateCronJobPolicyServerlessWithOptions(request *CreateCronJobPolicyServerlessRequest, runtime *dara.RuntimeOptions) (_result *CreateCronJobPolicyServerlessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowShutDown) {
		query["AllowShutDown"] = request.AllowShutDown
	}

	if !dara.IsNil(request.CronExpression) {
		query["CronExpression"] = request.CronExpression
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.ScaleApRoNumMax) {
		query["ScaleApRoNumMax"] = request.ScaleApRoNumMax
	}

	if !dara.IsNil(request.ScaleApRoNumMin) {
		query["ScaleApRoNumMin"] = request.ScaleApRoNumMin
	}

	if !dara.IsNil(request.ScaleMax) {
		query["ScaleMax"] = request.ScaleMax
	}

	if !dara.IsNil(request.ScaleMin) {
		query["ScaleMin"] = request.ScaleMin
	}

	if !dara.IsNil(request.ScaleRoNumMax) {
		query["ScaleRoNumMax"] = request.ScaleRoNumMax
	}

	if !dara.IsNil(request.ScaleRoNumMin) {
		query["ScaleRoNumMin"] = request.ScaleRoNumMin
	}

	if !dara.IsNil(request.SecondsUntilAutoPause) {
		query["SecondsUntilAutoPause"] = request.SecondsUntilAutoPause
	}

	if !dara.IsNil(request.ServerlessRuleCpuEnlargeThreshold) {
		query["ServerlessRuleCpuEnlargeThreshold"] = request.ServerlessRuleCpuEnlargeThreshold
	}

	if !dara.IsNil(request.ServerlessRuleCpuShrinkThreshold) {
		query["ServerlessRuleCpuShrinkThreshold"] = request.ServerlessRuleCpuShrinkThreshold
	}

	if !dara.IsNil(request.ServerlessRuleMode) {
		query["ServerlessRuleMode"] = request.ServerlessRuleMode
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCronJobPolicyServerless"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCronJobPolicyServerlessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// serverless创建周期任务策略
//
// @param request - CreateCronJobPolicyServerlessRequest
//
// @return CreateCronJobPolicyServerlessResponse
func (client *Client) CreateCronJobPolicyServerless(request *CreateCronJobPolicyServerlessRequest) (_result *CreateCronJobPolicyServerlessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCronJobPolicyServerlessResponse{}
	_body, _err := client.CreateCronJobPolicyServerlessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// CreateDBCluster.
//
// @param request - CreateDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBClusterResponse
func (client *Client) CreateDBClusterWithOptions(request *CreateDBClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateDBClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowShutDown) {
		query["AllowShutDown"] = request.AllowShutDown
	}

	if !dara.IsNil(request.Architecture) {
		query["Architecture"] = request.Architecture
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.BackupRetentionPolicyOnClusterDeletion) {
		query["BackupRetentionPolicyOnClusterDeletion"] = request.BackupRetentionPolicyOnClusterDeletion
	}

	if !dara.IsNil(request.BurstingEnabled) {
		query["BurstingEnabled"] = request.BurstingEnabled
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CloneDataPoint) {
		query["CloneDataPoint"] = request.CloneDataPoint
	}

	if !dara.IsNil(request.CloudProvider) {
		query["CloudProvider"] = request.CloudProvider
	}

	if !dara.IsNil(request.ClusterNetworkType) {
		query["ClusterNetworkType"] = request.ClusterNetworkType
	}

	if !dara.IsNil(request.CreationCategory) {
		query["CreationCategory"] = request.CreationCategory
	}

	if !dara.IsNil(request.CreationOption) {
		query["CreationOption"] = request.CreationOption
	}

	if !dara.IsNil(request.DBClusterDescription) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !dara.IsNil(request.DBMinorVersion) {
		query["DBMinorVersion"] = request.DBMinorVersion
	}

	if !dara.IsNil(request.DBNodeClass) {
		query["DBNodeClass"] = request.DBNodeClass
	}

	if !dara.IsNil(request.DBNodeNum) {
		query["DBNodeNum"] = request.DBNodeNum
	}

	if !dara.IsNil(request.DBType) {
		query["DBType"] = request.DBType
	}

	if !dara.IsNil(request.DBVersion) {
		query["DBVersion"] = request.DBVersion
	}

	if !dara.IsNil(request.DefaultTimeZone) {
		query["DefaultTimeZone"] = request.DefaultTimeZone
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.GDNId) {
		query["GDNId"] = request.GDNId
	}

	if !dara.IsNil(request.HotStandbyCluster) {
		query["HotStandbyCluster"] = request.HotStandbyCluster
	}

	if !dara.IsNil(request.LoosePolarLogBin) {
		query["LoosePolarLogBin"] = request.LoosePolarLogBin
	}

	if !dara.IsNil(request.LooseXEngine) {
		query["LooseXEngine"] = request.LooseXEngine
	}

	if !dara.IsNil(request.LooseXEngineUseMemoryPct) {
		query["LooseXEngineUseMemoryPct"] = request.LooseXEngineUseMemoryPct
	}

	if !dara.IsNil(request.LowerCaseTableNames) {
		query["LowerCaseTableNames"] = request.LowerCaseTableNames
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParameterGroupId) {
		query["ParameterGroupId"] = request.ParameterGroupId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.ProvisionedIops) {
		query["ProvisionedIops"] = request.ProvisionedIops
	}

	if !dara.IsNil(request.ProxyClass) {
		query["ProxyClass"] = request.ProxyClass
	}

	if !dara.IsNil(request.ProxyType) {
		query["ProxyType"] = request.ProxyType
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

	if !dara.IsNil(request.ScaleMax) {
		query["ScaleMax"] = request.ScaleMax
	}

	if !dara.IsNil(request.ScaleMin) {
		query["ScaleMin"] = request.ScaleMin
	}

	if !dara.IsNil(request.ScaleRoNumMax) {
		query["ScaleRoNumMax"] = request.ScaleRoNumMax
	}

	if !dara.IsNil(request.ScaleRoNumMin) {
		query["ScaleRoNumMin"] = request.ScaleRoNumMin
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !dara.IsNil(request.ServerlessType) {
		query["ServerlessType"] = request.ServerlessType
	}

	if !dara.IsNil(request.SourceResourceId) {
		query["SourceResourceId"] = request.SourceResourceId
	}

	if !dara.IsNil(request.SourceUid) {
		query["SourceUid"] = request.SourceUid
	}

	if !dara.IsNil(request.StandbyAZ) {
		query["StandbyAZ"] = request.StandbyAZ
	}

	if !dara.IsNil(request.StorageAutoScale) {
		query["StorageAutoScale"] = request.StorageAutoScale
	}

	if !dara.IsNil(request.StorageEncryption) {
		query["StorageEncryption"] = request.StorageEncryption
	}

	if !dara.IsNil(request.StorageEncryptionKey) {
		query["StorageEncryptionKey"] = request.StorageEncryptionKey
	}

	if !dara.IsNil(request.StoragePayType) {
		query["StoragePayType"] = request.StoragePayType
	}

	if !dara.IsNil(request.StorageSpace) {
		query["StorageSpace"] = request.StorageSpace
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.StorageUpperBound) {
		query["StorageUpperBound"] = request.StorageUpperBound
	}

	if !dara.IsNil(request.StrictConsistency) {
		query["StrictConsistency"] = request.StrictConsistency
	}

	if !dara.IsNil(request.TDEStatus) {
		query["TDEStatus"] = request.TDEStatus
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TargetMinorVersion) {
		query["TargetMinorVersion"] = request.TargetMinorVersion
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	if !dara.IsNil(request.VPCId) {
		query["VPCId"] = request.VPCId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBCluster"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// CreateDBCluster.
//
// @param request - CreateDBClusterRequest
//
// @return CreateDBClusterResponse
func (client *Client) CreateDBCluster(request *CreateDBClusterRequest) (_result *CreateDBClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.CreateDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom cluster endpoint for a PolarDB cluster.
//
// @param request - CreateDBClusterEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBClusterEndpointResponse
func (client *Client) CreateDBClusterEndpointWithOptions(request *CreateDBClusterEndpointRequest, runtime *dara.RuntimeOptions) (_result *CreateDBClusterEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoAddNewNodes) {
		query["AutoAddNewNodes"] = request.AutoAddNewNodes
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBEndpointDescription) {
		query["DBEndpointDescription"] = request.DBEndpointDescription
	}

	if !dara.IsNil(request.EndpointConfig) {
		query["EndpointConfig"] = request.EndpointConfig
	}

	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
	}

	if !dara.IsNil(request.Nodes) {
		query["Nodes"] = request.Nodes
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PolarFsInstanceId) {
		query["PolarFsInstanceId"] = request.PolarFsInstanceId
	}

	if !dara.IsNil(request.PolarSccTimeoutAction) {
		query["PolarSccTimeoutAction"] = request.PolarSccTimeoutAction
	}

	if !dara.IsNil(request.PolarSccWaitTimeout) {
		query["PolarSccWaitTimeout"] = request.PolarSccWaitTimeout
	}

	if !dara.IsNil(request.ReadWriteMode) {
		query["ReadWriteMode"] = request.ReadWriteMode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SccMode) {
		query["SccMode"] = request.SccMode
	}

	if !dara.IsNil(request.VPCId) {
		query["VPCId"] = request.VPCId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBClusterEndpoint"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBClusterEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom cluster endpoint for a PolarDB cluster.
//
// @param request - CreateDBClusterEndpointRequest
//
// @return CreateDBClusterEndpointResponse
func (client *Client) CreateDBClusterEndpoint(request *CreateDBClusterEndpointRequest) (_result *CreateDBClusterEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDBClusterEndpointResponse{}
	_body, _err := client.CreateDBClusterEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 边缘云创建链接地址接口
//
// @param request - CreateDBClusterEndpointZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBClusterEndpointZonalResponse
func (client *Client) CreateDBClusterEndpointZonalWithOptions(request *CreateDBClusterEndpointZonalRequest, runtime *dara.RuntimeOptions) (_result *CreateDBClusterEndpointZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoAddNewNodes) {
		query["AutoAddNewNodes"] = request.AutoAddNewNodes
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBEndpointDescription) {
		query["DBEndpointDescription"] = request.DBEndpointDescription
	}

	if !dara.IsNil(request.EndpointConfig) {
		query["EndpointConfig"] = request.EndpointConfig
	}

	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
	}

	if !dara.IsNil(request.Nodes) {
		query["Nodes"] = request.Nodes
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PolarSccTimeoutAction) {
		query["PolarSccTimeoutAction"] = request.PolarSccTimeoutAction
	}

	if !dara.IsNil(request.PolarSccWaitTimeout) {
		query["PolarSccWaitTimeout"] = request.PolarSccWaitTimeout
	}

	if !dara.IsNil(request.ReadWriteMode) {
		query["ReadWriteMode"] = request.ReadWriteMode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SccMode) {
		query["SccMode"] = request.SccMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBClusterEndpointZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBClusterEndpointZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 边缘云创建链接地址接口
//
// @param request - CreateDBClusterEndpointZonalRequest
//
// @return CreateDBClusterEndpointZonalResponse
func (client *Client) CreateDBClusterEndpointZonal(request *CreateDBClusterEndpointZonalRequest) (_result *CreateDBClusterEndpointZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDBClusterEndpointZonalResponse{}
	_body, _err := client.CreateDBClusterEndpointZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a public endpoint for the primary endpoint, the default cluster endpoint, or a custom cluster endpoint.
//
// Description:
//
// > You can create a public endpoint for the primary endpoint, the default cluster endpoint, or a custom cluster endpoint.
//
// @param request - CreateDBEndpointAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBEndpointAddressResponse
func (client *Client) CreateDBEndpointAddressWithOptions(request *CreateDBEndpointAddressRequest, runtime *dara.RuntimeOptions) (_result *CreateDBEndpointAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBEndpointId) {
		query["DBEndpointId"] = request.DBEndpointId
	}

	if !dara.IsNil(request.NetType) {
		query["NetType"] = request.NetType
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

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.VPCId) {
		query["VPCId"] = request.VPCId
	}

	if !dara.IsNil(request.ZoneInfo) {
		query["ZoneInfo"] = request.ZoneInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBEndpointAddress"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBEndpointAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a public endpoint for the primary endpoint, the default cluster endpoint, or a custom cluster endpoint.
//
// Description:
//
// > You can create a public endpoint for the primary endpoint, the default cluster endpoint, or a custom cluster endpoint.
//
// @param request - CreateDBEndpointAddressRequest
//
// @return CreateDBEndpointAddressResponse
func (client *Client) CreateDBEndpointAddress(request *CreateDBEndpointAddressRequest) (_result *CreateDBEndpointAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDBEndpointAddressResponse{}
	_body, _err := client.CreateDBEndpointAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a database link.
//
// Description:
//
// A database link can be used to connect two PolarDB for PostgreSQL(Compatible with Oracle) clusters, or connect a PolarDB for PostgreSQL(Compatible with Oracle) cluster to a user-created PostgreSQL database that is hosted on an Elastic Compute Service (ECS) instance. You can use database links to query data across clusters.
//
// > 	- You can create up to 10 database links for a cluster.
//
// > 	- Each database link connects a source cluster and a destination cluster.
//
// > 	- The source cluster and the destination cluster or the destination ECS instance must be located in the same region.
//
// @param request - CreateDBLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBLinkResponse
func (client *Client) CreateDBLinkWithOptions(request *CreateDBLinkRequest, runtime *dara.RuntimeOptions) (_result *CreateDBLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBLinkName) {
		query["DBLinkName"] = request.DBLinkName
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

	if !dara.IsNil(request.SourceDBName) {
		query["SourceDBName"] = request.SourceDBName
	}

	if !dara.IsNil(request.TargetDBAccount) {
		query["TargetDBAccount"] = request.TargetDBAccount
	}

	if !dara.IsNil(request.TargetDBInstanceName) {
		query["TargetDBInstanceName"] = request.TargetDBInstanceName
	}

	if !dara.IsNil(request.TargetDBName) {
		query["TargetDBName"] = request.TargetDBName
	}

	if !dara.IsNil(request.TargetDBPasswd) {
		query["TargetDBPasswd"] = request.TargetDBPasswd
	}

	if !dara.IsNil(request.TargetIp) {
		query["TargetIp"] = request.TargetIp
	}

	if !dara.IsNil(request.TargetPort) {
		query["TargetPort"] = request.TargetPort
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBLink"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a database link.
//
// Description:
//
// A database link can be used to connect two PolarDB for PostgreSQL(Compatible with Oracle) clusters, or connect a PolarDB for PostgreSQL(Compatible with Oracle) cluster to a user-created PostgreSQL database that is hosted on an Elastic Compute Service (ECS) instance. You can use database links to query data across clusters.
//
// > 	- You can create up to 10 database links for a cluster.
//
// > 	- Each database link connects a source cluster and a destination cluster.
//
// > 	- The source cluster and the destination cluster or the destination ECS instance must be located in the same region.
//
// @param request - CreateDBLinkRequest
//
// @return CreateDBLinkResponse
func (client *Client) CreateDBLink(request *CreateDBLinkRequest) (_result *CreateDBLinkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDBLinkResponse{}
	_body, _err := client.CreateDBLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a read-only node to a PolarDB cluster.
//
// @param request - CreateDBNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBNodesResponse
func (client *Client) CreateDBNodesWithOptions(request *CreateDBNodesRequest, runtime *dara.RuntimeOptions) (_result *CreateDBNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CloudProvider) {
		query["CloudProvider"] = request.CloudProvider
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNode) {
		query["DBNode"] = request.DBNode
	}

	if !dara.IsNil(request.DBNodeType) {
		query["DBNodeType"] = request.DBNodeType
	}

	if !dara.IsNil(request.EndpointBindList) {
		query["EndpointBindList"] = request.EndpointBindList
	}

	if !dara.IsNil(request.ImciSwitch) {
		query["ImciSwitch"] = request.ImciSwitch
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlannedEndTime) {
		query["PlannedEndTime"] = request.PlannedEndTime
	}

	if !dara.IsNil(request.PlannedStartTime) {
		query["PlannedStartTime"] = request.PlannedStartTime
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBNodes"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a read-only node to a PolarDB cluster.
//
// @param request - CreateDBNodesRequest
//
// @return CreateDBNodesResponse
func (client *Client) CreateDBNodes(request *CreateDBNodesRequest) (_result *CreateDBNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDBNodesResponse{}
	_body, _err := client.CreateDBNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a database for a PolarDB cluster.
//
// Description:
//
// Before you call this operation, make sure that the following requirements are met:
//
//   - The cluster is in the Running state.
//
//   - The cluster is unlocked.
//
// @param request - CreateDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatabaseResponse
func (client *Client) CreateDatabaseWithOptions(request *CreateDatabaseRequest, runtime *dara.RuntimeOptions) (_result *CreateDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPrivilege) {
		query["AccountPrivilege"] = request.AccountPrivilege
	}

	if !dara.IsNil(request.CharacterSetName) {
		query["CharacterSetName"] = request.CharacterSetName
	}

	if !dara.IsNil(request.Collate) {
		query["Collate"] = request.Collate
	}

	if !dara.IsNil(request.Ctype) {
		query["Ctype"] = request.Ctype
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBDescription) {
		query["DBDescription"] = request.DBDescription
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatabase"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a database for a PolarDB cluster.
//
// Description:
//
// Before you call this operation, make sure that the following requirements are met:
//
//   - The cluster is in the Running state.
//
//   - The cluster is unlocked.
//
// @param request - CreateDatabaseRequest
//
// @return CreateDatabaseResponse
func (client *Client) CreateDatabase(request *CreateDatabaseRequest) (_result *CreateDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CreateDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建PolarDB边缘云数据库
//
// @param request - CreateDatabaseZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatabaseZonalResponse
func (client *Client) CreateDatabaseZonalWithOptions(request *CreateDatabaseZonalRequest, runtime *dara.RuntimeOptions) (_result *CreateDatabaseZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPrivilege) {
		query["AccountPrivilege"] = request.AccountPrivilege
	}

	if !dara.IsNil(request.CharacterSetName) {
		query["CharacterSetName"] = request.CharacterSetName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Collate) {
		query["Collate"] = request.Collate
	}

	if !dara.IsNil(request.Ctype) {
		query["Ctype"] = request.Ctype
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBDescription) {
		query["DBDescription"] = request.DBDescription
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatabaseZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatabaseZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建PolarDB边缘云数据库
//
// @param request - CreateDatabaseZonalRequest
//
// @return CreateDatabaseZonalResponse
func (client *Client) CreateDatabaseZonal(request *CreateDatabaseZonalRequest) (_result *CreateDatabaseZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDatabaseZonalResponse{}
	_body, _err := client.CreateDatabaseZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建插件
//
// @param request - CreateExtensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExtensionsResponse
func (client *Client) CreateExtensionsWithOptions(request *CreateExtensionsRequest, runtime *dara.RuntimeOptions) (_result *CreateExtensionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNames) {
		query["DBNames"] = request.DBNames
	}

	if !dara.IsNil(request.Extensions) {
		query["Extensions"] = request.Extensions
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

	if !dara.IsNil(request.SourceDBName) {
		query["SourceDBName"] = request.SourceDBName
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExtensions"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExtensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建插件
//
// @param request - CreateExtensionsRequest
//
// @return CreateExtensionsResponse
func (client *Client) CreateExtensions(request *CreateExtensionsRequest) (_result *CreateExtensionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateExtensionsResponse{}
	_body, _err := client.CreateExtensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建网关地址
//
// @param request - CreateGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewayResponse
func (client *Client) CreateGatewayWithOptions(request *CreateGatewayRequest, runtime *dara.RuntimeOptions) (_result *CreateGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.DBClusterClass) {
		query["DBClusterClass"] = request.DBClusterClass
	}

	if !dara.IsNil(request.DBType) {
		query["DBType"] = request.DBType
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	if !dara.IsNil(request.VPCId) {
		query["VPCId"] = request.VPCId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGateway"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建网关地址
//
// @param request - CreateGatewayRequest
//
// @return CreateGatewayResponse
func (client *Client) CreateGateway(request *CreateGatewayRequest) (_result *CreateGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGatewayResponse{}
	_body, _err := client.CreateGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a global data network (GDN).
//
// @param request - CreateGlobalDataNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGlobalDataNetworkResponse
func (client *Client) CreateGlobalDataNetworkWithOptions(request *CreateGlobalDataNetworkRequest, runtime *dara.RuntimeOptions) (_result *CreateGlobalDataNetworkResponse, _err error) {
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

	if !dara.IsNil(request.DestinationFileSystemPath) {
		query["DestinationFileSystemPath"] = request.DestinationFileSystemPath
	}

	if !dara.IsNil(request.DestinationId) {
		query["DestinationId"] = request.DestinationId
	}

	if !dara.IsNil(request.DestinationRegion) {
		query["DestinationRegion"] = request.DestinationRegion
	}

	if !dara.IsNil(request.DestinationType) {
		query["DestinationType"] = request.DestinationType
	}

	if !dara.IsNil(request.FreezeSourceDuringSync) {
		query["FreezeSourceDuringSync"] = request.FreezeSourceDuringSync
	}

	if !dara.IsNil(request.SourceFileSystemPath) {
		query["SourceFileSystemPath"] = request.SourceFileSystemPath
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceRegion) {
		query["SourceRegion"] = request.SourceRegion
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGlobalDataNetwork"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGlobalDataNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a global data network (GDN).
//
// @param request - CreateGlobalDataNetworkRequest
//
// @return CreateGlobalDataNetworkResponse
func (client *Client) CreateGlobalDataNetwork(request *CreateGlobalDataNetworkRequest) (_result *CreateGlobalDataNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGlobalDataNetworkResponse{}
	_body, _err := client.CreateGlobalDataNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a global database network (GDN).
//
// Description:
//
// >  A cluster belongs to only one GDN.
//
// @param request - CreateGlobalDatabaseNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGlobalDatabaseNetworkResponse
func (client *Client) CreateGlobalDatabaseNetworkWithOptions(request *CreateGlobalDatabaseNetworkRequest, runtime *dara.RuntimeOptions) (_result *CreateGlobalDatabaseNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EnableGlobalDomainName) {
		query["EnableGlobalDomainName"] = request.EnableGlobalDomainName
	}

	if !dara.IsNil(request.GDNDescription) {
		query["GDNDescription"] = request.GDNDescription
	}

	if !dara.IsNil(request.GDNVersion) {
		query["GDNVersion"] = request.GDNVersion
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGlobalDatabaseNetwork"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGlobalDatabaseNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a global database network (GDN).
//
// Description:
//
// >  A cluster belongs to only one GDN.
//
// @param request - CreateGlobalDatabaseNetworkRequest
//
// @return CreateGlobalDatabaseNetworkResponse
func (client *Client) CreateGlobalDatabaseNetwork(request *CreateGlobalDatabaseNetworkRequest) (_result *CreateGlobalDatabaseNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGlobalDatabaseNetworkResponse{}
	_body, _err := client.CreateGlobalDatabaseNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a global IP whitelist template.
//
// @param request - CreateGlobalSecurityIPGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGlobalSecurityIPGroupResponse
func (client *Client) CreateGlobalSecurityIPGroupWithOptions(request *CreateGlobalSecurityIPGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateGlobalSecurityIPGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GIpList) {
		query["GIpList"] = request.GIpList
	}

	if !dara.IsNil(request.GlobalIgName) {
		query["GlobalIgName"] = request.GlobalIgName
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGlobalSecurityIPGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGlobalSecurityIPGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a global IP whitelist template.
//
// @param request - CreateGlobalSecurityIPGroupRequest
//
// @return CreateGlobalSecurityIPGroupResponse
func (client *Client) CreateGlobalSecurityIPGroup(request *CreateGlobalSecurityIPGroupRequest) (_result *CreateGlobalSecurityIPGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGlobalSecurityIPGroupResponse{}
	_body, _err := client.CreateGlobalSecurityIPGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建路由规则
//
// @param request - CreateModelApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelApiResponse
func (client *Client) CreateModelApiWithOptions(request *CreateModelApiRequest, runtime *dara.RuntimeOptions) (_result *CreateModelApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForceModel) {
		query["ForceModel"] = request.ForceModel
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.ModelCategory) {
		query["ModelCategory"] = request.ModelCategory
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PathPrefix) {
		query["PathPrefix"] = request.PathPrefix
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RecordInput) {
		query["RecordInput"] = request.RecordInput
	}

	if !dara.IsNil(request.RecordOutput) {
		query["RecordOutput"] = request.RecordOutput
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RouteRules) {
		query["RouteRules"] = request.RouteRules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModelApi"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建路由规则
//
// @param request - CreateModelApiRequest
//
// @return CreateModelApiResponse
func (client *Client) CreateModelApi(request *CreateModelApiRequest) (_result *CreateModelApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateModelApiResponse{}
	_body, _err := client.CreateModelApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建模型服务
//
// @param request - CreateModelServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelServiceResponse
func (client *Client) CreateModelServiceWithOptions(request *CreateModelServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateModelServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.BaseUrl) {
		query["BaseUrl"] = request.BaseUrl
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.InputCostPointsPerMillion) {
		query["InputCostPointsPerMillion"] = request.InputCostPointsPerMillion
	}

	if !dara.IsNil(request.ModelCategory) {
		query["ModelCategory"] = request.ModelCategory
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OutputCostPointsPerMillion) {
		query["OutputCostPointsPerMillion"] = request.OutputCostPointsPerMillion
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestCostPoints) {
		query["RequestCostPoints"] = request.RequestCostPoints
	}

	if !dara.IsNil(request.Vendor) {
		query["Vendor"] = request.Vendor
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModelService"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建模型服务
//
// @param request - CreateModelServiceRequest
//
// @return CreateModelServiceResponse
func (client *Client) CreateModelService(request *CreateModelServiceRequest) (_result *CreateModelServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateModelServiceResponse{}
	_body, _err := client.CreateModelServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建网络通道
//
// @param request - CreateNetworkChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkChannelResponse
func (client *Client) CreateNetworkChannelWithOptions(request *CreateNetworkChannelRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Notes) {
		query["Notes"] = request.Notes
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

	if !dara.IsNil(request.TargetDBClusterId) {
		query["TargetDBClusterId"] = request.TargetDBClusterId
	}

	if !dara.IsNil(request.TargetIp) {
		query["TargetIp"] = request.TargetIp
	}

	if !dara.IsNil(request.TargetPort) {
		query["TargetPort"] = request.TargetPort
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetworkChannel"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetworkChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建网络通道
//
// @param request - CreateNetworkChannelRequest
//
// @return CreateNetworkChannelResponse
func (client *Client) CreateNetworkChannel(request *CreateNetworkChannelRequest) (_result *CreateNetworkChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNetworkChannelResponse{}
	_body, _err := client.CreateNetworkChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or obtains a virtual license order.
//
// @param request - CreateOrGetVirtualLicenseOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrGetVirtualLicenseOrderResponse
func (client *Client) CreateOrGetVirtualLicenseOrderWithOptions(request *CreateOrGetVirtualLicenseOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateOrGetVirtualLicenseOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrGetVirtualLicenseOrder"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrGetVirtualLicenseOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or obtains a virtual license order.
//
// @param request - CreateOrGetVirtualLicenseOrderRequest
//
// @return CreateOrGetVirtualLicenseOrderResponse
func (client *Client) CreateOrGetVirtualLicenseOrder(request *CreateOrGetVirtualLicenseOrderRequest) (_result *CreateOrGetVirtualLicenseOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOrGetVirtualLicenseOrderResponse{}
	_body, _err := client.CreateOrGetVirtualLicenseOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a parameter template.
//
// Description:
//
// You can use parameter templates to manage multiple parameters at a time and apply existing parameters to a PolarDB cluster. For more information, see [Use a parameter template](https://help.aliyun.com/document_detail/207009.html).
//
// > You can call this operation only on a PolarDB for MySQL cluster.
//
// @param request - CreateParameterGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateParameterGroupResponse
func (client *Client) CreateParameterGroupWithOptions(request *CreateParameterGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateParameterGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBType) {
		query["DBType"] = request.DBType
	}

	if !dara.IsNil(request.DBVersion) {
		query["DBVersion"] = request.DBVersion
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParameterGroupDesc) {
		query["ParameterGroupDesc"] = request.ParameterGroupDesc
	}

	if !dara.IsNil(request.ParameterGroupName) {
		query["ParameterGroupName"] = request.ParameterGroupName
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateParameterGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateParameterGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a parameter template.
//
// Description:
//
// You can use parameter templates to manage multiple parameters at a time and apply existing parameters to a PolarDB cluster. For more information, see [Use a parameter template](https://help.aliyun.com/document_detail/207009.html).
//
// > You can call this operation only on a PolarDB for MySQL cluster.
//
// @param request - CreateParameterGroupRequest
//
// @return CreateParameterGroupResponse
func (client *Client) CreateParameterGroup(request *CreateParameterGroupRequest) (_result *CreateParameterGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateParameterGroupResponse{}
	_body, _err := client.CreateParameterGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建PolarClaw Agent
//
// @param request - CreatePolarClawAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolarClawAgentResponse
func (client *Client) CreatePolarClawAgentWithOptions(request *CreatePolarClawAgentRequest, runtime *dara.RuntimeOptions) (_result *CreatePolarClawAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.Avatar) {
		query["Avatar"] = request.Avatar
	}

	if !dara.IsNil(request.Emoji) {
		query["Emoji"] = request.Emoji
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	if !dara.IsNil(request.Workspace) {
		query["Workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolarClawAgent"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolarClawAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建PolarClaw Agent
//
// @param request - CreatePolarClawAgentRequest
//
// @return CreatePolarClawAgentResponse
func (client *Client) CreatePolarClawAgent(request *CreatePolarClawAgentRequest) (_result *CreatePolarClawAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePolarClawAgentResponse{}
	_body, _err := client.CreatePolarClawAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建PolarClaw Channel
//
// @param tmpReq - CreatePolarClawChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolarClawChannelResponse
func (client *Client) CreatePolarClawChannelWithOptions(tmpReq *CreatePolarClawChannelRequest, runtime *dara.RuntimeOptions) (_result *CreatePolarClawChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePolarClawChannelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChannelConfig) {
		request.ChannelConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelConfig, dara.String("ChannelConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ChannelConfigShrink) {
		query["ChannelConfig"] = request.ChannelConfigShrink
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.NpmPackage) {
		query["NpmPackage"] = request.NpmPackage
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolarClawChannel"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolarClawChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建PolarClaw Channel
//
// @param request - CreatePolarClawChannelRequest
//
// @return CreatePolarClawChannelResponse
func (client *Client) CreatePolarClawChannel(request *CreatePolarClawChannelRequest) (_result *CreatePolarClawChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePolarClawChannelResponse{}
	_body, _err := client.CreatePolarClawChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建PolarClaw定时任务
//
// @param tmpReq - CreatePolarClawCronJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolarClawCronJobResponse
func (client *Client) CreatePolarClawCronJobWithOptions(tmpReq *CreatePolarClawCronJobRequest, runtime *dara.RuntimeOptions) (_result *CreatePolarClawCronJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePolarClawCronJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Delivery) {
		request.DeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Delivery, dara.String("Delivery"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FailureAlert) {
		request.FailureAlertShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FailureAlert, dara.String("FailureAlert"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Schedule) {
		request.ScheduleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Schedule, dara.String("Schedule"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.DeleteAfterRun) {
		query["DeleteAfterRun"] = request.DeleteAfterRun
	}

	if !dara.IsNil(request.DeliveryShrink) {
		query["Delivery"] = request.DeliveryShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.FailureAlertShrink) {
		query["FailureAlert"] = request.FailureAlertShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PayloadShrink) {
		query["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	if !dara.IsNil(request.RunImmediately) {
		query["RunImmediately"] = request.RunImmediately
	}

	if !dara.IsNil(request.ScheduleShrink) {
		query["Schedule"] = request.ScheduleShrink
	}

	if !dara.IsNil(request.SessionKey) {
		query["SessionKey"] = request.SessionKey
	}

	if !dara.IsNil(request.SessionTarget) {
		query["SessionTarget"] = request.SessionTarget
	}

	if !dara.IsNil(request.WakeMode) {
		query["WakeMode"] = request.WakeMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolarClawCronJob"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolarClawCronJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建PolarClaw定时任务
//
// @param request - CreatePolarClawCronJobRequest
//
// @return CreatePolarClawCronJobResponse
func (client *Client) CreatePolarClawCronJob(request *CreatePolarClawCronJobRequest) (_result *CreatePolarClawCronJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePolarClawCronJobResponse{}
	_body, _err := client.CreatePolarClawCronJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用于在指定PolarFS实例中创建新的目录。
//
// Description:
//
// ## 请求说明
//
// - **Path**：需要创建的目录绝对路径。
//
// - **Recursive**：是否递归创建父目录，默认为 `false`。
//
// - 该接口支持在指定的PolarFS实例中创建单个或多个层级的目录结构。
//
// - 如果设置 `Recursive` 为 `true`，则会自动创建所有不存在的父目录。
//
// - 创建目录时，请确保具有足够的权限。
//
// @param request - CreatePolarFsObjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolarFsObjectResponse
func (client *Client) CreatePolarFsObjectWithOptions(request *CreatePolarFsObjectRequest, runtime *dara.RuntimeOptions) (_result *CreatePolarFsObjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.PolarFsInstanceId) {
		query["PolarFsInstanceId"] = request.PolarFsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolarFsObject"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolarFsObjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于在指定PolarFS实例中创建新的目录。
//
// Description:
//
// ## 请求说明
//
// - **Path**：需要创建的目录绝对路径。
//
// - **Recursive**：是否递归创建父目录，默认为 `false`。
//
// - 该接口支持在指定的PolarFS实例中创建单个或多个层级的目录结构。
//
// - 如果设置 `Recursive` 为 `true`，则会自动创建所有不存在的父目录。
//
// - 创建目录时，请确保具有足够的权限。
//
// @param request - CreatePolarFsObjectRequest
//
// @return CreatePolarFsObjectResponse
func (client *Client) CreatePolarFsObject(request *CreatePolarFsObjectRequest) (_result *CreatePolarFsObjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePolarFsObjectResponse{}
	_body, _err := client.CreatePolarFsObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建限流策略
//
// @param request - CreateRateLimitPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRateLimitPolicyResponse
func (client *Client) CreateRateLimitPolicyWithOptions(request *CreateRateLimitPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateRateLimitPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.RateLimitRpm) {
		query["RateLimitRpm"] = request.RateLimitRpm
	}

	if !dara.IsNil(request.RateLimitTpm) {
		query["RateLimitTpm"] = request.RateLimitTpm
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScopeRefId) {
		query["ScopeRefId"] = request.ScopeRefId
	}

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRateLimitPolicy"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRateLimitPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建限流策略
//
// @param request - CreateRateLimitPolicyRequest
//
// @return CreateRateLimitPolicyResponse
func (client *Client) CreateRateLimitPolicy(request *CreateRateLimitPolicyRequest) (_result *CreateRateLimitPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRateLimitPolicyResponse{}
	_body, _err := client.CreateRateLimitPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a service-linked role (SLR).
//
// @param request - CreateServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceLinkedRoleResponse
func (client *Client) CreateServiceLinkedRoleWithOptions(request *CreateServiceLinkedRoleRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceLinkedRole"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service-linked role (SLR).
//
// @param request - CreateServiceLinkedRoleRequest
//
// @return CreateServiceLinkedRoleResponse
func (client *Client) CreateServiceLinkedRole(request *CreateServiceLinkedRoleRequest) (_result *CreateServiceLinkedRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CreateServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Purchases a storage plan.
//
// @param request - CreateStoragePlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStoragePlanResponse
func (client *Client) CreateStoragePlanWithOptions(request *CreateStoragePlanRequest, runtime *dara.RuntimeOptions) (_result *CreateStoragePlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StorageClass) {
		query["StorageClass"] = request.StorageClass
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStoragePlan"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStoragePlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Purchases a storage plan.
//
// @param request - CreateStoragePlanRequest
//
// @return CreateStoragePlanResponse
func (client *Client) CreateStoragePlan(request *CreateStoragePlanRequest) (_result *CreateStoragePlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStoragePlanResponse{}
	_body, _err := client.CreateStoragePlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除AI集群实例
//
// @param request - DeleteAIDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAIDBClusterResponse
func (client *Client) DeleteAIDBClusterWithOptions(request *DeleteAIDBClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteAIDBClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAIDBCluster"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAIDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除AI集群实例
//
// @param request - DeleteAIDBClusterRequest
//
// @return DeleteAIDBClusterResponse
func (client *Client) DeleteAIDBCluster(request *DeleteAIDBClusterRequest) (_result *DeleteAIDBClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAIDBClusterResponse{}
	_body, _err := client.DeleteAIDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据集
//
// @param request - DeleteAIDBClusterDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAIDBClusterDatasetResponse
func (client *Client) DeleteAIDBClusterDatasetWithOptions(request *DeleteAIDBClusterDatasetRequest, runtime *dara.RuntimeOptions) (_result *DeleteAIDBClusterDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
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
		Action:      dara.String("DeleteAIDBClusterDataset"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAIDBClusterDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据集
//
// @param request - DeleteAIDBClusterDatasetRequest
//
// @return DeleteAIDBClusterDatasetResponse
func (client *Client) DeleteAIDBClusterDataset(request *DeleteAIDBClusterDatasetRequest) (_result *DeleteAIDBClusterDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAIDBClusterDatasetResponse{}
	_body, _err := client.DeleteAIDBClusterDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除任务实例
//
// @param request - DeleteAIDBClusterTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAIDBClusterTaskResponse
func (client *Client) DeleteAIDBClusterTaskWithOptions(request *DeleteAIDBClusterTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteAIDBClusterTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.RelativeDBClusterId) {
		query["RelativeDBClusterId"] = request.RelativeDBClusterId
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
		Action:      dara.String("DeleteAIDBClusterTask"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAIDBClusterTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除任务实例
//
// @param request - DeleteAIDBClusterTaskRequest
//
// @return DeleteAIDBClusterTaskResponse
func (client *Client) DeleteAIDBClusterTask(request *DeleteAIDBClusterTaskRequest) (_result *DeleteAIDBClusterTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAIDBClusterTaskResponse{}
	_body, _err := client.DeleteAIDBClusterTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除ai实例子节点
//
// @param request - DeleteAINodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAINodesResponse
func (client *Client) DeleteAINodesWithOptions(request *DeleteAINodesRequest, runtime *dara.RuntimeOptions) (_result *DeleteAINodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNodeId) {
		query["DBNodeId"] = request.DBNodeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAINodes"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAINodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除ai实例子节点
//
// @param request - DeleteAINodesRequest
//
// @return DeleteAINodesResponse
func (client *Client) DeleteAINodes(request *DeleteAINodesRequest) (_result *DeleteAINodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAINodesResponse{}
	_body, _err := client.DeleteAINodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a database account for a PolarDB cluster.
//
// Description:
//
// > Before you call this operation, make sure that the cluster is in the Running state. Otherwise, the operation fails.
//
// @param request - DeleteAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccountWithOptions(request *DeleteAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccount"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a database account for a PolarDB cluster.
//
// Description:
//
// > Before you call this operation, make sure that the cluster is in the Running state. Otherwise, the operation fails.
//
// @param request - DeleteAccountRequest
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccount(request *DeleteAccountRequest) (_result *DeleteAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DeleteAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除PolarDB边缘云集群账号
//
// @param request - DeleteAccountZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccountZonalResponse
func (client *Client) DeleteAccountZonalWithOptions(request *DeleteAccountZonalRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccountZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccountZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccountZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除PolarDB边缘云集群账号
//
// @param request - DeleteAccountZonalRequest
//
// @return DeleteAccountZonalResponse
func (client *Client) DeleteAccountZonal(request *DeleteAccountZonalRequest) (_result *DeleteAccountZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAccountZonalResponse{}
	_body, _err := client.DeleteAccountZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除PolarDB应用
//
// @param request - DeleteApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplicationWithOptions(request *DeleteApplicationRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplication"),
		Version:     dara.String("2017-08-01"),
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
// 删除PolarDB应用
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
// 删除PolarDB应用终端地址
//
// @param request - DeleteApplicationEndpointAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationEndpointAddressResponse
func (client *Client) DeleteApplicationEndpointAddressWithOptions(request *DeleteApplicationEndpointAddressRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationEndpointAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.NetType) {
		query["NetType"] = request.NetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplicationEndpointAddress"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApplicationEndpointAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除PolarDB应用终端地址
//
// @param request - DeleteApplicationEndpointAddressRequest
//
// @return DeleteApplicationEndpointAddressResponse
func (client *Client) DeleteApplicationEndpointAddress(request *DeleteApplicationEndpointAddressRequest) (_result *DeleteApplicationEndpointAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApplicationEndpointAddressResponse{}
	_body, _err := client.DeleteApplicationEndpointAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除应用提示词策略
//
// @param request - DeleteApplicationPromptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationPromptResponse
func (client *Client) DeleteApplicationPromptWithOptions(request *DeleteApplicationPromptRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationPromptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.PromptId) {
		query["PromptId"] = request.PromptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplicationPrompt"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApplicationPromptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除应用提示词策略
//
// @param request - DeleteApplicationPromptRequest
//
// @return DeleteApplicationPromptResponse
func (client *Client) DeleteApplicationPrompt(request *DeleteApplicationPromptRequest) (_result *DeleteApplicationPromptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApplicationPromptResponse{}
	_body, _err := client.DeleteApplicationPromptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the backup sets of a PolarDB cluster.
//
// Description:
//
// Before you call this operation, make sure that the cluster meets the following requirements:
//
//   - The cluster is in the Running state.
//
//   - The backup sets are in the Success state.
//
// > 	- You can call the [DescribeBackups](https://help.aliyun.com/document_detail/98102.html) operation to query the status of backup sets.
//
// >	- After you delete the backup set file, the storage space that is occupied by the file is released. The released storage space is smaller than the size of the file because your snapshots share some data blocks
//
// @param request - DeleteBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackupResponse
func (client *Client) DeleteBackupWithOptions(request *DeleteBackupRequest, runtime *dara.RuntimeOptions) (_result *DeleteBackupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBackup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the backup sets of a PolarDB cluster.
//
// Description:
//
// Before you call this operation, make sure that the cluster meets the following requirements:
//
//   - The cluster is in the Running state.
//
//   - The backup sets are in the Success state.
//
// > 	- You can call the [DescribeBackups](https://help.aliyun.com/document_detail/98102.html) operation to query the status of backup sets.
//
// >	- After you delete the backup set file, the storage space that is occupied by the file is released. The released storage space is smaller than the size of the file because your snapshots share some data blocks
//
// @param request - DeleteBackupRequest
//
// @return DeleteBackupResponse
func (client *Client) DeleteBackup(request *DeleteBackupRequest) (_result *DeleteBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBackupResponse{}
	_body, _err := client.DeleteBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除预算策略
//
// @param request - DeleteBudgetPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBudgetPolicyResponse
func (client *Client) DeleteBudgetPolicyWithOptions(request *DeleteBudgetPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteBudgetPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BudgetPolicyId) {
		query["BudgetPolicyId"] = request.BudgetPolicyId
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBudgetPolicy"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBudgetPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除预算策略
//
// @param request - DeleteBudgetPolicyRequest
//
// @return DeleteBudgetPolicyResponse
func (client *Client) DeleteBudgetPolicy(request *DeleteBudgetPolicyRequest) (_result *DeleteBudgetPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBudgetPolicyResponse{}
	_body, _err := client.DeleteBudgetPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除消费者
//
// @param request - DeleteConsumerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerResponse
func (client *Client) DeleteConsumerWithOptions(request *DeleteConsumerRequest, runtime *dara.RuntimeOptions) (_result *DeleteConsumerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerId) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConsumer"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConsumerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除消费者
//
// @param request - DeleteConsumerRequest
//
// @return DeleteConsumerResponse
func (client *Client) DeleteConsumer(request *DeleteConsumerRequest) (_result *DeleteConsumerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteConsumerResponse{}
	_body, _err := client.DeleteConsumerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除消费者组
//
// @param request - DeleteConsumerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerGroupResponse
func (client *Client) DeleteConsumerGroupWithOptions(request *DeleteConsumerGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteConsumerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerGroupName) {
		query["ConsumerGroupName"] = request.ConsumerGroupName
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConsumerGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除消费者组
//
// @param request - DeleteConsumerGroupRequest
//
// @return DeleteConsumerGroupResponse
func (client *Client) DeleteConsumerGroup(request *DeleteConsumerGroupRequest) (_result *DeleteConsumerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.DeleteConsumerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除限流策略
//
// @param request - DeleteCostRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCostRuleResponse
func (client *Client) DeleteCostRuleWithOptions(request *DeleteCostRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteCostRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CostRuleId) {
		query["CostRuleId"] = request.CostRuleId
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCostRule"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCostRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除限流策略
//
// @param request - DeleteCostRuleRequest
//
// @return DeleteCostRuleResponse
func (client *Client) DeleteCostRule(request *DeleteCostRuleRequest) (_result *DeleteCostRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCostRuleResponse{}
	_body, _err := client.DeleteCostRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a pay-as-you-go PolarDB cluster.
//
// @param request - DeleteDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBClusterResponse
func (client *Client) DeleteDBClusterWithOptions(request *DeleteDBClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupRetentionPolicyOnClusterDeletion) {
		query["BackupRetentionPolicyOnClusterDeletion"] = request.BackupRetentionPolicyOnClusterDeletion
	}

	if !dara.IsNil(request.CloudProvider) {
		query["CloudProvider"] = request.CloudProvider
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBCluster"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a pay-as-you-go PolarDB cluster.
//
// @param request - DeleteDBClusterRequest
//
// @return DeleteDBClusterResponse
func (client *Client) DeleteDBCluster(request *DeleteDBClusterRequest) (_result *DeleteDBClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.DeleteDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a custom cluster endpoint of a PolarDB cluster.
//
// @param request - DeleteDBClusterEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBClusterEndpointResponse
func (client *Client) DeleteDBClusterEndpointWithOptions(request *DeleteDBClusterEndpointRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBClusterEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBEndpointId) {
		query["DBEndpointId"] = request.DBEndpointId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PolarFsInstanceId) {
		query["PolarFsInstanceId"] = request.PolarFsInstanceId
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
		Action:      dara.String("DeleteDBClusterEndpoint"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBClusterEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a custom cluster endpoint of a PolarDB cluster.
//
// @param request - DeleteDBClusterEndpointRequest
//
// @return DeleteDBClusterEndpointResponse
func (client *Client) DeleteDBClusterEndpoint(request *DeleteDBClusterEndpointRequest) (_result *DeleteDBClusterEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDBClusterEndpointResponse{}
	_body, _err := client.DeleteDBClusterEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除PolarDB 边缘云集群的链接地址
//
// @param request - DeleteDBClusterEndpointZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBClusterEndpointZonalResponse
func (client *Client) DeleteDBClusterEndpointZonalWithOptions(request *DeleteDBClusterEndpointZonalRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBClusterEndpointZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBEndpointId) {
		query["DBEndpointId"] = request.DBEndpointId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBClusterEndpointZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBClusterEndpointZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除PolarDB 边缘云集群的链接地址
//
// @param request - DeleteDBClusterEndpointZonalRequest
//
// @return DeleteDBClusterEndpointZonalResponse
func (client *Client) DeleteDBClusterEndpointZonal(request *DeleteDBClusterEndpointZonalRequest) (_result *DeleteDBClusterEndpointZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDBClusterEndpointZonalResponse{}
	_body, _err := client.DeleteDBClusterEndpointZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases the public endpoints of a PolarDB cluster, including the primary endpoint, default cluster endpoint, and custom cluster endpoint.
//
// Description:
//
// > 	- You can delete a public-facing or classic network endpoint of the primary endpoint, the default cluster endpoint, or a custom cluster endpoint.
//
// > 	- Classic network endpoints are supported only on the China site (aliyun.com). Therefore, you do not need to delete classic network endpoints on the International site (alibabacloud.com).
//
// @param request - DeleteDBEndpointAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBEndpointAddressResponse
func (client *Client) DeleteDBEndpointAddressWithOptions(request *DeleteDBEndpointAddressRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBEndpointAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBEndpointId) {
		query["DBEndpointId"] = request.DBEndpointId
	}

	if !dara.IsNil(request.NetType) {
		query["NetType"] = request.NetType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBEndpointAddress"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBEndpointAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases the public endpoints of a PolarDB cluster, including the primary endpoint, default cluster endpoint, and custom cluster endpoint.
//
// Description:
//
// > 	- You can delete a public-facing or classic network endpoint of the primary endpoint, the default cluster endpoint, or a custom cluster endpoint.
//
// > 	- Classic network endpoints are supported only on the China site (aliyun.com). Therefore, you do not need to delete classic network endpoints on the International site (alibabacloud.com).
//
// @param request - DeleteDBEndpointAddressRequest
//
// @return DeleteDBEndpointAddressResponse
func (client *Client) DeleteDBEndpointAddress(request *DeleteDBEndpointAddressRequest) (_result *DeleteDBEndpointAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDBEndpointAddressResponse{}
	_body, _err := client.DeleteDBEndpointAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a database link from a PolarDB for PostgreSQL (Compatible with Oracle) cluster.
//
// @param request - DeleteDBLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBLinkResponse
func (client *Client) DeleteDBLinkWithOptions(request *DeleteDBLinkRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBLinkName) {
		query["DBLinkName"] = request.DBLinkName
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBLink"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a database link from a PolarDB for PostgreSQL (Compatible with Oracle) cluster.
//
// @param request - DeleteDBLinkRequest
//
// @return DeleteDBLinkResponse
func (client *Client) DeleteDBLink(request *DeleteDBLinkRequest) (_result *DeleteDBLinkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDBLinkResponse{}
	_body, _err := client.DeleteDBLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a read-only node from a PolarDB cluster.
//
// @param request - DeleteDBNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBNodesResponse
func (client *Client) DeleteDBNodesWithOptions(request *DeleteDBNodesRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CloudProvider) {
		query["CloudProvider"] = request.CloudProvider
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNodeId) {
		query["DBNodeId"] = request.DBNodeId
	}

	if !dara.IsNil(request.DBNodeType) {
		query["DBNodeType"] = request.DBNodeType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBNodes"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a read-only node from a PolarDB cluster.
//
// @param request - DeleteDBNodesRequest
//
// @return DeleteDBNodesResponse
func (client *Client) DeleteDBNodes(request *DeleteDBNodesRequest) (_result *DeleteDBNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDBNodesResponse{}
	_body, _err := client.DeleteDBNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a database from a PolarDB cluster.
//
// Description:
//
// >- The cluster must be in the Running state and unlocked. Otherwise, the specified database cannot be deleted.
//
// >- The delete operation is performed in an asynchronous manner. A long period of time may be required to delete a large database. A success response for this operation only indicates that the request to delete the database is sent. You must query the database to check whether the database is deleted.
//
// @param request - DeleteDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatabaseResponse
func (client *Client) DeleteDatabaseWithOptions(request *DeleteDatabaseRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatabase"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a database from a PolarDB cluster.
//
// Description:
//
// >- The cluster must be in the Running state and unlocked. Otherwise, the specified database cannot be deleted.
//
// >- The delete operation is performed in an asynchronous manner. A long period of time may be required to delete a large database. A success response for this operation only indicates that the request to delete the database is sent. You must query the database to check whether the database is deleted.
//
// @param request - DeleteDatabaseRequest
//
// @return DeleteDatabaseResponse
func (client *Client) DeleteDatabase(request *DeleteDatabaseRequest) (_result *DeleteDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.DeleteDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除PolarDB边缘云集群数据库
//
// @param request - DeleteDatabaseZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatabaseZonalResponse
func (client *Client) DeleteDatabaseZonalWithOptions(request *DeleteDatabaseZonalRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatabaseZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatabaseZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatabaseZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除PolarDB边缘云集群数据库
//
// @param request - DeleteDatabaseZonalRequest
//
// @return DeleteDatabaseZonalResponse
func (client *Client) DeleteDatabaseZonal(request *DeleteDatabaseZonalRequest) (_result *DeleteDatabaseZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDatabaseZonalResponse{}
	_body, _err := client.DeleteDatabaseZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除角色
//
// @param request - DeleteEncryptionDBRolePrivilegeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEncryptionDBRolePrivilegeResponse
func (client *Client) DeleteEncryptionDBRolePrivilegeWithOptions(request *DeleteEncryptionDBRolePrivilegeRequest, runtime *dara.RuntimeOptions) (_result *DeleteEncryptionDBRolePrivilegeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.RolePrivilegeNameList) {
		query["RolePrivilegeNameList"] = request.RolePrivilegeNameList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEncryptionDBRolePrivilege"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEncryptionDBRolePrivilegeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除角色
//
// @param request - DeleteEncryptionDBRolePrivilegeRequest
//
// @return DeleteEncryptionDBRolePrivilegeResponse
func (client *Client) DeleteEncryptionDBRolePrivilege(request *DeleteEncryptionDBRolePrivilegeRequest) (_result *DeleteEncryptionDBRolePrivilegeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEncryptionDBRolePrivilegeResponse{}
	_body, _err := client.DeleteEncryptionDBRolePrivilegeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除插件
//
// @param request - DeleteExtensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExtensionsResponse
func (client *Client) DeleteExtensionsWithOptions(request *DeleteExtensionsRequest, runtime *dara.RuntimeOptions) (_result *DeleteExtensionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNames) {
		query["DBNames"] = request.DBNames
	}

	if !dara.IsNil(request.Extensions) {
		query["Extensions"] = request.Extensions
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

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExtensions"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExtensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除插件
//
// @param request - DeleteExtensionsRequest
//
// @return DeleteExtensionsResponse
func (client *Client) DeleteExtensions(request *DeleteExtensionsRequest) (_result *DeleteExtensionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteExtensionsResponse{}
	_body, _err := client.DeleteExtensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除sql防火墙
//
// @param request - DeleteFirewallRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFirewallRulesResponse
func (client *Client) DeleteFirewallRulesWithOptions(request *DeleteFirewallRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteFirewallRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.RuleNameList) {
		query["RuleNameList"] = request.RuleNameList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFirewallRules"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFirewallRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除sql防火墙
//
// @param request - DeleteFirewallRulesRequest
//
// @return DeleteFirewallRulesResponse
func (client *Client) DeleteFirewallRules(request *DeleteFirewallRulesRequest) (_result *DeleteFirewallRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFirewallRulesResponse{}
	_body, _err := client.DeleteFirewallRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除网关实例
//
// @param request - DeleteGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayResponse
func (client *Client) DeleteGatewayWithOptions(request *DeleteGatewayRequest, runtime *dara.RuntimeOptions) (_result *DeleteGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGateway"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除网关实例
//
// @param request - DeleteGatewayRequest
//
// @return DeleteGatewayResponse
func (client *Client) DeleteGateway(request *DeleteGatewayRequest) (_result *DeleteGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.DeleteGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a global data network (GDN).
//
// @param request - DeleteGlobalDataNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGlobalDataNetworkResponse
func (client *Client) DeleteGlobalDataNetworkWithOptions(request *DeleteGlobalDataNetworkRequest, runtime *dara.RuntimeOptions) (_result *DeleteGlobalDataNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkId) {
		query["NetworkId"] = request.NetworkId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGlobalDataNetwork"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGlobalDataNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a global data network (GDN).
//
// @param request - DeleteGlobalDataNetworkRequest
//
// @return DeleteGlobalDataNetworkResponse
func (client *Client) DeleteGlobalDataNetwork(request *DeleteGlobalDataNetworkRequest) (_result *DeleteGlobalDataNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGlobalDataNetworkResponse{}
	_body, _err := client.DeleteGlobalDataNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a global database network (GDN).
//
// Description:
//
// >  You can delete a GDN only when the GDN includes only a primary cluster.
//
// @param request - DeleteGlobalDatabaseNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGlobalDatabaseNetworkResponse
func (client *Client) DeleteGlobalDatabaseNetworkWithOptions(request *DeleteGlobalDatabaseNetworkRequest, runtime *dara.RuntimeOptions) (_result *DeleteGlobalDatabaseNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GDNId) {
		query["GDNId"] = request.GDNId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGlobalDatabaseNetwork"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGlobalDatabaseNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a global database network (GDN).
//
// Description:
//
// >  You can delete a GDN only when the GDN includes only a primary cluster.
//
// @param request - DeleteGlobalDatabaseNetworkRequest
//
// @return DeleteGlobalDatabaseNetworkResponse
func (client *Client) DeleteGlobalDatabaseNetwork(request *DeleteGlobalDatabaseNetworkRequest) (_result *DeleteGlobalDatabaseNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGlobalDatabaseNetworkResponse{}
	_body, _err := client.DeleteGlobalDatabaseNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a global IP whitelist template.
//
// @param request - DeleteGlobalSecurityIPGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGlobalSecurityIPGroupResponse
func (client *Client) DeleteGlobalSecurityIPGroupWithOptions(request *DeleteGlobalSecurityIPGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteGlobalSecurityIPGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GlobalIgName) {
		query["GlobalIgName"] = request.GlobalIgName
	}

	if !dara.IsNil(request.GlobalSecurityGroupId) {
		query["GlobalSecurityGroupId"] = request.GlobalSecurityGroupId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGlobalSecurityIPGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGlobalSecurityIPGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a global IP whitelist template.
//
// @param request - DeleteGlobalSecurityIPGroupRequest
//
// @return DeleteGlobalSecurityIPGroupResponse
func (client *Client) DeleteGlobalSecurityIPGroup(request *DeleteGlobalSecurityIPGroupRequest) (_result *DeleteGlobalSecurityIPGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGlobalSecurityIPGroupResponse{}
	_body, _err := client.DeleteGlobalSecurityIPGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data masking rule.
//
// @param request - DeleteMaskingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMaskingRulesResponse
func (client *Client) DeleteMaskingRulesWithOptions(request *DeleteMaskingRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteMaskingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.InterfaceVersion) {
		query["InterfaceVersion"] = request.InterfaceVersion
	}

	if !dara.IsNil(request.RuleNameList) {
		query["RuleNameList"] = request.RuleNameList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMaskingRules"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMaskingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data masking rule.
//
// @param request - DeleteMaskingRulesRequest
//
// @return DeleteMaskingRulesResponse
func (client *Client) DeleteMaskingRules(request *DeleteMaskingRulesRequest) (_result *DeleteMaskingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMaskingRulesResponse{}
	_body, _err := client.DeleteMaskingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除路由规则
//
// @param request - DeleteModelApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelApiResponse
func (client *Client) DeleteModelApiWithOptions(request *DeleteModelApiRequest, runtime *dara.RuntimeOptions) (_result *DeleteModelApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.ModelApiId) {
		query["ModelApiId"] = request.ModelApiId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModelApi"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除路由规则
//
// @param request - DeleteModelApiRequest
//
// @return DeleteModelApiResponse
func (client *Client) DeleteModelApi(request *DeleteModelApiRequest) (_result *DeleteModelApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteModelApiResponse{}
	_body, _err := client.DeleteModelApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除模型服务
//
// @param request - DeleteModelServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelServiceResponse
func (client *Client) DeleteModelServiceWithOptions(request *DeleteModelServiceRequest, runtime *dara.RuntimeOptions) (_result *DeleteModelServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModelService"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除模型服务
//
// @param request - DeleteModelServiceRequest
//
// @return DeleteModelServiceResponse
func (client *Client) DeleteModelService(request *DeleteModelServiceRequest) (_result *DeleteModelServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteModelServiceResponse{}
	_body, _err := client.DeleteModelServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除网络通道
//
// @param request - DeleteNetworkChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkChannelResponse
func (client *Client) DeleteNetworkChannelWithOptions(request *DeleteNetworkChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNetworkChannel"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetworkChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除网络通道
//
// @param request - DeleteNetworkChannelRequest
//
// @return DeleteNetworkChannelResponse
func (client *Client) DeleteNetworkChannel(request *DeleteNetworkChannelRequest) (_result *DeleteNetworkChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNetworkChannelResponse{}
	_body, _err := client.DeleteNetworkChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a parameter template from a PolarDB cluster.
//
// Description:
//
// You can use parameter templates to manage multiple parameters at a time and quickly apply existing parameters to a PolarDB cluster. For more information, see [Use a parameter template](https://help.aliyun.com/document_detail/207009.html).
//
// >  When you delete a parameter template, the parameter settings that are applied to PolarDB clusters are not affected.
//
// @param request - DeleteParameterGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteParameterGroupResponse
func (client *Client) DeleteParameterGroupWithOptions(request *DeleteParameterGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteParameterGroupResponse, _err error) {
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

	if !dara.IsNil(request.ParameterGroupId) {
		query["ParameterGroupId"] = request.ParameterGroupId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteParameterGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteParameterGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a parameter template from a PolarDB cluster.
//
// Description:
//
// You can use parameter templates to manage multiple parameters at a time and quickly apply existing parameters to a PolarDB cluster. For more information, see [Use a parameter template](https://help.aliyun.com/document_detail/207009.html).
//
// >  When you delete a parameter template, the parameter settings that are applied to PolarDB clusters are not affected.
//
// @param request - DeleteParameterGroupRequest
//
// @return DeleteParameterGroupResponse
func (client *Client) DeleteParameterGroup(request *DeleteParameterGroupRequest) (_result *DeleteParameterGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteParameterGroupResponse{}
	_body, _err := client.DeleteParameterGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除PolarClaw Agent
//
// @param request - DeletePolarClawAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolarClawAgentResponse
func (client *Client) DeletePolarClawAgentWithOptions(request *DeletePolarClawAgentRequest, runtime *dara.RuntimeOptions) (_result *DeletePolarClawAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.DeleteFiles) {
		query["DeleteFiles"] = request.DeleteFiles
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolarClawAgent"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolarClawAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除PolarClaw Agent
//
// @param request - DeletePolarClawAgentRequest
//
// @return DeletePolarClawAgentResponse
func (client *Client) DeletePolarClawAgent(request *DeletePolarClawAgentRequest) (_result *DeletePolarClawAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePolarClawAgentResponse{}
	_body, _err := client.DeletePolarClawAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除PolarClaw Channel
//
// @param request - DeletePolarClawChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolarClawChannelResponse
func (client *Client) DeletePolarClawChannelWithOptions(request *DeletePolarClawChannelRequest, runtime *dara.RuntimeOptions) (_result *DeletePolarClawChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	if !dara.IsNil(request.UninstallPlugin) {
		query["UninstallPlugin"] = request.UninstallPlugin
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolarClawChannel"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolarClawChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除PolarClaw Channel
//
// @param request - DeletePolarClawChannelRequest
//
// @return DeletePolarClawChannelResponse
func (client *Client) DeletePolarClawChannel(request *DeletePolarClawChannelRequest) (_result *DeletePolarClawChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePolarClawChannelResponse{}
	_body, _err := client.DeletePolarClawChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除PolarClaw定时任务
//
// @param request - DeletePolarClawCronJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolarClawCronJobResponse
func (client *Client) DeletePolarClawCronJobWithOptions(request *DeletePolarClawCronJobRequest, runtime *dara.RuntimeOptions) (_result *DeletePolarClawCronJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolarClawCronJob"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolarClawCronJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除PolarClaw定时任务
//
// @param request - DeletePolarClawCronJobRequest
//
// @return DeletePolarClawCronJobResponse
func (client *Client) DeletePolarClawCronJob(request *DeletePolarClawCronJobRequest) (_result *DeletePolarClawCronJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePolarClawCronJobResponse{}
	_body, _err := client.DeletePolarClawCronJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除PolarFs文件
//
// Description:
//
// ## 请求说明
//
// - `PolarFsInstanceId` 是必须提供的参数，用来指定要操作的PolarFS实例。
//
// - `DBClusterId` 参数是可选的，如果提供，则表示与特定PolarDB集群关联的操作。
//
// - `Objects` 参数是一个字符串数组，列出了所有需要被删除的对象路径，并且是必需的。
//
// @param tmpReq - DeletePolarFsObjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolarFsObjectsResponse
func (client *Client) DeletePolarFsObjectsWithOptions(tmpReq *DeletePolarFsObjectsRequest, runtime *dara.RuntimeOptions) (_result *DeletePolarFsObjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeletePolarFsObjectsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ObjectsToDelete) {
		request.ObjectsToDeleteShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ObjectsToDelete, dara.String("ObjectsToDelete"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ObjectsToDeleteShrink) {
		query["ObjectsToDelete"] = request.ObjectsToDeleteShrink
	}

	if !dara.IsNil(request.PolarFsInstanceId) {
		query["PolarFsInstanceId"] = request.PolarFsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolarFsObjects"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolarFsObjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除PolarFs文件
//
// Description:
//
// ## 请求说明
//
// - `PolarFsInstanceId` 是必须提供的参数，用来指定要操作的PolarFS实例。
//
// - `DBClusterId` 参数是可选的，如果提供，则表示与特定PolarDB集群关联的操作。
//
// - `Objects` 参数是一个字符串数组，列出了所有需要被删除的对象路径，并且是必需的。
//
// @param request - DeletePolarFsObjectsRequest
//
// @return DeletePolarFsObjectsResponse
func (client *Client) DeletePolarFsObjects(request *DeletePolarFsObjectsRequest) (_result *DeletePolarFsObjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePolarFsObjectsResponse{}
	_body, _err := client.DeletePolarFsObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除polar fs bucket路径
//
// @param request - DeletePolarFsPathMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolarFsPathMappingResponse
func (client *Client) DeletePolarFsPathMappingWithOptions(request *DeletePolarFsPathMappingRequest, runtime *dara.RuntimeOptions) (_result *DeletePolarFsPathMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomBucketPathList) {
		query["CustomBucketPathList"] = request.CustomBucketPathList
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.PolarFsInstanceId) {
		query["PolarFsInstanceId"] = request.PolarFsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolarFsPathMapping"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolarFsPathMappingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除polar fs bucket路径
//
// @param request - DeletePolarFsPathMappingRequest
//
// @return DeletePolarFsPathMappingResponse
func (client *Client) DeletePolarFsPathMapping(request *DeletePolarFsPathMappingRequest) (_result *DeletePolarFsPathMappingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePolarFsPathMappingResponse{}
	_body, _err := client.DeletePolarFsPathMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除PolarFs Quota规则
//
// @param request - DeletePolarFsQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolarFsQuotaResponse
func (client *Client) DeletePolarFsQuotaWithOptions(request *DeletePolarFsQuotaRequest, runtime *dara.RuntimeOptions) (_result *DeletePolarFsQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.PolarFsInstanceId) {
		query["PolarFsInstanceId"] = request.PolarFsInstanceId
	}

	if !dara.IsNil(request.Quotas) {
		query["Quotas"] = request.Quotas
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolarFsQuota"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolarFsQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除PolarFs Quota规则
//
// @param request - DeletePolarFsQuotaRequest
//
// @return DeletePolarFsQuotaResponse
func (client *Client) DeletePolarFsQuota(request *DeletePolarFsQuotaRequest) (_result *DeletePolarFsQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePolarFsQuotaResponse{}
	_body, _err := client.DeletePolarFsQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除限流策略
//
// @param request - DeleteRateLimitPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRateLimitPolicyResponse
func (client *Client) DeleteRateLimitPolicyWithOptions(request *DeleteRateLimitPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteRateLimitPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRateLimitPolicy"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRateLimitPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除限流策略
//
// @param request - DeleteRateLimitPolicyRequest
//
// @return DeleteRateLimitPolicyResponse
func (client *Client) DeleteRateLimitPolicy(request *DeleteRateLimitPolicyRequest) (_result *DeleteRateLimitPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRateLimitPolicyResponse{}
	_body, _err := client.DeleteRateLimitPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除SQL限流规则
//
// @param request - DeleteSQLRateLimitingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSQLRateLimitingRulesResponse
func (client *Client) DeleteSQLRateLimitingRulesWithOptions(request *DeleteSQLRateLimitingRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteSQLRateLimitingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.RuleNameList) {
		query["RuleNameList"] = request.RuleNameList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSQLRateLimitingRules"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSQLRateLimitingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除SQL限流规则
//
// @param request - DeleteSQLRateLimitingRulesRequest
//
// @return DeleteSQLRateLimitingRulesResponse
func (client *Client) DeleteSQLRateLimitingRules(request *DeleteSQLRateLimitingRulesRequest) (_result *DeleteSQLRateLimitingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSQLRateLimitingRulesResponse{}
	_body, _err := client.DeleteSQLRateLimitingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看custom实例详情
//
// @param request - DescribeAIDBClusterAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAIDBClusterAttributeResponse
func (client *Client) DescribeAIDBClusterAttributeWithOptions(request *DescribeAIDBClusterAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeAIDBClusterAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAIDBClusterAttribute"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAIDBClusterAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看custom实例详情
//
// @param request - DescribeAIDBClusterAttributeRequest
//
// @return DescribeAIDBClusterAttributeResponse
func (client *Client) DescribeAIDBClusterAttribute(request *DescribeAIDBClusterAttributeRequest) (_result *DescribeAIDBClusterAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAIDBClusterAttributeResponse{}
	_body, _err := client.DescribeAIDBClusterAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据集列表
//
// @param request - DescribeAIDBClusterDatasetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAIDBClusterDatasetsResponse
func (client *Client) DescribeAIDBClusterDatasetsWithOptions(request *DescribeAIDBClusterDatasetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAIDBClusterDatasetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContinuationToken) {
		query["ContinuationToken"] = request.ContinuationToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetType) {
		query["DatasetType"] = request.DatasetType
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

	if !dara.IsNil(request.TrainMode) {
		query["TrainMode"] = request.TrainMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAIDBClusterDatasets"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAIDBClusterDatasetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据集列表
//
// @param request - DescribeAIDBClusterDatasetsRequest
//
// @return DescribeAIDBClusterDatasetsResponse
func (client *Client) DescribeAIDBClusterDatasets(request *DescribeAIDBClusterDatasetsRequest) (_result *DescribeAIDBClusterDatasetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAIDBClusterDatasetsResponse{}
	_body, _err := client.DescribeAIDBClusterDatasetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取ai容器性能指标
//
// @param request - DescribeAIDBClusterPerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAIDBClusterPerformanceResponse
func (client *Client) DescribeAIDBClusterPerformanceWithOptions(request *DescribeAIDBClusterPerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeAIDBClusterPerformanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAIDBClusterPerformance"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAIDBClusterPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取ai容器性能指标
//
// @param request - DescribeAIDBClusterPerformanceRequest
//
// @return DescribeAIDBClusterPerformanceResponse
func (client *Client) DescribeAIDBClusterPerformance(request *DescribeAIDBClusterPerformanceRequest) (_result *DescribeAIDBClusterPerformanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAIDBClusterPerformanceResponse{}
	_body, _err := client.DescribeAIDBClusterPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务实例详情
//
// @param request - DescribeAIDBClusterTaskAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAIDBClusterTaskAttributeResponse
func (client *Client) DescribeAIDBClusterTaskAttributeWithOptions(request *DescribeAIDBClusterTaskAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeAIDBClusterTaskAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DescribeType) {
		query["DescribeType"] = request.DescribeType
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
		Action:      dara.String("DescribeAIDBClusterTaskAttribute"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAIDBClusterTaskAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务实例详情
//
// @param request - DescribeAIDBClusterTaskAttributeRequest
//
// @return DescribeAIDBClusterTaskAttributeResponse
func (client *Client) DescribeAIDBClusterTaskAttribute(request *DescribeAIDBClusterTaskAttributeRequest) (_result *DescribeAIDBClusterTaskAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAIDBClusterTaskAttributeResponse{}
	_body, _err := client.DescribeAIDBClusterTaskAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询模型算子日志
//
// @param request - DescribeAIDBClusterTaskLogFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAIDBClusterTaskLogFilesResponse
func (client *Client) DescribeAIDBClusterTaskLogFilesWithOptions(request *DescribeAIDBClusterTaskLogFilesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAIDBClusterTaskLogFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.LogType) {
		query["LogType"] = request.LogType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RelativeDBClusterId) {
		query["RelativeDBClusterId"] = request.RelativeDBClusterId
	}

	if !dara.IsNil(request.Reverse) {
		query["Reverse"] = request.Reverse
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAIDBClusterTaskLogFiles"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAIDBClusterTaskLogFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询模型算子日志
//
// @param request - DescribeAIDBClusterTaskLogFilesRequest
//
// @return DescribeAIDBClusterTaskLogFilesResponse
func (client *Client) DescribeAIDBClusterTaskLogFiles(request *DescribeAIDBClusterTaskLogFilesRequest) (_result *DescribeAIDBClusterTaskLogFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAIDBClusterTaskLogFilesResponse{}
	_body, _err := client.DescribeAIDBClusterTaskLogFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询模型算子指标
//
// @param request - DescribeAIDBClusterTaskMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAIDBClusterTaskMetricsResponse
func (client *Client) DescribeAIDBClusterTaskMetricsWithOptions(request *DescribeAIDBClusterTaskMetricsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAIDBClusterTaskMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MetricType) {
		query["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RelativeDBClusterId) {
		query["RelativeDBClusterId"] = request.RelativeDBClusterId
	}

	if !dara.IsNil(request.Reverse) {
		query["Reverse"] = request.Reverse
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAIDBClusterTaskMetrics"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAIDBClusterTaskMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询模型算子指标
//
// @param request - DescribeAIDBClusterTaskMetricsRequest
//
// @return DescribeAIDBClusterTaskMetricsResponse
func (client *Client) DescribeAIDBClusterTaskMetrics(request *DescribeAIDBClusterTaskMetricsRequest) (_result *DescribeAIDBClusterTaskMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAIDBClusterTaskMetricsResponse{}
	_body, _err := client.DescribeAIDBClusterTaskMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务实例列表
//
// @param request - DescribeAIDBClusterTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAIDBClusterTasksResponse
func (client *Client) DescribeAIDBClusterTasksWithOptions(request *DescribeAIDBClusterTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeAIDBClusterTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KubeType) {
		query["KubeType"] = request.KubeType
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

	if !dara.IsNil(request.RelativeDBClusterId) {
		query["RelativeDBClusterId"] = request.RelativeDBClusterId
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
		Action:      dara.String("DescribeAIDBClusterTasks"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAIDBClusterTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务实例列表
//
// @param request - DescribeAIDBClusterTasksRequest
//
// @return DescribeAIDBClusterTasksResponse
func (client *Client) DescribeAIDBClusterTasks(request *DescribeAIDBClusterTasksRequest) (_result *DescribeAIDBClusterTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAIDBClusterTasksResponse{}
	_body, _err := client.DescribeAIDBClusterTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看custom集群列表
//
// @param request - DescribeAIDBClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAIDBClustersResponse
func (client *Client) DescribeAIDBClustersWithOptions(request *DescribeAIDBClustersRequest, runtime *dara.RuntimeOptions) (_result *DescribeAIDBClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AiNodeType) {
		query["AiNodeType"] = request.AiNodeType
	}

	if !dara.IsNil(request.DBClusterDescription) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !dara.IsNil(request.DBClusterIds) {
		query["DBClusterIds"] = request.DBClusterIds
	}

	if !dara.IsNil(request.DBClusterStatus) {
		query["DBClusterStatus"] = request.DBClusterStatus
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
		Action:      dara.String("DescribeAIDBClusters"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAIDBClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看custom集群列表
//
// @param request - DescribeAIDBClustersRequest
//
// @return DescribeAIDBClustersResponse
func (client *Client) DescribeAIDBClusters(request *DescribeAIDBClustersRequest) (_result *DescribeAIDBClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAIDBClustersResponse{}
	_body, _err := client.DescribeAIDBClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the PolarDB for AI feature.
//
// @param request - DescribeAITaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAITaskStatusResponse
func (client *Client) DescribeAITaskStatusWithOptions(request *DescribeAITaskStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAITaskStatusResponse, _err error) {
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
		Action:      dara.String("DescribeAITaskStatus"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAITaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the PolarDB for AI feature.
//
// @param request - DescribeAITaskStatusRequest
//
// @return DescribeAITaskStatusResponse
func (client *Client) DescribeAITaskStatus(request *DescribeAITaskStatusRequest) (_result *DescribeAITaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAITaskStatusResponse{}
	_body, _err := client.DescribeAITaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a database account of a PolarDB cluster.
//
// @param request - DescribeAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountsResponse
func (client *Client) DescribeAccountsWithOptions(request *DescribeAccountsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
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
		Action:      dara.String("DescribeAccounts"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a database account of a PolarDB cluster.
//
// @param request - DescribeAccountsRequest
//
// @return DescribeAccountsResponse
func (client *Client) DescribeAccounts(request *DescribeAccountsRequest) (_result *DescribeAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DescribeAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询PolarDB边缘云集群列表
//
// @param request - DescribeAccountsZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountsZonalResponse
func (client *Client) DescribeAccountsZonalWithOptions(request *DescribeAccountsZonalRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountsZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
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
		Action:      dara.String("DescribeAccountsZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccountsZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询PolarDB边缘云集群列表
//
// @param request - DescribeAccountsZonalRequest
//
// @return DescribeAccountsZonalResponse
func (client *Client) DescribeAccountsZonal(request *DescribeAccountsZonalRequest) (_result *DescribeAccountsZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccountsZonalResponse{}
	_body, _err := client.DescribeAccountsZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an activation code.
//
// @param request - DescribeActivationCodeDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActivationCodeDetailsResponse
func (client *Client) DescribeActivationCodeDetailsWithOptions(request *DescribeActivationCodeDetailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeActivationCodeDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivationCodeId) {
		query["ActivationCodeId"] = request.ActivationCodeId
	}

	if !dara.IsNil(request.AliyunOrderId) {
		query["AliyunOrderId"] = request.AliyunOrderId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeActivationCodeDetails"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActivationCodeDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an activation code.
//
// @param request - DescribeActivationCodeDetailsRequest
//
// @return DescribeActivationCodeDetailsResponse
func (client *Client) DescribeActivationCodeDetails(request *DescribeActivationCodeDetailsRequest) (_result *DescribeActivationCodeDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeActivationCodeDetailsResponse{}
	_body, _err := client.DescribeActivationCodeDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of activation codes.
//
// @param request - DescribeActivationCodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActivationCodesResponse
func (client *Client) DescribeActivationCodesWithOptions(request *DescribeActivationCodesRequest, runtime *dara.RuntimeOptions) (_result *DescribeActivationCodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunOrderId) {
		query["AliyunOrderId"] = request.AliyunOrderId
	}

	if !dara.IsNil(request.MacAddress) {
		query["MacAddress"] = request.MacAddress
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SystemIdentifier) {
		query["SystemIdentifier"] = request.SystemIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeActivationCodes"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActivationCodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of activation codes.
//
// @param request - DescribeActivationCodesRequest
//
// @return DescribeActivationCodesResponse
func (client *Client) DescribeActivationCodes(request *DescribeActivationCodesRequest) (_result *DescribeActivationCodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeActivationCodesResponse{}
	_body, _err := client.DescribeActivationCodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户的运维配置信息，目前包括主动运维窗口信息
//
// @param request - DescribeActiveOperationMaintainConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationMaintainConfResponse
func (client *Client) DescribeActiveOperationMaintainConfWithOptions(request *DescribeActiveOperationMaintainConfRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveOperationMaintainConfResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeActiveOperationMaintainConf"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActiveOperationMaintainConfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户的运维配置信息，目前包括主动运维窗口信息
//
// @param request - DescribeActiveOperationMaintainConfRequest
//
// @return DescribeActiveOperationMaintainConfResponse
func (client *Client) DescribeActiveOperationMaintainConf(request *DescribeActiveOperationMaintainConfRequest) (_result *DescribeActiveOperationMaintainConfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeActiveOperationMaintainConfResponse{}
	_body, _err := client.DescribeActiveOperationMaintainConfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the O\\\\\\&M event details of an instance.
//
// @param request - DescribeActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationTasksResponse
func (client *Client) DescribeActiveOperationTasksWithOptions(request *DescribeActiveOperationTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveOperationTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowCancel) {
		query["AllowCancel"] = request.AllowCancel
	}

	if !dara.IsNil(request.AllowChange) {
		query["AllowChange"] = request.AllowChange
	}

	if !dara.IsNil(request.ChangeLevel) {
		query["ChangeLevel"] = request.ChangeLevel
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBType) {
		query["DBType"] = request.DBType
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeActiveOperationTasks"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActiveOperationTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the O\\\\\\&M event details of an instance.
//
// @param request - DescribeActiveOperationTasksRequest
//
// @return DescribeActiveOperationTasksResponse
func (client *Client) DescribeActiveOperationTasks(request *DescribeActiveOperationTasksRequest) (_result *DescribeActiveOperationTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeActiveOperationTasksResponse{}
	_body, _err := client.DescribeActiveOperationTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取应用详情
//
// @param request - DescribeApplicationAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationAttributeResponse
func (client *Client) DescribeApplicationAttributeWithOptions(request *DescribeApplicationAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeApplicationAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplicationAttribute"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用详情
//
// @param request - DescribeApplicationAttributeRequest
//
// @return DescribeApplicationAttributeResponse
func (client *Client) DescribeApplicationAttribute(request *DescribeApplicationAttributeRequest) (_result *DescribeApplicationAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApplicationAttributeResponse{}
	_body, _err := client.DescribeApplicationAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # AI 应用日志明细
//
// @param request - DescribeApplicationLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationLogsResponse
func (client *Client) DescribeApplicationLogsWithOptions(request *DescribeApplicationLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApplicationLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ComponentName) {
		query["ComponentName"] = request.ComponentName
	}

	if !dara.IsNil(request.ContainerName) {
		query["ContainerName"] = request.ContainerName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
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
		Action:      dara.String("DescribeApplicationLogs"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AI 应用日志明细
//
// @param request - DescribeApplicationLogsRequest
//
// @return DescribeApplicationLogsResponse
func (client *Client) DescribeApplicationLogs(request *DescribeApplicationLogsRequest) (_result *DescribeApplicationLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApplicationLogsResponse{}
	_body, _err := client.DescribeApplicationLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取应用组件参数
//
// @param tmpReq - DescribeApplicationParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationParametersResponse
func (client *Client) DescribeApplicationParametersWithOptions(tmpReq *DescribeApplicationParametersRequest, runtime *dara.RuntimeOptions) (_result *DescribeApplicationParametersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeApplicationParametersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ComponentIdList) {
		request.ComponentIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ComponentIdList, dara.String("ComponentIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ComponentIdListShrink) {
		query["ComponentIdList"] = request.ComponentIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplicationParameters"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用组件参数
//
// @param request - DescribeApplicationParametersRequest
//
// @return DescribeApplicationParametersResponse
func (client *Client) DescribeApplicationParameters(request *DescribeApplicationParametersRequest) (_result *DescribeApplicationParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApplicationParametersResponse{}
	_body, _err := client.DescribeApplicationParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询当前应用下所有的应用提示词策略列表
//
// @param request - DescribeApplicationPromptsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationPromptsResponse
func (client *Client) DescribeApplicationPromptsWithOptions(request *DescribeApplicationPromptsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApplicationPromptsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplicationPrompts"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationPromptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询当前应用下所有的应用提示词策略列表
//
// @param request - DescribeApplicationPromptsRequest
//
// @return DescribeApplicationPromptsResponse
func (client *Client) DescribeApplicationPrompts(request *DescribeApplicationPromptsRequest) (_result *DescribeApplicationPromptsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApplicationPromptsResponse{}
	_body, _err := client.DescribeApplicationPromptsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取应用serverless配置
//
// @param request - DescribeApplicationServerlessConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationServerlessConfResponse
func (client *Client) DescribeApplicationServerlessConfWithOptions(request *DescribeApplicationServerlessConfRequest, runtime *dara.RuntimeOptions) (_result *DescribeApplicationServerlessConfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplicationServerlessConf"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationServerlessConfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用serverless配置
//
// @param request - DescribeApplicationServerlessConfRequest
//
// @return DescribeApplicationServerlessConfResponse
func (client *Client) DescribeApplicationServerlessConf(request *DescribeApplicationServerlessConfRequest) (_result *DescribeApplicationServerlessConfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApplicationServerlessConfResponse{}
	_body, _err := client.DescribeApplicationServerlessConfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取当前地域所有PolarDB实例的应用列表
//
// @param request - DescribeApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationsResponse
func (client *Client) DescribeApplicationsWithOptions(request *DescribeApplicationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApplicationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		query["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.ApplicationTypes) {
		query["ApplicationTypes"] = request.ApplicationTypes
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplications"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当前地域所有PolarDB实例的应用列表
//
// @param request - DescribeApplicationsRequest
//
// @return DescribeApplicationsResponse
func (client *Client) DescribeApplications(request *DescribeApplicationsRequest) (_result *DescribeApplicationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApplicationsResponse{}
	_body, _err := client.DescribeApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the auto-renewal attributes of a subscription PolarDB cluster.
//
// @param request - DescribeAutoRenewAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoRenewAttributeResponse
func (client *Client) DescribeAutoRenewAttributeWithOptions(request *DescribeAutoRenewAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoRenewAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloudProvider) {
		query["CloudProvider"] = request.CloudProvider
	}

	if !dara.IsNil(request.DBClusterIds) {
		query["DBClusterIds"] = request.DBClusterIds
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAutoRenewAttribute"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAutoRenewAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the auto-renewal attributes of a subscription PolarDB cluster.
//
// @param request - DescribeAutoRenewAttributeRequest
//
// @return DescribeAutoRenewAttributeResponse
func (client *Client) DescribeAutoRenewAttribute(request *DescribeAutoRenewAttributeRequest) (_result *DescribeAutoRenewAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAutoRenewAttributeResponse{}
	_body, _err := client.DescribeAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询可用跨地域备份地域列表
//
// @param request - DescribeAvailableCrossRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableCrossRegionsResponse
func (client *Client) DescribeAvailableCrossRegionsWithOptions(request *DescribeAvailableCrossRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailableCrossRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAvailableCrossRegions"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvailableCrossRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询可用跨地域备份地域列表
//
// @param request - DescribeAvailableCrossRegionsRequest
//
// @return DescribeAvailableCrossRegionsResponse
func (client *Client) DescribeAvailableCrossRegions(request *DescribeAvailableCrossRegionsRequest) (_result *DescribeAvailableCrossRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAvailableCrossRegionsResponse{}
	_body, _err := client.DescribeAvailableCrossRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取ai集群模型列表
//
// @param request - DescribeAvailableModelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableModelsResponse
func (client *Client) DescribeAvailableModelsWithOptions(request *DescribeAvailableModelsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailableModelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KubeType) {
		query["KubeType"] = request.KubeType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAvailableModels"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvailableModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取ai集群模型列表
//
// @param request - DescribeAvailableModelsRequest
//
// @return DescribeAvailableModelsResponse
func (client *Client) DescribeAvailableModels(request *DescribeAvailableModelsRequest) (_result *DescribeAvailableModelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAvailableModelsResponse{}
	_body, _err := client.DescribeAvailableModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries backup logs and the URLs to download the backup logs.
//
// @param request - DescribeBackupLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupLogsResponse
func (client *Client) DescribeBackupLogsWithOptions(request *DescribeBackupLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupRegion) {
		query["BackupRegion"] = request.BackupRegion
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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
		Action:      dara.String("DescribeBackupLogs"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries backup logs and the URLs to download the backup logs.
//
// @param request - DescribeBackupLogsRequest
//
// @return DescribeBackupLogsResponse
func (client *Client) DescribeBackupLogs(request *DescribeBackupLogsRequest) (_result *DescribeBackupLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupLogsResponse{}
	_body, _err := client.DescribeBackupLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the automatic backup policy of a PolarDB cluster.
//
// @param request - DescribeBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupPolicyResponse
func (client *Client) DescribeBackupPolicyWithOptions(request *DescribeBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupPolicy"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the automatic backup policy of a PolarDB cluster.
//
// @param request - DescribeBackupPolicyRequest
//
// @return DescribeBackupPolicyResponse
func (client *Client) DescribeBackupPolicy(request *DescribeBackupPolicyRequest) (_result *DescribeBackupPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DescribeBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询备份集所在地域信息
//
// @param request - DescribeBackupRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupRegionsResponse
func (client *Client) DescribeBackupRegionsWithOptions(request *DescribeBackupRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupRegions"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询备份集所在地域信息
//
// @param request - DescribeBackupRegionsRequest
//
// @return DescribeBackupRegionsResponse
func (client *Client) DescribeBackupRegions(request *DescribeBackupRegionsRequest) (_result *DescribeBackupRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupRegionsResponse{}
	_body, _err := client.DescribeBackupRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of backup tasks of a cluster.
//
// @param request - DescribeBackupTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupTasksResponse
func (client *Client) DescribeBackupTasksWithOptions(request *DescribeBackupTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupJobId) {
		query["BackupJobId"] = request.BackupJobId
	}

	if !dara.IsNil(request.BackupMode) {
		query["BackupMode"] = request.BackupMode
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupTasks"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of backup tasks of a cluster.
//
// @param request - DescribeBackupTasksRequest
//
// @return DescribeBackupTasksResponse
func (client *Client) DescribeBackupTasks(request *DescribeBackupTasksRequest) (_result *DescribeBackupTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupTasksResponse{}
	_body, _err := client.DescribeBackupTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the backup details of a PolarDB cluster.
//
// @param request - DescribeBackupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupsResponse
func (client *Client) DescribeBackupsWithOptions(request *DescribeBackupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.BackupMode) {
		query["BackupMode"] = request.BackupMode
	}

	if !dara.IsNil(request.BackupRegion) {
		query["BackupRegion"] = request.BackupRegion
	}

	if !dara.IsNil(request.BackupStatus) {
		query["BackupStatus"] = request.BackupStatus
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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
		Action:      dara.String("DescribeBackups"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backup details of a PolarDB cluster.
//
// @param request - DescribeBackupsRequest
//
// @return DescribeBackupsResponse
func (client *Client) DescribeBackups(request *DescribeBackupsRequest) (_result *DescribeBackupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.DescribeBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询预算策略
//
// @param request - DescribeBudgetPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBudgetPoliciesResponse
func (client *Client) DescribeBudgetPoliciesWithOptions(request *DescribeBudgetPoliciesRequest, runtime *dara.RuntimeOptions) (_result *DescribeBudgetPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BudgetDimensionRefId) {
		query["BudgetDimensionRefId"] = request.BudgetDimensionRefId
	}

	if !dara.IsNil(request.BudgetDimensionType) {
		query["BudgetDimensionType"] = request.BudgetDimensionType
	}

	if !dara.IsNil(request.BudgetPolicyId) {
		query["BudgetPolicyId"] = request.BudgetPolicyId
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
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

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBudgetPolicies"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBudgetPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询预算策略
//
// @param request - DescribeBudgetPoliciesRequest
//
// @return DescribeBudgetPoliciesResponse
func (client *Client) DescribeBudgetPolicies(request *DescribeBudgetPoliciesRequest) (_result *DescribeBudgetPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBudgetPoliciesResponse{}
	_body, _err := client.DescribeBudgetPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries character sets that are supported by a PolarDB for MySQL cluster.
//
// @param request - DescribeCharacterSetNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCharacterSetNameResponse
func (client *Client) DescribeCharacterSetNameWithOptions(request *DescribeCharacterSetNameRequest, runtime *dara.RuntimeOptions) (_result *DescribeCharacterSetNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("DescribeCharacterSetName"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCharacterSetNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries character sets that are supported by a PolarDB for MySQL cluster.
//
// @param request - DescribeCharacterSetNameRequest
//
// @return DescribeCharacterSetNameResponse
func (client *Client) DescribeCharacterSetName(request *DescribeCharacterSetNameRequest) (_result *DescribeCharacterSetNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCharacterSetNameResponse{}
	_body, _err := client.DescribeCharacterSetNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the specifications of a cluster.
//
// @param request - DescribeClassListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClassListResponse
func (client *Client) DescribeClassListWithOptions(request *DescribeClassListRequest, runtime *dara.RuntimeOptions) (_result *DescribeClassListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.MasterHa) {
		query["MasterHa"] = request.MasterHa
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClassList"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClassListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the specifications of a cluster.
//
// @param request - DescribeClassListRequest
//
// @return DescribeClassListResponse
func (client *Client) DescribeClassList(request *DescribeClassListRequest) (_result *DescribeClassListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClassListResponse{}
	_body, _err := client.DescribeClassListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询冷存储实例信息
//
// @param request - DescribeColdStorageInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColdStorageInstanceResponse
func (client *Client) DescribeColdStorageInstanceWithOptions(request *DescribeColdStorageInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeColdStorageInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
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

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeColdStorageInstance"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeColdStorageInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询冷存储实例信息
//
// @param request - DescribeColdStorageInstanceRequest
//
// @return DescribeColdStorageInstanceResponse
func (client *Client) DescribeColdStorageInstance(request *DescribeColdStorageInstanceRequest) (_result *DescribeColdStorageInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeColdStorageInstanceResponse{}
	_body, _err := client.DescribeColdStorageInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询消费者组列表
//
// @param request - DescribeConsumerGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConsumerGroupsResponse
func (client *Client) DescribeConsumerGroupsWithOptions(request *DescribeConsumerGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeConsumerGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerGroupId) {
		query["ConsumerGroupId"] = request.ConsumerGroupId
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeConsumerGroups"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeConsumerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询消费者组列表
//
// @param request - DescribeConsumerGroupsRequest
//
// @return DescribeConsumerGroupsResponse
func (client *Client) DescribeConsumerGroups(request *DescribeConsumerGroupsRequest) (_result *DescribeConsumerGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeConsumerGroupsResponse{}
	_body, _err := client.DescribeConsumerGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询消费者列表
//
// @param request - DescribeConsumersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConsumersResponse
func (client *Client) DescribeConsumersWithOptions(request *DescribeConsumersRequest, runtime *dara.RuntimeOptions) (_result *DescribeConsumersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerGroupId) {
		query["ConsumerGroupId"] = request.ConsumerGroupId
	}

	if !dara.IsNil(request.ConsumerId) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeConsumers"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeConsumersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询消费者列表
//
// @param request - DescribeConsumersRequest
//
// @return DescribeConsumersResponse
func (client *Client) DescribeConsumers(request *DescribeConsumersRequest) (_result *DescribeConsumersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeConsumersResponse{}
	_body, _err := client.DescribeConsumersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询限流策略
//
// @param request - DescribeCostRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCostRulesResponse
func (client *Client) DescribeCostRulesWithOptions(request *DescribeCostRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCostRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.ModelServiceId) {
		query["ModelServiceId"] = request.ModelServiceId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCostRules"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCostRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询限流策略
//
// @param request - DescribeCostRulesRequest
//
// @return DescribeCostRulesResponse
func (client *Client) DescribeCostRules(request *DescribeCostRulesRequest) (_result *DescribeCostRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCostRulesResponse{}
	_body, _err := client.DescribeCostRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询周期任务策略
//
// @param request - DescribeCronJobPolicyServerlessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCronJobPolicyServerlessResponse
func (client *Client) DescribeCronJobPolicyServerlessWithOptions(request *DescribeCronJobPolicyServerlessRequest, runtime *dara.RuntimeOptions) (_result *DescribeCronJobPolicyServerlessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCronJobPolicyServerless"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCronJobPolicyServerlessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询周期任务策略
//
// @param request - DescribeCronJobPolicyServerlessRequest
//
// @return DescribeCronJobPolicyServerlessResponse
func (client *Client) DescribeCronJobPolicyServerless(request *DescribeCronJobPolicyServerlessRequest) (_result *DescribeCronJobPolicyServerlessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCronJobPolicyServerlessResponse{}
	_body, _err := client.DescribeCronJobPolicyServerlessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询跨云支持的规格列表
//
// @param request - DescribeCrossCloudLevelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCrossCloudLevelsResponse
func (client *Client) DescribeCrossCloudLevelsWithOptions(request *DescribeCrossCloudLevelsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCrossCloudLevelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBType) {
		query["DBType"] = request.DBType
	}

	if !dara.IsNil(request.DBVersion) {
		query["DBVersion"] = request.DBVersion
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCrossCloudLevels"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCrossCloudLevelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询跨云支持的规格列表
//
// @param request - DescribeCrossCloudLevelsRequest
//
// @return DescribeCrossCloudLevelsResponse
func (client *Client) DescribeCrossCloudLevels(request *DescribeCrossCloudLevelsRequest) (_result *DescribeCrossCloudLevelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCrossCloudLevelsResponse{}
	_body, _err := client.DescribeCrossCloudLevelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看跨云开区信息
//
// @param request - DescribeCrossCloudRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCrossCloudRegionResponse
func (client *Client) DescribeCrossCloudRegionWithOptions(request *DescribeCrossCloudRegionRequest, runtime *dara.RuntimeOptions) (_result *DescribeCrossCloudRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloudProvider) {
		query["CloudProvider"] = request.CloudProvider
	}

	if !dara.IsNil(request.CrossCloudRegionId) {
		query["CrossCloudRegionId"] = request.CrossCloudRegionId
	}

	if !dara.IsNil(request.DBType) {
		query["DBType"] = request.DBType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCrossCloudRegion"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCrossCloudRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看跨云开区信息
//
// @param request - DescribeCrossCloudRegionRequest
//
// @return DescribeCrossCloudRegionResponse
func (client *Client) DescribeCrossCloudRegion(request *DescribeCrossCloudRegionRequest) (_result *DescribeCrossCloudRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCrossCloudRegionResponse{}
	_body, _err := client.DescribeCrossCloudRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询跨云地域映射
//
// @param request - DescribeCrossCloudRegionMappingToAliyunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCrossCloudRegionMappingToAliyunResponse
func (client *Client) DescribeCrossCloudRegionMappingToAliyunWithOptions(request *DescribeCrossCloudRegionMappingToAliyunRequest, runtime *dara.RuntimeOptions) (_result *DescribeCrossCloudRegionMappingToAliyunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunRegionId) {
		query["AliyunRegionId"] = request.AliyunRegionId
	}

	if !dara.IsNil(request.CloudProvider) {
		query["CloudProvider"] = request.CloudProvider
	}

	if !dara.IsNil(request.CrossCloudRegionId) {
		query["CrossCloudRegionId"] = request.CrossCloudRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCrossCloudRegionMappingToAliyun"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCrossCloudRegionMappingToAliyunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询跨云地域映射
//
// @param request - DescribeCrossCloudRegionMappingToAliyunRequest
//
// @return DescribeCrossCloudRegionMappingToAliyunResponse
func (client *Client) DescribeCrossCloudRegionMappingToAliyun(request *DescribeCrossCloudRegionMappingToAliyunRequest) (_result *DescribeCrossCloudRegionMappingToAliyunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCrossCloudRegionMappingToAliyunResponse{}
	_body, _err := client.DescribeCrossCloudRegionMappingToAliyunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IP address whitelists and security groups of a PolarDB cluster.
//
// @param request - DescribeDBClusterAccessWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterAccessWhitelistResponse
func (client *Client) DescribeDBClusterAccessWhitelistWithOptions(request *DescribeDBClusterAccessWhitelistRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterAccessWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterAccessWhitelist"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterAccessWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP address whitelists and security groups of a PolarDB cluster.
//
// @param request - DescribeDBClusterAccessWhitelistRequest
//
// @return DescribeDBClusterAccessWhitelistResponse
func (client *Client) DescribeDBClusterAccessWhitelist(request *DescribeDBClusterAccessWhitelistRequest) (_result *DescribeDBClusterAccessWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterAccessWhitelistResponse{}
	_body, _err := client.DescribeDBClusterAccessWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a PolarDB cluster.
//
// @param request - DescribeDBClusterAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterAttributeResponse
func (client *Client) DescribeDBClusterAttributeWithOptions(request *DescribeDBClusterAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DescribeType) {
		query["DescribeType"] = request.DescribeType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterAttribute"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a PolarDB cluster.
//
// @param request - DescribeDBClusterAttributeRequest
//
// @return DescribeDBClusterAttributeResponse
func (client *Client) DescribeDBClusterAttribute(request *DescribeDBClusterAttributeRequest) (_result *DescribeDBClusterAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterAttributeResponse{}
	_body, _err := client.DescribeDBClusterAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Describe SQL collector for a PolarDB cluster. Features related to SQL collector include audit log and SQL Explorer.
//
// @param request - DescribeDBClusterAuditLogCollectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterAuditLogCollectorResponse
func (client *Client) DescribeDBClusterAuditLogCollectorWithOptions(request *DescribeDBClusterAuditLogCollectorRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterAuditLogCollectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterAuditLogCollector"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterAuditLogCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describe SQL collector for a PolarDB cluster. Features related to SQL collector include audit log and SQL Explorer.
//
// @param request - DescribeDBClusterAuditLogCollectorRequest
//
// @return DescribeDBClusterAuditLogCollectorResponse
func (client *Client) DescribeDBClusterAuditLogCollector(request *DescribeDBClusterAuditLogCollectorRequest) (_result *DescribeDBClusterAuditLogCollectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterAuditLogCollectorResponse{}
	_body, _err := client.DescribeDBClusterAuditLogCollectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries available resources in a PolarDB cluster.
//
// @param request - DescribeDBClusterAvailableResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterAvailableResourcesResponse
func (client *Client) DescribeDBClusterAvailableResourcesWithOptions(request *DescribeDBClusterAvailableResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterAvailableResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBNodeClass) {
		query["DBNodeClass"] = request.DBNodeClass
	}

	if !dara.IsNil(request.DBType) {
		query["DBType"] = request.DBType
	}

	if !dara.IsNil(request.DBVersion) {
		query["DBVersion"] = request.DBVersion
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterAvailableResources"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterAvailableResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available resources in a PolarDB cluster.
//
// @param request - DescribeDBClusterAvailableResourcesRequest
//
// @return DescribeDBClusterAvailableResourcesResponse
func (client *Client) DescribeDBClusterAvailableResources(request *DescribeDBClusterAvailableResourcesRequest) (_result *DescribeDBClusterAvailableResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterAvailableResourcesResponse{}
	_body, _err := client.DescribeDBClusterAvailableResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether the source IP address can access a cluster.
//
// @param request - DescribeDBClusterConnectivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterConnectivityResponse
func (client *Client) DescribeDBClusterConnectivityWithOptions(request *DescribeDBClusterConnectivityRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterConnectivityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SourceIpAddress) {
		query["SourceIpAddress"] = request.SourceIpAddress
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterConnectivity"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterConnectivityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the source IP address can access a cluster.
//
// @param request - DescribeDBClusterConnectivityRequest
//
// @return DescribeDBClusterConnectivityResponse
func (client *Client) DescribeDBClusterConnectivity(request *DescribeDBClusterConnectivityRequest) (_result *DescribeDBClusterConnectivityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterConnectivityResponse{}
	_body, _err := client.DescribeDBClusterConnectivityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DescribeDBClusterEncryptionKey
//
// @param request - DescribeDBClusterEncryptionKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterEncryptionKeyResponse
func (client *Client) DescribeDBClusterEncryptionKeyWithOptions(request *DescribeDBClusterEncryptionKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterEncryptionKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterEncryptionKey"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterEncryptionKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeDBClusterEncryptionKey
//
// @param request - DescribeDBClusterEncryptionKeyRequest
//
// @return DescribeDBClusterEncryptionKeyResponse
func (client *Client) DescribeDBClusterEncryptionKey(request *DescribeDBClusterEncryptionKeyRequest) (_result *DescribeDBClusterEncryptionKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterEncryptionKeyResponse{}
	_body, _err := client.DescribeDBClusterEncryptionKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the endpoints of a PolarDB cluster.
//
// @param request - DescribeDBClusterEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterEndpointsResponse
func (client *Client) DescribeDBClusterEndpointsWithOptions(request *DescribeDBClusterEndpointsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterEndpointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBEndpointId) {
		query["DBEndpointId"] = request.DBEndpointId
	}

	if !dara.IsNil(request.DescribeType) {
		query["DescribeType"] = request.DescribeType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PolarFsInstanceId) {
		query["PolarFsInstanceId"] = request.PolarFsInstanceId
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
		Action:      dara.String("DescribeDBClusterEndpoints"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the endpoints of a PolarDB cluster.
//
// @param request - DescribeDBClusterEndpointsRequest
//
// @return DescribeDBClusterEndpointsResponse
func (client *Client) DescribeDBClusterEndpoints(request *DescribeDBClusterEndpointsRequest) (_result *DescribeDBClusterEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterEndpointsResponse{}
	_body, _err := client.DescribeDBClusterEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询PolarDB边缘集群的链接地址
//
// @param request - DescribeDBClusterEndpointsZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterEndpointsZonalResponse
func (client *Client) DescribeDBClusterEndpointsZonalWithOptions(request *DescribeDBClusterEndpointsZonalRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterEndpointsZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBEndpointId) {
		query["DBEndpointId"] = request.DBEndpointId
	}

	if !dara.IsNil(request.DescribeType) {
		query["DescribeType"] = request.DescribeType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterEndpointsZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterEndpointsZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询PolarDB边缘集群的链接地址
//
// @param request - DescribeDBClusterEndpointsZonalRequest
//
// @return DescribeDBClusterEndpointsZonalResponse
func (client *Client) DescribeDBClusterEndpointsZonal(request *DescribeDBClusterEndpointsZonalRequest) (_result *DescribeDBClusterEndpointsZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterEndpointsZonalResponse{}
	_body, _err := client.DescribeDBClusterEndpointsZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The ID of the synchronous task.
//
// Description:
//
// The ID of the request.
//
// @param request - DescribeDBClusterMigrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterMigrationResponse
func (client *Client) DescribeDBClusterMigrationWithOptions(request *DescribeDBClusterMigrationRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterMigrationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterMigration"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterMigrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the synchronous task.
//
// Description:
//
// The ID of the request.
//
// @param request - DescribeDBClusterMigrationRequest
//
// @return DescribeDBClusterMigrationResponse
func (client *Client) DescribeDBClusterMigration(request *DescribeDBClusterMigrationRequest) (_result *DescribeDBClusterMigrationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterMigrationResponse{}
	_body, _err := client.DescribeDBClusterMigrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the interval at which the monitoring data of a PolarDB cluster is collected.
//
// @param request - DescribeDBClusterMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterMonitorResponse
func (client *Client) DescribeDBClusterMonitorWithOptions(request *DescribeDBClusterMonitorRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterMonitor"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the interval at which the monitoring data of a PolarDB cluster is collected.
//
// @param request - DescribeDBClusterMonitorRequest
//
// @return DescribeDBClusterMonitorResponse
func (client *Client) DescribeDBClusterMonitor(request *DescribeDBClusterMonitorRequest) (_result *DescribeDBClusterMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterMonitorResponse{}
	_body, _err := client.DescribeDBClusterMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDBClusterNetInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterNetInfoResponse
func (client *Client) DescribeDBClusterNetInfoWithOptions(request *DescribeDBClusterNetInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterNetInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionStringType) {
		query["ConnectionStringType"] = request.ConnectionStringType
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterNetInfo"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterNetInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDBClusterNetInfoRequest
//
// @return DescribeDBClusterNetInfoResponse
func (client *Client) DescribeDBClusterNetInfo(request *DescribeDBClusterNetInfoRequest) (_result *DescribeDBClusterNetInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterNetInfoResponse{}
	_body, _err := client.DescribeDBClusterNetInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the parameters of a PolarDB cluster.
//
// @param request - DescribeDBClusterParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterParametersResponse
func (client *Client) DescribeDBClusterParametersWithOptions(request *DescribeDBClusterParametersRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterParametersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DescribeType) {
		query["DescribeType"] = request.DescribeType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterParameters"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the parameters of a PolarDB cluster.
//
// @param request - DescribeDBClusterParametersRequest
//
// @return DescribeDBClusterParametersResponse
func (client *Client) DescribeDBClusterParameters(request *DescribeDBClusterParametersRequest) (_result *DescribeDBClusterParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterParametersResponse{}
	_body, _err := client.DescribeDBClusterParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the performance data of a PolarDB cluster.
//
// Description:
//
//	  When the monitoring data is collected every 5 seconds:
//
//	    	- If the query time range is less than or equal to 1 hour, the data is displayed at intervals of 5 seconds.
//
//	    	- If the query time range is less than or equal to one day, the data is displayed at intervals of 1 minute.
//
//	    	- If the query time range is less than or equal to seven days, the data is displayed at intervals of 10 minutes.
//
//	    	- If the query time range is less than or equal to 30 days, the data is displayed at intervals of 1 hour.
//
//	    	- When the query time range is greater than 30 days, the data is displayed at intervals of 1 day.
//
//		- When the monitoring data is collected every 60 seconds:
//
//	    	- If the query time range is less than or equal to one day, the data is displayed at intervals of 1 minute.
//
//	    	- If the query time range is less than or equal to seven days, the data is displayed at intervals of 10 minutes.
//
//	    	- If the query time range is less than or equal to 30 days, the data is displayed at intervals of 1 hour.
//
//	    	- When the query time range is greater than 30 days, the data is displayed at intervals of 1 day.
//
// >  By default, the monitoring data is collected once every 60 seconds. You can call the [ModifyDBClusterMonitor](https://help.aliyun.com/document_detail/159557.html) operation to set the data collection interval to every 5 seconds.
//
// @param request - DescribeDBClusterPerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterPerformanceResponse
func (client *Client) DescribeDBClusterPerformanceWithOptions(request *DescribeDBClusterPerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.SubGroupName) {
		query["SubGroupName"] = request.SubGroupName
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterPerformance"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the performance data of a PolarDB cluster.
//
// Description:
//
//	  When the monitoring data is collected every 5 seconds:
//
//	    	- If the query time range is less than or equal to 1 hour, the data is displayed at intervals of 5 seconds.
//
//	    	- If the query time range is less than or equal to one day, the data is displayed at intervals of 1 minute.
//
//	    	- If the query time range is less than or equal to seven days, the data is displayed at intervals of 10 minutes.
//
//	    	- If the query time range is less than or equal to 30 days, the data is displayed at intervals of 1 hour.
//
//	    	- When the query time range is greater than 30 days, the data is displayed at intervals of 1 day.
//
//		- When the monitoring data is collected every 60 seconds:
//
//	    	- If the query time range is less than or equal to one day, the data is displayed at intervals of 1 minute.
//
//	    	- If the query time range is less than or equal to seven days, the data is displayed at intervals of 10 minutes.
//
//	    	- If the query time range is less than or equal to 30 days, the data is displayed at intervals of 1 hour.
//
//	    	- When the query time range is greater than 30 days, the data is displayed at intervals of 1 day.
//
// >  By default, the monitoring data is collected once every 60 seconds. You can call the [ModifyDBClusterMonitor](https://help.aliyun.com/document_detail/159557.html) operation to set the data collection interval to every 5 seconds.
//
// @param request - DescribeDBClusterPerformanceRequest
//
// @return DescribeDBClusterPerformanceResponse
func (client *Client) DescribeDBClusterPerformance(request *DescribeDBClusterPerformanceRequest) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.DescribeDBClusterPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询代理详情
//
// @param request - DescribeDBClusterProxyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterProxyResponse
func (client *Client) DescribeDBClusterProxyWithOptions(request *DescribeDBClusterProxyRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterProxyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterProxy"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询代理详情
//
// @param request - DescribeDBClusterProxyRequest
//
// @return DescribeDBClusterProxyResponse
func (client *Client) DescribeDBClusterProxy(request *DescribeDBClusterProxyRequest) (_result *DescribeDBClusterProxyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterProxyResponse{}
	_body, _err := client.DescribeDBClusterProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Secure Sockets Layer (SSL) settings of a PolarDB cluster.
//
// @param request - DescribeDBClusterSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterSSLResponse
func (client *Client) DescribeDBClusterSSLWithOptions(request *DescribeDBClusterSSLRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterSSLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterSSL"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Secure Sockets Layer (SSL) settings of a PolarDB cluster.
//
// @param request - DescribeDBClusterSSLRequest
//
// @return DescribeDBClusterSSLResponse
func (client *Client) DescribeDBClusterSSL(request *DescribeDBClusterSSLRequest) (_result *DescribeDBClusterSSLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterSSLResponse{}
	_body, _err := client.DescribeDBClusterSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the serverless configurations of a serverless cluster.
//
// @param request - DescribeDBClusterServerlessConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterServerlessConfResponse
func (client *Client) DescribeDBClusterServerlessConfWithOptions(request *DescribeDBClusterServerlessConfRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterServerlessConfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterServerlessConf"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterServerlessConfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the serverless configurations of a serverless cluster.
//
// @param request - DescribeDBClusterServerlessConfRequest
//
// @return DescribeDBClusterServerlessConfResponse
func (client *Client) DescribeDBClusterServerlessConf(request *DescribeDBClusterServerlessConfRequest) (_result *DescribeDBClusterServerlessConfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterServerlessConfResponse{}
	_body, _err := client.DescribeDBClusterServerlessConfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the transparent data encryption (TDE) settings of a PolarDB cluster.
//
// @param request - DescribeDBClusterTDERequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterTDEResponse
func (client *Client) DescribeDBClusterTDEWithOptions(request *DescribeDBClusterTDERequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterTDEResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterTDE"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterTDEResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the transparent data encryption (TDE) settings of a PolarDB cluster.
//
// @param request - DescribeDBClusterTDERequest
//
// @return DescribeDBClusterTDEResponse
func (client *Client) DescribeDBClusterTDE(request *DescribeDBClusterTDERequest) (_result *DescribeDBClusterTDEResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterTDEResponse{}
	_body, _err := client.DescribeDBClusterTDEWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the database engine version of a PolarDB for MySQL cluster.
//
// @param request - DescribeDBClusterVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterVersionResponse
func (client *Client) DescribeDBClusterVersionWithOptions(request *DescribeDBClusterVersionRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DescribeType) {
		query["DescribeType"] = request.DescribeType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterVersion"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the database engine version of a PolarDB for MySQL cluster.
//
// @param request - DescribeDBClusterVersionRequest
//
// @return DescribeDBClusterVersionResponse
func (client *Client) DescribeDBClusterVersion(request *DescribeDBClusterVersionRequest) (_result *DescribeDBClusterVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterVersionResponse{}
	_body, _err := client.DescribeDBClusterVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # PolarDB边缘集群查询版本
//
// @param request - DescribeDBClusterVersionZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterVersionZonalResponse
func (client *Client) DescribeDBClusterVersionZonalWithOptions(request *DescribeDBClusterVersionZonalRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterVersionZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DescribeType) {
		query["DescribeType"] = request.DescribeType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterVersionZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterVersionZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # PolarDB边缘集群查询版本
//
// @param request - DescribeDBClusterVersionZonalRequest
//
// @return DescribeDBClusterVersionZonalResponse
func (client *Client) DescribeDBClusterVersionZonal(request *DescribeDBClusterVersionZonalRequest) (_result *DescribeDBClusterVersionZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterVersionZonalResponse{}
	_body, _err := client.DescribeDBClusterVersionZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries PolarDB clusters or the clusters that can be accessed by an authorized RAM user.
//
// @param request - DescribeDBClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClustersResponse
func (client *Client) DescribeDBClustersWithOptions(request *DescribeDBClustersRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.DBClusterDescription) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !dara.IsNil(request.DBClusterIds) {
		query["DBClusterIds"] = request.DBClusterIds
	}

	if !dara.IsNil(request.DBClusterStatus) {
		query["DBClusterStatus"] = request.DBClusterStatus
	}

	if !dara.IsNil(request.DBNodeIds) {
		query["DBNodeIds"] = request.DBNodeIds
	}

	if !dara.IsNil(request.DBType) {
		query["DBType"] = request.DBType
	}

	if !dara.IsNil(request.DBVersion) {
		query["DBVersion"] = request.DBVersion
	}

	if !dara.IsNil(request.DescribeType) {
		query["DescribeType"] = request.DescribeType
	}

	if !dara.IsNil(request.Expired) {
		query["Expired"] = request.Expired
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

	if !dara.IsNil(request.RecentCreationInterval) {
		query["RecentCreationInterval"] = request.RecentCreationInterval
	}

	if !dara.IsNil(request.RecentExpirationInterval) {
		query["RecentExpirationInterval"] = request.RecentExpirationInterval
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
		Action:      dara.String("DescribeDBClusters"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries PolarDB clusters or the clusters that can be accessed by an authorized RAM user.
//
// @param request - DescribeDBClustersRequest
//
// @return DescribeDBClustersResponse
func (client *Client) DescribeDBClusters(request *DescribeDBClustersRequest) (_result *DescribeDBClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClustersResponse{}
	_body, _err := client.DescribeDBClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about PolarDB clusters that contain backup sets in a region.
//
// @param request - DescribeDBClustersWithBackupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClustersWithBackupsResponse
func (client *Client) DescribeDBClustersWithBackupsWithOptions(request *DescribeDBClustersWithBackupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClustersWithBackupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterDescription) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !dara.IsNil(request.DBClusterIds) {
		query["DBClusterIds"] = request.DBClusterIds
	}

	if !dara.IsNil(request.DBType) {
		query["DBType"] = request.DBType
	}

	if !dara.IsNil(request.DBVersion) {
		query["DBVersion"] = request.DBVersion
	}

	if !dara.IsNil(request.IsDeleted) {
		query["IsDeleted"] = request.IsDeleted
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClustersWithBackups"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClustersWithBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about PolarDB clusters that contain backup sets in a region.
//
// @param request - DescribeDBClustersWithBackupsRequest
//
// @return DescribeDBClustersWithBackupsResponse
func (client *Client) DescribeDBClustersWithBackups(request *DescribeDBClustersWithBackupsRequest) (_result *DescribeDBClustersWithBackupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClustersWithBackupsResponse{}
	_body, _err := client.DescribeDBClustersWithBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # MyBase中的PolarDB列表
//
// @param request - DescribeDBClustersZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClustersZonalResponse
func (client *Client) DescribeDBClustersZonalWithOptions(request *DescribeDBClustersZonalRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClustersZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloudProvider) {
		query["CloudProvider"] = request.CloudProvider
	}

	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.DBClusterDescription) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !dara.IsNil(request.DBClusterIds) {
		query["DBClusterIds"] = request.DBClusterIds
	}

	if !dara.IsNil(request.DBClusterStatus) {
		query["DBClusterStatus"] = request.DBClusterStatus
	}

	if !dara.IsNil(request.DBNodeIds) {
		query["DBNodeIds"] = request.DBNodeIds
	}

	if !dara.IsNil(request.DBType) {
		query["DBType"] = request.DBType
	}

	if !dara.IsNil(request.DBVersion) {
		query["DBVersion"] = request.DBVersion
	}

	if !dara.IsNil(request.DescribeType) {
		query["DescribeType"] = request.DescribeType
	}

	if !dara.IsNil(request.Expired) {
		query["Expired"] = request.Expired
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.RecentCreationInterval) {
		query["RecentCreationInterval"] = request.RecentCreationInterval
	}

	if !dara.IsNil(request.RecentExpirationInterval) {
		query["RecentExpirationInterval"] = request.RecentExpirationInterval
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
		Action:      dara.String("DescribeDBClustersZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClustersZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # MyBase中的PolarDB列表
//
// @param request - DescribeDBClustersZonalRequest
//
// @return DescribeDBClustersZonalResponse
func (client *Client) DescribeDBClustersZonal(request *DescribeDBClustersZonalRequest) (_result *DescribeDBClustersZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClustersZonalResponse{}
	_body, _err := client.DescribeDBClustersZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries attributes such as character sets and collations supported by a database in a PolarDB cluster.
//
// @param request - DescribeDBInitializeVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInitializeVariableResponse
func (client *Client) DescribeDBInitializeVariableWithOptions(request *DescribeDBInitializeVariableRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInitializeVariableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInitializeVariable"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInitializeVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries attributes such as character sets and collations supported by a database in a PolarDB cluster.
//
// @param request - DescribeDBInitializeVariableRequest
//
// @return DescribeDBInitializeVariableResponse
func (client *Client) DescribeDBInitializeVariable(request *DescribeDBInitializeVariableRequest) (_result *DescribeDBInitializeVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInitializeVariableResponse{}
	_body, _err := client.DescribeDBInitializeVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDBInstancePerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstancePerformanceResponse
func (client *Client) DescribeDBInstancePerformanceWithOptions(request *DescribeDBInstancePerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstancePerformanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstancePerformance"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstancePerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDBInstancePerformanceRequest
//
// @return DescribeDBInstancePerformanceResponse
func (client *Client) DescribeDBInstancePerformance(request *DescribeDBInstancePerformanceRequest) (_result *DescribeDBInstancePerformanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstancePerformanceResponse{}
	_body, _err := client.DescribeDBInstancePerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the database links of a PolarDB for Oracle cluster.
//
// Description:
//
// > You can query only the database links that use a PolarDB for Oracle cluster as the source.
//
// @param request - DescribeDBLinksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBLinksResponse
func (client *Client) DescribeDBLinksWithOptions(request *DescribeDBLinksRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBLinksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBLinkName) {
		query["DBLinkName"] = request.DBLinkName
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBLinks"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBLinksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the database links of a PolarDB for Oracle cluster.
//
// Description:
//
// > You can query only the database links that use a PolarDB for Oracle cluster as the source.
//
// @param request - DescribeDBLinksRequest
//
// @return DescribeDBLinksResponse
func (client *Client) DescribeDBLinks(request *DescribeDBLinksRequest) (_result *DescribeDBLinksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBLinksResponse{}
	_body, _err := client.DescribeDBLinksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries logs for a PolarDB cluster such as primary/secondary switchover logs.
//
// @param request - DescribeDBLogFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBLogFilesResponse
func (client *Client) DescribeDBLogFilesWithOptions(request *DescribeDBLogFilesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBLogFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNodeId) {
		query["DBNodeId"] = request.DBNodeId
	}

	if !dara.IsNil(request.DescribeSimulateSwitchMode) {
		query["DescribeSimulateSwitchMode"] = request.DescribeSimulateSwitchMode
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SimulateListId) {
		query["SimulateListId"] = request.SimulateListId
	}

	if !dara.IsNil(request.SimulateModeList) {
		query["SimulateModeList"] = request.SimulateModeList
	}

	if !dara.IsNil(request.SimulateStatusList) {
		query["SimulateStatusList"] = request.SimulateStatusList
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBLogFiles"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBLogFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries logs for a PolarDB cluster such as primary/secondary switchover logs.
//
// @param request - DescribeDBLogFilesRequest
//
// @return DescribeDBLogFilesResponse
func (client *Client) DescribeDBLogFiles(request *DescribeDBLogFilesRequest) (_result *DescribeDBLogFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBLogFilesResponse{}
	_body, _err := client.DescribeDBLogFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询可用内核小版本列表
//
// @param request - DescribeDBMiniEngineVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBMiniEngineVersionsResponse
func (client *Client) DescribeDBMiniEngineVersionsWithOptions(request *DescribeDBMiniEngineVersionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBMiniEngineVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Architecture) {
		query["Architecture"] = request.Architecture
	}

	if !dara.IsNil(request.CreationCategory) {
		query["CreationCategory"] = request.CreationCategory
	}

	if !dara.IsNil(request.DBMinorVersion) {
		query["DBMinorVersion"] = request.DBMinorVersion
	}

	if !dara.IsNil(request.DBType) {
		query["DBType"] = request.DBType
	}

	if !dara.IsNil(request.DBVersion) {
		query["DBVersion"] = request.DBVersion
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
		Action:      dara.String("DescribeDBMiniEngineVersions"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBMiniEngineVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询可用内核小版本列表
//
// @param request - DescribeDBMiniEngineVersionsRequest
//
// @return DescribeDBMiniEngineVersionsResponse
func (client *Client) DescribeDBMiniEngineVersions(request *DescribeDBMiniEngineVersionsRequest) (_result *DescribeDBMiniEngineVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBMiniEngineVersionsResponse{}
	_body, _err := client.DescribeDBMiniEngineVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the performance data of a node in a PolarDB cluster.
//
// Description:
//
//	  When the monitoring data is collected every 5 seconds:
//
//	    	- If the query time range is less than or equal to 1 hour, the data is displayed at intervals of 5 seconds.
//
//	    	- If the query time range is less than or equal to one day, the data is displayed at intervals of 1 minute.
//
//	    	- If the query time range is less than or equal to seven days, the data is displayed at intervals of 10 minutes.
//
//	    	- If the query time range is less than or equal to 30 days, the data is displayed at intervals of 1 hour.
//
//	    	- When the query time range is greater than 30 days, the data is displayed at intervals of 1 day.
//
//		- When the monitoring data is collected every 60 seconds:
//
//	    	- If the query time range is less than or equal to one day, the data is displayed at intervals of 1 minute.
//
//	    	- If the query time range is less than or equal to seven days, the data is displayed at intervals of 10 minutes.
//
//	    	- If the query time range is less than or equal to 30 days, the data is displayed at intervals of 1 hour.
//
//	    	- When the query time range is greater than 30 days, the data is displayed at intervals of 1 day.
//
// >  By default, the monitoring data is collected once every 60 seconds. You can call the [ModifyDBClusterMonitor](https://help.aliyun.com/document_detail/159557.html) operation to set the data collection interval to every 5 seconds.
//
// @param request - DescribeDBNodePerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBNodePerformanceResponse
func (client *Client) DescribeDBNodePerformanceWithOptions(request *DescribeDBNodePerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBNodePerformanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNodeId) {
		query["DBNodeId"] = request.DBNodeId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
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
		Action:      dara.String("DescribeDBNodePerformance"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBNodePerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the performance data of a node in a PolarDB cluster.
//
// Description:
//
//	  When the monitoring data is collected every 5 seconds:
//
//	    	- If the query time range is less than or equal to 1 hour, the data is displayed at intervals of 5 seconds.
//
//	    	- If the query time range is less than or equal to one day, the data is displayed at intervals of 1 minute.
//
//	    	- If the query time range is less than or equal to seven days, the data is displayed at intervals of 10 minutes.
//
//	    	- If the query time range is less than or equal to 30 days, the data is displayed at intervals of 1 hour.
//
//	    	- When the query time range is greater than 30 days, the data is displayed at intervals of 1 day.
//
//		- When the monitoring data is collected every 60 seconds:
//
//	    	- If the query time range is less than or equal to one day, the data is displayed at intervals of 1 minute.
//
//	    	- If the query time range is less than or equal to seven days, the data is displayed at intervals of 10 minutes.
//
//	    	- If the query time range is less than or equal to 30 days, the data is displayed at intervals of 1 hour.
//
//	    	- When the query time range is greater than 30 days, the data is displayed at intervals of 1 day.
//
// >  By default, the monitoring data is collected once every 60 seconds. You can call the [ModifyDBClusterMonitor](https://help.aliyun.com/document_detail/159557.html) operation to set the data collection interval to every 5 seconds.
//
// @param request - DescribeDBNodePerformanceRequest
//
// @return DescribeDBNodePerformanceResponse
func (client *Client) DescribeDBNodePerformance(request *DescribeDBNodePerformanceRequest) (_result *DescribeDBNodePerformanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBNodePerformanceResponse{}
	_body, _err := client.DescribeDBNodePerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the parameters of a specified node in a cluster.
//
// @param request - DescribeDBNodesParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBNodesParametersResponse
func (client *Client) DescribeDBNodesParametersWithOptions(request *DescribeDBNodesParametersRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBNodesParametersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNodeIds) {
		query["DBNodeIds"] = request.DBNodeIds
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBNodesParameters"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBNodesParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the parameters of a specified node in a cluster.
//
// @param request - DescribeDBNodesParametersRequest
//
// @return DescribeDBNodesParametersResponse
func (client *Client) DescribeDBNodesParameters(request *DescribeDBNodesParametersRequest) (_result *DescribeDBNodesParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBNodesParametersResponse{}
	_body, _err := client.DescribeDBNodesParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the performance data of PolarProxy.
//
// Description:
//
// > This operation is applicable only to PolarDB for MySQL clusters.
//
// @param request - DescribeDBProxyPerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBProxyPerformanceResponse
func (client *Client) DescribeDBProxyPerformanceWithOptions(request *DescribeDBProxyPerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBProxyPerformanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBEndpointId) {
		query["DBEndpointId"] = request.DBEndpointId
	}

	if !dara.IsNil(request.DBNodeId) {
		query["DBNodeId"] = request.DBNodeId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
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
		Action:      dara.String("DescribeDBProxyPerformance"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBProxyPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the performance data of PolarProxy.
//
// Description:
//
// > This operation is applicable only to PolarDB for MySQL clusters.
//
// @param request - DescribeDBProxyPerformanceRequest
//
// @return DescribeDBProxyPerformanceResponse
func (client *Client) DescribeDBProxyPerformance(request *DescribeDBProxyPerformanceRequest) (_result *DescribeDBProxyPerformanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBProxyPerformanceResponse{}
	_body, _err := client.DescribeDBProxyPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of a cluster in Database Autonomy Service (DAS).
//
// @param request - DescribeDasConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDasConfigResponse
func (client *Client) DescribeDasConfigWithOptions(request *DescribeDasConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDasConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDasConfig"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDasConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a cluster in Database Autonomy Service (DAS).
//
// @param request - DescribeDasConfigRequest
//
// @return DescribeDasConfigResponse
func (client *Client) DescribeDasConfig(request *DescribeDasConfigRequest) (_result *DescribeDasConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDasConfigResponse{}
	_body, _err := client.DescribeDasConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about databases in a PolarDB cluster.
//
// @param request - DescribeDatabasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDatabasesResponse
func (client *Client) DescribeDatabasesWithOptions(request *DescribeDatabasesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDatabasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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
		Action:      dara.String("DescribeDatabases"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about databases in a PolarDB cluster.
//
// @param request - DescribeDatabasesRequest
//
// @return DescribeDatabasesResponse
func (client *Client) DescribeDatabases(request *DescribeDatabasesRequest) (_result *DescribeDatabasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDatabasesResponse{}
	_body, _err := client.DescribeDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询PolarDB边缘云数据库
//
// @param request - DescribeDatabasesZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDatabasesZonalResponse
func (client *Client) DescribeDatabasesZonalWithOptions(request *DescribeDatabasesZonalRequest, runtime *dara.RuntimeOptions) (_result *DescribeDatabasesZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDatabasesZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDatabasesZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询PolarDB边缘云数据库
//
// @param request - DescribeDatabasesZonalRequest
//
// @return DescribeDatabasesZonalResponse
func (client *Client) DescribeDatabasesZonal(request *DescribeDatabasesZonalRequest) (_result *DescribeDatabasesZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDatabasesZonalResponse{}
	_body, _err := client.DescribeDatabasesZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # PolarDB的MyBase集群
//
// @param request - DescribeDbClusterAttributeZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDbClusterAttributeZonalResponse
func (client *Client) DescribeDbClusterAttributeZonalWithOptions(request *DescribeDbClusterAttributeZonalRequest, runtime *dara.RuntimeOptions) (_result *DescribeDbClusterAttributeZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DescribeType) {
		query["DescribeType"] = request.DescribeType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDbClusterAttributeZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDbClusterAttributeZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # PolarDB的MyBase集群
//
// @param request - DescribeDbClusterAttributeZonalRequest
//
// @return DescribeDbClusterAttributeZonalResponse
func (client *Client) DescribeDbClusterAttributeZonal(request *DescribeDbClusterAttributeZonalRequest) (_result *DescribeDbClusterAttributeZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDbClusterAttributeZonalResponse{}
	_body, _err := client.DescribeDbClusterAttributeZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the backup sets in a released PolarDB cluster.
//
// Description:
//
// Before you call this operation, make sure that the PolarDB cluster is in the **Released*	- state. You must also confirm that the **Retain All Backups Permanently*	- or **Retain Last Automatic Backup Permanently*	- backup retention policy takes effect after you release the cluster. If you delete all backup sets after the cluster is released, you cannot use this API operation to query the cluster.
//
// >  You can call the [DescribeDBClusterAttribute](https://help.aliyun.com/document_detail/98181.html) operation to query the cluster status.
//
// @param request - DescribeDetachedBackupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDetachedBackupsResponse
func (client *Client) DescribeDetachedBackupsWithOptions(request *DescribeDetachedBackupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDetachedBackupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.BackupMode) {
		query["BackupMode"] = request.BackupMode
	}

	if !dara.IsNil(request.BackupRegion) {
		query["BackupRegion"] = request.BackupRegion
	}

	if !dara.IsNil(request.BackupStatus) {
		query["BackupStatus"] = request.BackupStatus
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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
		Action:      dara.String("DescribeDetachedBackups"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDetachedBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the backup sets in a released PolarDB cluster.
//
// Description:
//
// Before you call this operation, make sure that the PolarDB cluster is in the **Released*	- state. You must also confirm that the **Retain All Backups Permanently*	- or **Retain Last Automatic Backup Permanently*	- backup retention policy takes effect after you release the cluster. If you delete all backup sets after the cluster is released, you cannot use this API operation to query the cluster.
//
// >  You can call the [DescribeDBClusterAttribute](https://help.aliyun.com/document_detail/98181.html) operation to query the cluster status.
//
// @param request - DescribeDetachedBackupsRequest
//
// @return DescribeDetachedBackupsResponse
func (client *Client) DescribeDetachedBackups(request *DescribeDetachedBackupsRequest) (_result *DescribeDetachedBackupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDetachedBackupsResponse{}
	_body, _err := client.DescribeDetachedBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取角色权限列表
//
// @param request - DescribeEncryptionDBRolePrivilegeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEncryptionDBRolePrivilegeResponse
func (client *Client) DescribeEncryptionDBRolePrivilegeWithOptions(request *DescribeEncryptionDBRolePrivilegeRequest, runtime *dara.RuntimeOptions) (_result *DescribeEncryptionDBRolePrivilegeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.RolePrivilegeNameList) {
		query["RolePrivilegeNameList"] = request.RolePrivilegeNameList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEncryptionDBRolePrivilege"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEncryptionDBRolePrivilegeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取角色权限列表
//
// @param request - DescribeEncryptionDBRolePrivilegeRequest
//
// @return DescribeEncryptionDBRolePrivilegeResponse
func (client *Client) DescribeEncryptionDBRolePrivilege(request *DescribeEncryptionDBRolePrivilegeRequest) (_result *DescribeEncryptionDBRolePrivilegeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEncryptionDBRolePrivilegeResponse{}
	_body, _err := client.DescribeEncryptionDBRolePrivilegeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取加密信息
//
// @param request - DescribeEncryptionDBSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEncryptionDBSecretResponse
func (client *Client) DescribeEncryptionDBSecretWithOptions(request *DescribeEncryptionDBSecretRequest, runtime *dara.RuntimeOptions) (_result *DescribeEncryptionDBSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEncryptionDBSecret"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEncryptionDBSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取加密信息
//
// @param request - DescribeEncryptionDBSecretRequest
//
// @return DescribeEncryptionDBSecretResponse
func (client *Client) DescribeEncryptionDBSecret(request *DescribeEncryptionDBSecretRequest) (_result *DescribeEncryptionDBSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEncryptionDBSecretResponse{}
	_body, _err := client.DescribeEncryptionDBSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询插件详情
//
// @param request - DescribeExtensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExtensionsResponse
func (client *Client) DescribeExtensionsWithOptions(request *DescribeExtensionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeExtensionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.ExtensionName) {
		query["ExtensionName"] = request.ExtensionName
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExtensions"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExtensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询插件详情
//
// @param request - DescribeExtensionsRequest
//
// @return DescribeExtensionsResponse
func (client *Client) DescribeExtensions(request *DescribeExtensionsRequest) (_result *DescribeExtensionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExtensionsResponse{}
	_body, _err := client.DescribeExtensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询sql防火墙信息
//
// @param request - DescribeFirewallRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFirewallRulesResponse
func (client *Client) DescribeFirewallRulesWithOptions(request *DescribeFirewallRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeFirewallRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.RuleNameList) {
		query["RuleNameList"] = request.RuleNameList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFirewallRules"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFirewallRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询sql防火墙信息
//
// @param request - DescribeFirewallRulesRequest
//
// @return DescribeFirewallRulesResponse
func (client *Client) DescribeFirewallRules(request *DescribeFirewallRulesRequest) (_result *DescribeFirewallRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFirewallRulesResponse{}
	_body, _err := client.DescribeFirewallRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询网关实例详情
//
// @param request - DescribeGatewayAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGatewayAttributeResponse
func (client *Client) DescribeGatewayAttributeWithOptions(request *DescribeGatewayAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeGatewayAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGatewayAttribute"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGatewayAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询网关实例详情
//
// @param request - DescribeGatewayAttributeRequest
//
// @return DescribeGatewayAttributeResponse
func (client *Client) DescribeGatewayAttribute(request *DescribeGatewayAttributeRequest) (_result *DescribeGatewayAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGatewayAttributeResponse{}
	_body, _err := client.DescribeGatewayAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询网关实例列表
//
// @param request - DescribeGatewayListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGatewayListResponse
func (client *Client) DescribeGatewayListWithOptions(request *DescribeGatewayListRequest, runtime *dara.RuntimeOptions) (_result *DescribeGatewayListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.GwDescription) {
		query["GwDescription"] = request.GwDescription
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGatewayList"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGatewayListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询网关实例列表
//
// @param request - DescribeGatewayListRequest
//
// @return DescribeGatewayListResponse
func (client *Client) DescribeGatewayList(request *DescribeGatewayListRequest) (_result *DescribeGatewayListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGatewayListResponse{}
	_body, _err := client.DescribeGatewayListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the PolarFS global data network (GDN) details in all regions.
//
// @param request - DescribeGlobalDataNetworkListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGlobalDataNetworkListResponse
func (client *Client) DescribeGlobalDataNetworkListWithOptions(request *DescribeGlobalDataNetworkListRequest, runtime *dara.RuntimeOptions) (_result *DescribeGlobalDataNetworkListResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGlobalDataNetworkList"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGlobalDataNetworkListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the PolarFS global data network (GDN) details in all regions.
//
// @param request - DescribeGlobalDataNetworkListRequest
//
// @return DescribeGlobalDataNetworkListResponse
func (client *Client) DescribeGlobalDataNetworkList(request *DescribeGlobalDataNetworkListRequest) (_result *DescribeGlobalDataNetworkListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGlobalDataNetworkListResponse{}
	_body, _err := client.DescribeGlobalDataNetworkListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a Global Database Network (GDN).
//
// @param request - DescribeGlobalDatabaseNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGlobalDatabaseNetworkResponse
func (client *Client) DescribeGlobalDatabaseNetworkWithOptions(request *DescribeGlobalDatabaseNetworkRequest, runtime *dara.RuntimeOptions) (_result *DescribeGlobalDatabaseNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GDNId) {
		query["GDNId"] = request.GDNId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGlobalDatabaseNetwork"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGlobalDatabaseNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a Global Database Network (GDN).
//
// @param request - DescribeGlobalDatabaseNetworkRequest
//
// @return DescribeGlobalDatabaseNetworkResponse
func (client *Client) DescribeGlobalDatabaseNetwork(request *DescribeGlobalDatabaseNetworkRequest) (_result *DescribeGlobalDatabaseNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGlobalDatabaseNetworkResponse{}
	_body, _err := client.DescribeGlobalDatabaseNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about all Global Database Networks (GDNs) that belong to an account.
//
// @param request - DescribeGlobalDatabaseNetworksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGlobalDatabaseNetworksResponse
func (client *Client) DescribeGlobalDatabaseNetworksWithOptions(request *DescribeGlobalDatabaseNetworksRequest, runtime *dara.RuntimeOptions) (_result *DescribeGlobalDatabaseNetworksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.FilterRegion) {
		query["FilterRegion"] = request.FilterRegion
	}

	if !dara.IsNil(request.GDNDescription) {
		query["GDNDescription"] = request.GDNDescription
	}

	if !dara.IsNil(request.GDNId) {
		query["GDNId"] = request.GDNId
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGlobalDatabaseNetworks"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGlobalDatabaseNetworksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all Global Database Networks (GDNs) that belong to an account.
//
// @param request - DescribeGlobalDatabaseNetworksRequest
//
// @return DescribeGlobalDatabaseNetworksResponse
func (client *Client) DescribeGlobalDatabaseNetworks(request *DescribeGlobalDatabaseNetworksRequest) (_result *DescribeGlobalDatabaseNetworksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGlobalDatabaseNetworksResponse{}
	_body, _err := client.DescribeGlobalDatabaseNetworksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries global IP whitelist templates.
//
// @param request - DescribeGlobalSecurityIPGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGlobalSecurityIPGroupResponse
func (client *Client) DescribeGlobalSecurityIPGroupWithOptions(request *DescribeGlobalSecurityIPGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeGlobalSecurityIPGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GlobalSecurityGroupId) {
		query["GlobalSecurityGroupId"] = request.GlobalSecurityGroupId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGlobalSecurityIPGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGlobalSecurityIPGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries global IP whitelist templates.
//
// @param request - DescribeGlobalSecurityIPGroupRequest
//
// @return DescribeGlobalSecurityIPGroupResponse
func (client *Client) DescribeGlobalSecurityIPGroup(request *DescribeGlobalSecurityIPGroupRequest) (_result *DescribeGlobalSecurityIPGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGlobalSecurityIPGroupResponse{}
	_body, _err := client.DescribeGlobalSecurityIPGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the relationship between a cluster and a global IP whitelist template.
//
// @param request - DescribeGlobalSecurityIPGroupRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGlobalSecurityIPGroupRelationResponse
func (client *Client) DescribeGlobalSecurityIPGroupRelationWithOptions(request *DescribeGlobalSecurityIPGroupRelationRequest, runtime *dara.RuntimeOptions) (_result *DescribeGlobalSecurityIPGroupRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGlobalSecurityIPGroupRelation"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGlobalSecurityIPGroupRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the relationship between a cluster and a global IP whitelist template.
//
// @param request - DescribeGlobalSecurityIPGroupRelationRequest
//
// @return DescribeGlobalSecurityIPGroupRelationResponse
func (client *Client) DescribeGlobalSecurityIPGroupRelation(request *DescribeGlobalSecurityIPGroupRelationRequest) (_result *DescribeGlobalSecurityIPGroupRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGlobalSecurityIPGroupRelationResponse{}
	_body, _err := client.DescribeGlobalSecurityIPGroupRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the HA logs of a cluster.
//
// @param request - DescribeHALogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHALogsResponse
func (client *Client) DescribeHALogsWithOptions(request *DescribeHALogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeHALogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNodeId) {
		query["DBNodeId"] = request.DBNodeId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.LogType) {
		query["LogType"] = request.LogType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHALogs"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHALogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the HA logs of a cluster.
//
// @param request - DescribeHALogsRequest
//
// @return DescribeHALogsResponse
func (client *Client) DescribeHALogs(request *DescribeHALogsRequest) (_result *DescribeHALogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHALogsResponse{}
	_body, _err := client.DescribeHALogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 事件中心事件列表
//
// @param request - DescribeHistoryEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHistoryEventsResponse
func (client *Client) DescribeHistoryEventsWithOptions(request *DescribeHistoryEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeHistoryEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArchiveStatus) {
		query["ArchiveStatus"] = request.ArchiveStatus
	}

	if !dara.IsNil(request.EventCategory) {
		query["EventCategory"] = request.EventCategory
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.EventLevel) {
		query["EventLevel"] = request.EventLevel
	}

	if !dara.IsNil(request.EventStatus) {
		query["EventStatus"] = request.EventStatus
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.FromStartTime) {
		query["FromStartTime"] = request.FromStartTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.ToStartTime) {
		query["ToStartTime"] = request.ToStartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHistoryEvents"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHistoryEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 事件中心事件列表
//
// @param request - DescribeHistoryEventsRequest
//
// @return DescribeHistoryEventsResponse
func (client *Client) DescribeHistoryEvents(request *DescribeHistoryEventsRequest) (_result *DescribeHistoryEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHistoryEventsResponse{}
	_body, _err := client.DescribeHistoryEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries historical tasks.
//
// @param request - DescribeHistoryTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHistoryTasksResponse
func (client *Client) DescribeHistoryTasksWithOptions(request *DescribeHistoryTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeHistoryTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FromExecTime) {
		query["FromExecTime"] = request.FromExecTime
	}

	if !dara.IsNil(request.FromStartTime) {
		query["FromStartTime"] = request.FromStartTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.ToExecTime) {
		query["ToExecTime"] = request.ToExecTime
	}

	if !dara.IsNil(request.ToStartTime) {
		query["ToStartTime"] = request.ToStartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHistoryTasks"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHistoryTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries historical tasks.
//
// @param request - DescribeHistoryTasksRequest
//
// @return DescribeHistoryTasksResponse
func (client *Client) DescribeHistoryTasks(request *DescribeHistoryTasksRequest) (_result *DescribeHistoryTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHistoryTasksResponse{}
	_body, _err := client.DescribeHistoryTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 任务中心任务统计
//
// @param request - DescribeHistoryTasksStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHistoryTasksStatResponse
func (client *Client) DescribeHistoryTasksStatWithOptions(request *DescribeHistoryTasksStatRequest, runtime *dara.RuntimeOptions) (_result *DescribeHistoryTasksStatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FromExecTime) {
		query["FromExecTime"] = request.FromExecTime
	}

	if !dara.IsNil(request.FromStartTime) {
		query["FromStartTime"] = request.FromStartTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.ToExecTime) {
		query["ToExecTime"] = request.ToExecTime
	}

	if !dara.IsNil(request.ToStartTime) {
		query["ToStartTime"] = request.ToStartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHistoryTasksStat"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHistoryTasksStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 任务中心任务统计
//
// @param request - DescribeHistoryTasksStatRequest
//
// @return DescribeHistoryTasksStatResponse
func (client *Client) DescribeHistoryTasksStat(request *DescribeHistoryTasksStatRequest) (_result *DescribeHistoryTasksStatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHistoryTasksStatResponse{}
	_body, _err := client.DescribeHistoryTasksStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of a license order.
//
// @param request - DescribeLicenseOrderDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLicenseOrderDetailsResponse
func (client *Client) DescribeLicenseOrderDetailsWithOptions(request *DescribeLicenseOrderDetailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLicenseOrderDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunOrderId) {
		query["AliyunOrderId"] = request.AliyunOrderId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLicenseOrderDetails"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLicenseOrderDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of a license order.
//
// @param request - DescribeLicenseOrderDetailsRequest
//
// @return DescribeLicenseOrderDetailsResponse
func (client *Client) DescribeLicenseOrderDetails(request *DescribeLicenseOrderDetailsRequest) (_result *DescribeLicenseOrderDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLicenseOrderDetailsResponse{}
	_body, _err := client.DescribeLicenseOrderDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of license orders.
//
// @param request - DescribeLicenseOrdersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLicenseOrdersResponse
func (client *Client) DescribeLicenseOrdersWithOptions(request *DescribeLicenseOrdersRequest, runtime *dara.RuntimeOptions) (_result *DescribeLicenseOrdersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunOrderId) {
		query["AliyunOrderId"] = request.AliyunOrderId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PackageType) {
		query["PackageType"] = request.PackageType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PurchaseChannel) {
		query["PurchaseChannel"] = request.PurchaseChannel
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VirtualOrder) {
		query["VirtualOrder"] = request.VirtualOrder
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLicenseOrders"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLicenseOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of license orders.
//
// @param request - DescribeLicenseOrdersRequest
//
// @return DescribeLicenseOrdersResponse
func (client *Client) DescribeLicenseOrders(request *DescribeLicenseOrdersRequest) (_result *DescribeLicenseOrdersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLicenseOrdersResponse{}
	_body, _err := client.DescribeLicenseOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询库表恢复可恢复时间范围
//
// @param request - DescribeLocalAvailableRecoveryTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLocalAvailableRecoveryTimeResponse
func (client *Client) DescribeLocalAvailableRecoveryTimeWithOptions(request *DescribeLocalAvailableRecoveryTimeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLocalAvailableRecoveryTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLocalAvailableRecoveryTime"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLocalAvailableRecoveryTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询库表恢复可恢复时间范围
//
// @param request - DescribeLocalAvailableRecoveryTimeRequest
//
// @return DescribeLocalAvailableRecoveryTimeResponse
func (client *Client) DescribeLocalAvailableRecoveryTime(request *DescribeLocalAvailableRecoveryTimeRequest) (_result *DescribeLocalAvailableRecoveryTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLocalAvailableRecoveryTimeResponse{}
	_body, _err := client.DescribeLocalAvailableRecoveryTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the retention policy of log backups in a PolarDB cluster.
//
// @param request - DescribeLogBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogBackupPolicyResponse
func (client *Client) DescribeLogBackupPolicyWithOptions(request *DescribeLogBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogBackupPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLogBackupPolicy"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the retention policy of log backups in a PolarDB cluster.
//
// @param request - DescribeLogBackupPolicyRequest
//
// @return DescribeLogBackupPolicyResponse
func (client *Client) DescribeLogBackupPolicy(request *DescribeLogBackupPolicyRequest) (_result *DescribeLogBackupPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLogBackupPolicyResponse{}
	_body, _err := client.DescribeLogBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the data masking rules of a PolarDB cluster or the information about a specified masking rule.
//
// @param request - DescribeMaskingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMaskingRulesResponse
func (client *Client) DescribeMaskingRulesWithOptions(request *DescribeMaskingRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeMaskingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.InterfaceVersion) {
		query["InterfaceVersion"] = request.InterfaceVersion
	}

	if !dara.IsNil(request.RuleNameList) {
		query["RuleNameList"] = request.RuleNameList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMaskingRules"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMaskingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data masking rules of a PolarDB cluster or the information about a specified masking rule.
//
// @param request - DescribeMaskingRulesRequest
//
// @return DescribeMaskingRulesResponse
func (client *Client) DescribeMaskingRules(request *DescribeMaskingRulesRequest) (_result *DescribeMaskingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMaskingRulesResponse{}
	_body, _err := client.DescribeMaskingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the databases or tables that can be restored.
//
// @param request - DescribeMetaListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetaListResponse
func (client *Client) DescribeMetaListWithOptions(request *DescribeMetaListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetaListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.GetDbName) {
		query["GetDbName"] = request.GetDbName
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

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetaList"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetaListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the databases or tables that can be restored.
//
// @param request - DescribeMetaListRequest
//
// @return DescribeMetaListResponse
func (client *Client) DescribeMetaList(request *DescribeMetaListRequest) (_result *DescribeMetaListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetaListResponse{}
	_body, _err := client.DescribeMetaListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询路由规则列表
//
// @param request - DescribeModelApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModelApisResponse
func (client *Client) DescribeModelApisWithOptions(request *DescribeModelApisRequest, runtime *dara.RuntimeOptions) (_result *DescribeModelApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.ModelApiIds) {
		query["ModelApiIds"] = request.ModelApiIds
	}

	if !dara.IsNil(request.ModelCategory) {
		query["ModelCategory"] = request.ModelCategory
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PathPrefix) {
		query["PathPrefix"] = request.PathPrefix
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModelApis"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModelApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询路由规则列表
//
// @param request - DescribeModelApisRequest
//
// @return DescribeModelApisResponse
func (client *Client) DescribeModelApis(request *DescribeModelApisRequest) (_result *DescribeModelApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeModelApisResponse{}
	_body, _err := client.DescribeModelApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询模型服务列表
//
// @param request - DescribeModelServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModelServicesResponse
func (client *Client) DescribeModelServicesWithOptions(request *DescribeModelServicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeModelServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.ModelCategory) {
		query["ModelCategory"] = request.ModelCategory
	}

	if !dara.IsNil(request.ModelServiceIds) {
		query["ModelServiceIds"] = request.ModelServiceIds
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModelServices"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModelServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询模型服务列表
//
// @param request - DescribeModelServicesRequest
//
// @return DescribeModelServicesResponse
func (client *Client) DescribeModelServices(request *DescribeModelServicesRequest) (_result *DescribeModelServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeModelServicesResponse{}
	_body, _err := client.DescribeModelServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询参数修改历史
//
// @param request - DescribeModifyParameterLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModifyParameterLogResponse
func (client *Client) DescribeModifyParameterLogWithOptions(request *DescribeModifyParameterLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeModifyParameterLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModifyParameterLog"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModifyParameterLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询参数修改历史
//
// @param request - DescribeModifyParameterLogRequest
//
// @return DescribeModifyParameterLogResponse
func (client *Client) DescribeModifyParameterLog(request *DescribeModifyParameterLogRequest) (_result *DescribeModifyParameterLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeModifyParameterLogResponse{}
	_body, _err := client.DescribeModifyParameterLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询网络通道
//
// @param request - DescribeNetworkChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkChannelResponse
func (client *Client) DescribeNetworkChannelWithOptions(request *DescribeNetworkChannelRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNetworkChannel"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNetworkChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询网络通道
//
// @param request - DescribeNetworkChannelRequest
//
// @return DescribeNetworkChannelResponse
func (client *Client) DescribeNetworkChannel(request *DescribeNetworkChannelRequest) (_result *DescribeNetworkChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNetworkChannelResponse{}
	_body, _err := client.DescribeNetworkChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a parameter template.
//
// Description:
//
// You can use parameter templates to manage multiple parameters at a time and apply existing parameters to a PolarDB cluster. For more information, see [Use a parameter template](https://help.aliyun.com/document_detail/207009.html).
//
// > This parameter is valid only for a PolarDB for MySQL cluster.
//
// @param request - DescribeParameterGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParameterGroupResponse
func (client *Client) DescribeParameterGroupWithOptions(request *DescribeParameterGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeParameterGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBType) {
		query["DBType"] = request.DBType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParameterGroupId) {
		query["ParameterGroupId"] = request.ParameterGroupId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParameterGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParameterGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a parameter template.
//
// Description:
//
// You can use parameter templates to manage multiple parameters at a time and apply existing parameters to a PolarDB cluster. For more information, see [Use a parameter template](https://help.aliyun.com/document_detail/207009.html).
//
// > This parameter is valid only for a PolarDB for MySQL cluster.
//
// @param request - DescribeParameterGroupRequest
//
// @return DescribeParameterGroupResponse
func (client *Client) DescribeParameterGroup(request *DescribeParameterGroupRequest) (_result *DescribeParameterGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeParameterGroupResponse{}
	_body, _err := client.DescribeParameterGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries parameter templates that are available in a specified region.
//
// Description:
//
// You can use parameter templates to manage multiple parameters at a time and apply existing parameters to a PolarDB cluster. For more information, see [Use a parameter template](https://help.aliyun.com/document_detail/207009.html).
//
// > This operation is applicable only to PolarDB for MySQL clusters.
//
// @param request - DescribeParameterGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParameterGroupsResponse
func (client *Client) DescribeParameterGroupsWithOptions(request *DescribeParameterGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeParameterGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBType) {
		query["DBType"] = request.DBType
	}

	if !dara.IsNil(request.DBVersion) {
		query["DBVersion"] = request.DBVersion
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParameterGroups"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParameterGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries parameter templates that are available in a specified region.
//
// Description:
//
// You can use parameter templates to manage multiple parameters at a time and apply existing parameters to a PolarDB cluster. For more information, see [Use a parameter template](https://help.aliyun.com/document_detail/207009.html).
//
// > This operation is applicable only to PolarDB for MySQL clusters.
//
// @param request - DescribeParameterGroupsRequest
//
// @return DescribeParameterGroupsResponse
func (client *Client) DescribeParameterGroups(request *DescribeParameterGroupsRequest) (_result *DescribeParameterGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeParameterGroupsResponse{}
	_body, _err := client.DescribeParameterGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the default parameters in a cluster.
//
// @param request - DescribeParameterTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParameterTemplatesResponse
func (client *Client) DescribeParameterTemplatesWithOptions(request *DescribeParameterTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeParameterTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBType) {
		query["DBType"] = request.DBType
	}

	if !dara.IsNil(request.DBVersion) {
		query["DBVersion"] = request.DBVersion
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParameterTemplates"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParameterTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the default parameters in a cluster.
//
// @param request - DescribeParameterTemplatesRequest
//
// @return DescribeParameterTemplatesResponse
func (client *Client) DescribeParameterTemplates(request *DescribeParameterTemplatesRequest) (_result *DescribeParameterTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeParameterTemplatesResponse{}
	_body, _err := client.DescribeParameterTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a pending event.
//
// @param request - DescribePendingMaintenanceActionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePendingMaintenanceActionResponse
func (client *Client) DescribePendingMaintenanceActionWithOptions(request *DescribePendingMaintenanceActionRequest, runtime *dara.RuntimeOptions) (_result *DescribePendingMaintenanceActionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsHistory) {
		query["IsHistory"] = request.IsHistory
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

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePendingMaintenanceAction"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePendingMaintenanceActionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a pending event.
//
// @param request - DescribePendingMaintenanceActionRequest
//
// @return DescribePendingMaintenanceActionResponse
func (client *Client) DescribePendingMaintenanceAction(request *DescribePendingMaintenanceActionRequest) (_result *DescribePendingMaintenanceActionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePendingMaintenanceActionResponse{}
	_body, _err := client.DescribePendingMaintenanceActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the numbers of scheduled events for different types of tasks.
//
// @param request - DescribePendingMaintenanceActionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePendingMaintenanceActionsResponse
func (client *Client) DescribePendingMaintenanceActionsWithOptions(request *DescribePendingMaintenanceActionsRequest, runtime *dara.RuntimeOptions) (_result *DescribePendingMaintenanceActionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsHistory) {
		query["IsHistory"] = request.IsHistory
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePendingMaintenanceActions"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePendingMaintenanceActionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the numbers of scheduled events for different types of tasks.
//
// @param request - DescribePendingMaintenanceActionsRequest
//
// @return DescribePendingMaintenanceActionsResponse
func (client *Client) DescribePendingMaintenanceActions(request *DescribePendingMaintenanceActionsRequest) (_result *DescribePendingMaintenanceActionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePendingMaintenanceActionsResponse{}
	_body, _err := client.DescribePendingMaintenanceActionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定会话明细
//
// @param request - DescribePolarAgentChatRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolarAgentChatRecordsResponse
func (client *Client) DescribePolarAgentChatRecordsWithOptions(request *DescribePolarAgentChatRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribePolarAgentChatRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolarAgentChatRecords"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolarAgentChatRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定会话明细
//
// @param request - DescribePolarAgentChatRecordsRequest
//
// @return DescribePolarAgentChatRecordsResponse
func (client *Client) DescribePolarAgentChatRecords(request *DescribePolarAgentChatRecordsRequest) (_result *DescribePolarAgentChatRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePolarAgentChatRecordsResponse{}
	_body, _err := client.DescribePolarAgentChatRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询会话状态
//
// @param request - DescribePolarAgentSessionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolarAgentSessionStatusResponse
func (client *Client) DescribePolarAgentSessionStatusWithOptions(request *DescribePolarAgentSessionStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribePolarAgentSessionStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolarAgentSessionStatus"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolarAgentSessionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会话状态
//
// @param request - DescribePolarAgentSessionStatusRequest
//
// @return DescribePolarAgentSessionStatusResponse
func (client *Client) DescribePolarAgentSessionStatus(request *DescribePolarAgentSessionStatusRequest) (_result *DescribePolarAgentSessionStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePolarAgentSessionStatusResponse{}
	_body, _err := client.DescribePolarAgentSessionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看历史会话记录
//
// @param request - DescribePolarAgentUserSessionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolarAgentUserSessionsResponse
func (client *Client) DescribePolarAgentUserSessionsWithOptions(request *DescribePolarAgentUserSessionsRequest, runtime *dara.RuntimeOptions) (_result *DescribePolarAgentUserSessionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolarAgentUserSessions"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolarAgentUserSessionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看历史会话记录
//
// @param request - DescribePolarAgentUserSessionsRequest
//
// @return DescribePolarAgentUserSessionsResponse
func (client *Client) DescribePolarAgentUserSessions(request *DescribePolarAgentUserSessionsRequest) (_result *DescribePolarAgentUserSessionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePolarAgentUserSessionsResponse{}
	_body, _err := client.DescribePolarAgentUserSessionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询PolarClaw Agent列表
//
// @param tmpReq - DescribePolarClawAgentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolarClawAgentsResponse
func (client *Client) DescribePolarClawAgentsWithOptions(tmpReq *DescribePolarClawAgentsRequest, runtime *dara.RuntimeOptions) (_result *DescribePolarClawAgentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribePolarClawAgentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentList) {
		request.AgentListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentList, dara.String("AgentList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentListShrink) {
		query["AgentList"] = request.AgentListShrink
	}

	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolarClawAgents"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolarClawAgentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询PolarClaw Agent列表
//
// @param request - DescribePolarClawAgentsRequest
//
// @return DescribePolarClawAgentsResponse
func (client *Client) DescribePolarClawAgents(request *DescribePolarClawAgentsRequest) (_result *DescribePolarClawAgentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePolarClawAgentsResponse{}
	_body, _err := client.DescribePolarClawAgentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询PolarClaw Channels
//
// @param tmpReq - DescribePolarClawChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolarClawChannelsResponse
func (client *Client) DescribePolarClawChannelsWithOptions(tmpReq *DescribePolarClawChannelsRequest, runtime *dara.RuntimeOptions) (_result *DescribePolarClawChannelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribePolarClawChannelsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChannelList) {
		request.ChannelListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelList, dara.String("ChannelList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ChannelListShrink) {
		query["ChannelList"] = request.ChannelListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolarClawChannels"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolarClawChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询PolarClaw Channels
//
// @param request - DescribePolarClawChannelsRequest
//
// @return DescribePolarClawChannelsResponse
func (client *Client) DescribePolarClawChannels(request *DescribePolarClawChannelsRequest) (_result *DescribePolarClawChannelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePolarClawChannelsResponse{}
	_body, _err := client.DescribePolarClawChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询PolarClaw定时任务列表
//
// @param tmpReq - DescribePolarClawCronJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolarClawCronJobsResponse
func (client *Client) DescribePolarClawCronJobsWithOptions(tmpReq *DescribePolarClawCronJobsRequest, runtime *dara.RuntimeOptions) (_result *DescribePolarClawCronJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribePolarClawCronJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JobIdList) {
		request.JobIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobIdList, dara.String("JobIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.IncludeDisabled) {
		query["IncludeDisabled"] = request.IncludeDisabled
	}

	if !dara.IsNil(request.IncludeRuns) {
		query["IncludeRuns"] = request.IncludeRuns
	}

	if !dara.IsNil(request.JobIdListShrink) {
		query["JobIdList"] = request.JobIdListShrink
	}

	if !dara.IsNil(request.RunLimit) {
		query["RunLimit"] = request.RunLimit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolarClawCronJobs"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolarClawCronJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询PolarClaw定时任务列表
//
// @param request - DescribePolarClawCronJobsRequest
//
// @return DescribePolarClawCronJobsResponse
func (client *Client) DescribePolarClawCronJobs(request *DescribePolarClawCronJobsRequest) (_result *DescribePolarClawCronJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePolarClawCronJobsResponse{}
	_body, _err := client.DescribePolarClawCronJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询PolarClaw MCP Servers
//
// @param request - DescribePolarClawMCPServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolarClawMCPServersResponse
func (client *Client) DescribePolarClawMCPServersWithOptions(request *DescribePolarClawMCPServersRequest, runtime *dara.RuntimeOptions) (_result *DescribePolarClawMCPServersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ServerList) {
		query["ServerList"] = request.ServerList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolarClawMCPServers"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolarClawMCPServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询PolarClaw MCP Servers
//
// @param request - DescribePolarClawMCPServersRequest
//
// @return DescribePolarClawMCPServersResponse
func (client *Client) DescribePolarClawMCPServers(request *DescribePolarClawMCPServersRequest) (_result *DescribePolarClawMCPServersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePolarClawMCPServersResponse{}
	_body, _err := client.DescribePolarClawMCPServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询PolarClaw Plugins
//
// @param request - DescribePolarClawPluginsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolarClawPluginsResponse
func (client *Client) DescribePolarClawPluginsWithOptions(request *DescribePolarClawPluginsRequest, runtime *dara.RuntimeOptions) (_result *DescribePolarClawPluginsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.PluginList) {
		query["PluginList"] = request.PluginList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolarClawPlugins"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolarClawPluginsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询PolarClaw Plugins
//
// @param request - DescribePolarClawPluginsRequest
//
// @return DescribePolarClawPluginsResponse
func (client *Client) DescribePolarClawPlugins(request *DescribePolarClawPluginsRequest) (_result *DescribePolarClawPluginsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePolarClawPluginsResponse{}
	_body, _err := client.DescribePolarClawPluginsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取PolarFS实例详情
//
// @param request - DescribePolarFsAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolarFsAttributeResponse
func (client *Client) DescribePolarFsAttributeWithOptions(request *DescribePolarFsAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribePolarFsAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.PolarFsInstanceId) {
		query["PolarFsInstanceId"] = request.PolarFsInstanceId
	}

	if !dara.IsNil(request.QueryFuseMountInfo) {
		query["QueryFuseMountInfo"] = request.QueryFuseMountInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolarFsAttribute"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolarFsAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取PolarFS实例详情
//
// @param request - DescribePolarFsAttributeRequest
//
// @return DescribePolarFsAttributeResponse
func (client *Client) DescribePolarFsAttribute(request *DescribePolarFsAttributeRequest) (_result *DescribePolarFsAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePolarFsAttributeResponse{}
	_body, _err := client.DescribePolarFsAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出指定路径下的文件和子目录信息。
//
// Description:
//
// ## 请求说明
//
// - **Path*	- 参数必须提供一个绝对路径。
//
// - **Recursive*	- 参数默认为 `false`，如果设置为 `true`，则会递归列出所有子目录的内容。
//
// - **Depth*	- 参数用于限制递归深度，默认值为 `1`。
//
// - **Filter*	- 参数支持通配符或正则表达式过滤结果。
//
// @param request - DescribePolarFsObjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolarFsObjectsResponse
func (client *Client) DescribePolarFsObjectsWithOptions(request *DescribePolarFsObjectsRequest, runtime *dara.RuntimeOptions) (_result *DescribePolarFsObjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.PolarFsInstanceId) {
		query["PolarFsInstanceId"] = request.PolarFsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolarFsObjects"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolarFsObjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出指定路径下的文件和子目录信息。
//
// Description:
//
// ## 请求说明
//
// - **Path*	- 参数必须提供一个绝对路径。
//
// - **Recursive*	- 参数默认为 `false`，如果设置为 `true`，则会递归列出所有子目录的内容。
//
// - **Depth*	- 参数用于限制递归深度，默认值为 `1`。
//
// - **Filter*	- 参数支持通配符或正则表达式过滤结果。
//
// @param request - DescribePolarFsObjectsRequest
//
// @return DescribePolarFsObjectsResponse
func (client *Client) DescribePolarFsObjects(request *DescribePolarFsObjectsRequest) (_result *DescribePolarFsObjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePolarFsObjectsResponse{}
	_body, _err := client.DescribePolarFsObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询配额规则
//
// @param request - DescribePolarFsQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolarFsQuotaResponse
func (client *Client) DescribePolarFsQuotaWithOptions(request *DescribePolarFsQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribePolarFsQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.PolarFsInstanceId) {
		query["PolarFsInstanceId"] = request.PolarFsInstanceId
	}

	if !dara.IsNil(request.QuotaType) {
		query["QuotaType"] = request.QuotaType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolarFsQuota"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolarFsQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询配额规则
//
// @param request - DescribePolarFsQuotaRequest
//
// @return DescribePolarFsQuotaResponse
func (client *Client) DescribePolarFsQuota(request *DescribePolarFsQuotaRequest) (_result *DescribePolarFsQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePolarFsQuotaResponse{}
	_body, _err := client.DescribePolarFsQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether the SQL Explorer feature is enabled for the cluster.
//
// @param request - DescribePolarSQLCollectorPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolarSQLCollectorPolicyResponse
func (client *Client) DescribePolarSQLCollectorPolicyWithOptions(request *DescribePolarSQLCollectorPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribePolarSQLCollectorPolicyResponse, _err error) {
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
		Action:      dara.String("DescribePolarSQLCollectorPolicy"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolarSQLCollectorPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the SQL Explorer feature is enabled for the cluster.
//
// @param request - DescribePolarSQLCollectorPolicyRequest
//
// @return DescribePolarSQLCollectorPolicyResponse
func (client *Client) DescribePolarSQLCollectorPolicy(request *DescribePolarSQLCollectorPolicyRequest) (_result *DescribePolarSQLCollectorPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePolarSQLCollectorPolicyResponse{}
	_body, _err := client.DescribePolarSQLCollectorPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询限流策略
//
// @param request - DescribeRateLimitPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRateLimitPolicyResponse
func (client *Client) DescribeRateLimitPolicyWithOptions(request *DescribeRateLimitPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeRateLimitPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScopeRefId) {
		query["ScopeRefId"] = request.ScopeRefId
	}

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRateLimitPolicy"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRateLimitPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询限流策略
//
// @param request - DescribeRateLimitPolicyRequest
//
// @return DescribeRateLimitPolicyResponse
func (client *Client) DescribeRateLimitPolicy(request *DescribeRateLimitPolicyRequest) (_result *DescribeRateLimitPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRateLimitPolicyResponse{}
	_body, _err := client.DescribeRateLimitPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRdsVSwitchsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRdsVSwitchsResponse
func (client *Client) DescribeRdsVSwitchsWithOptions(request *DescribeRdsVSwitchsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRdsVSwitchsResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRdsVSwitchs"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRdsVSwitchsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRdsVSwitchsRequest
//
// @return DescribeRdsVSwitchsResponse
func (client *Client) DescribeRdsVSwitchs(request *DescribeRdsVSwitchsRequest) (_result *DescribeRdsVSwitchsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRdsVSwitchsResponse{}
	_body, _err := client.DescribeRdsVSwitchsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询vpc
//
// @param request - DescribeRdsVpcsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRdsVpcsResponse
func (client *Client) DescribeRdsVpcsWithOptions(request *DescribeRdsVpcsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRdsVpcsResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRdsVpcs"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRdsVpcsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询vpc
//
// @param request - DescribeRdsVpcsRequest
//
// @return DescribeRdsVpcsResponse
func (client *Client) DescribeRdsVpcs(request *DescribeRdsVpcsRequest) (_result *DescribeRdsVpcsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRdsVpcsResponse{}
	_body, _err := client.DescribeRdsVpcsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions and zones that are supported by PolarDB.
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2017-08-01"),
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
// Queries the regions and zones that are supported by PolarDB.
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
// 查询资源包列表
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourcePackagesResponse
func (client *Client) DescribeResourcePackagesWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeResourcePackagesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourcePackages"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourcePackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资源包列表
//
// @return DescribeResourcePackagesResponse
func (client *Client) DescribeResourcePackages() (_result *DescribeResourcePackagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResourcePackagesResponse{}
	_body, _err := client.DescribeResourcePackagesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询SQL限流规则信息
//
// @param request - DescribeSQLRateLimitingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSQLRateLimitingRulesResponse
func (client *Client) DescribeSQLRateLimitingRulesWithOptions(request *DescribeSQLRateLimitingRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSQLRateLimitingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RuleNameList) {
		query["RuleNameList"] = request.RuleNameList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSQLRateLimitingRules"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSQLRateLimitingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询SQL限流规则信息
//
// @param request - DescribeSQLRateLimitingRulesRequest
//
// @return DescribeSQLRateLimitingRulesResponse
func (client *Client) DescribeSQLRateLimitingRules(request *DescribeSQLRateLimitingRulesRequest) (_result *DescribeSQLRateLimitingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSQLRateLimitingRulesResponse{}
	_body, _err := client.DescribeSQLRateLimitingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of all scheduled tasks.
//
// @param request - DescribeScheduleTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScheduleTasksResponse
func (client *Client) DescribeScheduleTasksWithOptions(request *DescribeScheduleTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeScheduleTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterDescription) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
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

	if !dara.IsNil(request.PlannedEndTime) {
		query["PlannedEndTime"] = request.PlannedEndTime
	}

	if !dara.IsNil(request.PlannedStartTime) {
		query["PlannedStartTime"] = request.PlannedStartTime
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

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskAction) {
		query["TaskAction"] = request.TaskAction
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScheduleTasks"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScheduleTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of all scheduled tasks.
//
// @param request - DescribeScheduleTasksRequest
//
// @return DescribeScheduleTasksResponse
func (client *Client) DescribeScheduleTasks(request *DescribeScheduleTasksRequest) (_result *DescribeScheduleTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeScheduleTasksResponse{}
	_body, _err := client.DescribeScheduleTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Slow Log Details
//
// Description:
//
// >- Only PolarDB MySQL Edition clusters support calling this interface.
//
// >- Starting from September 1, 2024, due to the optimization of the SQL template algorithm, when calling this interface, the value of the SQLHash field will change. For more details, please refer to [Notice] Optimization of Slow SQL Template Algorithm (~~2845725~~).
//
// @param request - DescribeSlowLogRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlowLogRecordsResponse
func (client *Client) DescribeSlowLogRecordsWithOptions(request *DescribeSlowLogRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlowLogRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
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

	if !dara.IsNil(request.SQLHASH) {
		query["SQLHASH"] = request.SQLHASH
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlowLogRecords"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Slow Log Details
//
// Description:
//
// >- Only PolarDB MySQL Edition clusters support calling this interface.
//
// >- Starting from September 1, 2024, due to the optimization of the SQL template algorithm, when calling this interface, the value of the SQLHash field will change. For more details, please refer to [Notice] Optimization of Slow SQL Template Algorithm (~~2845725~~).
//
// @param request - DescribeSlowLogRecordsRequest
//
// @return DescribeSlowLogRecordsResponse
func (client *Client) DescribeSlowLogRecords(request *DescribeSlowLogRecordsRequest) (_result *DescribeSlowLogRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.DescribeSlowLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics about the slow query logs of a PolarDB cluster.
//
// Description:
//
// > This operation is applicable only to PolarDB for MySQL clusters.
//
// @param request - DescribeSlowLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlowLogsResponse
func (client *Client) DescribeSlowLogsWithOptions(request *DescribeSlowLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlowLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlowLogs"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlowLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics about the slow query logs of a PolarDB cluster.
//
// Description:
//
// > This operation is applicable only to PolarDB for MySQL clusters.
//
// @param request - DescribeSlowLogsRequest
//
// @return DescribeSlowLogsResponse
func (client *Client) DescribeSlowLogs(request *DescribeSlowLogsRequest) (_result *DescribeSlowLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSlowLogsResponse{}
	_body, _err := client.DescribeSlowLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the tasks that are generated based on API operations, such as the status of instance creation tasks.
//
// Description:
//
//	  You can call this operation to view the details of a task that is generated by a specific API operation or in the PolarDB console. The system calls the specific API operation when you perform an operation in the PolarDB console. For example, you can view the details of the task when you call the [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) operation or [create a cluster](https://help.aliyun.com/document_detail/58769.html) in the PolarDB console.
//
//		- You can view the details of tasks that are generated only when you call the [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) operation to create a cluster and `CreationOption` is not set to `CreateGdnStandby`.
//
// @param request - DescribeTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTasksResponse
func (client *Client) DescribeTasksWithOptions(request *DescribeTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNodeId) {
		query["DBNodeId"] = request.DBNodeId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTasks"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the tasks that are generated based on API operations, such as the status of instance creation tasks.
//
// Description:
//
//	  You can call this operation to view the details of a task that is generated by a specific API operation or in the PolarDB console. The system calls the specific API operation when you perform an operation in the PolarDB console. For example, you can view the details of the task when you call the [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) operation or [create a cluster](https://help.aliyun.com/document_detail/58769.html) in the PolarDB console.
//
//		- You can view the details of tasks that are generated only when you call the [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) operation to create a cluster and `CreationOption` is not set to `CreateGdnStandby`.
//
// @param request - DescribeTasksRequest
//
// @return DescribeTasksResponse
func (client *Client) DescribeTasks(request *DescribeTasksRequest) (_result *DescribeTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTasksResponse{}
	_body, _err := client.DescribeTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询校验报告
//
// @param request - DescribeUpgradeReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUpgradeReportResponse
func (client *Client) DescribeUpgradeReportWithOptions(request *DescribeUpgradeReportRequest, runtime *dara.RuntimeOptions) (_result *DescribeUpgradeReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreationCategory) {
		query["CreationCategory"] = request.CreationCategory
	}

	if !dara.IsNil(request.DBType) {
		query["DBType"] = request.DBType
	}

	if !dara.IsNil(request.DBVersion) {
		query["DBVersion"] = request.DBVersion
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

	if !dara.IsNil(request.SourceDBClusterId) {
		query["SourceDBClusterId"] = request.SourceDBClusterId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUpgradeReport"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUpgradeReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询校验报告
//
// @param request - DescribeUpgradeReportRequest
//
// @return DescribeUpgradeReportResponse
func (client *Client) DescribeUpgradeReport(request *DescribeUpgradeReportRequest) (_result *DescribeUpgradeReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUpgradeReportResponse{}
	_body, _err := client.DescribeUpgradeReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Key Management Service (KMS)-managed customer master keys (CMKs) that are used to encrypt data in a PolarDB cluster.
//
// @param request - DescribeUserEncryptionKeyListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserEncryptionKeyListResponse
func (client *Client) DescribeUserEncryptionKeyListWithOptions(request *DescribeUserEncryptionKeyListRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserEncryptionKeyListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.TDERegion) {
		query["TDERegion"] = request.TDERegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserEncryptionKeyList"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserEncryptionKeyListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Key Management Service (KMS)-managed customer master keys (CMKs) that are used to encrypt data in a PolarDB cluster.
//
// @param request - DescribeUserEncryptionKeyListRequest
//
// @return DescribeUserEncryptionKeyListResponse
func (client *Client) DescribeUserEncryptionKeyList(request *DescribeUserEncryptionKeyListRequest) (_result *DescribeUserEncryptionKeyListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserEncryptionKeyListResponse{}
	_body, _err := client.DescribeUserEncryptionKeyListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询交换机信息
//
// @param request - DescribeVSwitchListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVSwitchListResponse
func (client *Client) DescribeVSwitchListWithOptions(request *DescribeVSwitchListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVSwitchListResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VSwitchIds) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVSwitchList"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVSwitchListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询交换机信息
//
// @param request - DescribeVSwitchListRequest
//
// @return DescribeVSwitchListResponse
func (client *Client) DescribeVSwitchList(request *DescribeVSwitchListRequest) (_result *DescribeVSwitchListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVSwitchListResponse{}
	_body, _err := client.DescribeVSwitchListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a vSwitch.
//
// @param request - DescribeVSwitchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVSwitchesResponse
func (client *Client) DescribeVSwitchesWithOptions(request *DescribeVSwitchesRequest, runtime *dara.RuntimeOptions) (_result *DescribeVSwitchesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DedicatedHostGroupId) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVSwitches"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVSwitchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a vSwitch.
//
// @param request - DescribeVSwitchesRequest
//
// @return DescribeVSwitchesResponse
func (client *Client) DescribeVSwitches(request *DescribeVSwitchesRequest) (_result *DescribeVSwitchesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVSwitchesResponse{}
	_body, _err := client.DescribeVSwitchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询vpc
//
// @param request - DescribeVpcsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcsResponse
func (client *Client) DescribeVpcsWithOptions(request *DescribeVpcsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcsResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcs"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询vpc
//
// @param request - DescribeVpcsRequest
//
// @return DescribeVpcsResponse
func (client *Client) DescribeVpcs(request *DescribeVpcsRequest) (_result *DescribeVpcsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVpcsResponse{}
	_body, _err := client.DescribeVpcsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取可用区
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
	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.Extra) {
		query["Extra"] = request.Extra
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZones"),
		Version:     dara.String("2017-08-01"),
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
// 获取可用区
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
// 关闭DynamoDB兼容性能力
//
// @param request - DisableDBClusterDynamoDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableDBClusterDynamoDBResponse
func (client *Client) DisableDBClusterDynamoDBWithOptions(request *DisableDBClusterDynamoDBRequest, runtime *dara.RuntimeOptions) (_result *DisableDBClusterDynamoDBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableDBClusterDynamoDB"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableDBClusterDynamoDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭DynamoDB兼容性能力
//
// @param request - DisableDBClusterDynamoDBRequest
//
// @return DisableDBClusterDynamoDBResponse
func (client *Client) DisableDBClusterDynamoDB(request *DisableDBClusterDynamoDBRequest) (_result *DisableDBClusterDynamoDBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableDBClusterDynamoDBResponse{}
	_body, _err := client.DisableDBClusterDynamoDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关闭集群的关系性
//
// @param request - DisableDBClusterOrcaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableDBClusterOrcaResponse
func (client *Client) DisableDBClusterOrcaWithOptions(request *DisableDBClusterOrcaRequest, runtime *dara.RuntimeOptions) (_result *DisableDBClusterOrcaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CleanData) {
		query["CleanData"] = request.CleanData
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableDBClusterOrca"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableDBClusterOrcaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭集群的关系性
//
// @param request - DisableDBClusterOrcaRequest
//
// @return DisableDBClusterOrcaResponse
func (client *Client) DisableDBClusterOrca(request *DisableDBClusterOrcaRequest) (_result *DisableDBClusterOrcaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableDBClusterOrcaResponse{}
	_body, _err := client.DisableDBClusterOrcaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a stable serverless cluster.
//
// @param request - DisableDBClusterServerlessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableDBClusterServerlessResponse
func (client *Client) DisableDBClusterServerlessWithOptions(request *DisableDBClusterServerlessRequest, runtime *dara.RuntimeOptions) (_result *DisableDBClusterServerlessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableDBClusterServerless"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableDBClusterServerlessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a stable serverless cluster.
//
// @param request - DisableDBClusterServerlessRequest
//
// @return DisableDBClusterServerlessResponse
func (client *Client) DisableDBClusterServerless(request *DisableDBClusterServerlessRequest) (_result *DisableDBClusterServerlessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableDBClusterServerlessResponse{}
	_body, _err := client.DisableDBClusterServerlessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 禁用PolarClaw Channel
//
// @param request - DisablePolarClawChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisablePolarClawChannelResponse
func (client *Client) DisablePolarClawChannelWithOptions(request *DisablePolarClawChannelRequest, runtime *dara.RuntimeOptions) (_result *DisablePolarClawChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisablePolarClawChannel"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisablePolarClawChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 禁用PolarClaw Channel
//
// @param request - DisablePolarClawChannelRequest
//
// @return DisablePolarClawChannelResponse
func (client *Client) DisablePolarClawChannel(request *DisablePolarClawChannelRequest) (_result *DisablePolarClawChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisablePolarClawChannelResponse{}
	_body, _err := client.DisablePolarClawChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 禁用PolarClaw定时任务
//
// @param request - DisablePolarClawCronJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisablePolarClawCronJobResponse
func (client *Client) DisablePolarClawCronJobWithOptions(request *DisablePolarClawCronJobRequest, runtime *dara.RuntimeOptions) (_result *DisablePolarClawCronJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisablePolarClawCronJob"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisablePolarClawCronJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 禁用PolarClaw定时任务
//
// @param request - DisablePolarClawCronJobRequest
//
// @return DisablePolarClawCronJobResponse
func (client *Client) DisablePolarClawCronJob(request *DisablePolarClawCronJobRequest) (_result *DisablePolarClawCronJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisablePolarClawCronJobResponse{}
	_body, _err := client.DisablePolarClawCronJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 禁用PolarClaw Plugin
//
// @param request - DisablePolarClawPluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisablePolarClawPluginResponse
func (client *Client) DisablePolarClawPluginWithOptions(request *DisablePolarClawPluginRequest, runtime *dara.RuntimeOptions) (_result *DisablePolarClawPluginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisablePolarClawPlugin"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisablePolarClawPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 禁用PolarClaw Plugin
//
// @param request - DisablePolarClawPluginRequest
//
// @return DisablePolarClawPluginResponse
func (client *Client) DisablePolarClawPlugin(request *DisablePolarClawPluginRequest) (_result *DisablePolarClawPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisablePolarClawPluginResponse{}
	_body, _err := client.DisablePolarClawPluginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启DynamoDB兼容性能力
//
// @param request - EnableDBClusterDynamoDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableDBClusterDynamoDBResponse
func (client *Client) EnableDBClusterDynamoDBWithOptions(request *EnableDBClusterDynamoDBRequest, runtime *dara.RuntimeOptions) (_result *EnableDBClusterDynamoDBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableDBClusterDynamoDB"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableDBClusterDynamoDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启DynamoDB兼容性能力
//
// @param request - EnableDBClusterDynamoDBRequest
//
// @return EnableDBClusterDynamoDBResponse
func (client *Client) EnableDBClusterDynamoDB(request *EnableDBClusterDynamoDBRequest) (_result *EnableDBClusterDynamoDBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableDBClusterDynamoDBResponse{}
	_body, _err := client.EnableDBClusterDynamoDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启集群的关系性
//
// @param request - EnableDBClusterOrcaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableDBClusterOrcaResponse
func (client *Client) EnableDBClusterOrcaWithOptions(request *EnableDBClusterOrcaRequest, runtime *dara.RuntimeOptions) (_result *EnableDBClusterOrcaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableDBClusterOrca"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableDBClusterOrcaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启集群的关系性
//
// @param request - EnableDBClusterOrcaRequest
//
// @return EnableDBClusterOrcaResponse
func (client *Client) EnableDBClusterOrca(request *EnableDBClusterOrcaRequest) (_result *EnableDBClusterOrcaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableDBClusterOrcaResponse{}
	_body, _err := client.EnableDBClusterOrcaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a stable serverless cluster.
//
// @param request - EnableDBClusterServerlessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableDBClusterServerlessResponse
func (client *Client) EnableDBClusterServerlessWithOptions(request *EnableDBClusterServerlessRequest, runtime *dara.RuntimeOptions) (_result *EnableDBClusterServerlessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.ScaleApRoNumMax) {
		query["ScaleApRoNumMax"] = request.ScaleApRoNumMax
	}

	if !dara.IsNil(request.ScaleApRoNumMin) {
		query["ScaleApRoNumMin"] = request.ScaleApRoNumMin
	}

	if !dara.IsNil(request.ScaleMax) {
		query["ScaleMax"] = request.ScaleMax
	}

	if !dara.IsNil(request.ScaleMin) {
		query["ScaleMin"] = request.ScaleMin
	}

	if !dara.IsNil(request.ScaleRoNumMax) {
		query["ScaleRoNumMax"] = request.ScaleRoNumMax
	}

	if !dara.IsNil(request.ScaleRoNumMin) {
		query["ScaleRoNumMin"] = request.ScaleRoNumMin
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableDBClusterServerless"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableDBClusterServerlessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a stable serverless cluster.
//
// @param request - EnableDBClusterServerlessRequest
//
// @return EnableDBClusterServerlessResponse
func (client *Client) EnableDBClusterServerless(request *EnableDBClusterServerlessRequest) (_result *EnableDBClusterServerlessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableDBClusterServerlessResponse{}
	_body, _err := client.EnableDBClusterServerlessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the status of SQL firewall rules for a cluster.
//
// @param request - EnableFirewallRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableFirewallRulesResponse
func (client *Client) EnableFirewallRulesWithOptions(request *EnableFirewallRulesRequest, runtime *dara.RuntimeOptions) (_result *EnableFirewallRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
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

	if !dara.IsNil(request.RuleNameList) {
		query["RuleNameList"] = request.RuleNameList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableFirewallRules"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableFirewallRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the status of SQL firewall rules for a cluster.
//
// @param request - EnableFirewallRulesRequest
//
// @return EnableFirewallRulesResponse
func (client *Client) EnableFirewallRules(request *EnableFirewallRulesRequest) (_result *EnableFirewallRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableFirewallRulesResponse{}
	_body, _err := client.EnableFirewallRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启用PolarClaw Channel
//
// @param request - EnablePolarClawChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnablePolarClawChannelResponse
func (client *Client) EnablePolarClawChannelWithOptions(request *EnablePolarClawChannelRequest, runtime *dara.RuntimeOptions) (_result *EnablePolarClawChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnablePolarClawChannel"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnablePolarClawChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启用PolarClaw Channel
//
// @param request - EnablePolarClawChannelRequest
//
// @return EnablePolarClawChannelResponse
func (client *Client) EnablePolarClawChannel(request *EnablePolarClawChannelRequest) (_result *EnablePolarClawChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnablePolarClawChannelResponse{}
	_body, _err := client.EnablePolarClawChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启用PolarClaw定时任务
//
// @param request - EnablePolarClawCronJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnablePolarClawCronJobResponse
func (client *Client) EnablePolarClawCronJobWithOptions(request *EnablePolarClawCronJobRequest, runtime *dara.RuntimeOptions) (_result *EnablePolarClawCronJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnablePolarClawCronJob"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnablePolarClawCronJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启用PolarClaw定时任务
//
// @param request - EnablePolarClawCronJobRequest
//
// @return EnablePolarClawCronJobResponse
func (client *Client) EnablePolarClawCronJob(request *EnablePolarClawCronJobRequest) (_result *EnablePolarClawCronJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnablePolarClawCronJobResponse{}
	_body, _err := client.EnablePolarClawCronJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启用PolarClaw Plugin
//
// @param request - EnablePolarClawPluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnablePolarClawPluginResponse
func (client *Client) EnablePolarClawPluginWithOptions(request *EnablePolarClawPluginRequest, runtime *dara.RuntimeOptions) (_result *EnablePolarClawPluginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnablePolarClawPlugin"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnablePolarClawPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启用PolarClaw Plugin
//
// @param request - EnablePolarClawPluginRequest
//
// @return EnablePolarClawPluginResponse
func (client *Client) EnablePolarClawPlugin(request *EnablePolarClawPluginRequest) (_result *EnablePolarClawPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnablePolarClawPluginResponse{}
	_body, _err := client.EnablePolarClawPluginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启用/禁用SQL限流规则
//
// @param request - EnableSQLRateLimitingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSQLRateLimitingRulesResponse
func (client *Client) EnableSQLRateLimitingRulesWithOptions(request *EnableSQLRateLimitingRulesRequest, runtime *dara.RuntimeOptions) (_result *EnableSQLRateLimitingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
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

	if !dara.IsNil(request.RuleNameList) {
		query["RuleNameList"] = request.RuleNameList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSQLRateLimitingRules"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSQLRateLimitingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启用/禁用SQL限流规则
//
// @param request - EnableSQLRateLimitingRulesRequest
//
// @return EnableSQLRateLimitingRulesResponse
func (client *Client) EnableSQLRateLimitingRules(request *EnableSQLRateLimitingRulesRequest) (_result *EnableSQLRateLimitingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableSQLRateLimitingRulesResponse{}
	_body, _err := client.EnableSQLRateLimitingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Evaluates available resources.
//
// @param request - EvaluateRegionResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EvaluateRegionResourceResponse
func (client *Client) EvaluateRegionResourceWithOptions(request *EvaluateRegionResourceRequest, runtime *dara.RuntimeOptions) (_result *EvaluateRegionResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceConnType) {
		query["DBInstanceConnType"] = request.DBInstanceConnType
	}

	if !dara.IsNil(request.DBNodeClass) {
		query["DBNodeClass"] = request.DBNodeClass
	}

	if !dara.IsNil(request.DBType) {
		query["DBType"] = request.DBType
	}

	if !dara.IsNil(request.DBVersion) {
		query["DBVersion"] = request.DBVersion
	}

	if !dara.IsNil(request.DispenseMode) {
		query["DispenseMode"] = request.DispenseMode
	}

	if !dara.IsNil(request.NeedMaxScaleLink) {
		query["NeedMaxScaleLink"] = request.NeedMaxScaleLink
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

	if !dara.IsNil(request.SubDomain) {
		query["SubDomain"] = request.SubDomain
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EvaluateRegionResource"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EvaluateRegionResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Evaluates available resources.
//
// @param request - EvaluateRegionResourceRequest
//
// @return EvaluateRegionResourceResponse
func (client *Client) EvaluateRegionResource(request *EvaluateRegionResourceRequest) (_result *EvaluateRegionResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EvaluateRegionResourceResponse{}
	_body, _err := client.EvaluateRegionResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用跨云 OpenAPI
//
// @param request - ExecuteCrossCloudOpenAPIRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteCrossCloudOpenAPIResponse
func (client *Client) ExecuteCrossCloudOpenAPIWithOptions(request *ExecuteCrossCloudOpenAPIRequest, runtime *dara.RuntimeOptions) (_result *ExecuteCrossCloudOpenAPIResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProxyInfo) {
		query["ProxyInfo"] = request.ProxyInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteCrossCloudOpenAPI"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteCrossCloudOpenAPIResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用跨云 OpenAPI
//
// @param request - ExecuteCrossCloudOpenAPIRequest
//
// @return ExecuteCrossCloudOpenAPIResponse
func (client *Client) ExecuteCrossCloudOpenAPI(request *ExecuteCrossCloudOpenAPIRequest) (_result *ExecuteCrossCloudOpenAPIResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteCrossCloudOpenAPIResponse{}
	_body, _err := client.ExecuteCrossCloudOpenAPIWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a manual failover to promote a read-only node to the primary node in a PolarDB cluster.
//
// @param request - FailoverDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FailoverDBClusterResponse
func (client *Client) FailoverDBClusterWithOptions(request *FailoverDBClusterRequest, runtime *dara.RuntimeOptions) (_result *FailoverDBClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.RollBackForDisaster) {
		query["RollBackForDisaster"] = request.RollBackForDisaster
	}

	if !dara.IsNil(request.TargetDBNodeId) {
		query["TargetDBNodeId"] = request.TargetDBNodeId
	}

	if !dara.IsNil(request.TargetZoneType) {
		query["TargetZoneType"] = request.TargetZoneType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FailoverDBCluster"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FailoverDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a manual failover to promote a read-only node to the primary node in a PolarDB cluster.
//
// @param request - FailoverDBClusterRequest
//
// @return FailoverDBClusterResponse
func (client *Client) FailoverDBCluster(request *FailoverDBClusterRequest) (_result *FailoverDBClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FailoverDBClusterResponse{}
	_body, _err := client.FailoverDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 角色切换
//
// @param request - FailoverDBClusterZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FailoverDBClusterZonalResponse
func (client *Client) FailoverDBClusterZonalWithOptions(request *FailoverDBClusterZonalRequest, runtime *dara.RuntimeOptions) (_result *FailoverDBClusterZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.RollBackForDisaster) {
		query["RollBackForDisaster"] = request.RollBackForDisaster
	}

	if !dara.IsNil(request.TargetDBNodeId) {
		query["TargetDBNodeId"] = request.TargetDBNodeId
	}

	if !dara.IsNil(request.TargetZoneType) {
		query["TargetZoneType"] = request.TargetZoneType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FailoverDBClusterZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FailoverDBClusterZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 角色切换
//
// @param request - FailoverDBClusterZonalRequest
//
// @return FailoverDBClusterZonalResponse
func (client *Client) FailoverDBClusterZonal(request *FailoverDBClusterZonalRequest) (_result *FailoverDBClusterZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FailoverDBClusterZonalResponse{}
	_body, _err := client.FailoverDBClusterZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成校验报告
//
// @param request - GenerateUpgradeReportForSyncCloneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateUpgradeReportForSyncCloneResponse
func (client *Client) GenerateUpgradeReportForSyncCloneWithOptions(request *GenerateUpgradeReportForSyncCloneRequest, runtime *dara.RuntimeOptions) (_result *GenerateUpgradeReportForSyncCloneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreationCategory) {
		query["CreationCategory"] = request.CreationCategory
	}

	if !dara.IsNil(request.CreationOption) {
		query["CreationOption"] = request.CreationOption
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.DBType) {
		query["DBType"] = request.DBType
	}

	if !dara.IsNil(request.DBVersion) {
		query["DBVersion"] = request.DBVersion
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

	if !dara.IsNil(request.Reserve) {
		query["Reserve"] = request.Reserve
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SourceDBClusterId) {
		query["SourceDBClusterId"] = request.SourceDBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateUpgradeReportForSyncClone"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateUpgradeReportForSyncCloneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成校验报告
//
// @param request - GenerateUpgradeReportForSyncCloneRequest
//
// @return GenerateUpgradeReportForSyncCloneResponse
func (client *Client) GenerateUpgradeReportForSyncClone(request *GenerateUpgradeReportForSyncCloneRequest) (_result *GenerateUpgradeReportForSyncCloneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateUpgradeReportForSyncCloneResponse{}
	_body, _err := client.GenerateUpgradeReportForSyncCloneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建聊天记录
//
// @param request - GetPolarAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolarAgentResponse
func (client *Client) GetPolarAgentWithSSE(request *GetPolarAgentRequest, runtime *dara.RuntimeOptions, _yield chan *GetPolarAgentResponse, _yieldErr chan error) {
	defer close(_yield)
	client.getPolarAgentWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
	return
}

// Summary:
//
// 创建聊天记录
//
// @param request - GetPolarAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolarAgentResponse
func (client *Client) GetPolarAgentWithOptions(request *GetPolarAgentRequest, runtime *dara.RuntimeOptions) (_result *GetPolarAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExtraInfo) {
		query["ExtraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolarAgent"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPolarAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建聊天记录
//
// @param request - GetPolarAgentRequest
//
// @return GetPolarAgentResponse
func (client *Client) GetPolarAgent(request *GetPolarAgentRequest) (_result *GetPolarAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPolarAgentResponse{}
	_body, _err := client.GetPolarAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取PolarClaw配置
//
// @param request - GetPolarClawConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolarClawConfigResponse
func (client *Client) GetPolarClawConfigWithOptions(request *GetPolarClawConfigRequest, runtime *dara.RuntimeOptions) (_result *GetPolarClawConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ConfigPath) {
		query["ConfigPath"] = request.ConfigPath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolarClawConfig"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPolarClawConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取PolarClaw配置
//
// @param request - GetPolarClawConfigRequest
//
// @return GetPolarClawConfigResponse
func (client *Client) GetPolarClawConfig(request *GetPolarClawConfigRequest) (_result *GetPolarClawConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPolarClawConfigResponse{}
	_body, _err := client.GetPolarClawConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants a standard account the permissions to access one or more databases in a specified PolarDB cluster.
//
// Description:
//
// > 	- An account can be authorized to access one or more databases.
//
// > 	- If the specified account already has the access permissions on the specified databases, the operation returns a successful response.
//
// > 	- Before you call this operation, make sure that the cluster is in the Running state. Otherwise, the operation fails.
//
// > 	- You can call this operation only on a PolarDB for MySQL cluster.
//
// > 	- By default, a privileged account for a cluster has all the permissions on the databases in the cluster.
//
// @param request - GrantAccountPrivilegeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantAccountPrivilegeResponse
func (client *Client) GrantAccountPrivilegeWithOptions(request *GrantAccountPrivilegeRequest, runtime *dara.RuntimeOptions) (_result *GrantAccountPrivilegeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPrivilege) {
		query["AccountPrivilege"] = request.AccountPrivilege
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantAccountPrivilege"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantAccountPrivilegeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants a standard account the permissions to access one or more databases in a specified PolarDB cluster.
//
// Description:
//
// > 	- An account can be authorized to access one or more databases.
//
// > 	- If the specified account already has the access permissions on the specified databases, the operation returns a successful response.
//
// > 	- Before you call this operation, make sure that the cluster is in the Running state. Otherwise, the operation fails.
//
// > 	- You can call this operation only on a PolarDB for MySQL cluster.
//
// > 	- By default, a privileged account for a cluster has all the permissions on the databases in the cluster.
//
// @param request - GrantAccountPrivilegeRequest
//
// @return GrantAccountPrivilegeResponse
func (client *Client) GrantAccountPrivilege(request *GrantAccountPrivilegeRequest) (_result *GrantAccountPrivilegeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GrantAccountPrivilegeResponse{}
	_body, _err := client.GrantAccountPrivilegeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改PolarDB边缘集群的账号权限
//
// @param request - GrantAccountPrivilegeZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantAccountPrivilegeZonalResponse
func (client *Client) GrantAccountPrivilegeZonalWithOptions(request *GrantAccountPrivilegeZonalRequest, runtime *dara.RuntimeOptions) (_result *GrantAccountPrivilegeZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPrivilege) {
		query["AccountPrivilege"] = request.AccountPrivilege
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantAccountPrivilegeZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantAccountPrivilegeZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改PolarDB边缘集群的账号权限
//
// @param request - GrantAccountPrivilegeZonalRequest
//
// @return GrantAccountPrivilegeZonalResponse
func (client *Client) GrantAccountPrivilegeZonal(request *GrantAccountPrivilegeZonalRequest) (_result *GrantAccountPrivilegeZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GrantAccountPrivilegeZonalResponse{}
	_body, _err := client.GrantAccountPrivilegeZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 安装PolarClaw Plugin
//
// @param request - InstallPolarClawPluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallPolarClawPluginResponse
func (client *Client) InstallPolarClawPluginWithOptions(request *InstallPolarClawPluginRequest, runtime *dara.RuntimeOptions) (_result *InstallPolarClawPluginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.NpmPackage) {
		query["NpmPackage"] = request.NpmPackage
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallPolarClawPlugin"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallPolarClawPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 安装PolarClaw Plugin
//
// @param request - InstallPolarClawPluginRequest
//
// @return InstallPolarClawPluginResponse
func (client *Client) InstallPolarClawPlugin(request *InstallPolarClawPluginRequest) (_result *InstallPolarClawPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InstallPolarClawPluginResponse{}
	_body, _err := client.InstallPolarClawPluginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询订单
//
// @param request - ListOrdersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrdersResponse
func (client *Client) ListOrdersWithOptions(request *ListOrdersRequest, runtime *dara.RuntimeOptions) (_result *ListOrdersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderStatus) {
		query["OrderStatus"] = request.OrderStatus
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
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
		Action:      dara.String("ListOrders"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询订单
//
// @param request - ListOrdersRequest
//
// @return ListOrdersResponse
func (client *Client) ListOrders(request *ListOrdersRequest) (_result *ListOrdersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOrdersResponse{}
	_body, _err := client.ListOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询PolarClaw绑定关系列表
//
// @param tmpReq - ListPolarClawBindingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolarClawBindingsResponse
func (client *Client) ListPolarClawBindingsWithOptions(tmpReq *ListPolarClawBindingsRequest, runtime *dara.RuntimeOptions) (_result *ListPolarClawBindingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPolarClawBindingsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentList) {
		request.AgentListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentList, dara.String("AgentList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentListShrink) {
		query["AgentList"] = request.AgentListShrink
	}

	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolarClawBindings"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPolarClawBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询PolarClaw绑定关系列表
//
// @param request - ListPolarClawBindingsRequest
//
// @return ListPolarClawBindingsResponse
func (client *Client) ListPolarClawBindings(request *ListPolarClawBindingsRequest) (_result *ListPolarClawBindingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPolarClawBindingsResponse{}
	_body, _err := client.ListPolarClawBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to one or more PolarDB clusters, or the PolarDB clusters to which one or more tags are added.
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
		Version:     dara.String("2017-08-01"),
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
// Queries the tags that are added to one or more PolarDB clusters, or the PolarDB clusters to which one or more tags are added.
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

// @param request - ListTagResourcesForRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesForRegionResponse
func (client *Client) ListTagResourcesForRegionWithOptions(request *ListTagResourcesForRegionRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesForRegionResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResourcesForRegion"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesForRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListTagResourcesForRegionRequest
//
// @return ListTagResourcesForRegionResponse
func (client *Client) ListTagResourcesForRegion(request *ListTagResourcesForRegionRequest) (_result *ListTagResourcesForRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagResourcesForRegionResponse{}
	_body, _err := client.ListTagResourcesForRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Manually starts a cluster.
//
// @param request - ManuallyStartDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManuallyStartDBClusterResponse
func (client *Client) ManuallyStartDBClusterWithOptions(request *ManuallyStartDBClusterRequest, runtime *dara.RuntimeOptions) (_result *ManuallyStartDBClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("ManuallyStartDBCluster"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ManuallyStartDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually starts a cluster.
//
// @param request - ManuallyStartDBClusterRequest
//
// @return ManuallyStartDBClusterResponse
func (client *Client) ManuallyStartDBCluster(request *ManuallyStartDBClusterRequest) (_result *ManuallyStartDBClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ManuallyStartDBClusterResponse{}
	_body, _err := client.ManuallyStartDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改AI实例名称
//
// @param request - ModifyAIDBClusterDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAIDBClusterDescriptionResponse
func (client *Client) ModifyAIDBClusterDescriptionWithOptions(request *ModifyAIDBClusterDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyAIDBClusterDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterDescription) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAIDBClusterDescription"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAIDBClusterDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改AI实例名称
//
// @param request - ModifyAIDBClusterDescriptionRequest
//
// @return ModifyAIDBClusterDescriptionResponse
func (client *Client) ModifyAIDBClusterDescription(request *ModifyAIDBClusterDescriptionRequest) (_result *ModifyAIDBClusterDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAIDBClusterDescriptionResponse{}
	_body, _err := client.ModifyAIDBClusterDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of a database account of a PolarDB cluster.
//
// @param request - ModifyAccountDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountDescriptionResponse
func (client *Client) ModifyAccountDescriptionWithOptions(request *ModifyAccountDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountDescription) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountDescription"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a database account of a PolarDB cluster.
//
// @param request - ModifyAccountDescriptionRequest
//
// @return ModifyAccountDescriptionResponse
func (client *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (_result *ModifyAccountDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.ModifyAccountDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改PolarDB边缘云账号的描述
//
// @param request - ModifyAccountDescriptionZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountDescriptionZonalResponse
func (client *Client) ModifyAccountDescriptionZonalWithOptions(request *ModifyAccountDescriptionZonalRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountDescriptionZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountDescription) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountDescriptionZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountDescriptionZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改PolarDB边缘云账号的描述
//
// @param request - ModifyAccountDescriptionZonalRequest
//
// @return ModifyAccountDescriptionZonalResponse
func (client *Client) ModifyAccountDescriptionZonal(request *ModifyAccountDescriptionZonalRequest) (_result *ModifyAccountDescriptionZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAccountDescriptionZonalResponse{}
	_body, _err := client.ModifyAccountDescriptionZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyAccountLockStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountLockStateResponse
func (client *Client) ModifyAccountLockStateWithOptions(request *ModifyAccountLockStateRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountLockStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountLockState) {
		query["AccountLockState"] = request.AccountLockState
	}

	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPasswordValidTime) {
		query["AccountPasswordValidTime"] = request.AccountPasswordValidTime
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountLockState"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountLockStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyAccountLockStateRequest
//
// @return ModifyAccountLockStateResponse
func (client *Client) ModifyAccountLockState(request *ModifyAccountLockStateRequest) (_result *ModifyAccountLockStateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAccountLockStateResponse{}
	_body, _err := client.ModifyAccountLockStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the password of a database account for a specified PolarDB cluster.
//
// @param request - ModifyAccountPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountPasswordResponse
func (client *Client) ModifyAccountPasswordWithOptions(request *ModifyAccountPasswordRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.NewAccountPassword) {
		query["NewAccountPassword"] = request.NewAccountPassword
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PasswordType) {
		query["PasswordType"] = request.PasswordType
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
		Action:      dara.String("ModifyAccountPassword"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the password of a database account for a specified PolarDB cluster.
//
// @param request - ModifyAccountPasswordRequest
//
// @return ModifyAccountPasswordResponse
func (client *Client) ModifyAccountPassword(request *ModifyAccountPasswordRequest) (_result *ModifyAccountPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAccountPasswordResponse{}
	_body, _err := client.ModifyAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改PolarDB边缘的账号密码
//
// @param request - ModifyAccountPasswordZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountPasswordZonalResponse
func (client *Client) ModifyAccountPasswordZonalWithOptions(request *ModifyAccountPasswordZonalRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountPasswordZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.NewAccountPassword) {
		query["NewAccountPassword"] = request.NewAccountPassword
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PasswordType) {
		query["PasswordType"] = request.PasswordType
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
		Action:      dara.String("ModifyAccountPasswordZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountPasswordZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改PolarDB边缘的账号密码
//
// @param request - ModifyAccountPasswordZonalRequest
//
// @return ModifyAccountPasswordZonalResponse
func (client *Client) ModifyAccountPasswordZonal(request *ModifyAccountPasswordZonalRequest) (_result *ModifyAccountPasswordZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAccountPasswordZonalResponse{}
	_body, _err := client.ModifyAccountPasswordZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置用户的运维信息，目前包括主动运维窗口信息
//
// @param request - ModifyActiveOperationMaintainConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyActiveOperationMaintainConfResponse
func (client *Client) ModifyActiveOperationMaintainConfWithOptions(request *ModifyActiveOperationMaintainConfRequest, runtime *dara.RuntimeOptions) (_result *ModifyActiveOperationMaintainConfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.CycleTime) {
		query["CycleTime"] = request.CycleTime
	}

	if !dara.IsNil(request.CycleType) {
		query["CycleType"] = request.CycleType
	}

	if !dara.IsNil(request.MaintainEndTime) {
		query["MaintainEndTime"] = request.MaintainEndTime
	}

	if !dara.IsNil(request.MaintainStartTime) {
		query["MaintainStartTime"] = request.MaintainStartTime
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyActiveOperationMaintainConf"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyActiveOperationMaintainConfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置用户的运维信息，目前包括主动运维窗口信息
//
// @param request - ModifyActiveOperationMaintainConfRequest
//
// @return ModifyActiveOperationMaintainConfResponse
func (client *Client) ModifyActiveOperationMaintainConf(request *ModifyActiveOperationMaintainConfRequest) (_result *ModifyActiveOperationMaintainConfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyActiveOperationMaintainConfResponse{}
	_body, _err := client.ModifyActiveOperationMaintainConfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the switching time of scheduled O\\\\\\&M events for an instance.
//
// @param request - ModifyActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyActiveOperationTasksResponse
func (client *Client) ModifyActiveOperationTasksWithOptions(request *ModifyActiveOperationTasksRequest, runtime *dara.RuntimeOptions) (_result *ModifyActiveOperationTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImmediateStart) {
		query["ImmediateStart"] = request.ImmediateStart
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !dara.IsNil(request.TaskIds) {
		query["TaskIds"] = request.TaskIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyActiveOperationTasks"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyActiveOperationTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the switching time of scheduled O\\\\\\&M events for an instance.
//
// @param request - ModifyActiveOperationTasksRequest
//
// @return ModifyActiveOperationTasksResponse
func (client *Client) ModifyActiveOperationTasks(request *ModifyActiveOperationTasksRequest) (_result *ModifyActiveOperationTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyActiveOperationTasksResponse{}
	_body, _err := client.ModifyActiveOperationTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改应用描述
//
// @param request - ModifyApplicationDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApplicationDescriptionResponse
func (client *Client) ModifyApplicationDescriptionWithOptions(request *ModifyApplicationDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyApplicationDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApplicationDescription"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApplicationDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用描述
//
// @param request - ModifyApplicationDescriptionRequest
//
// @return ModifyApplicationDescriptionResponse
func (client *Client) ModifyApplicationDescription(request *ModifyApplicationDescriptionRequest) (_result *ModifyApplicationDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApplicationDescriptionResponse{}
	_body, _err := client.ModifyApplicationDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改PolarDB应用参数
//
// @param tmpReq - ModifyApplicationParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApplicationParameterResponse
func (client *Client) ModifyApplicationParameterWithOptions(tmpReq *ModifyApplicationParameterRequest, runtime *dara.RuntimeOptions) (_result *ModifyApplicationParameterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyApplicationParameterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ParameterName) {
		query["ParameterName"] = request.ParameterName
	}

	if !dara.IsNil(request.ParameterValue) {
		query["ParameterValue"] = request.ParameterValue
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApplicationParameter"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApplicationParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改PolarDB应用参数
//
// @param request - ModifyApplicationParameterRequest
//
// @return ModifyApplicationParameterResponse
func (client *Client) ModifyApplicationParameter(request *ModifyApplicationParameterRequest) (_result *ModifyApplicationParameterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApplicationParameterResponse{}
	_body, _err := client.ModifyApplicationParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改应用提示词策略
//
// @param request - ModifyApplicationPromptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApplicationPromptResponse
func (client *Client) ModifyApplicationPromptWithOptions(request *ModifyApplicationPromptRequest, runtime *dara.RuntimeOptions) (_result *ModifyApplicationPromptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.PromptId) {
		query["PromptId"] = request.PromptId
	}

	if !dara.IsNil(request.PromptName) {
		query["PromptName"] = request.PromptName
	}

	if !dara.IsNil(request.PromptValue) {
		query["PromptValue"] = request.PromptValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApplicationPrompt"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApplicationPromptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用提示词策略
//
// @param request - ModifyApplicationPromptRequest
//
// @return ModifyApplicationPromptResponse
func (client *Client) ModifyApplicationPrompt(request *ModifyApplicationPromptRequest) (_result *ModifyApplicationPromptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApplicationPromptResponse{}
	_body, _err := client.ModifyApplicationPromptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改应用serverless配置
//
// @param request - ModifyApplicationServerlessConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApplicationServerlessConfResponse
func (client *Client) ModifyApplicationServerlessConfWithOptions(request *ModifyApplicationServerlessConfRequest, runtime *dara.RuntimeOptions) (_result *ModifyApplicationServerlessConfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ServerlessConfList) {
		query["ServerlessConfList"] = request.ServerlessConfList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApplicationServerlessConf"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApplicationServerlessConfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用serverless配置
//
// @param request - ModifyApplicationServerlessConfRequest
//
// @return ModifyApplicationServerlessConfResponse
func (client *Client) ModifyApplicationServerlessConf(request *ModifyApplicationServerlessConfRequest) (_result *ModifyApplicationServerlessConfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApplicationServerlessConfResponse{}
	_body, _err := client.ModifyApplicationServerlessConfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改应用白名单
//
// @param request - ModifyApplicationWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApplicationWhitelistResponse
func (client *Client) ModifyApplicationWhitelistWithOptions(request *ModifyApplicationWhitelistRequest, runtime *dara.RuntimeOptions) (_result *ModifyApplicationWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ComponentId) {
		query["ComponentId"] = request.ComponentId
	}

	if !dara.IsNil(request.ModifyMode) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !dara.IsNil(request.SecurityGroups) {
		query["SecurityGroups"] = request.SecurityGroups
	}

	if !dara.IsNil(request.SecurityIPArrayName) {
		query["SecurityIPArrayName"] = request.SecurityIPArrayName
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApplicationWhitelist"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApplicationWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用白名单
//
// @param request - ModifyApplicationWhitelistRequest
//
// @return ModifyApplicationWhitelistResponse
func (client *Client) ModifyApplicationWhitelist(request *ModifyApplicationWhitelistRequest) (_result *ModifyApplicationWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApplicationWhitelistResponse{}
	_body, _err := client.ModifyApplicationWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the auto-renewal attributes of a subscription PolarDB cluster.
//
// @param request - ModifyAutoRenewAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAutoRenewAttributeResponse
func (client *Client) ModifyAutoRenewAttributeWithOptions(request *ModifyAutoRenewAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyAutoRenewAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloudProvider) {
		query["CloudProvider"] = request.CloudProvider
	}

	if !dara.IsNil(request.DBClusterIds) {
		query["DBClusterIds"] = request.DBClusterIds
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RenewalStatus) {
		query["RenewalStatus"] = request.RenewalStatus
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAutoRenewAttribute"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAutoRenewAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the auto-renewal attributes of a subscription PolarDB cluster.
//
// @param request - ModifyAutoRenewAttributeRequest
//
// @return ModifyAutoRenewAttributeResponse
func (client *Client) ModifyAutoRenewAttribute(request *ModifyAutoRenewAttributeRequest) (_result *ModifyAutoRenewAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAutoRenewAttributeResponse{}
	_body, _err := client.ModifyAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the automatic backup policy of a PolarDB cluster.
//
// Description:
//
// > You can also modify the automatic backup policy of a PolarDB cluster in the console. For more information, see [Backup settings](https://help.aliyun.com/document_detail/280422.html).
//
// @param tmpReq - ModifyBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupPolicyResponse
func (client *Client) ModifyBackupPolicyWithOptions(tmpReq *ModifyBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyBackupPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AdvancedDataPolicies) {
		request.AdvancedDataPoliciesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdvancedDataPolicies, dara.String("AdvancedDataPolicies"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdvancedDataPoliciesShrink) {
		query["AdvancedDataPolicies"] = request.AdvancedDataPoliciesShrink
	}

	if !dara.IsNil(request.BackupFrequency) {
		query["BackupFrequency"] = request.BackupFrequency
	}

	if !dara.IsNil(request.BackupPolicyLevel) {
		query["BackupPolicyLevel"] = request.BackupPolicyLevel
	}

	if !dara.IsNil(request.BackupRetentionPolicyOnClusterDeletion) {
		query["BackupRetentionPolicyOnClusterDeletion"] = request.BackupRetentionPolicyOnClusterDeletion
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DataLevel1BackupFrequency) {
		query["DataLevel1BackupFrequency"] = request.DataLevel1BackupFrequency
	}

	if !dara.IsNil(request.DataLevel1BackupPeriod) {
		query["DataLevel1BackupPeriod"] = request.DataLevel1BackupPeriod
	}

	if !dara.IsNil(request.DataLevel1BackupRetentionPeriod) {
		query["DataLevel1BackupRetentionPeriod"] = request.DataLevel1BackupRetentionPeriod
	}

	if !dara.IsNil(request.DataLevel1BackupTime) {
		query["DataLevel1BackupTime"] = request.DataLevel1BackupTime
	}

	if !dara.IsNil(request.DataLevel2BackupAnotherRegionRegion) {
		query["DataLevel2BackupAnotherRegionRegion"] = request.DataLevel2BackupAnotherRegionRegion
	}

	if !dara.IsNil(request.DataLevel2BackupAnotherRegionRetentionPeriod) {
		query["DataLevel2BackupAnotherRegionRetentionPeriod"] = request.DataLevel2BackupAnotherRegionRetentionPeriod
	}

	if !dara.IsNil(request.DataLevel2BackupPeriod) {
		query["DataLevel2BackupPeriod"] = request.DataLevel2BackupPeriod
	}

	if !dara.IsNil(request.DataLevel2BackupRetentionPeriod) {
		query["DataLevel2BackupRetentionPeriod"] = request.DataLevel2BackupRetentionPeriod
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PreferredBackupPeriod) {
		query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	}

	if !dara.IsNil(request.PreferredBackupTime) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
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
		Action:      dara.String("ModifyBackupPolicy"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the automatic backup policy of a PolarDB cluster.
//
// Description:
//
// > You can also modify the automatic backup policy of a PolarDB cluster in the console. For more information, see [Backup settings](https://help.aliyun.com/document_detail/280422.html).
//
// @param request - ModifyBackupPolicyRequest
//
// @return ModifyBackupPolicyResponse
func (client *Client) ModifyBackupPolicy(request *ModifyBackupPolicyRequest) (_result *ModifyBackupPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.ModifyBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改预算策略
//
// @param request - ModifyBudgetPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBudgetPolicyResponse
func (client *Client) ModifyBudgetPolicyWithOptions(request *ModifyBudgetPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyBudgetPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertThresholdPct) {
		query["AlertThresholdPct"] = request.AlertThresholdPct
	}

	if !dara.IsNil(request.BudgetPoints) {
		query["BudgetPoints"] = request.BudgetPoints
	}

	if !dara.IsNil(request.BudgetPolicyId) {
		query["BudgetPolicyId"] = request.BudgetPolicyId
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResetDayOfMonth) {
		query["ResetDayOfMonth"] = request.ResetDayOfMonth
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBudgetPolicy"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBudgetPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改预算策略
//
// @param request - ModifyBudgetPolicyRequest
//
// @return ModifyBudgetPolicyResponse
func (client *Client) ModifyBudgetPolicy(request *ModifyBudgetPolicyRequest) (_result *ModifyBudgetPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBudgetPolicyResponse{}
	_body, _err := client.ModifyBudgetPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改消费者
//
// @param request - ModifyConsumerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyConsumerResponse
func (client *Client) ModifyConsumerWithOptions(request *ModifyConsumerRequest, runtime *dara.RuntimeOptions) (_result *ModifyConsumerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerGroupName) {
		query["ConsumerGroupName"] = request.ConsumerGroupName
	}

	if !dara.IsNil(request.ConsumerId) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.IsDefault) {
		query["IsDefault"] = request.IsDefault
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyConsumer"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyConsumerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改消费者
//
// @param request - ModifyConsumerRequest
//
// @return ModifyConsumerResponse
func (client *Client) ModifyConsumer(request *ModifyConsumerRequest) (_result *ModifyConsumerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyConsumerResponse{}
	_body, _err := client.ModifyConsumerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改消费者组
//
// @param request - ModifyConsumerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyConsumerGroupResponse
func (client *Client) ModifyConsumerGroupWithOptions(request *ModifyConsumerGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyConsumerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerGroupName) {
		query["ConsumerGroupName"] = request.ConsumerGroupName
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.IsDefault) {
		query["IsDefault"] = request.IsDefault
	}

	if !dara.IsNil(request.NickName) {
		query["NickName"] = request.NickName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyConsumerGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改消费者组
//
// @param request - ModifyConsumerGroupRequest
//
// @return ModifyConsumerGroupResponse
func (client *Client) ModifyConsumerGroup(request *ModifyConsumerGroupRequest) (_result *ModifyConsumerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyConsumerGroupResponse{}
	_body, _err := client.ModifyConsumerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改限流策略
//
// @param request - ModifyCostRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCostRuleResponse
func (client *Client) ModifyCostRuleWithOptions(request *ModifyCostRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyCostRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CacheCostPointsPerMillion) {
		query["CacheCostPointsPerMillion"] = request.CacheCostPointsPerMillion
	}

	if !dara.IsNil(request.CostRuleId) {
		query["CostRuleId"] = request.CostRuleId
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.InputCostPointsPerMillion) {
		query["InputCostPointsPerMillion"] = request.InputCostPointsPerMillion
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.ModelServiceId) {
		query["ModelServiceId"] = request.ModelServiceId
	}

	if !dara.IsNil(request.OutputCostPointsPerMillion) {
		query["OutputCostPointsPerMillion"] = request.OutputCostPointsPerMillion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCostRule"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCostRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改限流策略
//
// @param request - ModifyCostRuleRequest
//
// @return ModifyCostRuleResponse
func (client *Client) ModifyCostRule(request *ModifyCostRuleRequest) (_result *ModifyCostRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCostRuleResponse{}
	_body, _err := client.ModifyCostRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改周期任务策略
//
// @param request - ModifyCronJobPolicyServerlessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCronJobPolicyServerlessResponse
func (client *Client) ModifyCronJobPolicyServerlessWithOptions(request *ModifyCronJobPolicyServerlessRequest, runtime *dara.RuntimeOptions) (_result *ModifyCronJobPolicyServerlessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowShutDown) {
		query["AllowShutDown"] = request.AllowShutDown
	}

	if !dara.IsNil(request.CronExpression) {
		query["CronExpression"] = request.CronExpression
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
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

	if !dara.IsNil(request.ScaleApRoNumMax) {
		query["ScaleApRoNumMax"] = request.ScaleApRoNumMax
	}

	if !dara.IsNil(request.ScaleApRoNumMin) {
		query["ScaleApRoNumMin"] = request.ScaleApRoNumMin
	}

	if !dara.IsNil(request.ScaleMax) {
		query["ScaleMax"] = request.ScaleMax
	}

	if !dara.IsNil(request.ScaleMin) {
		query["ScaleMin"] = request.ScaleMin
	}

	if !dara.IsNil(request.ScaleRoNumMax) {
		query["ScaleRoNumMax"] = request.ScaleRoNumMax
	}

	if !dara.IsNil(request.ScaleRoNumMin) {
		query["ScaleRoNumMin"] = request.ScaleRoNumMin
	}

	if !dara.IsNil(request.SecondsUntilAutoPause) {
		query["SecondsUntilAutoPause"] = request.SecondsUntilAutoPause
	}

	if !dara.IsNil(request.ServerlessRuleCpuEnlargeThreshold) {
		query["ServerlessRuleCpuEnlargeThreshold"] = request.ServerlessRuleCpuEnlargeThreshold
	}

	if !dara.IsNil(request.ServerlessRuleCpuShrinkThreshold) {
		query["ServerlessRuleCpuShrinkThreshold"] = request.ServerlessRuleCpuShrinkThreshold
	}

	if !dara.IsNil(request.ServerlessRuleMode) {
		query["ServerlessRuleMode"] = request.ServerlessRuleMode
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCronJobPolicyServerless"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCronJobPolicyServerlessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改周期任务策略
//
// @param request - ModifyCronJobPolicyServerlessRequest
//
// @return ModifyCronJobPolicyServerlessResponse
func (client *Client) ModifyCronJobPolicyServerless(request *ModifyCronJobPolicyServerlessRequest) (_result *ModifyCronJobPolicyServerlessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCronJobPolicyServerlessResponse{}
	_body, _err := client.ModifyCronJobPolicyServerlessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a PolarDB for MySQL cluster.
//
// @param request - ModifyDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterResponse
func (client *Client) ModifyDBClusterWithOptions(request *ModifyDBClusterRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompressStorage) {
		query["CompressStorage"] = request.CompressStorage
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNodeCrashList) {
		query["DBNodeCrashList"] = request.DBNodeCrashList
	}

	if !dara.IsNil(request.DataSyncMode) {
		query["DataSyncMode"] = request.DataSyncMode
	}

	if !dara.IsNil(request.FaultInjectionType) {
		query["FaultInjectionType"] = request.FaultInjectionType
	}

	if !dara.IsNil(request.FaultSimulateMode) {
		query["FaultSimulateMode"] = request.FaultSimulateMode
	}

	if !dara.IsNil(request.ImciAutoIndex) {
		query["ImciAutoIndex"] = request.ImciAutoIndex
	}

	if !dara.IsNil(request.ModifyRowCompression) {
		query["ModifyRowCompression"] = request.ModifyRowCompression
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

	if !dara.IsNil(request.StandbyHAMode) {
		query["StandbyHAMode"] = request.StandbyHAMode
	}

	if !dara.IsNil(request.StorageAutoScale) {
		query["StorageAutoScale"] = request.StorageAutoScale
	}

	if !dara.IsNil(request.StorageUpperBound) {
		query["StorageUpperBound"] = request.StorageUpperBound
	}

	if !dara.IsNil(request.TableMeta) {
		query["TableMeta"] = request.TableMeta
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBCluster"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a PolarDB for MySQL cluster.
//
// @param request - ModifyDBClusterRequest
//
// @return ModifyDBClusterResponse
func (client *Client) ModifyDBCluster(request *ModifyDBClusterRequest) (_result *ModifyDBClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterResponse{}
	_body, _err := client.ModifyDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or modifies the whitelists (IP whitelists and security groups) of a specified cluster.
//
// @param request - ModifyDBClusterAccessWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterAccessWhitelistResponse
func (client *Client) ModifyDBClusterAccessWhitelistWithOptions(request *ModifyDBClusterAccessWhitelistRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterAccessWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterIPArrayAttribute) {
		query["DBClusterIPArrayAttribute"] = request.DBClusterIPArrayAttribute
	}

	if !dara.IsNil(request.DBClusterIPArrayName) {
		query["DBClusterIPArrayName"] = request.DBClusterIPArrayName
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ModifyMode) {
		query["ModifyMode"] = request.ModifyMode
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

	if !dara.IsNil(request.SecurityGroupIds) {
		query["SecurityGroupIds"] = request.SecurityGroupIds
	}

	if !dara.IsNil(request.SecurityIps) {
		query["SecurityIps"] = request.SecurityIps
	}

	if !dara.IsNil(request.WhiteListType) {
		query["WhiteListType"] = request.WhiteListType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterAccessWhitelist"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterAccessWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies the whitelists (IP whitelists and security groups) of a specified cluster.
//
// @param request - ModifyDBClusterAccessWhitelistRequest
//
// @return ModifyDBClusterAccessWhitelistResponse
func (client *Client) ModifyDBClusterAccessWhitelist(request *ModifyDBClusterAccessWhitelistRequest) (_result *ModifyDBClusterAccessWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterAccessWhitelistResponse{}
	_body, _err := client.ModifyDBClusterAccessWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies cluster parameters and applies them to specified nodes.
//
// @param request - ModifyDBClusterAndNodesParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterAndNodesParametersResponse
func (client *Client) ModifyDBClusterAndNodesParametersWithOptions(request *ModifyDBClusterAndNodesParametersRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterAndNodesParametersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClearBinlog) {
		query["ClearBinlog"] = request.ClearBinlog
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNodeIds) {
		query["DBNodeIds"] = request.DBNodeIds
	}

	if !dara.IsNil(request.FromTimeService) {
		query["FromTimeService"] = request.FromTimeService
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParameterGroupId) {
		query["ParameterGroupId"] = request.ParameterGroupId
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.PlannedEndTime) {
		query["PlannedEndTime"] = request.PlannedEndTime
	}

	if !dara.IsNil(request.PlannedStartTime) {
		query["PlannedStartTime"] = request.PlannedStartTime
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StandbyClusterIdListNeedToSync) {
		query["StandbyClusterIdListNeedToSync"] = request.StandbyClusterIdListNeedToSync
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterAndNodesParameters"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterAndNodesParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies cluster parameters and applies them to specified nodes.
//
// @param request - ModifyDBClusterAndNodesParametersRequest
//
// @return ModifyDBClusterAndNodesParametersResponse
func (client *Client) ModifyDBClusterAndNodesParameters(request *ModifyDBClusterAndNodesParametersRequest) (_result *ModifyDBClusterAndNodesParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterAndNodesParametersResponse{}
	_body, _err := client.ModifyDBClusterAndNodesParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the high availability mode of the cluster.
//
// @param request - ModifyDBClusterArchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterArchResponse
func (client *Client) ModifyDBClusterArchWithOptions(request *ModifyDBClusterArchRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterArchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.HotStandbyCluster) {
		query["HotStandbyCluster"] = request.HotStandbyCluster
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StandbyAZ) {
		query["StandbyAZ"] = request.StandbyAZ
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterArch"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterArchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the high availability mode of the cluster.
//
// @param request - ModifyDBClusterArchRequest
//
// @return ModifyDBClusterArchResponse
func (client *Client) ModifyDBClusterArch(request *ModifyDBClusterArchRequest) (_result *ModifyDBClusterArchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterArchResponse{}
	_body, _err := client.ModifyDBClusterArchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables SQL collector for a PolarDB cluster. The features related to SQL collector include Audit Logs and SQL Explorer.
//
// @param request - ModifyDBClusterAuditLogCollectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterAuditLogCollectorResponse
func (client *Client) ModifyDBClusterAuditLogCollectorWithOptions(request *ModifyDBClusterAuditLogCollectorRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterAuditLogCollectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CollectorStatus) {
		query["CollectorStatus"] = request.CollectorStatus
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterAuditLogCollector"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterAuditLogCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables SQL collector for a PolarDB cluster. The features related to SQL collector include Audit Logs and SQL Explorer.
//
// @param request - ModifyDBClusterAuditLogCollectorRequest
//
// @return ModifyDBClusterAuditLogCollectorResponse
func (client *Client) ModifyDBClusterAuditLogCollector(request *ModifyDBClusterAuditLogCollectorRequest) (_result *ModifyDBClusterAuditLogCollectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterAuditLogCollectorResponse{}
	_body, _err := client.ModifyDBClusterAuditLogCollectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the cluster lock feature for a PolarDB cluster.
//
// @param request - ModifyDBClusterDeletionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterDeletionResponse
func (client *Client) ModifyDBClusterDeletionWithOptions(request *ModifyDBClusterDeletionRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterDeletionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Protection) {
		query["Protection"] = request.Protection
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
		Action:      dara.String("ModifyDBClusterDeletion"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterDeletionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the cluster lock feature for a PolarDB cluster.
//
// @param request - ModifyDBClusterDeletionRequest
//
// @return ModifyDBClusterDeletionResponse
func (client *Client) ModifyDBClusterDeletion(request *ModifyDBClusterDeletionRequest) (_result *ModifyDBClusterDeletionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterDeletionResponse{}
	_body, _err := client.ModifyDBClusterDeletionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name of a PolarDB cluster.
//
// @param request - ModifyDBClusterDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterDescriptionResponse
func (client *Client) ModifyDBClusterDescriptionWithOptions(request *ModifyDBClusterDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterDescription) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterDescription"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of a PolarDB cluster.
//
// @param request - ModifyDBClusterDescriptionRequest
//
// @return ModifyDBClusterDescriptionResponse
func (client *Client) ModifyDBClusterDescription(request *ModifyDBClusterDescriptionRequest) (_result *ModifyDBClusterDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterDescriptionResponse{}
	_body, _err := client.ModifyDBClusterDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改集群描述
//
// @param request - ModifyDBClusterDescriptionZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterDescriptionZonalResponse
func (client *Client) ModifyDBClusterDescriptionZonalWithOptions(request *ModifyDBClusterDescriptionZonalRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterDescriptionZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterDescription) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterDescriptionZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterDescriptionZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改集群描述
//
// @param request - ModifyDBClusterDescriptionZonalRequest
//
// @return ModifyDBClusterDescriptionZonalResponse
func (client *Client) ModifyDBClusterDescriptionZonal(request *ModifyDBClusterDescriptionZonalRequest) (_result *ModifyDBClusterDescriptionZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterDescriptionZonalResponse{}
	_body, _err := client.ModifyDBClusterDescriptionZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a specified PolarDB cluster endpoint. For example, you can modify the following attributes for the specified cluster endpoint: read/write mode, consistency level, transaction splitting, primary node accepts read requests, and connection pool. You can also call the operation to specify whether newly added nodes are automatically associated with the specified cluster endpoint.
//
// @param request - ModifyDBClusterEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterEndpointResponse
func (client *Client) ModifyDBClusterEndpointWithOptions(request *ModifyDBClusterEndpointRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoAddNewNodes) {
		query["AutoAddNewNodes"] = request.AutoAddNewNodes
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBEndpointDescription) {
		query["DBEndpointDescription"] = request.DBEndpointDescription
	}

	if !dara.IsNil(request.DBEndpointId) {
		query["DBEndpointId"] = request.DBEndpointId
	}

	if !dara.IsNil(request.EndpointConfig) {
		query["EndpointConfig"] = request.EndpointConfig
	}

	if !dara.IsNil(request.Nodes) {
		query["Nodes"] = request.Nodes
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PolarSccTimeoutAction) {
		query["PolarSccTimeoutAction"] = request.PolarSccTimeoutAction
	}

	if !dara.IsNil(request.PolarSccWaitTimeout) {
		query["PolarSccWaitTimeout"] = request.PolarSccWaitTimeout
	}

	if !dara.IsNil(request.ReadWriteMode) {
		query["ReadWriteMode"] = request.ReadWriteMode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SccMode) {
		query["SccMode"] = request.SccMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterEndpoint"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a specified PolarDB cluster endpoint. For example, you can modify the following attributes for the specified cluster endpoint: read/write mode, consistency level, transaction splitting, primary node accepts read requests, and connection pool. You can also call the operation to specify whether newly added nodes are automatically associated with the specified cluster endpoint.
//
// @param request - ModifyDBClusterEndpointRequest
//
// @return ModifyDBClusterEndpointResponse
func (client *Client) ModifyDBClusterEndpoint(request *ModifyDBClusterEndpointRequest) (_result *ModifyDBClusterEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterEndpointResponse{}
	_body, _err := client.ModifyDBClusterEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # PolarDB边缘集群修改链接地址
//
// @param request - ModifyDBClusterEndpointZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterEndpointZonalResponse
func (client *Client) ModifyDBClusterEndpointZonalWithOptions(request *ModifyDBClusterEndpointZonalRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterEndpointZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoAddNewNodes) {
		query["AutoAddNewNodes"] = request.AutoAddNewNodes
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBEndpointDescription) {
		query["DBEndpointDescription"] = request.DBEndpointDescription
	}

	if !dara.IsNil(request.DBEndpointId) {
		query["DBEndpointId"] = request.DBEndpointId
	}

	if !dara.IsNil(request.EndpointConfig) {
		query["EndpointConfig"] = request.EndpointConfig
	}

	if !dara.IsNil(request.Nodes) {
		query["Nodes"] = request.Nodes
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PolarSccTimeoutAction) {
		query["PolarSccTimeoutAction"] = request.PolarSccTimeoutAction
	}

	if !dara.IsNil(request.PolarSccWaitTimeout) {
		query["PolarSccWaitTimeout"] = request.PolarSccWaitTimeout
	}

	if !dara.IsNil(request.ReadWriteMode) {
		query["ReadWriteMode"] = request.ReadWriteMode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SccMode) {
		query["SccMode"] = request.SccMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterEndpointZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterEndpointZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # PolarDB边缘集群修改链接地址
//
// @param request - ModifyDBClusterEndpointZonalRequest
//
// @return ModifyDBClusterEndpointZonalResponse
func (client *Client) ModifyDBClusterEndpointZonal(request *ModifyDBClusterEndpointZonalRequest) (_result *ModifyDBClusterEndpointZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterEndpointZonalResponse{}
	_body, _err := client.ModifyDBClusterEndpointZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the maintenance window of a PolarDB cluster.
//
// Description:
//
// >  We recommend that you set the routine maintenance window to off-peak hours. Alibaba Cloud maintains your cluster within the specified maintenance window to minimize the negative impacts on your business.
//
// @param request - ModifyDBClusterMaintainTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterMaintainTimeResponse
func (client *Client) ModifyDBClusterMaintainTimeWithOptions(request *ModifyDBClusterMaintainTimeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterMaintainTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.MaintainTime) {
		query["MaintainTime"] = request.MaintainTime
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterMaintainTime"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterMaintainTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the maintenance window of a PolarDB cluster.
//
// Description:
//
// >  We recommend that you set the routine maintenance window to off-peak hours. Alibaba Cloud maintains your cluster within the specified maintenance window to minimize the negative impacts on your business.
//
// @param request - ModifyDBClusterMaintainTimeRequest
//
// @return ModifyDBClusterMaintainTimeResponse
func (client *Client) ModifyDBClusterMaintainTime(request *ModifyDBClusterMaintainTimeRequest) (_result *ModifyDBClusterMaintainTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterMaintainTimeResponse{}
	_body, _err := client.ModifyDBClusterMaintainTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Switches or rolls back the task that migrates data from ApsaraDB for RDS to PolarDB.
//
// Description:
//
//	  You can call this operation to switch the task that migrates data from ApsaraDB for RDS to PolarDB.
//
//		- You can call this operation to roll back the task that migrates data from ApsaraDB for RDS to PolarDB.
//
// > Before you call this operation, ensure that a one-click upgrade task has been created for the cluster. You can call the [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) operation to create an upgrade task. Set the **CreationOption*	- parameter to **MigrationFromRDS**. For more information, see [Create a PolarDB for MySQL cluster by using the Migration from RDS method](https://help.aliyun.com/document_detail/121582.html).
//
// @param request - ModifyDBClusterMigrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterMigrationResponse
func (client *Client) ModifyDBClusterMigrationWithOptions(request *ModifyDBClusterMigrationRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterMigrationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionStrings) {
		query["ConnectionStrings"] = request.ConnectionStrings
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.NewMasterInstanceId) {
		query["NewMasterInstanceId"] = request.NewMasterInstanceId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SourceRDSDBInstanceId) {
		query["SourceRDSDBInstanceId"] = request.SourceRDSDBInstanceId
	}

	if !dara.IsNil(request.SwapConnectionString) {
		query["SwapConnectionString"] = request.SwapConnectionString
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterMigration"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterMigrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches or rolls back the task that migrates data from ApsaraDB for RDS to PolarDB.
//
// Description:
//
//	  You can call this operation to switch the task that migrates data from ApsaraDB for RDS to PolarDB.
//
//		- You can call this operation to roll back the task that migrates data from ApsaraDB for RDS to PolarDB.
//
// > Before you call this operation, ensure that a one-click upgrade task has been created for the cluster. You can call the [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) operation to create an upgrade task. Set the **CreationOption*	- parameter to **MigrationFromRDS**. For more information, see [Create a PolarDB for MySQL cluster by using the Migration from RDS method](https://help.aliyun.com/document_detail/121582.html).
//
// @param request - ModifyDBClusterMigrationRequest
//
// @return ModifyDBClusterMigrationResponse
func (client *Client) ModifyDBClusterMigration(request *ModifyDBClusterMigrationRequest) (_result *ModifyDBClusterMigrationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterMigrationResponse{}
	_body, _err := client.ModifyDBClusterMigrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改dts源实例、目标实例
//
// @param request - ModifyDBClusterMigrationEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterMigrationEndpointResponse
func (client *Client) ModifyDBClusterMigrationEndpointWithOptions(request *ModifyDBClusterMigrationEndpointRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterMigrationEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.MigrationConfig) {
		query["MigrationConfig"] = request.MigrationConfig
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterMigrationEndpoint"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterMigrationEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改dts源实例、目标实例
//
// @param request - ModifyDBClusterMigrationEndpointRequest
//
// @return ModifyDBClusterMigrationEndpointResponse
func (client *Client) ModifyDBClusterMigrationEndpoint(request *ModifyDBClusterMigrationEndpointRequest) (_result *ModifyDBClusterMigrationEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterMigrationEndpointResponse{}
	_body, _err := client.ModifyDBClusterMigrationEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the interval at which the monitoring data of a PolarDB cluster is collected.
//
// Description:
//
//	  When the monitoring data is collected every 5 seconds:
//
//	    	- If the query time range is less than or equal to 1 hour, the data is displayed at intervals of 5 seconds.
//
//	    	- If the query time range is less than or equal to one day, the data is displayed at intervals of 1 minute.
//
//	    	- If the query time range is less than or equal to seven days, the data is displayed at intervals of 10 minutes.
//
//	    	- If the query time range is less than or equal to 30 days, the data is displayed at intervals of 1 hour.
//
//	    	- When the query time range is greater than 30 days, the data is displayed at intervals of 1 day.
//
//		- When the monitoring data is collected every 60 seconds:
//
//	    	- If the query time range is less than or equal to one day, the data is displayed at intervals of 1 minute.
//
//	    	- If the query time range is less than or equal to seven days, the data is displayed at intervals of 10 minutes.
//
//	    	- If the query time range is less than or equal to 30 days, the data is displayed at intervals of 1 hour.
//
//	    	- When the query time range is greater than 30 days, the data is displayed at intervals of 1 day.
//
// @param request - ModifyDBClusterMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterMonitorResponse
func (client *Client) ModifyDBClusterMonitorWithOptions(request *ModifyDBClusterMonitorRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
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
		Action:      dara.String("ModifyDBClusterMonitor"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the interval at which the monitoring data of a PolarDB cluster is collected.
//
// Description:
//
//	  When the monitoring data is collected every 5 seconds:
//
//	    	- If the query time range is less than or equal to 1 hour, the data is displayed at intervals of 5 seconds.
//
//	    	- If the query time range is less than or equal to one day, the data is displayed at intervals of 1 minute.
//
//	    	- If the query time range is less than or equal to seven days, the data is displayed at intervals of 10 minutes.
//
//	    	- If the query time range is less than or equal to 30 days, the data is displayed at intervals of 1 hour.
//
//	    	- When the query time range is greater than 30 days, the data is displayed at intervals of 1 day.
//
//		- When the monitoring data is collected every 60 seconds:
//
//	    	- If the query time range is less than or equal to one day, the data is displayed at intervals of 1 minute.
//
//	    	- If the query time range is less than or equal to seven days, the data is displayed at intervals of 10 minutes.
//
//	    	- If the query time range is less than or equal to 30 days, the data is displayed at intervals of 1 hour.
//
//	    	- When the query time range is greater than 30 days, the data is displayed at intervals of 1 day.
//
// @param request - ModifyDBClusterMonitorRequest
//
// @return ModifyDBClusterMonitorResponse
func (client *Client) ModifyDBClusterMonitor(request *ModifyDBClusterMonitorRequest) (_result *ModifyDBClusterMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterMonitorResponse{}
	_body, _err := client.ModifyDBClusterMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the parameters of a specified PolarDB cluster or applies existing parameter templates to a specified cluster.
//
// Description:
//
// PolarDB supports the parameter template feature to centrally manage clusters. You can configure a number of parameters at a time by using a parameter template and apply the template to a PolarDB cluster. For more information, see [Use a parameter template](https://help.aliyun.com/document_detail/207009.html).
//
// **
//
// **Only PolarDB for MySQL clusters support parameter templates.
//
// @param request - ModifyDBClusterParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterParametersResponse
func (client *Client) ModifyDBClusterParametersWithOptions(request *ModifyDBClusterParametersRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterParametersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClearBinlog) {
		query["ClearBinlog"] = request.ClearBinlog
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.FromTimeService) {
		query["FromTimeService"] = request.FromTimeService
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParameterGroupId) {
		query["ParameterGroupId"] = request.ParameterGroupId
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.PlannedEndTime) {
		query["PlannedEndTime"] = request.PlannedEndTime
	}

	if !dara.IsNil(request.PlannedStartTime) {
		query["PlannedStartTime"] = request.PlannedStartTime
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
		Action:      dara.String("ModifyDBClusterParameters"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the parameters of a specified PolarDB cluster or applies existing parameter templates to a specified cluster.
//
// Description:
//
// PolarDB supports the parameter template feature to centrally manage clusters. You can configure a number of parameters at a time by using a parameter template and apply the template to a PolarDB cluster. For more information, see [Use a parameter template](https://help.aliyun.com/document_detail/207009.html).
//
// **
//
// **Only PolarDB for MySQL clusters support parameter templates.
//
// @param request - ModifyDBClusterParametersRequest
//
// @return ModifyDBClusterParametersResponse
func (client *Client) ModifyDBClusterParameters(request *ModifyDBClusterParametersRequest) (_result *ModifyDBClusterParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterParametersResponse{}
	_body, _err := client.ModifyDBClusterParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the primary zone of a PolarDB cluster.
//
// @param request - ModifyDBClusterPrimaryZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterPrimaryZoneResponse
func (client *Client) ModifyDBClusterPrimaryZoneWithOptions(request *ModifyDBClusterPrimaryZoneRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterPrimaryZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.FromTimeService) {
		query["FromTimeService"] = request.FromTimeService
	}

	if !dara.IsNil(request.IsSwitchOverForDisaster) {
		query["IsSwitchOverForDisaster"] = request.IsSwitchOverForDisaster
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlannedEndTime) {
		query["PlannedEndTime"] = request.PlannedEndTime
	}

	if !dara.IsNil(request.PlannedStartTime) {
		query["PlannedStartTime"] = request.PlannedStartTime
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VPCId) {
		query["VPCId"] = request.VPCId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	if !dara.IsNil(request.ZoneType) {
		query["ZoneType"] = request.ZoneType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterPrimaryZone"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterPrimaryZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the primary zone of a PolarDB cluster.
//
// @param request - ModifyDBClusterPrimaryZoneRequest
//
// @return ModifyDBClusterPrimaryZoneResponse
func (client *Client) ModifyDBClusterPrimaryZone(request *ModifyDBClusterPrimaryZoneRequest) (_result *ModifyDBClusterPrimaryZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterPrimaryZoneResponse{}
	_body, _err := client.ModifyDBClusterPrimaryZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a resource group for a database cluster.
//
// @param request - ModifyDBClusterResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterResourceGroupResponse
func (client *Client) ModifyDBClusterResourceGroupWithOptions(request *ModifyDBClusterResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterResourceGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a resource group for a database cluster.
//
// @param request - ModifyDBClusterResourceGroupRequest
//
// @return ModifyDBClusterResourceGroupResponse
func (client *Client) ModifyDBClusterResourceGroup(request *ModifyDBClusterResourceGroupRequest) (_result *ModifyDBClusterResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterResourceGroupResponse{}
	_body, _err := client.ModifyDBClusterResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the SSL encryption feature for a PolarDB cluster, or updates the CA certificate of the cluster.
//
// @param request - ModifyDBClusterSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterSSLResponse
func (client *Client) ModifyDBClusterSSLWithOptions(request *ModifyDBClusterSSLRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterSSLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBEndpointId) {
		query["DBEndpointId"] = request.DBEndpointId
	}

	if !dara.IsNil(request.NetType) {
		query["NetType"] = request.NetType
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

	if !dara.IsNil(request.SSLAutoRotate) {
		query["SSLAutoRotate"] = request.SSLAutoRotate
	}

	if !dara.IsNil(request.SSLEnabled) {
		query["SSLEnabled"] = request.SSLEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterSSL"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the SSL encryption feature for a PolarDB cluster, or updates the CA certificate of the cluster.
//
// @param request - ModifyDBClusterSSLRequest
//
// @return ModifyDBClusterSSLResponse
func (client *Client) ModifyDBClusterSSL(request *ModifyDBClusterSSLRequest) (_result *ModifyDBClusterSSLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterSSLResponse{}
	_body, _err := client.ModifyDBClusterSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a serverless cluster.
//
// @param request - ModifyDBClusterServerlessConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterServerlessConfResponse
func (client *Client) ModifyDBClusterServerlessConfWithOptions(request *ModifyDBClusterServerlessConfRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterServerlessConfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowShutDown) {
		query["AllowShutDown"] = request.AllowShutDown
	}

	if !dara.IsNil(request.CrontabJobId) {
		query["CrontabJobId"] = request.CrontabJobId
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.FromTimeService) {
		query["FromTimeService"] = request.FromTimeService
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlannedEndTime) {
		query["PlannedEndTime"] = request.PlannedEndTime
	}

	if !dara.IsNil(request.PlannedStartTime) {
		query["PlannedStartTime"] = request.PlannedStartTime
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScaleApRoNumMax) {
		query["ScaleApRoNumMax"] = request.ScaleApRoNumMax
	}

	if !dara.IsNil(request.ScaleApRoNumMin) {
		query["ScaleApRoNumMin"] = request.ScaleApRoNumMin
	}

	if !dara.IsNil(request.ScaleMax) {
		query["ScaleMax"] = request.ScaleMax
	}

	if !dara.IsNil(request.ScaleMin) {
		query["ScaleMin"] = request.ScaleMin
	}

	if !dara.IsNil(request.ScaleRoNumMax) {
		query["ScaleRoNumMax"] = request.ScaleRoNumMax
	}

	if !dara.IsNil(request.ScaleRoNumMin) {
		query["ScaleRoNumMin"] = request.ScaleRoNumMin
	}

	if !dara.IsNil(request.SecondsUntilAutoPause) {
		query["SecondsUntilAutoPause"] = request.SecondsUntilAutoPause
	}

	if !dara.IsNil(request.ServerlessRuleCpuEnlargeThreshold) {
		query["ServerlessRuleCpuEnlargeThreshold"] = request.ServerlessRuleCpuEnlargeThreshold
	}

	if !dara.IsNil(request.ServerlessRuleCpuShrinkThreshold) {
		query["ServerlessRuleCpuShrinkThreshold"] = request.ServerlessRuleCpuShrinkThreshold
	}

	if !dara.IsNil(request.ServerlessRuleMode) {
		query["ServerlessRuleMode"] = request.ServerlessRuleMode
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterServerlessConf"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterServerlessConfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a serverless cluster.
//
// @param request - ModifyDBClusterServerlessConfRequest
//
// @return ModifyDBClusterServerlessConfResponse
func (client *Client) ModifyDBClusterServerlessConf(request *ModifyDBClusterServerlessConfRequest) (_result *ModifyDBClusterServerlessConfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterServerlessConfResponse{}
	_body, _err := client.ModifyDBClusterServerlessConfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the storage performance of a PolarDB for MySQL cluster.
//
// @param request - ModifyDBClusterStoragePerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterStoragePerformanceResponse
func (client *Client) ModifyDBClusterStoragePerformanceWithOptions(request *ModifyDBClusterStoragePerformanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterStoragePerformanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.BurstingEnabled) {
		query["BurstingEnabled"] = request.BurstingEnabled
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ModifyType) {
		query["ModifyType"] = request.ModifyType
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.ProvisionedIops) {
		query["ProvisionedIops"] = request.ProvisionedIops
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterStoragePerformance"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterStoragePerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the storage performance of a PolarDB for MySQL cluster.
//
// @param request - ModifyDBClusterStoragePerformanceRequest
//
// @return ModifyDBClusterStoragePerformanceResponse
func (client *Client) ModifyDBClusterStoragePerformance(request *ModifyDBClusterStoragePerformanceRequest) (_result *ModifyDBClusterStoragePerformanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterStoragePerformanceResponse{}
	_body, _err := client.ModifyDBClusterStoragePerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the storage capacity of a pay-as-you-go cluster of Enterprise Edition or a cluster of Standard Edition.
//
// @param request - ModifyDBClusterStorageSpaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterStorageSpaceResponse
func (client *Client) ModifyDBClusterStorageSpaceWithOptions(request *ModifyDBClusterStorageSpaceRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterStorageSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CloudProvider) {
		query["CloudProvider"] = request.CloudProvider
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlannedEndTime) {
		query["PlannedEndTime"] = request.PlannedEndTime
	}

	if !dara.IsNil(request.PlannedStartTime) {
		query["PlannedStartTime"] = request.PlannedStartTime
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StorageSpace) {
		query["StorageSpace"] = request.StorageSpace
	}

	if !dara.IsNil(request.SubCategory) {
		query["SubCategory"] = request.SubCategory
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterStorageSpace"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterStorageSpaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the storage capacity of a pay-as-you-go cluster of Enterprise Edition or a cluster of Standard Edition.
//
// @param request - ModifyDBClusterStorageSpaceRequest
//
// @return ModifyDBClusterStorageSpaceResponse
func (client *Client) ModifyDBClusterStorageSpace(request *ModifyDBClusterStorageSpaceRequest) (_result *ModifyDBClusterStorageSpaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterStorageSpaceResponse{}
	_body, _err := client.ModifyDBClusterStorageSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the transparent data encryption (TDE) feature for a PolarDB cluster.
//
// Description:
//
// > 	- To perform this operation, you must activate KMS first. For more information, see [Purchase a dedicated KMS instance](https://help.aliyun.com/document_detail/153781.html).
//
// > 	- After TDE is enabled, you cannot disable TDE.
//
// @param request - ModifyDBClusterTDERequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterTDEResponse
func (client *Client) ModifyDBClusterTDEWithOptions(request *ModifyDBClusterTDERequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterTDEResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EnableAutomaticRotation) {
		query["EnableAutomaticRotation"] = request.EnableAutomaticRotation
	}

	if !dara.IsNil(request.EncryptNewTables) {
		query["EncryptNewTables"] = request.EncryptNewTables
	}

	if !dara.IsNil(request.EncryptionKey) {
		query["EncryptionKey"] = request.EncryptionKey
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

	if !dara.IsNil(request.RoleArn) {
		query["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.TDEStatus) {
		query["TDEStatus"] = request.TDEStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterTDE"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterTDEResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the transparent data encryption (TDE) feature for a PolarDB cluster.
//
// Description:
//
// > 	- To perform this operation, you must activate KMS first. For more information, see [Purchase a dedicated KMS instance](https://help.aliyun.com/document_detail/153781.html).
//
// > 	- After TDE is enabled, you cannot disable TDE.
//
// @param request - ModifyDBClusterTDERequest
//
// @return ModifyDBClusterTDEResponse
func (client *Client) ModifyDBClusterTDE(request *ModifyDBClusterTDERequest) (_result *ModifyDBClusterTDEResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterTDEResponse{}
	_body, _err := client.ModifyDBClusterTDEWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改集群VSwitch参数
//
// @param request - ModifyDBClusterVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterVpcResponse
func (client *Client) ModifyDBClusterVpcWithOptions(request *ModifyDBClusterVpcRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterVpcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ExistedEndpointSwitchType) {
		query["ExistedEndpointSwitchType"] = request.ExistedEndpointSwitchType
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

	if !dara.IsNil(request.VPCId) {
		query["VPCId"] = request.VPCId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterVpc"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改集群VSwitch参数
//
// @param request - ModifyDBClusterVpcRequest
//
// @return ModifyDBClusterVpcResponse
func (client *Client) ModifyDBClusterVpc(request *ModifyDBClusterVpcRequest) (_result *ModifyDBClusterVpcResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterVpcResponse{}
	_body, _err := client.ModifyDBClusterVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of a database in a PolarDB for MySQL cluster.
//
// @param request - ModifyDBDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBDescriptionResponse
func (client *Client) ModifyDBDescriptionWithOptions(request *ModifyDBDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBDescription) {
		query["DBDescription"] = request.DBDescription
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBDescription"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a database in a PolarDB for MySQL cluster.
//
// @param request - ModifyDBDescriptionRequest
//
// @return ModifyDBDescriptionResponse
func (client *Client) ModifyDBDescription(request *ModifyDBDescriptionRequest) (_result *ModifyDBDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBDescriptionResponse{}
	_body, _err := client.ModifyDBDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改PolarDB边缘云集群数据库描述
//
// @param request - ModifyDBDescriptionZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBDescriptionZonalResponse
func (client *Client) ModifyDBDescriptionZonalWithOptions(request *ModifyDBDescriptionZonalRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBDescriptionZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBDescription) {
		query["DBDescription"] = request.DBDescription
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBDescriptionZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBDescriptionZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改PolarDB边缘云集群数据库描述
//
// @param request - ModifyDBDescriptionZonalRequest
//
// @return ModifyDBDescriptionZonalResponse
func (client *Client) ModifyDBDescriptionZonal(request *ModifyDBDescriptionZonalRequest) (_result *ModifyDBDescriptionZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBDescriptionZonalResponse{}
	_body, _err := client.ModifyDBDescriptionZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the endpoints of a PolarDB cluster, including the primary endpoint, default cluster endpoint, custom cluster endpoint, and private domain name.
//
// @param request - ModifyDBEndpointAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBEndpointAddressResponse
func (client *Client) ModifyDBEndpointAddressWithOptions(request *ModifyDBEndpointAddressRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBEndpointAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBEndpointId) {
		query["DBEndpointId"] = request.DBEndpointId
	}

	if !dara.IsNil(request.NetType) {
		query["NetType"] = request.NetType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.PrivateZoneAddressPrefix) {
		query["PrivateZoneAddressPrefix"] = request.PrivateZoneAddressPrefix
	}

	if !dara.IsNil(request.PrivateZoneName) {
		query["PrivateZoneName"] = request.PrivateZoneName
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
		Action:      dara.String("ModifyDBEndpointAddress"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBEndpointAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the endpoints of a PolarDB cluster, including the primary endpoint, default cluster endpoint, custom cluster endpoint, and private domain name.
//
// @param request - ModifyDBEndpointAddressRequest
//
// @return ModifyDBEndpointAddressResponse
func (client *Client) ModifyDBEndpointAddress(request *ModifyDBEndpointAddressRequest) (_result *ModifyDBEndpointAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBEndpointAddressResponse{}
	_body, _err := client.ModifyDBEndpointAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the node specifications of a PolarDB cluster.
//
// @param request - ModifyDBNodeClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBNodeClassResponse
func (client *Client) ModifyDBNodeClassWithOptions(request *ModifyDBNodeClassRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBNodeClassResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CloudProvider) {
		query["CloudProvider"] = request.CloudProvider
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNodeTargetClass) {
		query["DBNodeTargetClass"] = request.DBNodeTargetClass
	}

	if !dara.IsNil(request.DBNodeType) {
		query["DBNodeType"] = request.DBNodeType
	}

	if !dara.IsNil(request.ModifyType) {
		query["ModifyType"] = request.ModifyType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlannedEndTime) {
		query["PlannedEndTime"] = request.PlannedEndTime
	}

	if !dara.IsNil(request.PlannedFlashingOffTime) {
		query["PlannedFlashingOffTime"] = request.PlannedFlashingOffTime
	}

	if !dara.IsNil(request.PlannedStartTime) {
		query["PlannedStartTime"] = request.PlannedStartTime
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SubCategory) {
		query["SubCategory"] = request.SubCategory
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBNodeClass"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBNodeClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the node specifications of a PolarDB cluster.
//
// @param request - ModifyDBNodeClassRequest
//
// @return ModifyDBNodeClassResponse
func (client *Client) ModifyDBNodeClass(request *ModifyDBNodeClassRequest) (_result *ModifyDBNodeClassResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBNodeClassResponse{}
	_body, _err := client.ModifyDBNodeClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyDBNodeConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBNodeConfigResponse
func (client *Client) ModifyDBNodeConfigWithOptions(request *ModifyDBNodeConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBNodeConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigName) {
		query["ConfigName"] = request.ConfigName
	}

	if !dara.IsNil(request.ConfigValue) {
		query["ConfigValue"] = request.ConfigValue
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNodeId) {
		query["DBNodeId"] = request.DBNodeId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBNodeConfig"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBNodeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyDBNodeConfigRequest
//
// @return ModifyDBNodeConfigResponse
func (client *Client) ModifyDBNodeConfig(request *ModifyDBNodeConfigRequest) (_result *ModifyDBNodeConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBNodeConfigResponse{}
	_body, _err := client.ModifyDBNodeConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改物理节点描述
//
// @param request - ModifyDBNodeDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBNodeDescriptionResponse
func (client *Client) ModifyDBNodeDescriptionWithOptions(request *ModifyDBNodeDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBNodeDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNodeDescription) {
		query["DBNodeDescription"] = request.DBNodeDescription
	}

	if !dara.IsNil(request.DBNodeId) {
		query["DBNodeId"] = request.DBNodeId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBNodeDescription"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBNodeDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改物理节点描述
//
// @param request - ModifyDBNodeDescriptionRequest
//
// @return ModifyDBNodeDescriptionResponse
func (client *Client) ModifyDBNodeDescription(request *ModifyDBNodeDescriptionRequest) (_result *ModifyDBNodeDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBNodeDescriptionResponse{}
	_body, _err := client.ModifyDBNodeDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the failover with hot replica feature for a node in a cluster.
//
// @param request - ModifyDBNodeHotReplicaModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBNodeHotReplicaModeResponse
func (client *Client) ModifyDBNodeHotReplicaModeWithOptions(request *ModifyDBNodeHotReplicaModeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBNodeHotReplicaModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNodeId) {
		query["DBNodeId"] = request.DBNodeId
	}

	if !dara.IsNil(request.HotReplicaMode) {
		query["HotReplicaMode"] = request.HotReplicaMode
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBNodeHotReplicaMode"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBNodeHotReplicaModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the failover with hot replica feature for a node in a cluster.
//
// @param request - ModifyDBNodeHotReplicaModeRequest
//
// @return ModifyDBNodeHotReplicaModeResponse
func (client *Client) ModifyDBNodeHotReplicaMode(request *ModifyDBNodeHotReplicaModeRequest) (_result *ModifyDBNodeHotReplicaModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBNodeHotReplicaModeResponse{}
	_body, _err := client.ModifyDBNodeHotReplicaModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改节点的Scc
//
// @param request - ModifyDBNodeSccModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBNodeSccModeResponse
func (client *Client) ModifyDBNodeSccModeWithOptions(request *ModifyDBNodeSccModeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBNodeSccModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNodeId) {
		query["DBNodeId"] = request.DBNodeId
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

	if !dara.IsNil(request.SccMode) {
		query["SccMode"] = request.SccMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBNodeSccMode"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBNodeSccModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改节点的Scc
//
// @param request - ModifyDBNodeSccModeRequest
//
// @return ModifyDBNodeSccModeResponse
func (client *Client) ModifyDBNodeSccMode(request *ModifyDBNodeSccModeRequest) (_result *ModifyDBNodeSccModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBNodeSccModeResponse{}
	_body, _err := client.ModifyDBNodeSccModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the specifications of a node in a PolarDB cluster.
//
// @param request - ModifyDBNodesClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBNodesClassResponse
func (client *Client) ModifyDBNodesClassWithOptions(request *ModifyDBNodesClassRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBNodesClassResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CloudProvider) {
		query["CloudProvider"] = request.CloudProvider
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNode) {
		query["DBNode"] = request.DBNode
	}

	if !dara.IsNil(request.ModifyType) {
		query["ModifyType"] = request.ModifyType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlannedEndTime) {
		query["PlannedEndTime"] = request.PlannedEndTime
	}

	if !dara.IsNil(request.PlannedFlashingOffTime) {
		query["PlannedFlashingOffTime"] = request.PlannedFlashingOffTime
	}

	if !dara.IsNil(request.PlannedStartTime) {
		query["PlannedStartTime"] = request.PlannedStartTime
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SubCategory) {
		query["SubCategory"] = request.SubCategory
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBNodesClass"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBNodesClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the specifications of a node in a PolarDB cluster.
//
// @param request - ModifyDBNodesClassRequest
//
// @return ModifyDBNodesClassResponse
func (client *Client) ModifyDBNodesClass(request *ModifyDBNodesClassRequest) (_result *ModifyDBNodesClassResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBNodesClassResponse{}
	_body, _err := client.ModifyDBNodesClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the parameters of a node and applies them to specified nodes.
//
// @param request - ModifyDBNodesParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBNodesParametersResponse
func (client *Client) ModifyDBNodesParametersWithOptions(request *ModifyDBNodesParametersRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBNodesParametersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNodeIds) {
		query["DBNodeIds"] = request.DBNodeIds
	}

	if !dara.IsNil(request.FromTimeService) {
		query["FromTimeService"] = request.FromTimeService
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParameterGroupId) {
		query["ParameterGroupId"] = request.ParameterGroupId
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.PlannedEndTime) {
		query["PlannedEndTime"] = request.PlannedEndTime
	}

	if !dara.IsNil(request.PlannedStartTime) {
		query["PlannedStartTime"] = request.PlannedStartTime
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
		Action:      dara.String("ModifyDBNodesParameters"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBNodesParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the parameters of a node and applies them to specified nodes.
//
// @param request - ModifyDBNodesParametersRequest
//
// @return ModifyDBNodesParametersResponse
func (client *Client) ModifyDBNodesParameters(request *ModifyDBNodesParametersRequest) (_result *ModifyDBNodesParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBNodesParametersResponse{}
	_body, _err := client.ModifyDBNodesParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变更角色权限
//
// @param request - ModifyEncryptionDBRolePrivilegeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEncryptionDBRolePrivilegeResponse
func (client *Client) ModifyEncryptionDBRolePrivilegeWithOptions(request *ModifyEncryptionDBRolePrivilegeRequest, runtime *dara.RuntimeOptions) (_result *ModifyEncryptionDBRolePrivilegeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.RolePrivilegeConfig) {
		query["RolePrivilegeConfig"] = request.RolePrivilegeConfig
	}

	if !dara.IsNil(request.RolePrivilegeName) {
		query["RolePrivilegeName"] = request.RolePrivilegeName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyEncryptionDBRolePrivilege"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEncryptionDBRolePrivilegeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更角色权限
//
// @param request - ModifyEncryptionDBRolePrivilegeRequest
//
// @return ModifyEncryptionDBRolePrivilegeResponse
func (client *Client) ModifyEncryptionDBRolePrivilege(request *ModifyEncryptionDBRolePrivilegeRequest) (_result *ModifyEncryptionDBRolePrivilegeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyEncryptionDBRolePrivilegeResponse{}
	_body, _err := client.ModifyEncryptionDBRolePrivilegeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变更加密策略
//
// @param request - ModifyEncryptionDBSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEncryptionDBSecretResponse
func (client *Client) ModifyEncryptionDBSecretWithOptions(request *ModifyEncryptionDBSecretRequest, runtime *dara.RuntimeOptions) (_result *ModifyEncryptionDBSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EncryptionDBStatus) {
		query["EncryptionDBStatus"] = request.EncryptionDBStatus
	}

	if !dara.IsNil(request.EncryptionKey) {
		query["EncryptionKey"] = request.EncryptionKey
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

	if !dara.IsNil(request.RoleArn) {
		query["RoleArn"] = request.RoleArn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyEncryptionDBSecret"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEncryptionDBSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更加密策略
//
// @param request - ModifyEncryptionDBSecretRequest
//
// @return ModifyEncryptionDBSecretResponse
func (client *Client) ModifyEncryptionDBSecret(request *ModifyEncryptionDBSecretRequest) (_result *ModifyEncryptionDBSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyEncryptionDBSecretResponse{}
	_body, _err := client.ModifyEncryptionDBSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改sql防火墙配置
//
// @param request - ModifyFirewallRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFirewallRulesResponse
func (client *Client) ModifyFirewallRulesWithOptions(request *ModifyFirewallRulesRequest, runtime *dara.RuntimeOptions) (_result *ModifyFirewallRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.RuleConfig) {
		query["RuleConfig"] = request.RuleConfig
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFirewallRules"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFirewallRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改sql防火墙配置
//
// @param request - ModifyFirewallRulesRequest
//
// @return ModifyFirewallRulesResponse
func (client *Client) ModifyFirewallRules(request *ModifyFirewallRulesRequest) (_result *ModifyFirewallRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyFirewallRulesResponse{}
	_body, _err := client.ModifyFirewallRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a global database network (GDN).
//
// @param request - ModifyGlobalDatabaseNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGlobalDatabaseNetworkResponse
func (client *Client) ModifyGlobalDatabaseNetworkWithOptions(request *ModifyGlobalDatabaseNetworkRequest, runtime *dara.RuntimeOptions) (_result *ModifyGlobalDatabaseNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableGlobalDomainName) {
		query["EnableGlobalDomainName"] = request.EnableGlobalDomainName
	}

	if !dara.IsNil(request.GDNDescription) {
		query["GDNDescription"] = request.GDNDescription
	}

	if !dara.IsNil(request.GDNId) {
		query["GDNId"] = request.GDNId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyGlobalDatabaseNetwork"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGlobalDatabaseNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a global database network (GDN).
//
// @param request - ModifyGlobalDatabaseNetworkRequest
//
// @return ModifyGlobalDatabaseNetworkResponse
func (client *Client) ModifyGlobalDatabaseNetwork(request *ModifyGlobalDatabaseNetworkRequest) (_result *ModifyGlobalDatabaseNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyGlobalDatabaseNetworkResponse{}
	_body, _err := client.ModifyGlobalDatabaseNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an IP whitelist template.
//
// @param request - ModifyGlobalSecurityIPGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGlobalSecurityIPGroupResponse
func (client *Client) ModifyGlobalSecurityIPGroupWithOptions(request *ModifyGlobalSecurityIPGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyGlobalSecurityIPGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GIpList) {
		query["GIpList"] = request.GIpList
	}

	if !dara.IsNil(request.GlobalIgName) {
		query["GlobalIgName"] = request.GlobalIgName
	}

	if !dara.IsNil(request.GlobalSecurityGroupId) {
		query["GlobalSecurityGroupId"] = request.GlobalSecurityGroupId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyGlobalSecurityIPGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGlobalSecurityIPGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an IP whitelist template.
//
// @param request - ModifyGlobalSecurityIPGroupRequest
//
// @return ModifyGlobalSecurityIPGroupResponse
func (client *Client) ModifyGlobalSecurityIPGroup(request *ModifyGlobalSecurityIPGroupRequest) (_result *ModifyGlobalSecurityIPGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyGlobalSecurityIPGroupResponse{}
	_body, _err := client.ModifyGlobalSecurityIPGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name of a global IP whitelist template.
//
// @param request - ModifyGlobalSecurityIPGroupNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGlobalSecurityIPGroupNameResponse
func (client *Client) ModifyGlobalSecurityIPGroupNameWithOptions(request *ModifyGlobalSecurityIPGroupNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyGlobalSecurityIPGroupNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GlobalIgName) {
		query["GlobalIgName"] = request.GlobalIgName
	}

	if !dara.IsNil(request.GlobalSecurityGroupId) {
		query["GlobalSecurityGroupId"] = request.GlobalSecurityGroupId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyGlobalSecurityIPGroupName"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGlobalSecurityIPGroupNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of a global IP whitelist template.
//
// @param request - ModifyGlobalSecurityIPGroupNameRequest
//
// @return ModifyGlobalSecurityIPGroupNameResponse
func (client *Client) ModifyGlobalSecurityIPGroupName(request *ModifyGlobalSecurityIPGroupNameRequest) (_result *ModifyGlobalSecurityIPGroupNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyGlobalSecurityIPGroupNameResponse{}
	_body, _err := client.ModifyGlobalSecurityIPGroupNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the relationship between a cluster and a global IP whitelist template.
//
// @param request - ModifyGlobalSecurityIPGroupRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGlobalSecurityIPGroupRelationResponse
func (client *Client) ModifyGlobalSecurityIPGroupRelationWithOptions(request *ModifyGlobalSecurityIPGroupRelationRequest, runtime *dara.RuntimeOptions) (_result *ModifyGlobalSecurityIPGroupRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.GlobalSecurityGroupId) {
		query["GlobalSecurityGroupId"] = request.GlobalSecurityGroupId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyGlobalSecurityIPGroupRelation"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGlobalSecurityIPGroupRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the relationship between a cluster and a global IP whitelist template.
//
// @param request - ModifyGlobalSecurityIPGroupRelationRequest
//
// @return ModifyGlobalSecurityIPGroupRelationResponse
func (client *Client) ModifyGlobalSecurityIPGroupRelation(request *ModifyGlobalSecurityIPGroupRelationRequest) (_result *ModifyGlobalSecurityIPGroupRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyGlobalSecurityIPGroupRelationResponse{}
	_body, _err := client.ModifyGlobalSecurityIPGroupRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the retention policy of the log backups in a PolarDB cluster.
//
// @param tmpReq - ModifyLogBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLogBackupPolicyResponse
func (client *Client) ModifyLogBackupPolicyWithOptions(tmpReq *ModifyLogBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyLogBackupPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyLogBackupPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AdvancedLogPolicies) {
		request.AdvancedLogPoliciesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdvancedLogPolicies, dara.String("AdvancedLogPolicies"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdvancedLogPoliciesShrink) {
		query["AdvancedLogPolicies"] = request.AdvancedLogPoliciesShrink
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.LogBackupAnotherRegionRegion) {
		query["LogBackupAnotherRegionRegion"] = request.LogBackupAnotherRegionRegion
	}

	if !dara.IsNil(request.LogBackupAnotherRegionRetentionPeriod) {
		query["LogBackupAnotherRegionRetentionPeriod"] = request.LogBackupAnotherRegionRetentionPeriod
	}

	if !dara.IsNil(request.LogBackupRetentionPeriod) {
		query["LogBackupRetentionPeriod"] = request.LogBackupRetentionPeriod
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLogBackupPolicy"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLogBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the retention policy of the log backups in a PolarDB cluster.
//
// @param request - ModifyLogBackupPolicyRequest
//
// @return ModifyLogBackupPolicyResponse
func (client *Client) ModifyLogBackupPolicy(request *ModifyLogBackupPolicyRequest) (_result *ModifyLogBackupPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyLogBackupPolicyResponse{}
	_body, _err := client.ModifyLogBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies or adds a data masking rule.
//
// @param request - ModifyMaskingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMaskingRulesResponse
func (client *Client) ModifyMaskingRulesWithOptions(request *ModifyMaskingRulesRequest, runtime *dara.RuntimeOptions) (_result *ModifyMaskingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DefaultAlgo) {
		query["DefaultAlgo"] = request.DefaultAlgo
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.InterfaceVersion) {
		query["InterfaceVersion"] = request.InterfaceVersion
	}

	if !dara.IsNil(request.MaskingAlgo) {
		query["MaskingAlgo"] = request.MaskingAlgo
	}

	if !dara.IsNil(request.RuleConfig) {
		query["RuleConfig"] = request.RuleConfig
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleNameList) {
		query["RuleNameList"] = request.RuleNameList
	}

	if !dara.IsNil(request.RuleVersion) {
		query["RuleVersion"] = request.RuleVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMaskingRules"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMaskingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies or adds a data masking rule.
//
// @param request - ModifyMaskingRulesRequest
//
// @return ModifyMaskingRulesResponse
func (client *Client) ModifyMaskingRules(request *ModifyMaskingRulesRequest) (_result *ModifyMaskingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMaskingRulesResponse{}
	_body, _err := client.ModifyMaskingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改模型API
//
// @param request - ModifyModelApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyModelApiResponse
func (client *Client) ModifyModelApiWithOptions(request *ModifyModelApiRequest, runtime *dara.RuntimeOptions) (_result *ModifyModelApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.ModelApiId) {
		query["ModelApiId"] = request.ModelApiId
	}

	if !dara.IsNil(request.ModelCategory) {
		query["ModelCategory"] = request.ModelCategory
	}

	if !dara.IsNil(request.PathPrefix) {
		query["PathPrefix"] = request.PathPrefix
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RecordInput) {
		query["RecordInput"] = request.RecordInput
	}

	if !dara.IsNil(request.RecordOutput) {
		query["RecordOutput"] = request.RecordOutput
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RouteRules) {
		query["RouteRules"] = request.RouteRules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyModelApi"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyModelApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改模型API
//
// @param request - ModifyModelApiRequest
//
// @return ModifyModelApiResponse
func (client *Client) ModifyModelApi(request *ModifyModelApiRequest) (_result *ModifyModelApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyModelApiResponse{}
	_body, _err := client.ModifyModelApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改模型服务
//
// @param request - ModifyModelServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyModelServiceResponse
func (client *Client) ModifyModelServiceWithOptions(request *ModifyModelServiceRequest, runtime *dara.RuntimeOptions) (_result *ModifyModelServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.BaseUrl) {
		query["BaseUrl"] = request.BaseUrl
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.InputCostPointsPerMillion) {
		query["InputCostPointsPerMillion"] = request.InputCostPointsPerMillion
	}

	if !dara.IsNil(request.ModelCategory) {
		query["ModelCategory"] = request.ModelCategory
	}

	if !dara.IsNil(request.ModelServiceId) {
		query["ModelServiceId"] = request.ModelServiceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OutputCostPointsPerMillion) {
		query["OutputCostPointsPerMillion"] = request.OutputCostPointsPerMillion
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestCostPoints) {
		query["RequestCostPoints"] = request.RequestCostPoints
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyModelService"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyModelServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改模型服务
//
// @param request - ModifyModelServiceRequest
//
// @return ModifyModelServiceResponse
func (client *Client) ModifyModelService(request *ModifyModelServiceRequest) (_result *ModifyModelServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyModelServiceResponse{}
	_body, _err := client.ModifyModelServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the switching time of a pending event.
//
// @param request - ModifyPendingMaintenanceActionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPendingMaintenanceActionResponse
func (client *Client) ModifyPendingMaintenanceActionWithOptions(request *ModifyPendingMaintenanceActionRequest, runtime *dara.RuntimeOptions) (_result *ModifyPendingMaintenanceActionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPendingMaintenanceAction"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPendingMaintenanceActionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the switching time of a pending event.
//
// @param request - ModifyPendingMaintenanceActionRequest
//
// @return ModifyPendingMaintenanceActionResponse
func (client *Client) ModifyPendingMaintenanceAction(request *ModifyPendingMaintenanceActionRequest) (_result *ModifyPendingMaintenanceActionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPendingMaintenanceActionResponse{}
	_body, _err := client.ModifyPendingMaintenanceActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改限流策略
//
// @param request - ModifyRateLimitPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRateLimitPolicyResponse
func (client *Client) ModifyRateLimitPolicyWithOptions(request *ModifyRateLimitPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyRateLimitPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RateLimitRpm) {
		query["RateLimitRpm"] = request.RateLimitRpm
	}

	if !dara.IsNil(request.RateLimitTpm) {
		query["RateLimitTpm"] = request.RateLimitTpm
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRateLimitPolicy"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRateLimitPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改限流策略
//
// @param request - ModifyRateLimitPolicyRequest
//
// @return ModifyRateLimitPolicyResponse
func (client *Client) ModifyRateLimitPolicy(request *ModifyRateLimitPolicyRequest) (_result *ModifyRateLimitPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyRateLimitPolicyResponse{}
	_body, _err := client.ModifyRateLimitPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置资源包
//
// @param request - ModifyResourcePackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyResourcePackageResponse
func (client *Client) ModifyResourcePackageWithOptions(request *ModifyResourcePackageRequest, runtime *dara.RuntimeOptions) (_result *ModifyResourcePackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoQuota) {
		query["AutoQuota"] = request.AutoQuota
	}

	if !dara.IsNil(request.ResourcePackageId) {
		query["ResourcePackageId"] = request.ResourcePackageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyResourcePackage"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyResourcePackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置资源包
//
// @param request - ModifyResourcePackageRequest
//
// @return ModifyResourcePackageResponse
func (client *Client) ModifyResourcePackage(request *ModifyResourcePackageRequest) (_result *ModifyResourcePackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyResourcePackageResponse{}
	_body, _err := client.ModifyResourcePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改SQL限流规则
//
// @param request - ModifySQLRateLimitingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySQLRateLimitingRulesResponse
func (client *Client) ModifySQLRateLimitingRulesWithOptions(request *ModifySQLRateLimitingRulesRequest, runtime *dara.RuntimeOptions) (_result *ModifySQLRateLimitingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.RuleConfig) {
		query["RuleConfig"] = request.RuleConfig
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySQLRateLimitingRules"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySQLRateLimitingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改SQL限流规则
//
// @param request - ModifySQLRateLimitingRulesRequest
//
// @return ModifySQLRateLimitingRulesResponse
func (client *Client) ModifySQLRateLimitingRules(request *ModifySQLRateLimitingRulesRequest) (_result *ModifySQLRateLimitingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySQLRateLimitingRulesResponse{}
	_body, _err := client.ModifySQLRateLimitingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改目标计划任务
//
// @param request - ModifyScheduleTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScheduleTaskResponse
func (client *Client) ModifyScheduleTaskWithOptions(request *ModifyScheduleTaskRequest, runtime *dara.RuntimeOptions) (_result *ModifyScheduleTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlannedEndTime) {
		query["PlannedEndTime"] = request.PlannedEndTime
	}

	if !dara.IsNil(request.PlannedStartTime) {
		query["PlannedStartTime"] = request.PlannedStartTime
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

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyScheduleTask"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyScheduleTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改目标计划任务
//
// @param request - ModifyScheduleTaskRequest
//
// @return ModifyScheduleTaskResponse
func (client *Client) ModifyScheduleTask(request *ModifyScheduleTaskRequest) (_result *ModifyScheduleTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyScheduleTaskResponse{}
	_body, _err := client.ModifyScheduleTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重命名或移动文件
//
// @param request - MovePolarFsObjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MovePolarFsObjectsResponse
func (client *Client) MovePolarFsObjectsWithOptions(request *MovePolarFsObjectsRequest, runtime *dara.RuntimeOptions) (_result *MovePolarFsObjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ObjectsToMove) {
		query["ObjectsToMove"] = request.ObjectsToMove
	}

	if !dara.IsNil(request.PolarFsInstanceId) {
		query["PolarFsInstanceId"] = request.PolarFsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MovePolarFsObjects"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MovePolarFsObjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重命名或移动文件
//
// @param request - MovePolarFsObjectsRequest
//
// @return MovePolarFsObjectsResponse
func (client *Client) MovePolarFsObjects(request *MovePolarFsObjectsRequest) (_result *MovePolarFsObjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MovePolarFsObjectsResponse{}
	_body, _err := client.MovePolarFsObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the PolarDB for AI feature for a cluster.
//
// @param request - OpenAITaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenAITaskResponse
func (client *Client) OpenAITaskWithOptions(request *OpenAITaskRequest, runtime *dara.RuntimeOptions) (_result *OpenAITaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
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

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenAITask"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenAITaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the PolarDB for AI feature for a cluster.
//
// @param request - OpenAITaskRequest
//
// @return OpenAITaskResponse
func (client *Client) OpenAITask(request *OpenAITaskRequest) (_result *OpenAITaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenAITaskResponse{}
	_body, _err := client.OpenAITaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新PolarClaw配置
//
// @param tmpReq - PatchPolarClawConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PatchPolarClawConfigResponse
func (client *Client) PatchPolarClawConfigWithOptions(tmpReq *PatchPolarClawConfigRequest, runtime *dara.RuntimeOptions) (_result *PatchPolarClawConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PatchPolarClawConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ConfigPatch) {
		request.ConfigPatchShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfigPatch, dara.String("ConfigPatch"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ConfigPatchShrink) {
		query["ConfigPatch"] = request.ConfigPatchShrink
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PatchPolarClawConfig"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PatchPolarClawConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新PolarClaw配置
//
// @param request - PatchPolarClawConfigRequest
//
// @return PatchPolarClawConfigResponse
func (client *Client) PatchPolarClawConfig(request *PatchPolarClawConfigRequest) (_result *PatchPolarClawConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PatchPolarClawConfigResponse{}
	_body, _err := client.PatchPolarClawConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reactivates the backup feature.
//
// @param request - ReactivateDBClusterBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReactivateDBClusterBackupResponse
func (client *Client) ReactivateDBClusterBackupWithOptions(request *ReactivateDBClusterBackupRequest, runtime *dara.RuntimeOptions) (_result *ReactivateDBClusterBackupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReactivateDBClusterBackup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReactivateDBClusterBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reactivates the backup feature.
//
// @param request - ReactivateDBClusterBackupRequest
//
// @return ReactivateDBClusterBackupResponse
func (client *Client) ReactivateDBClusterBackup(request *ReactivateDBClusterBackupRequest) (_result *ReactivateDBClusterBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReactivateDBClusterBackupResponse{}
	_body, _err := client.ReactivateDBClusterBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the storage usage of a cluster.
//
// @param request - RefreshDBClusterStorageUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshDBClusterStorageUsageResponse
func (client *Client) RefreshDBClusterStorageUsageWithOptions(request *RefreshDBClusterStorageUsageRequest, runtime *dara.RuntimeOptions) (_result *RefreshDBClusterStorageUsageResponse, _err error) {
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

	if !dara.IsNil(request.SyncRealTime) {
		query["SyncRealTime"] = request.SyncRealTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshDBClusterStorageUsage"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshDBClusterStorageUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the storage usage of a cluster.
//
// @param request - RefreshDBClusterStorageUsageRequest
//
// @return RefreshDBClusterStorageUsageResponse
func (client *Client) RefreshDBClusterStorageUsage(request *RefreshDBClusterStorageUsageRequest) (_result *RefreshDBClusterStorageUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefreshDBClusterStorageUsageResponse{}
	_body, _err := client.RefreshDBClusterStorageUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除应用环境变量
//
// @param tmpReq - RemoveApplicationEnvironmentVariablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveApplicationEnvironmentVariablesResponse
func (client *Client) RemoveApplicationEnvironmentVariablesWithOptions(tmpReq *RemoveApplicationEnvironmentVariablesRequest, runtime *dara.RuntimeOptions) (_result *RemoveApplicationEnvironmentVariablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveApplicationEnvironmentVariablesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.VariableNames) {
		request.VariableNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VariableNames, dara.String("VariableNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	if !dara.IsNil(request.VariableNamesShrink) {
		query["VariableNames"] = request.VariableNamesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveApplicationEnvironmentVariables"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveApplicationEnvironmentVariablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除应用环境变量
//
// @param request - RemoveApplicationEnvironmentVariablesRequest
//
// @return RemoveApplicationEnvironmentVariablesResponse
func (client *Client) RemoveApplicationEnvironmentVariables(request *RemoveApplicationEnvironmentVariablesRequest) (_result *RemoveApplicationEnvironmentVariablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveApplicationEnvironmentVariablesResponse{}
	_body, _err := client.RemoveApplicationEnvironmentVariablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a secondary cluster from a GDN.
//
// Description:
//
// >  You cannot remove the primary cluster from a GDN.
//
// @param request - RemoveDBClusterFromGDNRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDBClusterFromGDNResponse
func (client *Client) RemoveDBClusterFromGDNWithOptions(request *RemoveDBClusterFromGDNRequest, runtime *dara.RuntimeOptions) (_result *RemoveDBClusterFromGDNResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.GDNId) {
		query["GDNId"] = request.GDNId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TargetDBClusterId) {
		query["TargetDBClusterId"] = request.TargetDBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveDBClusterFromGDN"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveDBClusterFromGDNResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a secondary cluster from a GDN.
//
// Description:
//
// >  You cannot remove the primary cluster from a GDN.
//
// @param request - RemoveDBClusterFromGDNRequest
//
// @return RemoveDBClusterFromGDNResponse
func (client *Client) RemoveDBClusterFromGDN(request *RemoveDBClusterFromGDNRequest) (_result *RemoveDBClusterFromGDNResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveDBClusterFromGDNResponse{}
	_body, _err := client.RemoveDBClusterFromGDNWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除PolarClaw MCP Server
//
// @param request - RemovePolarClawMCPServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemovePolarClawMCPServerResponse
func (client *Client) RemovePolarClawMCPServerWithOptions(request *RemovePolarClawMCPServerRequest, runtime *dara.RuntimeOptions) (_result *RemovePolarClawMCPServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ServerName) {
		query["ServerName"] = request.ServerName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemovePolarClawMCPServer"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemovePolarClawMCPServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除PolarClaw MCP Server
//
// @param request - RemovePolarClawMCPServerRequest
//
// @return RemovePolarClawMCPServerResponse
func (client *Client) RemovePolarClawMCPServer(request *RemovePolarClawMCPServerRequest) (_result *RemovePolarClawMCPServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemovePolarClawMCPServerResponse{}
	_body, _err := client.RemovePolarClawMCPServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the permissions of a privileged account for a PolarDB cluster.
//
// Description:
//
// >- Only PolarDB for MySQL clusters support this operation.
//
// >- If the privileged account of your cluster encounters exceptions, you can call this operation to reset the permissions. For example, the permissions are accidentally revoked.
//
// @param request - ResetAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAccountResponse
func (client *Client) ResetAccountWithOptions(request *ResetAccountRequest, runtime *dara.RuntimeOptions) (_result *ResetAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAccount"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the permissions of a privileged account for a PolarDB cluster.
//
// Description:
//
// >- Only PolarDB for MySQL clusters support this operation.
//
// >- If the privileged account of your cluster encounters exceptions, you can call this operation to reset the permissions. For example, the permissions are accidentally revoked.
//
// @param request - ResetAccountRequest
//
// @return ResetAccountResponse
func (client *Client) ResetAccount(request *ResetAccountRequest) (_result *ResetAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetAccountResponse{}
	_body, _err := client.ResetAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ResetAccountPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAccountPasswordResponse
func (client *Client) ResetAccountPasswordWithOptions(request *ResetAccountPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetAccountPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAccountPassword"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResetAccountPasswordRequest
//
// @return ResetAccountPasswordResponse
func (client *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (_result *ResetAccountPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.ResetAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重置PolarDB边缘集群账号
//
// @param request - ResetAccountZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAccountZonalResponse
func (client *Client) ResetAccountZonalWithOptions(request *ResetAccountZonalRequest, runtime *dara.RuntimeOptions) (_result *ResetAccountZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAccountZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetAccountZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重置PolarDB边缘集群账号
//
// @param request - ResetAccountZonalRequest
//
// @return ResetAccountZonalResponse
func (client *Client) ResetAccountZonal(request *ResetAccountZonalRequest) (_result *ResetAccountZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetAccountZonalResponse{}
	_body, _err := client.ResetAccountZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重置API密钥
//
// @param request - ResetConsumerApiKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetConsumerApiKeyResponse
func (client *Client) ResetConsumerApiKeyWithOptions(request *ResetConsumerApiKeyRequest, runtime *dara.RuntimeOptions) (_result *ResetConsumerApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerId) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !dara.IsNil(request.GwClusterId) {
		query["GwClusterId"] = request.GwClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetConsumerApiKey"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetConsumerApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重置API密钥
//
// @param request - ResetConsumerApiKeyRequest
//
// @return ResetConsumerApiKeyResponse
func (client *Client) ResetConsumerApiKey(request *ResetConsumerApiKeyRequest) (_result *ResetConsumerApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetConsumerApiKeyResponse{}
	_body, _err := client.ResetConsumerApiKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rebuilds a secondary cluster in a Global Database Network (GDN).
//
// @param request - ResetGlobalDatabaseNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetGlobalDatabaseNetworkResponse
func (client *Client) ResetGlobalDatabaseNetworkWithOptions(request *ResetGlobalDatabaseNetworkRequest, runtime *dara.RuntimeOptions) (_result *ResetGlobalDatabaseNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.GDNId) {
		query["GDNId"] = request.GDNId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetGlobalDatabaseNetwork"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetGlobalDatabaseNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rebuilds a secondary cluster in a Global Database Network (GDN).
//
// @param request - ResetGlobalDatabaseNetworkRequest
//
// @return ResetGlobalDatabaseNetworkResponse
func (client *Client) ResetGlobalDatabaseNetwork(request *ResetGlobalDatabaseNetworkRequest) (_result *ResetGlobalDatabaseNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetGlobalDatabaseNetworkResponse{}
	_body, _err := client.ResetGlobalDatabaseNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts database links.
//
// @param request - RestartDBLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartDBLinkResponse
func (client *Client) RestartDBLinkWithOptions(request *RestartDBLinkRequest, runtime *dara.RuntimeOptions) (_result *RestartDBLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartDBLink"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartDBLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts database links.
//
// @param request - RestartDBLinkRequest
//
// @return RestartDBLinkResponse
func (client *Client) RestartDBLink(request *RestartDBLinkRequest) (_result *RestartDBLinkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestartDBLinkResponse{}
	_body, _err := client.RestartDBLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts a node in a PolarDB cluster.
//
// @param request - RestartDBNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartDBNodeResponse
func (client *Client) RestartDBNodeWithOptions(request *RestartDBNodeRequest, runtime *dara.RuntimeOptions) (_result *RestartDBNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBNodeId) {
		query["DBNodeId"] = request.DBNodeId
	}

	if !dara.IsNil(request.FromTimeService) {
		query["FromTimeService"] = request.FromTimeService
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlannedEndTime) {
		query["PlannedEndTime"] = request.PlannedEndTime
	}

	if !dara.IsNil(request.PlannedStartTime) {
		query["PlannedStartTime"] = request.PlannedStartTime
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
		Action:      dara.String("RestartDBNode"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartDBNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a node in a PolarDB cluster.
//
// @param request - RestartDBNodeRequest
//
// @return RestartDBNodeResponse
func (client *Client) RestartDBNode(request *RestartDBNodeRequest) (_result *RestartDBNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestartDBNodeResponse{}
	_body, _err := client.RestartDBNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重启poalrdb边缘云集群节点
//
// @param request - RestartDBNodeZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartDBNodeZonalResponse
func (client *Client) RestartDBNodeZonalWithOptions(request *RestartDBNodeZonalRequest, runtime *dara.RuntimeOptions) (_result *RestartDBNodeZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNodeId) {
		query["DBNodeId"] = request.DBNodeId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartDBNodeZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartDBNodeZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重启poalrdb边缘云集群节点
//
// @param request - RestartDBNodeZonalRequest
//
// @return RestartDBNodeZonalResponse
func (client *Client) RestartDBNodeZonal(request *RestartDBNodeZonalRequest) (_result *RestartDBNodeZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestartDBNodeZonalResponse{}
	_body, _err := client.RestartDBNodeZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores PolarDB databases and tables.
//
// @param request - RestoreTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreTableResponse
func (client *Client) RestoreTableWithOptions(request *RestoreTableRequest, runtime *dara.RuntimeOptions) (_result *RestoreTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TableMeta) {
		query["TableMeta"] = request.TableMeta
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestoreTable"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestoreTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores PolarDB databases and tables.
//
// @param request - RestoreTableRequest
//
// @return RestoreTableResponse
func (client *Client) RestoreTable(request *RestoreTableRequest) (_result *RestoreTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestoreTableResponse{}
	_body, _err := client.RestoreTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the access permissions on one or more databases from a specified PolarDB standard account.
//
// @param request - RevokeAccountPrivilegeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeAccountPrivilegeResponse
func (client *Client) RevokeAccountPrivilegeWithOptions(request *RevokeAccountPrivilegeRequest, runtime *dara.RuntimeOptions) (_result *RevokeAccountPrivilegeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeAccountPrivilege"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeAccountPrivilegeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the access permissions on one or more databases from a specified PolarDB standard account.
//
// @param request - RevokeAccountPrivilegeRequest
//
// @return RevokeAccountPrivilegeResponse
func (client *Client) RevokeAccountPrivilege(request *RevokeAccountPrivilegeRequest) (_result *RevokeAccountPrivilegeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeAccountPrivilegeResponse{}
	_body, _err := client.RevokeAccountPrivilegeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 撤销账号权限
//
// @param request - RevokeAccountPrivilegeZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeAccountPrivilegeZonalResponse
func (client *Client) RevokeAccountPrivilegeZonalWithOptions(request *RevokeAccountPrivilegeZonalRequest, runtime *dara.RuntimeOptions) (_result *RevokeAccountPrivilegeZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeAccountPrivilegeZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeAccountPrivilegeZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 撤销账号权限
//
// @param request - RevokeAccountPrivilegeZonalRequest
//
// @return RevokeAccountPrivilegeZonalResponse
func (client *Client) RevokeAccountPrivilegeZonal(request *RevokeAccountPrivilegeZonalRequest) (_result *RevokeAccountPrivilegeZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeAccountPrivilegeZonalResponse{}
	_body, _err := client.RevokeAccountPrivilegeZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 手动执行PolarClaw定时任务
//
// @param request - RunPolarClawCronJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunPolarClawCronJobResponse
func (client *Client) RunPolarClawCronJobWithOptions(request *RunPolarClawCronJobRequest, runtime *dara.RuntimeOptions) (_result *RunPolarClawCronJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunPolarClawCronJob"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunPolarClawCronJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 手动执行PolarClaw定时任务
//
// @param request - RunPolarClawCronJobRequest
//
// @return RunPolarClawCronJobResponse
func (client *Client) RunPolarClawCronJob(request *RunPolarClawCronJobRequest) (_result *RunPolarClawCronJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunPolarClawCronJobResponse{}
	_body, _err := client.RunPolarClawCronJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检索记忆
//
// @param request - SearchMemoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMemoriesResponse
func (client *Client) SearchMemoriesWithOptions(request *SearchMemoriesRequest, runtime *dara.RuntimeOptions) (_result *SearchMemoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.CreateTimeBegin) {
		query["CreateTimeBegin"] = request.CreateTimeBegin
	}

	if !dara.IsNil(request.CreateTimeEnd) {
		query["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.MemoryAgentId) {
		query["MemoryAgentId"] = request.MemoryAgentId
	}

	if !dara.IsNil(request.MemoryUserId) {
		query["MemoryUserId"] = request.MemoryUserId
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.TopK) {
		query["TopK"] = request.TopK
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMemories"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchMemoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检索记忆
//
// @param request - SearchMemoriesRequest
//
// @return SearchMemoriesResponse
func (client *Client) SearchMemories(request *SearchMemoriesRequest) (_result *SearchMemoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchMemoriesResponse{}
	_body, _err := client.SearchMemoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为目录配置配额或应用配额规则
//
// @param request - SetPolarFsFileQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPolarFsFileQuotaResponse
func (client *Client) SetPolarFsFileQuotaWithOptions(request *SetPolarFsFileQuotaRequest, runtime *dara.RuntimeOptions) (_result *SetPolarFsFileQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.FilePathQuotas) {
		query["FilePathQuotas"] = request.FilePathQuotas
	}

	if !dara.IsNil(request.PolarFsInstanceId) {
		query["PolarFsInstanceId"] = request.PolarFsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPolarFsFileQuota"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPolarFsFileQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为目录配置配额或应用配额规则
//
// @param request - SetPolarFsFileQuotaRequest
//
// @return SetPolarFsFileQuotaResponse
func (client *Client) SetPolarFsFileQuota(request *SetPolarFsFileQuotaRequest) (_result *SetPolarFsFileQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPolarFsFileQuotaResponse{}
	_body, _err := client.SetPolarFsFileQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Switches over the primary and secondary clusters in a global database network (GDN).
//
// @param request - SwitchOverGlobalDatabaseNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchOverGlobalDatabaseNetworkResponse
func (client *Client) SwitchOverGlobalDatabaseNetworkWithOptions(request *SwitchOverGlobalDatabaseNetworkRequest, runtime *dara.RuntimeOptions) (_result *SwitchOverGlobalDatabaseNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Forced) {
		query["Forced"] = request.Forced
	}

	if !dara.IsNil(request.GDNId) {
		query["GDNId"] = request.GDNId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchOverGlobalDatabaseNetwork"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchOverGlobalDatabaseNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches over the primary and secondary clusters in a global database network (GDN).
//
// @param request - SwitchOverGlobalDatabaseNetworkRequest
//
// @return SwitchOverGlobalDatabaseNetworkResponse
func (client *Client) SwitchOverGlobalDatabaseNetwork(request *SwitchOverGlobalDatabaseNetworkRequest) (_result *SwitchOverGlobalDatabaseNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchOverGlobalDatabaseNetworkResponse{}
	_body, _err := client.SwitchOverGlobalDatabaseNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates tags for a PolarDB cluster.
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
		Version:     dara.String("2017-08-01"),
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
// Creates tags for a PolarDB cluster.
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
// Temporarily changes the node configurations.
//
// @param request - TempModifyDBNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TempModifyDBNodeResponse
func (client *Client) TempModifyDBNodeWithOptions(request *TempModifyDBNodeRequest, runtime *dara.RuntimeOptions) (_result *TempModifyDBNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNode) {
		query["DBNode"] = request.DBNode
	}

	if !dara.IsNil(request.ModifyType) {
		query["ModifyType"] = request.ModifyType
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TempModifyDBNode"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TempModifyDBNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Temporarily changes the node configurations.
//
// @param request - TempModifyDBNodeRequest
//
// @return TempModifyDBNodeResponse
func (client *Client) TempModifyDBNode(request *TempModifyDBNodeRequest) (_result *TempModifyDBNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TempModifyDBNodeResponse{}
	_body, _err := client.TempModifyDBNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the billing method of a PolarDB cluster.
//
// Description:
//
// >
//
//   - PolarDB clusters support the subscription and pay-as-you-go billing methods. You can change the billing method from subscription to pay-as-you-go or from pay-as-you-go to subscription based on your business requirements. For more information, see [Change the billing method from subscription to pay-as-you-go](https://help.aliyun.com/document_detail/172886.html) and [Change the billing method from pay-as-you-go to subscription](https://help.aliyun.com/document_detail/84076.html).
//
//   - You cannot change the billing method from pay-as-you-go to subscription if your account balance is insufficient.
//
//   - If you change the billing method from subscription to pay-as-you-go, the system automatically refunds the balance of the prepaid subscription fees.
//
// @param request - TransformDBClusterPayTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransformDBClusterPayTypeResponse
func (client *Client) TransformDBClusterPayTypeWithOptions(request *TransformDBClusterPayTypeRequest, runtime *dara.RuntimeOptions) (_result *TransformDBClusterPayTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
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

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransformDBClusterPayType"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransformDBClusterPayTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the billing method of a PolarDB cluster.
//
// Description:
//
// >
//
//   - PolarDB clusters support the subscription and pay-as-you-go billing methods. You can change the billing method from subscription to pay-as-you-go or from pay-as-you-go to subscription based on your business requirements. For more information, see [Change the billing method from subscription to pay-as-you-go](https://help.aliyun.com/document_detail/172886.html) and [Change the billing method from pay-as-you-go to subscription](https://help.aliyun.com/document_detail/84076.html).
//
//   - You cannot change the billing method from pay-as-you-go to subscription if your account balance is insufficient.
//
//   - If you change the billing method from subscription to pay-as-you-go, the system automatically refunds the balance of the prepaid subscription fees.
//
// @param request - TransformDBClusterPayTypeRequest
//
// @return TransformDBClusterPayTypeResponse
func (client *Client) TransformDBClusterPayType(request *TransformDBClusterPayTypeRequest) (_result *TransformDBClusterPayTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TransformDBClusterPayTypeResponse{}
	_body, _err := client.TransformDBClusterPayTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑PolarClaw Agent通道
//
// @param request - UnbindPolarClawAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindPolarClawAgentResponse
func (client *Client) UnbindPolarClawAgentWithOptions(request *UnbindPolarClawAgentRequest, runtime *dara.RuntimeOptions) (_result *UnbindPolarClawAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.ChannelAccountId) {
		query["ChannelAccountId"] = request.ChannelAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindPolarClawAgent"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindPolarClawAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑PolarClaw Agent通道
//
// @param request - UnbindPolarClawAgentRequest
//
// @return UnbindPolarClawAgentResponse
func (client *Client) UnbindPolarClawAgent(request *UnbindPolarClawAgentRequest) (_result *UnbindPolarClawAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindPolarClawAgentResponse{}
	_body, _err := client.UnbindPolarClawAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 卸载PolarClaw Plugin
//
// @param request - UninstallPolarClawPluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallPolarClawPluginResponse
func (client *Client) UninstallPolarClawPluginWithOptions(request *UninstallPolarClawPluginRequest, runtime *dara.RuntimeOptions) (_result *UninstallPolarClawPluginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UninstallPolarClawPlugin"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallPolarClawPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 卸载PolarClaw Plugin
//
// @param request - UninstallPolarClawPluginRequest
//
// @return UninstallPolarClawPluginResponse
func (client *Client) UninstallPolarClawPlugin(request *UninstallPolarClawPluginRequest) (_result *UninstallPolarClawPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UninstallPolarClawPluginResponse{}
	_body, _err := client.UninstallPolarClawPluginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds tags from PolarDB clusters.
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
		Version:     dara.String("2017-08-01"),
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
// Unbinds tags from PolarDB clusters.
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
// 更新应用环境变量
//
// @param tmpReq - UpdateApplicationEnvironmentVariablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationEnvironmentVariablesResponse
func (client *Client) UpdateApplicationEnvironmentVariablesWithOptions(tmpReq *UpdateApplicationEnvironmentVariablesRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationEnvironmentVariablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateApplicationEnvironmentVariablesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Variables) {
		request.VariablesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Variables, dara.String("Variables"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	if !dara.IsNil(request.VariablesShrink) {
		query["Variables"] = request.VariablesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationEnvironmentVariables"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationEnvironmentVariablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新应用环境变量
//
// @param request - UpdateApplicationEnvironmentVariablesRequest
//
// @return UpdateApplicationEnvironmentVariablesResponse
func (client *Client) UpdateApplicationEnvironmentVariables(request *UpdateApplicationEnvironmentVariablesRequest) (_result *UpdateApplicationEnvironmentVariablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateApplicationEnvironmentVariablesResponse{}
	_body, _err := client.UpdateApplicationEnvironmentVariablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新插件
//
// @param request - UpdateExtensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExtensionsResponse
func (client *Client) UpdateExtensionsWithOptions(request *UpdateExtensionsRequest, runtime *dara.RuntimeOptions) (_result *UpdateExtensionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBNames) {
		query["DBNames"] = request.DBNames
	}

	if !dara.IsNil(request.Extensions) {
		query["Extensions"] = request.Extensions
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

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExtensions"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExtensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新插件
//
// @param request - UpdateExtensionsRequest
//
// @return UpdateExtensionsResponse
func (client *Client) UpdateExtensions(request *UpdateExtensionsRequest) (_result *UpdateExtensionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateExtensionsResponse{}
	_body, _err := client.UpdateExtensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新PolarClaw Agent
//
// @param tmpReq - UpdatePolarClawAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePolarClawAgentResponse
func (client *Client) UpdatePolarClawAgentWithOptions(tmpReq *UpdatePolarClawAgentRequest, runtime *dara.RuntimeOptions) (_result *UpdatePolarClawAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePolarClawAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Files) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, dara.String("Files"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.Avatar) {
		query["Avatar"] = request.Avatar
	}

	if !dara.IsNil(request.FilesShrink) {
		query["Files"] = request.FilesShrink
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	if !dara.IsNil(request.Workspace) {
		query["Workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePolarClawAgent"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePolarClawAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新PolarClaw Agent
//
// @param request - UpdatePolarClawAgentRequest
//
// @return UpdatePolarClawAgentResponse
func (client *Client) UpdatePolarClawAgent(request *UpdatePolarClawAgentRequest) (_result *UpdatePolarClawAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePolarClawAgentResponse{}
	_body, _err := client.UpdatePolarClawAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新PolarClaw定时任务
//
// @param tmpReq - UpdatePolarClawCronJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePolarClawCronJobResponse
func (client *Client) UpdatePolarClawCronJobWithOptions(tmpReq *UpdatePolarClawCronJobRequest, runtime *dara.RuntimeOptions) (_result *UpdatePolarClawCronJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePolarClawCronJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Delivery) {
		request.DeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Delivery, dara.String("Delivery"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FailureAlert) {
		request.FailureAlertShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FailureAlert, dara.String("FailureAlert"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Schedule) {
		request.ScheduleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Schedule, dara.String("Schedule"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.DeleteAfterRun) {
		query["DeleteAfterRun"] = request.DeleteAfterRun
	}

	if !dara.IsNil(request.DeliveryShrink) {
		query["Delivery"] = request.DeliveryShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.FailureAlertShrink) {
		query["FailureAlert"] = request.FailureAlertShrink
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PayloadShrink) {
		query["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	if !dara.IsNil(request.ScheduleShrink) {
		query["Schedule"] = request.ScheduleShrink
	}

	if !dara.IsNil(request.SessionKey) {
		query["SessionKey"] = request.SessionKey
	}

	if !dara.IsNil(request.SessionTarget) {
		query["SessionTarget"] = request.SessionTarget
	}

	if !dara.IsNil(request.WakeMode) {
		query["WakeMode"] = request.WakeMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePolarClawCronJob"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePolarClawCronJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新PolarClaw定时任务
//
// @param request - UpdatePolarClawCronJobRequest
//
// @return UpdatePolarClawCronJobResponse
func (client *Client) UpdatePolarClawCronJob(request *UpdatePolarClawCronJobRequest) (_result *UpdatePolarClawCronJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePolarClawCronJobResponse{}
	_body, _err := client.UpdatePolarClawCronJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades the kernel version of a PolarDB for MySQL cluster.
//
// Description:
//
// > 	- You can update only the revision version of a PolarDB for MySQL cluster, for example, from 8.0.1.1.3 to 8.0.1.1.4.
//
// >	- You can use only your Alibaba Cloud account to create scheduled tasks that update the kernel version of a PolarDB for MySQL cluster. RAM users are not authorized to update the kernel version of a PolarDB for MySQL cluster.
//
// @param request - UpgradeDBClusterVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeDBClusterVersionResponse
func (client *Client) UpgradeDBClusterVersionWithOptions(request *UpgradeDBClusterVersionRequest, runtime *dara.RuntimeOptions) (_result *UpgradeDBClusterVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.FromTimeService) {
		query["FromTimeService"] = request.FromTimeService
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlannedEndTime) {
		query["PlannedEndTime"] = request.PlannedEndTime
	}

	if !dara.IsNil(request.PlannedStartTime) {
		query["PlannedStartTime"] = request.PlannedStartTime
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TargetDBRevisionVersionCode) {
		query["TargetDBRevisionVersionCode"] = request.TargetDBRevisionVersionCode
	}

	if !dara.IsNil(request.TargetProxyRevisionVersionCode) {
		query["TargetProxyRevisionVersionCode"] = request.TargetProxyRevisionVersionCode
	}

	if !dara.IsNil(request.UpgradeLabel) {
		query["UpgradeLabel"] = request.UpgradeLabel
	}

	if !dara.IsNil(request.UpgradePolicy) {
		query["UpgradePolicy"] = request.UpgradePolicy
	}

	if !dara.IsNil(request.UpgradeType) {
		query["UpgradeType"] = request.UpgradeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeDBClusterVersion"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeDBClusterVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the kernel version of a PolarDB for MySQL cluster.
//
// Description:
//
// > 	- You can update only the revision version of a PolarDB for MySQL cluster, for example, from 8.0.1.1.3 to 8.0.1.1.4.
//
// >	- You can use only your Alibaba Cloud account to create scheduled tasks that update the kernel version of a PolarDB for MySQL cluster. RAM users are not authorized to update the kernel version of a PolarDB for MySQL cluster.
//
// @param request - UpgradeDBClusterVersionRequest
//
// @return UpgradeDBClusterVersionResponse
func (client *Client) UpgradeDBClusterVersion(request *UpgradeDBClusterVersionRequest) (_result *UpgradeDBClusterVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeDBClusterVersionResponse{}
	_body, _err := client.UpgradeDBClusterVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # PolarDB边缘云集群小版本升级
//
// @param request - UpgradeDBClusterVersionZonalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeDBClusterVersionZonalResponse
func (client *Client) UpgradeDBClusterVersionZonalWithOptions(request *UpgradeDBClusterVersionZonalRequest, runtime *dara.RuntimeOptions) (_result *UpgradeDBClusterVersionZonalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.FromTimeService) {
		query["FromTimeService"] = request.FromTimeService
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlannedEndTime) {
		query["PlannedEndTime"] = request.PlannedEndTime
	}

	if !dara.IsNil(request.PlannedStartTime) {
		query["PlannedStartTime"] = request.PlannedStartTime
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TargetDBRevisionVersionCode) {
		query["TargetDBRevisionVersionCode"] = request.TargetDBRevisionVersionCode
	}

	if !dara.IsNil(request.TargetProxyRevisionVersionCode) {
		query["TargetProxyRevisionVersionCode"] = request.TargetProxyRevisionVersionCode
	}

	if !dara.IsNil(request.UpgradeLabel) {
		query["UpgradeLabel"] = request.UpgradeLabel
	}

	if !dara.IsNil(request.UpgradePolicy) {
		query["UpgradePolicy"] = request.UpgradePolicy
	}

	if !dara.IsNil(request.UpgradeType) {
		query["UpgradeType"] = request.UpgradeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeDBClusterVersionZonal"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeDBClusterVersionZonalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # PolarDB边缘云集群小版本升级
//
// @param request - UpgradeDBClusterVersionZonalRequest
//
// @return UpgradeDBClusterVersionZonalResponse
func (client *Client) UpgradeDBClusterVersionZonal(request *UpgradeDBClusterVersionZonalRequest) (_result *UpgradeDBClusterVersionZonalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeDBClusterVersionZonalResponse{}
	_body, _err := client.UpgradeDBClusterVersionZonalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 升级PolarClaw Channel
//
// @param tmpReq - UpgradePolarClawChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradePolarClawChannelResponse
func (client *Client) UpgradePolarClawChannelWithOptions(tmpReq *UpgradePolarClawChannelRequest, runtime *dara.RuntimeOptions) (_result *UpgradePolarClawChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpgradePolarClawChannelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChannelConfig) {
		request.ChannelConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelConfig, dara.String("ChannelConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ChannelConfigShrink) {
		query["ChannelConfig"] = request.ChannelConfigShrink
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.NpmPackage) {
		query["NpmPackage"] = request.NpmPackage
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradePolarClawChannel"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradePolarClawChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 升级PolarClaw Channel
//
// @param request - UpgradePolarClawChannelRequest
//
// @return UpgradePolarClawChannelResponse
func (client *Client) UpgradePolarClawChannel(request *UpgradePolarClawChannelRequest) (_result *UpgradePolarClawChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradePolarClawChannelResponse{}
	_body, _err := client.UpgradePolarClawChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 升级PolarClaw Plugin
//
// @param request - UpgradePolarClawPluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradePolarClawPluginResponse
func (client *Client) UpgradePolarClawPluginWithOptions(request *UpgradePolarClawPluginRequest, runtime *dara.RuntimeOptions) (_result *UpgradePolarClawPluginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.NpmPackage) {
		query["NpmPackage"] = request.NpmPackage
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradePolarClawPlugin"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradePolarClawPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 升级PolarClaw Plugin
//
// @param request - UpgradePolarClawPluginRequest
//
// @return UpgradePolarClawPluginResponse
func (client *Client) UpgradePolarClawPlugin(request *UpgradePolarClawPluginRequest) (_result *UpgradePolarClawPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradePolarClawPluginResponse{}
	_body, _err := client.UpgradePolarClawPluginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 执行自定义命令
//
// @param request - UpgradePolarClawSkillsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradePolarClawSkillsResponse
func (client *Client) UpgradePolarClawSkillsWithOptions(request *UpgradePolarClawSkillsRequest, runtime *dara.RuntimeOptions) (_result *UpgradePolarClawSkillsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationType) {
		query["ApplicationType"] = request.ApplicationType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UpgradeMethod) {
		query["UpgradeMethod"] = request.UpgradeMethod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradePolarClawSkills"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradePolarClawSkillsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行自定义命令
//
// @param request - UpgradePolarClawSkillsRequest
//
// @return UpgradePolarClawSkillsResponse
func (client *Client) UpgradePolarClawSkills(request *UpgradePolarClawSkillsRequest) (_result *UpgradePolarClawSkillsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradePolarClawSkillsResponse{}
	_body, _err := client.UpgradePolarClawSkillsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) getPolarAgentWithSSE_opYieldFunc(_yield chan *GetPolarAgentResponse, _yieldErr chan error, request *GetPolarAgentRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExtraInfo) {
		query["ExtraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolarAgent"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}
