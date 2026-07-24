// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 发起实时外呼
//
// @param request - CreateCallOutboundInstantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCallOutboundInstantResponse
func (client *Client) CreateCallOutboundInstantWithContext(ctx context.Context, request *CreateCallOutboundInstantRequest, runtime *dara.RuntimeOptions) (_result *CreateCallOutboundInstantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CalledNumber) {
		body["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CustomerName) {
		body["CustomerName"] = request.CustomerName
	}

	if !dara.IsNil(request.EncryptCall) {
		body["EncryptCall"] = request.EncryptCall
	}

	if !dara.IsNil(request.PromptVariables) {
		body["PromptVariables"] = request.PromptVariables
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCallOutboundInstant"),
		Version:     dara.String("2025-11-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCallOutboundInstantResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询当前任务的并发数
//
// @param request - QueryTaskConcurrencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskConcurrencyResponse
func (client *Client) QueryTaskConcurrencyWithContext(ctx context.Context, request *QueryTaskConcurrencyRequest, runtime *dara.RuntimeOptions) (_result *QueryTaskConcurrencyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationCode) {
		body["ApplicationCode"] = request.ApplicationCode
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTaskConcurrency"),
		Version:     dara.String("2025-11-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTaskConcurrencyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 外呼任务通话列表查询
//
// @param tmpReq - ReadOutboundTaskCallListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadOutboundTaskCallListResponse
func (client *Client) ReadOutboundTaskCallListWithContext(ctx context.Context, tmpReq *ReadOutboundTaskCallListRequest, runtime *dara.RuntimeOptions) (_result *ReadOutboundTaskCallListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReadOutboundTaskCallListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DisplayStatusList) {
		request.DisplayStatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DisplayStatusList, dara.String("DisplayStatusList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LabelTags) {
		request.LabelTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelTags, dara.String("LabelTags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CallEndTimeBegin) {
		body["CallEndTimeBegin"] = request.CallEndTimeBegin
	}

	if !dara.IsNil(request.CallEndTimeEnd) {
		body["CallEndTimeEnd"] = request.CallEndTimeEnd
	}

	if !dara.IsNil(request.CallStartTimeBegin) {
		body["CallStartTimeBegin"] = request.CallStartTimeBegin
	}

	if !dara.IsNil(request.CallStartTimeEnd) {
		body["CallStartTimeEnd"] = request.CallStartTimeEnd
	}

	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.CustomerNameOrPhone) {
		body["CustomerNameOrPhone"] = request.CustomerNameOrPhone
	}

	if !dara.IsNil(request.DisplayStatusListShrink) {
		body["DisplayStatusList"] = request.DisplayStatusListShrink
	}

	if !dara.IsNil(request.LabelTagsShrink) {
		body["LabelTags"] = request.LabelTagsShrink
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadOutboundTaskCallList"),
		Version:     dara.String("2025-11-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadOutboundTaskCallListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
