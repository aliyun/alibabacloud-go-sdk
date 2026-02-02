// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Enables one or more rules in an account group. After a rule is enabled, the rule continues to automatically evaluate resources based on the trigger mechanism.
//
// Description:
//
// Enables one or more rules in an account group. After a rule is enabled, the rule continues to automatically evaluate resources based on the trigger mechanism.
//
// @param request - ActiveAggregateConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActiveAggregateConfigRulesResponse
func (client *Client) ActiveAggregateConfigRulesWithContext(ctx context.Context, request *ActiveAggregateConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *ActiveAggregateConfigRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a rule in Cloud Config. After a rule is enabled, Cloud Config automatically evaluates the compliance of a resource based on the trigger mechanism of the rule.
//
// Description:
//
// ### [](#)Prerequisites
//
// The rule is in the `INACTIVE` state.
//
// @param request - ActiveConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActiveConfigRulesResponse
func (client *Client) ActiveConfigRulesWithContext(ctx context.Context, request *ActiveConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *ActiveConfigRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds one or more rules in an account group to a compliance package.
//
// Description:
//
// The sample request in this topic shows you how to add the `cr-6cc4626622af00e7****` rule in the `ca-75b4626622af00c3****` account group to the `cp-5bb1626622af00bd****` compliance package.
//
// @param request - AttachAggregateConfigRuleToCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachAggregateConfigRuleToCompliancePackResponse
func (client *Client) AttachAggregateConfigRuleToCompliancePackWithContext(ctx context.Context, request *AttachAggregateConfigRuleToCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *AttachAggregateConfigRuleToCompliancePackResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds one or more rules to a compliance package.
//
// Description:
//
// This topic provides an example on how to add the `cr-6cc4626622af00e7****` rule to the `cp-5bb1626622af00bd****` compliance package.
//
// @param request - AttachConfigRuleToCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachConfigRuleToCompliancePackResponse
func (client *Client) AttachConfigRuleToCompliancePackWithContext(ctx context.Context, request *AttachConfigRuleToCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *AttachConfigRuleToCompliancePackResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replicates compliance packages.
//
// @param request - CopyCompliancePacksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyCompliancePacksResponse
func (client *Client) CopyCompliancePacksWithContext(ctx context.Context, request *CopyCompliancePacksRequest, runtime *dara.RuntimeOptions) (_result *CopyCompliancePacksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replicates rules.
//
// @param request - CopyConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyConfigRulesResponse
func (client *Client) CopyConfigRulesWithContext(ctx context.Context, request *CopyConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *CopyConfigRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a downloadable resource file for the current Alibaba Cloud account.
//
// @param request - CreateAdvancedSearchFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAdvancedSearchFileResponse
func (client *Client) CreateAdvancedSearchFileWithContext(ctx context.Context, request *CreateAdvancedSearchFileRequest, runtime *dara.RuntimeOptions) (_result *CreateAdvancedSearchFileResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a downloadable resource file for an account group.
//
// Description:
//
// This topic provides an example on how to create a downloadable resource file for an account group whose ID is `ca-edd3626622af00b3****`. The resource file includes all the ECS instances in the account group.
//
// @param request - CreateAggregateAdvancedSearchFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAggregateAdvancedSearchFileResponse
func (client *Client) CreateAggregateAdvancedSearchFileWithContext(ctx context.Context, request *CreateAggregateAdvancedSearchFileRequest, runtime *dara.RuntimeOptions) (_result *CreateAggregateAdvancedSearchFileResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a compliance package for an account group.
//
// Description:
//
// This topic provides an example on how to create a compliance package for the account group `ca-f632626622af0079****` by using the compliance package template `ClassifiedProtectionPreCheck`.
//
// @param tmpReq - CreateAggregateCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAggregateCompliancePackResponse
func (client *Client) CreateAggregateCompliancePackWithContext(ctx context.Context, tmpReq *CreateAggregateCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *CreateAggregateCompliancePackResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a delivery channel for an account group.
//
// Description:
//
// In this example, a delivery channel is created for an account group. The ID of the account group is `ca-a4e5626622af0079****`. The type of the delivery channel is `OSS` and the Alibaba Cloud Resource Name (ARN) of the delivery destination is `acs:oss:cn-shanghai:100931896542****:new-bucket`. The result indicates that the delivery channel is created. The ID of the delivery channel is `cdc-8e45ff4e06a3a8****`.
//
// @param request - CreateAggregateConfigDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAggregateConfigDeliveryChannelResponse
func (client *Client) CreateAggregateConfigDeliveryChannelWithContext(ctx context.Context, request *CreateAggregateConfigDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *CreateAggregateConfigDeliveryChannelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a rule for an account group.
//
// Description:
//
// ### Limits
//
// You can create up to 200 rules for each management account.
//
// ### Usage notes
//
// This topic provides an example on how to create a rule based on the required-tags managed rule in the `ca-a4e5626622af0079****` account group. The returned result shows that the rule is created and its ID is `cr-4e3d626622af0080****`.
//
// @param tmpReq - CreateAggregateConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAggregateConfigRuleResponse
func (client *Client) CreateAggregateConfigRuleWithContext(ctx context.Context, tmpReq *CreateAggregateConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateAggregateConfigRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a remediation template for a rule in an account group.
//
// Description:
//
// This topic provides an example on how to create a remediation template for the rule whose ID is `cr-6b7c626622af00b4****` in the account group whose ID is `ca-6b4a626622af0012****`. The returned result shows that a remediation template is created and the ID of the remediation template is `crr-909ba2d4716700eb****`.
//
// @param request - CreateAggregateRemediationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAggregateRemediationResponse
func (client *Client) CreateAggregateRemediationWithContext(ctx context.Context, request *CreateAggregateRemediationRequest, runtime *dara.RuntimeOptions) (_result *CreateAggregateRemediationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an account group.
//
// Description:
//
// Each management account can create a maximum of five account groups. Each account group can contain a maximum of 200 member accounts.
//
// Cloud Config supports the following types of account groups:
//
//   - Global account group: The global account group contains all the member accounts that are added to the resource directory. A management account can create only one global account group.
//
//   - Custom account group: If you create a custom account group, you must manually add all or specific member accounts from the resource directory to the custom account group.
//
// This topic provides an example on how to create an account group of the `CUSTOM` type. The custom account group is named `Test_Group`, and its description is `Test account group`. The custom account group contains the following two member accounts:
//
//   - Member account ID: `171322098523****`. Member account name: `Alice`.
//
//   - Member account ID: `100532098349****`. Member account name: `Tom`.
//
// @param tmpReq - CreateAggregatorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAggregatorResponse
func (client *Client) CreateAggregatorWithContext(ctx context.Context, tmpReq *CreateAggregatorRequest, runtime *dara.RuntimeOptions) (_result *CreateAggregatorResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a compliance package for the current account.
//
// Description:
//
// Each ordinary account can create up to five compliance packages.
//
// This topic provides an example on how to create a compliance package named ClassifiedProtectionPreCheck. The compliance package contains a managed rule named `eip-bandwidth-limit`.
//
// @param tmpReq - CreateCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCompliancePackResponse
func (client *Client) CreateCompliancePackWithContext(ctx context.Context, tmpReq *CreateCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *CreateCompliancePackResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a delivery channel.
//
// Description:
//
// In this example, a delivery channel is created. The type of the delivery channel is `OSS` and the Alibaba Cloud Resource Name (ARN) of the delivery destination is `acs:oss:cn-shanghai:100931896542****:new-bucket`. The result indicates that the delivery channel is created, and the ID of the delivery channel is `cdc-8e45ff4e06a3a8****`.
//
// @param request - CreateConfigDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConfigDeliveryChannelResponse
func (client *Client) CreateConfigDeliveryChannelWithContext(ctx context.Context, request *CreateConfigDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *CreateConfigDeliveryChannelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a rule for the current account.
//
// Description:
//
// ## Limits
//
// You can use a common account to create up to 200 rules.
//
// @param tmpReq - CreateConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConfigRuleResponse
func (client *Client) CreateConfigRuleWithContext(ctx context.Context, tmpReq *CreateConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateConfigRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a remediation template for a rule.
//
// Description:
//
// This topic provides an example on how to create a remediation template for the rule `cr-8a973ac2e2be00a2****`. The returned result shows that a remediation template is created and the ID of the remediation template is `crr-909ba2d4716700eb****`.
//
// @param request - CreateRemediationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRemediationResponse
func (client *Client) CreateRemediationWithContext(ctx context.Context, request *CreateRemediationRequest, runtime *dara.RuntimeOptions) (_result *CreateRemediationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为当前UID创建合规报告模版
//
// @param tmpReq - CreateReportTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReportTemplateResponse
func (client *Client) CreateReportTemplateWithContext(ctx context.Context, tmpReq *CreateReportTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateReportTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables one or more rules in an account group. After a rule is disabled, the resource in the rule is no longer evaluated. The compliance evaluation results before the rule is disabled are still displayed.
//
// Description:
//
// ### [](#)Prerequisites
//
// The status of the rule is `ACTIVE`.
//
// ### [](#)Description
//
// This topic provides an example on how to disable the `cr-5772ba41209e007b****` rule in the `ca-04b3fd170e340007****` account group.
//
// @param request - DeactiveAggregateConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeactiveAggregateConfigRulesResponse
func (client *Client) DeactiveAggregateConfigRulesWithContext(ctx context.Context, request *DeactiveAggregateConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *DeactiveAggregateConfigRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a rule. After a rule is disabled, the resource in the rule is no longer evaluated. The compliance evaluation results before the rule is disabled are still displayed.
//
// Description:
//
// ### [](#)Prerequisites
//
// The status of the rule is `ACTIVE`.
//
// ### [](#)Description
//
// This topic provides an example on how to disable the `cr-19a56457e0d90058****` rule.
//
// @param request - DeactiveConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeactiveConfigRulesResponse
func (client *Client) DeactiveConfigRulesWithContext(ctx context.Context, request *DeactiveConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *DeactiveConfigRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the compliance packages of an account group.
//
// Description:
//
// This topic provides an example on how to delete the `cp-541e626622af0087****` compliance package from the `ca-04b3fd170e340007****` account group.
//
// @param request - DeleteAggregateCompliancePacksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAggregateCompliancePacksResponse
func (client *Client) DeleteAggregateCompliancePacksWithContext(ctx context.Context, request *DeleteAggregateCompliancePacksRequest, runtime *dara.RuntimeOptions) (_result *DeleteAggregateCompliancePacksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a delivery channel from an account group.
//
// Description:
//
// This topic provides an example on how to delete the `cdc-38c3013b46c9002c****` delivery channel from the `ca-23c6626622af0041****` account group. The returned result shows that the `cdc-38c3013b46c9002c****` delivery channel is deleted.
//
// @param request - DeleteAggregateConfigDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAggregateConfigDeliveryChannelResponse
func (client *Client) DeleteAggregateConfigDeliveryChannelWithContext(ctx context.Context, request *DeleteAggregateConfigDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteAggregateConfigDeliveryChannelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more rules from an account group. You can delete a rule in the Cloud Config console. After you delete the rule, the configurations of the rule are deleted.
//
// Description:
//
// This topic provides an example on how to delete the `cr-4e3d626622af0080****` rule from the `ca-a4e5626622af0079****` account group.
//
// @param request - DeleteAggregateConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAggregateConfigRulesResponse
func (client *Client) DeleteAggregateConfigRulesWithContext(ctx context.Context, request *DeleteAggregateConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteAggregateConfigRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more remediation templates from a rule in an account group.
//
// Description:
//
// This topic provides an example on how to delete the remediation template whose ID is `crr-909ba2d4716700eb****` from the account group whose ID is `ca-6b4a626622af0012****`. The returned result shows that the remediation template whose ID is `crr-909ba2d4716700eb****` is deleted.
//
// @param request - DeleteAggregateRemediationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAggregateRemediationsResponse
func (client *Client) DeleteAggregateRemediationsWithContext(ctx context.Context, request *DeleteAggregateRemediationsRequest, runtime *dara.RuntimeOptions) (_result *DeleteAggregateRemediationsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The management account or delegated administrator account of a resource directory can delete an account group.
//
// Description:
//
// ### [](#)Background information
//
// After you delete an account group, the following changes occur to Cloud Config:
//
//   - The rules and compliance packages of the account group are deleted and cannot be recovered.
//
//   - All compliance results generated in the account group are automatically deleted and cannot be recovered.
//
//   - Service-linked roles for Cloud Config of member accounts in the account group are retained.
//
//   - If the account groups to which a member belongs are all deleted, the member account uses Cloud Config as an independent Alibaba Cloud account.
//
// ### [](#)Description
//
// This topic provides an example on how to delete the account group whose ID is `ca-9190626622af00a9****`.
//
// @param request - DeleteAggregatorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAggregatorsResponse
func (client *Client) DeleteAggregatorsWithContext(ctx context.Context, request *DeleteAggregatorsRequest, runtime *dara.RuntimeOptions) (_result *DeleteAggregatorsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more compliance packages.
//
// Description:
//
// This topic provides an example on how to delete the `cp-541e626622af0087****` compliance package.
//
// @param request - DeleteCompliancePacksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCompliancePacksResponse
func (client *Client) DeleteCompliancePacksWithContext(ctx context.Context, request *DeleteCompliancePacksRequest, runtime *dara.RuntimeOptions) (_result *DeleteCompliancePacksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a delivery channel.
//
// Description:
//
// This topic provides an example on how to delete the `cdc-38c3013b46c9002c****` delivery channel. The returned result shows that the `cdc-38c3013b46c9002c****` delivery channel is deleted.
//
// @param request - DeleteConfigDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConfigDeliveryChannelResponse
func (client *Client) DeleteConfigDeliveryChannelWithContext(ctx context.Context, request *DeleteConfigDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteConfigDeliveryChannelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes rules.
//
// Description:
//
// In this example, the rule whose ID is cr-9908626622af0035\\*\\*\\*\\	- is deleted.
//
// @param request - DeleteConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConfigRulesResponse
func (client *Client) DeleteConfigRulesWithContext(ctx context.Context, request *DeleteConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteConfigRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more configured remediation templates that are associated with a rule.
//
// Description:
//
// This topic provides an example on how to delete the remediation template `crr-909ba2d4716700eb****`. The returned result shows that the remediation template whose ID is `crr-909ba2d4716700eb****` is deleted.
//
// @param request - DeleteRemediationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRemediationsResponse
func (client *Client) DeleteRemediationsWithContext(ctx context.Context, request *DeleteRemediationsRequest, runtime *dara.RuntimeOptions) (_result *DeleteRemediationsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除合规报告模版
//
// @param request - DeleteReportTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteReportTemplateResponse
func (client *Client) DeleteReportTemplateWithContext(ctx context.Context, request *DeleteReportTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteReportTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取资源详情
//
// @param request - DescribeDiscoveredResourceBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiscoveredResourceBatchResponse
func (client *Client) DescribeDiscoveredResourceBatchWithContext(ctx context.Context, request *DescribeDiscoveredResourceBatchRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiscoveredResourceBatchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param request - DescribeIntegratedServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIntegratedServiceStatusResponse
func (client *Client) DescribeIntegratedServiceStatusWithContext(ctx context.Context, request *DescribeIntegratedServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeIntegratedServiceStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic provides an example on how to query the details of a remediation configuration whose ID is crr-f381cf0c1c2f004e\\*\\*\\*\\*.
//
// @param request - DescribeRemediationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRemediationResponse
func (client *Client) DescribeRemediationWithContext(ctx context.Context, request *DescribeRemediationRequest, runtime *dara.RuntimeOptions) (_result *DescribeRemediationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes one or more rules in an account group from a compliance package.
//
// Description:
//
// ### Prerequisites
//
// One or more rules are added to a compliance package.
//
// ### Usage notes
//
// The sample request in this topic shows you how to remove the `cr-6cc4626622af00e7****` rule in the `ca-75b4626622af00c3****` account group from the `cp-5bb1626622af00bd****` compliance package.
//
// @param request - DetachAggregateConfigRuleToCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachAggregateConfigRuleToCompliancePackResponse
func (client *Client) DetachAggregateConfigRuleToCompliancePackWithContext(ctx context.Context, request *DetachAggregateConfigRuleToCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *DetachAggregateConfigRuleToCompliancePackResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes one or more rules from a compliance package.
//
// Description:
//
// ### Prerequisites
//
// One or more rules are added to a compliance package.
//
// ### Usage notes
//
// This topic provides an example on how to remove the `cr-6cc4626622af00e7****` rule from the `cp-5bb1626622af00bd****` compliance package.
//
// @param request - DetachConfigRuleToCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachConfigRuleToCompliancePackResponse
func (client *Client) DetachConfigRuleToCompliancePackWithContext(ctx context.Context, request *DetachConfigRuleToCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *DetachConfigRuleToCompliancePackResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 试运行事前合规预检的规则
//
// @param request - DryRunConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DryRunConfigRuleResponse
func (client *Client) DryRunConfigRuleWithContext(ctx context.Context, request *DryRunConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *DryRunConfigRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes evaluation rules to evaluate resources.
//
// @param tmpReq - EvaluatePreConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EvaluatePreConfigRulesResponse
func (client *Client) EvaluatePreConfigRulesWithContext(ctx context.Context, tmpReq *EvaluatePreConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *EvaluatePreConfigRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a compliance evaluation report based on a compliance package in an account group.
//
// Description:
//
// > You can call this operation to generate the latest compliance evaluation report. To download the report, call the GetAggregateConfigRulesReport operation. For more information, see [GetAggregateCompliancePackReport](https://help.aliyun.com/document_detail/262699.html).
//
// This topic provides an example on how to generate a compliance evaluation report based on the `cp-fdc8626622af00f9****` compliance package in the `ca-f632626622af0079****` account group.
//
// @param request - GenerateAggregateCompliancePackReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateAggregateCompliancePackReportResponse
func (client *Client) GenerateAggregateCompliancePackReportWithContext(ctx context.Context, request *GenerateAggregateCompliancePackReportRequest, runtime *dara.RuntimeOptions) (_result *GenerateAggregateCompliancePackReportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a compliance evaluation report for the rules in a specified account group.
//
// Description:
//
// > You can call this operation to generate the latest compliance evaluation report. To download the report, call the GetAggregateConfigRulesReport operation. For more information, see [GetAggregateConfigRulesReport](https://help.aliyun.com/document_detail/262706.html).
//
// The topic provides an example on how to generate a compliance evaluation report based on all rules in the `ca-f632626622af0079****` account group.
//
// @param request - GenerateAggregateConfigRulesReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateAggregateConfigRulesReportResponse
func (client *Client) GenerateAggregateConfigRulesReportWithContext(ctx context.Context, request *GenerateAggregateConfigRulesReportRequest, runtime *dara.RuntimeOptions) (_result *GenerateAggregateConfigRulesReportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a downloadable inventory for global resources in an account group.
//
// Description:
//
// This topic provides an example to show how to generate a downloadable inventory for global resources in the account group ca-a91d626622af0035\\*\\*\\*\\*.
//
// @param request - GenerateAggregateResourceInventoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateAggregateResourceInventoryResponse
func (client *Client) GenerateAggregateResourceInventoryWithContext(ctx context.Context, request *GenerateAggregateResourceInventoryRequest, runtime *dara.RuntimeOptions) (_result *GenerateAggregateResourceInventoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a compliance evaluation report based on a compliance package.
//
// Description:
//
// > You can call this operation to generate the latest compliance evaluation report. To download the report, call the GetCompliancePackReport operation. For more information, see [GetCompliancePackReport](https://help.aliyun.com/document_detail/263347.html).
//
// This topic provides an example on how to generate a compliance evaluation report based on the `cp-a8a8626622af0082****` compliance package.
//
// @param request - GenerateCompliancePackReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCompliancePackReportResponse
func (client *Client) GenerateCompliancePackReportWithContext(ctx context.Context, request *GenerateCompliancePackReportRequest, runtime *dara.RuntimeOptions) (_result *GenerateCompliancePackReportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a compliance evaluation report for a rule.
//
// Description:
//
// >  You can call this operation to generate the latest compliance evaluation report. To download the report, call the GetConfigRulesReport operation. For more information, see [GetConfigRulesReport](https://help.aliyun.com/document_detail/263608.html).
//
// This topic provides an example of how to generate a compliance evaluation report based on all existing rules.
//
// @param request - GenerateConfigRulesReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateConfigRulesReportResponse
func (client *Client) GenerateConfigRulesReportWithContext(ctx context.Context, request *GenerateConfigRulesReportRequest, runtime *dara.RuntimeOptions) (_result *GenerateConfigRulesReportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 基于报告模版生成报告Id
//
// @param request - GenerateReportFromTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateReportFromTemplateResponse
func (client *Client) GenerateReportFromTemplateWithContext(ctx context.Context, request *GenerateReportFromTemplateRequest, runtime *dara.RuntimeOptions) (_result *GenerateReportFromTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a resource inventory for global resources.
//
// Description:
//
// This topic provides an example on how to generate a resource inventory for global resources of the current account.
//
// @param request - GenerateResourceInventoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateResourceInventoryResponse
func (client *Client) GenerateResourceInventoryWithContext(ctx context.Context, request *GenerateResourceInventoryRequest, runtime *dara.RuntimeOptions) (_result *GenerateResourceInventoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance evaluation results of member accounts for which a compliance package takes effect in an account group.
//
// Description:
//
// This topic provides an example on how to query the compliance evaluation results of member accounts for which the `cp-541e626622af0087****` compliance package takes effect in the `ca-04b3fd170e340007****` account group. The returned result shows that two member accounts are monitored by the compliance package and they are both evaluated as compliant.
//
// @param request - GetAggregateAccountComplianceByPackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateAccountComplianceByPackResponse
func (client *Client) GetAggregateAccountComplianceByPackWithContext(ctx context.Context, request *GetAggregateAccountComplianceByPackRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateAccountComplianceByPackResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the most recently generated resource file of an account group.
//
// @param request - GetAggregateAdvancedSearchFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateAdvancedSearchFileResponse
func (client *Client) GetAggregateAdvancedSearchFileWithContext(ctx context.Context, request *GetAggregateAdvancedSearchFileRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateAdvancedSearchFileResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a compliance package in an account group.
//
// Description:
//
// The topic provides an example on how to query the details of a compliance package whose ID is `cp-fdc8626622af00f9****` in an account group whose ID is `ca-f632626622af0079****`.
//
// @param tmpReq - GetAggregateCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateCompliancePackResponse
func (client *Client) GetAggregateCompliancePackWithContext(ctx context.Context, tmpReq *GetAggregateCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateCompliancePackResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance evaluation report that is generated based on a compliance package of an account group.
//
// Description:
//
// > Before you call this operation, you must call the GenerateAggregateCompliancePackReport operation to generate the latest compliance evaluation report based on a compliance package. For more information, see [GenerateAggregateCompliancePackReport](https://help.aliyun.com/document_detail/262687.html).
//
// This topic provides an example on how to query the compliance evaluation report that is generated based on the `cp-fdc8626622af00f9****` compliance package in the `ca-f632626622af0079****` account group.
//
// @param request - GetAggregateCompliancePackReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateCompliancePackReportResponse
func (client *Client) GetAggregateCompliancePackReportWithContext(ctx context.Context, request *GetAggregateCompliancePackReportRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateCompliancePackReportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance statistics of an account group.
//
// Description:
//
// This topic provides an example on how to query the compliance statistics of resources and rules in the account group ca-a91d626622af0035\\*\\*\\*\\*.
//
// @param request - GetAggregateComplianceSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateComplianceSummaryResponse
func (client *Client) GetAggregateComplianceSummaryWithContext(ctx context.Context, request *GetAggregateComplianceSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateComplianceSummaryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a delivery channel in an account group.
//
// @param request - GetAggregateConfigDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateConfigDeliveryChannelResponse
func (client *Client) GetAggregateConfigDeliveryChannelWithContext(ctx context.Context, request *GetAggregateConfigDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateConfigDeliveryChannelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a rule in an account group.
//
// Description:
//
// This example shows how to query the details of the `cr-7f7d626622af0041****` rule in the `ca-7f00626622af0041****` account group.
//
// @param tmpReq - GetAggregateConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateConfigRuleResponse
func (client *Client) GetAggregateConfigRuleWithContext(ctx context.Context, tmpReq *GetAggregateConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateConfigRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries compliance evaluation results based on the rules in a compliance package in an account group.
//
// Description:
//
// The sample request in this topic shows you how to query the compliance evaluation results based on rules in the `cp-541e626622af0087****` compliance package that is created for the `ca-04b3fd170e340007****` account group. The return result shows a total of `one` rule. `No resources` are evaluated as non-compliant based on the rule.
//
// @param request - GetAggregateConfigRuleComplianceByPackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateConfigRuleComplianceByPackResponse
func (client *Client) GetAggregateConfigRuleComplianceByPackWithContext(ctx context.Context, request *GetAggregateConfigRuleComplianceByPackRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateConfigRuleComplianceByPackResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the summary of compliance evaluation results by rule risk level in an account group.
//
// Description:
//
// This topic provides an example on how to query the summary of compliance evaluation results by rule risk level in the `ca-3a58626622af0005****` account group. The returned result shows four rules that are specified with the high risk level. One of the rules detects non-compliant resources, and the resources evaluated by the remaining three are compliant.
//
// @param request - GetAggregateConfigRuleSummaryByRiskLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateConfigRuleSummaryByRiskLevelResponse
func (client *Client) GetAggregateConfigRuleSummaryByRiskLevelWithContext(ctx context.Context, request *GetAggregateConfigRuleSummaryByRiskLevelRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateConfigRuleSummaryByRiskLevelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Downloads the compliance evaluation report in the Excel format to your on-premises machine. This allows you to assign tasks and modify incompliant resource configurations.
//
// Description:
//
// > Before you call this operation, you must call the GenerateAggregateConfigRulesReport operation to generate the latest compliance evaluation report based on all rules in an account group. For more information, see [GenerateAggregateConfigRulesReport](https://help.aliyun.com/document_detail/262701.html).
//
// This topic provides an example on how to query the compliance evaluation report that is generated based on all rules in the `ca-f632626622af0079****` account group.
//
// @param request - GetAggregateConfigRulesReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateConfigRulesReportResponse
func (client *Client) GetAggregateConfigRulesReportWithContext(ctx context.Context, request *GetAggregateConfigRulesReportRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateConfigRulesReportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// This topic provides an example on how to query the details of an Elastic Compute Service (ECS) instance `i-bp12g4xbl4i0brkn****` that resides in the China (Hangzhou) region in the account group `ca-5885626622af0008****`.
//
// @param request - GetAggregateDiscoveredResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateDiscoveredResourceResponse
func (client *Client) GetAggregateDiscoveredResourceWithContext(ctx context.Context, request *GetAggregateDiscoveredResourceRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateDiscoveredResourceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries compliance evaluation results based on the rules in a compliance package in an account group.
//
// Description:
//
// This topic provides an example on how to query the compliance evaluation results based on the `cr-d369626622af008e****` rule in the `ca-a4e5626622af0079****` account group. The returned result shows that a total of 10 resources are evaluated by the rule and five of them are evaluated as compliant.
//
// @param request - GetAggregateResourceComplianceByConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceComplianceByConfigRuleResponse
func (client *Client) GetAggregateResourceComplianceByConfigRuleWithContext(ctx context.Context, request *GetAggregateResourceComplianceByConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceComplianceByConfigRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance evaluation results of resources evaluated based on a compliance package of an account group.
//
// Description:
//
// This topic provides an example on how to query the compliance evaluation results of resources monitored based on the `cp-fdc8626622af00f9****` compliance package in the `ca-f632626622af0079****`account group. The returned result shows that the total number of monitored resources is `10` and the number of non-compliant resources is `7`.
//
// @param request - GetAggregateResourceComplianceByPackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceComplianceByPackResponse
func (client *Client) GetAggregateResourceComplianceByPackWithContext(ctx context.Context, request *GetAggregateResourceComplianceByPackRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceComplianceByPackResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the evaluation results grouped by resource type for an account group rule.
//
// @param request - GetAggregateResourceComplianceGroupByRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceComplianceGroupByRegionResponse
func (client *Client) GetAggregateResourceComplianceGroupByRegionWithContext(ctx context.Context, request *GetAggregateResourceComplianceGroupByRegionRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceComplianceGroupByRegionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the evaluation results grouped by resource type for an account group rule.
//
// @param request - GetAggregateResourceComplianceGroupByResourceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceComplianceGroupByResourceTypeResponse
func (client *Client) GetAggregateResourceComplianceGroupByResourceTypeWithContext(ctx context.Context, request *GetAggregateResourceComplianceGroupByResourceTypeRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceComplianceGroupByResourceTypeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance timeline of a resource in an account group.
//
// Description:
//
// The sample request in this topic shows you how to query the compliance timeline of the `new-bucket` resource that resides in the `cn-hangzhou` region within the `100931896542****` member account of the `ca-5885626622af0008****` account group. The new-bucket resource is an Object Storage Service (OSS) bucket. The return result shows the following two timestamps on the compliance timeline: `1625200295276` and `1625200228510`. The first timestamp indicates 12:31:35 on July 2, 2021 (UTC+8), and the second timestamp indicates 12:30:28 on July 2, 2021 (UTC+8).
//
// @param request - GetAggregateResourceComplianceTimelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceComplianceTimelineResponse
func (client *Client) GetAggregateResourceComplianceTimelineWithContext(ctx context.Context, request *GetAggregateResourceComplianceTimelineRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceComplianceTimelineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration timeline of a resource in an account group.
//
// Description:
//
// The sample request in this topic shows you how to query the configuration timeline of the `new-bucket` resource that resides in the `cn-hangzhou` region within the `100931896542****` member account of the `ca-5885626622af0008****` account group. The new-bucket resource is an Object Storage Service (OSS) bucket. The return result shows that the timestamp when the resource configuration changes is `1624961112000`. The timestamp indicates 18:05:12 on June 29, 2021 (UTC+8).
//
// @param request - GetAggregateResourceConfigurationTimelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceConfigurationTimelineResponse
func (client *Client) GetAggregateResourceConfigurationTimelineWithContext(ctx context.Context, request *GetAggregateResourceConfigurationTimelineRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceConfigurationTimelineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on the resources in an account group by region.
//
// Description:
//
// This topic provides an example on how to query the statistics on the resources in an account group named `ca-a260626622af0005****` by region. The returned result shows that a total of `10` resources exist in the `cn-hangzhou` region.
//
// @param request - GetAggregateResourceCountsGroupByRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceCountsGroupByRegionResponse
func (client *Client) GetAggregateResourceCountsGroupByRegionWithContext(ctx context.Context, request *GetAggregateResourceCountsGroupByRegionRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceCountsGroupByRegionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on the resources in an account group by resource type.
//
// Description:
//
// This topic provides an example on how to query the statistics on the resources in an account group whose ID is `ca-a260626622af0005****` by resource type. The returned result shows that the account group has a total of `seven` resources of the `ACS::RAM::Role` resource type.
//
// @param request - GetAggregateResourceCountsGroupByResourceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceCountsGroupByResourceTypeResponse
func (client *Client) GetAggregateResourceCountsGroupByResourceTypeWithContext(ctx context.Context, request *GetAggregateResourceCountsGroupByResourceTypeRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceCountsGroupByResourceTypeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the last resource inventory that is generated on the Global Resources page within the current account group.
//
// Description:
//
// ### [](#)Prerequisites
//
// The [GenerateAggregateResourceInventory](https://help.aliyun.com/document_detail/2398353.html) operation is called to generate a resource inventory. Then, this operation is called to obtain the URL of the resource inventory.
//
// ### [](#)Description
//
// This topic provides an example on how to obtain the last resource inventory that is generated within the account group ca-a91d626622af0035\\*\\*\\*\\*.
//
// @param request - GetAggregateResourceInventoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregateResourceInventoryResponse
func (client *Client) GetAggregateResourceInventoryWithContext(ctx context.Context, request *GetAggregateResourceInventoryRequest, runtime *dara.RuntimeOptions) (_result *GetAggregateResourceInventoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an account group. You can query the name, creation time, member, and type of an account group.
//
// Description:
//
// The sample request in this topic shows you how to query the details of the `ca-88ea626622af0055****` account group. The return result shows that the account group is named `Test_Group`, its description is `Test account group`, and it is of the `CUSTOM` type. The account group is in the `1` state, which indicates that it is created.
//
// @param tmpReq - GetAggregatorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggregatorResponse
func (client *Client) GetAggregatorWithContext(ctx context.Context, tmpReq *GetAggregatorRequest, runtime *dara.RuntimeOptions) (_result *GetAggregatorResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a compliance package.
//
// Description:
//
// This topic provides an example on how to query the details of a compliance package whose ID is `cp-fdc8626622af00f9****`. The returned result shows that the name of the compliance package is `ClassifiedProtectionPreCheck`, the compliance package is in the `ACTIVE` state, and the risk level of the rules in the compliance package is `1`, which indicates high risk level.
//
// @param tmpReq - GetCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCompliancePackResponse
func (client *Client) GetCompliancePackWithContext(ctx context.Context, tmpReq *GetCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *GetCompliancePackResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance evaluation report that is generated based on a compliance package.
//
// Description:
//
// > Before you call this operation, you must call the GenerateCompliancePackReport operation to generate the latest compliance evaluation report based on a compliance package. For more information, see [GenerateCompliancePackReport](https://help.aliyun.com/document_detail/263525.html).
//
// This topic provides an example on how to query the compliance evaluation report that is generated based on the `cp-fdc8626622af00f9****` compliance package.
//
// @param request - GetCompliancePackReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCompliancePackReportResponse
func (client *Client) GetCompliancePackReportWithContext(ctx context.Context, request *GetCompliancePackReportRequest, runtime *dara.RuntimeOptions) (_result *GetCompliancePackReportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a delivery channel.
//
// @param request - GetConfigDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfigDeliveryChannelResponse
func (client *Client) GetConfigDeliveryChannelWithContext(ctx context.Context, request *GetConfigDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *GetConfigDeliveryChannelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a rule.
//
// Description:
//
// This topic provides an example on how to query the details of the `cr-7f7d626622af0041****` rule.
//
// @param tmpReq - GetConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfigRuleResponse
func (client *Client) GetConfigRuleWithContext(ctx context.Context, tmpReq *GetConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *GetConfigRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries compliance evaluation results based on the rules in a compliance package.
//
// Description:
//
// In this topic, the `cp-541e626622af0087****` compliance package is used as an example. The return result shows a total of one rule against which specific resources are evaluated as compliant.
//
// @param request - GetConfigRuleComplianceByPackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfigRuleComplianceByPackResponse
func (client *Client) GetConfigRuleComplianceByPackWithContext(ctx context.Context, request *GetConfigRuleComplianceByPackRequest, runtime *dara.RuntimeOptions) (_result *GetConfigRuleComplianceByPackResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Downloads the compliance evaluation report in the Excel format to your on-premises machine. This allows you to assign tasks and modify incompliant resource configurations.
//
// Description:
//
// >  Before you call this operation, you must call the GenerateConfigRulesReport operation to generate the latest compliance evaluation report based on all existing rules. For more information, see [GenerateConfigRulesReport](https://help.aliyun.com/document_detail/263601.html).
//
// This topic provides an example of how to query the compliance evaluation report that is generated based on all existing rules.
//
// @param request - GetConfigRulesReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfigRulesReportResponse
func (client *Client) GetConfigRulesReportWithContext(ctx context.Context, request *GetConfigRulesReportRequest, runtime *dara.RuntimeOptions) (_result *GetConfigRulesReportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// This topic provides an example on how to query the details of the Elastic Compute Service (ECS) instance `i-bp12g4xbl4i0brkn****` that resides in the China (Hangzhou) region.
//
// @param request - GetDiscoveredResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDiscoveredResourceResponse
func (client *Client) GetDiscoveredResourceWithContext(ctx context.Context, request *GetDiscoveredResourceRequest, runtime *dara.RuntimeOptions) (_result *GetDiscoveredResourceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on resources by region.
//
// Description:
//
// This topic provides an example to demonstrate how to query the statistics on resources by region. The returned result shows that a total of 10 resources exist in the `cn-hangzhou` region.
//
// @param request - GetDiscoveredResourceCountsGroupByRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDiscoveredResourceCountsGroupByRegionResponse
func (client *Client) GetDiscoveredResourceCountsGroupByRegionWithContext(ctx context.Context, request *GetDiscoveredResourceCountsGroupByRegionRequest, runtime *dara.RuntimeOptions) (_result *GetDiscoveredResourceCountsGroupByRegionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on resources by resource type.
//
// Description:
//
// This topic describes how to query the statistics on resources by resource type. The returned result shows that a total of 10 resources of the `ACS::ECS::Instance` resource type exist.
//
// @param request - GetDiscoveredResourceCountsGroupByResourceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDiscoveredResourceCountsGroupByResourceTypeResponse
func (client *Client) GetDiscoveredResourceCountsGroupByResourceTypeWithContext(ctx context.Context, request *GetDiscoveredResourceCountsGroupByResourceTypeRequest, runtime *dara.RuntimeOptions) (_result *GetDiscoveredResourceCountsGroupByResourceTypeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the integration status of a specific cloud service.
//
// @param request - GetIntegratedServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIntegratedServiceStatusResponse
func (client *Client) GetIntegratedServiceStatusWithContext(ctx context.Context, request *GetIntegratedServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *GetIntegratedServiceStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specific managed rule.
//
// Description:
//
// This topic provides an example on how to query the details of the managed rule `cdn-domain-https-enabled`.
//
// @param request - GetManagedRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetManagedRuleResponse
func (client *Client) GetManagedRuleWithContext(ctx context.Context, request *GetManagedRuleRequest, runtime *dara.RuntimeOptions) (_result *GetManagedRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// This topic provides an example on how to query the details of the automatic remediation template ACS-ALB-BulkyEnableDeletionProtection.
//
// @param request - GetRemediationTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRemediationTemplateResponse
func (client *Client) GetRemediationTemplateWithContext(ctx context.Context, request *GetRemediationTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetRemediationTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取基于报告模版生成的报告
//
// @param request - GetReportFromTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReportFromTemplateResponse
func (client *Client) GetReportFromTemplateWithContext(ctx context.Context, request *GetReportFromTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetReportFromTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取合规报告模版详情
//
// @param request - GetReportTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReportTemplateResponse
func (client *Client) GetReportTemplateWithContext(ctx context.Context, request *GetReportTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetReportTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance summary based on the compliance evaluation result of a rule.
//
// Description:
//
// In this topic, the `cr-d369626622af008e****` rule is used as an example. The return result shows that a total of 10 resources are evaluated by the rule and `five` of them are evaluated as compliant.
//
// @param request - GetResourceComplianceByConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceComplianceByConfigRuleResponse
func (client *Client) GetResourceComplianceByConfigRuleWithContext(ctx context.Context, request *GetResourceComplianceByConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *GetResourceComplianceByConfigRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance evaluation results of resources evaluated based on a compliance package.
//
// Description:
//
// This topic provides an example on how to query the compliance evaluation results of resources monitored by using the `cp-541e626622af0087****` compliance package. The returned result shows a total of 10 resources and seven of them are evaluated as non-compliant.
//
// @param request - GetResourceComplianceByPackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceComplianceByPackResponse
func (client *Client) GetResourceComplianceByPackWithContext(ctx context.Context, request *GetResourceComplianceByPackRequest, runtime *dara.RuntimeOptions) (_result *GetResourceComplianceByPackResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the evaluation results grouped by region for a rule.
//
// @param request - GetResourceComplianceGroupByRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceComplianceGroupByRegionResponse
func (client *Client) GetResourceComplianceGroupByRegionWithContext(ctx context.Context, request *GetResourceComplianceGroupByRegionRequest, runtime *dara.RuntimeOptions) (_result *GetResourceComplianceGroupByRegionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the evaluation results grouped by resource type for a rule.
//
// @param request - GetResourceComplianceGroupByResourceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceComplianceGroupByResourceTypeResponse
func (client *Client) GetResourceComplianceGroupByResourceTypeWithContext(ctx context.Context, request *GetResourceComplianceGroupByResourceTypeRequest, runtime *dara.RuntimeOptions) (_result *GetResourceComplianceGroupByResourceTypeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance timeline of a resource. The compliance timeline of a resource indicates the compliance evaluation record of the resource. A compliance timeline includes points and the content on the compliance timeline.
//
// Description:
//
// In Cloud Config, each resource has a compliance timeline. Cloud Config generates a compliance evaluation record for a resource each time the resource is evaluated based on a rule. The compliance evaluation records of a resource are displayed in a compliance timeline. You can configure Cloud Config to execute a rule to evaluate a resource on a regular basis or each time you change the resource configuration. You can also manually execute a rule to evaluate a resource.
//
// This topic provides an example on how to query the compliance timeline of the `new-bucket` resource that resides in the `cn-hangzhou` region. The resource is an Object Storage Service (OSS) bucket. The returned result shows the following two timestamps on the compliance timeline: `1625200295276` and `1625200228510`. The first timestamp indicates 12:31:35 on July 2, 2021 (UTC+8) and the second timestamp indicates 12:30:28 on July 2, 2021 (UTC+8).
//
// @param request - GetResourceComplianceTimelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceComplianceTimelineResponse
func (client *Client) GetResourceComplianceTimelineWithContext(ctx context.Context, request *GetResourceComplianceTimelineRequest, runtime *dara.RuntimeOptions) (_result *GetResourceComplianceTimelineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源配置样例
//
// @param request - GetResourceConfigurationSampleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceConfigurationSampleResponse
func (client *Client) GetResourceConfigurationSampleWithContext(ctx context.Context, request *GetResourceConfigurationSampleRequest, runtime *dara.RuntimeOptions) (_result *GetResourceConfigurationSampleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration timeline of a resource.
//
// Description:
//
// The sample request in this topic shows you how to query the configuration timeline of the `new-bucket` resource that resides in the `cn-hangzhou` region. The new-bucket resource is an Object Storage Service (OSS) bucket. The return result shows that the timestamp when the resource configuration changes is `1624961112000`. The timestamp indicates 18:05:12 on June 29, 2021 (UTC+8).
//
// @param request - GetResourceConfigurationTimelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceConfigurationTimelineResponse
func (client *Client) GetResourceConfigurationTimelineWithContext(ctx context.Context, request *GetResourceConfigurationTimelineRequest, runtime *dara.RuntimeOptions) (_result *GetResourceConfigurationTimelineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源类型的属性配置
//
// @param request - GetResourceTypePropertiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceTypePropertiesResponse
func (client *Client) GetResourceTypePropertiesWithContext(ctx context.Context, request *GetResourceTypePropertiesRequest, runtime *dara.RuntimeOptions) (_result *GetResourceTypePropertiesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource relationships supported by a resource type.
//
// Description:
//
// This topic provides an example to show how to query the resource relationships that are supported by the ACS::ECS::Instance resource type.
//
// @param request - GetSupportedResourceRelationConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupportedResourceRelationConfigResponse
func (client *Client) GetSupportedResourceRelationConfigWithContext(ctx context.Context, request *GetSupportedResourceRelationConfigRequest, runtime *dara.RuntimeOptions) (_result *GetSupportedResourceRelationConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Ignores the evaluation results of some resources in an account group based on a specific rule. You can also specify a time period for ignoring the evaluation results. After the period elapses, the evaluation results of the resources based on the rule are automatically displayed.
//
// Description:
//
// After you ignore a resource that is evaluated as incompliant by using a rule, the resource is still evaluated by using the rule, but the compliance result is Ignored.
//
// This example shows how to ignore the `lb-hp3a3b4ztyfm2plgm****` incompliant resource that is evaluated by using the `cr-7e72626622af0051***` rule in the `120886317861****` member account of the `ca-5b6c626622af008f****` account group. The ID of the region where the resource resides is `cn-beijing`, and the type of the resource is `ACS::SLB::LoadBalancer`.
//
// @param tmpReq - IgnoreAggregateEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IgnoreAggregateEvaluationResultsResponse
func (client *Client) IgnoreAggregateEvaluationResultsWithContext(ctx context.Context, tmpReq *IgnoreAggregateEvaluationResultsRequest, runtime *dara.RuntimeOptions) (_result *IgnoreAggregateEvaluationResultsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Ignores the evaluation results of some resources based on a specific rule. You can also specify a time period for ignoring the evaluation results. After the period elapses, the evaluation results of the resources based on the rule are automatically displayed.
//
// Description:
//
// After you ignore a resource that is evaluated as incompliant by using a rule, the resource is still evaluated by using the rule, but the compliance result is Ignored.
//
// This example shows how to ignore the `lb-hp3a3b4ztyfm2plgm****` resource that is evaluated as incompliant by using the `cr-7e72626622af0051****` rule in the `100931896542****` account. The ID of the region in which the resource resides is `cn-beijing`, and the type of the resource is `ACS::SLB::LoadBalancer`.
//
// @param tmpReq - IgnoreEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IgnoreEvaluationResultsResponse
func (client *Client) IgnoreEvaluationResultsWithContext(ctx context.Context, tmpReq *IgnoreEvaluationResultsRequest, runtime *dara.RuntimeOptions) (_result *IgnoreEvaluationResultsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of compliance packages in an account group.
//
// Description:
//
// In this topic, the `ca-f632626622af0079****` account group is used as an example. The return result shows one compliance package whose ID is `cp-fdc8626622af00f9****`.
//
// @param tmpReq - ListAggregateCompliancePacksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateCompliancePacksResponse
func (client *Client) ListAggregateCompliancePacksWithContext(ctx context.Context, tmpReq *ListAggregateCompliancePacksRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateCompliancePacksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all delivery channels in an account group.
//
// @param request - ListAggregateConfigDeliveryChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateConfigDeliveryChannelsResponse
func (client *Client) ListAggregateConfigDeliveryChannelsWithContext(ctx context.Context, request *ListAggregateConfigDeliveryChannelsRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateConfigDeliveryChannelsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance evaluation results of resources based on a rule in an account group.
//
// Description:
//
// This topic provides an example on how to query the compliance evaluation results of resources based on the `cr-888f626622af00ae****` rule in the `ca-d1e3326622af00cb****` account group. The returned result indicates that the `Bucket-test` resource is evaluated as `NON_COMPLIANT` by using the rule. The resource is an Object Storage Service (OSS) bucket.
//
// @param request - ListAggregateConfigRuleEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateConfigRuleEvaluationResultsResponse
func (client *Client) ListAggregateConfigRuleEvaluationResultsWithContext(ctx context.Context, request *ListAggregateConfigRuleEvaluationResultsRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateConfigRuleEvaluationResultsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics of compliance evaluation results of an account group.
//
// Description:
//
// This topic provides an example on how to query the statistics of compliance evaluation results of an account group whose ID is ca-edd3626622af00b3\\*\\*\\*\\*.
//
// @param request - ListAggregateConfigRuleEvaluationStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateConfigRuleEvaluationStatisticsResponse
func (client *Client) ListAggregateConfigRuleEvaluationStatisticsWithContext(ctx context.Context, request *ListAggregateConfigRuleEvaluationStatisticsRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateConfigRuleEvaluationStatisticsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of rules in an account group.
//
// Description:
//
// This topic provides an example on how to query the rules in an account group whose ID is `ca-f632626622af0079****`. The returned result shows a total of one rule and two evaluated resources. The resources are both evaluated as `COMPLIANT`.
//
// @param tmpReq - ListAggregateConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateConfigRulesResponse
func (client *Client) ListAggregateConfigRulesWithContext(ctx context.Context, tmpReq *ListAggregateConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateConfigRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of resources aggregated across regions within all member accounts of a specific account group.
//
// Description:
//
// This topic provides an example on how to query the resources within the member account `100931896542****` of the account group `ca-c560626622af0005****`. The result indicates that eight resources are queried.
//
// @param request - ListAggregateDiscoveredResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateDiscoveredResourcesResponse
func (client *Client) ListAggregateDiscoveredResourcesWithContext(ctx context.Context, request *ListAggregateDiscoveredResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateDiscoveredResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取推荐的托管规则列表-多账号
//
// @param request - ListAggregateRecommendManagedRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateRecommendManagedRulesResponse
func (client *Client) ListAggregateRecommendManagedRulesWithContext(ctx context.Context, request *ListAggregateRecommendManagedRulesRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateRecommendManagedRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the remediation records of a rule in an account group.
//
// Description:
//
// This topic provides an example on how to query the remediation records of the `cr-d04a626622af00af****` rule in the `ca-edd3626622af00b3****` account group.
//
// @param request - ListAggregateRemediationExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateRemediationExecutionsResponse
func (client *Client) ListAggregateRemediationExecutionsWithContext(ctx context.Context, request *ListAggregateRemediationExecutionsRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateRemediationExecutionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of remediation templates for a rule in an account group.
//
// Description:
//
// This topic provides an example on how to query the remediation templates of the rule whose ID is `cr-6b7c626622af00b4****` in the account group whose ID is `ca-6b4a626622af0012****`.
//
// @param request - ListAggregateRemediationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateRemediationsResponse
func (client *Client) ListAggregateRemediationsWithContext(ctx context.Context, request *ListAggregateRemediationsRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateRemediationsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance evaluation results of resources in an account group.
//
// Description:
//
// This example shows how to query the compliance evaluation result of the `23642660635396****` resource in the `ca-7f00626622af0041****` account group. The resource is a RAM user. The returned result indicates that the resource is evaluated as `NON_COMPLIANT` by using the `cr-7f7d626622af0041****` rule.
//
// @param request - ListAggregateResourceEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateResourceEvaluationResultsResponse
func (client *Client) ListAggregateResourceEvaluationResultsWithContext(ctx context.Context, request *ListAggregateResourceEvaluationResultsRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateResourceEvaluationResultsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of the resources of a specific resource in an account group.
//
// Description:
//
// This topic provides an example on how to query the disks that are associated with an Elastic Compute Service (ECS) instance in an account group.
//
// @param request - ListAggregateResourceRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateResourceRelationsResponse
func (client *Client) ListAggregateResourceRelationsWithContext(ctx context.Context, request *ListAggregateResourceRelationsRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateResourceRelationsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains resources in a specific account group based on the fields in the resource properties by using a SELECT statement.
//
// Description:
//
// When you write a `SELECT` statement, you must obtain the fields and the data types of the fields from the property file of the resource type. For more information about property files, see[ Alibaba Cloud Config Resource Schema](https://github.com/aliyun/alibabacloud-config-resource-schema)
//
// >
//
//   - Each resource type supported by Cloud Config has a property file. Property files are named based on the related resource types. For example, the property file of the `ACS::ECS::Instance` resource type is named `ACS_ECS_Instance.properties.json`. Property files of different resource types are placed under the `config/properties/resource-types` path.
//
//   - For more information about the examples and limits on SQL query statements, see [Examples of SQL query statements](https://help.aliyun.com/document_detail/398718.html) and [Limits on SQL query statements](https://help.aliyun.com/document_detail/398750.html).
//
// This topic provides an example on how to obtain all resources whose tag key is `business` and whose tag value is `online` in the account group `ca-4b05626622af000c****` by using the advanced search feature.
//
// @param request - ListAggregateResourcesByAdvancedSearchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregateResourcesByAdvancedSearchResponse
func (client *Client) ListAggregateResourcesByAdvancedSearchWithContext(ctx context.Context, request *ListAggregateResourcesByAdvancedSearchRequest, runtime *dara.RuntimeOptions) (_result *ListAggregateResourcesByAdvancedSearchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all account groups within the current management account or delegated administrator account.
//
// Description:
//
// The sample request in this topic shows you how to query account groups. A maximum of 10 entries can be returned for the request. As shown in the responses, the account group returned is named as `Test_Group`, its description is `Test account group`, and it is of the `CUSTOM` type, which indicates a custom account group. The account group contains two member accounts.
//
// @param tmpReq - ListAggregatorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggregatorsResponse
func (client *Client) ListAggregatorsWithContext(ctx context.Context, tmpReq *ListAggregatorsRequest, runtime *dara.RuntimeOptions) (_result *ListAggregatorsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all compliance package templates provided by Cloud Config and the details of the compliance package templates.
//
// Description:
//
// A compliance package template is a collection of rules that Cloud Config can create based on compliance scenarios.
//
// @param request - ListCompliancePackTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCompliancePackTemplatesResponse
func (client *Client) ListCompliancePackTemplatesWithContext(ctx context.Context, request *ListCompliancePackTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListCompliancePackTemplatesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of compliance packages.
//
// Description:
//
// This topic provides an example of how to query compliance packages. The return result shows the details of the `cp-fdc8626622af00f9****` compliance package.
//
// @param tmpReq - ListCompliancePacksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCompliancePacksResponse
func (client *Client) ListCompliancePacksWithContext(ctx context.Context, tmpReq *ListCompliancePacksRequest, runtime *dara.RuntimeOptions) (_result *ListCompliancePacksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of delivery channels.
//
// @param request - ListConfigDeliveryChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConfigDeliveryChannelsResponse
func (client *Client) ListConfigDeliveryChannelsWithContext(ctx context.Context, request *ListConfigDeliveryChannelsRequest, runtime *dara.RuntimeOptions) (_result *ListConfigDeliveryChannelsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance evaluation results of resources based on a rule.
//
// Description:
//
// This topic provides an example on how to query the compliance evaluation result of resources based on a rule whose ID is `cr-cac56457e0d900d3****`. The returned result indicates that the `i-hp3e4kvhzqn2s11t****` resource is evaluated as `NON_COMPLIANT` by using the rule. The resource is an Elastic Compute Service (ECS) instance.
//
// @param request - ListConfigRuleEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConfigRuleEvaluationResultsResponse
func (client *Client) ListConfigRuleEvaluationResultsWithContext(ctx context.Context, request *ListConfigRuleEvaluationResultsRequest, runtime *dara.RuntimeOptions) (_result *ListConfigRuleEvaluationResultsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the rules of the current account.
//
// Description:
//
// This topic provides an example on how to query the rules of the current account. The response shows that the current account has a total of one rule and three evaluated resources. The resources are evaluated as compliant.
//
// @param tmpReq - ListConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConfigRulesResponse
func (client *Client) ListConfigRulesWithContext(ctx context.Context, tmpReq *ListConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *ListConfigRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of resources aggregated across regions within an Alibaba Cloud account.
//
// Description:
//
// This topic provides an example on how to call the ListDiscoveredResources operation to query the resources in the current Alibaba Cloud account. The returned result indicates that a total of eight resources exist in the account.
//
// @param request - ListDiscoveredResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiscoveredResourcesResponse
func (client *Client) ListDiscoveredResourcesWithContext(ctx context.Context, request *ListDiscoveredResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListDiscoveredResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of managed rules supported by Cloud Config.
//
// Description:
//
// ### [](#)Background information
//
// For more information about how to define, execute, and integrate a managed rule, see [Definition and execution of rules](https://help.aliyun.com/document_detail/128273.html).
//
// ### [](#)Description
//
// This topic provides an example on how to query all managed rules whose keyword is `CDN`. The response shows that 21 managed rules exist.
//
// @param request - ListManagedRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListManagedRulesResponse
func (client *Client) ListManagedRulesWithContext(ctx context.Context, request *ListManagedRulesRequest, runtime *dara.RuntimeOptions) (_result *ListManagedRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of evaluation rules supported by Cloud Config.
//
// Description:
//
// For more information about how to define, execute, and integrate an evaluation rule, see [Definition and execution of evaluation rules](https://help.aliyun.com/document_detail/470802.html).
//
// After you create an evaluation rule, a managed rule that has the same settings as the evaluation rule is created. After you create a resource, the managed rule can be used to continuously check the compliance of the resource.
//
// @param tmpReq - ListPreManagedRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPreManagedRulesResponse
func (client *Client) ListPreManagedRulesWithContext(ctx context.Context, tmpReq *ListPreManagedRulesRequest, runtime *dara.RuntimeOptions) (_result *ListPreManagedRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取推荐的托管规则列表
//
// @param request - ListRecommendManagedRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecommendManagedRulesResponse
func (client *Client) ListRecommendManagedRulesWithContext(ctx context.Context, request *ListRecommendManagedRulesRequest, runtime *dara.RuntimeOptions) (_result *ListRecommendManagedRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the remediation records of a rule.
//
// Description:
//
// This topic provides an example on how to query the remediation records of the rule cr-5392626622af0000\\*\\*\\*\\*.
//
// @param request - ListRemediationExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRemediationExecutionsResponse
func (client *Client) ListRemediationExecutionsWithContext(ctx context.Context, request *ListRemediationExecutionsRequest, runtime *dara.RuntimeOptions) (_result *ListRemediationExecutionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of remediation templates for a managed rule.
//
// Description:
//
// In this topic, the `oss-bucket-public-write-prohibited` managed rule is used as an example. The return result shows the details of the remediation template of the `OOS` type for the managed rule. OOS represents Operation Orchestration Service.
//
// @param request - ListRemediationTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRemediationTemplatesResponse
func (client *Client) ListRemediationTemplatesWithContext(ctx context.Context, request *ListRemediationTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListRemediationTemplatesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the execution of remediation templates.
//
// Description:
//
// This topic provides an example on how to query the remediation templates for the rule whose ID is `cr-6b7c626622af00b4****`.
//
// @param request - ListRemediationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRemediationsResponse
func (client *Client) ListRemediationsWithContext(ctx context.Context, request *ListRemediationsRequest, runtime *dara.RuntimeOptions) (_result *ListRemediationsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取合规报告模版详情
//
// @param request - ListReportTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReportTemplatesResponse
func (client *Client) ListReportTemplatesWithContext(ctx context.Context, request *ListReportTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListReportTemplatesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance evaluation results of resources.
//
// Description:
//
// In this example, the compliance evaluation result of the `23642660635396****` resource is queried and the resource is a RAM user. The returned result indicates that the resource is evaluated as `NON_COMPLIANT` by using the `cr-7f7d626622af0041****` rule.
//
// @param request - ListResourceEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceEvaluationResultsResponse
func (client *Client) ListResourceEvaluationResultsWithContext(ctx context.Context, request *ListResourceEvaluationResultsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceEvaluationResultsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of resources that associate with a specific resource.
//
// Description:
//
// For information about the Alibaba Cloud services and resource types supported by Cloud Config, see [Alibaba Cloud services and resource types supported by Cloud Config](https://help.aliyun.com/document_detail/127411.html).
//
// This topic provides an example on how to query the information about the disks that are attached to an Elastic Compute Service (ECS) instance named `i-j6cajg9yrfoh4sas****` that is created by the current Alibaba Cloud account in the China (Shanghai) region.
//
// @param request - ListResourceRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceRelationsResponse
func (client *Client) ListResourceRelationsWithContext(ctx context.Context, request *ListResourceRelationsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceRelationsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains resources based on the fields in the resource properties by using a SELECT statement.
//
// Description:
//
// When you write a `SELECT` statement, you must obtain the fields and the data types of the fields from the property file of the resource type. For more information about property files, see [Alibaba Cloud Config Resource Schema](https://github.com/aliyun/alibabacloud-config-resource-schema).
//
// >
//
//   - Each resource type supported by Cloud Config has a property file. Property files are named based on the related resource types. For example, the property file of the `ACS::ECS::Instance` resource type is named `ACS_ECS_Instance.properties.json`. Property files of different resource types are placed under the `config/properties/resource-types` path.
//
//   - For more information about the examples and limits on SQL query statements, see [Examples of SQL query statements](https://help.aliyun.com/document_detail/398718.html) and [Limits on SQL query statements](https://help.aliyun.com/document_detail/398750.html).
//
// This topic provides an example on how to obtain all resources whose tag key is `business` and whose tag value is `online` within the current account by using the advanced search feature.
//
// @param request - ListResourcesByAdvancedSearchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesByAdvancedSearchResponse
func (client *Client) ListResourcesByAdvancedSearchWithContext(ctx context.Context, request *ListResourcesByAdvancedSearchRequest, runtime *dara.RuntimeOptions) (_result *ListResourcesByAdvancedSearchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSupportedProductsResponse
func (client *Client) ListSupportedProductsWithContext(ctx context.Context, request *ListSupportedProductsRequest, runtime *dara.RuntimeOptions) (_result *ListSupportedProductsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tags that are added to specified resources.
//
// @param tmpReq - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, tmpReq *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits the evaluation results of a rule from Function Compute.
//
// Description:
//
// For more information about the definition, use scenarios, and execution of custom function rules, see [Definition and execution of custom function rules](https://help.aliyun.com/document_detail/127405.html).
//
// @param request - PutEvaluationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutEvaluationsResponse
func (client *Client) PutEvaluationsWithContext(ctx context.Context, request *PutEvaluationsRequest, runtime *dara.RuntimeOptions) (_result *PutEvaluationsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Re-evaluates resources that are evaluated based on a rule after the evaluation results on some resources of an ignored rule in an account group are resumed.
//
// Description:
//
// ### [](#)Prerequisites
//
// One or more non-compliant resources that are evaluated by a rule are ignored. For more information, see [IgnoreAggregateEvaluationResults](https://help.aliyun.com/document_detail/607054.html).
//
// ### [](#)Description
//
// This topic provides an example on how to re-evaluate the non-compliant resource that is evaluated by the `cr-7e72626622af0051****` rule of the `120886317861****` member in the `ca-5b6c626622af008f****` group account. The ID of the region in which the resource resides is `cn-beijing`, the type of the resource is `ACS::SLB::LoadBalancer`, and the ID of the resource is `lb-hp3a3b4ztyfm2plgm****`.
//
// @param tmpReq - RevertAggregateEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevertAggregateEvaluationResultsResponse
func (client *Client) RevertAggregateEvaluationResultsWithContext(ctx context.Context, tmpReq *RevertAggregateEvaluationResultsRequest, runtime *dara.RuntimeOptions) (_result *RevertAggregateEvaluationResultsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Re-evaluates resources that are evaluated based on a rule after the evaluation results on some resources of an ignored rule are resumed.
//
// Description:
//
// ### [](#)Prerequisites
//
// One or more non-compliant resources that are evaluated by a rule are ignored. For more information, see [IgnoreEvaluationResults](https://help.aliyun.com/document_detail/606990.html).
//
// ### [](#)Description
//
// This topic provides an example on how to re-evaluate the `lb-hp3a3b4ztyfm2plgm****` non-compliant resource that is evaluated by the `cr-7e72626622af0051****` rule. The ID of the region in which the resource resides is`cn-beijing`, the type of the resource is `ACS::SLB::LoadBalancer`, and the ID of the resource is `lb-hp3a3b4ztyfm2plgm****`.
//
// @param tmpReq - RevertEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevertEvaluationResultsResponse
func (client *Client) RevertEvaluationResultsWithContext(ctx context.Context, tmpReq *RevertEvaluationResultsRequest, runtime *dara.RuntimeOptions) (_result *RevertEvaluationResultsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Re-evaluates the compliance of resources based on a rule or based on all rules in a compliance package in a specific account group.
//
// Description:
//
// > After you call this operation, the compliance evaluation is performed only once. To query the compliance evaluation results returned by the rule, call the ListAggregateConfigRuleEvaluationResults operation. For more information, see [ListAggregateConfigRuleEvaluationResults](https://help.aliyun.com/document_detail/265979.html).
//
// The sample request in this topic shows how to use the `cr-c169626622af009f****` rule in the `ca-3a58626622af0005****` account group to evaluate resources.
//
// @param request - StartAggregateConfigRuleEvaluationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAggregateConfigRuleEvaluationResponse
func (client *Client) StartAggregateConfigRuleEvaluationWithContext(ctx context.Context, request *StartAggregateConfigRuleEvaluationRequest, runtime *dara.RuntimeOptions) (_result *StartAggregateConfigRuleEvaluationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a remediation operation by using a rule in an account group.
//
// Description:
//
// This topic provides an example on how to manually perform a remediation operation by using the rule whose ID is `cr-6b7c626622af00b4****` in the account group whose ID is `ca-6b4a626622af0012****`. The return result shows that the manual execution is successful.
//
// @param request - StartAggregateRemediationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAggregateRemediationResponse
func (client *Client) StartAggregateRemediationWithContext(ctx context.Context, request *StartAggregateRemediationRequest, runtime *dara.RuntimeOptions) (_result *StartAggregateRemediationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Re-evaluates the compliance of resources based on a rule or based on all rules in a compliance package.
//
// Description:
//
// In this example, the cr-9920626622af0035\\*\\*\\*\\	- rule is used to re-evaluate the compliance of resources.
//
// @param request - StartConfigRuleEvaluationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartConfigRuleEvaluationResponse
func (client *Client) StartConfigRuleEvaluationWithContext(ctx context.Context, request *StartConfigRuleEvaluationRequest, runtime *dara.RuntimeOptions) (_result *StartConfigRuleEvaluationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 触发单资源重新评估
//
// @param request - StartConfigRuleEvaluationByResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartConfigRuleEvaluationByResourceResponse
func (client *Client) StartConfigRuleEvaluationByResourceWithContext(ctx context.Context, request *StartConfigRuleEvaluationByResourceRequest, runtime *dara.RuntimeOptions) (_result *StartConfigRuleEvaluationByResourceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a remediation operation by using a rule.
//
// Description:
//
// This topic provides an example on how to perform a remediation operation by using the rule whose ID is `cr-8a973ac2e2be00a2****`. The returned result shows that the manual execution is successful.
//
// @param request - StartRemediationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRemediationResponse
func (client *Client) StartRemediationWithContext(ctx context.Context, request *StartRemediationRequest, runtime *dara.RuntimeOptions) (_result *StartRemediationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to resources.
//
// @param tmpReq - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, tmpReq *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实时测试通知
//
// @param request - TriggerReportSendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerReportSendResponse
func (client *Client) TriggerReportSendWithContext(ctx context.Context, request *TriggerReportSendRequest, runtime *dara.RuntimeOptions) (_result *TriggerReportSendResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from specified resources.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a compliance package in an account group.
//
// Description:
//
// This topic provides an example on how to change the value of the `eip-bandwidth-limit` parameter in the rule template of the compliance package `cp-fdc8626622af00f9****` in the account group `ca-f632626622af0079****` to `20`.
//
// @param tmpReq - UpdateAggregateCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggregateCompliancePackResponse
func (client *Client) UpdateAggregateCompliancePackWithContext(ctx context.Context, tmpReq *UpdateAggregateCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *UpdateAggregateCompliancePackResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a delivery channel in an account group.
//
// Description:
//
// This topic provides an example on how to disable a delivery channel in an account group. The ID of the account group is `ca-a4e5626622af0079****`, and the ID of the delivery channel is `cdc-8e45ff4e06a3a8****`. The Status parameter is set to `0`. After the delivery channel is disabled, Cloud Config retains the most recent delivery configuration and stops resource data delivery.
//
// @param request - UpdateAggregateConfigDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggregateConfigDeliveryChannelResponse
func (client *Client) UpdateAggregateConfigDeliveryChannelWithContext(ctx context.Context, request *UpdateAggregateConfigDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *UpdateAggregateConfigDeliveryChannelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description, input parameters, and risk level of a rule in a specific account group.
//
// Description:
//
// This topic provides an example on how to change the risk level of the rule `cr-4e3d626622af0080****` in an account group `ca-a4e5626622af0079****` to `3`, which indicates low risk level.
//
// @param tmpReq - UpdateAggregateConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggregateConfigRuleResponse
func (client *Client) UpdateAggregateConfigRuleWithContext(ctx context.Context, tmpReq *UpdateAggregateConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateAggregateConfigRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a remediation template for a rule in an account group.
//
// Description:
//
// This topic describes how to change the execution mode of the `crr-909ba2d4716700eb****` remediation setting for a rule in the `ca-6b4a626622af0012****` account group to `AUTO_EXECUTION`, which specifies automatic remediation. This topic also provides a sample request.
//
// @param request - UpdateAggregateRemediationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggregateRemediationResponse
func (client *Client) UpdateAggregateRemediationWithContext(ctx context.Context, request *UpdateAggregateRemediationRequest, runtime *dara.RuntimeOptions) (_result *UpdateAggregateRemediationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The management account or delegated administrator account of a resource directory can be used to modify the name and description of an account group. The management account or delegated administrator account can also be used to add or remove members from the account group.
//
// Description:
//
// This topic provides an example on how to add a member to the account group `ca-dacf86d8314e00eb****`. The member ID is `173808452267****`, the member name is `Tony`, and the member belongs to the resource directory `ResourceDirectory`.
//
// @param tmpReq - UpdateAggregatorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggregatorResponse
func (client *Client) UpdateAggregatorWithContext(ctx context.Context, tmpReq *UpdateAggregatorRequest, runtime *dara.RuntimeOptions) (_result *UpdateAggregatorResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a specific compliance package in the current account.
//
// Description:
//
// This topic provides an example on how to change the value of the `eip-bandwidth-limit` parameter of a rule in the compliance package `cp-a8a8626622af0082****` to `20`.
//
// @param tmpReq - UpdateCompliancePackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCompliancePackResponse
func (client *Client) UpdateCompliancePackWithContext(ctx context.Context, tmpReq *UpdateCompliancePackRequest, runtime *dara.RuntimeOptions) (_result *UpdateCompliancePackResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a delivery channel by using the current account.
//
// Description:
//
// In this example, a delivery channel is disabled. The ID of the delivery channel is `cdc-8e45ff4e06a3a8****```. The Status parameter is set to 0. After the delivery channel is disabled, Cloud Config retains the most recent delivery configuration and stops the delivery of resource data.
//
// @param request - UpdateConfigDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConfigDeliveryChannelResponse
func (client *Client) UpdateConfigDeliveryChannelWithContext(ctx context.Context, request *UpdateConfigDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *UpdateConfigDeliveryChannelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// This topic provides an example on how to change the risk level of the rule `cr-a260626622af0005****` to `3`, which indicates low risk level.
//
// @param tmpReq - UpdateConfigRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConfigRuleResponse
func (client *Client) UpdateConfigRuleWithContext(ctx context.Context, tmpReq *UpdateConfigRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateConfigRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the resource monitoring scope of the current account.
//
// Description:
//
// This topic provides an example on how to change the resource monitoring scope of the current account to ACS::ECS::Instance.
//
// @param request - UpdateConfigurationRecorderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConfigurationRecorderResponse
func (client *Client) UpdateConfigurationRecorderWithContext(ctx context.Context, request *UpdateConfigurationRecorderRequest, runtime *dara.RuntimeOptions) (_result *UpdateConfigurationRecorderResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the integration of a cloud service.
//
// @param request - UpdateIntegratedServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIntegratedServiceStatusResponse
func (client *Client) UpdateIntegratedServiceStatusWithContext(ctx context.Context, request *UpdateIntegratedServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateIntegratedServiceStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a remediation template for a rule.
//
// Description:
//
// This topic describes how to change the execution mode of the `crr-909ba2d4716700eb****` remediation setting to `AUTO_EXECUTION`, which specifies automatic remediation. This topic also provides a sample request.
//
// @param request - UpdateRemediationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRemediationResponse
func (client *Client) UpdateRemediationWithContext(ctx context.Context, request *UpdateRemediationRequest, runtime *dara.RuntimeOptions) (_result *UpdateRemediationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新合规报告模版
//
// @param tmpReq - UpdateReportTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateReportTemplateResponse
func (client *Client) UpdateReportTemplateWithContext(ctx context.Context, tmpReq *UpdateReportTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateReportTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
