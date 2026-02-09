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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("intelligentcreation"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 检查Turing任务
//
// @param request - CheckTuringTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckTuringTaskResponse
func (client *Client) CheckTuringTaskWithOptions(request *CheckTuringTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckTuringTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckTuringTask"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/aigc-turing/openService/v1/task/checkTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckTuringTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查Turing任务
//
// @param request - CheckTuringTaskRequest
//
// @return CheckTuringTaskResponse
func (client *Client) CheckTuringTask(request *CheckTuringTaskRequest) (_result *CheckTuringTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckTuringTaskResponse{}
	_body, _err := client.CheckTuringTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 任务提交接口
//
// @param request - SubmitTuringTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTuringTaskResponse
func (client *Client) SubmitTuringTaskWithOptions(request *SubmitTuringTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitTuringTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Duration) {
		body["duration"] = request.Duration
	}

	if !dara.IsNil(request.IdempotentKey) {
		body["idempotentKey"] = request.IdempotentKey
	}

	if !dara.IsNil(request.ImgUrl) {
		body["imgUrl"] = request.ImgUrl
	}

	if !dara.IsNil(request.Resolution) {
		body["resolution"] = request.Resolution
	}

	if !dara.IsNil(request.ResourceType) {
		body["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TaskType) {
		body["taskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitTuringTask"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/aigc-turing/openService/v1/task/submitTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitTuringTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 任务提交接口
//
// @param request - SubmitTuringTaskRequest
//
// @return SubmitTuringTaskResponse
func (client *Client) SubmitTuringTask(request *SubmitTuringTaskRequest) (_result *SubmitTuringTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitTuringTaskResponse{}
	_body, _err := client.SubmitTuringTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
