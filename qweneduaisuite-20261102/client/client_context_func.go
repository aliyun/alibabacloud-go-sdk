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
// Activates the learning capability for the current Alibaba Cloud account and supports querying the activation status.
//
// Description:
//
// ## Request description
//
// - This operation activates the Qwen learning AI suite capability for the current Alibaba Cloud account.
//
// - You must call this operation first. You can call business API operations such as problem solving and grading only after `ACTIVE` is returned.
//
// ### Precautions
//
// - Both Alibaba Cloud accounts and RAM users can call this operation.
//
// - When a RAM user calls this operation, ensure that the RAM user has the required RAM permissions. The service is activated for the Alibaba Cloud account to which the RAM user belongs.
//
// - This operation is idempotent based on the Alibaba Cloud account ID. Repeated calls do not result in multiple activations.
//
// @param request - ActivateQwenEduApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateQwenEduApiResponse
func (client *Client) ActivateQwenEduApiWithContext(ctx context.Context, request *ActivateQwenEduApiRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ActivateQwenEduApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActivateQwenEduApi"),
		Version:     dara.String("2026-11-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/service/activate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateQwenEduApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a single-page homework image and asynchronously retrieves the grading results.
//
// Description:
//
// ## Request description
//
// - This operation submits a homework image to the server for question segmentation, recognition, and grading.
//
// - You can submit the image by using `imageBase64` or `imageUrl`. `imageBase64` takes precedence.
//
// - The optional `dewarp` parameter specifies whether to correct the image distortion. Default value: `true`.
//
// - Provide `requestId` to facilitate troubleshooting and reconciliation.
//
// - Image size limit: The Base64-encoded image size cannot exceed 15 MB.
//
// - All requests and responses use the JSON format.
//
// @param request - CreateEduHomeworkCorrectSubmitRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEduHomeworkCorrectSubmitResponse
func (client *Client) CreateEduHomeworkCorrectSubmitWithContext(ctx context.Context, request *CreateEduHomeworkCorrectSubmitRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEduHomeworkCorrectSubmitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Dewarp) {
		body["dewarp"] = request.Dewarp
	}

	if !dara.IsNil(request.ImageBase64) {
		body["imageBase64"] = request.ImageBase64
	}

	if !dara.IsNil(request.ImageUrl) {
		body["imageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEduHomeworkCorrectSubmit"),
		Version:     dara.String("2026-11-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/homework/correct/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEduHomeworkCorrectSubmitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides question resolution services by using image URLs or Base64 encoding.
//
// Description:
//
// ## Request description
//
// - This API supports obtaining question resolution results by providing the URL or Base64 encoding of a question image.
//
// - Supports processing a single image. The image format must be JPG, JPEG, or PNG.
//
// - The image must be clear and complete. Avoid tilted, obscured, or irrelevant content.
//
// - Each request can process only one image. Multi-image or plain text requests are not supported.
//
// - The recommended image size is up to 1 MB, with a maximum of 10 MB. The image width and height must not exceed 8,000 pixels, and the total pixel count must not exceed 2.6 million.
//
// - You must use the HTTPS protocol to call the API and provide image URLs.
//
// - The user_id parameter must be an anonymized identifier. Do not pass direct identity information such as names or phone numbers.
//
// - The API uses the HTTP POST method combined with Server-Sent Events (SSE) streaming response. The client must read the response event by event, rather than waiting for the connection to close to read the entire response.
//
// - A successful response returns the question resolution process through a series of message events and ends with a complete event. If an error occurs, an error event is returned.
//
// - For retryable errors such as TOO_MANY_ACTIVE_STREAMS, use an exponential backoff strategy to retry the request.
//
// - Set a reasonable timeout period of at least 3 minutes on the client, and ensure that the proxy layer does not cache SSE responses.
//
// @param tmpReq - CreateEduProblemSolveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEduProblemSolveResponse
func (client *Client) CreateEduProblemSolveWithSSECtx(ctx context.Context, tmpReq *CreateEduProblemSolveRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *CreateEduProblemSolveResponse, _yieldErr chan error) {
	defer close(_yield)
	client.createEduProblemSolveWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, headers, runtime)
	return
}

// Summary:
//
// Provides question resolution services by using image URLs or Base64 encoding.
//
// Description:
//
// ## Request description
//
// - This API supports obtaining question resolution results by providing the URL or Base64 encoding of a question image.
//
// - Supports processing a single image. The image format must be JPG, JPEG, or PNG.
//
// - The image must be clear and complete. Avoid tilted, obscured, or irrelevant content.
//
// - Each request can process only one image. Multi-image or plain text requests are not supported.
//
// - The recommended image size is up to 1 MB, with a maximum of 10 MB. The image width and height must not exceed 8,000 pixels, and the total pixel count must not exceed 2.6 million.
//
// - You must use the HTTPS protocol to call the API and provide image URLs.
//
// - The user_id parameter must be an anonymized identifier. Do not pass direct identity information such as names or phone numbers.
//
// - The API uses the HTTP POST method combined with Server-Sent Events (SSE) streaming response. The client must read the response event by event, rather than waiting for the connection to close to read the entire response.
//
// - A successful response returns the question resolution process through a series of message events and ends with a complete event. If an error occurs, an error event is returned.
//
// - For retryable errors such as TOO_MANY_ACTIVE_STREAMS, use an exponential backoff strategy to retry the request.
//
// - Set a reasonable timeout period of at least 3 minutes on the client, and ensure that the proxy layer does not cache SSE responses.
//
// @param tmpReq - CreateEduProblemSolveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEduProblemSolveResponse
func (client *Client) CreateEduProblemSolveWithContext(ctx context.Context, tmpReq *CreateEduProblemSolveRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEduProblemSolveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateEduProblemSolveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Query) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, dara.String("query"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.QueryShrink) {
		body["query"] = request.QueryShrink
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEduProblemSolve"),
		Version:     dara.String("2026-11-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/problem/solve/stream"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEduProblemSolveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the progress and final results of homework grading based on the task ID.
//
// Description:
//
// ## Request description
//
// - This operation queries the status and detailed results of single-page homework grading by using the `taskId` returned when the task is submitted.
//
// - The caller should poll this operation at an interval of 1 to 2 seconds until the task status changes to `finished` or `failed`.
//
// - When the task status is `finished`, the response contains complete grading information, including question recognition, grading conclusions, dewarped images, and rendered images.
//
// - If the task fails (the status is `failed`), check the `message` field in the response or the specific fault information in `data`.
//
// - All requests and responses are in JSON format.
//
// @param request - GetEduHomeworkCorrectionResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEduHomeworkCorrectionResultResponse
func (client *Client) GetEduHomeworkCorrectionResultWithContext(ctx context.Context, request *GetEduHomeworkCorrectionResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEduHomeworkCorrectionResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEduHomeworkCorrectionResult"),
		Version:     dara.String("2026-11-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/homework/correct/result"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEduHomeworkCorrectionResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the current account has activated Qwen education-related services.
//
// Description:
//
// ## Request description
//
// This operation queries whether a specified Alibaba Cloud account or RAM user has successfully activated the Qwen learning AI suite. Before calling any business API operations, you must use this operation to confirm that the service status is `ACTIVE`. If the service is not activated, perform the corresponding operations based on the returned status information, such as waiting for a period of time and retrying.
//
// ### Precautions
//
// - This operation supports calls from Alibaba Cloud accounts and RAM users.
//
// - When a RAM user calls this operation, make sure that the RAM user has the required RAM permissions. The service is activated for the Alibaba Cloud account to which the RAM user belongs.
//
// - This operation is idempotent based on the Alibaba Cloud account ID. Repeated calls do not result in multiple activations.
//
// - If the status is `ACTIVATING`, retry the query after the time interval recommended by the returned `retry_after_seconds` field.
//
// - If a failure occurs (`FAILED`), refer to the specific error code, take corresponding measures, and then try again.
//
// @param request - GetQwenEduApiActivationStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQwenEduApiActivationStatusResponse
func (client *Client) GetQwenEduApiActivationStatusWithContext(ctx context.Context, request *GetQwenEduApiActivationStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetQwenEduApiActivationStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQwenEduApiActivationStatus"),
		Version:     dara.String("2026-11-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/service/activation/status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQwenEduApiActivationStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) createEduProblemSolveWithSSECtx_opYieldFunc(_yield chan *CreateEduProblemSolveResponse, _yieldErr chan error, ctx context.Context, tmpReq *CreateEduProblemSolveRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &CreateEduProblemSolveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Query) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, dara.String("query"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.QueryShrink) {
		body["query"] = request.QueryShrink
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEduProblemSolve"),
		Version:     dara.String("2026-11-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/problem/solve/stream"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
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
