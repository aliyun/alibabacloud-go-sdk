// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 修改Workbench终端配置
//
// Description:
//
// 修改Workbench终端配置
//
// @param tmpReq - ModifyTerminalSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTerminalSettingsResponse
func (client *Client) ModifyTerminalSettingsWithContext(ctx context.Context, tmpReq *ModifyTerminalSettingsRequest, runtime *dara.RuntimeOptions) (_result *ModifyTerminalSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyTerminalSettingsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PasswordlessLoginConfig) {
		request.PasswordlessLoginConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PasswordlessLoginConfig, dara.String("PasswordlessLoginConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PasswordlessLoginConfigShrink) {
		query["PasswordlessLoginConfig"] = request.PasswordlessLoginConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTerminalSettings"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTerminalSettingsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
