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
		"eu-central-1":   dara.String("modelstudio.eu-central-1.aliyuncs.com"),
		"cn-hongkong":    dara.String("modelstudio.cn-hongkong.aliyuncs.com"),
		"cn-beijing":     dara.String("modelstudio.cn-beijing.aliyuncs.com"),
		"ap-southeast-1": dara.String("modelstudio.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("modelstudio"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Creates an account and directly adds it as a member.
//
// @param request - AddOrganizationMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOrganizationMemberResponse
func (client *Client) AddOrganizationMemberWithOptions(request *AddOrganizationMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddOrganizationMemberResponse, _err error) {
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

	if !dara.IsNil(request.CallerUacAccountId) {
		query["CallerUacAccountId"] = request.CallerUacAccountId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.OrgId) {
		query["OrgId"] = request.OrgId
	}

	if !dara.IsNil(request.OrgRoleCode) {
		query["OrgRoleCode"] = request.OrgRoleCode
	}

	if !dara.IsNil(request.SpecType) {
		query["SpecType"] = request.SpecType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddOrganizationMember"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/organization/member-additions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddOrganizationMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an account and directly adds it as a member.
//
// @param request - AddOrganizationMemberRequest
//
// @return AddOrganizationMemberResponse
func (client *Client) AddOrganizationMember(request *AddOrganizationMemberRequest) (_result *AddOrganizationMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddOrganizationMemberResponse{}
	_body, _err := client.AddOrganizationMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Assigns seats in bulk to the member level.
//
// @param request - BatchAssignSeatsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAssignSeatsResponse
func (client *Client) BatchAssignSeatsWithOptions(request *BatchAssignSeatsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchAssignSeatsResponse, _err error) {
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

	if !dara.IsNil(request.AccountIdsStr) {
		query["AccountIdsStr"] = request.AccountIdsStr
	}

	if !dara.IsNil(request.CallerUacAccountId) {
		query["CallerUacAccountId"] = request.CallerUacAccountId
	}

	if !dara.IsNil(request.Locale) {
		query["Locale"] = request.Locale
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SeatType) {
		query["SeatType"] = request.SeatType
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchAssignSeats"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/subscription/seat-assignments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchAssignSeatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns seats in bulk to the member level.
//
// @param request - BatchAssignSeatsRequest
//
// @return BatchAssignSeatsResponse
func (client *Client) BatchAssignSeats(request *BatchAssignSeatsRequest) (_result *BatchAssignSeatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchAssignSeatsResponse{}
	_body, _err := client.BatchAssignSeatsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Before using large models or applications in Alibaba Cloud Model Studio, create an API key as an authentication credential.
//
// @param request - CreateApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApiKeyResponse
func (client *Client) CreateApiKeyWithOptions(request *CreateApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Auth) {
		body["auth"] = request.Auth
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApiKey"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/apikeys"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Before using large models or applications in Alibaba Cloud Model Studio, create an API key as an authentication credential.
//
// @param request - CreateApiKeyRequest
//
// @return CreateApiKeyResponse
func (client *Client) CreateApiKey(request *CreateApiKeyRequest) (_result *CreateApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateApiKeyResponse{}
	_body, _err := client.CreateApiKeyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a UAC API key.
//
// @param request - CreateTokenPlanKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTokenPlanKeyResponse
func (client *Client) CreateTokenPlanKeyWithOptions(request *CreateTokenPlanKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTokenPlanKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.CallerUacAccountId) {
		query["CallerUacAccountId"] = request.CallerUacAccountId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTokenPlanKey"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/api-keys"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTokenPlanKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a UAC API key.
//
// @param request - CreateTokenPlanKeyRequest
//
// @return CreateTokenPlanKeyResponse
func (client *Client) CreateTokenPlanKey(request *CreateTokenPlanKeyRequest) (_result *CreateTokenPlanKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTokenPlanKeyResponse{}
	_body, _err := client.CreateTokenPlanKeyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a business workspace.
//
// @param request - CreateWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspaceWithOptions(request *CreateWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceSite) {
		query["serviceSite"] = request.ServiceSite
	}

	if !dara.IsNil(request.WorkspaceName) {
		query["workspaceName"] = request.WorkspaceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkspace"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/workspaces"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a business workspace.
//
// @param request - CreateWorkspaceRequest
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspace(request *CreateWorkspaceRequest) (_result *CreateWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CreateWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an authentication credential API key.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApiKeyResponse
func (client *Client) DeleteApiKeyWithOptions(apiKeyId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteApiKeyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApiKey"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/apikeys/" + dara.PercentEncode(dara.StringValue(apiKeyId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an authentication credential API key.
//
// @return DeleteApiKeyResponse
func (client *Client) DeleteApiKey(apiKeyId *string) (_result *DeleteApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteApiKeyResponse{}
	_body, _err := client.DeleteApiKeyWithOptions(apiKeyId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a workspace.
//
// Description:
//
// A workspace can be deleted only if the following conditional requirements are met:
//
// 1. The workspace is not the default workspace.
//
// 2. The workspace is not used to purchase other products, such as AMB.
//
// 3. In permission management, the workspace is not granted to Resource Access Management (RAM) users or RAM roles.
//
// 4. The workspace does not contain any resources, such as API keys, model deployments, or knowledge bases.
//
// @param request - DeleteWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspaceWithOptions(workspaceId *string, request *DeleteWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkspace"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a workspace.
//
// Description:
//
// A workspace can be deleted only if the following conditional requirements are met:
//
// 1. The workspace is not the default workspace.
//
// 2. The workspace is not used to purchase other products, such as AMB.
//
// 3. In permission management, the workspace is not granted to Resource Access Management (RAM) users or RAM roles.
//
// 4. The workspace does not contain any resources, such as API keys, model deployments, or knowledge bases.
//
// @param request - DeleteWorkspaceRequest
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspace(workspaceId *string, request *DeleteWorkspaceRequest) (_result *DeleteWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.DeleteWorkspaceWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables an API key.
//
// Description:
//
// An API key cannot be disabled if it is already disabled.
//
// @param request - DisableApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApiKeyResponse
func (client *Client) DisableApiKeyWithOptions(apiKeyId *string, request *DisableApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableApiKey"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/apikeys/" + dara.PercentEncode(dara.StringValue(apiKeyId)) + "/disable"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an API key.
//
// Description:
//
// An API key cannot be disabled if it is already disabled.
//
// @param request - DisableApiKeyRequest
//
// @return DisableApiKeyResponse
func (client *Client) DisableApiKey(apiKeyId *string, request *DisableApiKeyRequest) (_result *DisableApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableApiKeyResponse{}
	_body, _err := client.DisableApiKeyWithOptions(apiKeyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables an API key.
//
// Description:
//
// An API key that is already enabled cannot be enabled again.
//
// @param request - EnableApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApiKeyResponse
func (client *Client) EnableApiKeyWithOptions(apiKeyId *string, request *EnableApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableApiKey"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/apikeys/" + dara.PercentEncode(dara.StringValue(apiKeyId)) + "/enable"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an API key.
//
// Description:
//
// An API key that is already enabled cannot be enabled again.
//
// @param request - EnableApiKeyRequest
//
// @return EnableApiKeyResponse
func (client *Client) EnableApiKey(apiKeyId *string, request *EnableApiKeyRequest) (_result *EnableApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableApiKeyResponse{}
	_body, _err := client.EnableApiKeyWithOptions(apiKeyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the information of a specified authentication credential API key.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApiKeyResponse
func (client *Client) GetApiKeyWithOptions(apiKeyId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetApiKeyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApiKey"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/apikeys/" + dara.PercentEncode(dara.StringValue(apiKeyId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the information of a specified authentication credential API key.
//
// @return GetApiKeyResponse
func (client *Client) GetApiKey(apiKeyId *string) (_result *GetApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetApiKeyResponse{}
	_body, _err := client.GetApiKeyWithOptions(apiKeyId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries seat details by paging.
//
// @param request - GetSubscriptionSeatDetailsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubscriptionSeatDetailsResponse
func (client *Client) GetSubscriptionSeatDetailsWithOptions(request *GetSubscriptionSeatDetailsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSubscriptionSeatDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallerUacAccountId) {
		query["CallerUacAccountId"] = request.CallerUacAccountId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryAssigned) {
		query["QueryAssigned"] = request.QueryAssigned
	}

	if !dara.IsNil(request.SeatId) {
		query["SeatId"] = request.SeatId
	}

	if !dara.IsNil(request.SeatType) {
		query["SeatType"] = request.SeatType
	}

	if !dara.IsNil(request.StatusList) {
		query["StatusList"] = request.StatusList
	}

	if !dara.IsNil(request.StatusListStr) {
		query["StatusListStr"] = request.StatusListStr
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSubscriptionSeatDetails"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/subscription/seat-detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSubscriptionSeatDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries seat details by paging.
//
// @param request - GetSubscriptionSeatDetailsRequest
//
// @return GetSubscriptionSeatDetailsResponse
func (client *Client) GetSubscriptionSeatDetails(request *GetSubscriptionSeatDetailsRequest) (_result *GetSubscriptionSeatDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSubscriptionSeatDetailsResponse{}
	_body, _err := client.GetSubscriptionSeatDetailsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the list of authentication credential API Key information.
//
// @param request - ListApiKeysRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiKeysResponse
func (client *Client) ListApiKeysWithOptions(request *ListApiKeysRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListApiKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["apiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApiKeys"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/apikeys"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApiKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the list of authentication credential API Key information.
//
// @param request - ListApiKeysRequest
//
// @return ListApiKeysResponse
func (client *Client) ListApiKeys(request *ListApiKeysRequest) (_result *ListApiKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListApiKeysResponse{}
	_body, _err := client.ListApiKeysWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of business workspaces.
//
// @param request - ListWorkspacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspacesWithOptions(request *ListWorkspacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WorkspaceName) {
		query["workspaceName"] = request.WorkspaceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkspaces"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/workspaces"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of business workspaces.
//
// @param request - ListWorkspacesRequest
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (_result *ListWorkspacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkspacesResponse{}
	_body, _err := client.ListWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets an API key.
//
// Description:
//
// Only the API key value changes. The API key ID remains unchanged.
//
// @param request - ResetApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetApiKeyResponse
func (client *Client) ResetApiKeyWithOptions(apiKeyId *string, request *ResetApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResetApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetApiKey"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/apikeys/" + dara.PercentEncode(dara.StringValue(apiKeyId)) + "/reset"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets an API key.
//
// Description:
//
// Only the API key value changes. The API key ID remains unchanged.
//
// @param request - ResetApiKeyRequest
//
// @return ResetApiKeyResponse
func (client *Client) ResetApiKey(apiKeyId *string, request *ResetApiKeyRequest) (_result *ResetApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResetApiKeyResponse{}
	_body, _err := client.ResetApiKeyWithOptions(apiKeyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Edits the information of an authentication credential API key.
//
// @param request - UpdateApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApiKeyResponse
func (client *Client) UpdateApiKeyWithOptions(apiKeyId *string, request *UpdateApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Auth) {
		body["auth"] = request.Auth
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApiKey"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/apikeys/" + dara.PercentEncode(dara.StringValue(apiKeyId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits the information of an authentication credential API key.
//
// @param request - UpdateApiKeyRequest
//
// @return UpdateApiKeyResponse
func (client *Client) UpdateApiKey(apiKeyId *string, request *UpdateApiKeyRequest) (_result *UpdateApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateApiKeyResponse{}
	_body, _err := client.UpdateApiKeyWithOptions(apiKeyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
