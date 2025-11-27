// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// /api/v1/sessions
//
// @param request - SearchInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchInfoResponse
func (client *Client) SearchInfoWithContext(ctx context.Context, request *SearchInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.KnowledgeBaseFilters) {
		body["KnowledgeBaseFilters"] = request.KnowledgeBaseFilters
	}

	if !dara.IsNil(request.WebFilters) {
		body["WebFilters"] = request.WebFilters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchInfo"),
		Version:     dara.String("2025-07-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/searches"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
