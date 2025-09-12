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
	client.Endpoint, _err = client.GetEndpoint(dara.String("umeng-finplus"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 提交数据集任务，校验数据集
//
// @param request - BuildStsAKRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BuildStsAKResponse
func (client *Client) BuildStsAKWithOptions(request *BuildStsAKRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BuildStsAKResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("BuildStsAK"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/buildStsAK"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BuildStsAKResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交数据集任务，校验数据集
//
// @param request - BuildStsAKRequest
//
// @return BuildStsAKResponse
func (client *Client) BuildStsAK(request *BuildStsAKRequest) (_result *BuildStsAKResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BuildStsAKResponse{}
	_body, _err := client.BuildStsAKWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交数据集任务，校验数据集
//
// @param request - BuildStsAK2Request
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BuildStsAK2Response
func (client *Client) BuildStsAK2WithOptions(request *BuildStsAK2Request, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BuildStsAK2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.DataSetId) {
		body["dataSetId"] = request.DataSetId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BuildStsAK2"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/buildStsAK2"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BuildStsAK2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交数据集任务，校验数据集
//
// @param request - BuildStsAK2Request
//
// @return BuildStsAK2Response
func (client *Client) BuildStsAK2(request *BuildStsAK2Request) (_result *BuildStsAK2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BuildStsAK2Response{}
	_body, _err := client.BuildStsAK2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消批量计算任务
//
// @param request - CancelTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelTaskResponse
func (client *Client) CancelTaskWithOptions(request *CancelTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelTask"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/cancelTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消批量计算任务
//
// @param request - CancelTaskRequest
//
// @return CancelTaskResponse
func (client *Client) CancelTask(request *CancelTaskRequest) (_result *CancelTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelTaskResponse{}
	_body, _err := client.CancelTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消批量计算任务
//
// @param request - CancelTask2Request
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelTask2Response
func (client *Client) CancelTask2WithOptions(request *CancelTask2Request, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelTask2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BcId) {
		body["bcId"] = request.BcId
	}

	if !dara.IsNil(request.ClientId) {
		body["clientId"] = request.ClientId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelTask2"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/cancelTask2"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelTask2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消批量计算任务
//
// @param request - CancelTask2Request
//
// @return CancelTask2Response
func (client *Client) CancelTask2(request *CancelTask2Request) (_result *CancelTask2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelTask2Response{}
	_body, _err := client.CancelTask2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建批量计算任务
//
// @param request - CreateComputeTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComputeTaskResponse
func (client *Client) CreateComputeTaskWithOptions(request *CreateComputeTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateComputeTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.DataSetIds) {
		body["dataSetIds"] = request.DataSetIds
	}

	if !dara.IsNil(request.MorseInfoList) {
		body["morseInfoList"] = request.MorseInfoList
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Remarks) {
		body["remarks"] = request.Remarks
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComputeTask"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/createComputeTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComputeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建批量计算任务
//
// @param request - CreateComputeTaskRequest
//
// @return CreateComputeTaskResponse
func (client *Client) CreateComputeTask(request *CreateComputeTaskRequest) (_result *CreateComputeTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateComputeTaskResponse{}
	_body, _err := client.CreateComputeTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建批量计算任务
//
// @param request - CreateComputeTask2Request
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComputeTask2Response
func (client *Client) CreateComputeTask2WithOptions(request *CreateComputeTask2Request, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateComputeTask2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.ClientId) {
		body["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.DataSetIds) {
		body["dataSetIds"] = request.DataSetIds
	}

	if !dara.IsNil(request.MorseInfoList) {
		body["morseInfoList"] = request.MorseInfoList
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Remarks) {
		body["remarks"] = request.Remarks
	}

	if !dara.IsNil(request.TaskSource) {
		body["taskSource"] = request.TaskSource
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComputeTask2"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/createComputeTask2"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComputeTask2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建批量计算任务
//
// @param request - CreateComputeTask2Request
//
// @return CreateComputeTask2Response
func (client *Client) CreateComputeTask2(request *CreateComputeTask2Request) (_result *CreateComputeTask2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateComputeTask2Response{}
	_body, _err := client.CreateComputeTask2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据集
//
// @param request - CreateDataSetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSetResponse
func (client *Client) CreateDataSetWithOptions(request *CreateDataSetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDataSetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataSet"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/createDataSet"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据集
//
// @param request - CreateDataSetRequest
//
// @return CreateDataSetResponse
func (client *Client) CreateDataSet(request *CreateDataSetRequest) (_result *CreateDataSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDataSetResponse{}
	_body, _err := client.CreateDataSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据集
//
// @param request - CreateDataSet2Request
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSet2Response
func (client *Client) CreateDataSet2WithOptions(request *CreateDataSet2Request, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDataSet2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataSet2"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/createDataSet2"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataSet2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据集
//
// @param request - CreateDataSet2Request
//
// @return CreateDataSet2Response
func (client *Client) CreateDataSet2(request *CreateDataSet2Request) (_result *CreateDataSet2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDataSet2Response{}
	_body, _err := client.CreateDataSet2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 友准达-创建知识库
//
// @param tmpReq - CreateKnowLedgeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKnowLedgeResponse
func (client *Client) CreateKnowLedgeWithOptions(tmpReq *CreateKnowLedgeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateKnowLedgeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateKnowLedgeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKnowLedge"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yzd/createKnowLedge"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKnowLedgeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 友准达-创建知识库
//
// @param request - CreateKnowLedgeRequest
//
// @return CreateKnowLedgeResponse
func (client *Client) CreateKnowLedge(request *CreateKnowLedgeRequest) (_result *CreateKnowLedgeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateKnowLedgeResponse{}
	_body, _err := client.CreateKnowLedgeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 加密调用OpenAPI
//
// @param request - EncryptInvokeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EncryptInvokeResponse
func (client *Client) EncryptInvokeWithOptions(request *EncryptInvokeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EncryptInvokeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.Data) {
		body["data"] = request.Data
	}

	if !dara.IsNil(request.EncryptKey) {
		body["encryptKey"] = request.EncryptKey
	}

	if !dara.IsNil(request.MethodName) {
		body["methodName"] = request.MethodName
	}

	if !dara.IsNil(request.Sign) {
		body["sign"] = request.Sign
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EncryptInvoke"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/encryptInvoke"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EncryptInvokeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 加密调用OpenAPI
//
// @param request - EncryptInvokeRequest
//
// @return EncryptInvokeResponse
func (client *Client) EncryptInvoke(request *EncryptInvokeRequest) (_result *EncryptInvokeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EncryptInvokeResponse{}
	_body, _err := client.EncryptInvokeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取人群集信息
//
// @param tmpReq - GetCrowdDatasetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCrowdDatasetResponse
func (client *Client) GetCrowdDatasetWithOptions(tmpReq *GetCrowdDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCrowdDatasetResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetCrowdDatasetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		query["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCrowdDataset"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yzd/getCrowdDataset"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCrowdDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取人群集信息
//
// @param request - GetCrowdDatasetRequest
//
// @return GetCrowdDatasetResponse
func (client *Client) GetCrowdDataset(request *GetCrowdDatasetRequest) (_result *GetCrowdDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCrowdDatasetResponse{}
	_body, _err := client.GetCrowdDatasetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取知识库数据数据
//
// @param tmpReq - GetKnowledgeDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKnowledgeDataResponse
func (client *Client) GetKnowledgeDataWithOptions(tmpReq *GetKnowledgeDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetKnowledgeDataResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetKnowledgeDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		query["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKnowledgeData"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yzd/getKnowledgeData"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKnowledgeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取知识库数据数据
//
// @param request - GetKnowledgeDataRequest
//
// @return GetKnowledgeDataResponse
func (client *Client) GetKnowledgeData(request *GetKnowledgeDataRequest) (_result *GetKnowledgeDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetKnowledgeDataResponse{}
	_body, _err := client.GetKnowledgeDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例结果
//
// @param tmpReq - GetYzdInstanceTaskResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYzdInstanceTaskResultResponse
func (client *Client) GetYzdInstanceTaskResultWithOptions(tmpReq *GetYzdInstanceTaskResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetYzdInstanceTaskResultResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetYzdInstanceTaskResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		query["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetYzdInstanceTaskResult"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yzd/getYzdInstanceTaskResult"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetYzdInstanceTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例结果
//
// @param request - GetYzdInstanceTaskResultRequest
//
// @return GetYzdInstanceTaskResultResponse
func (client *Client) GetYzdInstanceTaskResult(request *GetYzdInstanceTaskResultRequest) (_result *GetYzdInstanceTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetYzdInstanceTaskResultResponse{}
	_body, _err := client.GetYzdInstanceTaskResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 友准达知识库获取上传OSS数据文件的TOKEN
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYzdStsAKResponse
func (client *Client) GetYzdStsAKWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetYzdStsAKResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetYzdStsAK"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yzd/getYzdStsAK"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetYzdStsAKResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 友准达知识库获取上传OSS数据文件的TOKEN
//
// @return GetYzdStsAKResponse
func (client *Client) GetYzdStsAK() (_result *GetYzdStsAKResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetYzdStsAKResponse{}
	_body, _err := client.GetYzdStsAKWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个批量任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeTaskResponse
func (client *Client) ListComputeTaskWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListComputeTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComputeTask"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/listComputeTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComputeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个批量任务
//
// @return ListComputeTaskResponse
func (client *Client) ListComputeTask() (_result *ListComputeTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListComputeTaskResponse{}
	_body, _err := client.ListComputeTaskWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个批量任务
//
// @param request - ListComputeTask2Request
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeTask2Response
func (client *Client) ListComputeTask2WithOptions(request *ListComputeTask2Request, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListComputeTask2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.PageNum) {
		body["pageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComputeTask2"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/listComputeTask2"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComputeTask2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个批量任务
//
// @param request - ListComputeTask2Request
//
// @return ListComputeTask2Response
func (client *Client) ListComputeTask2(request *ListComputeTask2Request) (_result *ListComputeTask2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListComputeTask2Response{}
	_body, _err := client.ListComputeTask2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个数据集
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSetResponse
func (client *Client) ListDataSetWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataSetResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSet"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/listDataSet"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个数据集
//
// @return ListDataSetResponse
func (client *Client) ListDataSet() (_result *ListDataSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSetResponse{}
	_body, _err := client.ListDataSetWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个数据集
//
// @param request - ListDataSet2Request
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSet2Response
func (client *Client) ListDataSet2WithOptions(request *ListDataSet2Request, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataSet2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.PageNo) {
		body["pageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSet2"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/listDataSet2"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSet2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个数据集
//
// @param request - ListDataSet2Request
//
// @return ListDataSet2Response
func (client *Client) ListDataSet2(request *ListDataSet2Request) (_result *ListDataSet2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSet2Response{}
	_body, _err := client.ListDataSet2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据集
//
// @param request - RemoveDataSetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDataSetResponse
func (client *Client) RemoveDataSetWithOptions(request *RemoveDataSetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveDataSetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveDataSet"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/removeDataSet"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据集
//
// @param request - RemoveDataSetRequest
//
// @return RemoveDataSetResponse
func (client *Client) RemoveDataSet(request *RemoveDataSetRequest) (_result *RemoveDataSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveDataSetResponse{}
	_body, _err := client.RemoveDataSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据集
//
// @param request - RemoveDataSet2Request
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDataSet2Response
func (client *Client) RemoveDataSet2WithOptions(request *RemoveDataSet2Request, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveDataSet2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.DataSetId) {
		body["dataSetId"] = request.DataSetId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveDataSet2"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/removeDataSet2"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveDataSet2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据集
//
// @param request - RemoveDataSet2Request
//
// @return RemoveDataSet2Response
func (client *Client) RemoveDataSet2(request *RemoveDataSet2Request) (_result *RemoveDataSet2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveDataSet2Response{}
	_body, _err := client.RemoveDataSet2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建人群集并绑定数据集
//
// @param tmpReq - SaveCrowdDatasetAndBindingDatasetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveCrowdDatasetAndBindingDatasetResponse
func (client *Client) SaveCrowdDatasetAndBindingDatasetWithOptions(tmpReq *SaveCrowdDatasetAndBindingDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SaveCrowdDatasetAndBindingDatasetResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SaveCrowdDatasetAndBindingDatasetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveCrowdDatasetAndBindingDataset"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yzd/saveCrowdDatasetAndBindingDataset"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveCrowdDatasetAndBindingDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建人群集并绑定数据集
//
// @param request - SaveCrowdDatasetAndBindingDatasetRequest
//
// @return SaveCrowdDatasetAndBindingDatasetResponse
func (client *Client) SaveCrowdDatasetAndBindingDataset(request *SaveCrowdDatasetAndBindingDatasetRequest) (_result *SaveCrowdDatasetAndBindingDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveCrowdDatasetAndBindingDatasetResponse{}
	_body, _err := client.SaveCrowdDatasetAndBindingDatasetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个批量任务
//
// @param request - SelectComputeTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SelectComputeTaskResponse
func (client *Client) SelectComputeTaskWithOptions(request *SelectComputeTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SelectComputeTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("SelectComputeTask"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/selectComputeTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SelectComputeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个批量任务
//
// @param request - SelectComputeTaskRequest
//
// @return SelectComputeTaskResponse
func (client *Client) SelectComputeTask(request *SelectComputeTaskRequest) (_result *SelectComputeTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SelectComputeTaskResponse{}
	_body, _err := client.SelectComputeTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个批量任务
//
// @param request - SelectComputeTask2Request
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SelectComputeTask2Response
func (client *Client) SelectComputeTask2WithOptions(request *SelectComputeTask2Request, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SelectComputeTask2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BcConfirm) {
		body["bcConfirm"] = request.BcConfirm
	}

	if !dara.IsNil(request.BcId) {
		body["bcId"] = request.BcId
	}

	if !dara.IsNil(request.ClientId) {
		body["clientId"] = request.ClientId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SelectComputeTask2"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/selectComputeTask2"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SelectComputeTask2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个批量任务
//
// @param request - SelectComputeTask2Request
//
// @return SelectComputeTask2Response
func (client *Client) SelectComputeTask2(request *SelectComputeTask2Request) (_result *SelectComputeTask2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SelectComputeTask2Response{}
	_body, _err := client.SelectComputeTask2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个数据集
//
// @param request - SelectDataSetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SelectDataSetResponse
func (client *Client) SelectDataSetWithOptions(request *SelectDataSetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SelectDataSetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("SelectDataSet"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/selectDataSet"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SelectDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个数据集
//
// @param request - SelectDataSetRequest
//
// @return SelectDataSetResponse
func (client *Client) SelectDataSet(request *SelectDataSetRequest) (_result *SelectDataSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SelectDataSetResponse{}
	_body, _err := client.SelectDataSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个数据集
//
// @param request - SelectDataSet2Request
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SelectDataSet2Response
func (client *Client) SelectDataSet2WithOptions(request *SelectDataSet2Request, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SelectDataSet2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.DataSetId) {
		body["dataSetId"] = request.DataSetId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SelectDataSet2"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/selectDataSet2"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SelectDataSet2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个数据集
//
// @param request - SelectDataSet2Request
//
// @return SelectDataSet2Response
func (client *Client) SelectDataSet2(request *SelectDataSet2Request) (_result *SelectDataSet2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SelectDataSet2Response{}
	_body, _err := client.SelectDataSet2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交数据集任务，校验数据集
//
// @param request - SubmitDataSetTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDataSetTaskResponse
func (client *Client) SubmitDataSetTaskWithOptions(request *SubmitDataSetTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitDataSetTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDataSetTask"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/submitDataSetTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDataSetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交数据集任务，校验数据集
//
// @param request - SubmitDataSetTaskRequest
//
// @return SubmitDataSetTaskResponse
func (client *Client) SubmitDataSetTask(request *SubmitDataSetTaskRequest) (_result *SubmitDataSetTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitDataSetTaskResponse{}
	_body, _err := client.SubmitDataSetTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交数据集任务，校验数据集
//
// @param request - SubmitDataSetTask2Request
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDataSetTask2Response
func (client *Client) SubmitDataSetTask2WithOptions(request *SubmitDataSetTask2Request, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitDataSetTask2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.DataSetId) {
		body["dataSetId"] = request.DataSetId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDataSetTask2"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bc/submitDataSetTask2"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDataSetTask2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交数据集任务，校验数据集
//
// @param request - SubmitDataSetTask2Request
//
// @return SubmitDataSetTask2Response
func (client *Client) SubmitDataSetTask2(request *SubmitDataSetTask2Request) (_result *SubmitDataSetTask2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitDataSetTask2Response{}
	_body, _err := client.SubmitDataSetTask2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交知识库校验任务
//
// @param tmpReq - ValidateKnowLedgeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateKnowLedgeResponse
func (client *Client) ValidateKnowLedgeWithOptions(tmpReq *ValidateKnowLedgeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateKnowLedgeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ValidateKnowLedgeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateKnowLedge"),
		Version:     dara.String("2022-09-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yzd/validateKnowLedge"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateKnowLedgeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交知识库校验任务
//
// @param request - ValidateKnowLedgeRequest
//
// @return ValidateKnowLedgeResponse
func (client *Client) ValidateKnowLedge(request *ValidateKnowLedgeRequest) (_result *ValidateKnowLedgeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateKnowLedgeResponse{}
	_body, _err := client.ValidateKnowLedgeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
