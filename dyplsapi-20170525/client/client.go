// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
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
	client.ProductId = dara.String("Dyplsapi")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = dara.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("dyplsapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds a tracking number for a private number in the AXN binding.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddAxnTrackNoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAxnTrackNoResponse
func (client *Client) AddAxnTrackNoWithOptions(request *AddAxnTrackNoRequest, runtime *dara.RuntimeOptions) (_result *AddAxnTrackNoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNoX) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SubsId) {
		query["SubsId"] = request.SubsId
	}

	if !dara.IsNil(request.TrackNo) {
		query["trackNo"] = request.TrackNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAxnTrackNo"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAxnTrackNoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a tracking number for a private number in the AXN binding.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddAxnTrackNoRequest
//
// @return AddAxnTrackNoResponse
func (client *Client) AddAxnTrackNo(request *AddAxnTrackNoRequest) (_result *AddAxnTrackNoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddAxnTrackNoResponse{}
	_body, _err := client.AddAxnTrackNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a phone number to the blacklist of a China number pool. Only AXN privacy number, AXN extension number, and 95AXN privacy number pool types support blacklist settings.
//
// Description:
//
// - This operation is not a batch operation. You can add only one pair of **point-to-point blacklist*	- numbers or one **number pool blacklist*	- number at a time.
//
// - To delete a blacklist entry, call the [DeleteSecretBlacklist](https://www.alibabacloud.com/help/en/pnp/developer-reference/api-dyplsapi-2017-05-25-deletesecretblacklist) operation.
//
// ### Quantity limit
//
// The maximum number of blacklist entries per number pool is 2,000,000. If this limit is exceeded, API calls to add blacklist numbers will fail, which may affect your business. If you have special requirements, contact your Alibaba Cloud account manager.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 100 calls per second. If this limit is exceeded, API calls will be throttled, which may affect your business. Call this operation appropriately.
//
// @param request - AddSecretBlacklistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSecretBlacklistResponse
func (client *Client) AddSecretBlacklistWithOptions(request *AddSecretBlacklistRequest, runtime *dara.RuntimeOptions) (_result *AddSecretBlacklistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BlackNo) {
		query["BlackNo"] = request.BlackNo
	}

	if !dara.IsNil(request.BlackType) {
		query["BlackType"] = request.BlackType
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.WayControl) {
		query["WayControl"] = request.WayControl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSecretBlacklist"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSecretBlacklistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a phone number to the blacklist of a China number pool. Only AXN privacy number, AXN extension number, and 95AXN privacy number pool types support blacklist settings.
//
// Description:
//
// - This operation is not a batch operation. You can add only one pair of **point-to-point blacklist*	- numbers or one **number pool blacklist*	- number at a time.
//
// - To delete a blacklist entry, call the [DeleteSecretBlacklist](https://www.alibabacloud.com/help/en/pnp/developer-reference/api-dyplsapi-2017-05-25-deletesecretblacklist) operation.
//
// ### Quantity limit
//
// The maximum number of blacklist entries per number pool is 2,000,000. If this limit is exceeded, API calls to add blacklist numbers will fail, which may affect your business. If you have special requirements, contact your Alibaba Cloud account manager.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 100 calls per second. If this limit is exceeded, API calls will be throttled, which may affect your business. Call this operation appropriately.
//
// @param request - AddSecretBlacklistRequest
//
// @return AddSecretBlacklistResponse
func (client *Client) AddSecretBlacklist(request *AddSecretBlacklistRequest) (_result *AddSecretBlacklistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddSecretBlacklistResponse{}
	_body, _err := client.AddSecretBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Establishes a call binding relationship (B, X, A) for work number X of employee B on the work number platform to allow employee B to call customer A through work number X.
//
// @param request - BindAXBCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAXBCallResponse
func (client *Client) BindAXBCallWithOptions(request *BindAXBCallRequest, runtime *dara.RuntimeOptions) (_result *BindAXBCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthId) {
		query["AuthId"] = request.AuthId
	}

	if !dara.IsNil(request.CallerParentId) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !dara.IsNil(request.CustomerPoolKey) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReqId) {
		query["ReqId"] = request.ReqId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TelA) {
		query["TelA"] = request.TelA
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAXBCall"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAXBCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Establishes a call binding relationship (B, X, A) for work number X of employee B on the work number platform to allow employee B to call customer A through work number X.
//
// @param request - BindAXBCallRequest
//
// @return BindAXBCallResponse
func (client *Client) BindAXBCall(request *BindAXBCallRequest) (_result *BindAXBCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindAXBCallResponse{}
	_body, _err := client.BindAXBCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an AXB binding.
//
// Description:
//
// Before you add an AXB binding, we recommend that you specify role A and role B in the AXB device certificate (ProductKey, DeviceName, and DeviceSecret) in your business scenario. For example, in a taxi-hailing scenario, role A is the passenger and role B is the driver.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - BindAxbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxbResponse
func (client *Client) BindAxbWithOptions(request *BindAxbRequest, runtime *dara.RuntimeOptions) (_result *BindAxbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ASRModelId) {
		query["ASRModelId"] = request.ASRModelId
	}

	if !dara.IsNil(request.ASRStatus) {
		query["ASRStatus"] = request.ASRStatus
	}

	if !dara.IsNil(request.CallDisplayType) {
		query["CallDisplayType"] = request.CallDisplayType
	}

	if !dara.IsNil(request.CallRestrict) {
		query["CallRestrict"] = request.CallRestrict
	}

	if !dara.IsNil(request.CallTimeout) {
		query["CallTimeout"] = request.CallTimeout
	}

	if !dara.IsNil(request.DtmfConfig) {
		query["DtmfConfig"] = request.DtmfConfig
	}

	if !dara.IsNil(request.ExpectCity) {
		query["ExpectCity"] = request.ExpectCity
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.IsRecordingEnabled) {
		query["IsRecordingEnabled"] = request.IsRecordingEnabled
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["OutOrderId"] = request.OutOrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNoA) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !dara.IsNil(request.PhoneNoB) {
		query["PhoneNoB"] = request.PhoneNoB
	}

	if !dara.IsNil(request.PhoneNoX) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RingConfig) {
		query["RingConfig"] = request.RingConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAxb"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAxbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an AXB binding.
//
// Description:
//
// Before you add an AXB binding, we recommend that you specify role A and role B in the AXB device certificate (ProductKey, DeviceName, and DeviceSecret) in your business scenario. For example, in a taxi-hailing scenario, role A is the passenger and role B is the driver.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - BindAxbRequest
//
// @return BindAxbResponse
func (client *Client) BindAxb(request *BindAxbRequest) (_result *BindAxbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindAxbResponse{}
	_body, _err := client.BindAxbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds phone numbers for a 700AXB private number. A 700AXB private number restricts the calling relationship to a one-to-one relationship. User A can contact User B through the intermediate number X that starts with 700, and User B can also contact User A through the same intermediate number X.
//
// Description:
//
// Before invoking this operation, make sure that you have purchased a private number in the console.
//
// You can use this operation to attach numbers for a 700AXB private number, or attach A and B numbers for a 700AXB private number on the Number Management page in the console.
//
// Before attaching numbers for a 700AXB private number, determine the A role and B role in your business scenario. For example, in a ride-hailing scenario, A is the passenger and B is the driver.
//
// The 700AXB private number is mainly used for privacy protection in one-to-one scenarios. You attach number A and number B through the Phone Number Protection service. When A calls X, the call is forwarded to B. When B calls X, the call is forwarded to A. A single X number can be reused by up to 200 binding pairs, such as A1XB1, A2XB2, A3XB3, and so on. As long as the phone numbers of A and B users are unique, a single X number can be bound to 200 groups of phone numbers simultaneously, and calls can also be made simultaneously.
//
// # QPS limit
//
// The single-user QPS limit for this operation is 5,000 calls per second. Exceeding this limit will cause throttling, which may affect your business. Invoke this operation appropriately.
//
// @param request - BindAxb700Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxb700Response
func (client *Client) BindAxb700WithOptions(request *BindAxb700Request, runtime *dara.RuntimeOptions) (_result *BindAxb700Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AsrModelId) {
		query["AsrModelId"] = request.AsrModelId
	}

	if !dara.IsNil(request.Audio) {
		query["Audio"] = request.Audio
	}

	if !dara.IsNil(request.CallRestrict) {
		query["CallRestrict"] = request.CallRestrict
	}

	if !dara.IsNil(request.CallTimeout) {
		query["CallTimeout"] = request.CallTimeout
	}

	if !dara.IsNil(request.DtmfConfig) {
		query["DtmfConfig"] = request.DtmfConfig
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.IndustrialId) {
		query["IndustrialId"] = request.IndustrialId
	}

	if !dara.IsNil(request.NeedAsr) {
		query["NeedAsr"] = request.NeedAsr
	}

	if !dara.IsNil(request.NeedRecord) {
		query["NeedRecord"] = request.NeedRecord
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["OutOrderId"] = request.OutOrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.RecType) {
		query["RecType"] = request.RecType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TelA) {
		query["TelA"] = request.TelA
	}

	if !dara.IsNil(request.TelB) {
		query["TelB"] = request.TelB
	}

	if !dara.IsNil(request.TelX) {
		query["TelX"] = request.TelX
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAxb700"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAxb700Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds phone numbers for a 700AXB private number. A 700AXB private number restricts the calling relationship to a one-to-one relationship. User A can contact User B through the intermediate number X that starts with 700, and User B can also contact User A through the same intermediate number X.
//
// Description:
//
// Before invoking this operation, make sure that you have purchased a private number in the console.
//
// You can use this operation to attach numbers for a 700AXB private number, or attach A and B numbers for a 700AXB private number on the Number Management page in the console.
//
// Before attaching numbers for a 700AXB private number, determine the A role and B role in your business scenario. For example, in a ride-hailing scenario, A is the passenger and B is the driver.
//
// The 700AXB private number is mainly used for privacy protection in one-to-one scenarios. You attach number A and number B through the Phone Number Protection service. When A calls X, the call is forwarded to B. When B calls X, the call is forwarded to A. A single X number can be reused by up to 200 binding pairs, such as A1XB1, A2XB2, A3XB3, and so on. As long as the phone numbers of A and B users are unique, a single X number can be bound to 200 groups of phone numbers simultaneously, and calls can also be made simultaneously.
//
// # QPS limit
//
// The single-user QPS limit for this operation is 5,000 calls per second. Exceeding this limit will cause throttling, which may affect your business. Invoke this operation appropriately.
//
// @param request - BindAxb700Request
//
// @return BindAxb700Response
func (client *Client) BindAxb700(request *BindAxb700Request) (_result *BindAxb700Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindAxb700Response{}
	_body, _err := client.BindAxb700WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches an AXB privacy number to a fixed line to restrict the calling relationship to a one-to-one relationship, allowing User A to contact User B through the fixed-line intermediate number X, and User B to contact User A through the same intermediate number X. Fixed-line privacy numbers are currently available by invitation only. Contact your technical support to activate the service and obtain integration support.
//
// Description:
//
// - Before you invoke this operation, make sure that you have purchased a privacy number in the console. You can use this operation to bind numbers to a fixed-line AXB intermediate number. Currently, you cannot bind Number A and Number B to a fixed-line AXB intermediate number on the Number Management page in the console.
//
// - Before binding numbers, we recommend that you clarify the roles of User A and User B in your business scenario. For example, in common scenarios such as ride-hailing, User A is the passenger and User B is the driver.
//
// - Fixed-line AXB intermediate numbers are mainly used for privacy protection in one-to-one scenarios. Number A and Number B are bound through the Phone Number Protection service. When User A calls Number X, the call is forwarded to User B. When User B calls Number X, the call is forwarded to User A. A single fixed-line Number X can be reused for N binding relationships, such as A1-X-B1, A2-X-B2, and A3-X-B3. As long as the mobile phone numbers of Users A and B are not duplicated, a single fixed-line Number X can be bound to N pairs of mobile phone numbers simultaneously, and calls can be made concurrently.
//
// - The number of times a number can be reused varies based on the number you purchased. You can confirm the specific reuse limit with your technical support when purchasing the number.
//
// @param tmpReq - BindAxbFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxbFixedLineResponse
func (client *Client) BindAxbFixedLineWithOptions(tmpReq *BindAxbFixedLineRequest, runtime *dara.RuntimeOptions) (_result *BindAxbFixedLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BindAxbFixedLineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("Extra"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Anucode) {
		query["Anucode"] = request.Anucode
	}

	if !dara.IsNil(request.Anucodecalled) {
		query["Anucodecalled"] = request.Anucodecalled
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Areacode) {
		query["Areacode"] = request.Areacode
	}

	if !dara.IsNil(request.BindType) {
		query["BindType"] = request.BindType
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.ExtraShrink) {
		query["Extra"] = request.ExtraShrink
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Subts) {
		query["Subts"] = request.Subts
	}

	if !dara.IsNil(request.TAnucodeConnect) {
		query["TAnucodeConnect"] = request.TAnucodeConnect
	}

	if !dara.IsNil(request.TelA) {
		query["TelA"] = request.TelA
	}

	if !dara.IsNil(request.TelB) {
		query["TelB"] = request.TelB
	}

	if !dara.IsNil(request.TelX) {
		query["TelX"] = request.TelX
	}

	if !dara.IsNil(request.Ts) {
		query["Ts"] = request.Ts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAxbFixedLine"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAxbFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches an AXB privacy number to a fixed line to restrict the calling relationship to a one-to-one relationship, allowing User A to contact User B through the fixed-line intermediate number X, and User B to contact User A through the same intermediate number X. Fixed-line privacy numbers are currently available by invitation only. Contact your technical support to activate the service and obtain integration support.
//
// Description:
//
// - Before you invoke this operation, make sure that you have purchased a privacy number in the console. You can use this operation to bind numbers to a fixed-line AXB intermediate number. Currently, you cannot bind Number A and Number B to a fixed-line AXB intermediate number on the Number Management page in the console.
//
// - Before binding numbers, we recommend that you clarify the roles of User A and User B in your business scenario. For example, in common scenarios such as ride-hailing, User A is the passenger and User B is the driver.
//
// - Fixed-line AXB intermediate numbers are mainly used for privacy protection in one-to-one scenarios. Number A and Number B are bound through the Phone Number Protection service. When User A calls Number X, the call is forwarded to User B. When User B calls Number X, the call is forwarded to User A. A single fixed-line Number X can be reused for N binding relationships, such as A1-X-B1, A2-X-B2, and A3-X-B3. As long as the mobile phone numbers of Users A and B are not duplicated, a single fixed-line Number X can be bound to N pairs of mobile phone numbers simultaneously, and calls can be made concurrently.
//
// - The number of times a number can be reused varies based on the number you purchased. You can confirm the specific reuse limit with your technical support when purchasing the number.
//
// @param request - BindAxbFixedLineRequest
//
// @return BindAxbFixedLineResponse
func (client *Client) BindAxbFixedLine(request *BindAxbFixedLineRequest) (_result *BindAxbFixedLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindAxbFixedLineResponse{}
	_body, _err := client.BindAxbFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an AXG binding.
//
// Description:
//
// An AXG protection solution can be configured to meet the requirements for grading users, limiting the scope of calls, and restricting order snatching. The letter G represents a phone number group to which you can add phone numbers as needed.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - BindAxgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxgResponse
func (client *Client) BindAxgWithOptions(request *BindAxgRequest, runtime *dara.RuntimeOptions) (_result *BindAxgResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ASRModelId) {
		query["ASRModelId"] = request.ASRModelId
	}

	if !dara.IsNil(request.ASRStatus) {
		query["ASRStatus"] = request.ASRStatus
	}

	if !dara.IsNil(request.CallDisplayType) {
		query["CallDisplayType"] = request.CallDisplayType
	}

	if !dara.IsNil(request.CallRestrict) {
		query["CallRestrict"] = request.CallRestrict
	}

	if !dara.IsNil(request.ExpectCity) {
		query["ExpectCity"] = request.ExpectCity
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IsRecordingEnabled) {
		query["IsRecordingEnabled"] = request.IsRecordingEnabled
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["OutOrderId"] = request.OutOrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNoA) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !dara.IsNil(request.PhoneNoB) {
		query["PhoneNoB"] = request.PhoneNoB
	}

	if !dara.IsNil(request.PhoneNoX) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RingConfig) {
		query["RingConfig"] = request.RingConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAxg"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAxgResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an AXG binding.
//
// Description:
//
// An AXG protection solution can be configured to meet the requirements for grading users, limiting the scope of calls, and restricting order snatching. The letter G represents a phone number group to which you can add phone numbers as needed.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - BindAxgRequest
//
// @return BindAxgResponse
func (client *Client) BindAxg(request *BindAxgRequest) (_result *BindAxgResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindAxgResponse{}
	_body, _err := client.BindAxgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a binding relationship for an AXN privacy number. An AXN privacy number assigns an exclusive privacy number to User A, allowing any Number N to reach User A by dialing Number X.
//
// Description:
//
// - Before calling this operation, make sure that you have [purchased a privacy number](https://dyplsnext.console.aliyun.com/numberbuy) in the console.
//
// - You can use this operation to bind numbers to an AXN privacy number, or bind Number A and Number B to an AXN privacy number on the [Number Management](https://dyplsnext.console.aliyun.com/number/list?spm=a2c4g.11186623.0.0.64e649bbTzmb6j&resType=1&billId=FC100000165392498) page in the console.
//
// - Before binding numbers to an AXN privacy number, clarify the role of User A in your business scenario. For example, in a used car information publishing scenario, User A represents the contact information of the landlord or agent.
//
// ### QPS limits
//
// The QPS limit per user is 5,000 calls per second. If this limit is exceeded, API calls are throttled, which may affect your services. Call the API at a reasonable frequency.
//
// @param request - BindAxnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxnResponse
func (client *Client) BindAxnWithOptions(request *BindAxnRequest, runtime *dara.RuntimeOptions) (_result *BindAxnResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ASRModelId) {
		query["ASRModelId"] = request.ASRModelId
	}

	if !dara.IsNil(request.ASRStatus) {
		query["ASRStatus"] = request.ASRStatus
	}

	if !dara.IsNil(request.CallDisplayType) {
		query["CallDisplayType"] = request.CallDisplayType
	}

	if !dara.IsNil(request.CallRestrict) {
		query["CallRestrict"] = request.CallRestrict
	}

	if !dara.IsNil(request.CallTimeout) {
		query["CallTimeout"] = request.CallTimeout
	}

	if !dara.IsNil(request.ExpectCity) {
		query["ExpectCity"] = request.ExpectCity
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.IsRecordingEnabled) {
		query["IsRecordingEnabled"] = request.IsRecordingEnabled
	}

	if !dara.IsNil(request.NoType) {
		query["NoType"] = request.NoType
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["OutOrderId"] = request.OutOrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNoA) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !dara.IsNil(request.PhoneNoB) {
		query["PhoneNoB"] = request.PhoneNoB
	}

	if !dara.IsNil(request.PhoneNoX) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RingConfig) {
		query["RingConfig"] = request.RingConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAxn"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAxnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a binding relationship for an AXN privacy number. An AXN privacy number assigns an exclusive privacy number to User A, allowing any Number N to reach User A by dialing Number X.
//
// Description:
//
// - Before calling this operation, make sure that you have [purchased a privacy number](https://dyplsnext.console.aliyun.com/numberbuy) in the console.
//
// - You can use this operation to bind numbers to an AXN privacy number, or bind Number A and Number B to an AXN privacy number on the [Number Management](https://dyplsnext.console.aliyun.com/number/list?spm=a2c4g.11186623.0.0.64e649bbTzmb6j&resType=1&billId=FC100000165392498) page in the console.
//
// - Before binding numbers to an AXN privacy number, clarify the role of User A in your business scenario. For example, in a used car information publishing scenario, User A represents the contact information of the landlord or agent.
//
// ### QPS limits
//
// The QPS limit per user is 5,000 calls per second. If this limit is exceeded, API calls are throttled, which may affect your services. Call the API at a reasonable frequency.
//
// @param request - BindAxnRequest
//
// @return BindAxnResponse
func (client *Client) BindAxn(request *BindAxnRequest) (_result *BindAxnResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindAxnResponse{}
	_body, _err := client.BindAxnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a binding relationship for an AXN extension number. The AXN extension number feature is similar to the AXN privacy number feature but introduces extension numbers to reduce the usage of X numbers. When making a call, the caller dials the X number first and then enters the corresponding extension number after hearing the prompt tone to connect to number A.
//
// Description:
//
// - Before calling this operation, make sure that you have [purchased privacy numbers](https://dyplsnext.console.aliyun.com/numberbuy) in the console.
//
// - You can use this operation to bind numbers to an AXN extension number. You can also bind numbers A and B to an AXN extension number on the [Number Management](https://dyplsnext.console.aliyun.com/number/list?spm=a2c4g.11186623.0.0.64e649bbTzmb6j&resType=1&billId=FC100000165392498) page in the console.
//
// - Before binding numbers for an AXN extension number, we recommend that you clarify the roles of A and N in your business scenario. For example, if number A is a customer, number X is the privacy number assigned to the customer. Calls from any other number to number X plus the extension number are forwarded to number A. When A calls X, the call is forwarded to the default number B specified during binding.
//
// ### QPS limits
//
// The queries per second (QPS) limit for a single user is 5,000 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - BindAxnExtensionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxnExtensionResponse
func (client *Client) BindAxnExtensionWithOptions(request *BindAxnExtensionRequest, runtime *dara.RuntimeOptions) (_result *BindAxnExtensionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ASRModelId) {
		query["ASRModelId"] = request.ASRModelId
	}

	if !dara.IsNil(request.ASRStatus) {
		query["ASRStatus"] = request.ASRStatus
	}

	if !dara.IsNil(request.CallDisplayType) {
		query["CallDisplayType"] = request.CallDisplayType
	}

	if !dara.IsNil(request.CallRestrict) {
		query["CallRestrict"] = request.CallRestrict
	}

	if !dara.IsNil(request.ExpectCity) {
		query["ExpectCity"] = request.ExpectCity
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.Extension) {
		query["Extension"] = request.Extension
	}

	if !dara.IsNil(request.IsRecordingEnabled) {
		query["IsRecordingEnabled"] = request.IsRecordingEnabled
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["OutOrderId"] = request.OutOrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNoA) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !dara.IsNil(request.PhoneNoB) {
		query["PhoneNoB"] = request.PhoneNoB
	}

	if !dara.IsNil(request.PhoneNoX) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RingConfig) {
		query["RingConfig"] = request.RingConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAxnExtension"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAxnExtensionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a binding relationship for an AXN extension number. The AXN extension number feature is similar to the AXN privacy number feature but introduces extension numbers to reduce the usage of X numbers. When making a call, the caller dials the X number first and then enters the corresponding extension number after hearing the prompt tone to connect to number A.
//
// Description:
//
// - Before calling this operation, make sure that you have [purchased privacy numbers](https://dyplsnext.console.aliyun.com/numberbuy) in the console.
//
// - You can use this operation to bind numbers to an AXN extension number. You can also bind numbers A and B to an AXN extension number on the [Number Management](https://dyplsnext.console.aliyun.com/number/list?spm=a2c4g.11186623.0.0.64e649bbTzmb6j&resType=1&billId=FC100000165392498) page in the console.
//
// - Before binding numbers for an AXN extension number, we recommend that you clarify the roles of A and N in your business scenario. For example, if number A is a customer, number X is the privacy number assigned to the customer. Calls from any other number to number X plus the extension number are forwarded to number A. When A calls X, the call is forwarded to the default number B specified during binding.
//
// ### QPS limits
//
// The queries per second (QPS) limit for a single user is 5,000 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - BindAxnExtensionRequest
//
// @return BindAxnExtensionResponse
func (client *Client) BindAxnExtension(request *BindAxnExtensionRequest) (_result *BindAxnExtensionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindAxnExtensionResponse{}
	_body, _err := client.BindAxnExtensionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The fixed-line AXN extension number introduces extension numbers on top of the fixed-line AXN privacy number, thereby reducing the usage of X numbers. When making a call, first dial the X number, and then enter the corresponding extension number after hearing the prompt tone to be connected to A. The fixed-line AXN extension number service is currently invitation-only. Contact your technical support to activate the service and obtain integration support.
//
// Description:
//
// - Before calling this API, make sure that you have [purchased privacy numbers](https://dyplsnext.console.aliyun.com/numberbuy) in the console.
//
// - You can use this API to bind numbers for the fixed-line AXN extension number. Binding numbers for the fixed-line AXN extension number through the console is not supported.
//
// - Before binding numbers, we recommend that you first clarify the A role and N role in your business scenario. For example, if the A number represents the customer, the X number is the privacy number assigned to the customer. Any other number calling the X number plus the extension number will be forwarded to the A number. When A calls X, the call will be forwarded to the default B number specified during binding.
//
// @param tmpReq - BindAxnExtensionFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxnExtensionFixedLineResponse
func (client *Client) BindAxnExtensionFixedLineWithOptions(tmpReq *BindAxnExtensionFixedLineRequest, runtime *dara.RuntimeOptions) (_result *BindAxnExtensionFixedLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BindAxnExtensionFixedLineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Extraaxx) {
		request.ExtraaxxShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extraaxx, dara.String("Extraaxx"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Anucode) {
		query["Anucode"] = request.Anucode
	}

	if !dara.IsNil(request.Anucodecalled) {
		query["Anucodecalled"] = request.Anucodecalled
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Areacode) {
		query["Areacode"] = request.Areacode
	}

	if !dara.IsNil(request.BindType) {
		query["BindType"] = request.BindType
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.ExtraaxxShrink) {
		query["Extraaxx"] = request.ExtraaxxShrink
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Subts) {
		query["Subts"] = request.Subts
	}

	if !dara.IsNil(request.TAnucodeConnect) {
		query["TAnucodeConnect"] = request.TAnucodeConnect
	}

	if !dara.IsNil(request.TelA) {
		query["TelA"] = request.TelA
	}

	if !dara.IsNil(request.TelX) {
		query["TelX"] = request.TelX
	}

	if !dara.IsNil(request.TelXext) {
		query["TelXext"] = request.TelXext
	}

	if !dara.IsNil(request.Ts) {
		query["Ts"] = request.Ts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAxnExtensionFixedLine"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAxnExtensionFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The fixed-line AXN extension number introduces extension numbers on top of the fixed-line AXN privacy number, thereby reducing the usage of X numbers. When making a call, first dial the X number, and then enter the corresponding extension number after hearing the prompt tone to be connected to A. The fixed-line AXN extension number service is currently invitation-only. Contact your technical support to activate the service and obtain integration support.
//
// Description:
//
// - Before calling this API, make sure that you have [purchased privacy numbers](https://dyplsnext.console.aliyun.com/numberbuy) in the console.
//
// - You can use this API to bind numbers for the fixed-line AXN extension number. Binding numbers for the fixed-line AXN extension number through the console is not supported.
//
// - Before binding numbers, we recommend that you first clarify the A role and N role in your business scenario. For example, if the A number represents the customer, the X number is the privacy number assigned to the customer. Any other number calling the X number plus the extension number will be forwarded to the A number. When A calls X, the call will be forwarded to the default B number specified during binding.
//
// @param request - BindAxnExtensionFixedLineRequest
//
// @return BindAxnExtensionFixedLineResponse
func (client *Client) BindAxnExtensionFixedLine(request *BindAxnExtensionFixedLineRequest) (_result *BindAxnExtensionFixedLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindAxnExtensionFixedLineResponse{}
	_body, _err := client.BindAxnExtensionFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a fixed-line AXN privacy number to assign a dedicated privacy number to User A, which allows any number N to be forwarded to User A by dialing number X. Fixed-line privacy numbers are currently available by invitation only. Contact your technical support to activate the service and obtain integration assistance.
//
// Description:
//
// - Before you call this operation, make sure that you have purchased a privacy number in the console. You can use this operation to bind a number to a fixed-line AXN privacy number. The console does not support binding for fixed-line AXN privacy numbers.
//
// - Before binding a number, we recommend that you clarify the role of User A in your business scenario. For example, in a used car listing scenario, User A represents the contact information of the property owner or broker.
//
// @param tmpReq - BindAxnFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxnFixedLineResponse
func (client *Client) BindAxnFixedLineWithOptions(tmpReq *BindAxnFixedLineRequest, runtime *dara.RuntimeOptions) (_result *BindAxnFixedLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BindAxnFixedLineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("Extra"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Anucode) {
		query["Anucode"] = request.Anucode
	}

	if !dara.IsNil(request.Anucodecalled) {
		query["Anucodecalled"] = request.Anucodecalled
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Areacode) {
		query["Areacode"] = request.Areacode
	}

	if !dara.IsNil(request.BindType) {
		query["BindType"] = request.BindType
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.ExtraShrink) {
		query["Extra"] = request.ExtraShrink
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Subts) {
		query["Subts"] = request.Subts
	}

	if !dara.IsNil(request.TAnucodeConnect) {
		query["TAnucodeConnect"] = request.TAnucodeConnect
	}

	if !dara.IsNil(request.TelA) {
		query["TelA"] = request.TelA
	}

	if !dara.IsNil(request.TelB) {
		query["TelB"] = request.TelB
	}

	if !dara.IsNil(request.TelX) {
		query["TelX"] = request.TelX
	}

	if !dara.IsNil(request.Ts) {
		query["Ts"] = request.Ts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAxnFixedLine"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAxnFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a fixed-line AXN privacy number to assign a dedicated privacy number to User A, which allows any number N to be forwarded to User A by dialing number X. Fixed-line privacy numbers are currently available by invitation only. Contact your technical support to activate the service and obtain integration assistance.
//
// Description:
//
// - Before you call this operation, make sure that you have purchased a privacy number in the console. You can use this operation to bind a number to a fixed-line AXN privacy number. The console does not support binding for fixed-line AXN privacy numbers.
//
// - Before binding a number, we recommend that you clarify the role of User A in your business scenario. For example, in a used car listing scenario, User A represents the contact information of the property owner or broker.
//
// @param request - BindAxnFixedLineRequest
//
// @return BindAxnFixedLineResponse
func (client *Client) BindAxnFixedLine(request *BindAxnFixedLineRequest) (_result *BindAxnFixedLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindAxnFixedLineResponse{}
	_body, _err := client.BindAxnFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds AXG bindings in batches. The AXG model is a Phone Number Protection Solution for scenarios requiring User Tiering, Call Scope Restriction, or Order Grabbing Restriction. In this model, G represents a Number Group, to which you can add phone numbers.
//
// Description:
//
// - Before you bind AXG Numbers, you must purchase X Numbers, create a G Number Group, and add Numbers to the group. To create a G Number Group, see [Set up G Number Group](https://help.aliyun.com/zh/pnp/user-guide/manage-phone-number-pools#section-9bb-xrg-c78) or call the [CreateAxgGroup](https://help.aliyun.com/zh/pnp/developer-reference/api-dyplsapi-2017-05-25-createaxggroup) API.
//
// - Before binding AXG Numbers, define Role A and Role G for your business scenario. For example, on a recruitment platform, Role A can represent job applicants and Role G can represent the hiring team.
//
// - When you call this API for Batch Binding of AXG Numbers, a single failed request does not affect others in the batch. The API returns a status code for each binding request.
//
// ### QPS limit
//
// The QPS limit for this API is 5,000 per user. Exceeding this limit results in throttling. Plan your API calls accordingly to avoid impacting your business.
//
// @param tmpReq - BindBatchAxgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindBatchAxgResponse
func (client *Client) BindBatchAxgWithOptions(tmpReq *BindBatchAxgRequest, runtime *dara.RuntimeOptions) (_result *BindBatchAxgResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BindBatchAxgShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AxgBindList) {
		request.AxgBindListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AxgBindList, dara.String("AxgBindList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AxgBindListShrink) {
		query["AxgBindList"] = request.AxgBindListShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindBatchAxg"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindBatchAxgResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds AXG bindings in batches. The AXG model is a Phone Number Protection Solution for scenarios requiring User Tiering, Call Scope Restriction, or Order Grabbing Restriction. In this model, G represents a Number Group, to which you can add phone numbers.
//
// Description:
//
// - Before you bind AXG Numbers, you must purchase X Numbers, create a G Number Group, and add Numbers to the group. To create a G Number Group, see [Set up G Number Group](https://help.aliyun.com/zh/pnp/user-guide/manage-phone-number-pools#section-9bb-xrg-c78) or call the [CreateAxgGroup](https://help.aliyun.com/zh/pnp/developer-reference/api-dyplsapi-2017-05-25-createaxggroup) API.
//
// - Before binding AXG Numbers, define Role A and Role G for your business scenario. For example, on a recruitment platform, Role A can represent job applicants and Role G can represent the hiring team.
//
// - When you call this API for Batch Binding of AXG Numbers, a single failed request does not affect others in the batch. The API returns a status code for each binding request.
//
// ### QPS limit
//
// The QPS limit for this API is 5,000 per user. Exceeding this limit results in throttling. Plan your API calls accordingly to avoid impacting your business.
//
// @param request - BindBatchAxgRequest
//
// @return BindBatchAxgResponse
func (client *Client) BindBatchAxg(request *BindBatchAxgRequest) (_result *BindBatchAxgResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindBatchAxgResponse{}
	_body, _err := client.BindBatchAxgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a binding relationship for a 700GXB private number. The 700GXB product is a privacy protection solutions designed for scenarios that require user-level classification, call scope restrictions, and order-grabbing restrictions. In this product, G represents a number group, and you can add numbers to the group as needed.
//
// Description:
//
// Before calling this operation, make sure that you have purchased a private number in the console.
//
// You can use this operation to bind numbers to a 700GXB private number, or bind numbers on the Number Management page in the console.
//
// Before binding numbers for a 700GXB private number, clarify the G role and B role in your business scenario. For example, on a recruitment platform, B is the job applicant and G is the recruitment team.
//
// # QPS limit
//
// The single-user QPS limit for this operation is 5,000 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - BindGxb700Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindGxb700Response
func (client *Client) BindGxb700WithOptions(request *BindGxb700Request, runtime *dara.RuntimeOptions) (_result *BindGxb700Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AsrModelId) {
		query["AsrModelId"] = request.AsrModelId
	}

	if !dara.IsNil(request.Audio) {
		query["Audio"] = request.Audio
	}

	if !dara.IsNil(request.CallRestrict) {
		query["CallRestrict"] = request.CallRestrict
	}

	if !dara.IsNil(request.CallTimeout) {
		query["CallTimeout"] = request.CallTimeout
	}

	if !dara.IsNil(request.DefaultA) {
		query["DefaultA"] = request.DefaultA
	}

	if !dara.IsNil(request.DtmfConfig) {
		query["DtmfConfig"] = request.DtmfConfig
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IndustrialId) {
		query["IndustrialId"] = request.IndustrialId
	}

	if !dara.IsNil(request.NeedAsr) {
		query["NeedAsr"] = request.NeedAsr
	}

	if !dara.IsNil(request.NeedRecord) {
		query["NeedRecord"] = request.NeedRecord
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["OutOrderId"] = request.OutOrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.RecType) {
		query["RecType"] = request.RecType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TelB) {
		query["TelB"] = request.TelB
	}

	if !dara.IsNil(request.TelX) {
		query["TelX"] = request.TelX
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindGxb700"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindGxb700Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a binding relationship for a 700GXB private number. The 700GXB product is a privacy protection solutions designed for scenarios that require user-level classification, call scope restrictions, and order-grabbing restrictions. In this product, G represents a number group, and you can add numbers to the group as needed.
//
// Description:
//
// Before calling this operation, make sure that you have purchased a private number in the console.
//
// You can use this operation to bind numbers to a 700GXB private number, or bind numbers on the Number Management page in the console.
//
// Before binding numbers for a 700GXB private number, clarify the G role and B role in your business scenario. For example, on a recruitment platform, B is the job applicant and G is the recruitment team.
//
// # QPS limit
//
// The single-user QPS limit for this operation is 5,000 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - BindGxb700Request
//
// @return BindGxb700Response
func (client *Client) BindGxb700(request *BindGxb700Request) (_result *BindGxb700Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindGxb700Response{}
	_body, _err := client.BindGxb700WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Establishes a binding relationship between a specified work number X and an employee number B on the work number platform. The binding takes effect only after the employee completes real-name verification for the work number. After the binding succeeds, all calls to the work number are forwarded to the employee number. The platform pushes the binding result to the caller through the asynchronous notification API.
//
// @param request - BindXBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindXBResponse
func (client *Client) BindXBWithOptions(request *BindXBRequest, runtime *dara.RuntimeOptions) (_result *BindXBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallerParentId) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !dara.IsNil(request.CustomerPoolKey) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReqId) {
		query["ReqId"] = request.ReqId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TelB) {
		query["TelB"] = request.TelB
	}

	if !dara.IsNil(request.TelX) {
		query["TelX"] = request.TelX
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindXB"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindXBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Establishes a binding relationship between a specified work number X and an employee number B on the work number platform. The binding takes effect only after the employee completes real-name verification for the work number. After the binding succeeds, all calls to the work number are forwarded to the employee number. The platform pushes the binding result to the caller through the asynchronous notification API.
//
// @param request - BindXBRequest
//
// @return BindXBResponse
func (client *Client) BindXB(request *BindXBRequest) (_result *BindXBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindXBResponse{}
	_body, _err := client.BindXBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Purchases a phone number.
//
// Description:
//
//	  After you create a phone number pool in the Phone Number Protection console, the phone number pool is empty by default. You must purchase phone numbers and add them to the phone number pool.
//
//		- Before you call this operation, make sure that you are familiar with the [pricing](https://help.aliyun.com/document_detail/59825.html) of Phone Number Protection.
//
//		- When purchasing a phone number, specify the home location. If no sufficient phone numbers are available for purchase in the home location, the purchase of the phone number fails. Before you call this operation to purchase a phone number, check the quantity of phone numbers available for purchase by using the [QuerySecretNoRemain](https://help.aliyun.com/document_detail/111699.html) operation.
//
//		- The account used to purchase a phone number must be an enterprise account that has passed real-name verification. For more information about how to perform real-name verification, see [Enterprise verification FAQs](https://help.aliyun.com/document_detail/37172.html).
//
// @param request - BuySecretNoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BuySecretNoResponse
func (client *Client) BuySecretNoWithOptions(request *BuySecretNoRequest, runtime *dara.RuntimeOptions) (_result *BuySecretNoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.DisplayPool) {
		query["DisplayPool"] = request.DisplayPool
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecretNo) {
		query["SecretNo"] = request.SecretNo
	}

	if !dara.IsNil(request.SpecId) {
		query["SpecId"] = request.SpecId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BuySecretNo"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BuySecretNoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Purchases a phone number.
//
// Description:
//
//	  After you create a phone number pool in the Phone Number Protection console, the phone number pool is empty by default. You must purchase phone numbers and add them to the phone number pool.
//
//		- Before you call this operation, make sure that you are familiar with the [pricing](https://help.aliyun.com/document_detail/59825.html) of Phone Number Protection.
//
//		- When purchasing a phone number, specify the home location. If no sufficient phone numbers are available for purchase in the home location, the purchase of the phone number fails. Before you call this operation to purchase a phone number, check the quantity of phone numbers available for purchase by using the [QuerySecretNoRemain](https://help.aliyun.com/document_detail/111699.html) operation.
//
//		- The account used to purchase a phone number must be an enterprise account that has passed real-name verification. For more information about how to perform real-name verification, see [Enterprise verification FAQs](https://help.aliyun.com/document_detail/37172.html).
//
// @param request - BuySecretNoRequest
//
// @return BuySecretNoResponse
func (client *Client) BuySecretNo(request *BuySecretNoRequest) (_result *BuySecretNoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BuySecretNoResponse{}
	_body, _err := client.BuySecretNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a door-to-door delivery order.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CancelPickUpWaybillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelPickUpWaybillResponse
func (client *Client) CancelPickUpWaybillWithOptions(request *CancelPickUpWaybillRequest, runtime *dara.RuntimeOptions) (_result *CancelPickUpWaybillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CancelDesc) {
		query["CancelDesc"] = request.CancelDesc
	}

	if !dara.IsNil(request.OuterOrderCode) {
		query["OuterOrderCode"] = request.OuterOrderCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelPickUpWaybill"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelPickUpWaybillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a door-to-door delivery order.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CancelPickUpWaybillRequest
//
// @return CancelPickUpWaybillResponse
func (client *Client) CancelPickUpWaybill(request *CancelPickUpWaybillRequest) (_result *CancelPickUpWaybillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelPickUpWaybillResponse{}
	_body, _err := client.CancelPickUpWaybillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures an X number to independently manage communication features such as voice calls and enterprise business cards for a work number.
//
// @param tmpReq - ConfigXRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigXResponse
func (client *Client) ConfigXWithOptions(tmpReq *ConfigXRequest, runtime *dara.RuntimeOptions) (_result *ConfigXResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ConfigXShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SequenceCalls) {
		request.SequenceCallsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SequenceCalls, dara.String("SequenceCalls"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CallAbility) {
		query["CallAbility"] = request.CallAbility
	}

	if !dara.IsNil(request.CallerParentId) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !dara.IsNil(request.CustomerPoolKey) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !dara.IsNil(request.GNFlag) {
		query["GNFlag"] = request.GNFlag
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReqId) {
		query["ReqId"] = request.ReqId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SequenceCallsShrink) {
		query["SequenceCalls"] = request.SequenceCallsShrink
	}

	if !dara.IsNil(request.SequenceMode) {
		query["SequenceMode"] = request.SequenceMode
	}

	if !dara.IsNil(request.SmsAbility) {
		query["SmsAbility"] = request.SmsAbility
	}

	if !dara.IsNil(request.SmsSignMode) {
		query["SmsSignMode"] = request.SmsSignMode
	}

	if !dara.IsNil(request.TelX) {
		query["TelX"] = request.TelX
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigX"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigXResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures an X number to independently manage communication features such as voice calls and enterprise business cards for a work number.
//
// @param request - ConfigXRequest
//
// @return ConfigXResponse
func (client *Client) ConfigX(request *ConfigXRequest) (_result *ConfigXResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigXResponse{}
	_body, _err := client.ConfigXWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates number group G.
//
// Description:
//
// Before you add an AXG binding, you must purchase phone number X, create number group G, and then add phone numbers to number group G. If you do not add phone numbers to number group G after you create number group G, you can call the [OperateAxgGroup](https://help.aliyun.com/document_detail/110252.htm) operation to add phone numbers to number group G.
//
// >  Up to 2,000 number groups G can be added for a single phone number pool.
//
// @param request - CreateAxgGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAxgGroupResponse
func (client *Client) CreateAxgGroupWithOptions(request *CreateAxgGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateAxgGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Numbers) {
		query["Numbers"] = request.Numbers
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAxgGroup"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAxgGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates number group G.
//
// Description:
//
// Before you add an AXG binding, you must purchase phone number X, create number group G, and then add phone numbers to number group G. If you do not add phone numbers to number group G after you create number group G, you can call the [OperateAxgGroup](https://help.aliyun.com/document_detail/110252.htm) operation to add phone numbers to number group G.
//
// >  Up to 2,000 number groups G can be added for a single phone number pool.
//
// @param request - CreateAxgGroupRequest
//
// @return CreateAxgGroupResponse
func (client *Client) CreateAxgGroup(request *CreateAxgGroupRequest) (_result *CreateAxgGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAxgGroupResponse{}
	_body, _err := client.CreateAxgGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submit Individual Information of a landline number A through an API.
//
// Description:
//
// Before invoking this API, you need to first call [GetDyplsOSSInfoForUploadFile](https://help.aliyun.com/document_detail/2880885.html) to upload the individual\\"s photo information to OSS. For the detailed operation flow, see [Submit Individual Real-name Information for Number A via API](https://help.aliyun.com/document_detail/2880077.html).
//
// ### queries per second (QPS) limit
//
// The QPS limit for this API is 50 queries per second per user. If this limit is exceeded, Rate Limiting will be applied to your API Call, which may Impact your business. Please invoke the API appropriately.
//
// @param request - CreateFixedNoAReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFixedNoAReportResponse
func (client *Client) CreateFixedNoAReportWithOptions(request *CreateFixedNoAReportRequest, runtime *dara.RuntimeOptions) (_result *CreateFixedNoAReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ANoWhiteGroupId) {
		query["ANoWhiteGroupId"] = request.ANoWhiteGroupId
	}

	if !dara.IsNil(request.CustName) {
		query["CustName"] = request.CustName
	}

	if !dara.IsNil(request.CustPhoneNo) {
		query["CustPhoneNo"] = request.CustPhoneNo
	}

	if !dara.IsNil(request.CustType) {
		query["CustType"] = request.CustType
	}

	if !dara.IsNil(request.DocumentNumber) {
		query["DocumentNumber"] = request.DocumentNumber
	}

	if !dara.IsNil(request.DocumentType) {
		query["DocumentType"] = request.DocumentType
	}

	if !dara.IsNil(request.FixedLineWorkId) {
		query["FixedLineWorkId"] = request.FixedLineWorkId
	}

	if !dara.IsNil(request.FixedNoA) {
		query["FixedNoA"] = request.FixedNoA
	}

	if !dara.IsNil(request.IdCardAlivePhoto) {
		query["IdCardAlivePhoto"] = request.IdCardAlivePhoto
	}

	if !dara.IsNil(request.IdCardBackPhoto) {
		query["IdCardBackPhoto"] = request.IdCardBackPhoto
	}

	if !dara.IsNil(request.IdCardFrontPhoto) {
		query["IdCardFrontPhoto"] = request.IdCardFrontPhoto
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFixedNoAReport"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFixedNoAReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submit Individual Information of a landline number A through an API.
//
// Description:
//
// Before invoking this API, you need to first call [GetDyplsOSSInfoForUploadFile](https://help.aliyun.com/document_detail/2880885.html) to upload the individual\\"s photo information to OSS. For the detailed operation flow, see [Submit Individual Real-name Information for Number A via API](https://help.aliyun.com/document_detail/2880077.html).
//
// ### queries per second (QPS) limit
//
// The QPS limit for this API is 50 queries per second per user. If this limit is exceeded, Rate Limiting will be applied to your API Call, which may Impact your business. Please invoke the API appropriately.
//
// @param request - CreateFixedNoAReportRequest
//
// @return CreateFixedNoAReportResponse
func (client *Client) CreateFixedNoAReport(request *CreateFixedNoAReportRequest) (_result *CreateFixedNoAReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFixedNoAReportResponse{}
	_body, _err := client.CreateFixedNoAReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submit Individual Information of an A-number mobile phone number through an API.
//
// Description:
//
// Before invoking this API, you must first call [GetDyplsOSSInfoForUploadFile](https://help.aliyun.com/document_detail/2880885.html) to upload the individual\\"s photo information to OSS. For the detailed operation flow, see [Submit Individual Real-name Information of an A-number Mobile Phone Number via API](https://help.aliyun.com/document_detail/2880077.html).
//
// ### QPS Limit
//
// The queries per second (QPS) limit for this API is 50 requests per second per user. If this limit is exceeded, API calls will be subject to Rate Limiting, which may Impact your business. Please invoke the API appropriately.
//
// @param request - CreatePhoneNoAReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePhoneNoAReportResponse
func (client *Client) CreatePhoneNoAReportWithOptions(request *CreatePhoneNoAReportRequest, runtime *dara.RuntimeOptions) (_result *CreatePhoneNoAReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ANoWhiteGroupId) {
		query["ANoWhiteGroupId"] = request.ANoWhiteGroupId
	}

	if !dara.IsNil(request.CustName) {
		query["CustName"] = request.CustName
	}

	if !dara.IsNil(request.DocumentNumber) {
		query["DocumentNumber"] = request.DocumentNumber
	}

	if !dara.IsNil(request.DocumentType) {
		query["DocumentType"] = request.DocumentType
	}

	if !dara.IsNil(request.IdCardAlivePhoto) {
		query["IdCardAlivePhoto"] = request.IdCardAlivePhoto
	}

	if !dara.IsNil(request.IdCardBackPhoto) {
		query["IdCardBackPhoto"] = request.IdCardBackPhoto
	}

	if !dara.IsNil(request.IdCardFrontPhoto) {
		query["IdCardFrontPhoto"] = request.IdCardFrontPhoto
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNoA) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePhoneNoAReport"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePhoneNoAReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submit Individual Information of an A-number mobile phone number through an API.
//
// Description:
//
// Before invoking this API, you must first call [GetDyplsOSSInfoForUploadFile](https://help.aliyun.com/document_detail/2880885.html) to upload the individual\\"s photo information to OSS. For the detailed operation flow, see [Submit Individual Real-name Information of an A-number Mobile Phone Number via API](https://help.aliyun.com/document_detail/2880077.html).
//
// ### QPS Limit
//
// The queries per second (QPS) limit for this API is 50 requests per second per user. If this limit is exceeded, API calls will be subject to Rate Limiting, which may Impact your business. Please invoke the API appropriately.
//
// @param request - CreatePhoneNoAReportRequest
//
// @return CreatePhoneNoAReportResponse
func (client *Client) CreatePhoneNoAReport(request *CreatePhoneNoAReportRequest) (_result *CreatePhoneNoAReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePhoneNoAReportResponse{}
	_body, _err := client.CreatePhoneNoAReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a door-to-door pickup order.
//
// Description:
//
// ### QPS limits
//
// The queries per second (QPS) limit for a single user is 200 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call the API at a reasonable frequency.
//
// @param tmpReq - CreatePickUpWaybillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePickUpWaybillResponse
func (client *Client) CreatePickUpWaybillWithOptions(tmpReq *CreatePickUpWaybillRequest, runtime *dara.RuntimeOptions) (_result *CreatePickUpWaybillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePickUpWaybillShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ConsigneeAddress) {
		request.ConsigneeAddressShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConsigneeAddress, dara.String("ConsigneeAddress"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.GoodsInfos) {
		request.GoodsInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GoodsInfos, dara.String("GoodsInfos"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SendAddress) {
		request.SendAddressShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SendAddress, dara.String("SendAddress"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppointGotEndTime) {
		query["AppointGotEndTime"] = request.AppointGotEndTime
	}

	if !dara.IsNil(request.AppointGotStartTime) {
		query["AppointGotStartTime"] = request.AppointGotStartTime
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ConsigneeAddressShrink) {
		query["ConsigneeAddress"] = request.ConsigneeAddressShrink
	}

	if !dara.IsNil(request.ConsigneeMobile) {
		query["ConsigneeMobile"] = request.ConsigneeMobile
	}

	if !dara.IsNil(request.ConsigneeName) {
		query["ConsigneeName"] = request.ConsigneeName
	}

	if !dara.IsNil(request.ConsigneePhone) {
		query["ConsigneePhone"] = request.ConsigneePhone
	}

	if !dara.IsNil(request.CpCode) {
		query["CpCode"] = request.CpCode
	}

	if !dara.IsNil(request.GoodsInfosShrink) {
		query["GoodsInfos"] = request.GoodsInfosShrink
	}

	if !dara.IsNil(request.OrderChannels) {
		query["OrderChannels"] = request.OrderChannels
	}

	if !dara.IsNil(request.OuterOrderCode) {
		query["OuterOrderCode"] = request.OuterOrderCode
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SendAddressShrink) {
		query["SendAddress"] = request.SendAddressShrink
	}

	if !dara.IsNil(request.SendMobile) {
		query["SendMobile"] = request.SendMobile
	}

	if !dara.IsNil(request.SendName) {
		query["SendName"] = request.SendName
	}

	if !dara.IsNil(request.SendPhone) {
		query["SendPhone"] = request.SendPhone
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePickUpWaybill"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePickUpWaybillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a door-to-door pickup order.
//
// Description:
//
// ### QPS limits
//
// The queries per second (QPS) limit for a single user is 200 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call the API at a reasonable frequency.
//
// @param request - CreatePickUpWaybillRequest
//
// @return CreatePickUpWaybillResponse
func (client *Client) CreatePickUpWaybill(request *CreatePickUpWaybillRequest) (_result *CreatePickUpWaybillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePickUpWaybillResponse{}
	_body, _err := client.CreatePickUpWaybillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pre-queries shipping order information before placing an order.
//
// Description:
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 100 calls per second. Exceeding this limit will trigger throttling, which may affect your business. Call this operation appropriately.
//
// @param tmpReq - CreatePickUpWaybillPreQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePickUpWaybillPreQueryResponse
func (client *Client) CreatePickUpWaybillPreQueryWithOptions(tmpReq *CreatePickUpWaybillPreQueryRequest, runtime *dara.RuntimeOptions) (_result *CreatePickUpWaybillPreQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePickUpWaybillPreQueryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ConsigneeInfo) {
		request.ConsigneeInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConsigneeInfo, dara.String("ConsigneeInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SenderInfo) {
		request.SenderInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SenderInfo, dara.String("SenderInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsigneeInfoShrink) {
		query["ConsigneeInfo"] = request.ConsigneeInfoShrink
	}

	if !dara.IsNil(request.CpCode) {
		query["CpCode"] = request.CpCode
	}

	if !dara.IsNil(request.OrderChannels) {
		query["OrderChannels"] = request.OrderChannels
	}

	if !dara.IsNil(request.OuterOrderCode) {
		query["OuterOrderCode"] = request.OuterOrderCode
	}

	if !dara.IsNil(request.PreWeight) {
		query["PreWeight"] = request.PreWeight
	}

	if !dara.IsNil(request.SenderInfoShrink) {
		query["SenderInfo"] = request.SenderInfoShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePickUpWaybillPreQuery"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePickUpWaybillPreQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pre-queries shipping order information before placing an order.
//
// Description:
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 100 calls per second. Exceeding this limit will trigger throttling, which may affect your business. Call this operation appropriately.
//
// @param request - CreatePickUpWaybillPreQueryRequest
//
// @return CreatePickUpWaybillPreQueryResponse
func (client *Client) CreatePickUpWaybillPreQuery(request *CreatePickUpWaybillPreQueryRequest) (_result *CreatePickUpWaybillPreQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePickUpWaybillPreQueryResponse{}
	_body, _err := client.CreatePickUpWaybillPreQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 对客提交号码迁移任务，两侧号池须为同一资质、同一场景且均为开启状态。
//
// @param request - CreateSecretTransferRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecretTransferRecordResponse
func (client *Client) CreateSecretTransferRecordWithOptions(request *CreateSecretTransferRecordRequest, runtime *dara.RuntimeOptions) (_result *CreateSecretTransferRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LastPartnerKey) {
		query["LastPartnerKey"] = request.LastPartnerKey
	}

	if !dara.IsNil(request.NumberList) {
		query["NumberList"] = request.NumberList
	}

	if !dara.IsNil(request.OriginPartnerKey) {
		query["OriginPartnerKey"] = request.OriginPartnerKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSecretTransferRecord"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecretTransferRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对客提交号码迁移任务，两侧号池须为同一资质、同一场景且均为开启状态。
//
// @param request - CreateSecretTransferRecordRequest
//
// @return CreateSecretTransferRecordResponse
func (client *Client) CreateSecretTransferRecord(request *CreateSecretTransferRecordRequest) (_result *CreateSecretTransferRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSecretTransferRecordResponse{}
	_body, _err := client.CreateSecretTransferRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a text message delivery tag for scenarios where an employee sends a text message to a customer by using a work number. The client calls this operation to obtain the tag and appends it to the end of the message content. The work number platform then parses the actual called number from the tag and delivers the message to the customer.
//
// @param request - CreateSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmsSignResponse
func (client *Client) CreateSmsSignWithOptions(request *CreateSmsSignRequest, runtime *dara.RuntimeOptions) (_result *CreateSmsSignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalledNo) {
		query["CalledNo"] = request.CalledNo
	}

	if !dara.IsNil(request.CallerParentId) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !dara.IsNil(request.CallingNo) {
		query["CallingNo"] = request.CallingNo
	}

	if !dara.IsNil(request.CustomerPoolKey) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReqId) {
		query["ReqId"] = request.ReqId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSmsSign"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a text message delivery tag for scenarios where an employee sends a text message to a customer by using a work number. The client calls this operation to obtain the tag and appends it to the end of the message content. The work number platform then parses the actual called number from the tag and delivers the message to the customer.
//
// @param request - CreateSmsSignRequest
//
// @return CreateSmsSignResponse
func (client *Client) CreateSmsSign(request *CreateSmsSignRequest) (_result *CreateSmsSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSmsSignResponse{}
	_body, _err := client.CreateSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbind a fixed-line AXB privacy number.
//
// Description:
//
// - After unbinding a fixed-line AXB privacy number from an existing binding, your business continuity may be affected. Proceed with caution.
//
// - Before releasing a number, you must unbind the number first.
//
// - Unbinding a fixed-line AXB privacy number is supported only through this API. Unbinding through the console is currently not supported.
//
// @param request - DeleteAxbBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAxbBindFixedLineResponse
func (client *Client) DeleteAxbBindFixedLineWithOptions(request *DeleteAxbBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *DeleteAxbBindFixedLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SubId) {
		query["SubId"] = request.SubId
	}

	if !dara.IsNil(request.Ts) {
		query["Ts"] = request.Ts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAxbBindFixedLine"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAxbBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbind a fixed-line AXB privacy number.
//
// Description:
//
// - After unbinding a fixed-line AXB privacy number from an existing binding, your business continuity may be affected. Proceed with caution.
//
// - Before releasing a number, you must unbind the number first.
//
// - Unbinding a fixed-line AXB privacy number is supported only through this API. Unbinding through the console is currently not supported.
//
// @param request - DeleteAxbBindFixedLineRequest
//
// @return DeleteAxbBindFixedLineResponse
func (client *Client) DeleteAxbBindFixedLine(request *DeleteAxbBindFixedLineRequest) (_result *DeleteAxbBindFixedLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAxbBindFixedLineResponse{}
	_body, _err := client.DeleteAxbBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a G group. G groups are supported only in Number Pools designated for the AXG privacy number type.
//
// Description:
//
// - Deleted G groups cannot be recovered. Proceed with caution.
//
// - Before calling this operation, verify the G group to be deleted and the Number Pool it belongs to.
//
// - You cannot delete a G group that has active bindings. Before you delete the group, you must unbind the associated AXG privacy number. To do this, log on to the Phone Number Privacy Protection [Console](https://dyplsnext.console.aliyun.com/number/list), go to the **Number Management*	- page, filter by **AXG privacy number**, and click **Details*	- > Unbind. Alternatively, you can call the [UnbindSubscription](https://help.aliyun.com/zh/pnp/developer-reference/api-dyplsapi-2017-05-25-unbindsubscription) API to unbind the number.
//
// - When you delete a G group, all phone numbers in the group are also deleted.
//
// ### QPS limit
//
// The single-user QPS limit for this API is 5,000 requests per second. API calls that exceed this limit are subject to throttling. Plan your API calls accordingly to prevent disruptions to your business.
//
// @param request - DeleteAxgGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAxgGroupResponse
func (client *Client) DeleteAxgGroupWithOptions(request *DeleteAxgGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteAxgGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAxgGroup"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAxgGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a G group. G groups are supported only in Number Pools designated for the AXG privacy number type.
//
// Description:
//
// - Deleted G groups cannot be recovered. Proceed with caution.
//
// - Before calling this operation, verify the G group to be deleted and the Number Pool it belongs to.
//
// - You cannot delete a G group that has active bindings. Before you delete the group, you must unbind the associated AXG privacy number. To do this, log on to the Phone Number Privacy Protection [Console](https://dyplsnext.console.aliyun.com/number/list), go to the **Number Management*	- page, filter by **AXG privacy number**, and click **Details*	- > Unbind. Alternatively, you can call the [UnbindSubscription](https://help.aliyun.com/zh/pnp/developer-reference/api-dyplsapi-2017-05-25-unbindsubscription) API to unbind the number.
//
// - When you delete a G group, all phone numbers in the group are also deleted.
//
// ### QPS limit
//
// The single-user QPS limit for this API is 5,000 requests per second. API calls that exceed this limit are subject to throttling. Plan your API calls accordingly to prevent disruptions to your business.
//
// @param request - DeleteAxgGroupRequest
//
// @return DeleteAxgGroupResponse
func (client *Client) DeleteAxgGroup(request *DeleteAxgGroupRequest) (_result *DeleteAxgGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAxgGroupResponse{}
	_body, _err := client.DeleteAxgGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds a fixed-line AXN private number binding relationship.
//
// Description:
//
// - After you unbind an existing fixed-line AXN private number binding relationship, your service continuity may be affected. Proceed with caution.
//
// - Before releasing a number, unbind the number binding relationship.
//
// - Unbinding a fixed-line AXN private number binding relationship is supported only through this API operation. The console does not support this operation.
//
// @param request - DeleteAxnBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAxnBindFixedLineResponse
func (client *Client) DeleteAxnBindFixedLineWithOptions(request *DeleteAxnBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *DeleteAxnBindFixedLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SubId) {
		query["SubId"] = request.SubId
	}

	if !dara.IsNil(request.Ts) {
		query["Ts"] = request.Ts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAxnBindFixedLine"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAxnBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a fixed-line AXN private number binding relationship.
//
// Description:
//
// - After you unbind an existing fixed-line AXN private number binding relationship, your service continuity may be affected. Proceed with caution.
//
// - Before releasing a number, unbind the number binding relationship.
//
// - Unbinding a fixed-line AXN private number binding relationship is supported only through this API operation. The console does not support this operation.
//
// @param request - DeleteAxnBindFixedLineRequest
//
// @return DeleteAxnBindFixedLineResponse
func (client *Client) DeleteAxnBindFixedLine(request *DeleteAxnBindFixedLineRequest) (_result *DeleteAxnBindFixedLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAxnBindFixedLineResponse{}
	_body, _err := client.DeleteAxnBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbind a fixed-line AXN extension number binding.
//
// @param request - DeleteAxnExtensionBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAxnExtensionBindFixedLineResponse
func (client *Client) DeleteAxnExtensionBindFixedLineWithOptions(request *DeleteAxnExtensionBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *DeleteAxnExtensionBindFixedLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SubId) {
		query["SubId"] = request.SubId
	}

	if !dara.IsNil(request.Ts) {
		query["Ts"] = request.Ts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAxnExtensionBindFixedLine"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAxnExtensionBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbind a fixed-line AXN extension number binding.
//
// @param request - DeleteAxnExtensionBindFixedLineRequest
//
// @return DeleteAxnExtensionBindFixedLineResponse
func (client *Client) DeleteAxnExtensionBindFixedLine(request *DeleteAxnExtensionBindFixedLineRequest) (_result *DeleteAxnExtensionBindFixedLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAxnExtensionBindFixedLineResponse{}
	_body, _err := client.DeleteAxnExtensionBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete A-number filing data.
//
// Description:
//
// - This API can only delete A-number filing records whose filing status is either "Reviewed and Passed" or "Review Failed." It does not support deleting A-number filing records with a filing status of "Under Review."
//
// - After deleting an A-number filing record, the carrier will immediately impose call restrictions or deactivate the number, which may impact your business continuity. Proceed with caution.
//
// @param request - DeleteSecretAPhoneNoToCustRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecretAPhoneNoToCustResponse
func (client *Client) DeleteSecretAPhoneNoToCustWithOptions(request *DeleteSecretAPhoneNoToCustRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecretAPhoneNoToCustResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ANoWhiteGroupId) {
		query["ANoWhiteGroupId"] = request.ANoWhiteGroupId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNoA) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecretAPhoneNoToCust"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecretAPhoneNoToCustResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete A-number filing data.
//
// Description:
//
// - This API can only delete A-number filing records whose filing status is either "Reviewed and Passed" or "Review Failed." It does not support deleting A-number filing records with a filing status of "Under Review."
//
// - After deleting an A-number filing record, the carrier will immediately impose call restrictions or deactivate the number, which may impact your business continuity. Proceed with caution.
//
// @param request - DeleteSecretAPhoneNoToCustRequest
//
// @return DeleteSecretAPhoneNoToCustResponse
func (client *Client) DeleteSecretAPhoneNoToCust(request *DeleteSecretAPhoneNoToCustRequest) (_result *DeleteSecretAPhoneNoToCustResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSecretAPhoneNoToCustResponse{}
	_body, _err := client.DeleteSecretAPhoneNoToCustWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the blacklist of a number pool. Only number pools of the AXN privacy number, AXN extension number, and 95AXN privacy number types support blacklists.
//
// Description:
//
// - This operation is not a batch operation. It supports deleting only one pair of point-to-point blacklist numbers or one number pool blacklist number at a time.
//
// - To add a blacklist, call the [AddSecretBlacklist](https://www.alibabacloud.com/help/en/pnp/developer-reference/api-dyplsapi-2017-05-25-addsecretblacklist) operation.
//
// ### QPS limit
//
// The QPS limit for a single user is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - DeleteSecretBlacklistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecretBlacklistResponse
func (client *Client) DeleteSecretBlacklistWithOptions(request *DeleteSecretBlacklistRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecretBlacklistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BlackNo) {
		query["BlackNo"] = request.BlackNo
	}

	if !dara.IsNil(request.BlackType) {
		query["BlackType"] = request.BlackType
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.WayControl) {
		query["WayControl"] = request.WayControl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecretBlacklist"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecretBlacklistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the blacklist of a number pool. Only number pools of the AXN privacy number, AXN extension number, and 95AXN privacy number types support blacklists.
//
// Description:
//
// - This operation is not a batch operation. It supports deleting only one pair of point-to-point blacklist numbers or one number pool blacklist number at a time.
//
// - To add a blacklist, call the [AddSecretBlacklist](https://www.alibabacloud.com/help/en/pnp/developer-reference/api-dyplsapi-2017-05-25-addsecretblacklist) operation.
//
// ### QPS limit
//
// The QPS limit for a single user is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - DeleteSecretBlacklistRequest
//
// @return DeleteSecretBlacklistResponse
func (client *Client) DeleteSecretBlacklist(request *DeleteSecretBlacklistRequest) (_result *DeleteSecretBlacklistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSecretBlacklistResponse{}
	_body, _err := client.DeleteSecretBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain OSS resource configuration information. This configuration information will be used for subsequent OSS image upload operations.
//
// Description:
//
// - When uploading personal filing photos for an A-number mobile phone number, you must upload a clear frontal portrait photo with the face unobscured, and optionally select to upload photos of both sides of your ID card. Uploading these photos helps reviewers understand your personal identity information. Supported formats are PNG, JPG, and JPEG. The image size must not exceed 2 MB.
//
// - For the specific operation flow when submitting personal real-name filing for an A-number mobile phone number through the API, see [Submit Personal Real-Name Information for an A-Number Mobile Phone Number via API](https://help.aliyun.com/document_detail/2880077.html).
//
// @param request - GetDyplsOSSInfoForUploadFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDyplsOSSInfoForUploadFileResponse
func (client *Client) GetDyplsOSSInfoForUploadFileWithOptions(request *GetDyplsOSSInfoForUploadFileRequest, runtime *dara.RuntimeOptions) (_result *GetDyplsOSSInfoForUploadFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDyplsOSSInfoForUploadFile"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDyplsOSSInfoForUploadFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain OSS resource configuration information. This configuration information will be used for subsequent OSS image upload operations.
//
// Description:
//
// - When uploading personal filing photos for an A-number mobile phone number, you must upload a clear frontal portrait photo with the face unobscured, and optionally select to upload photos of both sides of your ID card. Uploading these photos helps reviewers understand your personal identity information. Supported formats are PNG, JPG, and JPEG. The image size must not exceed 2 MB.
//
// - For the specific operation flow when submitting personal real-name filing for an A-number mobile phone number through the API, see [Submit Personal Real-Name Information for an A-Number Mobile Phone Number via API](https://help.aliyun.com/document_detail/2880077.html).
//
// @param request - GetDyplsOSSInfoForUploadFileRequest
//
// @return GetDyplsOSSInfoForUploadFileResponse
func (client *Client) GetDyplsOSSInfoForUploadFile(request *GetDyplsOSSInfoForUploadFileRequest) (_result *GetDyplsOSSInfoForUploadFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDyplsOSSInfoForUploadFileResponse{}
	_body, _err := client.GetDyplsOSSInfoForUploadFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of the automatic speech recognition (ASR) result.
//
// Description:
//
// Before you call the GetSecretAsrDetail operation, set the ASRStatus parameter to true in the [BindAxn operation](https://help.aliyun.com/document_detail/400483.html). This ensures that you can obtain the ASR result properly.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetSecretAsrDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecretAsrDetailResponse
func (client *Client) GetSecretAsrDetailWithOptions(request *GetSecretAsrDetailRequest, runtime *dara.RuntimeOptions) (_result *GetSecretAsrDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CallTime) {
		query["CallTime"] = request.CallTime
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecretAsrDetail"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecretAsrDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of the automatic speech recognition (ASR) result.
//
// Description:
//
// Before you call the GetSecretAsrDetail operation, set the ASRStatus parameter to true in the [BindAxn operation](https://help.aliyun.com/document_detail/400483.html). This ensures that you can obtain the ASR result properly.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetSecretAsrDetailRequest
//
// @return GetSecretAsrDetailResponse
func (client *Client) GetSecretAsrDetail(request *GetSecretAsrDetailRequest) (_result *GetSecretAsrDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSecretAsrDetailResponse{}
	_body, _err := client.GetSecretAsrDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the download URL of a recorded ringing tone.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetTotalPublicUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTotalPublicUrlResponse
func (client *Client) GetTotalPublicUrlWithOptions(request *GetTotalPublicUrlRequest, runtime *dara.RuntimeOptions) (_result *GetTotalPublicUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CallTime) {
		query["CallTime"] = request.CallTime
	}

	if !dara.IsNil(request.CheckSubs) {
		query["CheckSubs"] = request.CheckSubs
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PartnerKey) {
		query["PartnerKey"] = request.PartnerKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTotalPublicUrl"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTotalPublicUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the download URL of a recorded ringing tone.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetTotalPublicUrlRequest
//
// @return GetTotalPublicUrlResponse
func (client *Client) GetTotalPublicUrl(request *GetTotalPublicUrlRequest) (_result *GetTotalPublicUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTotalPublicUrlResponse{}
	_body, _err := client.GetTotalPublicUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the configuration information of the X number.
//
// @param request - GetXConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetXConfigResponse
func (client *Client) GetXConfigWithOptions(request *GetXConfigRequest, runtime *dara.RuntimeOptions) (_result *GetXConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallerParentId) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !dara.IsNil(request.CustomerPoolKey) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReqId) {
		query["ReqId"] = request.ReqId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TelX) {
		query["TelX"] = request.TelX
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetXConfig"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetXConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the configuration information of the X number.
//
// @param request - GetXConfigRequest
//
// @return GetXConfigResponse
func (client *Client) GetXConfig(request *GetXConfigRequest) (_result *GetXConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetXConfigResponse{}
	_body, _err := client.GetXConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the default configurations of an X number.
//
// @param request - GetXDefaultConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetXDefaultConfigResponse
func (client *Client) GetXDefaultConfigWithOptions(request *GetXDefaultConfigRequest, runtime *dara.RuntimeOptions) (_result *GetXDefaultConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallerParentId) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !dara.IsNil(request.CustomerPoolKey) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReqId) {
		query["ReqId"] = request.ReqId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TelX) {
		query["TelX"] = request.TelX
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetXDefaultConfig"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetXDefaultConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the default configurations of an X number.
//
// @param request - GetXDefaultConfigRequest
//
// @return GetXDefaultConfigResponse
func (client *Client) GetXDefaultConfig(request *GetXDefaultConfigRequest) (_result *GetXDefaultConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetXDefaultConfigResponse{}
	_body, _err := client.GetXDefaultConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of X numbers under a customer account.
//
// @param request - ListXTelephonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListXTelephonesResponse
func (client *Client) ListXTelephonesWithOptions(request *ListXTelephonesRequest, runtime *dara.RuntimeOptions) (_result *ListXTelephonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallerParentId) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !dara.IsNil(request.CustomerPoolKey) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReqId) {
		query["ReqId"] = request.ReqId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListXTelephones"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListXTelephonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of X numbers under a customer account.
//
// @param request - ListXTelephonesRequest
//
// @return ListXTelephonesResponse
func (client *Client) ListXTelephones(request *ListXTelephonesRequest) (_result *ListXTelephonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListXTelephonesResponse{}
	_body, _err := client.ListXTelephonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Locks a phone number.
//
// Description:
//
// After a phone number is locked, the locked phone number cannot be selected when you call an operation to create a binding.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - LockSecretNoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LockSecretNoResponse
func (client *Client) LockSecretNoWithOptions(request *LockSecretNoRequest, runtime *dara.RuntimeOptions) (_result *LockSecretNoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecretNo) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LockSecretNo"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LockSecretNoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Locks a phone number.
//
// Description:
//
// After a phone number is locked, the locked phone number cannot be selected when you call an operation to create a binding.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - LockSecretNoRequest
//
// @return LockSecretNoResponse
func (client *Client) LockSecretNo(request *LockSecretNoRequest) (_result *LockSecretNoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LockSecretNoResponse{}
	_body, _err := client.LockSecretNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies number group G.
//
// Description:
//
// After you create number group G, you can call the OperateAxgGroup operation to modify number group G. For example, you can add phone numbers to number group G, delete phone numbers from number group G, and replace all phone numbers in number group G.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - OperateAxgGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateAxgGroupResponse
func (client *Client) OperateAxgGroupWithOptions(request *OperateAxgGroupRequest, runtime *dara.RuntimeOptions) (_result *OperateAxgGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Numbers) {
		query["Numbers"] = request.Numbers
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateAxgGroup"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateAxgGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies number group G.
//
// Description:
//
// After you create number group G, you can call the OperateAxgGroup operation to modify number group G. For example, you can add phone numbers to number group G, delete phone numbers from number group G, and replace all phone numbers in number group G.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - OperateAxgGroupRequest
//
// @return OperateAxgGroupResponse
func (client *Client) OperateAxgGroup(request *OperateAxgGroupRequest) (_result *OperateAxgGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateAxgGroupResponse{}
	_body, _err := client.OperateAxgGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a phone number to a blacklist or deletes a phone number from a blacklist.
//
// Description:
//
// The OperateBlackNo operation supports the following number pool types: AXN, AXN extension, and 95AXN.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - OperateBlackNoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateBlackNoResponse
func (client *Client) OperateBlackNoWithOptions(request *OperateBlackNoRequest, runtime *dara.RuntimeOptions) (_result *OperateBlackNoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BlackNo) {
		query["BlackNo"] = request.BlackNo
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tips) {
		query["Tips"] = request.Tips
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateBlackNo"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateBlackNoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a phone number to a blacklist or deletes a phone number from a blacklist.
//
// Description:
//
// The OperateBlackNo operation supports the following number pool types: AXN, AXN extension, and 95AXN.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - OperateBlackNoRequest
//
// @return OperateBlackNoResponse
func (client *Client) OperateBlackNo(request *OperateBlackNoRequest) (_result *OperateBlackNoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateBlackNoResponse{}
	_body, _err := client.OperateBlackNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the binding information of an AXB landline privacy number.
//
// @param request - QueryAxbBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAxbBindFixedLineResponse
func (client *Client) QueryAxbBindFixedLineWithOptions(request *QueryAxbBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *QueryAxbBindFixedLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SubId) {
		query["SubId"] = request.SubId
	}

	if !dara.IsNil(request.TelX) {
		query["TelX"] = request.TelX
	}

	if !dara.IsNil(request.Ts) {
		query["Ts"] = request.Ts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAxbBindFixedLine"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAxbBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the binding information of an AXB landline privacy number.
//
// @param request - QueryAxbBindFixedLineRequest
//
// @return QueryAxbBindFixedLineResponse
func (client *Client) QueryAxbBindFixedLine(request *QueryAxbBindFixedLineRequest) (_result *QueryAxbBindFixedLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAxbBindFixedLineResponse{}
	_body, _err := client.QueryAxbBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the binding information of an AXN landline privacy number.
//
// @param request - QueryAxnBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAxnBindFixedLineResponse
func (client *Client) QueryAxnBindFixedLineWithOptions(request *QueryAxnBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *QueryAxnBindFixedLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SubId) {
		query["SubId"] = request.SubId
	}

	if !dara.IsNil(request.TelX) {
		query["TelX"] = request.TelX
	}

	if !dara.IsNil(request.Ts) {
		query["Ts"] = request.Ts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAxnBindFixedLine"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAxnBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the binding information of an AXN landline privacy number.
//
// @param request - QueryAxnBindFixedLineRequest
//
// @return QueryAxnBindFixedLineResponse
func (client *Client) QueryAxnBindFixedLine(request *QueryAxnBindFixedLineRequest) (_result *QueryAxnBindFixedLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAxnBindFixedLineResponse{}
	_body, _err := client.QueryAxnBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the binding information of a landline AXN extension number.
//
// @param request - QueryAxnExtensionBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAxnExtensionBindFixedLineResponse
func (client *Client) QueryAxnExtensionBindFixedLineWithOptions(request *QueryAxnExtensionBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *QueryAxnExtensionBindFixedLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SubId) {
		query["SubId"] = request.SubId
	}

	if !dara.IsNil(request.TelA) {
		query["TelA"] = request.TelA
	}

	if !dara.IsNil(request.Ts) {
		query["Ts"] = request.Ts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAxnExtensionBindFixedLine"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAxnExtensionBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the binding information of a landline AXN extension number.
//
// @param request - QueryAxnExtensionBindFixedLineRequest
//
// @return QueryAxnExtensionBindFixedLineResponse
func (client *Client) QueryAxnExtensionBindFixedLine(request *QueryAxnExtensionBindFixedLineRequest) (_result *QueryAxnExtensionBindFixedLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAxnExtensionBindFixedLineResponse{}
	_body, _err := client.QueryAxnExtensionBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about a tracking number.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryPhoneNoAByTrackNoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPhoneNoAByTrackNoResponse
func (client *Client) QueryPhoneNoAByTrackNoWithOptions(request *QueryPhoneNoAByTrackNoRequest, runtime *dara.RuntimeOptions) (_result *QueryPhoneNoAByTrackNoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CabinetNo) {
		query["CabinetNo"] = request.CabinetNo
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNoX) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TrackNo) {
		query["trackNo"] = request.TrackNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPhoneNoAByTrackNo"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPhoneNoAByTrackNoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a tracking number.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryPhoneNoAByTrackNoRequest
//
// @return QueryPhoneNoAByTrackNoResponse
func (client *Client) QueryPhoneNoAByTrackNo(request *QueryPhoneNoAByTrackNoRequest) (_result *QueryPhoneNoAByTrackNoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPhoneNoAByTrackNoResponse{}
	_body, _err := client.QueryPhoneNoAByTrackNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the download URL of a recording file.
//
// Description:
//
// If the recording feature is enabled for a binding, all calls made by the bound phone numbers are recorded. You can obtain the download URL of a recording file by calling the QueryRecordFileDownloadUrl operation and download the recording file.
//
// >  We recommend that you subscribe to [the recording status report SecretRecording](https://help.aliyun.com/document_detail/109198.html). The values of the response parameters in SecretRecording can be used as the values of the request parameters for downloading a recording file.
//
// ### [](#)Procedure for obtaining a recording file
//
// 1.  Specify the request parameter in an update or binding operation to enable the recording feature.
//
// 2.  Subscribe to recording message receipts in the Phone Number Protection console.
//
// 3.  After a recording message receipt is returned, call the QueryRecordFileDownloadUrl operation to obtain the download URL of the recording file, and download the recording file.
//
// >
//
//   - A download URL is valid for 2 hours. Download the recording file as soon as possible after obtaining a download URL.
//
//   - The storage period of recording files is 30 days. You can download only the recording files of calls recorded in the last 30 days.
//
// @param request - QueryRecordFileDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecordFileDownloadUrlResponse
func (client *Client) QueryRecordFileDownloadUrlWithOptions(request *QueryRecordFileDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordFileDownloadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CallTime) {
		query["CallTime"] = request.CallTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecordFileDownloadUrl"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecordFileDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the download URL of a recording file.
//
// Description:
//
// If the recording feature is enabled for a binding, all calls made by the bound phone numbers are recorded. You can obtain the download URL of a recording file by calling the QueryRecordFileDownloadUrl operation and download the recording file.
//
// >  We recommend that you subscribe to [the recording status report SecretRecording](https://help.aliyun.com/document_detail/109198.html). The values of the response parameters in SecretRecording can be used as the values of the request parameters for downloading a recording file.
//
// ### [](#)Procedure for obtaining a recording file
//
// 1.  Specify the request parameter in an update or binding operation to enable the recording feature.
//
// 2.  Subscribe to recording message receipts in the Phone Number Protection console.
//
// 3.  After a recording message receipt is returned, call the QueryRecordFileDownloadUrl operation to obtain the download URL of the recording file, and download the recording file.
//
// >
//
//   - A download URL is valid for 2 hours. Download the recording file as soon as possible after obtaining a download URL.
//
//   - The storage period of recording files is 30 days. You can download only the recording files of calls recorded in the last 30 days.
//
// @param request - QueryRecordFileDownloadUrlRequest
//
// @return QueryRecordFileDownloadUrlResponse
func (client *Client) QueryRecordFileDownloadUrl(request *QueryRecordFileDownloadUrlRequest) (_result *QueryRecordFileDownloadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRecordFileDownloadUrlResponse{}
	_body, _err := client.QueryRecordFileDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the reporting status of an A number.
//
// Description:
//
// Before you invoke this API to query the reporting status of an A number, you must first submit the real-name information for the mobile or landline A number by calling the [CreatePhoneNoAReport](https://help.aliyun.com/document_detail/2880908.html) or [CreateFixedNoAReport](https://help.aliyun.com/document_detail/2881429.html) API. For detailed instructions, see [Submit real-name information for an A number via API](https://help.aliyun.com/document_detail/2880077.html).
//
// @param request - QuerySecretAPhoneNoToCustRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySecretAPhoneNoToCustResponse
func (client *Client) QuerySecretAPhoneNoToCustWithOptions(request *QuerySecretAPhoneNoToCustRequest, runtime *dara.RuntimeOptions) (_result *QuerySecretAPhoneNoToCustResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ANoWhiteGroupId) {
		query["ANoWhiteGroupId"] = request.ANoWhiteGroupId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNoA) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySecretAPhoneNoToCust"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySecretAPhoneNoToCustResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the reporting status of an A number.
//
// Description:
//
// Before you invoke this API to query the reporting status of an A number, you must first submit the real-name information for the mobile or landline A number by calling the [CreatePhoneNoAReport](https://help.aliyun.com/document_detail/2880908.html) or [CreateFixedNoAReport](https://help.aliyun.com/document_detail/2881429.html) API. For detailed instructions, see [Submit real-name information for an A number via API](https://help.aliyun.com/document_detail/2880077.html).
//
// @param request - QuerySecretAPhoneNoToCustRequest
//
// @return QuerySecretAPhoneNoToCustResponse
func (client *Client) QuerySecretAPhoneNoToCust(request *QuerySecretAPhoneNoToCustRequest) (_result *QuerySecretAPhoneNoToCustResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySecretAPhoneNoToCustResponse{}
	_body, _err := client.QuerySecretAPhoneNoToCustWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the property information of a private number.
//
// Description:
//
// - You can query information about a private number (X number), including the purchase time, carrier, province, city, and status.
//
// - Before you call this operation, make sure that you have the private number (X number) that you want to query.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 1,000 calls per second. If the limit is exceeded, the API call is throttled. Call this operation properly.
//
// @param request - QuerySecretNoDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySecretNoDetailResponse
func (client *Client) QuerySecretNoDetailWithOptions(request *QuerySecretNoDetailRequest, runtime *dara.RuntimeOptions) (_result *QuerySecretNoDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecretNo) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySecretNoDetail"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySecretNoDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the property information of a private number.
//
// Description:
//
// - You can query information about a private number (X number), including the purchase time, carrier, province, city, and status.
//
// - Before you call this operation, make sure that you have the private number (X number) that you want to query.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 1,000 calls per second. If the limit is exceeded, the API call is throttled. Call this operation properly.
//
// @param request - QuerySecretNoDetailRequest
//
// @return QuerySecretNoDetailResponse
func (client *Client) QuerySecretNoDetail(request *QuerySecretNoDetailRequest) (_result *QuerySecretNoDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySecretNoDetailResponse{}
	_body, _err := client.QuerySecretNoDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the quantity of remaining phone numbers available for online purchase.
//
// Description:
//
// When purchasing a phone number, specify the home location. If no sufficient phone numbers are available for purchase in the home location, the purchase of the phone number fails. Before calling the [BuySecretNo](~~BuySecretNo~~) operation to purchase a phone number, call the [QuerySecretNoRemain](~~QuerySecretNoRemain~~) operation to query the quantity of remaining phone numbers available for online purchase.
//
// @param request - QuerySecretNoRemainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySecretNoRemainResponse
func (client *Client) QuerySecretNoRemainWithOptions(request *QuerySecretNoRemainRequest, runtime *dara.RuntimeOptions) (_result *QuerySecretNoRemainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecretNo) {
		query["SecretNo"] = request.SecretNo
	}

	if !dara.IsNil(request.SpecId) {
		query["SpecId"] = request.SpecId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySecretNoRemain"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySecretNoRemainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quantity of remaining phone numbers available for online purchase.
//
// Description:
//
// When purchasing a phone number, specify the home location. If no sufficient phone numbers are available for purchase in the home location, the purchase of the phone number fails. Before calling the [BuySecretNo](~~BuySecretNo~~) operation to purchase a phone number, call the [QuerySecretNoRemain](~~QuerySecretNoRemain~~) operation to query the quantity of remaining phone numbers available for online purchase.
//
// @param request - QuerySecretNoRemainRequest
//
// @return QuerySecretNoRemainResponse
func (client *Client) QuerySecretNoRemain(request *QuerySecretNoRemainRequest) (_result *QuerySecretNoRemainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySecretNoRemainResponse{}
	_body, _err := client.QuerySecretNoRemainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 对客分页查询指定迁移记录的明细，记录的来源与目标号池须均归属 aliUid。
//
// @param request - QuerySecretTransferDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySecretTransferDetailsResponse
func (client *Client) QuerySecretTransferDetailsWithOptions(request *QuerySecretTransferDetailsRequest, runtime *dara.RuntimeOptions) (_result *QuerySecretTransferDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySecretTransferDetails"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySecretTransferDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对客分页查询指定迁移记录的明细，记录的来源与目标号池须均归属 aliUid。
//
// @param request - QuerySecretTransferDetailsRequest
//
// @return QuerySecretTransferDetailsResponse
func (client *Client) QuerySecretTransferDetails(request *QuerySecretTransferDetailsRequest) (_result *QuerySecretTransferDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySecretTransferDetailsResponse{}
	_body, _err := client.QuerySecretTransferDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the downloadable URL of a call recording by call ID.
//
// @param request - QuerySoundRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySoundRecordResponse
func (client *Client) QuerySoundRecordWithOptions(request *QuerySoundRecordRequest, runtime *dara.RuntimeOptions) (_result *QuerySoundRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CallerParentId) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !dara.IsNil(request.CustomerPoolKey) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReqId) {
		query["ReqId"] = request.ReqId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySoundRecord"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySoundRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the downloadable URL of a call recording by call ID.
//
// @param request - QuerySoundRecordRequest
//
// @return QuerySoundRecordResponse
func (client *Client) QuerySoundRecord(request *QuerySoundRecordRequest) (_result *QuerySoundRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySoundRecordResponse{}
	_body, _err := client.QuerySoundRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries binding IDs.
//
// Description:
//
// You can query binding IDs by phone number X. In the AXB product, multiple bindings may exist for the same phone number X. In this case, multiple binding IDs may be obtained for the same phone number X.
//
// @param request - QuerySubsIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySubsIdResponse
func (client *Client) QuerySubsIdWithOptions(request *QuerySubsIdRequest, runtime *dara.RuntimeOptions) (_result *QuerySubsIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNoX) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySubsId"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySubsIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries binding IDs.
//
// Description:
//
// You can query binding IDs by phone number X. In the AXB product, multiple bindings may exist for the same phone number X. In this case, multiple binding IDs may be obtained for the same phone number X.
//
// @param request - QuerySubsIdRequest
//
// @return QuerySubsIdResponse
func (client *Client) QuerySubsId(request *QuerySubsIdRequest) (_result *QuerySubsIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySubsIdResponse{}
	_body, _err := client.QuerySubsIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the binding details of a phone number. Retrieves information such as the binding relationship status, extension number, recording status, and ASR status.
//
// Description:
//
// Before you call this operation, make sure that you have the private number (X number) to query and that the number has a binding relationship.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 1,000. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// ### PoolKey and ProductType limit
//
// Specify at least one of the PoolKey and ProductType parameters. If neither parameter is specified, an error is returned.
//
// Former Alibaba Dayu users should specify the ProductType parameter, and Alibaba Cloud users should specify the PoolKey parameter.
//
// @param request - QuerySubscriptionDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySubscriptionDetailResponse
func (client *Client) QuerySubscriptionDetailWithOptions(request *QuerySubscriptionDetailRequest, runtime *dara.RuntimeOptions) (_result *QuerySubscriptionDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNoX) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SubsId) {
		query["SubsId"] = request.SubsId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySubscriptionDetail"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySubscriptionDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the binding details of a phone number. Retrieves information such as the binding relationship status, extension number, recording status, and ASR status.
//
// Description:
//
// Before you call this operation, make sure that you have the private number (X number) to query and that the number has a binding relationship.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 1,000. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// ### PoolKey and ProductType limit
//
// Specify at least one of the PoolKey and ProductType parameters. If neither parameter is specified, an error is returned.
//
// Former Alibaba Dayu users should specify the ProductType parameter, and Alibaba Cloud users should specify the PoolKey parameter.
//
// @param request - QuerySubscriptionDetailRequest
//
// @return QuerySubscriptionDetailResponse
func (client *Client) QuerySubscriptionDetail(request *QuerySubscriptionDetailRequest) (_result *QuerySubscriptionDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySubscriptionDetailResponse{}
	_body, _err := client.QuerySubscriptionDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a phone number.
//
// Description:
//
//	  After a phone number is released, it will no longer be charged from the following month.
//
//		- Before you release a phone number, log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) to check whether the phone number is bound to other phone numbers. The phone number can be released only if it is not bound to other phone numbers.
//
// @param request - ReleaseSecretNoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseSecretNoResponse
func (client *Client) ReleaseSecretNoWithOptions(request *ReleaseSecretNoRequest, runtime *dara.RuntimeOptions) (_result *ReleaseSecretNoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecretNo) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseSecretNo"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseSecretNoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a phone number.
//
// Description:
//
//	  After a phone number is released, it will no longer be charged from the following month.
//
//		- Before you release a phone number, log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) to check whether the phone number is bound to other phone numbers. The phone number can be released only if it is not bound to other phone numbers.
//
// @param request - ReleaseSecretNoRequest
//
// @return ReleaseSecretNoResponse
func (client *Client) ReleaseSecretNo(request *ReleaseSecretNoRequest) (_result *ReleaseSecretNoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseSecretNoResponse{}
	_body, _err := client.ReleaseSecretNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Dissociates a specified call binding relationship (A, X, and B). After the call binding relationship is dissociated, Employee B cannot call Customer A through Work Number X.
//
// @param request - UnBindAXBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnBindAXBResponse
func (client *Client) UnBindAXBWithOptions(request *UnBindAXBRequest, runtime *dara.RuntimeOptions) (_result *UnBindAXBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BindId) {
		query["BindId"] = request.BindId
	}

	if !dara.IsNil(request.CallerParentId) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !dara.IsNil(request.CustomerPoolKey) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReqId) {
		query["ReqId"] = request.ReqId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnBindAXB"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnBindAXBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Dissociates a specified call binding relationship (A, X, and B). After the call binding relationship is dissociated, Employee B cannot call Customer A through Work Number X.
//
// @param request - UnBindAXBRequest
//
// @return UnBindAXBResponse
func (client *Client) UnBindAXB(request *UnBindAXBRequest) (_result *UnBindAXBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnBindAXBResponse{}
	_body, _err := client.UnBindAXBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the binding between work number X and employee number B. After the binding is canceled, calls to X are no longer forwarded to B.
//
// @param request - UnBindXBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnBindXBResponse
func (client *Client) UnBindXBWithOptions(request *UnBindXBRequest, runtime *dara.RuntimeOptions) (_result *UnBindXBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthId) {
		query["AuthId"] = request.AuthId
	}

	if !dara.IsNil(request.CallerParentId) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !dara.IsNil(request.CustomerPoolKey) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReqId) {
		query["ReqId"] = request.ReqId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TelX) {
		query["TelX"] = request.TelX
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnBindXB"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnBindXBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the binding between work number X and employee number B. After the binding is canceled, calls to X are no longer forwarded to B.
//
// @param request - UnBindXBRequest
//
// @return UnBindXBResponse
func (client *Client) UnBindXB(request *UnBindXBRequest) (_result *UnBindXBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnBindXBResponse{}
	_body, _err := client.UnBindXBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds a 700 number.
//
// @param request - UnbindSubs700Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindSubs700Response
func (client *Client) UnbindSubs700WithOptions(request *UnbindSubs700Request, runtime *dara.RuntimeOptions) (_result *UnbindSubs700Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IndustrialId) {
		query["IndustrialId"] = request.IndustrialId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SubsId) {
		query["SubsId"] = request.SubsId
	}

	if !dara.IsNil(request.TelX) {
		query["TelX"] = request.TelX
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindSubs700"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindSubs700Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a 700 number.
//
// @param request - UnbindSubs700Request
//
// @return UnbindSubs700Response
func (client *Client) UnbindSubs700(request *UnbindSubs700Request) (_result *UnbindSubs700Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindSubs700Response{}
	_body, _err := client.UnbindSubs700WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds a phone number.
//
// Description:
//
// Before releasing a phone number, you must call the UnbindSubscription operation to unbind the phone number.
//
// @param request - UnbindSubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindSubscriptionResponse
func (client *Client) UnbindSubscriptionWithOptions(request *UnbindSubscriptionRequest, runtime *dara.RuntimeOptions) (_result *UnbindSubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecretNo) {
		query["SecretNo"] = request.SecretNo
	}

	if !dara.IsNil(request.SubsId) {
		query["SubsId"] = request.SubsId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindSubscription"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a phone number.
//
// Description:
//
// Before releasing a phone number, you must call the UnbindSubscription operation to unbind the phone number.
//
// @param request - UnbindSubscriptionRequest
//
// @return UnbindSubscriptionResponse
func (client *Client) UnbindSubscription(request *UnbindSubscriptionRequest) (_result *UnbindSubscriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindSubscriptionResponse{}
	_body, _err := client.UnbindSubscriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unlocks a phone number.
//
// Description:
//
// After a phone number is unlocked, you can reselect the unlocked phone number when you call an operation to create a binding.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UnlockSecretNoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockSecretNoResponse
func (client *Client) UnlockSecretNoWithOptions(request *UnlockSecretNoRequest, runtime *dara.RuntimeOptions) (_result *UnlockSecretNoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecretNo) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnlockSecretNo"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnlockSecretNoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unlocks a phone number.
//
// Description:
//
// After a phone number is unlocked, you can reselect the unlocked phone number when you call an operation to create a binding.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UnlockSecretNoRequest
//
// @return UnlockSecretNoResponse
func (client *Client) UnlockSecretNo(request *UnlockSecretNoRequest) (_result *UnlockSecretNoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnlockSecretNoResponse{}
	_body, _err := client.UnlockSecretNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the binding information of an AXB privacy number for landlines.
//
// Description:
//
// When you update an AXB privacy number for landlines, you cannot modify Number A and Number B at the same time. For example, if the original binding relationship is A1 and B1, you can change it to A1 and B2, or A2 and B1.
//
// @param tmpReq - UpdateAxbBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAxbBindFixedLineResponse
func (client *Client) UpdateAxbBindFixedLineWithOptions(tmpReq *UpdateAxbBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *UpdateAxbBindFixedLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAxbBindFixedLineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("Extra"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Anucode) {
		query["Anucode"] = request.Anucode
	}

	if !dara.IsNil(request.Anucodecalled) {
		query["Anucodecalled"] = request.Anucodecalled
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.ExtraShrink) {
		query["Extra"] = request.ExtraShrink
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SubId) {
		query["SubId"] = request.SubId
	}

	if !dara.IsNil(request.Subts) {
		query["Subts"] = request.Subts
	}

	if !dara.IsNil(request.TAnucodeConnect) {
		query["TAnucodeConnect"] = request.TAnucodeConnect
	}

	if !dara.IsNil(request.TelA) {
		query["TelA"] = request.TelA
	}

	if !dara.IsNil(request.TelB) {
		query["TelB"] = request.TelB
	}

	if !dara.IsNil(request.Ts) {
		query["Ts"] = request.Ts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAxbBindFixedLine"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAxbBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the binding information of an AXB privacy number for landlines.
//
// Description:
//
// When you update an AXB privacy number for landlines, you cannot modify Number A and Number B at the same time. For example, if the original binding relationship is A1 and B1, you can change it to A1 and B2, or A2 and B1.
//
// @param request - UpdateAxbBindFixedLineRequest
//
// @return UpdateAxbBindFixedLineResponse
func (client *Client) UpdateAxbBindFixedLine(request *UpdateAxbBindFixedLineRequest) (_result *UpdateAxbBindFixedLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAxbBindFixedLineResponse{}
	_body, _err := client.UpdateAxbBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the binding information of a landline AXN privacy number.
//
// @param tmpReq - UpdateAxnBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAxnBindFixedLineResponse
func (client *Client) UpdateAxnBindFixedLineWithOptions(tmpReq *UpdateAxnBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *UpdateAxnBindFixedLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAxnBindFixedLineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("Extra"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Anucode) {
		query["Anucode"] = request.Anucode
	}

	if !dara.IsNil(request.Anucodecalled) {
		query["Anucodecalled"] = request.Anucodecalled
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.ExtraShrink) {
		query["Extra"] = request.ExtraShrink
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SubId) {
		query["SubId"] = request.SubId
	}

	if !dara.IsNil(request.Subts) {
		query["Subts"] = request.Subts
	}

	if !dara.IsNil(request.TAnucodeConnect) {
		query["TAnucodeConnect"] = request.TAnucodeConnect
	}

	if !dara.IsNil(request.TelA) {
		query["TelA"] = request.TelA
	}

	if !dara.IsNil(request.TelB) {
		query["TelB"] = request.TelB
	}

	if !dara.IsNil(request.Ts) {
		query["Ts"] = request.Ts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAxnBindFixedLine"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAxnBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the binding information of a landline AXN privacy number.
//
// @param request - UpdateAxnBindFixedLineRequest
//
// @return UpdateAxnBindFixedLineResponse
func (client *Client) UpdateAxnBindFixedLine(request *UpdateAxnBindFixedLineRequest) (_result *UpdateAxnBindFixedLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAxnBindFixedLineResponse{}
	_body, _err := client.UpdateAxnBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the binding relationship of a landline AXN extension number.
//
// @param tmpReq - UpdateAxnExtensionBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAxnExtensionBindFixedLineResponse
func (client *Client) UpdateAxnExtensionBindFixedLineWithOptions(tmpReq *UpdateAxnExtensionBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *UpdateAxnExtensionBindFixedLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAxnExtensionBindFixedLineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Extraaxx) {
		request.ExtraaxxShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extraaxx, dara.String("Extraaxx"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Anucode) {
		query["Anucode"] = request.Anucode
	}

	if !dara.IsNil(request.Anucodecalled) {
		query["Anucodecalled"] = request.Anucodecalled
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.ExtraaxxShrink) {
		query["Extraaxx"] = request.ExtraaxxShrink
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SubId) {
		query["SubId"] = request.SubId
	}

	if !dara.IsNil(request.Subts) {
		query["Subts"] = request.Subts
	}

	if !dara.IsNil(request.TAnucodeConnect) {
		query["TAnucodeConnect"] = request.TAnucodeConnect
	}

	if !dara.IsNil(request.TelA) {
		query["TelA"] = request.TelA
	}

	if !dara.IsNil(request.Ts) {
		query["Ts"] = request.Ts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAxnExtensionBindFixedLine"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAxnExtensionBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the binding relationship of a landline AXN extension number.
//
// @param request - UpdateAxnExtensionBindFixedLineRequest
//
// @return UpdateAxnExtensionBindFixedLineResponse
func (client *Client) UpdateAxnExtensionBindFixedLine(request *UpdateAxnExtensionBindFixedLineRequest) (_result *UpdateAxnExtensionBindFixedLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAxnExtensionBindFixedLineResponse{}
	_body, _err := client.UpdateAxnExtensionBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the binding relationship of a 700 number.
//
// @param request - UpdateSubs700Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSubs700Response
func (client *Client) UpdateSubs700WithOptions(request *UpdateSubs700Request, runtime *dara.RuntimeOptions) (_result *UpdateSubs700Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AsrModelId) {
		query["AsrModelId"] = request.AsrModelId
	}

	if !dara.IsNil(request.Audio) {
		query["Audio"] = request.Audio
	}

	if !dara.IsNil(request.CallRestrict) {
		query["CallRestrict"] = request.CallRestrict
	}

	if !dara.IsNil(request.DefaultA) {
		query["DefaultA"] = request.DefaultA
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IndustrialId) {
		query["IndustrialId"] = request.IndustrialId
	}

	if !dara.IsNil(request.NeedAsr) {
		query["NeedAsr"] = request.NeedAsr
	}

	if !dara.IsNil(request.NeedRecord) {
		query["NeedRecord"] = request.NeedRecord
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SubsId) {
		query["SubsId"] = request.SubsId
	}

	if !dara.IsNil(request.TelA) {
		query["TelA"] = request.TelA
	}

	if !dara.IsNil(request.TelB) {
		query["TelB"] = request.TelB
	}

	if !dara.IsNil(request.TelX) {
		query["TelX"] = request.TelX
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSubs700"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSubs700Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the binding relationship of a 700 number.
//
// @param request - UpdateSubs700Request
//
// @return UpdateSubs700Response
func (client *Client) UpdateSubs700(request *UpdateSubs700Request) (_result *UpdateSubs700Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSubs700Response{}
	_body, _err := client.UpdateSubs700WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a phone number binding.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 10,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UpdateSubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSubscriptionResponse
func (client *Client) UpdateSubscriptionWithOptions(request *UpdateSubscriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateSubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ASRModelId) {
		query["ASRModelId"] = request.ASRModelId
	}

	if !dara.IsNil(request.ASRStatus) {
		query["ASRStatus"] = request.ASRStatus
	}

	if !dara.IsNil(request.CallDisplayType) {
		query["CallDisplayType"] = request.CallDisplayType
	}

	if !dara.IsNil(request.CallRestrict) {
		query["CallRestrict"] = request.CallRestrict
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IsRecordingEnabled) {
		query["IsRecordingEnabled"] = request.IsRecordingEnabled
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNoA) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !dara.IsNil(request.PhoneNoB) {
		query["PhoneNoB"] = request.PhoneNoB
	}

	if !dara.IsNil(request.PhoneNoX) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !dara.IsNil(request.PoolKey) {
		query["PoolKey"] = request.PoolKey
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RingConfig) {
		query["RingConfig"] = request.RingConfig
	}

	if !dara.IsNil(request.SubsId) {
		query["SubsId"] = request.SubsId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSubscription"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a phone number binding.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 10,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UpdateSubscriptionRequest
//
// @return UpdateSubscriptionResponse
func (client *Client) UpdateSubscription(request *UpdateSubscriptionRequest) (_result *UpdateSubscriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSubscriptionResponse{}
	_body, _err := client.UpdateSubscriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
