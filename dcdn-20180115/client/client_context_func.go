// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds a domain name to accelerate. You can specify only one domain name in each request.
//
// Description:
//
// > 	- Dynamic Content Delivery Network (DCDN) is activated.
//
// > 	- Internet content provider (ICP) filing is complete for the accelerated domain name.
//
// > 	- If the content of the origin server is not stored on Alibaba Cloud, the content must be reviewed. After you submit the request, the review is complete by the end of the following business day.
//
// > 	- You can call this operation up to 30 times per second per account.
//
// @param request - AddDcdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDcdnDomainResponse
func (client *Client) AddDcdnDomainWithContext(ctx context.Context, request *AddDcdnDomainRequest, runtime *dara.RuntimeOptions) (_result *AddDcdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckUrl) {
		query["CheckUrl"] = request.CheckUrl
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionType) {
		query["FunctionType"] = request.FunctionType
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

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TopLevelDomain) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDcdnDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDcdnDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a domain name to IPA. You can specify only one domain name in each request.
//
// Description:
//
// >
//
//   - Make sure that the IPA service is activated before you add a domain name to accelerate.
//
//   - Make sure that the Internet content provider (ICP) filling is complete for the domain name to accelerate.
//
//   - If the content on the origin server is not stored on Alibaba Cloud, the content must be reviewed. The review is complete by the end of the next business day after you submit the request.
//
//   - You can call this operation up to 10 times per second per user.
//
// @param request - AddDcdnIpaDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDcdnIpaDomainResponse
func (client *Client) AddDcdnIpaDomainWithContext(ctx context.Context, request *AddDcdnIpaDomainRequest, runtime *dara.RuntimeOptions) (_result *AddDcdnIpaDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckUrl) {
		query["CheckUrl"] = request.CheckUrl
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	if !dara.IsNil(request.TopLevelDomain) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDcdnIpaDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDcdnIpaDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds one or more domain names to DCDN at a time.
//
// Description:
//
// *Prerequisites**:
//
//   - The [DCDN service is activated](https://help.aliyun.com/document_detail/64926.html).
//
//   - Internet content provider (ICP) filing is complete for the accelerated domain names.
//
// > 	- If the content of the origin server is not stored on Alibaba Cloud, the content must be reviewed. After you submit the request, the review is complete by the end of the following business day.
//
// >	- You can specify up to 50 domain names in each request.
//
// >	- You can call this operation up to 30 times per second per account.
//
// @param request - BatchAddDcdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddDcdnDomainResponse
func (client *Client) BatchAddDcdnDomainWithContext(ctx context.Context, request *BatchAddDcdnDomainRequest, runtime *dara.RuntimeOptions) (_result *BatchAddDcdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckUrl) {
		query["CheckUrl"] = request.CheckUrl
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
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

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	if !dara.IsNil(request.TopLevelDomain) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchAddDcdnDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchAddDcdnDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates Web Application Firewall (WAF) protection rules.
//
// Description:
//
// >  You can call this operation up to 20 times per second per account.
//
// @param request - BatchCreateDcdnWafRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateDcdnWafRulesResponse
func (client *Client) BatchCreateDcdnWafRulesWithContext(ctx context.Context, request *BatchCreateDcdnWafRulesRequest, runtime *dara.RuntimeOptions) (_result *BatchCreateDcdnWafRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RuleConfigs) {
		body["RuleConfigs"] = request.RuleConfigs
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCreateDcdnWafRules"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCreateDcdnWafRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes configurations of multiple accelerated domain names at a time.
//
// Description:
//
// > - You can specify up to 50 domain names in each request.
//
// > - You can call this operation up to 30 times per second per account.
//
// @param request - BatchDeleteDcdnDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteDcdnDomainConfigsResponse
func (client *Client) BatchDeleteDcdnDomainConfigsWithContext(ctx context.Context, request *BatchDeleteDcdnDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteDcdnDomainConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.FunctionNames) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteDcdnDomainConfigs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteDcdnDomainConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除kv数据，支持最大2M的请求体
//
// @param tmpReq - BatchDeleteDcdnKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteDcdnKvResponse
func (client *Client) BatchDeleteDcdnKvWithContext(ctx context.Context, tmpReq *BatchDeleteDcdnKvRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteDcdnKvResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &BatchDeleteDcdnKvShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Keys) {
		request.KeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keys, dara.String("Keys"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.KeysShrink) {
		body["Keys"] = request.KeysShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteDcdnKv"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteDcdnKvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除kv数据，支持最大100M的请求体
//
// @param request - BatchDeleteDcdnKvWithHighCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteDcdnKvWithHighCapacityResponse
func (client *Client) BatchDeleteDcdnKvWithHighCapacityWithContext(ctx context.Context, request *BatchDeleteDcdnKvWithHighCapacityRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteDcdnKvWithHighCapacityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteDcdnKvWithHighCapacity"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteDcdnKvWithHighCapacityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple Web Application Firewall (WAF) protection rules at a time.
//
// Description:
//
//	  You can call this operation up to 20 times per second per account.
//
//		- Alibaba Cloud Dynamic Content Delivery Network (DCDN) supports POST requests.
//
// @param request - BatchDeleteDcdnWafRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteDcdnWafRulesResponse
func (client *Client) BatchDeleteDcdnWafRulesWithContext(ctx context.Context, request *BatchDeleteDcdnWafRulesRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteDcdnWafRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RuleIds) {
		body["RuleIds"] = request.RuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteDcdnWafRules"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteDcdnWafRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies multiple Web Application Firewall (WAF) protection rules. Only Bot management rules can be modified.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - BatchModifyDcdnWafRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchModifyDcdnWafRulesResponse
func (client *Client) BatchModifyDcdnWafRulesWithContext(ctx context.Context, request *BatchModifyDcdnWafRulesRequest, runtime *dara.RuntimeOptions) (_result *BatchModifyDcdnWafRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RuleConfigs) {
		body["RuleConfigs"] = request.RuleConfigs
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchModifyDcdnWafRules"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchModifyDcdnWafRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures multiple key-value (KV) pairs for a namespace.
//
// @param tmpReq - BatchPutDcdnKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchPutDcdnKvResponse
func (client *Client) BatchPutDcdnKvWithContext(ctx context.Context, tmpReq *BatchPutDcdnKvRequest, runtime *dara.RuntimeOptions) (_result *BatchPutDcdnKvResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &BatchPutDcdnKvShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.KvList) {
		request.KvListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KvList, dara.String("KvList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.KvListShrink) {
		body["KvList"] = request.KvListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchPutDcdnKv"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchPutDcdnKvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量写入kv数据，支持最大100M的请求体
//
// @param request - BatchPutDcdnKvWithHighCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchPutDcdnKvWithHighCapacityResponse
func (client *Client) BatchPutDcdnKvWithHighCapacityWithContext(ctx context.Context, request *BatchPutDcdnKvWithHighCapacityRequest, runtime *dara.RuntimeOptions) (_result *BatchPutDcdnKvWithHighCapacityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchPutDcdnKvWithHighCapacity"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchPutDcdnKvWithHighCapacityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures features for one or more domain names.
//
// Description:
//
//	  You can specify up to 50 domain names in each request.
//
//		- You can call this operation up to 30 times per second per account.
//
// @param request - BatchSetDcdnDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSetDcdnDomainConfigsResponse
func (client *Client) BatchSetDcdnDomainConfigsWithContext(ctx context.Context, request *BatchSetDcdnDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *BatchSetDcdnDomainConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.Functions) {
		query["Functions"] = request.Functions
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchSetDcdnDomainConfigs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSetDcdnDomainConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures multiple domain names to be accelerated by IP Application Accelerator (IPA).
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - BatchSetDcdnIpaDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSetDcdnIpaDomainConfigsResponse
func (client *Client) BatchSetDcdnIpaDomainConfigsWithContext(ctx context.Context, request *BatchSetDcdnIpaDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *BatchSetDcdnIpaDomainConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.Functions) {
		query["Functions"] = request.Functions
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchSetDcdnIpaDomainConfigs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSetDcdnIpaDomainConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the protection status of multiple domain names at a time.
//
// Description:
//
// #
//
//   - You can call this operation up to 20 times per second.
//
//   - Alibaba Cloud Dynamic Content Delivery Network (DCDN) supports POST requests.
//
// @param request - BatchSetDcdnWafDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSetDcdnWafDomainConfigsResponse
func (client *Client) BatchSetDcdnWafDomainConfigsWithContext(ctx context.Context, request *BatchSetDcdnWafDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *BatchSetDcdnWafDomainConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientIpTag) {
		body["ClientIpTag"] = request.ClientIpTag
	}

	if !dara.IsNil(request.DefenseStatus) {
		body["DefenseStatus"] = request.DefenseStatus
	}

	if !dara.IsNil(request.DomainNames) {
		body["DomainNames"] = request.DomainNames
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchSetDcdnWafDomainConfigs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSetDcdnWafDomainConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables one or more accelerated domain names. After the accelerated domain names are enabled, the value of the DomainStatus parameter for the domain names changes to Online.
//
// Description:
//
// >
//
//   - If an accelerated domain name is in an invalid state or your account has an overdue payment, the accelerated domain name cannot be enabled.
//
//   - You can specify up to 50 domain names in each request.
//
//   - You can call this operation up to 30 times per second per account.
//
// @param request - BatchStartDcdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStartDcdnDomainResponse
func (client *Client) BatchStartDcdnDomainWithContext(ctx context.Context, request *BatchStartDcdnDomainRequest, runtime *dara.RuntimeOptions) (_result *BatchStartDcdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchStartDcdnDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchStartDcdnDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables one or more accelerated domain names. After an accelerated domain name is disabled, the value of the DomainStatus parameter for the domain name changes to Offline.
//
// Description:
//
// > 	- After an accelerated domain name is disabled, Dynamic Content Delivery Network (DCDN) retains the domain name information. The system automatically reroutes all requests that are destined for the accelerated domain name to the origin.
//
// >	- You can specify up to 50 domain names in each request.
//
// >	- You can call this operation up to 30 times per second per account.
//
// @param request - BatchStopDcdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStopDcdnDomainResponse
func (client *Client) BatchStopDcdnDomainWithContext(ctx context.Context, request *BatchStopDcdnDomainRequest, runtime *dara.RuntimeOptions) (_result *BatchStopDcdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchStopDcdnDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchStopDcdnDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a real-time log delivery project exists.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - CheckDcdnProjectExistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDcdnProjectExistResponse
func (client *Client) CheckDcdnProjectExistWithContext(ctx context.Context, request *CheckDcdnProjectExistRequest, runtime *dara.RuntimeOptions) (_result *CheckDcdnProjectExistResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDcdnProjectExist"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDcdnProjectExistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates an official code version from unstable JavaScript code that is in the staging
//
//	environment. The version can be used in the canary release or production environment.
//
// Description:
//
// >  The call frequency of the API is no more than 100 queries per second.
//
// @param request - CommitStagingRoutineCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CommitStagingRoutineCodeResponse
func (client *Client) CommitStagingRoutineCodeWithContext(ctx context.Context, request *CommitStagingRoutineCodeRequest, runtime *dara.RuntimeOptions) (_result *CommitStagingRoutineCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CodeDescription) {
		body["CodeDescription"] = request.CodeDescription
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CommitStagingRoutineCode"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CommitStagingRoutineCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a certificate signing request (CSR) file.
//
// @param request - CreateDcdnCertificateSigningRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDcdnCertificateSigningRequestResponse
func (client *Client) CreateDcdnCertificateSigningRequestWithContext(ctx context.Context, request *CreateDcdnCertificateSigningRequestRequest, runtime *dara.RuntimeOptions) (_result *CreateDcdnCertificateSigningRequestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.CommonName) {
		query["CommonName"] = request.CommonName
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Organization) {
		query["Organization"] = request.Organization
	}

	if !dara.IsNil(request.OrganizationUnit) {
		query["OrganizationUnit"] = request.OrganizationUnit
	}

	if !dara.IsNil(request.SANs) {
		query["SANs"] = request.SANs
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDcdnCertificateSigningRequest"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDcdnCertificateSigningRequestResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a tracking task. After you create a tracking task, the system periodically sends operations reports to you by email.
//
// Description:
//
// *
//
// **You can call this operation up to three times per second.
//
// @param request - CreateDcdnDeliverTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDcdnDeliverTaskResponse
func (client *Client) CreateDcdnDeliverTaskWithContext(ctx context.Context, request *CreateDcdnDeliverTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDcdnDeliverTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Deliver) {
		body["Deliver"] = request.Deliver
	}

	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Reports) {
		body["Reports"] = request.Reports
	}

	if !dara.IsNil(request.Schedule) {
		body["Schedule"] = request.Schedule
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDcdnDeliverTask"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDcdnDeliverTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a real-time log delivery project.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - CreateDcdnSLSRealTimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDcdnSLSRealTimeLogDeliveryResponse
func (client *Client) CreateDcdnSLSRealTimeLogDeliveryWithContext(ctx context.Context, request *CreateDcdnSLSRealTimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *CreateDcdnSLSRealTimeLogDeliveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		body["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.DataCenter) {
		body["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SLSLogStore) {
		body["SLSLogStore"] = request.SLSLogStore
	}

	if !dara.IsNil(request.SLSProject) {
		body["SLSProject"] = request.SLSProject
	}

	if !dara.IsNil(request.SLSRegion) {
		body["SLSRegion"] = request.SLSRegion
	}

	if !dara.IsNil(request.SamplingRate) {
		body["SamplingRate"] = request.SamplingRate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDcdnSLSRealTimeLogDelivery"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDcdnSLSRealTimeLogDeliveryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom operations report.
//
// Description:
//
// > 	- This operation allows you to create a custom operations report for a specific domain name. You can view the statistics about the domain name in the report.
//
// > 	- You can call this operation up to three times per second per account.
//
// @param request - CreateDcdnSubTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDcdnSubTaskResponse
func (client *Client) CreateDcdnSubTaskWithContext(ctx context.Context, request *CreateDcdnSubTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDcdnSubTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.ReportIds) {
		body["ReportIds"] = request.ReportIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDcdnSubTask"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDcdnSubTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a custom WAF rule group.
//
// @param request - CreateDcdnWafGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDcdnWafGroupResponse
func (client *Client) CreateDcdnWafGroupWithContext(ctx context.Context, request *CreateDcdnWafGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDcdnWafGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Subscribe) {
		body["Subscribe"] = request.Subscribe
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDcdnWafGroup"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDcdnWafGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Web Application Firewall (WAF) protection policy.
//
// Description:
//
//	  You can call this operation up to 20 times per second per user.
//
//		- Alibaba Cloud Dynamic Route for CDN (DCDN) supports POST requests.
//
// @param request - CreateDcdnWafPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDcdnWafPolicyResponse
func (client *Client) CreateDcdnWafPolicyWithContext(ctx context.Context, request *CreateDcdnWafPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateDcdnWafPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DefenseScene) {
		body["DefenseScene"] = request.DefenseScene
	}

	if !dara.IsNil(request.PolicyName) {
		body["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicyStatus) {
		body["PolicyStatus"] = request.PolicyStatus
	}

	if !dara.IsNil(request.PolicyType) {
		body["PolicyType"] = request.PolicyType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDcdnWafPolicy"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDcdnWafPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a routine.
//
// Description:
//
// > 	- The parameters must comply with the rules of EnvConf. The description of a routine cannot exceed 50 characters in length.
//
// >	- You can only specify the production and staging environments when you call this operation.
//
// >	- You can call this operation up to 100 times per second per account.
//
// @param tmpReq - CreateRoutineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoutineResponse
func (client *Client) CreateRoutineWithContext(ctx context.Context, tmpReq *CreateRoutineRequest, runtime *dara.RuntimeOptions) (_result *CreateRoutineResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateRoutineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EnvConf) {
		request.EnvConfShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EnvConf, dara.String("EnvConf"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnvConfShrink) {
		body["EnvConf"] = request.EnvConfShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRoutine"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRoutineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service-linked role (SLR) and a Log Service project.
//
// Description:
//
// >  You can call this operation up to 100 times per second per account.
//
// @param request - CreateSlrAndSlsProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSlrAndSlsProjectResponse
func (client *Client) CreateSlrAndSlsProjectWithContext(ctx context.Context, request *CreateSlrAndSlsProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateSlrAndSlsProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		body["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSlrAndSlsProject"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSlrAndSlsProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # A客户定制实时日志删除接口
//
// @param request - DeleteCustomDomainSampleRateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomDomainSampleRateResponse
func (client *Client) DeleteCustomDomainSampleRateWithContext(ctx context.Context, request *DeleteCustomDomainSampleRateRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomDomainSampleRateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		body["DomainNames"] = request.DomainNames
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomDomainSampleRate"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomDomainSampleRateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes tracking tasks by task ID.
//
// Description:
//
// >  The maximum number of times that each user can call this operation per second is 3.
//
// @param request - DeleteDcdnDeliverTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnDeliverTaskResponse
func (client *Client) DeleteDcdnDeliverTaskWithContext(ctx context.Context, request *DeleteDcdnDeliverTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnDeliverTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliverId) {
		query["DeliverId"] = request.DeliverId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnDeliverTask"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnDeliverTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified accelerated domain name.
//
// Description:
//
// > 	- Before you delete your domain name, you need to request the Domain Name System (DNS) provider to restore the A record of the domain name. Otherwise, the domain name may become inaccessible after you delete it.
//
// > 	- If you call the **DeleteDcdnDomain*	- operation, all the information about the accelerated domain name is deleted. If you want to disable an accelerated domain name, call the [StopDcdnDomain](https://help.aliyun.com/document_detail/130622.html) operation.
//
// > 	- You can call this operation up to 10 times per second per account.
//
// @param request - DeleteDcdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnDomainResponse
func (client *Client) DeleteDcdnDomainWithContext(ctx context.Context, request *DeleteDcdnDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an accelerated domain name from IP Application Accelerator (IPA).
//
// Description:
//
// >
//
//   - Before you delete your domain name, we recommend that you request the Domain Name System (DNS) provider to restore the A record of the domain name. Otherwise, the domain name may become inaccessible after you delete it.
//
//   - This operation deletes all records of the specified accelerated domain name. If you want to temporarily disable an accelerated domain name, call the **StopDcdnIpaDomain*	- operation.****
//
//   - You can call this operation up to 10 times per second per account.
//
// @param request - DeleteDcdnIpaDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnIpaDomainResponse
func (client *Client) DeleteDcdnIpaDomainWithContext(ctx context.Context, request *DeleteDcdnIpaDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnIpaDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnIpaDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnIpaDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes specific configurations of an accelerated domain name from IP Application Accelerator (IPA).
//
// Description:
//
// > You can call this operation up to 10 times per second per account.
//
// @param request - DeleteDcdnIpaSpecificConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnIpaSpecificConfigResponse
func (client *Client) DeleteDcdnIpaSpecificConfigWithContext(ctx context.Context, request *DeleteDcdnIpaSpecificConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnIpaSpecificConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnIpaSpecificConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnIpaSpecificConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the key-value pairs in a namespace that you specify when you call the PutDcdnKvNamespace operation. EdgeKV provides a global key-value database for Dynamic Route for CDN (DCDN) points of presence (POPs).
//
// @param request - DeleteDcdnKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnKvResponse
func (client *Client) DeleteDcdnKvWithContext(ctx context.Context, request *DeleteDcdnKvRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnKvResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnKv"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnKvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a namespace that belongs to your account.
//
// @param request - DeleteDcdnKvNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnKvNamespaceResponse
func (client *Client) DeleteDcdnKvNamespaceWithContext(ctx context.Context, request *DeleteDcdnKvNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnKvNamespaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnKvNamespace"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnKvNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the request.
//
// Description:
//
// >  You can call this operation up to 100 times per second per account.
//
// @param request - DeleteDcdnRealTimeLogProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnRealTimeLogProjectResponse
func (client *Client) DeleteDcdnRealTimeLogProjectWithContext(ctx context.Context, request *DeleteDcdnRealTimeLogProjectRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnRealTimeLogProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnRealTimeLogProject"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnRealTimeLogProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes configurations of a domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DeleteDcdnSpecificConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnSpecificConfigResponse
func (client *Client) DeleteDcdnSpecificConfigWithContext(ctx context.Context, request *DeleteDcdnSpecificConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnSpecificConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnSpecificConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnSpecificConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the configurations of an accelerated domain name in the canary release environment.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DeleteDcdnSpecificStagingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnSpecificStagingConfigResponse
func (client *Client) DeleteDcdnSpecificStagingConfigWithContext(ctx context.Context, request *DeleteDcdnSpecificStagingConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnSpecificStagingConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnSpecificStagingConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnSpecificStagingConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes feature configurations by user.
//
// @param request - DeleteDcdnUserConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnUserConfigResponse
func (client *Client) DeleteDcdnUserConfigWithContext(ctx context.Context, request *DeleteDcdnUserConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnUserConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnUserConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnUserConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom WAF rule group.
//
// @param request - DeleteDcdnWafGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnWafGroupResponse
func (client *Client) DeleteDcdnWafGroupWithContext(ctx context.Context, request *DeleteDcdnWafGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnWafGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnWafGroup"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnWafGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a protection policy.
//
// Description:
//
//	  You can call this operation up to 20 times per second per account.
//
//		- Alibaba Cloud Dynamic Content Delivery Network (DCDN) supports POST requests.
//
// @param request - DeleteDcdnWafPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnWafPolicyResponse
func (client *Client) DeleteDcdnWafPolicyWithContext(ctx context.Context, request *DeleteDcdnWafPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnWafPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnWafPolicy"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnWafPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a routine.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DeleteRoutineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoutineResponse
func (client *Client) DeleteRoutineWithContext(ctx context.Context, request *DeleteRoutineRequest, runtime *dara.RuntimeOptions) (_result *DeleteRoutineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRoutine"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRoutineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the code of the specified version from a routine.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DeleteRoutineCodeRevisionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoutineCodeRevisionResponse
func (client *Client) DeleteRoutineCodeRevisionWithContext(ctx context.Context, request *DeleteRoutineCodeRevisionRequest, runtime *dara.RuntimeOptions) (_result *DeleteRoutineCodeRevisionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SelectCodeRevision) {
		body["SelectCodeRevision"] = request.SelectCodeRevision
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRoutineCodeRevision"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRoutineCodeRevisionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes canary release environments from a routine.
//
// Description:
//
// >
//
//   - This operation deletes only custom preset canary release environments. You cannot delete production or staging environments.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param tmpReq - DeleteRoutineConfEnvsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoutineConfEnvsResponse
func (client *Client) DeleteRoutineConfEnvsWithContext(ctx context.Context, tmpReq *DeleteRoutineConfEnvsRequest, runtime *dara.RuntimeOptions) (_result *DeleteRoutineConfEnvsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteRoutineConfEnvsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Envs) {
		request.EnvsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Envs, dara.String("Envs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EnvsShrink) {
		body["Envs"] = request.EnvsShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRoutineConfEnvs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRoutineConfEnvsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # A客户定制查询域名采样率
//
// @param request - DescribeCustomDomainSampleRateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomDomainSampleRateResponse
func (client *Client) DescribeCustomDomainSampleRateWithContext(ctx context.Context, request *DescribeCustomDomainSampleRateRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomDomainSampleRateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
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
		Action:      dara.String("DescribeCustomDomainSampleRate"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomDomainSampleRateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries precise access control rules.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - DescribeDcdnAclFieldsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnAclFieldsResponse
func (client *Client) DescribeDcdnAclFieldsWithContext(ctx context.Context, request *DescribeDcdnAclFieldsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnAclFieldsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnAclFields"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnAclFieldsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries bandwidth data for Border Gateway Protocol (BGP) accelerated domain names. Data is collected every 5 minutes.
//
// Description:
//
//	  If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both the StartTime and EndTime parameters, the request returns the data collected within the specified time range. You must set both parameters or leave both of them empty.
//
//		- If you specify multiple Internet service providers (ISPs), the data for the ISPs is aggregated.
//
//		- You can query data in the last 90 days.
//
//		- The maximum time range from the start time to the end time is 31 days. The start time is specified by the StartTime parameter and the end time is specified by the EndTime parameter.
//
//		- If the time range from the start time to the end time is 72 hours or shorter, you can specify the interval as 5 minutes. If the time range is longer than 72 hours, you must specify the interval as 1 hour.
//
//		- You can call this operation up to five times per second per account.
//
// @param request - DescribeDcdnBgpBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnBgpBpsDataResponse
func (client *Client) DescribeDcdnBgpBpsDataWithContext(ctx context.Context, request *DescribeDcdnBgpBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnBgpBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.DevicePort) {
		query["DevicePort"] = request.DevicePort
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnBgpBpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnBgpBpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries traffic data for BGP accelerated domain names. Data is collected every 5 minutes.
//
// Description:
//
//	  If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range. You must set both parameters or leave both parameters empty.
//
//		- If you specify multiple Internet service providers (ISPs), the data for the ISPs is aggregated.
//
//		- You can query data in the last 90 days.
//
//		- The maximum time range that you can specify is 31 days. StartTime specifies the start time and EndTime specifies the end time of the time range.
//
//		- If the time range from the start time to the end time is 72 hours or shorter, you can specify the interval as 5 minutes. If the time range is longer than 72 hours, you must specify the interval as 1 hour.
//
//		- You can call this operation up to five times per second per account.
//
// @param request - DescribeDcdnBgpTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnBgpTrafficDataResponse
func (client *Client) DescribeDcdnBgpTrafficDataWithContext(ctx context.Context, request *DescribeDcdnBgpTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnBgpTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnBgpTrafficData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnBgpTrafficDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries countries and regions that can be added to the blacklist.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnBlockedRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnBlockedRegionsResponse
func (client *Client) DescribeDcdnBlockedRegionsWithContext(ctx context.Context, request *DescribeDcdnBlockedRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnBlockedRegionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnBlockedRegions"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnBlockedRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about a certificate.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnCertificateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnCertificateDetailResponse
func (client *Client) DescribeDcdnCertificateDetailWithContext(ctx context.Context, request *DescribeDcdnCertificateDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnCertificateDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnCertificateDetail"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnCertificateDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeDcdnCertificateList is deprecated, please use dcdn::2018-01-15::DescribeDcdnSSLCertificateList instead.
//
// Summary:
//
// Queries the certificates of one or more accelerated domain names.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnCertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnCertificateListResponse
func (client *Client) DescribeDcdnCertificateListWithContext(ctx context.Context, request *DescribeDcdnCertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnCertificateListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnCertificateList"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnCertificateListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of DCDN DDoS mitigation.
//
// @param request - DescribeDcdnDdosServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDdosServiceResponse
func (client *Client) DescribeDcdnDdosServiceWithContext(ctx context.Context, request *DescribeDcdnDdosServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDdosServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDdosService"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDdosServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the domain names that are deleted from your Alibaba Cloud account.
//
// Description:
//
// > You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDcdnDeletedDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDeletedDomainsResponse
func (client *Client) DescribeDcdnDeletedDomainsWithContext(ctx context.Context, request *DescribeDcdnDeletedDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDeletedDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DescribeDcdnDeletedDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDeletedDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all tracking tasks of operations reports.
//
// Description:
//
// >You can call this operation up to three times per second.
//
// @param request - DescribeDcdnDeliverListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDeliverListResponse
func (client *Client) DescribeDcdnDeliverListWithContext(ctx context.Context, request *DescribeDcdnDeliverListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDeliverListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliverId) {
		query["DeliverId"] = request.DeliverId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDeliverList"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDeliverListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring data of network bandwidth for one or more accelerated domain names. You can query data in the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainBpsDataResponse
func (client *Client) DescribeDcdnDomainBpsDataWithContext(ctx context.Context, request *DescribeDcdnDomainBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainBpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainBpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries bandwidth data of accelerated domain names.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
//   - If you do not set **StartTime*	- or **EndTime**, the request returns the data collected in the last 24 hours. If you set both **StartTime*	- and **EndTime**, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the time range to query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainBpsDataByLayerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainBpsDataByLayerResponse
func (client *Client) DescribeDcdnDomainBpsDataByLayerWithContext(ctx context.Context, request *DescribeDcdnDomainBpsDataByLayerRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainBpsDataByLayerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.Layer) {
		query["Layer"] = request.Layer
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainBpsDataByLayer"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainBpsDataByLayerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries accelerated domain names by SSL certificate.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainByCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainByCertificateResponse
func (client *Client) DescribeDcdnDomainByCertificateWithContext(ctx context.Context, request *DescribeDcdnDomainByCertificateRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainByCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Exact) {
		query["Exact"] = request.Exact
	}

	if !dara.IsNil(request.SSLPub) {
		query["SSLPub"] = request.SSLPub
	}

	if !dara.IsNil(request.SSLStatus) {
		query["SSLStatus"] = request.SSLStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainByCertificate"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainByCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries logs of rate limiting.
//
// Description:
//
// >
//
//   - If you do not configure the StartTime or EndTime parameter, data collected over the last 24 hours is queried. If you configure both the StartTime and EndTime parameters, data collected within the specified time range is queried.
//
//   - You can query data collected over the last 30 days.
//
//   - You can call the RefreshObjectCaches operation up to 50 times per second per account.
//
// @param request - DescribeDcdnDomainCcActivityLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainCcActivityLogResponse
func (client *Client) DescribeDcdnDomainCcActivityLogWithContext(ctx context.Context, request *DescribeDcdnDomainCcActivityLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainCcActivityLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TriggerObject) {
		query["TriggerObject"] = request.TriggerObject
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainCcActivityLog"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainCcActivityLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificate information about an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainCertificateInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainCertificateInfoResponse
func (client *Client) DescribeDcdnDomainCertificateInfoWithContext(ctx context.Context, request *DescribeDcdnDomainCertificateInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainCertificateInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainCertificateInfo"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainCertificateInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether CNAME records are configured for one or more accelerated domain names.
//
// Description:
//
// > You can call this operation up to 80 times per second per account.
//
// @param request - DescribeDcdnDomainCnameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainCnameResponse
func (client *Client) DescribeDcdnDomainCnameWithContext(ctx context.Context, request *DescribeDcdnDomainCnameRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainCnameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainCname"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainCnameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of an accelerated domain name.
//
// Description:
//
// > 	- You can query the configurations of one or more features in a request.
//
// > 	- You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainConfigsResponse
func (client *Client) DescribeDcdnDomainConfigsWithContext(ctx context.Context, request *DescribeDcdnDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionNames) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainConfigs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic configuration information about an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnDomainDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainDetailResponse
func (client *Client) DescribeDcdnDomainDetailWithContext(ctx context.Context, request *DescribeDcdnDomainDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainDetail"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request hit ratios of one or more accelerated domain names. You can query data collected within the last 90 days.
//
// Description:
//
// #
//
//   - You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity*	- The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table. |Time granularity |Maximum time range per query |Historical data available |Data delay | -------------- | -------------- | ------ |5 minutes |3 days |93 days |15 minutes |1 hour |31 days |186 days |4 hours |1 day |366 days |366 days |04:00 on the next day
//
// @param request - DescribeDcdnDomainHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainHitRateDataResponse
func (client *Client) DescribeDcdnDomainHitRateDataWithContext(ctx context.Context, request *DescribeDcdnDomainHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainHitRateDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainHitRateData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainHitRateDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total number and proportions of HTTP status codes returned from one or more accelerated domain names. Data is collected every 5 minutes. You can query data in the last 90 days.
//
// Description:
//
// If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
//   - You can call this operation up to 100 times per second per account.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainHttpCodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainHttpCodeDataResponse
func (client *Client) DescribeDcdnDomainHttpCodeDataWithContext(ctx context.Context, request *DescribeDcdnDomainHttpCodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainHttpCodeDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainHttpCodeData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainHttpCodeDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the distribution of HTTP status codes by protocol.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
//   - You cannot query the distribution of HTTP status codes by IP protocol.
//
//   - If you do not specify the **StartTime*	- or **EndTime*	- parameter, the data that is collected within the last 24 hours is collected. If you specify both the **StartTime*	- and **EndTime*	- parameters, the data that is collected within the time range that you specify is collected.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainHttpCodeDataByLayerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainHttpCodeDataByLayerResponse
func (client *Client) DescribeDcdnDomainHttpCodeDataByLayerWithContext(ctx context.Context, request *DescribeDcdnDomainHttpCodeDataByLayerRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainHttpCodeDataByLayerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.Layer) {
		query["Layer"] = request.Layer
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainHttpCodeDataByLayer"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainHttpCodeDataByLayerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries bandwidth of accelerated domain names for which Layer 4 acceleration is enabled. You can query the data that is collected over the last 90 days.
//
// Description:
//
// >
//
//   - The bandwidth is measured in bit/s.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainIpaBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainIpaBpsDataResponse
func (client *Client) DescribeDcdnDomainIpaBpsDataWithContext(ctx context.Context, request *DescribeDcdnDomainIpaBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainIpaBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FixTimeGap) {
		query["FixTimeGap"] = request.FixTimeGap
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TimeMerge) {
		query["TimeMerge"] = request.TimeMerge
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainIpaBpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainIpaBpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of IPA user connections.
//
// Description:
//
//	  You can call this operation up to 10 times per second per user.
//
//		- If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
//		- The minimum time granularity at which the data is queried is 5 minutes. The maximum time range for a single query is 31 days. The period within which historical data is available is 366 days. The data latency is no more than 10 minutes.
//
// @param request - DescribeDcdnDomainIpaConnDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainIpaConnDataResponse
func (client *Client) DescribeDcdnDomainIpaConnDataWithContext(ctx context.Context, request *DescribeDcdnDomainIpaConnDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainIpaConnDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.SplitBy) {
		query["SplitBy"] = request.SplitBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainIpaConnData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainIpaConnDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries traffic of one or more accelerated domain names for which Layer 4 acceleration is enabled. You can query the data that is collected over the last 90 days.
//
// Description:
//
// >
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
//   - Unit: bytes.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainIpaTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainIpaTrafficDataResponse
func (client *Client) DescribeDcdnDomainIpaTrafficDataWithContext(ctx context.Context, request *DescribeDcdnDomainIpaTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainIpaTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FixTimeGap) {
		query["FixTimeGap"] = request.FixTimeGap
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TimeMerge) {
		query["TimeMerge"] = request.TimeMerge
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainIpaTrafficData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainIpaTrafficDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the proportions of data usage of different Internet service providers (ISPs). You can query data within the last 90 days.
//
// Description:
//
// >
//
//   - You can call this operation up to 100 times per second per account.
//
//   - If **StartTime*	- is set but **EndTime*	- is not set, the data within the hour that starts from **StartTime*	- is queried.
//
//   - If **EndTime*	- is set but **StartTime*	- is not set, the data within the last hour that precedes **EndTime*	- is queried.
//
//   - You can query data of a domain name or all domain names that belong to your account.
//
//   - You can view data that is collected over the last seven days. The interval at which data is queried is based on the time range specified by **StartTime*	- and **EndTime**.
//
//   - **If the time range is shorter than or equal to one hour**, data is queried every minute.
//
//   - **If the time range is longer than 1 hour but shorter than or equal to three days**, data is queried every five minutes.
//
//   - **If the time range is longer than three days but shorter than or equal to seven days**, data is queried every hour.
//
// @param request - DescribeDcdnDomainIspDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainIspDataResponse
func (client *Client) DescribeDcdnDomainIspDataWithContext(ctx context.Context, request *DescribeDcdnDomainIspDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainIspDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainIspData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainIspDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the address where you can download the log data of a domain name.
//
// Description:
//
// >
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.********
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainLogResponse
func (client *Client) DescribeDcdnDomainLogWithContext(ctx context.Context, request *DescribeDcdnDomainLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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
		Action:      dara.String("DescribeDcdnDomainLog"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDcdnDomainLogExTtlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainLogExTtlResponse
func (client *Client) DescribeDcdnDomainLogExTtlWithContext(ctx context.Context, request *DescribeDcdnDomainLogExTtlRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainLogExTtlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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
		Action:      dara.String("DescribeDcdnDomainLogExTtl"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainLogExTtlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the billable items of accelerated domain names. The data is collected at least every 5 minutes. The billable items do not include non-request items.
//
// Description:
//
//	  If you do not set the StartTime or EndTime parameter, data within the last 10 minutes is queried. You can set both the StartTime and EndTime parameters to specify a time range.
//
//		- You can specify one or more accelerated domain names. Separate domain names with commas (,).
//
//		- You can query data within the last 90 days.
//
//		- The time range cannot exceed 1 hour.
//
// @param request - DescribeDcdnDomainMultiUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainMultiUsageDataResponse
func (client *Client) DescribeDcdnDomainMultiUsageDataWithContext(ctx context.Context, request *DescribeDcdnDomainMultiUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainMultiUsageDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainMultiUsageData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainMultiUsageDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the origin bandwidth data for one or more accelerated domain names. You can query data in the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainOriginBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainOriginBpsDataResponse
func (client *Client) DescribeDcdnDomainOriginBpsDataWithContext(ctx context.Context, request *DescribeDcdnDomainOriginBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainOriginBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainOriginBpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainOriginBpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the back-to-origin traffic of one or more accelerated domain names.
//
// Description:
//
// - You can call this operation up to 100 times per second per account.
//
// - If you do not set the **StartTime*	- or **EndTime*	- parameters, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter varies with the maximum time range per query. The following table describes the time period within which historical data is available and the data delay.
//
// | Time granularity | Maximum time range per query | Historical data available | Data delay |
//
// | ---------------- | ---------------------------- | ------------------------- | ---------- |
//
// | 5 minutes | 3 days | 93 days | 15 minutes |
//
// | 1 hour | 31 days | 186 days | 4 hours |
//
// | 1 day | 366 days | 366 days | 04:00 on the next day |
//
// @param request - DescribeDcdnDomainOriginTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainOriginTrafficDataResponse
func (client *Client) DescribeDcdnDomainOriginTrafficDataWithContext(ctx context.Context, request *DescribeDcdnDomainOriginTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainOriginTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainOriginTrafficData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainOriginTrafficDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the protocol type of IP Application Accelerator (IPA).
//
// Description:
//
// > You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDcdnDomainPropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainPropertyResponse
func (client *Client) DescribeDcdnDomainPropertyWithContext(ctx context.Context, request *DescribeDcdnDomainPropertyRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainPropertyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainProperty"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainPropertyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries page view (PV) data of an accelerated domain name. Data can be collected at minimum intervals of one hour.
//
// @param request - DescribeDcdnDomainPvDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainPvDataResponse
func (client *Client) DescribeDcdnDomainPvDataWithContext(ctx context.Context, request *DescribeDcdnDomainPvDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainPvDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainPvData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainPvDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of requests to an accelerated domain name per second. You can query data in the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainQpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainQpsDataResponse
func (client *Client) DescribeDcdnDomainQpsDataWithContext(ctx context.Context, request *DescribeDcdnDomainQpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainQpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainQpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainQpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The number of queries per second in the Chinese mainland.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the time range to query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainQpsDataByLayerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainQpsDataByLayerResponse
func (client *Client) DescribeDcdnDomainQpsDataByLayerWithContext(ctx context.Context, request *DescribeDcdnDomainQpsDataByLayerRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainQpsDataByLayerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.Layer) {
		query["Layer"] = request.Layer
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainQpsDataByLayer"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainQpsDataByLayerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the real-time network bandwidth of a domain name.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not specify **StartTime*	- or **EndTime**, the request returns the data collected in the last hour by default. If you specify both parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|3 to 4 hours|
//
// @param request - DescribeDcdnDomainRealTimeBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeBpsDataResponse
func (client *Client) DescribeDcdnDomainRealTimeBpsDataWithContext(ctx context.Context, request *DescribeDcdnDomainRealTimeBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeBpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeBpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries byte hit ratios at a time granularity of 1 minute.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - The network traffic destined for different domain names may be redirected to the same origin server. Therefore, the byte hit ratios may be inaccurate. The accuracy of query results is based on the actual configurations.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last hour. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// | Time granularity | Maximum time range per query | Historical data available | Data delay |
//
// |----|------|-----|--------|
//
// | 1 minute | 1 hour | 7 days | 5 minutes |
//
// | 5 minutes | 3 days | 93 days | 15 minutes |
//
// | 1 hour | 31 days | 186 days | 4 hours |
//
// @param request - DescribeDcdnDomainRealTimeByteHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeByteHitRateDataResponse
func (client *Client) DescribeDcdnDomainRealTimeByteHitRateDataWithContext(ctx context.Context, request *DescribeDcdnDomainRealTimeByteHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeByteHitRateDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeByteHitRateData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeByteHitRateDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries traffic usage through each Internet service provider (ISP) and the number of visits in each region. The resolution of the data is one minute. The maximum time range to query for this operation is 10 minutes.
//
// Description:
//
// >
//
// > - You can call this operation up to 10 times per second per account.
//
// > - This operation is available only to users whose daily peak bandwidth value is higher than 1 Gbit/s. If you meet this requirement, you can [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket-intl) to apply for permissions to use this operation.
//
// @param request - DescribeDcdnDomainRealTimeDetailDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeDetailDataResponse
func (client *Client) DescribeDcdnDomainRealTimeDetailDataWithContext(ctx context.Context, request *DescribeDcdnDomainRealTimeDetailDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeDetailDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeDetailData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeDetailDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total number and proportions of HTTP status codes returned from one or more accelerated domain names.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeHttpCodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeHttpCodeDataResponse
func (client *Client) DescribeDcdnDomainRealTimeHttpCodeDataWithContext(ctx context.Context, request *DescribeDcdnDomainRealTimeHttpCodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeHttpCodeDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeHttpCodeData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeHttpCodeDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The number of QPS for one or more accelerated domain names is queried. Data is collected every minute.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not specify the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you specify both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeQpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeQpsDataResponse
func (client *Client) DescribeDcdnDomainRealTimeQpsDataWithContext(ctx context.Context, request *DescribeDcdnDomainRealTimeQpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeQpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeQpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeQpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request hit rates with a time granularity of 1 minute.
//
// Description:
//
// You can call this operation up to 10 times per second per user.
//
//   - The network traffic destined for different domain names may be redirected to the same origin server. Therefore, the byte hit ratios may be inaccurate. The accuracy of query results is based on the actual configurations.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last hour. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity*	- The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeReqHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeReqHitRateDataResponse
func (client *Client) DescribeDcdnDomainRealTimeReqHitRateDataWithContext(ctx context.Context, request *DescribeDcdnDomainRealTimeReqHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeReqHitRateDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeReqHitRateData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeReqHitRateDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bandwidth data of back-to-origin requests. Data is collected every minute. You can query data collected in the last 7 days.
//
// Description:
//
//	  You can call this operation up to 10 times per second per account.
//
//		- If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both the parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// | Time granularity | Maximum time range per query | Historical data available | Data delay |
//
// |-----|-----|-----|--------|
//
// | 1 minute | 1 hour | 7 days | 5 minutes |
//
// | 5 minutes | 3 days | 93 days | 15 minutes | | 1 hour | 31 days | 186 days | 4 hours |
//
// @param request - DescribeDcdnDomainRealTimeSrcBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeSrcBpsDataResponse
func (client *Client) DescribeDcdnDomainRealTimeSrcBpsDataWithContext(ctx context.Context, request *DescribeDcdnDomainRealTimeSrcBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeSrcBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeSrcBpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeSrcBpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the proportions of HTTP status codes based on back-to-origin statistics with a time granularity of one minute.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not specify the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you specify both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse
func (client *Client) DescribeDcdnDomainRealTimeSrcHttpCodeDataWithContext(ctx context.Context, request *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeSrcHttpCodeData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the origin traffic monitoring data for an accelerated domain name. Data is collected every minute. You can query data in the last 90 days.
//
// Description:
//
// If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeSrcTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeSrcTrafficDataResponse
func (client *Client) DescribeDcdnDomainRealTimeSrcTrafficDataWithContext(ctx context.Context, request *DescribeDcdnDomainRealTimeSrcTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeSrcTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeSrcTrafficData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeSrcTrafficDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic monitoring data of an accelerated domain name. Data is collected every minute.
//
// Description:
//
// You can call this operation up to 50 times per second per user.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeTrafficDataResponse
func (client *Client) DescribeDcdnDomainRealTimeTrafficDataWithContext(ctx context.Context, request *DescribeDcdnDomainRealTimeTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeTrafficData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeTrafficDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries regional distribution of users. Data is collected every day. You can query data within the last 90 days.
//
// Description:
//
// >
//
//   - If you do not specify the StartTime and EndTime parameters, the data within the last 24 hours is queried. If you specify the StartTime and EndTime parameters, the data within the specified time range is queried.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainRegionDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRegionDataResponse
func (client *Client) DescribeDcdnDomainRegionDataWithContext(ctx context.Context, request *DescribeDcdnDomainRegionDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRegionDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRegionData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRegionDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the environment configuration in the canary release environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnDomainStagingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainStagingConfigResponse
func (client *Client) DescribeDcdnDomainStagingConfigWithContext(ctx context.Context, request *DescribeDcdnDomainStagingConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainStagingConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionNames) {
		query["FunctionNames"] = request.FunctionNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainStagingConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainStagingConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries and sorts frequently requested web pages on a specified day. You can query data collected within the last 90 days.
//
// Description:
//
//	  If you do not set the StartTime parameter, the data on the previous day is queried.
//
//		- You can specify only one domain name.
//
// @param request - DescribeDcdnDomainTopReferVisitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainTopReferVisitResponse
func (client *Client) DescribeDcdnDomainTopReferVisitWithContext(ctx context.Context, request *DescribeDcdnDomainTopReferVisitRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainTopReferVisitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainTopReferVisit"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainTopReferVisitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries frequently requested URLs on a day.
//
// Description:
//
// > You can query data in the last seven days.
//
// @param request - DescribeDcdnDomainTopUrlVisitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainTopUrlVisitResponse
func (client *Client) DescribeDcdnDomainTopUrlVisitWithContext(ctx context.Context, request *DescribeDcdnDomainTopUrlVisitRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainTopUrlVisitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainTopUrlVisit"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainTopUrlVisitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the network traffic of accelerated domain names. You can query data collected in the last 90 days.
//
// Description:
//
// If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
//   - You can call this operation up to 100 times per second per account.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainTrafficDataResponse
func (client *Client) DescribeDcdnDomainTrafficDataWithContext(ctx context.Context, request *DescribeDcdnDomainTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainTrafficData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainTrafficDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resource usage about domain names in a billable region.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - Usage data includes traffic (measured in bytes), bandwidth values (measured in bit/s), and the number of requests.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainUsageDataResponse
func (client *Client) DescribeDcdnDomainUsageDataWithContext(ctx context.Context, request *DescribeDcdnDomainUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainUsageDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.DataProtocol) {
		query["DataProtocol"] = request.DataProtocol
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Field) {
		query["Field"] = request.Field
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
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
		Action:      dara.String("DescribeDcdnDomainUsageData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainUsageDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of unique visitors (UVs) to an accelerated domain name. Data is collected every hour. You can query data within the last 90 days.
//
// Description:
//
//	  If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
//		- You can specify only one accelerated domain name or all the accelerated domain names that belong to your Alibaba Cloud account.
//
// @param request - DescribeDcdnDomainUvDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainUvDataResponse
func (client *Client) DescribeDcdnDomainUvDataWithContext(ctx context.Context, request *DescribeDcdnDomainUvDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainUvDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainUvData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainUvDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries bandwidth of one or more accelerated domain names for which WebSocket is enabled. You can query the data that is collected over the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainWebsocketBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainWebsocketBpsDataResponse
func (client *Client) DescribeDcdnDomainWebsocketBpsDataWithContext(ctx context.Context, request *DescribeDcdnDomainWebsocketBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainWebsocketBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainWebsocketBpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainWebsocketBpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The total number and proportions of HTTP status codes returned from one or more accelerated domain names for which WebSocket is enabled are queried. Data can be collected at minimum intervals of 5 minutes.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the time range to query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainWebsocketHttpCodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainWebsocketHttpCodeDataResponse
func (client *Client) DescribeDcdnDomainWebsocketHttpCodeDataWithContext(ctx context.Context, request *DescribeDcdnDomainWebsocketHttpCodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainWebsocketHttpCodeDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainWebsocketHttpCodeData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainWebsocketHttpCodeDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic monitoring data for an accelerated domain name with WebSocket enabled. You can query data in the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainWebsocketTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainWebsocketTrafficDataResponse
func (client *Client) DescribeDcdnDomainWebsocketTrafficDataWithContext(ctx context.Context, request *DescribeDcdnDomainWebsocketTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainWebsocketTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainWebsocketTrafficData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainWebsocketTrafficDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries DCDN-accelerated domain names by origin server.
//
// @param request - DescribeDcdnDomainsBySourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainsBySourceResponse
func (client *Client) DescribeDcdnDomainsBySourceWithContext(ctx context.Context, request *DescribeDcdnDomainsBySourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainsBySourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainsBySource"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainsBySourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of times that a routine is executed within a specified period of time.
//
// Description:
//
//	  You can call this operation up to 10 times per second per account.
//
//		- The minimum time granularity for a query is 1 hour. The maximum time span for a query is 24 hours. The time period within which historical data is available for a query is 366 days.
//
// @param request - DescribeDcdnErUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnErUsageDataResponse
func (client *Client) DescribeDcdnErUsageDataWithContext(ctx context.Context, request *DescribeDcdnErUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnErUsageDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.RoutineID) {
		query["RoutineID"] = request.RoutineID
	}

	if !dara.IsNil(request.Spec) {
		query["Spec"] = request.Spec
	}

	if !dara.IsNil(request.SplitBy) {
		query["SplitBy"] = request.SplitBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnErUsageData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnErUsageDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of blocked IP addresses.
//
// Description:
//
// > 	- If you specify IP addresses or CIDR blocks, IP addresses that are effective and the corresponding expiration time are returned. If you do not specify IP addresses or CIDR blocks, all effective IP addresses and the corresponding expiration time are returned.
//
// > 	- The results are written to OSS and returned as OSS URLs. The content in OSS objects is in the format of `IP address-Corresponding expiration time`. The expiration time is in the YYYY-MM-DD hh:mm:ss format.
//
// > 	- You can share OSS URLs with others. The shared URLs are valid for three days.
//
// @param request - DescribeDcdnFullDomainsBlockIPConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnFullDomainsBlockIPConfigResponse
func (client *Client) DescribeDcdnFullDomainsBlockIPConfigWithContext(ctx context.Context, request *DescribeDcdnFullDomainsBlockIPConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnFullDomainsBlockIPConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnFullDomainsBlockIPConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnFullDomainsBlockIPConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户海量封禁历史
//
// Description:
//
//	  For a specified IP addresses and time range, the time when the IP address was delivered to the edge and the corresponding result are returned.
//
//		- If a specified IP address or CIDR block has multiple blocking records in a specified time range, the records are sorted by delivery time in descending order.
//
//		- The maximum time range to query is 90 days.
//
//		- If no blocking record exists or delivery fails for the given IP address and time range, the delivery time is empty.
//
// @param request - DescribeDcdnFullDomainsBlockIPHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnFullDomainsBlockIPHistoryResponse
func (client *Client) DescribeDcdnFullDomainsBlockIPHistoryWithContext(ctx context.Context, request *DescribeDcdnFullDomainsBlockIPHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnFullDomainsBlockIPHistoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IPList) {
		body["IPList"] = request.IPList
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnFullDomainsBlockIPHistory"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnFullDomainsBlockIPHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about all certificates that belong to your account.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnHttpsDomainListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnHttpsDomainListResponse
func (client *Client) DescribeDcdnHttpsDomainListWithContext(ctx context.Context, request *DescribeDcdnHttpsDomainListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnHttpsDomainListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
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
		Action:      dara.String("DescribeDcdnHttpsDomainList"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnHttpsDomainListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether an IP address belongs to a POP.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnIpInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnIpInfoResponse
func (client *Client) DescribeDcdnIpInfoWithContext(ctx context.Context, request *DescribeDcdnIpInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnIpInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IP) {
		query["IP"] = request.IP
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnIpInfo"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnIpInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the back-to-origin CIDR blocks of IPA-accelerated domain names. If you want to call this API operation, you must submit a ticket to apply for the required permissions.
//
// Description:
//
// >  This operation can be called globally up to 50 times per second. This operation can be called up to 10 times per second per account.
//
// @param request - DescribeDcdnIpaDomainCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnIpaDomainCidrResponse
func (client *Client) DescribeDcdnIpaDomainCidrWithContext(ctx context.Context, request *DescribeDcdnIpaDomainCidrRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnIpaDomainCidrResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnIpaDomainCidr"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnIpaDomainCidrResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of an accelerated domain name. You can query the configurations of one or more features in each request.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnIpaDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnIpaDomainConfigsResponse
func (client *Client) DescribeDcdnIpaDomainConfigsWithContext(ctx context.Context, request *DescribeDcdnIpaDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnIpaDomainConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionNames) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnIpaDomainConfigs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnIpaDomainConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic configuration information about an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnIpaDomainDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnIpaDomainDetailResponse
func (client *Client) DescribeDcdnIpaDomainDetailWithContext(ctx context.Context, request *DescribeDcdnIpaDomainDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnIpaDomainDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnIpaDomainDetail"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnIpaDomainDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of IPA. The information includes the time when the service was activated, the current service status, the current billing method, and the billing method of the next cycle.
//
// Description:
//
// *
//
// **The maximum number of times that each user can call this operation per second is 20.
//
// @param request - DescribeDcdnIpaServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnIpaServiceResponse
func (client *Client) DescribeDcdnIpaServiceWithContext(ctx context.Context, request *DescribeDcdnIpaServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnIpaServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnIpaService"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnIpaServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about all domain names that are accelerated by IP Application Accelerator (IPA) in your account. Fuzzy search and filtering by domain status are supported.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnIpaUserDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnIpaUserDomainsResponse
func (client *Client) DescribeDcdnIpaUserDomainsWithContext(ctx context.Context, request *DescribeDcdnIpaUserDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnIpaUserDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckDomainShow) {
		query["CheckDomainShow"] = request.CheckDomainShow
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainSearchType) {
		query["DomainSearchType"] = request.DomainSearchType
	}

	if !dara.IsNil(request.DomainStatus) {
		query["DomainStatus"] = request.DomainStatus
	}

	if !dara.IsNil(request.FuncFilter) {
		query["FuncFilter"] = request.FuncFilter
	}

	if !dara.IsNil(request.FuncId) {
		query["FuncId"] = request.FuncId
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnIpaUserDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnIpaUserDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a namespace.
//
// @param request - DescribeDcdnKvNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnKvNamespaceResponse
func (client *Client) DescribeDcdnKvNamespaceWithContext(ctx context.Context, request *DescribeDcdnKvNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnKvNamespaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnKvNamespace"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnKvNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the origin CIDR blocks by domain name. The CIDR blocks include IPv4 and IPv6 CIDR blocks.
//
// @param request - DescribeDcdnL2VipsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnL2VipsResponse
func (client *Client) DescribeDcdnL2VipsWithContext(ctx context.Context, request *DescribeDcdnL2VipsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnL2VipsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnL2Vips"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnL2VipsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the origin server for a DCDN-accelerated domain name.
//
// @param request - DescribeDcdnOriginSiteHealthStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnOriginSiteHealthStatusResponse
func (client *Client) DescribeDcdnOriginSiteHealthStatusWithContext(ctx context.Context, request *DescribeDcdnOriginSiteHealthStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnOriginSiteHealthStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnOriginSiteHealthStatus"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnOriginSiteHealthStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the fields in real-time log entries.
//
// Description:
//
// >  You can call this API operation up to 100 times per second per account.
//
// @param request - DescribeDcdnRealTimeDeliveryFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnRealTimeDeliveryFieldResponse
func (client *Client) DescribeDcdnRealTimeDeliveryFieldWithContext(ctx context.Context, request *DescribeDcdnRealTimeDeliveryFieldRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnRealTimeDeliveryFieldResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnRealTimeDeliveryField"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnRealTimeDeliveryFieldResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the maximum number and the remaining number of URLs and directories that can be refreshed or the maximum number and the remaining number of URLs that can be prefetched per day.
//
// Description:
//
// >
//
//   - You can call the **RefreshDcdnObjectCaches*	- operation to refresh content and call the **PreloadDcdnObjectCaches*	- operation to prefetch content.
//
//   - You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnRefreshQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnRefreshQuotaResponse
func (client *Client) DescribeDcdnRefreshQuotaWithContext(ctx context.Context, request *DescribeDcdnRefreshQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnRefreshQuotaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnRefreshQuota"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnRefreshQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of refresh or prefetch tasks by task ID.
//
// Description:
//
// >
//
//   - You can query data within the last three days.
//
//   - You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnRefreshTaskByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnRefreshTaskByIdResponse
func (client *Client) DescribeDcdnRefreshTaskByIdWithContext(ctx context.Context, request *DescribeDcdnRefreshTaskByIdRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnRefreshTaskByIdResponse, _err error) {
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
		Action:      dara.String("DescribeDcdnRefreshTaskById"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnRefreshTaskByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the refresh or prefetch tasks. You can query the tasks in the last three days.
//
// Description:
//
//	  You can query the refresh or prefetch tasks by ID or URL.
//
//		- You can set both **TaskId*	- and **ObjectPath*	- in a request. If you do not set **TaskId*	- or **ObjectPath**, the data in the last 3 days on the first page is returned. By default, a maximum of 20 entries can be displayed on each page.
//
//		- If you specify **DomainName*	- or **Status**, you must also specify **ObjectType**.
//
//		- You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDcdnRefreshTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnRefreshTasksResponse
func (client *Client) DescribeDcdnRefreshTasksWithContext(ctx context.Context, request *DescribeDcdnRefreshTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnRefreshTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ObjectPath) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnRefreshTasks"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnRefreshTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of Internet service providers (ISPs) and regions.
//
// Description:
//
// >  You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnRegionAndIspRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnRegionAndIspResponse
func (client *Client) DescribeDcdnRegionAndIspWithContext(ctx context.Context, request *DescribeDcdnRegionAndIspRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnRegionAndIspResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnRegionAndIsp"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnRegionAndIspResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the content of an operations report.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - DescribeDcdnReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnReportResponse
func (client *Client) DescribeDcdnReportWithContext(ctx context.Context, request *DescribeDcdnReportRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.HttpCode) {
		query["HttpCode"] = request.HttpCode
	}

	if !dara.IsNil(request.IsOverseas) {
		query["IsOverseas"] = request.IsOverseas
	}

	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnReport"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom operations reports.
//
// Description:
//
// > 	- This operation queries the metadata of all operations reports. The statistics in the reports are not returned.
//
// > 	- You can call this operation up to three times per second per account.
//
// @param request - DescribeDcdnReportListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnReportListResponse
func (client *Client) DescribeDcdnReportListWithContext(ctx context.Context, request *DescribeDcdnReportListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnReportListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnReportList"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnReportListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a real-time log delivery project.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnSLSRealtimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnSLSRealtimeLogDeliveryResponse
func (client *Client) DescribeDcdnSLSRealtimeLogDeliveryWithContext(ctx context.Context, request *DescribeDcdnSLSRealtimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnSLSRealtimeLogDeliveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnSLSRealtimeLogDelivery"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnSLSRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a ShangMi (SM) certificate.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnSMCertificateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnSMCertificateDetailResponse
func (client *Client) DescribeDcdnSMCertificateDetailWithContext(ctx context.Context, request *DescribeDcdnSMCertificateDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnSMCertificateDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnSMCertificateDetail"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnSMCertificateDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ShangMi (SM) certificates of an accelerated domain name.
//
// Description:
//
// >  You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnSMCertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnSMCertificateListResponse
func (client *Client) DescribeDcdnSMCertificateListWithContext(ctx context.Context, request *DescribeDcdnSMCertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnSMCertificateListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnSMCertificateList"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnSMCertificateListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificates of accelerated domain names.
//
// @param request - DescribeDcdnSSLCertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnSSLCertificateListResponse
func (client *Client) DescribeDcdnSSLCertificateListWithContext(ctx context.Context, request *DescribeDcdnSSLCertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnSSLCertificateListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
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

	if !dara.IsNil(request.SearchKeyword) {
		query["SearchKeyword"] = request.SearchKeyword
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnSSLCertificateList"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnSSLCertificateListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an edge security drop-down list in the Dynamic Content Delivery Network (DCDN) console.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnSecFuncInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnSecFuncInfoResponse
func (client *Client) DescribeDcdnSecFuncInfoWithContext(ctx context.Context, request *DescribeDcdnSecFuncInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnSecFuncInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SecFuncType) {
		query["SecFuncType"] = request.SecFuncType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnSecFuncInfo"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnSecFuncInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the Dynamic Content Delivery Network (DCDN) service. The information includes the time when the service was activated, the current service status, the current billing method, and the billing method of the next cycle.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnServiceResponse
func (client *Client) DescribeDcdnServiceWithContext(ctx context.Context, request *DescribeDcdnServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnService"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags of one or more resources.
//
// Description:
//
// > You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDcdnTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnTagResourcesResponse
func (client *Client) DescribeDcdnTagResourcesWithContext(ctx context.Context, request *DescribeDcdnTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
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
		Action:      dara.String("DescribeDcdnTagResources"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnTagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries domain names ranked by network traffic. You can query data within the last 90 days.
//
// Description:
//
// If you do not specify the StartTime and EndTime parameters, the data within the current month is queried. If you specify the StartTime and EndTime parameters, the data within the specified time range is queried.
//
// @param request - DescribeDcdnTopDomainsByFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnTopDomainsByFlowResponse
func (client *Client) DescribeDcdnTopDomainsByFlowWithContext(ctx context.Context, request *DescribeDcdnTopDomainsByFlowRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnTopDomainsByFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnTopDomainsByFlow"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnTopDomainsByFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the billing records of an Alibaba Cloud account. The maximum time range that you can specify is one month.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnUserBillHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserBillHistoryResponse
func (client *Client) DescribeDcdnUserBillHistoryWithContext(ctx context.Context, request *DescribeDcdnUserBillHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserBillHistoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserBillHistory"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserBillHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metering method that is used in Dynamic Content Delivery Network (DCDN).
//
// @param request - DescribeDcdnUserBillTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserBillTypeResponse
func (client *Client) DescribeDcdnUserBillTypeWithContext(ctx context.Context, request *DescribeDcdnUserBillTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserBillTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserBillType"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserBillTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of security features.
//
// Description:
//
// You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnUserConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserConfigsResponse
func (client *Client) DescribeDcdnUserConfigsWithContext(ctx context.Context, request *DescribeDcdnUserConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserConfigs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the accelerated domain names that belong to your Alibaba Cloud account. You can filter domains by name or by status. Fuzzy match is supported when you filter domains by name.
//
// Description:
//
// > You can call this operation up to 80 times per second per account.
//
// @param request - DescribeDcdnUserDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserDomainsResponse
func (client *Client) DescribeDcdnUserDomainsWithContext(ctx context.Context, request *DescribeDcdnUserDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeEndTime) {
		query["ChangeEndTime"] = request.ChangeEndTime
	}

	if !dara.IsNil(request.ChangeStartTime) {
		query["ChangeStartTime"] = request.ChangeStartTime
	}

	if !dara.IsNil(request.CheckDomainShow) {
		query["CheckDomainShow"] = request.CheckDomainShow
	}

	if !dara.IsNil(request.Coverage) {
		query["Coverage"] = request.Coverage
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainSearchType) {
		query["DomainSearchType"] = request.DomainSearchType
	}

	if !dara.IsNil(request.DomainStatus) {
		query["DomainStatus"] = request.DomainStatus
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.WebSiteType) {
		query["WebSiteType"] = request.WebSiteType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all accelerated domain names with specified features configured that belong to your Alibaba Cloud account based on the FuncId parameter.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnUserDomainsByFuncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserDomainsByFuncResponse
func (client *Client) DescribeDcdnUserDomainsByFuncWithContext(ctx context.Context, request *DescribeDcdnUserDomainsByFuncRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserDomainsByFuncResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FuncFilter) {
		query["FuncFilter"] = request.FuncFilter
	}

	if !dara.IsNil(request.FuncId) {
		query["FuncId"] = request.FuncId
	}

	if !dara.IsNil(request.MatchType) {
		query["MatchType"] = request.MatchType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserDomainsByFunc"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserDomainsByFuncResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource quotas and the used resources.
//
// Description:
//
// >  The maximum number of times that each user can call this operation per second is 30.
//
// @param request - DescribeDcdnUserQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserQuotaResponse
func (client *Client) DescribeDcdnUserQuotaWithContext(ctx context.Context, request *DescribeDcdnUserQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserQuotaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserQuota"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the fields that are selected.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnUserRealTimeDeliveryFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserRealTimeDeliveryFieldResponse
func (client *Client) DescribeDcdnUserRealTimeDeliveryFieldWithContext(ctx context.Context, request *DescribeDcdnUserRealTimeDeliveryFieldRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserRealTimeDeliveryFieldResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserRealTimeDeliveryField"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserRealTimeDeliveryFieldResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the resource plans in your Alibaba Cloud account.
//
// Description:
//
// The maximum number of times that each user can call this operation per second is 30.
//
// @param request - DescribeDcdnUserResourcePackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserResourcePackageResponse
func (client *Client) DescribeDcdnUserResourcePackageWithContext(ctx context.Context, request *DescribeDcdnUserResourcePackageRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserResourcePackageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeDcdnUserResourcePackage"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserResourcePackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of packets blocked by a specified security feature.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnUserSecDropRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserSecDropResponse
func (client *Client) DescribeDcdnUserSecDropWithContext(ctx context.Context, request *DescribeDcdnUserSecDropRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserSecDropResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		query["Data"] = request.Data
	}

	if !dara.IsNil(request.Metric) {
		query["Metric"] = request.Metric
	}

	if !dara.IsNil(request.SecFunc) {
		query["SecFunc"] = request.SecFunc
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserSecDrop"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserSecDropResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of packets that are blocked by security features at the application layer in a specific time range.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnUserSecDropByMinuteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserSecDropByMinuteResponse
func (client *Client) DescribeDcdnUserSecDropByMinuteWithContext(ctx context.Context, request *DescribeDcdnUserSecDropByMinuteRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserSecDropByMinuteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Object) {
		query["Object"] = request.Object
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SecFunc) {
		query["SecFunc"] = request.SecFunc
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserSecDropByMinute"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserSecDropByMinuteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries virtual IP addresses of the POPs by domain name.
//
// Description:
//
// >  You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnUserVipsByDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserVipsByDomainResponse
func (client *Client) DescribeDcdnUserVipsByDomainWithContext(ctx context.Context, request *DescribeDcdnUserVipsByDomainRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserVipsByDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Available) {
		query["Available"] = request.Available
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserVipsByDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserVipsByDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ownership verification content of a domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnVerifyContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnVerifyContentResponse
func (client *Client) DescribeDcdnVerifyContentWithContext(ctx context.Context, request *DescribeDcdnVerifyContentRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnVerifyContentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnVerifyContent"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnVerifyContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the default configurations of a WAF rule.
//
// @param request - DescribeDcdnWafDefaultRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafDefaultRulesResponse
func (client *Client) DescribeDcdnWafDefaultRulesWithContext(ctx context.Context, request *DescribeDcdnWafDefaultRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafDefaultRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QueryArgs) {
		query["QueryArgs"] = request.QueryArgs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafDefaultRules"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafDefaultRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries domain names that have Web Application Firewall (WAF) enabled and the relevant information, including the status of the access control list (ACL), protection against HTTP flood attacks, domain name, and WAF.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnWafDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafDomainResponse
func (client *Client) DescribeDcdnWafDomainWithContext(ctx context.Context, request *DescribeDcdnWafDomainRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
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
		Action:      dara.String("DescribeDcdnWafDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the protection policy of a domain name.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafDomainDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafDomainDetailResponse
func (client *Client) DescribeDcdnWafDomainDetailWithContext(ctx context.Context, request *DescribeDcdnWafDomainDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafDomainDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafDomainDetail"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafDomainDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the accelerated domain names that are protected by Web Application Firewall (WAF). Fuzzy search is supported.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafDomainsResponse
func (client *Client) DescribeDcdnWafDomainsWithContext(ctx context.Context, request *DescribeDcdnWafDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryArgs) {
		query["QueryArgs"] = request.QueryArgs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about match conditions in a custom protection rule, such as the match fields, logical characters, and match content.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafFilterInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafFilterInfoResponse
func (client *Client) DescribeDcdnWafFilterInfoWithContext(ctx context.Context, request *DescribeDcdnWafFilterInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafFilterInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefenseScenes) {
		query["DefenseScenes"] = request.DefenseScenes
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafFilterInfo"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafFilterInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the countries and regions that can be added to the blacklist of Web Application Firewall (WAF).
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafGeoInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafGeoInfoResponse
func (client *Client) DescribeDcdnWafGeoInfoWithContext(ctx context.Context, request *DescribeDcdnWafGeoInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafGeoInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafGeoInfo"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafGeoInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a custom WAF rule group by page.
//
// @param request - DescribeDcdnWafGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafGroupResponse
func (client *Client) DescribeDcdnWafGroupWithContext(ctx context.Context, request *DescribeDcdnWafGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryArgs) {
		query["QueryArgs"] = request.QueryArgs
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafGroup"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom Web Application Firewall (WAF) rule groups.
//
// @param request - DescribeDcdnWafGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafGroupsResponse
func (client *Client) DescribeDcdnWafGroupsWithContext(ctx context.Context, request *DescribeDcdnWafGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryArgs) {
		query["QueryArgs"] = request.QueryArgs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafGroups"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the address from which you can download the Web Application Firewall (WAF) logs of a domain name.
//
// Description:
//
// >
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
//   - The log data is collected every hour.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnWafLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafLogsResponse
func (client *Client) DescribeDcdnWafLogsWithContext(ctx context.Context, request *DescribeDcdnWafLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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
		Action:      dara.String("DescribeDcdnWafLogs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the Web Application Firewall (WAF) protection policies that you configured.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafPoliciesResponse
func (client *Client) DescribeDcdnWafPoliciesWithContext(ctx context.Context, request *DescribeDcdnWafPoliciesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafPoliciesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryArgs) {
		query["QueryArgs"] = request.QueryArgs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafPolicies"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a protection policy.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafPolicyResponse
func (client *Client) DescribeDcdnWafPolicyWithContext(ctx context.Context, request *DescribeDcdnWafPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafPolicy"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the accelerated domain names that are protected by a specified Web Application Firewall (WAF) protection policy.
//
// Description:
//
// You can call this operation up to 20 times per second per user.
//
// @param request - DescribeDcdnWafPolicyDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafPolicyDomainsResponse
func (client *Client) DescribeDcdnWafPolicyDomainsWithContext(ctx context.Context, request *DescribeDcdnWafPolicyDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafPolicyDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafPolicyDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafPolicyDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the domain names that can be bound to a custom protection policy.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafPolicyValidDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafPolicyValidDomainsResponse
func (client *Client) DescribeDcdnWafPolicyValidDomainsWithContext(ctx context.Context, request *DescribeDcdnWafPolicyValidDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafPolicyValidDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefenseScene) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !dara.IsNil(request.DomainNameLike) {
		query["DomainNameLike"] = request.DomainNameLike
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
		Action:      dara.String("DescribeDcdnWafPolicyValidDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafPolicyValidDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified protection rule.
//
// Description:
//
// #
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafRuleResponse
func (client *Client) DescribeDcdnWafRuleWithContext(ctx context.Context, request *DescribeDcdnWafRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafRule"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the protection rules that you configured.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafRulesResponse
func (client *Client) DescribeDcdnWafRulesWithContext(ctx context.Context, request *DescribeDcdnWafRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryArgs) {
		query["QueryArgs"] = request.QueryArgs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafRules"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the type of the protection policy that you use.
//
// Description:
//
// You can call this operation up to 20 times per second per user.
//
// @param request - DescribeDcdnWafScenesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafScenesResponse
func (client *Client) DescribeDcdnWafScenesWithContext(ctx context.Context, request *DescribeDcdnWafScenesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafScenesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefenseScenes) {
		query["DefenseScenes"] = request.DefenseScenes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafScenes"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafScenesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about Dynamic Content Delivery Network (DCDN) Web Application Firewall WAF), including the time when WAF is enabled, edition of WAF, current status of WAF, metering method for requests, and metering method for rules.
//
// Description:
//
// # Usage notes
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafServiceResponse
func (client *Client) DescribeDcdnWafServiceWithContext(ctx context.Context, request *DescribeDcdnWafServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafService"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The number of used SeCUs.
//
// Description:
//
//	  You can call this operation up to 10 times per second per account.
//
//		- The minimum time granularity for a query is 5 minutes. The maximum time span for a query is 31 days. The time period within which historical data is available for a query is 90 days.
//
// @param request - DescribeDcdnWafUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafUsageDataResponse
func (client *Client) DescribeDcdnWafUsageDataWithContext(ctx context.Context, request *DescribeDcdnWafUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafUsageDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.SplitBy) {
		query["SplitBy"] = request.SplitBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafUsageData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafUsageDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about Dynamic Content Delivery Network (DCDN), such as the service activation time, the expiration time, and the current status.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnsecServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnsecServiceResponse
func (client *Client) DescribeDcdnsecServiceWithContext(ctx context.Context, request *DescribeDcdnsecServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnsecServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnsecService"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnsecServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries attack events.
//
// @param request - DescribeDdosAllEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDdosAllEventListResponse
func (client *Client) DescribeDdosAllEventListWithContext(ctx context.Context, request *DescribeDdosAllEventListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDdosAllEventListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
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
		Action:      dara.String("DescribeDdosAllEventList"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDdosAllEventListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the highlighted data of attack details. You can query the reasons for which requests are blocked based on TraceIDs in logs of requests that are blocked by Basic Web Protection. The highlighted data matches the content blocked by the basic web protection module.
//
// @param request - DescribeHighlightInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHighlightInfoResponse
func (client *Client) DescribeHighlightInfoWithContext(ctx context.Context, request *DescribeHighlightInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeHighlightInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TraceId) {
		query["TraceId"] = request.TraceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHighlightInfo"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHighlightInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// kv存储实时Qps监控数据
//
// @param request - DescribeKvRealTimeQpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKvRealTimeQpsDataResponse
func (client *Client) DescribeKvRealTimeQpsDataWithContext(ctx context.Context, request *DescribeKvRealTimeQpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeKvRealTimeQpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessType) {
		query["AccessType"] = request.AccessType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SplitBy) {
		query["SplitBy"] = request.SplitBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeKvRealTimeQpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKvRealTimeQpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage data of KV storage.
//
// Description:
//
// *Note**	- You can call this operation up to 5 times per second per account.
//
//   - The usage data indicates the number of requests.
//
// **Time granularity:*	- This operation supports only the time granularity of 1 hour.
//
// |Time granularity|Time range to query|Historical data available|Data latency|
//
// |---|---|---|---|
//
// |1 hour|31 days|90 days|3 to 4 hours|
//
// @param request - DescribeKvUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKvUsageDataResponse
func (client *Client) DescribeKvUsageDataWithContext(ctx context.Context, request *DescribeKvUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeKvUsageDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessType) {
		query["AccessType"] = request.AccessType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Field) {
		query["Field"] = request.Field
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.ResponseType) {
		query["ResponseType"] = request.ResponseType
	}

	if !dara.IsNil(request.SplitBy) {
		query["SplitBy"] = request.SplitBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeKvUsageData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKvUsageDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the feature configurations of an accelerated domain name in the resource directory.
//
// @param request - DescribeRDDomainConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRDDomainConfigResponse
func (client *Client) DescribeRDDomainConfigWithContext(ctx context.Context, request *DescribeRDDomainConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeRDDomainConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRDDomainConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRDDomainConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all domain names of Alibaba Cloud CDN and Dynamic Content Delivery Network (DCDN) in a Resource Directory (RD).
//
// Description:
//
// A domain name can be in one of the following states:
//
//   - online
//
//   - offline
//
//   - configuring
//
//   - configure_failed
//
//   - checking
//
//   - check_failed
//
// @param request - DescribeRDDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRDDomainsResponse
func (client *Client) DescribeRDDomainsWithContext(ctx context.Context, request *DescribeRDDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRDDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DescribeRDDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRDDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metadata of a specified routine. The metadata includes the routine configuration, configuration version, and code version.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeRoutineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRoutineResponse
func (client *Client) DescribeRoutineWithContext(ctx context.Context, request *DescribeRoutineRequest, runtime *dara.RuntimeOptions) (_result *DescribeRoutineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRoutine"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRoutineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the JavaScript code version of a routine.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeRoutineCodeRevisionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRoutineCodeRevisionResponse
func (client *Client) DescribeRoutineCodeRevisionWithContext(ctx context.Context, request *DescribeRoutineCodeRevisionRequest, runtime *dara.RuntimeOptions) (_result *DescribeRoutineCodeRevisionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SelectCodeRevision) {
		body["SelectCodeRevision"] = request.SelectCodeRevision
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRoutineCodeRevision"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRoutineCodeRevisionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of domain names that are associated with a routine.
//
// @param request - DescribeRoutineRelatedDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRoutineRelatedDomainsResponse
func (client *Client) DescribeRoutineRelatedDomainsWithContext(ctx context.Context, request *DescribeRoutineRelatedDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRoutineRelatedDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRoutineRelatedDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRoutineRelatedDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Whether IPA is enabled and whether you have overdue payments for your IPA are queried.
//
// Description:
//
// *
//
// **The maximum number of times that each user can call this operation per second is 20.
//
// @param request - DescribeUserDcdnIpaStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserDcdnIpaStatusResponse
func (client *Client) DescribeUserDcdnIpaStatusWithContext(ctx context.Context, request *DescribeUserDcdnIpaStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserDcdnIpaStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserDcdnIpaStatus"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserDcdnIpaStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether DCDN is activated and whether your account has overdue payments.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeUserDcdnStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserDcdnStatusResponse
func (client *Client) DescribeUserDcdnStatusWithContext(ctx context.Context, request *DescribeUserDcdnStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserDcdnStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserDcdnStatus"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserDcdnStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether EdgeRoutine (ER) is activated or has an overdue payment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeUserErStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserErStatusResponse
func (client *Client) DescribeUserErStatusWithContext(ctx context.Context, request *DescribeUserErStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserErStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserErStatus"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserErStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether Log Service is activated and whether you have overdue payments for your Log Service.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeUserLogserviceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserLogserviceStatusResponse
func (client *Client) DescribeUserLogserviceStatusWithContext(ctx context.Context, request *DescribeUserLogserviceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserLogserviceStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserLogserviceStatus"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserLogserviceStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a routine.
//
// Description:
//
// >
//
//   - This operation modifies only the specified configurations. Other configurations remain unchanged.
//
//   - If you want to delete a setting, delete the parameter value.
//
//   - This operation can add canary release environments. Make sure that the environment names comply with the naming rules. Otherwise, you will fail to add the environments.
//
//   - Dynamic Route for CDN (DCDN) provides 35 canary release environments. Among these environments, 34 are deployed in China and 1 is deployed outside China. The canary release environments are:
//
//   - Outside China: presetCanaryOverseas.
//
//   - In China: The 34 canary release environments are named in the format of presetCanaryXX. For example, presetCanaryBeijing represents the canary release environment in Beijing. A canary release environment is in each of the following regions: Anhui, Beijing, Chongqing, Fujian, Gansu, Guangdong, Guangxi, Guizhou, Hainan, Hebei, Heilongjiang, Henan, Hong Kong, Hubei, Hunan, Jiangsu, Jiangxi, Jilin, Liaoning, Macao, Neimenggu, Ningxia, Qinghai, Shaanxi, Shandong, Shanghai, Shanxi, Sichuan, Taiwan, Tianjin, Xinjiang, Xizang, Yunan, and Zhejiang.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param tmpReq - EditRoutineConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditRoutineConfResponse
func (client *Client) EditRoutineConfWithContext(ctx context.Context, tmpReq *EditRoutineConfRequest, runtime *dara.RuntimeOptions) (_result *EditRoutineConfResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &EditRoutineConfShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EnvConf) {
		request.EnvConfShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EnvConf, dara.String("EnvConf"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnvConfShrink) {
		body["EnvConf"] = request.EnvConfShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditRoutineConf"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditRoutineConfResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the value of a key in a key-value pair.
//
// @param request - GetDcdnKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDcdnKvResponse
func (client *Client) GetDcdnKvWithContext(ctx context.Context, request *GetDcdnKvRequest, runtime *dara.RuntimeOptions) (_result *GetDcdnKvResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDcdnKv"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDcdnKvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询KV对的值以及TTL信息
//
// @param request - GetDcdnKvDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDcdnKvDetailResponse
func (client *Client) GetDcdnKvDetailWithContext(ctx context.Context, request *GetDcdnKvDetailRequest, runtime *dara.RuntimeOptions) (_result *GetDcdnKvDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDcdnKvDetail"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDcdnKvDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the KV status by key value.
//
// @param request - GetDcdnKvStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDcdnKvStatusResponse
func (client *Client) GetDcdnKvStatusWithContext(ctx context.Context, request *GetDcdnKvStatusRequest, runtime *dara.RuntimeOptions) (_result *GetDcdnKvStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDcdnKvStatus"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDcdnKvStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Traverses the values of keys in a namespace.
//
// @param request - ListDcdnKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDcdnKvResponse
func (client *Client) ListDcdnKvWithContext(ctx context.Context, request *ListDcdnKvRequest, runtime *dara.RuntimeOptions) (_result *ListDcdnKvResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDcdnKv"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDcdnKvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a real-time log delivery project.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ListDcdnRealTimeDeliveryProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDcdnRealTimeDeliveryProjectResponse
func (client *Client) ListDcdnRealTimeDeliveryProjectWithContext(ctx context.Context, request *ListDcdnRealTimeDeliveryProjectRequest, runtime *dara.RuntimeOptions) (_result *ListDcdnRealTimeDeliveryProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
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
		Action:      dara.String("ListDcdnRealTimeDeliveryProject"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDcdnRealTimeDeliveryProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # A客户定制新增修改域名采样率接口
//
// @param request - ModifyCustomDomainSampleRateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCustomDomainSampleRateResponse
func (client *Client) ModifyCustomDomainSampleRateWithContext(ctx context.Context, request *ModifyCustomDomainSampleRateRequest, runtime *dara.RuntimeOptions) (_result *ModifyCustomDomainSampleRateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseConfigID) {
		body["BaseConfigID"] = request.BaseConfigID
	}

	if !dara.IsNil(request.DomainNames) {
		body["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.SampleRate) {
		body["SampleRate"] = request.SampleRate
	}

	if !dara.IsNil(request.SinkID) {
		body["SinkID"] = request.SinkID
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCustomDomainSampleRate"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCustomDomainSampleRateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the acceleration region.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ModifyDCdnDomainSchdmByPropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDCdnDomainSchdmByPropertyResponse
func (client *Client) ModifyDCdnDomainSchdmByPropertyWithContext(ctx context.Context, request *ModifyDCdnDomainSchdmByPropertyRequest, runtime *dara.RuntimeOptions) (_result *ModifyDCdnDomainSchdmByPropertyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Property) {
		query["Property"] = request.Property
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDCdnDomainSchdmByProperty"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDCdnDomainSchdmByPropertyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a custom Web Application Firewall (WAF) rule group.
//
// @param request - ModifyDcdnWafGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDcdnWafGroupResponse
func (client *Client) ModifyDcdnWafGroupWithContext(ctx context.Context, request *ModifyDcdnWafGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDcdnWafGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Rules) {
		body["Rules"] = request.Rules
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDcdnWafGroup"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDcdnWafGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the name or the status of a protection policy.
//
// Description:
//
//	  You can call this operation up to 20 times per second per account.
//
//		- Alibaba Cloud Dynamic Content Delivery Network (DCDN) supports POST requests.
//
// @param request - ModifyDcdnWafPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDcdnWafPolicyResponse
func (client *Client) ModifyDcdnWafPolicyWithContext(ctx context.Context, request *ModifyDcdnWafPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyDcdnWafPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PolicyName) {
		body["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicyStatus) {
		body["PolicyStatus"] = request.PolicyStatus
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDcdnWafPolicy"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDcdnWafPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the accelerated domain names that are bound to a protection policy.
//
// Description:
//
// # Usage notes
//
//   - You can call this operation up to 20 times per second per account.
//
//   - Alibaba Cloud Dynamic Route for CDN (DCDN) supports POST requests.
//
// @param request - ModifyDcdnWafPolicyDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDcdnWafPolicyDomainsResponse
func (client *Client) ModifyDcdnWafPolicyDomainsWithContext(ctx context.Context, request *ModifyDcdnWafPolicyDomainsRequest, runtime *dara.RuntimeOptions) (_result *ModifyDcdnWafPolicyDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BindDomains) {
		body["BindDomains"] = request.BindDomains
	}

	if !dara.IsNil(request.Method) {
		body["Method"] = request.Method
	}

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.UnbindDomains) {
		body["UnbindDomains"] = request.UnbindDomains
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDcdnWafPolicyDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDcdnWafPolicyDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the name, status, or configurations of a protection rule.
//
// Description:
//
//	  You can call this operation up to 20 times per second per account.
//
//		- Alibaba Cloud Dynamic Content Delivery Network (DCDN) supports POST requests.
//
//		- You must configure at least one of the **RuleStatus**, **RuleName*	- and **RuleConfig*	- parameters.
//
// @param request - ModifyDcdnWafRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDcdnWafRuleResponse
func (client *Client) ModifyDcdnWafRuleWithContext(ctx context.Context, request *ModifyDcdnWafRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyDcdnWafRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RuleConfig) {
		body["RuleConfig"] = request.RuleConfig
	}

	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		body["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleStatus) {
		body["RuleStatus"] = request.RuleStatus
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDcdnWafRule"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDcdnWafRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates Dynamic Route for CDN (DCDN).
//
// Description:
//
// >
//
//   - DCDN can be activated only once per Alibaba Cloud account. The Alibaba Cloud account must pass real-name verification.
//
//   - You can call this operation up to five times per second per user.
//
// @param request - OpenDcdnServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenDcdnServiceResponse
func (client *Client) OpenDcdnServiceWithContext(ctx context.Context, request *OpenDcdnServiceRequest, runtime *dara.RuntimeOptions) (_result *OpenDcdnServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillType) {
		query["BillType"] = request.BillType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.WebsocketBillType) {
		query["WebsocketBillType"] = request.WebsocketBillType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenDcdnService"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenDcdnServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Prefetches content from origin servers to points of presence (POPs). This reduces workloads on origin servers because users can hit cache upon their first visits.
//
// Description:
//
//	  You can call the [RefreshDcdnObjectCaches](https://help.aliyun.com/document_detail/130620.html) operation to refresh content and call the [PreloadDcdnObjectCaches](https://help.aliyun.com/document_detail/130636.html) operation to prefetch content.
//
//		- Dynamic Content Delivery Network (DCDN) supports POST requests in which parameters are sent as a form.
//
//		- By default, each Alibaba Cloud account can submit up to 1,000 URLs per day. If the daily peak bandwidth value of your workloads exceeds 200 Mbit/s, you can [submit a ticket](https://account.alibabacloud.com/login/login.htm?oauth_callback=https%3A//ticket-intl.console.aliyun.com/%23/ticket/createIndex) to increase your daily quota. Alibaba Cloud reviews your application and then increases the quota accordingly.
//
//		- You can specify up to 100 URLs to prefetch.
//
//		- The prefetch queue of each Alibaba Cloud account can contain up to 100,000 URLs. DCDN executes prefetch tasks based on the time at which you submit the URLs.
//
//		- You can call this operation up to 15 times per second per account.
//
// ## Description
//
//   - After a refresh task is submitted and completed, the POPs immediately start to retrieve resources from the origin server. Therefore, a large number of refresh tasks cause a large number of concurrent download tasks. This increases the number of requests that are redirected to the origin server. The back-to-origin routing process consumes more bandwidth resources and the origin server may be overwhelmed.
//
//   - The time required for a prefetch task to complete is proportional to the size of the prefetched file. In actual practice, most prefetch tasks require 5 to 30 minutes to complete. A task with a smaller average file size requires less time.
//
//   - To allow RAM users to perform this operation, you need to first grant them the required permissions. For more information, see [Authorize a RAM user to prefetch and refresh resources](https://help.aliyun.com/document_detail/445051.html).
//
// @param request - PreloadDcdnObjectCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreloadDcdnObjectCachesResponse
func (client *Client) PreloadDcdnObjectCachesWithContext(ctx context.Context, request *PreloadDcdnObjectCachesRequest, runtime *dara.RuntimeOptions) (_result *PreloadDcdnObjectCachesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.L2Preload) {
		query["L2Preload"] = request.L2Preload
	}

	if !dara.IsNil(request.ObjectPath) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QueryHashkey) {
		query["QueryHashkey"] = request.QueryHashkey
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.WithHeader) {
		query["WithHeader"] = request.WithHeader
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreloadDcdnObjectCaches"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreloadDcdnObjectCachesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes the configurations of an accelerated domain name from the staging environment to the production environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - PublishDcdnStagingConfigToProductionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishDcdnStagingConfigToProductionResponse
func (client *Client) PublishDcdnStagingConfigToProductionWithContext(ctx context.Context, request *PublishDcdnStagingConfigToProductionRequest, runtime *dara.RuntimeOptions) (_result *PublishDcdnStagingConfigToProductionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishDcdnStagingConfigToProduction"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishDcdnStagingConfigToProductionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a specified version of routine code to an environment.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param tmpReq - PublishRoutineCodeRevisionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishRoutineCodeRevisionResponse
func (client *Client) PublishRoutineCodeRevisionWithContext(ctx context.Context, tmpReq *PublishRoutineCodeRevisionRequest, runtime *dara.RuntimeOptions) (_result *PublishRoutineCodeRevisionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &PublishRoutineCodeRevisionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Envs) {
		request.EnvsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Envs, dara.String("Envs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EnvsShrink) {
		body["Envs"] = request.EnvsShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SelectCodeRevision) {
		body["SelectCodeRevision"] = request.SelectCodeRevision
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishRoutineCodeRevision"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishRoutineCodeRevisionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets key-value pairs in a namespace.
//
// @param request - PutDcdnKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutDcdnKvResponse
func (client *Client) PutDcdnKvWithContext(ctx context.Context, request *PutDcdnKvRequest, runtime *dara.RuntimeOptions) (_result *PutDcdnKvResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.ExpirationTtl) {
		query["ExpirationTtl"] = request.ExpirationTtl
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Value) {
		body["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutDcdnKv"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutDcdnKvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds namespaces to your account.
//
// @param request - PutDcdnKvNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutDcdnKvNamespaceResponse
func (client *Client) PutDcdnKvNamespaceWithContext(ctx context.Context, request *PutDcdnKvNamespaceRequest, runtime *dara.RuntimeOptions) (_result *PutDcdnKvNamespaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutDcdnKvNamespace"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutDcdnKvNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置Namespace的key-value对，支持最大25M的请求体
//
// @param request - PutDcdnKvWithHighCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutDcdnKvWithHighCapacityResponse
func (client *Client) PutDcdnKvWithHighCapacityWithContext(ctx context.Context, request *PutDcdnKvWithHighCapacityRequest, runtime *dara.RuntimeOptions) (_result *PutDcdnKvWithHighCapacityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutDcdnKvWithHighCapacity"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutDcdnKvWithHighCapacityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定缓存tag刷新节点上的文件内容
//
// @param request - RefreshDcdnObjectCacheByCacheTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshDcdnObjectCacheByCacheTagResponse
func (client *Client) RefreshDcdnObjectCacheByCacheTagWithContext(ctx context.Context, request *RefreshDcdnObjectCacheByCacheTagRequest, runtime *dara.RuntimeOptions) (_result *RefreshDcdnObjectCacheByCacheTagResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CacheTag) {
		query["CacheTag"] = request.CacheTag
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshDcdnObjectCacheByCacheTag"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshDcdnObjectCacheByCacheTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Refreshes specified objects on points of presence (POPs). The objects can be included in the content of files or URLs. You can refresh multiple URLs in each request.
//
// Description:
//
//	  Dynamic Content Delivery Network (DCDN) supports POST requests in which parameters are sent as a form.
//
//		- You can call the [RefreshDcdnObjectCaches](https://help.aliyun.com/document_detail/130620.html) operation to purge content and call the [PreloadDcdnObjectCaches](https://help.aliyun.com/document_detail/130636.html) operation to prefetch content.
//
//		- By default, each Alibaba Cloud account can purge content from a maximum of 10,000 URLs and 100 directories including subdirectories per day. If the daily peak bandwidth of your Alibaba Cloud account exceeds 200 Mbit/s, [submit a ticket](https://account.alibabacloud.com/login/login.htm?oauth_callback=https%3A//ticket-intl.console.aliyun.com/%23/ticket/createIndex) to request a quota increase. Alibaba Cloud determines whether to approve your application based on your workloads.
//
//		- You can specify up to 1,000 URLs or 100 directories that you want to purge in each request.
//
//		- You can specify up to 1,000 URLs that you want to purge per minute for each domain name.
//
//		- You can call this operation up to 30 times per second per account.
//
// #### [](#)Precautions
//
//   - After a purge task is completed, your resources that are cached on points of presence (POPs) are removed. When a POP receives a request for your resources, the request is redirected to the origin server to retrieve the resources. Then, the resources are returned to the client and cached on POPs. If you frequently run purge tasks, more requests are redirected to the origin server for resources. This results in high bandwidth costs and more loads on the origin server.
//
//   - A purge task takes effect 5 to 6 minutes after being submitted. If the resource you want to purge has a TTL of less than 5 minutes, you wait for it to expire instead of manually running a purge task.
//
//   - To allow RAM users to perform this operation, you need to first grant them the required permissions. For more information, see [Authorize a RAM user to prefetch and refresh resources](https://help.aliyun.com/document_detail/445051.html).
//
// @param request - RefreshDcdnObjectCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshDcdnObjectCachesResponse
func (client *Client) RefreshDcdnObjectCachesWithContext(ctx context.Context, request *RefreshDcdnObjectCachesRequest, runtime *dara.RuntimeOptions) (_result *RefreshDcdnObjectCachesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.ObjectPath) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshDcdnObjectCaches"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshDcdnObjectCachesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Refreshes the cache that is written by calling the cache operation of EdgeRoutine. You can refresh multiple URLs in each request.
//
// Description:
//
// > 	- Dynamic Content Delivery Network (DCDN) supports POST requests in which parameters are sent as a form.
//
// > 	- Related operation: [RefreshDcdnObjectCaches](https://help.aliyun.com/document_detail/130620.html).
//
// > 	- By default, each Alibaba Cloud account can purge content from a maxim> um of 10,000 URLs and 100 directories including subdirectories per day.
//
// > 	- You can specify up to 1,000 URLs or 100 directories that you want to purge in each request.
//
// > 	- You can specify up to 1,000 URLs that you want to purge per minute for each domain name.
//
// > 	- You can call this operation up to 30 times per second per account.
//
// @param request - RefreshErObjectCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshErObjectCachesResponse
func (client *Client) RefreshErObjectCachesWithContext(ctx context.Context, request *RefreshErObjectCachesRequest, runtime *dara.RuntimeOptions) (_result *RefreshErObjectCachesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.MergeDomainName) {
		query["MergeDomainName"] = request.MergeDomainName
	}

	if !dara.IsNil(request.ObjectPath) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.RoutineId) {
		query["RoutineId"] = request.RoutineId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshErObjectCaches"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshErObjectCachesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back the configurations of an accelerated domain name from the staging environment to the production environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - RollbackDcdnStagingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackDcdnStagingConfigResponse
func (client *Client) RollbackDcdnStagingConfigWithContext(ctx context.Context, request *RollbackDcdnStagingConfigRequest, runtime *dara.RuntimeOptions) (_result *RollbackDcdnStagingConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackDcdnStagingConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackDcdnStagingConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures an SSL certificate for a specified domain name.
//
// @param request - SetDcdnDomainCSRCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDcdnDomainCSRCertificateResponse
func (client *Client) SetDcdnDomainCSRCertificateWithContext(ctx context.Context, request *SetDcdnDomainCSRCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetDcdnDomainCSRCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.ServerCertificate) {
		query["ServerCertificate"] = request.ServerCertificate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDcdnDomainCSRCertificate"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDcdnDomainCSRCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the ShangMi (SM) certificate for a domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - SetDcdnDomainSMCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDcdnDomainSMCertificateResponse
func (client *Client) SetDcdnDomainSMCertificateWithContext(ctx context.Context, request *SetDcdnDomainSMCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetDcdnDomainSMCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SSLProtocol) {
		query["SSLProtocol"] = request.SSLProtocol
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDcdnDomainSMCertificate"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDcdnDomainSMCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the SSL certificate for a domain name and updates certificate details.
//
// @param request - SetDcdnDomainSSLCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDcdnDomainSSLCertificateResponse
func (client *Client) SetDcdnDomainSSLCertificateWithContext(ctx context.Context, request *SetDcdnDomainSSLCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetDcdnDomainSSLCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertId) {
		query["CertId"] = request.CertId
	}

	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

	if !dara.IsNil(request.CertRegion) {
		query["CertRegion"] = request.CertRegion
	}

	if !dara.IsNil(request.CertType) {
		query["CertType"] = request.CertType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SSLPri) {
		query["SSLPri"] = request.SSLPri
	}

	if !dara.IsNil(request.SSLProtocol) {
		query["SSLProtocol"] = request.SSLProtocol
	}

	if !dara.IsNil(request.SSLPub) {
		query["SSLPub"] = request.SSLPub
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDcdnDomainSSLCertificate"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDcdnDomainSSLCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets or modifies the domain name configuration in the canary release environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - SetDcdnDomainStagingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDcdnDomainStagingConfigResponse
func (client *Client) SetDcdnDomainStagingConfigWithContext(ctx context.Context, request *SetDcdnDomainStagingConfigRequest, runtime *dara.RuntimeOptions) (_result *SetDcdnDomainStagingConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Functions) {
		query["Functions"] = request.Functions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDcdnDomainStagingConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDcdnDomainStagingConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Blocks or unblocks IP addresses or CIDR blocks.
//
// Description:
//
// >  You can call this operation to block or unblock a large number of IP addresses or CIDR blocks. You can block or unblock up to 1,000 IP addresses or CIDR blocks in a request.
//
// @param request - SetDcdnFullDomainsBlockIPRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDcdnFullDomainsBlockIPResponse
func (client *Client) SetDcdnFullDomainsBlockIPWithContext(ctx context.Context, request *SetDcdnFullDomainsBlockIPRequest, runtime *dara.RuntimeOptions) (_result *SetDcdnFullDomainsBlockIPResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BlockInterval) {
		body["BlockInterval"] = request.BlockInterval
	}

	if !dara.IsNil(request.IPList) {
		body["IPList"] = request.IPList
	}

	if !dara.IsNil(request.OperationType) {
		body["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.UpdateType) {
		body["UpdateType"] = request.UpdateType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDcdnFullDomainsBlockIP"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDcdnFullDomainsBlockIPResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures features for a user.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - SetDcdnUserConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDcdnUserConfigResponse
func (client *Client) SetDcdnUserConfigWithContext(ctx context.Context, request *SetDcdnUserConfigRequest, runtime *dara.RuntimeOptions) (_result *SetDcdnUserConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Configs) {
		query["Configs"] = request.Configs
	}

	if !dara.IsNil(request.FunctionId) {
		query["FunctionId"] = request.FunctionId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDcdnUserConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDcdnUserConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a subdomain for a routine.
//
// Description:
//
// >
//
//   - Each subdomain is globally unique. Resource Access Management (RAM) users cannot create duplicate subdomains.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param tmpReq - SetRoutineSubdomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetRoutineSubdomainResponse
func (client *Client) SetRoutineSubdomainWithContext(ctx context.Context, tmpReq *SetRoutineSubdomainRequest, runtime *dara.RuntimeOptions) (_result *SetRoutineSubdomainResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SetRoutineSubdomainShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Subdomains) {
		request.SubdomainsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Subdomains, dara.String("Subdomains"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SubdomainsShrink) {
		body["Subdomains"] = request.SubdomainsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetRoutineSubdomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetRoutineSubdomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a specified disabled accelerated domain. After the accelerated domain is enabled, the value of the DomainStatus parameter changes to Online for the domain.
//
// Description:
//
// >
//
//   - If an accelerated domain name is in invalid state or your account has an overdue payment, the accelerated domain name cannot be enabled.
//
//   - You can call this operation up to 30 times per second per account.
//
// @param request - StartDcdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDcdnDomainResponse
func (client *Client) StartDcdnDomainWithContext(ctx context.Context, request *StartDcdnDomainRequest, runtime *dara.RuntimeOptions) (_result *StartDcdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDcdnDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDcdnDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables one or more accelerated domain names. After the accelerated domain names are enabled, the value of the DomainStatus parameter for the domain names changes to Online.
//
// Description:
//
//	  If an accelerated domain name is in invalid state or your account has an overdue payment, the accelerated domain name cannot be enabled.
//
//		- You can call this operation up to 20 times per second per account.
//
// @param request - StartDcdnIpaDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDcdnIpaDomainResponse
func (client *Client) StartDcdnIpaDomainWithContext(ctx context.Context, request *StartDcdnIpaDomainRequest, runtime *dara.RuntimeOptions) (_result *StartDcdnIpaDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDcdnIpaDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDcdnIpaDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a specified accelerated domain. After the accelerated domain is disabled,
//
//	the value of the DomainStatus parameter changes to Offline for the domain.
//
// Description:
//
// >
//
//   - After an accelerated domain is disabled, Dynamic Content Delivery Network (DCDN) retains its information and routes all the requests that are destined for the accelerated domain to the origin server.
//
//   - You can call this operation up to 30 times per second per account.
//
// @param request - StopDcdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDcdnDomainResponse
func (client *Client) StopDcdnDomainWithContext(ctx context.Context, request *StopDcdnDomainRequest, runtime *dara.RuntimeOptions) (_result *StopDcdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDcdnDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDcdnDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an accelerated domain name. After an accelerated domain name is disabled, the value of the DomainStatus parameter for the domain name changes to Offline.
//
// Description:
//
// >
//
//   - If you disable an accelerated domain, the configurations of the accelerated domain are still retained. The system automatically forwards all the requests that are destined for this domain to the origin.
//
//   - You can call this operation up to 20 times per second per account.
//
// @param request - StopDcdnIpaDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDcdnIpaDomainResponse
func (client *Client) StopDcdnIpaDomainWithContext(ctx context.Context, request *StopDcdnIpaDomainRequest, runtime *dara.RuntimeOptions) (_result *StopDcdnIpaDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDcdnIpaDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDcdnIpaDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds one or more tags to a resource.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - TagDcdnResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagDcdnResourcesResponse
func (client *Client) TagDcdnResourcesWithContext(ctx context.Context, request *TagDcdnResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagDcdnResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
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
		Action:      dara.String("TagDcdnResources"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagDcdnResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes one or more tags from a resource.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - UntagDcdnResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagDcdnResourcesResponse
func (client *Client) UntagDcdnResourcesWithContext(ctx context.Context, request *UntagDcdnResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagDcdnResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
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
		Action:      dara.String("UntagDcdnResources"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagDcdnResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a tracking task by task ID.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - UpdateDcdnDeliverTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDcdnDeliverTaskResponse
func (client *Client) UpdateDcdnDeliverTaskWithContext(ctx context.Context, request *UpdateDcdnDeliverTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateDcdnDeliverTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Deliver) {
		body["Deliver"] = request.Deliver
	}

	if !dara.IsNil(request.DeliverId) {
		body["DeliverId"] = request.DeliverId
	}

	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Reports) {
		body["Reports"] = request.Reports
	}

	if !dara.IsNil(request.Schedule) {
		body["Schedule"] = request.Schedule
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDcdnDeliverTask"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDcdnDeliverTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - UpdateDcdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDcdnDomainResponse
func (client *Client) UpdateDcdnDomainWithContext(ctx context.Context, request *UpdateDcdnDomainRequest, runtime *dara.RuntimeOptions) (_result *UpdateDcdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	if !dara.IsNil(request.TopLevelDomain) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDcdnDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDcdnDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a domain name that is accelerated by IP Application Accelerator (IPA).
//
// Description:
//
// >  You can call this operation up to 20 times per second per account.
//
// @param request - UpdateDcdnIpaDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDcdnIpaDomainResponse
func (client *Client) UpdateDcdnIpaDomainWithContext(ctx context.Context, request *UpdateDcdnIpaDomainRequest, runtime *dara.RuntimeOptions) (_result *UpdateDcdnIpaDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	if !dara.IsNil(request.TopLevelDomain) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDcdnIpaDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDcdnIpaDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a real-time log delivery project.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - UpdateDcdnSLSRealtimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDcdnSLSRealtimeLogDeliveryResponse
func (client *Client) UpdateDcdnSLSRealtimeLogDeliveryWithContext(ctx context.Context, request *UpdateDcdnSLSRealtimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *UpdateDcdnSLSRealtimeLogDeliveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataCenter) {
		body["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SLSLogStore) {
		body["SLSLogStore"] = request.SLSLogStore
	}

	if !dara.IsNil(request.SLSProject) {
		body["SLSProject"] = request.SLSProject
	}

	if !dara.IsNil(request.SLSRegion) {
		body["SLSRegion"] = request.SLSRegion
	}

	if !dara.IsNil(request.SamplingRate) {
		body["SamplingRate"] = request.SamplingRate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDcdnSLSRealtimeLogDelivery"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDcdnSLSRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates one or more operations reports.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - UpdateDcdnSubTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDcdnSubTaskResponse
func (client *Client) UpdateDcdnSubTaskWithContext(ctx context.Context, request *UpdateDcdnSubTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateDcdnSubTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ReportIds) {
		body["ReportIds"] = request.ReportIds
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDcdnSubTask"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDcdnSubTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the fields in real-time log entries.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - UpdateDcdnUserRealTimeDeliveryFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDcdnUserRealTimeDeliveryFieldResponse
func (client *Client) UpdateDcdnUserRealTimeDeliveryFieldWithContext(ctx context.Context, request *UpdateDcdnUserRealTimeDeliveryFieldRequest, runtime *dara.RuntimeOptions) (_result *UpdateDcdnUserRealTimeDeliveryFieldResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.Fields) {
		query["Fields"] = request.Fields
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDcdnUserRealTimeDeliveryField"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDcdnUserRealTimeDeliveryFieldResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads code to EdgeRoutine (ER).
//
// Description:
//
// >
//
//   - Each time you submit code, a version of the code is generated. You can manage and publish code by version.
//
//   - Each routine can retain at most 10 versions. If the upper limit is reached, you must call the DeleteRoutineCodeRevision operation to manually delete versions that are no longer needed before new versions can be saved.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - UploadRoutineCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadRoutineCodeResponse
func (client *Client) UploadRoutineCodeWithContext(ctx context.Context, request *UploadRoutineCodeRequest, runtime *dara.RuntimeOptions) (_result *UploadRoutineCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CodeDescription) {
		body["CodeDescription"] = request.CodeDescription
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadRoutineCode"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadRoutineCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads code to a routine for testing.
//
// Description:
//
// >
//
//   - Each time you upload code to a routine, a version is generated. The number of versions is counted by CodeRev. The uploaded code is used only for testing.
//
//   - The code is automatically published to a staging environment.
//
//   - Each routine can retain at most 10 versions. If the upper limit is reached, you need to call the DeleteRoutineCodeRevision operation to manually delete versions that are no longer needed before new versions can be saved.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - UploadStagingRoutineCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadStagingRoutineCodeResponse
func (client *Client) UploadStagingRoutineCodeWithContext(ctx context.Context, request *UploadStagingRoutineCodeRequest, runtime *dara.RuntimeOptions) (_result *UploadStagingRoutineCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CodeDescription) {
		body["CodeDescription"] = request.CodeDescription
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadStagingRoutineCode"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadStagingRoutineCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the ownership of a domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - VerifyDcdnDomainOwnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyDcdnDomainOwnerResponse
func (client *Client) VerifyDcdnDomainOwnerWithContext(ctx context.Context, request *VerifyDcdnDomainOwnerRequest, runtime *dara.RuntimeOptions) (_result *VerifyDcdnDomainOwnerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.VerifyType) {
		query["VerifyType"] = request.VerifyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyDcdnDomainOwner"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyDcdnDomainOwnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
