// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 全员营销
//
// @param request - GetEveryOneSellsFormListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEveryOneSellsFormListResponse
func (client *Client) GetEveryOneSellsFormListWithContext(ctx context.Context, request *GetEveryOneSellsFormListRequest, runtime *dara.RuntimeOptions) (_result *GetEveryOneSellsFormListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEveryOneSellsFormList"),
		Version:     dara.String("2025-07-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("array"),
	}
	_result = &GetEveryOneSellsFormListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送钉钉消息
//
// @param tmpReq - PushEveryOneSellMsgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushEveryOneSellMsgResponse
func (client *Client) PushEveryOneSellMsgWithContext(ctx context.Context, tmpReq *PushEveryOneSellMsgRequest, runtime *dara.RuntimeOptions) (_result *PushEveryOneSellMsgResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &PushEveryOneSellMsgShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DingIdList) {
		request.DingIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DingIdList, dara.String("DingIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DingIdListShrink) {
		body["DingIdList"] = request.DingIdListShrink
	}

	if !dara.IsNil(request.PushMsg) {
		body["PushMsg"] = request.PushMsg
	}

	if !dara.IsNil(request.PushType) {
		body["PushType"] = request.PushType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushEveryOneSellMsg"),
		Version:     dara.String("2025-07-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &PushEveryOneSellMsgResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
