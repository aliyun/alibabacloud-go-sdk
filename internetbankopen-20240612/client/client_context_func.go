// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 数据产品统一调用入口
//
// @param request - InvokeDataProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeDataProductResponse
func (client *Client) InvokeDataProductWithContext(ctx context.Context, request *InvokeDataProductRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InvokeDataProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ProductCode) {
		query["productCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ServiceParameters) {
		query["serviceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeDataProduct"),
		Version:     dara.String("2024-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/api/v1/rpc"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InvokeDataProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
