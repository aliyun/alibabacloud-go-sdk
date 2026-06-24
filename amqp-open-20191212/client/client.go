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
		"us-west-1":             dara.String("amqp-open.us-west-1.aliyuncs.com"),
		"us-east-1":             dara.String("amqp-open.us-east-1.aliyuncs.com"),
		"me-central-1":          dara.String("amqp-open.me-central-1.aliyuncs.com"),
		"eu-central-1":          dara.String("amqp-open.eu-central-1.aliyuncs.com"),
		"cn-zhengzhou-jva":      dara.String("amqp-open.cn-zhengzhou-jva.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("amqp-open.cn-zhangjiakou.aliyuncs.com"),
		"cn-wulanchabu":         dara.String("amqp-open.cn-wulanchabu.aliyuncs.com"),
		"cn-shenzhen":           dara.String("amqp-open.cn-shenzhen.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("amqp-open.cn-shanghai-finance-1.aliyuncs.com"),
		"cn-shanghai":           dara.String("amqp-open.cn-shanghai.aliyuncs.com"),
		"cn-qingdao":            dara.String("amqp-open.cn-qingdao.aliyuncs.com"),
		"cn-huhehaote":          dara.String("amqp-open.cn-huhehaote.aliyuncs.com"),
		"cn-hongkong":           dara.String("amqp-open.cn-hongkong.aliyuncs.com"),
		"cn-hangzhou":           dara.String("amqp-open.cn-hangzhou.aliyuncs.com"),
		"cn-guangzhou":          dara.String("amqp-open.cn-guangzhou.aliyuncs.com"),
		"cn-chengdu":            dara.String("amqp-open.cn-chengdu.aliyuncs.com"),
		"cn-beijing-finance-1":  dara.String("amqp-open.cn-beijing-finance-1.aliyuncs.com"),
		"cn-beijing":            dara.String("amqp-open.cn-beijing.aliyuncs.com"),
		"ap-southeast-7":        dara.String("amqp-open.ap-southeast-7.aliyuncs.com"),
		"ap-southeast-6":        dara.String("amqp-open.ap-southeast-6.aliyuncs.com"),
		"ap-southeast-5":        dara.String("amqp-open.ap-southeast-5.aliyuncs.com"),
		"ap-southeast-3":        dara.String("amqp-open.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-1":        dara.String("amqp-open.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-1":        dara.String("amqp-open.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("amqp-open"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds an entry to the whitelist of an instance.
//
// @param tmpReq - AddInstanceWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddInstanceWhiteListResponse
func (client *Client) AddInstanceWhiteListWithOptions(tmpReq *AddInstanceWhiteListRequest, runtime *dara.RuntimeOptions) (_result *AddInstanceWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddInstanceWhiteListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WhiteListItem) {
		request.WhiteListItemShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WhiteListItem, dara.String("WhiteListItem"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.WhiteListItemShrink) {
		query["WhiteListItem"] = request.WhiteListItemShrink
	}

	if !dara.IsNil(request.WhiteListType) {
		query["WhiteListType"] = request.WhiteListType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddInstanceWhiteList"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddInstanceWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an entry to the whitelist of an instance.
//
// @param request - AddInstanceWhiteListRequest
//
// @return AddInstanceWhiteListResponse
func (client *Client) AddInstanceWhiteList(request *AddInstanceWhiteListRequest) (_result *AddInstanceWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddInstanceWhiteListResponse{}
	_body, _err := client.AddInstanceWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// When an open-source client accesses an ApsaraMQ for RabbitMQ server, it must provide a username and password for authentication. ApsaraMQ for RabbitMQ lets you generate a username and password from an AccessKey ID and AccessKey secret provided by Resource Access Management (RAM).
//
// @param request - CreateAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccountResponse
func (client *Client) CreateAccountWithOptions(request *CreateAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.AccountAccessKey) {
		query["accountAccessKey"] = request.AccountAccessKey
	}

	if !dara.IsNil(request.CreateTimestamp) {
		query["createTimestamp"] = request.CreateTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecretSign) {
		query["secretSign"] = request.SecretSign
	}

	if !dara.IsNil(request.Signature) {
		query["signature"] = request.Signature
	}

	if !dara.IsNil(request.UserName) {
		query["userName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccount"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// When an open-source client accesses an ApsaraMQ for RabbitMQ server, it must provide a username and password for authentication. ApsaraMQ for RabbitMQ lets you generate a username and password from an AccessKey ID and AccessKey secret provided by Resource Access Management (RAM).
//
// @param request - CreateAccountRequest
//
// @return CreateAccountResponse
func (client *Client) CreateAccount(request *CreateAccountRequest) (_result *CreateAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAccountResponse{}
	_body, _err := client.CreateAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// A producer sends a message to an exchange, which then routes the message to a specified queue or another exchange based on the binding and routing rules.
//
// @param request - CreateBindingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBindingResponse
func (client *Client) CreateBindingWithOptions(request *CreateBindingRequest, runtime *dara.RuntimeOptions) (_result *CreateBindingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Argument) {
		body["Argument"] = request.Argument
	}

	if !dara.IsNil(request.BindingKey) {
		body["BindingKey"] = request.BindingKey
	}

	if !dara.IsNil(request.BindingType) {
		body["BindingType"] = request.BindingType
	}

	if !dara.IsNil(request.DestinationName) {
		body["DestinationName"] = request.DestinationName
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SourceExchange) {
		body["SourceExchange"] = request.SourceExchange
	}

	if !dara.IsNil(request.VirtualHost) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBinding"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// A producer sends a message to an exchange, which then routes the message to a specified queue or another exchange based on the binding and routing rules.
//
// @param request - CreateBindingRequest
//
// @return CreateBindingResponse
func (client *Client) CreateBinding(request *CreateBindingRequest) (_result *CreateBindingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBindingResponse{}
	_body, _err := client.CreateBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// A producer sends a message to an exchange. The exchange then routes the message to one or more queues based on the routing key and the binding key, or discards the message.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlternateExchange) {
		body["AlternateExchange"] = request.AlternateExchange
	}

	if !dara.IsNil(request.AutoDeleteState) {
		body["AutoDeleteState"] = request.AutoDeleteState
	}

	if !dara.IsNil(request.ExchangeName) {
		body["ExchangeName"] = request.ExchangeName
	}

	if !dara.IsNil(request.ExchangeType) {
		body["ExchangeType"] = request.ExchangeType
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Internal) {
		body["Internal"] = request.Internal
	}

	if !dara.IsNil(request.VirtualHost) {
		body["VirtualHost"] = request.VirtualHost
	}

	if !dara.IsNil(request.XDelayedType) {
		body["XDelayedType"] = request.XDelayedType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExchange"),
		Version:     dara.String("2019-12-12"),
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
// A producer sends a message to an exchange. The exchange then routes the message to one or more queues based on the routing key and the binding key, or discards the message.
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
// Creates an ApsaraMQ for RabbitMQ instance.
//
// @param tmpReq - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(tmpReq *CreateInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VswitchIds) {
		request.VswitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VswitchIds, dara.String("VswitchIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthModel) {
		query["AuthModel"] = request.AuthModel
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.EncryptedInstance) {
		query["EncryptedInstance"] = request.EncryptedInstance
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.KmsKeyId) {
		query["KmsKeyId"] = request.KmsKeyId
	}

	if !dara.IsNil(request.ListenerMode) {
		query["ListenerMode"] = request.ListenerMode
	}

	if !dara.IsNil(request.MaxConnections) {
		query["MaxConnections"] = request.MaxConnections
	}

	if !dara.IsNil(request.MaxEipTps) {
		query["MaxEipTps"] = request.MaxEipTps
	}

	if !dara.IsNil(request.MaxPrivateTps) {
		query["MaxPrivateTps"] = request.MaxPrivateTps
	}

	if !dara.IsNil(request.PaymentType) {
		query["PaymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodCycle) {
		query["PeriodCycle"] = request.PeriodCycle
	}

	if !dara.IsNil(request.ProvisionedCapacity) {
		query["ProvisionedCapacity"] = request.ProvisionedCapacity
	}

	if !dara.IsNil(request.QueueCapacity) {
		query["QueueCapacity"] = request.QueueCapacity
	}

	if !dara.IsNil(request.RenewStatus) {
		query["RenewStatus"] = request.RenewStatus
	}

	if !dara.IsNil(request.RenewalDurationUnit) {
		query["RenewalDurationUnit"] = request.RenewalDurationUnit
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.ServerlessChargeType) {
		query["ServerlessChargeType"] = request.ServerlessChargeType
	}

	if !dara.IsNil(request.ServerlessSwitch) {
		query["ServerlessSwitch"] = request.ServerlessSwitch
	}

	if !dara.IsNil(request.StorageSize) {
		query["StorageSize"] = request.StorageSize
	}

	if !dara.IsNil(request.SupportEip) {
		query["SupportEip"] = request.SupportEip
	}

	if !dara.IsNil(request.SupportTracing) {
		query["SupportTracing"] = request.SupportTracing
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TracingStorageTime) {
		query["TracingStorageTime"] = request.TracingStorageTime
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VswitchIdsShrink) {
		query["VswitchIds"] = request.VswitchIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an ApsaraMQ for RabbitMQ instance.
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an open-source username and password.
//
// @param request - CreateOpenSourceAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOpenSourceAccountResponse
func (client *Client) CreateOpenSourceAccountWithOptions(request *CreateOpenSourceAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateOpenSourceAccountResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOpenSourceAccount"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOpenSourceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an open-source username and password.
//
// @param request - CreateOpenSourceAccountRequest
//
// @return CreateOpenSourceAccountResponse
func (client *Client) CreateOpenSourceAccount(request *CreateOpenSourceAccountRequest) (_result *CreateOpenSourceAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOpenSourceAccountResponse{}
	_body, _err := client.CreateOpenSourceAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an open source permission.
//
// @param request - CreateOpenSourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOpenSourcePermissionResponse
func (client *Client) CreateOpenSourcePermissionWithOptions(request *CreateOpenSourcePermissionRequest, runtime *dara.RuntimeOptions) (_result *CreateOpenSourcePermissionResponse, _err error) {
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

	if !dara.IsNil(request.Configure) {
		query["Configure"] = request.Configure
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Read) {
		query["Read"] = request.Read
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.Vhost) {
		query["Vhost"] = request.Vhost
	}

	if !dara.IsNil(request.Write) {
		query["Write"] = request.Write
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOpenSourcePermission"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOpenSourcePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an open source permission.
//
// @param request - CreateOpenSourcePermissionRequest
//
// @return CreateOpenSourcePermissionResponse
func (client *Client) CreateOpenSourcePermission(request *CreateOpenSourcePermissionRequest) (_result *CreateOpenSourcePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOpenSourcePermissionResponse{}
	_body, _err := client.CreateOpenSourcePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// A queue is a buffer that stores messages. In ApsaraMQ for RabbitMQ, messages are sent to a specified exchange and then routed to a bound queue.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoDeleteState) {
		body["AutoDeleteState"] = request.AutoDeleteState
	}

	if !dara.IsNil(request.AutoExpireState) {
		body["AutoExpireState"] = request.AutoExpireState
	}

	if !dara.IsNil(request.DeadLetterExchange) {
		body["DeadLetterExchange"] = request.DeadLetterExchange
	}

	if !dara.IsNil(request.DeadLetterRoutingKey) {
		body["DeadLetterRoutingKey"] = request.DeadLetterRoutingKey
	}

	if !dara.IsNil(request.ExclusiveState) {
		body["ExclusiveState"] = request.ExclusiveState
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxLength) {
		body["MaxLength"] = request.MaxLength
	}

	if !dara.IsNil(request.MaximumPriority) {
		body["MaximumPriority"] = request.MaximumPriority
	}

	if !dara.IsNil(request.MessageTTL) {
		body["MessageTTL"] = request.MessageTTL
	}

	if !dara.IsNil(request.QueueName) {
		body["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.VirtualHost) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQueue"),
		Version:     dara.String("2019-12-12"),
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
// A queue is a buffer that stores messages. In ApsaraMQ for RabbitMQ, messages are sent to a specified exchange and then routed to a bound queue.
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
// Creates a vhost. A vhost is used to logically isolate resources. Each vhost manages its own exchanges, queues, and bindings. Applications can run on independent vhosts in a secure manner. This way, the business of an application is not affected by other applications. Before you connect producers and consumers to an ApsaraMQ for RabbitMQ instance, you must specify vhosts for the producers and consumers.
//
// @param request - CreateVirtualHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVirtualHostResponse
func (client *Client) CreateVirtualHostWithOptions(request *CreateVirtualHostRequest, runtime *dara.RuntimeOptions) (_result *CreateVirtualHostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.VirtualHost) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVirtualHost"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVirtualHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a vhost. A vhost is used to logically isolate resources. Each vhost manages its own exchanges, queues, and bindings. Applications can run on independent vhosts in a secure manner. This way, the business of an application is not affected by other applications. Before you connect producers and consumers to an ApsaraMQ for RabbitMQ instance, you must specify vhosts for the producers and consumers.
//
// @param request - CreateVirtualHostRequest
//
// @return CreateVirtualHostResponse
func (client *Client) CreateVirtualHost(request *CreateVirtualHostRequest) (_result *CreateVirtualHostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVirtualHostResponse{}
	_body, _err := client.CreateVirtualHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a static username and password.
//
// @param request - DeleteAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccountWithOptions(request *DeleteAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimestamp) {
		query["CreateTimestamp"] = request.CreateTimestamp
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccount"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a static username and password.
//
// @param request - DeleteAccountRequest
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccount(request *DeleteAccountRequest) (_result *DeleteAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DeleteAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The DeleteBinding operation detaches a source exchange from a target queue or another exchange.
//
// @param request - DeleteBindingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBindingResponse
func (client *Client) DeleteBindingWithOptions(request *DeleteBindingRequest, runtime *dara.RuntimeOptions) (_result *DeleteBindingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BindingKey) {
		body["BindingKey"] = request.BindingKey
	}

	if !dara.IsNil(request.BindingType) {
		body["BindingType"] = request.BindingType
	}

	if !dara.IsNil(request.DestinationName) {
		body["DestinationName"] = request.DestinationName
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SourceExchange) {
		body["SourceExchange"] = request.SourceExchange
	}

	if !dara.IsNil(request.VirtualHost) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBinding"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The DeleteBinding operation detaches a source exchange from a target queue or another exchange.
//
// @param request - DeleteBindingRequest
//
// @return DeleteBindingResponse
func (client *Client) DeleteBinding(request *DeleteBindingRequest) (_result *DeleteBindingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBindingResponse{}
	_body, _err := client.DeleteBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an exchange.
//
// Description:
//
// ## Usage notes
//
// - You cannot delete exchanges with the type **headers**.
//
// - You cannot delete the three built-in exchanges in a vhost: amq.direct, amq.topic, or amq.fanout.
//
// @param request - DeleteExchangeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExchangeResponse
func (client *Client) DeleteExchangeWithOptions(request *DeleteExchangeRequest, runtime *dara.RuntimeOptions) (_result *DeleteExchangeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExchangeName) {
		body["ExchangeName"] = request.ExchangeName
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.VirtualHost) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExchange"),
		Version:     dara.String("2019-12-12"),
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
// Deletes an exchange.
//
// Description:
//
// ## Usage notes
//
// - You cannot delete exchanges with the type **headers**.
//
// - You cannot delete the three built-in exchanges in a vhost: amq.direct, amq.topic, or amq.fanout.
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
// Deletes the username and password of an open-source user.
//
// @param request - DeleteOpenSourceAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOpenSourceAccountResponse
func (client *Client) DeleteOpenSourceAccountWithOptions(request *DeleteOpenSourceAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteOpenSourceAccountResponse, _err error) {
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

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOpenSourceAccount"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOpenSourceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the username and password of an open-source user.
//
// @param request - DeleteOpenSourceAccountRequest
//
// @return DeleteOpenSourceAccountResponse
func (client *Client) DeleteOpenSourceAccount(request *DeleteOpenSourceAccountRequest) (_result *DeleteOpenSourceAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteOpenSourceAccountResponse{}
	_body, _err := client.DeleteOpenSourceAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an open source permission.
//
// @param request - DeleteOpenSourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOpenSourcePermissionResponse
func (client *Client) DeleteOpenSourcePermissionWithOptions(request *DeleteOpenSourcePermissionRequest, runtime *dara.RuntimeOptions) (_result *DeleteOpenSourcePermissionResponse, _err error) {
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

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.Vhost) {
		query["Vhost"] = request.Vhost
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOpenSourcePermission"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOpenSourcePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an open source permission.
//
// @param request - DeleteOpenSourcePermissionRequest
//
// @return DeleteOpenSourcePermissionResponse
func (client *Client) DeleteOpenSourcePermission(request *DeleteOpenSourcePermissionRequest) (_result *DeleteOpenSourcePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteOpenSourcePermissionResponse{}
	_body, _err := client.DeleteOpenSourcePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a queue.
//
// @param request - DeleteQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQueueResponse
func (client *Client) DeleteQueueWithOptions(request *DeleteQueueRequest, runtime *dara.RuntimeOptions) (_result *DeleteQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.QueueName) {
		body["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.VirtualHost) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQueue"),
		Version:     dara.String("2019-12-12"),
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
// Deletes a queue.
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
// Deletes a virtual host (vhost).
//
// Description:
//
// Before you delete a vhost, you must delete all exchanges and queues in it.
//
// @param request - DeleteVirtualHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVirtualHostResponse
func (client *Client) DeleteVirtualHostWithOptions(request *DeleteVirtualHostRequest, runtime *dara.RuntimeOptions) (_result *DeleteVirtualHostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.VirtualHost) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVirtualHost"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVirtualHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a virtual host (vhost).
//
// Description:
//
// Before you delete a vhost, you must delete all exchanges and queues in it.
//
// @param request - DeleteVirtualHostRequest
//
// @return DeleteVirtualHostResponse
func (client *Client) DeleteVirtualHost(request *DeleteVirtualHostRequest) (_result *DeleteVirtualHostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVirtualHostResponse{}
	_body, _err := client.DeleteVirtualHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of an ApsaraMQ for RabbitMQ instance.
//
// @param request - GetInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(request *GetInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
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
		Action:      dara.String("GetInstance"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an ApsaraMQ for RabbitMQ instance.
//
// @param request - GetInstanceRequest
//
// @return GetInstanceResponse
func (client *Client) GetInstance(request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the current and maximum numbers of vhosts, exchanges, and queues for a specified ApsaraMQ for RabbitMQ instance.
//
// @param request - GetMetadataAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetadataAmountResponse
func (client *Client) GetMetadataAmountWithOptions(request *GetMetadataAmountRequest, runtime *dara.RuntimeOptions) (_result *GetMetadataAmountResponse, _err error) {
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
		Action:      dara.String("GetMetadataAmount"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetadataAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the current and maximum numbers of vhosts, exchanges, and queues for a specified ApsaraMQ for RabbitMQ instance.
//
// @param request - GetMetadataAmountRequest
//
// @return GetMetadataAmountResponse
func (client *Client) GetMetadataAmount(request *GetMetadataAmountRequest) (_result *GetMetadataAmountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetadataAmountResponse{}
	_body, _err := client.GetMetadataAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the usernames and passwords for a specified ApsaraMQ for RabbitMQ instance.
//
// @param request - ListAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountsResponse
func (client *Client) ListAccountsWithOptions(request *ListAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListAccountsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccounts"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the usernames and passwords for a specified ApsaraMQ for RabbitMQ instance.
//
// @param request - ListAccountsRequest
//
// @return ListAccountsResponse
func (client *Client) ListAccounts(request *ListAccountsRequest) (_result *ListAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAccountsResponse{}
	_body, _err := client.ListAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all bindings that are created in a vhost of a specified ApsaraMQ for RabbitMQ instance.
//
// @param request - ListBindingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBindingsResponse
func (client *Client) ListBindingsWithOptions(request *ListBindingsRequest, runtime *dara.RuntimeOptions) (_result *ListBindingsResponse, _err error) {
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
		Action:      dara.String("ListBindings"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all bindings that are created in a vhost of a specified ApsaraMQ for RabbitMQ instance.
//
// @param request - ListBindingsRequest
//
// @return ListBindingsResponse
func (client *Client) ListBindings(request *ListBindingsRequest) (_result *ListBindingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBindingsResponse{}
	_body, _err := client.ListBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the exchanges or queues that are bound to a specified exchange.
//
// @param request - ListDownStreamBindingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDownStreamBindingsResponse
func (client *Client) ListDownStreamBindingsWithOptions(request *ListDownStreamBindingsRequest, runtime *dara.RuntimeOptions) (_result *ListDownStreamBindingsResponse, _err error) {
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
		Action:      dara.String("ListDownStreamBindings"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDownStreamBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the exchanges or queues that are bound to a specified exchange.
//
// @param request - ListDownStreamBindingsRequest
//
// @return ListDownStreamBindingsResponse
func (client *Client) ListDownStreamBindings(request *ListDownStreamBindingsRequest) (_result *ListDownStreamBindingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDownStreamBindingsResponse{}
	_body, _err := client.ListDownStreamBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the queues or other exchanges that are bound to a specified exchange.
//
// @param request - ListExchangeUpStreamBindingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExchangeUpStreamBindingsResponse
func (client *Client) ListExchangeUpStreamBindingsWithOptions(request *ListExchangeUpStreamBindingsRequest, runtime *dara.RuntimeOptions) (_result *ListExchangeUpStreamBindingsResponse, _err error) {
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
		Action:      dara.String("ListExchangeUpStreamBindings"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExchangeUpStreamBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the queues or other exchanges that are bound to a specified exchange.
//
// @param request - ListExchangeUpStreamBindingsRequest
//
// @return ListExchangeUpStreamBindingsResponse
func (client *Client) ListExchangeUpStreamBindings(request *ListExchangeUpStreamBindingsRequest) (_result *ListExchangeUpStreamBindingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExchangeUpStreamBindingsResponse{}
	_body, _err := client.ListExchangeUpStreamBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all exchanges in a specified vhost of an instance.
//
// @param request - ListExchangesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExchangesResponse
func (client *Client) ListExchangesWithOptions(request *ListExchangesRequest, runtime *dara.RuntimeOptions) (_result *ListExchangesResponse, _err error) {
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
		Action:      dara.String("ListExchanges"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExchangesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all exchanges in a specified vhost of an instance.
//
// @param request - ListExchangesRequest
//
// @return ListExchangesResponse
func (client *Client) ListExchanges(request *ListExchangesRequest) (_result *ListExchangesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExchangesResponse{}
	_body, _err := client.ListExchangesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IP or VPC whitelist for an instance.
//
// @param request - ListInstanceWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceWhiteListResponse
func (client *Client) ListInstanceWhiteListWithOptions(request *ListInstanceWhiteListRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceWhiteListResponse, _err error) {
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

	if !dara.IsNil(request.WhiteListType) {
		query["whiteListType"] = request.WhiteListType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceWhiteList"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP or VPC whitelist for an instance.
//
// @param request - ListInstanceWhiteListRequest
//
// @return ListInstanceWhiteListResponse
func (client *Client) ListInstanceWhiteList(request *ListInstanceWhiteListRequest) (_result *ListInstanceWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceWhiteListResponse{}
	_body, _err := client.ListInstanceWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of ApsaraMQ for RabbitMQ instances in a specified region and returns basic information about each instance, such as its endpoints and specification limits.
//
// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
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
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
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
// Queries a list of ApsaraMQ for RabbitMQ instances in a specified region and returns basic information about each instance, such as its endpoints and specification limits.
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
// Enumerates open-source usernames and passwords.
//
// @param request - ListOpenSourceAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOpenSourceAccountsResponse
func (client *Client) ListOpenSourceAccountsWithOptions(request *ListOpenSourceAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListOpenSourceAccountsResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOpenSourceAccounts"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOpenSourceAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enumerates open-source usernames and passwords.
//
// @param request - ListOpenSourceAccountsRequest
//
// @return ListOpenSourceAccountsResponse
func (client *Client) ListOpenSourceAccounts(request *ListOpenSourceAccountsRequest) (_result *ListOpenSourceAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOpenSourceAccountsResponse{}
	_body, _err := client.ListOpenSourceAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists open source permissions.
//
// @param request - ListOpenSourcePermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOpenSourcePermissionsResponse
func (client *Client) ListOpenSourcePermissionsWithOptions(request *ListOpenSourcePermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListOpenSourcePermissionsResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOpenSourcePermissions"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOpenSourcePermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists open source permissions.
//
// @param request - ListOpenSourcePermissionsRequest
//
// @return ListOpenSourcePermissionsResponse
func (client *Client) ListOpenSourcePermissions(request *ListOpenSourcePermissionsRequest) (_result *ListOpenSourcePermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOpenSourcePermissionsResponse{}
	_body, _err := client.ListOpenSourcePermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the online consumer clients of a specified queue.
//
// Description:
//
// ApsaraMQ for RabbitMQ lets you query only online consumer clients. You cannot query offline consumer clients.
//
// @param request - ListQueueConsumersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueueConsumersResponse
func (client *Client) ListQueueConsumersWithOptions(request *ListQueueConsumersRequest, runtime *dara.RuntimeOptions) (_result *ListQueueConsumersResponse, _err error) {
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
		Action:      dara.String("ListQueueConsumers"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQueueConsumersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the online consumer clients of a specified queue.
//
// Description:
//
// ApsaraMQ for RabbitMQ lets you query only online consumer clients. You cannot query offline consumer clients.
//
// @param request - ListQueueConsumersRequest
//
// @return ListQueueConsumersResponse
func (client *Client) ListQueueConsumers(request *ListQueueConsumersRequest) (_result *ListQueueConsumersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQueueConsumersResponse{}
	_body, _err := client.ListQueueConsumersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the exchanges that are bound to a specified queue.
//
// @param request - ListQueueUpStreamBindingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueueUpStreamBindingsResponse
func (client *Client) ListQueueUpStreamBindingsWithOptions(request *ListQueueUpStreamBindingsRequest, runtime *dara.RuntimeOptions) (_result *ListQueueUpStreamBindingsResponse, _err error) {
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
		Action:      dara.String("ListQueueUpStreamBindings"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQueueUpStreamBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the exchanges that are bound to a specified queue.
//
// @param request - ListQueueUpStreamBindingsRequest
//
// @return ListQueueUpStreamBindingsResponse
func (client *Client) ListQueueUpStreamBindings(request *ListQueueUpStreamBindingsRequest) (_result *ListQueueUpStreamBindingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQueueUpStreamBindingsResponse{}
	_body, _err := client.ListQueueUpStreamBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about all queues in a vhost on a specified ApsaraMQ for RabbitMQ instance.
//
// @param request - ListQueuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueuesResponse
func (client *Client) ListQueuesWithOptions(request *ListQueuesRequest, runtime *dara.RuntimeOptions) (_result *ListQueuesResponse, _err error) {
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
		Action:      dara.String("ListQueues"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQueuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about all queues in a vhost on a specified ApsaraMQ for RabbitMQ instance.
//
// @param request - ListQueuesRequest
//
// @return ListQueuesResponse
func (client *Client) ListQueues(request *ListQueuesRequest) (_result *ListQueuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQueuesResponse{}
	_body, _err := client.ListQueuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all vhosts in a specified ApsaraMQ for RabbitMQ instance.
//
// @param request - ListVirtualHostsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVirtualHostsResponse
func (client *Client) ListVirtualHostsWithOptions(request *ListVirtualHostsRequest, runtime *dara.RuntimeOptions) (_result *ListVirtualHostsResponse, _err error) {
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
		Action:      dara.String("ListVirtualHosts"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVirtualHostsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all vhosts in a specified ApsaraMQ for RabbitMQ instance.
//
// @param request - ListVirtualHostsRequest
//
// @return ListVirtualHostsResponse
func (client *Client) ListVirtualHosts(request *ListVirtualHostsRequest) (_result *ListVirtualHostsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVirtualHostsResponse{}
	_body, _err := client.ListVirtualHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes an IP address or a VPC from an instance\\"s whitelist.
//
// @param request - RemoveInstanceWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveInstanceWhiteListResponse
func (client *Client) RemoveInstanceWhiteListWithOptions(request *RemoveInstanceWhiteListRequest, runtime *dara.RuntimeOptions) (_result *RemoveInstanceWhiteListResponse, _err error) {
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

	if !dara.IsNil(request.WhiteListItemId) {
		query["whiteListItemId"] = request.WhiteListItemId
	}

	if !dara.IsNil(request.WhiteListType) {
		query["whiteListType"] = request.WhiteListType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveInstanceWhiteList"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveInstanceWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an IP address or a VPC from an instance\\"s whitelist.
//
// @param request - RemoveInstanceWhiteListRequest
//
// @return RemoveInstanceWhiteListResponse
func (client *Client) RemoveInstanceWhiteList(request *RemoveInstanceWhiteListRequest) (_result *RemoveInstanceWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveInstanceWhiteListResponse{}
	_body, _err := client.RemoveInstanceWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades or downgrades the specifications of an ApsaraMQ for RabbitMQ instance.
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
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.EncryptedInstance) {
		query["EncryptedInstance"] = request.EncryptedInstance
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.KmsKeyId) {
		query["KmsKeyId"] = request.KmsKeyId
	}

	if !dara.IsNil(request.MaxConnections) {
		query["MaxConnections"] = request.MaxConnections
	}

	if !dara.IsNil(request.MaxEipTps) {
		query["MaxEipTps"] = request.MaxEipTps
	}

	if !dara.IsNil(request.MaxPrivateTps) {
		query["MaxPrivateTps"] = request.MaxPrivateTps
	}

	if !dara.IsNil(request.ModifyType) {
		query["ModifyType"] = request.ModifyType
	}

	if !dara.IsNil(request.ProvisionedCapacity) {
		query["ProvisionedCapacity"] = request.ProvisionedCapacity
	}

	if !dara.IsNil(request.QueueCapacity) {
		query["QueueCapacity"] = request.QueueCapacity
	}

	if !dara.IsNil(request.ServerlessChargeType) {
		query["ServerlessChargeType"] = request.ServerlessChargeType
	}

	if !dara.IsNil(request.StorageSize) {
		query["StorageSize"] = request.StorageSize
	}

	if !dara.IsNil(request.SupportEip) {
		query["SupportEip"] = request.SupportEip
	}

	if !dara.IsNil(request.SupportTracing) {
		query["SupportTracing"] = request.SupportTracing
	}

	if !dara.IsNil(request.TracingStorageTime) {
		query["TracingStorageTime"] = request.TracingStorageTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstance"),
		Version:     dara.String("2019-12-12"),
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
// Upgrades or downgrades the specifications of an ApsaraMQ for RabbitMQ instance.
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
// An ApsaraMQ for RabbitMQ instance is named after its instance ID by default. You can change the name for easier identification.
//
// @param request - UpdateInstanceNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceNameResponse
func (client *Client) UpdateInstanceNameWithOptions(request *UpdateInstanceNameRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceNameResponse, _err error) {
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

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceName"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// An ApsaraMQ for RabbitMQ instance is named after its instance ID by default. You can change the name for easier identification.
//
// @param request - UpdateInstanceNameRequest
//
// @return UpdateInstanceNameResponse
func (client *Client) UpdateInstanceName(request *UpdateInstanceNameRequest) (_result *UpdateInstanceNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInstanceNameResponse{}
	_body, _err := client.UpdateInstanceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update serverless switch
//
// @param request - UpdateInstanceServerlessSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceServerlessSwitchResponse
func (client *Client) UpdateInstanceServerlessSwitchWithOptions(request *UpdateInstanceServerlessSwitchRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceServerlessSwitchResponse, _err error) {
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
		Action:      dara.String("UpdateInstanceServerlessSwitch"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceServerlessSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update serverless switch
//
// @param request - UpdateInstanceServerlessSwitchRequest
//
// @return UpdateInstanceServerlessSwitchResponse
func (client *Client) UpdateInstanceServerlessSwitch(request *UpdateInstanceServerlessSwitchRequest) (_result *UpdateInstanceServerlessSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInstanceServerlessSwitchResponse{}
	_body, _err := client.UpdateInstanceServerlessSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the username and password for open-source access.
//
// @param request - UpdateOpenSourceAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOpenSourceAccountResponse
func (client *Client) UpdateOpenSourceAccountWithOptions(request *UpdateOpenSourceAccountRequest, runtime *dara.RuntimeOptions) (_result *UpdateOpenSourceAccountResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOpenSourceAccount"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOpenSourceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the username and password for open-source access.
//
// @param request - UpdateOpenSourceAccountRequest
//
// @return UpdateOpenSourceAccountResponse
func (client *Client) UpdateOpenSourceAccount(request *UpdateOpenSourceAccountRequest) (_result *UpdateOpenSourceAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOpenSourceAccountResponse{}
	_body, _err := client.UpdateOpenSourceAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates open source permissions.
//
// @param request - UpdateOpenSourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOpenSourcePermissionResponse
func (client *Client) UpdateOpenSourcePermissionWithOptions(request *UpdateOpenSourcePermissionRequest, runtime *dara.RuntimeOptions) (_result *UpdateOpenSourcePermissionResponse, _err error) {
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

	if !dara.IsNil(request.Configure) {
		query["Configure"] = request.Configure
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Read) {
		query["Read"] = request.Read
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.Vhost) {
		query["Vhost"] = request.Vhost
	}

	if !dara.IsNil(request.Write) {
		query["Write"] = request.Write
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOpenSourcePermission"),
		Version:     dara.String("2019-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOpenSourcePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates open source permissions.
//
// @param request - UpdateOpenSourcePermissionRequest
//
// @return UpdateOpenSourcePermissionResponse
func (client *Client) UpdateOpenSourcePermission(request *UpdateOpenSourcePermissionRequest) (_result *UpdateOpenSourcePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOpenSourcePermissionResponse{}
	_body, _err := client.UpdateOpenSourcePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
