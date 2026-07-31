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
		"us-west-1":             dara.String("ecd.us-west-1.aliyuncs.com"),
		"us-east-1":             dara.String("ecd.us-east-1.aliyuncs.com"),
		"me-east-1":             dara.String("ecd.me-east-1.aliyuncs.com"),
		"me-central-1":          dara.String("ecd.me-central-1.aliyuncs.com"),
		"eu-west-1":             dara.String("ecd.eu-west-1.aliyuncs.com"),
		"eu-central-1":          dara.String("ecd.eu-central-1.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("ecd.cn-zhangjiakou.aliyuncs.com"),
		"cn-wulanchabu":         dara.String("ecd.cn-wulanchabu.aliyuncs.com"),
		"cn-shenzhen":           dara.String("ecd.cn-shenzhen.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("ecd.cn-shanghai-finance-1.aliyuncs.com"),
		"cn-shanghai":           dara.String("ecd.cn-shanghai.aliyuncs.com"),
		"cn-qingdao":            dara.String("ecd.cn-qingdao.aliyuncs.com"),
		"cn-nanjing":            dara.String("ecd.cn-nanjing.aliyuncs.com"),
		"cn-hongkong":           dara.String("ecd.cn-hongkong.aliyuncs.com"),
		"cn-hangzhou-finance":   dara.String("ecd.cn-hangzhou-finance.aliyuncs.com"),
		"cn-hangzhou":           dara.String("ecd.cn-hangzhou.aliyuncs.com"),
		"cn-guangzhou":          dara.String("ecd.cn-guangzhou.aliyuncs.com"),
		"cn-chengdu":            dara.String("ecd.cn-chengdu.aliyuncs.com"),
		"cn-beijing":            dara.String("ecd.cn-beijing.aliyuncs.com"),
		"ap-southeast-7":        dara.String("ecd.ap-southeast-7.aliyuncs.com"),
		"ap-southeast-6":        dara.String("ecd.ap-southeast-6.aliyuncs.com"),
		"ap-southeast-5":        dara.String("ecd.ap-southeast-5.aliyuncs.com"),
		"ap-southeast-1":        dara.String("ecd.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-1":        dara.String("ecd.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("ecd"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Creates a tenant skill.
//
// @param tmpReq - CreateTenantSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTenantSkillResponse
func (client *Client) CreateTenantSkillWithOptions(tmpReq *CreateTenantSkillRequest, runtime *dara.RuntimeOptions) (_result *CreateTenantSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTenantSkillShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EnvVars) {
		request.EnvVarsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EnvVars, dara.String("EnvVars"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.EnvVarsShrink) {
		query["EnvVars"] = request.EnvVarsShrink
	}

	if !dara.IsNil(request.IconETag) {
		query["IconETag"] = request.IconETag
	}

	if !dara.IsNil(request.SkillChannel) {
		query["SkillChannel"] = request.SkillChannel
	}

	if !dara.IsNil(request.SkillIcon) {
		query["SkillIcon"] = request.SkillIcon
	}

	if !dara.IsNil(request.SkillVersion) {
		query["SkillVersion"] = request.SkillVersion
	}

	if !dara.IsNil(request.Slug) {
		query["Slug"] = request.Slug
	}

	if !dara.IsNil(request.TaskKey) {
		query["TaskKey"] = request.TaskKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTenantSkill"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTenantSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a tenant skill.
//
// @param request - CreateTenantSkillRequest
//
// @return CreateTenantSkillResponse
func (client *Client) CreateTenantSkill(request *CreateTenantSkillRequest) (_result *CreateTenantSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTenantSkillResponse{}
	_body, _err := client.CreateTenantSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes skills in batches.
//
// @param request - DeleteTenantSkillsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTenantSkillsResponse
func (client *Client) DeleteTenantSkillsWithOptions(request *DeleteTenantSkillsRequest, runtime *dara.RuntimeOptions) (_result *DeleteTenantSkillsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SkillChannel) {
		query["SkillChannel"] = request.SkillChannel
	}

	if !dara.IsNil(request.SkillIds) {
		query["SkillIds"] = request.SkillIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTenantSkills"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTenantSkillsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes skills in batches.
//
// @param request - DeleteTenantSkillsRequest
//
// @return DeleteTenantSkillsResponse
func (client *Client) DeleteTenantSkills(request *DeleteTenantSkillsRequest) (_result *DeleteTenantSkillsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTenantSkillsResponse{}
	_body, _err := client.DeleteTenantSkillsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a temporary OSS token for authentication.
//
// Description:
//
// The obtained SecurityToken is valid for 15 minutes.
//
// @param request - GetOssStsTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssStsTokenResponse
func (client *Client) GetOssStsTokenWithOptions(request *GetOssStsTokenRequest, runtime *dara.RuntimeOptions) (_result *GetOssStsTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOssStsToken"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOssStsTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a temporary OSS token for authentication.
//
// Description:
//
// The obtained SecurityToken is valid for 15 minutes.
//
// @param request - GetOssStsTokenRequest
//
// @return GetOssStsTokenResponse
func (client *Client) GetOssStsToken(request *GetOssStsTokenRequest) (_result *GetOssStsTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOssStsTokenResponse{}
	_body, _err := client.GetOssStsTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the parsed content of a skill package.
//
// Description:
//
// Call the ParseSkillPackage operation first. Poll this operation every 3 seconds.
//
// @param request - GetParseProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetParseProgressResponse
func (client *Client) GetParseProgressWithOptions(request *GetParseProgressRequest, runtime *dara.RuntimeOptions) (_result *GetParseProgressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskKey) {
		query["TaskKey"] = request.TaskKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetParseProgress"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetParseProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the parsed content of a skill package.
//
// Description:
//
// Call the ParseSkillPackage operation first. Poll this operation every 3 seconds.
//
// @param request - GetParseProgressRequest
//
// @return GetParseProgressResponse
func (client *Client) GetParseProgress(request *GetParseProgressRequest) (_result *GetParseProgressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetParseProgressResponse{}
	_body, _err := client.GetParseProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of identities for which security policies are enabled.
//
// Description:
//
// The resource type supports only cloud computers.
//
// @param request - ListSecureSkillIdentitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecureSkillIdentitiesResponse
func (client *Client) ListSecureSkillIdentitiesWithOptions(request *ListSecureSkillIdentitiesRequest, runtime *dara.RuntimeOptions) (_result *ListSecureSkillIdentitiesResponse, _err error) {
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

	if !dara.IsNil(request.SkillChannel) {
		query["SkillChannel"] = request.SkillChannel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSecureSkillIdentities"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecureSkillIdentitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of identities for which security policies are enabled.
//
// Description:
//
// The resource type supports only cloud computers.
//
// @param request - ListSecureSkillIdentitiesRequest
//
// @return ListSecureSkillIdentitiesResponse
func (client *Client) ListSecureSkillIdentities(request *ListSecureSkillIdentitiesRequest) (_result *ListSecureSkillIdentitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSecureSkillIdentitiesResponse{}
	_body, _err := client.ListSecureSkillIdentitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of identities authorized for a skill.
//
// Description:
//
// Authorized objects support only cloud computers.
//
// @param request - ListSkillAuthedIdentitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillAuthedIdentitiesResponse
func (client *Client) ListSkillAuthedIdentitiesWithOptions(request *ListSkillAuthedIdentitiesRequest, runtime *dara.RuntimeOptions) (_result *ListSkillAuthedIdentitiesResponse, _err error) {
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

	if !dara.IsNil(request.SkillChannel) {
		query["SkillChannel"] = request.SkillChannel
	}

	if !dara.IsNil(request.SkillId) {
		query["SkillId"] = request.SkillId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkillAuthedIdentities"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillAuthedIdentitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of identities authorized for a skill.
//
// Description:
//
// Authorized objects support only cloud computers.
//
// @param request - ListSkillAuthedIdentitiesRequest
//
// @return ListSkillAuthedIdentitiesResponse
func (client *Client) ListSkillAuthedIdentities(request *ListSkillAuthedIdentitiesRequest) (_result *ListSkillAuthedIdentitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSkillAuthedIdentitiesResponse{}
	_body, _err := client.ListSkillAuthedIdentitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of skills.
//
// @param request - ListSkillsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillsResponse
func (client *Client) ListSkillsWithOptions(request *ListSkillsRequest, runtime *dara.RuntimeOptions) (_result *ListSkillsResponse, _err error) {
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

	if !dara.IsNil(request.SkillChannel) {
		query["SkillChannel"] = request.SkillChannel
	}

	if !dara.IsNil(request.SkillIds) {
		query["SkillIds"] = request.SkillIds
	}

	if !dara.IsNil(request.SupplierType) {
		query["SupplierType"] = request.SupplierType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkills"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of skills.
//
// @param request - ListSkillsRequest
//
// @return ListSkillsResponse
func (client *Client) ListSkills(request *ListSkillsRequest) (_result *ListSkillsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSkillsResponse{}
	_body, _err := client.ListSkillsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Parses a skill package.
//
// @param request - ParseSkillPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ParseSkillPackageResponse
func (client *Client) ParseSkillPackageWithOptions(request *ParseSkillPackageRequest, runtime *dara.RuntimeOptions) (_result *ParseSkillPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OssObjectETag) {
		query["OssObjectETag"] = request.OssObjectETag
	}

	if !dara.IsNil(request.OssObjectKey) {
		query["OssObjectKey"] = request.OssObjectKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ParseSkillPackage"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ParseSkillPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Parses a skill package.
//
// @param request - ParseSkillPackageRequest
//
// @return ParseSkillPackageResponse
func (client *Client) ParseSkillPackage(request *ParseSkillPackageRequest) (_result *ParseSkillPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ParseSkillPackageResponse{}
	_body, _err := client.ParseSkillPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets skill permissions for an identity.
//
// Description:
//
// The authorized object supports only cloud computers.
//
// @param request - SetIdentitySkillAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetIdentitySkillAuthResponse
func (client *Client) SetIdentitySkillAuthWithOptions(request *SetIdentitySkillAuthRequest, runtime *dara.RuntimeOptions) (_result *SetIdentitySkillAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoInstall) {
		query["AutoInstall"] = request.AutoInstall
	}

	if !dara.IsNil(request.Identities) {
		query["Identities"] = request.Identities
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.SkillChannel) {
		query["SkillChannel"] = request.SkillChannel
	}

	if !dara.IsNil(request.SkillIds) {
		query["SkillIds"] = request.SkillIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetIdentitySkillAuth"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetIdentitySkillAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets skill permissions for an identity.
//
// Description:
//
// The authorized object supports only cloud computers.
//
// @param request - SetIdentitySkillAuthRequest
//
// @return SetIdentitySkillAuthResponse
func (client *Client) SetIdentitySkillAuth(request *SetIdentitySkillAuthRequest) (_result *SetIdentitySkillAuthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetIdentitySkillAuthResponse{}
	_body, _err := client.SetIdentitySkillAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the security policy for identity skills.
//
// Description:
//
// The resource type supports only cloud computers.
//
// @param request - SetIdentitySkillSecurityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetIdentitySkillSecurityResponse
func (client *Client) SetIdentitySkillSecurityWithOptions(request *SetIdentitySkillSecurityRequest, runtime *dara.RuntimeOptions) (_result *SetIdentitySkillSecurityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.IdentityIds) {
		query["IdentityIds"] = request.IdentityIds
	}

	if !dara.IsNil(request.SkillChannel) {
		query["SkillChannel"] = request.SkillChannel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetIdentitySkillSecurity"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetIdentitySkillSecurityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the security policy for identity skills.
//
// Description:
//
// The resource type supports only cloud computers.
//
// @param request - SetIdentitySkillSecurityRequest
//
// @return SetIdentitySkillSecurityResponse
func (client *Client) SetIdentitySkillSecurity(request *SetIdentitySkillSecurityRequest) (_result *SetIdentitySkillSecurityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetIdentitySkillSecurityResponse{}
	_body, _err := client.SetIdentitySkillSecurityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the enabling status of skills at the tenant level.
//
// @param request - SetTenantSkillEnabledRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetTenantSkillEnabledResponse
func (client *Client) SetTenantSkillEnabledWithOptions(request *SetTenantSkillEnabledRequest, runtime *dara.RuntimeOptions) (_result *SetTenantSkillEnabledResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.SkillChannel) {
		query["SkillChannel"] = request.SkillChannel
	}

	if !dara.IsNil(request.SkillIds) {
		query["SkillIds"] = request.SkillIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetTenantSkillEnabled"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetTenantSkillEnabledResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the enabling status of skills at the tenant level.
//
// @param request - SetTenantSkillEnabledRequest
//
// @return SetTenantSkillEnabledResponse
func (client *Client) SetTenantSkillEnabled(request *SetTenantSkillEnabledRequest) (_result *SetTenantSkillEnabledResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetTenantSkillEnabledResponse{}
	_body, _err := client.SetTenantSkillEnabledWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
