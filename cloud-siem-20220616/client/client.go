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
	client.Endpoint, _err = client.GetEndpoint(dara.String("cloud-siem"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds a data source to a cloud account that is added to the threat analysis feature.
//
// @param request - AddDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDataSourceResponse
func (client *Client) AddDataSourceWithOptions(request *AddDataSourceRequest, runtime *dara.RuntimeOptions) (_result *AddDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		body["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.DataSourceInstanceName) {
		body["DataSourceInstanceName"] = request.DataSourceInstanceName
	}

	if !dara.IsNil(request.DataSourceInstanceParams) {
		body["DataSourceInstanceParams"] = request.DataSourceInstanceParams
	}

	if !dara.IsNil(request.DataSourceInstanceRemark) {
		body["DataSourceInstanceRemark"] = request.DataSourceInstanceRemark
	}

	if !dara.IsNil(request.DataSourceType) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDataSource"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a data source to a cloud account that is added to the threat analysis feature.
//
// @param request - AddDataSourceRequest
//
// @return AddDataSourceResponse
func (client *Client) AddDataSource(request *AddDataSourceRequest) (_result *AddDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDataSourceResponse{}
	_body, _err := client.AddDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds logs of a cloud account to the threat analysis feature.
//
// @param request - AddDataSourceLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDataSourceLogResponse
func (client *Client) AddDataSourceLogWithOptions(request *AddDataSourceLogRequest, runtime *dara.RuntimeOptions) (_result *AddDataSourceLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		body["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.DataSourceInstanceId) {
		body["DataSourceInstanceId"] = request.DataSourceInstanceId
	}

	if !dara.IsNil(request.DataSourceInstanceLogs) {
		body["DataSourceInstanceLogs"] = request.DataSourceInstanceLogs
	}

	if !dara.IsNil(request.LogCode) {
		body["LogCode"] = request.LogCode
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDataSourceLog"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDataSourceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds logs of a cloud account to the threat analysis feature.
//
// @param request - AddDataSourceLogRequest
//
// @return AddDataSourceLogResponse
func (client *Client) AddDataSourceLog(request *AddDataSourceLogRequest) (_result *AddDataSourceLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDataSourceLogResponse{}
	_body, _err := client.AddDataSourceLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds the logs of a cloud service within a cloud account to the threat analysis feature for alert and event anslysis.
//
// @param request - AddUserSourceLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserSourceLogConfigResponse
func (client *Client) AddUserSourceLogConfigWithOptions(request *AddUserSourceLogConfigRequest, runtime *dara.RuntimeOptions) (_result *AddUserSourceLogConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Deleted) {
		body["Deleted"] = request.Deleted
	}

	if !dara.IsNil(request.DisPlayLine) {
		body["DisPlayLine"] = request.DisPlayLine
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourceLogCode) {
		body["SourceLogCode"] = request.SourceLogCode
	}

	if !dara.IsNil(request.SourceLogInfo) {
		body["SourceLogInfo"] = request.SourceLogInfo
	}

	if !dara.IsNil(request.SourceProdCode) {
		body["SourceProdCode"] = request.SourceProdCode
	}

	if !dara.IsNil(request.SubUserId) {
		body["SubUserId"] = request.SubUserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserSourceLogConfig"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserSourceLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds the logs of a cloud service within a cloud account to the threat analysis feature for alert and event anslysis.
//
// @param request - AddUserSourceLogConfigRequest
//
// @return AddUserSourceLogConfigResponse
func (client *Client) AddUserSourceLogConfig(request *AddUserSourceLogConfigRequest) (_result *AddUserSourceLogConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUserSourceLogConfigResponse{}
	_body, _err := client.AddUserSourceLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a third-party cloud account that is displayed on the Multi-cloud assets tab of the Feature Settings page to the threat analysis feature.
//
// @param request - BindAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAccountResponse
func (client *Client) BindAccountWithOptions(request *BindAccountRequest, runtime *dara.RuntimeOptions) (_result *BindAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessId) {
		body["AccessId"] = request.AccessId
	}

	if !dara.IsNil(request.AccountId) {
		body["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAccount"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a third-party cloud account that is displayed on the Multi-cloud assets tab of the Feature Settings page to the threat analysis feature.
//
// @param request - BindAccountRequest
//
// @return BindAccountResponse
func (client *Client) BindAccount(request *BindAccountRequest) (_result *BindAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindAccountResponse{}
	_body, _err := client.BindAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the log delivery feature for a cloud service.
//
// @param request - CloseDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseDeliveryResponse
func (client *Client) CloseDeliveryWithOptions(request *CloseDeliveryRequest, runtime *dara.RuntimeOptions) (_result *CloseDeliveryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LogCode) {
		body["LogCode"] = request.LogCode
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseDelivery"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the log delivery feature for a cloud service.
//
// @param request - CloseDeliveryRequest
//
// @return CloseDeliveryResponse
func (client *Client) CloseDelivery(request *CloseDeliveryRequest) (_result *CloseDeliveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloseDeliveryResponse{}
	_body, _err := client.CloseDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the automated response rule with a specified ID.
//
// @param request - DeleteAutomateResponseConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutomateResponseConfigResponse
func (client *Client) DeleteAutomateResponseConfigWithOptions(request *DeleteAutomateResponseConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteAutomateResponseConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAutomateResponseConfig"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAutomateResponseConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the automated response rule with a specified ID.
//
// @param request - DeleteAutomateResponseConfigRequest
//
// @return DeleteAutomateResponseConfigResponse
func (client *Client) DeleteAutomateResponseConfig(request *DeleteAutomateResponseConfigRequest) (_result *DeleteAutomateResponseConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAutomateResponseConfigResponse{}
	_body, _err := client.DeleteAutomateResponseConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a third-party cloud account that is added to the threat analysis feature by using its AccessKey ID. You can add another cloud account based on your business requirements.
//
// @param request - DeleteBindAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBindAccountResponse
func (client *Client) DeleteBindAccountWithOptions(request *DeleteBindAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteBindAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessId) {
		body["AccessId"] = request.AccessId
	}

	if !dara.IsNil(request.AccountId) {
		body["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.BindId) {
		body["BindId"] = request.BindId
	}

	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBindAccount"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBindAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a third-party cloud account that is added to the threat analysis feature by using its AccessKey ID. You can add another cloud account based on your business requirements.
//
// @param request - DeleteBindAccountRequest
//
// @return DeleteBindAccountResponse
func (client *Client) DeleteBindAccount(request *DeleteBindAccountRequest) (_result *DeleteBindAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBindAccountResponse{}
	_body, _err := client.DeleteBindAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a rule by rule ID.
//
// @param request - DeleteCustomizeRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomizeRuleResponse
func (client *Client) DeleteCustomizeRuleWithOptions(request *DeleteCustomizeRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomizeRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomizeRule"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomizeRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a rule by rule ID.
//
// @param request - DeleteCustomizeRuleRequest
//
// @return DeleteCustomizeRuleResponse
func (client *Client) DeleteCustomizeRule(request *DeleteCustomizeRuleRequest) (_result *DeleteCustomizeRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomizeRuleResponse{}
	_body, _err := client.DeleteCustomizeRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a data source that is no longer required.
//
// @param request - DeleteDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSourceWithOptions(request *DeleteDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		body["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.DataSourceInstanceId) {
		body["DataSourceInstanceId"] = request.DataSourceInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataSource"),
		Version:     dara.String("2022-06-16"),
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
// Removes a data source that is no longer required.
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
// Removes a log.
//
// @param request - DeleteDataSourceLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceLogResponse
func (client *Client) DeleteDataSourceLogWithOptions(request *DeleteDataSourceLogRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataSourceLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		body["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.DataSourceInstanceId) {
		body["DataSourceInstanceId"] = request.DataSourceInstanceId
	}

	if !dara.IsNil(request.LogInstanceId) {
		body["LogInstanceId"] = request.LogInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataSourceLog"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataSourceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a log.
//
// @param request - DeleteDataSourceLogRequest
//
// @return DeleteDataSourceLogResponse
func (client *Client) DeleteDataSourceLog(request *DeleteDataSourceLogRequest) (_result *DeleteDataSourceLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataSourceLogResponse{}
	_body, _err := client.DeleteDataSourceLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an alert whitelist rule with a specified ID.
//
// @param request - DeleteWhiteRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWhiteRuleListResponse
func (client *Client) DeleteWhiteRuleListWithOptions(request *DeleteWhiteRuleListRequest, runtime *dara.RuntimeOptions) (_result *DeleteWhiteRuleListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWhiteRuleList"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWhiteRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert whitelist rule with a specified ID.
//
// @param request - DeleteWhiteRuleListRequest
//
// @return DeleteWhiteRuleListResponse
func (client *Client) DeleteWhiteRuleList(request *DeleteWhiteRuleListRequest) (_result *DeleteWhiteRuleListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWhiteRuleListResponse{}
	_body, _err := client.DeleteWhiteRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the aggregate functions that are supported for a custom rule.
//
// @param request - DescribeAggregateFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAggregateFunctionResponse
func (client *Client) DescribeAggregateFunctionWithOptions(request *DescribeAggregateFunctionRequest, runtime *dara.RuntimeOptions) (_result *DescribeAggregateFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAggregateFunction"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAggregateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the aggregate functions that are supported for a custom rule.
//
// @param request - DescribeAggregateFunctionRequest
//
// @return DescribeAggregateFunctionResponse
func (client *Client) DescribeAggregateFunction(request *DescribeAggregateFunctionRequest) (_result *DescribeAggregateFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAggregateFunctionResponse{}
	_body, _err := client.DescribeAggregateFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the scenarios in which an alert needs to be added to the whitelist.
//
// @param request - DescribeAlertSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertSceneResponse
func (client *Client) DescribeAlertSceneWithOptions(request *DescribeAlertSceneRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertScene"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the scenarios in which an alert needs to be added to the whitelist.
//
// @param request - DescribeAlertSceneRequest
//
// @return DescribeAlertSceneResponse
func (client *Client) DescribeAlertScene(request *DescribeAlertSceneRequest) (_result *DescribeAlertSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAlertSceneResponse{}
	_body, _err := client.DescribeAlertSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the scenarios and objects that can be added to an alert whitelist rule.
//
// @param request - DescribeAlertSceneByEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertSceneByEventResponse
func (client *Client) DescribeAlertSceneByEventWithOptions(request *DescribeAlertSceneByEventRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertSceneByEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertSceneByEvent"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertSceneByEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the scenarios and objects that can be added to an alert whitelist rule.
//
// @param request - DescribeAlertSceneByEventRequest
//
// @return DescribeAlertSceneByEventResponse
func (client *Client) DescribeAlertSceneByEvent(request *DescribeAlertSceneByEventRequest) (_result *DescribeAlertSceneByEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAlertSceneByEventResponse{}
	_body, _err := client.DescribeAlertSceneByEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries alert data sources.
//
// @param request - DescribeAlertSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertSourceResponse
func (client *Client) DescribeAlertSourceWithOptions(request *DescribeAlertSourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Level) {
		body["Level"] = request.Level
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertSource"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert data sources.
//
// @param request - DescribeAlertSourceRequest
//
// @return DescribeAlertSourceResponse
func (client *Client) DescribeAlertSource(request *DescribeAlertSourceRequest) (_result *DescribeAlertSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAlertSourceResponse{}
	_body, _err := client.DescribeAlertSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the data sources of the alert that is associated with an event.
//
// @param request - DescribeAlertSourceWithEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertSourceWithEventResponse
func (client *Client) DescribeAlertSourceWithEventWithOptions(request *DescribeAlertSourceWithEventRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertSourceWithEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertSourceWithEvent"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertSourceWithEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data sources of the alert that is associated with an event.
//
// @param request - DescribeAlertSourceWithEventRequest
//
// @return DescribeAlertSourceWithEventResponse
func (client *Client) DescribeAlertSourceWithEvent(request *DescribeAlertSourceWithEventRequest) (_result *DescribeAlertSourceWithEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAlertSourceWithEventResponse{}
	_body, _err := client.DescribeAlertSourceWithEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the threat types that you can select when you create a custom rule.
//
// @param request - DescribeAlertTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertTypeResponse
func (client *Client) DescribeAlertTypeWithOptions(request *DescribeAlertTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.RuleType) {
		body["RuleType"] = request.RuleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertType"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the threat types that you can select when you create a custom rule.
//
// @param request - DescribeAlertTypeRequest
//
// @return DescribeAlertTypeResponse
func (client *Client) DescribeAlertType(request *DescribeAlertTypeRequest) (_result *DescribeAlertTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAlertTypeResponse{}
	_body, _err := client.DescribeAlertTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries alerts within your account.
//
// @param request - DescribeAlertsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertsResponse
func (client *Client) DescribeAlertsWithOptions(request *DescribeAlertsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertName) {
		body["AlertName"] = request.AlertName
	}

	if !dara.IsNil(request.AlertStatus) {
		body["AlertStatus"] = request.AlertStatus
	}

	if !dara.IsNil(request.AlertTitle) {
		body["AlertTitle"] = request.AlertTitle
	}

	if !dara.IsNil(request.AlertType) {
		body["AlertType"] = request.AlertType
	}

	if !dara.IsNil(request.AlertUuid) {
		body["AlertUuid"] = request.AlertUuid
	}

	if !dara.IsNil(request.AssetId) {
		body["AssetId"] = request.AssetId
	}

	if !dara.IsNil(request.AssetName) {
		body["AssetName"] = request.AssetName
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.EntityName) {
		body["EntityName"] = request.EntityName
	}

	if !dara.IsNil(request.IsDefend) {
		body["IsDefend"] = request.IsDefend
	}

	if !dara.IsNil(request.LabelType) {
		body["LabelType"] = request.LabelType
	}

	if !dara.IsNil(request.Level) {
		body["Level"] = request.Level
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.Source) {
		body["Source"] = request.Source
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.SubUserId) {
		body["SubUserId"] = request.SubUserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlerts"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alerts within your account.
//
// @param request - DescribeAlertsRequest
//
// @return DescribeAlertsResponse
func (client *Client) DescribeAlerts(request *DescribeAlertsRequest) (_result *DescribeAlertsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAlertsResponse{}
	_body, _err := client.DescribeAlertsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of alerts of different severities.
//
// @param request - DescribeAlertsCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertsCountResponse
func (client *Client) DescribeAlertsCountWithOptions(request *DescribeAlertsCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertsCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.QueryType) {
		body["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertsCount"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertsCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of alerts of different severities.
//
// @param request - DescribeAlertsCountRequest
//
// @return DescribeAlertsCountResponse
func (client *Client) DescribeAlertsCount(request *DescribeAlertsCountRequest) (_result *DescribeAlertsCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAlertsCountResponse{}
	_body, _err := client.DescribeAlertsCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the alerts that are associated with an entity.
//
// @param request - DescribeAlertsWithEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertsWithEntityResponse
func (client *Client) DescribeAlertsWithEntityWithOptions(request *DescribeAlertsWithEntityRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertsWithEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.EntityUuid) {
		body["EntityUuid"] = request.EntityUuid
	}

	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.SophonTaskId) {
		body["SophonTaskId"] = request.SophonTaskId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertsWithEntity"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertsWithEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alerts that are associated with an entity.
//
// @param request - DescribeAlertsWithEntityRequest
//
// @return DescribeAlertsWithEntityResponse
func (client *Client) DescribeAlertsWithEntity(request *DescribeAlertsWithEntityRequest) (_result *DescribeAlertsWithEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAlertsWithEntityResponse{}
	_body, _err := client.DescribeAlertsWithEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the alerts that are associated with an event.
//
// @param request - DescribeAlertsWithEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertsWithEventResponse
func (client *Client) DescribeAlertsWithEventWithOptions(request *DescribeAlertsWithEventRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertsWithEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertName) {
		body["AlertName"] = request.AlertName
	}

	if !dara.IsNil(request.AlertTitle) {
		body["AlertTitle"] = request.AlertTitle
	}

	if !dara.IsNil(request.AlertType) {
		body["AlertType"] = request.AlertType
	}

	if !dara.IsNil(request.AssetId) {
		body["AssetId"] = request.AssetId
	}

	if !dara.IsNil(request.AssetName) {
		body["AssetName"] = request.AssetName
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.EntityName) {
		body["EntityName"] = request.EntityName
	}

	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.IsDefend) {
		body["IsDefend"] = request.IsDefend
	}

	if !dara.IsNil(request.Level) {
		body["Level"] = request.Level
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.Source) {
		body["Source"] = request.Source
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.SubUserId) {
		body["SubUserId"] = request.SubUserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertsWithEvent"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertsWithEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alerts that are associated with an event.
//
// @param request - DescribeAlertsWithEventRequest
//
// @return DescribeAlertsWithEventResponse
func (client *Client) DescribeAlertsWithEvent(request *DescribeAlertsWithEventRequest) (_result *DescribeAlertsWithEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAlertsWithEventResponse{}
	_body, _err := client.DescribeAlertsWithEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether the security information and event management (SIEM) system is granted the required permissions to access other cloud resources within your Alibaba Cloud account and whether the AliyunServiceRoleForSasCloudSiem service-linked role is created.
//
// @param request - DescribeAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuthResponse
func (client *Client) DescribeAuthWithOptions(request *DescribeAuthRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuth"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the security information and event management (SIEM) system is granted the required permissions to access other cloud resources within your Alibaba Cloud account and whether the AliyunServiceRoleForSasCloudSiem service-linked role is created.
//
// @param request - DescribeAuthRequest
//
// @return DescribeAuthResponse
func (client *Client) DescribeAuth(request *DescribeAuthRequest) (_result *DescribeAuthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAuthResponse{}
	_body, _err := client.DescribeAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of automated response rules.
//
// @param request - DescribeAutomateResponseConfigCounterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutomateResponseConfigCounterResponse
func (client *Client) DescribeAutomateResponseConfigCounterWithOptions(request *DescribeAutomateResponseConfigCounterRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutomateResponseConfigCounterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAutomateResponseConfigCounter"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAutomateResponseConfigCounterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of automated response rules.
//
// @param request - DescribeAutomateResponseConfigCounterRequest
//
// @return DescribeAutomateResponseConfigCounterResponse
func (client *Client) DescribeAutomateResponseConfigCounter(request *DescribeAutomateResponseConfigCounterRequest) (_result *DescribeAutomateResponseConfigCounterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAutomateResponseConfigCounterResponse{}
	_body, _err := client.DescribeAutomateResponseConfigCounterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurable fields and operators of an automated response rule.
//
// @param request - DescribeAutomateResponseConfigFeatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutomateResponseConfigFeatureResponse
func (client *Client) DescribeAutomateResponseConfigFeatureWithOptions(request *DescribeAutomateResponseConfigFeatureRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutomateResponseConfigFeatureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoResponseType) {
		body["AutoResponseType"] = request.AutoResponseType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAutomateResponseConfigFeature"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAutomateResponseConfigFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurable fields and operators of an automated response rule.
//
// @param request - DescribeAutomateResponseConfigFeatureRequest
//
// @return DescribeAutomateResponseConfigFeatureResponse
func (client *Client) DescribeAutomateResponseConfigFeature(request *DescribeAutomateResponseConfigFeatureRequest) (_result *DescribeAutomateResponseConfigFeatureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAutomateResponseConfigFeatureResponse{}
	_body, _err := client.DescribeAutomateResponseConfigFeatureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the assets that are associated with an event.
//
// @param request - DescribeCloudSiemAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudSiemAssetsResponse
func (client *Client) DescribeCloudSiemAssetsWithOptions(request *DescribeCloudSiemAssetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudSiemAssetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssetName) {
		body["AssetName"] = request.AssetName
	}

	if !dara.IsNil(request.AssetType) {
		body["AssetType"] = request.AssetType
	}

	if !dara.IsNil(request.AssetUuid) {
		body["AssetUuid"] = request.AssetUuid
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudSiemAssets"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudSiemAssetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the assets that are associated with an event.
//
// @param request - DescribeCloudSiemAssetsRequest
//
// @return DescribeCloudSiemAssetsResponse
func (client *Client) DescribeCloudSiemAssets(request *DescribeCloudSiemAssetsRequest) (_result *DescribeCloudSiemAssetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudSiemAssetsResponse{}
	_body, _err := client.DescribeCloudSiemAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of assets that are associated with an event by asset type.
//
// @param request - DescribeCloudSiemAssetsCounterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudSiemAssetsCounterResponse
func (client *Client) DescribeCloudSiemAssetsCounterWithOptions(request *DescribeCloudSiemAssetsCounterRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudSiemAssetsCounterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudSiemAssetsCounter"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudSiemAssetsCounterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of assets that are associated with an event by asset type.
//
// @param request - DescribeCloudSiemAssetsCounterRequest
//
// @return DescribeCloudSiemAssetsCounterResponse
func (client *Client) DescribeCloudSiemAssetsCounter(request *DescribeCloudSiemAssetsCounterRequest) (_result *DescribeCloudSiemAssetsCounterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudSiemAssetsCounterResponse{}
	_body, _err := client.DescribeCloudSiemAssetsCounterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an event.
//
// @param request - DescribeCloudSiemEventDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudSiemEventDetailResponse
func (client *Client) DescribeCloudSiemEventDetailWithOptions(request *DescribeCloudSiemEventDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudSiemEventDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudSiemEventDetail"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudSiemEventDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an event.
//
// @param request - DescribeCloudSiemEventDetailRequest
//
// @return DescribeCloudSiemEventDetailResponse
func (client *Client) DescribeCloudSiemEventDetail(request *DescribeCloudSiemEventDetailRequest) (_result *DescribeCloudSiemEventDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudSiemEventDetailResponse{}
	_body, _err := client.DescribeCloudSiemEventDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries events in SIEM.
//
// @param request - DescribeCloudSiemEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudSiemEventsResponse
func (client *Client) DescribeCloudSiemEventsWithOptions(request *DescribeCloudSiemEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudSiemEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssetId) {
		body["AssetId"] = request.AssetId
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EntityUuid) {
		body["EntityUuid"] = request.EntityUuid
	}

	if !dara.IsNil(request.EventName) {
		body["EventName"] = request.EventName
	}

	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.Order) {
		body["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderField) {
		body["OrderField"] = request.OrderField
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.ThreadLevel) {
		body["ThreadLevel"] = request.ThreadLevel
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudSiemEvents"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudSiemEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries events in SIEM.
//
// @param request - DescribeCloudSiemEventsRequest
//
// @return DescribeCloudSiemEventsResponse
func (client *Client) DescribeCloudSiemEvents(request *DescribeCloudSiemEventsRequest) (_result *DescribeCloudSiemEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudSiemEventsResponse{}
	_body, _err := client.DescribeCloudSiemEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of custom rules.
//
// @param request - DescribeCustomizeRuleCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomizeRuleCountResponse
func (client *Client) DescribeCustomizeRuleCountWithOptions(request *DescribeCustomizeRuleCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomizeRuleCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomizeRuleCount"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomizeRuleCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of custom rules.
//
// @param request - DescribeCustomizeRuleCountRequest
//
// @return DescribeCustomizeRuleCountResponse
func (client *Client) DescribeCustomizeRuleCount(request *DescribeCustomizeRuleCountRequest) (_result *DescribeCustomizeRuleCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomizeRuleCountResponse{}
	_body, _err := client.DescribeCustomizeRuleCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the historical simulation data that is used in a simulation test scenario.
//
// @param request - DescribeCustomizeRuleTestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomizeRuleTestResponse
func (client *Client) DescribeCustomizeRuleTestWithOptions(request *DescribeCustomizeRuleTestRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomizeRuleTestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomizeRuleTest"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomizeRuleTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical simulation data that is used in a simulation test scenario.
//
// @param request - DescribeCustomizeRuleTestRequest
//
// @return DescribeCustomizeRuleTestResponse
func (client *Client) DescribeCustomizeRuleTest(request *DescribeCustomizeRuleTestRequest) (_result *DescribeCustomizeRuleTestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomizeRuleTestResponse{}
	_body, _err := client.DescribeCustomizeRuleTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the chart that displays the test results of business data for a custom rule.
//
// @param request - DescribeCustomizeRuleTestHistogramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomizeRuleTestHistogramResponse
func (client *Client) DescribeCustomizeRuleTestHistogramWithOptions(request *DescribeCustomizeRuleTestHistogramRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomizeRuleTestHistogramResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomizeRuleTestHistogram"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomizeRuleTestHistogramResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the chart that displays the test results of business data for a custom rule.
//
// @param request - DescribeCustomizeRuleTestHistogramRequest
//
// @return DescribeCustomizeRuleTestHistogramResponse
func (client *Client) DescribeCustomizeRuleTestHistogram(request *DescribeCustomizeRuleTestHistogramRequest) (_result *DescribeCustomizeRuleTestHistogramResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomizeRuleTestHistogramResponse{}
	_body, _err := client.DescribeCustomizeRuleTestHistogramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a data source.
//
// @param request - DescribeDataSourceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataSourceInstanceResponse
func (client *Client) DescribeDataSourceInstanceWithOptions(request *DescribeDataSourceInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataSourceInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		body["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.DataSourceInstanceId) {
		body["DataSourceInstanceId"] = request.DataSourceInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataSourceInstance"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataSourceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a data source.
//
// @param request - DescribeDataSourceInstanceRequest
//
// @return DescribeDataSourceInstanceResponse
func (client *Client) DescribeDataSourceInstance(request *DescribeDataSourceInstanceRequest) (_result *DescribeDataSourceInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataSourceInstanceResponse{}
	_body, _err := client.DescribeDataSourceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the parameters of a data source.
//
// @param request - DescribeDataSourceParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataSourceParametersResponse
func (client *Client) DescribeDataSourceParametersWithOptions(request *DescribeDataSourceParametersRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataSourceParametersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.DataSourceType) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataSourceParameters"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataSourceParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the parameters of a data source.
//
// @param request - DescribeDataSourceParametersRequest
//
// @return DescribeDataSourceParametersResponse
func (client *Client) DescribeDataSourceParameters(request *DescribeDataSourceParametersRequest) (_result *DescribeDataSourceParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataSourceParametersResponse{}
	_body, _err := client.DescribeDataSourceParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of entities and playbooks that need to be handled.
//
// @param request - DescribeDisposeAndPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDisposeAndPlaybookResponse
func (client *Client) DescribeDisposeAndPlaybookWithOptions(request *DescribeDisposeAndPlaybookRequest, runtime *dara.RuntimeOptions) (_result *DescribeDisposeAndPlaybookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EntityType) {
		body["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.EntityUuid) {
		body["EntityUuid"] = request.EntityUuid
	}

	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDisposeAndPlaybook"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDisposeAndPlaybookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of entities and playbooks that need to be handled.
//
// @param request - DescribeDisposeAndPlaybookRequest
//
// @return DescribeDisposeAndPlaybookResponse
func (client *Client) DescribeDisposeAndPlaybook(request *DescribeDisposeAndPlaybookRequest) (_result *DescribeDisposeAndPlaybookResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDisposeAndPlaybookResponse{}
	_body, _err := client.DescribeDisposeAndPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of playbooks that are used by a handling policy.
//
// @param request - DescribeDisposeStrategyPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDisposeStrategyPlaybookResponse
func (client *Client) DescribeDisposeStrategyPlaybookWithOptions(request *DescribeDisposeStrategyPlaybookRequest, runtime *dara.RuntimeOptions) (_result *DescribeDisposeStrategyPlaybookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDisposeStrategyPlaybook"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDisposeStrategyPlaybookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of playbooks that are used by a handling policy.
//
// @param request - DescribeDisposeStrategyPlaybookRequest
//
// @return DescribeDisposeStrategyPlaybookResponse
func (client *Client) DescribeDisposeStrategyPlaybook(request *DescribeDisposeStrategyPlaybookRequest) (_result *DescribeDisposeStrategyPlaybookResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDisposeStrategyPlaybookResponse{}
	_body, _err := client.DescribeDisposeStrategyPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an entity.
//
// @param request - DescribeEntityInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEntityInfoResponse
func (client *Client) DescribeEntityInfoWithOptions(request *DescribeEntityInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeEntityInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.EntityIdentity) {
		body["EntityIdentity"] = request.EntityIdentity
	}

	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.SophonTaskId) {
		body["SophonTaskId"] = request.SophonTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEntityInfo"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEntityInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an entity.
//
// @param request - DescribeEntityInfoRequest
//
// @return DescribeEntityInfoResponse
func (client *Client) DescribeEntityInfo(request *DescribeEntityInfoRequest) (_result *DescribeEntityInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEntityInfoResponse{}
	_body, _err := client.DescribeEntityInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of events by type.
//
// @param request - DescribeEventCountByThreatLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventCountByThreatLevelResponse
func (client *Client) DescribeEventCountByThreatLevelWithOptions(request *DescribeEventCountByThreatLevelRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventCountByThreatLevelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventCountByThreatLevel"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventCountByThreatLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of events by type.
//
// @param request - DescribeEventCountByThreatLevelRequest
//
// @return DescribeEventCountByThreatLevelResponse
func (client *Client) DescribeEventCountByThreatLevel(request *DescribeEventCountByThreatLevelRequest) (_result *DescribeEventCountByThreatLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventCountByThreatLevelResponse{}
	_body, _err := client.DescribeEventCountByThreatLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the handling policies of a historical event.
//
// @param request - DescribeEventDisposeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventDisposeResponse
func (client *Client) DescribeEventDisposeWithOptions(request *DescribeEventDisposeRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventDisposeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventDispose"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventDisposeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the handling policies of a historical event.
//
// @param request - DescribeEventDisposeRequest
//
// @return DescribeEventDisposeResponse
func (client *Client) DescribeEventDispose(request *DescribeEventDisposeRequest) (_result *DescribeEventDisposeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventDisposeResponse{}
	_body, _err := client.DescribeEventDisposeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of logs that are added to the threat analysis feature.
//
// @param request - DescribeImportedLogCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImportedLogCountResponse
func (client *Client) DescribeImportedLogCountWithOptions(request *DescribeImportedLogCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeImportedLogCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeImportedLogCount"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImportedLogCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of logs that are added to the threat analysis feature.
//
// @param request - DescribeImportedLogCountRequest
//
// @return DescribeImportedLogCountResponse
func (client *Client) DescribeImportedLogCount(request *DescribeImportedLogCountRequest) (_result *DescribeImportedLogCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeImportedLogCountResponse{}
	_body, _err := client.DescribeImportedLogCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the fields that can be configured for a custom rule.
//
// @param request - DescribeLogFieldsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogFieldsResponse
func (client *Client) DescribeLogFieldsWithOptions(request *DescribeLogFieldsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogFieldsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LogSource) {
		body["LogSource"] = request.LogSource
	}

	if !dara.IsNil(request.LogType) {
		body["LogType"] = request.LogType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLogFields"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the fields that can be configured for a custom rule.
//
// @param request - DescribeLogFieldsRequest
//
// @return DescribeLogFieldsResponse
func (client *Client) DescribeLogFields(request *DescribeLogFieldsRequest) (_result *DescribeLogFieldsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLogFieldsResponse{}
	_body, _err := client.DescribeLogFieldsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the log sources that can be configured for a custom rule.
//
// @param request - DescribeLogSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogSourceResponse
func (client *Client) DescribeLogSourceWithOptions(request *DescribeLogSourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LogType) {
		body["LogType"] = request.LogType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLogSource"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log sources that can be configured for a custom rule.
//
// @param request - DescribeLogSourceRequest
//
// @return DescribeLogSourceResponse
func (client *Client) DescribeLogSource(request *DescribeLogSourceRequest) (_result *DescribeLogSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLogSourceResponse{}
	_body, _err := client.DescribeLogSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the log types that can be configured for a custom rule.
//
// @param request - DescribeLogTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogTypeResponse
func (client *Client) DescribeLogTypeWithOptions(request *DescribeLogTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLogType"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log types that can be configured for a custom rule.
//
// @param request - DescribeLogTypeRequest
//
// @return DescribeLogTypeResponse
func (client *Client) DescribeLogType(request *DescribeLogTypeRequest) (_result *DescribeLogTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLogTypeResponse{}
	_body, _err := client.DescribeLogTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operator of a custom rule.
//
// @param request - DescribeOperatorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOperatorsResponse
func (client *Client) DescribeOperatorsWithOptions(request *DescribeOperatorsRequest, runtime *dara.RuntimeOptions) (_result *DescribeOperatorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.SceneType) {
		body["SceneType"] = request.SceneType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOperators"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOperatorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operator of a custom rule.
//
// @param request - DescribeOperatorsRequest
//
// @return DescribeOperatorsResponse
func (client *Client) DescribeOperators(request *DescribeOperatorsRequest) (_result *DescribeOperatorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOperatorsResponse{}
	_body, _err := client.DescribeOperatorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of services that can be added to the threat analysis feature in Alibaba Cloud, Tenant Cloud, and Huawei Cloud.
//
// @param request - DescribeProdCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProdCountResponse
func (client *Client) DescribeProdCountWithOptions(request *DescribeProdCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeProdCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProdCount"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProdCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of services that can be added to the threat analysis feature in Alibaba Cloud, Tenant Cloud, and Huawei Cloud.
//
// @param request - DescribeProdCountRequest
//
// @return DescribeProdCountResponse
func (client *Client) DescribeProdCount(request *DescribeProdCountRequest) (_result *DescribeProdCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProdCountResponse{}
	_body, _err := client.DescribeProdCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of users in the playbook scope.
//
// @param request - DescribeScopeUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScopeUsersResponse
func (client *Client) DescribeScopeUsersWithOptions(request *DescribeScopeUsersRequest, runtime *dara.RuntimeOptions) (_result *DescribeScopeUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScopeUsers"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScopeUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of users in the playbook scope.
//
// @param request - DescribeScopeUsersRequest
//
// @return DescribeScopeUsersResponse
func (client *Client) DescribeScopeUsers(request *DescribeScopeUsersRequest) (_result *DescribeScopeUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeScopeUsersResponse{}
	_body, _err := client.DescribeScopeUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether the threat analysis feature is authorized to access a resource directory.
//
// @param request - DescribeServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceStatusResponse
func (client *Client) DescribeServiceStatusWithOptions(request *DescribeServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceStatus"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the threat analysis feature is authorized to access a resource directory.
//
// @param request - DescribeServiceStatusRequest
//
// @return DescribeServiceStatusResponse
func (client *Client) DescribeServiceStatus(request *DescribeServiceStatusRequest) (_result *DescribeServiceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeServiceStatusResponse{}
	_body, _err := client.DescribeServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the Logstores for the threat analysis feature in Simple Log Service on the user side.
//
// @param request - DescribeStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStorageResponse
func (client *Client) DescribeStorageWithOptions(request *DescribeStorageRequest, runtime *dara.RuntimeOptions) (_result *DescribeStorageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStorage"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the Logstores for the threat analysis feature in Simple Log Service on the user side.
//
// @param request - DescribeStorageRequest
//
// @return DescribeStorageResponse
func (client *Client) DescribeStorage(request *DescribeStorageRequest) (_result *DescribeStorageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeStorageResponse{}
	_body, _err := client.DescribeStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether the current Alibaba Cloud account or the management account of a resource directory is used to purchase the threat analysis feature.
//
// @param request - DescribeUserBuyStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserBuyStatusResponse
func (client *Client) DescribeUserBuyStatusWithOptions(request *DescribeUserBuyStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserBuyStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SubUserId) {
		body["SubUserId"] = request.SubUserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserBuyStatus"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserBuyStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the current Alibaba Cloud account or the management account of a resource directory is used to purchase the threat analysis feature.
//
// @param request - DescribeUserBuyStatusRequest
//
// @return DescribeUserBuyStatusResponse
func (client *Client) DescribeUserBuyStatus(request *DescribeUserBuyStatusRequest) (_result *DescribeUserBuyStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserBuyStatusResponse{}
	_body, _err := client.DescribeUserBuyStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the protected domain names of the WAF instance for a user to which an entity belongs.
//
// @param request - DescribeWafScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWafScopeResponse
func (client *Client) DescribeWafScopeWithOptions(request *DescribeWafScopeRequest, runtime *dara.RuntimeOptions) (_result *DescribeWafScopeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWafScope"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWafScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the protected domain names of the WAF instance for a user to which an entity belongs.
//
// @param request - DescribeWafScopeRequest
//
// @return DescribeWafScopeResponse
func (client *Client) DescribeWafScope(request *DescribeWafScopeRequest) (_result *DescribeWafScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWafScopeResponse{}
	_body, _err := client.DescribeWafScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of whitelist rules for alerts.
//
// @param request - DescribeWhiteRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWhiteRuleListResponse
func (client *Client) DescribeWhiteRuleListWithOptions(request *DescribeWhiteRuleListRequest, runtime *dara.RuntimeOptions) (_result *DescribeWhiteRuleListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertName) {
		body["AlertName"] = request.AlertName
	}

	if !dara.IsNil(request.AlertType) {
		body["AlertType"] = request.AlertType
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWhiteRuleList"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWhiteRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of whitelist rules for alerts.
//
// @param request - DescribeWhiteRuleListRequest
//
// @return DescribeWhiteRuleListResponse
func (client *Client) DescribeWhiteRuleList(request *DescribeWhiteRuleListRequest) (_result *DescribeWhiteRuleListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWhiteRuleListResponse{}
	_body, _err := client.DescribeWhiteRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a service-linked role named AliyunServiceRoleForSasCloudSiem for the threat analysis feature. The feature can assume this role to access cloud services.
//
// @param request - EnableAccessForCloudSiemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAccessForCloudSiemResponse
func (client *Client) EnableAccessForCloudSiemWithOptions(request *EnableAccessForCloudSiemRequest, runtime *dara.RuntimeOptions) (_result *EnableAccessForCloudSiemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoSubmit) {
		body["AutoSubmit"] = request.AutoSubmit
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableAccessForCloudSiem"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableAccessForCloudSiemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service-linked role named AliyunServiceRoleForSasCloudSiem for the threat analysis feature. The feature can assume this role to access cloud services.
//
// @param request - EnableAccessForCloudSiemRequest
//
// @return EnableAccessForCloudSiemResponse
func (client *Client) EnableAccessForCloudSiem(request *EnableAccessForCloudSiemRequest) (_result *EnableAccessForCloudSiemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableAccessForCloudSiemResponse{}
	_body, _err := client.EnableAccessForCloudSiemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Authorizes the threat analysis feature to access a resource directory. This operation must be called by the management account of the resource directory.
//
// @param request - EnableServiceForCloudSiemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableServiceForCloudSiemResponse
func (client *Client) EnableServiceForCloudSiemWithOptions(request *EnableServiceForCloudSiemRequest, runtime *dara.RuntimeOptions) (_result *EnableServiceForCloudSiemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableServiceForCloudSiem"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableServiceForCloudSiemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes the threat analysis feature to access a resource directory. This operation must be called by the management account of the resource directory.
//
// @param request - EnableServiceForCloudSiemRequest
//
// @return EnableServiceForCloudSiemResponse
func (client *Client) EnableServiceForCloudSiem(request *EnableServiceForCloudSiemRequest) (_result *EnableServiceForCloudSiemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableServiceForCloudSiemResponse{}
	_body, _err := client.EnableServiceForCloudSiemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Queries the storage capacity usage of the threat analysis feature and the purchased storage capacity
//
// @param request - GetCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCapacityResponse
func (client *Client) GetCapacityWithOptions(request *GetCapacityRequest, runtime *dara.RuntimeOptions) (_result *GetCapacityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCapacity"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCapacityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Queries the storage capacity usage of the threat analysis feature and the purchased storage capacity
//
// @param request - GetCapacityRequest
//
// @return GetCapacityResponse
func (client *Client) GetCapacity(request *GetCapacityRequest) (_result *GetCapacityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCapacityResponse{}
	_body, _err := client.GetCapacityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the storage configurations for the threat analysis feature on the user side.
//
// @param request - GetStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStorageResponse
func (client *Client) GetStorageWithOptions(request *GetStorageRequest, runtime *dara.RuntimeOptions) (_result *GetStorageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStorage"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the storage configurations for the threat analysis feature on the user side.
//
// @param request - GetStorageRequest
//
// @return GetStorageResponse
func (client *Client) GetStorage(request *GetStorageRequest) (_result *GetStorageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStorageResponse{}
	_body, _err := client.GetStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of AccessKey IDs of third-party cloud accounts that are added to the threat analysis feature.
//
// @param request - ListAccountAccessIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountAccessIdResponse
func (client *Client) ListAccountAccessIdWithOptions(request *ListAccountAccessIdRequest, runtime *dara.RuntimeOptions) (_result *ListAccountAccessIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccountAccessId"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccountAccessIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of AccessKey IDs of third-party cloud accounts that are added to the threat analysis feature.
//
// @param request - ListAccountAccessIdRequest
//
// @return ListAccountAccessIdResponse
func (client *Client) ListAccountAccessId(request *ListAccountAccessIdRequest) (_result *ListAccountAccessIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAccountAccessIdResponse{}
	_body, _err := client.ListAccountAccessIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query accounts by log.
//
// @param request - ListAccountsByLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountsByLogResponse
func (client *Client) ListAccountsByLogWithOptions(request *ListAccountsByLogRequest, runtime *dara.RuntimeOptions) (_result *ListAccountsByLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.LogCodes) {
		body["LogCodes"] = request.LogCodes
	}

	if !dara.IsNil(request.ProdCode) {
		body["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccountsByLog"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccountsByLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query accounts by log.
//
// @param request - ListAccountsByLogRequest
//
// @return ListAccountsByLogResponse
func (client *Client) ListAccountsByLog(request *ListAccountsByLogRequest) (_result *ListAccountsByLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAccountsByLogResponse{}
	_body, _err := client.ListAccountsByLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of cloud services that can be added to the threat analysis feature.
//
// @param request - ListAllProdsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllProdsResponse
func (client *Client) ListAllProdsWithOptions(request *ListAllProdsRequest, runtime *dara.RuntimeOptions) (_result *ListAllProdsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAllProds"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllProdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of cloud services that can be added to the threat analysis feature.
//
// @param request - ListAllProdsRequest
//
// @return ListAllProdsResponse
func (client *Client) ListAllProds(request *ListAllProdsRequest) (_result *ListAllProdsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAllProdsResponse{}
	_body, _err := client.ListAllProdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries automated response rules.
//
// @param request - ListAutomateResponseConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAutomateResponseConfigsResponse
func (client *Client) ListAutomateResponseConfigsWithOptions(request *ListAutomateResponseConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListAutomateResponseConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ActionType) {
		body["ActionType"] = request.ActionType
	}

	if !dara.IsNil(request.AutoResponseType) {
		body["AutoResponseType"] = request.AutoResponseType
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PlaybookUuid) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResponseRuleType) {
		body["ResponseRuleType"] = request.ResponseRuleType
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.RuleName) {
		body["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.SubUserId) {
		body["SubUserId"] = request.SubUserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAutomateResponseConfigs"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAutomateResponseConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries automated response rules.
//
// @param request - ListAutomateResponseConfigsRequest
//
// @return ListAutomateResponseConfigsResponse
func (client *Client) ListAutomateResponseConfigs(request *ListAutomateResponseConfigsRequest) (_result *ListAutomateResponseConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAutomateResponseConfigsResponse{}
	_body, _err := client.ListAutomateResponseConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of cloud accounts that are added to the threat analysis feature.
//
// @param request - ListBindAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBindAccountResponse
func (client *Client) ListBindAccountWithOptions(request *ListBindAccountRequest, runtime *dara.RuntimeOptions) (_result *ListBindAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBindAccount"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBindAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of cloud accounts that are added to the threat analysis feature.
//
// @param request - ListBindAccountRequest
//
// @return ListBindAccountResponse
func (client *Client) ListBindAccount(request *ListBindAccountRequest) (_result *ListBindAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBindAccountResponse{}
	_body, _err := client.ListBindAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of data sources that are added to the threat analysis feature.
//
// @param request - ListBindDataSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBindDataSourcesResponse
func (client *Client) ListBindDataSourcesWithOptions(request *ListBindDataSourcesRequest, runtime *dara.RuntimeOptions) (_result *ListBindDataSourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		body["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBindDataSources"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBindDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of data sources that are added to the threat analysis feature.
//
// @param request - ListBindDataSourcesRequest
//
// @return ListBindDataSourcesResponse
func (client *Client) ListBindDataSources(request *ListBindDataSourcesRequest) (_result *ListBindDataSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBindDataSourcesResponse{}
	_body, _err := client.ListBindDataSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries custom rules.
//
// @param request - ListCloudSiemCustomizeRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudSiemCustomizeRulesResponse
func (client *Client) ListCloudSiemCustomizeRulesWithOptions(request *ListCloudSiemCustomizeRulesRequest, runtime *dara.RuntimeOptions) (_result *ListCloudSiemCustomizeRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertType) {
		body["AlertType"] = request.AlertType
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Order) {
		body["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderField) {
		body["OrderField"] = request.OrderField
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.RuleName) {
		body["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleType) {
		body["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.ThreatLevel) {
		body["ThreatLevel"] = request.ThreatLevel
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudSiemCustomizeRules"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudSiemCustomizeRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom rules.
//
// @param request - ListCloudSiemCustomizeRulesRequest
//
// @return ListCloudSiemCustomizeRulesResponse
func (client *Client) ListCloudSiemCustomizeRules(request *ListCloudSiemCustomizeRulesRequest) (_result *ListCloudSiemCustomizeRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudSiemCustomizeRulesResponse{}
	_body, _err := client.ListCloudSiemCustomizeRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries predefined rules.
//
// @param request - ListCloudSiemPredefinedRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudSiemPredefinedRulesResponse
func (client *Client) ListCloudSiemPredefinedRulesWithOptions(request *ListCloudSiemPredefinedRulesRequest, runtime *dara.RuntimeOptions) (_result *ListCloudSiemPredefinedRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertType) {
		body["AlertType"] = request.AlertType
	}

	if !dara.IsNil(request.AttCk) {
		body["AttCk"] = request.AttCk
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventTransferType) {
		body["EventTransferType"] = request.EventTransferType
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.LogSource) {
		body["LogSource"] = request.LogSource
	}

	if !dara.IsNil(request.Order) {
		body["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderField) {
		body["OrderField"] = request.OrderField
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.RuleName) {
		body["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleType) {
		body["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.ThreatLevel) {
		body["ThreatLevel"] = request.ThreatLevel
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudSiemPredefinedRules"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudSiemPredefinedRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries predefined rules.
//
// @param request - ListCloudSiemPredefinedRulesRequest
//
// @return ListCloudSiemPredefinedRulesResponse
func (client *Client) ListCloudSiemPredefinedRules(request *ListCloudSiemPredefinedRulesRequest) (_result *ListCloudSiemPredefinedRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudSiemPredefinedRulesResponse{}
	_body, _err := client.ListCloudSiemPredefinedRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the test results of a custom rule.
//
// @param request - ListCustomizeRuleTestResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomizeRuleTestResultResponse
func (client *Client) ListCustomizeRuleTestResultWithOptions(request *ListCustomizeRuleTestResultRequest, runtime *dara.RuntimeOptions) (_result *ListCustomizeRuleTestResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DetectionRuleId) {
		body["DetectionRuleId"] = request.DetectionRuleId
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.VerifyType) {
		body["VerifyType"] = request.VerifyType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomizeRuleTestResult"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomizeRuleTestResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the test results of a custom rule.
//
// @param request - ListCustomizeRuleTestResultRequest
//
// @return ListCustomizeRuleTestResultResponse
func (client *Client) ListCustomizeRuleTestResult(request *ListCustomizeRuleTestResultRequest) (_result *ListCustomizeRuleTestResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCustomizeRuleTestResultResponse{}
	_body, _err := client.ListCustomizeRuleTestResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the logs of a data source.
//
// @param request - ListDataSourceLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceLogsResponse
func (client *Client) ListDataSourceLogsWithOptions(request *ListDataSourceLogsRequest, runtime *dara.RuntimeOptions) (_result *ListDataSourceLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		body["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.DataSourceInstanceId) {
		body["DataSourceInstanceId"] = request.DataSourceInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSourceLogs"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSourceLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logs of a data source.
//
// @param request - ListDataSourceLogsRequest
//
// @return ListDataSourceLogsResponse
func (client *Client) ListDataSourceLogs(request *ListDataSourceLogsRequest) (_result *ListDataSourceLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataSourceLogsResponse{}
	_body, _err := client.ListDataSourceLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of data source types in third-party cloud services that can be added to the threat analysis feature.
//
// @param request - ListDataSourceTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceTypesResponse
func (client *Client) ListDataSourceTypesWithOptions(request *ListDataSourceTypesRequest, runtime *dara.RuntimeOptions) (_result *ListDataSourceTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSourceTypes"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSourceTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of data source types in third-party cloud services that can be added to the threat analysis feature.
//
// @param request - ListDataSourceTypesRequest
//
// @return ListDataSourceTypesResponse
func (client *Client) ListDataSourceTypes(request *ListDataSourceTypesRequest) (_result *ListDataSourceTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataSourceTypesResponse{}
	_body, _err := client.ListDataSourceTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the cloud services that are integrated with the threat analysis feature, the logs of the cloud services, and the delivery of the logs.
//
// @param request - ListDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeliveryResponse
func (client *Client) ListDeliveryWithOptions(request *ListDeliveryRequest, runtime *dara.RuntimeOptions) (_result *ListDeliveryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDelivery"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the cloud services that are integrated with the threat analysis feature, the logs of the cloud services, and the delivery of the logs.
//
// @param request - ListDeliveryRequest
//
// @return ListDeliveryResponse
func (client *Client) ListDelivery(request *ListDeliveryRequest) (_result *ListDeliveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDeliveryResponse{}
	_body, _err := client.ListDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries handling policies.
//
// @param request - ListDisposeStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDisposeStrategyResponse
func (client *Client) ListDisposeStrategyWithOptions(request *ListDisposeStrategyRequest, runtime *dara.RuntimeOptions) (_result *ListDisposeStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EffectiveStatus) {
		body["EffectiveStatus"] = request.EffectiveStatus
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EntityIdentity) {
		body["EntityIdentity"] = request.EntityIdentity
	}

	if !dara.IsNil(request.EntityType) {
		body["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.Order) {
		body["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderField) {
		body["OrderField"] = request.OrderField
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PlaybookName) {
		body["PlaybookName"] = request.PlaybookName
	}

	if !dara.IsNil(request.PlaybookTypes) {
		body["PlaybookTypes"] = request.PlaybookTypes
	}

	if !dara.IsNil(request.PlaybookUuid) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.SophonTaskId) {
		body["SophonTaskId"] = request.SophonTaskId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDisposeStrategy"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDisposeStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries handling policies.
//
// @param request - ListDisposeStrategyRequest
//
// @return ListDisposeStrategyResponse
func (client *Client) ListDisposeStrategy(request *ListDisposeStrategyRequest) (_result *ListDisposeStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDisposeStrategyResponse{}
	_body, _err := client.ListDisposeStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries entities.
//
// @param request - ListEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEntitiesResponse
func (client *Client) ListEntitiesWithOptions(request *ListEntitiesRequest, runtime *dara.RuntimeOptions) (_result *ListEntitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EntityName) {
		body["EntityName"] = request.EntityName
	}

	if !dara.IsNil(request.EntityType) {
		body["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.EntityUuid) {
		body["EntityUuid"] = request.EntityUuid
	}

	if !dara.IsNil(request.EntityUuids) {
		body["EntityUuids"] = request.EntityUuids
	}

	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.IsMalwareEntity) {
		body["IsMalwareEntity"] = request.IsMalwareEntity
	}

	if !dara.IsNil(request.MalwareType) {
		body["MalwareType"] = request.MalwareType
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.Tags) {
		body["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEntities"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries entities.
//
// @param request - ListEntitiesRequest
//
// @return ListEntitiesResponse
func (client *Client) ListEntities(request *ListEntitiesRequest) (_result *ListEntitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEntitiesResponse{}
	_body, _err := client.ListEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the logs in a cloud service that is added to the threat analysis feature.
//
// @param request - ListImportedLogsByProdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImportedLogsByProdResponse
func (client *Client) ListImportedLogsByProdWithOptions(request *ListImportedLogsByProdRequest, runtime *dara.RuntimeOptions) (_result *ListImportedLogsByProdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.ProdCode) {
		body["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImportedLogsByProd"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImportedLogsByProdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the logs in a cloud service that is added to the threat analysis feature.
//
// @param request - ListImportedLogsByProdRequest
//
// @return ListImportedLogsByProdResponse
func (client *Client) ListImportedLogsByProd(request *ListImportedLogsByProdRequest) (_result *ListImportedLogsByProdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListImportedLogsByProdResponse{}
	_body, _err := client.ListImportedLogsByProdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the dedicated Simple Log Service project and Logstore for a cloud service based on the patterns of the project and Logstore names.
//
// @param request - ListProjectLogStoresRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectLogStoresResponse
func (client *Client) ListProjectLogStoresWithOptions(request *ListProjectLogStoresRequest, runtime *dara.RuntimeOptions) (_result *ListProjectLogStoresResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourceLogCode) {
		body["SourceLogCode"] = request.SourceLogCode
	}

	if !dara.IsNil(request.SourceProdCode) {
		body["SourceProdCode"] = request.SourceProdCode
	}

	if !dara.IsNil(request.SubUserId) {
		body["SubUserId"] = request.SubUserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectLogStores"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectLogStoresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the dedicated Simple Log Service project and Logstore for a cloud service based on the patterns of the project and Logstore names.
//
// @param request - ListProjectLogStoresRequest
//
// @return ListProjectLogStoresResponse
func (client *Client) ListProjectLogStores(request *ListProjectLogStoresRequest) (_result *ListProjectLogStoresResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectLogStoresResponse{}
	_body, _err := client.ListProjectLogStoresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of Alibaba Cloud accounts that are added to the threat analysis feature for centralized management. These accounts can be used to perform operations supported by the threat analysis feature, such as adding logs and handling events.
//
// @param request - ListRdUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRdUsersResponse
func (client *Client) ListRdUsersWithOptions(request *ListRdUsersRequest, runtime *dara.RuntimeOptions) (_result *ListRdUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRdUsers"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRdUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Alibaba Cloud accounts that are added to the threat analysis feature for centralized management. These accounts can be used to perform operations supported by the threat analysis feature, such as adding logs and handling events.
//
// @param request - ListRdUsersRequest
//
// @return ListRdUsersResponse
func (client *Client) ListRdUsers(request *ListRdUsersRequest) (_result *ListRdUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRdUsersResponse{}
	_body, _err := client.ListRdUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a third-party cloud account that is added to the threat analysis feature.
//
// @param request - ModifyBindAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBindAccountResponse
func (client *Client) ModifyBindAccountWithOptions(request *ModifyBindAccountRequest, runtime *dara.RuntimeOptions) (_result *ModifyBindAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessId) {
		body["AccessId"] = request.AccessId
	}

	if !dara.IsNil(request.AccountId) {
		body["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.BindId) {
		body["BindId"] = request.BindId
	}

	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBindAccount"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBindAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a third-party cloud account that is added to the threat analysis feature.
//
// @param request - ModifyBindAccountRequest
//
// @return ModifyBindAccountResponse
func (client *Client) ModifyBindAccount(request *ModifyBindAccountRequest) (_result *ModifyBindAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBindAccountResponse{}
	_body, _err := client.ModifyBindAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a data source that is added to the threat analysis feature.
//
// @param request - ModifyDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDataSourceResponse
func (client *Client) ModifyDataSourceWithOptions(request *ModifyDataSourceRequest, runtime *dara.RuntimeOptions) (_result *ModifyDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		body["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.DataSourceInstanceId) {
		body["DataSourceInstanceId"] = request.DataSourceInstanceId
	}

	if !dara.IsNil(request.DataSourceInstanceName) {
		body["DataSourceInstanceName"] = request.DataSourceInstanceName
	}

	if !dara.IsNil(request.DataSourceInstanceParams) {
		body["DataSourceInstanceParams"] = request.DataSourceInstanceParams
	}

	if !dara.IsNil(request.DataSourceInstanceRemark) {
		body["DataSourceInstanceRemark"] = request.DataSourceInstanceRemark
	}

	if !dara.IsNil(request.DataSourceType) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDataSource"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a data source that is added to the threat analysis feature.
//
// @param request - ModifyDataSourceRequest
//
// @return ModifyDataSourceResponse
func (client *Client) ModifyDataSource(request *ModifyDataSourceRequest) (_result *ModifyDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDataSourceResponse{}
	_body, _err := client.ModifyDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of the logs that are added to the threat analysis feature for a data source within a cloud account.
//
// @param request - ModifyDataSourceLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDataSourceLogResponse
func (client *Client) ModifyDataSourceLogWithOptions(request *ModifyDataSourceLogRequest, runtime *dara.RuntimeOptions) (_result *ModifyDataSourceLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		body["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.DataSourceInstanceId) {
		body["DataSourceInstanceId"] = request.DataSourceInstanceId
	}

	if !dara.IsNil(request.DataSourceInstanceLogs) {
		body["DataSourceInstanceLogs"] = request.DataSourceInstanceLogs
	}

	if !dara.IsNil(request.DataSourceType) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.LogCode) {
		body["LogCode"] = request.LogCode
	}

	if !dara.IsNil(request.LogInstanceId) {
		body["LogInstanceId"] = request.LogInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDataSourceLog"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDataSourceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of the logs that are added to the threat analysis feature for a data source within a cloud account.
//
// @param request - ModifyDataSourceLogRequest
//
// @return ModifyDataSourceLogResponse
func (client *Client) ModifyDataSourceLog(request *ModifyDataSourceLogRequest) (_result *ModifyDataSourceLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDataSourceLogResponse{}
	_body, _err := client.ModifyDataSourceLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the log delivery feature for a cloud service that is integrated with Simple Log Service.
//
// @param request - OpenDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenDeliveryResponse
func (client *Client) OpenDeliveryWithOptions(request *OpenDeliveryRequest, runtime *dara.RuntimeOptions) (_result *OpenDeliveryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LogCode) {
		body["LogCode"] = request.LogCode
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenDelivery"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the log delivery feature for a cloud service that is integrated with Simple Log Service.
//
// @param request - OpenDeliveryRequest
//
// @return OpenDeliveryResponse
func (client *Client) OpenDelivery(request *OpenDeliveryRequest) (_result *OpenDeliveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenDeliveryResponse{}
	_body, _err := client.OpenDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or updates an automatic response rule.
//
// @param request - PostAutomateResponseConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostAutomateResponseConfigResponse
func (client *Client) PostAutomateResponseConfigWithOptions(request *PostAutomateResponseConfigRequest, runtime *dara.RuntimeOptions) (_result *PostAutomateResponseConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ActionConfig) {
		body["ActionConfig"] = request.ActionConfig
	}

	if !dara.IsNil(request.ActionType) {
		body["ActionType"] = request.ActionType
	}

	if !dara.IsNil(request.AutoResponseType) {
		body["AutoResponseType"] = request.AutoResponseType
	}

	if !dara.IsNil(request.ExecutionCondition) {
		body["ExecutionCondition"] = request.ExecutionCondition
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.RuleName) {
		body["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SubUserId) {
		body["SubUserId"] = request.SubUserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PostAutomateResponseConfig"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PostAutomateResponseConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates an automatic response rule.
//
// @param request - PostAutomateResponseConfigRequest
//
// @return PostAutomateResponseConfigResponse
func (client *Client) PostAutomateResponseConfig(request *PostAutomateResponseConfigRequest) (_result *PostAutomateResponseConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PostAutomateResponseConfigResponse{}
	_body, _err := client.PostAutomateResponseConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or updates a custom rule.
//
// @param request - PostCustomizeRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostCustomizeRuleResponse
func (client *Client) PostCustomizeRuleWithOptions(request *PostCustomizeRuleRequest, runtime *dara.RuntimeOptions) (_result *PostCustomizeRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertType) {
		body["AlertType"] = request.AlertType
	}

	if !dara.IsNil(request.AlertTypeMds) {
		body["AlertTypeMds"] = request.AlertTypeMds
	}

	if !dara.IsNil(request.AttCk) {
		body["AttCk"] = request.AttCk
	}

	if !dara.IsNil(request.EventTransferExt) {
		body["EventTransferExt"] = request.EventTransferExt
	}

	if !dara.IsNil(request.EventTransferSwitch) {
		body["EventTransferSwitch"] = request.EventTransferSwitch
	}

	if !dara.IsNil(request.EventTransferType) {
		body["EventTransferType"] = request.EventTransferType
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.LogSource) {
		body["LogSource"] = request.LogSource
	}

	if !dara.IsNil(request.LogSourceMds) {
		body["LogSourceMds"] = request.LogSourceMds
	}

	if !dara.IsNil(request.LogType) {
		body["LogType"] = request.LogType
	}

	if !dara.IsNil(request.LogTypeMds) {
		body["LogTypeMds"] = request.LogTypeMds
	}

	if !dara.IsNil(request.QueryCycle) {
		body["QueryCycle"] = request.QueryCycle
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.RuleCondition) {
		body["RuleCondition"] = request.RuleCondition
	}

	if !dara.IsNil(request.RuleDesc) {
		body["RuleDesc"] = request.RuleDesc
	}

	if !dara.IsNil(request.RuleGroup) {
		body["RuleGroup"] = request.RuleGroup
	}

	if !dara.IsNil(request.RuleName) {
		body["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleThreshold) {
		body["RuleThreshold"] = request.RuleThreshold
	}

	if !dara.IsNil(request.ThreatLevel) {
		body["ThreatLevel"] = request.ThreatLevel
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PostCustomizeRule"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PostCustomizeRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates a custom rule.
//
// @param request - PostCustomizeRuleRequest
//
// @return PostCustomizeRuleResponse
func (client *Client) PostCustomizeRule(request *PostCustomizeRuleRequest) (_result *PostCustomizeRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PostCustomizeRuleResponse{}
	_body, _err := client.PostCustomizeRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a custom rule for testing.
//
// @param request - PostCustomizeRuleTestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostCustomizeRuleTestResponse
func (client *Client) PostCustomizeRuleTestWithOptions(request *PostCustomizeRuleTestRequest, runtime *dara.RuntimeOptions) (_result *PostCustomizeRuleTestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.SimulatedData) {
		body["SimulatedData"] = request.SimulatedData
	}

	if !dara.IsNil(request.TestType) {
		body["TestType"] = request.TestType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PostCustomizeRuleTest"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PostCustomizeRuleTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a custom rule for testing.
//
// @param request - PostCustomizeRuleTestRequest
//
// @return PostCustomizeRuleTestResponse
func (client *Client) PostCustomizeRuleTest(request *PostCustomizeRuleTestRequest) (_result *PostCustomizeRuleTestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PostCustomizeRuleTestResponse{}
	_body, _err := client.PostCustomizeRuleTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits event handling information.
//
// @param request - PostEventDisposeAndWhiteruleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostEventDisposeAndWhiteruleListResponse
func (client *Client) PostEventDisposeAndWhiteruleListWithOptions(request *PostEventDisposeAndWhiteruleListRequest, runtime *dara.RuntimeOptions) (_result *PostEventDisposeAndWhiteruleListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DisposeStrategyIds) {
		body["DisposeStrategyIds"] = request.DisposeStrategyIds
	}

	if !dara.IsNil(request.EventDispose) {
		body["EventDispose"] = request.EventDispose
	}

	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ReceiverInfo) {
		body["ReceiverInfo"] = request.ReceiverInfo
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		body["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResponseSource) {
		body["ResponseSource"] = request.ResponseSource
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.ThreatLevel) {
		body["ThreatLevel"] = request.ThreatLevel
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PostEventDisposeAndWhiteruleList"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PostEventDisposeAndWhiteruleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits event handling information.
//
// @param request - PostEventDisposeAndWhiteruleListRequest
//
// @return PostEventDisposeAndWhiteruleListResponse
func (client *Client) PostEventDisposeAndWhiteruleList(request *PostEventDisposeAndWhiteruleListRequest) (_result *PostEventDisposeAndWhiteruleListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PostEventDisposeAndWhiteruleListResponse{}
	_body, _err := client.PostEventDisposeAndWhiteruleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits an alert whitelist rule.
//
// @param request - PostEventWhiteruleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostEventWhiteruleListResponse
func (client *Client) PostEventWhiteruleListWithOptions(request *PostEventWhiteruleListRequest, runtime *dara.RuntimeOptions) (_result *PostEventWhiteruleListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.WhiteruleList) {
		body["WhiteruleList"] = request.WhiteruleList
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PostEventWhiteruleList"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PostEventWhiteruleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an alert whitelist rule.
//
// @param request - PostEventWhiteruleListRequest
//
// @return PostEventWhiteruleListResponse
func (client *Client) PostEventWhiteruleList(request *PostEventWhiteruleListRequest) (_result *PostEventWhiteruleListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PostEventWhiteruleListResponse{}
	_body, _err := client.PostEventWhiteruleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Ends the test of a custom rule.
//
// @param request - PostFinishCustomizeRuleTestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostFinishCustomizeRuleTestResponse
func (client *Client) PostFinishCustomizeRuleTestWithOptions(request *PostFinishCustomizeRuleTestRequest, runtime *dara.RuntimeOptions) (_result *PostFinishCustomizeRuleTestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PostFinishCustomizeRuleTest"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PostFinishCustomizeRuleTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Ends the test of a custom rule.
//
// @param request - PostFinishCustomizeRuleTestRequest
//
// @return PostFinishCustomizeRuleTestResponse
func (client *Client) PostFinishCustomizeRuleTest(request *PostFinishCustomizeRuleTestRequest) (_result *PostFinishCustomizeRuleTestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PostFinishCustomizeRuleTestResponse{}
	_body, _err := client.PostFinishCustomizeRuleTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the status of a custom rule.
//
// @param request - PostRuleStatusChangeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostRuleStatusChangeResponse
func (client *Client) PostRuleStatusChangeWithOptions(request *PostRuleStatusChangeRequest, runtime *dara.RuntimeOptions) (_result *PostRuleStatusChangeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Ids) {
		body["Ids"] = request.Ids
	}

	if !dara.IsNil(request.InUse) {
		body["InUse"] = request.InUse
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.RuleType) {
		body["RuleType"] = request.RuleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PostRuleStatusChange"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PostRuleStatusChangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of a custom rule.
//
// @param request - PostRuleStatusChangeRequest
//
// @return PostRuleStatusChangeResponse
func (client *Client) PostRuleStatusChange(request *PostRuleStatusChangeRequest) (_result *PostRuleStatusChangeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PostRuleStatusChangeResponse{}
	_body, _err := client.PostRuleStatusChangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases storage to reduce the storage usage. The release operation is irreversible and may cause data loss. Proceed with caution.
//
// @param request - RestoreCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreCapacityResponse
func (client *Client) RestoreCapacityWithOptions(request *RestoreCapacityRequest, runtime *dara.RuntimeOptions) (_result *RestoreCapacityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestoreCapacity"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestoreCapacityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases storage to reduce the storage usage. The release operation is irreversible and may cause data loss. Proceed with caution.
//
// @param request - RestoreCapacityRequest
//
// @return RestoreCapacityResponse
func (client *Client) RestoreCapacity(request *RestoreCapacityRequest) (_result *RestoreCapacityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestoreCapacityResponse{}
	_body, _err := client.RestoreCapacityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the settings of log storage, such as the storage duration and storage region.
//
// @param request - SetStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetStorageResponse
func (client *Client) SetStorageWithOptions(request *SetStorageRequest, runtime *dara.RuntimeOptions) (_result *SetStorageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.Ttl) {
		body["Ttl"] = request.Ttl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetStorage"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the settings of log storage, such as the storage duration and storage region.
//
// @param request - SetStorageRequest
//
// @return SetStorageResponse
func (client *Client) SetStorage(request *SetStorageRequest) (_result *SetStorageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetStorageResponse{}
	_body, _err := client.SetStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits log collection tasks at a time.
//
// @param request - SubmitImportLogTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitImportLogTasksResponse
func (client *Client) SubmitImportLogTasksWithOptions(request *SubmitImportLogTasksRequest, runtime *dara.RuntimeOptions) (_result *SubmitImportLogTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accounts) {
		body["Accounts"] = request.Accounts
	}

	if !dara.IsNil(request.AutoImported) {
		body["AutoImported"] = request.AutoImported
	}

	if !dara.IsNil(request.CloudCode) {
		body["CloudCode"] = request.CloudCode
	}

	if !dara.IsNil(request.LogCodes) {
		body["LogCodes"] = request.LogCodes
	}

	if !dara.IsNil(request.ProdCode) {
		body["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitImportLogTasks"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitImportLogTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits log collection tasks at a time.
//
// @param request - SubmitImportLogTasksRequest
//
// @return SubmitImportLogTasksResponse
func (client *Client) SubmitImportLogTasks(request *SubmitImportLogTasksRequest) (_result *SubmitImportLogTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitImportLogTasksResponse{}
	_body, _err := client.SubmitImportLogTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the status of an automatic response rule.
//
// @param request - UpdateAutomateResponseConfigStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAutomateResponseConfigStatusResponse
func (client *Client) UpdateAutomateResponseConfigStatusWithOptions(request *UpdateAutomateResponseConfigStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateAutomateResponseConfigStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Ids) {
		body["Ids"] = request.Ids
	}

	if !dara.IsNil(request.InUse) {
		body["InUse"] = request.InUse
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAutomateResponseConfigStatus"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAutomateResponseConfigStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of an automatic response rule.
//
// @param request - UpdateAutomateResponseConfigStatusRequest
//
// @return UpdateAutomateResponseConfigStatusResponse
func (client *Client) UpdateAutomateResponseConfigStatus(request *UpdateAutomateResponseConfigStatusRequest) (_result *UpdateAutomateResponseConfigStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAutomateResponseConfigStatusResponse{}
	_body, _err := client.UpdateAutomateResponseConfigStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or updates an alert whitelist rule.
//
// @param request - UpdateWhiteRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWhiteRuleListResponse
func (client *Client) UpdateWhiteRuleListWithOptions(request *UpdateWhiteRuleListRequest, runtime *dara.RuntimeOptions) (_result *UpdateWhiteRuleListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Expression) {
		body["Expression"] = request.Expression
	}

	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.WhiteRuleId) {
		body["WhiteRuleId"] = request.WhiteRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWhiteRuleList"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWhiteRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates an alert whitelist rule.
//
// @param request - UpdateWhiteRuleListRequest
//
// @return UpdateWhiteRuleListResponse
func (client *Client) UpdateWhiteRuleList(request *UpdateWhiteRuleListRequest) (_result *UpdateWhiteRuleListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWhiteRuleListResponse{}
	_body, _err := client.UpdateWhiteRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
