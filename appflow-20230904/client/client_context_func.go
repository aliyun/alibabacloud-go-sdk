// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// # Generate Login Session Token
//
// @param request - GenerateUserSessionTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateUserSessionTokenResponse
func (client *Client) GenerateUserSessionTokenWithContext(ctx context.Context, request *GenerateUserSessionTokenRequest, runtime *dara.RuntimeOptions) (_result *GenerateUserSessionTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChatbotId) {
		query["ChatbotId"] = request.ChatbotId
	}

	if !dara.IsNil(request.ExpireSecond) {
		query["ExpireSecond"] = request.ExpireSecond
	}

	if !dara.IsNil(request.ExtraInfo) {
		query["ExtraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.IntegrateId) {
		query["IntegrateId"] = request.IntegrateId
	}

	if !dara.IsNil(request.UserAvatar) {
		query["UserAvatar"] = request.UserAvatar
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateUserSessionToken"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateUserSessionTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
