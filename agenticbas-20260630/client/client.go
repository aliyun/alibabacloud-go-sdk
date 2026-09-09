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
	client.Endpoint, _err = client.GetEndpoint(dara.String("agenticbas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建渗透测试任务
//
// @param tmpReq - CreatePentestTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePentestTaskResponse
func (client *Client) CreatePentestTaskWithOptions(tmpReq *CreatePentestTaskRequest, runtime *dara.RuntimeOptions) (_result *CreatePentestTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePentestTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OperationInput) {
		request.OperationInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationInput, dara.String("OperationInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationInputShrink) {
		query["OperationInput"] = request.OperationInputShrink
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePentestTask"),
		Version:     dara.String("2026-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePentestTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建渗透测试任务
//
// @param request - CreatePentestTaskRequest
//
// @return CreatePentestTaskResponse
func (client *Client) CreatePentestTask(request *CreatePentestTaskRequest) (_result *CreatePentestTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePentestTaskResponse{}
	_body, _err := client.CreatePentestTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询渗透测试报告内容
//
// @param tmpReq - DescribePentestReportContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePentestReportContentResponse
func (client *Client) DescribePentestReportContentWithOptions(tmpReq *DescribePentestReportContentRequest, runtime *dara.RuntimeOptions) (_result *DescribePentestReportContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribePentestReportContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OperationInput) {
		request.OperationInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationInput, dara.String("OperationInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationInputShrink) {
		query["OperationInput"] = request.OperationInputShrink
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePentestReportContent"),
		Version:     dara.String("2026-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePentestReportContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询渗透测试报告内容
//
// @param request - DescribePentestReportContentRequest
//
// @return DescribePentestReportContentResponse
func (client *Client) DescribePentestReportContent(request *DescribePentestReportContentRequest) (_result *DescribePentestReportContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePentestReportContentResponse{}
	_body, _err := client.DescribePentestReportContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询渗透测试任务列表
//
// @param tmpReq - DescribePentestTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePentestTaskListResponse
func (client *Client) DescribePentestTaskListWithOptions(tmpReq *DescribePentestTaskListRequest, runtime *dara.RuntimeOptions) (_result *DescribePentestTaskListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribePentestTaskListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OperationInput) {
		request.OperationInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationInput, dara.String("OperationInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationInputShrink) {
		query["OperationInput"] = request.OperationInputShrink
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePentestTaskList"),
		Version:     dara.String("2026-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePentestTaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询渗透测试任务列表
//
// @param request - DescribePentestTaskListRequest
//
// @return DescribePentestTaskListResponse
func (client *Client) DescribePentestTaskList(request *DescribePentestTaskListRequest) (_result *DescribePentestTaskListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePentestTaskListResponse{}
	_body, _err := client.DescribePentestTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询渗透测试漏洞列表
//
// @param tmpReq - DescribePentestVulnListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePentestVulnListResponse
func (client *Client) DescribePentestVulnListWithOptions(tmpReq *DescribePentestVulnListRequest, runtime *dara.RuntimeOptions) (_result *DescribePentestVulnListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribePentestVulnListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OperationInput) {
		request.OperationInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationInput, dara.String("OperationInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationInputShrink) {
		query["OperationInput"] = request.OperationInputShrink
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePentestVulnList"),
		Version:     dara.String("2026-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePentestVulnListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询渗透测试漏洞列表
//
// @param request - DescribePentestVulnListRequest
//
// @return DescribePentestVulnListResponse
func (client *Client) DescribePentestVulnList(request *DescribePentestVulnListRequest) (_result *DescribePentestVulnListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePentestVulnListResponse{}
	_body, _err := client.DescribePentestVulnListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
