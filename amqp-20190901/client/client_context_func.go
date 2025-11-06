// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 路由绑定
//
// @param request - BindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindResponse
func (client *Client) BindWithContext(ctx context.Context, request *BindRequest, runtime *dara.RuntimeOptions) (_result *BindResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelUserSettingResponse
func (client *Client) CancelUserSettingWithContext(ctx context.Context, request *CancelUserSettingRequest, runtime *dara.RuntimeOptions) (_result *CancelUserSettingResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureUserSettingResponse
func (client *Client) ConfigureUserSettingWithContext(ctx context.Context, request *ConfigureUserSettingRequest, runtime *dara.RuntimeOptions) (_result *ConfigureUserSettingResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConsoleClearPretendStatusResponse
func (client *Client) ConsoleClearPretendStatusWithContext(ctx context.Context, request *ConsoleClearPretendStatusRequest, runtime *dara.RuntimeOptions) (_result *ConsoleClearPretendStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConsoleSavePretendStatusResponse
func (client *Client) ConsoleSavePretendStatusWithContext(ctx context.Context, request *ConsoleSavePretendStatusRequest, runtime *dara.RuntimeOptions) (_result *ConsoleSavePretendStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudMonitorSLRResponse
func (client *Client) CreateCloudMonitorSLRWithContext(ctx context.Context, request *CreateCloudMonitorSLRRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudMonitorSLRResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExchangeResponse
func (client *Client) CreateExchangeWithContext(ctx context.Context, request *CreateExchangeRequest, runtime *dara.RuntimeOptions) (_result *CreateExchangeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLogDeliverySLRResponse
func (client *Client) CreateLogDeliverySLRWithContext(ctx context.Context, request *CreateLogDeliverySLRRequest, runtime *dara.RuntimeOptions) (_result *CreateLogDeliverySLRResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQueueResponse
func (client *Client) CreateQueueWithContext(ctx context.Context, request *CreateQueueRequest, runtime *dara.RuntimeOptions) (_result *CreateQueueResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVhostResponse
func (client *Client) CreateVhostWithContext(ctx context.Context, request *CreateVhostRequest, runtime *dara.RuntimeOptions) (_result *CreateVhostResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DashboardCheckServiceStatusResponse
func (client *Client) DashboardCheckServiceStatusWithContext(ctx context.Context, request *DashboardCheckServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *DashboardCheckServiceStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DashboardListResponse
func (client *Client) DashboardListWithContext(ctx context.Context, request *DashboardListRequest, runtime *dara.RuntimeOptions) (_result *DashboardListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DeleteExchangeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExchangeResponse
func (client *Client) DeleteExchangeWithContext(ctx context.Context, tmpReq *DeleteExchangeRequest, runtime *dara.RuntimeOptions) (_result *DeleteExchangeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DeleteQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQueueResponse
func (client *Client) DeleteQueueWithContext(ctx context.Context, tmpReq *DeleteQueueRequest, runtime *dara.RuntimeOptions) (_result *DeleteQueueResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStaticAccountResponse
func (client *Client) DeleteStaticAccountWithContext(ctx context.Context, request *DeleteStaticAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteStaticAccountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DeleteVhostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVhostResponse
func (client *Client) DeleteVhostWithContext(ctx context.Context, tmpReq *DeleteVhostRequest, runtime *dara.RuntimeOptions) (_result *DeleteVhostResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportResponse
func (client *Client) ExportWithContext(ctx context.Context, request *ExportRequest, runtime *dara.RuntimeOptions) (_result *ExportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchStaticAccountResponse
func (client *Client) FetchStaticAccountWithContext(ctx context.Context, request *FetchStaticAccountRequest, runtime *dara.RuntimeOptions) (_result *FetchStaticAccountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAckInfoByIntervalResponse
func (client *Client) GetAckInfoByIntervalWithContext(ctx context.Context, request *GetAckInfoByIntervalRequest, runtime *dara.RuntimeOptions) (_result *GetAckInfoByIntervalResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAckInfoOfMessageResponse
func (client *Client) GetAckInfoOfMessageWithContext(ctx context.Context, request *GetAckInfoOfMessageRequest, runtime *dara.RuntimeOptions) (_result *GetAckInfoOfMessageResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBindingCountResponse
func (client *Client) GetBindingCountWithContext(ctx context.Context, request *GetBindingCountRequest, runtime *dara.RuntimeOptions) (_result *GetBindingCountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBindingErrorByTaskIdResponse
func (client *Client) GetBindingErrorByTaskIdWithContext(ctx context.Context, request *GetBindingErrorByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetBindingErrorByTaskIdResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetCommonBuyUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCommonBuyUrlResponse
func (client *Client) GetCommonBuyUrlWithContext(ctx context.Context, request *GetCommonBuyUrlRequest, runtime *dara.RuntimeOptions) (_result *GetCommonBuyUrlResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumeTraceByQueueAndRocketMqMsgIdResponse
func (client *Client) GetConsumeTraceByQueueAndRocketMqMsgIdWithContext(ctx context.Context, request *GetConsumeTraceByQueueAndRocketMqMsgIdRequest, runtime *dara.RuntimeOptions) (_result *GetConsumeTraceByQueueAndRocketMqMsgIdResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExchangeErrorByTaskIdResponse
func (client *Client) GetExchangeErrorByTaskIdWithContext(ctx context.Context, request *GetExchangeErrorByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetExchangeErrorByTaskIdResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - GetExchangeRateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExchangeRateResponse
func (client *Client) GetExchangeRateWithContext(ctx context.Context, tmpReq *GetExchangeRateRequest, runtime *dara.RuntimeOptions) (_result *GetExchangeRateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMsgIdListByQueueResponse
func (client *Client) GetMsgIdListByQueueWithContext(ctx context.Context, request *GetMsgIdListByQueueRequest, runtime *dara.RuntimeOptions) (_result *GetMsgIdListByQueueResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueueConsumersResponse
func (client *Client) GetQueueConsumersWithContext(ctx context.Context, request *GetQueueConsumersRequest, runtime *dara.RuntimeOptions) (_result *GetQueueConsumersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueueErrorByTaskIdResponse
func (client *Client) GetQueueErrorByTaskIdWithContext(ctx context.Context, request *GetQueueErrorByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetQueueErrorByTaskIdResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - GetQueueRateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueueRateResponse
func (client *Client) GetQueueRateWithContext(ctx context.Context, tmpReq *GetQueueRateRequest, runtime *dara.RuntimeOptions) (_result *GetQueueRateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSendTraceByConnectionAndChannelAndDeliveryTagResponse
func (client *Client) GetSendTraceByConnectionAndChannelAndDeliveryTagWithContext(ctx context.Context, request *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest, runtime *dara.RuntimeOptions) (_result *GetSendTraceByConnectionAndChannelAndDeliveryTagResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSendTraceByMsgIdResponse
func (client *Client) GetSendTraceByMsgIdWithContext(ctx context.Context, request *GetSendTraceByMsgIdRequest, runtime *dara.RuntimeOptions) (_result *GetSendTraceByMsgIdResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSendTraceByQueueResponse
func (client *Client) GetSendTraceByQueueWithContext(ctx context.Context, request *GetSendTraceByQueueRequest, runtime *dara.RuntimeOptions) (_result *GetSendTraceByQueueResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStatisticsByVhostResponse
func (client *Client) GetStatisticsByVhostWithContext(ctx context.Context, request *GetStatisticsByVhostRequest, runtime *dara.RuntimeOptions) (_result *GetStatisticsByVhostResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskByUidResponse
func (client *Client) GetTaskByUidWithContext(ctx context.Context, request *GetTaskByUidRequest, runtime *dara.RuntimeOptions) (_result *GetTaskByUidResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTpsByTimeResponse
func (client *Client) GetTpsByTimeWithContext(ctx context.Context, request *GetTpsByTimeRequest, runtime *dara.RuntimeOptions) (_result *GetTpsByTimeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserSettingResponse
func (client *Client) GetUserSettingWithContext(ctx context.Context, request *GetUserSettingRequest, runtime *dara.RuntimeOptions) (_result *GetUserSettingResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVhostErrorByTaskIdResponse
func (client *Client) GetVhostErrorByTaskIdWithContext(ctx context.Context, request *GetVhostErrorByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetVhostErrorByTaskIdResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - GetVhostRateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVhostRateResponse
func (client *Client) GetVhostRateWithContext(ctx context.Context, tmpReq *GetVhostRateRequest, runtime *dara.RuntimeOptions) (_result *GetVhostRateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportDefinitionAsynchronousResponse
func (client *Client) ImportDefinitionAsynchronousWithContext(ctx context.Context, request *ImportDefinitionAsynchronousRequest, runtime *dara.RuntimeOptions) (_result *ImportDefinitionAsynchronousResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstancePreivewResponse
func (client *Client) InstancePreivewWithContext(ctx context.Context, request *InstancePreivewRequest, runtime *dara.RuntimeOptions) (_result *InstancePreivewResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExchangeResponse
func (client *Client) ListExchangeWithContext(ctx context.Context, request *ListExchangeRequest, runtime *dara.RuntimeOptions) (_result *ListExchangeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExchangeDownstreamBindingsResponse
func (client *Client) ListExchangeDownstreamBindingsWithContext(ctx context.Context, request *ListExchangeDownstreamBindingsRequest, runtime *dara.RuntimeOptions) (_result *ListExchangeDownstreamBindingsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExchangeUpstreamBindingsResponse
func (client *Client) ListExchangeUpstreamBindingsWithContext(ctx context.Context, request *ListExchangeUpstreamBindingsRequest, runtime *dara.RuntimeOptions) (_result *ListExchangeUpstreamBindingsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceResponse
func (client *Client) ListInstanceWithContext(ctx context.Context, request *ListInstanceRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceAlarmResponse
func (client *Client) ListInstanceAlarmWithContext(ctx context.Context, request *ListInstanceAlarmRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceAlarmResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogstoreResponse
func (client *Client) ListLogstoreWithContext(ctx context.Context, request *ListLogstoreRequest, runtime *dara.RuntimeOptions) (_result *ListLogstoreResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectResponse
func (client *Client) ListProjectWithContext(ctx context.Context, request *ListProjectRequest, runtime *dara.RuntimeOptions) (_result *ListProjectResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueueResponse
func (client *Client) ListQueueWithContext(ctx context.Context, request *ListQueueRequest, runtime *dara.RuntimeOptions) (_result *ListQueueResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueueUpstreamBindingsResponse
func (client *Client) ListQueueUpstreamBindingsWithContext(ctx context.Context, request *ListQueueUpstreamBindingsRequest, runtime *dara.RuntimeOptions) (_result *ListQueueUpstreamBindingsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStaticAccountsResponse
func (client *Client) ListStaticAccountsWithContext(ctx context.Context, request *ListStaticAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListStaticAccountsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVhostResponse
func (client *Client) ListVhostWithContext(ctx context.Context, request *ListVhostRequest, runtime *dara.RuntimeOptions) (_result *ListVhostResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MetadataResponse
func (client *Client) MetadataWithContext(ctx context.Context, request *MetadataRequest, runtime *dara.RuntimeOptions) (_result *MetadataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - PurgeQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PurgeQueueResponse
func (client *Client) PurgeQueueWithContext(ctx context.Context, tmpReq *PurgeQueueRequest, runtime *dara.RuntimeOptions) (_result *PurgeQueueResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMessageByMessageIdResponse
func (client *Client) QueryMessageByMessageIdWithContext(ctx context.Context, request *QueryMessageByMessageIdRequest, runtime *dara.RuntimeOptions) (_result *QueryMessageByMessageIdResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMessageByQueueNameResponse
func (client *Client) QueryMessageByQueueNameWithContext(ctx context.Context, request *QueryMessageByQueueNameRequest, runtime *dara.RuntimeOptions) (_result *QueryMessageByQueueNameResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstanceResponse
func (client *Client) ReleaseInstanceWithContext(ctx context.Context, request *ReleaseInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageResponse
func (client *Client) SendMessageWithContext(ctx context.Context, request *SendMessageRequest, runtime *dara.RuntimeOptions) (_result *SendMessageResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageCopyResponse
func (client *Client) SendMessageCopyWithContext(ctx context.Context, request *SendMessageCopyRequest, runtime *dara.RuntimeOptions) (_result *SendMessageCopyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindResponse
func (client *Client) UnbindWithContext(ctx context.Context, request *UnbindRequest, runtime *dara.RuntimeOptions) (_result *UnbindResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithContext(ctx context.Context, request *UpdateInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceRetryStrategyResponse
func (client *Client) UpdateInstanceRetryStrategyWithContext(ctx context.Context, request *UpdateInstanceRetryStrategyRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceRetryStrategyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServerlessSwitchResponse
func (client *Client) UpdateServerlessSwitchWithContext(ctx context.Context, request *UpdateServerlessSwitchRequest, runtime *dara.RuntimeOptions) (_result *UpdateServerlessSwitchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeLimitsResponse
func (client *Client) UpgradeLimitsWithContext(ctx context.Context, request *UpgradeLimitsRequest, runtime *dara.RuntimeOptions) (_result *UpgradeLimitsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
