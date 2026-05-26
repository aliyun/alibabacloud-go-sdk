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
	client.Endpoint, _err = client.GetEndpoint(dara.String("aisearchengine"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// # AI搜索
//
// @param request - EngineSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EngineSearchResponse
func (client *Client) EngineSearchWithOptions(request *EngineSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EngineSearchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.Grey) {
		body["grey"] = request.Grey
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.Recall) {
		body["recall"] = request.Recall
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.StrategyId) {
		body["strategyId"] = request.StrategyId
	}

	if !dara.IsNil(request.User) {
		body["user"] = request.User
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EngineSearch"),
		Version:     dara.String("2026-04-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/platform/app/search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EngineSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AI搜索
//
// @param request - EngineSearchRequest
//
// @return EngineSearchResponse
func (client *Client) EngineSearch(request *EngineSearchRequest) (_result *EngineSearchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EngineSearchResponse{}
	_body, _err := client.EngineSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据集资源 OSS 访问地址
//
// @param request - GetDatasetResourceUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetResourceUrlResponse
func (client *Client) GetDatasetResourceUrlWithOptions(request *GetDatasetResourceUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatasetResourceUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		body["datasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.PrimaryKey) {
		body["primaryKey"] = request.PrimaryKey
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatasetResourceUrl"),
		Version:     dara.String("2026-04-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/dataset/open/resources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasetResourceUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据集资源 OSS 访问地址
//
// @param request - GetDatasetResourceUrlRequest
//
// @return GetDatasetResourceUrlResponse
func (client *Client) GetDatasetResourceUrl(request *GetDatasetResourceUrlRequest) (_result *GetDatasetResourceUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatasetResourceUrlResponse{}
	_body, _err := client.GetDatasetResourceUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交单条记录导入任务
//
// @param request - ImportDatasetDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportDatasetDataResponse
func (client *Client) ImportDatasetDataWithOptions(request *ImportDatasetDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ImportDatasetDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		body["datasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.Records) {
		body["records"] = request.Records
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportDatasetData"),
		Version:     dara.String("2026-04-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/dataset/open/upsert"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportDatasetDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交单条记录导入任务
//
// @param request - ImportDatasetDataRequest
//
// @return ImportDatasetDataResponse
func (client *Client) ImportDatasetData(request *ImportDatasetDataRequest) (_result *ImportDatasetDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ImportDatasetDataResponse{}
	_body, _err := client.ImportDatasetDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # AI问答对话
//
// @param request - QaChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QaChatResponse
func (client *Client) QaChatWithSSE(request *QaChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *QaChatResponse, _yieldErr chan error) {
	defer close(_yield)
	client.qaChatWithSSE_opYieldFunc(_yield, _yieldErr, request, headers, runtime)
	return
}

// Summary:
//
// # AI问答对话
//
// @param request - QaChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QaChatResponse
func (client *Client) QaChatWithOptions(request *QaChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QaChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.Message) {
		body["message"] = request.Message
	}

	if !dara.IsNil(request.Options) {
		body["options"] = request.Options
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QaChat"),
		Version:     dara.String("2026-04-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/platform/app/chat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QaChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AI问答对话
//
// @param request - QaChatRequest
//
// @return QaChatResponse
func (client *Client) QaChat(request *QaChatRequest) (_result *QaChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QaChatResponse{}
	_body, _err := client.QaChatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) qaChatWithSSE_opYieldFunc(_yield chan *QaChatResponse, _yieldErr chan error, request *QaChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.Message) {
		body["message"] = request.Message
	}

	if !dara.IsNil(request.Options) {
		body["options"] = request.Options
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QaChat"),
		Version:     dara.String("2026-04-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/platform/app/chat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}
