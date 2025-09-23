// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Used for external deletion of community samples in risk identification services.
//
// @param request - RevokeFeedbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeFeedbackResponse
func (client *Client) RevokeFeedbackWithContext(ctx context.Context, request *RevokeFeedbackRequest, runtime *dara.RuntimeOptions) (_result *RevokeFeedbackResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SampleType) {
		body["SampleType"] = request.SampleType
	}

	if !dara.IsNil(request.Value) {
		body["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeFeedback"),
		Version:     dara.String("2021-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeFeedbackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Used for the external creation of community samples in risk identification services.
//
// @param request - SendFeedbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendFeedbackResponse
func (client *Client) SendFeedbackWithContext(ctx context.Context, request *SendFeedbackRequest, runtime *dara.RuntimeOptions) (_result *SendFeedbackResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Reason) {
		query["Reason"] = request.Reason
	}

	if !dara.IsNil(request.RiskLabel) {
		query["RiskLabel"] = request.RiskLabel
	}

	if !dara.IsNil(request.SampleType) {
		query["SampleType"] = request.SampleType
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendFeedback"),
		Version:     dara.String("2021-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendFeedbackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Single User API for Sample Creation
//
// @param request - UploadSampleApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadSampleApiResponse
func (client *Client) UploadSampleApiWithContext(ctx context.Context, request *UploadSampleApiRequest, runtime *dara.RuntimeOptions) (_result *UploadSampleApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataType) {
		query["DataType"] = request.DataType
	}

	if !dara.IsNil(request.DataValue) {
		query["DataValue"] = request.DataValue
	}

	if !dara.IsNil(request.SampleType) {
		query["SampleType"] = request.SampleType
	}

	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadSampleApi"),
		Version:     dara.String("2021-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadSampleApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
