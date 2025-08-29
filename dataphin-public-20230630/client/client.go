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
// 添加数据服务项目用户并设置角色。
//
// @param tmpReq - AddDataServiceProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDataServiceProjectMemberResponse
func (client *Client) AddDataServiceProjectMemberWithOptions(tmpReq *AddDataServiceProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *AddDataServiceProjectMemberResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 添加数据服务项目用户并设置角色。
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
// 增加项目成员。
//
// @param tmpReq - AddProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddProjectMemberResponse
func (client *Client) AddProjectMemberWithOptions(tmpReq *AddProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *AddProjectMemberResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 增加项目成员。
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
// 新增租户成员
//
// @param tmpReq - AddTenantMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTenantMembersResponse
func (client *Client) AddTenantMembersWithOptions(tmpReq *AddTenantMembersRequest, runtime *dara.RuntimeOptions) (_result *AddTenantMembersResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 新增租户成员
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
// 通过原始用户添加租户成员.
//
// @param tmpReq - AddTenantMembersBySourceUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTenantMembersBySourceUserResponse
func (client *Client) AddTenantMembersBySourceUserWithOptions(tmpReq *AddTenantMembersBySourceUserRequest, runtime *dara.RuntimeOptions) (_result *AddTenantMembersBySourceUserResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 通过原始用户添加租户成员.
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
// 添加用户组成员.
//
// @param tmpReq - AddUserGroupMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserGroupMemberResponse
func (client *Client) AddUserGroupMemberWithOptions(tmpReq *AddUserGroupMemberRequest, runtime *dara.RuntimeOptions) (_result *AddUserGroupMemberResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 添加用户组成员.
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
// 申请API权限。
//
// @param tmpReq - ApplyDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyDataServiceApiResponse
func (client *Client) ApplyDataServiceApiWithOptions(tmpReq *ApplyDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *ApplyDataServiceApiResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 申请API权限。
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
// 申请应用权限。
//
// @param tmpReq - ApplyDataServiceAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyDataServiceAppResponse
func (client *Client) ApplyDataServiceAppWithOptions(tmpReq *ApplyDataServiceAppRequest, runtime *dara.RuntimeOptions) (_result *ApplyDataServiceAppResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 申请应用权限。
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
// 项目计算源连通性检查。
//
// @param tmpReq - CheckComputeSourceConnectivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckComputeSourceConnectivityResponse
func (client *Client) CheckComputeSourceConnectivityWithOptions(tmpReq *CheckComputeSourceConnectivityRequest, runtime *dara.RuntimeOptions) (_result *CheckComputeSourceConnectivityResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 项目计算源连通性检查。
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
// 计算源连通性检查。
//
// @param request - CheckComputeSourceConnectivityByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckComputeSourceConnectivityByIdResponse
func (client *Client) CheckComputeSourceConnectivityByIdWithOptions(request *CheckComputeSourceConnectivityByIdRequest, runtime *dara.RuntimeOptions) (_result *CheckComputeSourceConnectivityByIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 计算源连通性检查。
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
// 检查数据源连通性
//
// @param tmpReq - CheckDataSourceConnectivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDataSourceConnectivityResponse
func (client *Client) CheckDataSourceConnectivityWithOptions(tmpReq *CheckDataSourceConnectivityRequest, runtime *dara.RuntimeOptions) (_result *CheckDataSourceConnectivityResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 检查数据源连通性
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
// 检查已创建的数据源是否正常连通
//
// @param request - CheckDataSourceConnectivityByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDataSourceConnectivityByIdResponse
func (client *Client) CheckDataSourceConnectivityByIdWithOptions(request *CheckDataSourceConnectivityByIdRequest, runtime *dara.RuntimeOptions) (_result *CheckDataSourceConnectivityByIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 检查已创建的数据源是否正常连通
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
// 检查项目是否存在依赖。
//
// @param request - CheckProjectHasDependencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckProjectHasDependencyResponse
func (client *Client) CheckProjectHasDependencyWithOptions(request *CheckProjectHasDependencyRequest, runtime *dara.RuntimeOptions) (_result *CheckProjectHasDependencyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 检查项目是否存在依赖。
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
// 校验用户是否有指定资源权限点.
//
// @param tmpReq - CheckResourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckResourcePermissionResponse
func (client *Client) CheckResourcePermissionWithOptions(tmpReq *CheckResourcePermissionRequest, runtime *dara.RuntimeOptions) (_result *CheckResourcePermissionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 校验用户是否有指定资源权限点.
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
// 创建即席查询文件
//
// @param tmpReq - CreateAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAdHocFileResponse
func (client *Client) CreateAdHocFileWithOptions(tmpReq *CreateAdHocFileRequest, runtime *dara.RuntimeOptions) (_result *CreateAdHocFileResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 创建即席查询文件
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
// 创建离线计算任务。
//
// @param tmpReq - CreateBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBatchTaskResponse
func (client *Client) CreateBatchTaskWithOptions(tmpReq *CreateBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateBatchTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 创建离线计算任务。
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
// 创建业务实体。
//
// @param tmpReq - CreateBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBizEntityResponse
func (client *Client) CreateBizEntityWithOptions(tmpReq *CreateBizEntityRequest, runtime *dara.RuntimeOptions) (_result *CreateBizEntityResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 创建业务实体。
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
// 创建数据板块。
//
// @param tmpReq - CreateBizUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBizUnitResponse
func (client *Client) CreateBizUnitWithOptions(tmpReq *CreateBizUnitRequest, runtime *dara.RuntimeOptions) (_result *CreateBizUnitResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 创建数据板块。
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
// 创建计算源。
//
// @param tmpReq - CreateComputeSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComputeSourceResponse
func (client *Client) CreateComputeSourceWithOptions(tmpReq *CreateComputeSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateComputeSourceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 创建计算源。
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
// 创建主题域。
//
// @param tmpReq - CreateDataDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataDomainResponse
func (client *Client) CreateDataDomainWithOptions(tmpReq *CreateDataDomainRequest, runtime *dara.RuntimeOptions) (_result *CreateDataDomainResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 创建主题域。
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
// 创建新的数据服务API并提交。
//
// @param tmpReq - CreateDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataServiceApiResponse
func (client *Client) CreateDataServiceApiWithOptions(tmpReq *CreateDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *CreateDataServiceApiResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 创建新的数据服务API并提交。
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
// 新建数据源
//
// @param tmpReq - CreateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSourceWithOptions(tmpReq *CreateDataSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateDataSourceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 新建数据源
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
// 创建菜单树文件目录
//
// @param tmpReq - CreateDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDirectoryResponse
func (client *Client) CreateDirectoryWithOptions(tmpReq *CreateDirectoryRequest, runtime *dara.RuntimeOptions) (_result *CreateDirectoryResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 创建菜单树文件目录
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
// 通用补数据接口 1.会生成补数据实例运行：影响相关产产出表数据 2.会进行任务运行：造成计算的费用以及存储的费用
//
// @param tmpReq - CreateNodeSupplementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeSupplementResponse
func (client *Client) CreateNodeSupplementWithOptions(tmpReq *CreateNodeSupplementRequest, runtime *dara.RuntimeOptions) (_result *CreateNodeSupplementResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 通用补数据接口 1.会生成补数据实例运行：影响相关产产出表数据 2.会进行任务运行：造成计算的费用以及存储的费用
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
// 创建数据集成任务。
//
// @param tmpReq - CreatePipelineNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineNodeResponse
func (client *Client) CreatePipelineNodeWithOptions(tmpReq *CreatePipelineNodeRequest, runtime *dara.RuntimeOptions) (_result *CreatePipelineNodeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 创建数据集成任务。
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
// 创建资源文件。
//
// @param tmpReq - CreateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceResponse
func (client *Client) CreateResourceWithOptions(tmpReq *CreateResourceRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 创建资源文件。
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
// 创建行级权限
//
// @param tmpReq - CreateRowPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRowPermissionResponse
func (client *Client) CreateRowPermissionWithOptions(tmpReq *CreateRowPermissionRequest, runtime *dara.RuntimeOptions) (_result *CreateRowPermissionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 创建行级权限
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
// 创建流批一体任务
//
// @param tmpReq - CreateStreamBatchJobMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStreamBatchJobMappingResponse
func (client *Client) CreateStreamBatchJobMappingWithOptions(tmpReq *CreateStreamBatchJobMappingRequest, runtime *dara.RuntimeOptions) (_result *CreateStreamBatchJobMappingResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 创建流批一体任务
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
// 创建自定义函数。
//
// @param tmpReq - CreateUdfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUdfResponse
func (client *Client) CreateUdfWithOptions(tmpReq *CreateUdfRequest, runtime *dara.RuntimeOptions) (_result *CreateUdfResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 创建自定义函数。
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
// 新建用户组.
//
// @param tmpReq - CreateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroupWithOptions(tmpReq *CreateUserGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateUserGroupResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 新建用户组.
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
// 删除菜单树即席查询文件
//
// @param request - DeleteAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAdHocFileResponse
func (client *Client) DeleteAdHocFileWithOptions(request *DeleteAdHocFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteAdHocFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除菜单树即席查询文件
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
// 删除离线计算任务，如果任务还没下线需要先下线再删除。
//
// @param tmpReq - DeleteBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBatchTaskResponse
func (client *Client) DeleteBatchTaskWithOptions(tmpReq *DeleteBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteBatchTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 删除离线计算任务，如果任务还没下线需要先下线再删除。
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
// 删除业务实体。
//
// @param request - DeleteBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBizEntityResponse
func (client *Client) DeleteBizEntityWithOptions(request *DeleteBizEntityRequest, runtime *dara.RuntimeOptions) (_result *DeleteBizEntityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除业务实体。
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
// 删除数据板块。
//
// @param request - DeleteBizUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBizUnitResponse
func (client *Client) DeleteBizUnitWithOptions(request *DeleteBizUnitRequest, runtime *dara.RuntimeOptions) (_result *DeleteBizUnitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除数据板块。
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
// 删除计算源。
//
// @param request - DeleteComputeSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComputeSourceResponse
func (client *Client) DeleteComputeSourceWithOptions(request *DeleteComputeSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteComputeSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除计算源。
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
// 删除主题域。
//
// @param request - DeleteDataDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataDomainResponse
func (client *Client) DeleteDataDomainWithOptions(request *DeleteDataDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除主题域。
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
// 删除数据源
//
// @param tmpReq - DeleteDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSourceWithOptions(tmpReq *DeleteDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 删除数据源
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
// 删除菜单树文件目录
//
// @param request - DeleteDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDirectoryResponse
func (client *Client) DeleteDirectoryWithOptions(request *DeleteDirectoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteDirectoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除菜单树文件目录
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
// 删除资源文件。
//
// @param request - DeleteResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceResponse
func (client *Client) DeleteResourceWithOptions(request *DeleteResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除资源文件。
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
// 删除行级权限
//
// @param tmpReq - DeleteRowPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRowPermissionResponse
func (client *Client) DeleteRowPermissionWithOptions(tmpReq *DeleteRowPermissionRequest, runtime *dara.RuntimeOptions) (_result *DeleteRowPermissionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 删除行级权限
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
// 删除自定义函数。
//
// @param request - DeleteUdfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUdfResponse
func (client *Client) DeleteUdfWithOptions(request *DeleteUdfRequest, runtime *dara.RuntimeOptions) (_result *DeleteUdfResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除自定义函数。
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
// 删除用户组.
//
// @param request - DeleteUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserGroupResponse
func (client *Client) DeleteUserGroupWithOptions(request *DeleteUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除用户组.
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
// 执行即席查询任务。
//
// @param tmpReq - ExecuteAdHocTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAdHocTaskResponse
func (client *Client) ExecuteAdHocTaskWithOptions(tmpReq *ExecuteAdHocTaskRequest, runtime *dara.RuntimeOptions) (_result *ExecuteAdHocTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 执行即席查询任务。
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
// 运行手动调度节点。
//
// @param tmpReq - ExecuteManualNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteManualNodeResponse
func (client *Client) ExecuteManualNodeWithOptions(tmpReq *ExecuteManualNodeRequest, runtime *dara.RuntimeOptions) (_result *ExecuteManualNodeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 运行手动调度节点。
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
// 重跑下游(修复链路数据), 支持强制重跑下游。影响范围: 1. 会产生计算成本；2. 会影响数据产出
//
// @param tmpReq - FixDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FixDataResponse
func (client *Client) FixDataWithOptions(tmpReq *FixDataRequest, runtime *dara.RuntimeOptions) (_result *FixDataResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 重跑下游(修复链路数据), 支持强制重跑下游。影响范围: 1. 会产生计算成本；2. 会影响数据产出
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
// 根据行级权限ID获取某一行级权限下的所有授权账号
//
// @param tmpReq - GetAccountByRowPermissionIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountByRowPermissionIdResponse
func (client *Client) GetAccountByRowPermissionIdWithOptions(tmpReq *GetAccountByRowPermissionIdRequest, runtime *dara.RuntimeOptions) (_result *GetAccountByRowPermissionIdResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 根据行级权限ID获取某一行级权限下的所有授权账号
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
// 查询即席查询文件。
//
// @param request - GetAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdHocFileResponse
func (client *Client) GetAdHocFileWithOptions(request *GetAdHocFileRequest, runtime *dara.RuntimeOptions) (_result *GetAdHocFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询即席查询文件。
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
// 获取即席查询任务运行日志。
//
// @param request - GetAdHocTaskLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdHocTaskLogResponse
func (client *Client) GetAdHocTaskLogWithOptions(request *GetAdHocTaskLogRequest, runtime *dara.RuntimeOptions) (_result *GetAdHocTaskLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取即席查询任务运行日志。
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
// 获取即席查询的任务运行结果。
//
// @param request - GetAdHocTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdHocTaskResultResponse
func (client *Client) GetAdHocTaskResultWithOptions(request *GetAdHocTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetAdHocTaskResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取即席查询的任务运行结果。
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
// 获取告警事件详情
//
// @param request - GetAlertEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlertEventResponse
func (client *Client) GetAlertEventWithOptions(request *GetAlertEventRequest, runtime *dara.RuntimeOptions) (_result *GetAlertEventResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取告警事件详情
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
// 获取离线计算任务详情。
//
// @param request - GetBatchTaskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTaskInfoResponse
func (client *Client) GetBatchTaskInfoWithOptions(request *GetBatchTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTaskInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取离线计算任务详情。
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
// 获取离线计算任务指定版本任务详情。
//
// @param request - GetBatchTaskInfoByVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTaskInfoByVersionResponse
func (client *Client) GetBatchTaskInfoByVersionWithOptions(request *GetBatchTaskInfoByVersionRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTaskInfoByVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取离线计算任务指定版本任务详情。
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
// 获取离线任务自定义血缘。
//
// @param request - GetBatchTaskUdfLineagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTaskUdfLineagesResponse
func (client *Client) GetBatchTaskUdfLineagesWithOptions(request *GetBatchTaskUdfLineagesRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTaskUdfLineagesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取离线任务自定义血缘。
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
// 获取离线计算任务版本列表。
//
// @param request - GetBatchTaskVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTaskVersionsResponse
func (client *Client) GetBatchTaskVersionsWithOptions(request *GetBatchTaskVersionsRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTaskVersionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取离线计算任务版本列表。
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
// 获取业务实体详情。
//
// @param request - GetBizEntityInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBizEntityInfoResponse
func (client *Client) GetBizEntityInfoWithOptions(request *GetBizEntityInfoRequest, runtime *dara.RuntimeOptions) (_result *GetBizEntityInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取业务实体详情。
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
// 查询指定版本的业务实体的详情。
//
// @param request - GetBizEntityInfoByVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBizEntityInfoByVersionResponse
func (client *Client) GetBizEntityInfoByVersionWithOptions(request *GetBizEntityInfoByVersionRequest, runtime *dara.RuntimeOptions) (_result *GetBizEntityInfoByVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询指定版本的业务实体的详情。
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
// 获取数据板块详情。
//
// @param request - GetBizUnitInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBizUnitInfoResponse
func (client *Client) GetBizUnitInfoWithOptions(request *GetBizUnitInfoRequest, runtime *dara.RuntimeOptions) (_result *GetBizUnitInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取数据板块详情。
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
// 根据环境获取集群信息
//
// @param request - GetClusterQueueInfoByEnvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterQueueInfoByEnvResponse
func (client *Client) GetClusterQueueInfoByEnvWithOptions(request *GetClusterQueueInfoByEnvRequest, runtime *dara.RuntimeOptions) (_result *GetClusterQueueInfoByEnvResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 根据环境获取集群信息
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
// 获取计算源详情。
//
// @param request - GetComputeSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComputeSourceResponse
func (client *Client) GetComputeSourceWithOptions(request *GetComputeSourceRequest, runtime *dara.RuntimeOptions) (_result *GetComputeSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取计算源详情。
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
// 获取主题域详情。
//
// @param request - GetDataDomainInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataDomainInfoResponse
func (client *Client) GetDataDomainInfoWithOptions(request *GetDataDomainInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDataDomainInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取主题域详情。
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
// 运维监控Api调用汇总统计。
//
// @param request - GetDataServiceApiCallSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiCallSummaryResponse
func (client *Client) GetDataServiceApiCallSummaryWithOptions(request *GetDataServiceApiCallSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiCallSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 运维监控Api调用汇总统计。
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
// 运维监控Api访问趋势分析。
//
// @param request - GetDataServiceApiCallTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiCallTrendResponse
func (client *Client) GetDataServiceApiCallTrendWithOptions(request *GetDataServiceApiCallTrendRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiCallTrendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 运维监控Api访问趋势分析。
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
// 获取API文档。
//
// @param request - GetDataServiceApiDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiDocumentResponse
func (client *Client) GetDataServiceApiDocumentWithOptions(request *GetDataServiceApiDocumentRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiDocumentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取API文档。
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
// 获取API异常影响汇总。
//
// @param request - GetDataServiceApiErrorImpactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiErrorImpactResponse
func (client *Client) GetDataServiceApiErrorImpactWithOptions(request *GetDataServiceApiErrorImpactRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiErrorImpactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取API异常影响汇总。
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
// 查询数据服务API分组列表。
//
// @param request - GetDataServiceApiGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiGroupsResponse
func (client *Client) GetDataServiceApiGroupsWithOptions(request *GetDataServiceApiGroupsRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询数据服务API分组列表。
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
// 获取应用有权限的用户列表。
//
// @param request - GetDataServiceAppAuthorizedUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAppAuthorizedUsersResponse
func (client *Client) GetDataServiceAppAuthorizedUsersWithOptions(request *GetDataServiceAppAuthorizedUsersRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAppAuthorizedUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取应用有权限的用户列表。
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
// 查询数据服务项目的应用分组列表。
//
// @param request - GetDataServiceAppGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAppGroupsResponse
func (client *Client) GetDataServiceAppGroupsWithOptions(request *GetDataServiceAppGroupsRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAppGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询数据服务项目的应用分组列表。
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
// 查询分组下应用列表。
//
// @param request - GetDataServiceAppsByGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAppsByGroupIdResponse
func (client *Client) GetDataServiceAppsByGroupIdWithOptions(request *GetDataServiceAppsByGroupIdRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAppsByGroupIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询分组下应用列表。
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
// 根据App分组Id查询账号有权限的应用列表。
//
// @param request - GetDataServiceAuthorizedAppsByGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAuthorizedAppsByGroupIdResponse
func (client *Client) GetDataServiceAuthorizedAppsByGroupIdWithOptions(request *GetDataServiceAuthorizedAppsByGroupIdRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAuthorizedAppsByGroupIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 根据App分组Id查询账号有权限的应用列表。
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
// 查询有权限的项目列表。
//
// @param request - GetDataServiceAuthorizedProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAuthorizedProjectsResponse
func (client *Client) GetDataServiceAuthorizedProjectsWithOptions(request *GetDataServiceAuthorizedProjectsRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAuthorizedProjectsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询有权限的项目列表。
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
// 当前登录当前用户作为负责人的项目列表。
//
// @param request - GetDataServiceMyProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceMyProjectsResponse
func (client *Client) GetDataServiceMyProjectsWithOptions(request *GetDataServiceMyProjectsRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceMyProjectsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 当前登录当前用户作为负责人的项目列表。
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
// 获取可添加到项目成员的用户列表。
//
// @param request - GetDataServiceProjectAddableUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceProjectAddableUsersResponse
func (client *Client) GetDataServiceProjectAddableUsersWithOptions(request *GetDataServiceProjectAddableUsersRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceProjectAddableUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取可添加到项目成员的用户列表。
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
// 查询数据源变更影响的集成任务及数据库SQL任务。
//
// @param request - GetDataSourceDependenciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataSourceDependenciesResponse
func (client *Client) GetDataSourceDependenciesWithOptions(request *GetDataSourceDependenciesRequest, runtime *dara.RuntimeOptions) (_result *GetDataSourceDependenciesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询数据源变更影响的集成任务及数据库SQL任务。
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
// 查询开发态对象上游依赖。
//
// @param request - GetDevObjectDependencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDevObjectDependencyResponse
func (client *Client) GetDevObjectDependencyWithOptions(request *GetDevObjectDependencyRequest, runtime *dara.RuntimeOptions) (_result *GetDevObjectDependencyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询开发态对象上游依赖。
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
// 获取文件夹目录树
//
// @param request - GetDirectoryTreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDirectoryTreeResponse
func (client *Client) GetDirectoryTreeWithOptions(request *GetDirectoryTreeRequest, runtime *dara.RuntimeOptions) (_result *GetDirectoryTreeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取文件夹目录树
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
// 获取文件存储临时读写授权。
//
// @param request - GetFileStorageCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileStorageCredentialResponse
func (client *Client) GetFileStorageCredentialWithOptions(request *GetFileStorageCredentialRequest, runtime *dara.RuntimeOptions) (_result *GetFileStorageCredentialResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取文件存储临时读写授权。
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
// 根据起始的实例查询该实例的下游
//
// @param tmpReq - GetInstanceDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceDownStreamResponse
func (client *Client) GetInstanceDownStreamWithOptions(tmpReq *GetInstanceDownStreamRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceDownStreamResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 根据起始的实例查询该实例的下游
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
// 查询实例的上下游，支持逻辑表和代码任务。
//
// @param tmpReq - GetInstanceUpDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceUpDownStreamResponse
func (client *Client) GetInstanceUpDownStreamWithOptions(tmpReq *GetInstanceUpDownStreamRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceUpDownStreamResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 查询实例的上下游，支持逻辑表和代码任务。
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
// 获取最新的待发布记录详情
//
// @param tmpReq - GetLatestSubmitDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLatestSubmitDetailResponse
func (client *Client) GetLatestSubmitDetailWithOptions(tmpReq *GetLatestSubmitDetailRequest, runtime *dara.RuntimeOptions) (_result *GetLatestSubmitDetailResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 获取最新的待发布记录详情
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
// 获取用户角色列表
//
// @param request - GetMyRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMyRolesResponse
func (client *Client) GetMyRolesWithOptions(request *GetMyRolesRequest, runtime *dara.RuntimeOptions) (_result *GetMyRolesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取用户角色列表
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
// 获取当前用户归属租户.
//
// @param tmpReq - GetMyTenantsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMyTenantsResponse
func (client *Client) GetMyTenantsWithOptions(tmpReq *GetMyTenantsRequest, runtime *dara.RuntimeOptions) (_result *GetMyTenantsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 获取当前用户归属租户.
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
// 通用查询节点上下游接口
//
// @param tmpReq - GetNodeUpDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeUpDownStreamResponse
func (client *Client) GetNodeUpDownStreamWithOptions(tmpReq *GetNodeUpDownStreamRequest, runtime *dara.RuntimeOptions) (_result *GetNodeUpDownStreamResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 通用查询节点上下游接口
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
// 查询补数据提交的状态
//
// @param request - GetOperationSubmitStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOperationSubmitStatusResponse
func (client *Client) GetOperationSubmitStatusWithOptions(request *GetOperationSubmitStatusRequest, runtime *dara.RuntimeOptions) (_result *GetOperationSubmitStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询补数据提交的状态
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
// 查询脚本的实例信息, 包括实例状态、运行时间等信息.
//
// @param request - GetPhysicalInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalInstanceResponse
func (client *Client) GetPhysicalInstanceWithOptions(request *GetPhysicalInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询脚本的实例信息, 包括实例状态、运行时间等信息.
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
// 获取实例执行的日志，如果实例重跑了多次，则会有多条日志
//
// @param request - GetPhysicalInstanceLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalInstanceLogResponse
func (client *Client) GetPhysicalInstanceLogWithOptions(request *GetPhysicalInstanceLogRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalInstanceLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取实例执行的日志，如果实例重跑了多次，则会有多条日志
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
// 查询物理调度节点。
//
// @param request - GetPhysicalNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeResponse
func (client *Client) GetPhysicalNodeWithOptions(request *GetPhysicalNodeRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalNodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询物理调度节点。
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
// 根据输出名查询对应的物理节点。
//
// @param request - GetPhysicalNodeByOutputNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeByOutputNameResponse
func (client *Client) GetPhysicalNodeByOutputNameWithOptions(request *GetPhysicalNodeByOutputNameRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalNodeByOutputNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 根据输出名查询对应的物理节点。
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
// 查询调度节点代码内容。
//
// @param request - GetPhysicalNodeContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeContentResponse
func (client *Client) GetPhysicalNodeContentWithOptions(request *GetPhysicalNodeContentRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalNodeContentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询调度节点代码内容。
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
// 查询节点的操作日志。
//
// @param request - GetPhysicalNodeOperationLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeOperationLogResponse
func (client *Client) GetPhysicalNodeOperationLogWithOptions(request *GetPhysicalNodeOperationLogRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalNodeOperationLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询节点的操作日志。
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
// 获取项目详情。
//
// @param request - GetProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithOptions(request *GetProjectRequest, runtime *dara.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取项目详情。
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
// 通过项目名获取项目详情。
//
// @param request - GetProjectByNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectByNameResponse
func (client *Client) GetProjectByNameWithOptions(request *GetProjectByNameRequest, runtime *dara.RuntimeOptions) (_result *GetProjectByNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 通过项目名获取项目详情。
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
// 获取项目生产账号
//
// @param request - GetProjectProduceUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectProduceUserResponse
func (client *Client) GetProjectProduceUserWithOptions(request *GetProjectProduceUserRequest, runtime *dara.RuntimeOptions) (_result *GetProjectProduceUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取项目生产账号
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
// 获取项目白名单。
//
// @param request - GetProjectWhiteListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectWhiteListsResponse
func (client *Client) GetProjectWhiteListsWithOptions(request *GetProjectWhiteListsRequest, runtime *dara.RuntimeOptions) (_result *GetProjectWhiteListsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取项目白名单。
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
// 根据集群ID获取集群版本
//
// @param request - GetQueueEngineVersionByEnvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueueEngineVersionByEnvResponse
func (client *Client) GetQueueEngineVersionByEnvWithOptions(request *GetQueueEngineVersionByEnvRequest, runtime *dara.RuntimeOptions) (_result *GetQueueEngineVersionByEnvResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 根据集群ID获取集群版本
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
// 获取资源文件详情。
//
// @param request - GetResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceResponse
func (client *Client) GetResourceWithOptions(request *GetResourceRequest, runtime *dara.RuntimeOptions) (_result *GetResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取资源文件详情。
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
// 获取资源文件指定版本详情。
//
// @param request - GetResourceByVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceByVersionResponse
func (client *Client) GetResourceByVersionWithOptions(request *GetResourceByVersionRequest, runtime *dara.RuntimeOptions) (_result *GetResourceByVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取资源文件指定版本详情。
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
// 获取计算源对应集群的spark客户信息
//
// @param request - GetSparkLocalClientInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkLocalClientInfoResponse
func (client *Client) GetSparkLocalClientInfoWithOptions(request *GetSparkLocalClientInfoRequest, runtime *dara.RuntimeOptions) (_result *GetSparkLocalClientInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取计算源对应集群的spark客户信息
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
// 获取dataphin实时研发任务集合
//
// @param request - GetStreamJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStreamJobsResponse
func (client *Client) GetStreamJobsWithOptions(request *GetStreamJobsRequest, runtime *dara.RuntimeOptions) (_result *GetStreamJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取dataphin实时研发任务集合
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
// 获取补数据工作流所有业务日期的Dagrun信息。
//
// @param request - GetSupplementDagrunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupplementDagrunResponse
func (client *Client) GetSupplementDagrunWithOptions(request *GetSupplementDagrunRequest, runtime *dara.RuntimeOptions) (_result *GetSupplementDagrunResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取补数据工作流所有业务日期的Dagrun信息。
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
// 列出补数据工作流下具体一个业务日期的所有节点的实例。
//
// @param request - GetSupplementDagrunInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupplementDagrunInstanceResponse
func (client *Client) GetSupplementDagrunInstanceWithOptions(request *GetSupplementDagrunInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetSupplementDagrunInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 列出补数据工作流下具体一个业务日期的所有节点的实例。
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
// 查询表字段血缘信息
//
// @param tmpReq - GetTableColumnLineageByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableColumnLineageByTaskIdResponse
func (client *Client) GetTableColumnLineageByTaskIdWithOptions(tmpReq *GetTableColumnLineageByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetTableColumnLineageByTaskIdResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 查询表字段血缘信息
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
// 查询表血缘信息
//
// @param tmpReq - GetTableLineageByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableLineageByTaskIdResponse
func (client *Client) GetTableLineageByTaskIdWithOptions(tmpReq *GetTableLineageByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetTableLineageByTaskIdResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 查询表血缘信息
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
// 根据转交任务ID查询转交任务的进度
//
// @param request - GetTransferInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTransferInfoResponse
func (client *Client) GetTransferInfoWithOptions(request *GetTransferInfoRequest, runtime *dara.RuntimeOptions) (_result *GetTransferInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 根据转交任务ID查询转交任务的进度
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
// 获取自定义函数详情。
//
// @param request - GetUdfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUdfResponse
func (client *Client) GetUdfWithOptions(request *GetUdfRequest, runtime *dara.RuntimeOptions) (_result *GetUdfResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取自定义函数详情。
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
// 查询自定义函数版本详情。
//
// @param request - GetUdfByVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUdfByVersionResponse
func (client *Client) GetUdfByVersionWithOptions(request *GetUdfByVersionRequest, runtime *dara.RuntimeOptions) (_result *GetUdfByVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询自定义函数版本详情。
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
// 通过用户原始Id（如阿里云Id）获取用户详情
//
// @param request - GetUserBySourceIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserBySourceIdResponse
func (client *Client) GetUserBySourceIdWithOptions(request *GetUserBySourceIdRequest, runtime *dara.RuntimeOptions) (_result *GetUserBySourceIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 通过用户原始Id（如阿里云Id）获取用户详情
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
// 获取用户组详情.
//
// @param request - GetUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroupWithOptions(request *GetUserGroupRequest, runtime *dara.RuntimeOptions) (_result *GetUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取用户组详情.
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
// 获取用户详情
//
// @param tmpReq - GetUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUsersResponse
func (client *Client) GetUsersWithOptions(tmpReq *GetUsersRequest, runtime *dara.RuntimeOptions) (_result *GetUsersResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 获取用户详情
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
// API授权。
//
// @param tmpReq - GrantDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantDataServiceApiResponse
func (client *Client) GrantDataServiceApiWithOptions(tmpReq *GrantDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *GrantDataServiceApiResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// API授权。
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
// 通过资源点对用户授权
//
// @param tmpReq - GrantResourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantResourcePermissionResponse
func (client *Client) GrantResourcePermissionWithOptions(tmpReq *GrantResourcePermissionRequest, runtime *dara.RuntimeOptions) (_result *GrantResourcePermissionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 通过资源点对用户授权
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
// 获取用户角色列表
//
// @param request - ListAddableRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddableRolesResponse
func (client *Client) ListAddableRolesWithOptions(request *ListAddableRolesRequest, runtime *dara.RuntimeOptions) (_result *ListAddableRolesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取用户角色列表
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
// 获取可加入租户成员列表的用户
//
// @param tmpReq - ListAddableUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddableUsersResponse
func (client *Client) ListAddableUsersWithOptions(tmpReq *ListAddableUsersRequest, runtime *dara.RuntimeOptions) (_result *ListAddableUsersResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 获取可加入租户成员列表的用户
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
// 根据条件查询多个告警事件
//
// @param tmpReq - ListAlertEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertEventsResponse
func (client *Client) ListAlertEventsWithOptions(tmpReq *ListAlertEventsRequest, runtime *dara.RuntimeOptions) (_result *ListAlertEventsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 根据条件查询多个告警事件
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
// 根据条件查询多个推送记录
//
// @param tmpReq - ListAlertNotificationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertNotificationsResponse
func (client *Client) ListAlertNotificationsWithOptions(tmpReq *ListAlertNotificationsRequest, runtime *dara.RuntimeOptions) (_result *ListAlertNotificationsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 根据条件查询多个推送记录
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
// 根据app查询api列表
//
// @param tmpReq - ListApiByAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiByAppResponse
func (client *Client) ListApiByAppWithOptions(tmpReq *ListApiByAppRequest, runtime *dara.RuntimeOptions) (_result *ListApiByAppResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 根据app查询api列表
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
// 查询应用已申请的API的具体的字段列表
//
// @param tmpReq - ListAuthorizedDataServiceApiDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizedDataServiceApiDetailsResponse
func (client *Client) ListAuthorizedDataServiceApiDetailsWithOptions(tmpReq *ListAuthorizedDataServiceApiDetailsRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizedDataServiceApiDetailsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 查询应用已申请的API的具体的字段列表
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
// 查询业务实体列表。
//
// @param tmpReq - ListBizEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBizEntitiesResponse
func (client *Client) ListBizEntitiesWithOptions(tmpReq *ListBizEntitiesRequest, runtime *dara.RuntimeOptions) (_result *ListBizEntitiesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 查询业务实体列表。
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
// 获取当前租户下的所有数据板块
//
// @param request - ListBizUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBizUnitsResponse
func (client *Client) ListBizUnitsWithOptions(request *ListBizUnitsRequest, runtime *dara.RuntimeOptions) (_result *ListBizUnitsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取当前租户下的所有数据板块
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
// 查询计算源列表。
//
// @param tmpReq - ListComputeSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeSourcesResponse
func (client *Client) ListComputeSourcesWithOptions(tmpReq *ListComputeSourcesRequest, runtime *dara.RuntimeOptions) (_result *ListComputeSourcesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 查询计算源列表。
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
// 获取主题域列表。
//
// @param tmpReq - ListDataDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataDomainsResponse
func (client *Client) ListDataDomainsWithOptions(tmpReq *ListDataDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListDataDomainsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 获取主题域列表。
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
// 查询API运维统计信息。
//
// @param tmpReq - ListDataServiceApiCallStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceApiCallStatisticsResponse
func (client *Client) ListDataServiceApiCallStatisticsWithOptions(tmpReq *ListDataServiceApiCallStatisticsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceApiCallStatisticsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 查询API运维统计信息。
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
// 分页查询数据服务调用日志。
//
// @param tmpReq - ListDataServiceApiCallsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceApiCallsResponse
func (client *Client) ListDataServiceApiCallsWithOptions(tmpReq *ListDataServiceApiCallsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceApiCallsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 分页查询数据服务调用日志。
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
// API影响分析列表。
//
// @param tmpReq - ListDataServiceApiImpactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceApiImpactsResponse
func (client *Client) ListDataServiceApiImpactsWithOptions(tmpReq *ListDataServiceApiImpactsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceApiImpactsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// API影响分析列表。
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
// 查询当前用户有权限的应用列表。
//
// @param tmpReq - ListDataServiceAuthorizedAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceAuthorizedAppsResponse
func (client *Client) ListDataServiceAuthorizedAppsWithOptions(tmpReq *ListDataServiceAuthorizedAppsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceAuthorizedAppsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 查询当前用户有权限的应用列表。
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
// 获取我管理的API权限列表。
//
// @param tmpReq - ListDataServiceMyApiPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceMyApiPermissionsResponse
func (client *Client) ListDataServiceMyApiPermissionsWithOptions(tmpReq *ListDataServiceMyApiPermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceMyApiPermissionsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 获取我管理的API权限列表。
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
// 查询当前用户有权限的应用。
//
// @param tmpReq - ListDataServiceMyAppPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceMyAppPermissionsResponse
func (client *Client) ListDataServiceMyAppPermissionsWithOptions(tmpReq *ListDataServiceMyAppPermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceMyAppPermissionsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 查询当前用户有权限的应用。
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
// 分页查询已发布的API列表。
//
// @param tmpReq - ListDataServicePublishedApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServicePublishedApisResponse
func (client *Client) ListDataServicePublishedApisWithOptions(tmpReq *ListDataServicePublishedApisRequest, runtime *dara.RuntimeOptions) (_result *ListDataServicePublishedApisResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 分页查询已发布的API列表。
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
// 搜索数据源，所属结果包含数据源配置项
//
// @param tmpReq - ListDataSourceWithConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceWithConfigResponse
func (client *Client) ListDataSourceWithConfigWithOptions(tmpReq *ListDataSourceWithConfigRequest, runtime *dara.RuntimeOptions) (_result *ListDataSourceWithConfigResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 搜索数据源，所属结果包含数据源配置项
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
// 遍历菜单树目录文件。
//
// @param tmpReq - ListFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFilesResponse
func (client *Client) ListFilesWithOptions(tmpReq *ListFilesRequest, runtime *dara.RuntimeOptions) (_result *ListFilesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 遍历菜单树目录文件。
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
// 分页查询实例。
//
// @param tmpReq - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(tmpReq *ListInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 分页查询实例。
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
// 查询节点下游，创建补数据工作流时可以作为数据参考
//
// @param tmpReq - ListNodeDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeDownStreamResponse
func (client *Client) ListNodeDownStreamWithOptions(tmpReq *ListNodeDownStreamRequest, runtime *dara.RuntimeOptions) (_result *ListNodeDownStreamResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 查询节点下游，创建补数据工作流时可以作为数据参考
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
// 查询调度节点列表。
//
// @param tmpReq - ListNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithOptions(tmpReq *ListNodesRequest, runtime *dara.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 查询调度节点列表。
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
// 获取项目成员列表。
//
// @param tmpReq - ListProjectMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectMembersResponse
func (client *Client) ListProjectMembersWithOptions(tmpReq *ListProjectMembersRequest, runtime *dara.RuntimeOptions) (_result *ListProjectMembersResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 获取项目成员列表。
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
// 获取项目列表。
//
// @param tmpReq - ListProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithOptions(tmpReq *ListProjectsRequest, runtime *dara.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 获取项目列表。
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
// 分页获取发布记录列表
//
// @param tmpReq - ListPublishRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublishRecordsResponse
func (client *Client) ListPublishRecordsWithOptions(tmpReq *ListPublishRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListPublishRecordsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 分页获取发布记录列表
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
// 分页获取权限操作列表
//
// @param tmpReq - ListResourcePermissionOperationLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcePermissionOperationLogResponse
func (client *Client) ListResourcePermissionOperationLogWithOptions(tmpReq *ListResourcePermissionOperationLogRequest, runtime *dara.RuntimeOptions) (_result *ListResourcePermissionOperationLogResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 分页获取权限操作列表
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
// 分页获取权限记录列表
//
// @param tmpReq - ListResourcePermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcePermissionsResponse
func (client *Client) ListResourcePermissionsWithOptions(tmpReq *ListResourcePermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListResourcePermissionsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 分页获取权限记录列表
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
// 分页查询行级权限
//
// @param tmpReq - ListRowPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRowPermissionResponse
func (client *Client) ListRowPermissionWithOptions(tmpReq *ListRowPermissionRequest, runtime *dara.RuntimeOptions) (_result *ListRowPermissionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 分页查询行级权限
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
// 分页查询指定用户行级权限
//
// @param tmpReq - ListRowPermissionByUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRowPermissionByUserIdResponse
func (client *Client) ListRowPermissionByUserIdWithOptions(tmpReq *ListRowPermissionByUserIdRequest, runtime *dara.RuntimeOptions) (_result *ListRowPermissionByUserIdResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 分页查询指定用户行级权限
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
// 分页获取待发布记录列表
//
// @param tmpReq - ListSubmitRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubmitRecordsResponse
func (client *Client) ListSubmitRecordsWithOptions(tmpReq *ListSubmitRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListSubmitRecordsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 分页获取待发布记录列表
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
// 查询租户成员列表
//
// @param tmpReq - ListTenantMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTenantMembersResponse
func (client *Client) ListTenantMembersWithOptions(tmpReq *ListTenantMembersRequest, runtime *dara.RuntimeOptions) (_result *ListTenantMembersResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 查询租户成员列表
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
// 用户组成员列表分页查询.
//
// @param tmpReq - ListUserGroupMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupMembersResponse
func (client *Client) ListUserGroupMembersWithOptions(tmpReq *ListUserGroupMembersRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupMembersResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 用户组成员列表分页查询.
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
// 用户组列表分页查询.
//
// @param tmpReq - ListUserGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroupsWithOptions(tmpReq *ListUserGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 用户组列表分页查询.
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
// 下线离线计算任务。
//
// @param request - OfflineBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineBatchTaskResponse
func (client *Client) OfflineBatchTaskWithOptions(request *OfflineBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *OfflineBatchTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 下线离线计算任务。
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
// 下线业务实体、
//
// @param tmpReq - OfflineBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineBizEntityResponse
func (client *Client) OfflineBizEntityWithOptions(tmpReq *OfflineBizEntityRequest, runtime *dara.RuntimeOptions) (_result *OfflineBizEntityResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 下线业务实体、
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
// 上线业务实体。
//
// @param tmpReq - OnlineBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineBizEntityResponse
func (client *Client) OnlineBizEntityWithOptions(tmpReq *OnlineBizEntityRequest, runtime *dara.RuntimeOptions) (_result *OnlineBizEntityResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 上线业务实体。
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
// 运维实例。
//
// @param tmpReq - OperateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateInstanceResponse
func (client *Client) OperateInstanceWithOptions(tmpReq *OperateInstanceRequest, runtime *dara.RuntimeOptions) (_result *OperateInstanceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 运维实例。
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
// 解析离线计算任务的逻辑表依赖，注意解析结果上游依赖信息中可能包含自依赖节点（上游节点ID和解析代码的任务节点ID相同）需要用户自己进行处理。
//
// @param tmpReq - ParseBatchTaskDependencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ParseBatchTaskDependencyResponse
func (client *Client) ParseBatchTaskDependencyWithOptions(tmpReq *ParseBatchTaskDependencyRequest, runtime *dara.RuntimeOptions) (_result *ParseBatchTaskDependencyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 解析离线计算任务的逻辑表依赖，注意解析结果上游依赖信息中可能包含自依赖节点（上游节点ID和解析代码的任务节点ID相同）需要用户自己进行处理。
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
// 暂停物理节点调度。
//
// @param tmpReq - PausePhysicalNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PausePhysicalNodeResponse
func (client *Client) PausePhysicalNodeWithOptions(tmpReq *PausePhysicalNodeRequest, runtime *dara.RuntimeOptions) (_result *PausePhysicalNodeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 暂停物理节点调度。
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
// 发布数据服务API到生产环境。
//
// @param request - PublishDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishDataServiceApiResponse
func (client *Client) PublishDataServiceApiWithOptions(request *PublishDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *PublishDataServiceApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 发布数据服务API到生产环境。
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
// 批量发布对象
//
// @param tmpReq - PublishObjectListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishObjectListResponse
func (client *Client) PublishObjectListWithOptions(tmpReq *PublishObjectListRequest, runtime *dara.RuntimeOptions) (_result *PublishObjectListResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 批量发布对象
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
// 删除项目成员。
//
// @param tmpReq - RemoveProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveProjectMemberResponse
func (client *Client) RemoveProjectMemberWithOptions(tmpReq *RemoveProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *RemoveProjectMemberResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 删除项目成员。
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
// 删除租户成员
//
// @param tmpReq - RemoveTenantMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTenantMemberResponse
func (client *Client) RemoveTenantMemberWithOptions(tmpReq *RemoveTenantMemberRequest, runtime *dara.RuntimeOptions) (_result *RemoveTenantMemberResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 删除租户成员
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
// 移除用户组成员.
//
// @param tmpReq - RemoveUserGroupMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserGroupMemberResponse
func (client *Client) RemoveUserGroupMemberWithOptions(tmpReq *RemoveUserGroupMemberRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserGroupMemberResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 移除用户组成员.
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
// 更新项目白名单。
//
// @param tmpReq - ReplaceProjectWhiteListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceProjectWhiteListsResponse
func (client *Client) ReplaceProjectWhiteListsWithOptions(tmpReq *ReplaceProjectWhiteListsRequest, runtime *dara.RuntimeOptions) (_result *ReplaceProjectWhiteListsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 更新项目白名单。
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
// 恢复物理节点调度。
//
// @param tmpReq - ResumePhysicalNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumePhysicalNodeResponse
func (client *Client) ResumePhysicalNodeWithOptions(tmpReq *ResumePhysicalNodeRequest, runtime *dara.RuntimeOptions) (_result *ResumePhysicalNodeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 恢复物理节点调度。
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
// 重新转交运行失败的转交任务
//
// @param tmpReq - RetryTransferOwnershipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryTransferOwnershipResponse
func (client *Client) RetryTransferOwnershipWithOptions(tmpReq *RetryTransferOwnershipRequest, runtime *dara.RuntimeOptions) (_result *RetryTransferOwnershipResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 重新转交运行失败的转交任务
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
// 回收API授权。
//
// @param tmpReq - RevokeDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeDataServiceApiResponse
func (client *Client) RevokeDataServiceApiWithOptions(tmpReq *RevokeDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *RevokeDataServiceApiResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 回收API授权。
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
// 回收用户资源授权
//
// @param tmpReq - RevokeResourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeResourcePermissionResponse
func (client *Client) RevokeResourcePermissionWithOptions(tmpReq *RevokeResourcePermissionRequest, runtime *dara.RuntimeOptions) (_result *RevokeResourcePermissionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 回收用户资源授权
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
// 终止即席查询任务。
//
// @param request - StopAdHocTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAdHocTaskResponse
func (client *Client) StopAdHocTaskWithOptions(request *StopAdHocTaskRequest, runtime *dara.RuntimeOptions) (_result *StopAdHocTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 终止即席查询任务。
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
// 提交离线计算任务。
//
// @param tmpReq - SubmitBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitBatchTaskResponse
func (client *Client) SubmitBatchTaskWithOptions(tmpReq *SubmitBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitBatchTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 提交离线计算任务。
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
// 一键转交负责人
//
// @param tmpReq - TransferOwnershipForAllObjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferOwnershipForAllObjectResponse
func (client *Client) TransferOwnershipForAllObjectWithOptions(tmpReq *TransferOwnershipForAllObjectRequest, runtime *dara.RuntimeOptions) (_result *TransferOwnershipForAllObjectResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 一键转交负责人
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
// 编辑即席查询文件。
//
// @param tmpReq - UpdateAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAdHocFileResponse
func (client *Client) UpdateAdHocFileWithOptions(tmpReq *UpdateAdHocFileRequest, runtime *dara.RuntimeOptions) (_result *UpdateAdHocFileResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 编辑即席查询文件。
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
// 编辑离线计算任务。
//
// @param tmpReq - UpdateBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBatchTaskResponse
func (client *Client) UpdateBatchTaskWithOptions(tmpReq *UpdateBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateBatchTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 编辑离线计算任务。
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
// 编辑离线计算任务自定义血缘。
//
// @param tmpReq - UpdateBatchTaskUdfLineagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBatchTaskUdfLineagesResponse
func (client *Client) UpdateBatchTaskUdfLineagesWithOptions(tmpReq *UpdateBatchTaskUdfLineagesRequest, runtime *dara.RuntimeOptions) (_result *UpdateBatchTaskUdfLineagesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 编辑离线计算任务自定义血缘。
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
// 更新业务实体、
//
// @param tmpReq - UpdateBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBizEntityResponse
func (client *Client) UpdateBizEntityWithOptions(tmpReq *UpdateBizEntityRequest, runtime *dara.RuntimeOptions) (_result *UpdateBizEntityResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 更新业务实体、
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
// 更新数据板块。
//
// @param tmpReq - UpdateBizUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBizUnitResponse
func (client *Client) UpdateBizUnitWithOptions(tmpReq *UpdateBizUnitRequest, runtime *dara.RuntimeOptions) (_result *UpdateBizUnitResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 更新数据板块。
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
// 修改计算源。
//
// @param tmpReq - UpdateComputeSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeSourceResponse
func (client *Client) UpdateComputeSourceWithOptions(tmpReq *UpdateComputeSourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateComputeSourceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 修改计算源。
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
// 更新主题域。
//
// @param tmpReq - UpdateDataDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataDomainResponse
func (client *Client) UpdateDataDomainWithOptions(tmpReq *UpdateDataDomainRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataDomainResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 更新主题域。
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
// 编辑数据源基本信息
//
// @param tmpReq - UpdateDataSourceBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceBasicInfoResponse
func (client *Client) UpdateDataSourceBasicInfoWithOptions(tmpReq *UpdateDataSourceBasicInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataSourceBasicInfoResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 编辑数据源基本信息
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
// 编辑数据源连接配置项
//
// @param tmpReq - UpdateDataSourceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceConfigResponse
func (client *Client) UpdateDataSourceConfigWithOptions(tmpReq *UpdateDataSourceConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataSourceConfigResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 编辑数据源连接配置项
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
// 修改菜单树文件所在目录
//
// @param request - UpdateFileDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileDirectoryResponse
func (client *Client) UpdateFileDirectoryWithOptions(request *UpdateFileDirectoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateFileDirectoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 修改菜单树文件所在目录
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
// 修改菜单树文件名称
//
// @param request - UpdateFileNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileNameResponse
func (client *Client) UpdateFileNameWithOptions(request *UpdateFileNameRequest, runtime *dara.RuntimeOptions) (_result *UpdateFileNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 修改菜单树文件名称
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
// 增加项目成员。
//
// @param tmpReq - UpdateProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectMemberResponse
func (client *Client) UpdateProjectMemberWithOptions(tmpReq *UpdateProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *UpdateProjectMemberResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 增加项目成员。
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
// 编辑资源文件。
//
// @param tmpReq - UpdateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceResponse
func (client *Client) UpdateResourceWithOptions(tmpReq *UpdateResourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 编辑资源文件。
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
// 更新行级权限
//
// @param tmpReq - UpdateRowPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRowPermissionResponse
func (client *Client) UpdateRowPermissionWithOptions(tmpReq *UpdateRowPermissionRequest, runtime *dara.RuntimeOptions) (_result *UpdateRowPermissionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 更新行级权限
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
// 修改租户计算设置。
//
// @param tmpReq - UpdateTenantComputeEngineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTenantComputeEngineResponse
func (client *Client) UpdateTenantComputeEngineWithOptions(tmpReq *UpdateTenantComputeEngineRequest, runtime *dara.RuntimeOptions) (_result *UpdateTenantComputeEngineResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 修改租户计算设置。
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
// 编辑租户成员
//
// @param tmpReq - UpdateTenantMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTenantMemberResponse
func (client *Client) UpdateTenantMemberWithOptions(tmpReq *UpdateTenantMemberRequest, runtime *dara.RuntimeOptions) (_result *UpdateTenantMemberResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 编辑租户成员
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
// 编辑自定义函数。
//
// @param tmpReq - UpdateUdfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUdfResponse
func (client *Client) UpdateUdfWithOptions(tmpReq *UpdateUdfRequest, runtime *dara.RuntimeOptions) (_result *UpdateUdfResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 编辑自定义函数。
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
// 编辑用户组.
//
// @param tmpReq - UpdateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserGroupResponse
func (client *Client) UpdateUserGroupWithOptions(tmpReq *UpdateUserGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserGroupResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 编辑用户组.
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
// 编辑用户组启用开关.
//
// @param request - UpdateUserGroupSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserGroupSwitchResponse
func (client *Client) UpdateUserGroupSwitchWithOptions(request *UpdateUserGroupSwitchRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserGroupSwitchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 编辑用户组启用开关.
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
