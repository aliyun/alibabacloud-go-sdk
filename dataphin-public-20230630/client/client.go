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
		"cn-shenzhen": dara.String("dataphin-public.cn-shenzhen.aliyuncs.com"),
		"cn-shanghai": dara.String("dataphin-public.cn-shanghai.aliyuncs.com"),
		"cn-hangzhou": dara.String("dataphin-public.cn-hangzhou.aliyuncs.com"),
		"cn-chengdu":  dara.String("dataphin-public.cn-chengdu.aliyuncs.com"),
		"cn-beijing":  dara.String("dataphin-public.cn-beijing.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("dataphin-public"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds a regular member to a data service application. Only the application owner can perform this operation.
//
// Online version: v6.0.0.
//
// @param tmpReq - AddDataServiceAppMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDataServiceAppMemberResponse
func (client *Client) AddDataServiceAppMemberWithOptions(tmpReq *AddDataServiceAppMemberRequest, runtime *dara.RuntimeOptions) (_result *AddDataServiceAppMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddDataServiceAppMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddCommand) {
		request.AddCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddCommand, dara.String("AddCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddCommandShrink) {
		body["AddCommand"] = request.AddCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDataServiceAppMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDataServiceAppMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a regular member to a data service application. Only the application owner can perform this operation.
//
// Online version: v6.0.0.
//
// @param request - AddDataServiceAppMemberRequest
//
// @return AddDataServiceAppMemberResponse
func (client *Client) AddDataServiceAppMember(request *AddDataServiceAppMemberRequest) (_result *AddDataServiceAppMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDataServiceAppMemberResponse{}
	_body, _err := client.AddDataServiceAppMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds users to a data service project and assigns roles to them.
//
// @param tmpReq - AddDataServiceProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDataServiceProjectMemberResponse
func (client *Client) AddDataServiceProjectMemberWithOptions(tmpReq *AddDataServiceProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *AddDataServiceProjectMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddDataServiceProjectMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddCommand) {
		request.AddCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddCommand, dara.String("AddCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddCommandShrink) {
		body["AddCommand"] = request.AddCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDataServiceProjectMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDataServiceProjectMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds users to a data service project and assigns roles to them.
//
// @param request - AddDataServiceProjectMemberRequest
//
// @return AddDataServiceProjectMemberResponse
func (client *Client) AddDataServiceProjectMember(request *AddDataServiceProjectMemberRequest) (_result *AddDataServiceProjectMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDataServiceProjectMemberResponse{}
	_body, _err := client.AddDataServiceProjectMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds members to a project.
//
// @param tmpReq - AddProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddProjectMemberResponse
func (client *Client) AddProjectMemberWithOptions(tmpReq *AddProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *AddProjectMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddProjectMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddCommand) {
		request.AddCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddCommand, dara.String("AddCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddCommandShrink) {
		body["AddCommand"] = request.AddCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddProjectMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddProjectMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds members to a project.
//
// @param request - AddProjectMemberRequest
//
// @return AddProjectMemberResponse
func (client *Client) AddProjectMember(request *AddProjectMemberRequest) (_result *AddProjectMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddProjectMemberResponse{}
	_body, _err := client.AddProjectMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers data lineage. Available since version v5.4.0.
//
// @param tmpReq - AddRegisterLineageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRegisterLineageResponse
func (client *Client) AddRegisterLineageWithOptions(tmpReq *AddRegisterLineageRequest, runtime *dara.RuntimeOptions) (_result *AddRegisterLineageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddRegisterLineageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddRegisterLineageCommand) {
		request.AddRegisterLineageCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddRegisterLineageCommand, dara.String("AddRegisterLineageCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddRegisterLineageCommandShrink) {
		body["AddRegisterLineageCommand"] = request.AddRegisterLineageCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRegisterLineage"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRegisterLineageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers data lineage. Available since version v5.4.0.
//
// @param request - AddRegisterLineageRequest
//
// @return AddRegisterLineageResponse
func (client *Client) AddRegisterLineage(request *AddRegisterLineageRequest) (_result *AddRegisterLineageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRegisterLineageResponse{}
	_body, _err := client.AddRegisterLineageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds users to a tenant in batches. Only super administrators (SuperAdmin) and system administrators can invoke this API operation.
//
// @param tmpReq - AddTenantMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTenantMembersResponse
func (client *Client) AddTenantMembersWithOptions(tmpReq *AddTenantMembersRequest, runtime *dara.RuntimeOptions) (_result *AddTenantMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddTenantMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddCommand) {
		request.AddCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddCommand, dara.String("AddCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddCommandShrink) {
		body["AddCommand"] = request.AddCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTenantMembers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTenantMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds users to a tenant in batches. Only super administrators (SuperAdmin) and system administrators can invoke this API operation.
//
// @param request - AddTenantMembersRequest
//
// @return AddTenantMembersResponse
func (client *Client) AddTenantMembers(request *AddTenantMembersRequest) (_result *AddTenantMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddTenantMembersResponse{}
	_body, _err := client.AddTenantMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tenant members by using original user identities.
//
// @param tmpReq - AddTenantMembersBySourceUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTenantMembersBySourceUserResponse
func (client *Client) AddTenantMembersBySourceUserWithOptions(tmpReq *AddTenantMembersBySourceUserRequest, runtime *dara.RuntimeOptions) (_result *AddTenantMembersBySourceUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddTenantMembersBySourceUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddCommand) {
		request.AddCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddCommand, dara.String("AddCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddCommandShrink) {
		body["AddCommand"] = request.AddCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTenantMembersBySourceUser"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTenantMembersBySourceUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tenant members by using original user identities.
//
// @param request - AddTenantMembersBySourceUserRequest
//
// @return AddTenantMembersBySourceUserResponse
func (client *Client) AddTenantMembersBySourceUser(request *AddTenantMembersBySourceUserRequest) (_result *AddTenantMembersBySourceUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddTenantMembersBySourceUserResponse{}
	_body, _err := client.AddTenantMembersBySourceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds members to a user group.
//
// @param tmpReq - AddUserGroupMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserGroupMemberResponse
func (client *Client) AddUserGroupMemberWithOptions(tmpReq *AddUserGroupMemberRequest, runtime *dara.RuntimeOptions) (_result *AddUserGroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddUserGroupMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddCommand) {
		request.AddCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddCommand, dara.String("AddCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddCommandShrink) {
		body["AddCommand"] = request.AddCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserGroupMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserGroupMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds members to a user group.
//
// @param request - AddUserGroupMemberRequest
//
// @return AddUserGroupMemberResponse
func (client *Client) AddUserGroupMember(request *AddUserGroupMemberRequest) (_result *AddUserGroupMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUserGroupMemberResponse{}
	_body, _err := client.AddUserGroupMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Applies for API permissions.
//
// @param tmpReq - ApplyDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyDataServiceApiResponse
func (client *Client) ApplyDataServiceApiWithOptions(tmpReq *ApplyDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *ApplyDataServiceApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ApplyDataServiceApiShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApplyCommand) {
		request.ApplyCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplyCommand, dara.String("ApplyCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplyCommandShrink) {
		body["ApplyCommand"] = request.ApplyCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyDataServiceApi"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyDataServiceApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies for API permissions.
//
// @param request - ApplyDataServiceApiRequest
//
// @return ApplyDataServiceApiResponse
func (client *Client) ApplyDataServiceApi(request *ApplyDataServiceApiRequest) (_result *ApplyDataServiceApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyDataServiceApiResponse{}
	_body, _err := client.ApplyDataServiceApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Applies for application permissions.
//
// @param tmpReq - ApplyDataServiceAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyDataServiceAppResponse
func (client *Client) ApplyDataServiceAppWithOptions(tmpReq *ApplyDataServiceAppRequest, runtime *dara.RuntimeOptions) (_result *ApplyDataServiceAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ApplyDataServiceAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApplyCommand) {
		request.ApplyCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplyCommand, dara.String("ApplyCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplyCommandShrink) {
		body["ApplyCommand"] = request.ApplyCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyDataServiceApp"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyDataServiceAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies for application permissions.
//
// @param request - ApplyDataServiceAppRequest
//
// @return ApplyDataServiceAppResponse
func (client *Client) ApplyDataServiceApp(request *ApplyDataServiceAppRequest) (_result *ApplyDataServiceAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyDataServiceAppResponse{}
	_body, _err := client.ApplyDataServiceAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds specified quality rules to schedule settings.
//
// Release version: v5.4.2.
//
// @param tmpReq - AssignQualityRuleOfAllRuleScopeSchedulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignQualityRuleOfAllRuleScopeSchedulesResponse
func (client *Client) AssignQualityRuleOfAllRuleScopeSchedulesWithOptions(tmpReq *AssignQualityRuleOfAllRuleScopeSchedulesRequest, runtime *dara.RuntimeOptions) (_result *AssignQualityRuleOfAllRuleScopeSchedulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AssignCommand) {
		request.AssignCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssignCommand, dara.String("AssignCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AssignCommandShrink) {
		body["AssignCommand"] = request.AssignCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssignQualityRuleOfAllRuleScopeSchedules"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssignQualityRuleOfAllRuleScopeSchedulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds specified quality rules to schedule settings.
//
// Release version: v5.4.2.
//
// @param request - AssignQualityRuleOfAllRuleScopeSchedulesRequest
//
// @return AssignQualityRuleOfAllRuleScopeSchedulesResponse
func (client *Client) AssignQualityRuleOfAllRuleScopeSchedules(request *AssignQualityRuleOfAllRuleScopeSchedulesRequest) (_result *AssignQualityRuleOfAllRuleScopeSchedulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssignQualityRuleOfAllRuleScopeSchedulesResponse{}
	_body, _err := client.AssignQualityRuleOfAllRuleScopeSchedulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks the connectivity of a compute source.
//
// @param tmpReq - CheckComputeSourceConnectivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckComputeSourceConnectivityResponse
func (client *Client) CheckComputeSourceConnectivityWithOptions(tmpReq *CheckComputeSourceConnectivityRequest, runtime *dara.RuntimeOptions) (_result *CheckComputeSourceConnectivityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CheckComputeSourceConnectivityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CheckCommand) {
		request.CheckCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckCommand, dara.String("CheckCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckCommandShrink) {
		body["CheckCommand"] = request.CheckCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckComputeSourceConnectivity"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckComputeSourceConnectivityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the connectivity of a compute source.
//
// @param request - CheckComputeSourceConnectivityRequest
//
// @return CheckComputeSourceConnectivityResponse
func (client *Client) CheckComputeSourceConnectivity(request *CheckComputeSourceConnectivityRequest) (_result *CheckComputeSourceConnectivityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckComputeSourceConnectivityResponse{}
	_body, _err := client.CheckComputeSourceConnectivityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks the connectivity of an existing compute source by compute source ID.
//
// @param request - CheckComputeSourceConnectivityByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckComputeSourceConnectivityByIdResponse
func (client *Client) CheckComputeSourceConnectivityByIdWithOptions(request *CheckComputeSourceConnectivityByIdRequest, runtime *dara.RuntimeOptions) (_result *CheckComputeSourceConnectivityByIdResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckComputeSourceConnectivityById"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckComputeSourceConnectivityByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the connectivity of an existing compute source by compute source ID.
//
// @param request - CheckComputeSourceConnectivityByIdRequest
//
// @return CheckComputeSourceConnectivityByIdResponse
func (client *Client) CheckComputeSourceConnectivityById(request *CheckComputeSourceConnectivityByIdRequest) (_result *CheckComputeSourceConnectivityByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckComputeSourceConnectivityByIdResponse{}
	_body, _err := client.CheckComputeSourceConnectivityByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks the connectivity of a data source.
//
// @param tmpReq - CheckDataSourceConnectivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDataSourceConnectivityResponse
func (client *Client) CheckDataSourceConnectivityWithOptions(tmpReq *CheckDataSourceConnectivityRequest, runtime *dara.RuntimeOptions) (_result *CheckDataSourceConnectivityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CheckDataSourceConnectivityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CheckCommand) {
		request.CheckCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckCommand, dara.String("CheckCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckCommandShrink) {
		body["CheckCommand"] = request.CheckCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDataSourceConnectivity"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDataSourceConnectivityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the connectivity of a data source.
//
// @param request - CheckDataSourceConnectivityRequest
//
// @return CheckDataSourceConnectivityResponse
func (client *Client) CheckDataSourceConnectivity(request *CheckDataSourceConnectivityRequest) (_result *CheckDataSourceConnectivityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckDataSourceConnectivityResponse{}
	_body, _err := client.CheckDataSourceConnectivityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks the connectivity of a data source.
//
// @param request - CheckDataSourceConnectivityByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDataSourceConnectivityByIdResponse
func (client *Client) CheckDataSourceConnectivityByIdWithOptions(request *CheckDataSourceConnectivityByIdRequest, runtime *dara.RuntimeOptions) (_result *CheckDataSourceConnectivityByIdResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDataSourceConnectivityById"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDataSourceConnectivityByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the connectivity of a data source.
//
// @param request - CheckDataSourceConnectivityByIdRequest
//
// @return CheckDataSourceConnectivityByIdResponse
func (client *Client) CheckDataSourceConnectivityById(request *CheckDataSourceConnectivityByIdRequest) (_result *CheckDataSourceConnectivityByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckDataSourceConnectivityByIdResponse{}
	_body, _err := client.CheckDataSourceConnectivityByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether a project has data dependencies such as tasks.
//
// @param request - CheckProjectHasDependencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckProjectHasDependencyResponse
func (client *Client) CheckProjectHasDependencyWithOptions(request *CheckProjectHasDependencyRequest, runtime *dara.RuntimeOptions) (_result *CheckProjectHasDependencyResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckProjectHasDependency"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckProjectHasDependencyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a project has data dependencies such as tasks.
//
// @param request - CheckProjectHasDependencyRequest
//
// @return CheckProjectHasDependencyResponse
func (client *Client) CheckProjectHasDependency(request *CheckProjectHasDependencyRequest) (_result *CheckProjectHasDependencyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckProjectHasDependencyResponse{}
	_body, _err := client.CheckProjectHasDependencyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether a user has the permission on a specified resource.
//
// @param tmpReq - CheckResourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckResourcePermissionResponse
func (client *Client) CheckResourcePermissionWithOptions(tmpReq *CheckResourcePermissionRequest, runtime *dara.RuntimeOptions) (_result *CheckResourcePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CheckResourcePermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CheckCommand) {
		request.CheckCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckCommand, dara.String("CheckCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckCommandShrink) {
		body["CheckCommand"] = request.CheckCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckResourcePermission"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckResourcePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a user has the permission on a specified resource.
//
// @param request - CheckResourcePermissionRequest
//
// @return CheckResourcePermissionResponse
func (client *Client) CheckResourcePermission(request *CheckResourcePermissionRequest) (_result *CheckResourcePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckResourcePermissionResponse{}
	_body, _err := client.CheckResourcePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an ad hoc query file.
//
// @param tmpReq - CreateAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAdHocFileResponse
func (client *Client) CreateAdHocFileWithOptions(tmpReq *CreateAdHocFileRequest, runtime *dara.RuntimeOptions) (_result *CreateAdHocFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAdHocFileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAdHocFile"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAdHocFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an ad hoc query file.
//
// @param request - CreateAdHocFileRequest
//
// @return CreateAdHocFileResponse
func (client *Client) CreateAdHocFile(request *CreateAdHocFileRequest) (_result *CreateAdHocFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAdHocFileResponse{}
	_body, _err := client.CreateAdHocFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a batch task.
//
// @param tmpReq - CreateBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBatchTaskResponse
func (client *Client) CreateBatchTaskWithOptions(tmpReq *CreateBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateBatchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateBatchTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBatchTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBatchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a batch task.
//
// @param request - CreateBatchTaskRequest
//
// @return CreateBatchTaskResponse
func (client *Client) CreateBatchTask(request *CreateBatchTaskRequest) (_result *CreateBatchTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBatchTaskResponse{}
	_body, _err := client.CreateBatchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a business entity.
//
// @param tmpReq - CreateBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBizEntityResponse
func (client *Client) CreateBizEntityWithOptions(tmpReq *CreateBizEntityRequest, runtime *dara.RuntimeOptions) (_result *CreateBizEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateBizEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBizEntity"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBizEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a business entity.
//
// @param request - CreateBizEntityRequest
//
// @return CreateBizEntityResponse
func (client *Client) CreateBizEntity(request *CreateBizEntityRequest) (_result *CreateBizEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBizEntityResponse{}
	_body, _err := client.CreateBizEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a business metric.
//
// Release version: v5.5.0.
//
// @param tmpReq - CreateBizMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBizMetricResponse
func (client *Client) CreateBizMetricWithOptions(tmpReq *CreateBizMetricRequest, runtime *dara.RuntimeOptions) (_result *CreateBizMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateBizMetricShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateBizMetricCommand) {
		request.CreateBizMetricCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateBizMetricCommand, dara.String("CreateBizMetricCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateBizMetricCommandShrink) {
		body["CreateBizMetricCommand"] = request.CreateBizMetricCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBizMetric"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBizMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a business metric.
//
// Release version: v5.5.0.
//
// @param request - CreateBizMetricRequest
//
// @return CreateBizMetricResponse
func (client *Client) CreateBizMetric(request *CreateBizMetricRequest) (_result *CreateBizMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBizMetricResponse{}
	_body, _err := client.CreateBizMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data domain.
//
// @param tmpReq - CreateBizUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBizUnitResponse
func (client *Client) CreateBizUnitWithOptions(tmpReq *CreateBizUnitRequest, runtime *dara.RuntimeOptions) (_result *CreateBizUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateBizUnitShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBizUnit"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBizUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data domain.
//
// @param request - CreateBizUnitRequest
//
// @return CreateBizUnitResponse
func (client *Client) CreateBizUnit(request *CreateBizUnitRequest) (_result *CreateBizUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBizUnitResponse{}
	_body, _err := client.CreateBizUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a compute source. Business unit administrators and project administrators have permissions to perform this operation.
//
// @param tmpReq - CreateComputeSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComputeSourceResponse
func (client *Client) CreateComputeSourceWithOptions(tmpReq *CreateComputeSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateComputeSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateComputeSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComputeSource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComputeSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a compute source. Business unit administrators and project administrators have permissions to perform this operation.
//
// @param request - CreateComputeSourceRequest
//
// @return CreateComputeSourceResponse
func (client *Client) CreateComputeSource(request *CreateComputeSourceRequest) (_result *CreateComputeSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateComputeSourceResponse{}
	_body, _err := client.CreateComputeSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data domain.
//
// @param tmpReq - CreateDataDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataDomainResponse
func (client *Client) CreateDataDomainWithOptions(tmpReq *CreateDataDomainRequest, runtime *dara.RuntimeOptions) (_result *CreateDataDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataDomainShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataDomain"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data domain.
//
// @param request - CreateDataDomainRequest
//
// @return CreateDataDomainResponse
func (client *Client) CreateDataDomain(request *CreateDataDomainRequest) (_result *CreateDataDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataDomainResponse{}
	_body, _err := client.CreateDataDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data service API and submits it.
//
// @param tmpReq - CreateDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataServiceApiResponse
func (client *Client) CreateDataServiceApiWithOptions(tmpReq *CreateDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *CreateDataServiceApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataServiceApiShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataServiceApi"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataServiceApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data service API and submits it.
//
// @param request - CreateDataServiceApiRequest
//
// @return CreateDataServiceApiResponse
func (client *Client) CreateDataServiceApi(request *CreateDataServiceApiRequest) (_result *CreateDataServiceApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataServiceApiResponse{}
	_body, _err := client.CreateDataServiceApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data service application. Only super administrators or system administrators can perform this operation.
//
// Online version: v6.0.0.
//
// @param tmpReq - CreateDataServiceAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataServiceAppResponse
func (client *Client) CreateDataServiceAppWithOptions(tmpReq *CreateDataServiceAppRequest, runtime *dara.RuntimeOptions) (_result *CreateDataServiceAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataServiceAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataServiceApp"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataServiceAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data service application. Only super administrators or system administrators can perform this operation.
//
// Online version: v6.0.0.
//
// @param request - CreateDataServiceAppRequest
//
// @return CreateDataServiceAppResponse
func (client *Client) CreateDataServiceApp(request *CreateDataServiceAppRequest) (_result *CreateDataServiceAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataServiceAppResponse{}
	_body, _err := client.CreateDataServiceAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data service application group. Only superusers and system administrators can perform this operation.
//
// Online version: v6.0.0.
//
// @param request - CreateDataServiceAppGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataServiceAppGroupResponse
func (client *Client) CreateDataServiceAppGroupWithOptions(request *CreateDataServiceAppGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDataServiceAppGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataServiceAppGroup"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataServiceAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data service application group. Only superusers and system administrators can perform this operation.
//
// Online version: v6.0.0.
//
// @param request - CreateDataServiceAppGroupRequest
//
// @return CreateDataServiceAppGroupResponse
func (client *Client) CreateDataServiceAppGroup(request *CreateDataServiceAppGroupRequest) (_result *CreateDataServiceAppGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataServiceAppGroupResponse{}
	_body, _err := client.CreateDataServiceAppGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create Data Source: Tenant administrators, data administrators, business unit administrators, project administrators, and operations administrators have permission to perform this operation.
//
// @param tmpReq - CreateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSourceWithOptions(tmpReq *CreateDataSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataSource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create Data Source: Tenant administrators, data administrators, business unit administrators, project administrators, and operations administrators have permission to perform this operation.
//
// @param request - CreateDataSourceRequest
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSource(request *CreateDataSourceRequest) (_result *CreateDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CreateDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a menu tree directory. This operation supports features such as compute nodes, data integration, and synchronization tasks.
//
// @param tmpReq - CreateDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDirectoryResponse
func (client *Client) CreateDirectoryWithOptions(tmpReq *CreateDirectoryRequest, runtime *dara.RuntimeOptions) (_result *CreateDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDirectoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDirectory"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a menu tree directory. This operation supports features such as compute nodes, data integration, and synchronization tasks.
//
// @param request - CreateDirectoryRequest
//
// @return CreateDirectoryResponse
func (client *Client) CreateDirectory(request *CreateDirectoryRequest) (_result *CreateDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDirectoryResponse{}
	_body, _err := client.CreateDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// General-purpose backfill API that supports both list-mode and bulk-mode backfill:
//
// 1. Backfill instances will be generated and executed, affecting the data output of related tables.
//
// 2. Task execution will incur computing costs and storage costs.
//
// @param tmpReq - CreateNodeSupplementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeSupplementResponse
func (client *Client) CreateNodeSupplementWithOptions(tmpReq *CreateNodeSupplementRequest, runtime *dara.RuntimeOptions) (_result *CreateNodeSupplementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateNodeSupplementShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNodeSupplement"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNodeSupplementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// General-purpose backfill API that supports both list-mode and bulk-mode backfill:
//
// 1. Backfill instances will be generated and executed, affecting the data output of related tables.
//
// 2. Task execution will incur computing costs and storage costs.
//
// @param request - CreateNodeSupplementRequest
//
// @return CreateNodeSupplementResponse
func (client *Client) CreateNodeSupplement(request *CreateNodeSupplementRequest) (_result *CreateNodeSupplementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNodeSupplementResponse{}
	_body, _err := client.CreateNodeSupplementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create an integration pipeline/unstructured workflow task.
//
// @param tmpReq - CreatePipelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineResponse
func (client *Client) CreatePipelineWithOptions(tmpReq *CreatePipelineRequest, runtime *dara.RuntimeOptions) (_result *CreatePipelineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePipelineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Context) {
		request.ContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Context, dara.String("Context"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextShrink) {
		body["Context"] = request.ContextShrink
	}

	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePipeline"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an integration pipeline/unstructured workflow task.
//
// @param request - CreatePipelineRequest
//
// @return CreatePipelineResponse
func (client *Client) CreatePipeline(request *CreatePipelineRequest) (_result *CreatePipelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePipelineResponse{}
	_body, _err := client.CreatePipelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously create a pipeline/unstructured workflow.
//
// @param tmpReq - CreatePipelineByAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineByAsyncResponse
func (client *Client) CreatePipelineByAsyncWithOptions(tmpReq *CreatePipelineByAsyncRequest, runtime *dara.RuntimeOptions) (_result *CreatePipelineByAsyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePipelineByAsyncShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Context) {
		request.ContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Context, dara.String("Context"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextShrink) {
		body["Context"] = request.ContextShrink
	}

	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePipelineByAsync"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePipelineByAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously create a pipeline/unstructured workflow.
//
// @param request - CreatePipelineByAsyncRequest
//
// @return CreatePipelineByAsyncResponse
func (client *Client) CreatePipelineByAsync(request *CreatePipelineByAsyncRequest) (_result *CreatePipelineByAsyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePipelineByAsyncResponse{}
	_body, _err := client.CreatePipelineByAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data integration task. Note: This operation is deprecated starting from Dataphin v5.3.1. Use CreatePipeline instead.
//
// @param tmpReq - CreatePipelineNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineNodeResponse
func (client *Client) CreatePipelineNodeWithOptions(tmpReq *CreatePipelineNodeRequest, runtime *dara.RuntimeOptions) (_result *CreatePipelineNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePipelineNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreatePipelineNodeCommand) {
		request.CreatePipelineNodeCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreatePipelineNodeCommand, dara.String("CreatePipelineNodeCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreatePipelineNodeCommandShrink) {
		body["CreatePipelineNodeCommand"] = request.CreatePipelineNodeCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePipelineNode"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePipelineNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data integration task. Note: This operation is deprecated starting from Dataphin v5.3.1. Use CreatePipeline instead.
//
// @param request - CreatePipelineNodeRequest
//
// @return CreatePipelineNodeResponse
func (client *Client) CreatePipelineNode(request *CreatePipelineNodeRequest) (_result *CreatePipelineNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePipelineNodeResponse{}
	_body, _err := client.CreatePipelineNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a resource file.
//
// @param tmpReq - CreateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceResponse
func (client *Client) CreateResourceWithOptions(tmpReq *CreateResourceRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a resource file.
//
// @param request - CreateResourceRequest
//
// @return CreateResourceResponse
func (client *Client) CreateResource(request *CreateResourceRequest) (_result *CreateResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateResourceResponse{}
	_body, _err := client.CreateResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a row-level permission.
//
// Description:
//
// You can query detailed information about published APIs based on the appKey.
//
// @param tmpReq - CreateRowPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRowPermissionResponse
func (client *Client) CreateRowPermissionWithOptions(tmpReq *CreateRowPermissionRequest, runtime *dara.RuntimeOptions) (_result *CreateRowPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateRowPermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateRowPermissionCommand) {
		request.CreateRowPermissionCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateRowPermissionCommand, dara.String("CreateRowPermissionCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateRowPermissionCommandShrink) {
		body["CreateRowPermissionCommand"] = request.CreateRowPermissionCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRowPermission"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRowPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a row-level permission.
//
// Description:
//
// You can query detailed information about published APIs based on the appKey.
//
// @param request - CreateRowPermissionRequest
//
// @return CreateRowPermissionResponse
func (client *Client) CreateRowPermission(request *CreateRowPermissionRequest) (_result *CreateRowPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRowPermissionResponse{}
	_body, _err := client.CreateRowPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data classification. Available since v5.4.2.
//
// @param tmpReq - CreateSecurityClassifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecurityClassifyResponse
func (client *Client) CreateSecurityClassifyWithOptions(tmpReq *CreateSecurityClassifyRequest, runtime *dara.RuntimeOptions) (_result *CreateSecurityClassifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSecurityClassifyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSecurityClassify"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecurityClassifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data classification. Available since v5.4.2.
//
// @param request - CreateSecurityClassifyRequest
//
// @return CreateSecurityClassifyResponse
func (client *Client) CreateSecurityClassify(request *CreateSecurityClassifyRequest) (_result *CreateSecurityClassifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSecurityClassifyResponse{}
	_body, _err := client.CreateSecurityClassifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data classification folder. Available since v5.4.2.
//
// @param tmpReq - CreateSecurityClassifyCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecurityClassifyCatalogResponse
func (client *Client) CreateSecurityClassifyCatalogWithOptions(tmpReq *CreateSecurityClassifyCatalogRequest, runtime *dara.RuntimeOptions) (_result *CreateSecurityClassifyCatalogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSecurityClassifyCatalogShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSecurityClassifyCatalog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecurityClassifyCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data classification folder. Available since v5.4.2.
//
// @param request - CreateSecurityClassifyCatalogRequest
//
// @return CreateSecurityClassifyCatalogResponse
func (client *Client) CreateSecurityClassifyCatalog(request *CreateSecurityClassifyCatalogRequest) (_result *CreateSecurityClassifyCatalogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSecurityClassifyCatalogResponse{}
	_body, _err := client.CreateSecurityClassifyCatalogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a security identification result.
//
// Release version: v5.4.2.
//
// @param tmpReq - CreateSecurityIdentifyResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecurityIdentifyResultResponse
func (client *Client) CreateSecurityIdentifyResultWithOptions(tmpReq *CreateSecurityIdentifyResultRequest, runtime *dara.RuntimeOptions) (_result *CreateSecurityIdentifyResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSecurityIdentifyResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSecurityIdentifyResult"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecurityIdentifyResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a security identification result.
//
// Release version: v5.4.2.
//
// @param request - CreateSecurityIdentifyResultRequest
//
// @return CreateSecurityIdentifyResultResponse
func (client *Client) CreateSecurityIdentifyResult(request *CreateSecurityIdentifyResultRequest) (_result *CreateSecurityIdentifyResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSecurityIdentifyResultResponse{}
	_body, _err := client.CreateSecurityIdentifyResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data classification level. Available since v5.4.2.
//
// @param tmpReq - CreateSecurityLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecurityLevelResponse
func (client *Client) CreateSecurityLevelWithOptions(tmpReq *CreateSecurityLevelRequest, runtime *dara.RuntimeOptions) (_result *CreateSecurityLevelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSecurityLevelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSecurityLevel"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecurityLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data classification level. Available since v5.4.2.
//
// @param request - CreateSecurityLevelRequest
//
// @return CreateSecurityLevelResponse
func (client *Client) CreateSecurityLevel(request *CreateSecurityLevelRequest) (_result *CreateSecurityLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSecurityLevelResponse{}
	_body, _err := client.CreateSecurityLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a standard.
//
// Release version: v5.4.2.
//
// @param tmpReq - CreateStandardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStandardResponse
func (client *Client) CreateStandardWithOptions(tmpReq *CreateStandardRequest, runtime *dara.RuntimeOptions) (_result *CreateStandardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStandardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStandard"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStandardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a standard.
//
// Release version: v5.4.2.
//
// @param request - CreateStandardRequest
//
// @return CreateStandardResponse
func (client *Client) CreateStandard(request *CreateStandardRequest) (_result *CreateStandardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStandardResponse{}
	_body, _err := client.CreateStandardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data standard lookup table.
//
// Release version: v5.4.2.
//
// @param tmpReq - CreateStandardLookupTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStandardLookupTableResponse
func (client *Client) CreateStandardLookupTableWithOptions(tmpReq *CreateStandardLookupTableRequest, runtime *dara.RuntimeOptions) (_result *CreateStandardLookupTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStandardLookupTableShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStandardLookupTable"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStandardLookupTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data standard lookup table.
//
// Release version: v5.4.2.
//
// @param request - CreateStandardLookupTableRequest
//
// @return CreateStandardLookupTableResponse
func (client *Client) CreateStandardLookupTable(request *CreateStandardLookupTableRequest) (_result *CreateStandardLookupTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStandardLookupTableResponse{}
	_body, _err := client.CreateStandardLookupTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates standard mapping relationships, including valid mappings and invalid mappings.
//
// Release version: v5.4.2.
//
// @param tmpReq - CreateStandardMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStandardMappingResponse
func (client *Client) CreateStandardMappingWithOptions(tmpReq *CreateStandardMappingRequest, runtime *dara.RuntimeOptions) (_result *CreateStandardMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStandardMappingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStandardMapping"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStandardMappingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates standard mapping relationships, including valid mappings and invalid mappings.
//
// Release version: v5.4.2.
//
// @param request - CreateStandardMappingRequest
//
// @return CreateStandardMappingResponse
func (client *Client) CreateStandardMapping(request *CreateStandardMappingRequest) (_result *CreateStandardMappingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStandardMappingResponse{}
	_body, _err := client.CreateStandardMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a standard association. Release version: v5.4.2.
//
// @param tmpReq - CreateStandardRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStandardRelationsResponse
func (client *Client) CreateStandardRelationsWithOptions(tmpReq *CreateStandardRelationsRequest, runtime *dara.RuntimeOptions) (_result *CreateStandardRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStandardRelationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStandardRelations"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStandardRelationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a standard association. Release version: v5.4.2.
//
// @param request - CreateStandardRelationsRequest
//
// @return CreateStandardRelationsResponse
func (client *Client) CreateStandardRelations(request *CreateStandardRelationsRequest) (_result *CreateStandardRelationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStandardRelationsResponse{}
	_body, _err := client.CreateStandardRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a standard set.
//
// Available since: v5.4.2.
//
// @param tmpReq - CreateStandardSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStandardSetResponse
func (client *Client) CreateStandardSetWithOptions(tmpReq *CreateStandardSetRequest, runtime *dara.RuntimeOptions) (_result *CreateStandardSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStandardSetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStandardSet"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStandardSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a standard set.
//
// Available since: v5.4.2.
//
// @param request - CreateStandardSetRequest
//
// @return CreateStandardSetResponse
func (client *Client) CreateStandardSet(request *CreateStandardSetRequest) (_result *CreateStandardSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStandardSetResponse{}
	_body, _err := client.CreateStandardSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a data standard template.
//
// Release version: v5.4.2.
//
// @param tmpReq - CreateStandardTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStandardTemplateResponse
func (client *Client) CreateStandardTemplateWithOptions(tmpReq *CreateStandardTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateStandardTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStandardTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStandardTemplate"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStandardTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a data standard template.
//
// Release version: v5.4.2.
//
// @param request - CreateStandardTemplateRequest
//
// @return CreateStandardTemplateResponse
func (client *Client) CreateStandardTemplate(request *CreateStandardTemplateRequest) (_result *CreateStandardTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStandardTemplateResponse{}
	_body, _err := client.CreateStandardTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data standard root word.
//
// Release version: v5.4.2.
//
// @param tmpReq - CreateStandardWordRootRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStandardWordRootResponse
func (client *Client) CreateStandardWordRootWithOptions(tmpReq *CreateStandardWordRootRequest, runtime *dara.RuntimeOptions) (_result *CreateStandardWordRootResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStandardWordRootShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStandardWordRoot"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStandardWordRootResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data standard root word.
//
// Release version: v5.4.2.
//
// @param request - CreateStandardWordRootRequest
//
// @return CreateStandardWordRootResponse
func (client *Client) CreateStandardWordRoot(request *CreateStandardWordRootRequest) (_result *CreateStandardWordRootResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStandardWordRootResponse{}
	_body, _err := client.CreateStandardWordRootWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a stream-batch integrated node.
//
// @param tmpReq - CreateStreamBatchJobMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStreamBatchJobMappingResponse
func (client *Client) CreateStreamBatchJobMappingWithOptions(tmpReq *CreateStreamBatchJobMappingRequest, runtime *dara.RuntimeOptions) (_result *CreateStreamBatchJobMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStreamBatchJobMappingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StreamBatchJobMappingCreateCommand) {
		request.StreamBatchJobMappingCreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamBatchJobMappingCreateCommand, dara.String("StreamBatchJobMappingCreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.StreamBatchJobMappingCreateCommandShrink) {
		body["StreamBatchJobMappingCreateCommand"] = request.StreamBatchJobMappingCreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStreamBatchJobMapping"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStreamBatchJobMappingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a stream-batch integrated node.
//
// @param request - CreateStreamBatchJobMappingRequest
//
// @return CreateStreamBatchJobMappingResponse
func (client *Client) CreateStreamBatchJobMapping(request *CreateStreamBatchJobMappingRequest) (_result *CreateStreamBatchJobMappingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStreamBatchJobMappingResponse{}
	_body, _err := client.CreateStreamBatchJobMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a user-defined function.
//
// @param tmpReq - CreateUdfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUdfResponse
func (client *Client) CreateUdfWithOptions(tmpReq *CreateUdfRequest, runtime *dara.RuntimeOptions) (_result *CreateUdfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateUdfShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUdf"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUdfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user-defined function.
//
// @param request - CreateUdfRequest
//
// @return CreateUdfResponse
func (client *Client) CreateUdf(request *CreateUdfRequest) (_result *CreateUdfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUdfResponse{}
	_body, _err := client.CreateUdfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a user group.
//
// @param tmpReq - CreateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroupWithOptions(tmpReq *CreateUserGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateUserGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserGroup"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user group.
//
// @param request - CreateUserGroupRequest
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroup(request *CreateUserGroupRequest) (_result *CreateUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CreateUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an ad hoc query file from the menu tree.
//
// @param request - DeleteAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAdHocFileResponse
func (client *Client) DeleteAdHocFileWithOptions(request *DeleteAdHocFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteAdHocFileResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAdHocFile"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAdHocFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an ad hoc query file from the menu tree.
//
// @param request - DeleteAdHocFileRequest
//
// @return DeleteAdHocFileResponse
func (client *Client) DeleteAdHocFile(request *DeleteAdHocFileRequest) (_result *DeleteAdHocFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAdHocFileResponse{}
	_body, _err := client.DeleteAdHocFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a batch task. If the node has not been offlined, you must offline it before deleting it.
//
// @param tmpReq - DeleteBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBatchTaskResponse
func (client *Client) DeleteBatchTaskWithOptions(tmpReq *DeleteBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteBatchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteBatchTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteCommand) {
		request.DeleteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteCommand, dara.String("DeleteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCommandShrink) {
		body["DeleteCommand"] = request.DeleteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBatchTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBatchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a batch task. If the node has not been offlined, you must offline it before deleting it.
//
// @param request - DeleteBatchTaskRequest
//
// @return DeleteBatchTaskResponse
func (client *Client) DeleteBatchTask(request *DeleteBatchTaskRequest) (_result *DeleteBatchTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBatchTaskResponse{}
	_body, _err := client.DeleteBatchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a business entity.
//
// @param request - DeleteBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBizEntityResponse
func (client *Client) DeleteBizEntityWithOptions(request *DeleteBizEntityRequest, runtime *dara.RuntimeOptions) (_result *DeleteBizEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizUnitId) {
		query["BizUnitId"] = request.BizUnitId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBizEntity"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBizEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a business entity.
//
// @param request - DeleteBizEntityRequest
//
// @return DeleteBizEntityResponse
func (client *Client) DeleteBizEntity(request *DeleteBizEntityRequest) (_result *DeleteBizEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBizEntityResponse{}
	_body, _err := client.DeleteBizEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a business metric.
//
// Release version: v5.5.0.
//
// @param tmpReq - DeleteBizMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBizMetricResponse
func (client *Client) DeleteBizMetricWithOptions(tmpReq *DeleteBizMetricRequest, runtime *dara.RuntimeOptions) (_result *DeleteBizMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteBizMetricShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteBizMetricCommand) {
		request.DeleteBizMetricCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteBizMetricCommand, dara.String("DeleteBizMetricCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteBizMetricCommandShrink) {
		body["DeleteBizMetricCommand"] = request.DeleteBizMetricCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBizMetric"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBizMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a business metric.
//
// Release version: v5.5.0.
//
// @param request - DeleteBizMetricRequest
//
// @return DeleteBizMetricResponse
func (client *Client) DeleteBizMetric(request *DeleteBizMetricRequest) (_result *DeleteBizMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBizMetricResponse{}
	_body, _err := client.DeleteBizMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data domain.
//
// @param request - DeleteBizUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBizUnitResponse
func (client *Client) DeleteBizUnitWithOptions(request *DeleteBizUnitRequest, runtime *dara.RuntimeOptions) (_result *DeleteBizUnitResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBizUnit"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBizUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data domain.
//
// @param request - DeleteBizUnitRequest
//
// @return DeleteBizUnitResponse
func (client *Client) DeleteBizUnit(request *DeleteBizUnitRequest) (_result *DeleteBizUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBizUnitResponse{}
	_body, _err := client.DeleteBizUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a compute source.
//
// @param request - DeleteComputeSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComputeSourceResponse
func (client *Client) DeleteComputeSourceWithOptions(request *DeleteComputeSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteComputeSourceResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComputeSource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteComputeSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a compute source.
//
// @param request - DeleteComputeSourceRequest
//
// @return DeleteComputeSourceResponse
func (client *Client) DeleteComputeSource(request *DeleteComputeSourceRequest) (_result *DeleteComputeSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteComputeSourceResponse{}
	_body, _err := client.DeleteComputeSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a subject domain.
//
// @param request - DeleteDataDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataDomainResponse
func (client *Client) DeleteDataDomainWithOptions(request *DeleteDataDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizUnitId) {
		query["BizUnitId"] = request.BizUnitId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataDomain"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a subject domain.
//
// @param request - DeleteDataDomainRequest
//
// @return DeleteDataDomainResponse
func (client *Client) DeleteDataDomain(request *DeleteDataDomainRequest) (_result *DeleteDataDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataDomainResponse{}
	_body, _err := client.DeleteDataDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data service application. Only superusers, system administrators, or application owners can perform this operation.
//
// Online version: v6.0.0.
//
// @param request - DeleteDataServiceAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataServiceAppResponse
func (client *Client) DeleteDataServiceAppWithOptions(request *DeleteDataServiceAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataServiceAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataServiceApp"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataServiceAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data service application. Only superusers, system administrators, or application owners can perform this operation.
//
// Online version: v6.0.0.
//
// @param request - DeleteDataServiceAppRequest
//
// @return DeleteDataServiceAppResponse
func (client *Client) DeleteDataServiceApp(request *DeleteDataServiceAppRequest) (_result *DeleteDataServiceAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataServiceAppResponse{}
	_body, _err := client.DeleteDataServiceAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data service application group. Only superusers and system administrators can perform this operation.
//
// Online version: v6.0.0.
//
// @param request - DeleteDataServiceAppGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataServiceAppGroupResponse
func (client *Client) DeleteDataServiceAppGroupWithOptions(request *DeleteDataServiceAppGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataServiceAppGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataServiceAppGroup"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataServiceAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data service application group. Only superusers and system administrators can perform this operation.
//
// Online version: v6.0.0.
//
// @param request - DeleteDataServiceAppGroupRequest
//
// @return DeleteDataServiceAppGroupResponse
func (client *Client) DeleteDataServiceAppGroup(request *DeleteDataServiceAppGroupRequest) (_result *DeleteDataServiceAppGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataServiceAppGroupResponse{}
	_body, _err := client.DeleteDataServiceAppGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data source.
//
// @param tmpReq - DeleteDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSourceWithOptions(tmpReq *DeleteDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteDataSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteCommand) {
		request.DeleteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteCommand, dara.String("DeleteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCommandShrink) {
		body["DeleteCommand"] = request.DeleteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataSource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data source.
//
// @param request - DeleteDataSourceRequest
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSource(request *DeleteDataSourceRequest) (_result *DeleteDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.DeleteDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a file directory from the menu tree.
//
// @param request - DeleteDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDirectoryResponse
func (client *Client) DeleteDirectoryWithOptions(request *DeleteDirectoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteDirectoryResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDirectory"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a file directory from the menu tree.
//
// @param request - DeleteDirectoryRequest
//
// @return DeleteDirectoryResponse
func (client *Client) DeleteDirectory(request *DeleteDirectoryRequest) (_result *DeleteDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDirectoryResponse{}
	_body, _err := client.DeleteDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes quality rule objects in batches.
//
// Release version: v5.4.2.
//
// @param tmpReq - DeleteQualityRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualityRulesResponse
func (client *Client) DeleteQualityRulesWithOptions(tmpReq *DeleteQualityRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteQualityRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteQualityRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteCommand) {
		request.DeleteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteCommand, dara.String("DeleteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCommandShrink) {
		body["DeleteCommand"] = request.DeleteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQualityRules"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQualityRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes quality rule objects in batches.
//
// Release version: v5.4.2.
//
// @param request - DeleteQualityRulesRequest
//
// @return DeleteQualityRulesResponse
func (client *Client) DeleteQualityRules(request *DeleteQualityRulesRequest) (_result *DeleteQualityRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteQualityRulesResponse{}
	_body, _err := client.DeleteQualityRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes quality scheduling objects in batches.
//
// Release version: v5.4.2.
//
// @param tmpReq - DeleteQualitySchedulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualitySchedulesResponse
func (client *Client) DeleteQualitySchedulesWithOptions(tmpReq *DeleteQualitySchedulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteQualitySchedulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteQualitySchedulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteCommand) {
		request.DeleteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteCommand, dara.String("DeleteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCommandShrink) {
		body["DeleteCommand"] = request.DeleteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQualitySchedules"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQualitySchedulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes quality scheduling objects in batches.
//
// Release version: v5.4.2.
//
// @param request - DeleteQualitySchedulesRequest
//
// @return DeleteQualitySchedulesResponse
func (client *Client) DeleteQualitySchedules(request *DeleteQualitySchedulesRequest) (_result *DeleteQualitySchedulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteQualitySchedulesResponse{}
	_body, _err := client.DeleteQualitySchedulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes quality template objects in batches.
//
// Online version: v5.4.2.
//
// @param tmpReq - DeleteQualityTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualityTemplatesResponse
func (client *Client) DeleteQualityTemplatesWithOptions(tmpReq *DeleteQualityTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DeleteQualityTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteQualityTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteCommand) {
		request.DeleteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteCommand, dara.String("DeleteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCommandShrink) {
		body["DeleteCommand"] = request.DeleteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQualityTemplates"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQualityTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes quality template objects in batches.
//
// Online version: v5.4.2.
//
// @param request - DeleteQualityTemplatesRequest
//
// @return DeleteQualityTemplatesResponse
func (client *Client) DeleteQualityTemplates(request *DeleteQualityTemplatesRequest) (_result *DeleteQualityTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteQualityTemplatesResponse{}
	_body, _err := client.DeleteQualityTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes monitored objects in batches.
//
// Release version: v5.4.2.
//
// @param tmpReq - DeleteQualityWatchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualityWatchesResponse
func (client *Client) DeleteQualityWatchesWithOptions(tmpReq *DeleteQualityWatchesRequest, runtime *dara.RuntimeOptions) (_result *DeleteQualityWatchesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteQualityWatchesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteCommand) {
		request.DeleteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteCommand, dara.String("DeleteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCommandShrink) {
		body["DeleteCommand"] = request.DeleteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQualityWatches"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQualityWatchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes monitored objects in batches.
//
// Release version: v5.4.2.
//
// @param request - DeleteQualityWatchesRequest
//
// @return DeleteQualityWatchesResponse
func (client *Client) DeleteQualityWatches(request *DeleteQualityWatchesRequest) (_result *DeleteQualityWatchesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteQualityWatchesResponse{}
	_body, _err := client.DeleteQualityWatchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes registered lineage. Available since version v5.4.0.
//
// @param tmpReq - DeleteRegisterLineageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRegisterLineageResponse
func (client *Client) DeleteRegisterLineageWithOptions(tmpReq *DeleteRegisterLineageRequest, runtime *dara.RuntimeOptions) (_result *DeleteRegisterLineageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteRegisterLineageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteRegisterLineageCommand) {
		request.DeleteRegisterLineageCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteRegisterLineageCommand, dara.String("DeleteRegisterLineageCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteRegisterLineageCommandShrink) {
		body["DeleteRegisterLineageCommand"] = request.DeleteRegisterLineageCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRegisterLineage"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRegisterLineageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes registered lineage. Available since version v5.4.0.
//
// @param request - DeleteRegisterLineageRequest
//
// @return DeleteRegisterLineageResponse
func (client *Client) DeleteRegisterLineage(request *DeleteRegisterLineageRequest) (_result *DeleteRegisterLineageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRegisterLineageResponse{}
	_body, _err := client.DeleteRegisterLineageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete a resource file.
//
// @param request - DeleteResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceResponse
func (client *Client) DeleteResourceWithOptions(request *DeleteResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a resource file.
//
// @param request - DeleteResourceRequest
//
// @return DeleteResourceResponse
func (client *Client) DeleteResource(request *DeleteResourceRequest) (_result *DeleteResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteResourceResponse{}
	_body, _err := client.DeleteResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a row-level permission.
//
// @param tmpReq - DeleteRowPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRowPermissionResponse
func (client *Client) DeleteRowPermissionWithOptions(tmpReq *DeleteRowPermissionRequest, runtime *dara.RuntimeOptions) (_result *DeleteRowPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteRowPermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteRowPermissionCommand) {
		request.DeleteRowPermissionCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteRowPermissionCommand, dara.String("DeleteRowPermissionCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteRowPermissionCommandShrink) {
		body["DeleteRowPermissionCommand"] = request.DeleteRowPermissionCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRowPermission"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRowPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a row-level permission.
//
// @param request - DeleteRowPermissionRequest
//
// @return DeleteRowPermissionResponse
func (client *Client) DeleteRowPermission(request *DeleteRowPermissionRequest) (_result *DeleteRowPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRowPermissionResponse{}
	_body, _err := client.DeleteRowPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data categorization. Available since v5.4.2.
//
// @param tmpReq - DeleteSecurityClassifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecurityClassifyResponse
func (client *Client) DeleteSecurityClassifyWithOptions(tmpReq *DeleteSecurityClassifyRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecurityClassifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteSecurityClassifyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteCommand) {
		request.DeleteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteCommand, dara.String("DeleteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCommandShrink) {
		body["DeleteCommand"] = request.DeleteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecurityClassify"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecurityClassifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data categorization. Available since v5.4.2.
//
// @param request - DeleteSecurityClassifyRequest
//
// @return DeleteSecurityClassifyResponse
func (client *Client) DeleteSecurityClassify(request *DeleteSecurityClassifyRequest) (_result *DeleteSecurityClassifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSecurityClassifyResponse{}
	_body, _err := client.DeleteSecurityClassifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data classification catalog. Release version: v5.4.2.
//
// @param tmpReq - DeleteSecurityClassifyCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecurityClassifyCatalogResponse
func (client *Client) DeleteSecurityClassifyCatalogWithOptions(tmpReq *DeleteSecurityClassifyCatalogRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecurityClassifyCatalogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteSecurityClassifyCatalogShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteCommand) {
		request.DeleteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteCommand, dara.String("DeleteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCommandShrink) {
		body["DeleteCommand"] = request.DeleteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecurityClassifyCatalog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecurityClassifyCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data classification catalog. Release version: v5.4.2.
//
// @param request - DeleteSecurityClassifyCatalogRequest
//
// @return DeleteSecurityClassifyCatalogResponse
func (client *Client) DeleteSecurityClassifyCatalog(request *DeleteSecurityClassifyCatalogRequest) (_result *DeleteSecurityClassifyCatalogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSecurityClassifyCatalogResponse{}
	_body, _err := client.DeleteSecurityClassifyCatalogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes security identification results in batches. Release version: v5.4.2.
//
// @param tmpReq - DeleteSecurityIdentifyResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecurityIdentifyResultsResponse
func (client *Client) DeleteSecurityIdentifyResultsWithOptions(tmpReq *DeleteSecurityIdentifyResultsRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecurityIdentifyResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteSecurityIdentifyResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteCommand) {
		request.DeleteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteCommand, dara.String("DeleteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCommandShrink) {
		body["DeleteCommand"] = request.DeleteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecurityIdentifyResults"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecurityIdentifyResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes security identification results in batches. Release version: v5.4.2.
//
// @param request - DeleteSecurityIdentifyResultsRequest
//
// @return DeleteSecurityIdentifyResultsResponse
func (client *Client) DeleteSecurityIdentifyResults(request *DeleteSecurityIdentifyResultsRequest) (_result *DeleteSecurityIdentifyResultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSecurityIdentifyResultsResponse{}
	_body, _err := client.DeleteSecurityIdentifyResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data classification level. Available since v5.4.2.
//
// @param tmpReq - DeleteSecurityLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecurityLevelResponse
func (client *Client) DeleteSecurityLevelWithOptions(tmpReq *DeleteSecurityLevelRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecurityLevelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteSecurityLevelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteCommand) {
		request.DeleteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteCommand, dara.String("DeleteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCommandShrink) {
		body["DeleteCommand"] = request.DeleteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecurityLevel"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecurityLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data classification level. Available since v5.4.2.
//
// @param request - DeleteSecurityLevelRequest
//
// @return DeleteSecurityLevelResponse
func (client *Client) DeleteSecurityLevel(request *DeleteSecurityLevelRequest) (_result *DeleteSecurityLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSecurityLevelResponse{}
	_body, _err := client.DeleteSecurityLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a standard.
//
// Online version: v5.4.2.
//
// @param tmpReq - DeleteStandardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStandardResponse
func (client *Client) DeleteStandardWithOptions(tmpReq *DeleteStandardRequest, runtime *dara.RuntimeOptions) (_result *DeleteStandardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteStandardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteCommand) {
		request.DeleteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteCommand, dara.String("DeleteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCommandShrink) {
		body["DeleteCommand"] = request.DeleteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStandard"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStandardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a standard.
//
// Online version: v5.4.2.
//
// @param request - DeleteStandardRequest
//
// @return DeleteStandardResponse
func (client *Client) DeleteStandard(request *DeleteStandardRequest) (_result *DeleteStandardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStandardResponse{}
	_body, _err := client.DeleteStandardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes invalid mapping relationships.
//
// Online version: v5.4.2.
//
// @param tmpReq - DeleteStandardInValidMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStandardInValidMappingResponse
func (client *Client) DeleteStandardInValidMappingWithOptions(tmpReq *DeleteStandardInValidMappingRequest, runtime *dara.RuntimeOptions) (_result *DeleteStandardInValidMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteStandardInValidMappingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteCommand) {
		request.DeleteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteCommand, dara.String("DeleteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCommandShrink) {
		body["DeleteCommand"] = request.DeleteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStandardInValidMapping"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStandardInValidMappingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes invalid mapping relationships.
//
// Online version: v5.4.2.
//
// @param request - DeleteStandardInValidMappingRequest
//
// @return DeleteStandardInValidMappingResponse
func (client *Client) DeleteStandardInValidMapping(request *DeleteStandardInValidMappingRequest) (_result *DeleteStandardInValidMappingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStandardInValidMappingResponse{}
	_body, _err := client.DeleteStandardInValidMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data standard lookup table. Release version: v5.4.2.
//
// @param request - DeleteStandardLookupTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStandardLookupTableResponse
func (client *Client) DeleteStandardLookupTableWithOptions(request *DeleteStandardLookupTableRequest, runtime *dara.RuntimeOptions) (_result *DeleteStandardLookupTableResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStandardLookupTable"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStandardLookupTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data standard lookup table. Release version: v5.4.2.
//
// @param request - DeleteStandardLookupTableRequest
//
// @return DeleteStandardLookupTableResponse
func (client *Client) DeleteStandardLookupTable(request *DeleteStandardLookupTableRequest) (_result *DeleteStandardLookupTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStandardLookupTableResponse{}
	_body, _err := client.DeleteStandardLookupTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes standard associations in batches.
//
// Release version: v5.4.2.
//
// @param tmpReq - DeleteStandardRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStandardRelationsResponse
func (client *Client) DeleteStandardRelationsWithOptions(tmpReq *DeleteStandardRelationsRequest, runtime *dara.RuntimeOptions) (_result *DeleteStandardRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteStandardRelationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteCommand) {
		request.DeleteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteCommand, dara.String("DeleteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCommandShrink) {
		body["DeleteCommand"] = request.DeleteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStandardRelations"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStandardRelationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes standard associations in batches.
//
// Release version: v5.4.2.
//
// @param request - DeleteStandardRelationsRequest
//
// @return DeleteStandardRelationsResponse
func (client *Client) DeleteStandardRelations(request *DeleteStandardRelationsRequest) (_result *DeleteStandardRelationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStandardRelationsResponse{}
	_body, _err := client.DeleteStandardRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a standard set.
//
// Online version: v5.4.2.
//
// @param request - DeleteStandardSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStandardSetResponse
func (client *Client) DeleteStandardSetWithOptions(request *DeleteStandardSetRequest, runtime *dara.RuntimeOptions) (_result *DeleteStandardSetResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStandardSet"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStandardSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a standard set.
//
// Online version: v5.4.2.
//
// @param request - DeleteStandardSetRequest
//
// @return DeleteStandardSetResponse
func (client *Client) DeleteStandardSet(request *DeleteStandardSetRequest) (_result *DeleteStandardSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStandardSetResponse{}
	_body, _err := client.DeleteStandardSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes valid mapping relationships.
//
// Release version: v5.4.2.
//
// @param tmpReq - DeleteStandardValidMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStandardValidMappingResponse
func (client *Client) DeleteStandardValidMappingWithOptions(tmpReq *DeleteStandardValidMappingRequest, runtime *dara.RuntimeOptions) (_result *DeleteStandardValidMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteStandardValidMappingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteCommand) {
		request.DeleteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteCommand, dara.String("DeleteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCommandShrink) {
		body["DeleteCommand"] = request.DeleteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStandardValidMapping"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStandardValidMappingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes valid mapping relationships.
//
// Release version: v5.4.2.
//
// @param request - DeleteStandardValidMappingRequest
//
// @return DeleteStandardValidMappingResponse
func (client *Client) DeleteStandardValidMapping(request *DeleteStandardValidMappingRequest) (_result *DeleteStandardValidMappingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStandardValidMappingResponse{}
	_body, _err := client.DeleteStandardValidMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data standard root word.
//
// Online version: v5.4.2.
//
// @param request - DeleteStandardWordRootRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStandardWordRootResponse
func (client *Client) DeleteStandardWordRootWithOptions(request *DeleteStandardWordRootRequest, runtime *dara.RuntimeOptions) (_result *DeleteStandardWordRootResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStandardWordRoot"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStandardWordRootResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data standard root word.
//
// Online version: v5.4.2.
//
// @param request - DeleteStandardWordRootRequest
//
// @return DeleteStandardWordRootResponse
func (client *Client) DeleteStandardWordRoot(request *DeleteStandardWordRootRequest) (_result *DeleteStandardWordRootResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStandardWordRootResponse{}
	_body, _err := client.DeleteStandardWordRootWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a user-defined function.
//
// @param request - DeleteUdfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUdfResponse
func (client *Client) DeleteUdfWithOptions(request *DeleteUdfRequest, runtime *dara.RuntimeOptions) (_result *DeleteUdfResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUdf"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUdfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user-defined function.
//
// @param request - DeleteUdfRequest
//
// @return DeleteUdfResponse
func (client *Client) DeleteUdf(request *DeleteUdfRequest) (_result *DeleteUdfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUdfResponse{}
	_body, _err := client.DeleteUdfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a user group.
//
// @param request - DeleteUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserGroupResponse
func (client *Client) DeleteUserGroupWithOptions(request *DeleteUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserGroup"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user group.
//
// @param request - DeleteUserGroupRequest
//
// @return DeleteUserGroupResponse
func (client *Client) DeleteUserGroup(request *DeleteUserGroupRequest) (_result *DeleteUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.DeleteUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes an ad hoc query task.
//
// @param tmpReq - ExecuteAdHocTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAdHocTaskResponse
func (client *Client) ExecuteAdHocTaskWithOptions(tmpReq *ExecuteAdHocTaskRequest, runtime *dara.RuntimeOptions) (_result *ExecuteAdHocTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExecuteAdHocTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExecuteCommand) {
		request.ExecuteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecuteCommand, dara.String("ExecuteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExecuteCommandShrink) {
		body["ExecuteCommand"] = request.ExecuteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAdHocTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteAdHocTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes an ad hoc query task.
//
// @param request - ExecuteAdHocTaskRequest
//
// @return ExecuteAdHocTaskResponse
func (client *Client) ExecuteAdHocTask(request *ExecuteAdHocTaskRequest) (_result *ExecuteAdHocTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteAdHocTaskResponse{}
	_body, _err := client.ExecuteAdHocTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Runs a manually scheduled node.
//
// @param tmpReq - ExecuteManualNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteManualNodeResponse
func (client *Client) ExecuteManualNodeWithOptions(tmpReq *ExecuteManualNodeRequest, runtime *dara.RuntimeOptions) (_result *ExecuteManualNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExecuteManualNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExecuteCommand) {
		request.ExecuteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecuteCommand, dara.String("ExecuteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExecuteCommandShrink) {
		body["ExecuteCommand"] = request.ExecuteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteManualNode"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteManualNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Runs a manually scheduled node.
//
// @param request - ExecuteManualNodeRequest
//
// @return ExecuteManualNodeResponse
func (client *Client) ExecuteManualNode(request *ExecuteManualNodeRequest) (_result *ExecuteManualNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteManualNodeResponse{}
	_body, _err := client.ExecuteManualNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Runs a trigger-based node.
//
// @param request - ExecuteTriggerNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTriggerNodeResponse
func (client *Client) ExecuteTriggerNodeWithOptions(request *ExecuteTriggerNodeRequest, runtime *dara.RuntimeOptions) (_result *ExecuteTriggerNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizDate) {
		query["BizDate"] = request.BizDate
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.Index) {
		query["Index"] = request.Index
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteTriggerNode"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteTriggerNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Runs a trigger-based node.
//
// @param request - ExecuteTriggerNodeRequest
//
// @return ExecuteTriggerNodeResponse
func (client *Client) ExecuteTriggerNode(request *ExecuteTriggerNodeRequest) (_result *ExecuteTriggerNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteTriggerNodeResponse{}
	_body, _err := client.ExecuteTriggerNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reruns downstream nodes to fix data link issues. Supports forced rerun of downstream nodes. Impact: incurs compute costs and affects data output.
//
// @param tmpReq - FixDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FixDataResponse
func (client *Client) FixDataWithOptions(tmpReq *FixDataRequest, runtime *dara.RuntimeOptions) (_result *FixDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &FixDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FixDataCommand) {
		request.FixDataCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FixDataCommand, dara.String("FixDataCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FixDataCommandShrink) {
		body["FixDataCommand"] = request.FixDataCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FixData"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FixDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reruns downstream nodes to fix data link issues. Supports forced rerun of downstream nodes. Impact: incurs compute costs and affects data output.
//
// @param request - FixDataRequest
//
// @return FixDataResponse
func (client *Client) FixData(request *FixDataRequest) (_result *FixDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FixDataResponse{}
	_body, _err := client.FixDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves all authorized accounts under a specific row-level permission by row-level permission ID.
//
// @param tmpReq - GetAccountByRowPermissionIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountByRowPermissionIdResponse
func (client *Client) GetAccountByRowPermissionIdWithOptions(tmpReq *GetAccountByRowPermissionIdRequest, runtime *dara.RuntimeOptions) (_result *GetAccountByRowPermissionIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAccountByRowPermissionIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GetAccountByRowPermissionIdQuery) {
		request.GetAccountByRowPermissionIdQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GetAccountByRowPermissionIdQuery, dara.String("GetAccountByRowPermissionIdQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GetAccountByRowPermissionIdQueryShrink) {
		body["GetAccountByRowPermissionIdQuery"] = request.GetAccountByRowPermissionIdQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccountByRowPermissionId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccountByRowPermissionIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all authorized accounts under a specific row-level permission by row-level permission ID.
//
// @param request - GetAccountByRowPermissionIdRequest
//
// @return GetAccountByRowPermissionIdResponse
func (client *Client) GetAccountByRowPermissionId(request *GetAccountByRowPermissionIdRequest) (_result *GetAccountByRowPermissionIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccountByRowPermissionIdResponse{}
	_body, _err := client.GetAccountByRowPermissionIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a custom query file in the directory tree.
//
// @param request - GetAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdHocFileResponse
func (client *Client) GetAdHocFileWithOptions(request *GetAdHocFileRequest, runtime *dara.RuntimeOptions) (_result *GetAdHocFileResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAdHocFile"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAdHocFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a custom query file in the directory tree.
//
// @param request - GetAdHocFileRequest
//
// @return GetAdHocFileResponse
func (client *Client) GetAdHocFile(request *GetAdHocFileRequest) (_result *GetAdHocFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAdHocFileResponse{}
	_body, _err := client.GetAdHocFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the runtime logs of an ad hoc query task.
//
// @param request - GetAdHocTaskLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdHocTaskLogResponse
func (client *Client) GetAdHocTaskLogWithOptions(request *GetAdHocTaskLogRequest, runtime *dara.RuntimeOptions) (_result *GetAdHocTaskLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SubTaskId) {
		query["SubTaskId"] = request.SubTaskId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAdHocTaskLog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAdHocTaskLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the runtime logs of an ad hoc query task.
//
// @param request - GetAdHocTaskLogRequest
//
// @return GetAdHocTaskLogResponse
func (client *Client) GetAdHocTaskLog(request *GetAdHocTaskLogRequest) (_result *GetAdHocTaskLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAdHocTaskLogResponse{}
	_body, _err := client.GetAdHocTaskLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the task execution result of an ad hoc query.
//
// @param request - GetAdHocTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdHocTaskResultResponse
func (client *Client) GetAdHocTaskResultWithOptions(request *GetAdHocTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetAdHocTaskResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SubTaskId) {
		query["SubTaskId"] = request.SubTaskId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAdHocTaskResult"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAdHocTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the task execution result of an ad hoc query.
//
// @param request - GetAdHocTaskResultRequest
//
// @return GetAdHocTaskResultResponse
func (client *Client) GetAdHocTaskResult(request *GetAdHocTaskResultRequest) (_result *GetAdHocTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAdHocTaskResultResponse{}
	_body, _err := client.GetAdHocTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of an alert event.
//
// @param request - GetAlertEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlertEventResponse
func (client *Client) GetAlertEventWithOptions(request *GetAlertEventRequest, runtime *dara.RuntimeOptions) (_result *GetAlertEventResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlertEvent"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlertEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an alert event.
//
// @param request - GetAlertEventRequest
//
// @return GetAlertEventResponse
func (client *Client) GetAlertEvent(request *GetAlertEventRequest) (_result *GetAlertEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAlertEventResponse{}
	_body, _err := client.GetAlertEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries mapping relationships by asset object GUID.
//
// Available since: v5.4.2.
//
// @param tmpReq - GetAssetMappingRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAssetMappingRelationsResponse
func (client *Client) GetAssetMappingRelationsWithOptions(tmpReq *GetAssetMappingRelationsRequest, runtime *dara.RuntimeOptions) (_result *GetAssetMappingRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAssetMappingRelationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AssetMappingQuery) {
		request.AssetMappingQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssetMappingQuery, dara.String("AssetMappingQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AssetMappingQueryShrink) {
		body["AssetMappingQuery"] = request.AssetMappingQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAssetMappingRelations"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAssetMappingRelationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries mapping relationships by asset object GUID.
//
// Available since: v5.4.2.
//
// @param request - GetAssetMappingRelationsRequest
//
// @return GetAssetMappingRelationsResponse
func (client *Client) GetAssetMappingRelations(request *GetAssetMappingRelationsRequest) (_result *GetAssetMappingRelationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAssetMappingRelationsResponse{}
	_body, _err := client.GetAssetMappingRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of an offline compute node.
//
// @param request - GetBatchTaskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTaskInfoResponse
func (client *Client) GetBatchTaskInfoWithOptions(request *GetBatchTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTaskInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.IncludeAllUpStreams) {
		query["IncludeAllUpStreams"] = request.IncludeAllUpStreams
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBatchTaskInfo"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBatchTaskInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an offline compute node.
//
// @param request - GetBatchTaskInfoRequest
//
// @return GetBatchTaskInfoResponse
func (client *Client) GetBatchTaskInfo(request *GetBatchTaskInfoRequest) (_result *GetBatchTaskInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBatchTaskInfoResponse{}
	_body, _err := client.GetBatchTaskInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified version of a batch task.
//
// @param request - GetBatchTaskInfoByVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTaskInfoByVersionResponse
func (client *Client) GetBatchTaskInfoByVersionWithOptions(request *GetBatchTaskInfoByVersionRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTaskInfoByVersionResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBatchTaskInfoByVersion"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBatchTaskInfoByVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified version of a batch task.
//
// @param request - GetBatchTaskInfoByVersionRequest
//
// @return GetBatchTaskInfoByVersionResponse
func (client *Client) GetBatchTaskInfoByVersion(request *GetBatchTaskInfoByVersionRequest) (_result *GetBatchTaskInfoByVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBatchTaskInfoByVersionResponse{}
	_body, _err := client.GetBatchTaskInfoByVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the custom lineage of an offline task.
//
// @param request - GetBatchTaskUdfLineagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTaskUdfLineagesResponse
func (client *Client) GetBatchTaskUdfLineagesWithOptions(request *GetBatchTaskUdfLineagesRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTaskUdfLineagesResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBatchTaskUdfLineages"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBatchTaskUdfLineagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the custom lineage of an offline task.
//
// @param request - GetBatchTaskUdfLineagesRequest
//
// @return GetBatchTaskUdfLineagesResponse
func (client *Client) GetBatchTaskUdfLineages(request *GetBatchTaskUdfLineagesRequest) (_result *GetBatchTaskUdfLineagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBatchTaskUdfLineagesResponse{}
	_body, _err := client.GetBatchTaskUdfLineagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the version list of a batch task.
//
// @param request - GetBatchTaskVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTaskVersionsResponse
func (client *Client) GetBatchTaskVersionsWithOptions(request *GetBatchTaskVersionsRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTaskVersionsResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBatchTaskVersions"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBatchTaskVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the version list of a batch task.
//
// @param request - GetBatchTaskVersionsRequest
//
// @return GetBatchTaskVersionsResponse
func (client *Client) GetBatchTaskVersions(request *GetBatchTaskVersionsRequest) (_result *GetBatchTaskVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBatchTaskVersionsResponse{}
	_body, _err := client.GetBatchTaskVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query mapping relationships by belonging asset GUID.
//
// Release version: v5.4.2.
//
// @param tmpReq - GetBelongAssetMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBelongAssetMappingResponse
func (client *Client) GetBelongAssetMappingWithOptions(tmpReq *GetBelongAssetMappingRequest, runtime *dara.RuntimeOptions) (_result *GetBelongAssetMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetBelongAssetMappingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AssetMappingQuery) {
		request.AssetMappingQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssetMappingQuery, dara.String("AssetMappingQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AssetMappingQueryShrink) {
		body["AssetMappingQuery"] = request.AssetMappingQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBelongAssetMapping"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBelongAssetMappingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query mapping relationships by belonging asset GUID.
//
// Release version: v5.4.2.
//
// @param request - GetBelongAssetMappingRequest
//
// @return GetBelongAssetMappingResponse
func (client *Client) GetBelongAssetMapping(request *GetBelongAssetMappingRequest) (_result *GetBelongAssetMappingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBelongAssetMappingResponse{}
	_body, _err := client.GetBelongAssetMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a business entity.
//
// @param request - GetBizEntityInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBizEntityInfoResponse
func (client *Client) GetBizEntityInfoWithOptions(request *GetBizEntityInfoRequest, runtime *dara.RuntimeOptions) (_result *GetBizEntityInfoResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBizEntityInfo"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBizEntityInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a business entity.
//
// @param request - GetBizEntityInfoRequest
//
// @return GetBizEntityInfoResponse
func (client *Client) GetBizEntityInfo(request *GetBizEntityInfoRequest) (_result *GetBizEntityInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBizEntityInfoResponse{}
	_body, _err := client.GetBizEntityInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a business entity of a specified version.
//
// @param request - GetBizEntityInfoByVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBizEntityInfoByVersionResponse
func (client *Client) GetBizEntityInfoByVersionWithOptions(request *GetBizEntityInfoByVersionRequest, runtime *dara.RuntimeOptions) (_result *GetBizEntityInfoByVersionResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBizEntityInfoByVersion"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBizEntityInfoByVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a business entity of a specified version.
//
// @param request - GetBizEntityInfoByVersionRequest
//
// @return GetBizEntityInfoByVersionResponse
func (client *Client) GetBizEntityInfoByVersion(request *GetBizEntityInfoByVersionRequest) (_result *GetBizEntityInfoByVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBizEntityInfoByVersionResponse{}
	_body, _err := client.GetBizEntityInfoByVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query business metric details by name.
//
// Release version: v5.5.0.
//
// @param tmpReq - GetBizMetricByNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBizMetricByNameResponse
func (client *Client) GetBizMetricByNameWithOptions(tmpReq *GetBizMetricByNameRequest, runtime *dara.RuntimeOptions) (_result *GetBizMetricByNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetBizMetricByNameShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizMetricByNameQuery) {
		request.BizMetricByNameQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizMetricByNameQuery, dara.String("BizMetricByNameQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizMetricByNameQueryShrink) {
		body["BizMetricByNameQuery"] = request.BizMetricByNameQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBizMetricByName"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBizMetricByNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query business metric details by name.
//
// Release version: v5.5.0.
//
// @param request - GetBizMetricByNameRequest
//
// @return GetBizMetricByNameResponse
func (client *Client) GetBizMetricByName(request *GetBizMetricByNameRequest) (_result *GetBizMetricByNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBizMetricByNameResponse{}
	_body, _err := client.GetBizMetricByNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data domain.
//
// @param request - GetBizUnitInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBizUnitInfoResponse
func (client *Client) GetBizUnitInfoWithOptions(request *GetBizUnitInfoRequest, runtime *dara.RuntimeOptions) (_result *GetBizUnitInfoResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBizUnitInfo"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBizUnitInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data domain.
//
// @param request - GetBizUnitInfoRequest
//
// @return GetBizUnitInfoResponse
func (client *Client) GetBizUnitInfo(request *GetBizUnitInfoRequest) (_result *GetBizUnitInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBizUnitInfoResponse{}
	_body, _err := client.GetBizUnitInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries asset details. Release version: v6.1.0.
//
// @param tmpReq - GetCatalogAssetDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCatalogAssetDetailsResponse
func (client *Client) GetCatalogAssetDetailsWithOptions(tmpReq *GetCatalogAssetDetailsRequest, runtime *dara.RuntimeOptions) (_result *GetCatalogAssetDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetCatalogAssetDetailsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GetCatalogAssetDetailsQuery) {
		request.GetCatalogAssetDetailsQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GetCatalogAssetDetailsQuery, dara.String("GetCatalogAssetDetailsQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GetCatalogAssetDetailsQueryShrink) {
		body["GetCatalogAssetDetailsQuery"] = request.GetCatalogAssetDetailsQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCatalogAssetDetails"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCatalogAssetDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries asset details. Release version: v6.1.0.
//
// @param request - GetCatalogAssetDetailsRequest
//
// @return GetCatalogAssetDetailsResponse
func (client *Client) GetCatalogAssetDetails(request *GetCatalogAssetDetailsRequest) (_result *GetCatalogAssetDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCatalogAssetDetailsResponse{}
	_body, _err := client.GetCatalogAssetDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of connectivity check tasks for a specified data source ID. This operation includes null value validation and tenant permission verification to prevent cross-tenant access.
//
// Release version: v5.5.0.
//
// Description:
//
// Queries the details of connectivity tasks that have been tested for a specified data source ID.
//
// @param request - GetCheckConnectivityJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCheckConnectivityJobsResponse
func (client *Client) GetCheckConnectivityJobsWithOptions(request *GetCheckConnectivityJobsRequest, runtime *dara.RuntimeOptions) (_result *GetCheckConnectivityJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCheckConnectivityJobs"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCheckConnectivityJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of connectivity check tasks for a specified data source ID. This operation includes null value validation and tenant permission verification to prevent cross-tenant access.
//
// Release version: v5.5.0.
//
// Description:
//
// Queries the details of connectivity tasks that have been tested for a specified data source ID.
//
// @param request - GetCheckConnectivityJobsRequest
//
// @return GetCheckConnectivityJobsResponse
func (client *Client) GetCheckConnectivityJobs(request *GetCheckConnectivityJobsRequest) (_result *GetCheckConnectivityJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCheckConnectivityJobsResponse{}
	_body, _err := client.GetCheckConnectivityJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves cluster information based on the environment.
//
// @param request - GetClusterQueueInfoByEnvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterQueueInfoByEnvResponse
func (client *Client) GetClusterQueueInfoByEnvWithOptions(request *GetClusterQueueInfoByEnvRequest, runtime *dara.RuntimeOptions) (_result *GetClusterQueueInfoByEnvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.StreamBatchMode) {
		query["StreamBatchMode"] = request.StreamBatchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClusterQueueInfoByEnv"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterQueueInfoByEnvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves cluster information based on the environment.
//
// @param request - GetClusterQueueInfoByEnvRequest
//
// @return GetClusterQueueInfoByEnvResponse
func (client *Client) GetClusterQueueInfoByEnv(request *GetClusterQueueInfoByEnvRequest) (_result *GetClusterQueueInfoByEnvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetClusterQueueInfoByEnvResponse{}
	_body, _err := client.GetClusterQueueInfoByEnvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a compute source by compute source ID.
//
// @param request - GetComputeSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComputeSourceResponse
func (client *Client) GetComputeSourceWithOptions(request *GetComputeSourceRequest, runtime *dara.RuntimeOptions) (_result *GetComputeSourceResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetComputeSource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetComputeSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a compute source by compute source ID.
//
// @param request - GetComputeSourceRequest
//
// @return GetComputeSourceResponse
func (client *Client) GetComputeSource(request *GetComputeSourceRequest) (_result *GetComputeSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetComputeSourceResponse{}
	_body, _err := client.GetComputeSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data domain.
//
// @param request - GetDataDomainInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataDomainInfoResponse
func (client *Client) GetDataDomainInfoWithOptions(request *GetDataDomainInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDataDomainInfoResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataDomainInfo"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataDomainInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data domain.
//
// @param request - GetDataDomainInfoRequest
//
// @return GetDataDomainInfoResponse
func (client *Client) GetDataDomainInfo(request *GetDataDomainInfoRequest) (_result *GetDataDomainInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataDomainInfoResponse{}
	_body, _err := client.GetDataDomainInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Service Monitoring: Retrieves the aggregate statistics of API calls.
//
// @param request - GetDataServiceApiCallSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiCallSummaryResponse
func (client *Client) GetDataServiceApiCallSummaryWithOptions(request *GetDataServiceApiCallSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiCallSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceApiCallSummary"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceApiCallSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Service Monitoring: Retrieves the aggregate statistics of API calls.
//
// @param request - GetDataServiceApiCallSummaryRequest
//
// @return GetDataServiceApiCallSummaryResponse
func (client *Client) GetDataServiceApiCallSummary(request *GetDataServiceApiCallSummaryRequest) (_result *GetDataServiceApiCallSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceApiCallSummaryResponse{}
	_body, _err := client.GetDataServiceApiCallSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Service Monitoring: Analyzes API access trends.
//
// @param request - GetDataServiceApiCallTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiCallTrendResponse
func (client *Client) GetDataServiceApiCallTrendWithOptions(request *GetDataServiceApiCallTrendRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiCallTrendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceApiCallTrend"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceApiCallTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Service Monitoring: Analyzes API access trends.
//
// @param request - GetDataServiceApiCallTrendRequest
//
// @return GetDataServiceApiCallTrendResponse
func (client *Client) GetDataServiceApiCallTrend(request *GetDataServiceApiCallTrendRequest) (_result *GetDataServiceApiCallTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceApiCallTrendResponse{}
	_body, _err := client.GetDataServiceApiCallTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves API documentation.
//
// @param request - GetDataServiceApiDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiDocumentResponse
func (client *Client) GetDataServiceApiDocumentWithOptions(request *GetDataServiceApiDocumentRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiDocumentResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceApiDocument"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceApiDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves API documentation.
//
// @param request - GetDataServiceApiDocumentRequest
//
// @return GetDataServiceApiDocumentResponse
func (client *Client) GetDataServiceApiDocument(request *GetDataServiceApiDocumentRequest) (_result *GetDataServiceApiDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceApiDocumentResponse{}
	_body, _err := client.GetDataServiceApiDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the summary of API exception impacts.
//
// @param request - GetDataServiceApiErrorImpactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiErrorImpactResponse
func (client *Client) GetDataServiceApiErrorImpactWithOptions(request *GetDataServiceApiErrorImpactRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiErrorImpactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceApiErrorImpact"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceApiErrorImpactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the summary of API exception impacts.
//
// @param request - GetDataServiceApiErrorImpactRequest
//
// @return GetDataServiceApiErrorImpactResponse
func (client *Client) GetDataServiceApiErrorImpact(request *GetDataServiceApiErrorImpactRequest) (_result *GetDataServiceApiErrorImpactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceApiErrorImpactResponse{}
	_body, _err := client.GetDataServiceApiErrorImpactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of API groups in Data Service.
//
// @param request - GetDataServiceApiGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiGroupsResponse
func (client *Client) GetDataServiceApiGroupsWithOptions(request *GetDataServiceApiGroupsRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceApiGroups"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceApiGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of API groups in Data Service.
//
// @param request - GetDataServiceApiGroupsRequest
//
// @return GetDataServiceApiGroupsResponse
func (client *Client) GetDataServiceApiGroups(request *GetDataServiceApiGroupsRequest) (_result *GetDataServiceApiGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceApiGroupsResponse{}
	_body, _err := client.GetDataServiceApiGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a data service application, including the project, application name, authentication information, and IP whitelist. Only application members can view the details.
//
// Release version: v6.0.0.
//
// @param request - GetDataServiceAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAppResponse
func (client *Client) GetDataServiceAppWithOptions(request *GetDataServiceAppRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceApp"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a data service application, including the project, application name, authentication information, and IP whitelist. Only application members can view the details.
//
// Release version: v6.0.0.
//
// @param request - GetDataServiceAppRequest
//
// @return GetDataServiceAppResponse
func (client *Client) GetDataServiceApp(request *GetDataServiceAppRequest) (_result *GetDataServiceAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceAppResponse{}
	_body, _err := client.GetDataServiceAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of users who have permissions on an application.
//
// @param request - GetDataServiceAppAuthorizedUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAppAuthorizedUsersResponse
func (client *Client) GetDataServiceAppAuthorizedUsersWithOptions(request *GetDataServiceAppAuthorizedUsersRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAppAuthorizedUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceAppAuthorizedUsers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceAppAuthorizedUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of users who have permissions on an application.
//
// @param request - GetDataServiceAppAuthorizedUsersRequest
//
// @return GetDataServiceAppAuthorizedUsersResponse
func (client *Client) GetDataServiceAppAuthorizedUsers(request *GetDataServiceAppAuthorizedUsersRequest) (_result *GetDataServiceAppAuthorizedUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceAppAuthorizedUsersResponse{}
	_body, _err := client.GetDataServiceAppAuthorizedUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of application groups for a data service project.
//
// @param request - GetDataServiceAppGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAppGroupsResponse
func (client *Client) GetDataServiceAppGroupsWithOptions(request *GetDataServiceAppGroupsRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAppGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceAppGroups"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceAppGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of application groups for a data service project.
//
// @param request - GetDataServiceAppGroupsRequest
//
// @return GetDataServiceAppGroupsResponse
func (client *Client) GetDataServiceAppGroups(request *GetDataServiceAppGroupsRequest) (_result *GetDataServiceAppGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceAppGroupsResponse{}
	_body, _err := client.GetDataServiceAppGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the member list of a data service application, including regular members and owners. Only application owners can call this operation.
//
// Online version: v6.0.0.
//
// @param request - GetDataServiceAppMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAppMembersResponse
func (client *Client) GetDataServiceAppMembersWithOptions(request *GetDataServiceAppMembersRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAppMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceAppMembers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceAppMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the member list of a data service application, including regular members and owners. Only application owners can call this operation.
//
// Online version: v6.0.0.
//
// @param request - GetDataServiceAppMembersRequest
//
// @return GetDataServiceAppMembersResponse
func (client *Client) GetDataServiceAppMembers(request *GetDataServiceAppMembersRequest) (_result *GetDataServiceAppMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceAppMembersResponse{}
	_body, _err := client.GetDataServiceAppMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of applications in a group.
//
// @param request - GetDataServiceAppsByGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAppsByGroupIdResponse
func (client *Client) GetDataServiceAppsByGroupIdWithOptions(request *GetDataServiceAppsByGroupIdRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAppsByGroupIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceAppsByGroupId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceAppsByGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of applications in a group.
//
// @param request - GetDataServiceAppsByGroupIdRequest
//
// @return GetDataServiceAppsByGroupIdResponse
func (client *Client) GetDataServiceAppsByGroupId(request *GetDataServiceAppsByGroupIdRequest) (_result *GetDataServiceAppsByGroupIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceAppsByGroupIdResponse{}
	_body, _err := client.GetDataServiceAppsByGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of applications that the account has permissions to access based on the app group ID.
//
// @param request - GetDataServiceAuthorizedAppsByGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAuthorizedAppsByGroupIdResponse
func (client *Client) GetDataServiceAuthorizedAppsByGroupIdWithOptions(request *GetDataServiceAuthorizedAppsByGroupIdRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAuthorizedAppsByGroupIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceAuthorizedAppsByGroupId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceAuthorizedAppsByGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of applications that the account has permissions to access based on the app group ID.
//
// @param request - GetDataServiceAuthorizedAppsByGroupIdRequest
//
// @return GetDataServiceAuthorizedAppsByGroupIdResponse
func (client *Client) GetDataServiceAuthorizedAppsByGroupId(request *GetDataServiceAuthorizedAppsByGroupIdRequest) (_result *GetDataServiceAuthorizedAppsByGroupIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceAuthorizedAppsByGroupIdResponse{}
	_body, _err := client.GetDataServiceAuthorizedAppsByGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of projects that the current user has permissions to access.
//
// @param request - GetDataServiceAuthorizedProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAuthorizedProjectsResponse
func (client *Client) GetDataServiceAuthorizedProjectsWithOptions(request *GetDataServiceAuthorizedProjectsRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAuthorizedProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceAuthorizedProjects"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceAuthorizedProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of projects that the current user has permissions to access.
//
// @param request - GetDataServiceAuthorizedProjectsRequest
//
// @return GetDataServiceAuthorizedProjectsResponse
func (client *Client) GetDataServiceAuthorizedProjects(request *GetDataServiceAuthorizedProjectsRequest) (_result *GetDataServiceAuthorizedProjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceAuthorizedProjectsResponse{}
	_body, _err := client.GetDataServiceAuthorizedProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of projects for which the current user is the owner.
//
// @param request - GetDataServiceMyProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceMyProjectsResponse
func (client *Client) GetDataServiceMyProjectsWithOptions(request *GetDataServiceMyProjectsRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceMyProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceMyProjects"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceMyProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of projects for which the current user is the owner.
//
// @param request - GetDataServiceMyProjectsRequest
//
// @return GetDataServiceMyProjectsResponse
func (client *Client) GetDataServiceMyProjects(request *GetDataServiceMyProjectsRequest) (_result *GetDataServiceMyProjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceMyProjectsResponse{}
	_body, _err := client.GetDataServiceMyProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of users who can be added as project members.
//
// @param request - GetDataServiceProjectAddableUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceProjectAddableUsersResponse
func (client *Client) GetDataServiceProjectAddableUsersWithOptions(request *GetDataServiceProjectAddableUsersRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceProjectAddableUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceProjectAddableUsers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceProjectAddableUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of users who can be added as project members.
//
// @param request - GetDataServiceProjectAddableUsersRequest
//
// @return GetDataServiceProjectAddableUsersResponse
func (client *Client) GetDataServiceProjectAddableUsers(request *GetDataServiceProjectAddableUsersRequest) (_result *GetDataServiceProjectAddableUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceProjectAddableUsersResponse{}
	_body, _err := client.GetDataServiceProjectAddableUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the integration tasks and database SQL tasks affected by data source changes.
//
// @param request - GetDataSourceDependenciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataSourceDependenciesResponse
func (client *Client) GetDataSourceDependenciesWithOptions(request *GetDataSourceDependenciesRequest, runtime *dara.RuntimeOptions) (_result *GetDataSourceDependenciesResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataSourceDependencies"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataSourceDependenciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the integration tasks and database SQL tasks affected by data source changes.
//
// @param request - GetDataSourceDependenciesRequest
//
// @return GetDataSourceDependenciesResponse
func (client *Client) GetDataSourceDependencies(request *GetDataSourceDependenciesRequest) (_result *GetDataSourceDependenciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataSourceDependenciesResponse{}
	_body, _err := client.GetDataSourceDependenciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query upstream dependencies of development objects.
//
// @param request - GetDevObjectDependencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDevObjectDependencyResponse
func (client *Client) GetDevObjectDependencyWithOptions(request *GetDevObjectDependencyRequest, runtime *dara.RuntimeOptions) (_result *GetDevObjectDependencyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ObjectFrom) {
		query["ObjectFrom"] = request.ObjectFrom
	}

	if !dara.IsNil(request.ObjectId) {
		query["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDevObjectDependency"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDevObjectDependencyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query upstream dependencies of development objects.
//
// @param request - GetDevObjectDependencyRequest
//
// @return GetDevObjectDependencyResponse
func (client *Client) GetDevObjectDependency(request *GetDevObjectDependencyRequest) (_result *GetDevObjectDependencyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDevObjectDependencyResponse{}
	_body, _err := client.GetDevObjectDependencyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the folder directory tree.
//
// @param request - GetDirectoryTreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDirectoryTreeResponse
func (client *Client) GetDirectoryTreeWithOptions(request *GetDirectoryTreeRequest, runtime *dara.RuntimeOptions) (_result *GetDirectoryTreeResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDirectoryTree"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDirectoryTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the folder directory tree.
//
// @param request - GetDirectoryTreeRequest
//
// @return GetDirectoryTreeResponse
func (client *Client) GetDirectoryTree(request *GetDirectoryTreeRequest) (_result *GetDirectoryTreeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDirectoryTreeResponse{}
	_body, _err := client.GetDirectoryTreeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains temporary read/write authorization for file storage.
//
// @param request - GetFileStorageCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileStorageCredentialResponse
func (client *Client) GetFileStorageCredentialWithOptions(request *GetFileStorageCredentialRequest, runtime *dara.RuntimeOptions) (_result *GetFileStorageCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Purpose) {
		query["Purpose"] = request.Purpose
	}

	if !dara.IsNil(request.UseVpcEndpoint) {
		query["UseVpcEndpoint"] = request.UseVpcEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileStorageCredential"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileStorageCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains temporary read/write authorization for file storage.
//
// @param request - GetFileStorageCredentialRequest
//
// @return GetFileStorageCredentialResponse
func (client *Client) GetFileStorageCredential(request *GetFileStorageCredentialRequest) (_result *GetFileStorageCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFileStorageCredentialResponse{}
	_body, _err := client.GetFileStorageCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the downstream instances of a specified instance.
//
// @param tmpReq - GetInstanceDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceDownStreamResponse
func (client *Client) GetInstanceDownStreamWithOptions(tmpReq *GetInstanceDownStreamRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceDownStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetInstanceDownStreamShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceGet) {
		request.InstanceGetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceGet, dara.String("InstanceGet"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DownStreamDepth) {
		query["DownStreamDepth"] = request.DownStreamDepth
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.RunStatus) {
		query["RunStatus"] = request.RunStatus
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceGetShrink) {
		body["InstanceGet"] = request.InstanceGetShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceDownStream"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceDownStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the downstream instances of a specified instance.
//
// @param request - GetInstanceDownStreamRequest
//
// @return GetInstanceDownStreamResponse
func (client *Client) GetInstanceDownStream(request *GetInstanceDownStreamRequest) (_result *GetInstanceDownStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceDownStreamResponse{}
	_body, _err := client.GetInstanceDownStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the dag of an instance. Logical tables and code nodes are supported.
//
// @param tmpReq - GetInstanceUpDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceUpDownStreamResponse
func (client *Client) GetInstanceUpDownStreamWithOptions(tmpReq *GetInstanceUpDownStreamRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceUpDownStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetInstanceUpDownStreamShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceId) {
		request.InstanceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceId, dara.String("InstanceId"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DownStreamDepth) {
		query["DownStreamDepth"] = request.DownStreamDepth
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.UpStreamDepth) {
		query["UpStreamDepth"] = request.UpStreamDepth
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdShrink) {
		body["InstanceId"] = request.InstanceIdShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceUpDownStream"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceUpDownStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the dag of an instance. Logical tables and code nodes are supported.
//
// @param request - GetInstanceUpDownStreamRequest
//
// @return GetInstanceUpDownStreamResponse
func (client *Client) GetInstanceUpDownStream(request *GetInstanceUpDownStreamRequest) (_result *GetInstanceUpDownStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceUpDownStreamResponse{}
	_body, _err := client.GetInstanceUpDownStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of the latest pending submit record.
//
// @param tmpReq - GetLatestSubmitDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLatestSubmitDetailResponse
func (client *Client) GetLatestSubmitDetailWithOptions(tmpReq *GetLatestSubmitDetailRequest, runtime *dara.RuntimeOptions) (_result *GetLatestSubmitDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetLatestSubmitDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SubmitDetailQuery) {
		request.SubmitDetailQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubmitDetailQuery, dara.String("SubmitDetailQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SubmitDetailQueryShrink) {
		body["SubmitDetailQuery"] = request.SubmitDetailQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLatestSubmitDetail"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLatestSubmitDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of the latest pending submit record.
//
// @param request - GetLatestSubmitDetailRequest
//
// @return GetLatestSubmitDetailResponse
func (client *Client) GetLatestSubmitDetail(request *GetLatestSubmitDetailRequest) (_result *GetLatestSubmitDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLatestSubmitDetailResponse{}
	_body, _err := client.GetLatestSubmitDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of roles for the current user.
//
// @param request - GetMyRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMyRolesResponse
func (client *Client) GetMyRolesWithOptions(request *GetMyRolesRequest, runtime *dara.RuntimeOptions) (_result *GetMyRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMyRoles"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMyRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of roles for the current user.
//
// @param request - GetMyRolesRequest
//
// @return GetMyRolesResponse
func (client *Client) GetMyRoles(request *GetMyRolesRequest) (_result *GetMyRolesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMyRolesResponse{}
	_body, _err := client.GetMyRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the tenants to which the current user belongs.
//
// @param tmpReq - GetMyTenantsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMyTenantsResponse
func (client *Client) GetMyTenantsWithOptions(tmpReq *GetMyTenantsRequest, runtime *dara.RuntimeOptions) (_result *GetMyTenantsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetMyTenantsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FeatureCodeList) {
		request.FeatureCodeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FeatureCodeList, dara.String("FeatureCodeList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FeatureCodeListShrink) {
		body["FeatureCodeList"] = request.FeatureCodeListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMyTenants"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMyTenantsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the tenants to which the current user belongs.
//
// @param request - GetMyTenantsRequest
//
// @return GetMyTenantsResponse
func (client *Client) GetMyTenants(request *GetMyTenantsRequest) (_result *GetMyTenantsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMyTenantsResponse{}
	_body, _err := client.GetMyTenantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the dag of a node. This is a general-purpose operation.
//
// @param tmpReq - GetNodeUpDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeUpDownStreamResponse
func (client *Client) GetNodeUpDownStreamWithOptions(tmpReq *GetNodeUpDownStreamRequest, runtime *dara.RuntimeOptions) (_result *GetNodeUpDownStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetNodeUpDownStreamShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodeId) {
		request.NodeIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeId, dara.String("NodeId"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DownStreamDepth) {
		query["DownStreamDepth"] = request.DownStreamDepth
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.UpStreamDepth) {
		query["UpStreamDepth"] = request.UpStreamDepth
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeIdShrink) {
		body["NodeId"] = request.NodeIdShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodeUpDownStream"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeUpDownStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the dag of a node. This is a general-purpose operation.
//
// @param request - GetNodeUpDownStreamRequest
//
// @return GetNodeUpDownStreamResponse
func (client *Client) GetNodeUpDownStream(request *GetNodeUpDownStreamRequest) (_result *GetNodeUpDownStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNodeUpDownStreamResponse{}
	_body, _err := client.GetNodeUpDownStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the submit status of a data backfill request.
//
// @param request - GetOperationSubmitStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOperationSubmitStatusResponse
func (client *Client) GetOperationSubmitStatusWithOptions(request *GetOperationSubmitStatusRequest, runtime *dara.RuntimeOptions) (_result *GetOperationSubmitStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOperationSubmitStatus"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOperationSubmitStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the submit status of a data backfill request.
//
// @param request - GetOperationSubmitStatusRequest
//
// @return GetOperationSubmitStatusResponse
func (client *Client) GetOperationSubmitStatus(request *GetOperationSubmitStatusRequest) (_result *GetOperationSubmitStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOperationSubmitStatusResponse{}
	_body, _err := client.GetOperationSubmitStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets instance information.
//
// @param request - GetPhysicalInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalInstanceResponse
func (client *Client) GetPhysicalInstanceWithOptions(request *GetPhysicalInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPhysicalInstance"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPhysicalInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets instance information.
//
// @param request - GetPhysicalInstanceRequest
//
// @return GetPhysicalInstanceResponse
func (client *Client) GetPhysicalInstance(request *GetPhysicalInstanceRequest) (_result *GetPhysicalInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPhysicalInstanceResponse{}
	_body, _err := client.GetPhysicalInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the execution logs of an instance. If the instance has been rerun multiple times, multiple log entries are returned.
//
// @param request - GetPhysicalInstanceLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalInstanceLogResponse
func (client *Client) GetPhysicalInstanceLogWithOptions(request *GetPhysicalInstanceLogRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalInstanceLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPhysicalInstanceLog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPhysicalInstanceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the execution logs of an instance. If the instance has been rerun multiple times, multiple log entries are returned.
//
// @param request - GetPhysicalInstanceLogRequest
//
// @return GetPhysicalInstanceLogResponse
func (client *Client) GetPhysicalInstanceLog(request *GetPhysicalInstanceLogRequest) (_result *GetPhysicalInstanceLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPhysicalInstanceLogResponse{}
	_body, _err := client.GetPhysicalInstanceLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a physical schedule resource.
//
// @param request - GetPhysicalNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeResponse
func (client *Client) GetPhysicalNodeWithOptions(request *GetPhysicalNodeRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPhysicalNode"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPhysicalNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a physical schedule resource.
//
// @param request - GetPhysicalNodeRequest
//
// @return GetPhysicalNodeResponse
func (client *Client) GetPhysicalNode(request *GetPhysicalNodeRequest) (_result *GetPhysicalNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPhysicalNodeResponse{}
	_body, _err := client.GetPhysicalNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a physical node by output name. Only offline code nodes and integration task nodes are supported.
//
// @param request - GetPhysicalNodeByOutputNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeByOutputNameResponse
func (client *Client) GetPhysicalNodeByOutputNameWithOptions(request *GetPhysicalNodeByOutputNameRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalNodeByOutputNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.OutputName) {
		query["OutputName"] = request.OutputName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPhysicalNodeByOutputName"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPhysicalNodeByOutputNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a physical node by output name. Only offline code nodes and integration task nodes are supported.
//
// @param request - GetPhysicalNodeByOutputNameRequest
//
// @return GetPhysicalNodeByOutputNameResponse
func (client *Client) GetPhysicalNodeByOutputName(request *GetPhysicalNodeByOutputNameRequest) (_result *GetPhysicalNodeByOutputNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPhysicalNodeByOutputNameResponse{}
	_body, _err := client.GetPhysicalNodeByOutputNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the code content of a schedule resource node.
//
// @param request - GetPhysicalNodeContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeContentResponse
func (client *Client) GetPhysicalNodeContentWithOptions(request *GetPhysicalNodeContentRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalNodeContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPhysicalNodeContent"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPhysicalNodeContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the code content of a schedule resource node.
//
// @param request - GetPhysicalNodeContentRequest
//
// @return GetPhysicalNodeContentResponse
func (client *Client) GetPhysicalNodeContent(request *GetPhysicalNodeContentRequest) (_result *GetPhysicalNodeContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPhysicalNodeContentResponse{}
	_body, _err := client.GetPhysicalNodeContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operation logs of a node.
//
// @param request - GetPhysicalNodeOperationLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeOperationLogResponse
func (client *Client) GetPhysicalNodeOperationLogWithOptions(request *GetPhysicalNodeOperationLogRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalNodeOperationLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPhysicalNodeOperationLog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPhysicalNodeOperationLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operation logs of a node.
//
// @param request - GetPhysicalNodeOperationLogRequest
//
// @return GetPhysicalNodeOperationLogResponse
func (client *Client) GetPhysicalNodeOperationLog(request *GetPhysicalNodeOperationLogRequest) (_result *GetPhysicalNodeOperationLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPhysicalNodeOperationLogResponse{}
	_body, _err := client.GetPhysicalNodeOperationLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution result of an asynchronous pipeline task.
//
// @param tmpReq - GetPipelineAsyncResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineAsyncResultResponse
func (client *Client) GetPipelineAsyncResultWithOptions(tmpReq *GetPipelineAsyncResultRequest, runtime *dara.RuntimeOptions) (_result *GetPipelineAsyncResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetPipelineAsyncResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Context) {
		request.ContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Context, dara.String("Context"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AsyncId) {
		query["AsyncId"] = request.AsyncId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextShrink) {
		body["Context"] = request.ContextShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPipelineAsyncResult"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPipelineAsyncResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution result of an asynchronous pipeline task.
//
// @param request - GetPipelineAsyncResultRequest
//
// @return GetPipelineAsyncResultResponse
func (client *Client) GetPipelineAsyncResult(request *GetPipelineAsyncResultRequest) (_result *GetPipelineAsyncResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPipelineAsyncResultResponse{}
	_body, _err := client.GetPipelineAsyncResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a pipeline task by pipeline task ID.
//
// @param tmpReq - GetPipelineByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineByIdResponse
func (client *Client) GetPipelineByIdWithOptions(tmpReq *GetPipelineByIdRequest, runtime *dara.RuntimeOptions) (_result *GetPipelineByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetPipelineByIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Context) {
		request.ContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Context, dara.String("Context"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.QueryId) {
		request.QueryIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryId, dara.String("QueryId"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextShrink) {
		body["Context"] = request.ContextShrink
	}

	if !dara.IsNil(request.QueryIdShrink) {
		body["QueryId"] = request.QueryIdShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPipelineById"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPipelineByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a pipeline task by pipeline task ID.
//
// @param request - GetPipelineByIdRequest
//
// @return GetPipelineByIdResponse
func (client *Client) GetPipelineById(request *GetPipelineByIdRequest) (_result *GetPipelineByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPipelineByIdResponse{}
	_body, _err := client.GetPipelineByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get project details by project ID.
//
// @param request - GetProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithOptions(request *GetProjectRequest, runtime *dara.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProject"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get project details by project ID.
//
// @param request - GetProjectRequest
//
// @return GetProjectResponse
func (client *Client) GetProject(request *GetProjectRequest) (_result *GetProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves project details by project name.
//
// @param request - GetProjectByNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectByNameResponse
func (client *Client) GetProjectByNameWithOptions(request *GetProjectByNameRequest, runtime *dara.RuntimeOptions) (_result *GetProjectByNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectByName"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectByNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves project details by project name.
//
// @param request - GetProjectByNameRequest
//
// @return GetProjectByNameResponse
func (client *Client) GetProjectByName(request *GetProjectByNameRequest) (_result *GetProjectByNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProjectByNameResponse{}
	_body, _err := client.GetProjectByNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the production account of a project. Only a super administrator (SuperAdmin) can call this API operation.
//
// @param request - GetProjectProduceUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectProduceUserResponse
func (client *Client) GetProjectProduceUserWithOptions(request *GetProjectProduceUserRequest, runtime *dara.RuntimeOptions) (_result *GetProjectProduceUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectProduceUser"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectProduceUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the production account of a project. Only a super administrator (SuperAdmin) can call this API operation.
//
// @param request - GetProjectProduceUserRequest
//
// @return GetProjectProduceUserResponse
func (client *Client) GetProjectProduceUser(request *GetProjectProduceUserRequest) (_result *GetProjectProduceUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProjectProduceUserResponse{}
	_body, _err := client.GetProjectProduceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the whitelist of a project.
//
// @param request - GetProjectWhiteListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectWhiteListsResponse
func (client *Client) GetProjectWhiteListsWithOptions(request *GetProjectWhiteListsRequest, runtime *dara.RuntimeOptions) (_result *GetProjectWhiteListsResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectWhiteLists"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectWhiteListsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the whitelist of a project.
//
// @param request - GetProjectWhiteListsRequest
//
// @return GetProjectWhiteListsResponse
func (client *Client) GetProjectWhiteLists(request *GetProjectWhiteListsRequest) (_result *GetProjectWhiteListsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProjectWhiteListsResponse{}
	_body, _err := client.GetProjectWhiteListsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves alert settings by monitored object ID. Release version: v5.4.2.
//
// @param request - GetQualityAlertOfAllRuleScopeByWatchIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityAlertOfAllRuleScopeByWatchIdResponse
func (client *Client) GetQualityAlertOfAllRuleScopeByWatchIdWithOptions(request *GetQualityAlertOfAllRuleScopeByWatchIdRequest, runtime *dara.RuntimeOptions) (_result *GetQualityAlertOfAllRuleScopeByWatchIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.WatchId) {
		query["WatchId"] = request.WatchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityAlertOfAllRuleScopeByWatchId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityAlertOfAllRuleScopeByWatchIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves alert settings by monitored object ID. Release version: v5.4.2.
//
// @param request - GetQualityAlertOfAllRuleScopeByWatchIdRequest
//
// @return GetQualityAlertOfAllRuleScopeByWatchIdResponse
func (client *Client) GetQualityAlertOfAllRuleScopeByWatchId(request *GetQualityAlertOfAllRuleScopeByWatchIdRequest) (_result *GetQualityAlertOfAllRuleScopeByWatchIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQualityAlertOfAllRuleScopeByWatchIdResponse{}
	_body, _err := client.GetQualityAlertOfAllRuleScopeByWatchIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a quality rule object.
//
// Release version: v5.4.2.
//
// @param request - GetQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityRuleResponse
func (client *Client) GetQualityRuleWithOptions(request *GetQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *GetQualityRuleResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityRule"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a quality rule object.
//
// Release version: v5.4.2.
//
// @param request - GetQualityRuleRequest
//
// @return GetQualityRuleResponse
func (client *Client) GetQualityRule(request *GetQualityRuleRequest) (_result *GetQualityRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQualityRuleResponse{}
	_body, _err := client.GetQualityRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a quality node task object. Online version: v5.4.2.
//
// @param request - GetQualityRuleTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityRuleTaskResponse
func (client *Client) GetQualityRuleTaskWithOptions(request *GetQualityRuleTaskRequest, runtime *dara.RuntimeOptions) (_result *GetQualityRuleTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.RuleTaskId) {
		query["RuleTaskId"] = request.RuleTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityRuleTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityRuleTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a quality node task object. Online version: v5.4.2.
//
// @param request - GetQualityRuleTaskRequest
//
// @return GetQualityRuleTaskResponse
func (client *Client) GetQualityRuleTask(request *GetQualityRuleTaskRequest) (_result *GetQualityRuleTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQualityRuleTaskResponse{}
	_body, _err := client.GetQualityRuleTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the log content of a quality node task object.
//
// Online version: v5.4.2.
//
// @param request - GetQualityRuleTaskLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityRuleTaskLogResponse
func (client *Client) GetQualityRuleTaskLogWithOptions(request *GetQualityRuleTaskLogRequest, runtime *dara.RuntimeOptions) (_result *GetQualityRuleTaskLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.RuleTaskId) {
		query["RuleTaskId"] = request.RuleTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityRuleTaskLog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityRuleTaskLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the log content of a quality node task object.
//
// Online version: v5.4.2.
//
// @param request - GetQualityRuleTaskLogRequest
//
// @return GetQualityRuleTaskLogResponse
func (client *Client) GetQualityRuleTaskLog(request *GetQualityRuleTaskLogRequest) (_result *GetQualityRuleTaskLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQualityRuleTaskLogResponse{}
	_body, _err := client.GetQualityRuleTaskLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a quality schedule object.
//
// Release version: v5.4.2.
//
// @param request - GetQualityScheduleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityScheduleResponse
func (client *Client) GetQualityScheduleWithOptions(request *GetQualityScheduleRequest, runtime *dara.RuntimeOptions) (_result *GetQualityScheduleResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualitySchedule"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a quality schedule object.
//
// Release version: v5.4.2.
//
// @param request - GetQualityScheduleRequest
//
// @return GetQualityScheduleResponse
func (client *Client) GetQualitySchedule(request *GetQualityScheduleRequest) (_result *GetQualityScheduleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQualityScheduleResponse{}
	_body, _err := client.GetQualityScheduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of schedule settings by monitored object ID.
//
// Release version: v5.4.2.
//
// @param request - GetQualitySchedulesByWatchIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualitySchedulesByWatchIdResponse
func (client *Client) GetQualitySchedulesByWatchIdWithOptions(request *GetQualitySchedulesByWatchIdRequest, runtime *dara.RuntimeOptions) (_result *GetQualitySchedulesByWatchIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.WatchId) {
		query["WatchId"] = request.WatchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualitySchedulesByWatchId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualitySchedulesByWatchIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of schedule settings by monitored object ID.
//
// Release version: v5.4.2.
//
// @param request - GetQualitySchedulesByWatchIdRequest
//
// @return GetQualitySchedulesByWatchIdResponse
func (client *Client) GetQualitySchedulesByWatchId(request *GetQualitySchedulesByWatchIdRequest) (_result *GetQualitySchedulesByWatchIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQualitySchedulesByWatchIdResponse{}
	_body, _err := client.GetQualitySchedulesByWatchIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a quality template object.
//
// Release version: v5.4.2.
//
// @param request - GetQualityTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityTemplateResponse
func (client *Client) GetQualityTemplateWithOptions(request *GetQualityTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetQualityTemplateResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityTemplate"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a quality template object.
//
// Release version: v5.4.2.
//
// @param request - GetQualityTemplateRequest
//
// @return GetQualityTemplateResponse
func (client *Client) GetQualityTemplate(request *GetQualityTemplateRequest) (_result *GetQualityTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQualityTemplateResponse{}
	_body, _err := client.GetQualityTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a quality monitored object.
//
// Release version: v5.4.2.
//
// @param request - GetQualityWatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityWatchResponse
func (client *Client) GetQualityWatchWithOptions(request *GetQualityWatchRequest, runtime *dara.RuntimeOptions) (_result *GetQualityWatchResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityWatch"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityWatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a quality monitored object.
//
// Release version: v5.4.2.
//
// @param request - GetQualityWatchRequest
//
// @return GetQualityWatchResponse
func (client *Client) GetQualityWatch(request *GetQualityWatchRequest) (_result *GetQualityWatchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQualityWatchResponse{}
	_body, _err := client.GetQualityWatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a quality watchtask record by the original ID of the monitored object, such as the ID of a datasource, table, or metric.
//
// Release version: v5.4.2.
//
// @param request - GetQualityWatchByObjectIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityWatchByObjectIdResponse
func (client *Client) GetQualityWatchByObjectIdWithOptions(request *GetQualityWatchByObjectIdRequest, runtime *dara.RuntimeOptions) (_result *GetQualityWatchByObjectIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.WatchObjectId) {
		query["WatchObjectId"] = request.WatchObjectId
	}

	if !dara.IsNil(request.WatchType) {
		query["WatchType"] = request.WatchType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityWatchByObjectId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityWatchByObjectIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a quality watchtask record by the original ID of the monitored object, such as the ID of a datasource, table, or metric.
//
// Release version: v5.4.2.
//
// @param request - GetQualityWatchByObjectIdRequest
//
// @return GetQualityWatchByObjectIdResponse
func (client *Client) GetQualityWatchByObjectId(request *GetQualityWatchByObjectIdRequest) (_result *GetQualityWatchByObjectIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQualityWatchByObjectIdResponse{}
	_body, _err := client.GetQualityWatchByObjectIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a monitoring node task object.
//
// Online version: v5.4.2.
//
// @param request - GetQualityWatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityWatchTaskResponse
func (client *Client) GetQualityWatchTaskWithOptions(request *GetQualityWatchTaskRequest, runtime *dara.RuntimeOptions) (_result *GetQualityWatchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.WatchTaskId) {
		query["WatchTaskId"] = request.WatchTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityWatchTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityWatchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a monitoring node task object.
//
// Online version: v5.4.2.
//
// @param request - GetQualityWatchTaskRequest
//
// @return GetQualityWatchTaskResponse
func (client *Client) GetQualityWatchTask(request *GetQualityWatchTaskRequest) (_result *GetQualityWatchTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQualityWatchTaskResponse{}
	_body, _err := client.GetQualityWatchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the log content of a monitoring task object.
//
// Online version: v5.4.2.
//
// @param request - GetQualityWatchTaskLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityWatchTaskLogResponse
func (client *Client) GetQualityWatchTaskLogWithOptions(request *GetQualityWatchTaskLogRequest, runtime *dara.RuntimeOptions) (_result *GetQualityWatchTaskLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.WatchTaskId) {
		query["WatchTaskId"] = request.WatchTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityWatchTaskLog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityWatchTaskLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the log content of a monitoring task object.
//
// Online version: v5.4.2.
//
// @param request - GetQualityWatchTaskLogRequest
//
// @return GetQualityWatchTaskLogResponse
func (client *Client) GetQualityWatchTaskLog(request *GetQualityWatchTaskLogRequest) (_result *GetQualityWatchTaskLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQualityWatchTaskLogResponse{}
	_body, _err := client.GetQualityWatchTaskLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the cluster version based on the cluster ID.
//
// @param request - GetQueueEngineVersionByEnvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueueEngineVersionByEnvResponse
func (client *Client) GetQueueEngineVersionByEnvWithOptions(request *GetQueueEngineVersionByEnvRequest, runtime *dara.RuntimeOptions) (_result *GetQueueEngineVersionByEnvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.StreamBatchMode) {
		query["StreamBatchMode"] = request.StreamBatchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQueueEngineVersionByEnv"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueueEngineVersionByEnvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the cluster version based on the cluster ID.
//
// @param request - GetQueueEngineVersionByEnvRequest
//
// @return GetQueueEngineVersionByEnvResponse
func (client *Client) GetQueueEngineVersionByEnv(request *GetQueueEngineVersionByEnvRequest) (_result *GetQueueEngineVersionByEnvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQueueEngineVersionByEnvResponse{}
	_body, _err := client.GetQueueEngineVersionByEnvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a resource file.
//
// @param request - GetResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceResponse
func (client *Client) GetResourceWithOptions(request *GetResourceRequest, runtime *dara.RuntimeOptions) (_result *GetResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a resource file.
//
// @param request - GetResourceRequest
//
// @return GetResourceResponse
func (client *Client) GetResource(request *GetResourceRequest) (_result *GetResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceResponse{}
	_body, _err := client.GetResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets the details of a specified version of a resource file.
//
// @param request - GetResourceByVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceByVersionResponse
func (client *Client) GetResourceByVersionWithOptions(request *GetResourceByVersionRequest, runtime *dara.RuntimeOptions) (_result *GetResourceByVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceByVersion"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceByVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the details of a specified version of a resource file.
//
// @param request - GetResourceByVersionRequest
//
// @return GetResourceByVersionResponse
func (client *Client) GetResourceByVersion(request *GetResourceByVersionRequest) (_result *GetResourceByVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceByVersionResponse{}
	_body, _err := client.GetResourceByVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Release version: v5.4.2.
//
// @param tmpReq - GetRowPermissionByTableGuidsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRowPermissionByTableGuidsResponse
func (client *Client) GetRowPermissionByTableGuidsWithOptions(tmpReq *GetRowPermissionByTableGuidsRequest, runtime *dara.RuntimeOptions) (_result *GetRowPermissionByTableGuidsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetRowPermissionByTableGuidsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GetRowPermissionByTableGuidsQuery) {
		request.GetRowPermissionByTableGuidsQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GetRowPermissionByTableGuidsQuery, dara.String("GetRowPermissionByTableGuidsQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GetRowPermissionByTableGuidsQueryShrink) {
		body["GetRowPermissionByTableGuidsQuery"] = request.GetRowPermissionByTableGuidsQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRowPermissionByTableGuids"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRowPermissionByTableGuidsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Release version: v5.4.2.
//
// @param request - GetRowPermissionByTableGuidsRequest
//
// @return GetRowPermissionByTableGuidsResponse
func (client *Client) GetRowPermissionByTableGuids(request *GetRowPermissionByTableGuidsRequest) (_result *GetRowPermissionByTableGuidsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRowPermissionByTableGuidsResponse{}
	_body, _err := client.GetRowPermissionByTableGuidsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data classification. Release version: v5.4.2.
//
// @param request - GetSecurityClassifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecurityClassifyResponse
func (client *Client) GetSecurityClassifyWithOptions(request *GetSecurityClassifyRequest, runtime *dara.RuntimeOptions) (_result *GetSecurityClassifyResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecurityClassify"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecurityClassifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data classification. Release version: v5.4.2.
//
// @param request - GetSecurityClassifyRequest
//
// @return GetSecurityClassifyResponse
func (client *Client) GetSecurityClassify(request *GetSecurityClassifyRequest) (_result *GetSecurityClassifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSecurityClassifyResponse{}
	_body, _err := client.GetSecurityClassifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of an identification result.
//
// Release version: v5.4.2.
//
// @param request - GetSecurityIdentifyResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecurityIdentifyResultResponse
func (client *Client) GetSecurityIdentifyResultWithOptions(request *GetSecurityIdentifyResultRequest, runtime *dara.RuntimeOptions) (_result *GetSecurityIdentifyResultResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecurityIdentifyResult"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecurityIdentifyResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an identification result.
//
// Release version: v5.4.2.
//
// @param request - GetSecurityIdentifyResultRequest
//
// @return GetSecurityIdentifyResultResponse
func (client *Client) GetSecurityIdentifyResult(request *GetSecurityIdentifyResultRequest) (_result *GetSecurityIdentifyResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSecurityIdentifyResultResponse{}
	_body, _err := client.GetSecurityIdentifyResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data classification level. Available since v5.4.2.
//
// @param request - GetSecurityLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecurityLevelResponse
func (client *Client) GetSecurityLevelWithOptions(request *GetSecurityLevelRequest, runtime *dara.RuntimeOptions) (_result *GetSecurityLevelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Index) {
		query["Index"] = request.Index
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecurityLevel"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecurityLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data classification level. Available since v5.4.2.
//
// @param request - GetSecurityLevelRequest
//
// @return GetSecurityLevelResponse
func (client *Client) GetSecurityLevel(request *GetSecurityLevelRequest) (_result *GetSecurityLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSecurityLevelResponse{}
	_body, _err := client.GetSecurityLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a key value by key name. Online version: v5.4.2.
//
// @param request - GetSecuritySecretKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecuritySecretKeyResponse
func (client *Client) GetSecuritySecretKeyWithOptions(request *GetSecuritySecretKeyRequest, runtime *dara.RuntimeOptions) (_result *GetSecuritySecretKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecuritySecretKey"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecuritySecretKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a key value by key name. Online version: v5.4.2.
//
// @param request - GetSecuritySecretKeyRequest
//
// @return GetSecuritySecretKeyResponse
func (client *Client) GetSecuritySecretKey(request *GetSecuritySecretKeyRequest) (_result *GetSecuritySecretKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSecuritySecretKeyResponse{}
	_body, _err := client.GetSecuritySecretKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the Spark client information of the cluster associated with a compute source.
//
// @param request - GetSparkLocalClientInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkLocalClientInfoResponse
func (client *Client) GetSparkLocalClientInfoWithOptions(request *GetSparkLocalClientInfoRequest, runtime *dara.RuntimeOptions) (_result *GetSparkLocalClientInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvEnum) {
		query["EnvEnum"] = request.EnvEnum
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSparkLocalClientInfo"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSparkLocalClientInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the Spark client information of the cluster associated with a compute source.
//
// @param request - GetSparkLocalClientInfoRequest
//
// @return GetSparkLocalClientInfoResponse
func (client *Client) GetSparkLocalClientInfo(request *GetSparkLocalClientInfoRequest) (_result *GetSparkLocalClientInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSparkLocalClientInfoResponse{}
	_body, _err := client.GetSparkLocalClientInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a standard.
//
// Release version: v5.4.2.
//
// @param tmpReq - GetStandardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStandardResponse
func (client *Client) GetStandardWithOptions(tmpReq *GetStandardRequest, runtime *dara.RuntimeOptions) (_result *GetStandardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetStandardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StandardGetQuery) {
		request.StandardGetQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StandardGetQuery, dara.String("StandardGetQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.StandardGetQueryShrink) {
		body["StandardGetQuery"] = request.StandardGetQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStandard"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStandardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a standard.
//
// Release version: v5.4.2.
//
// @param request - GetStandardRequest
//
// @return GetStandardResponse
func (client *Client) GetStandard(request *GetStandardRequest) (_result *GetStandardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStandardResponse{}
	_body, _err := client.GetStandardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data standard lookup table.
//
// Online version: v5.4.2.
//
// @param request - GetStandardLookupTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStandardLookupTableResponse
func (client *Client) GetStandardLookupTableWithOptions(request *GetStandardLookupTableRequest, runtime *dara.RuntimeOptions) (_result *GetStandardLookupTableResponse, _err error) {
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

	if !dara.IsNil(request.Nullable) {
		query["Nullable"] = request.Nullable
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStandardLookupTable"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStandardLookupTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data standard lookup table.
//
// Online version: v5.4.2.
//
// @param request - GetStandardLookupTableRequest
//
// @return GetStandardLookupTableResponse
func (client *Client) GetStandardLookupTable(request *GetStandardLookupTableRequest) (_result *GetStandardLookupTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStandardLookupTableResponse{}
	_body, _err := client.GetStandardLookupTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a standard set.
//
// Release version: v5.4.2.
//
// @param request - GetStandardSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStandardSetResponse
func (client *Client) GetStandardSetWithOptions(request *GetStandardSetRequest, runtime *dara.RuntimeOptions) (_result *GetStandardSetResponse, _err error) {
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

	if !dara.IsNil(request.Nullable) {
		query["Nullable"] = request.Nullable
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStandardSet"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStandardSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a standard set.
//
// Release version: v5.4.2.
//
// @param request - GetStandardSetRequest
//
// @return GetStandardSetResponse
func (client *Client) GetStandardSet(request *GetStandardSetRequest) (_result *GetStandardSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStandardSetResponse{}
	_body, _err := client.GetStandardSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of standards grouped by standard type under a specified folder.
//
// Online version: v5.4.2.
//
// @param tmpReq - GetStandardStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStandardStatisticsResponse
func (client *Client) GetStandardStatisticsWithOptions(tmpReq *GetStandardStatisticsRequest, runtime *dara.RuntimeOptions) (_result *GetStandardStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetStandardStatisticsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StatisticsQuery) {
		request.StatisticsQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StatisticsQuery, dara.String("StatisticsQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.StatisticsQueryShrink) {
		body["StatisticsQuery"] = request.StatisticsQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStandardStatistics"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStandardStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of standards grouped by standard type under a specified folder.
//
// Online version: v5.4.2.
//
// @param request - GetStandardStatisticsRequest
//
// @return GetStandardStatisticsResponse
func (client *Client) GetStandardStatistics(request *GetStandardStatisticsRequest) (_result *GetStandardStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStandardStatisticsResponse{}
	_body, _err := client.GetStandardStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data standard template.
//
// Online version: v5.4.2.
//
// @param tmpReq - GetStandardTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStandardTemplateResponse
func (client *Client) GetStandardTemplateWithOptions(tmpReq *GetStandardTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetStandardTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetStandardTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterQuery) {
		request.FilterQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterQuery, dara.String("FilterQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Nullable) {
		query["Nullable"] = request.Nullable
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FilterQueryShrink) {
		body["FilterQuery"] = request.FilterQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStandardTemplate"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStandardTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data standard template.
//
// Online version: v5.4.2.
//
// @param request - GetStandardTemplateRequest
//
// @return GetStandardTemplateResponse
func (client *Client) GetStandardTemplate(request *GetStandardTemplateRequest) (_result *GetStandardTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStandardTemplateResponse{}
	_body, _err := client.GetStandardTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data standard word root.
//
// Online version: v5.4.2.
//
// @param request - GetStandardWordRootRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStandardWordRootResponse
func (client *Client) GetStandardWordRootWithOptions(request *GetStandardWordRootRequest, runtime *dara.RuntimeOptions) (_result *GetStandardWordRootResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Nullable) {
		query["Nullable"] = request.Nullable
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStandardWordRoot"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStandardWordRootResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data standard word root.
//
// Online version: v5.4.2.
//
// @param request - GetStandardWordRootRequest
//
// @return GetStandardWordRootResponse
func (client *Client) GetStandardWordRoot(request *GetStandardWordRootRequest) (_result *GetStandardWordRootResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStandardWordRootResponse{}
	_body, _err := client.GetStandardWordRootWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of real-time development nodes.
//
// @param request - GetStreamJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStreamJobsResponse
func (client *Client) GetStreamJobsWithOptions(request *GetStreamJobsRequest, runtime *dara.RuntimeOptions) (_result *GetStreamJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStreamJobs"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStreamJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of real-time development nodes.
//
// @param request - GetStreamJobsRequest
//
// @return GetStreamJobsResponse
func (client *Client) GetStreamJobs(request *GetStreamJobsRequest) (_result *GetStreamJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStreamJobsResponse{}
	_body, _err := client.GetStreamJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves dagrun information for all business dates of a data backfill instance workflow.
//
// @param request - GetSupplementDagrunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupplementDagrunResponse
func (client *Client) GetSupplementDagrunWithOptions(request *GetSupplementDagrunRequest, runtime *dara.RuntimeOptions) (_result *GetSupplementDagrunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.SupplementId) {
		query["SupplementId"] = request.SupplementId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSupplementDagrun"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSupplementDagrunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves dagrun information for all business dates of a data backfill instance workflow.
//
// @param request - GetSupplementDagrunRequest
//
// @return GetSupplementDagrunResponse
func (client *Client) GetSupplementDagrun(request *GetSupplementDagrunRequest) (_result *GetSupplementDagrunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSupplementDagrunResponse{}
	_body, _err := client.GetSupplementDagrunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the instances of all nodes for a specific business date in a data backfill workflow.
//
// @param request - GetSupplementDagrunInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupplementDagrunInstanceResponse
func (client *Client) GetSupplementDagrunInstanceWithOptions(request *GetSupplementDagrunInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetSupplementDagrunInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagrunId) {
		query["DagrunId"] = request.DagrunId
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSupplementDagrunInstance"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSupplementDagrunInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the instances of all nodes for a specific business date in a data backfill workflow.
//
// @param request - GetSupplementDagrunInstanceRequest
//
// @return GetSupplementDagrunInstanceResponse
func (client *Client) GetSupplementDagrunInstance(request *GetSupplementDagrunInstanceRequest) (_result *GetSupplementDagrunInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSupplementDagrunInstanceResponse{}
	_body, _err := client.GetSupplementDagrunInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries table column lineage information.
//
// @param tmpReq - GetTableColumnLineageByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableColumnLineageByTaskIdResponse
func (client *Client) GetTableColumnLineageByTaskIdWithOptions(tmpReq *GetTableColumnLineageByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetTableColumnLineageByTaskIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetTableColumnLineageByTaskIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TableColumnLineageByTaskIdQuery) {
		request.TableColumnLineageByTaskIdQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableColumnLineageByTaskIdQuery, dara.String("TableColumnLineageByTaskIdQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TableColumnLineageByTaskIdQueryShrink) {
		body["TableColumnLineageByTaskIdQuery"] = request.TableColumnLineageByTaskIdQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableColumnLineageByTaskId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableColumnLineageByTaskIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries table column lineage information.
//
// @param request - GetTableColumnLineageByTaskIdRequest
//
// @return GetTableColumnLineageByTaskIdResponse
func (client *Client) GetTableColumnLineageByTaskId(request *GetTableColumnLineageByTaskIdRequest) (_result *GetTableColumnLineageByTaskIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTableColumnLineageByTaskIdResponse{}
	_body, _err := client.GetTableColumnLineageByTaskIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the column-level data lineage of an asset table.
//
// Online version: v5.4.2.
//
// @param tmpReq - GetTableColumnLineagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableColumnLineagesResponse
func (client *Client) GetTableColumnLineagesWithOptions(tmpReq *GetTableColumnLineagesRequest, runtime *dara.RuntimeOptions) (_result *GetTableColumnLineagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetTableColumnLineagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterQuery) {
		request.FilterQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterQuery, dara.String("FilterQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FilterQueryShrink) {
		body["FilterQuery"] = request.FilterQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableColumnLineages"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableColumnLineagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the column-level data lineage of an asset table.
//
// Online version: v5.4.2.
//
// @param request - GetTableColumnLineagesRequest
//
// @return GetTableColumnLineagesResponse
func (client *Client) GetTableColumnLineages(request *GetTableColumnLineagesRequest) (_result *GetTableColumnLineagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTableColumnLineagesResponse{}
	_body, _err := client.GetTableColumnLineagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries columns of a Dataphin table in the asset inventory. Supported table types: dimension logical table, fact logical table, aggregate logical table, tag logical table, logical table view, physical table, physical view, and materialized view.
//
// Release version: v5.4.2.
//
// @param request - GetTableColumnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableColumnsResponse
func (client *Client) GetTableColumnsWithOptions(request *GetTableColumnsRequest, runtime *dara.RuntimeOptions) (_result *GetTableColumnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Catalog) {
		query["Catalog"] = request.Catalog
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableColumns"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableColumnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries columns of a Dataphin table in the asset inventory. Supported table types: dimension logical table, fact logical table, aggregate logical table, tag logical table, logical table view, physical table, physical view, and materialized view.
//
// Release version: v5.4.2.
//
// @param request - GetTableColumnsRequest
//
// @return GetTableColumnsResponse
func (client *Client) GetTableColumns(request *GetTableColumnsRequest) (_result *GetTableColumnsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTableColumnsResponse{}
	_body, _err := client.GetTableColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries table lineage information.
//
// @param tmpReq - GetTableLineageByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableLineageByTaskIdResponse
func (client *Client) GetTableLineageByTaskIdWithOptions(tmpReq *GetTableLineageByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetTableLineageByTaskIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetTableLineageByTaskIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TableLineageByTaskIdQuery) {
		request.TableLineageByTaskIdQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableLineageByTaskIdQuery, dara.String("TableLineageByTaskIdQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TableLineageByTaskIdQueryShrink) {
		body["TableLineageByTaskIdQuery"] = request.TableLineageByTaskIdQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableLineageByTaskId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableLineageByTaskIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries table lineage information.
//
// @param request - GetTableLineageByTaskIdRequest
//
// @return GetTableLineageByTaskIdResponse
func (client *Client) GetTableLineageByTaskId(request *GetTableLineageByTaskIdRequest) (_result *GetTableLineageByTaskIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTableLineageByTaskIdResponse{}
	_body, _err := client.GetTableLineageByTaskIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries lineage information of an asset table.
//
// Release version: v5.4.2.
//
// @param tmpReq - GetTableLineagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableLineagesResponse
func (client *Client) GetTableLineagesWithOptions(tmpReq *GetTableLineagesRequest, runtime *dara.RuntimeOptions) (_result *GetTableLineagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetTableLineagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterQuery) {
		request.FilterQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterQuery, dara.String("FilterQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FilterQueryShrink) {
		body["FilterQuery"] = request.FilterQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableLineages"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableLineagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries lineage information of an asset table.
//
// Release version: v5.4.2.
//
// @param request - GetTableLineagesRequest
//
// @return GetTableLineagesResponse
func (client *Client) GetTableLineages(request *GetTableLineagesRequest) (_result *GetTableLineagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTableLineagesResponse{}
	_body, _err := client.GetTableLineagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the progress of a transfer task by transfer task ID.
//
// @param request - GetTransferInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTransferInfoResponse
func (client *Client) GetTransferInfoWithOptions(request *GetTransferInfoRequest, runtime *dara.RuntimeOptions) (_result *GetTransferInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProposalId) {
		query["ProposalId"] = request.ProposalId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTransferInfo"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTransferInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the progress of a transfer task by transfer task ID.
//
// @param request - GetTransferInfoRequest
//
// @return GetTransferInfoResponse
func (client *Client) GetTransferInfo(request *GetTransferInfoRequest) (_result *GetTransferInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTransferInfoResponse{}
	_body, _err := client.GetTransferInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a user-defined function.
//
// @param request - GetUdfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUdfResponse
func (client *Client) GetUdfWithOptions(request *GetUdfRequest, runtime *dara.RuntimeOptions) (_result *GetUdfResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUdf"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUdfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a user-defined function.
//
// @param request - GetUdfRequest
//
// @return GetUdfResponse
func (client *Client) GetUdf(request *GetUdfRequest) (_result *GetUdfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUdfResponse{}
	_body, _err := client.GetUdfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specific version of a user-defined function.
//
// @param request - GetUdfByVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUdfByVersionResponse
func (client *Client) GetUdfByVersionWithOptions(request *GetUdfByVersionRequest, runtime *dara.RuntimeOptions) (_result *GetUdfByVersionResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUdfByVersion"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUdfByVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specific version of a user-defined function.
//
// @param request - GetUdfByVersionRequest
//
// @return GetUdfByVersionResponse
func (client *Client) GetUdfByVersion(request *GetUdfByVersionRequest) (_result *GetUdfByVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUdfByVersionResponse{}
	_body, _err := client.GetUdfByVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves user details by original user ID.
//
// @param request - GetUserBySourceIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserBySourceIdResponse
func (client *Client) GetUserBySourceIdWithOptions(request *GetUserBySourceIdRequest, runtime *dara.RuntimeOptions) (_result *GetUserBySourceIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserBySourceId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserBySourceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves user details by original user ID.
//
// @param request - GetUserBySourceIdRequest
//
// @return GetUserBySourceIdResponse
func (client *Client) GetUserBySourceId(request *GetUserBySourceIdRequest) (_result *GetUserBySourceIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserBySourceIdResponse{}
	_body, _err := client.GetUserBySourceIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a user group.
//
// @param request - GetUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroupWithOptions(request *GetUserGroupRequest, runtime *dara.RuntimeOptions) (_result *GetUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserGroup"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a user group.
//
// @param request - GetUserGroupRequest
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroup(request *GetUserGroupRequest) (_result *GetUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserGroupResponse{}
	_body, _err := client.GetUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves user information in batches by user ID.
//
// @param tmpReq - GetUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUsersResponse
func (client *Client) GetUsersWithOptions(tmpReq *GetUsersRequest, runtime *dara.RuntimeOptions) (_result *GetUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserIdList) {
		request.UserIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIdList, dara.String("UserIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UserIdListShrink) {
		body["UserIdList"] = request.UserIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUsers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves user information in batches by user ID.
//
// @param request - GetUsersRequest
//
// @return GetUsersResponse
func (client *Client) GetUsers(request *GetUsersRequest) (_result *GetUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUsersResponse{}
	_body, _err := client.GetUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants API authorization.
//
// @param tmpReq - GrantDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantDataServiceApiResponse
func (client *Client) GrantDataServiceApiWithOptions(tmpReq *GrantDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *GrantDataServiceApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GrantDataServiceApiShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GrantCommand) {
		request.GrantCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GrantCommand, dara.String("GrantCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GrantCommandShrink) {
		body["GrantCommand"] = request.GrantCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantDataServiceApi"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantDataServiceApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants API authorization.
//
// @param request - GrantDataServiceApiRequest
//
// @return GrantDataServiceApiResponse
func (client *Client) GrantDataServiceApi(request *GrantDataServiceApiRequest) (_result *GrantDataServiceApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GrantDataServiceApiResponse{}
	_body, _err := client.GrantDataServiceApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants permissions on resources to users by resource point.
//
// @param tmpReq - GrantResourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantResourcePermissionResponse
func (client *Client) GrantResourcePermissionWithOptions(tmpReq *GrantResourcePermissionRequest, runtime *dara.RuntimeOptions) (_result *GrantResourcePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GrantResourcePermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GrantCommand) {
		request.GrantCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GrantCommand, dara.String("GrantCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GrantCommandShrink) {
		body["GrantCommand"] = request.GrantCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantResourcePermission"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantResourcePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants permissions on resources to users by resource point.
//
// @param request - GrantResourcePermissionRequest
//
// @return GrantResourcePermissionResponse
func (client *Client) GrantResourcePermission(request *GrantResourcePermissionRequest) (_result *GrantResourcePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GrantResourcePermissionResponse{}
	_body, _err := client.GrantResourcePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the global roles that can be assigned to tenant members. Only built-in global roles are supported. Custom global roles are not supported.
//
// @param request - ListAddableRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddableRolesResponse
func (client *Client) ListAddableRolesWithOptions(request *ListAddableRolesRequest, runtime *dara.RuntimeOptions) (_result *ListAddableRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAddableRoles"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAddableRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the global roles that can be assigned to tenant members. Only built-in global roles are supported. Custom global roles are not supported.
//
// @param request - ListAddableRolesRequest
//
// @return ListAddableRolesResponse
func (client *Client) ListAddableRoles(request *ListAddableRolesRequest) (_result *ListAddableRolesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAddableRolesResponse{}
	_body, _err := client.ListAddableRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries users that can be added to a tenant. Only the super administrator (SuperAdmin) and system administrator can call this operation. The users must already exist in the Dataphin instance member list but have not yet been added to the tenant member list.
//
// @param tmpReq - ListAddableUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddableUsersResponse
func (client *Client) ListAddableUsersWithOptions(tmpReq *ListAddableUsersRequest, runtime *dara.RuntimeOptions) (_result *ListAddableUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAddableUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAddableUsers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAddableUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries users that can be added to a tenant. Only the super administrator (SuperAdmin) and system administrator can call this operation. The users must already exist in the Dataphin instance member list but have not yet been added to the tenant member list.
//
// @param request - ListAddableUsersRequest
//
// @return ListAddableUsersResponse
func (client *Client) ListAddableUsers(request *ListAddableUsersRequest) (_result *ListAddableUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAddableUsersResponse{}
	_body, _err := client.ListAddableUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a conditional query to list multiple alerting events.
//
// @param tmpReq - ListAlertEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertEventsResponse
func (client *Client) ListAlertEventsWithOptions(tmpReq *ListAlertEventsRequest, runtime *dara.RuntimeOptions) (_result *ListAlertEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAlertEventsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlertEvents"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlertEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a conditional query to list multiple alerting events.
//
// @param request - ListAlertEventsRequest
//
// @return ListAlertEventsResponse
func (client *Client) ListAlertEvents(request *ListAlertEventsRequest) (_result *ListAlertEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAlertEventsResponse{}
	_body, _err := client.ListAlertEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a conditional query to list multiple push records.
//
// @param tmpReq - ListAlertNotificationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertNotificationsResponse
func (client *Client) ListAlertNotificationsWithOptions(tmpReq *ListAlertNotificationsRequest, runtime *dara.RuntimeOptions) (_result *ListAlertNotificationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAlertNotificationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlertNotifications"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlertNotificationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a conditional query to list multiple push records.
//
// @param request - ListAlertNotificationsRequest
//
// @return ListAlertNotificationsResponse
func (client *Client) ListAlertNotifications(request *ListAlertNotificationsRequest) (_result *ListAlertNotificationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAlertNotificationsResponse{}
	_body, _err := client.ListAlertNotificationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of APIs by application.
//
// Description:
//
// Queries the detailed information of published APIs by appKey.
//
// @param tmpReq - ListApiByAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiByAppResponse
func (client *Client) ListApiByAppWithOptions(tmpReq *ListApiByAppRequest, runtime *dara.RuntimeOptions) (_result *ListApiByAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListApiByAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PageQuery) {
		request.PageQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PageQuery, dara.String("PageQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PageQueryShrink) {
		body["PageQuery"] = request.PageQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApiByApp"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApiByAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of APIs by application.
//
// Description:
//
// Queries the detailed information of published APIs by appKey.
//
// @param request - ListApiByAppRequest
//
// @return ListApiByAppResponse
func (client *Client) ListApiByApp(request *ListApiByAppRequest) (_result *ListApiByAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApiByAppResponse{}
	_body, _err := client.ListApiByAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of specific fields for APIs that an application has requested.
//
// @param tmpReq - ListAuthorizedDataServiceApiDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizedDataServiceApiDetailsResponse
func (client *Client) ListAuthorizedDataServiceApiDetailsWithOptions(tmpReq *ListAuthorizedDataServiceApiDetailsRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizedDataServiceApiDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAuthorizedDataServiceApiDetailsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizedDataServiceApiDetails"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizedDataServiceApiDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of specific fields for APIs that an application has requested.
//
// @param request - ListAuthorizedDataServiceApiDetailsRequest
//
// @return ListAuthorizedDataServiceApiDetailsResponse
func (client *Client) ListAuthorizedDataServiceApiDetails(request *ListAuthorizedDataServiceApiDetailsRequest) (_result *ListAuthorizedDataServiceApiDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAuthorizedDataServiceApiDetailsResponse{}
	_body, _err := client.ListAuthorizedDataServiceApiDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of business entities.
//
// @param tmpReq - ListBizEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBizEntitiesResponse
func (client *Client) ListBizEntitiesWithOptions(tmpReq *ListBizEntitiesRequest, runtime *dara.RuntimeOptions) (_result *ListBizEntitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListBizEntitiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBizEntities"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBizEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of business entities.
//
// @param request - ListBizEntitiesRequest
//
// @return ListBizEntitiesResponse
func (client *Client) ListBizEntities(request *ListBizEntitiesRequest) (_result *ListBizEntitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBizEntitiesResponse{}
	_body, _err := client.ListBizEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves all business units under the current tenant.
//
// @param request - ListBizUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBizUnitsResponse
func (client *Client) ListBizUnitsWithOptions(request *ListBizUnitsRequest, runtime *dara.RuntimeOptions) (_result *ListBizUnitsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBizUnits"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBizUnitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all business units under the current tenant.
//
// @param request - ListBizUnitsRequest
//
// @return ListBizUnitsResponse
func (client *Client) ListBizUnits(request *ListBizUnitsRequest) (_result *ListBizUnitsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBizUnitsResponse{}
	_body, _err := client.ListBizUnitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of asset catalog entries. Online version: v6.1.0.
//
// @param tmpReq - ListCatalogAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCatalogAssetsResponse
func (client *Client) ListCatalogAssetsWithOptions(tmpReq *ListCatalogAssetsRequest, runtime *dara.RuntimeOptions) (_result *ListCatalogAssetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListCatalogAssetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListCatalogAssetsQuery) {
		request.ListCatalogAssetsQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListCatalogAssetsQuery, dara.String("ListCatalogAssetsQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListCatalogAssetsQueryShrink) {
		body["ListCatalogAssetsQuery"] = request.ListCatalogAssetsQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCatalogAssets"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCatalogAssetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of asset catalog entries. Online version: v6.1.0.
//
// @param request - ListCatalogAssetsRequest
//
// @return ListCatalogAssetsResponse
func (client *Client) ListCatalogAssets(request *ListCatalogAssetsRequest) (_result *ListCatalogAssetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCatalogAssetsResponse{}
	_body, _err := client.ListCatalogAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the list of compute sources.
//
// @param tmpReq - ListComputeSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeSourcesResponse
func (client *Client) ListComputeSourcesWithOptions(tmpReq *ListComputeSourcesRequest, runtime *dara.RuntimeOptions) (_result *ListComputeSourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListComputeSourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComputeSources"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComputeSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of compute sources.
//
// @param request - ListComputeSourcesRequest
//
// @return ListComputeSourcesResponse
func (client *Client) ListComputeSources(request *ListComputeSourcesRequest) (_result *ListComputeSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListComputeSourcesResponse{}
	_body, _err := client.ListComputeSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of data domains.
//
// @param tmpReq - ListDataDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataDomainsResponse
func (client *Client) ListDataDomainsWithOptions(tmpReq *ListDataDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListDataDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataDomainsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataDomains"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of data domains.
//
// @param request - ListDataDomainsRequest
//
// @return ListDataDomainsResponse
func (client *Client) ListDataDomains(request *ListDataDomainsRequest) (_result *ListDataDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataDomainsResponse{}
	_body, _err := client.ListDataDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// O&M analysis: API call statistics.
//
// @param tmpReq - ListDataServiceApiCallStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceApiCallStatisticsResponse
func (client *Client) ListDataServiceApiCallStatisticsWithOptions(tmpReq *ListDataServiceApiCallStatisticsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceApiCallStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataServiceApiCallStatisticsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceApiCallStatistics"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceApiCallStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// O&M analysis: API call statistics.
//
// @param request - ListDataServiceApiCallStatisticsRequest
//
// @return ListDataServiceApiCallStatisticsResponse
func (client *Client) ListDataServiceApiCallStatistics(request *ListDataServiceApiCallStatisticsRequest) (_result *ListDataServiceApiCallStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataServiceApiCallStatisticsResponse{}
	_body, _err := client.ListDataServiceApiCallStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries data service call logs with pagination.
//
// @param tmpReq - ListDataServiceApiCallsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceApiCallsResponse
func (client *Client) ListDataServiceApiCallsWithOptions(tmpReq *ListDataServiceApiCallsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceApiCallsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataServiceApiCallsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceApiCalls"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceApiCallsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries data service call logs with pagination.
//
// @param request - ListDataServiceApiCallsRequest
//
// @return ListDataServiceApiCallsResponse
func (client *Client) ListDataServiceApiCalls(request *ListDataServiceApiCallsRequest) (_result *ListDataServiceApiCallsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataServiceApiCallsResponse{}
	_body, _err := client.ListDataServiceApiCallsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Operations analysis: analyzes the impact of abnormal API calls.
//
// @param tmpReq - ListDataServiceApiImpactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceApiImpactsResponse
func (client *Client) ListDataServiceApiImpactsWithOptions(tmpReq *ListDataServiceApiImpactsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceApiImpactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataServiceApiImpactsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceApiImpacts"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceApiImpactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Operations analysis: analyzes the impact of abnormal API calls.
//
// @param request - ListDataServiceApiImpactsRequest
//
// @return ListDataServiceApiImpactsResponse
func (client *Client) ListDataServiceApiImpacts(request *ListDataServiceApiImpactsRequest) (_result *ListDataServiceApiImpactsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataServiceApiImpactsResponse{}
	_body, _err := client.ListDataServiceApiImpactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of all applications under a data service tenant. All tenant members can perform this operation.
//
// Release version: v6.0.0.
//
// @param tmpReq - ListDataServiceAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceAppsResponse
func (client *Client) ListDataServiceAppsWithOptions(tmpReq *ListDataServiceAppsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataServiceAppsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceApps"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of all applications under a data service tenant. All tenant members can perform this operation.
//
// Release version: v6.0.0.
//
// @param request - ListDataServiceAppsRequest
//
// @return ListDataServiceAppsResponse
func (client *Client) ListDataServiceApps(request *ListDataServiceAppsRequest) (_result *ListDataServiceAppsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataServiceAppsResponse{}
	_body, _err := client.ListDataServiceAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of applications that the current user has permissions to access.
//
// @param tmpReq - ListDataServiceAuthorizedAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceAuthorizedAppsResponse
func (client *Client) ListDataServiceAuthorizedAppsWithOptions(tmpReq *ListDataServiceAuthorizedAppsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceAuthorizedAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataServiceAuthorizedAppsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceAuthorizedApps"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceAuthorizedAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of applications that the current user has permissions to access.
//
// @param request - ListDataServiceAuthorizedAppsRequest
//
// @return ListDataServiceAuthorizedAppsResponse
func (client *Client) ListDataServiceAuthorizedApps(request *ListDataServiceAuthorizedAppsRequest) (_result *ListDataServiceAuthorizedAppsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataServiceAuthorizedAppsResponse{}
	_body, _err := client.ListDataServiceAuthorizedAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the list of API permissions managed by me.
//
// @param tmpReq - ListDataServiceMyApiPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceMyApiPermissionsResponse
func (client *Client) ListDataServiceMyApiPermissionsWithOptions(tmpReq *ListDataServiceMyApiPermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceMyApiPermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataServiceMyApiPermissionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceMyApiPermissions"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceMyApiPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the list of API permissions managed by me.
//
// @param request - ListDataServiceMyApiPermissionsRequest
//
// @return ListDataServiceMyApiPermissionsResponse
func (client *Client) ListDataServiceMyApiPermissions(request *ListDataServiceMyApiPermissionsRequest) (_result *ListDataServiceMyApiPermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataServiceMyApiPermissionsResponse{}
	_body, _err := client.ListDataServiceMyApiPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the applications that the current user has permissions to access.
//
// @param tmpReq - ListDataServiceMyAppPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceMyAppPermissionsResponse
func (client *Client) ListDataServiceMyAppPermissionsWithOptions(tmpReq *ListDataServiceMyAppPermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceMyAppPermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataServiceMyAppPermissionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceMyAppPermissions"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceMyAppPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the applications that the current user has permissions to access.
//
// @param request - ListDataServiceMyAppPermissionsRequest
//
// @return ListDataServiceMyAppPermissionsResponse
func (client *Client) ListDataServiceMyAppPermissions(request *ListDataServiceMyAppPermissionsRequest) (_result *ListDataServiceMyAppPermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataServiceMyAppPermissionsResponse{}
	_body, _err := client.ListDataServiceMyAppPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of published APIs by page.
//
// @param tmpReq - ListDataServicePublishedApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServicePublishedApisResponse
func (client *Client) ListDataServicePublishedApisWithOptions(tmpReq *ListDataServicePublishedApisRequest, runtime *dara.RuntimeOptions) (_result *ListDataServicePublishedApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataServicePublishedApisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServicePublishedApis"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServicePublishedApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of published APIs by page.
//
// @param request - ListDataServicePublishedApisRequest
//
// @return ListDataServicePublishedApisResponse
func (client *Client) ListDataServicePublishedApis(request *ListDataServicePublishedApisRequest) (_result *ListDataServicePublishedApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataServicePublishedApisResponse{}
	_body, _err := client.ListDataServicePublishedApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Search for data sources. The results include data source configuration items.
//
// @param tmpReq - ListDataSourceWithConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceWithConfigResponse
func (client *Client) ListDataSourceWithConfigWithOptions(tmpReq *ListDataSourceWithConfigRequest, runtime *dara.RuntimeOptions) (_result *ListDataSourceWithConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataSourceWithConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSourceWithConfig"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSourceWithConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Search for data sources. The results include data source configuration items.
//
// @param request - ListDataSourceWithConfigRequest
//
// @return ListDataSourceWithConfigResponse
func (client *Client) ListDataSourceWithConfig(request *ListDataSourceWithConfigRequest) (_result *ListDataSourceWithConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataSourceWithConfigResponse{}
	_body, _err := client.ListDataSourceWithConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the directory tree file list.
//
// @param tmpReq - ListFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFilesResponse
func (client *Client) ListFilesWithOptions(tmpReq *ListFilesRequest, runtime *dara.RuntimeOptions) (_result *ListFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListFilesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFiles"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the directory tree file list.
//
// @param request - ListFilesRequest
//
// @return ListFilesResponse
func (client *Client) ListFiles(request *ListFilesRequest) (_result *ListFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFilesResponse{}
	_body, _err := client.ListFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Paginate and query instances.
//
// @param tmpReq - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(tmpReq *ListInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Paginate and query instances.
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the downstream nodes of a node. The query results can be used as a data reference when you create a data backfill workflow.
//
// @param tmpReq - ListNodeDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeDownStreamResponse
func (client *Client) ListNodeDownStreamWithOptions(tmpReq *ListNodeDownStreamRequest, runtime *dara.RuntimeOptions) (_result *ListNodeDownStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListNodeDownStreamShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodeDownStream"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodeDownStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the downstream nodes of a node. The query results can be used as a data reference when you create a data backfill workflow.
//
// @param request - ListNodeDownStreamRequest
//
// @return ListNodeDownStreamResponse
func (client *Client) ListNodeDownStream(request *ListNodeDownStreamRequest) (_result *ListNodeDownStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNodeDownStreamResponse{}
	_body, _err := client.ListNodeDownStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of scheduling nodes.
//
// @param tmpReq - ListNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithOptions(tmpReq *ListNodesRequest, runtime *dara.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodes"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of scheduling nodes.
//
// @param request - ListNodesRequest
//
// @return ListNodesResponse
func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of project members.
//
// @param tmpReq - ListProjectMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectMembersResponse
func (client *Client) ListProjectMembersWithOptions(tmpReq *ListProjectMembersRequest, runtime *dara.RuntimeOptions) (_result *ListProjectMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListProjectMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectMembers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of project members.
//
// @param request - ListProjectMembersRequest
//
// @return ListProjectMembersResponse
func (client *Client) ListProjectMembers(request *ListProjectMembersRequest) (_result *ListProjectMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectMembersResponse{}
	_body, _err := client.ListProjectMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of projects.
//
// @param tmpReq - ListProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithOptions(tmpReq *ListProjectsRequest, runtime *dara.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListProjectsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjects"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of projects.
//
// @param request - ListProjectsRequest
//
// @return ListProjectsResponse
func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of publish records.
//
// @param tmpReq - ListPublishRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublishRecordsResponse
func (client *Client) ListPublishRecordsWithOptions(tmpReq *ListPublishRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListPublishRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPublishRecordsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPublishRecords"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPublishRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of publish records.
//
// @param request - ListPublishRecordsRequest
//
// @return ListPublishRecordsResponse
func (client *Client) ListPublishRecords(request *ListPublishRecordsRequest) (_result *ListPublishRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPublishRecordsResponse{}
	_body, _err := client.ListPublishRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries quality rule tasks by paging.
//
// Online version: v5.4.2.
//
// @param tmpReq - ListQualityRuleTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQualityRuleTasksResponse
func (client *Client) ListQualityRuleTasksWithOptions(tmpReq *ListQualityRuleTasksRequest, runtime *dara.RuntimeOptions) (_result *ListQualityRuleTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListQualityRuleTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQualityRuleTasks"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQualityRuleTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quality rule tasks by paging.
//
// Online version: v5.4.2.
//
// @param request - ListQualityRuleTasksRequest
//
// @return ListQualityRuleTasksResponse
func (client *Client) ListQualityRuleTasks(request *ListQualityRuleTasksRequest) (_result *ListQualityRuleTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQualityRuleTasksResponse{}
	_body, _err := client.ListQualityRuleTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries quality rules by paging.
//
// Online version: v5.4.2.
//
// @param tmpReq - ListQualityRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQualityRulesResponse
func (client *Client) ListQualityRulesWithOptions(tmpReq *ListQualityRulesRequest, runtime *dara.RuntimeOptions) (_result *ListQualityRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListQualityRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQualityRules"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQualityRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quality rules by paging.
//
// Online version: v5.4.2.
//
// @param request - ListQualityRulesRequest
//
// @return ListQualityRulesResponse
func (client *Client) ListQualityRules(request *ListQualityRulesRequest) (_result *ListQualityRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQualityRulesResponse{}
	_body, _err := client.ListQualityRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries quality templates by paging.
//
// Online version: v5.4.2.
//
// @param tmpReq - ListQualityTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQualityTemplatesResponse
func (client *Client) ListQualityTemplatesWithOptions(tmpReq *ListQualityTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListQualityTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListQualityTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQualityTemplates"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQualityTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quality templates by paging.
//
// Online version: v5.4.2.
//
// @param request - ListQualityTemplatesRequest
//
// @return ListQualityTemplatesResponse
func (client *Client) ListQualityTemplates(request *ListQualityTemplatesRequest) (_result *ListQualityTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQualityTemplatesResponse{}
	_body, _err := client.ListQualityTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries quality monitoring nodes by paged query.
//
// Online version: v5.4.2.
//
// @param tmpReq - ListQualityWatchTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQualityWatchTasksResponse
func (client *Client) ListQualityWatchTasksWithOptions(tmpReq *ListQualityWatchTasksRequest, runtime *dara.RuntimeOptions) (_result *ListQualityWatchTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListQualityWatchTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQualityWatchTasks"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQualityWatchTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quality monitoring nodes by paged query.
//
// Online version: v5.4.2.
//
// @param request - ListQualityWatchTasksRequest
//
// @return ListQualityWatchTasksResponse
func (client *Client) ListQualityWatchTasks(request *ListQualityWatchTasksRequest) (_result *ListQualityWatchTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQualityWatchTasksResponse{}
	_body, _err := client.ListQualityWatchTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paged query of quality monitored objects.
//
// Online version: v5.4.2.
//
// @param tmpReq - ListQualityWatchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQualityWatchesResponse
func (client *Client) ListQualityWatchesWithOptions(tmpReq *ListQualityWatchesRequest, runtime *dara.RuntimeOptions) (_result *ListQualityWatchesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListQualityWatchesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQualityWatches"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQualityWatchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query of quality monitored objects.
//
// Online version: v5.4.2.
//
// @param request - ListQualityWatchesRequest
//
// @return ListQualityWatchesResponse
func (client *Client) ListQualityWatches(request *ListQualityWatchesRequest) (_result *ListQualityWatchesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQualityWatchesResponse{}
	_body, _err := client.ListQualityWatchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of permission operation logs.
//
// @param tmpReq - ListResourcePermissionOperationLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcePermissionOperationLogResponse
func (client *Client) ListResourcePermissionOperationLogWithOptions(tmpReq *ListResourcePermissionOperationLogRequest, runtime *dara.RuntimeOptions) (_result *ListResourcePermissionOperationLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListResourcePermissionOperationLogShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourcePermissionOperationLog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourcePermissionOperationLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of permission operation logs.
//
// @param request - ListResourcePermissionOperationLogRequest
//
// @return ListResourcePermissionOperationLogResponse
func (client *Client) ListResourcePermissionOperationLog(request *ListResourcePermissionOperationLogRequest) (_result *ListResourcePermissionOperationLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourcePermissionOperationLogResponse{}
	_body, _err := client.ListResourcePermissionOperationLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves permission authorization records with pagination.
//
// @param tmpReq - ListResourcePermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcePermissionsResponse
func (client *Client) ListResourcePermissionsWithOptions(tmpReq *ListResourcePermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListResourcePermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListResourcePermissionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourcePermissions"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourcePermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves permission authorization records with pagination.
//
// @param request - ListResourcePermissionsRequest
//
// @return ListResourcePermissionsResponse
func (client *Client) ListResourcePermissions(request *ListResourcePermissionsRequest) (_result *ListResourcePermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourcePermissionsResponse{}
	_body, _err := client.ListResourcePermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paged query of row-level permissions.
//
// @param tmpReq - ListRowPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRowPermissionResponse
func (client *Client) ListRowPermissionWithOptions(tmpReq *ListRowPermissionRequest, runtime *dara.RuntimeOptions) (_result *ListRowPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListRowPermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PageRowPermissionQuery) {
		request.PageRowPermissionQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PageRowPermissionQuery, dara.String("PageRowPermissionQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PageRowPermissionQueryShrink) {
		body["PageRowPermissionQuery"] = request.PageRowPermissionQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRowPermission"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRowPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query of row-level permissions.
//
// @param request - ListRowPermissionRequest
//
// @return ListRowPermissionResponse
func (client *Client) ListRowPermission(request *ListRowPermissionRequest) (_result *ListRowPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRowPermissionResponse{}
	_body, _err := client.ListRowPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries row-level permissions of a specified user by paging.
//
// @param tmpReq - ListRowPermissionByUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRowPermissionByUserIdResponse
func (client *Client) ListRowPermissionByUserIdWithOptions(tmpReq *ListRowPermissionByUserIdRequest, runtime *dara.RuntimeOptions) (_result *ListRowPermissionByUserIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListRowPermissionByUserIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListRowPermissionByUserIdQuery) {
		request.ListRowPermissionByUserIdQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListRowPermissionByUserIdQuery, dara.String("ListRowPermissionByUserIdQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListRowPermissionByUserIdQueryShrink) {
		body["ListRowPermissionByUserIdQuery"] = request.ListRowPermissionByUserIdQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRowPermissionByUserId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRowPermissionByUserIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries row-level permissions of a specified user by paging.
//
// @param request - ListRowPermissionByUserIdRequest
//
// @return ListRowPermissionByUserIdResponse
func (client *Client) ListRowPermissionByUserId(request *ListRowPermissionByUserIdRequest) (_result *ListRowPermissionByUserIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRowPermissionByUserIdResponse{}
	_body, _err := client.ListRowPermissionByUserIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries identification records of security identification results by paging.
//
// Online version: v5.4.2.
//
// @param tmpReq - ListSecurityIdentifyRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecurityIdentifyRecordsResponse
func (client *Client) ListSecurityIdentifyRecordsWithOptions(tmpReq *ListSecurityIdentifyRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListSecurityIdentifyRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSecurityIdentifyRecordsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSecurityIdentifyRecords"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecurityIdentifyRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries identification records of security identification results by paging.
//
// Online version: v5.4.2.
//
// @param request - ListSecurityIdentifyRecordsRequest
//
// @return ListSecurityIdentifyRecordsResponse
func (client *Client) ListSecurityIdentifyRecords(request *ListSecurityIdentifyRecordsRequest) (_result *ListSecurityIdentifyRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSecurityIdentifyRecordsResponse{}
	_body, _err := client.ListSecurityIdentifyRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query security identification results by page.
//
// Release version: v5.4.2.
//
// @param tmpReq - ListSecurityIdentifyResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecurityIdentifyResultsResponse
func (client *Client) ListSecurityIdentifyResultsWithOptions(tmpReq *ListSecurityIdentifyResultsRequest, runtime *dara.RuntimeOptions) (_result *ListSecurityIdentifyResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSecurityIdentifyResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSecurityIdentifyResults"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecurityIdentifyResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query security identification results by page.
//
// Release version: v5.4.2.
//
// @param request - ListSecurityIdentifyResultsRequest
//
// @return ListSecurityIdentifyResultsResponse
func (client *Client) ListSecurityIdentifyResults(request *ListSecurityIdentifyResultsRequest) (_result *ListSecurityIdentifyResultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSecurityIdentifyResultsResponse{}
	_body, _err := client.ListSecurityIdentifyResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the standard list by page.
//
// Release version: v5.4.2.
//
// @param tmpReq - ListStandardsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStandardsResponse
func (client *Client) ListStandardsWithOptions(tmpReq *ListStandardsRequest, runtime *dara.RuntimeOptions) (_result *ListStandardsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListStandardsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStandards"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStandardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the standard list by page.
//
// Release version: v5.4.2.
//
// @param request - ListStandardsRequest
//
// @return ListStandardsResponse
func (client *Client) ListStandards(request *ListStandardsRequest) (_result *ListStandardsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListStandardsResponse{}
	_body, _err := client.ListStandardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Paginate and retrieve the list of pending deployment records.
//
// @param tmpReq - ListSubmitRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubmitRecordsResponse
func (client *Client) ListSubmitRecordsWithOptions(tmpReq *ListSubmitRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListSubmitRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSubmitRecordsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSubmitRecords"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSubmitRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Paginate and retrieve the list of pending deployment records.
//
// @param request - ListSubmitRecordsRequest
//
// @return ListSubmitRecordsResponse
func (client *Client) ListSubmitRecords(request *ListSubmitRecordsRequest) (_result *ListSubmitRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSubmitRecordsResponse{}
	_body, _err := client.ListSubmitRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paged query to retrieve asset table metadata.
//
// Online version: v5.4.2.
//
// @param tmpReq - ListTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTablesResponse
func (client *Client) ListTablesWithOptions(tmpReq *ListTablesRequest, runtime *dara.RuntimeOptions) (_result *ListTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTablesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTables"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query to retrieve asset table metadata.
//
// Online version: v5.4.2.
//
// @param request - ListTablesRequest
//
// @return ListTablesResponse
func (client *Client) ListTables(request *ListTablesRequest) (_result *ListTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTablesResponse{}
	_body, _err := client.ListTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of tenant members.
//
// @param tmpReq - ListTenantMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTenantMembersResponse
func (client *Client) ListTenantMembersWithOptions(tmpReq *ListTenantMembersRequest, runtime *dara.RuntimeOptions) (_result *ListTenantMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTenantMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTenantMembers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTenantMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of tenant members.
//
// @param request - ListTenantMembersRequest
//
// @return ListTenantMembersResponse
func (client *Client) ListTenantMembers(request *ListTenantMembersRequest) (_result *ListTenantMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTenantMembersResponse{}
	_body, _err := client.ListTenantMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paging query of user group members.
//
// @param tmpReq - ListUserGroupMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupMembersResponse
func (client *Client) ListUserGroupMembersWithOptions(tmpReq *ListUserGroupMembersRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListUserGroupMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserGroupMembers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserGroupMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paging query of user group members.
//
// @param request - ListUserGroupMembersRequest
//
// @return ListUserGroupMembersResponse
func (client *Client) ListUserGroupMembers(request *ListUserGroupMembersRequest) (_result *ListUserGroupMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserGroupMembersResponse{}
	_body, _err := client.ListUserGroupMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries user groups by paging.
//
// @param tmpReq - ListUserGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroupsWithOptions(tmpReq *ListUserGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListUserGroupsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserGroups"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user groups by paging.
//
// @param request - ListUserGroupsRequest
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroups(request *ListUserGroupsRequest) (_result *ListUserGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.ListUserGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Offlines a batch task.
//
// @param request - OfflineBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineBatchTaskResponse
func (client *Client) OfflineBatchTaskWithOptions(request *OfflineBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *OfflineBatchTaskResponse, _err error) {
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

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineBatchTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineBatchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Offlines a batch task.
//
// @param request - OfflineBatchTaskRequest
//
// @return OfflineBatchTaskResponse
func (client *Client) OfflineBatchTask(request *OfflineBatchTaskRequest) (_result *OfflineBatchTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OfflineBatchTaskResponse{}
	_body, _err := client.OfflineBatchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Offline a business entity.
//
// @param tmpReq - OfflineBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineBizEntityResponse
func (client *Client) OfflineBizEntityWithOptions(tmpReq *OfflineBizEntityRequest, runtime *dara.RuntimeOptions) (_result *OfflineBizEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OfflineBizEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OfflineCommand) {
		request.OfflineCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OfflineCommand, dara.String("OfflineCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OfflineCommandShrink) {
		body["OfflineCommand"] = request.OfflineCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineBizEntity"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineBizEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Offline a business entity.
//
// @param request - OfflineBizEntityRequest
//
// @return OfflineBizEntityResponse
func (client *Client) OfflineBizEntity(request *OfflineBizEntityRequest) (_result *OfflineBizEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OfflineBizEntityResponse{}
	_body, _err := client.OfflineBizEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Offlines an integration pipeline node.
//
// @param tmpReq - OfflinePipelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflinePipelineResponse
func (client *Client) OfflinePipelineWithOptions(tmpReq *OfflinePipelineRequest, runtime *dara.RuntimeOptions) (_result *OfflinePipelineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OfflinePipelineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Context) {
		request.ContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Context, dara.String("Context"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OfflineCommand) {
		request.OfflineCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OfflineCommand, dara.String("OfflineCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextShrink) {
		body["Context"] = request.ContextShrink
	}

	if !dara.IsNil(request.OfflineCommandShrink) {
		body["OfflineCommand"] = request.OfflineCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflinePipeline"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflinePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Offlines an integration pipeline node.
//
// @param request - OfflinePipelineRequest
//
// @return OfflinePipelineResponse
func (client *Client) OfflinePipeline(request *OfflinePipelineRequest) (_result *OfflinePipelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OfflinePipelineResponse{}
	_body, _err := client.OfflinePipelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously offlines an integration pipeline node.
//
// @param tmpReq - OfflinePipelineByAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflinePipelineByAsyncResponse
func (client *Client) OfflinePipelineByAsyncWithOptions(tmpReq *OfflinePipelineByAsyncRequest, runtime *dara.RuntimeOptions) (_result *OfflinePipelineByAsyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OfflinePipelineByAsyncShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Context) {
		request.ContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Context, dara.String("Context"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OfflineCommand) {
		request.OfflineCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OfflineCommand, dara.String("OfflineCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextShrink) {
		body["Context"] = request.ContextShrink
	}

	if !dara.IsNil(request.OfflineCommandShrink) {
		body["OfflineCommand"] = request.OfflineCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflinePipelineByAsync"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflinePipelineByAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously offlines an integration pipeline node.
//
// @param request - OfflinePipelineByAsyncRequest
//
// @return OfflinePipelineByAsyncResponse
func (client *Client) OfflinePipelineByAsync(request *OfflinePipelineByAsyncRequest) (_result *OfflinePipelineByAsyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OfflinePipelineByAsyncResponse{}
	_body, _err := client.OfflinePipelineByAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Offlines a standard.
//
// Online version: v5.4.2.
//
// @param tmpReq - OfflineStandardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineStandardResponse
func (client *Client) OfflineStandardWithOptions(tmpReq *OfflineStandardRequest, runtime *dara.RuntimeOptions) (_result *OfflineStandardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OfflineStandardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OfflineCommand) {
		request.OfflineCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OfflineCommand, dara.String("OfflineCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OfflineCommandShrink) {
		body["OfflineCommand"] = request.OfflineCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineStandard"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineStandardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Offlines a standard.
//
// Online version: v5.4.2.
//
// @param request - OfflineStandardRequest
//
// @return OfflineStandardResponse
func (client *Client) OfflineStandard(request *OfflineStandardRequest) (_result *OfflineStandardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OfflineStandardResponse{}
	_body, _err := client.OfflineStandardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Brings a business entity online.
//
// @param tmpReq - OnlineBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineBizEntityResponse
func (client *Client) OnlineBizEntityWithOptions(tmpReq *OnlineBizEntityRequest, runtime *dara.RuntimeOptions) (_result *OnlineBizEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OnlineBizEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OnlineCommand) {
		request.OnlineCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OnlineCommand, dara.String("OnlineCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OnlineCommandShrink) {
		body["OnlineCommand"] = request.OnlineCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OnlineBizEntity"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OnlineBizEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Brings a business entity online.
//
// @param request - OnlineBizEntityRequest
//
// @return OnlineBizEntityResponse
func (client *Client) OnlineBizEntity(request *OnlineBizEntityRequest) (_result *OnlineBizEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OnlineBizEntityResponse{}
	_body, _err := client.OnlineBizEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs batch O&M operations on instances. Both physical instances and logical table instances are supported.
//
// @param tmpReq - OperateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateInstanceResponse
func (client *Client) OperateInstanceWithOptions(tmpReq *OperateInstanceRequest, runtime *dara.RuntimeOptions) (_result *OperateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OperateInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OperateCommand) {
		request.OperateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperateCommand, dara.String("OperateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OperateCommandShrink) {
		body["OperateCommand"] = request.OperateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateInstance"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs batch O&M operations on instances. Both physical instances and logical table instances are supported.
//
// @param request - OperateInstanceRequest
//
// @return OperateInstanceResponse
func (client *Client) OperateInstance(request *OperateInstanceRequest) (_result *OperateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateInstanceResponse{}
	_body, _err := client.OperateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Parses the logical table dependencies of an offline compute node. The parsing result may contain self-dependent nodes in the upstream dependency information, where the upstream node ID is the same as the node ID of the parsed code. You must handle such cases on your own.
//
// @param tmpReq - ParseBatchTaskDependencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ParseBatchTaskDependencyResponse
func (client *Client) ParseBatchTaskDependencyWithOptions(tmpReq *ParseBatchTaskDependencyRequest, runtime *dara.RuntimeOptions) (_result *ParseBatchTaskDependencyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ParseBatchTaskDependencyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ParseCommand) {
		request.ParseCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParseCommand, dara.String("ParseCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ParseCommandShrink) {
		body["ParseCommand"] = request.ParseCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ParseBatchTaskDependency"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ParseBatchTaskDependencyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Parses the logical table dependencies of an offline compute node. The parsing result may contain self-dependent nodes in the upstream dependency information, where the upstream node ID is the same as the node ID of the parsed code. You must handle such cases on your own.
//
// @param request - ParseBatchTaskDependencyRequest
//
// @return ParseBatchTaskDependencyResponse
func (client *Client) ParseBatchTaskDependency(request *ParseBatchTaskDependencyRequest) (_result *ParseBatchTaskDependencyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ParseBatchTaskDependencyResponse{}
	_body, _err := client.ParseBatchTaskDependencyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pauses the scheduling of physical nodes. This stops the scheduling of nodes, and downstream nodes cannot be triggered. Currently, only offline code nodes and integration nodes are supported.
//
// @param tmpReq - PausePhysicalNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PausePhysicalNodeResponse
func (client *Client) PausePhysicalNodeWithOptions(tmpReq *PausePhysicalNodeRequest, runtime *dara.RuntimeOptions) (_result *PausePhysicalNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PausePhysicalNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PauseCommand) {
		request.PauseCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PauseCommand, dara.String("PauseCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PauseCommandShrink) {
		body["PauseCommand"] = request.PauseCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PausePhysicalNode"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PausePhysicalNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pauses the scheduling of physical nodes. This stops the scheduling of nodes, and downstream nodes cannot be triggered. Currently, only offline code nodes and integration nodes are supported.
//
// @param request - PausePhysicalNodeRequest
//
// @return PausePhysicalNodeResponse
func (client *Client) PausePhysicalNode(request *PausePhysicalNodeRequest) (_result *PausePhysicalNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PausePhysicalNodeResponse{}
	_body, _err := client.PausePhysicalNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes a data service API to the production environment.
//
// @param request - PublishDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishDataServiceApiResponse
func (client *Client) PublishDataServiceApiWithOptions(request *PublishDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *PublishDataServiceApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishDataServiceApi"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishDataServiceApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a data service API to the production environment.
//
// @param request - PublishDataServiceApiRequest
//
// @return PublishDataServiceApiResponse
func (client *Client) PublishDataServiceApi(request *PublishDataServiceApiRequest) (_result *PublishDataServiceApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishDataServiceApiResponse{}
	_body, _err := client.PublishDataServiceApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes objects in batches.
//
// @param tmpReq - PublishObjectListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishObjectListResponse
func (client *Client) PublishObjectListWithOptions(tmpReq *PublishObjectListRequest, runtime *dara.RuntimeOptions) (_result *PublishObjectListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PublishObjectListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PublishCommand) {
		request.PublishCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PublishCommand, dara.String("PublishCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PublishCommandShrink) {
		body["PublishCommand"] = request.PublishCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishObjectList"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishObjectListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes objects in batches.
//
// @param request - PublishObjectListRequest
//
// @return PublishObjectListResponse
func (client *Client) PublishObjectList(request *PublishObjectListRequest) (_result *PublishObjectListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishObjectListResponse{}
	_body, _err := client.PublishObjectListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes a standard.
//
// Release version: v5.4.2.
//
// @param tmpReq - PublishStandardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishStandardResponse
func (client *Client) PublishStandardWithOptions(tmpReq *PublishStandardRequest, runtime *dara.RuntimeOptions) (_result *PublishStandardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PublishStandardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PublishCommand) {
		request.PublishCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PublishCommand, dara.String("PublishCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PublishCommandShrink) {
		body["PublishCommand"] = request.PublishCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishStandard"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishStandardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a standard.
//
// Release version: v5.4.2.
//
// @param request - PublishStandardRequest
//
// @return PublishStandardResponse
func (client *Client) PublishStandard(request *PublishStandardRequest) (_result *PublishStandardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishStandardResponse{}
	_body, _err := client.PublishStandardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Remove regular members from a data service application. Only the application owner can perform this operation.
//
// Released version: v6.0.0.
//
// @param tmpReq - RemoveDataServiceAppMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDataServiceAppMemberResponse
func (client *Client) RemoveDataServiceAppMemberWithOptions(tmpReq *RemoveDataServiceAppMemberRequest, runtime *dara.RuntimeOptions) (_result *RemoveDataServiceAppMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveDataServiceAppMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RemoveCommand) {
		request.RemoveCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RemoveCommand, dara.String("RemoveCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RemoveCommandShrink) {
		body["RemoveCommand"] = request.RemoveCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveDataServiceAppMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveDataServiceAppMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Remove regular members from a data service application. Only the application owner can perform this operation.
//
// Released version: v6.0.0.
//
// @param request - RemoveDataServiceAppMemberRequest
//
// @return RemoveDataServiceAppMemberResponse
func (client *Client) RemoveDataServiceAppMember(request *RemoveDataServiceAppMemberRequest) (_result *RemoveDataServiceAppMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveDataServiceAppMemberResponse{}
	_body, _err := client.RemoveDataServiceAppMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a project member.
//
// @param tmpReq - RemoveProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveProjectMemberResponse
func (client *Client) RemoveProjectMemberWithOptions(tmpReq *RemoveProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *RemoveProjectMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveProjectMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RemoveCommand) {
		request.RemoveCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RemoveCommand, dara.String("RemoveCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RemoveCommandShrink) {
		body["RemoveCommand"] = request.RemoveCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveProjectMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveProjectMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a project member.
//
// @param request - RemoveProjectMemberRequest
//
// @return RemoveProjectMemberResponse
func (client *Client) RemoveProjectMember(request *RemoveProjectMemberRequest) (_result *RemoveProjectMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveProjectMemberResponse{}
	_body, _err := client.RemoveProjectMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the bindings between quality rules and schedules in batches.
//
// Release version: v5.4.2.
//
// @param tmpReq - RemoveQualityRuleSchedulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveQualityRuleSchedulesResponse
func (client *Client) RemoveQualityRuleSchedulesWithOptions(tmpReq *RemoveQualityRuleSchedulesRequest, runtime *dara.RuntimeOptions) (_result *RemoveQualityRuleSchedulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveQualityRuleSchedulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RemoveCommand) {
		request.RemoveCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RemoveCommand, dara.String("RemoveCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RemoveCommandShrink) {
		body["RemoveCommand"] = request.RemoveCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveQualityRuleSchedules"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveQualityRuleSchedulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the bindings between quality rules and schedules in batches.
//
// Release version: v5.4.2.
//
// @param request - RemoveQualityRuleSchedulesRequest
//
// @return RemoveQualityRuleSchedulesResponse
func (client *Client) RemoveQualityRuleSchedules(request *RemoveQualityRuleSchedulesRequest) (_result *RemoveQualityRuleSchedulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveQualityRuleSchedulesResponse{}
	_body, _err := client.RemoveQualityRuleSchedulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a tenant member. Only superusers and system administrators can call this API operation.
//
// @param tmpReq - RemoveTenantMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTenantMemberResponse
func (client *Client) RemoveTenantMemberWithOptions(tmpReq *RemoveTenantMemberRequest, runtime *dara.RuntimeOptions) (_result *RemoveTenantMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveTenantMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RemoveCommand) {
		request.RemoveCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RemoveCommand, dara.String("RemoveCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RemoveCommandShrink) {
		body["RemoveCommand"] = request.RemoveCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTenantMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTenantMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a tenant member. Only superusers and system administrators can call this API operation.
//
// @param request - RemoveTenantMemberRequest
//
// @return RemoveTenantMemberResponse
func (client *Client) RemoveTenantMember(request *RemoveTenantMemberRequest) (_result *RemoveTenantMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveTenantMemberResponse{}
	_body, _err := client.RemoveTenantMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes members from a user group.
//
// @param tmpReq - RemoveUserGroupMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserGroupMemberResponse
func (client *Client) RemoveUserGroupMemberWithOptions(tmpReq *RemoveUserGroupMemberRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserGroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveUserGroupMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RemoveCommand) {
		request.RemoveCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RemoveCommand, dara.String("RemoveCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RemoveCommandShrink) {
		body["RemoveCommand"] = request.RemoveCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUserGroupMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUserGroupMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes members from a user group.
//
// @param request - RemoveUserGroupMemberRequest
//
// @return RemoveUserGroupMemberResponse
func (client *Client) RemoveUserGroupMember(request *RemoveUserGroupMemberRequest) (_result *RemoveUserGroupMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveUserGroupMemberResponse{}
	_body, _err := client.RemoveUserGroupMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the whitelist of a project.
//
// @param tmpReq - ReplaceProjectWhiteListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceProjectWhiteListsResponse
func (client *Client) ReplaceProjectWhiteListsWithOptions(tmpReq *ReplaceProjectWhiteListsRequest, runtime *dara.RuntimeOptions) (_result *ReplaceProjectWhiteListsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReplaceProjectWhiteListsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReplaceCommand) {
		request.ReplaceCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReplaceCommand, dara.String("ReplaceCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ReplaceCommandShrink) {
		body["ReplaceCommand"] = request.ReplaceCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplaceProjectWhiteLists"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplaceProjectWhiteListsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the whitelist of a project.
//
// @param request - ReplaceProjectWhiteListsRequest
//
// @return ReplaceProjectWhiteListsResponse
func (client *Client) ReplaceProjectWhiteLists(request *ReplaceProjectWhiteListsRequest) (_result *ReplaceProjectWhiteListsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReplaceProjectWhiteListsResponse{}
	_body, _err := client.ReplaceProjectWhiteListsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reset the Data Service application key. Only the application owner can perform this operation.
//
// Release version: v6.0.0.
//
// @param tmpReq - ResetDataServiceAppSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetDataServiceAppSecretResponse
func (client *Client) ResetDataServiceAppSecretWithOptions(tmpReq *ResetDataServiceAppSecretRequest, runtime *dara.RuntimeOptions) (_result *ResetDataServiceAppSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ResetDataServiceAppSecretShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetDataServiceAppSecret"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetDataServiceAppSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reset the Data Service application key. Only the application owner can perform this operation.
//
// Release version: v6.0.0.
//
// @param request - ResetDataServiceAppSecretRequest
//
// @return ResetDataServiceAppSecretResponse
func (client *Client) ResetDataServiceAppSecret(request *ResetDataServiceAppSecretRequest) (_result *ResetDataServiceAppSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetDataServiceAppSecretResponse{}
	_body, _err := client.ResetDataServiceAppSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resume physical node scheduling.
//
// @param tmpReq - ResumePhysicalNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumePhysicalNodeResponse
func (client *Client) ResumePhysicalNodeWithOptions(tmpReq *ResumePhysicalNodeRequest, runtime *dara.RuntimeOptions) (_result *ResumePhysicalNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ResumePhysicalNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResumeCommand) {
		request.ResumeCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResumeCommand, dara.String("ResumeCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ResumeCommandShrink) {
		body["ResumeCommand"] = request.ResumeCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumePhysicalNode"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumePhysicalNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resume physical node scheduling.
//
// @param request - ResumePhysicalNodeRequest
//
// @return ResumePhysicalNodeResponse
func (client *Client) ResumePhysicalNode(request *ResumePhysicalNodeRequest) (_result *ResumePhysicalNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumePhysicalNodeResponse{}
	_body, _err := client.ResumePhysicalNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retransfers a failed transfer task.
//
// @param tmpReq - RetryTransferOwnershipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryTransferOwnershipResponse
func (client *Client) RetryTransferOwnershipWithOptions(tmpReq *RetryTransferOwnershipRequest, runtime *dara.RuntimeOptions) (_result *RetryTransferOwnershipResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RetryTransferOwnershipShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PrivilegeTransferRecord) {
		request.PrivilegeTransferRecordShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PrivilegeTransferRecord, dara.String("PrivilegeTransferRecord"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PrivilegeTransferRecordShrink) {
		body["PrivilegeTransferRecord"] = request.PrivilegeTransferRecordShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryTransferOwnership"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryTransferOwnershipResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retransfers a failed transfer task.
//
// @param request - RetryTransferOwnershipRequest
//
// @return RetryTransferOwnershipResponse
func (client *Client) RetryTransferOwnership(request *RetryTransferOwnershipRequest) (_result *RetryTransferOwnershipResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RetryTransferOwnershipResponse{}
	_body, _err := client.RetryTransferOwnershipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes API authorization.
//
// @param tmpReq - RevokeDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeDataServiceApiResponse
func (client *Client) RevokeDataServiceApiWithOptions(tmpReq *RevokeDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *RevokeDataServiceApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RevokeDataServiceApiShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RevokeCommand) {
		request.RevokeCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RevokeCommand, dara.String("RevokeCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RevokeCommandShrink) {
		body["RevokeCommand"] = request.RevokeCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeDataServiceApi"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeDataServiceApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes API authorization.
//
// @param request - RevokeDataServiceApiRequest
//
// @return RevokeDataServiceApiResponse
func (client *Client) RevokeDataServiceApi(request *RevokeDataServiceApiRequest) (_result *RevokeDataServiceApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeDataServiceApiResponse{}
	_body, _err := client.RevokeDataServiceApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes resource authorization from a user.
//
// @param tmpReq - RevokeResourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeResourcePermissionResponse
func (client *Client) RevokeResourcePermissionWithOptions(tmpReq *RevokeResourcePermissionRequest, runtime *dara.RuntimeOptions) (_result *RevokeResourcePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RevokeResourcePermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RevokeCommand) {
		request.RevokeCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RevokeCommand, dara.String("RevokeCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RevokeCommandShrink) {
		body["RevokeCommand"] = request.RevokeCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeResourcePermission"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeResourcePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes resource authorization from a user.
//
// @param request - RevokeResourcePermissionRequest
//
// @return RevokeResourcePermissionResponse
func (client *Client) RevokeResourcePermission(request *RevokeResourcePermissionRequest) (_result *RevokeResourcePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeResourcePermissionResponse{}
	_body, _err := client.RevokeResourcePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops an ad hoc query task.
//
// @param request - StopAdHocTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAdHocTaskResponse
func (client *Client) StopAdHocTaskWithOptions(request *StopAdHocTaskRequest, runtime *dara.RuntimeOptions) (_result *StopAdHocTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopAdHocTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopAdHocTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an ad hoc query task.
//
// @param request - StopAdHocTaskRequest
//
// @return StopAdHocTaskResponse
func (client *Client) StopAdHocTask(request *StopAdHocTaskRequest) (_result *StopAdHocTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopAdHocTaskResponse{}
	_body, _err := client.StopAdHocTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a batch task.
//
// @param tmpReq - SubmitBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitBatchTaskResponse
func (client *Client) SubmitBatchTaskWithOptions(tmpReq *SubmitBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitBatchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitBatchTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SubmitCommand) {
		request.SubmitCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubmitCommand, dara.String("SubmitCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SubmitCommandShrink) {
		body["SubmitCommand"] = request.SubmitCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitBatchTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitBatchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a batch task.
//
// @param request - SubmitBatchTaskRequest
//
// @return SubmitBatchTaskResponse
func (client *Client) SubmitBatchTask(request *SubmitBatchTaskRequest) (_result *SubmitBatchTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitBatchTaskResponse{}
	_body, _err := client.SubmitBatchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch submit rule tasks with support for test runs.
//
// Release version: v5.4.2.
//
// @param tmpReq - SubmitQualityRuleTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitQualityRuleTasksResponse
func (client *Client) SubmitQualityRuleTasksWithOptions(tmpReq *SubmitQualityRuleTasksRequest, runtime *dara.RuntimeOptions) (_result *SubmitQualityRuleTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitQualityRuleTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SubmitCommand) {
		request.SubmitCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubmitCommand, dara.String("SubmitCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SubmitCommandShrink) {
		body["SubmitCommand"] = request.SubmitCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitQualityRuleTasks"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitQualityRuleTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch submit rule tasks with support for test runs.
//
// Release version: v5.4.2.
//
// @param request - SubmitQualityRuleTasksRequest
//
// @return SubmitQualityRuleTasksResponse
func (client *Client) SubmitQualityRuleTasks(request *SubmitQualityRuleTasksRequest) (_result *SubmitQualityRuleTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitQualityRuleTasksResponse{}
	_body, _err := client.SubmitQualityRuleTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits quality watchtask check tasks in batches.
//
// Online version: v5.4.2.
//
// @param tmpReq - SubmitQualityWatchTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitQualityWatchTasksResponse
func (client *Client) SubmitQualityWatchTasksWithOptions(tmpReq *SubmitQualityWatchTasksRequest, runtime *dara.RuntimeOptions) (_result *SubmitQualityWatchTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitQualityWatchTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SubmitCommand) {
		request.SubmitCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubmitCommand, dara.String("SubmitCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SubmitCommandShrink) {
		body["SubmitCommand"] = request.SubmitCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitQualityWatchTasks"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitQualityWatchTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits quality watchtask check tasks in batches.
//
// Online version: v5.4.2.
//
// @param request - SubmitQualityWatchTasksRequest
//
// @return SubmitQualityWatchTasksResponse
func (client *Client) SubmitQualityWatchTasks(request *SubmitQualityWatchTasksRequest) (_result *SubmitQualityWatchTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitQualityWatchTasksResponse{}
	_body, _err := client.SubmitQualityWatchTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Synchronizes department information.
//
// Description:
//
// Queries the details of a published API operation by AppKey.
//
// @param tmpReq - SyncDepartmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncDepartmentResponse
func (client *Client) SyncDepartmentWithOptions(tmpReq *SyncDepartmentRequest, runtime *dara.RuntimeOptions) (_result *SyncDepartmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SyncDepartmentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SyncDepartmentCommand) {
		request.SyncDepartmentCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SyncDepartmentCommand, dara.String("SyncDepartmentCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SyncDepartmentCommandShrink) {
		body["SyncDepartmentCommand"] = request.SyncDepartmentCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncDepartment"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncDepartmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes department information.
//
// Description:
//
// Queries the details of a published API operation by AppKey.
//
// @param request - SyncDepartmentRequest
//
// @return SyncDepartmentResponse
func (client *Client) SyncDepartment(request *SyncDepartmentRequest) (_result *SyncDepartmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SyncDepartmentResponse{}
	_body, _err := client.SyncDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Synchronizes department member information.
//
// @param tmpReq - SyncDepartmentUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncDepartmentUserResponse
func (client *Client) SyncDepartmentUserWithOptions(tmpReq *SyncDepartmentUserRequest, runtime *dara.RuntimeOptions) (_result *SyncDepartmentUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SyncDepartmentUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SyncDepartmentUserCommand) {
		request.SyncDepartmentUserCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SyncDepartmentUserCommand, dara.String("SyncDepartmentUserCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SyncDepartmentUserCommandShrink) {
		body["SyncDepartmentUserCommand"] = request.SyncDepartmentUserCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncDepartmentUser"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncDepartmentUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes department member information.
//
// @param request - SyncDepartmentUserRequest
//
// @return SyncDepartmentUserResponse
func (client *Client) SyncDepartmentUser(request *SyncDepartmentUserRequest) (_result *SyncDepartmentUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SyncDepartmentUserResponse{}
	_body, _err := client.SyncDepartmentUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Transfers ownership to a new owner in one click.
//
// @param tmpReq - TransferOwnershipForAllObjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferOwnershipForAllObjectResponse
func (client *Client) TransferOwnershipForAllObjectWithOptions(tmpReq *TransferOwnershipForAllObjectRequest, runtime *dara.RuntimeOptions) (_result *TransferOwnershipForAllObjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &TransferOwnershipForAllObjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PrivilegeTransferRecord) {
		request.PrivilegeTransferRecordShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PrivilegeTransferRecord, dara.String("PrivilegeTransferRecord"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PrivilegeTransferRecordShrink) {
		body["PrivilegeTransferRecord"] = request.PrivilegeTransferRecordShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferOwnershipForAllObject"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferOwnershipForAllObjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Transfers ownership to a new owner in one click.
//
// @param request - TransferOwnershipForAllObjectRequest
//
// @return TransferOwnershipForAllObjectResponse
func (client *Client) TransferOwnershipForAllObject(request *TransferOwnershipForAllObjectRequest) (_result *TransferOwnershipForAllObjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TransferOwnershipForAllObjectResponse{}
	_body, _err := client.TransferOwnershipForAllObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an ad hoc query file.
//
// @param tmpReq - UpdateAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAdHocFileResponse
func (client *Client) UpdateAdHocFileWithOptions(tmpReq *UpdateAdHocFileRequest, runtime *dara.RuntimeOptions) (_result *UpdateAdHocFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAdHocFileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAdHocFile"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAdHocFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an ad hoc query file.
//
// @param request - UpdateAdHocFileRequest
//
// @return UpdateAdHocFileResponse
func (client *Client) UpdateAdHocFile(request *UpdateAdHocFileRequest) (_result *UpdateAdHocFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAdHocFileResponse{}
	_body, _err := client.UpdateAdHocFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an offline compute node.
//
// @param tmpReq - UpdateBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBatchTaskResponse
func (client *Client) UpdateBatchTaskWithOptions(tmpReq *UpdateBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateBatchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateBatchTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBatchTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBatchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an offline compute node.
//
// @param request - UpdateBatchTaskRequest
//
// @return UpdateBatchTaskResponse
func (client *Client) UpdateBatchTask(request *UpdateBatchTaskRequest) (_result *UpdateBatchTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBatchTaskResponse{}
	_body, _err := client.UpdateBatchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Edits the custom data lineage of a batch task.
//
// @param tmpReq - UpdateBatchTaskUdfLineagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBatchTaskUdfLineagesResponse
func (client *Client) UpdateBatchTaskUdfLineagesWithOptions(tmpReq *UpdateBatchTaskUdfLineagesRequest, runtime *dara.RuntimeOptions) (_result *UpdateBatchTaskUdfLineagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateBatchTaskUdfLineagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBatchTaskUdfLineages"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBatchTaskUdfLineagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits the custom data lineage of a batch task.
//
// @param request - UpdateBatchTaskUdfLineagesRequest
//
// @return UpdateBatchTaskUdfLineagesResponse
func (client *Client) UpdateBatchTaskUdfLineages(request *UpdateBatchTaskUdfLineagesRequest) (_result *UpdateBatchTaskUdfLineagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBatchTaskUdfLineagesResponse{}
	_body, _err := client.UpdateBatchTaskUdfLineagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a business entity.
//
// @param tmpReq - UpdateBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBizEntityResponse
func (client *Client) UpdateBizEntityWithOptions(tmpReq *UpdateBizEntityRequest, runtime *dara.RuntimeOptions) (_result *UpdateBizEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateBizEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBizEntity"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBizEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a business entity.
//
// @param request - UpdateBizEntityRequest
//
// @return UpdateBizEntityResponse
func (client *Client) UpdateBizEntity(request *UpdateBizEntityRequest) (_result *UpdateBizEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBizEntityResponse{}
	_body, _err := client.UpdateBizEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update a business metric.
//
// Release version: v5.5.0.
//
// @param tmpReq - UpdateBizMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBizMetricResponse
func (client *Client) UpdateBizMetricWithOptions(tmpReq *UpdateBizMetricRequest, runtime *dara.RuntimeOptions) (_result *UpdateBizMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateBizMetricShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateBizMetricCommand) {
		request.UpdateBizMetricCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateBizMetricCommand, dara.String("UpdateBizMetricCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateBizMetricCommandShrink) {
		body["UpdateBizMetricCommand"] = request.UpdateBizMetricCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBizMetric"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBizMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update a business metric.
//
// Release version: v5.5.0.
//
// @param request - UpdateBizMetricRequest
//
// @return UpdateBizMetricResponse
func (client *Client) UpdateBizMetric(request *UpdateBizMetricRequest) (_result *UpdateBizMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBizMetricResponse{}
	_body, _err := client.UpdateBizMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a data domain.
//
// @param tmpReq - UpdateBizUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBizUnitResponse
func (client *Client) UpdateBizUnitWithOptions(tmpReq *UpdateBizUnitRequest, runtime *dara.RuntimeOptions) (_result *UpdateBizUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateBizUnitShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBizUnit"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBizUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data domain.
//
// @param request - UpdateBizUnitRequest
//
// @return UpdateBizUnitResponse
func (client *Client) UpdateBizUnit(request *UpdateBizUnitRequest) (_result *UpdateBizUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBizUnitResponse{}
	_body, _err := client.UpdateBizUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Edits a compute source. Business unit administrators and project administrators have permissions to perform this operation.
//
// @param tmpReq - UpdateComputeSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeSourceResponse
func (client *Client) UpdateComputeSourceWithOptions(tmpReq *UpdateComputeSourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateComputeSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateComputeSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComputeSource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComputeSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits a compute source. Business unit administrators and project administrators have permissions to perform this operation.
//
// @param request - UpdateComputeSourceRequest
//
// @return UpdateComputeSourceResponse
func (client *Client) UpdateComputeSource(request *UpdateComputeSourceRequest) (_result *UpdateComputeSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateComputeSourceResponse{}
	_body, _err := client.UpdateComputeSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a data domain.
//
// @param tmpReq - UpdateDataDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataDomainResponse
func (client *Client) UpdateDataDomainWithOptions(tmpReq *UpdateDataDomainRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataDomainShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataDomain"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data domain.
//
// @param request - UpdateDataDomainRequest
//
// @return UpdateDataDomainResponse
func (client *Client) UpdateDataDomain(request *UpdateDataDomainRequest) (_result *UpdateDataDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataDomainResponse{}
	_body, _err := client.UpdateDataDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a data service application. Only super administrators, system administrators, and application owners can perform this operation.
//
// Release version: v6.0.0.
//
// @param tmpReq - UpdateDataServiceAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataServiceAppResponse
func (client *Client) UpdateDataServiceAppWithOptions(tmpReq *UpdateDataServiceAppRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataServiceAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataServiceAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataServiceApp"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataServiceAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data service application. Only super administrators, system administrators, and application owners can perform this operation.
//
// Release version: v6.0.0.
//
// @param request - UpdateDataServiceAppRequest
//
// @return UpdateDataServiceAppResponse
func (client *Client) UpdateDataServiceApp(request *UpdateDataServiceAppRequest) (_result *UpdateDataServiceAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataServiceAppResponse{}
	_body, _err := client.UpdateDataServiceAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a data service application group. Only superusers and system administrators can perform this operation.
//
// Online version: v6.0.0.
//
// @param tmpReq - UpdateDataServiceAppGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataServiceAppGroupResponse
func (client *Client) UpdateDataServiceAppGroupWithOptions(tmpReq *UpdateDataServiceAppGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataServiceAppGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataServiceAppGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataServiceAppGroup"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataServiceAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data service application group. Only superusers and system administrators can perform this operation.
//
// Online version: v6.0.0.
//
// @param request - UpdateDataServiceAppGroupRequest
//
// @return UpdateDataServiceAppGroupResponse
func (client *Client) UpdateDataServiceAppGroup(request *UpdateDataServiceAppGroupRequest) (_result *UpdateDataServiceAppGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataServiceAppGroupResponse{}
	_body, _err := client.UpdateDataServiceAppGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the regular members of a data service application. Only the application owner can perform this operation.
//
// Online version: v6.0.0.
//
// @param tmpReq - UpdateDataServiceAppMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataServiceAppMemberResponse
func (client *Client) UpdateDataServiceAppMemberWithOptions(tmpReq *UpdateDataServiceAppMemberRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataServiceAppMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataServiceAppMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataServiceAppMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataServiceAppMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the regular members of a data service application. Only the application owner can perform this operation.
//
// Online version: v6.0.0.
//
// @param request - UpdateDataServiceAppMemberRequest
//
// @return UpdateDataServiceAppMemberResponse
func (client *Client) UpdateDataServiceAppMember(request *UpdateDataServiceAppMemberRequest) (_result *UpdateDataServiceAppMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataServiceAppMemberResponse{}
	_body, _err := client.UpdateDataServiceAppMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Edits the basic information of a data source. Tenant administrators, data administrators, business segment administrators, project administrators, and operations administrators have permissions to perform this operation.
//
// @param tmpReq - UpdateDataSourceBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceBasicInfoResponse
func (client *Client) UpdateDataSourceBasicInfoWithOptions(tmpReq *UpdateDataSourceBasicInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataSourceBasicInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataSourceBasicInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataSourceBasicInfo"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataSourceBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits the basic information of a data source. Tenant administrators, data administrators, business segment administrators, project administrators, and operations administrators have permissions to perform this operation.
//
// @param request - UpdateDataSourceBasicInfoRequest
//
// @return UpdateDataSourceBasicInfoResponse
func (client *Client) UpdateDataSourceBasicInfo(request *UpdateDataSourceBasicInfoRequest) (_result *UpdateDataSourceBasicInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataSourceBasicInfoResponse{}
	_body, _err := client.UpdateDataSourceBasicInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Edits the connection configuration items of a data source. Tenant administrators, data administrators, business unit administrators, project administrators, and operations administrators have permissions to perform this operation.
//
// @param tmpReq - UpdateDataSourceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceConfigResponse
func (client *Client) UpdateDataSourceConfigWithOptions(tmpReq *UpdateDataSourceConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataSourceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataSourceConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataSourceConfig"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataSourceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits the connection configuration items of a data source. Tenant administrators, data administrators, business unit administrators, project administrators, and operations administrators have permissions to perform this operation.
//
// @param request - UpdateDataSourceConfigRequest
//
// @return UpdateDataSourceConfigResponse
func (client *Client) UpdateDataSourceConfig(request *UpdateDataSourceConfigRequest) (_result *UpdateDataSourceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataSourceConfigResponse{}
	_body, _err := client.UpdateDataSourceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves the file position in the menu tree.
//
// @param request - UpdateFileDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileDirectoryResponse
func (client *Client) UpdateFileDirectoryWithOptions(request *UpdateFileDirectoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateFileDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Directory) {
		query["Directory"] = request.Directory
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFileDirectory"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFileDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves the file position in the menu tree.
//
// @param request - UpdateFileDirectoryRequest
//
// @return UpdateFileDirectoryResponse
func (client *Client) UpdateFileDirectory(request *UpdateFileDirectoryRequest) (_result *UpdateFileDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFileDirectoryResponse{}
	_body, _err := client.UpdateFileDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a file name.
//
// @param request - UpdateFileNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileNameResponse
func (client *Client) UpdateFileNameWithOptions(request *UpdateFileNameRequest, runtime *dara.RuntimeOptions) (_result *UpdateFileNameResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFileName"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFileNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a file name.
//
// @param request - UpdateFileNameRequest
//
// @return UpdateFileNameResponse
func (client *Client) UpdateFileName(request *UpdateFileNameRequest) (_result *UpdateFileNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFileNameResponse{}
	_body, _err := client.UpdateFileNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an integration pipeline or unstructured workflow node.
//
// @param tmpReq - UpdatePipelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePipelineResponse
func (client *Client) UpdatePipelineWithOptions(tmpReq *UpdatePipelineRequest, runtime *dara.RuntimeOptions) (_result *UpdatePipelineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePipelineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Context) {
		request.ContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Context, dara.String("Context"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextShrink) {
		body["Context"] = request.ContextShrink
	}

	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePipeline"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an integration pipeline or unstructured workflow node.
//
// @param request - UpdatePipelineRequest
//
// @return UpdatePipelineResponse
func (client *Client) UpdatePipeline(request *UpdatePipelineRequest) (_result *UpdatePipelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePipelineResponse{}
	_body, _err := client.UpdatePipelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously updates a pipeline or unstructured workflow node.
//
// @param tmpReq - UpdatePipelineByAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePipelineByAsyncResponse
func (client *Client) UpdatePipelineByAsyncWithOptions(tmpReq *UpdatePipelineByAsyncRequest, runtime *dara.RuntimeOptions) (_result *UpdatePipelineByAsyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePipelineByAsyncShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Context) {
		request.ContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Context, dara.String("Context"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextShrink) {
		body["Context"] = request.ContextShrink
	}

	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePipelineByAsync"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePipelineByAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously updates a pipeline or unstructured workflow node.
//
// @param request - UpdatePipelineByAsyncRequest
//
// @return UpdatePipelineByAsyncResponse
func (client *Client) UpdatePipelineByAsync(request *UpdatePipelineByAsyncRequest) (_result *UpdatePipelineByAsyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePipelineByAsyncResponse{}
	_body, _err := client.UpdatePipelineByAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Edits project members.
//
// @param tmpReq - UpdateProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectMemberResponse
func (client *Client) UpdateProjectMemberWithOptions(tmpReq *UpdateProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *UpdateProjectMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateProjectMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProjectMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProjectMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits project members.
//
// @param request - UpdateProjectMemberRequest
//
// @return UpdateProjectMemberResponse
func (client *Client) UpdateProjectMember(request *UpdateProjectMemberRequest) (_result *UpdateProjectMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateProjectMemberResponse{}
	_body, _err := client.UpdateProjectMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables quality rules in batches.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateQualityRuleSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQualityRuleSwitchResponse
func (client *Client) UpdateQualityRuleSwitchWithOptions(tmpReq *UpdateQualityRuleSwitchRequest, runtime *dara.RuntimeOptions) (_result *UpdateQualityRuleSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateQualityRuleSwitchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateQualityRuleSwitch"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQualityRuleSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables quality rules in batches.
//
// Release version: v5.4.2.
//
// @param request - UpdateQualityRuleSwitchRequest
//
// @return UpdateQualityRuleSwitchResponse
func (client *Client) UpdateQualityRuleSwitch(request *UpdateQualityRuleSwitchRequest) (_result *UpdateQualityRuleSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateQualityRuleSwitchResponse{}
	_body, _err := client.UpdateQualityRuleSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts or stops quality monitored objects in batches.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateQualityWatchSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQualityWatchSwitchResponse
func (client *Client) UpdateQualityWatchSwitchWithOptions(tmpReq *UpdateQualityWatchSwitchRequest, runtime *dara.RuntimeOptions) (_result *UpdateQualityWatchSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateQualityWatchSwitchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateQualityWatchSwitch"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQualityWatchSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts or stops quality monitored objects in batches.
//
// Release version: v5.4.2.
//
// @param request - UpdateQualityWatchSwitchRequest
//
// @return UpdateQualityWatchSwitchResponse
func (client *Client) UpdateQualityWatchSwitch(request *UpdateQualityWatchSwitchRequest) (_result *UpdateQualityWatchSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateQualityWatchSwitchResponse{}
	_body, _err := client.UpdateQualityWatchSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Edits a resource file.
//
// @param tmpReq - UpdateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceResponse
func (client *Client) UpdateResourceWithOptions(tmpReq *UpdateResourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits a resource file.
//
// @param request - UpdateResourceRequest
//
// @return UpdateResourceResponse
func (client *Client) UpdateResource(request *UpdateResourceRequest) (_result *UpdateResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateResourceResponse{}
	_body, _err := client.UpdateResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a row-level permission.
//
// @param tmpReq - UpdateRowPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRowPermissionResponse
func (client *Client) UpdateRowPermissionWithOptions(tmpReq *UpdateRowPermissionRequest, runtime *dara.RuntimeOptions) (_result *UpdateRowPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateRowPermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateRowPermissionCommand) {
		request.UpdateRowPermissionCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateRowPermissionCommand, dara.String("UpdateRowPermissionCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateRowPermissionCommandShrink) {
		body["UpdateRowPermissionCommand"] = request.UpdateRowPermissionCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRowPermission"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRowPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a row-level permission.
//
// @param request - UpdateRowPermissionRequest
//
// @return UpdateRowPermissionResponse
func (client *Client) UpdateRowPermission(request *UpdateRowPermissionRequest) (_result *UpdateRowPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRowPermissionResponse{}
	_body, _err := client.UpdateRowPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a data classification.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateSecurityClassifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecurityClassifyResponse
func (client *Client) UpdateSecurityClassifyWithOptions(tmpReq *UpdateSecurityClassifyRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecurityClassifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSecurityClassifyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSecurityClassify"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecurityClassifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data classification.
//
// Release version: v5.4.2.
//
// @param request - UpdateSecurityClassifyRequest
//
// @return UpdateSecurityClassifyResponse
func (client *Client) UpdateSecurityClassify(request *UpdateSecurityClassifyRequest) (_result *UpdateSecurityClassifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSecurityClassifyResponse{}
	_body, _err := client.UpdateSecurityClassifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a data classification folder. Release version: v5.4.2.
//
// @param tmpReq - UpdateSecurityClassifyCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecurityClassifyCatalogResponse
func (client *Client) UpdateSecurityClassifyCatalogWithOptions(tmpReq *UpdateSecurityClassifyCatalogRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecurityClassifyCatalogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSecurityClassifyCatalogShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSecurityClassifyCatalog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecurityClassifyCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data classification folder. Release version: v5.4.2.
//
// @param request - UpdateSecurityClassifyCatalogRequest
//
// @return UpdateSecurityClassifyCatalogResponse
func (client *Client) UpdateSecurityClassifyCatalog(request *UpdateSecurityClassifyCatalogRequest) (_result *UpdateSecurityClassifyCatalogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSecurityClassifyCatalogResponse{}
	_body, _err := client.UpdateSecurityClassifyCatalogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the effective status of security identification results.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateSecurityIdentifyResultStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecurityIdentifyResultStatusResponse
func (client *Client) UpdateSecurityIdentifyResultStatusWithOptions(tmpReq *UpdateSecurityIdentifyResultStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecurityIdentifyResultStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSecurityIdentifyResultStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSecurityIdentifyResultStatus"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecurityIdentifyResultStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the effective status of security identification results.
//
// Release version: v5.4.2.
//
// @param request - UpdateSecurityIdentifyResultStatusRequest
//
// @return UpdateSecurityIdentifyResultStatusResponse
func (client *Client) UpdateSecurityIdentifyResultStatus(request *UpdateSecurityIdentifyResultStatusRequest) (_result *UpdateSecurityIdentifyResultStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSecurityIdentifyResultStatusResponse{}
	_body, _err := client.UpdateSecurityIdentifyResultStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates data classification.
//
// Online version: v5.4.2.
//
// @param tmpReq - UpdateSecurityLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecurityLevelResponse
func (client *Client) UpdateSecurityLevelWithOptions(tmpReq *UpdateSecurityLevelRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecurityLevelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSecurityLevelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSecurityLevel"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecurityLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates data classification.
//
// Online version: v5.4.2.
//
// @param request - UpdateSecurityLevelRequest
//
// @return UpdateSecurityLevelResponse
func (client *Client) UpdateSecurityLevel(request *UpdateSecurityLevelRequest) (_result *UpdateSecurityLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSecurityLevelResponse{}
	_body, _err := client.UpdateSecurityLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a standard.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateStandardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStandardResponse
func (client *Client) UpdateStandardWithOptions(tmpReq *UpdateStandardRequest, runtime *dara.RuntimeOptions) (_result *UpdateStandardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStandardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStandard"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStandardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a standard.
//
// Release version: v5.4.2.
//
// @param request - UpdateStandardRequest
//
// @return UpdateStandardResponse
func (client *Client) UpdateStandard(request *UpdateStandardRequest) (_result *UpdateStandardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateStandardResponse{}
	_body, _err := client.UpdateStandardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a data standard lookup table.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateStandardLookupTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStandardLookupTableResponse
func (client *Client) UpdateStandardLookupTableWithOptions(tmpReq *UpdateStandardLookupTableRequest, runtime *dara.RuntimeOptions) (_result *UpdateStandardLookupTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStandardLookupTableShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStandardLookupTable"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStandardLookupTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data standard lookup table.
//
// Release version: v5.4.2.
//
// @param request - UpdateStandardLookupTableRequest
//
// @return UpdateStandardLookupTableResponse
func (client *Client) UpdateStandardLookupTable(request *UpdateStandardLookupTableRequest) (_result *UpdateStandardLookupTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateStandardLookupTableResponse{}
	_body, _err := client.UpdateStandardLookupTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the standard mapping relationship to invalid mapping.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateStandardMappingToInvalidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStandardMappingToInvalidResponse
func (client *Client) UpdateStandardMappingToInvalidWithOptions(tmpReq *UpdateStandardMappingToInvalidRequest, runtime *dara.RuntimeOptions) (_result *UpdateStandardMappingToInvalidResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStandardMappingToInvalidShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStandardMappingToInvalid"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStandardMappingToInvalidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the standard mapping relationship to invalid mapping.
//
// Release version: v5.4.2.
//
// @param request - UpdateStandardMappingToInvalidRequest
//
// @return UpdateStandardMappingToInvalidResponse
func (client *Client) UpdateStandardMappingToInvalid(request *UpdateStandardMappingToInvalidRequest) (_result *UpdateStandardMappingToInvalidResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateStandardMappingToInvalidResponse{}
	_body, _err := client.UpdateStandardMappingToInvalidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update standard set.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateStandardSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStandardSetResponse
func (client *Client) UpdateStandardSetWithOptions(tmpReq *UpdateStandardSetRequest, runtime *dara.RuntimeOptions) (_result *UpdateStandardSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStandardSetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStandardSet"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStandardSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update standard set.
//
// Release version: v5.4.2.
//
// @param request - UpdateStandardSetRequest
//
// @return UpdateStandardSetResponse
func (client *Client) UpdateStandardSet(request *UpdateStandardSetRequest) (_result *UpdateStandardSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateStandardSetResponse{}
	_body, _err := client.UpdateStandardSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a data standard template.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateStandardTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStandardTemplateResponse
func (client *Client) UpdateStandardTemplateWithOptions(tmpReq *UpdateStandardTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateStandardTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStandardTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStandardTemplate"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStandardTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data standard template.
//
// Release version: v5.4.2.
//
// @param request - UpdateStandardTemplateRequest
//
// @return UpdateStandardTemplateResponse
func (client *Client) UpdateStandardTemplate(request *UpdateStandardTemplateRequest) (_result *UpdateStandardTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateStandardTemplateResponse{}
	_body, _err := client.UpdateStandardTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a data standard root word.
//
// Online version: v5.4.2.
//
// @param tmpReq - UpdateStandardWordRootRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStandardWordRootResponse
func (client *Client) UpdateStandardWordRootWithOptions(tmpReq *UpdateStandardWordRootRequest, runtime *dara.RuntimeOptions) (_result *UpdateStandardWordRootResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStandardWordRootShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStandardWordRoot"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStandardWordRootResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data standard root word.
//
// Online version: v5.4.2.
//
// @param request - UpdateStandardWordRootRequest
//
// @return UpdateStandardWordRootResponse
func (client *Client) UpdateStandardWordRoot(request *UpdateStandardWordRootRequest) (_result *UpdateStandardWordRootResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateStandardWordRootResponse{}
	_body, _err := client.UpdateStandardWordRootWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the compute settings of a tenant.
//
// @param tmpReq - UpdateTenantComputeEngineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTenantComputeEngineResponse
func (client *Client) UpdateTenantComputeEngineWithOptions(tmpReq *UpdateTenantComputeEngineRequest, runtime *dara.RuntimeOptions) (_result *UpdateTenantComputeEngineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTenantComputeEngineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTenantComputeEngine"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTenantComputeEngineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the compute settings of a tenant.
//
// @param request - UpdateTenantComputeEngineRequest
//
// @return UpdateTenantComputeEngineResponse
func (client *Client) UpdateTenantComputeEngine(request *UpdateTenantComputeEngineRequest) (_result *UpdateTenantComputeEngineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTenantComputeEngineResponse{}
	_body, _err := client.UpdateTenantComputeEngineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Edits tenant members.
//
// @param tmpReq - UpdateTenantMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTenantMemberResponse
func (client *Client) UpdateTenantMemberWithOptions(tmpReq *UpdateTenantMemberRequest, runtime *dara.RuntimeOptions) (_result *UpdateTenantMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTenantMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTenantMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTenantMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits tenant members.
//
// @param request - UpdateTenantMemberRequest
//
// @return UpdateTenantMemberResponse
func (client *Client) UpdateTenantMember(request *UpdateTenantMemberRequest) (_result *UpdateTenantMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTenantMemberResponse{}
	_body, _err := client.UpdateTenantMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Edits a user-defined function.
//
// @param tmpReq - UpdateUdfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUdfResponse
func (client *Client) UpdateUdfWithOptions(tmpReq *UpdateUdfRequest, runtime *dara.RuntimeOptions) (_result *UpdateUdfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateUdfShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUdf"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUdfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits a user-defined function.
//
// @param request - UpdateUdfRequest
//
// @return UpdateUdfResponse
func (client *Client) UpdateUdf(request *UpdateUdfRequest) (_result *UpdateUdfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUdfResponse{}
	_body, _err := client.UpdateUdfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Edits a user group.
//
// @param tmpReq - UpdateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserGroupResponse
func (client *Client) UpdateUserGroupWithOptions(tmpReq *UpdateUserGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateUserGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserGroup"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits a user group.
//
// @param request - UpdateUserGroupRequest
//
// @return UpdateUserGroupResponse
func (client *Client) UpdateUserGroup(request *UpdateUserGroupRequest) (_result *UpdateUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.UpdateUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the enabled status of a user group.
//
// @param request - UpdateUserGroupSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserGroupSwitchResponse
func (client *Client) UpdateUserGroupSwitchWithOptions(request *UpdateUserGroupSwitchRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserGroupSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Active) {
		query["Active"] = request.Active
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserGroupSwitch"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserGroupSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the enabled status of a user group.
//
// @param request - UpdateUserGroupSwitchRequest
//
// @return UpdateUserGroupSwitchResponse
func (client *Client) UpdateUserGroupSwitch(request *UpdateUserGroupSwitchRequest) (_result *UpdateUserGroupSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserGroupSwitchResponse{}
	_body, _err := client.UpdateUserGroupSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create or modify a quality rule.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpsertQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertQualityRuleResponse
func (client *Client) UpsertQualityRuleWithOptions(tmpReq *UpsertQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *UpsertQualityRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpsertQualityRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpsertCommand) {
		request.UpsertCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpsertCommand, dara.String("UpsertCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpsertCommandShrink) {
		body["UpsertCommand"] = request.UpsertCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpsertQualityRule"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpsertQualityRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create or modify a quality rule.
//
// Release version: v5.4.2.
//
// @param request - UpsertQualityRuleRequest
//
// @return UpsertQualityRuleResponse
func (client *Client) UpsertQualityRule(request *UpsertQualityRuleRequest) (_result *UpsertQualityRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpsertQualityRuleResponse{}
	_body, _err := client.UpsertQualityRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or updates scheduling settings.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpsertQualityScheduleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertQualityScheduleResponse
func (client *Client) UpsertQualityScheduleWithOptions(tmpReq *UpsertQualityScheduleRequest, runtime *dara.RuntimeOptions) (_result *UpsertQualityScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpsertQualityScheduleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpsertCommand) {
		request.UpsertCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpsertCommand, dara.String("UpsertCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpsertCommandShrink) {
		body["UpsertCommand"] = request.UpsertCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpsertQualitySchedule"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpsertQualityScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates scheduling settings.
//
// Release version: v5.4.2.
//
// @param request - UpsertQualityScheduleRequest
//
// @return UpsertQualityScheduleResponse
func (client *Client) UpsertQualitySchedule(request *UpsertQualityScheduleRequest) (_result *UpsertQualityScheduleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpsertQualityScheduleResponse{}
	_body, _err := client.UpsertQualityScheduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or updates a quality template.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpsertQualityTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertQualityTemplateResponse
func (client *Client) UpsertQualityTemplateWithOptions(tmpReq *UpsertQualityTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpsertQualityTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpsertQualityTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpsertCommand) {
		request.UpsertCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpsertCommand, dara.String("UpsertCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpsertCommandShrink) {
		body["UpsertCommand"] = request.UpsertCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpsertQualityTemplate"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpsertQualityTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates a quality template.
//
// Release version: v5.4.2.
//
// @param request - UpsertQualityTemplateRequest
//
// @return UpsertQualityTemplateResponse
func (client *Client) UpsertQualityTemplate(request *UpsertQualityTemplateRequest) (_result *UpsertQualityTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpsertQualityTemplateResponse{}
	_body, _err := client.UpsertQualityTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or updates a quality monitored object. You can add multiple types of quality monitored objects, including Dataphin tables, global tables, data sources, metrics, and real-time meta tables.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpsertQualityWatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertQualityWatchResponse
func (client *Client) UpsertQualityWatchWithOptions(tmpReq *UpsertQualityWatchRequest, runtime *dara.RuntimeOptions) (_result *UpsertQualityWatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpsertQualityWatchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpsertCommand) {
		request.UpsertCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpsertCommand, dara.String("UpsertCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpsertCommandShrink) {
		body["UpsertCommand"] = request.UpsertCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpsertQualityWatch"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpsertQualityWatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates a quality monitored object. You can add multiple types of quality monitored objects, including Dataphin tables, global tables, data sources, metrics, and real-time meta tables.
//
// Release version: v5.4.2.
//
// @param request - UpsertQualityWatchRequest
//
// @return UpsertQualityWatchResponse
func (client *Client) UpsertQualityWatch(request *UpsertQualityWatchRequest) (_result *UpsertQualityWatchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpsertQualityWatchResponse{}
	_body, _err := client.UpsertQualityWatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or updates alert settings for a monitored object.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpsertQualityWatchAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertQualityWatchAlertResponse
func (client *Client) UpsertQualityWatchAlertWithOptions(tmpReq *UpsertQualityWatchAlertRequest, runtime *dara.RuntimeOptions) (_result *UpsertQualityWatchAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpsertQualityWatchAlertShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpsertCommand) {
		request.UpsertCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpsertCommand, dara.String("UpsertCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpsertCommandShrink) {
		body["UpsertCommand"] = request.UpsertCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpsertQualityWatchAlert"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpsertQualityWatchAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates alert settings for a monitored object.
//
// Release version: v5.4.2.
//
// @param request - UpsertQualityWatchAlertRequest
//
// @return UpsertQualityWatchAlertResponse
func (client *Client) UpsertQualityWatchAlert(request *UpsertQualityWatchAlertRequest) (_result *UpsertQualityWatchAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpsertQualityWatchAlertResponse{}
	_body, _err := client.UpsertQualityWatchAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
