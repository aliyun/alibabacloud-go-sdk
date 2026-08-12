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
		"public": dara.String("riskmanagement.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("riskmanagement"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Bindss authorization to machines in Security Center.
//
// @param tmpReq - BindAuthToMachineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAuthToMachineResponse
func (client *Client) BindAuthToMachineWithOptions(tmpReq *BindAuthToMachineRequest, runtime *dara.RuntimeOptions) (_result *BindAuthToMachineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BindAuthToMachineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAuthToMachine"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAuthToMachineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Bindss authorization to machines in Security Center.
//
// @param request - BindAuthToMachineRequest
//
// @return BindAuthToMachineResponse
func (client *Client) BindAuthToMachine(request *BindAuthToMachineRequest) (_result *BindAuthToMachineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindAuthToMachineResponse{}
	_body, _err := client.BindAuthToMachineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initializes Security Center module rules.
//
// @param tmpReq - CreateSasTrialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSasTrialResponse
func (client *Client) CreateSasTrialWithOptions(tmpReq *CreateSasTrialRequest, runtime *dara.RuntimeOptions) (_result *CreateSasTrialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSasTrialShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSasTrial"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSasTrialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initializes Security Center module rules.
//
// @param request - CreateSasTrialRequest
//
// @return CreateSasTrialResponse
func (client *Client) CreateSasTrial(request *CreateSasTrialRequest) (_result *CreateSasTrialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSasTrialResponse{}
	_body, _err := client.CreateSasTrialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a service-linked role for Security Center and authorizes Security Center to access cloud resources.
//
// @param tmpReq - CreateServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceLinkedRoleResponse
func (client *Client) CreateServiceLinkedRoleWithOptions(tmpReq *CreateServiceLinkedRoleRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateServiceLinkedRoleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceLinkedRole"),
		Version:     dara.String("2026-04-24"),
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
// Creates a service-linked role for Security Center and authorizes Security Center to access cloud resources.
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
// Creates a full disk scan task.
//
// @param request - CreateVirusScanOnceTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVirusScanOnceTaskResponse
func (client *Client) CreateVirusScanOnceTaskWithOptions(request *CreateVirusScanOnceTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateVirusScanOnceTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVirusScanOnceTask"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVirusScanOnceTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a full disk scan task.
//
// @param request - CreateVirusScanOnceTaskRequest
//
// @return CreateVirusScanOnceTaskResponse
func (client *Client) CreateVirusScanOnceTask(request *CreateVirusScanOnceTaskRequest) (_result *CreateVirusScanOnceTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVirusScanOnceTaskResponse{}
	_body, _err := client.CreateVirusScanOnceTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of Security Center instances.
//
// @param tmpReq - DescribeCloudCenterInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudCenterInstancesResponse
func (client *Client) DescribeCloudCenterInstancesWithOptions(tmpReq *DescribeCloudCenterInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudCenterInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeCloudCenterInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudCenterInstances"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudCenterInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of Security Center instances.
//
// @param request - DescribeCloudCenterInstancesRequest
//
// @return DescribeCloudCenterInstancesResponse
func (client *Client) DescribeCloudCenterInstances(request *DescribeCloudCenterInstancesRequest) (_result *DescribeCloudCenterInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudCenterInstancesResponse{}
	_body, _err := client.DescribeCloudCenterInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a service-linked role for Security Center.
//
// @param tmpReq - DescribeServiceLinkedRoleStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceLinkedRoleStatusResponse
func (client *Client) DescribeServiceLinkedRoleStatusWithOptions(tmpReq *DescribeServiceLinkedRoleStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceLinkedRoleStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeServiceLinkedRoleStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceLinkedRoleStatus"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceLinkedRoleStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a service-linked role for Security Center.
//
// @param request - DescribeServiceLinkedRoleStatusRequest
//
// @return DescribeServiceLinkedRoleStatusResponse
func (client *Client) DescribeServiceLinkedRoleStatus(request *DescribeServiceLinkedRoleStatusRequest) (_result *DescribeServiceLinkedRoleStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeServiceLinkedRoleStatusResponse{}
	_body, _err := client.DescribeServiceLinkedRoleStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries security alert events in Security Center.
//
// @param tmpReq - DescribeSuspEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSuspEventsResponse
func (client *Client) DescribeSuspEventsWithOptions(tmpReq *DescribeSuspEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSuspEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeSuspEventsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSuspEvents"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSuspEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries security alert events in Security Center.
//
// @param request - DescribeSuspEventsRequest
//
// @return DescribeSuspEventsResponse
func (client *Client) DescribeSuspEvents(request *DescribeSuspEventsRequest) (_result *DescribeSuspEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSuspEventsResponse{}
	_body, _err := client.DescribeSuspEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the edition details of a purchased Security Center instance.
//
// @param tmpReq - DescribeVersionConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVersionConfigResponse
func (client *Client) DescribeVersionConfigWithOptions(tmpReq *DescribeVersionConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeVersionConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeVersionConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVersionConfig"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVersionConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the edition details of a purchased Security Center instance.
//
// @param request - DescribeVersionConfigRequest
//
// @return DescribeVersionConfigResponse
func (client *Client) DescribeVersionConfig(request *DescribeVersionConfigRequest) (_result *DescribeVersionConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVersionConfigResponse{}
	_body, _err := client.DescribeVersionConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the analysis results of alert records.
//
// @param tmpReq - GetAlertRecordAnalysisResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlertRecordAnalysisResultResponse
func (client *Client) GetAlertRecordAnalysisResultWithOptions(tmpReq *GetAlertRecordAnalysisResultRequest, runtime *dara.RuntimeOptions) (_result *GetAlertRecordAnalysisResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAlertRecordAnalysisResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UniqueTagList) {
		request.UniqueTagListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UniqueTagList, dara.String("UniqueTagList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmUniqueInfo) {
		query["AlarmUniqueInfo"] = request.AlarmUniqueInfo
	}

	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.UniqueInfo) {
		query["UniqueInfo"] = request.UniqueInfo
	}

	if !dara.IsNil(request.UniqueTagListShrink) {
		query["UniqueTagList"] = request.UniqueTagListShrink
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlertRecordAnalysisResult"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlertRecordAnalysisResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the analysis results of alert records.
//
// @param request - GetAlertRecordAnalysisResultRequest
//
// @return GetAlertRecordAnalysisResultResponse
func (client *Client) GetAlertRecordAnalysisResult(request *GetAlertRecordAnalysisResultRequest) (_result *GetAlertRecordAnalysisResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAlertRecordAnalysisResultResponse{}
	_body, _err := client.GetAlertRecordAnalysisResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls specific Security Center API operations.
//
// @param tmpReq - GetAliYunSafeCenterResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAliYunSafeCenterResultResponse
func (client *Client) GetAliYunSafeCenterResultWithOptions(tmpReq *GetAliYunSafeCenterResultRequest, runtime *dara.RuntimeOptions) (_result *GetAliYunSafeCenterResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAliYunSafeCenterResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateSimilarSecurityEventsQueryTaskRequest) {
		request.CreateSimilarSecurityEventsQueryTaskRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateSimilarSecurityEventsQueryTaskRequest, dara.String("CreateSimilarSecurityEventsQueryTaskRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DescribeInstancesFullStatusRequest) {
		request.DescribeInstancesFullStatusRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DescribeInstancesFullStatusRequest, dara.String("DescribeInstancesFullStatusRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DescribeSecurityEventOperationStatusRequest) {
		request.DescribeSecurityEventOperationStatusRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DescribeSecurityEventOperationStatusRequest, dara.String("DescribeSecurityEventOperationStatusRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DescribeSimilarSecurityEventsRequest) {
		request.DescribeSimilarSecurityEventsRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DescribeSimilarSecurityEventsRequest, dara.String("DescribeSimilarSecurityEventsRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.GetAssetDetailByUuidRequest) {
		request.GetAssetDetailByUuidRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GetAssetDetailByUuidRequest, dara.String("GetAssetDetailByUuidRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HandleSecurityEventsRequest) {
		request.HandleSecurityEventsRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HandleSecurityEventsRequest, dara.String("HandleSecurityEventsRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HandleSimilarSecurityEventsRequest) {
		request.HandleSimilarSecurityEventsRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HandleSimilarSecurityEventsRequest, dara.String("HandleSimilarSecurityEventsRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ListInstancesRequest) {
		request.ListInstancesRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListInstancesRequest, dara.String("ListInstancesRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateSimilarSecurityEventsQueryTaskRequestShrink) {
		query["CreateSimilarSecurityEventsQueryTaskRequest"] = request.CreateSimilarSecurityEventsQueryTaskRequestShrink
	}

	if !dara.IsNil(request.DescribeInstancesFullStatusRequestShrink) {
		query["DescribeInstancesFullStatusRequest"] = request.DescribeInstancesFullStatusRequestShrink
	}

	if !dara.IsNil(request.DescribeSecurityEventOperationStatusRequestShrink) {
		query["DescribeSecurityEventOperationStatusRequest"] = request.DescribeSecurityEventOperationStatusRequestShrink
	}

	if !dara.IsNil(request.DescribeSimilarSecurityEventsRequestShrink) {
		query["DescribeSimilarSecurityEventsRequest"] = request.DescribeSimilarSecurityEventsRequestShrink
	}

	if !dara.IsNil(request.GetAssetDetailByUuidRequestShrink) {
		query["GetAssetDetailByUuidRequest"] = request.GetAssetDetailByUuidRequestShrink
	}

	if !dara.IsNil(request.HandleSecurityEventsRequestShrink) {
		query["HandleSecurityEventsRequest"] = request.HandleSecurityEventsRequestShrink
	}

	if !dara.IsNil(request.HandleSimilarSecurityEventsRequestShrink) {
		query["HandleSimilarSecurityEventsRequest"] = request.HandleSimilarSecurityEventsRequestShrink
	}

	if !dara.IsNil(request.InterfaceCode) {
		query["InterfaceCode"] = request.InterfaceCode
	}

	if !dara.IsNil(request.ListInstancesRequestShrink) {
		query["ListInstancesRequest"] = request.ListInstancesRequestShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAliYunSafeCenterResult"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAliYunSafeCenterResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls specific Security Center API operations.
//
// @param request - GetAliYunSafeCenterResultRequest
//
// @return GetAliYunSafeCenterResultResponse
func (client *Client) GetAliYunSafeCenterResult(request *GetAliYunSafeCenterResultRequest) (_result *GetAliYunSafeCenterResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAliYunSafeCenterResultResponse{}
	_body, _err := client.GetAliYunSafeCenterResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether a user is eligible for a Security Center free trial.
//
// @param tmpReq - GetCanTrySasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCanTrySasResponse
func (client *Client) GetCanTrySasWithOptions(tmpReq *GetCanTrySasRequest, runtime *dara.RuntimeOptions) (_result *GetCanTrySasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetCanTrySasShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCanTrySas"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCanTrySasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether a user is eligible for a Security Center free trial.
//
// @param request - GetCanTrySasRequest
//
// @return GetCanTrySasResponse
func (client *Client) GetCanTrySas(request *GetCanTrySasRequest) (_result *GetCanTrySasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCanTrySasResponse{}
	_body, _err := client.GetCanTrySasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the security compliance package ID.
//
// @param request - GetCompliancePackIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCompliancePackIdResponse
func (client *Client) GetCompliancePackIdWithOptions(request *GetCompliancePackIdRequest, runtime *dara.RuntimeOptions) (_result *GetCompliancePackIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetCompliancePackId"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCompliancePackIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the security compliance package ID.
//
// @param request - GetCompliancePackIdRequest
//
// @return GetCompliancePackIdResponse
func (client *Client) GetCompliancePackId(request *GetCompliancePackIdRequest) (_result *GetCompliancePackIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCompliancePackIdResponse{}
	_body, _err := client.GetCompliancePackIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the authorization status for one-click disposal.
//
// @param request - GetDisposalToolStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDisposalToolStatusResponse
func (client *Client) GetDisposalToolStatusWithOptions(request *GetDisposalToolStatusRequest, runtime *dara.RuntimeOptions) (_result *GetDisposalToolStatusResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDisposalToolStatus"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDisposalToolStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the authorization status for one-click disposal.
//
// @param request - GetDisposalToolStatusRequest
//
// @return GetDisposalToolStatusResponse
func (client *Client) GetDisposalToolStatus(request *GetDisposalToolStatusRequest) (_result *GetDisposalToolStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDisposalToolStatusResponse{}
	_body, _err := client.GetDisposalToolStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the record of a user confirming a security contact.
//
// @param request - GetNotificationClickRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNotificationClickRecordResponse
func (client *Client) GetNotificationClickRecordWithOptions(request *GetNotificationClickRecordRequest, runtime *dara.RuntimeOptions) (_result *GetNotificationClickRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetNotificationClickRecord"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNotificationClickRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the record of a user confirming a security contact.
//
// @param request - GetNotificationClickRecordRequest
//
// @return GetNotificationClickRecordResponse
func (client *Client) GetNotificationClickRecord(request *GetNotificationClickRecordRequest) (_result *GetNotificationClickRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNotificationClickRecordResponse{}
	_body, _err := client.GetNotificationClickRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves all information about security contacts.
//
// @param request - GetNotificationContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNotificationContactsResponse
func (client *Client) GetNotificationContactsWithOptions(request *GetNotificationContactsRequest, runtime *dara.RuntimeOptions) (_result *GetNotificationContactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetNotificationContacts"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNotificationContactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all information about security contacts.
//
// @param request - GetNotificationContactsRequest
//
// @return GetNotificationContactsResponse
func (client *Client) GetNotificationContacts(request *GetNotificationContactsRequest) (_result *GetNotificationContactsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNotificationContactsResponse{}
	_body, _err := client.GetNotificationContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the number of pending items for security contacts.
//
// @param request - GetNotificationPendNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNotificationPendNumberResponse
func (client *Client) GetNotificationPendNumberWithOptions(request *GetNotificationPendNumberRequest, runtime *dara.RuntimeOptions) (_result *GetNotificationPendNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetNotificationPendNumber"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNotificationPendNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the number of pending items for security contacts.
//
// @param request - GetNotificationPendNumberRequest
//
// @return GetNotificationPendNumberResponse
func (client *Client) GetNotificationPendNumber(request *GetNotificationPendNumberRequest) (_result *GetNotificationPendNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNotificationPendNumberResponse{}
	_body, _err := client.GetNotificationPendNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of cloud resource control events.
//
// @param tmpReq - GetResourceControlEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceControlEventResponse
func (client *Client) GetResourceControlEventWithOptions(tmpReq *GetResourceControlEventRequest, runtime *dara.RuntimeOptions) (_result *GetResourceControlEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetResourceControlEventShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EventIdList) {
		request.EventIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventIdList, dara.String("EventIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.EventIdListShrink) {
		query["EventIdList"] = request.EventIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceControlEvent"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceControlEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of cloud resource control events.
//
// @param request - GetResourceControlEventRequest
//
// @return GetResourceControlEventResponse
func (client *Client) GetResourceControlEvent(request *GetResourceControlEventRequest) (_result *GetResourceControlEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceControlEventResponse{}
	_body, _err := client.GetResourceControlEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the basic information of a security check.
//
// @param request - GetSecurityCheckBaseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecurityCheckBaseInfoResponse
func (client *Client) GetSecurityCheckBaseInfoWithOptions(request *GetSecurityCheckBaseInfoRequest, runtime *dara.RuntimeOptions) (_result *GetSecurityCheckBaseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecurityCheckBaseInfo"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecurityCheckBaseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the basic information of a security check.
//
// @param request - GetSecurityCheckBaseInfoRequest
//
// @return GetSecurityCheckBaseInfoResponse
func (client *Client) GetSecurityCheckBaseInfo(request *GetSecurityCheckBaseInfoRequest) (_result *GetSecurityCheckBaseInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSecurityCheckBaseInfoResponse{}
	_body, _err := client.GetSecurityCheckBaseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the basic information of security check results.
//
// @param request - GetSecurityCheckResultBaseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecurityCheckResultBaseInfoResponse
func (client *Client) GetSecurityCheckResultBaseInfoWithOptions(request *GetSecurityCheckResultBaseInfoRequest, runtime *dara.RuntimeOptions) (_result *GetSecurityCheckResultBaseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecurityCheckResultBaseInfo"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecurityCheckResultBaseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the basic information of security check results.
//
// @param request - GetSecurityCheckResultBaseInfoRequest
//
// @return GetSecurityCheckResultBaseInfoResponse
func (client *Client) GetSecurityCheckResultBaseInfo(request *GetSecurityCheckResultBaseInfoRequest) (_result *GetSecurityCheckResultBaseInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSecurityCheckResultBaseInfoResponse{}
	_body, _err := client.GetSecurityCheckResultBaseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of security optimization suggestions.
//
// @param tmpReq - GetSecuritySuggestionListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecuritySuggestionListResponse
func (client *Client) GetSecuritySuggestionListWithOptions(tmpReq *GetSecuritySuggestionListRequest, runtime *dara.RuntimeOptions) (_result *GetSecuritySuggestionListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetSecuritySuggestionListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListConfigRulesRequest) {
		request.ListConfigRulesRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListConfigRulesRequest, dara.String("ListConfigRulesRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ListConfigRulesRequestShrink) {
		query["ListConfigRulesRequest"] = request.ListConfigRulesRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecuritySuggestionList"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecuritySuggestionListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of security optimization suggestions.
//
// @param request - GetSecuritySuggestionListRequest
//
// @return GetSecuritySuggestionListResponse
func (client *Client) GetSecuritySuggestionList(request *GetSecuritySuggestionListRequest) (_result *GetSecuritySuggestionListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSecuritySuggestionListResponse{}
	_body, _err := client.GetSecuritySuggestionListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the number of security optimization suggestions.
//
// @param request - GetSecuritySuggestionNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecuritySuggestionNumberResponse
func (client *Client) GetSecuritySuggestionNumberWithOptions(request *GetSecuritySuggestionNumberRequest, runtime *dara.RuntimeOptions) (_result *GetSecuritySuggestionNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecuritySuggestionNumber"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecuritySuggestionNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the number of security optimization suggestions.
//
// @param request - GetSecuritySuggestionNumberRequest
//
// @return GetSecuritySuggestionNumberResponse
func (client *Client) GetSecuritySuggestionNumber(request *GetSecuritySuggestionNumberRequest) (_result *GetSecuritySuggestionNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSecuritySuggestionNumberResponse{}
	_body, _err := client.GetSecuritySuggestionNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the status of a service-linked role.
//
// @param request - GetServiceLinkedRoleStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceLinkedRoleStatusResponse
func (client *Client) GetServiceLinkedRoleStatusWithOptions(request *GetServiceLinkedRoleStatusRequest, runtime *dara.RuntimeOptions) (_result *GetServiceLinkedRoleStatusResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceLinkedRoleStatus"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceLinkedRoleStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the status of a service-linked role.
//
// @param request - GetServiceLinkedRoleStatusRequest
//
// @return GetServiceLinkedRoleStatusResponse
func (client *Client) GetServiceLinkedRoleStatus(request *GetServiceLinkedRoleStatusRequest) (_result *GetServiceLinkedRoleStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceLinkedRoleStatusResponse{}
	_body, _err := client.GetServiceLinkedRoleStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves valid deductible instances for Security Center.
//
// @param tmpReq - GetValidDeductInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetValidDeductInstancesResponse
func (client *Client) GetValidDeductInstancesWithOptions(tmpReq *GetValidDeductInstancesRequest, runtime *dara.RuntimeOptions) (_result *GetValidDeductInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetValidDeductInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetValidDeductInstances"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetValidDeductInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves valid deductible instances for Security Center.
//
// @param request - GetValidDeductInstancesRequest
//
// @return GetValidDeductInstancesResponse
func (client *Client) GetValidDeductInstances(request *GetValidDeductInstancesRequest) (_result *GetValidDeductInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetValidDeductInstancesResponse{}
	_body, _err := client.GetValidDeductInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initializes module rules for Security Center.
//
// @param tmpReq - InitSasModuleRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitSasModuleRuleResponse
func (client *Client) InitSasModuleRuleWithOptions(tmpReq *InitSasModuleRuleRequest, runtime *dara.RuntimeOptions) (_result *InitSasModuleRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InitSasModuleRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Instances) {
		request.InstancesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Instances, dara.String("Instances"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoBind) {
		query["AutoBind"] = request.AutoBind
	}

	if !dara.IsNil(request.InstancesShrink) {
		query["Instances"] = request.InstancesShrink
	}

	if !dara.IsNil(request.IsTrial) {
		query["IsTrial"] = request.IsTrial
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitSasModuleRule"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitSasModuleRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initializes module rules for Security Center.
//
// @param request - InitSasModuleRuleRequest
//
// @return InitSasModuleRuleResponse
func (client *Client) InitSasModuleRule(request *InitSasModuleRuleRequest) (_result *InitSasModuleRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitSasModuleRuleResponse{}
	_body, _err := client.InitSasModuleRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the results of a full disk scan.
//
// @param request - ListVirusScanMachineEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVirusScanMachineEventResponse
func (client *Client) ListVirusScanMachineEventWithOptions(request *ListVirusScanMachineEventRequest, runtime *dara.RuntimeOptions) (_result *ListVirusScanMachineEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OperateTaskId) {
		query["OperateTaskId"] = request.OperateTaskId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVirusScanMachineEvent"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVirusScanMachineEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the results of a full disk scan.
//
// @param request - ListVirusScanMachineEventRequest
//
// @return ListVirusScanMachineEventResponse
func (client *Client) ListVirusScanMachineEvent(request *ListVirusScanMachineEventRequest) (_result *ListVirusScanMachineEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVirusScanMachineEventResponse{}
	_body, _err := client.ListVirusScanMachineEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a free trial of Security Center.
//
// @param request - OpenTrialPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenTrialPackageResponse
func (client *Client) OpenTrialPackageWithOptions(request *OpenTrialPackageRequest, runtime *dara.RuntimeOptions) (_result *OpenTrialPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoCloseSwitch) {
		query["AutoCloseSwitch"] = request.AutoCloseSwitch
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenTrialPackage"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenTrialPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a free trial of Security Center.
//
// @param request - OpenTrialPackageRequest
//
// @return OpenTrialPackageResponse
func (client *Client) OpenTrialPackage(request *OpenTrialPackageRequest) (_result *OpenTrialPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenTrialPackageResponse{}
	_body, _err := client.OpenTrialPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries account security events.
//
// @param request - QueryAccountSafetyIncidentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccountSafetyIncidentResponse
func (client *Client) QueryAccountSafetyIncidentWithOptions(request *QueryAccountSafetyIncidentRequest, runtime *dara.RuntimeOptions) (_result *QueryAccountSafetyIncidentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.CaseCode) {
		query["CaseCode"] = request.CaseCode
	}

	if !dara.IsNil(request.Current) {
		query["Current"] = request.Current
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PunishEndTime) {
		query["PunishEndTime"] = request.PunishEndTime
	}

	if !dara.IsNil(request.PunishStartTime) {
		query["PunishStartTime"] = request.PunishStartTime
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAccountSafetyIncident"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAccountSafetyIncidentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries account security events.
//
// @param request - QueryAccountSafetyIncidentRequest
//
// @return QueryAccountSafetyIncidentResponse
func (client *Client) QueryAccountSafetyIncident(request *QueryAccountSafetyIncidentRequest) (_result *QueryAccountSafetyIncidentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAccountSafetyIncidentResponse{}
	_body, _err := client.QueryAccountSafetyIncidentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the subscription status of the cloud security guide.
//
// @param request - QueryGuideSubStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGuideSubStatusResponse
func (client *Client) QueryGuideSubStatusWithOptions(request *QueryGuideSubStatusRequest, runtime *dara.RuntimeOptions) (_result *QueryGuideSubStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("QueryGuideSubStatus"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryGuideSubStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the subscription status of the cloud security guide.
//
// @param request - QueryGuideSubStatusRequest
//
// @return QueryGuideSubStatusResponse
func (client *Client) QueryGuideSubStatus(request *QueryGuideSubStatusRequest) (_result *QueryGuideSubStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryGuideSubStatusResponse{}
	_body, _err := client.QueryGuideSubStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries cloud resource control events.
//
// @param tmpReq - QueryResourceControlEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryResourceControlEventsResponse
func (client *Client) QueryResourceControlEventsWithOptions(tmpReq *QueryResourceControlEventsRequest, runtime *dara.RuntimeOptions) (_result *QueryResourceControlEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryResourceControlEventsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ActionCodes) {
		request.ActionCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ActionCodes, dara.String("ActionCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.BusinessCodes) {
		request.BusinessCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BusinessCodes, dara.String("BusinessCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CaseCodesPrefix) {
		request.CaseCodesPrefixShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CaseCodesPrefix, dara.String("CaseCodesPrefix"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.EventCodes) {
		request.EventCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventCodes, dara.String("EventCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.EventIdList) {
		request.EventIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventIdList, dara.String("EventIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExcludeActionCodes) {
		request.ExcludeActionCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeActionCodes, dara.String("ExcludeActionCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExcludeEventCodes) {
		request.ExcludeEventCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeEventCodes, dara.String("ExcludeEventCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExcludeReasons) {
		request.ExcludeReasonsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeReasons, dara.String("ExcludeReasons"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IncludeReasons) {
		request.IncludeReasonsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IncludeReasons, dara.String("IncludeReasons"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceCodes) {
		request.SourceCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceCodes, dara.String("SourceCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StatusList) {
		request.StatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StatusList, dara.String("StatusList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionCode) {
		query["ActionCode"] = request.ActionCode
	}

	if !dara.IsNil(request.ActionCodesShrink) {
		query["ActionCodes"] = request.ActionCodesShrink
	}

	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.BusinessCode) {
		query["BusinessCode"] = request.BusinessCode
	}

	if !dara.IsNil(request.BusinessCodesShrink) {
		query["BusinessCodes"] = request.BusinessCodesShrink
	}

	if !dara.IsNil(request.CaseCodesPrefixShrink) {
		query["CaseCodesPrefix"] = request.CaseCodesPrefixShrink
	}

	if !dara.IsNil(request.Current) {
		query["Current"] = request.Current
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.EventCode) {
		query["EventCode"] = request.EventCode
	}

	if !dara.IsNil(request.EventCodesShrink) {
		query["EventCodes"] = request.EventCodesShrink
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.EventIdListShrink) {
		query["EventIdList"] = request.EventIdListShrink
	}

	if !dara.IsNil(request.ExcludeActionCodesShrink) {
		query["ExcludeActionCodes"] = request.ExcludeActionCodesShrink
	}

	if !dara.IsNil(request.ExcludeEventCodesShrink) {
		query["ExcludeEventCodes"] = request.ExcludeEventCodesShrink
	}

	if !dara.IsNil(request.ExcludeReasonsShrink) {
		query["ExcludeReasons"] = request.ExcludeReasonsShrink
	}

	if !dara.IsNil(request.IncludeReasonsShrink) {
		query["IncludeReasons"] = request.IncludeReasonsShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PunishEndTime) {
		query["PunishEndTime"] = request.PunishEndTime
	}

	if !dara.IsNil(request.PunishStartTime) {
		query["PunishStartTime"] = request.PunishStartTime
	}

	if !dara.IsNil(request.Reason) {
		query["Reason"] = request.Reason
	}

	if !dara.IsNil(request.SourceCodesShrink) {
		query["SourceCodes"] = request.SourceCodesShrink
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StatusListShrink) {
		query["StatusList"] = request.StatusListShrink
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryResourceControlEvents"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryResourceControlEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries cloud resource control events.
//
// @param request - QueryResourceControlEventsRequest
//
// @return QueryResourceControlEventsResponse
func (client *Client) QueryResourceControlEvents(request *QueryResourceControlEventsRequest) (_result *QueryResourceControlEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryResourceControlEventsResponse{}
	_body, _err := client.QueryResourceControlEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the security check brief.
//
// @param request - QuerySecurityCheckReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySecurityCheckReportResponse
func (client *Client) QuerySecurityCheckReportWithOptions(request *QuerySecurityCheckReportRequest, runtime *dara.RuntimeOptions) (_result *QuerySecurityCheckReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySecurityCheckReport"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySecurityCheckReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the security check brief.
//
// @param request - QuerySecurityCheckReportRequest
//
// @return QuerySecurityCheckReportResponse
func (client *Client) QuerySecurityCheckReport(request *QuerySecurityCheckReportRequest) (_result *QuerySecurityCheckReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySecurityCheckReportResponse{}
	_body, _err := client.QuerySecurityCheckReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables disposal tool authorization.
//
// @param request - StartDisposalToolServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDisposalToolServiceResponse
func (client *Client) StartDisposalToolServiceWithOptions(request *StartDisposalToolServiceRequest, runtime *dara.RuntimeOptions) (_result *StartDisposalToolServiceResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDisposalToolService"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDisposalToolServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables disposal tool authorization.
//
// @param request - StartDisposalToolServiceRequest
//
// @return StartDisposalToolServiceResponse
func (client *Client) StartDisposalToolService(request *StartDisposalToolServiceRequest) (_result *StartDisposalToolServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartDisposalToolServiceResponse{}
	_body, _err := client.StartDisposalToolServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables security check.
//
// @param request - StartSecurityCheckServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSecurityCheckServiceResponse
func (client *Client) StartSecurityCheckServiceWithOptions(request *StartSecurityCheckServiceRequest, runtime *dara.RuntimeOptions) (_result *StartSecurityCheckServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("StartSecurityCheckService"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartSecurityCheckServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables security check.
//
// @param request - StartSecurityCheckServiceRequest
//
// @return StartSecurityCheckServiceResponse
func (client *Client) StartSecurityCheckService(request *StartSecurityCheckServiceRequest) (_result *StartSecurityCheckServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartSecurityCheckServiceResponse{}
	_body, _err := client.StartSecurityCheckServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits an unblocking request.
//
// @param tmpReq - SubmitApplyRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitApplyRecordResponse
func (client *Client) SubmitApplyRecordWithOptions(tmpReq *SubmitApplyRecordRequest, runtime *dara.RuntimeOptions) (_result *SubmitApplyRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitApplyRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EventIdList) {
		request.EventIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventIdList, dara.String("EventIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplyRequest) {
		query["ApplyRequest"] = request.ApplyRequest
	}

	if !dara.IsNil(request.CommitmentLetter) {
		query["CommitmentLetter"] = request.CommitmentLetter
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EventIdListShrink) {
		query["EventIdList"] = request.EventIdListShrink
	}

	if !dara.IsNil(request.QualificationProof) {
		query["QualificationProof"] = request.QualificationProof
	}

	if !dara.IsNil(request.Trial) {
		query["Trial"] = request.Trial
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitApplyRecord"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitApplyRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an unblocking request.
//
// @param request - SubmitApplyRecordRequest
//
// @return SubmitApplyRecordResponse
func (client *Client) SubmitApplyRecord(request *SubmitApplyRecordRequest) (_result *SubmitApplyRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitApplyRecordResponse{}
	_body, _err := client.SubmitApplyRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the pay-as-you-go binding relationship for Security Center.
//
// @param tmpReq - UpdatePostPaidBindRelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePostPaidBindRelResponse
func (client *Client) UpdatePostPaidBindRelWithOptions(tmpReq *UpdatePostPaidBindRelRequest, runtime *dara.RuntimeOptions) (_result *UpdatePostPaidBindRelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePostPaidBindRelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePostPaidBindRel"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePostPaidBindRelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the pay-as-you-go binding relationship for Security Center.
//
// @param request - UpdatePostPaidBindRelRequest
//
// @return UpdatePostPaidBindRelResponse
func (client *Client) UpdatePostPaidBindRel(request *UpdatePostPaidBindRelRequest) (_result *UpdatePostPaidBindRelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePostPaidBindRelResponse{}
	_body, _err := client.UpdatePostPaidBindRelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the health check result.
//
// @param request - UpdateSecurityCheckResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecurityCheckResultResponse
func (client *Client) UpdateSecurityCheckResultWithOptions(request *UpdateSecurityCheckResultRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecurityCheckResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSecurityCheckResult"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecurityCheckResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the health check result.
//
// @param request - UpdateSecurityCheckResultRequest
//
// @return UpdateSecurityCheckResultResponse
func (client *Client) UpdateSecurityCheckResult(request *UpdateSecurityCheckResultRequest) (_result *UpdateSecurityCheckResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSecurityCheckResultResponse{}
	_body, _err := client.UpdateSecurityCheckResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
