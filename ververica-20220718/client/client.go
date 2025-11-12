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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("ververica"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 执行定时计划
//
// @param headers - ApplyScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyScheduledPlanResponse
func (client *Client) ApplyScheduledPlanWithOptions(namespace *string, scheduledPlanId *string, headers *ApplyScheduledPlanHeaders, runtime *dara.RuntimeOptions) (_result *ApplyScheduledPlanResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyScheduledPlan"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/scheduled-plans/" + dara.PercentEncode(dara.StringValue(scheduledPlanId)) + "%3Aapply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyScheduledPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行定时计划
//
// @return ApplyScheduledPlanResponse
func (client *Client) ApplyScheduledPlan(namespace *string, scheduledPlanId *string) (_result *ApplyScheduledPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ApplyScheduledPlanHeaders{}
	_result = &ApplyScheduledPlanResponse{}
	_body, _err := client.ApplyScheduledPlanWithOptions(namespace, scheduledPlanId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消调试
//
// @param request - CancelSqlPreviewRequest
//
// @param headers - CancelSqlPreviewHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelSqlPreviewResponse
func (client *Client) CancelSqlPreviewWithOptions(namespace *string, request *CancelSqlPreviewRequest, headers *CancelSqlPreviewHeaders, runtime *dara.RuntimeOptions) (_result *CancelSqlPreviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QueryId) {
		query["queryId"] = request.QueryId
	}

	if !dara.IsNil(request.SessionClusterName) {
		query["sessionClusterName"] = request.SessionClusterName
	}

	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelSqlPreview"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sql-preview/cancel"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelSqlPreviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消调试
//
// @param request - CancelSqlPreviewRequest
//
// @return CancelSqlPreviewResponse
func (client *Client) CancelSqlPreview(namespace *string, request *CancelSqlPreviewRequest) (_result *CancelSqlPreviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CancelSqlPreviewHeaders{}
	_result = &CancelSqlPreviewResponse{}
	_body, _err := client.CancelSqlPreviewWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a deployment.
//
// @param request - CreateDeploymentRequest
//
// @param headers - CreateDeploymentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeploymentResponse
func (client *Client) CreateDeploymentWithOptions(namespace *string, request *CreateDeploymentRequest, headers *CreateDeploymentHeaders, runtime *dara.RuntimeOptions) (_result *CreateDeploymentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeployment"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a deployment.
//
// @param request - CreateDeploymentRequest
//
// @return CreateDeploymentResponse
func (client *Client) CreateDeployment(namespace *string, request *CreateDeploymentRequest) (_result *CreateDeploymentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateDeploymentHeaders{}
	_result = &CreateDeploymentResponse{}
	_body, _err := client.CreateDeploymentWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// create a deploymentDraft
//
// @param request - CreateDeploymentDraftRequest
//
// @param headers - CreateDeploymentDraftHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeploymentDraftResponse
func (client *Client) CreateDeploymentDraftWithOptions(namespace *string, request *CreateDeploymentDraftRequest, headers *CreateDeploymentDraftHeaders, runtime *dara.RuntimeOptions) (_result *CreateDeploymentDraftResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeploymentDraft"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-drafts"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeploymentDraftResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// create a deploymentDraft
//
// @param request - CreateDeploymentDraftRequest
//
// @return CreateDeploymentDraftResponse
func (client *Client) CreateDeploymentDraft(namespace *string, request *CreateDeploymentDraftRequest) (_result *CreateDeploymentDraftResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateDeploymentDraftHeaders{}
	_result = &CreateDeploymentDraftResponse{}
	_body, _err := client.CreateDeploymentDraftWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建deploymentTarget
//
// @param request - CreateDeploymentTargetRequest
//
// @param headers - CreateDeploymentTargetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeploymentTargetResponse
func (client *Client) CreateDeploymentTargetWithOptions(namespace *string, request *CreateDeploymentTargetRequest, headers *CreateDeploymentTargetHeaders, runtime *dara.RuntimeOptions) (_result *CreateDeploymentTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentTargetName) {
		query["deploymentTargetName"] = request.DeploymentTargetName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeploymentTarget"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-targets"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeploymentTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建deploymentTarget
//
// @param request - CreateDeploymentTargetRequest
//
// @return CreateDeploymentTargetResponse
func (client *Client) CreateDeploymentTarget(namespace *string, request *CreateDeploymentTargetRequest) (_result *CreateDeploymentTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateDeploymentTargetHeaders{}
	_result = &CreateDeploymentTargetResponse{}
	_body, _err := client.CreateDeploymentTargetWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建部署目标V2
//
// @param request - CreateDeploymentTargetV2Request
//
// @param headers - CreateDeploymentTargetV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeploymentTargetV2Response
func (client *Client) CreateDeploymentTargetV2WithOptions(namespace *string, request *CreateDeploymentTargetV2Request, headers *CreateDeploymentTargetV2Headers, runtime *dara.RuntimeOptions) (_result *CreateDeploymentTargetV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentTargetName) {
		query["deploymentTargetName"] = request.DeploymentTargetName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeploymentTargetV2"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-targets/support-elastic"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeploymentTargetV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建部署目标V2
//
// @param request - CreateDeploymentTargetV2Request
//
// @return CreateDeploymentTargetV2Response
func (client *Client) CreateDeploymentTargetV2(namespace *string, request *CreateDeploymentTargetV2Request) (_result *CreateDeploymentTargetV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateDeploymentTargetV2Headers{}
	_result = &CreateDeploymentTargetV2Response{}
	_body, _err := client.CreateDeploymentTargetV2WithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// create a folder
//
// @param request - CreateFolderRequest
//
// @param headers - CreateFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFolderResponse
func (client *Client) CreateFolderWithOptions(namespace *string, request *CreateFolderRequest, headers *CreateFolderHeaders, runtime *dara.RuntimeOptions) (_result *CreateFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFolder"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/folder"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// create a folder
//
// @param request - CreateFolderRequest
//
// @return CreateFolderResponse
func (client *Client) CreateFolder(namespace *string, request *CreateFolderRequest) (_result *CreateFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateFolderHeaders{}
	_result = &CreateFolderResponse{}
	_body, _err := client.CreateFolderWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a user to a namespace as a member and grants permissions to the user.
//
// @param request - CreateMemberRequest
//
// @param headers - CreateMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemberResponse
func (client *Client) CreateMemberWithOptions(namespace *string, request *CreateMemberRequest, headers *CreateMemberHeaders, runtime *dara.RuntimeOptions) (_result *CreateMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMember"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/gateway/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/members"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a user to a namespace as a member and grants permissions to the user.
//
// @param request - CreateMemberRequest
//
// @return CreateMemberResponse
func (client *Client) CreateMember(namespace *string, request *CreateMemberRequest) (_result *CreateMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateMemberHeaders{}
	_result = &CreateMemberResponse{}
	_body, _err := client.CreateMemberWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a savepoint.
//
// @param request - CreateSavepointRequest
//
// @param headers - CreateSavepointHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSavepointResponse
func (client *Client) CreateSavepointWithOptions(namespace *string, request *CreateSavepointRequest, headers *CreateSavepointHeaders, runtime *dara.RuntimeOptions) (_result *CreateSavepointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentId) {
		body["deploymentId"] = request.DeploymentId
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.NativeFormat) {
		body["nativeFormat"] = request.NativeFormat
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSavepoint"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/savepoints"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSavepointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a savepoint.
//
// @param request - CreateSavepointRequest
//
// @return CreateSavepointResponse
func (client *Client) CreateSavepoint(namespace *string, request *CreateSavepointRequest) (_result *CreateSavepointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateSavepointHeaders{}
	_result = &CreateSavepointResponse{}
	_body, _err := client.CreateSavepointWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建定时执行计划
//
// @param request - CreateScheduledPlanRequest
//
// @param headers - CreateScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduledPlanResponse
func (client *Client) CreateScheduledPlanWithOptions(namespace *string, request *CreateScheduledPlanRequest, headers *CreateScheduledPlanHeaders, runtime *dara.RuntimeOptions) (_result *CreateScheduledPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScheduledPlan"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/scheduled-plans"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScheduledPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建定时执行计划
//
// @param request - CreateScheduledPlanRequest
//
// @return CreateScheduledPlanResponse
func (client *Client) CreateScheduledPlan(namespace *string, request *CreateScheduledPlanRequest) (_result *CreateScheduledPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateScheduledPlanHeaders{}
	_result = &CreateScheduledPlanResponse{}
	_body, _err := client.CreateScheduledPlanWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建session集群
//
// @param request - CreateSessionClusterRequest
//
// @param headers - CreateSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSessionClusterResponse
func (client *Client) CreateSessionClusterWithOptions(namespace *string, request *CreateSessionClusterRequest, headers *CreateSessionClusterHeaders, runtime *dara.RuntimeOptions) (_result *CreateSessionClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSessionCluster"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sessionclusters"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSessionClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建session集群
//
// @param request - CreateSessionClusterRequest
//
// @return CreateSessionClusterResponse
func (client *Client) CreateSessionCluster(namespace *string, request *CreateSessionClusterRequest) (_result *CreateSessionClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateSessionClusterHeaders{}
	_result = &CreateSessionClusterResponse{}
	_body, _err := client.CreateSessionClusterWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Parses all user-defined function (UDF) methods in your JAR or Python file and creates an artifact configuration for a UDF.
//
// @param request - CreateUdfArtifactRequest
//
// @param headers - CreateUdfArtifactHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUdfArtifactResponse
func (client *Client) CreateUdfArtifactWithOptions(namespace *string, request *CreateUdfArtifactRequest, headers *CreateUdfArtifactHeaders, runtime *dara.RuntimeOptions) (_result *CreateUdfArtifactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUdfArtifact"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/udfartifacts"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUdfArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Parses all user-defined function (UDF) methods in your JAR or Python file and creates an artifact configuration for a UDF.
//
// @param request - CreateUdfArtifactRequest
//
// @return CreateUdfArtifactResponse
func (client *Client) CreateUdfArtifact(namespace *string, request *CreateUdfArtifactRequest) (_result *CreateUdfArtifactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateUdfArtifactHeaders{}
	_result = &CreateUdfArtifactResponse{}
	_body, _err := client.CreateUdfArtifactWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a variable.
//
// @param request - CreateVariableRequest
//
// @param headers - CreateVariableHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVariableResponse
func (client *Client) CreateVariableWithOptions(namespace *string, request *CreateVariableRequest, headers *CreateVariableHeaders, runtime *dara.RuntimeOptions) (_result *CreateVariableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVariable"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/variables"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a variable.
//
// @param request - CreateVariableRequest
//
// @return CreateVariableResponse
func (client *Client) CreateVariable(namespace *string, request *CreateVariableRequest) (_result *CreateVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateVariableHeaders{}
	_result = &CreateVariableResponse{}
	_body, _err := client.CreateVariableWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a registered custom connector from a workspace.
//
// @param headers - DeleteCustomConnectorHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomConnectorResponse
func (client *Client) DeleteCustomConnectorWithOptions(namespace *string, connectorName *string, headers *DeleteCustomConnectorHeaders, runtime *dara.RuntimeOptions) (_result *DeleteCustomConnectorResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomConnector"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/connectors/" + dara.PercentEncode(dara.StringValue(connectorName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a registered custom connector from a workspace.
//
// @return DeleteCustomConnectorResponse
func (client *Client) DeleteCustomConnector(namespace *string, connectorName *string) (_result *DeleteCustomConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteCustomConnectorHeaders{}
	_result = &DeleteCustomConnectorResponse{}
	_body, _err := client.DeleteCustomConnectorWithOptions(namespace, connectorName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a deployment based on the deployment ID.
//
// @param headers - DeleteDeploymentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeploymentResponse
func (client *Client) DeleteDeploymentWithOptions(namespace *string, deploymentId *string, headers *DeleteDeploymentHeaders, runtime *dara.RuntimeOptions) (_result *DeleteDeploymentResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDeployment"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments/" + dara.PercentEncode(dara.StringValue(deploymentId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a deployment based on the deployment ID.
//
// @return DeleteDeploymentResponse
func (client *Client) DeleteDeployment(namespace *string, deploymentId *string) (_result *DeleteDeploymentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteDeploymentHeaders{}
	_result = &DeleteDeploymentResponse{}
	_body, _err := client.DeleteDeploymentWithOptions(namespace, deploymentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// delete a deploymentDraft
//
// @param headers - DeleteDeploymentDraftHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeploymentDraftResponse
func (client *Client) DeleteDeploymentDraftWithOptions(namespace *string, deploymentDraftId *string, headers *DeleteDeploymentDraftHeaders, runtime *dara.RuntimeOptions) (_result *DeleteDeploymentDraftResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDeploymentDraft"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-drafts/" + dara.PercentEncode(dara.StringValue(deploymentDraftId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDeploymentDraftResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// delete a deploymentDraft
//
// @return DeleteDeploymentDraftResponse
func (client *Client) DeleteDeploymentDraft(namespace *string, deploymentDraftId *string) (_result *DeleteDeploymentDraftResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteDeploymentDraftHeaders{}
	_result = &DeleteDeploymentDraftResponse{}
	_body, _err := client.DeleteDeploymentDraftWithOptions(namespace, deploymentDraftId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除deploymentTarget
//
// @param headers - DeleteDeploymentTargetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeploymentTargetResponse
func (client *Client) DeleteDeploymentTargetWithOptions(namespace *string, deploymentTargetName *string, headers *DeleteDeploymentTargetHeaders, runtime *dara.RuntimeOptions) (_result *DeleteDeploymentTargetResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDeploymentTarget"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-targets/" + dara.PercentEncode(dara.StringValue(deploymentTargetName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDeploymentTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除deploymentTarget
//
// @return DeleteDeploymentTargetResponse
func (client *Client) DeleteDeploymentTarget(namespace *string, deploymentTargetName *string) (_result *DeleteDeploymentTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteDeploymentTargetHeaders{}
	_result = &DeleteDeploymentTargetResponse{}
	_body, _err := client.DeleteDeploymentTargetWithOptions(namespace, deploymentTargetName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// delete a folder
//
// @param headers - DeleteFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFolderResponse
func (client *Client) DeleteFolderWithOptions(namespace *string, folderId *string, headers *DeleteFolderHeaders, runtime *dara.RuntimeOptions) (_result *DeleteFolderResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFolder"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/folder/" + dara.PercentEncode(dara.StringValue(folderId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// delete a folder
//
// @return DeleteFolderResponse
func (client *Client) DeleteFolder(namespace *string, folderId *string) (_result *DeleteFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteFolderHeaders{}
	_result = &DeleteFolderResponse{}
	_body, _err := client.DeleteFolderWithOptions(namespace, folderId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the information about a job that is not in the running state in a deployment.
//
// @param headers - DeleteJobHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJobResponse
func (client *Client) DeleteJobWithOptions(namespace *string, jobId *string, headers *DeleteJobHeaders, runtime *dara.RuntimeOptions) (_result *DeleteJobResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteJob"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the information about a job that is not in the running state in a deployment.
//
// @return DeleteJobResponse
func (client *Client) DeleteJob(namespace *string, jobId *string) (_result *DeleteJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteJobHeaders{}
	_result = &DeleteJobResponse{}
	_body, _err := client.DeleteJobWithOptions(namespace, jobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the permissions from a member.
//
// @param headers - DeleteMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemberResponse
func (client *Client) DeleteMemberWithOptions(namespace *string, member *string, headers *DeleteMemberHeaders, runtime *dara.RuntimeOptions) (_result *DeleteMemberResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMember"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/gateway/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/members/" + dara.PercentEncode(dara.StringValue(member))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the permissions from a member.
//
// @return DeleteMemberResponse
func (client *Client) DeleteMember(namespace *string, member *string) (_result *DeleteMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteMemberHeaders{}
	_result = &DeleteMemberResponse{}
	_body, _err := client.DeleteMemberWithOptions(namespace, member, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a savepoint.
//
// @param headers - DeleteSavepointHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSavepointResponse
func (client *Client) DeleteSavepointWithOptions(namespace *string, savepointId *string, headers *DeleteSavepointHeaders, runtime *dara.RuntimeOptions) (_result *DeleteSavepointResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSavepoint"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/savepoints/" + dara.PercentEncode(dara.StringValue(savepointId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSavepointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a savepoint.
//
// @return DeleteSavepointResponse
func (client *Client) DeleteSavepoint(namespace *string, savepointId *string) (_result *DeleteSavepointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteSavepointHeaders{}
	_result = &DeleteSavepointResponse{}
	_body, _err := client.DeleteSavepointWithOptions(namespace, savepointId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除定时执行计划
//
// @param headers - DeleteScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScheduledPlanResponse
func (client *Client) DeleteScheduledPlanWithOptions(namespace *string, scheduledPlanId *string, headers *DeleteScheduledPlanHeaders, runtime *dara.RuntimeOptions) (_result *DeleteScheduledPlanResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScheduledPlan"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/scheduled-plans/" + dara.PercentEncode(dara.StringValue(scheduledPlanId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScheduledPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除定时执行计划
//
// @return DeleteScheduledPlanResponse
func (client *Client) DeleteScheduledPlan(namespace *string, scheduledPlanId *string) (_result *DeleteScheduledPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteScheduledPlanHeaders{}
	_result = &DeleteScheduledPlanResponse{}
	_body, _err := client.DeleteScheduledPlanWithOptions(namespace, scheduledPlanId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除session集群
//
// @param headers - DeleteSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSessionClusterResponse
func (client *Client) DeleteSessionClusterWithOptions(namespace *string, sessionClusterName *string, headers *DeleteSessionClusterHeaders, runtime *dara.RuntimeOptions) (_result *DeleteSessionClusterResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSessionCluster"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sessionclusters/" + dara.PercentEncode(dara.StringValue(sessionClusterName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSessionClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除session集群
//
// @return DeleteSessionClusterResponse
func (client *Client) DeleteSessionCluster(namespace *string, sessionClusterName *string) (_result *DeleteSessionClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteSessionClusterHeaders{}
	_result = &DeleteSessionClusterResponse{}
	_body, _err := client.DeleteSessionClusterWithOptions(namespace, sessionClusterName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除UdfArtifact
//
// @param headers - DeleteUdfArtifactHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUdfArtifactResponse
func (client *Client) DeleteUdfArtifactWithOptions(namespace *string, udfArtifactName *string, headers *DeleteUdfArtifactHeaders, runtime *dara.RuntimeOptions) (_result *DeleteUdfArtifactResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUdfArtifact"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/udfartifacts/" + dara.PercentEncode(dara.StringValue(udfArtifactName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUdfArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除UdfArtifact
//
// @return DeleteUdfArtifactResponse
func (client *Client) DeleteUdfArtifact(namespace *string, udfArtifactName *string) (_result *DeleteUdfArtifactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteUdfArtifactHeaders{}
	_result = &DeleteUdfArtifactResponse{}
	_body, _err := client.DeleteUdfArtifactWithOptions(namespace, udfArtifactName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an existing user-defined function (UDF) from a Realtime Compute for Apache Flink workspace.
//
// @param request - DeleteUdfFunctionRequest
//
// @param headers - DeleteUdfFunctionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUdfFunctionResponse
func (client *Client) DeleteUdfFunctionWithOptions(namespace *string, functionName *string, request *DeleteUdfFunctionRequest, headers *DeleteUdfFunctionHeaders, runtime *dara.RuntimeOptions) (_result *DeleteUdfFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClassName) {
		query["className"] = request.ClassName
	}

	if !dara.IsNil(request.UdfArtifactName) {
		query["udfArtifactName"] = request.UdfArtifactName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUdfFunction"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/udfartifacts/function/" + dara.PercentEncode(dara.StringValue(functionName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUdfFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an existing user-defined function (UDF) from a Realtime Compute for Apache Flink workspace.
//
// @param request - DeleteUdfFunctionRequest
//
// @return DeleteUdfFunctionResponse
func (client *Client) DeleteUdfFunction(namespace *string, functionName *string, request *DeleteUdfFunctionRequest) (_result *DeleteUdfFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteUdfFunctionHeaders{}
	_result = &DeleteUdfFunctionResponse{}
	_body, _err := client.DeleteUdfFunctionWithOptions(namespace, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a variable.
//
// @param headers - DeleteVariableHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVariableResponse
func (client *Client) DeleteVariableWithOptions(namespace *string, name *string, headers *DeleteVariableHeaders, runtime *dara.RuntimeOptions) (_result *DeleteVariableResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVariable"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/variables/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a variable.
//
// @return DeleteVariableResponse
func (client *Client) DeleteVariable(namespace *string, name *string) (_result *DeleteVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteVariableHeaders{}
	_result = &DeleteVariableResponse{}
	_body, _err := client.DeleteVariableWithOptions(namespace, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// deploy deploymentDraft async
//
// @param request - DeployDeploymentDraftAsyncRequest
//
// @param headers - DeployDeploymentDraftAsyncHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployDeploymentDraftAsyncResponse
func (client *Client) DeployDeploymentDraftAsyncWithOptions(namespace *string, request *DeployDeploymentDraftAsyncRequest, headers *DeployDeploymentDraftAsyncHeaders, runtime *dara.RuntimeOptions) (_result *DeployDeploymentDraftAsyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployDeploymentDraftAsync"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-drafts/async-deploy"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployDeploymentDraftAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// deploy deploymentDraft async
//
// @param request - DeployDeploymentDraftAsyncRequest
//
// @return DeployDeploymentDraftAsyncResponse
func (client *Client) DeployDeploymentDraftAsync(namespace *string, request *DeployDeploymentDraftAsyncRequest) (_result *DeployDeploymentDraftAsyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeployDeploymentDraftAsyncHeaders{}
	_result = &DeployDeploymentDraftAsyncResponse{}
	_body, _err := client.DeployDeploymentDraftAsyncWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 执行sql语句
//
// @param request - ExecuteSqlStatementRequest
//
// @param headers - ExecuteSqlStatementHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteSqlStatementResponse
func (client *Client) ExecuteSqlStatementWithOptions(namespace *string, request *ExecuteSqlStatementRequest, headers *ExecuteSqlStatementHeaders, runtime *dara.RuntimeOptions) (_result *ExecuteSqlStatementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteSqlStatement"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sql-statement/execute"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteSqlStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行sql语句
//
// @param request - ExecuteSqlStatementRequest
//
// @return ExecuteSqlStatementResponse
func (client *Client) ExecuteSqlStatement(namespace *string, request *ExecuteSqlStatementRequest) (_result *ExecuteSqlStatementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ExecuteSqlStatementHeaders{}
	_result = &ExecuteSqlStatementResponse{}
	_body, _err := client.ExecuteSqlStatementWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取调试结果
//
// @param request - FetchSqlPreviewResultsRequest
//
// @param headers - FetchSqlPreviewResultsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchSqlPreviewResultsResponse
func (client *Client) FetchSqlPreviewResultsWithOptions(namespace *string, request *FetchSqlPreviewResultsRequest, headers *FetchSqlPreviewResultsHeaders, runtime *dara.RuntimeOptions) (_result *FetchSqlPreviewResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QueryId) {
		query["queryId"] = request.QueryId
	}

	if !dara.IsNil(request.SessionClusterName) {
		query["sessionClusterName"] = request.SessionClusterName
	}

	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FetchSqlPreviewResults"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sql-preview/fetchResults"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FetchSqlPreviewResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取调试结果
//
// @param request - FetchSqlPreviewResultsRequest
//
// @return FetchSqlPreviewResultsResponse
func (client *Client) FetchSqlPreviewResults(namespace *string, request *FetchSqlPreviewResultsRequest) (_result *FetchSqlPreviewResultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &FetchSqlPreviewResultsHeaders{}
	_result = &FetchSqlPreviewResultsResponse{}
	_body, _err := client.FetchSqlPreviewResultsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Provides a Flink request proxy.
//
// @param request - FlinkApiProxyRequest
//
// @param headers - FlinkApiProxyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlinkApiProxyResponse
func (client *Client) FlinkApiProxyWithOptions(request *FlinkApiProxyRequest, headers *FlinkApiProxyHeaders, runtime *dara.RuntimeOptions) (_result *FlinkApiProxyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlinkApiPath) {
		query["flinkApiPath"] = request.FlinkApiPath
	}

	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ResourceId) {
		query["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlinkApiProxy"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/flink-ui/v2/proxy"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlinkApiProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides a Flink request proxy.
//
// @param request - FlinkApiProxyRequest
//
// @return FlinkApiProxyResponse
func (client *Client) FlinkApiProxy(request *FlinkApiProxyRequest) (_result *FlinkApiProxyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &FlinkApiProxyHeaders{}
	_result = &FlinkApiProxyResponse{}
	_body, _err := client.FlinkApiProxyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a ticket that applies for asynchronous generation of the fine-grained resources. This operation returns the ID of the ticket for you to query the asynchronous generation result.
//
// @param request - GenerateResourcePlanWithFlinkConfAsyncRequest
//
// @param headers - GenerateResourcePlanWithFlinkConfAsyncHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateResourcePlanWithFlinkConfAsyncResponse
func (client *Client) GenerateResourcePlanWithFlinkConfAsyncWithOptions(namespace *string, deploymentId *string, request *GenerateResourcePlanWithFlinkConfAsyncRequest, headers *GenerateResourcePlanWithFlinkConfAsyncHeaders, runtime *dara.RuntimeOptions) (_result *GenerateResourcePlanWithFlinkConfAsyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateResourcePlanWithFlinkConfAsync"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments/" + dara.PercentEncode(dara.StringValue(deploymentId)) + "/resource-plan%3AasyncGenerate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateResourcePlanWithFlinkConfAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a ticket that applies for asynchronous generation of the fine-grained resources. This operation returns the ID of the ticket for you to query the asynchronous generation result.
//
// @param request - GenerateResourcePlanWithFlinkConfAsyncRequest
//
// @return GenerateResourcePlanWithFlinkConfAsyncResponse
func (client *Client) GenerateResourcePlanWithFlinkConfAsync(namespace *string, deploymentId *string, request *GenerateResourcePlanWithFlinkConfAsyncRequest) (_result *GenerateResourcePlanWithFlinkConfAsyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GenerateResourcePlanWithFlinkConfAsyncHeaders{}
	_result = &GenerateResourcePlanWithFlinkConfAsyncResponse{}
	_body, _err := client.GenerateResourcePlanWithFlinkConfAsyncWithOptions(namespace, deploymentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取应用中的执行定时计划
//
// @param request - GetAppliedScheduledPlanRequest
//
// @param headers - GetAppliedScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppliedScheduledPlanResponse
func (client *Client) GetAppliedScheduledPlanWithOptions(namespace *string, request *GetAppliedScheduledPlanRequest, headers *GetAppliedScheduledPlanHeaders, runtime *dara.RuntimeOptions) (_result *GetAppliedScheduledPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentId) {
		query["deploymentId"] = request.DeploymentId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppliedScheduledPlan"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/scheduled-plans%3AgetExecutedScheduledPlan"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppliedScheduledPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用中的执行定时计划
//
// @param request - GetAppliedScheduledPlanRequest
//
// @return GetAppliedScheduledPlanResponse
func (client *Client) GetAppliedScheduledPlan(namespace *string, request *GetAppliedScheduledPlanRequest) (_result *GetAppliedScheduledPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetAppliedScheduledPlanHeaders{}
	_result = &GetAppliedScheduledPlanResponse{}
	_body, _err := client.GetAppliedScheduledPlanWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取catalog
//
// @param request - GetCatalogsRequest
//
// @param headers - GetCatalogsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCatalogsResponse
func (client *Client) GetCatalogsWithOptions(namespace *string, request *GetCatalogsRequest, headers *GetCatalogsHeaders, runtime *dara.RuntimeOptions) (_result *GetCatalogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["catalogName"] = request.CatalogName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCatalogs"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/catalogs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCatalogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取catalog
//
// @param request - GetCatalogsRequest
//
// @return GetCatalogsResponse
func (client *Client) GetCatalogs(namespace *string, request *GetCatalogsRequest) (_result *GetCatalogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetCatalogsHeaders{}
	_result = &GetCatalogsResponse{}
	_body, _err := client.GetCatalogsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取database
//
// @param request - GetDatabasesRequest
//
// @param headers - GetDatabasesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatabasesResponse
func (client *Client) GetDatabasesWithOptions(namespace *string, catalogName *string, request *GetDatabasesRequest, headers *GetDatabasesHeaders, runtime *dara.RuntimeOptions) (_result *GetDatabasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseName) {
		query["databaseName"] = request.DatabaseName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatabases"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/catalogs/" + dara.PercentEncode(dara.StringValue(catalogName)) + "/databases"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取database
//
// @param request - GetDatabasesRequest
//
// @return GetDatabasesResponse
func (client *Client) GetDatabases(namespace *string, catalogName *string, request *GetDatabasesRequest) (_result *GetDatabasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetDatabasesHeaders{}
	_result = &GetDatabasesResponse{}
	_body, _err := client.GetDatabasesWithOptions(namespace, catalogName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// get deploy deploymentDraft result
//
// @param headers - GetDeployDeploymentDraftResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeployDeploymentDraftResultResponse
func (client *Client) GetDeployDeploymentDraftResultWithOptions(namespace *string, ticketId *string, headers *GetDeployDeploymentDraftResultHeaders, runtime *dara.RuntimeOptions) (_result *GetDeployDeploymentDraftResultResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeployDeploymentDraftResult"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-drafts/tickets/" + dara.PercentEncode(dara.StringValue(ticketId)) + "/async-deploy"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeployDeploymentDraftResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// get deploy deploymentDraft result
//
// @return GetDeployDeploymentDraftResultResponse
func (client *Client) GetDeployDeploymentDraftResult(namespace *string, ticketId *string) (_result *GetDeployDeploymentDraftResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetDeployDeploymentDraftResultHeaders{}
	_result = &GetDeployDeploymentDraftResultResponse{}
	_body, _err := client.GetDeployDeploymentDraftResultWithOptions(namespace, ticketId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of a deployment.
//
// @param headers - GetDeploymentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeploymentResponse
func (client *Client) GetDeploymentWithOptions(namespace *string, deploymentId *string, headers *GetDeploymentHeaders, runtime *dara.RuntimeOptions) (_result *GetDeploymentResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeployment"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments/" + dara.PercentEncode(dara.StringValue(deploymentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a deployment.
//
// @return GetDeploymentResponse
func (client *Client) GetDeployment(namespace *string, deploymentId *string) (_result *GetDeploymentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetDeploymentHeaders{}
	_result = &GetDeploymentResponse{}
	_body, _err := client.GetDeploymentWithOptions(namespace, deploymentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// get a deploymentDraft
//
// @param headers - GetDeploymentDraftHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeploymentDraftResponse
func (client *Client) GetDeploymentDraftWithOptions(namespace *string, deploymentDraftId *string, headers *GetDeploymentDraftHeaders, runtime *dara.RuntimeOptions) (_result *GetDeploymentDraftResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeploymentDraft"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-drafts/" + dara.PercentEncode(dara.StringValue(deploymentDraftId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeploymentDraftResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// get a deploymentDraft
//
// @return GetDeploymentDraftResponse
func (client *Client) GetDeploymentDraft(namespace *string, deploymentDraftId *string) (_result *GetDeploymentDraftResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetDeploymentDraftHeaders{}
	_result = &GetDeploymentDraftResponse{}
	_body, _err := client.GetDeploymentDraftWithOptions(namespace, deploymentDraftId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// get deploymentDraft lock
//
// @param request - GetDeploymentDraftLockRequest
//
// @param headers - GetDeploymentDraftLockHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeploymentDraftLockResponse
func (client *Client) GetDeploymentDraftLockWithOptions(namespace *string, request *GetDeploymentDraftLockRequest, headers *GetDeploymentDraftLockHeaders, runtime *dara.RuntimeOptions) (_result *GetDeploymentDraftLockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentDraftId) {
		query["deploymentDraftId"] = request.DeploymentDraftId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeploymentDraftLock"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-drafts/getLock"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeploymentDraftLockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// get deploymentDraft lock
//
// @param request - GetDeploymentDraftLockRequest
//
// @return GetDeploymentDraftLockResponse
func (client *Client) GetDeploymentDraftLock(namespace *string, request *GetDeploymentDraftLockRequest) (_result *GetDeploymentDraftLockResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetDeploymentDraftLockHeaders{}
	_result = &GetDeploymentDraftLockResponse{}
	_body, _err := client.GetDeploymentDraftLockWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取运行事件
//
// @param request - GetEventsRequest
//
// @param headers - GetEventsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEventsResponse
func (client *Client) GetEventsWithOptions(namespace *string, request *GetEventsRequest, headers *GetEventsHeaders, runtime *dara.RuntimeOptions) (_result *GetEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentId) {
		query["deploymentId"] = request.DeploymentId
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEvents"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/events"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取运行事件
//
// @param request - GetEventsRequest
//
// @return GetEventsResponse
func (client *Client) GetEvents(namespace *string, request *GetEventsRequest) (_result *GetEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetEventsHeaders{}
	_result = &GetEventsResponse{}
	_body, _err := client.GetEventsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// get a folder
//
// @param request - GetFolderRequest
//
// @param headers - GetFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFolderResponse
func (client *Client) GetFolderWithOptions(namespace *string, request *GetFolderRequest, headers *GetFolderHeaders, runtime *dara.RuntimeOptions) (_result *GetFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		query["folderId"] = request.FolderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFolder"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/folder"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// get a folder
//
// @param request - GetFolderRequest
//
// @return GetFolderResponse
func (client *Client) GetFolder(namespace *string, request *GetFolderRequest) (_result *GetFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetFolderHeaders{}
	_result = &GetFolderResponse{}
	_body, _err := client.GetFolderWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the asynchronous generation result of fine-grained resources based on the ID of the ticket that applies for an asynchronous generation.
//
// @param headers - GetGenerateResourcePlanResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGenerateResourcePlanResultResponse
func (client *Client) GetGenerateResourcePlanResultWithOptions(namespace *string, ticketId *string, headers *GetGenerateResourcePlanResultHeaders, runtime *dara.RuntimeOptions) (_result *GetGenerateResourcePlanResultResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGenerateResourcePlanResult"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments/tickets/" + dara.PercentEncode(dara.StringValue(ticketId)) + "/resource-plan%3AasyncGenerate"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGenerateResourcePlanResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the asynchronous generation result of fine-grained resources based on the ID of the ticket that applies for an asynchronous generation.
//
// @return GetGenerateResourcePlanResultResponse
func (client *Client) GetGenerateResourcePlanResult(namespace *string, ticketId *string) (_result *GetGenerateResourcePlanResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetGenerateResourcePlanResultHeaders{}
	_result = &GetGenerateResourcePlanResultResponse{}
	_body, _err := client.GetGenerateResourcePlanResultWithOptions(namespace, ticketId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询动态更新结果
//
// @param headers - GetHotUpdateJobResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotUpdateJobResultResponse
func (client *Client) GetHotUpdateJobResultWithOptions(namespace *string, jobHotUpdateId *string, headers *GetHotUpdateJobResultHeaders, runtime *dara.RuntimeOptions) (_result *GetHotUpdateJobResultResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotUpdateJobResult"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/jobs/hot-updates/" + dara.PercentEncode(dara.StringValue(jobHotUpdateId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotUpdateJobResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询动态更新结果
//
// @return GetHotUpdateJobResultResponse
func (client *Client) GetHotUpdateJobResult(namespace *string, jobHotUpdateId *string) (_result *GetHotUpdateJobResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetHotUpdateJobResultHeaders{}
	_result = &GetHotUpdateJobResultResponse{}
	_body, _err := client.GetHotUpdateJobResultWithOptions(namespace, jobHotUpdateId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of a job.
//
// @param headers - GetJobHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
func (client *Client) GetJobWithOptions(namespace *string, jobId *string, headers *GetJobHeaders, runtime *dara.RuntimeOptions) (_result *GetJobResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJob"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a job.
//
// @return GetJobResponse
func (client *Client) GetJob(namespace *string, jobId *string) (_result *GetJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetJobHeaders{}
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(namespace, jobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取作业诊断信息
//
// @param headers - GetJobDiagnosisHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobDiagnosisResponse
func (client *Client) GetJobDiagnosisWithOptions(namespace *string, deploymentId *string, jobId *string, headers *GetJobDiagnosisHeaders, runtime *dara.RuntimeOptions) (_result *GetJobDiagnosisResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobDiagnosis"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments/" + dara.PercentEncode(dara.StringValue(deploymentId)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId)) + "/job-diagnoses/lite"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取作业诊断信息
//
// @return GetJobDiagnosisResponse
func (client *Client) GetJobDiagnosis(namespace *string, deploymentId *string, jobId *string) (_result *GetJobDiagnosisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetJobDiagnosisHeaders{}
	_result = &GetJobDiagnosisResponse{}
	_body, _err := client.GetJobDiagnosisWithOptions(namespace, deploymentId, jobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the latest startup logs of a job.
//
// @param headers - GetLatestJobStartLogHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLatestJobStartLogResponse
func (client *Client) GetLatestJobStartLogWithOptions(namespace *string, deploymentId *string, headers *GetLatestJobStartLogHeaders, runtime *dara.RuntimeOptions) (_result *GetLatestJobStartLogResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLatestJobStartLog"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments/" + dara.PercentEncode(dara.StringValue(deploymentId)) + "/latest_jobmanager_start_log"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLatestJobStartLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the latest startup logs of a job.
//
// @return GetLatestJobStartLogResponse
func (client *Client) GetLatestJobStartLog(namespace *string, deploymentId *string) (_result *GetLatestJobStartLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetLatestJobStartLogHeaders{}
	_result = &GetLatestJobStartLogResponse{}
	_body, _err := client.GetLatestJobStartLogWithOptions(namespace, deploymentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the lineage information of a deployment.
//
// @param request - GetLineageInfoRequest
//
// @param headers - GetLineageInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLineageInfoResponse
func (client *Client) GetLineageInfoWithOptions(request *GetLineageInfoRequest, headers *GetLineageInfoHeaders, runtime *dara.RuntimeOptions) (_result *GetLineageInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLineageInfo"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/meta/v2/lineage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLineageInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the lineage information of a deployment.
//
// @param request - GetLineageInfoRequest
//
// @return GetLineageInfoResponse
func (client *Client) GetLineageInfo(request *GetLineageInfoRequest) (_result *GetLineageInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetLineageInfoHeaders{}
	_result = &GetLineageInfoResponse{}
	_body, _err := client.GetLineageInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the permissions of a member.
//
// @param headers - GetMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemberResponse
func (client *Client) GetMemberWithOptions(namespace *string, member *string, headers *GetMemberHeaders, runtime *dara.RuntimeOptions) (_result *GetMemberResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMember"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/gateway/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/members/" + dara.PercentEncode(dara.StringValue(member))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the permissions of a member.
//
// @return GetMemberResponse
func (client *Client) GetMember(namespace *string, member *string) (_result *GetMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetMemberHeaders{}
	_result = &GetMemberResponse{}
	_body, _err := client.GetMemberWithOptions(namespace, member, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取上传文件URL
//
// @param request - GetPreSignedUrlForPutObjectRequest
//
// @param headers - GetPreSignedUrlForPutObjectHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPreSignedUrlForPutObjectResponse
func (client *Client) GetPreSignedUrlForPutObjectWithOptions(namespace *string, request *GetPreSignedUrlForPutObjectRequest, headers *GetPreSignedUrlForPutObjectHeaders, runtime *dara.RuntimeOptions) (_result *GetPreSignedUrlForPutObjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["fileName"] = request.FileName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPreSignedUrlForPutObject"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/artifacts/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/getPreSignedUrlForPutObject"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPreSignedUrlForPutObjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取上传文件URL
//
// @param request - GetPreSignedUrlForPutObjectRequest
//
// @return GetPreSignedUrlForPutObjectResponse
func (client *Client) GetPreSignedUrlForPutObject(namespace *string, request *GetPreSignedUrlForPutObjectRequest) (_result *GetPreSignedUrlForPutObjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetPreSignedUrlForPutObjectHeaders{}
	_result = &GetPreSignedUrlForPutObjectResponse{}
	_body, _err := client.GetPreSignedUrlForPutObjectWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details of a savepoint and checkpoint.
//
// @param headers - GetSavepointHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSavepointResponse
func (client *Client) GetSavepointWithOptions(namespace *string, savepointId *string, headers *GetSavepointHeaders, runtime *dara.RuntimeOptions) (_result *GetSavepointResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSavepoint"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/savepoints/" + dara.PercentEncode(dara.StringValue(savepointId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSavepointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details of a savepoint and checkpoint.
//
// @return GetSavepointResponse
func (client *Client) GetSavepoint(namespace *string, savepointId *string) (_result *GetSavepointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetSavepointHeaders{}
	_result = &GetSavepointResponse{}
	_body, _err := client.GetSavepointWithOptions(namespace, savepointId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取session集群
//
// @param headers - GetSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSessionClusterResponse
func (client *Client) GetSessionClusterWithOptions(namespace *string, sessionClusterName *string, headers *GetSessionClusterHeaders, runtime *dara.RuntimeOptions) (_result *GetSessionClusterResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSessionCluster"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sessionclusters/" + dara.PercentEncode(dara.StringValue(sessionClusterName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSessionClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取session集群
//
// @return GetSessionClusterResponse
func (client *Client) GetSessionCluster(namespace *string, sessionClusterName *string) (_result *GetSessionClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetSessionClusterHeaders{}
	_result = &GetSessionClusterResponse{}
	_body, _err := client.GetSessionClusterWithOptions(namespace, sessionClusterName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取table
//
// @param request - GetTablesRequest
//
// @param headers - GetTablesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTablesResponse
func (client *Client) GetTablesWithOptions(namespace *string, catalogName *string, databaseName *string, request *GetTablesRequest, headers *GetTablesHeaders, runtime *dara.RuntimeOptions) (_result *GetTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TableName) {
		query["tableName"] = request.TableName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTables"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/catalogs/" + dara.PercentEncode(dara.StringValue(catalogName)) + "/databases/" + dara.PercentEncode(dara.StringValue(databaseName)) + "/tables"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取table
//
// @param request - GetTablesRequest
//
// @return GetTablesResponse
func (client *Client) GetTables(namespace *string, catalogName *string, databaseName *string, request *GetTablesRequest) (_result *GetTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetTablesHeaders{}
	_result = &GetTablesResponse{}
	_body, _err := client.GetTablesWithOptions(namespace, catalogName, databaseName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of the JAR or Python file that corresponds to the user-defined function (UDF) that you upload and create.
//
// @param request - GetUdfArtifactsRequest
//
// @param headers - GetUdfArtifactsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUdfArtifactsResponse
func (client *Client) GetUdfArtifactsWithOptions(namespace *string, request *GetUdfArtifactsRequest, headers *GetUdfArtifactsHeaders, runtime *dara.RuntimeOptions) (_result *GetUdfArtifactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UdfArtifactName) {
		query["udfArtifactName"] = request.UdfArtifactName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUdfArtifacts"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/udfartifacts"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUdfArtifactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of the JAR or Python file that corresponds to the user-defined function (UDF) that you upload and create.
//
// @param request - GetUdfArtifactsRequest
//
// @return GetUdfArtifactsResponse
func (client *Client) GetUdfArtifacts(namespace *string, request *GetUdfArtifactsRequest) (_result *GetUdfArtifactsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetUdfArtifactsHeaders{}
	_result = &GetUdfArtifactsResponse{}
	_body, _err := client.GetUdfArtifactsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Dynamically updates parameters or resources of a deployment that is running.
//
// @param headers - HotUpdateJobHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotUpdateJobResponse
func (client *Client) HotUpdateJobWithOptions(namespace *string, jobId *string, headers *HotUpdateJobHeaders, runtime *dara.RuntimeOptions) (_result *HotUpdateJobResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotUpdateJob"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId)) + "%3AhotUpdate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotUpdateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Dynamically updates parameters or resources of a deployment that is running.
//
// @return HotUpdateJobResponse
func (client *Client) HotUpdateJob(namespace *string, jobId *string) (_result *HotUpdateJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &HotUpdateJobHeaders{}
	_result = &HotUpdateJobResponse{}
	_body, _err := client.HotUpdateJobWithOptions(namespace, jobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of existing custom connectors.
//
// @param headers - ListCustomConnectorsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomConnectorsResponse
func (client *Client) ListCustomConnectorsWithOptions(namespace *string, headers *ListCustomConnectorsHeaders, runtime *dara.RuntimeOptions) (_result *ListCustomConnectorsResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomConnectors"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/connectors"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomConnectorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of existing custom connectors.
//
// @return ListCustomConnectorsResponse
func (client *Client) ListCustomConnectors(namespace *string) (_result *ListCustomConnectorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListCustomConnectorsHeaders{}
	_result = &ListCustomConnectorsResponse{}
	_body, _err := client.ListCustomConnectorsWithOptions(namespace, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// list deploymentDrafts
//
// @param request - ListDeploymentDraftsRequest
//
// @param headers - ListDeploymentDraftsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentDraftsResponse
func (client *Client) ListDeploymentDraftsWithOptions(namespace *string, request *ListDeploymentDraftsRequest, headers *ListDeploymentDraftsHeaders, runtime *dara.RuntimeOptions) (_result *ListDeploymentDraftsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeploymentDrafts"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-drafts"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeploymentDraftsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// list deploymentDrafts
//
// @param request - ListDeploymentDraftsRequest
//
// @return ListDeploymentDraftsResponse
func (client *Client) ListDeploymentDrafts(namespace *string, request *ListDeploymentDraftsRequest) (_result *ListDeploymentDraftsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListDeploymentDraftsHeaders{}
	_result = &ListDeploymentDraftsResponse{}
	_body, _err := client.ListDeploymentDraftsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of clusters in which deployments can be deployed. The cluster can be a session cluster or a per-job cluster.
//
// @param request - ListDeploymentTargetsRequest
//
// @param headers - ListDeploymentTargetsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentTargetsResponse
func (client *Client) ListDeploymentTargetsWithOptions(namespace *string, request *ListDeploymentTargetsRequest, headers *ListDeploymentTargetsHeaders, runtime *dara.RuntimeOptions) (_result *ListDeploymentTargetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeploymentTargets"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-targets"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeploymentTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of clusters in which deployments can be deployed. The cluster can be a session cluster or a per-job cluster.
//
// @param request - ListDeploymentTargetsRequest
//
// @return ListDeploymentTargetsResponse
func (client *Client) ListDeploymentTargets(namespace *string, request *ListDeploymentTargetsRequest) (_result *ListDeploymentTargetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListDeploymentTargetsHeaders{}
	_result = &ListDeploymentTargetsResponse{}
	_body, _err := client.ListDeploymentTargetsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains information about all deployments.
//
// @param request - ListDeploymentsRequest
//
// @param headers - ListDeploymentsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentsResponse
func (client *Client) ListDeploymentsWithOptions(namespace *string, request *ListDeploymentsRequest, headers *ListDeploymentsHeaders, runtime *dara.RuntimeOptions) (_result *ListDeploymentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Creator) {
		query["creator"] = request.Creator
	}

	if !dara.IsNil(request.ExecutionMode) {
		query["executionMode"] = request.ExecutionMode
	}

	if !dara.IsNil(request.LabelKey) {
		query["labelKey"] = request.LabelKey
	}

	if !dara.IsNil(request.LabelValueArray) {
		query["labelValueArray"] = request.LabelValueArray
	}

	if !dara.IsNil(request.Modifier) {
		query["modifier"] = request.Modifier
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortName) {
		query["sortName"] = request.SortName
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeployments"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeploymentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains information about all deployments.
//
// @param request - ListDeploymentsRequest
//
// @return ListDeploymentsResponse
func (client *Client) ListDeployments(namespace *string, request *ListDeploymentsRequest) (_result *ListDeploymentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListDeploymentsHeaders{}
	_result = &ListDeploymentsResponse{}
	_body, _err := client.ListDeploymentsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出有编辑权限的项目空间。
//
// @param request - ListEditableNamespaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEditableNamespaceResponse
func (client *Client) ListEditableNamespaceWithOptions(request *ListEditableNamespaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEditableNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEditableNamespace"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/gateway/v2/namespaces/editable"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEditableNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出有编辑权限的项目空间。
//
// @param request - ListEditableNamespaceRequest
//
// @return ListEditableNamespaceResponse
func (client *Client) ListEditableNamespace(request *ListEditableNamespaceRequest) (_result *ListEditableNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEditableNamespaceResponse{}
	_body, _err := client.ListEditableNamespaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of engine versions that are supported by Realtime Compute for Apache Flink.
//
// @param headers - ListEngineVersionMetadataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEngineVersionMetadataResponse
func (client *Client) ListEngineVersionMetadataWithOptions(headers *ListEngineVersionMetadataHeaders, runtime *dara.RuntimeOptions) (_result *ListEngineVersionMetadataResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEngineVersionMetadata"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/engine-version-meta.json"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEngineVersionMetadataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of engine versions that are supported by Realtime Compute for Apache Flink.
//
// @return ListEngineVersionMetadataResponse
func (client *Client) ListEngineVersionMetadata() (_result *ListEngineVersionMetadataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListEngineVersionMetadataHeaders{}
	_result = &ListEngineVersionMetadataResponse{}
	_body, _err := client.ListEngineVersionMetadataWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about all jobs in a deployment.
//
// @param request - ListJobsRequest
//
// @param headers - ListJobsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
func (client *Client) ListJobsWithOptions(namespace *string, request *ListJobsRequest, headers *ListJobsHeaders, runtime *dara.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentId) {
		query["deploymentId"] = request.DeploymentId
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortName) {
		query["sortName"] = request.SortName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobs"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/jobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all jobs in a deployment.
//
// @param request - ListJobsRequest
//
// @return ListJobsResponse
func (client *Client) ListJobs(namespace *string, request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListJobsHeaders{}
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the mappings between the ID and permissions of a member in a specific namespace.
//
// @param request - ListMembersRequest
//
// @param headers - ListMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMembersResponse
func (client *Client) ListMembersWithOptions(namespace *string, request *ListMembersRequest, headers *ListMembersHeaders, runtime *dara.RuntimeOptions) (_result *ListMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMembers"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/gateway/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/members"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the mappings between the ID and permissions of a member in a specific namespace.
//
// @param request - ListMembersRequest
//
// @return ListMembersResponse
func (client *Client) ListMembers(namespace *string, request *ListMembersRequest) (_result *ListMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListMembersHeaders{}
	_result = &ListMembersResponse{}
	_body, _err := client.ListMembersWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of savepoints or checkpoints.
//
// @param request - ListSavepointsRequest
//
// @param headers - ListSavepointsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSavepointsResponse
func (client *Client) ListSavepointsWithOptions(namespace *string, request *ListSavepointsRequest, headers *ListSavepointsHeaders, runtime *dara.RuntimeOptions) (_result *ListSavepointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentId) {
		query["deploymentId"] = request.DeploymentId
	}

	if !dara.IsNil(request.JobId) {
		query["jobId"] = request.JobId
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSavepoints"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/savepoints"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSavepointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of savepoints or checkpoints.
//
// @param request - ListSavepointsRequest
//
// @return ListSavepointsResponse
func (client *Client) ListSavepoints(namespace *string, request *ListSavepointsRequest) (_result *ListSavepointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListSavepointsHeaders{}
	_result = &ListSavepointsResponse{}
	_body, _err := client.ListSavepointsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列表定时执行计划
//
// @param request - ListScheduledPlanRequest
//
// @param headers - ListScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduledPlanResponse
func (client *Client) ListScheduledPlanWithOptions(namespace *string, request *ListScheduledPlanRequest, headers *ListScheduledPlanHeaders, runtime *dara.RuntimeOptions) (_result *ListScheduledPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentId) {
		query["deploymentId"] = request.DeploymentId
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScheduledPlan"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/scheduled-plans"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScheduledPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表定时执行计划
//
// @param request - ListScheduledPlanRequest
//
// @return ListScheduledPlanResponse
func (client *Client) ListScheduledPlan(namespace *string, request *ListScheduledPlanRequest) (_result *ListScheduledPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListScheduledPlanHeaders{}
	_result = &ListScheduledPlanResponse{}
	_body, _err := client.ListScheduledPlanWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取作业资源变更历史
//
// @param request - ListScheduledPlanExecutedHistoryRequest
//
// @param headers - ListScheduledPlanExecutedHistoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduledPlanExecutedHistoryResponse
func (client *Client) ListScheduledPlanExecutedHistoryWithOptions(namespace *string, request *ListScheduledPlanExecutedHistoryRequest, headers *ListScheduledPlanExecutedHistoryHeaders, runtime *dara.RuntimeOptions) (_result *ListScheduledPlanExecutedHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentId) {
		query["deploymentId"] = request.DeploymentId
	}

	if !dara.IsNil(request.Origin) {
		query["origin"] = request.Origin
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScheduledPlanExecutedHistory"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/job-resource-upgradings"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScheduledPlanExecutedHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取作业资源变更历史
//
// @param request - ListScheduledPlanExecutedHistoryRequest
//
// @return ListScheduledPlanExecutedHistoryResponse
func (client *Client) ListScheduledPlanExecutedHistory(namespace *string, request *ListScheduledPlanExecutedHistoryRequest) (_result *ListScheduledPlanExecutedHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListScheduledPlanExecutedHistoryHeaders{}
	_result = &ListScheduledPlanExecutedHistoryResponse{}
	_body, _err := client.ListScheduledPlanExecutedHistoryWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举session集群
//
// @param headers - ListSessionClustersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSessionClustersResponse
func (client *Client) ListSessionClustersWithOptions(namespace *string, headers *ListSessionClustersHeaders, runtime *dara.RuntimeOptions) (_result *ListSessionClustersResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSessionClusters"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sessionclusters"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSessionClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举session集群
//
// @return ListSessionClustersResponse
func (client *Client) ListSessionClusters(namespace *string) (_result *ListSessionClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListSessionClustersHeaders{}
	_result = &ListSessionClustersResponse{}
	_body, _err := client.ListSessionClustersWithOptions(namespace, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of variables.
//
// @param request - ListVariablesRequest
//
// @param headers - ListVariablesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVariablesResponse
func (client *Client) ListVariablesWithOptions(namespace *string, request *ListVariablesRequest, headers *ListVariablesHeaders, runtime *dara.RuntimeOptions) (_result *ListVariablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVariables"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/variables"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVariablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of variables.
//
// @param request - ListVariablesRequest
//
// @return ListVariablesResponse
func (client *Client) ListVariables(namespace *string, request *ListVariablesRequest) (_result *ListVariablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListVariablesHeaders{}
	_result = &ListVariablesResponse{}
	_body, _err := client.ListVariablesWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers a custom connector in a namespace. The registered custom connector can be used in SQL statements.
//
// @param request - RegisterCustomConnectorRequest
//
// @param headers - RegisterCustomConnectorHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterCustomConnectorResponse
func (client *Client) RegisterCustomConnectorWithOptions(namespace *string, request *RegisterCustomConnectorRequest, headers *RegisterCustomConnectorHeaders, runtime *dara.RuntimeOptions) (_result *RegisterCustomConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JarUrl) {
		query["jarUrl"] = request.JarUrl
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterCustomConnector"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/connectors%3Aregister"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterCustomConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers a custom connector in a namespace. The registered custom connector can be used in SQL statements.
//
// @param request - RegisterCustomConnectorRequest
//
// @return RegisterCustomConnectorResponse
func (client *Client) RegisterCustomConnector(namespace *string, request *RegisterCustomConnectorRequest) (_result *RegisterCustomConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &RegisterCustomConnectorHeaders{}
	_result = &RegisterCustomConnectorResponse{}
	_body, _err := client.RegisterCustomConnectorWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers specific or all of the user-defined functions (UDFs) that are parsed from the JAR files. The registered functions can be used in SQL statements.
//
// @param request - RegisterUdfFunctionRequest
//
// @param headers - RegisterUdfFunctionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterUdfFunctionResponse
func (client *Client) RegisterUdfFunctionWithOptions(namespace *string, request *RegisterUdfFunctionRequest, headers *RegisterUdfFunctionHeaders, runtime *dara.RuntimeOptions) (_result *RegisterUdfFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClassName) {
		query["className"] = request.ClassName
	}

	if !dara.IsNil(request.FunctionName) {
		query["functionName"] = request.FunctionName
	}

	if !dara.IsNil(request.UdfArtifactName) {
		query["udfArtifactName"] = request.UdfArtifactName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterUdfFunction"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/udfartifacts/function%3AregisterUdfFunction"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterUdfFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers specific or all of the user-defined functions (UDFs) that are parsed from the JAR files. The registered functions can be used in SQL statements.
//
// @param request - RegisterUdfFunctionRequest
//
// @return RegisterUdfFunctionResponse
func (client *Client) RegisterUdfFunction(namespace *string, request *RegisterUdfFunctionRequest) (_result *RegisterUdfFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &RegisterUdfFunctionHeaders{}
	_result = &RegisterUdfFunctionResponse{}
	_body, _err := client.RegisterUdfFunctionWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI StartJob is deprecated
//
// Summary:
//
// Creates and starts a job.
//
// @param request - StartJobRequest
//
// @param headers - StartJobHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartJobResponse
func (client *Client) StartJobWithOptions(namespace *string, request *StartJobRequest, headers *StartJobHeaders, runtime *dara.RuntimeOptions) (_result *StartJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartJob"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/jobs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI StartJob is deprecated
//
// Summary:
//
// Creates and starts a job.
//
// @param request - StartJobRequest
//
// @return StartJobResponse
// Deprecated
func (client *Client) StartJob(namespace *string, request *StartJobRequest) (_result *StartJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &StartJobHeaders{}
	_result = &StartJobResponse{}
	_body, _err := client.StartJobWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a job.
//
// @param request - StartJobWithParamsRequest
//
// @param headers - StartJobWithParamsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartJobWithParamsResponse
func (client *Client) StartJobWithParamsWithOptions(namespace *string, request *StartJobWithParamsRequest, headers *StartJobWithParamsHeaders, runtime *dara.RuntimeOptions) (_result *StartJobWithParamsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartJobWithParams"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/jobs%3Astart"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartJobWithParamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a job.
//
// @param request - StartJobWithParamsRequest
//
// @return StartJobWithParamsResponse
func (client *Client) StartJobWithParams(namespace *string, request *StartJobWithParamsRequest) (_result *StartJobWithParamsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &StartJobWithParamsHeaders{}
	_result = &StartJobWithParamsResponse{}
	_body, _err := client.StartJobWithParamsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动session集群
//
// @param headers - StartSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSessionClusterResponse
func (client *Client) StartSessionClusterWithOptions(namespace *string, sessionClusterName *string, headers *StartSessionClusterHeaders, runtime *dara.RuntimeOptions) (_result *StartSessionClusterResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartSessionCluster"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sessionclusters/" + dara.PercentEncode(dara.StringValue(sessionClusterName)) + "%3Astart"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartSessionClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动session集群
//
// @return StartSessionClusterResponse
func (client *Client) StartSessionCluster(namespace *string, sessionClusterName *string) (_result *StartSessionClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &StartSessionClusterHeaders{}
	_result = &StartSessionClusterResponse{}
	_body, _err := client.StartSessionClusterWithOptions(namespace, sessionClusterName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止应用执行定时计划
//
// @param headers - StopApplyScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopApplyScheduledPlanResponse
func (client *Client) StopApplyScheduledPlanWithOptions(namespace *string, scheduledPlanId *string, headers *StopApplyScheduledPlanHeaders, runtime *dara.RuntimeOptions) (_result *StopApplyScheduledPlanResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopApplyScheduledPlan"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/scheduled-plans/" + dara.PercentEncode(dara.StringValue(scheduledPlanId)) + "%3Astop"),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopApplyScheduledPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止应用执行定时计划
//
// @return StopApplyScheduledPlanResponse
func (client *Client) StopApplyScheduledPlan(namespace *string, scheduledPlanId *string) (_result *StopApplyScheduledPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &StopApplyScheduledPlanHeaders{}
	_result = &StopApplyScheduledPlanResponse{}
	_body, _err := client.StopApplyScheduledPlanWithOptions(namespace, scheduledPlanId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a job.
//
// @param request - StopJobRequest
//
// @param headers - StopJobHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopJobResponse
func (client *Client) StopJobWithOptions(namespace *string, jobId *string, request *StopJobRequest, headers *StopJobHeaders, runtime *dara.RuntimeOptions) (_result *StopJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopJob"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId)) + "%3Astop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a job.
//
// @param request - StopJobRequest
//
// @return StopJobResponse
func (client *Client) StopJob(namespace *string, jobId *string, request *StopJobRequest) (_result *StopJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &StopJobHeaders{}
	_result = &StopJobResponse{}
	_body, _err := client.StopJobWithOptions(namespace, jobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止session集群
//
// @param headers - StopSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopSessionClusterResponse
func (client *Client) StopSessionClusterWithOptions(namespace *string, sessionClusterName *string, headers *StopSessionClusterHeaders, runtime *dara.RuntimeOptions) (_result *StopSessionClusterResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopSessionCluster"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sessionclusters/" + dara.PercentEncode(dara.StringValue(sessionClusterName)) + "%3Astop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopSessionClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止session集群
//
// @return StopSessionClusterResponse
func (client *Client) StopSessionCluster(namespace *string, sessionClusterName *string) (_result *StopSessionClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &StopSessionClusterHeaders{}
	_result = &StopSessionClusterResponse{}
	_body, _err := client.StopSessionClusterWithOptions(namespace, sessionClusterName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交sql调试
//
// @param request - SubmitSqlPreviewRequest
//
// @param headers - SubmitSqlPreviewHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSqlPreviewResponse
func (client *Client) SubmitSqlPreviewWithOptions(namespace *string, request *SubmitSqlPreviewRequest, headers *SubmitSqlPreviewHeaders, runtime *dara.RuntimeOptions) (_result *SubmitSqlPreviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SessionClusterName) {
		query["sessionClusterName"] = request.SessionClusterName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSqlPreview"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sql-preview/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSqlPreviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交sql调试
//
// @param request - SubmitSqlPreviewRequest
//
// @return SubmitSqlPreviewResponse
func (client *Client) SubmitSqlPreview(namespace *string, request *SubmitSqlPreviewRequest) (_result *SubmitSqlPreviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SubmitSqlPreviewHeaders{}
	_result = &SubmitSqlPreviewResponse{}
	_body, _err := client.SubmitSqlPreviewWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates information about a deployment.
//
// @param request - UpdateDeploymentRequest
//
// @param headers - UpdateDeploymentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeploymentResponse
func (client *Client) UpdateDeploymentWithOptions(namespace *string, deploymentId *string, request *UpdateDeploymentRequest, headers *UpdateDeploymentHeaders, runtime *dara.RuntimeOptions) (_result *UpdateDeploymentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDeployment"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments/" + dara.PercentEncode(dara.StringValue(deploymentId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates information about a deployment.
//
// @param request - UpdateDeploymentRequest
//
// @return UpdateDeploymentResponse
func (client *Client) UpdateDeployment(namespace *string, deploymentId *string, request *UpdateDeploymentRequest) (_result *UpdateDeploymentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateDeploymentHeaders{}
	_result = &UpdateDeploymentResponse{}
	_body, _err := client.UpdateDeploymentWithOptions(namespace, deploymentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// update a deploymentDraft
//
// @param request - UpdateDeploymentDraftRequest
//
// @param headers - UpdateDeploymentDraftHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeploymentDraftResponse
func (client *Client) UpdateDeploymentDraftWithOptions(namespace *string, deploymentDraftId *string, request *UpdateDeploymentDraftRequest, headers *UpdateDeploymentDraftHeaders, runtime *dara.RuntimeOptions) (_result *UpdateDeploymentDraftResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDeploymentDraft"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-drafts/" + dara.PercentEncode(dara.StringValue(deploymentDraftId))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDeploymentDraftResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// update a deploymentDraft
//
// @param request - UpdateDeploymentDraftRequest
//
// @return UpdateDeploymentDraftResponse
func (client *Client) UpdateDeploymentDraft(namespace *string, deploymentDraftId *string, request *UpdateDeploymentDraftRequest) (_result *UpdateDeploymentDraftResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateDeploymentDraftHeaders{}
	_result = &UpdateDeploymentDraftResponse{}
	_body, _err := client.UpdateDeploymentDraftWithOptions(namespace, deploymentDraftId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改deploymentTarget
//
// @param request - UpdateDeploymentTargetRequest
//
// @param headers - UpdateDeploymentTargetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeploymentTargetResponse
func (client *Client) UpdateDeploymentTargetWithOptions(namespace *string, deploymentTargetName *string, request *UpdateDeploymentTargetRequest, headers *UpdateDeploymentTargetHeaders, runtime *dara.RuntimeOptions) (_result *UpdateDeploymentTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDeploymentTarget"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-targets/" + dara.PercentEncode(dara.StringValue(deploymentTargetName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDeploymentTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改deploymentTarget
//
// @param request - UpdateDeploymentTargetRequest
//
// @return UpdateDeploymentTargetResponse
func (client *Client) UpdateDeploymentTarget(namespace *string, deploymentTargetName *string, request *UpdateDeploymentTargetRequest) (_result *UpdateDeploymentTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateDeploymentTargetHeaders{}
	_result = &UpdateDeploymentTargetResponse{}
	_body, _err := client.UpdateDeploymentTargetWithOptions(namespace, deploymentTargetName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新部署目标
//
// @param request - UpdateDeploymentTargetV2Request
//
// @param headers - UpdateDeploymentTargetV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeploymentTargetV2Response
func (client *Client) UpdateDeploymentTargetV2WithOptions(namespace *string, deploymentTargetName *string, request *UpdateDeploymentTargetV2Request, headers *UpdateDeploymentTargetV2Headers, runtime *dara.RuntimeOptions) (_result *UpdateDeploymentTargetV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDeploymentTargetV2"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-targets/support-elastic/" + dara.PercentEncode(dara.StringValue(deploymentTargetName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDeploymentTargetV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新部署目标
//
// @param request - UpdateDeploymentTargetV2Request
//
// @return UpdateDeploymentTargetV2Response
func (client *Client) UpdateDeploymentTargetV2(namespace *string, deploymentTargetName *string, request *UpdateDeploymentTargetV2Request) (_result *UpdateDeploymentTargetV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateDeploymentTargetV2Headers{}
	_result = &UpdateDeploymentTargetV2Response{}
	_body, _err := client.UpdateDeploymentTargetV2WithOptions(namespace, deploymentTargetName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// update a folder
//
// @param request - UpdateFolderRequest
//
// @param headers - UpdateFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFolderResponse
func (client *Client) UpdateFolderWithOptions(namespace *string, folderId *string, request *UpdateFolderRequest, headers *UpdateFolderHeaders, runtime *dara.RuntimeOptions) (_result *UpdateFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFolder"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/folder/" + dara.PercentEncode(dara.StringValue(folderId))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// update a folder
//
// @param request - UpdateFolderRequest
//
// @return UpdateFolderResponse
func (client *Client) UpdateFolder(namespace *string, folderId *string, request *UpdateFolderRequest) (_result *UpdateFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateFolderHeaders{}
	_result = &UpdateFolderResponse{}
	_body, _err := client.UpdateFolderWithOptions(namespace, folderId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the permissions of one or more members in a specific namespace.
//
// @param request - UpdateMemberRequest
//
// @param headers - UpdateMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemberResponse
func (client *Client) UpdateMemberWithOptions(namespace *string, request *UpdateMemberRequest, headers *UpdateMemberHeaders, runtime *dara.RuntimeOptions) (_result *UpdateMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMember"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/gateway/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/members"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the permissions of one or more members in a specific namespace.
//
// @param request - UpdateMemberRequest
//
// @return UpdateMemberResponse
func (client *Client) UpdateMember(namespace *string, request *UpdateMemberRequest) (_result *UpdateMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateMemberHeaders{}
	_result = &UpdateMemberResponse{}
	_body, _err := client.UpdateMemberWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新定时执行计划
//
// @param request - UpdateScheduledPlanRequest
//
// @param headers - UpdateScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScheduledPlanResponse
func (client *Client) UpdateScheduledPlanWithOptions(namespace *string, scheduledPlanId *string, request *UpdateScheduledPlanRequest, headers *UpdateScheduledPlanHeaders, runtime *dara.RuntimeOptions) (_result *UpdateScheduledPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateScheduledPlan"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/scheduled-plans/" + dara.PercentEncode(dara.StringValue(scheduledPlanId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateScheduledPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新定时执行计划
//
// @param request - UpdateScheduledPlanRequest
//
// @return UpdateScheduledPlanResponse
func (client *Client) UpdateScheduledPlan(namespace *string, scheduledPlanId *string, request *UpdateScheduledPlanRequest) (_result *UpdateScheduledPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateScheduledPlanHeaders{}
	_result = &UpdateScheduledPlanResponse{}
	_body, _err := client.UpdateScheduledPlanWithOptions(namespace, scheduledPlanId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新session集群
//
// @param request - UpdateSessionClusterRequest
//
// @param headers - UpdateSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSessionClusterResponse
func (client *Client) UpdateSessionClusterWithOptions(namespace *string, sessionClusterName *string, request *UpdateSessionClusterRequest, headers *UpdateSessionClusterHeaders, runtime *dara.RuntimeOptions) (_result *UpdateSessionClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSessionCluster"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sessionclusters/" + dara.PercentEncode(dara.StringValue(sessionClusterName))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSessionClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新session集群
//
// @param request - UpdateSessionClusterRequest
//
// @return UpdateSessionClusterResponse
func (client *Client) UpdateSessionCluster(namespace *string, sessionClusterName *string, request *UpdateSessionClusterRequest) (_result *UpdateSessionClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateSessionClusterHeaders{}
	_result = &UpdateSessionClusterResponse{}
	_body, _err := client.UpdateSessionClusterWithOptions(namespace, sessionClusterName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the JAR file of the user-defined function (UDF) that you create.
//
// @param request - UpdateUdfArtifactRequest
//
// @param headers - UpdateUdfArtifactHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUdfArtifactResponse
func (client *Client) UpdateUdfArtifactWithOptions(namespace *string, udfArtifactName *string, request *UpdateUdfArtifactRequest, headers *UpdateUdfArtifactHeaders, runtime *dara.RuntimeOptions) (_result *UpdateUdfArtifactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUdfArtifact"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/udfartifacts/" + dara.PercentEncode(dara.StringValue(udfArtifactName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUdfArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the JAR file of the user-defined function (UDF) that you create.
//
// @param request - UpdateUdfArtifactRequest
//
// @return UpdateUdfArtifactResponse
func (client *Client) UpdateUdfArtifact(namespace *string, udfArtifactName *string, request *UpdateUdfArtifactRequest) (_result *UpdateUdfArtifactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateUdfArtifactHeaders{}
	_result = &UpdateUdfArtifactResponse{}
	_body, _err := client.UpdateUdfArtifactWithOptions(namespace, udfArtifactName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新秘钥
//
// @param request - UpdateVariableRequest
//
// @param headers - UpdateVariableHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVariableResponse
func (client *Client) UpdateVariableWithOptions(namespace *string, name *string, request *UpdateVariableRequest, headers *UpdateVariableHeaders, runtime *dara.RuntimeOptions) (_result *UpdateVariableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVariable"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/variables/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新秘钥
//
// @param request - UpdateVariableRequest
//
// @return UpdateVariableResponse
func (client *Client) UpdateVariable(namespace *string, name *string, request *UpdateVariableRequest) (_result *UpdateVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateVariableHeaders{}
	_result = &UpdateVariableResponse{}
	_body, _err := client.UpdateVariableWithOptions(namespace, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the code of an SQL deployment.
//
// @param request - ValidateSqlStatementRequest
//
// @param headers - ValidateSqlStatementHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateSqlStatementResponse
func (client *Client) ValidateSqlStatementWithOptions(namespace *string, request *ValidateSqlStatementRequest, headers *ValidateSqlStatementHeaders, runtime *dara.RuntimeOptions) (_result *ValidateSqlStatementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateSqlStatement"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sql-statement/validate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateSqlStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the code of an SQL deployment.
//
// @param request - ValidateSqlStatementRequest
//
// @return ValidateSqlStatementResponse
func (client *Client) ValidateSqlStatement(namespace *string, request *ValidateSqlStatementRequest) (_result *ValidateSqlStatementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ValidateSqlStatementHeaders{}
	_result = &ValidateSqlStatementResponse{}
	_body, _err := client.ValidateSqlStatementWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
