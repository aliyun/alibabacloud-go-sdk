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
	client.Endpoint, _err = client.GetEndpoint(dara.String("amqp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 路由绑定
//
// @param request - BindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindResponse
func (client *Client) BindWithOptions(request *BindRequest, runtime *dara.RuntimeOptions) (_result *BindResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Argument) {
		query["Argument"] = request.Argument
	}

	if !dara.IsNil(request.BindingKey) {
		query["BindingKey"] = request.BindingKey
	}

	if !dara.IsNil(request.BindingType) {
		query["BindingType"] = request.BindingType
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.DstName) {
		query["DstName"] = request.DstName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SrcName) {
		query["SrcName"] = request.SrcName
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Bind"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 路由绑定
//
// @param request - BindRequest
//
// @return BindResponse
func (client *Client) Bind(request *BindRequest) (_result *BindResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindResponse{}
	_body, _err := client.BindWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除用户配置
//
// @param request - CancelUserSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelUserSettingResponse
func (client *Client) CancelUserSettingWithOptions(request *CancelUserSettingRequest, runtime *dara.RuntimeOptions) (_result *CancelUserSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelUserSetting"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelUserSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除用户配置
//
// @param request - CancelUserSettingRequest
//
// @return CancelUserSettingResponse
func (client *Client) CancelUserSetting(request *CancelUserSettingRequest) (_result *CancelUserSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelUserSettingResponse{}
	_body, _err := client.CancelUserSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增用户配置
//
// @param request - ConfigureUserSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureUserSettingResponse
func (client *Client) ConfigureUserSettingWithOptions(request *ConfigureUserSettingRequest, runtime *dara.RuntimeOptions) (_result *ConfigureUserSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Logstore) {
		query["Logstore"] = request.Logstore
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.PutType) {
		query["PutType"] = request.PutType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigureUserSetting"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigureUserSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增用户配置
//
// @param request - ConfigureUserSettingRequest
//
// @return ConfigureUserSettingResponse
func (client *Client) ConfigureUserSetting(request *ConfigureUserSettingRequest) (_result *ConfigureUserSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigureUserSettingResponse{}
	_body, _err := client.ConfigureUserSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 清除售后视角状态
//
// @param request - ConsoleClearPretendStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConsoleClearPretendStatusResponse
func (client *Client) ConsoleClearPretendStatusWithOptions(request *ConsoleClearPretendStatusRequest, runtime *dara.RuntimeOptions) (_result *ConsoleClearPretendStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConsoleClearPretendStatus"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConsoleClearPretendStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 清除售后视角状态
//
// @param request - ConsoleClearPretendStatusRequest
//
// @return ConsoleClearPretendStatusResponse
func (client *Client) ConsoleClearPretendStatus(request *ConsoleClearPretendStatusRequest) (_result *ConsoleClearPretendStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConsoleClearPretendStatusResponse{}
	_body, _err := client.ConsoleClearPretendStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存售后视角状态
//
// @param request - ConsoleSavePretendStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConsoleSavePretendStatusResponse
func (client *Client) ConsoleSavePretendStatusWithOptions(request *ConsoleSavePretendStatusRequest, runtime *dara.RuntimeOptions) (_result *ConsoleSavePretendStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConsoleSavePretendStatus"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConsoleSavePretendStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存售后视角状态
//
// @param request - ConsoleSavePretendStatusRequest
//
// @return ConsoleSavePretendStatusResponse
func (client *Client) ConsoleSavePretendStatus(request *ConsoleSavePretendStatusRequest) (_result *ConsoleSavePretendStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConsoleSavePretendStatusResponse{}
	_body, _err := client.ConsoleSavePretendStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建云监控相关角色
//
// @param request - CreateCloudMonitorSLRRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudMonitorSLRResponse
func (client *Client) CreateCloudMonitorSLRWithOptions(request *CreateCloudMonitorSLRRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudMonitorSLRResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudMonitorSLR"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudMonitorSLRResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建云监控相关角色
//
// @param request - CreateCloudMonitorSLRRequest
//
// @return CreateCloudMonitorSLRResponse
func (client *Client) CreateCloudMonitorSLR(request *CreateCloudMonitorSLRRequest) (_result *CreateCloudMonitorSLRResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudMonitorSLRResponse{}
	_body, _err := client.CreateCloudMonitorSLRWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Exchange
//
// @param request - CreateExchangeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExchangeResponse
func (client *Client) CreateExchangeWithOptions(request *CreateExchangeRequest, runtime *dara.RuntimeOptions) (_result *CreateExchangeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlternateExchange) {
		query["AlternateExchange"] = request.AlternateExchange
	}

	if !dara.IsNil(request.AutoDelete) {
		query["AutoDelete"] = request.AutoDelete
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.ExchangeName) {
		query["ExchangeName"] = request.ExchangeName
	}

	if !dara.IsNil(request.ExchangeType) {
		query["ExchangeType"] = request.ExchangeType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Internal) {
		query["Internal"] = request.Internal
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	if !dara.IsNil(request.XDelayedType) {
		query["XDelayedType"] = request.XDelayedType
	}

	if !dara.IsNil(request.XHashHeader) {
		query["XHashHeader"] = request.XHashHeader
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExchange"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExchangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Exchange
//
// @param request - CreateExchangeRequest
//
// @return CreateExchangeResponse
func (client *Client) CreateExchange(request *CreateExchangeRequest) (_result *CreateExchangeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateExchangeResponse{}
	_body, _err := client.CreateExchangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建日志相关角色
//
// @param request - CreateLogDeliverySLRRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLogDeliverySLRResponse
func (client *Client) CreateLogDeliverySLRWithOptions(request *CreateLogDeliverySLRRequest, runtime *dara.RuntimeOptions) (_result *CreateLogDeliverySLRResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLogDeliverySLR"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLogDeliverySLRResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建日志相关角色
//
// @param request - CreateLogDeliverySLRRequest
//
// @return CreateLogDeliverySLRResponse
func (client *Client) CreateLogDeliverySLR(request *CreateLogDeliverySLRRequest) (_result *CreateLogDeliverySLRResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLogDeliverySLRResponse{}
	_body, _err := client.CreateLogDeliverySLRWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建队列
//
// @param request - CreateQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQueueResponse
func (client *Client) CreateQueueWithOptions(request *CreateQueueRequest, runtime *dara.RuntimeOptions) (_result *CreateQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoDelete) {
		query["AutoDelete"] = request.AutoDelete
	}

	if !dara.IsNil(request.AutoExpire) {
		query["AutoExpire"] = request.AutoExpire
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.DeadLetterExchange) {
		query["DeadLetterExchange"] = request.DeadLetterExchange
	}

	if !dara.IsNil(request.DeadLetterRoutingKey) {
		query["DeadLetterRoutingKey"] = request.DeadLetterRoutingKey
	}

	if !dara.IsNil(request.Exclusive) {
		query["Exclusive"] = request.Exclusive
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxLength) {
		query["MaxLength"] = request.MaxLength
	}

	if !dara.IsNil(request.MaximunPrioty) {
		query["MaximunPrioty"] = request.MaximunPrioty
	}

	if !dara.IsNil(request.MessageTTL) {
		query["MessageTTL"] = request.MessageTTL
	}

	if !dara.IsNil(request.Ordered) {
		query["Ordered"] = request.Ordered
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.RetryInherit) {
		query["RetryInherit"] = request.RetryInherit
	}

	if !dara.IsNil(request.RetryInterval) {
		query["RetryInterval"] = request.RetryInterval
	}

	if !dara.IsNil(request.RetryTimes) {
		query["RetryTimes"] = request.RetryTimes
	}

	if !dara.IsNil(request.SingleActiveConsumer) {
		query["SingleActiveConsumer"] = request.SingleActiveConsumer
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQueue"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建队列
//
// @param request - CreateQueueRequest
//
// @return CreateQueueResponse
func (client *Client) CreateQueue(request *CreateQueueRequest) (_result *CreateQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateQueueResponse{}
	_body, _err := client.CreateQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Vhost
//
// @param request - CreateVhostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVhostResponse
func (client *Client) CreateVhostWithOptions(request *CreateVhostRequest, runtime *dara.RuntimeOptions) (_result *CreateVhostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVhost"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVhostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Vhost
//
// @param request - CreateVhostRequest
//
// @return CreateVhostResponse
func (client *Client) CreateVhost(request *CreateVhostRequest) (_result *CreateVhostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVhostResponse{}
	_body, _err := client.CreateVhostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// prometheus Dashboard 检查相关服务开通状态
//
// @param request - DashboardCheckServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DashboardCheckServiceStatusResponse
func (client *Client) DashboardCheckServiceStatusWithOptions(request *DashboardCheckServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *DashboardCheckServiceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DashboardCheckServiceStatus"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DashboardCheckServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// prometheus Dashboard 检查相关服务开通状态
//
// @param request - DashboardCheckServiceStatusRequest
//
// @return DashboardCheckServiceStatusResponse
func (client *Client) DashboardCheckServiceStatus(request *DashboardCheckServiceStatusRequest) (_result *DashboardCheckServiceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DashboardCheckServiceStatusResponse{}
	_body, _err := client.DashboardCheckServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取 arms grafana 大盘 http url
//
// @param request - DashboardListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DashboardListResponse
func (client *Client) DashboardListWithOptions(request *DashboardListRequest, runtime *dara.RuntimeOptions) (_result *DashboardListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.DashboardName) {
		query["DashboardName"] = request.DashboardName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DashboardList"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DashboardListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取 arms grafana 大盘 http url
//
// @param request - DashboardListRequest
//
// @return DashboardListResponse
func (client *Client) DashboardList(request *DashboardListRequest) (_result *DashboardListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DashboardListResponse{}
	_body, _err := client.DashboardListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Exchange
//
// @param tmpReq - DeleteExchangeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExchangeResponse
func (client *Client) DeleteExchangeWithOptions(tmpReq *DeleteExchangeRequest, runtime *dara.RuntimeOptions) (_result *DeleteExchangeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteExchangeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExchangeNames) {
		request.ExchangeNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExchangeNames, dara.String("ExchangeNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Collina) {
		query["Collina"] = request.Collina
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.ExchangeName) {
		query["ExchangeName"] = request.ExchangeName
	}

	if !dara.IsNil(request.ExchangeNamesShrink) {
		query["ExchangeNames"] = request.ExchangeNamesShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UmidToken) {
		query["UmidToken"] = request.UmidToken
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExchange"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExchangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Exchange
//
// @param request - DeleteExchangeRequest
//
// @return DeleteExchangeResponse
func (client *Client) DeleteExchange(request *DeleteExchangeRequest) (_result *DeleteExchangeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteExchangeResponse{}
	_body, _err := client.DeleteExchangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除实例
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除实例
//
// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除队列
//
// @param tmpReq - DeleteQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQueueResponse
func (client *Client) DeleteQueueWithOptions(tmpReq *DeleteQueueRequest, runtime *dara.RuntimeOptions) (_result *DeleteQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteQueueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.QueueNames) {
		request.QueueNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueueNames, dara.String("QueueNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Collina) {
		query["Collina"] = request.Collina
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.QueueNamesShrink) {
		query["QueueNames"] = request.QueueNamesShrink
	}

	if !dara.IsNil(request.UmidToken) {
		query["UmidToken"] = request.UmidToken
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQueue"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除队列
//
// @param request - DeleteQueueRequest
//
// @return DeleteQueueResponse
func (client *Client) DeleteQueue(request *DeleteQueueRequest) (_result *DeleteQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteQueueResponse{}
	_body, _err := client.DeleteQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除静态账户
//
// @param request - DeleteStaticAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStaticAccountResponse
func (client *Client) DeleteStaticAccountWithOptions(request *DeleteStaticAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteStaticAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CreateTimeStamp) {
		query["CreateTimeStamp"] = request.CreateTimeStamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStaticAccount"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStaticAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除静态账户
//
// @param request - DeleteStaticAccountRequest
//
// @return DeleteStaticAccountResponse
func (client *Client) DeleteStaticAccount(request *DeleteStaticAccountRequest) (_result *DeleteStaticAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStaticAccountResponse{}
	_body, _err := client.DeleteStaticAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Vhost
//
// @param tmpReq - DeleteVhostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVhostResponse
func (client *Client) DeleteVhostWithOptions(tmpReq *DeleteVhostRequest, runtime *dara.RuntimeOptions) (_result *DeleteVhostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteVhostShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.VhostNames) {
		request.VhostNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VhostNames, dara.String("VhostNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	if !dara.IsNil(request.VhostNamesShrink) {
		query["VhostNames"] = request.VhostNamesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVhost"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVhostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Vhost
//
// @param request - DeleteVhostRequest
//
// @return DeleteVhostResponse
func (client *Client) DeleteVhost(request *DeleteVhostRequest) (_result *DeleteVhostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVhostResponse{}
	_body, _err := client.DeleteVhostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出元数据
//
// @param request - ExportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportResponse
func (client *Client) ExportWithOptions(request *ExportRequest, runtime *dara.RuntimeOptions) (_result *ExportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.ExportType) {
		query["ExportType"] = request.ExportType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Export"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出元数据
//
// @param request - ExportRequest
//
// @return ExportResponse
func (client *Client) Export(request *ExportRequest) (_result *ExportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportResponse{}
	_body, _err := client.ExportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新静态账户
//
// @param request - FetchStaticAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchStaticAccountResponse
func (client *Client) FetchStaticAccountWithOptions(request *FetchStaticAccountRequest, runtime *dara.RuntimeOptions) (_result *FetchStaticAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountAccessKey) {
		query["AccountAccessKey"] = request.AccountAccessKey
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CreateTimeStamp) {
		query["CreateTimeStamp"] = request.CreateTimeStamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SKey) {
		query["SKey"] = request.SKey
	}

	if !dara.IsNil(request.SecretSign) {
		query["SecretSign"] = request.SecretSign
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FetchStaticAccount"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FetchStaticAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新静态账户
//
// @param request - FetchStaticAccountRequest
//
// @return FetchStaticAccountResponse
func (client *Client) FetchStaticAccount(request *FetchStaticAccountRequest) (_result *FetchStaticAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FetchStaticAccountResponse{}
	_body, _err := client.FetchStaticAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据耗时查询ack信息
//
// @param request - GetAckInfoByIntervalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAckInfoByIntervalResponse
func (client *Client) GetAckInfoByIntervalWithOptions(request *GetAckInfoByIntervalRequest, runtime *dara.RuntimeOptions) (_result *GetAckInfoByIntervalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntervalSec) {
		query["IntervalSec"] = request.IntervalSec
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAckInfoByInterval"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAckInfoByIntervalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据耗时查询ack信息
//
// @param request - GetAckInfoByIntervalRequest
//
// @return GetAckInfoByIntervalResponse
func (client *Client) GetAckInfoByInterval(request *GetAckInfoByIntervalRequest) (_result *GetAckInfoByIntervalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAckInfoByIntervalResponse{}
	_body, _err := client.GetAckInfoByIntervalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一个PushMessage（PullMessage）对应的Ack行为
//
// @param request - GetAckInfoOfMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAckInfoOfMessageResponse
func (client *Client) GetAckInfoOfMessageWithOptions(request *GetAckInfoOfMessageRequest, runtime *dara.RuntimeOptions) (_result *GetAckInfoOfMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.ConsumeStatus) {
		query["ConsumeStatus"] = request.ConsumeStatus
	}

	if !dara.IsNil(request.DeliveryTag) {
		query["DeliveryTag"] = request.DeliveryTag
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MsgId) {
		query["MsgId"] = request.MsgId
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TimeStamp) {
		query["TimeStamp"] = request.TimeStamp
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAckInfoOfMessage"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAckInfoOfMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一个PushMessage（PullMessage）对应的Ack行为
//
// @param request - GetAckInfoOfMessageRequest
//
// @return GetAckInfoOfMessageResponse
func (client *Client) GetAckInfoOfMessage(request *GetAckInfoOfMessageRequest) (_result *GetAckInfoOfMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAckInfoOfMessageResponse{}
	_body, _err := client.GetAckInfoOfMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取绑定数量
//
// @param request - GetBindingCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBindingCountResponse
func (client *Client) GetBindingCountWithOptions(request *GetBindingCountRequest, runtime *dara.RuntimeOptions) (_result *GetBindingCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BindingType) {
		query["BindingType"] = request.BindingType
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.Upstream) {
		query["Upstream"] = request.Upstream
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBindingCount"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBindingCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取绑定数量
//
// @param request - GetBindingCountRequest
//
// @return GetBindingCountResponse
func (client *Client) GetBindingCount(request *GetBindingCountRequest) (_result *GetBindingCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBindingCountResponse{}
	_body, _err := client.GetBindingCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取绑定错误
//
// @param request - GetBindingErrorByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBindingErrorByTaskIdResponse
func (client *Client) GetBindingErrorByTaskIdWithOptions(request *GetBindingErrorByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetBindingErrorByTaskIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBindingErrorByTaskId"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBindingErrorByTaskIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取绑定错误
//
// @param request - GetBindingErrorByTaskIdRequest
//
// @return GetBindingErrorByTaskIdResponse
func (client *Client) GetBindingErrorByTaskId(request *GetBindingErrorByTaskIdRequest) (_result *GetBindingErrorByTaskIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBindingErrorByTaskIdResponse{}
	_body, _err := client.GetBindingErrorByTaskIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetCommonBuyUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCommonBuyUrlResponse
func (client *Client) GetCommonBuyUrlWithOptions(request *GetCommonBuyUrlRequest, runtime *dara.RuntimeOptions) (_result *GetCommonBuyUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionType) {
		query["ActionType"] = request.ActionType
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCommonBuyUrl"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCommonBuyUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetCommonBuyUrlRequest
//
// @return GetCommonBuyUrlResponse
func (client *Client) GetCommonBuyUrl(request *GetCommonBuyUrlRequest) (_result *GetCommonBuyUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCommonBuyUrlResponse{}
	_body, _err := client.GetCommonBuyUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过rocketMqMsgId查询消息消费轨迹
//
// @param request - GetConsumeTraceByQueueAndRocketMqMsgIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumeTraceByQueueAndRocketMqMsgIdResponse
func (client *Client) GetConsumeTraceByQueueAndRocketMqMsgIdWithOptions(request *GetConsumeTraceByQueueAndRocketMqMsgIdRequest, runtime *dara.RuntimeOptions) (_result *GetConsumeTraceByQueueAndRocketMqMsgIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MsgId) {
		query["MsgId"] = request.MsgId
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConsumeTraceByQueueAndRocketMqMsgId"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConsumeTraceByQueueAndRocketMqMsgIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过rocketMqMsgId查询消息消费轨迹
//
// @param request - GetConsumeTraceByQueueAndRocketMqMsgIdRequest
//
// @return GetConsumeTraceByQueueAndRocketMqMsgIdResponse
func (client *Client) GetConsumeTraceByQueueAndRocketMqMsgId(request *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) (_result *GetConsumeTraceByQueueAndRocketMqMsgIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetConsumeTraceByQueueAndRocketMqMsgIdResponse{}
	_body, _err := client.GetConsumeTraceByQueueAndRocketMqMsgIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Exchange错误
//
// @param request - GetExchangeErrorByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExchangeErrorByTaskIdResponse
func (client *Client) GetExchangeErrorByTaskIdWithOptions(request *GetExchangeErrorByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetExchangeErrorByTaskIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExchangeErrorByTaskId"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExchangeErrorByTaskIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Exchange错误
//
// @param request - GetExchangeErrorByTaskIdRequest
//
// @return GetExchangeErrorByTaskIdResponse
func (client *Client) GetExchangeErrorByTaskId(request *GetExchangeErrorByTaskIdRequest) (_result *GetExchangeErrorByTaskIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetExchangeErrorByTaskIdResponse{}
	_body, _err := client.GetExchangeErrorByTaskIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Exchange Rate
//
// @param tmpReq - GetExchangeRateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExchangeRateResponse
func (client *Client) GetExchangeRateWithOptions(tmpReq *GetExchangeRateRequest, runtime *dara.RuntimeOptions) (_result *GetExchangeRateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetExchangeRateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExchangeNames) {
		request.ExchangeNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExchangeNames, dara.String("ExchangeNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.ExchangeNamesShrink) {
		query["ExchangeNames"] = request.ExchangeNamesShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExchangeRate"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExchangeRateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Exchange Rate
//
// @param request - GetExchangeRateRequest
//
// @return GetExchangeRateResponse
func (client *Client) GetExchangeRate(request *GetExchangeRateRequest) (_result *GetExchangeRateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetExchangeRateResponse{}
	_body, _err := client.GetExchangeRateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过queueName查询一段时间内的消息id列表
//
// @param request - GetMsgIdListByQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMsgIdListByQueueResponse
func (client *Client) GetMsgIdListByQueueWithOptions(request *GetMsgIdListByQueueRequest, runtime *dara.RuntimeOptions) (_result *GetMsgIdListByQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMsgIdListByQueue"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMsgIdListByQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过queueName查询一段时间内的消息id列表
//
// @param request - GetMsgIdListByQueueRequest
//
// @return GetMsgIdListByQueueResponse
func (client *Client) GetMsgIdListByQueue(request *GetMsgIdListByQueueRequest) (_result *GetMsgIdListByQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMsgIdListByQueueResponse{}
	_body, _err := client.GetMsgIdListByQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetQueueConsumers
//
// @param request - GetQueueConsumersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueueConsumersResponse
func (client *Client) GetQueueConsumersWithOptions(request *GetQueueConsumersRequest, runtime *dara.RuntimeOptions) (_result *GetQueueConsumersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQueueConsumers"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueueConsumersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetQueueConsumers
//
// @param request - GetQueueConsumersRequest
//
// @return GetQueueConsumersResponse
func (client *Client) GetQueueConsumers(request *GetQueueConsumersRequest) (_result *GetQueueConsumersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQueueConsumersResponse{}
	_body, _err := client.GetQueueConsumersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取队列错误
//
// @param request - GetQueueErrorByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueueErrorByTaskIdResponse
func (client *Client) GetQueueErrorByTaskIdWithOptions(request *GetQueueErrorByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetQueueErrorByTaskIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQueueErrorByTaskId"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueueErrorByTaskIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取队列错误
//
// @param request - GetQueueErrorByTaskIdRequest
//
// @return GetQueueErrorByTaskIdResponse
func (client *Client) GetQueueErrorByTaskId(request *GetQueueErrorByTaskIdRequest) (_result *GetQueueErrorByTaskIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQueueErrorByTaskIdResponse{}
	_body, _err := client.GetQueueErrorByTaskIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Queue Rate
//
// @param tmpReq - GetQueueRateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueueRateResponse
func (client *Client) GetQueueRateWithOptions(tmpReq *GetQueueRateRequest, runtime *dara.RuntimeOptions) (_result *GetQueueRateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetQueueRateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.QueueNames) {
		request.QueueNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueueNames, dara.String("QueueNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.QueueNamesShrink) {
		query["QueueNames"] = request.QueueNamesShrink
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQueueRate"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueueRateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Queue Rate
//
// @param request - GetQueueRateRequest
//
// @return GetQueueRateResponse
func (client *Client) GetQueueRate(request *GetQueueRateRequest) (_result *GetQueueRateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQueueRateResponse{}
	_body, _err := client.GetQueueRateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据connectionId,channelId,deliveryTag查询SendMessage
//
// @param request - GetSendTraceByConnectionAndChannelAndDeliveryTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSendTraceByConnectionAndChannelAndDeliveryTagResponse
func (client *Client) GetSendTraceByConnectionAndChannelAndDeliveryTagWithOptions(request *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest, runtime *dara.RuntimeOptions) (_result *GetSendTraceByConnectionAndChannelAndDeliveryTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConnectionId) {
		query["ConnectionId"] = request.ConnectionId
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.DeliveryTag) {
		query["DeliveryTag"] = request.DeliveryTag
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSendTraceByConnectionAndChannelAndDeliveryTag"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSendTraceByConnectionAndChannelAndDeliveryTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据connectionId,channelId,deliveryTag查询SendMessage
//
// @param request - GetSendTraceByConnectionAndChannelAndDeliveryTagRequest
//
// @return GetSendTraceByConnectionAndChannelAndDeliveryTagResponse
func (client *Client) GetSendTraceByConnectionAndChannelAndDeliveryTag(request *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) (_result *GetSendTraceByConnectionAndChannelAndDeliveryTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSendTraceByConnectionAndChannelAndDeliveryTagResponse{}
	_body, _err := client.GetSendTraceByConnectionAndChannelAndDeliveryTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过用户msgId查询消息发送轨迹
//
// @param request - GetSendTraceByMsgIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSendTraceByMsgIdResponse
func (client *Client) GetSendTraceByMsgIdWithOptions(request *GetSendTraceByMsgIdRequest, runtime *dara.RuntimeOptions) (_result *GetSendTraceByMsgIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MsgId) {
		query["MsgId"] = request.MsgId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSendTraceByMsgId"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSendTraceByMsgIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过用户msgId查询消息发送轨迹
//
// @param request - GetSendTraceByMsgIdRequest
//
// @return GetSendTraceByMsgIdResponse
func (client *Client) GetSendTraceByMsgId(request *GetSendTraceByMsgIdRequest) (_result *GetSendTraceByMsgIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSendTraceByMsgIdResponse{}
	_body, _err := client.GetSendTraceByMsgIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据queue 查询SendMessage
//
// @param request - GetSendTraceByQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSendTraceByQueueResponse
func (client *Client) GetSendTraceByQueueWithOptions(request *GetSendTraceByQueueRequest, runtime *dara.RuntimeOptions) (_result *GetSendTraceByQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSendTraceByQueue"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSendTraceByQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据queue 查询SendMessage
//
// @param request - GetSendTraceByQueueRequest
//
// @return GetSendTraceByQueueResponse
func (client *Client) GetSendTraceByQueue(request *GetSendTraceByQueueRequest) (_result *GetSendTraceByQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSendTraceByQueueResponse{}
	_body, _err := client.GetSendTraceByQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetStatisticsByVhost
//
// @param request - GetStatisticsByVhostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStatisticsByVhostResponse
func (client *Client) GetStatisticsByVhostWithOptions(request *GetStatisticsByVhostRequest, runtime *dara.RuntimeOptions) (_result *GetStatisticsByVhostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStatisticsByVhost"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStatisticsByVhostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetStatisticsByVhost
//
// @param request - GetStatisticsByVhostRequest
//
// @return GetStatisticsByVhostResponse
func (client *Client) GetStatisticsByVhost(request *GetStatisticsByVhostRequest) (_result *GetStatisticsByVhostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStatisticsByVhostResponse{}
	_body, _err := client.GetStatisticsByVhostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务
//
// @param request - GetTaskByUidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskByUidResponse
func (client *Client) GetTaskByUidWithOptions(request *GetTaskByUidRequest, runtime *dara.RuntimeOptions) (_result *GetTaskByUidResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskByUid"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskByUidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务
//
// @param request - GetTaskByUidRequest
//
// @return GetTaskByUidResponse
func (client *Client) GetTaskByUid(request *GetTaskByUidRequest) (_result *GetTaskByUidResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTaskByUidResponse{}
	_body, _err := client.GetTaskByUidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询一段时间内某个实例或是vhost或是queue的tps
//
// @param request - GetTpsByTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTpsByTimeResponse
func (client *Client) GetTpsByTimeWithOptions(request *GetTpsByTimeRequest, runtime *dara.RuntimeOptions) (_result *GetTpsByTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Api) {
		query["Api"] = request.Api
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTpsByTime"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTpsByTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询一段时间内某个实例或是vhost或是queue的tps
//
// @param request - GetTpsByTimeRequest
//
// @return GetTpsByTimeResponse
func (client *Client) GetTpsByTime(request *GetTpsByTimeRequest) (_result *GetTpsByTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTpsByTimeResponse{}
	_body, _err := client.GetTpsByTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户配置
//
// @param request - GetUserSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserSettingResponse
func (client *Client) GetUserSettingWithOptions(request *GetUserSettingRequest, runtime *dara.RuntimeOptions) (_result *GetUserSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserSetting"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户配置
//
// @param request - GetUserSettingRequest
//
// @return GetUserSettingResponse
func (client *Client) GetUserSetting(request *GetUserSettingRequest) (_result *GetUserSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserSettingResponse{}
	_body, _err := client.GetUserSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Vhost错误
//
// @param request - GetVhostErrorByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVhostErrorByTaskIdResponse
func (client *Client) GetVhostErrorByTaskIdWithOptions(request *GetVhostErrorByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetVhostErrorByTaskIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVhostErrorByTaskId"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVhostErrorByTaskIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Vhost错误
//
// @param request - GetVhostErrorByTaskIdRequest
//
// @return GetVhostErrorByTaskIdResponse
func (client *Client) GetVhostErrorByTaskId(request *GetVhostErrorByTaskIdRequest) (_result *GetVhostErrorByTaskIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVhostErrorByTaskIdResponse{}
	_body, _err := client.GetVhostErrorByTaskIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Vhost Rate
//
// @param tmpReq - GetVhostRateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVhostRateResponse
func (client *Client) GetVhostRateWithOptions(tmpReq *GetVhostRateRequest, runtime *dara.RuntimeOptions) (_result *GetVhostRateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetVhostRateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.VhostNames) {
		request.VhostNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VhostNames, dara.String("VhostNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.VhostNamesShrink) {
		query["VhostNames"] = request.VhostNamesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVhostRate"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVhostRateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Vhost Rate
//
// @param request - GetVhostRateRequest
//
// @return GetVhostRateResponse
func (client *Client) GetVhostRate(request *GetVhostRateRequest) (_result *GetVhostRateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVhostRateResponse{}
	_body, _err := client.GetVhostRateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 异步导入元数据
//
// @param request - ImportDefinitionAsynchronousRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportDefinitionAsynchronousResponse
func (client *Client) ImportDefinitionAsynchronousWithOptions(request *ImportDefinitionAsynchronousRequest, runtime *dara.RuntimeOptions) (_result *ImportDefinitionAsynchronousResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.ImportType) {
		query["ImportType"] = request.ImportType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.OssUrl) {
		query["OssUrl"] = request.OssUrl
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportDefinitionAsynchronous"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportDefinitionAsynchronousResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 异步导入元数据
//
// @param request - ImportDefinitionAsynchronousRequest
//
// @return ImportDefinitionAsynchronousResponse
func (client *Client) ImportDefinitionAsynchronous(request *ImportDefinitionAsynchronousRequest) (_result *ImportDefinitionAsynchronousResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportDefinitionAsynchronousResponse{}
	_body, _err := client.ImportDefinitionAsynchronousWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例列表
//
// @param request - InstancePreivewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstancePreivewResponse
func (client *Client) InstancePreivewWithOptions(request *InstancePreivewRequest, runtime *dara.RuntimeOptions) (_result *InstancePreivewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstancePreivew"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstancePreivewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例列表
//
// @param request - InstancePreivewRequest
//
// @return InstancePreivewResponse
func (client *Client) InstancePreivew(request *InstancePreivewRequest) (_result *InstancePreivewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InstancePreivewResponse{}
	_body, _err := client.InstancePreivewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Exchange列表
//
// @param request - ListExchangeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExchangeResponse
func (client *Client) ListExchangeWithOptions(request *ListExchangeRequest, runtime *dara.RuntimeOptions) (_result *ListExchangeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.ExchangeNamePrefix) {
		query["ExchangeNamePrefix"] = request.ExchangeNamePrefix
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExchange"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExchangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Exchange列表
//
// @param request - ListExchangeRequest
//
// @return ListExchangeResponse
func (client *Client) ListExchange(request *ListExchangeRequest) (_result *ListExchangeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExchangeResponse{}
	_body, _err := client.ListExchangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Exchange下游列表
//
// @param request - ListExchangeDownstreamBindingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExchangeDownstreamBindingsResponse
func (client *Client) ListExchangeDownstreamBindingsWithOptions(request *ListExchangeDownstreamBindingsRequest, runtime *dara.RuntimeOptions) (_result *ListExchangeDownstreamBindingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.ExchangeName) {
		query["ExchangeName"] = request.ExchangeName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExchangeDownstreamBindings"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExchangeDownstreamBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Exchange下游列表
//
// @param request - ListExchangeDownstreamBindingsRequest
//
// @return ListExchangeDownstreamBindingsResponse
func (client *Client) ListExchangeDownstreamBindings(request *ListExchangeDownstreamBindingsRequest) (_result *ListExchangeDownstreamBindingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExchangeDownstreamBindingsResponse{}
	_body, _err := client.ListExchangeDownstreamBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Exchange上游绑定列表
//
// @param request - ListExchangeUpstreamBindingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExchangeUpstreamBindingsResponse
func (client *Client) ListExchangeUpstreamBindingsWithOptions(request *ListExchangeUpstreamBindingsRequest, runtime *dara.RuntimeOptions) (_result *ListExchangeUpstreamBindingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.ExchangeName) {
		query["ExchangeName"] = request.ExchangeName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExchangeUpstreamBindings"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExchangeUpstreamBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Exchange上游绑定列表
//
// @param request - ListExchangeUpstreamBindingsRequest
//
// @return ListExchangeUpstreamBindingsResponse
func (client *Client) ListExchangeUpstreamBindings(request *ListExchangeUpstreamBindingsRequest) (_result *ListExchangeUpstreamBindingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExchangeUpstreamBindingsResponse{}
	_body, _err := client.ListExchangeUpstreamBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例列表
//
// @param request - ListInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceResponse
func (client *Client) ListInstanceWithOptions(request *ListInstanceRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstance"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例列表
//
// @param request - ListInstanceRequest
//
// @return ListInstanceResponse
func (client *Client) ListInstance(request *ListInstanceRequest) (_result *ListInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceResponse{}
	_body, _err := client.ListInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例告警
//
// @param request - ListInstanceAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceAlarmResponse
func (client *Client) ListInstanceAlarmWithOptions(request *ListInstanceAlarmRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceAlarm"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例告警
//
// @param request - ListInstanceAlarmRequest
//
// @return ListInstanceAlarmResponse
func (client *Client) ListInstanceAlarm(request *ListInstanceAlarmRequest) (_result *ListInstanceAlarmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceAlarmResponse{}
	_body, _err := client.ListInstanceAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取日志Logstore
//
// @param request - ListLogstoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogstoreResponse
func (client *Client) ListLogstoreWithOptions(request *ListLogstoreRequest, runtime *dara.RuntimeOptions) (_result *ListLogstoreResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLogstore"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLogstoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取日志Logstore
//
// @param request - ListLogstoreRequest
//
// @return ListLogstoreResponse
func (client *Client) ListLogstore(request *ListLogstoreRequest) (_result *ListLogstoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLogstoreResponse{}
	_body, _err := client.ListLogstoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取日志Project
//
// @param request - ListProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectResponse
func (client *Client) ListProjectWithOptions(request *ListProjectRequest, runtime *dara.RuntimeOptions) (_result *ListProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProject"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取日志Project
//
// @param request - ListProjectRequest
//
// @return ListProjectResponse
func (client *Client) ListProject(request *ListProjectRequest) (_result *ListProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectResponse{}
	_body, _err := client.ListProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取队列列表
//
// @param request - ListQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueueResponse
func (client *Client) ListQueueWithOptions(request *ListQueueRequest, runtime *dara.RuntimeOptions) (_result *ListQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueueNamePrefix) {
		query["QueueNamePrefix"] = request.QueueNamePrefix
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQueue"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取队列列表
//
// @param request - ListQueueRequest
//
// @return ListQueueResponse
func (client *Client) ListQueue(request *ListQueueRequest) (_result *ListQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQueueResponse{}
	_body, _err := client.ListQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取队列上游绑定列表
//
// @param request - ListQueueUpstreamBindingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueueUpstreamBindingsResponse
func (client *Client) ListQueueUpstreamBindingsWithOptions(request *ListQueueUpstreamBindingsRequest, runtime *dara.RuntimeOptions) (_result *ListQueueUpstreamBindingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQueueUpstreamBindings"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQueueUpstreamBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取队列上游绑定列表
//
// @param request - ListQueueUpstreamBindingsRequest
//
// @return ListQueueUpstreamBindingsResponse
func (client *Client) ListQueueUpstreamBindings(request *ListQueueUpstreamBindingsRequest) (_result *ListQueueUpstreamBindingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQueueUpstreamBindingsResponse{}
	_body, _err := client.ListQueueUpstreamBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取静态账户列表
//
// @param request - ListStaticAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStaticAccountsResponse
func (client *Client) ListStaticAccountsWithOptions(request *ListStaticAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListStaticAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStaticAccounts"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStaticAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取静态账户列表
//
// @param request - ListStaticAccountsRequest
//
// @return ListStaticAccountsResponse
func (client *Client) ListStaticAccounts(request *ListStaticAccountsRequest) (_result *ListStaticAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListStaticAccountsResponse{}
	_body, _err := client.ListStaticAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Vhost列表
//
// @param request - ListVhostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVhostResponse
func (client *Client) ListVhostWithOptions(request *ListVhostRequest, runtime *dara.RuntimeOptions) (_result *ListVhostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.VhostNamePrefix) {
		query["VhostNamePrefix"] = request.VhostNamePrefix
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVhost"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVhostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Vhost列表
//
// @param request - ListVhostRequest
//
// @return ListVhostResponse
func (client *Client) ListVhost(request *ListVhostRequest) (_result *ListVhostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVhostResponse{}
	_body, _err := client.ListVhostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取元数据
//
// @param request - MetadataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MetadataResponse
func (client *Client) MetadataWithOptions(request *MetadataRequest, runtime *dara.RuntimeOptions) (_result *MetadataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Metadata"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MetadataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取元数据
//
// @param request - MetadataRequest
//
// @return MetadataResponse
func (client *Client) Metadata(request *MetadataRequest) (_result *MetadataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MetadataResponse{}
	_body, _err := client.MetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 清空队列
//
// @param tmpReq - PurgeQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PurgeQueueResponse
func (client *Client) PurgeQueueWithOptions(tmpReq *PurgeQueueRequest, runtime *dara.RuntimeOptions) (_result *PurgeQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PurgeQueueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.QueueNames) {
		request.QueueNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueueNames, dara.String("QueueNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Collina) {
		query["Collina"] = request.Collina
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.QueueNamesShrink) {
		query["QueueNames"] = request.QueueNamesShrink
	}

	if !dara.IsNil(request.UmidToken) {
		query["UmidToken"] = request.UmidToken
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PurgeQueue"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PurgeQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 清空队列
//
// @param request - PurgeQueueRequest
//
// @return PurgeQueueResponse
func (client *Client) PurgeQueue(request *PurgeQueueRequest) (_result *PurgeQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PurgeQueueResponse{}
	_body, _err := client.PurgeQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据Message Id查询消息
//
// @param request - QueryMessageByMessageIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMessageByMessageIdResponse
func (client *Client) QueryMessageByMessageIdWithOptions(request *QueryMessageByMessageIdRequest, runtime *dara.RuntimeOptions) (_result *QueryMessageByMessageIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMessageByMessageId"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMessageByMessageIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据Message Id查询消息
//
// @param request - QueryMessageByMessageIdRequest
//
// @return QueryMessageByMessageIdResponse
func (client *Client) QueryMessageByMessageId(request *QueryMessageByMessageIdRequest) (_result *QueryMessageByMessageIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryMessageByMessageIdResponse{}
	_body, _err := client.QueryMessageByMessageIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据队列查询消息
//
// @param request - QueryMessageByQueueNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMessageByQueueNameResponse
func (client *Client) QueryMessageByQueueNameWithOptions(request *QueryMessageByQueueNameRequest, runtime *dara.RuntimeOptions) (_result *QueryMessageByQueueNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMessageByQueueName"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMessageByQueueNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据队列查询消息
//
// @param request - QueryMessageByQueueNameRequest
//
// @return QueryMessageByQueueNameResponse
func (client *Client) QueryMessageByQueueName(request *QueryMessageByQueueNameRequest) (_result *QueryMessageByQueueNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryMessageByQueueNameResponse{}
	_body, _err := client.QueryMessageByQueueNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实例释放
//
// @param request - ReleaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstanceResponse
func (client *Client) ReleaseInstanceWithOptions(request *ReleaseInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseInstance"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实例释放
//
// @param request - ReleaseInstanceRequest
//
// @return ReleaseInstanceResponse
func (client *Client) ReleaseInstance(request *ReleaseInstanceRequest) (_result *ReleaseInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.ReleaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送消息
//
// @param request - SendMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageResponse
func (client *Client) SendMessageWithOptions(request *SendMessageRequest, runtime *dara.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		query["Body"] = request.Body
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.ExchangeName) {
		query["ExchangeName"] = request.ExchangeName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	if !dara.IsNil(request.Props) {
		query["Props"] = request.Props
	}

	if !dara.IsNil(request.RoutingKey) {
		query["RoutingKey"] = request.RoutingKey
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendMessage"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送消息
//
// @param request - SendMessageRequest
//
// @return SendMessageResponse
func (client *Client) SendMessage(request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送消息
//
// @param request - SendMessageCopyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageCopyResponse
func (client *Client) SendMessageCopyWithOptions(request *SendMessageCopyRequest, runtime *dara.RuntimeOptions) (_result *SendMessageCopyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProcessToken) {
		query["ProcessToken"] = request.ProcessToken
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendMessageCopy"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendMessageCopyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送消息
//
// @param request - SendMessageCopyRequest
//
// @return SendMessageCopyResponse
func (client *Client) SendMessageCopy(request *SendMessageCopyRequest) (_result *SendMessageCopyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendMessageCopyResponse{}
	_body, _err := client.SendMessageCopyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消绑定
//
// @param request - UnbindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindResponse
func (client *Client) UnbindWithOptions(request *UnbindRequest, runtime *dara.RuntimeOptions) (_result *UnbindResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BindingKey) {
		query["BindingKey"] = request.BindingKey
	}

	if !dara.IsNil(request.BindingType) {
		query["BindingType"] = request.BindingType
	}

	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.DstName) {
		query["DstName"] = request.DstName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SrcName) {
		query["SrcName"] = request.SrcName
	}

	if !dara.IsNil(request.VhostName) {
		query["VhostName"] = request.VhostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Unbind"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消绑定
//
// @param request - UnbindRequest
//
// @return UnbindResponse
func (client *Client) Unbind(request *UnbindRequest) (_result *UnbindResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindResponse{}
	_body, _err := client.UnbindWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新实例
//
// @param request - UpdateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithOptions(request *UpdateInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstance"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实例
//
// @param request - UpdateInstanceRequest
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstance(request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改实例的重试策略
//
// @param request - UpdateInstanceRetryStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceRetryStrategyResponse
func (client *Client) UpdateInstanceRetryStrategyWithOptions(request *UpdateInstanceRetryStrategyRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceRetryStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RetryInterval) {
		query["RetryInterval"] = request.RetryInterval
	}

	if !dara.IsNil(request.RetryTimes) {
		query["RetryTimes"] = request.RetryTimes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceRetryStrategy"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceRetryStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改实例的重试策略
//
// @param request - UpdateInstanceRetryStrategyRequest
//
// @return UpdateInstanceRetryStrategyResponse
func (client *Client) UpdateInstanceRetryStrategy(request *UpdateInstanceRetryStrategyRequest) (_result *UpdateInstanceRetryStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInstanceRetryStrategyResponse{}
	_body, _err := client.UpdateInstanceRetryStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新serverless开关
//
// @param request - UpdateServerlessSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServerlessSwitchResponse
func (client *Client) UpdateServerlessSwitchWithOptions(request *UpdateServerlessSwitchRequest, runtime *dara.RuntimeOptions) (_result *UpdateServerlessSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ServerlessSwitch) {
		query["ServerlessSwitch"] = request.ServerlessSwitch
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServerlessSwitch"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServerlessSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新serverless开关
//
// @param request - UpdateServerlessSwitchRequest
//
// @return UpdateServerlessSwitchResponse
func (client *Client) UpdateServerlessSwitch(request *UpdateServerlessSwitchRequest) (_result *UpdateServerlessSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateServerlessSwitchResponse{}
	_body, _err := client.UpdateServerlessSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 升级实例配额
//
// @param request - UpgradeLimitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeLimitsResponse
func (client *Client) UpgradeLimitsWithOptions(request *UpgradeLimitsRequest, runtime *dara.RuntimeOptions) (_result *UpgradeLimitsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleSessionId) {
		query["ConsoleSessionId"] = request.ConsoleSessionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeLimits"),
		Version:     dara.String("2019-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeLimitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 升级实例配额
//
// @param request - UpgradeLimitsRequest
//
// @return UpgradeLimitsResponse
func (client *Client) UpgradeLimits(request *UpgradeLimitsRequest) (_result *UpgradeLimitsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeLimitsResponse{}
	_body, _err := client.UpgradeLimitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
