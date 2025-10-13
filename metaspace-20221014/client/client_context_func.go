// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 协同者发起流协同请求并获取协同ticket的API接口。
//
// Description:
//
// ## 请求说明
//
// - 该API为内部使用，主要用于协同者发起流协同。
//
// - `DependOnMainStream`参数指定了是否依赖主流的状态来建立或断开协同流。
//
// - 当`InitiatorType`设置为强制协同类型时（如`ADMIN_INITIATE_FORCE`或`COORDINATOR_INITIATE_FORCE`），将返回一个`CoordinateTicket`。
//
// - 协同资源列表`CoordinationResourceCandidates`中必须包含至少一项资源信息，并且所有提供的资源ID、类型和地区等信息需准确无误。
//
// - 确保`OwnerAliUid`与协同者的租户ID相同，否则可能无法成功发起协同请求。
//
// - 对于AD用户，请务必填写`AdDomainName`字段。
//
// @param request - ApplyCoordinationForCoordinatorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyCoordinationForCoordinatorResponse
func (client *Client) ApplyCoordinationForCoordinatorWithContext(ctx context.Context, request *ApplyCoordinationForCoordinatorRequest, runtime *dara.RuntimeOptions) (_result *ApplyCoordinationForCoordinatorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.CoordinationResourceCandidates) {
		bodyFlat["CoordinationResourceCandidates"] = request.CoordinationResourceCandidates
	}

	if !dara.IsNil(request.InitiatorType) {
		body["InitiatorType"] = request.InitiatorType
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyCoordinationForCoordinator"),
		Version:     dara.String("2022-10-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyCoordinationForCoordinatorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消协同请求
//
// @param request - CancelCoordinationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelCoordinationResponse
func (client *Client) CancelCoordinationWithContext(ctx context.Context, request *CancelCoordinationRequest, runtime *dara.RuntimeOptions) (_result *CancelCoordinationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.CoIds) {
		bodyFlat["CoIds"] = request.CoIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelCoordination"),
		Version:     dara.String("2022-10-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelCoordinationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取协同流连接ticket
//
// @param request - GetCoordinationTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCoordinationTicketResponse
func (client *Client) GetCoordinationTicketWithContext(ctx context.Context, request *GetCoordinationTicketRequest, runtime *dara.RuntimeOptions) (_result *GetCoordinationTicketResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CoId) {
		body["CoId"] = request.CoId
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCoordinationTicket"),
		Version:     dara.String("2022-10-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCoordinationTicketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
