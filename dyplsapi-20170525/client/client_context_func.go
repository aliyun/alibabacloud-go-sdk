// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

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
func (client *Client) AddAxnTrackNoWithContext(ctx context.Context, request *AddAxnTrackNoRequest, runtime *dara.RuntimeOptions) (_result *AddAxnTrackNoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a blacklist.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddSecretBlacklistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSecretBlacklistResponse
func (client *Client) AddSecretBlacklistWithContext(ctx context.Context, request *AddSecretBlacklistRequest, runtime *dara.RuntimeOptions) (_result *AddSecretBlacklistResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用本接口向工作号平台请求为员工B的工作号X建立呼叫绑定（B，X，A），允许B通过X呼叫客户A
//
// @param request - BindAXBCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAXBCallResponse
func (client *Client) BindAXBCallWithContext(ctx context.Context, request *BindAXBCallRequest, runtime *dara.RuntimeOptions) (_result *BindAXBCallResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxbResponse
func (client *Client) BindAxbWithContext(ctx context.Context, request *BindAxbRequest, runtime *dara.RuntimeOptions) (_result *BindAxbResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建700绑定关系
//
// @param request - BindAxb700Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxb700Response
func (client *Client) BindAxb700WithContext(ctx context.Context, request *BindAxb700Request, runtime *dara.RuntimeOptions) (_result *BindAxb700Response, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 固话AxB绑定
//
// @param tmpReq - BindAxbFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxbFixedLineResponse
func (client *Client) BindAxbFixedLineWithContext(ctx context.Context, tmpReq *BindAxbFixedLineRequest, runtime *dara.RuntimeOptions) (_result *BindAxbFixedLineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxgResponse
func (client *Client) BindAxgWithContext(ctx context.Context, request *BindAxgRequest, runtime *dara.RuntimeOptions) (_result *BindAxgResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an AXN binding.
//
// Description:
//
// >  An AXN private number is a dedicated private number assigned to phone number A. When an N-side number is used to call phone number X, the call is forwarded to phone number A.
//
// @param request - BindAxnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxnResponse
func (client *Client) BindAxnWithContext(ctx context.Context, request *BindAxnRequest, runtime *dara.RuntimeOptions) (_result *BindAxnResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an AXN extension binding.
//
// Description:
//
// Before you add an AXN extension binding, confirm phone number A and phone number N in the business scenario. Phone number A belongs to a customer, and phone number X is the private number assigned to the customer. When any other phone number is used to call phone number X and the extension, the call is transferred to phone number A. When phone number A is used to call phone number X, the call is transferred to the default phone number B that is specified during the phone number binding.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - BindAxnExtensionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxnExtensionResponse
func (client *Client) BindAxnExtensionWithContext(ctx context.Context, request *BindAxnExtensionRequest, runtime *dara.RuntimeOptions) (_result *BindAxnExtensionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AXN分机号-号码绑定
//
// @param tmpReq - BindAxnExtensionFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxnExtensionFixedLineResponse
func (client *Client) BindAxnExtensionFixedLineWithContext(ctx context.Context, tmpReq *BindAxnExtensionFixedLineRequest, runtime *dara.RuntimeOptions) (_result *BindAxnExtensionFixedLineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AXN模式绑定，分配X号码
//
// @param tmpReq - BindAxnFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxnFixedLineResponse
func (client *Client) BindAxnFixedLineWithContext(ctx context.Context, tmpReq *BindAxnFixedLineRequest, runtime *dara.RuntimeOptions) (_result *BindAxnFixedLineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - BindBatchAxgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindBatchAxgResponse
func (client *Client) BindBatchAxgWithContext(ctx context.Context, tmpReq *BindBatchAxgRequest, runtime *dara.RuntimeOptions) (_result *BindBatchAxgResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建700Gxb绑定关系
//
// @param request - BindGxb700Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindGxb700Response
func (client *Client) BindGxb700WithContext(ctx context.Context, request *BindGxb700Request, runtime *dara.RuntimeOptions) (_result *BindGxb700Response, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 平台指定工作号X 和员工号B建立关联，完成X 实名认证，绑定生效后，所有X 的呼叫都会转接到B
//
// @param request - BindXBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindXBResponse
func (client *Client) BindXBWithContext(ctx context.Context, request *BindXBRequest, runtime *dara.RuntimeOptions) (_result *BindXBResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BuySecretNoResponse
func (client *Client) BuySecretNoWithContext(ctx context.Context, request *BuySecretNoRequest, runtime *dara.RuntimeOptions) (_result *BuySecretNoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelPickUpWaybillResponse
func (client *Client) CancelPickUpWaybillWithContext(ctx context.Context, request *CancelPickUpWaybillRequest, runtime *dara.RuntimeOptions) (_result *CancelPickUpWaybillResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置X号码，单独对工作号的话音呼叫、企业名片等通信功能进行配置操作
//
// @param tmpReq - ConfigXRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigXResponse
func (client *Client) ConfigXWithContext(ctx context.Context, tmpReq *ConfigXRequest, runtime *dara.RuntimeOptions) (_result *ConfigXResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAxgGroupResponse
func (client *Client) CreateAxgGroupWithContext(ctx context.Context, request *CreateAxgGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateAxgGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过API收集小号a号码固话
//
// @param request - CreateFixedNoAReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFixedNoAReportResponse
func (client *Client) CreateFixedNoAReportWithContext(ctx context.Context, request *CreateFixedNoAReportRequest, runtime *dara.RuntimeOptions) (_result *CreateFixedNoAReportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过API收集小号a号码手机号
//
// @param request - CreatePhoneNoAReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePhoneNoAReportResponse
func (client *Client) CreatePhoneNoAReportWithContext(ctx context.Context, request *CreatePhoneNoAReportRequest, runtime *dara.RuntimeOptions) (_result *CreatePhoneNoAReportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a door-to-door delivery order.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - CreatePickUpWaybillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePickUpWaybillResponse
func (client *Client) CreatePickUpWaybillWithContext(ctx context.Context, tmpReq *CreatePickUpWaybillRequest, runtime *dara.RuntimeOptions) (_result *CreatePickUpWaybillResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a pickup order.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - CreatePickUpWaybillPreQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePickUpWaybillPreQueryResponse
func (client *Client) CreatePickUpWaybillPreQueryWithContext(ctx context.Context, tmpReq *CreatePickUpWaybillPreQueryRequest, runtime *dara.RuntimeOptions) (_result *CreatePickUpWaybillPreQueryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// B向A 发短信，客户端获取“短信标签”，尾部添加“标签”。通过“标签”解析被叫A，发短信到A。
//
// @param request - CreateSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmsSignResponse
func (client *Client) CreateSmsSignWithContext(ctx context.Context, request *CreateSmsSignRequest, runtime *dara.RuntimeOptions) (_result *CreateSmsSignResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 固话AxB解绑
//
// @param request - DeleteAxbBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAxbBindFixedLineResponse
func (client *Client) DeleteAxbBindFixedLineWithContext(ctx context.Context, request *DeleteAxbBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *DeleteAxbBindFixedLineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAxgGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAxgGroupResponse
func (client *Client) DeleteAxgGroupWithContext(ctx context.Context, request *DeleteAxgGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteAxgGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑已有Axn绑定
//
// @param request - DeleteAxnBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAxnBindFixedLineResponse
func (client *Client) DeleteAxnBindFixedLineWithContext(ctx context.Context, request *DeleteAxnBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *DeleteAxnBindFixedLineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑已有AXN分机号绑定
//
// @param request - DeleteAxnExtensionBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAxnExtensionBindFixedLineResponse
func (client *Client) DeleteAxnExtensionBindFixedLineWithContext(ctx context.Context, request *DeleteAxnExtensionBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *DeleteAxnExtensionBindFixedLineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # A号码报备数据删除
//
// @param request - DeleteSecretAPhoneNoToCustRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecretAPhoneNoToCustResponse
func (client *Client) DeleteSecretAPhoneNoToCustWithContext(ctx context.Context, request *DeleteSecretAPhoneNoToCustRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecretAPhoneNoToCustResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a blacklist.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteSecretBlacklistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecretBlacklistResponse
func (client *Client) DeleteSecretBlacklistWithContext(ctx context.Context, request *DeleteSecretBlacklistRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecretBlacklistResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 隐私号上传文件，获取 OSS 信息
//
// @param request - GetDyplsOSSInfoForUploadFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDyplsOSSInfoForUploadFileResponse
func (client *Client) GetDyplsOSSInfoForUploadFileWithContext(ctx context.Context, request *GetDyplsOSSInfoForUploadFileRequest, runtime *dara.RuntimeOptions) (_result *GetDyplsOSSInfoForUploadFileResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecretAsrDetailResponse
func (client *Client) GetSecretAsrDetailWithContext(ctx context.Context, request *GetSecretAsrDetailRequest, runtime *dara.RuntimeOptions) (_result *GetSecretAsrDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTotalPublicUrlResponse
func (client *Client) GetTotalPublicUrlWithContext(ctx context.Context, request *GetTotalPublicUrlRequest, runtime *dara.RuntimeOptions) (_result *GetTotalPublicUrlResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取X号码配置信息
//
// @param request - GetXConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetXConfigResponse
func (client *Client) GetXConfigWithContext(ctx context.Context, request *GetXConfigRequest, runtime *dara.RuntimeOptions) (_result *GetXConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取X号码默认配置信息
//
// @param request - GetXDefaultConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetXDefaultConfigResponse
func (client *Client) GetXDefaultConfigWithContext(ctx context.Context, request *GetXDefaultConfigRequest, runtime *dara.RuntimeOptions) (_result *GetXDefaultConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询客户名下X号码列表
//
// @param request - ListXTelephonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListXTelephonesResponse
func (client *Client) ListXTelephonesWithContext(ctx context.Context, request *ListXTelephonesRequest, runtime *dara.RuntimeOptions) (_result *ListXTelephonesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LockSecretNoResponse
func (client *Client) LockSecretNoWithContext(ctx context.Context, request *LockSecretNoRequest, runtime *dara.RuntimeOptions) (_result *LockSecretNoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateAxgGroupResponse
func (client *Client) OperateAxgGroupWithContext(ctx context.Context, request *OperateAxgGroupRequest, runtime *dara.RuntimeOptions) (_result *OperateAxgGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateBlackNoResponse
func (client *Client) OperateBlackNoWithContext(ctx context.Context, request *OperateBlackNoRequest, runtime *dara.RuntimeOptions) (_result *OperateBlackNoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 固话AxB查询
//
// @param request - QueryAxbBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAxbBindFixedLineResponse
func (client *Client) QueryAxbBindFixedLineWithContext(ctx context.Context, request *QueryAxbBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *QueryAxbBindFixedLineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Axn绑定关系
//
// @param request - QueryAxnBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAxnBindFixedLineResponse
func (client *Client) QueryAxnBindFixedLineWithContext(ctx context.Context, request *QueryAxnBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *QueryAxnBindFixedLineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询AXN分机号绑定关系
//
// @param request - QueryAxnExtensionBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAxnExtensionBindFixedLineResponse
func (client *Client) QueryAxnExtensionBindFixedLineWithContext(ctx context.Context, request *QueryAxnExtensionBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *QueryAxnExtensionBindFixedLineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPhoneNoAByTrackNoResponse
func (client *Client) QueryPhoneNoAByTrackNoWithContext(ctx context.Context, request *QueryPhoneNoAByTrackNoRequest, runtime *dara.RuntimeOptions) (_result *QueryPhoneNoAByTrackNoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecordFileDownloadUrlResponse
func (client *Client) QueryRecordFileDownloadUrlWithContext(ctx context.Context, request *QueryRecordFileDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordFileDownloadUrlResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # A号码报备状态查询
//
// @param request - QuerySecretAPhoneNoToCustRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySecretAPhoneNoToCustResponse
func (client *Client) QuerySecretAPhoneNoToCustWithContext(ctx context.Context, request *QuerySecretAPhoneNoToCustRequest, runtime *dara.RuntimeOptions) (_result *QuerySecretAPhoneNoToCustResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attributes of a private number.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySecretNoDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySecretNoDetailResponse
func (client *Client) QuerySecretNoDetailWithContext(ctx context.Context, request *QuerySecretNoDetailRequest, runtime *dara.RuntimeOptions) (_result *QuerySecretNoDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySecretNoRemainResponse
func (client *Client) QuerySecretNoRemainWithContext(ctx context.Context, request *QuerySecretNoRemainRequest, runtime *dara.RuntimeOptions) (_result *QuerySecretNoRemainResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询通话录音链接
//
// @param request - QuerySoundRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySoundRecordResponse
func (client *Client) QuerySoundRecordWithContext(ctx context.Context, request *QuerySoundRecordRequest, runtime *dara.RuntimeOptions) (_result *QuerySoundRecordResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySubsIdResponse
func (client *Client) QuerySubsIdWithContext(ctx context.Context, request *QuerySubsIdRequest, runtime *dara.RuntimeOptions) (_result *QuerySubsIdResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a phone number binding.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// ### [](#poolkeyproducttype)Limits on PoolKey and ProductType
//
// You must specify either PoolKey or ProductType. If both parameters are not specified, an error is reported when you call the QuerySubscriptionDetail operation. We recommend that you specify the ProductType parameter for the original key accounts of Alibaba Cloud and the PoolKey parameter for Alibaba Cloud users.
//
// @param request - QuerySubscriptionDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySubscriptionDetailResponse
func (client *Client) QuerySubscriptionDetailWithContext(ctx context.Context, request *QuerySubscriptionDetailRequest, runtime *dara.RuntimeOptions) (_result *QuerySubscriptionDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseSecretNoResponse
func (client *Client) ReleaseSecretNoWithContext(ctx context.Context, request *ReleaseSecretNoRequest, runtime *dara.RuntimeOptions) (_result *ReleaseSecretNoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解除指定的呼叫绑定关系（A，X，B），解决呼叫绑定关系后，员工B不能通过工作号X呼叫到客户A。
//
// @param request - UnBindAXBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnBindAXBResponse
func (client *Client) UnBindAXBWithContext(ctx context.Context, request *UnBindAXBRequest, runtime *dara.RuntimeOptions) (_result *UnBindAXBResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用本接口可取消工作号X与员工号码B的绑定。绑定解除后，对X的呼叫都不会转接给B。
//
// @param request - UnBindXBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnBindXBResponse
func (client *Client) UnBindXBWithContext(ctx context.Context, request *UnBindXBRequest, runtime *dara.RuntimeOptions) (_result *UnBindXBResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解除700绑定关系
//
// @param request - UnbindSubs700Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindSubs700Response
func (client *Client) UnbindSubs700WithContext(ctx context.Context, request *UnbindSubs700Request, runtime *dara.RuntimeOptions) (_result *UnbindSubs700Response, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindSubscriptionResponse
func (client *Client) UnbindSubscriptionWithContext(ctx context.Context, request *UnbindSubscriptionRequest, runtime *dara.RuntimeOptions) (_result *UnbindSubscriptionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockSecretNoResponse
func (client *Client) UnlockSecretNoWithContext(ctx context.Context, request *UnlockSecretNoRequest, runtime *dara.RuntimeOptions) (_result *UnlockSecretNoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 固话AxB绑定更新
//
// @param tmpReq - UpdateAxbBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAxbBindFixedLineResponse
func (client *Client) UpdateAxbBindFixedLineWithContext(ctx context.Context, tmpReq *UpdateAxbBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *UpdateAxbBindFixedLineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Axn绑定关系
//
// @param tmpReq - UpdateAxnBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAxnBindFixedLineResponse
func (client *Client) UpdateAxnBindFixedLineWithContext(ctx context.Context, tmpReq *UpdateAxnBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *UpdateAxnBindFixedLineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新AXN分机号绑定关系
//
// @param tmpReq - UpdateAxnExtensionBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAxnExtensionBindFixedLineResponse
func (client *Client) UpdateAxnExtensionBindFixedLineWithContext(ctx context.Context, tmpReq *UpdateAxnExtensionBindFixedLineRequest, runtime *dara.RuntimeOptions) (_result *UpdateAxnExtensionBindFixedLineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新700绑定关系
//
// @param request - UpdateSubs700Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSubs700Response
func (client *Client) UpdateSubs700WithContext(ctx context.Context, request *UpdateSubs700Request, runtime *dara.RuntimeOptions) (_result *UpdateSubs700Response, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSubscriptionResponse
func (client *Client) UpdateSubscriptionWithContext(ctx context.Context, request *UpdateSubscriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateSubscriptionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
