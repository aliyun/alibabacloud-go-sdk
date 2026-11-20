// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Initiates a conversation with the Qwen medical agent.
//
// @param request - MedicalAgentChatRequest
//
// @param headers - MedicalAgentChatHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MedicalAgentChatResponse
func (client *Client) MedicalAgentChatWithSSECtx(ctx context.Context, request *MedicalAgentChatRequest, headers *MedicalAgentChatHeaders, runtime *dara.RuntimeOptions, _yield chan *MedicalAgentChatResponse, _yieldErr chan error) {
	defer close(_yield)
	client.medicalAgentChatWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// Initiates a conversation with the Qwen medical agent.
//
// @param request - MedicalAgentChatRequest
//
// @param headers - MedicalAgentChatHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MedicalAgentChatResponse
func (client *Client) MedicalAgentChatWithContext(ctx context.Context, request *MedicalAgentChatRequest, headers *MedicalAgentChatHeaders, runtime *dara.RuntimeOptions) (_result *MedicalAgentChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Communication) {
		body["communication"] = request.Communication
	}

	if !dara.IsNil(request.Data) {
		body["data"] = request.Data
	}

	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XGWAccess) {
		realHeaders["X-GW-Access"] = dara.String(dara.ToString(dara.StringValue(headers.XGWAccess)))
	}

	if !dara.IsNil(headers.XGWNonce) {
		realHeaders["X-GW-Nonce"] = dara.String(dara.ToString(dara.StringValue(headers.XGWNonce)))
	}

	if !dara.IsNil(headers.XGWProtocol) {
		realHeaders["X-GW-Protocol"] = dara.String(dara.ToString(dara.StringValue(headers.XGWProtocol)))
	}

	if !dara.IsNil(headers.XGWSignature) {
		realHeaders["X-GW-Signature"] = dara.String(dara.ToString(dara.StringValue(headers.XGWSignature)))
	}

	if !dara.IsNil(headers.XGWTimestamp) {
		realHeaders["X-GW-Timestamp"] = dara.String(dara.ToString(dara.StringValue(headers.XGWTimestamp)))
	}

	if !dara.IsNil(headers.XRequestId) {
		realHeaders["X-Request-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XRequestId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MedicalAgentChat"),
		Version:     dara.String("2026-11-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/api/medicalQa"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MedicalAgentChatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs inference using the external Qwen medical foundation model.
//
// @param request - MedicalModelInferenceRequest
//
// @param headers - MedicalModelInferenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MedicalModelInferenceResponse
func (client *Client) MedicalModelInferenceWithSSECtx(ctx context.Context, request *MedicalModelInferenceRequest, headers *MedicalModelInferenceHeaders, runtime *dara.RuntimeOptions, _yield chan *MedicalModelInferenceResponse, _yieldErr chan error) {
	defer close(_yield)
	client.medicalModelInferenceWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// Performs inference using the external Qwen medical foundation model.
//
// @param request - MedicalModelInferenceRequest
//
// @param headers - MedicalModelInferenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MedicalModelInferenceResponse
func (client *Client) MedicalModelInferenceWithContext(ctx context.Context, request *MedicalModelInferenceRequest, headers *MedicalModelInferenceHeaders, runtime *dara.RuntimeOptions) (_result *MedicalModelInferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Communication) {
		body["communication"] = request.Communication
	}

	if !dara.IsNil(request.Data) {
		body["data"] = request.Data
	}

	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XGWAccess) {
		realHeaders["X-GW-Access"] = dara.String(dara.ToString(dara.StringValue(headers.XGWAccess)))
	}

	if !dara.IsNil(headers.XGWNonce) {
		realHeaders["X-GW-Nonce"] = dara.String(dara.ToString(dara.StringValue(headers.XGWNonce)))
	}

	if !dara.IsNil(headers.XGWProtocol) {
		realHeaders["X-GW-Protocol"] = dara.String(dara.ToString(dara.StringValue(headers.XGWProtocol)))
	}

	if !dara.IsNil(headers.XGWSignature) {
		realHeaders["X-GW-Signature"] = dara.String(dara.ToString(dara.StringValue(headers.XGWSignature)))
	}

	if !dara.IsNil(headers.XGWTimestamp) {
		realHeaders["X-GW-Timestamp"] = dara.String(dara.ToString(dara.StringValue(headers.XGWTimestamp)))
	}

	if !dara.IsNil(headers.XRequestId) {
		realHeaders["X-Request-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XRequestId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MedicalModelInference"),
		Version:     dara.String("2026-11-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/api/directModel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MedicalModelInferenceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates the Qwen medical service and returns the current activation status.
//
// @param request - OpenServiceRequest
//
// @param headers - OpenServiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenServiceResponse
func (client *Client) OpenServiceWithContext(ctx context.Context, request *OpenServiceRequest, headers *OpenServiceHeaders, runtime *dara.RuntimeOptions) (_result *OpenServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Communication) {
		body["communication"] = request.Communication
	}

	if !dara.IsNil(request.OpenBenefits) {
		body["openBenefits"] = request.OpenBenefits
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XGWAccess) {
		realHeaders["X-GW-Access"] = dara.String(dara.ToString(dara.StringValue(headers.XGWAccess)))
	}

	if !dara.IsNil(headers.XGWNonce) {
		realHeaders["X-GW-Nonce"] = dara.String(dara.ToString(dara.StringValue(headers.XGWNonce)))
	}

	if !dara.IsNil(headers.XGWSignature) {
		realHeaders["X-GW-Signature"] = dara.String(dara.ToString(dara.StringValue(headers.XGWSignature)))
	}

	if !dara.IsNil(headers.XGWTimestamp) {
		realHeaders["X-GW-Timestamp"] = dara.String(dara.ToString(dara.StringValue(headers.XGWTimestamp)))
	}

	if !dara.IsNil(headers.XRequestId) {
		realHeaders["X-Request-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XRequestId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenService"),
		Version:     dara.String("2026-11-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/api/medical/openService"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) medicalAgentChatWithSSECtx_opYieldFunc(_yield chan *MedicalAgentChatResponse, _yieldErr chan error, ctx context.Context, request *MedicalAgentChatRequest, headers *MedicalAgentChatHeaders, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Communication) {
		body["communication"] = request.Communication
	}

	if !dara.IsNil(request.Data) {
		body["data"] = request.Data
	}

	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XGWAccess) {
		realHeaders["X-GW-Access"] = dara.String(dara.ToString(dara.StringValue(headers.XGWAccess)))
	}

	if !dara.IsNil(headers.XGWNonce) {
		realHeaders["X-GW-Nonce"] = dara.String(dara.ToString(dara.StringValue(headers.XGWNonce)))
	}

	if !dara.IsNil(headers.XGWProtocol) {
		realHeaders["X-GW-Protocol"] = dara.String(dara.ToString(dara.StringValue(headers.XGWProtocol)))
	}

	if !dara.IsNil(headers.XGWSignature) {
		realHeaders["X-GW-Signature"] = dara.String(dara.ToString(dara.StringValue(headers.XGWSignature)))
	}

	if !dara.IsNil(headers.XGWTimestamp) {
		realHeaders["X-GW-Timestamp"] = dara.String(dara.ToString(dara.StringValue(headers.XGWTimestamp)))
	}

	if !dara.IsNil(headers.XRequestId) {
		realHeaders["X-Request-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XRequestId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MedicalAgentChat"),
		Version:     dara.String("2026-11-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/api/medicalQa"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}

func (client *Client) medicalModelInferenceWithSSECtx_opYieldFunc(_yield chan *MedicalModelInferenceResponse, _yieldErr chan error, ctx context.Context, request *MedicalModelInferenceRequest, headers *MedicalModelInferenceHeaders, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Communication) {
		body["communication"] = request.Communication
	}

	if !dara.IsNil(request.Data) {
		body["data"] = request.Data
	}

	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XGWAccess) {
		realHeaders["X-GW-Access"] = dara.String(dara.ToString(dara.StringValue(headers.XGWAccess)))
	}

	if !dara.IsNil(headers.XGWNonce) {
		realHeaders["X-GW-Nonce"] = dara.String(dara.ToString(dara.StringValue(headers.XGWNonce)))
	}

	if !dara.IsNil(headers.XGWProtocol) {
		realHeaders["X-GW-Protocol"] = dara.String(dara.ToString(dara.StringValue(headers.XGWProtocol)))
	}

	if !dara.IsNil(headers.XGWSignature) {
		realHeaders["X-GW-Signature"] = dara.String(dara.ToString(dara.StringValue(headers.XGWSignature)))
	}

	if !dara.IsNil(headers.XGWTimestamp) {
		realHeaders["X-GW-Timestamp"] = dara.String(dara.ToString(dara.StringValue(headers.XGWTimestamp)))
	}

	if !dara.IsNil(headers.XRequestId) {
		realHeaders["X-Request-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XRequestId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MedicalModelInference"),
		Version:     dara.String("2026-11-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/api/directModel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}
