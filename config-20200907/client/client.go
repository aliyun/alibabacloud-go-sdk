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
	client.EndpointRule = dara.String("central")
	client.EndpointMap = map[string]*string{
		"cn-shanghai":    dara.String("config.cn-shanghai.aliyuncs.com"),
		"ap-southeast-1": dara.String("config.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("config"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 启用账号组规则
//
// @param request - ActiveAggregateConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActiveAggregateConfigRulesResponse
func (client *Client) ActiveAggregateConfigRulesWithOptions(request *ActiveAggregateConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *ActiveAggregateConfigRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.CompliancePackId) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !dara.IsNil(request.ConfigRuleIds) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActiveAggregateConfigRules"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActiveAggregateConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启用账号组规则
//
// @param request - ActiveAggregateConfigRulesRequest
//
// @return ActiveAggregateConfigRulesResponse
func (client *Client) ActiveAggregateConfigRules(request *ActiveAggregateConfigRulesRequest) (_result *ActiveAggregateConfigRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ActiveAggregateConfigRulesResponse{}
	_body, _err := client.ActiveAggregateConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启用指定规则
//
// @param request - ActiveConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActiveConfigRulesResponse
func (client *Client) ActiveConfigRulesWithOptions(request *ActiveConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *ActiveConfigRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompliancePackId) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !dara.IsNil(request.ConfigRuleIds) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActiveConfigRules"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActiveConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启用指定规则
//
// @param request - ActiveConfigRulesRequest
//
// @return ActiveConfigRulesResponse
func (client *Client) ActiveConfigRules(request *ActiveConfigRulesRequest) (_result *ActiveConfigRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ActiveConfigRulesResponse{}
	_body, _err := client.ActiveConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将指定账号组规则加入指定账号组合规包
//
// @param request - AttachAggregateConfigRuleToCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachAggregateConfigRuleToCompliancePackResponse
func (client *Client) AttachAggregateConfigRuleToCompliancePackWithOptions(request *AttachAggregateConfigRuleToCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *AttachAggregateConfigRuleToCompliancePackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.CompliancePackId) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !dara.IsNil(request.ConfigRuleIds) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachAggregateConfigRuleToCompliancePack"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachAggregateConfigRuleToCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将指定账号组规则加入指定账号组合规包
//
// @param request - AttachAggregateConfigRuleToCompliancePackRequest
//
// @return AttachAggregateConfigRuleToCompliancePackResponse
func (client *Client) AttachAggregateConfigRuleToCompliancePack(request *AttachAggregateConfigRuleToCompliancePackRequest) (_result *AttachAggregateConfigRuleToCompliancePackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachAggregateConfigRuleToCompliancePackResponse{}
	_body, _err := client.AttachAggregateConfigRuleToCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将指定规则加入指定合规包
//
// @param request - AttachConfigRuleToCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachConfigRuleToCompliancePackResponse
func (client *Client) AttachConfigRuleToCompliancePackWithOptions(request *AttachConfigRuleToCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *AttachConfigRuleToCompliancePackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompliancePackId) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !dara.IsNil(request.ConfigRuleIds) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachConfigRuleToCompliancePack"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachConfigRuleToCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将指定规则加入指定合规包
//
// @param request - AttachConfigRuleToCompliancePackRequest
//
// @return AttachConfigRuleToCompliancePackResponse
func (client *Client) AttachConfigRuleToCompliancePack(request *AttachConfigRuleToCompliancePackRequest) (_result *AttachConfigRuleToCompliancePackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachConfigRuleToCompliancePackResponse{}
	_body, _err := client.AttachConfigRuleToCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 合规包复制
//
// @param request - CopyCompliancePacksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyCompliancePacksResponse
func (client *Client) CopyCompliancePacksWithOptions(request *CopyCompliancePacksRequest, runtime *dara.RuntimeOptions) (_result *CopyCompliancePacksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesAggregatorIds) {
		query["DesAggregatorIds"] = request.DesAggregatorIds
	}

	if !dara.IsNil(request.SrcAggregatorId) {
		query["SrcAggregatorId"] = request.SrcAggregatorId
	}

	if !dara.IsNil(request.SrcCompliancePackIds) {
		query["SrcCompliancePackIds"] = request.SrcCompliancePackIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyCompliancePacks"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyCompliancePacksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合规包复制
//
// @param request - CopyCompliancePacksRequest
//
// @return CopyCompliancePacksResponse
func (client *Client) CopyCompliancePacks(request *CopyCompliancePacksRequest) (_result *CopyCompliancePacksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CopyCompliancePacksResponse{}
	_body, _err := client.CopyCompliancePacksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 规则复制
//
// @param request - CopyConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyConfigRulesResponse
func (client *Client) CopyConfigRulesWithOptions(request *CopyConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *CopyConfigRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesAggregatorIds) {
		query["DesAggregatorIds"] = request.DesAggregatorIds
	}

	if !dara.IsNil(request.SrcAggregatorId) {
		query["SrcAggregatorId"] = request.SrcAggregatorId
	}

	if !dara.IsNil(request.SrcConfigRuleIds) {
		query["SrcConfigRuleIds"] = request.SrcConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyConfigRules"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 规则复制
//
// @param request - CopyConfigRulesRequest
//
// @return CopyConfigRulesResponse
func (client *Client) CopyConfigRules(request *CopyConfigRulesRequest) (_result *CopyConfigRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CopyConfigRulesResponse{}
	_body, _err := client.CopyConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成当前账号搜索结果下载文件
//
// @param request - CreateAdvancedSearchFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAdvancedSearchFileResponse
func (client *Client) CreateAdvancedSearchFileWithOptions(request *CreateAdvancedSearchFileRequest, runtime *dara.RuntimeOptions) (_result *CreateAdvancedSearchFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Sql) {
		query["Sql"] = request.Sql
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAdvancedSearchFile"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAdvancedSearchFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成当前账号搜索结果下载文件
//
// @param request - CreateAdvancedSearchFileRequest
//
// @return CreateAdvancedSearchFileResponse
func (client *Client) CreateAdvancedSearchFile(request *CreateAdvancedSearchFileRequest) (_result *CreateAdvancedSearchFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAdvancedSearchFileResponse{}
	_body, _err := client.CreateAdvancedSearchFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a downloadable file of advanced search results for resources in an account group.
//
// Description:
//
// This topic provides an example of how to query for ECS instances in the account group `ca-edd3626622af00b3****` and create a downloadable file of the search results.
//
// @param request - CreateAggregateAdvancedSearchFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAggregateAdvancedSearchFileResponse
func (client *Client) CreateAggregateAdvancedSearchFileWithOptions(request *CreateAggregateAdvancedSearchFileRequest, runtime *dara.RuntimeOptions) (_result *CreateAggregateAdvancedSearchFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.Sql) {
		query["Sql"] = request.Sql
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAggregateAdvancedSearchFile"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAggregateAdvancedSearchFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a downloadable file of advanced search results for resources in an account group.
//
// Description:
//
// This topic provides an example of how to query for ECS instances in the account group `ca-edd3626622af00b3****` and create a downloadable file of the search results.
//
// @param request - CreateAggregateAdvancedSearchFileRequest
//
// @return CreateAggregateAdvancedSearchFileResponse
func (client *Client) CreateAggregateAdvancedSearchFile(request *CreateAggregateAdvancedSearchFileRequest) (_result *CreateAggregateAdvancedSearchFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAggregateAdvancedSearchFileResponse{}
	_body, _err := client.CreateAggregateAdvancedSearchFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a compliance pack for a specified account group.
//
// Description:
//
// A compliance pack is a collection of rules. When you create a compliance pack, you can select default rules from a compliance pack template. You can also select rules from rule templates and the list of existing rules.
//
// After a compliance pack is created, its rules are evaluated once by default. Subsequent evaluations are automatically triggered based on the trigger mechanism of the rules. You can also manually trigger an evaluation.
//
// A compliance pack template is a collection of rules created by CloudConfig for a specific compliance scenario.
//
// This topic provides an example of how to create a compliance pack for the account group `ca-f632626622af0079****` using the compliance pack template `ClassifiedProtectionPreCheck` (Classified Protection Level 3 Pre-check).
//
// @param tmpReq - CreateAggregateCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAggregateCompliancePackResponse
func (client *Client) CreateAggregateCompliancePackWithOptions(tmpReq *CreateAggregateCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *CreateAggregateCompliancePackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAggregateCompliancePackShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ConfigRules) {
		request.ConfigRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfigRules, dara.String("ConfigRules"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CompliancePackName) {
		body["CompliancePackName"] = request.CompliancePackName
	}

	if !dara.IsNil(request.CompliancePackTemplateId) {
		body["CompliancePackTemplateId"] = request.CompliancePackTemplateId
	}

	if !dara.IsNil(request.ConfigRulesShrink) {
		body["ConfigRules"] = request.ConfigRulesShrink
	}

	if !dara.IsNil(request.DefaultEnable) {
		body["DefaultEnable"] = request.DefaultEnable
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExcludeRegionIdsScope) {
		body["ExcludeRegionIdsScope"] = request.ExcludeRegionIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceGroupIdsScope) {
		body["ExcludeResourceGroupIdsScope"] = request.ExcludeResourceGroupIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceIdsScope) {
		body["ExcludeResourceIdsScope"] = request.ExcludeResourceIdsScope
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ExcludeTagsScope) {
		bodyFlat["ExcludeTagsScope"] = request.ExcludeTagsScope
	}

	if !dara.IsNil(request.RegionIdsScope) {
		body["RegionIdsScope"] = request.RegionIdsScope
	}

	if !dara.IsNil(request.ResourceGroupIdsScope) {
		body["ResourceGroupIdsScope"] = request.ResourceGroupIdsScope
	}

	if !dara.IsNil(request.ResourceIdsScope) {
		body["ResourceIdsScope"] = request.ResourceIdsScope
	}

	if !dara.IsNil(request.RiskLevel) {
		body["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.TagKeyScope) {
		body["TagKeyScope"] = request.TagKeyScope
	}

	if !dara.IsNil(request.TagValueScope) {
		body["TagValueScope"] = request.TagValueScope
	}

	if !dara.IsNil(request.TagsScope) {
		bodyFlat["TagsScope"] = request.TagsScope
	}

	if !dara.IsNil(request.TemplateContent) {
		body["TemplateContent"] = request.TemplateContent
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAggregateCompliancePack"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAggregateCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a compliance pack for a specified account group.
//
// Description:
//
// A compliance pack is a collection of rules. When you create a compliance pack, you can select default rules from a compliance pack template. You can also select rules from rule templates and the list of existing rules.
//
// After a compliance pack is created, its rules are evaluated once by default. Subsequent evaluations are automatically triggered based on the trigger mechanism of the rules. You can also manually trigger an evaluation.
//
// A compliance pack template is a collection of rules created by CloudConfig for a specific compliance scenario.
//
// This topic provides an example of how to create a compliance pack for the account group `ca-f632626622af0079****` using the compliance pack template `ClassifiedProtectionPreCheck` (Classified Protection Level 3 Pre-check).
//
// @param request - CreateAggregateCompliancePackRequest
//
// @return CreateAggregateCompliancePackResponse
func (client *Client) CreateAggregateCompliancePack(request *CreateAggregateCompliancePackRequest) (_result *CreateAggregateCompliancePackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAggregateCompliancePackResponse{}
	_body, _err := client.CreateAggregateCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a delivery channel for a specified account group to deliver resource data to Simple Log Service (SLS), Object Storage Service (OSS), or Simple Message Queue (formerly MNS).
//
// Description:
//
// ### Background information
//
// - Deliver to Simple Log Service (SLS)
//
//	To deliver configuration history, non-compliance events, and scheduled resource snapshots to a Logstore in Simple Log Service (SLS), you must first create a log project and a Logstore. This lets you query and analyze logs. For examples of the content in JSON format, see [Example of resource configuration change history](https://help.aliyun.com/document_detail/308347.html), [Example of resource non-compliance event](https://help.aliyun.com/document_detail/307122.html), and [Example of scheduled resource snapshot](https://help.aliyun.com/document_detail/611894.html).
//
// - Deliver to Object Storage Service (OSS)
//
//	To deliver scheduled resource snapshots or configuration history to a specified location in Object Storage Service (OSS), you must first create a bucket. This lets you view or download files in JSON format. For examples of the content in JSON format, see [Example of scheduled resource snapshot](https://help.aliyun.com/document_detail/305669.html) and [Example of resource configuration change history](https://help.aliyun.com/document_detail/189738.html).
//
// - Deliver to Simple Message Queue (MNS)
//
//	To deliver resource configuration change history and resource non-compliance events to a specified topic in Simple Message Queue (formerly MNS), you must first create a topic. This lets you configure the push method and content for the topic. For examples of the content in JSON format, see [Example of resource configuration change history](https://help.aliyun.com/document_detail/309462.html) and [Example of resource non-compliance event](https://help.aliyun.com/document_detail/309463.html).
//
// ### Limits
//
// You can create a maximum of five delivery channels for each account group.
//
// ### Usage notes
//
// This example shows how to create a delivery channel of the `OSS` type for the account group `ca-a4e5626622af0079****`. The Amazon Resource Name (ARN) of the delivery destination is `acs:oss:cn-shanghai:100931896542****:new-bucket`. The response shows that the delivery channel is created and its ID is `cdc-8e45ff4e06a3a8****`.
//
// @param request - CreateAggregateConfigDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAggregateConfigDeliveryChannelResponse
func (client *Client) CreateAggregateConfigDeliveryChannelWithOptions(request *CreateAggregateConfigDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *CreateAggregateConfigDeliveryChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CompliantSnapshot) {
		query["CompliantSnapshot"] = request.CompliantSnapshot
	}

	if !dara.IsNil(request.ConfigurationItemChangeNotification) {
		query["ConfigurationItemChangeNotification"] = request.ConfigurationItemChangeNotification
	}

	if !dara.IsNil(request.ConfigurationSnapshot) {
		query["ConfigurationSnapshot"] = request.ConfigurationSnapshot
	}

	if !dara.IsNil(request.DeliveryChannelCondition) {
		query["DeliveryChannelCondition"] = request.DeliveryChannelCondition
	}

	if !dara.IsNil(request.DeliveryChannelName) {
		query["DeliveryChannelName"] = request.DeliveryChannelName
	}

	if !dara.IsNil(request.DeliveryChannelTargetArn) {
		query["DeliveryChannelTargetArn"] = request.DeliveryChannelTargetArn
	}

	if !dara.IsNil(request.DeliveryChannelType) {
		query["DeliveryChannelType"] = request.DeliveryChannelType
	}

	if !dara.IsNil(request.DeliverySnapshotTime) {
		query["DeliverySnapshotTime"] = request.DeliverySnapshotTime
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.NonCompliantNotification) {
		query["NonCompliantNotification"] = request.NonCompliantNotification
	}

	if !dara.IsNil(request.OversizedDataOSSTargetArn) {
		query["OversizedDataOSSTargetArn"] = request.OversizedDataOSSTargetArn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAggregateConfigDeliveryChannel"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAggregateConfigDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a delivery channel for a specified account group to deliver resource data to Simple Log Service (SLS), Object Storage Service (OSS), or Simple Message Queue (formerly MNS).
//
// Description:
//
// ### Background information
//
// - Deliver to Simple Log Service (SLS)
//
//	To deliver configuration history, non-compliance events, and scheduled resource snapshots to a Logstore in Simple Log Service (SLS), you must first create a log project and a Logstore. This lets you query and analyze logs. For examples of the content in JSON format, see [Example of resource configuration change history](https://help.aliyun.com/document_detail/308347.html), [Example of resource non-compliance event](https://help.aliyun.com/document_detail/307122.html), and [Example of scheduled resource snapshot](https://help.aliyun.com/document_detail/611894.html).
//
// - Deliver to Object Storage Service (OSS)
//
//	To deliver scheduled resource snapshots or configuration history to a specified location in Object Storage Service (OSS), you must first create a bucket. This lets you view or download files in JSON format. For examples of the content in JSON format, see [Example of scheduled resource snapshot](https://help.aliyun.com/document_detail/305669.html) and [Example of resource configuration change history](https://help.aliyun.com/document_detail/189738.html).
//
// - Deliver to Simple Message Queue (MNS)
//
//	To deliver resource configuration change history and resource non-compliance events to a specified topic in Simple Message Queue (formerly MNS), you must first create a topic. This lets you configure the push method and content for the topic. For examples of the content in JSON format, see [Example of resource configuration change history](https://help.aliyun.com/document_detail/309462.html) and [Example of resource non-compliance event](https://help.aliyun.com/document_detail/309463.html).
//
// ### Limits
//
// You can create a maximum of five delivery channels for each account group.
//
// ### Usage notes
//
// This example shows how to create a delivery channel of the `OSS` type for the account group `ca-a4e5626622af0079****`. The Amazon Resource Name (ARN) of the delivery destination is `acs:oss:cn-shanghai:100931896542****:new-bucket`. The response shows that the delivery channel is created and its ID is `cdc-8e45ff4e06a3a8****`.
//
// @param request - CreateAggregateConfigDeliveryChannelRequest
//
// @return CreateAggregateConfigDeliveryChannelResponse
func (client *Client) CreateAggregateConfigDeliveryChannel(request *CreateAggregateConfigDeliveryChannelRequest) (_result *CreateAggregateConfigDeliveryChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAggregateConfigDeliveryChannelResponse{}
	_body, _err := client.CreateAggregateConfigDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a rule for a specified account group. You can create a rule from a template or create a custom rule using Function Compute. The rule checks your resources for compliance. After a rule is created, it automatically runs once. Cloud Config then runs evaluations based on the rule\\"s trigger. You can also run evaluations manually.
//
// Description:
//
// ### Limits
//
// Each management account and delegated administrator account can have up to 200 rules.
//
// ### Background information
//
// Cloud Config supports the following methods for creating rules:
//
// - Create rules from templates
//
//	Rule templates are predefined rule functions that Cloud Config provides in Function Compute (FC). Use rule templates to create rules quickly. For more information about rules, see [the referenced document](https://help.aliyun.com/document_detail/128273.html).
//
// - Create rules based on functions in Function Compute
//
//	Custom function rules are rules whose code is hosted in FC functions. If the predefined rule templates in Cloud Config do not meet your compliance requirements, write function code to check compliance in complex scenarios. For more information about custom function rules, see [the referenced document](https://help.aliyun.com/document_detail/127405.html).
//
// ### Usage notes
//
// This example shows how to create a rule for the account group `ca-a4e5626622af0079****` using the \\`required-tags\\` template. The response shows that the rule is created with the ID `cr-4e3d626622af0080****`.
//
// @param tmpReq - CreateAggregateConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAggregateConfigRuleResponse
func (client *Client) CreateAggregateConfigRuleWithOptions(tmpReq *CreateAggregateConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateAggregateConfigRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAggregateConfigRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InputParameters) {
		request.InputParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InputParameters, dara.String("InputParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceTypesScope) {
		request.ResourceTypesScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceTypesScope, dara.String("ResourceTypesScope"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceNameScope) {
		query["ResourceNameScope"] = request.ResourceNameScope
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountIdsScope) {
		body["AccountIdsScope"] = request.AccountIdsScope
	}

	if !dara.IsNil(request.AggregatorId) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Conditions) {
		body["Conditions"] = request.Conditions
	}

	if !dara.IsNil(request.ConfigRuleName) {
		body["ConfigRuleName"] = request.ConfigRuleName
	}

	if !dara.IsNil(request.ConfigRuleTriggerTypes) {
		body["ConfigRuleTriggerTypes"] = request.ConfigRuleTriggerTypes
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExcludeAccountIdsScope) {
		body["ExcludeAccountIdsScope"] = request.ExcludeAccountIdsScope
	}

	if !dara.IsNil(request.ExcludeFolderIdsScope) {
		body["ExcludeFolderIdsScope"] = request.ExcludeFolderIdsScope
	}

	if !dara.IsNil(request.ExcludeRegionIdsScope) {
		body["ExcludeRegionIdsScope"] = request.ExcludeRegionIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceGroupIdsScope) {
		body["ExcludeResourceGroupIdsScope"] = request.ExcludeResourceGroupIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceIdsScope) {
		body["ExcludeResourceIdsScope"] = request.ExcludeResourceIdsScope
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ExcludeTagsScope) {
		bodyFlat["ExcludeTagsScope"] = request.ExcludeTagsScope
	}

	if !dara.IsNil(request.ExtendContent) {
		body["ExtendContent"] = request.ExtendContent
	}

	if !dara.IsNil(request.FolderIdsScope) {
		body["FolderIdsScope"] = request.FolderIdsScope
	}

	if !dara.IsNil(request.InputParametersShrink) {
		body["InputParameters"] = request.InputParametersShrink
	}

	if !dara.IsNil(request.MaximumExecutionFrequency) {
		body["MaximumExecutionFrequency"] = request.MaximumExecutionFrequency
	}

	if !dara.IsNil(request.RegionIdsScope) {
		body["RegionIdsScope"] = request.RegionIdsScope
	}

	if !dara.IsNil(request.ResourceGroupIdsScope) {
		body["ResourceGroupIdsScope"] = request.ResourceGroupIdsScope
	}

	if !dara.IsNil(request.ResourceIdsScope) {
		body["ResourceIdsScope"] = request.ResourceIdsScope
	}

	if !dara.IsNil(request.ResourceTypesScopeShrink) {
		body["ResourceTypesScope"] = request.ResourceTypesScopeShrink
	}

	if !dara.IsNil(request.RiskLevel) {
		body["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.SourceIdentifier) {
		body["SourceIdentifier"] = request.SourceIdentifier
	}

	if !dara.IsNil(request.SourceOwner) {
		body["SourceOwner"] = request.SourceOwner
	}

	if !dara.IsNil(request.TagKeyLogicScope) {
		body["TagKeyLogicScope"] = request.TagKeyLogicScope
	}

	if !dara.IsNil(request.TagKeyScope) {
		body["TagKeyScope"] = request.TagKeyScope
	}

	if !dara.IsNil(request.TagValueScope) {
		body["TagValueScope"] = request.TagValueScope
	}

	if !dara.IsNil(request.TagsScope) {
		bodyFlat["TagsScope"] = request.TagsScope
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAggregateConfigRule"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAggregateConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a rule for a specified account group. You can create a rule from a template or create a custom rule using Function Compute. The rule checks your resources for compliance. After a rule is created, it automatically runs once. Cloud Config then runs evaluations based on the rule\\"s trigger. You can also run evaluations manually.
//
// Description:
//
// ### Limits
//
// Each management account and delegated administrator account can have up to 200 rules.
//
// ### Background information
//
// Cloud Config supports the following methods for creating rules:
//
// - Create rules from templates
//
//	Rule templates are predefined rule functions that Cloud Config provides in Function Compute (FC). Use rule templates to create rules quickly. For more information about rules, see [the referenced document](https://help.aliyun.com/document_detail/128273.html).
//
// - Create rules based on functions in Function Compute
//
//	Custom function rules are rules whose code is hosted in FC functions. If the predefined rule templates in Cloud Config do not meet your compliance requirements, write function code to check compliance in complex scenarios. For more information about custom function rules, see [the referenced document](https://help.aliyun.com/document_detail/127405.html).
//
// ### Usage notes
//
// This example shows how to create a rule for the account group `ca-a4e5626622af0079****` using the \\`required-tags\\` template. The response shows that the rule is created with the ID `cr-4e3d626622af0080****`.
//
// @param request - CreateAggregateConfigRuleRequest
//
// @return CreateAggregateConfigRuleResponse
func (client *Client) CreateAggregateConfigRule(request *CreateAggregateConfigRuleRequest) (_result *CreateAggregateConfigRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAggregateConfigRuleResponse{}
	_body, _err := client.CreateAggregateConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a remediation for a rule in a specified account group.
//
// Description:
//
// ### Background information
//
// - Template-based remediation: Uses public templates provided by Operation Orchestration Service (OOS) to quickly remediate non-compliant resources.
//
//	Only one remediation can be created for a rule. This type of remediation is supported only for rules that are created from specific templates.
//
// - Custom remediation: Runs custom code in Function Compute (FC) to quickly remediate non-compliant resources.
//
//	Only one remediation can be created for a rule. This type of remediation is supported for rules created from templates and for custom rules.
//
// ### Usage notes
//
// This topic provides an example of how to create a remediation for the rule `cr-6b7c626622af00b4****` in the account group `ca-6b4a626622af0012****`. The response shows that the remediation is created and its ID is `crr-909ba2d4716700eb****`.
//
// @param request - CreateAggregateRemediationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAggregateRemediationResponse
func (client *Client) CreateAggregateRemediationWithOptions(request *CreateAggregateRemediationRequest, runtime *dara.RuntimeOptions) (_result *CreateAggregateRemediationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigRuleId) {
		body["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.InvokeType) {
		body["InvokeType"] = request.InvokeType
	}

	if !dara.IsNil(request.Params) {
		body["Params"] = request.Params
	}

	if !dara.IsNil(request.RemediationTemplateId) {
		body["RemediationTemplateId"] = request.RemediationTemplateId
	}

	if !dara.IsNil(request.RemediationType) {
		body["RemediationType"] = request.RemediationType
	}

	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAggregateRemediation"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAggregateRemediationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a remediation for a rule in a specified account group.
//
// Description:
//
// ### Background information
//
// - Template-based remediation: Uses public templates provided by Operation Orchestration Service (OOS) to quickly remediate non-compliant resources.
//
//	Only one remediation can be created for a rule. This type of remediation is supported only for rules that are created from specific templates.
//
// - Custom remediation: Runs custom code in Function Compute (FC) to quickly remediate non-compliant resources.
//
//	Only one remediation can be created for a rule. This type of remediation is supported for rules created from templates and for custom rules.
//
// ### Usage notes
//
// This topic provides an example of how to create a remediation for the rule `cr-6b7c626622af00b4****` in the account group `ca-6b4a626622af0012****`. The response shows that the remediation is created and its ID is `crr-909ba2d4716700eb****`.
//
// @param request - CreateAggregateRemediationRequest
//
// @return CreateAggregateRemediationResponse
func (client *Client) CreateAggregateRemediation(request *CreateAggregateRemediationRequest) (_result *CreateAggregateRemediationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAggregateRemediationResponse{}
	_body, _err := client.CreateAggregateRemediationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// A management account or a delegated administrator account in a resource directory can create an account group to centrally manage resources, compliance packages, and rules across multiple member accounts.
//
// Description:
//
// ### Limits
//
// A management account or a delegated administrator account can create a maximum of 5 account groups. Each account group can contain a maximum of 200 member accounts.
//
// ### Background information
//
// For more information about account groups, including their concepts, use cases, and the impact of member account changes on Cloud Config, see [Overview](https://help.aliyun.com/document_detail/211534.html).
//
// Cloud Config supports the following types of account groups:
//
// - Global account group: A global account group contains all members in a resource directory and automatically synchronizes member changes. A management account or a delegated administrator account can create only one global account group.
//
// - Custom account group: To create a custom account group, a management account or a delegated administrator account selects some or all member accounts from the resource directory.
//
//   - If a new member is added to the resource directory, the change is not automatically synchronized. The management account or delegated administrator account must manually add the new member to the account group.
//
//   - If a member is removed from the resource directory, the management account or delegated administrator account loses the permissions to manage that member\\"s compliance. The custom account group automatically detects this change and removes the member from the group.
//
// - Folder account group: When an account group is created from a folder, it automatically detects and synchronizes changes to the members within that folder. The members in a folder account group are always consistent with the members in the selected folder.
//
//	A management account or a delegated administrator account can select only one non-empty folder to create a folder account group.
//
// ### Usage notes
//
// This topic provides an example of how to use a management account to create a custom account group of the `CUSTOM` type. The account group is named `Test_Group` and has the description `Test Group`. The member accounts are as follows:
//
// - The member account ID is `171322098523****` and the member account name is `Alice`.
//
// - The member account ID is `100532098349****` and the member account name is `Tom`.
//
// @param tmpReq - CreateAggregatorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAggregatorResponse
func (client *Client) CreateAggregatorWithOptions(tmpReq *CreateAggregatorRequest, runtime *dara.RuntimeOptions) (_result *CreateAggregatorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAggregatorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AggregatorAccounts) {
		request.AggregatorAccountsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AggregatorAccounts, dara.String("AggregatorAccounts"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorAccountsShrink) {
		body["AggregatorAccounts"] = request.AggregatorAccountsShrink
	}

	if !dara.IsNil(request.AggregatorName) {
		body["AggregatorName"] = request.AggregatorName
	}

	if !dara.IsNil(request.AggregatorType) {
		body["AggregatorType"] = request.AggregatorType
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAggregator"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAggregatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// A management account or a delegated administrator account in a resource directory can create an account group to centrally manage resources, compliance packages, and rules across multiple member accounts.
//
// Description:
//
// ### Limits
//
// A management account or a delegated administrator account can create a maximum of 5 account groups. Each account group can contain a maximum of 200 member accounts.
//
// ### Background information
//
// For more information about account groups, including their concepts, use cases, and the impact of member account changes on Cloud Config, see [Overview](https://help.aliyun.com/document_detail/211534.html).
//
// Cloud Config supports the following types of account groups:
//
// - Global account group: A global account group contains all members in a resource directory and automatically synchronizes member changes. A management account or a delegated administrator account can create only one global account group.
//
// - Custom account group: To create a custom account group, a management account or a delegated administrator account selects some or all member accounts from the resource directory.
//
//   - If a new member is added to the resource directory, the change is not automatically synchronized. The management account or delegated administrator account must manually add the new member to the account group.
//
//   - If a member is removed from the resource directory, the management account or delegated administrator account loses the permissions to manage that member\\"s compliance. The custom account group automatically detects this change and removes the member from the group.
//
// - Folder account group: When an account group is created from a folder, it automatically detects and synchronizes changes to the members within that folder. The members in a folder account group are always consistent with the members in the selected folder.
//
//	A management account or a delegated administrator account can select only one non-empty folder to create a folder account group.
//
// ### Usage notes
//
// This topic provides an example of how to use a management account to create a custom account group of the `CUSTOM` type. The account group is named `Test_Group` and has the description `Test Group`. The member accounts are as follows:
//
// - The member account ID is `171322098523****` and the member account name is `Alice`.
//
// - The member account ID is `100532098349****` and the member account name is `Tom`.
//
// @param request - CreateAggregatorRequest
//
// @return CreateAggregatorResponse
func (client *Client) CreateAggregator(request *CreateAggregatorRequest) (_result *CreateAggregatorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAggregatorResponse{}
	_body, _err := client.CreateAggregatorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为当前账号创建合规包
//
// @param tmpReq - CreateCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCompliancePackResponse
func (client *Client) CreateCompliancePackWithOptions(tmpReq *CreateCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *CreateCompliancePackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCompliancePackShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ConfigRules) {
		request.ConfigRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfigRules, dara.String("ConfigRules"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CompliancePackName) {
		body["CompliancePackName"] = request.CompliancePackName
	}

	if !dara.IsNil(request.CompliancePackTemplateId) {
		body["CompliancePackTemplateId"] = request.CompliancePackTemplateId
	}

	if !dara.IsNil(request.ConfigRulesShrink) {
		body["ConfigRules"] = request.ConfigRulesShrink
	}

	if !dara.IsNil(request.DefaultEnable) {
		body["DefaultEnable"] = request.DefaultEnable
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExcludeRegionIdsScope) {
		body["ExcludeRegionIdsScope"] = request.ExcludeRegionIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceGroupIdsScope) {
		body["ExcludeResourceGroupIdsScope"] = request.ExcludeResourceGroupIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceIdsScope) {
		body["ExcludeResourceIdsScope"] = request.ExcludeResourceIdsScope
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ExcludeTagsScope) {
		bodyFlat["ExcludeTagsScope"] = request.ExcludeTagsScope
	}

	if !dara.IsNil(request.RegionIdsScope) {
		body["RegionIdsScope"] = request.RegionIdsScope
	}

	if !dara.IsNil(request.ResourceGroupIdsScope) {
		body["ResourceGroupIdsScope"] = request.ResourceGroupIdsScope
	}

	if !dara.IsNil(request.ResourceIdsScope) {
		body["ResourceIdsScope"] = request.ResourceIdsScope
	}

	if !dara.IsNil(request.RiskLevel) {
		body["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.TagKeyScope) {
		body["TagKeyScope"] = request.TagKeyScope
	}

	if !dara.IsNil(request.TagValueScope) {
		body["TagValueScope"] = request.TagValueScope
	}

	if !dara.IsNil(request.TagsScope) {
		bodyFlat["TagsScope"] = request.TagsScope
	}

	if !dara.IsNil(request.TemplateContent) {
		body["TemplateContent"] = request.TemplateContent
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCompliancePack"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为当前账号创建合规包
//
// @param request - CreateCompliancePackRequest
//
// @return CreateCompliancePackResponse
func (client *Client) CreateCompliancePack(request *CreateCompliancePackRequest) (_result *CreateCompliancePackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCompliancePackResponse{}
	_body, _err := client.CreateCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a delivery channel to deliver resource data to Simple Log Service (SLS), Object Storage Service (OSS), or Simple Message Queue (formerly MNS).
//
// Description:
//
// ### Background information
//
// - Deliver to Simple Log Service (SLS)
//
//	When you deliver resource configuration histories, non-compliant events, and scheduled snapshots to a Logstore in SLS, you must first create a project and a Logstore. This lets you query and analyze logs. For examples of the content of the JSON files, see [Resource configuration history examples](https://help.aliyun.com/document_detail/308347.html), [Non-compliant event examples](https://help.aliyun.com/document_detail/307122.html), and [Scheduled resource snapshot examples](https://help.aliyun.com/document_detail/611894.html).
//
// - Deliver to Object Storage Service (OSS)
//
//	When you deliver scheduled resource snapshots or configuration histories to a specified location in OSS, you must first create a bucket. This lets you view or download the JSON files. For examples of the content of the JSON files, see [Scheduled resource snapshot examples](https://help.aliyun.com/document_detail/305669.html) and [Resource configuration history examples](https://help.aliyun.com/document_detail/189738.html).
//
// - Deliver to Simple Message Queue (formerly MNS)
//
//	When you deliver resource configuration histories and non-compliant events to a specified topic in MNS, you must first create a topic. This lets you set the push method and content for the topic. For examples of the content of the JSON files, see [Resource configuration history examples](https://help.aliyun.com/document_detail/309462.html) and [Non-compliant event examples](https://help.aliyun.com/document_detail/309463.html).
//
// ### Limits
//
// You can create a maximum of 5 delivery channels.
//
// ### Usage notes
//
// This topic provides an example of how to create a delivery channel. In this example, the channel type is `OSS` and the destination ARN is `acs:oss:cn-shanghai:100931896542****:new-bucket`. The response shows that a delivery channel with the ID `cdc-8e45ff4e06a3a8****` is created.
//
// @param request - CreateConfigDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConfigDeliveryChannelResponse
func (client *Client) CreateConfigDeliveryChannelWithOptions(request *CreateConfigDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *CreateConfigDeliveryChannelResponse, _err error) {
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

	if !dara.IsNil(request.CompliantSnapshot) {
		query["CompliantSnapshot"] = request.CompliantSnapshot
	}

	if !dara.IsNil(request.ConfigurationItemChangeNotification) {
		query["ConfigurationItemChangeNotification"] = request.ConfigurationItemChangeNotification
	}

	if !dara.IsNil(request.ConfigurationSnapshot) {
		query["ConfigurationSnapshot"] = request.ConfigurationSnapshot
	}

	if !dara.IsNil(request.DeliveryChannelCondition) {
		query["DeliveryChannelCondition"] = request.DeliveryChannelCondition
	}

	if !dara.IsNil(request.DeliveryChannelName) {
		query["DeliveryChannelName"] = request.DeliveryChannelName
	}

	if !dara.IsNil(request.DeliveryChannelTargetArn) {
		query["DeliveryChannelTargetArn"] = request.DeliveryChannelTargetArn
	}

	if !dara.IsNil(request.DeliveryChannelType) {
		query["DeliveryChannelType"] = request.DeliveryChannelType
	}

	if !dara.IsNil(request.DeliverySnapshotTime) {
		query["DeliverySnapshotTime"] = request.DeliverySnapshotTime
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.NonCompliantNotification) {
		query["NonCompliantNotification"] = request.NonCompliantNotification
	}

	if !dara.IsNil(request.OversizedDataOSSTargetArn) {
		query["OversizedDataOSSTargetArn"] = request.OversizedDataOSSTargetArn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConfigDeliveryChannel"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConfigDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a delivery channel to deliver resource data to Simple Log Service (SLS), Object Storage Service (OSS), or Simple Message Queue (formerly MNS).
//
// Description:
//
// ### Background information
//
// - Deliver to Simple Log Service (SLS)
//
//	When you deliver resource configuration histories, non-compliant events, and scheduled snapshots to a Logstore in SLS, you must first create a project and a Logstore. This lets you query and analyze logs. For examples of the content of the JSON files, see [Resource configuration history examples](https://help.aliyun.com/document_detail/308347.html), [Non-compliant event examples](https://help.aliyun.com/document_detail/307122.html), and [Scheduled resource snapshot examples](https://help.aliyun.com/document_detail/611894.html).
//
// - Deliver to Object Storage Service (OSS)
//
//	When you deliver scheduled resource snapshots or configuration histories to a specified location in OSS, you must first create a bucket. This lets you view or download the JSON files. For examples of the content of the JSON files, see [Scheduled resource snapshot examples](https://help.aliyun.com/document_detail/305669.html) and [Resource configuration history examples](https://help.aliyun.com/document_detail/189738.html).
//
// - Deliver to Simple Message Queue (formerly MNS)
//
//	When you deliver resource configuration histories and non-compliant events to a specified topic in MNS, you must first create a topic. This lets you set the push method and content for the topic. For examples of the content of the JSON files, see [Resource configuration history examples](https://help.aliyun.com/document_detail/309462.html) and [Non-compliant event examples](https://help.aliyun.com/document_detail/309463.html).
//
// ### Limits
//
// You can create a maximum of 5 delivery channels.
//
// ### Usage notes
//
// This topic provides an example of how to create a delivery channel. In this example, the channel type is `OSS` and the destination ARN is `acs:oss:cn-shanghai:100931896542****:new-bucket`. The response shows that a delivery channel with the ID `cdc-8e45ff4e06a3a8****` is created.
//
// @param request - CreateConfigDeliveryChannelRequest
//
// @return CreateConfigDeliveryChannelResponse
func (client *Client) CreateConfigDeliveryChannel(request *CreateConfigDeliveryChannelRequest) (_result *CreateConfigDeliveryChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateConfigDeliveryChannelResponse{}
	_body, _err := client.CreateConfigDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a rule from a template or a custom rule using Function Compute to check resource compliance. After you create a rule, Cloud Config runs an initial evaluation and then automatically triggers subsequent evaluations based on the rule\\"s trigger. You can also run evaluations manually.
//
// Description:
//
// ### Limits
//
// You can create up to 200 rules per account.
//
// ### Background information
//
// You can create rules in Cloud Config in two ways:
//
// - Create rules from templates
//
//	Rule templates are predefined rule functions provided by Cloud Config in Function Compute. You can use rule templates to quickly create rules. For more information, see [Definition and working principles of rules](https://help.aliyun.com/document_detail/128273.html).
//
// - Create custom rules using Function Compute
//
//	Custom rules use Function Compute functions to host your rule code. If Cloud Config\\"s predefined rule templates do not meet your compliance requirements, you can write your own function code to check compliance in complex scenarios. For more information, see [Definition and working principles of custom rules](https://help.aliyun.com/document_detail/127405.html).
//
// ### Usage notes
//
// This topic demonstrates how to create a rule from the \\`required-tags\\` template. The response confirms that the rule was created successfully. Its ID is `cr-5772ba41209e007b****`.
//
// @param tmpReq - CreateConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConfigRuleResponse
func (client *Client) CreateConfigRuleWithOptions(tmpReq *CreateConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateConfigRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateConfigRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InputParameters) {
		request.InputParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InputParameters, dara.String("InputParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceTypesScope) {
		request.ResourceTypesScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceTypesScope, dara.String("ResourceTypesScope"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Conditions) {
		body["Conditions"] = request.Conditions
	}

	if !dara.IsNil(request.ConfigRuleName) {
		body["ConfigRuleName"] = request.ConfigRuleName
	}

	if !dara.IsNil(request.ConfigRuleTriggerTypes) {
		body["ConfigRuleTriggerTypes"] = request.ConfigRuleTriggerTypes
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExcludeRegionIdsScope) {
		body["ExcludeRegionIdsScope"] = request.ExcludeRegionIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceGroupIdsScope) {
		body["ExcludeResourceGroupIdsScope"] = request.ExcludeResourceGroupIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceIdsScope) {
		body["ExcludeResourceIdsScope"] = request.ExcludeResourceIdsScope
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ExcludeTagsScope) {
		bodyFlat["ExcludeTagsScope"] = request.ExcludeTagsScope
	}

	if !dara.IsNil(request.ExtendContent) {
		body["ExtendContent"] = request.ExtendContent
	}

	if !dara.IsNil(request.InputParametersShrink) {
		body["InputParameters"] = request.InputParametersShrink
	}

	if !dara.IsNil(request.MaximumExecutionFrequency) {
		body["MaximumExecutionFrequency"] = request.MaximumExecutionFrequency
	}

	if !dara.IsNil(request.RegionIdsScope) {
		body["RegionIdsScope"] = request.RegionIdsScope
	}

	if !dara.IsNil(request.ResourceGroupIdsScope) {
		body["ResourceGroupIdsScope"] = request.ResourceGroupIdsScope
	}

	if !dara.IsNil(request.ResourceIdsScope) {
		body["ResourceIdsScope"] = request.ResourceIdsScope
	}

	if !dara.IsNil(request.ResourceNameScope) {
		body["ResourceNameScope"] = request.ResourceNameScope
	}

	if !dara.IsNil(request.ResourceTypesScopeShrink) {
		body["ResourceTypesScope"] = request.ResourceTypesScopeShrink
	}

	if !dara.IsNil(request.RiskLevel) {
		body["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.SourceIdentifier) {
		body["SourceIdentifier"] = request.SourceIdentifier
	}

	if !dara.IsNil(request.SourceOwner) {
		body["SourceOwner"] = request.SourceOwner
	}

	if !dara.IsNil(request.TagKeyLogicScope) {
		body["TagKeyLogicScope"] = request.TagKeyLogicScope
	}

	if !dara.IsNil(request.TagKeyScope) {
		body["TagKeyScope"] = request.TagKeyScope
	}

	if !dara.IsNil(request.TagValueScope) {
		body["TagValueScope"] = request.TagValueScope
	}

	if !dara.IsNil(request.TagsScope) {
		bodyFlat["TagsScope"] = request.TagsScope
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConfigRule"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a rule from a template or a custom rule using Function Compute to check resource compliance. After you create a rule, Cloud Config runs an initial evaluation and then automatically triggers subsequent evaluations based on the rule\\"s trigger. You can also run evaluations manually.
//
// Description:
//
// ### Limits
//
// You can create up to 200 rules per account.
//
// ### Background information
//
// You can create rules in Cloud Config in two ways:
//
// - Create rules from templates
//
//	Rule templates are predefined rule functions provided by Cloud Config in Function Compute. You can use rule templates to quickly create rules. For more information, see [Definition and working principles of rules](https://help.aliyun.com/document_detail/128273.html).
//
// - Create custom rules using Function Compute
//
//	Custom rules use Function Compute functions to host your rule code. If Cloud Config\\"s predefined rule templates do not meet your compliance requirements, you can write your own function code to check compliance in complex scenarios. For more information, see [Definition and working principles of custom rules](https://help.aliyun.com/document_detail/127405.html).
//
// ### Usage notes
//
// This topic demonstrates how to create a rule from the \\`required-tags\\` template. The response confirms that the rule was created successfully. Its ID is `cr-5772ba41209e007b****`.
//
// @param request - CreateConfigRuleRequest
//
// @return CreateConfigRuleResponse
func (client *Client) CreateConfigRule(request *CreateConfigRuleRequest) (_result *CreateConfigRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateConfigRuleResponse{}
	_body, _err := client.CreateConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a remediation for an existing rule.
//
// Description:
//
// ### Background information
//
// - Template-based remediation: Uses public templates from Operation Orchestration Service (OOS) to remediate non-compliant resources.
//
//	Each rule supports only one remediation. This remediation type is available only for rules created from specific templates.
//
// - Custom remediation: Uses custom code in Function Compute (FC) to remediate non-compliant resources.
//
//	Each rule supports only one remediation. This remediation type is available for rules created from templates and custom rules.
//
// ### Usage notes
//
// In this example, a remediation is created for the rule `cr-8a973ac2e2be00a2****`. The sample response shows that the remediation is created and has the ID `crr-909ba2d4716700eb****`.
//
// @param request - CreateRemediationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRemediationResponse
func (client *Client) CreateRemediationWithOptions(request *CreateRemediationRequest, runtime *dara.RuntimeOptions) (_result *CreateRemediationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigRuleId) {
		body["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.InvokeType) {
		body["InvokeType"] = request.InvokeType
	}

	if !dara.IsNil(request.Params) {
		body["Params"] = request.Params
	}

	if !dara.IsNil(request.RemediationTemplateId) {
		body["RemediationTemplateId"] = request.RemediationTemplateId
	}

	if !dara.IsNil(request.RemediationType) {
		body["RemediationType"] = request.RemediationType
	}

	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRemediation"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRemediationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a remediation for an existing rule.
//
// Description:
//
// ### Background information
//
// - Template-based remediation: Uses public templates from Operation Orchestration Service (OOS) to remediate non-compliant resources.
//
//	Each rule supports only one remediation. This remediation type is available only for rules created from specific templates.
//
// - Custom remediation: Uses custom code in Function Compute (FC) to remediate non-compliant resources.
//
//	Each rule supports only one remediation. This remediation type is available for rules created from templates and custom rules.
//
// ### Usage notes
//
// In this example, a remediation is created for the rule `cr-8a973ac2e2be00a2****`. The sample response shows that the remediation is created and has the ID `crr-909ba2d4716700eb****`.
//
// @param request - CreateRemediationRequest
//
// @return CreateRemediationResponse
func (client *Client) CreateRemediation(request *CreateRemediationRequest) (_result *CreateRemediationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRemediationResponse{}
	_body, _err := client.CreateRemediationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a compliance report template for the current UID.
//
// @param tmpReq - CreateReportTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReportTemplateResponse
func (client *Client) CreateReportTemplateWithOptions(tmpReq *CreateReportTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateReportTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateReportTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReportScope) {
		request.ReportScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReportScope, dara.String("ReportScope"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ReportFileFormats) {
		body["ReportFileFormats"] = request.ReportFileFormats
	}

	if !dara.IsNil(request.ReportGranularity) {
		body["ReportGranularity"] = request.ReportGranularity
	}

	if !dara.IsNil(request.ReportLanguage) {
		body["ReportLanguage"] = request.ReportLanguage
	}

	if !dara.IsNil(request.ReportScopeShrink) {
		body["ReportScope"] = request.ReportScopeShrink
	}

	if !dara.IsNil(request.ReportTemplateDescription) {
		body["ReportTemplateDescription"] = request.ReportTemplateDescription
	}

	if !dara.IsNil(request.ReportTemplateName) {
		body["ReportTemplateName"] = request.ReportTemplateName
	}

	if !dara.IsNil(request.SubscriptionFrequency) {
		body["SubscriptionFrequency"] = request.SubscriptionFrequency
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateReportTemplate"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateReportTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a compliance report template for the current UID.
//
// @param request - CreateReportTemplateRequest
//
// @return CreateReportTemplateResponse
func (client *Client) CreateReportTemplate(request *CreateReportTemplateRequest) (_result *CreateReportTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateReportTemplateResponse{}
	_body, _err := client.CreateReportTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停用账号组规则
//
// @param request - DeactiveAggregateConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeactiveAggregateConfigRulesResponse
func (client *Client) DeactiveAggregateConfigRulesWithOptions(request *DeactiveAggregateConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *DeactiveAggregateConfigRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.CompliancePackId) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !dara.IsNil(request.ConfigRuleIds) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeactiveAggregateConfigRules"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeactiveAggregateConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停用账号组规则
//
// @param request - DeactiveAggregateConfigRulesRequest
//
// @return DeactiveAggregateConfigRulesResponse
func (client *Client) DeactiveAggregateConfigRules(request *DeactiveAggregateConfigRulesRequest) (_result *DeactiveAggregateConfigRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeactiveAggregateConfigRulesResponse{}
	_body, _err := client.DeactiveAggregateConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停用指定规则
//
// @param request - DeactiveConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeactiveConfigRulesResponse
func (client *Client) DeactiveConfigRulesWithOptions(request *DeactiveConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *DeactiveConfigRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompliancePackId) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !dara.IsNil(request.ConfigRuleIds) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeactiveConfigRules"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeactiveConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停用指定规则
//
// @param request - DeactiveConfigRulesRequest
//
// @return DeactiveConfigRulesResponse
func (client *Client) DeactiveConfigRules(request *DeactiveConfigRulesRequest) (_result *DeactiveConfigRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeactiveConfigRulesResponse{}
	_body, _err := client.DeactiveConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定账号组内的合规包。删除合规包后，该合规包的检查结果和检查报告也自动被删除
//
// @param request - DeleteAggregateCompliancePacksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAggregateCompliancePacksResponse
func (client *Client) DeleteAggregateCompliancePacksWithOptions(request *DeleteAggregateCompliancePacksRequest, runtime *dara.RuntimeOptions) (_result *DeleteAggregateCompliancePacksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CompliancePackIds) {
		body["CompliancePackIds"] = request.CompliancePackIds
	}

	if !dara.IsNil(request.DeleteRule) {
		body["DeleteRule"] = request.DeleteRule
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAggregateCompliancePacks"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAggregateCompliancePacksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定账号组内的合规包。删除合规包后，该合规包的检查结果和检查报告也自动被删除
//
// @param request - DeleteAggregateCompliancePacksRequest
//
// @return DeleteAggregateCompliancePacksResponse
func (client *Client) DeleteAggregateCompliancePacks(request *DeleteAggregateCompliancePacksRequest) (_result *DeleteAggregateCompliancePacksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAggregateCompliancePacksResponse{}
	_body, _err := client.DeleteAggregateCompliancePacksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 账号组删除投递渠道
//
// @param request - DeleteAggregateConfigDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAggregateConfigDeliveryChannelResponse
func (client *Client) DeleteAggregateConfigDeliveryChannelWithOptions(request *DeleteAggregateConfigDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteAggregateConfigDeliveryChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.DeliveryChannelId) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAggregateConfigDeliveryChannel"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAggregateConfigDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 账号组删除投递渠道
//
// @param request - DeleteAggregateConfigDeliveryChannelRequest
//
// @return DeleteAggregateConfigDeliveryChannelResponse
func (client *Client) DeleteAggregateConfigDeliveryChannel(request *DeleteAggregateConfigDeliveryChannelRequest) (_result *DeleteAggregateConfigDeliveryChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAggregateConfigDeliveryChannelResponse{}
	_body, _err := client.DeleteAggregateConfigDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定账号组内的规则。删除规则后，其配置信息不再保留
//
// @param request - DeleteAggregateConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAggregateConfigRulesResponse
func (client *Client) DeleteAggregateConfigRulesWithOptions(request *DeleteAggregateConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteAggregateConfigRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ConfigRuleIds) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAggregateConfigRules"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAggregateConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定账号组内的规则。删除规则后，其配置信息不再保留
//
// @param request - DeleteAggregateConfigRulesRequest
//
// @return DeleteAggregateConfigRulesResponse
func (client *Client) DeleteAggregateConfigRules(request *DeleteAggregateConfigRulesRequest) (_result *DeleteAggregateConfigRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAggregateConfigRulesResponse{}
	_body, _err := client.DeleteAggregateConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the remediation settings for a rule in a specified account group.
//
// Description:
//
// This topic provides an example of how to delete the remediation setting `crr-909ba2d4716700eb****` for a rule in the account group `ca-6b4a626622af0012****`. The response shows that the remediation setting is deleted.
//
// @param request - DeleteAggregateRemediationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAggregateRemediationsResponse
func (client *Client) DeleteAggregateRemediationsWithOptions(request *DeleteAggregateRemediationsRequest, runtime *dara.RuntimeOptions) (_result *DeleteAggregateRemediationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.RemediationIds) {
		body["RemediationIds"] = request.RemediationIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAggregateRemediations"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAggregateRemediationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the remediation settings for a rule in a specified account group.
//
// Description:
//
// This topic provides an example of how to delete the remediation setting `crr-909ba2d4716700eb****` for a rule in the account group `ca-6b4a626622af0012****`. The response shows that the remediation setting is deleted.
//
// @param request - DeleteAggregateRemediationsRequest
//
// @return DeleteAggregateRemediationsResponse
func (client *Client) DeleteAggregateRemediations(request *DeleteAggregateRemediationsRequest) (_result *DeleteAggregateRemediationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAggregateRemediationsResponse{}
	_body, _err := client.DeleteAggregateRemediationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除账号组
//
// @param request - DeleteAggregatorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAggregatorsResponse
func (client *Client) DeleteAggregatorsWithOptions(request *DeleteAggregatorsRequest, runtime *dara.RuntimeOptions) (_result *DeleteAggregatorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorIds) {
		body["AggregatorIds"] = request.AggregatorIds
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAggregators"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAggregatorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除账号组
//
// @param request - DeleteAggregatorsRequest
//
// @return DeleteAggregatorsResponse
func (client *Client) DeleteAggregators(request *DeleteAggregatorsRequest) (_result *DeleteAggregatorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAggregatorsResponse{}
	_body, _err := client.DeleteAggregatorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除合规包
//
// @param request - DeleteCompliancePacksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCompliancePacksResponse
func (client *Client) DeleteCompliancePacksWithOptions(request *DeleteCompliancePacksRequest, runtime *dara.RuntimeOptions) (_result *DeleteCompliancePacksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CompliancePackIds) {
		body["CompliancePackIds"] = request.CompliancePackIds
	}

	if !dara.IsNil(request.DeleteRule) {
		body["DeleteRule"] = request.DeleteRule
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCompliancePacks"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCompliancePacksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除合规包
//
// @param request - DeleteCompliancePacksRequest
//
// @return DeleteCompliancePacksResponse
func (client *Client) DeleteCompliancePacks(request *DeleteCompliancePacksRequest) (_result *DeleteCompliancePacksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCompliancePacksResponse{}
	_body, _err := client.DeleteCompliancePacksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 当前账号删除投递渠道
//
// @param request - DeleteConfigDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConfigDeliveryChannelResponse
func (client *Client) DeleteConfigDeliveryChannelWithOptions(request *DeleteConfigDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteConfigDeliveryChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryChannelId) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConfigDeliveryChannel"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConfigDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 当前账号删除投递渠道
//
// @param request - DeleteConfigDeliveryChannelRequest
//
// @return DeleteConfigDeliveryChannelResponse
func (client *Client) DeleteConfigDeliveryChannel(request *DeleteConfigDeliveryChannelRequest) (_result *DeleteConfigDeliveryChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteConfigDeliveryChannelResponse{}
	_body, _err := client.DeleteConfigDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除规则
//
// @param request - DeleteConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConfigRulesResponse
func (client *Client) DeleteConfigRulesWithOptions(request *DeleteConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteConfigRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigRuleIds) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConfigRules"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除规则
//
// @param request - DeleteConfigRulesRequest
//
// @return DeleteConfigRulesResponse
func (client *Client) DeleteConfigRules(request *DeleteConfigRulesRequest) (_result *DeleteConfigRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteConfigRulesResponse{}
	_body, _err := client.DeleteConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes specified remediation settings.
//
// Description:
//
// This topic provides an example of how to delete a remediation setting for a rule. In this example, the remediation setting with the ID `crr-909ba2d4716700eb****` is deleted. The response indicates that the operation is successful.
//
// @param request - DeleteRemediationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRemediationsResponse
func (client *Client) DeleteRemediationsWithOptions(request *DeleteRemediationsRequest, runtime *dara.RuntimeOptions) (_result *DeleteRemediationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RemediationIds) {
		body["RemediationIds"] = request.RemediationIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRemediations"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRemediationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes specified remediation settings.
//
// Description:
//
// This topic provides an example of how to delete a remediation setting for a rule. In this example, the remediation setting with the ID `crr-909ba2d4716700eb****` is deleted. The response indicates that the operation is successful.
//
// @param request - DeleteRemediationsRequest
//
// @return DeleteRemediationsResponse
func (client *Client) DeleteRemediations(request *DeleteRemediationsRequest) (_result *DeleteRemediationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRemediationsResponse{}
	_body, _err := client.DeleteRemediationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a compliance report template.
//
// @param request - DeleteReportTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteReportTemplateResponse
func (client *Client) DeleteReportTemplateWithOptions(request *DeleteReportTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteReportTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ReportTemplateId) {
		body["ReportTemplateId"] = request.ReportTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteReportTemplate"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteReportTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a compliance report template.
//
// @param request - DeleteReportTemplateRequest
//
// @return DeleteReportTemplateResponse
func (client *Client) DeleteReportTemplate(request *DeleteReportTemplateRequest) (_result *DeleteReportTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteReportTemplateResponse{}
	_body, _err := client.DeleteReportTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation retrieves the details of multiple resources in a batch.
//
// @param request - DescribeDiscoveredResourceBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiscoveredResourceBatchResponse
func (client *Client) DescribeDiscoveredResourceBatchWithOptions(request *DescribeDiscoveredResourceBatchRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiscoveredResourceBatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Regions) {
		query["Regions"] = request.Regions
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiscoveredResourceBatch"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiscoveredResourceBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation retrieves the details of multiple resources in a batch.
//
// @param request - DescribeDiscoveredResourceBatchRequest
//
// @return DescribeDiscoveredResourceBatchResponse
func (client *Client) DescribeDiscoveredResourceBatch(request *DescribeDiscoveredResourceBatchRequest) (_result *DescribeDiscoveredResourceBatchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDiscoveredResourceBatchResponse{}
	_body, _err := client.DescribeDiscoveredResourceBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the authorization status of an integrated cloud service.
//
// @param request - DescribeIntegratedServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIntegratedServiceStatusResponse
func (client *Client) DescribeIntegratedServiceStatusWithOptions(request *DescribeIntegratedServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeIntegratedServiceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ServiceCode) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIntegratedServiceStatus"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIntegratedServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authorization status of an integrated cloud service.
//
// @param request - DescribeIntegratedServiceStatusRequest
//
// @return DescribeIntegratedServiceStatusResponse
func (client *Client) DescribeIntegratedServiceStatus(request *DescribeIntegratedServiceStatusRequest) (_result *DescribeIntegratedServiceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIntegratedServiceStatusResponse{}
	_body, _err := client.DescribeIntegratedServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取修正详情
//
// @param request - DescribeRemediationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRemediationResponse
func (client *Client) DescribeRemediationWithOptions(request *DescribeRemediationRequest, runtime *dara.RuntimeOptions) (_result *DescribeRemediationResponse, _err error) {
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
		Action:      dara.String("DescribeRemediation"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRemediationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取修正详情
//
// @param request - DescribeRemediationRequest
//
// @return DescribeRemediationResponse
func (client *Client) DescribeRemediation(request *DescribeRemediationRequest) (_result *DescribeRemediationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRemediationResponse{}
	_body, _err := client.DescribeRemediationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将指定账号组规则从指定账号组合规包中移出
//
// @param request - DetachAggregateConfigRuleToCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachAggregateConfigRuleToCompliancePackResponse
func (client *Client) DetachAggregateConfigRuleToCompliancePackWithOptions(request *DetachAggregateConfigRuleToCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *DetachAggregateConfigRuleToCompliancePackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.CompliancePackId) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !dara.IsNil(request.ConfigRuleIds) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachAggregateConfigRuleToCompliancePack"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachAggregateConfigRuleToCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将指定账号组规则从指定账号组合规包中移出
//
// @param request - DetachAggregateConfigRuleToCompliancePackRequest
//
// @return DetachAggregateConfigRuleToCompliancePackResponse
func (client *Client) DetachAggregateConfigRuleToCompliancePack(request *DetachAggregateConfigRuleToCompliancePackRequest) (_result *DetachAggregateConfigRuleToCompliancePackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachAggregateConfigRuleToCompliancePackResponse{}
	_body, _err := client.DetachAggregateConfigRuleToCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation detaches one or more rules from a compliance package.
//
// Description:
//
// ### Prerequisites
//
// Make sure that the destination compliance package contains rules.
//
// ### Usage notes
//
// This topic provides an example of detaching the rule `cr-6cc4626622af00e7****` from the compliance package `cp-5bb1626622af00bd****`.
//
// @param request - DetachConfigRuleToCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachConfigRuleToCompliancePackResponse
func (client *Client) DetachConfigRuleToCompliancePackWithOptions(request *DetachConfigRuleToCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *DetachConfigRuleToCompliancePackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompliancePackId) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !dara.IsNil(request.ConfigRuleIds) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachConfigRuleToCompliancePack"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachConfigRuleToCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation detaches one or more rules from a compliance package.
//
// Description:
//
// ### Prerequisites
//
// Make sure that the destination compliance package contains rules.
//
// ### Usage notes
//
// This topic provides an example of detaching the rule `cr-6cc4626622af00e7****` from the compliance package `cp-5bb1626622af00bd****`.
//
// @param request - DetachConfigRuleToCompliancePackRequest
//
// @return DetachConfigRuleToCompliancePackResponse
func (client *Client) DetachConfigRuleToCompliancePack(request *DetachConfigRuleToCompliancePackRequest) (_result *DetachConfigRuleToCompliancePackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachConfigRuleToCompliancePackResponse{}
	_body, _err := client.DetachConfigRuleToCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This API performs a dry run on rules for proactive compliance pre-checks.
//
// @param request - DryRunConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DryRunConfigRuleResponse
func (client *Client) DryRunConfigRuleWithOptions(request *DryRunConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *DryRunConfigRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigurationItem) {
		body["ConfigurationItem"] = request.ConfigurationItem
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DryRunConfigRule"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DryRunConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API performs a dry run on rules for proactive compliance pre-checks.
//
// @param request - DryRunConfigRuleRequest
//
// @return DryRunConfigRuleResponse
func (client *Client) DryRunConfigRule(request *DryRunConfigRuleRequest) (_result *DryRunConfigRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DryRunConfigRuleResponse{}
	_body, _err := client.DryRunConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes evaluation rules to perform compliance pre-checks on resources.
//
// Description:
//
// For more information about the concepts, operating principles, and integration process of evaluation rules, see [Definition and operating principles of evaluation rules](https://help.aliyun.com/document_detail/470802.html).
//
// After you create an evaluation rule, a managed rule with the same settings is created. This managed rule can continuously check the compliance of resources after they are created.
//
// @param tmpReq - EvaluatePreConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EvaluatePreConfigRulesResponse
func (client *Client) EvaluatePreConfigRulesWithOptions(tmpReq *EvaluatePreConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *EvaluatePreConfigRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &EvaluatePreConfigRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceEvaluateItems) {
		request.ResourceEvaluateItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceEvaluateItems, dara.String("ResourceEvaluateItems"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EnableManagedRules) {
		body["EnableManagedRules"] = request.EnableManagedRules
	}

	if !dara.IsNil(request.ResourceEvaluateItemsShrink) {
		body["ResourceEvaluateItems"] = request.ResourceEvaluateItemsShrink
	}

	if !dara.IsNil(request.ResourceTypeFormat) {
		body["ResourceTypeFormat"] = request.ResourceTypeFormat
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EvaluatePreConfigRules"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EvaluatePreConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes evaluation rules to perform compliance pre-checks on resources.
//
// Description:
//
// For more information about the concepts, operating principles, and integration process of evaluation rules, see [Definition and operating principles of evaluation rules](https://help.aliyun.com/document_detail/470802.html).
//
// After you create an evaluation rule, a managed rule with the same settings is created. This managed rule can continuously check the compliance of resources after they are created.
//
// @param request - EvaluatePreConfigRulesRequest
//
// @return EvaluatePreConfigRulesResponse
func (client *Client) EvaluatePreConfigRules(request *EvaluatePreConfigRulesRequest) (_result *EvaluatePreConfigRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EvaluatePreConfigRulesResponse{}
	_body, _err := client.EvaluatePreConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates an assessment report for a specified compliance package in a specified account group.
//
// Description:
//
// > This operation only generates the latest assessment report. You need to call the GetAggregateCompliancePackReport operation to download the assessment report. For more information, see [GetAggregateCompliancePackReport](https://help.aliyun.com/document_detail/262699.html).
//
// This topic provides an example that shows how to generate an assessment report for the compliance package `cp-fdc8626622af00f9****` in the account group `ca-f632626622af0079****`.
//
// @param request - GenerateAggregateCompliancePackReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateAggregateCompliancePackReportResponse
func (client *Client) GenerateAggregateCompliancePackReportWithOptions(request *GenerateAggregateCompliancePackReportRequest, runtime *dara.RuntimeOptions) (_result *GenerateAggregateCompliancePackReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CompliancePackId) {
		body["CompliancePackId"] = request.CompliancePackId
	}

	if !dara.IsNil(request.MultiFiles) {
		body["MultiFiles"] = request.MultiFiles
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateAggregateCompliancePackReport"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateAggregateCompliancePackReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates an assessment report for a specified compliance package in a specified account group.
//
// Description:
//
// > This operation only generates the latest assessment report. You need to call the GetAggregateCompliancePackReport operation to download the assessment report. For more information, see [GetAggregateCompliancePackReport](https://help.aliyun.com/document_detail/262699.html).
//
// This topic provides an example that shows how to generate an assessment report for the compliance package `cp-fdc8626622af00f9****` in the account group `ca-f632626622af0079****`.
//
// @param request - GenerateAggregateCompliancePackReportRequest
//
// @return GenerateAggregateCompliancePackReportResponse
func (client *Client) GenerateAggregateCompliancePackReport(request *GenerateAggregateCompliancePackReportRequest) (_result *GenerateAggregateCompliancePackReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateAggregateCompliancePackReportResponse{}
	_body, _err := client.GenerateAggregateCompliancePackReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成指定账号组的合规报告
//
// @param request - GenerateAggregateConfigRulesReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateAggregateConfigRulesReportResponse
func (client *Client) GenerateAggregateConfigRulesReportWithOptions(request *GenerateAggregateConfigRulesReportRequest, runtime *dara.RuntimeOptions) (_result *GenerateAggregateConfigRulesReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigRuleIds) {
		body["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateAggregateConfigRulesReport"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateAggregateConfigRulesReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成指定账号组的合规报告
//
// @param request - GenerateAggregateConfigRulesReportRequest
//
// @return GenerateAggregateConfigRulesReportResponse
func (client *Client) GenerateAggregateConfigRulesReport(request *GenerateAggregateConfigRulesReportRequest) (_result *GenerateAggregateConfigRulesReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateAggregateConfigRulesReportResponse{}
	_body, _err := client.GenerateAggregateConfigRulesReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 账号组资源清单生成
//
// @param request - GenerateAggregateResourceInventoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateAggregateResourceInventoryResponse
func (client *Client) GenerateAggregateResourceInventoryWithOptions(request *GenerateAggregateResourceInventoryRequest, runtime *dara.RuntimeOptions) (_result *GenerateAggregateResourceInventoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountIds) {
		query["AccountIds"] = request.AccountIds
	}

	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.Regions) {
		query["Regions"] = request.Regions
	}

	if !dara.IsNil(request.ResourceDeleted) {
		query["ResourceDeleted"] = request.ResourceDeleted
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateAggregateResourceInventory"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateAggregateResourceInventoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 账号组资源清单生成
//
// @param request - GenerateAggregateResourceInventoryRequest
//
// @return GenerateAggregateResourceInventoryResponse
func (client *Client) GenerateAggregateResourceInventory(request *GenerateAggregateResourceInventoryRequest) (_result *GenerateAggregateResourceInventoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateAggregateResourceInventoryResponse{}
	_body, _err := client.GenerateAggregateResourceInventoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成合规包的合规包报告
//
// @param request - GenerateCompliancePackReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCompliancePackReportResponse
func (client *Client) GenerateCompliancePackReportWithOptions(request *GenerateCompliancePackReportRequest, runtime *dara.RuntimeOptions) (_result *GenerateCompliancePackReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CompliancePackId) {
		body["CompliancePackId"] = request.CompliancePackId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateCompliancePackReport"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateCompliancePackReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成合规包的合规包报告
//
// @param request - GenerateCompliancePackReportRequest
//
// @return GenerateCompliancePackReportResponse
func (client *Client) GenerateCompliancePackReport(request *GenerateCompliancePackReportRequest) (_result *GenerateCompliancePackReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateCompliancePackReportResponse{}
	_body, _err := client.GenerateCompliancePackReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成规则的合规评估报告
//
// @param request - GenerateConfigRulesReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateConfigRulesReportResponse
func (client *Client) GenerateConfigRulesReportWithOptions(request *GenerateConfigRulesReportRequest, runtime *dara.RuntimeOptions) (_result *GenerateConfigRulesReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigRuleIds) {
		body["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateConfigRulesReport"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateConfigRulesReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成规则的合规评估报告
//
// @param request - GenerateConfigRulesReportRequest
//
// @return GenerateConfigRulesReportResponse
func (client *Client) GenerateConfigRulesReport(request *GenerateConfigRulesReportRequest) (_result *GenerateConfigRulesReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateConfigRulesReportResponse{}
	_body, _err := client.GenerateConfigRulesReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generate a report ID from a report template.
//
// @param request - GenerateReportFromTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateReportFromTemplateResponse
func (client *Client) GenerateReportFromTemplateWithOptions(request *GenerateReportFromTemplateRequest, runtime *dara.RuntimeOptions) (_result *GenerateReportFromTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ReportTemplateId) {
		body["ReportTemplateId"] = request.ReportTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateReportFromTemplate"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateReportFromTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generate a report ID from a report template.
//
// @param request - GenerateReportFromTemplateRequest
//
// @return GenerateReportFromTemplateResponse
func (client *Client) GenerateReportFromTemplate(request *GenerateReportFromTemplateRequest) (_result *GenerateReportFromTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateReportFromTemplateResponse{}
	_body, _err := client.GenerateReportFromTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a downloadable inventory of global resources.
//
// Description:
//
// This topic provides an example of how to generate a downloadable inventory of global resources for the current account.
//
// @param request - GenerateResourceInventoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateResourceInventoryResponse
func (client *Client) GenerateResourceInventoryWithOptions(request *GenerateResourceInventoryRequest, runtime *dara.RuntimeOptions) (_result *GenerateResourceInventoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Regions) {
		query["Regions"] = request.Regions
	}

	if !dara.IsNil(request.ResourceDeleted) {
		query["ResourceDeleted"] = request.ResourceDeleted
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateResourceInventory"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateResourceInventoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a downloadable inventory of global resources.
//
// Description:
//
// This topic provides an example of how to generate a downloadable inventory of global resources for the current account.
//
// @param request - GenerateResourceInventoryRequest
//
// @return GenerateResourceInventoryResponse
func (client *Client) GenerateResourceInventory(request *GenerateResourceInventoryRequest) (_result *GenerateResourceInventoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateResourceInventoryResponse{}
	_body, _err := client.GenerateResourceInventoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 高级搜索 获取搜索结果下载文件地址
//
// @param request - GetAdvancedSearchFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdvancedSearchFileResponse
func (client *Client) GetAdvancedSearchFileWithOptions(runtime *dara.RuntimeOptions) (_result *GetAdvancedSearchFileResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetAdvancedSearchFile"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAdvancedSearchFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 高级搜索 获取搜索结果下载文件地址
//
// @return GetAdvancedSearchFileResponse
func (client *Client) GetAdvancedSearchFile() (_result *GetAdvancedSearchFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAdvancedSearchFileResponse{}
	_body, _err := client.GetAdvancedSearchFileWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定账号组内指定合规包中成员账号的合规结果
//
// @param request - GetAggregateAccountComplianceByPackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateAccountComplianceByPackResponse
func (client *Client) GetAggregateAccountComplianceByPackWithOptions(request *GetAggregateAccountComplianceByPackRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateAccountComplianceByPackResponse, _err error) {
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
		Action:      dara.String("GetAggregateAccountComplianceByPack"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateAccountComplianceByPackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定账号组内指定合规包中成员账号的合规结果
//
// @param request - GetAggregateAccountComplianceByPackRequest
//
// @return GetAggregateAccountComplianceByPackResponse
func (client *Client) GetAggregateAccountComplianceByPack(request *GetAggregateAccountComplianceByPackRequest) (_result *GetAggregateAccountComplianceByPackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateAccountComplianceByPackResponse{}
	_body, _err := client.GetAggregateAccountComplianceByPackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 高级搜获取搜索结果下载文件地址
//
// @param request - GetAggregateAdvancedSearchFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateAdvancedSearchFileResponse
func (client *Client) GetAggregateAdvancedSearchFileWithOptions(request *GetAggregateAdvancedSearchFileRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateAdvancedSearchFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAggregateAdvancedSearchFile"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateAdvancedSearchFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 高级搜获取搜索结果下载文件地址
//
// @param request - GetAggregateAdvancedSearchFileRequest
//
// @return GetAggregateAdvancedSearchFileResponse
func (client *Client) GetAggregateAdvancedSearchFile(request *GetAggregateAdvancedSearchFileRequest) (_result *GetAggregateAdvancedSearchFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateAdvancedSearchFileResponse{}
	_body, _err := client.GetAggregateAdvancedSearchFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a compliance pack in a specified account group.
//
// Description:
//
// This topic provides an example of how to retrieve the details of the compliance pack `cp-fdc8626622af00f9****` in the account group `ca-f632626622af0079****`.
//
// @param tmpReq - GetAggregateCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateCompliancePackResponse
func (client *Client) GetAggregateCompliancePackWithOptions(tmpReq *GetAggregateCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateCompliancePackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAggregateCompliancePackShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAggregateCompliancePack"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a compliance pack in a specified account group.
//
// Description:
//
// This topic provides an example of how to retrieve the details of the compliance pack `cp-fdc8626622af00f9****` in the account group `ca-f632626622af0079****`.
//
// @param request - GetAggregateCompliancePackRequest
//
// @return GetAggregateCompliancePackResponse
func (client *Client) GetAggregateCompliancePack(request *GetAggregateCompliancePackRequest) (_result *GetAggregateCompliancePackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateCompliancePackResponse{}
	_body, _err := client.GetAggregateCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定账号组内指定合规包的评估报告
//
// @param request - GetAggregateCompliancePackReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateCompliancePackReportResponse
func (client *Client) GetAggregateCompliancePackReportWithOptions(request *GetAggregateCompliancePackReportRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateCompliancePackReportResponse, _err error) {
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
		Action:      dara.String("GetAggregateCompliancePackReport"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateCompliancePackReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定账号组内指定合规包的评估报告
//
// @param request - GetAggregateCompliancePackReportRequest
//
// @return GetAggregateCompliancePackReportResponse
func (client *Client) GetAggregateCompliancePackReport(request *GetAggregateCompliancePackReportRequest) (_result *GetAggregateCompliancePackReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateCompliancePackReportResponse{}
	_body, _err := client.GetAggregateCompliancePackReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the compliance summary for a specified account group.
//
// Description:
//
// This topic provides an example of how to query the compliance summary by resource and by rule for the account group ca-a91d626622af0035\\*\\*\\*\\*.
//
// @param request - GetAggregateComplianceSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateComplianceSummaryResponse
func (client *Client) GetAggregateComplianceSummaryWithOptions(request *GetAggregateComplianceSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateComplianceSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAggregateComplianceSummary"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateComplianceSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance summary for a specified account group.
//
// Description:
//
// This topic provides an example of how to query the compliance summary by resource and by rule for the account group ca-a91d626622af0035\\*\\*\\*\\*.
//
// @param request - GetAggregateComplianceSummaryRequest
//
// @return GetAggregateComplianceSummaryResponse
func (client *Client) GetAggregateComplianceSummary(request *GetAggregateComplianceSummaryRequest) (_result *GetAggregateComplianceSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateComplianceSummaryResponse{}
	_body, _err := client.GetAggregateComplianceSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 账号组查询单个投递渠道
//
// @param request - GetAggregateConfigDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateConfigDeliveryChannelResponse
func (client *Client) GetAggregateConfigDeliveryChannelWithOptions(request *GetAggregateConfigDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateConfigDeliveryChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.DeliveryChannelId) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAggregateConfigDeliveryChannel"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateConfigDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 账号组查询单个投递渠道
//
// @param request - GetAggregateConfigDeliveryChannelRequest
//
// @return GetAggregateConfigDeliveryChannelResponse
func (client *Client) GetAggregateConfigDeliveryChannel(request *GetAggregateConfigDeliveryChannelRequest) (_result *GetAggregateConfigDeliveryChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateConfigDeliveryChannelResponse{}
	_body, _err := client.GetAggregateConfigDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specific rule in a specified account group.
//
// Description:
//
// This topic provides an example of how to query the details of the rule `cr-7f7d626622af0041****` in the account group `ca-7f00626622af0041****`.
//
// @param tmpReq - GetAggregateConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateConfigRuleResponse
func (client *Client) GetAggregateConfigRuleWithOptions(tmpReq *GetAggregateConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateConfigRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAggregateConfigRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ConfigRuleId) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAggregateConfigRule"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specific rule in a specified account group.
//
// Description:
//
// This topic provides an example of how to query the details of the rule `cr-7f7d626622af0041****` in the account group `ca-7f00626622af0041****`.
//
// @param request - GetAggregateConfigRuleRequest
//
// @return GetAggregateConfigRuleResponse
func (client *Client) GetAggregateConfigRule(request *GetAggregateConfigRuleRequest) (_result *GetAggregateConfigRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateConfigRuleResponse{}
	_body, _err := client.GetAggregateConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the compliance results of rules in a specified compliance pack within a specified account group.
//
// Description:
//
// This topic provides an example of how to query the compliance results of rules in the compliance pack `cp-541e626622af0087****` for the account group `ca-04b3fd170e340007****`. The response returns a total of `1` rule and `0` non-compliant rules.
//
// @param request - GetAggregateConfigRuleComplianceByPackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateConfigRuleComplianceByPackResponse
func (client *Client) GetAggregateConfigRuleComplianceByPackWithOptions(request *GetAggregateConfigRuleComplianceByPackRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateConfigRuleComplianceByPackResponse, _err error) {
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
		Action:      dara.String("GetAggregateConfigRuleComplianceByPack"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateConfigRuleComplianceByPackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance results of rules in a specified compliance pack within a specified account group.
//
// Description:
//
// This topic provides an example of how to query the compliance results of rules in the compliance pack `cp-541e626622af0087****` for the account group `ca-04b3fd170e340007****`. The response returns a total of `1` rule and `0` non-compliant rules.
//
// @param request - GetAggregateConfigRuleComplianceByPackRequest
//
// @return GetAggregateConfigRuleComplianceByPackResponse
func (client *Client) GetAggregateConfigRuleComplianceByPack(request *GetAggregateConfigRuleComplianceByPackRequest) (_result *GetAggregateConfigRuleComplianceByPackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateConfigRuleComplianceByPackResponse{}
	_body, _err := client.GetAggregateConfigRuleComplianceByPackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定风险等级的账号组规则合规统计
//
// @param request - GetAggregateConfigRuleSummaryByRiskLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateConfigRuleSummaryByRiskLevelResponse
func (client *Client) GetAggregateConfigRuleSummaryByRiskLevelWithOptions(request *GetAggregateConfigRuleSummaryByRiskLevelRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateConfigRuleSummaryByRiskLevelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAggregateConfigRuleSummaryByRiskLevel"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateConfigRuleSummaryByRiskLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定风险等级的账号组规则合规统计
//
// @param request - GetAggregateConfigRuleSummaryByRiskLevelRequest
//
// @return GetAggregateConfigRuleSummaryByRiskLevelResponse
func (client *Client) GetAggregateConfigRuleSummaryByRiskLevel(request *GetAggregateConfigRuleSummaryByRiskLevelRequest) (_result *GetAggregateConfigRuleSummaryByRiskLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateConfigRuleSummaryByRiskLevelResponse{}
	_body, _err := client.GetAggregateConfigRuleSummaryByRiskLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 下载Excel格式的规则合规评估报告到本地，便于您云下分配任务并跟进不合规资源配置的修改
//
// @param request - GetAggregateConfigRulesReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateConfigRulesReportResponse
func (client *Client) GetAggregateConfigRulesReportWithOptions(request *GetAggregateConfigRulesReportRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateConfigRulesReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAggregateConfigRulesReport"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateConfigRulesReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下载Excel格式的规则合规评估报告到本地，便于您云下分配任务并跟进不合规资源配置的修改
//
// @param request - GetAggregateConfigRulesReportRequest
//
// @return GetAggregateConfigRulesReportResponse
func (client *Client) GetAggregateConfigRulesReport(request *GetAggregateConfigRulesReportRequest) (_result *GetAggregateConfigRulesReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateConfigRulesReportResponse{}
	_body, _err := client.GetAggregateConfigRulesReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specific resource in an account group.
//
// Description:
//
// This topic provides an example on how to query the details of an ECS instance `i-bp12g4xbl4i0brkn****` in the Hangzhou region within the account group `ca-5885626622af0008****`.
//
// @param request - GetAggregateDiscoveredResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateDiscoveredResourceResponse
func (client *Client) GetAggregateDiscoveredResourceWithOptions(request *GetAggregateDiscoveredResourceRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateDiscoveredResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ComplianceOption) {
		query["ComplianceOption"] = request.ComplianceOption
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceAccountId) {
		query["ResourceAccountId"] = request.ResourceAccountId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
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
		Action:      dara.String("GetAggregateDiscoveredResource"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateDiscoveredResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specific resource in an account group.
//
// Description:
//
// This topic provides an example on how to query the details of an ECS instance `i-bp12g4xbl4i0brkn****` in the Hangzhou region within the account group `ca-5885626622af0008****`.
//
// @param request - GetAggregateDiscoveredResourceRequest
//
// @return GetAggregateDiscoveredResourceResponse
func (client *Client) GetAggregateDiscoveredResource(request *GetAggregateDiscoveredResourceRequest) (_result *GetAggregateDiscoveredResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateDiscoveredResourceResponse{}
	_body, _err := client.GetAggregateDiscoveredResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询多账号规则合规情况
//
// @param request - GetAggregateResourceComplianceByConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceComplianceByConfigRuleResponse
func (client *Client) GetAggregateResourceComplianceByConfigRuleWithOptions(request *GetAggregateResourceComplianceByConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceComplianceByConfigRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ComplianceType) {
		query["ComplianceType"] = request.ComplianceType
	}

	if !dara.IsNil(request.ConfigRuleId) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.ResourceAccountId) {
		query["ResourceAccountId"] = request.ResourceAccountId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAggregateResourceComplianceByConfigRule"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateResourceComplianceByConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询多账号规则合规情况
//
// @param request - GetAggregateResourceComplianceByConfigRuleRequest
//
// @return GetAggregateResourceComplianceByConfigRuleResponse
func (client *Client) GetAggregateResourceComplianceByConfigRule(request *GetAggregateResourceComplianceByConfigRuleRequest) (_result *GetAggregateResourceComplianceByConfigRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateResourceComplianceByConfigRuleResponse{}
	_body, _err := client.GetAggregateResourceComplianceByConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the compliance statistics for resources in a specified compliance package within a specified account group.
//
// Description:
//
// This topic provides an example of how to query the compliance results for resources in the compliance package `cp-fdc8626622af00f9****` within the account group `ca-f632626622af0079****`. The response shows that of a total of `10` resources, `7` are non-compliant.
//
// @param request - GetAggregateResourceComplianceByPackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceComplianceByPackResponse
func (client *Client) GetAggregateResourceComplianceByPackWithOptions(request *GetAggregateResourceComplianceByPackRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceComplianceByPackResponse, _err error) {
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
		Action:      dara.String("GetAggregateResourceComplianceByPack"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateResourceComplianceByPackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance statistics for resources in a specified compliance package within a specified account group.
//
// Description:
//
// This topic provides an example of how to query the compliance results for resources in the compliance package `cp-fdc8626622af00f9****` within the account group `ca-f632626622af0079****`. The response shows that of a total of `10` resources, `7` are non-compliant.
//
// @param request - GetAggregateResourceComplianceByPackRequest
//
// @return GetAggregateResourceComplianceByPackResponse
func (client *Client) GetAggregateResourceComplianceByPack(request *GetAggregateResourceComplianceByPackRequest) (_result *GetAggregateResourceComplianceByPackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateResourceComplianceByPackResponse{}
	_body, _err := client.GetAggregateResourceComplianceByPackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定账号组内规则对资源的评估结果，评估结果按资源所属地域进行分组展示
//
// @param request - GetAggregateResourceComplianceGroupByRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceComplianceGroupByRegionResponse
func (client *Client) GetAggregateResourceComplianceGroupByRegionWithOptions(request *GetAggregateResourceComplianceGroupByRegionRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceComplianceGroupByRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ConfigRuleIds) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAggregateResourceComplianceGroupByRegion"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateResourceComplianceGroupByRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定账号组内规则对资源的评估结果，评估结果按资源所属地域进行分组展示
//
// @param request - GetAggregateResourceComplianceGroupByRegionRequest
//
// @return GetAggregateResourceComplianceGroupByRegionResponse
func (client *Client) GetAggregateResourceComplianceGroupByRegion(request *GetAggregateResourceComplianceGroupByRegionRequest) (_result *GetAggregateResourceComplianceGroupByRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateResourceComplianceGroupByRegionResponse{}
	_body, _err := client.GetAggregateResourceComplianceGroupByRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定账号组内规则对资源的评估结果，评估结果按资源类型进行分组展示
//
// @param request - GetAggregateResourceComplianceGroupByResourceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceComplianceGroupByResourceTypeResponse
func (client *Client) GetAggregateResourceComplianceGroupByResourceTypeWithOptions(request *GetAggregateResourceComplianceGroupByResourceTypeRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceComplianceGroupByResourceTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ConfigRuleIds) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAggregateResourceComplianceGroupByResourceType"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateResourceComplianceGroupByResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定账号组内规则对资源的评估结果，评估结果按资源类型进行分组展示
//
// @param request - GetAggregateResourceComplianceGroupByResourceTypeRequest
//
// @return GetAggregateResourceComplianceGroupByResourceTypeResponse
func (client *Client) GetAggregateResourceComplianceGroupByResourceType(request *GetAggregateResourceComplianceGroupByResourceTypeRequest) (_result *GetAggregateResourceComplianceGroupByResourceTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateResourceComplianceGroupByResourceTypeResponse{}
	_body, _err := client.GetAggregateResourceComplianceGroupByResourceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the compliance timeline of a specific resource in an account group. A compliance timeline is a set of compliance evaluation records for a resource. Each record includes the time and content of an evaluation.
//
// Description:
//
// In Cloud Config, each resource has its own compliance timeline, which is composed of compliance evaluation records. A record is generated each time a rule is triggered to evaluate the resource. Rules can be triggered by configuration changes, periodic executions, or manual executions.
//
// This topic provides an example of how to query the compliance timeline for the resource `new-bucket` (an OSS bucket). The resource is in the `cn-hangzhou` region and belongs to the member account `100931896542****` within the account group `ca-5885626622af0008****`. The response shows that the resource\\"s compliance timeline includes records with the timestamps `1625200295276` (UTC+8: 2021-07-02 12:31:35) and `1625200228510` (UTC+8: 2021-07-02 12:30:28).
//
// @param request - GetAggregateResourceComplianceTimelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceComplianceTimelineResponse
func (client *Client) GetAggregateResourceComplianceTimelineWithOptions(request *GetAggregateResourceComplianceTimelineRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceComplianceTimelineResponse, _err error) {
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
		Action:      dara.String("GetAggregateResourceComplianceTimeline"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateResourceComplianceTimelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance timeline of a specific resource in an account group. A compliance timeline is a set of compliance evaluation records for a resource. Each record includes the time and content of an evaluation.
//
// Description:
//
// In Cloud Config, each resource has its own compliance timeline, which is composed of compliance evaluation records. A record is generated each time a rule is triggered to evaluate the resource. Rules can be triggered by configuration changes, periodic executions, or manual executions.
//
// This topic provides an example of how to query the compliance timeline for the resource `new-bucket` (an OSS bucket). The resource is in the `cn-hangzhou` region and belongs to the member account `100931896542****` within the account group `ca-5885626622af0008****`. The response shows that the resource\\"s compliance timeline includes records with the timestamps `1625200295276` (UTC+8: 2021-07-02 12:31:35) and `1625200228510` (UTC+8: 2021-07-02 12:30:28).
//
// @param request - GetAggregateResourceComplianceTimelineRequest
//
// @return GetAggregateResourceComplianceTimelineResponse
func (client *Client) GetAggregateResourceComplianceTimeline(request *GetAggregateResourceComplianceTimelineRequest) (_result *GetAggregateResourceComplianceTimelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateResourceComplianceTimelineResponse{}
	_body, _err := client.GetAggregateResourceComplianceTimelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration history of a specific resource in a specified account group. After you enable Cloud Config, the service records all configuration and relationship changes for your resources and organizes them into a configuration history. This history is saved for 10 years by default.
//
// Description:
//
// Cloud Config provides a configuration history for each resource within the monitoring scope:
//
// - For resources that already exist when you enable Cloud Config, the configuration history begins when the service is enabled.
//
// - For resources that are created after you enable Cloud Config, the configuration history begins when the resources are created. Cloud Config records resource configuration changes every 10 minutes. When the configuration of a resource changes, a node appears in the configuration history. This node contains the resource configuration details, change details, and the related management event.
//
// This topic provides an example of how to query the configuration history of an OSS bucket named `new-bucket`. The bucket is in the `cn-hangzhou` region, belongs to the member account `100931896542****`, and is part of the account group `ca-5885626622af0008****`. The returned result indicates that a configuration change for the resource was recorded at the UNIX timestamp `1624961112000` (UTC+8: 2021-06-29 18:05:12).
//
// @param request - GetAggregateResourceConfigurationTimelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceConfigurationTimelineResponse
func (client *Client) GetAggregateResourceConfigurationTimelineWithOptions(request *GetAggregateResourceConfigurationTimelineRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceConfigurationTimelineResponse, _err error) {
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
		Action:      dara.String("GetAggregateResourceConfigurationTimeline"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateResourceConfigurationTimelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration history of a specific resource in a specified account group. After you enable Cloud Config, the service records all configuration and relationship changes for your resources and organizes them into a configuration history. This history is saved for 10 years by default.
//
// Description:
//
// Cloud Config provides a configuration history for each resource within the monitoring scope:
//
// - For resources that already exist when you enable Cloud Config, the configuration history begins when the service is enabled.
//
// - For resources that are created after you enable Cloud Config, the configuration history begins when the resources are created. Cloud Config records resource configuration changes every 10 minutes. When the configuration of a resource changes, a node appears in the configuration history. This node contains the resource configuration details, change details, and the related management event.
//
// This topic provides an example of how to query the configuration history of an OSS bucket named `new-bucket`. The bucket is in the `cn-hangzhou` region, belongs to the member account `100931896542****`, and is part of the account group `ca-5885626622af0008****`. The returned result indicates that a configuration change for the resource was recorded at the UNIX timestamp `1624961112000` (UTC+8: 2021-06-29 18:05:12).
//
// @param request - GetAggregateResourceConfigurationTimelineRequest
//
// @return GetAggregateResourceConfigurationTimelineResponse
func (client *Client) GetAggregateResourceConfigurationTimeline(request *GetAggregateResourceConfigurationTimelineRequest) (_result *GetAggregateResourceConfigurationTimelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateResourceConfigurationTimelineResponse{}
	_body, _err := client.GetAggregateResourceConfigurationTimelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从地域维度查询指定账号组内资源的统计结果
//
// @param request - GetAggregateResourceCountsGroupByRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceCountsGroupByRegionResponse
func (client *Client) GetAggregateResourceCountsGroupByRegionWithOptions(request *GetAggregateResourceCountsGroupByRegionRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceCountsGroupByRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.FolderId) {
		query["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.ResourceAccountId) {
		query["ResourceAccountId"] = request.ResourceAccountId
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
		Action:      dara.String("GetAggregateResourceCountsGroupByRegion"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateResourceCountsGroupByRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从地域维度查询指定账号组内资源的统计结果
//
// @param request - GetAggregateResourceCountsGroupByRegionRequest
//
// @return GetAggregateResourceCountsGroupByRegionResponse
func (client *Client) GetAggregateResourceCountsGroupByRegion(request *GetAggregateResourceCountsGroupByRegionRequest) (_result *GetAggregateResourceCountsGroupByRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateResourceCountsGroupByRegionResponse{}
	_body, _err := client.GetAggregateResourceCountsGroupByRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从资源类型维度查询指定账号组内资源的统计结果
//
// @param request - GetAggregateResourceCountsGroupByResourceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceCountsGroupByResourceTypeResponse
func (client *Client) GetAggregateResourceCountsGroupByResourceTypeWithOptions(request *GetAggregateResourceCountsGroupByResourceTypeRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceCountsGroupByResourceTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.FolderId) {
		query["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceAccountId) {
		query["ResourceAccountId"] = request.ResourceAccountId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAggregateResourceCountsGroupByResourceType"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateResourceCountsGroupByResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从资源类型维度查询指定账号组内资源的统计结果
//
// @param request - GetAggregateResourceCountsGroupByResourceTypeRequest
//
// @return GetAggregateResourceCountsGroupByResourceTypeResponse
func (client *Client) GetAggregateResourceCountsGroupByResourceType(request *GetAggregateResourceCountsGroupByResourceTypeRequest) (_result *GetAggregateResourceCountsGroupByResourceTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateResourceCountsGroupByResourceTypeResponse{}
	_body, _err := client.GetAggregateResourceCountsGroupByResourceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询账号组内全局资源下载清单信息。
//
// @param request - GetAggregateResourceInventoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceInventoryResponse
func (client *Client) GetAggregateResourceInventoryWithOptions(request *GetAggregateResourceInventoryRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceInventoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAggregateResourceInventory"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregateResourceInventoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询账号组内全局资源下载清单信息。
//
// @param request - GetAggregateResourceInventoryRequest
//
// @return GetAggregateResourceInventoryResponse
func (client *Client) GetAggregateResourceInventory(request *GetAggregateResourceInventoryRequest) (_result *GetAggregateResourceInventoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregateResourceInventoryResponse{}
	_body, _err := client.GetAggregateResourceInventoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the name, creation time, members, and type of an account group.
//
// Description:
//
// This topic provides an example of how to query the details of the account group `ca-88ea626622af0055****`. The response shows that the account group name is `Test_Group`, the description is `Test Group`, the type is `CUSTOM` (custom account group), and the status is `1` (created).
//
// @param tmpReq - GetAggregatorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregatorResponse
func (client *Client) GetAggregatorWithOptions(tmpReq *GetAggregatorRequest, runtime *dara.RuntimeOptions) (_result *GetAggregatorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAggregatorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAggregator"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggregatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the name, creation time, members, and type of an account group.
//
// Description:
//
// This topic provides an example of how to query the details of the account group `ca-88ea626622af0055****`. The response shows that the account group name is `Test_Group`, the description is `Test Group`, the type is `CUSTOM` (custom account group), and the status is `1` (created).
//
// @param request - GetAggregatorRequest
//
// @return GetAggregatorResponse
func (client *Client) GetAggregator(request *GetAggregatorRequest) (_result *GetAggregatorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAggregatorResponse{}
	_body, _err := client.GetAggregatorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specific compliance package.
//
// Description:
//
// This topic provides an example of how to query the details of the compliance package `cp-fdc8626622af00f9****`. The response indicates that the compliance package is named `MLPS 2.0 Level 3 Pre-check Compliance Package`, its status is `ACTIVE`, and the risk level of its rules is `1` (high risk).
//
// @param tmpReq - GetCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCompliancePackResponse
func (client *Client) GetCompliancePackWithOptions(tmpReq *GetCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *GetCompliancePackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetCompliancePackShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCompliancePack"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specific compliance package.
//
// Description:
//
// This topic provides an example of how to query the details of the compliance package `cp-fdc8626622af00f9****`. The response indicates that the compliance package is named `MLPS 2.0 Level 3 Pre-check Compliance Package`, its status is `ACTIVE`, and the risk level of its rules is `1` (high risk).
//
// @param request - GetCompliancePackRequest
//
// @return GetCompliancePackResponse
func (client *Client) GetCompliancePack(request *GetCompliancePackRequest) (_result *GetCompliancePackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCompliancePackResponse{}
	_body, _err := client.GetCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取合规包的评估报告
//
// @param request - GetCompliancePackReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCompliancePackReportResponse
func (client *Client) GetCompliancePackReportWithOptions(request *GetCompliancePackReportRequest, runtime *dara.RuntimeOptions) (_result *GetCompliancePackReportResponse, _err error) {
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
		Action:      dara.String("GetCompliancePackReport"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCompliancePackReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取合规包的评估报告
//
// @param request - GetCompliancePackReportRequest
//
// @return GetCompliancePackReportResponse
func (client *Client) GetCompliancePackReport(request *GetCompliancePackReportRequest) (_result *GetCompliancePackReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCompliancePackReportResponse{}
	_body, _err := client.GetCompliancePackReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the compliance summary for the current account.
//
// Description:
//
// This topic provides an example of how to query the compliance summary by resource and rule for an account group.
//
// @param request - GetComplianceSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComplianceSummaryResponse
func (client *Client) GetComplianceSummaryWithOptions(runtime *dara.RuntimeOptions) (_result *GetComplianceSummaryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetComplianceSummary"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetComplianceSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance summary for the current account.
//
// Description:
//
// This topic provides an example of how to query the compliance summary by resource and rule for an account group.
//
// @return GetComplianceSummaryResponse
func (client *Client) GetComplianceSummary() (_result *GetComplianceSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetComplianceSummaryResponse{}
	_body, _err := client.GetComplianceSummaryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified delivery channel.
//
// @param request - GetConfigDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfigDeliveryChannelResponse
func (client *Client) GetConfigDeliveryChannelWithOptions(request *GetConfigDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *GetConfigDeliveryChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryChannelId) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConfigDeliveryChannel"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConfigDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified delivery channel.
//
// @param request - GetConfigDeliveryChannelRequest
//
// @return GetConfigDeliveryChannelResponse
func (client *Client) GetConfigDeliveryChannel(request *GetConfigDeliveryChannelRequest) (_result *GetConfigDeliveryChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetConfigDeliveryChannelResponse{}
	_body, _err := client.GetConfigDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified rule.
//
// Description:
//
// This topic provides an example of how to query the details of the rule `cr-7f7d626622af0041****`.
//
// @param tmpReq - GetConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfigRuleResponse
func (client *Client) GetConfigRuleWithOptions(tmpReq *GetConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *GetConfigRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetConfigRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigRuleId) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConfigRule"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified rule.
//
// Description:
//
// This topic provides an example of how to query the details of the rule `cr-7f7d626622af0041****`.
//
// @param request - GetConfigRuleRequest
//
// @return GetConfigRuleResponse
func (client *Client) GetConfigRule(request *GetConfigRuleRequest) (_result *GetConfigRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetConfigRuleResponse{}
	_body, _err := client.GetConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the compliance statistics for rules in a specified compliance package.
//
// Description:
//
// This example shows how to query the compliance results for rules in the compliance package `cp-541e626622af0087****`. The response shows that the total number of rules is 1 and the number of non-compliant rules is 0.
//
// @param request - GetConfigRuleComplianceByPackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfigRuleComplianceByPackResponse
func (client *Client) GetConfigRuleComplianceByPackWithOptions(request *GetConfigRuleComplianceByPackRequest, runtime *dara.RuntimeOptions) (_result *GetConfigRuleComplianceByPackResponse, _err error) {
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
		Action:      dara.String("GetConfigRuleComplianceByPack"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConfigRuleComplianceByPackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance statistics for rules in a specified compliance package.
//
// Description:
//
// This example shows how to query the compliance results for rules in the compliance package `cp-541e626622af0087****`. The response shows that the total number of rules is 1 and the number of non-compliant rules is 0.
//
// @param request - GetConfigRuleComplianceByPackRequest
//
// @return GetConfigRuleComplianceByPackResponse
func (client *Client) GetConfigRuleComplianceByPack(request *GetConfigRuleComplianceByPackRequest) (_result *GetConfigRuleComplianceByPackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetConfigRuleComplianceByPackResponse{}
	_body, _err := client.GetConfigRuleComplianceByPackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the compliance summary based on the risk level of a rule.
//
// Description:
//
// This topic provides an example of how to query the summary of compliance evaluation results by rule risk level. The return result shows four rules that are specified with the high risk level. One of them detects non-compliant resources, and the resources evaluated by the remaining three are all compliant.
//
// @param request - GetConfigRuleSummaryByRiskLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfigRuleSummaryByRiskLevelResponse
func (client *Client) GetConfigRuleSummaryByRiskLevelWithOptions(runtime *dara.RuntimeOptions) (_result *GetConfigRuleSummaryByRiskLevelResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetConfigRuleSummaryByRiskLevel"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConfigRuleSummaryByRiskLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance summary based on the risk level of a rule.
//
// Description:
//
// This topic provides an example of how to query the summary of compliance evaluation results by rule risk level. The return result shows four rules that are specified with the high risk level. One of them detects non-compliant resources, and the resources evaluated by the remaining three are all compliant.
//
// @return GetConfigRuleSummaryByRiskLevelResponse
func (client *Client) GetConfigRuleSummaryByRiskLevel() (_result *GetConfigRuleSummaryByRiskLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetConfigRuleSummaryByRiskLevelResponse{}
	_body, _err := client.GetConfigRuleSummaryByRiskLevelWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 下载Excel格式的规则合规评估报告到本地，便于您云下分配任务并跟进不合规资源配置的修改
//
// @param request - GetConfigRulesReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfigRulesReportResponse
func (client *Client) GetConfigRulesReportWithOptions(request *GetConfigRulesReportRequest, runtime *dara.RuntimeOptions) (_result *GetConfigRulesReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConfigRulesReport"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConfigRulesReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下载Excel格式的规则合规评估报告到本地，便于您云下分配任务并跟进不合规资源配置的修改
//
// @param request - GetConfigRulesReportRequest
//
// @return GetConfigRulesReportResponse
func (client *Client) GetConfigRulesReport(request *GetConfigRulesReportRequest) (_result *GetConfigRulesReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetConfigRulesReportResponse{}
	_body, _err := client.GetConfigRulesReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the Cloud Config service and the resource monitoring scope for the current account.
//
// Description:
//
// This topic provides an example of how to query the status of the Cloud Config service and the resource monitoring scope for the current account.
//
// @param request - GetConfigurationRecorderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfigurationRecorderResponse
func (client *Client) GetConfigurationRecorderWithOptions(runtime *dara.RuntimeOptions) (_result *GetConfigurationRecorderResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetConfigurationRecorder"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConfigurationRecorderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the Cloud Config service and the resource monitoring scope for the current account.
//
// Description:
//
// This topic provides an example of how to query the status of the Cloud Config service and the resource monitoring scope for the current account.
//
// @return GetConfigurationRecorderResponse
func (client *Client) GetConfigurationRecorder() (_result *GetConfigurationRecorderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetConfigurationRecorderResponse{}
	_body, _err := client.GetConfigurationRecorderWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specific resource.
//
// Description:
//
// This topic provides an example of how to query the details of the ECS instance `i-bp12g4xbl4i0brkn****` in the China (Hangzhou) region.
//
// @param request - GetDiscoveredResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDiscoveredResourceResponse
func (client *Client) GetDiscoveredResourceWithOptions(request *GetDiscoveredResourceRequest, runtime *dara.RuntimeOptions) (_result *GetDiscoveredResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComplianceOption) {
		query["ComplianceOption"] = request.ComplianceOption
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
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
		Action:      dara.String("GetDiscoveredResource"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDiscoveredResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specific resource.
//
// Description:
//
// This topic provides an example of how to query the details of the ECS instance `i-bp12g4xbl4i0brkn****` in the China (Hangzhou) region.
//
// @param request - GetDiscoveredResourceRequest
//
// @return GetDiscoveredResourceResponse
func (client *Client) GetDiscoveredResource(request *GetDiscoveredResourceRequest) (_result *GetDiscoveredResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDiscoveredResourceResponse{}
	_body, _err := client.GetDiscoveredResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从地域维度查询资源的统计结果
//
// @param request - GetDiscoveredResourceCountsGroupByRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDiscoveredResourceCountsGroupByRegionResponse
func (client *Client) GetDiscoveredResourceCountsGroupByRegionWithOptions(request *GetDiscoveredResourceCountsGroupByRegionRequest, runtime *dara.RuntimeOptions) (_result *GetDiscoveredResourceCountsGroupByRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDiscoveredResourceCountsGroupByRegion"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDiscoveredResourceCountsGroupByRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从地域维度查询资源的统计结果
//
// @param request - GetDiscoveredResourceCountsGroupByRegionRequest
//
// @return GetDiscoveredResourceCountsGroupByRegionResponse
func (client *Client) GetDiscoveredResourceCountsGroupByRegion(request *GetDiscoveredResourceCountsGroupByRegionRequest) (_result *GetDiscoveredResourceCountsGroupByRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDiscoveredResourceCountsGroupByRegionResponse{}
	_body, _err := client.GetDiscoveredResourceCountsGroupByRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从资源类型维度查询资源的统计结果
//
// @param request - GetDiscoveredResourceCountsGroupByResourceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDiscoveredResourceCountsGroupByResourceTypeResponse
func (client *Client) GetDiscoveredResourceCountsGroupByResourceTypeWithOptions(request *GetDiscoveredResourceCountsGroupByResourceTypeRequest, runtime *dara.RuntimeOptions) (_result *GetDiscoveredResourceCountsGroupByResourceTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDiscoveredResourceCountsGroupByResourceType"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDiscoveredResourceCountsGroupByResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从资源类型维度查询资源的统计结果
//
// @param request - GetDiscoveredResourceCountsGroupByResourceTypeRequest
//
// @return GetDiscoveredResourceCountsGroupByResourceTypeResponse
func (client *Client) GetDiscoveredResourceCountsGroupByResourceType(request *GetDiscoveredResourceCountsGroupByResourceTypeRequest) (_result *GetDiscoveredResourceCountsGroupByResourceTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDiscoveredResourceCountsGroupByResourceTypeResponse{}
	_body, _err := client.GetDiscoveredResourceCountsGroupByResourceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户集成云产品的授权状态
//
// @param request - GetIntegratedServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIntegratedServiceStatusResponse
func (client *Client) GetIntegratedServiceStatusWithOptions(request *GetIntegratedServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *GetIntegratedServiceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ServiceCode) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIntegratedServiceStatus"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIntegratedServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户集成云产品的授权状态
//
// @param request - GetIntegratedServiceStatusRequest
//
// @return GetIntegratedServiceStatusResponse
func (client *Client) GetIntegratedServiceStatus(request *GetIntegratedServiceStatusRequest) (_result *GetIntegratedServiceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIntegratedServiceStatusResponse{}
	_body, _err := client.GetIntegratedServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified rule template.
//
// Description:
//
// This topic provides an example of how to query the details of the rule template `cdn-domain-https-enabled`.
//
// @param request - GetManagedRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetManagedRuleResponse
func (client *Client) GetManagedRuleWithOptions(request *GetManagedRuleRequest, runtime *dara.RuntimeOptions) (_result *GetManagedRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Identifier) {
		query["Identifier"] = request.Identifier
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetManagedRule"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetManagedRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified rule template.
//
// Description:
//
// This topic provides an example of how to query the details of the rule template `cdn-domain-https-enabled`.
//
// @param request - GetManagedRuleRequest
//
// @return GetManagedRuleResponse
func (client *Client) GetManagedRule(request *GetManagedRuleRequest) (_result *GetManagedRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetManagedRuleResponse{}
	_body, _err := client.GetManagedRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an automatic remediation template.
//
// Description:
//
// This topic describes how to query the details of the `ACS-ALB-BulkyEnableDeletionProtection` automatic remediation template.
//
// @param request - GetRemediationTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRemediationTemplateResponse
func (client *Client) GetRemediationTemplateWithOptions(request *GetRemediationTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetRemediationTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateIdentifier) {
		query["TemplateIdentifier"] = request.TemplateIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRemediationTemplate"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRemediationTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an automatic remediation template.
//
// Description:
//
// This topic describes how to query the details of the `ACS-ALB-BulkyEnableDeletionProtection` automatic remediation template.
//
// @param request - GetRemediationTemplateRequest
//
// @return GetRemediationTemplateResponse
func (client *Client) GetRemediationTemplate(request *GetRemediationTemplateRequest) (_result *GetRemediationTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRemediationTemplateResponse{}
	_body, _err := client.GetRemediationTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a report generated from a report template.
//
// @param request - GetReportFromTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReportFromTemplateResponse
func (client *Client) GetReportFromTemplateWithOptions(request *GetReportFromTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetReportFromTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReportTemplateId) {
		query["ReportTemplateId"] = request.ReportTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetReportFromTemplate"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetReportFromTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a report generated from a report template.
//
// @param request - GetReportFromTemplateRequest
//
// @return GetReportFromTemplateResponse
func (client *Client) GetReportFromTemplate(request *GetReportFromTemplateRequest) (_result *GetReportFromTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetReportFromTemplateResponse{}
	_body, _err := client.GetReportFromTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve details of a compliance report template.
//
// @param request - GetReportTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReportTemplateResponse
func (client *Client) GetReportTemplateWithOptions(request *GetReportTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetReportTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReportTemplateId) {
		query["ReportTemplateId"] = request.ReportTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetReportTemplate"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetReportTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve details of a compliance report template.
//
// @param request - GetReportTemplateRequest
//
// @return GetReportTemplateResponse
func (client *Client) GetReportTemplate(request *GetReportTemplateRequest) (_result *GetReportTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetReportTemplateResponse{}
	_body, _err := client.GetReportTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从规则的合规评估结果维度查询合规概要
//
// @param request - GetResourceComplianceByConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceComplianceByConfigRuleResponse
func (client *Client) GetResourceComplianceByConfigRuleWithOptions(request *GetResourceComplianceByConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *GetResourceComplianceByConfigRuleResponse, _err error) {
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
		Action:      dara.String("GetResourceComplianceByConfigRule"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceComplianceByConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从规则的合规评估结果维度查询合规概要
//
// @param request - GetResourceComplianceByConfigRuleRequest
//
// @return GetResourceComplianceByConfigRuleResponse
func (client *Client) GetResourceComplianceByConfigRule(request *GetResourceComplianceByConfigRuleRequest) (_result *GetResourceComplianceByConfigRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceComplianceByConfigRuleResponse{}
	_body, _err := client.GetResourceComplianceByConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the compliance results for resources in a compliance package.
//
// Description:
//
// This topic provides an example of how to query the compliance results for resources in the compliance package `cp-541e626622af0087****`. The response shows that 7 of the 10 resources are non-compliant.
//
// @param request - GetResourceComplianceByPackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceComplianceByPackResponse
func (client *Client) GetResourceComplianceByPackWithOptions(request *GetResourceComplianceByPackRequest, runtime *dara.RuntimeOptions) (_result *GetResourceComplianceByPackResponse, _err error) {
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
		Action:      dara.String("GetResourceComplianceByPack"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceComplianceByPackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance results for resources in a compliance package.
//
// Description:
//
// This topic provides an example of how to query the compliance results for resources in the compliance package `cp-541e626622af0087****`. The response shows that 7 of the 10 resources are non-compliant.
//
// @param request - GetResourceComplianceByPackRequest
//
// @return GetResourceComplianceByPackResponse
func (client *Client) GetResourceComplianceByPack(request *GetResourceComplianceByPackRequest) (_result *GetResourceComplianceByPackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceComplianceByPackResponse{}
	_body, _err := client.GetResourceComplianceByPackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询合规情况按照地域分组统计
//
// @param request - GetResourceComplianceGroupByRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceComplianceGroupByRegionResponse
func (client *Client) GetResourceComplianceGroupByRegionWithOptions(request *GetResourceComplianceGroupByRegionRequest, runtime *dara.RuntimeOptions) (_result *GetResourceComplianceGroupByRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigRuleIds) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceComplianceGroupByRegion"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceComplianceGroupByRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询合规情况按照地域分组统计
//
// @param request - GetResourceComplianceGroupByRegionRequest
//
// @return GetResourceComplianceGroupByRegionResponse
func (client *Client) GetResourceComplianceGroupByRegion(request *GetResourceComplianceGroupByRegionRequest) (_result *GetResourceComplianceGroupByRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceComplianceGroupByRegionResponse{}
	_body, _err := client.GetResourceComplianceGroupByRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询规则对资源的评估结果，评估结果按资源类型进行分组展示
//
// @param request - GetResourceComplianceGroupByResourceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceComplianceGroupByResourceTypeResponse
func (client *Client) GetResourceComplianceGroupByResourceTypeWithOptions(request *GetResourceComplianceGroupByResourceTypeRequest, runtime *dara.RuntimeOptions) (_result *GetResourceComplianceGroupByResourceTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigRuleIds) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceComplianceGroupByResourceType"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceComplianceGroupByResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询规则对资源的评估结果，评估结果按资源类型进行分组展示
//
// @param request - GetResourceComplianceGroupByResourceTypeRequest
//
// @return GetResourceComplianceGroupByResourceTypeResponse
func (client *Client) GetResourceComplianceGroupByResourceType(request *GetResourceComplianceGroupByResourceTypeRequest) (_result *GetResourceComplianceGroupByResourceTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceComplianceGroupByResourceTypeResponse{}
	_body, _err := client.GetResourceComplianceGroupByResourceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the compliance evaluation history of a specified resource. The history is a set of compliance evaluation records that contain the timestamp and details of each evaluation.
//
// Description:
//
// In Cloud Config, each resource has its own compliance evaluation history. A compliance evaluation record is generated when a rule is triggered to evaluate a resource. The collection of these records forms the compliance evaluation history of the resource. Rules can be triggered by configuration changes, periodic execution, or manual execution.
//
// This topic provides an example of how to query the compliance evaluation history of the resource `new-bucket`, which is an Object Storage Service (OSS) bucket in the `cn-hangzhou` region. The returned result shows that the compliance evaluation history of the resource includes records with the timestamps `1625200295276` (UTC+8: 2021-07-02 12:31:35) and `1625200228510` (UTC+8: 2021-07-02 12:30:28).
//
// @param request - GetResourceComplianceTimelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceComplianceTimelineResponse
func (client *Client) GetResourceComplianceTimelineWithOptions(request *GetResourceComplianceTimelineRequest, runtime *dara.RuntimeOptions) (_result *GetResourceComplianceTimelineResponse, _err error) {
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
		Action:      dara.String("GetResourceComplianceTimeline"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceComplianceTimelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance evaluation history of a specified resource. The history is a set of compliance evaluation records that contain the timestamp and details of each evaluation.
//
// Description:
//
// In Cloud Config, each resource has its own compliance evaluation history. A compliance evaluation record is generated when a rule is triggered to evaluate a resource. The collection of these records forms the compliance evaluation history of the resource. Rules can be triggered by configuration changes, periodic execution, or manual execution.
//
// This topic provides an example of how to query the compliance evaluation history of the resource `new-bucket`, which is an Object Storage Service (OSS) bucket in the `cn-hangzhou` region. The returned result shows that the compliance evaluation history of the resource includes records with the timestamps `1625200295276` (UTC+8: 2021-07-02 12:31:35) and `1625200228510` (UTC+8: 2021-07-02 12:30:28).
//
// @param request - GetResourceComplianceTimelineRequest
//
// @return GetResourceComplianceTimelineResponse
func (client *Client) GetResourceComplianceTimeline(request *GetResourceComplianceTimelineRequest) (_result *GetResourceComplianceTimelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceComplianceTimelineResponse{}
	_body, _err := client.GetResourceComplianceTimelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation obtains configuration samples for a specified resource type.
//
// @param request - GetResourceConfigurationSampleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceConfigurationSampleResponse
func (client *Client) GetResourceConfigurationSampleWithOptions(request *GetResourceConfigurationSampleRequest, runtime *dara.RuntimeOptions) (_result *GetResourceConfigurationSampleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MockOnly) {
		query["MockOnly"] = request.MockOnly
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceConfigurationSample"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceConfigurationSampleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation obtains configuration samples for a specified resource type.
//
// @param request - GetResourceConfigurationSampleRequest
//
// @return GetResourceConfigurationSampleResponse
func (client *Client) GetResourceConfigurationSample(request *GetResourceConfigurationSampleRequest) (_result *GetResourceConfigurationSampleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceConfigurationSampleResponse{}
	_body, _err := client.GetResourceConfigurationSampleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration history of a specified resource. Cloud Config records every configuration and relationship change for a resource in a configuration history. Recording starts after you enable the Cloud Config service. By default, the history is retained for 10 years.
//
// Description:
//
// Cloud Config provides a configuration history for each resource that it monitors. The details are as follows:
//
// - For existing resources, the configuration history starts when you enable the Cloud Config service.
//
// - For new resources created after you enable the service, the configuration history starts when the resource is created. Cloud Config records configuration changes every 10 minutes. When a configuration changes, a new node appears in the history. This node contains the resource configuration details, change details, and the associated management event.
//
// This topic provides an example of how to query the configuration history for a resource named `new-bucket`. The resource is a bucket in the `cn-hangzhou` region. The response shows that the creation time of the resource is `1624961112000` (18:05:12 on June 29, 2021, UTC+8).
//
// @param request - GetResourceConfigurationTimelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceConfigurationTimelineResponse
func (client *Client) GetResourceConfigurationTimelineWithOptions(request *GetResourceConfigurationTimelineRequest, runtime *dara.RuntimeOptions) (_result *GetResourceConfigurationTimelineResponse, _err error) {
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
		Action:      dara.String("GetResourceConfigurationTimeline"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceConfigurationTimelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration history of a specified resource. Cloud Config records every configuration and relationship change for a resource in a configuration history. Recording starts after you enable the Cloud Config service. By default, the history is retained for 10 years.
//
// Description:
//
// Cloud Config provides a configuration history for each resource that it monitors. The details are as follows:
//
// - For existing resources, the configuration history starts when you enable the Cloud Config service.
//
// - For new resources created after you enable the service, the configuration history starts when the resource is created. Cloud Config records configuration changes every 10 minutes. When a configuration changes, a new node appears in the history. This node contains the resource configuration details, change details, and the associated management event.
//
// This topic provides an example of how to query the configuration history for a resource named `new-bucket`. The resource is a bucket in the `cn-hangzhou` region. The response shows that the creation time of the resource is `1624961112000` (18:05:12 on June 29, 2021, UTC+8).
//
// @param request - GetResourceConfigurationTimelineRequest
//
// @return GetResourceConfigurationTimelineResponse
func (client *Client) GetResourceConfigurationTimeline(request *GetResourceConfigurationTimelineRequest) (_result *GetResourceConfigurationTimelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceConfigurationTimelineResponse{}
	_body, _err := client.GetResourceConfigurationTimelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the most recently generated global resource inventory for the current account.
//
// Description:
//
// ### Prerequisites
//
// You can call the [GenerateResourceInventory](https://help.aliyun.com/document_detail/2398354.html) operation to generate a global resource inventory. You can then call this operation to obtain the URL of the global resource inventory.
//
// ### Usage notes
//
// This topic provides an example of how to query the most recently generated global resource inventory for the current account.
//
// @param request - GetResourceInventoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceInventoryResponse
func (client *Client) GetResourceInventoryWithOptions(runtime *dara.RuntimeOptions) (_result *GetResourceInventoryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceInventory"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceInventoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the most recently generated global resource inventory for the current account.
//
// Description:
//
// ### Prerequisites
//
// You can call the [GenerateResourceInventory](https://help.aliyun.com/document_detail/2398354.html) operation to generate a global resource inventory. You can then call this operation to obtain the URL of the global resource inventory.
//
// ### Usage notes
//
// This topic provides an example of how to query the most recently generated global resource inventory for the current account.
//
// @return GetResourceInventoryResponse
func (client *Client) GetResourceInventory() (_result *GetResourceInventoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceInventoryResponse{}
	_body, _err := client.GetResourceInventoryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation retrieves the property descriptions for a specified resource type.
//
// @param request - GetResourceTypePropertiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceTypePropertiesResponse
func (client *Client) GetResourceTypePropertiesWithOptions(request *GetResourceTypePropertiesRequest, runtime *dara.RuntimeOptions) (_result *GetResourceTypePropertiesResponse, _err error) {
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
		Action:      dara.String("GetResourceTypeProperties"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceTypePropertiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation retrieves the property descriptions for a specified resource type.
//
// @param request - GetResourceTypePropertiesRequest
//
// @return GetResourceTypePropertiesResponse
func (client *Client) GetResourceTypeProperties(request *GetResourceTypePropertiesRequest) (_result *GetResourceTypePropertiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceTypePropertiesResponse{}
	_body, _err := client.GetResourceTypePropertiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the supported resource relations for a specified resource type.
//
// Description:
//
// This topic provides an example of how to query the resource relations supported by the ACS::ECS::Instance resource type.
//
// @param request - GetSupportedResourceRelationConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupportedResourceRelationConfigResponse
func (client *Client) GetSupportedResourceRelationConfigWithOptions(request *GetSupportedResourceRelationConfigRequest, runtime *dara.RuntimeOptions) (_result *GetSupportedResourceRelationConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSupportedResourceRelationConfig"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSupportedResourceRelationConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the supported resource relations for a specified resource type.
//
// Description:
//
// This topic provides an example of how to query the resource relations supported by the ACS::ECS::Instance resource type.
//
// @param request - GetSupportedResourceRelationConfigRequest
//
// @return GetSupportedResourceRelationConfigResponse
func (client *Client) GetSupportedResourceRelationConfig(request *GetSupportedResourceRelationConfigRequest) (_result *GetSupportedResourceRelationConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSupportedResourceRelationConfigResponse{}
	_body, _err := client.GetSupportedResourceRelationConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 忽略评估结果增加截止时间
//
// @param tmpReq - IgnoreAggregateEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IgnoreAggregateEvaluationResultsResponse
func (client *Client) IgnoreAggregateEvaluationResultsWithOptions(tmpReq *IgnoreAggregateEvaluationResultsRequest, runtime *dara.RuntimeOptions) (_result *IgnoreAggregateEvaluationResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &IgnoreAggregateEvaluationResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Resources) {
		request.ResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Resources, dara.String("Resources"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ConfigRuleId) {
		body["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.IgnoreDate) {
		body["IgnoreDate"] = request.IgnoreDate
	}

	if !dara.IsNil(request.Reason) {
		body["Reason"] = request.Reason
	}

	if !dara.IsNil(request.ResourcesShrink) {
		body["Resources"] = request.ResourcesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IgnoreAggregateEvaluationResults"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IgnoreAggregateEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 忽略评估结果增加截止时间
//
// @param request - IgnoreAggregateEvaluationResultsRequest
//
// @return IgnoreAggregateEvaluationResultsResponse
func (client *Client) IgnoreAggregateEvaluationResults(request *IgnoreAggregateEvaluationResultsRequest) (_result *IgnoreAggregateEvaluationResultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &IgnoreAggregateEvaluationResultsResponse{}
	_body, _err := client.IgnoreAggregateEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Ignores the evaluation results of a rule for specific resources. You can also set a time period to ignore the rule. When the ignore period expires, the system automatically resumes displaying the evaluation results of the rule for the resources.
//
// Description:
//
// After a non-compliant resource is ignored, the rule still evaluates the resource. The evaluation result is Ignored.
//
// This topic provides an example of how to ignore the evaluation results of the rule `cr-7e72626622af0051****` for a specified non-compliant resource in the Alibaba Cloud account `100931896542****`. The region ID of the resource is `cn-beijing`, the resource type is `ACS::SLB::LoadBalancer`, and the resource ID is `lb-hp3a3b4ztyfm2plgm****`.
//
// @param tmpReq - IgnoreEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IgnoreEvaluationResultsResponse
func (client *Client) IgnoreEvaluationResultsWithOptions(tmpReq *IgnoreEvaluationResultsRequest, runtime *dara.RuntimeOptions) (_result *IgnoreEvaluationResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &IgnoreEvaluationResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Resources) {
		request.ResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Resources, dara.String("Resources"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigRuleId) {
		body["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.IgnoreDate) {
		body["IgnoreDate"] = request.IgnoreDate
	}

	if !dara.IsNil(request.Reason) {
		body["Reason"] = request.Reason
	}

	if !dara.IsNil(request.ResourcesShrink) {
		body["Resources"] = request.ResourcesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IgnoreEvaluationResults"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IgnoreEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Ignores the evaluation results of a rule for specific resources. You can also set a time period to ignore the rule. When the ignore period expires, the system automatically resumes displaying the evaluation results of the rule for the resources.
//
// Description:
//
// After a non-compliant resource is ignored, the rule still evaluates the resource. The evaluation result is Ignored.
//
// This topic provides an example of how to ignore the evaluation results of the rule `cr-7e72626622af0051****` for a specified non-compliant resource in the Alibaba Cloud account `100931896542****`. The region ID of the resource is `cn-beijing`, the resource type is `ACS::SLB::LoadBalancer`, and the resource ID is `lb-hp3a3b4ztyfm2plgm****`.
//
// @param request - IgnoreEvaluationResultsRequest
//
// @return IgnoreEvaluationResultsResponse
func (client *Client) IgnoreEvaluationResults(request *IgnoreEvaluationResultsRequest) (_result *IgnoreEvaluationResultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &IgnoreEvaluationResultsResponse{}
	_body, _err := client.IgnoreEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the compliance packs in a specified account group.
//
// Description:
//
// This topic provides an example of how to query the compliance packs in the account group `ca-f632626622af0079****`. The response shows that the account group contains the compliance pack `cp-fdc8626622af00f9****`.
//
// @param tmpReq - ListAggregateCompliancePacksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateCompliancePacksResponse
func (client *Client) ListAggregateCompliancePacksWithOptions(tmpReq *ListAggregateCompliancePacksRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateCompliancePacksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAggregateCompliancePacksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAggregateCompliancePacks"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAggregateCompliancePacksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance packs in a specified account group.
//
// Description:
//
// This topic provides an example of how to query the compliance packs in the account group `ca-f632626622af0079****`. The response shows that the account group contains the compliance pack `cp-fdc8626622af00f9****`.
//
// @param request - ListAggregateCompliancePacksRequest
//
// @return ListAggregateCompliancePacksResponse
func (client *Client) ListAggregateCompliancePacks(request *ListAggregateCompliancePacksRequest) (_result *ListAggregateCompliancePacksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAggregateCompliancePacksResponse{}
	_body, _err := client.ListAggregateCompliancePacksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 账号组查询投递渠道列表
//
// @param request - ListAggregateConfigDeliveryChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateConfigDeliveryChannelsResponse
func (client *Client) ListAggregateConfigDeliveryChannelsWithOptions(request *ListAggregateConfigDeliveryChannelsRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateConfigDeliveryChannelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.DeliveryChannelIds) {
		query["DeliveryChannelIds"] = request.DeliveryChannelIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAggregateConfigDeliveryChannels"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAggregateConfigDeliveryChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 账号组查询投递渠道列表
//
// @param request - ListAggregateConfigDeliveryChannelsRequest
//
// @return ListAggregateConfigDeliveryChannelsResponse
func (client *Client) ListAggregateConfigDeliveryChannels(request *ListAggregateConfigDeliveryChannelsRequest) (_result *ListAggregateConfigDeliveryChannelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAggregateConfigDeliveryChannelsResponse{}
	_body, _err := client.ListAggregateConfigDeliveryChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the evaluation results of a rule for resources in a specified account group.
//
// @param request - ListAggregateConfigRuleEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateConfigRuleEvaluationResultsResponse
func (client *Client) ListAggregateConfigRuleEvaluationResultsWithOptions(request *ListAggregateConfigRuleEvaluationResultsRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateConfigRuleEvaluationResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.CompliancePackId) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !dara.IsNil(request.ComplianceType) {
		query["ComplianceType"] = request.ComplianceType
	}

	if !dara.IsNil(request.ConfigRuleId) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Regions) {
		query["Regions"] = request.Regions
	}

	if !dara.IsNil(request.ResourceAccountId) {
		query["ResourceAccountId"] = request.ResourceAccountId
	}

	if !dara.IsNil(request.ResourceGroupIds) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAggregateConfigRuleEvaluationResults"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAggregateConfigRuleEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the evaluation results of a rule for resources in a specified account group.
//
// @param request - ListAggregateConfigRuleEvaluationResultsRequest
//
// @return ListAggregateConfigRuleEvaluationResultsResponse
func (client *Client) ListAggregateConfigRuleEvaluationResults(request *ListAggregateConfigRuleEvaluationResultsRequest) (_result *ListAggregateConfigRuleEvaluationResultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAggregateConfigRuleEvaluationResultsResponse{}
	_body, _err := client.ListAggregateConfigRuleEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 透出用户账号组维度的评估结果统计
//
// @param request - ListAggregateConfigRuleEvaluationStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateConfigRuleEvaluationStatisticsResponse
func (client *Client) ListAggregateConfigRuleEvaluationStatisticsWithOptions(request *ListAggregateConfigRuleEvaluationStatisticsRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateConfigRuleEvaluationStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAggregateConfigRuleEvaluationStatistics"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAggregateConfigRuleEvaluationStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 透出用户账号组维度的评估结果统计
//
// @param request - ListAggregateConfigRuleEvaluationStatisticsRequest
//
// @return ListAggregateConfigRuleEvaluationStatisticsResponse
func (client *Client) ListAggregateConfigRuleEvaluationStatistics(request *ListAggregateConfigRuleEvaluationStatisticsRequest) (_result *ListAggregateConfigRuleEvaluationStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAggregateConfigRuleEvaluationStatisticsResponse{}
	_body, _err := client.ListAggregateConfigRuleEvaluationStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the rules in a specified account group.
//
// Description:
//
// This topic provides an example of how to query the rules in the account group `ca-f632626622af0079****`. The response shows that the account group contains one rule. This rule evaluates two resources, and the compliance result is `COMPLIANT`.
//
// @param tmpReq - ListAggregateConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateConfigRulesResponse
func (client *Client) ListAggregateConfigRulesWithOptions(tmpReq *ListAggregateConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateConfigRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAggregateConfigRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.CompliancePackId) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !dara.IsNil(request.ComplianceType) {
		query["ComplianceType"] = request.ComplianceType
	}

	if !dara.IsNil(request.ConfigRuleName) {
		query["ConfigRuleName"] = request.ConfigRuleName
	}

	if !dara.IsNil(request.ConfigRuleState) {
		query["ConfigRuleState"] = request.ConfigRuleState
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAggregateConfigRules"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAggregateConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the rules in a specified account group.
//
// Description:
//
// This topic provides an example of how to query the rules in the account group `ca-f632626622af0079****`. The response shows that the account group contains one rule. This rule evaluates two resources, and the compliance result is `COMPLIANT`.
//
// @param request - ListAggregateConfigRulesRequest
//
// @return ListAggregateConfigRulesResponse
func (client *Client) ListAggregateConfigRules(request *ListAggregateConfigRulesRequest) (_result *ListAggregateConfigRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAggregateConfigRulesResponse{}
	_body, _err := client.ListAggregateConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resources in a specified account group.
//
// Description:
//
// ### Limits
//
// The resource checklist displays only some of your resources because Cloud Config supports only specific Alibaba Cloud services and resource types. For more information about the supported services and resource types, see [Supported resource types and resource relationships](https://help.aliyun.com/document_detail/127411.html).
//
// ### Usage notes
//
// This topic provides an example of how to query the resources of a member with the ID `100931896542****` in the account group `ca-c560626622af0005****`. The response indicates that there are eight resources in total.
//
// @param request - ListAggregateDiscoveredResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateDiscoveredResourcesResponse
func (client *Client) ListAggregateDiscoveredResourcesWithOptions(request *ListAggregateDiscoveredResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateDiscoveredResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.EndUpdateTimestamp) {
		query["EndUpdateTimestamp"] = request.EndUpdateTimestamp
	}

	if !dara.IsNil(request.ExcludeResourceTypes) {
		query["ExcludeResourceTypes"] = request.ExcludeResourceTypes
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Regions) {
		query["Regions"] = request.Regions
	}

	if !dara.IsNil(request.ResourceAccountId) {
		query["ResourceAccountId"] = request.ResourceAccountId
	}

	if !dara.IsNil(request.ResourceDeleted) {
		query["ResourceDeleted"] = request.ResourceDeleted
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	if !dara.IsNil(request.StartUpdateTimestamp) {
		query["StartUpdateTimestamp"] = request.StartUpdateTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAggregateDiscoveredResources"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAggregateDiscoveredResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resources in a specified account group.
//
// Description:
//
// ### Limits
//
// The resource checklist displays only some of your resources because Cloud Config supports only specific Alibaba Cloud services and resource types. For more information about the supported services and resource types, see [Supported resource types and resource relationships](https://help.aliyun.com/document_detail/127411.html).
//
// ### Usage notes
//
// This topic provides an example of how to query the resources of a member with the ID `100931896542****` in the account group `ca-c560626622af0005****`. The response indicates that there are eight resources in total.
//
// @param request - ListAggregateDiscoveredResourcesRequest
//
// @return ListAggregateDiscoveredResourcesResponse
func (client *Client) ListAggregateDiscoveredResources(request *ListAggregateDiscoveredResourcesRequest) (_result *ListAggregateDiscoveredResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAggregateDiscoveredResourcesResponse{}
	_body, _err := client.ListAggregateDiscoveredResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation retrieves the list of recommended managed rules for an account group.
//
// Description:
//
// For more information about the concepts, working principles, and integration process of dry run rules, see [Definition and working principles of dry run rules](https://help.aliyun.com/document_detail/470802.html).
//
// Dry run rules and rule templates are based on the same underlying rule definitions. After you create a resource, a dry run rule continuously checks the resource for compliance.
//
// @param request - ListAggregateRecommendManagedRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateRecommendManagedRulesResponse
func (client *Client) ListAggregateRecommendManagedRulesWithOptions(request *ListAggregateRecommendManagedRulesRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateRecommendManagedRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ExcludeRegionIdsScope) {
		query["ExcludeRegionIdsScope"] = request.ExcludeRegionIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceGroupIdsScope) {
		query["ExcludeResourceGroupIdsScope"] = request.ExcludeResourceGroupIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceIdsScope) {
		query["ExcludeResourceIdsScope"] = request.ExcludeResourceIdsScope
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionIdsScope) {
		query["RegionIdsScope"] = request.RegionIdsScope
	}

	if !dara.IsNil(request.ResourceGroupIdsScope) {
		query["ResourceGroupIdsScope"] = request.ResourceGroupIdsScope
	}

	if !dara.IsNil(request.ResourceIdsScope) {
		query["ResourceIdsScope"] = request.ResourceIdsScope
	}

	if !dara.IsNil(request.SelectedManagedRuleIdentifiers) {
		query["SelectedManagedRuleIdentifiers"] = request.SelectedManagedRuleIdentifiers
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAggregateRecommendManagedRules"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAggregateRecommendManagedRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation retrieves the list of recommended managed rules for an account group.
//
// Description:
//
// For more information about the concepts, working principles, and integration process of dry run rules, see [Definition and working principles of dry run rules](https://help.aliyun.com/document_detail/470802.html).
//
// Dry run rules and rule templates are based on the same underlying rule definitions. After you create a resource, a dry run rule continuously checks the resource for compliance.
//
// @param request - ListAggregateRecommendManagedRulesRequest
//
// @return ListAggregateRecommendManagedRulesResponse
func (client *Client) ListAggregateRecommendManagedRules(request *ListAggregateRecommendManagedRulesRequest) (_result *ListAggregateRecommendManagedRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAggregateRecommendManagedRulesResponse{}
	_body, _err := client.ListAggregateRecommendManagedRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 账号组规则修正执行历史
//
// @param request - ListAggregateRemediationExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateRemediationExecutionsResponse
func (client *Client) ListAggregateRemediationExecutionsWithOptions(request *ListAggregateRemediationExecutionsRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateRemediationExecutionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ConfigRuleId) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.ExecutionStatus) {
		query["ExecutionStatus"] = request.ExecutionStatus
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceAccountId) {
		query["ResourceAccountId"] = request.ResourceAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAggregateRemediationExecutions"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAggregateRemediationExecutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 账号组规则修正执行历史
//
// @param request - ListAggregateRemediationExecutionsRequest
//
// @return ListAggregateRemediationExecutionsResponse
func (client *Client) ListAggregateRemediationExecutions(request *ListAggregateRemediationExecutionsRequest) (_result *ListAggregateRemediationExecutionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAggregateRemediationExecutionsResponse{}
	_body, _err := client.ListAggregateRemediationExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the remediations in a specified account group.
//
// Description:
//
// This topic provides an example of how to query the remediation settings for the rule `cr-6b7c626622af00b4****` in the account group `ca-6b4a626622af0012****`.
//
// @param request - ListAggregateRemediationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateRemediationsResponse
func (client *Client) ListAggregateRemediationsWithOptions(request *ListAggregateRemediationsRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateRemediationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ConfigRuleIds) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAggregateRemediations"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAggregateRemediationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the remediations in a specified account group.
//
// Description:
//
// This topic provides an example of how to query the remediation settings for the rule `cr-6b7c626622af00b4****` in the account group `ca-6b4a626622af0012****`.
//
// @param request - ListAggregateRemediationsRequest
//
// @return ListAggregateRemediationsResponse
func (client *Client) ListAggregateRemediations(request *ListAggregateRemediationsRequest) (_result *ListAggregateRemediationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAggregateRemediationsResponse{}
	_body, _err := client.ListAggregateRemediationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the rule evaluation results for resources in a specified account group.
//
// @param request - ListAggregateResourceEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateResourceEvaluationResultsResponse
func (client *Client) ListAggregateResourceEvaluationResultsWithOptions(request *ListAggregateResourceEvaluationResultsRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateResourceEvaluationResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ComplianceType) {
		query["ComplianceType"] = request.ComplianceType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAggregateResourceEvaluationResults"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAggregateResourceEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the rule evaluation results for resources in a specified account group.
//
// @param request - ListAggregateResourceEvaluationResultsRequest
//
// @return ListAggregateResourceEvaluationResultsResponse
func (client *Client) ListAggregateResourceEvaluationResults(request *ListAggregateResourceEvaluationResultsRequest) (_result *ListAggregateResourceEvaluationResultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAggregateResourceEvaluationResultsResponse{}
	_body, _err := client.ListAggregateResourceEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取资源关系列表
//
// @param request - ListAggregateResourceRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateResourceRelationsResponse
func (client *Client) ListAggregateResourceRelationsWithOptions(request *ListAggregateResourceRelationsRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateResourceRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RelationType) {
		query["RelationType"] = request.RelationType
	}

	if !dara.IsNil(request.ResourceAccountId) {
		query["ResourceAccountId"] = request.ResourceAccountId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TargetResourceId) {
		query["TargetResourceId"] = request.TargetResourceId
	}

	if !dara.IsNil(request.TargetResourceType) {
		query["TargetResourceType"] = request.TargetResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAggregateResourceRelations"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAggregateResourceRelationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源关系列表
//
// @param request - ListAggregateResourceRelationsRequest
//
// @return ListAggregateResourceRelationsResponse
func (client *Client) ListAggregateResourceRelations(request *ListAggregateResourceRelationsRequest) (_result *ListAggregateResourceRelationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAggregateResourceRelationsResponse{}
	_body, _err := client.ListAggregateResourceRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can use SQL Select statements to search for resources in a specific account group based on fields in the resource properties.
//
// Description:
//
// When you write an SQL `Select` statement, you can retrieve the search fields and their types from the property file of the target resource type. For more information about resource property files, see [alibabacloud-config-resource-schema](https://github.com/aliyun/alibabacloud-config-resource-schema).
//
// > - The resource property files contain all resource types that are supported by Cloud Config. These files are named after their corresponding resource types. For example, the property file for the `ACS::ECS::Instance` resource type is `ACS_ECS_Instance.properties.json`. The path to the property files is `config/properties/resource-types`.
//
// >
//
// > - For more information about SQL search examples and limits, see [SQL search examples](https://help.aliyun.com/document_detail/398718.html) and [Limits of SQL search](https://help.aliyun.com/document_detail/398750.html).
//
// This topic provides an example of an advanced search query. The query finds all resources in the account group `ca-4b05626622af000c****` that have a tag key of `business` and a tag value of `online`.
//
// @param request - ListAggregateResourcesByAdvancedSearchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateResourcesByAdvancedSearchResponse
func (client *Client) ListAggregateResourcesByAdvancedSearchWithOptions(request *ListAggregateResourcesByAdvancedSearchRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateResourcesByAdvancedSearchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.Sql) {
		query["Sql"] = request.Sql
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAggregateResourcesByAdvancedSearch"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAggregateResourcesByAdvancedSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can use SQL Select statements to search for resources in a specific account group based on fields in the resource properties.
//
// Description:
//
// When you write an SQL `Select` statement, you can retrieve the search fields and their types from the property file of the target resource type. For more information about resource property files, see [alibabacloud-config-resource-schema](https://github.com/aliyun/alibabacloud-config-resource-schema).
//
// > - The resource property files contain all resource types that are supported by Cloud Config. These files are named after their corresponding resource types. For example, the property file for the `ACS::ECS::Instance` resource type is `ACS_ECS_Instance.properties.json`. The path to the property files is `config/properties/resource-types`.
//
// >
//
// > - For more information about SQL search examples and limits, see [SQL search examples](https://help.aliyun.com/document_detail/398718.html) and [Limits of SQL search](https://help.aliyun.com/document_detail/398750.html).
//
// This topic provides an example of an advanced search query. The query finds all resources in the account group `ca-4b05626622af000c****` that have a tag key of `business` and a tag value of `online`.
//
// @param request - ListAggregateResourcesByAdvancedSearchRequest
//
// @return ListAggregateResourcesByAdvancedSearchResponse
func (client *Client) ListAggregateResourcesByAdvancedSearch(request *ListAggregateResourcesByAdvancedSearchRequest) (_result *ListAggregateResourcesByAdvancedSearchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAggregateResourcesByAdvancedSearchResponse{}
	_body, _err := client.ListAggregateResourcesByAdvancedSearchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取账号组列表
//
// @param tmpReq - ListAggregatorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregatorsResponse
func (client *Client) ListAggregatorsWithOptions(tmpReq *ListAggregatorsRequest, runtime *dara.RuntimeOptions) (_result *ListAggregatorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAggregatorsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAggregators"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAggregatorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取账号组列表
//
// @param request - ListAggregatorsRequest
//
// @return ListAggregatorsResponse
func (client *Client) ListAggregators(request *ListAggregatorsRequest) (_result *ListAggregatorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAggregatorsResponse{}
	_body, _err := client.ListAggregatorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of all compliance pack templates provided by CloudConfig.
//
// Description:
//
// A compliance pack template is a collection of rules customized by CloudConfig for compliance scenarios.
//
// @param request - ListCompliancePackTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCompliancePackTemplatesResponse
func (client *Client) ListCompliancePackTemplatesWithOptions(request *ListCompliancePackTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListCompliancePackTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompliancePackTemplateId) {
		query["CompliancePackTemplateId"] = request.CompliancePackTemplateId
	}

	if !dara.IsNil(request.FilterType) {
		query["FilterType"] = request.FilterType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	if !dara.IsNil(request.RuleRiskLevel) {
		query["RuleRiskLevel"] = request.RuleRiskLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCompliancePackTemplates"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCompliancePackTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of all compliance pack templates provided by CloudConfig.
//
// Description:
//
// A compliance pack template is a collection of rules customized by CloudConfig for compliance scenarios.
//
// @param request - ListCompliancePackTemplatesRequest
//
// @return ListCompliancePackTemplatesResponse
func (client *Client) ListCompliancePackTemplates(request *ListCompliancePackTemplatesRequest) (_result *ListCompliancePackTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCompliancePackTemplatesResponse{}
	_body, _err := client.ListCompliancePackTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the compliance packs for the current account.
//
// Description:
//
// This topic provides an example of how to list compliance packs. The response returns one compliance pack record: `cp-fdc8626622af00f9****`.
//
// @param tmpReq - ListCompliancePacksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCompliancePacksResponse
func (client *Client) ListCompliancePacksWithOptions(tmpReq *ListCompliancePacksRequest, runtime *dara.RuntimeOptions) (_result *ListCompliancePacksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListCompliancePacksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCompliancePacks"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCompliancePacksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the compliance packs for the current account.
//
// Description:
//
// This topic provides an example of how to list compliance packs. The response returns one compliance pack record: `cp-fdc8626622af00f9****`.
//
// @param request - ListCompliancePacksRequest
//
// @return ListCompliancePacksResponse
func (client *Client) ListCompliancePacks(request *ListCompliancePacksRequest) (_result *ListCompliancePacksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCompliancePacksResponse{}
	_body, _err := client.ListCompliancePacksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Returns a list of delivery channels.
//
// @param request - ListConfigDeliveryChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConfigDeliveryChannelsResponse
func (client *Client) ListConfigDeliveryChannelsWithOptions(request *ListConfigDeliveryChannelsRequest, runtime *dara.RuntimeOptions) (_result *ListConfigDeliveryChannelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryChannelIds) {
		query["DeliveryChannelIds"] = request.DeliveryChannelIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConfigDeliveryChannels"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConfigDeliveryChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns a list of delivery channels.
//
// @param request - ListConfigDeliveryChannelsRequest
//
// @return ListConfigDeliveryChannelsResponse
func (client *Client) ListConfigDeliveryChannels(request *ListConfigDeliveryChannelsRequest) (_result *ListConfigDeliveryChannelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListConfigDeliveryChannelsResponse{}
	_body, _err := client.ListConfigDeliveryChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the compliance evaluation results of a rule.
//
// @param request - ListConfigRuleEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConfigRuleEvaluationResultsResponse
func (client *Client) ListConfigRuleEvaluationResultsWithOptions(request *ListConfigRuleEvaluationResultsRequest, runtime *dara.RuntimeOptions) (_result *ListConfigRuleEvaluationResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompliancePackId) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !dara.IsNil(request.ComplianceType) {
		query["ComplianceType"] = request.ComplianceType
	}

	if !dara.IsNil(request.ConfigRuleId) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Regions) {
		query["Regions"] = request.Regions
	}

	if !dara.IsNil(request.ResourceGroupIds) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConfigRuleEvaluationResults"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConfigRuleEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance evaluation results of a rule.
//
// @param request - ListConfigRuleEvaluationResultsRequest
//
// @return ListConfigRuleEvaluationResultsResponse
func (client *Client) ListConfigRuleEvaluationResults(request *ListConfigRuleEvaluationResultsRequest) (_result *ListConfigRuleEvaluationResultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListConfigRuleEvaluationResultsResponse{}
	_body, _err := client.ListConfigRuleEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 透出用户维度的历史评估结果统计
//
// @param request - ListConfigRuleEvaluationStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConfigRuleEvaluationStatisticsResponse
func (client *Client) ListConfigRuleEvaluationStatisticsWithOptions(runtime *dara.RuntimeOptions) (_result *ListConfigRuleEvaluationStatisticsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListConfigRuleEvaluationStatistics"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConfigRuleEvaluationStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 透出用户维度的历史评估结果统计
//
// @return ListConfigRuleEvaluationStatisticsResponse
func (client *Client) ListConfigRuleEvaluationStatistics() (_result *ListConfigRuleEvaluationStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListConfigRuleEvaluationStatisticsResponse{}
	_body, _err := client.ListConfigRuleEvaluationStatisticsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the operators available for Cloud Config rules.
//
// @param request - ListConfigRuleOperatorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConfigRuleOperatorsResponse
func (client *Client) ListConfigRuleOperatorsWithOptions(runtime *dara.RuntimeOptions) (_result *ListConfigRuleOperatorsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListConfigRuleOperators"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConfigRuleOperatorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the operators available for Cloud Config rules.
//
// @return ListConfigRuleOperatorsResponse
func (client *Client) ListConfigRuleOperators() (_result *ListConfigRuleOperatorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListConfigRuleOperatorsResponse{}
	_body, _err := client.ListConfigRuleOperatorsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of rules.
//
// Description:
//
// This topic provides an example of how to query the list of rules for the current account. The sample response indicates that the rule list contains one rule, three resources are evaluated, and the compliance result is COMPLIANT.
//
// @param tmpReq - ListConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConfigRulesResponse
func (client *Client) ListConfigRulesWithOptions(tmpReq *ListConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *ListConfigRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListConfigRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CompliancePackId) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !dara.IsNil(request.ComplianceType) {
		query["ComplianceType"] = request.ComplianceType
	}

	if !dara.IsNil(request.ConfigRuleName) {
		query["ConfigRuleName"] = request.ConfigRuleName
	}

	if !dara.IsNil(request.ConfigRuleState) {
		query["ConfigRuleState"] = request.ConfigRuleState
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConfigRules"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of rules.
//
// Description:
//
// This topic provides an example of how to query the list of rules for the current account. The sample response indicates that the rule list contains one rule, three resources are evaluated, and the compliance result is COMPLIANT.
//
// @param request - ListConfigRulesRequest
//
// @return ListConfigRulesResponse
func (client *Client) ListConfigRules(request *ListConfigRulesRequest) (_result *ListConfigRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListConfigRulesResponse{}
	_body, _err := client.ListConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of resources that are aggregated across regions in the current Alibaba Cloud account.
//
// Description:
//
// ### Limits
//
// Cloud Config supports only specific Alibaba Cloud services and resource types. The returned resource list includes only these supported resources. For more information about supported services and resource types, see [Supported resource types and resource relationships](https://help.aliyun.com/document_detail/127411.html).
//
// ### Usage notes
//
// This topic provides an example of how to query the resources in your account. The sample response shows that eight resources are returned.
//
// @param request - ListDiscoveredResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiscoveredResourcesResponse
func (client *Client) ListDiscoveredResourcesWithOptions(request *ListDiscoveredResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListDiscoveredResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndUpdateTimestamp) {
		query["EndUpdateTimestamp"] = request.EndUpdateTimestamp
	}

	if !dara.IsNil(request.ExcludeResourceTypes) {
		query["ExcludeResourceTypes"] = request.ExcludeResourceTypes
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Regions) {
		query["Regions"] = request.Regions
	}

	if !dara.IsNil(request.ResourceDeleted) {
		query["ResourceDeleted"] = request.ResourceDeleted
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	if !dara.IsNil(request.StartUpdateTimestamp) {
		query["StartUpdateTimestamp"] = request.StartUpdateTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDiscoveredResources"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDiscoveredResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of resources that are aggregated across regions in the current Alibaba Cloud account.
//
// Description:
//
// ### Limits
//
// Cloud Config supports only specific Alibaba Cloud services and resource types. The returned resource list includes only these supported resources. For more information about supported services and resource types, see [Supported resource types and resource relationships](https://help.aliyun.com/document_detail/127411.html).
//
// ### Usage notes
//
// This topic provides an example of how to query the resources in your account. The sample response shows that eight resources are returned.
//
// @param request - ListDiscoveredResourcesRequest
//
// @return ListDiscoveredResourcesResponse
func (client *Client) ListDiscoveredResources(request *ListDiscoveredResourcesRequest) (_result *ListDiscoveredResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDiscoveredResourcesResponse{}
	_body, _err := client.ListDiscoveredResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the services integrated with CloudConfig and their statuses.
//
// @param request - ListIntegratedServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegratedServiceResponse
func (client *Client) ListIntegratedServiceWithOptions(runtime *dara.RuntimeOptions) (_result *ListIntegratedServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntegratedService"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntegratedServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the services integrated with CloudConfig and their statuses.
//
// @return ListIntegratedServiceResponse
func (client *Client) ListIntegratedService() (_result *ListIntegratedServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIntegratedServiceResponse{}
	_body, _err := client.ListIntegratedServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the rule templates that are supported by Cloud Config.
//
// Description:
//
// ### Background information
//
// For more information about rule definitions, working principles, and templates, see [Definition and working principles of rules](https://help.aliyun.com/document_detail/128273.html).
//
// ### Usage notes
//
// This topic provides an example of how to query all rule templates that contain the keyword `CDN`. The sample response shows that a total of 21 rule templates are returned.
//
// @param request - ListManagedRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListManagedRulesResponse
func (client *Client) ListManagedRulesWithOptions(request *ListManagedRulesRequest, runtime *dara.RuntimeOptions) (_result *ListManagedRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterType) {
		query["FilterType"] = request.FilterType
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListManagedRules"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListManagedRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the rule templates that are supported by Cloud Config.
//
// Description:
//
// ### Background information
//
// For more information about rule definitions, working principles, and templates, see [Definition and working principles of rules](https://help.aliyun.com/document_detail/128273.html).
//
// ### Usage notes
//
// This topic provides an example of how to query all rule templates that contain the keyword `CDN`. The sample response shows that a total of 21 rule templates are returned.
//
// @param request - ListManagedRulesRequest
//
// @return ListManagedRulesResponse
func (client *Client) ListManagedRules(request *ListManagedRulesRequest) (_result *ListManagedRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListManagedRulesResponse{}
	_body, _err := client.ListManagedRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the proactive rules supported by Cloud Config.
//
// Description:
//
// For more information about the concepts, working principles, and integration procedure of proactive rules, see [Definitions and working principles of proactive rules](https://help.aliyun.com/document_detail/470802.html).
//
// Proactive rules and rule templates originate from the same source rules. After you create resources, proactive rules continuously check your resources for compliance.
//
// @param tmpReq - ListPreManagedRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPreManagedRulesResponse
func (client *Client) ListPreManagedRulesWithOptions(tmpReq *ListPreManagedRulesRequest, runtime *dara.RuntimeOptions) (_result *ListPreManagedRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPreManagedRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceTypes) {
		request.ResourceTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceTypes, dara.String("ResourceTypes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceTypesShrink) {
		body["ResourceTypes"] = request.ResourceTypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPreManagedRules"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPreManagedRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the proactive rules supported by Cloud Config.
//
// Description:
//
// For more information about the concepts, working principles, and integration procedure of proactive rules, see [Definitions and working principles of proactive rules](https://help.aliyun.com/document_detail/470802.html).
//
// Proactive rules and rule templates originate from the same source rules. After you create resources, proactive rules continuously check your resources for compliance.
//
// @param request - ListPreManagedRulesRequest
//
// @return ListPreManagedRulesResponse
func (client *Client) ListPreManagedRules(request *ListPreManagedRulesRequest) (_result *ListPreManagedRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPreManagedRulesResponse{}
	_body, _err := client.ListPreManagedRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation obtains a list of recommended managed rules.
//
// @param request - ListRecommendManagedRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecommendManagedRulesResponse
func (client *Client) ListRecommendManagedRulesWithOptions(request *ListRecommendManagedRulesRequest, runtime *dara.RuntimeOptions) (_result *ListRecommendManagedRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExcludeRegionIdsScope) {
		query["ExcludeRegionIdsScope"] = request.ExcludeRegionIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceGroupIdsScope) {
		query["ExcludeResourceGroupIdsScope"] = request.ExcludeResourceGroupIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceIdsScope) {
		query["ExcludeResourceIdsScope"] = request.ExcludeResourceIdsScope
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionIdsScope) {
		query["RegionIdsScope"] = request.RegionIdsScope
	}

	if !dara.IsNil(request.ResourceGroupIdsScope) {
		query["ResourceGroupIdsScope"] = request.ResourceGroupIdsScope
	}

	if !dara.IsNil(request.ResourceIdsScope) {
		query["ResourceIdsScope"] = request.ResourceIdsScope
	}

	if !dara.IsNil(request.SelectedManagedRuleIdentifiers) {
		query["SelectedManagedRuleIdentifiers"] = request.SelectedManagedRuleIdentifiers
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecommendManagedRules"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecommendManagedRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation obtains a list of recommended managed rules.
//
// @param request - ListRecommendManagedRulesRequest
//
// @return ListRecommendManagedRulesResponse
func (client *Client) ListRecommendManagedRules(request *ListRecommendManagedRulesRequest) (_result *ListRecommendManagedRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRecommendManagedRulesResponse{}
	_body, _err := client.ListRecommendManagedRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修正执行历史
//
// @param request - ListRemediationExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRemediationExecutionsResponse
func (client *Client) ListRemediationExecutionsWithOptions(request *ListRemediationExecutionsRequest, runtime *dara.RuntimeOptions) (_result *ListRemediationExecutionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigRuleId) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.ExecutionStatus) {
		query["ExecutionStatus"] = request.ExecutionStatus
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRemediationExecutions"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRemediationExecutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修正执行历史
//
// @param request - ListRemediationExecutionsRequest
//
// @return ListRemediationExecutionsResponse
func (client *Client) ListRemediationExecutions(request *ListRemediationExecutionsRequest) (_result *ListRemediationExecutionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRemediationExecutionsResponse{}
	_body, _err := client.ListRemediationExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修正模版列表
//
// @param request - ListRemediationTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRemediationTemplatesResponse
func (client *Client) ListRemediationTemplatesWithOptions(request *ListRemediationTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListRemediationTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ManagedRuleIdentifier) {
		query["ManagedRuleIdentifier"] = request.ManagedRuleIdentifier
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RemediationType) {
		query["RemediationType"] = request.RemediationType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRemediationTemplates"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRemediationTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修正模版列表
//
// @param request - ListRemediationTemplatesRequest
//
// @return ListRemediationTemplatesResponse
func (client *Client) ListRemediationTemplates(request *ListRemediationTemplatesRequest) (_result *ListRemediationTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRemediationTemplatesResponse{}
	_body, _err := client.ListRemediationTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询规则的修正
//
// @param request - ListRemediationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRemediationsResponse
func (client *Client) ListRemediationsWithOptions(request *ListRemediationsRequest, runtime *dara.RuntimeOptions) (_result *ListRemediationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigRuleIds) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
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
		Action:      dara.String("ListRemediations"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRemediationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询规则的修正
//
// @param request - ListRemediationsRequest
//
// @return ListRemediationsResponse
func (client *Client) ListRemediations(request *ListRemediationsRequest) (_result *ListRemediationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRemediationsResponse{}
	_body, _err := client.ListRemediationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of compliance report templates in batches.
//
// @param request - ListReportTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReportTemplatesResponse
func (client *Client) ListReportTemplatesWithOptions(request *ListReportTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListReportTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListReportTemplates"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListReportTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of compliance report templates in batches.
//
// @param request - ListReportTemplatesRequest
//
// @return ListReportTemplatesResponse
func (client *Client) ListReportTemplates(request *ListReportTemplatesRequest) (_result *ListReportTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListReportTemplatesResponse{}
	_body, _err := client.ListReportTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the evaluation results for a resource based on a rule.
//
// @param request - ListResourceEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceEvaluationResultsResponse
func (client *Client) ListResourceEvaluationResultsWithOptions(request *ListResourceEvaluationResultsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceEvaluationResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComplianceType) {
		query["ComplianceType"] = request.ComplianceType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceEvaluationResults"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the evaluation results for a resource based on a rule.
//
// @param request - ListResourceEvaluationResultsRequest
//
// @return ListResourceEvaluationResultsResponse
func (client *Client) ListResourceEvaluationResults(request *ListResourceEvaluationResultsRequest) (_result *ListResourceEvaluationResultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceEvaluationResultsResponse{}
	_body, _err := client.ListResourceEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取资源关系列表
//
// @param request - ListResourceRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceRelationsResponse
func (client *Client) ListResourceRelationsWithOptions(request *ListResourceRelationsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RelationType) {
		query["RelationType"] = request.RelationType
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TargetResourceId) {
		query["TargetResourceId"] = request.TargetResourceId
	}

	if !dara.IsNil(request.TargetResourceType) {
		query["TargetResourceType"] = request.TargetResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceRelations"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceRelationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源关系列表
//
// @param request - ListResourceRelationsRequest
//
// @return ListResourceRelationsResponse
func (client *Client) ListResourceRelations(request *ListResourceRelationsRequest) (_result *ListResourceRelationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceRelationsResponse{}
	_body, _err := client.ListResourceRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 资源列表高级搜索接口
//
// @param request - ListResourcesByAdvancedSearchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesByAdvancedSearchResponse
func (client *Client) ListResourcesByAdvancedSearchWithOptions(request *ListResourcesByAdvancedSearchRequest, runtime *dara.RuntimeOptions) (_result *ListResourcesByAdvancedSearchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Sql) {
		query["Sql"] = request.Sql
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourcesByAdvancedSearch"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourcesByAdvancedSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 资源列表高级搜索接口
//
// @param request - ListResourcesByAdvancedSearchRequest
//
// @return ListResourcesByAdvancedSearchResponse
func (client *Client) ListResourcesByAdvancedSearch(request *ListResourcesByAdvancedSearchRequest) (_result *ListResourcesByAdvancedSearchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourcesByAdvancedSearchResponse{}
	_body, _err := client.ListResourcesByAdvancedSearchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the cloud services and resource types that are supported by Cloud Config.
//
// Description:
//
// This topic provides an example on how to query the Alibaba Cloud services and resource types supported by a Cloud Config.
//
// @param request - ListSupportedProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSupportedProductsResponse
func (client *Client) ListSupportedProductsWithOptions(request *ListSupportedProductsRequest, runtime *dara.RuntimeOptions) (_result *ListSupportedProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSupportedProducts"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSupportedProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cloud services and resource types that are supported by Cloud Config.
//
// Description:
//
// This topic provides an example on how to query the Alibaba Cloud services and resource types supported by a Cloud Config.
//
// @param request - ListSupportedProductsRequest
//
// @return ListSupportedProductsResponse
func (client *Client) ListSupportedProducts(request *ListSupportedProductsRequest) (_result *ListSupportedProductsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSupportedProductsResponse{}
	_body, _err := client.ListSupportedProductsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags attached to resources in Cloud Config.
//
// @param tmpReq - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(tmpReq *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagShrink) {
		body["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2020-09-07"),
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
// Queries the tags attached to resources in Cloud Config.
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
// 自定义规则评估结果回调
//
// @param request - PutEvaluationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutEvaluationsResponse
func (client *Client) PutEvaluationsWithOptions(request *PutEvaluationsRequest, runtime *dara.RuntimeOptions) (_result *PutEvaluationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteMode) {
		body["DeleteMode"] = request.DeleteMode
	}

	if !dara.IsNil(request.Evaluations) {
		body["Evaluations"] = request.Evaluations
	}

	if !dara.IsNil(request.ResultToken) {
		body["ResultToken"] = request.ResultToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutEvaluations"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutEvaluationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 自定义规则评估结果回调
//
// @param request - PutEvaluationsRequest
//
// @return PutEvaluationsResponse
func (client *Client) PutEvaluations(request *PutEvaluationsRequest) (_result *PutEvaluationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutEvaluationsResponse{}
	_body, _err := client.PutEvaluationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消评估结果的忽略状态
//
// @param tmpReq - RevertAggregateEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevertAggregateEvaluationResultsResponse
func (client *Client) RevertAggregateEvaluationResultsWithOptions(tmpReq *RevertAggregateEvaluationResultsRequest, runtime *dara.RuntimeOptions) (_result *RevertAggregateEvaluationResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RevertAggregateEvaluationResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Resources) {
		request.ResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Resources, dara.String("Resources"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ConfigRuleId) {
		body["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.ResourcesShrink) {
		body["Resources"] = request.ResourcesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevertAggregateEvaluationResults"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevertAggregateEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消评估结果的忽略状态
//
// @param request - RevertAggregateEvaluationResultsRequest
//
// @return RevertAggregateEvaluationResultsResponse
func (client *Client) RevertAggregateEvaluationResults(request *RevertAggregateEvaluationResultsRequest) (_result *RevertAggregateEvaluationResultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevertAggregateEvaluationResultsResponse{}
	_body, _err := client.RevertAggregateEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 恢复已忽略的某条规则对某些资源的评估结果后，继续显示该规则对资源的评估结果
//
// @param tmpReq - RevertEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevertEvaluationResultsResponse
func (client *Client) RevertEvaluationResultsWithOptions(tmpReq *RevertEvaluationResultsRequest, runtime *dara.RuntimeOptions) (_result *RevertEvaluationResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RevertEvaluationResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Resources) {
		request.ResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Resources, dara.String("Resources"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigRuleId) {
		body["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.ResourcesShrink) {
		body["Resources"] = request.ResourcesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevertEvaluationResults"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevertEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 恢复已忽略的某条规则对某些资源的评估结果后，继续显示该规则对资源的评估结果
//
// @param request - RevertEvaluationResultsRequest
//
// @return RevertEvaluationResultsResponse
func (client *Client) RevertEvaluationResults(request *RevertEvaluationResultsRequest) (_result *RevertEvaluationResultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevertEvaluationResultsResponse{}
	_body, _err := client.RevertEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重新审计指定账号组内的某条规则或某个合规包中的所有规则
//
// @param request - StartAggregateConfigRuleEvaluationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAggregateConfigRuleEvaluationResponse
func (client *Client) StartAggregateConfigRuleEvaluationWithOptions(request *StartAggregateConfigRuleEvaluationRequest, runtime *dara.RuntimeOptions) (_result *StartAggregateConfigRuleEvaluationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.CompliancePackId) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !dara.IsNil(request.ConfigRuleId) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.RevertEvaluation) {
		query["RevertEvaluation"] = request.RevertEvaluation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartAggregateConfigRuleEvaluation"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartAggregateConfigRuleEvaluationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重新审计指定账号组内的某条规则或某个合规包中的所有规则
//
// @param request - StartAggregateConfigRuleEvaluationRequest
//
// @return StartAggregateConfigRuleEvaluationResponse
func (client *Client) StartAggregateConfigRuleEvaluation(request *StartAggregateConfigRuleEvaluationRequest) (_result *StartAggregateConfigRuleEvaluationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartAggregateConfigRuleEvaluationResponse{}
	_body, _err := client.StartAggregateConfigRuleEvaluationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Manually run remediation for a specified account group rule.
//
// Description:
//
// This topic describes how to manually run remediation once for the rule `cr-6b7c626622af00b4****` in the account group `ca-6b4a626622af0012****`. The response indicates that remediation completed successfully.
//
// @param request - StartAggregateRemediationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAggregateRemediationResponse
func (client *Client) StartAggregateRemediationWithOptions(request *StartAggregateRemediationRequest, runtime *dara.RuntimeOptions) (_result *StartAggregateRemediationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ConfigRuleId) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.ResourceAccountId) {
		query["ResourceAccountId"] = request.ResourceAccountId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartAggregateRemediation"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartAggregateRemediationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually run remediation for a specified account group rule.
//
// Description:
//
// This topic describes how to manually run remediation once for the rule `cr-6b7c626622af00b4****` in the account group `ca-6b4a626622af0012****`. The response indicates that remediation completed successfully.
//
// @param request - StartAggregateRemediationRequest
//
// @return StartAggregateRemediationResponse
func (client *Client) StartAggregateRemediation(request *StartAggregateRemediationRequest) (_result *StartAggregateRemediationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartAggregateRemediationResponse{}
	_body, _err := client.StartAggregateRemediationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Re-evaluates a specific rule or all rules in a compliance package.
//
// Description:
//
// This topic provides an example of how to re-evaluate the rule cr-9920626622af0035\\*\\*\\*\\*.
//
// @param request - StartConfigRuleEvaluationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartConfigRuleEvaluationResponse
func (client *Client) StartConfigRuleEvaluationWithOptions(request *StartConfigRuleEvaluationRequest, runtime *dara.RuntimeOptions) (_result *StartConfigRuleEvaluationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompliancePackId) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !dara.IsNil(request.ConfigRuleId) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.RevertEvaluation) {
		query["RevertEvaluation"] = request.RevertEvaluation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartConfigRuleEvaluation"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartConfigRuleEvaluationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Re-evaluates a specific rule or all rules in a compliance package.
//
// Description:
//
// This topic provides an example of how to re-evaluate the rule cr-9920626622af0035\\*\\*\\*\\*.
//
// @param request - StartConfigRuleEvaluationRequest
//
// @return StartConfigRuleEvaluationResponse
func (client *Client) StartConfigRuleEvaluation(request *StartConfigRuleEvaluationRequest) (_result *StartConfigRuleEvaluationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartConfigRuleEvaluationResponse{}
	_body, _err := client.StartConfigRuleEvaluationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a re-evaluation of a single resource.
//
// @param request - StartConfigRuleEvaluationByResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartConfigRuleEvaluationByResourceResponse
func (client *Client) StartConfigRuleEvaluationByResourceWithOptions(request *StartConfigRuleEvaluationByResourceRequest, runtime *dara.RuntimeOptions) (_result *StartConfigRuleEvaluationByResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartConfigRuleEvaluationByResource"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartConfigRuleEvaluationByResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a re-evaluation of a single resource.
//
// @param request - StartConfigRuleEvaluationByResourceRequest
//
// @return StartConfigRuleEvaluationByResourceResponse
func (client *Client) StartConfigRuleEvaluationByResource(request *StartConfigRuleEvaluationByResourceRequest) (_result *StartConfigRuleEvaluationByResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartConfigRuleEvaluationByResourceResponse{}
	_body, _err := client.StartConfigRuleEvaluationByResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启用配置审计服务
//
// @param request - StartConfigurationRecorderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartConfigurationRecorderResponse
func (client *Client) StartConfigurationRecorderWithOptions(runtime *dara.RuntimeOptions) (_result *StartConfigurationRecorderResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("StartConfigurationRecorder"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartConfigurationRecorderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启用配置审计服务
//
// @return StartConfigurationRecorderResponse
func (client *Client) StartConfigurationRecorder() (_result *StartConfigurationRecorderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartConfigurationRecorderResponse{}
	_body, _err := client.StartConfigurationRecorderWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Manually execute the specified rule remediation.
//
// Description:
//
// This topic provides an example of a manual remediation for rule `cr-8a973ac2e2be00a2****`. The returned result indicates a successful manual remediation.
//
// @param request - StartRemediationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRemediationResponse
func (client *Client) StartRemediationWithOptions(request *StartRemediationRequest, runtime *dara.RuntimeOptions) (_result *StartRemediationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigRuleId) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartRemediation"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartRemediationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually execute the specified rule remediation.
//
// Description:
//
// This topic provides an example of a manual remediation for rule `cr-8a973ac2e2be00a2****`. The returned result indicates a successful manual remediation.
//
// @param request - StartRemediationRequest
//
// @return StartRemediationResponse
func (client *Client) StartRemediation(request *StartRemediationRequest) (_result *StartRemediationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartRemediationResponse{}
	_body, _err := client.StartRemediationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停用配置审计服务
//
// @param request - StopConfigurationRecorderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopConfigurationRecorderResponse
func (client *Client) StopConfigurationRecorderWithOptions(runtime *dara.RuntimeOptions) (_result *StopConfigurationRecorderResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("StopConfigurationRecorder"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopConfigurationRecorderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停用配置审计服务
//
// @return StopConfigurationRecorderResponse
func (client *Client) StopConfigurationRecorder() (_result *StopConfigurationRecorderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopConfigurationRecorderResponse{}
	_body, _err := client.StopConfigurationRecorderWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches tags to CloudConfig rules, account groups, and compliance packages.
//
// @param tmpReq - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(tmpReq *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &TagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagShrink) {
		body["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2020-09-07"),
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
// Attaches tags to CloudConfig rules, account groups, and compliance packages.
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
// Sends a real-time test notification.
//
// @param request - TriggerReportSendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerReportSendResponse
func (client *Client) TriggerReportSendWithOptions(request *TriggerReportSendRequest, runtime *dara.RuntimeOptions) (_result *TriggerReportSendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ReportTemplateId) {
		body["ReportTemplateId"] = request.ReportTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TriggerReportSend"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TriggerReportSendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a real-time test notification.
//
// @param request - TriggerReportSendRequest
//
// @return TriggerReportSendResponse
func (client *Client) TriggerReportSend(request *TriggerReportSendRequest) (_result *TriggerReportSendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TriggerReportSendResponse{}
	_body, _err := client.TriggerReportSendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detaches tags from resources in Cloud Config.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		body["All"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
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
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2020-09-07"),
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
// Detaches tags from resources in Cloud Config.
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
// Updates a compliance pack in a specified account group.
//
// Description:
//
// This topic provides an example of how to change the value of a parameter for the `eip-bandwidth-limit` rule template to `20` in the `cp-fdc8626622af00f9****` compliance pack that belongs to the `ca-f632626622af0079****` account group.
//
// @param tmpReq - UpdateAggregateCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggregateCompliancePackResponse
func (client *Client) UpdateAggregateCompliancePackWithOptions(tmpReq *UpdateAggregateCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *UpdateAggregateCompliancePackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAggregateCompliancePackShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ConfigRules) {
		request.ConfigRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfigRules, dara.String("ConfigRules"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CompliancePackId) {
		body["CompliancePackId"] = request.CompliancePackId
	}

	if !dara.IsNil(request.CompliancePackName) {
		body["CompliancePackName"] = request.CompliancePackName
	}

	if !dara.IsNil(request.ConfigRulesShrink) {
		body["ConfigRules"] = request.ConfigRulesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExcludeRegionIdsScope) {
		body["ExcludeRegionIdsScope"] = request.ExcludeRegionIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceGroupIdsScope) {
		body["ExcludeResourceGroupIdsScope"] = request.ExcludeResourceGroupIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceIdsScope) {
		body["ExcludeResourceIdsScope"] = request.ExcludeResourceIdsScope
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ExcludeTagsScope) {
		bodyFlat["ExcludeTagsScope"] = request.ExcludeTagsScope
	}

	if !dara.IsNil(request.RegionIdsScope) {
		body["RegionIdsScope"] = request.RegionIdsScope
	}

	if !dara.IsNil(request.ResourceGroupIdsScope) {
		body["ResourceGroupIdsScope"] = request.ResourceGroupIdsScope
	}

	if !dara.IsNil(request.ResourceIdsScope) {
		body["ResourceIdsScope"] = request.ResourceIdsScope
	}

	if !dara.IsNil(request.RiskLevel) {
		body["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.TagKeyScope) {
		body["TagKeyScope"] = request.TagKeyScope
	}

	if !dara.IsNil(request.TagValueScope) {
		body["TagValueScope"] = request.TagValueScope
	}

	if !dara.IsNil(request.TagsScope) {
		bodyFlat["TagsScope"] = request.TagsScope
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAggregateCompliancePack"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAggregateCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a compliance pack in a specified account group.
//
// Description:
//
// This topic provides an example of how to change the value of a parameter for the `eip-bandwidth-limit` rule template to `20` in the `cp-fdc8626622af00f9****` compliance pack that belongs to the `ca-f632626622af0079****` account group.
//
// @param request - UpdateAggregateCompliancePackRequest
//
// @return UpdateAggregateCompliancePackResponse
func (client *Client) UpdateAggregateCompliancePack(request *UpdateAggregateCompliancePackRequest) (_result *UpdateAggregateCompliancePackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAggregateCompliancePackResponse{}
	_body, _err := client.UpdateAggregateCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 账号组修改投递渠道
//
// @param request - UpdateAggregateConfigDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggregateConfigDeliveryChannelResponse
func (client *Client) UpdateAggregateConfigDeliveryChannelWithOptions(request *UpdateAggregateConfigDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *UpdateAggregateConfigDeliveryChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CompliantSnapshot) {
		query["CompliantSnapshot"] = request.CompliantSnapshot
	}

	if !dara.IsNil(request.ConfigurationItemChangeNotification) {
		query["ConfigurationItemChangeNotification"] = request.ConfigurationItemChangeNotification
	}

	if !dara.IsNil(request.ConfigurationSnapshot) {
		query["ConfigurationSnapshot"] = request.ConfigurationSnapshot
	}

	if !dara.IsNil(request.DeliveryChannelCondition) {
		query["DeliveryChannelCondition"] = request.DeliveryChannelCondition
	}

	if !dara.IsNil(request.DeliveryChannelId) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	if !dara.IsNil(request.DeliveryChannelName) {
		query["DeliveryChannelName"] = request.DeliveryChannelName
	}

	if !dara.IsNil(request.DeliveryChannelTargetArn) {
		query["DeliveryChannelTargetArn"] = request.DeliveryChannelTargetArn
	}

	if !dara.IsNil(request.DeliverySnapshotTime) {
		query["DeliverySnapshotTime"] = request.DeliverySnapshotTime
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.NonCompliantNotification) {
		query["NonCompliantNotification"] = request.NonCompliantNotification
	}

	if !dara.IsNil(request.OversizedDataOSSTargetArn) {
		query["OversizedDataOSSTargetArn"] = request.OversizedDataOSSTargetArn
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAggregateConfigDeliveryChannel"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAggregateConfigDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 账号组修改投递渠道
//
// @param request - UpdateAggregateConfigDeliveryChannelRequest
//
// @return UpdateAggregateConfigDeliveryChannelResponse
func (client *Client) UpdateAggregateConfigDeliveryChannel(request *UpdateAggregateConfigDeliveryChannelRequest) (_result *UpdateAggregateConfigDeliveryChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAggregateConfigDeliveryChannelResponse{}
	_body, _err := client.UpdateAggregateConfigDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description, input parameters, and risk level of a rule in a specified account group.
//
// Description:
//
// This topic provides an example of how to change the risk level of the rule `cr-4e3d626622af0080****` in the account group `ca-a4e5626622af0079****` to `3` (low risk).
//
// @param tmpReq - UpdateAggregateConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggregateConfigRuleResponse
func (client *Client) UpdateAggregateConfigRuleWithOptions(tmpReq *UpdateAggregateConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateAggregateConfigRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAggregateConfigRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InputParameters) {
		request.InputParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InputParameters, dara.String("InputParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceTypesScope) {
		request.ResourceTypesScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceTypesScope, dara.String("ResourceTypesScope"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountIdsScope) {
		body["AccountIdsScope"] = request.AccountIdsScope
	}

	if !dara.IsNil(request.AggregatorId) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Conditions) {
		body["Conditions"] = request.Conditions
	}

	if !dara.IsNil(request.ConfigRuleId) {
		body["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.ConfigRuleName) {
		body["ConfigRuleName"] = request.ConfigRuleName
	}

	if !dara.IsNil(request.ConfigRuleTriggerTypes) {
		body["ConfigRuleTriggerTypes"] = request.ConfigRuleTriggerTypes
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExcludeAccountIdsScope) {
		body["ExcludeAccountIdsScope"] = request.ExcludeAccountIdsScope
	}

	if !dara.IsNil(request.ExcludeFolderIdsScope) {
		body["ExcludeFolderIdsScope"] = request.ExcludeFolderIdsScope
	}

	if !dara.IsNil(request.ExcludeRegionIdsScope) {
		body["ExcludeRegionIdsScope"] = request.ExcludeRegionIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceGroupIdsScope) {
		body["ExcludeResourceGroupIdsScope"] = request.ExcludeResourceGroupIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceIdsScope) {
		body["ExcludeResourceIdsScope"] = request.ExcludeResourceIdsScope
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ExcludeTagsScope) {
		bodyFlat["ExcludeTagsScope"] = request.ExcludeTagsScope
	}

	if !dara.IsNil(request.FolderIdsScope) {
		body["FolderIdsScope"] = request.FolderIdsScope
	}

	if !dara.IsNil(request.InputParametersShrink) {
		body["InputParameters"] = request.InputParametersShrink
	}

	if !dara.IsNil(request.MaximumExecutionFrequency) {
		body["MaximumExecutionFrequency"] = request.MaximumExecutionFrequency
	}

	if !dara.IsNil(request.RegionIdsScope) {
		body["RegionIdsScope"] = request.RegionIdsScope
	}

	if !dara.IsNil(request.ResourceGroupIdsScope) {
		body["ResourceGroupIdsScope"] = request.ResourceGroupIdsScope
	}

	if !dara.IsNil(request.ResourceIdsScope) {
		body["ResourceIdsScope"] = request.ResourceIdsScope
	}

	if !dara.IsNil(request.ResourceNameScope) {
		body["ResourceNameScope"] = request.ResourceNameScope
	}

	if !dara.IsNil(request.ResourceTypesScopeShrink) {
		body["ResourceTypesScope"] = request.ResourceTypesScopeShrink
	}

	if !dara.IsNil(request.RiskLevel) {
		body["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.TagKeyLogicScope) {
		body["TagKeyLogicScope"] = request.TagKeyLogicScope
	}

	if !dara.IsNil(request.TagKeyScope) {
		body["TagKeyScope"] = request.TagKeyScope
	}

	if !dara.IsNil(request.TagValueScope) {
		body["TagValueScope"] = request.TagValueScope
	}

	if !dara.IsNil(request.TagsScope) {
		bodyFlat["TagsScope"] = request.TagsScope
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAggregateConfigRule"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAggregateConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description, input parameters, and risk level of a rule in a specified account group.
//
// Description:
//
// This topic provides an example of how to change the risk level of the rule `cr-4e3d626622af0080****` in the account group `ca-a4e5626622af0079****` to `3` (low risk).
//
// @param request - UpdateAggregateConfigRuleRequest
//
// @return UpdateAggregateConfigRuleResponse
func (client *Client) UpdateAggregateConfigRule(request *UpdateAggregateConfigRuleRequest) (_result *UpdateAggregateConfigRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAggregateConfigRuleResponse{}
	_body, _err := client.UpdateAggregateConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a specified rule remediation.
//
// Description:
//
// This topic provides an example of how to change the execution mode of the remediation `crr-909ba2d4716700eb****` to `AUTO_EXECUTION` (automatic execution) for a rule in the account group `ca-6b4a626622af0012****`.
//
// @param request - UpdateAggregateRemediationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggregateRemediationResponse
func (client *Client) UpdateAggregateRemediationWithOptions(request *UpdateAggregateRemediationRequest, runtime *dara.RuntimeOptions) (_result *UpdateAggregateRemediationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorId) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.InvokeType) {
		body["InvokeType"] = request.InvokeType
	}

	if !dara.IsNil(request.Params) {
		body["Params"] = request.Params
	}

	if !dara.IsNil(request.RemediationId) {
		body["RemediationId"] = request.RemediationId
	}

	if !dara.IsNil(request.RemediationTemplateId) {
		body["RemediationTemplateId"] = request.RemediationTemplateId
	}

	if !dara.IsNil(request.RemediationType) {
		body["RemediationType"] = request.RemediationType
	}

	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAggregateRemediation"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAggregateRemediationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a specified rule remediation.
//
// Description:
//
// This topic provides an example of how to change the execution mode of the remediation `crr-909ba2d4716700eb****` to `AUTO_EXECUTION` (automatic execution) for a rule in the account group `ca-6b4a626622af0012****`.
//
// @param request - UpdateAggregateRemediationRequest
//
// @return UpdateAggregateRemediationResponse
func (client *Client) UpdateAggregateRemediation(request *UpdateAggregateRemediationRequest) (_result *UpdateAggregateRemediationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAggregateRemediationResponse{}
	_body, _err := client.UpdateAggregateRemediationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The management account or a delegated administrator account of a resource directory can modify the name and description of an account group, and add or remove members.
//
// Description:
//
// This topic provides an example of how to add a member to the account group `ca-dacf86d8314e00eb****`. The member has an ID of `173808452267****`, a name of `Tony`, and an account type of `ResourceDirectory`.
//
// @param tmpReq - UpdateAggregatorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggregatorResponse
func (client *Client) UpdateAggregatorWithOptions(tmpReq *UpdateAggregatorRequest, runtime *dara.RuntimeOptions) (_result *UpdateAggregatorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAggregatorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AggregatorAccounts) {
		request.AggregatorAccountsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AggregatorAccounts, dara.String("AggregatorAccounts"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorAccountsShrink) {
		body["AggregatorAccounts"] = request.AggregatorAccountsShrink
	}

	if !dara.IsNil(request.AggregatorId) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !dara.IsNil(request.AggregatorName) {
		body["AggregatorName"] = request.AggregatorName
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAggregator"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAggregatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The management account or a delegated administrator account of a resource directory can modify the name and description of an account group, and add or remove members.
//
// Description:
//
// This topic provides an example of how to add a member to the account group `ca-dacf86d8314e00eb****`. The member has an ID of `173808452267****`, a name of `Tony`, and an account type of `ResourceDirectory`.
//
// @param request - UpdateAggregatorRequest
//
// @return UpdateAggregatorResponse
func (client *Client) UpdateAggregator(request *UpdateAggregatorRequest) (_result *UpdateAggregatorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAggregatorResponse{}
	_body, _err := client.UpdateAggregatorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a specified compliance pack in the current account.
//
// Description:
//
// This topic provides an example of how to change the parameter value for the `eip-bandwidth-limit` rule to `20` in the compliance pack `cp-a8a8626622af0082****`.
//
// @param tmpReq - UpdateCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCompliancePackResponse
func (client *Client) UpdateCompliancePackWithOptions(tmpReq *UpdateCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *UpdateCompliancePackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCompliancePackShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ConfigRules) {
		request.ConfigRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfigRules, dara.String("ConfigRules"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CompliancePackId) {
		body["CompliancePackId"] = request.CompliancePackId
	}

	if !dara.IsNil(request.CompliancePackName) {
		body["CompliancePackName"] = request.CompliancePackName
	}

	if !dara.IsNil(request.ConfigRulesShrink) {
		body["ConfigRules"] = request.ConfigRulesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExcludeRegionIdsScope) {
		body["ExcludeRegionIdsScope"] = request.ExcludeRegionIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceGroupIdsScope) {
		body["ExcludeResourceGroupIdsScope"] = request.ExcludeResourceGroupIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceIdsScope) {
		body["ExcludeResourceIdsScope"] = request.ExcludeResourceIdsScope
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ExcludeTagsScope) {
		bodyFlat["ExcludeTagsScope"] = request.ExcludeTagsScope
	}

	if !dara.IsNil(request.RegionIdsScope) {
		body["RegionIdsScope"] = request.RegionIdsScope
	}

	if !dara.IsNil(request.ResourceGroupIdsScope) {
		body["ResourceGroupIdsScope"] = request.ResourceGroupIdsScope
	}

	if !dara.IsNil(request.ResourceIdsScope) {
		body["ResourceIdsScope"] = request.ResourceIdsScope
	}

	if !dara.IsNil(request.RiskLevel) {
		body["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.TagKeyScope) {
		body["TagKeyScope"] = request.TagKeyScope
	}

	if !dara.IsNil(request.TagValueScope) {
		body["TagValueScope"] = request.TagValueScope
	}

	if !dara.IsNil(request.TagsScope) {
		bodyFlat["TagsScope"] = request.TagsScope
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCompliancePack"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a specified compliance pack in the current account.
//
// Description:
//
// This topic provides an example of how to change the parameter value for the `eip-bandwidth-limit` rule to `20` in the compliance pack `cp-a8a8626622af0082****`.
//
// @param request - UpdateCompliancePackRequest
//
// @return UpdateCompliancePackResponse
func (client *Client) UpdateCompliancePack(request *UpdateCompliancePackRequest) (_result *UpdateCompliancePackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCompliancePackResponse{}
	_body, _err := client.UpdateCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation modifies a delivery channel for the current account.
//
// Description:
//
// This topic provides an example of how to change the status of the delivery channel `cdc-8e45ff4e06a3a8****` to `0` (disabled). After you disable the delivery channel, Cloud Config retains the most recent delivery configuration and stops delivering resource data.
//
// @param request - UpdateConfigDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConfigDeliveryChannelResponse
func (client *Client) UpdateConfigDeliveryChannelWithOptions(request *UpdateConfigDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *UpdateConfigDeliveryChannelResponse, _err error) {
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

	if !dara.IsNil(request.CompliantSnapshot) {
		query["CompliantSnapshot"] = request.CompliantSnapshot
	}

	if !dara.IsNil(request.ConfigurationItemChangeNotification) {
		query["ConfigurationItemChangeNotification"] = request.ConfigurationItemChangeNotification
	}

	if !dara.IsNil(request.ConfigurationSnapshot) {
		query["ConfigurationSnapshot"] = request.ConfigurationSnapshot
	}

	if !dara.IsNil(request.DeliveryChannelCondition) {
		query["DeliveryChannelCondition"] = request.DeliveryChannelCondition
	}

	if !dara.IsNil(request.DeliveryChannelId) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	if !dara.IsNil(request.DeliveryChannelName) {
		query["DeliveryChannelName"] = request.DeliveryChannelName
	}

	if !dara.IsNil(request.DeliveryChannelTargetArn) {
		query["DeliveryChannelTargetArn"] = request.DeliveryChannelTargetArn
	}

	if !dara.IsNil(request.DeliverySnapshotTime) {
		query["DeliverySnapshotTime"] = request.DeliverySnapshotTime
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.NonCompliantNotification) {
		query["NonCompliantNotification"] = request.NonCompliantNotification
	}

	if !dara.IsNil(request.OversizedDataOSSTargetArn) {
		query["OversizedDataOSSTargetArn"] = request.OversizedDataOSSTargetArn
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConfigDeliveryChannel"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConfigDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation modifies a delivery channel for the current account.
//
// Description:
//
// This topic provides an example of how to change the status of the delivery channel `cdc-8e45ff4e06a3a8****` to `0` (disabled). After you disable the delivery channel, Cloud Config retains the most recent delivery configuration and stops delivering resource data.
//
// @param request - UpdateConfigDeliveryChannelRequest
//
// @return UpdateConfigDeliveryChannelResponse
func (client *Client) UpdateConfigDeliveryChannel(request *UpdateConfigDeliveryChannelRequest) (_result *UpdateConfigDeliveryChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateConfigDeliveryChannelResponse{}
	_body, _err := client.UpdateConfigDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description, input parameters, and risk level of a rule.
//
// Description:
//
// This topic shows how to change the risk level of rule `cr-a260626622af0005****` to `3`, which indicates low risk.
//
// @param tmpReq - UpdateConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConfigRuleResponse
func (client *Client) UpdateConfigRuleWithOptions(tmpReq *UpdateConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateConfigRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateConfigRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InputParameters) {
		request.InputParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InputParameters, dara.String("InputParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceTypesScope) {
		request.ResourceTypesScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceTypesScope, dara.String("ResourceTypesScope"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Conditions) {
		body["Conditions"] = request.Conditions
	}

	if !dara.IsNil(request.ConfigRuleId) {
		body["ConfigRuleId"] = request.ConfigRuleId
	}

	if !dara.IsNil(request.ConfigRuleName) {
		body["ConfigRuleName"] = request.ConfigRuleName
	}

	if !dara.IsNil(request.ConfigRuleTriggerTypes) {
		body["ConfigRuleTriggerTypes"] = request.ConfigRuleTriggerTypes
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExcludeRegionIdsScope) {
		body["ExcludeRegionIdsScope"] = request.ExcludeRegionIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceGroupIdsScope) {
		body["ExcludeResourceGroupIdsScope"] = request.ExcludeResourceGroupIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceIdsScope) {
		body["ExcludeResourceIdsScope"] = request.ExcludeResourceIdsScope
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ExcludeTagsScope) {
		bodyFlat["ExcludeTagsScope"] = request.ExcludeTagsScope
	}

	if !dara.IsNil(request.ExtendContent) {
		body["ExtendContent"] = request.ExtendContent
	}

	if !dara.IsNil(request.InputParametersShrink) {
		body["InputParameters"] = request.InputParametersShrink
	}

	if !dara.IsNil(request.MaximumExecutionFrequency) {
		body["MaximumExecutionFrequency"] = request.MaximumExecutionFrequency
	}

	if !dara.IsNil(request.RegionIdsScope) {
		body["RegionIdsScope"] = request.RegionIdsScope
	}

	if !dara.IsNil(request.ResourceGroupIdsScope) {
		body["ResourceGroupIdsScope"] = request.ResourceGroupIdsScope
	}

	if !dara.IsNil(request.ResourceIdsScope) {
		body["ResourceIdsScope"] = request.ResourceIdsScope
	}

	if !dara.IsNil(request.ResourceNameScope) {
		body["ResourceNameScope"] = request.ResourceNameScope
	}

	if !dara.IsNil(request.ResourceTypesScopeShrink) {
		body["ResourceTypesScope"] = request.ResourceTypesScopeShrink
	}

	if !dara.IsNil(request.RiskLevel) {
		body["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.TagKeyLogicScope) {
		body["TagKeyLogicScope"] = request.TagKeyLogicScope
	}

	if !dara.IsNil(request.TagKeyScope) {
		body["TagKeyScope"] = request.TagKeyScope
	}

	if !dara.IsNil(request.TagValueScope) {
		body["TagValueScope"] = request.TagValueScope
	}

	if !dara.IsNil(request.TagsScope) {
		bodyFlat["TagsScope"] = request.TagsScope
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConfigRule"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description, input parameters, and risk level of a rule.
//
// Description:
//
// This topic shows how to change the risk level of rule `cr-a260626622af0005****` to `3`, which indicates low risk.
//
// @param request - UpdateConfigRuleRequest
//
// @return UpdateConfigRuleResponse
func (client *Client) UpdateConfigRule(request *UpdateConfigRuleRequest) (_result *UpdateConfigRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateConfigRuleResponse{}
	_body, _err := client.UpdateConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改配置审计监控资源范围
//
// @param request - UpdateConfigurationRecorderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConfigurationRecorderResponse
func (client *Client) UpdateConfigurationRecorderWithOptions(request *UpdateConfigurationRecorderRequest, runtime *dara.RuntimeOptions) (_result *UpdateConfigurationRecorderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceTypes) {
		body["ResourceTypes"] = request.ResourceTypes
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConfigurationRecorder"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConfigurationRecorderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改配置审计监控资源范围
//
// @param request - UpdateConfigurationRecorderRequest
//
// @return UpdateConfigurationRecorderResponse
func (client *Client) UpdateConfigurationRecorder(request *UpdateConfigurationRecorderRequest) (_result *UpdateConfigurationRecorderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateConfigurationRecorderResponse{}
	_body, _err := client.UpdateConfigurationRecorderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改云产品集成用户状态
//
// @param request - UpdateIntegratedServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIntegratedServiceStatusResponse
func (client *Client) UpdateIntegratedServiceStatusWithOptions(request *UpdateIntegratedServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateIntegratedServiceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AggregatorDeliveryDataType) {
		body["AggregatorDeliveryDataType"] = request.AggregatorDeliveryDataType
	}

	if !dara.IsNil(request.IntegratedTypes) {
		body["IntegratedTypes"] = request.IntegratedTypes
	}

	if !dara.IsNil(request.ServiceCode) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIntegratedServiceStatus"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIntegratedServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改云产品集成用户状态
//
// @param request - UpdateIntegratedServiceStatusRequest
//
// @return UpdateIntegratedServiceStatusResponse
func (client *Client) UpdateIntegratedServiceStatus(request *UpdateIntegratedServiceStatusRequest) (_result *UpdateIntegratedServiceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIntegratedServiceStatusResponse{}
	_body, _err := client.UpdateIntegratedServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the specified remediation setting.
//
// Description:
//
// This topic provides an example of how to change the execution mode for the remediation setting `crr-909ba2d4716700eb****` to `AUTO_EXECUTION` (automatic execution).
//
// @param request - UpdateRemediationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRemediationResponse
func (client *Client) UpdateRemediationWithOptions(request *UpdateRemediationRequest, runtime *dara.RuntimeOptions) (_result *UpdateRemediationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InvokeType) {
		body["InvokeType"] = request.InvokeType
	}

	if !dara.IsNil(request.Params) {
		body["Params"] = request.Params
	}

	if !dara.IsNil(request.RemediationId) {
		body["RemediationId"] = request.RemediationId
	}

	if !dara.IsNil(request.RemediationTemplateId) {
		body["RemediationTemplateId"] = request.RemediationTemplateId
	}

	if !dara.IsNil(request.RemediationType) {
		body["RemediationType"] = request.RemediationType
	}

	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRemediation"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRemediationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the specified remediation setting.
//
// Description:
//
// This topic provides an example of how to change the execution mode for the remediation setting `crr-909ba2d4716700eb****` to `AUTO_EXECUTION` (automatic execution).
//
// @param request - UpdateRemediationRequest
//
// @return UpdateRemediationResponse
func (client *Client) UpdateRemediation(request *UpdateRemediationRequest) (_result *UpdateRemediationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRemediationResponse{}
	_body, _err := client.UpdateRemediationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a compliance report template.
//
// @param tmpReq - UpdateReportTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateReportTemplateResponse
func (client *Client) UpdateReportTemplateWithOptions(tmpReq *UpdateReportTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateReportTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateReportTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReportScope) {
		request.ReportScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReportScope, dara.String("ReportScope"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ReportFileFormats) {
		body["ReportFileFormats"] = request.ReportFileFormats
	}

	if !dara.IsNil(request.ReportGranularity) {
		body["ReportGranularity"] = request.ReportGranularity
	}

	if !dara.IsNil(request.ReportLanguage) {
		body["ReportLanguage"] = request.ReportLanguage
	}

	if !dara.IsNil(request.ReportScopeShrink) {
		body["ReportScope"] = request.ReportScopeShrink
	}

	if !dara.IsNil(request.ReportTemplateDescription) {
		body["ReportTemplateDescription"] = request.ReportTemplateDescription
	}

	if !dara.IsNil(request.ReportTemplateId) {
		body["ReportTemplateId"] = request.ReportTemplateId
	}

	if !dara.IsNil(request.ReportTemplateName) {
		body["ReportTemplateName"] = request.ReportTemplateName
	}

	if !dara.IsNil(request.SubscriptionFrequency) {
		body["SubscriptionFrequency"] = request.SubscriptionFrequency
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateReportTemplate"),
		Version:     dara.String("2020-09-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateReportTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a compliance report template.
//
// @param request - UpdateReportTemplateRequest
//
// @return UpdateReportTemplateResponse
func (client *Client) UpdateReportTemplate(request *UpdateReportTemplateRequest) (_result *UpdateReportTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateReportTemplateResponse{}
	_body, _err := client.UpdateReportTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
