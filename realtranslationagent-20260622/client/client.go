// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
	EnableValidate  *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"public": dara.String("realtranslationagent.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("realtranslationagent"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) _postOSSObject(bucketName *string, form map[string]interface{}, runtime *dara.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_runtime := dara.NewRuntimeObject(map[string]interface{}{
		"key":            dara.ToString(dara.Default(dara.StringValue(runtime.Key), dara.StringValue(client.Key))),
		"cert":           dara.ToString(dara.Default(dara.StringValue(runtime.Cert), dara.StringValue(client.Cert))),
		"ca":             dara.ToString(dara.Default(dara.StringValue(runtime.Ca), dara.StringValue(client.Ca))),
		"readTimeout":    dara.ForceInt(dara.Default(dara.IntValue(runtime.ReadTimeout), dara.IntValue(client.ReadTimeout))),
		"connectTimeout": dara.ForceInt(dara.Default(dara.IntValue(runtime.ConnectTimeout), dara.IntValue(client.ConnectTimeout))),
		"httpProxy":      dara.ToString(dara.Default(dara.StringValue(runtime.HttpProxy), dara.StringValue(client.HttpProxy))),
		"httpsProxy":     dara.ToString(dara.Default(dara.StringValue(runtime.HttpsProxy), dara.StringValue(client.HttpsProxy))),
		"noProxy":        dara.ToString(dara.Default(dara.StringValue(runtime.NoProxy), dara.StringValue(client.NoProxy))),
		"socks5Proxy":    dara.ToString(dara.Default(dara.StringValue(runtime.Socks5Proxy), dara.StringValue(client.Socks5Proxy))),
		"socks5NetWork":  dara.ToString(dara.Default(dara.StringValue(runtime.Socks5NetWork), dara.StringValue(client.Socks5NetWork))),
		"maxIdleConns":   dara.ForceInt(dara.Default(dara.IntValue(runtime.MaxIdleConns), dara.IntValue(client.MaxIdleConns))),
		"retryOptions":   client.RetryOptions,
		"ignoreSSL":      dara.ForceBoolean(dara.Default(dara.BoolValue(runtime.IgnoreSSL), false)),
		"tlsMinVersion":  dara.StringValue(client.TlsMinVersion),
	})

	var retryPolicyContext *dara.RetryPolicyContext
	var request_ *dara.Request
	var response_ *dara.Response
	var _resultErr error
	retriesAttempted := int(0)
	retryPolicyContext = &dara.RetryPolicyContext{
		RetriesAttempted: retriesAttempted,
	}

	_result = make(map[string]interface{})
	for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {
		_resultErr = nil
		_backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)
		dara.Sleep(_backoffDelayTime)

		request_ = dara.NewRequest()
		boundary := dara.GetBoundary()
		tmp := dara.ToString(form["host"])
		host := dara.StringValue(bucketName) + "." + tmp
		request_.Protocol = dara.String("HTTPS")
		request_.Method = dara.String("POST")
		request_.Pathname = dara.String("/")
		request_.Headers = map[string]*string{
			"host":       dara.String(host),
			"date":       openapiutil.GetDateUTCString(),
			"user-agent": openapiutil.GetUserAgent(dara.String("")),
		}
		request_.Headers["content-type"] = dara.String("multipart/form-data; boundary=" + boundary)
		request_.Body = dara.ToFileForm(form, boundary)
		response_, _err = dara.DoRequest(request_, _runtime)
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		_result, _err = _postOSSObject_opResponse(response_)
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		return _result, _err
	}
	if dara.BoolValue(client.DisableSDKError) != true {
		_resultErr = dara.TeaSDKError(_resultErr)
	}
	return _result, _resultErr
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a translation task that is currently running.
//
// Description:
//
// *Billing description**
//
// After the task is successfully canceled, the Credits frozen for this translation task will be fully refunded to your account.
//
// **Before you begin**
//
// - This operation only supports canceling translation tasks that are in the processing state. Tasks that are completed or failed cannot be canceled.
//
// @param request - CancelTranslationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelTranslationTaskResponse
func (client *Client) CancelTranslationTaskWithOptions(request *CancelTranslationTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelTranslationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		query["APIKey"] = request.APIKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelTranslationTask"),
		Version:     dara.String("2026-06-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelTranslationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a translation task that is currently running.
//
// Description:
//
// *Billing description**
//
// After the task is successfully canceled, the Credits frozen for this translation task will be fully refunded to your account.
//
// **Before you begin**
//
// - This operation only supports canceling translation tasks that are in the processing state. Tasks that are completed or failed cannot be canceled.
//
// @param request - CancelTranslationTaskRequest
//
// @return CancelTranslationTaskResponse
func (client *Client) CancelTranslationTask(request *CancelTranslationTaskRequest) (_result *CancelTranslationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelTranslationTaskResponse{}
	_body, _err := client.CancelTranslationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the download URL of the original file for a translation task.
//
// @param request - GetOriginalFileUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOriginalFileUrlResponse
func (client *Client) GetOriginalFileUrlWithOptions(request *GetOriginalFileUrlRequest, runtime *dara.RuntimeOptions) (_result *GetOriginalFileUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		query["APIKey"] = request.APIKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOriginalFileUrl"),
		Version:     dara.String("2026-06-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOriginalFileUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the download URL of the original file for a translation task.
//
// @param request - GetOriginalFileUrlRequest
//
// @return GetOriginalFileUrlResponse
func (client *Client) GetOriginalFileUrl(request *GetOriginalFileUrlRequest) (_result *GetOriginalFileUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOriginalFileUrlResponse{}
	_body, _err := client.GetOriginalFileUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the download URL of the translated file for a translation task.
//
// @param request - GetTranslatedFileUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranslatedFileUrlResponse
func (client *Client) GetTranslatedFileUrlWithOptions(request *GetTranslatedFileUrlRequest, runtime *dara.RuntimeOptions) (_result *GetTranslatedFileUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		query["APIKey"] = request.APIKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTranslatedFileUrl"),
		Version:     dara.String("2026-06-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTranslatedFileUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the download URL of the translated file for a translation task.
//
// @param request - GetTranslatedFileUrlRequest
//
// @return GetTranslatedFileUrlResponse
func (client *Client) GetTranslatedFileUrl(request *GetTranslatedFileUrlRequest) (_result *GetTranslatedFileUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTranslatedFileUrlResponse{}
	_body, _err := client.GetTranslatedFileUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a translation task.
//
// @param request - GetTranslationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranslationTaskResponse
func (client *Client) GetTranslationTaskWithOptions(request *GetTranslationTaskRequest, runtime *dara.RuntimeOptions) (_result *GetTranslationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		query["APIKey"] = request.APIKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTranslationTask"),
		Version:     dara.String("2026-06-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTranslationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a translation task.
//
// @param request - GetTranslationTaskRequest
//
// @return GetTranslationTaskResponse
func (client *Client) GetTranslationTask(request *GetTranslationTaskRequest) (_result *GetTranslationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTranslationTaskResponse{}
	_body, _err := client.GetTranslationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries translation tasks by paging.
//
// @param request - ListTranslationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTranslationTasksResponse
func (client *Client) ListTranslationTasksWithOptions(request *ListTranslationTasksRequest, runtime *dara.RuntimeOptions) (_result *ListTranslationTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		query["APIKey"] = request.APIKey
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OriginalFileName) {
		query["OriginalFileName"] = request.OriginalFileName
	}

	if !dara.IsNil(request.SourceLanguage) {
		query["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TargetLanguage) {
		query["TargetLanguage"] = request.TargetLanguage
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTranslationTasks"),
		Version:     dara.String("2026-06-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTranslationTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries translation tasks by paging.
//
// @param request - ListTranslationTasksRequest
//
// @return ListTranslationTasksResponse
func (client *Client) ListTranslationTasks(request *ListTranslationTasksRequest) (_result *ListTranslationTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTranslationTasksResponse{}
	_body, _err := client.ListTranslationTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a translation task. You can submit a new translation task by passing in a TaskId, or resubmit a historical task for translation by passing in a BaseTaskId. After successful submission, the translation task ID and current task status are returned. You can use the task ID to call subsequent operations to query translation progress and results.
//
// Description:
//
// *Billing description**
//
// This operation involves Credits consumption. Before submitting a translation task, ensure that your account has sufficient Credits balance. After calling `UploadTranslationFile`, you can check the `CreditsAvailable` field in the response to confirm whether your current balance meets the requirements of this translation task. For detailed billing information, refer to the `CreditBreakdown` field.
//
// **Task submission description**
//
// - To submit a new translation task, pass in the `TaskId` returned by the `UploadTranslationFile` operation.
//
// - To resubmit a historical task for translation, pass in the task ID of a previously submitted translation task, which is the `BaseTaskId`.
//
// - You must pass in either `TaskId` or `BaseTaskId`. You cannot pass in both at the same time.
//
// **Precautions**
//
// - The `Style` parameter takes effect only when the translation file is a PPT file. Passing in this parameter for files in other formats has no effect.
//
// - For new tasks, you can obtain the list of available fonts from the `Fonts` field in the response of `UploadTranslationFile`. For retranslation of historical tasks, you can obtain the list of available fonts by calling the `GetTranslationTask` operation.
//
// @param tmpReq - SubmitTranslationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTranslationTaskResponse
func (client *Client) SubmitTranslationTaskWithOptions(tmpReq *SubmitTranslationTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitTranslationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitTranslationTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Config) {
		request.ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Config, dara.String("Config"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CustomTerms) {
		request.CustomTermsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomTerms, dara.String("CustomTerms"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		query["APIKey"] = request.APIKey
	}

	if !dara.IsNil(request.CustomTermsShrink) {
		query["CustomTerms"] = request.CustomTermsShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseTaskId) {
		body["BaseTaskId"] = request.BaseTaskId
	}

	if !dara.IsNil(request.ConfigShrink) {
		body["Config"] = request.ConfigShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitTranslationTask"),
		Version:     dara.String("2026-06-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitTranslationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a translation task. You can submit a new translation task by passing in a TaskId, or resubmit a historical task for translation by passing in a BaseTaskId. After successful submission, the translation task ID and current task status are returned. You can use the task ID to call subsequent operations to query translation progress and results.
//
// Description:
//
// *Billing description**
//
// This operation involves Credits consumption. Before submitting a translation task, ensure that your account has sufficient Credits balance. After calling `UploadTranslationFile`, you can check the `CreditsAvailable` field in the response to confirm whether your current balance meets the requirements of this translation task. For detailed billing information, refer to the `CreditBreakdown` field.
//
// **Task submission description**
//
// - To submit a new translation task, pass in the `TaskId` returned by the `UploadTranslationFile` operation.
//
// - To resubmit a historical task for translation, pass in the task ID of a previously submitted translation task, which is the `BaseTaskId`.
//
// - You must pass in either `TaskId` or `BaseTaskId`. You cannot pass in both at the same time.
//
// **Precautions**
//
// - The `Style` parameter takes effect only when the translation file is a PPT file. Passing in this parameter for files in other formats has no effect.
//
// - For new tasks, you can obtain the list of available fonts from the `Fonts` field in the response of `UploadTranslationFile`. For retranslation of historical tasks, you can obtain the list of available fonts by calling the `GetTranslationTask` operation.
//
// @param request - SubmitTranslationTaskRequest
//
// @return SubmitTranslationTaskResponse
func (client *Client) SubmitTranslationTask(request *SubmitTranslationTaskRequest) (_result *SubmitTranslationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitTranslationTaskResponse{}
	_body, _err := client.SubmitTranslationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads a document, parses document-related information, and generates a translation task. After a successful upload, the task ID and document parsing results are returned, including word count, page count, estimated Credits consumption, estimated translation time, detected language type, and font list. The system also performs sensitive information detection on the uploaded document, and you can decide whether to proceed with submitting the translation task based on the detection results.
//
// Description:
//
// > - This operation only involves document upload and information estimation. **No fees are incurred.*	- Credits consumption starts only after you **officially submit the translation*	- task.
//
// **Language detection**
//
// The system automatically detects the language type of the uploaded document. Currently, Chinese is supported.
//
// **Sensitive information detection**
//
// The system performs sensitive information detection on the uploaded document. If sensitive information is detected, the `SensitiveDetected` field in the response is set to `true`, and the `SensitiveTags` field returns the list of matched keywords.
//
// >  - You can decide whether to proceed with submitting the translation task based on your actual needs.
//
// >  - If the translation quality setting is set to ultimate mode when you submit the task, the system automatically switches the **portions containing sensitive information*	- to auto mode.
//
// **Notes**
//
// - Make sure the uploaded document format is supported by the system. Otherwise, parsing may fail.
//
// - The `EstimatedCostCredits` value in the response is the estimated Credits consumption. The actual consumption is based on the settlement after the translation task is officially submitted.
//
// - The `EstimatedTime` value in the response is the estimated translation duration in milliseconds. The actual translation duration may vary depending on document complexity.
//
// - The `Fonts` field in the response contains the languages that support font modification and the corresponding font lists. You can select an appropriate font based on the target language.
//
// @param request - UploadTranslationFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadTranslationFileResponse
func (client *Client) UploadTranslationFileWithOptions(request *UploadTranslationFileRequest, runtime *dara.RuntimeOptions) (_result *UploadTranslationFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		query["APIKey"] = request.APIKey
	}

	if !dara.IsNil(request.File) {
		query["File"] = request.File
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadTranslationFile"),
		Version:     dara.String("2026-06-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadTranslationFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a document, parses document-related information, and generates a translation task. After a successful upload, the task ID and document parsing results are returned, including word count, page count, estimated Credits consumption, estimated translation time, detected language type, and font list. The system also performs sensitive information detection on the uploaded document, and you can decide whether to proceed with submitting the translation task based on the detection results.
//
// Description:
//
// > - This operation only involves document upload and information estimation. **No fees are incurred.*	- Credits consumption starts only after you **officially submit the translation*	- task.
//
// **Language detection**
//
// The system automatically detects the language type of the uploaded document. Currently, Chinese is supported.
//
// **Sensitive information detection**
//
// The system performs sensitive information detection on the uploaded document. If sensitive information is detected, the `SensitiveDetected` field in the response is set to `true`, and the `SensitiveTags` field returns the list of matched keywords.
//
// >  - You can decide whether to proceed with submitting the translation task based on your actual needs.
//
// >  - If the translation quality setting is set to ultimate mode when you submit the task, the system automatically switches the **portions containing sensitive information*	- to auto mode.
//
// **Notes**
//
// - Make sure the uploaded document format is supported by the system. Otherwise, parsing may fail.
//
// - The `EstimatedCostCredits` value in the response is the estimated Credits consumption. The actual consumption is based on the settlement after the translation task is officially submitted.
//
// - The `EstimatedTime` value in the response is the estimated translation duration in milliseconds. The actual translation duration may vary depending on document complexity.
//
// - The `Fonts` field in the response contains the languages that support font modification and the corresponding font lists. You can select an appropriate font based on the target language.
//
// @param request - UploadTranslationFileRequest
//
// @return UploadTranslationFileResponse
func (client *Client) UploadTranslationFile(request *UploadTranslationFileRequest) (_result *UploadTranslationFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadTranslationFileResponse{}
	_body, _err := client.UploadTranslationFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadTranslationFileAdvance(request *UploadTranslationFileAdvanceRequest, runtime *dara.RuntimeOptions) (_result *UploadTranslationFileResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("RealTranslationAgent"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	uploadTranslationFileReq := &UploadTranslationFileRequest{}
	openapiutil.Convert(request, uploadTranslationFileReq)
	if !dara.IsNil(request.FileObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader, runtime)
		if _err != nil {
			return _result, _err
		}
		uploadTranslationFileReq.File = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	uploadTranslationFileResp, _err := client.UploadTranslationFileWithOptions(uploadTranslationFileReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = uploadTranslationFileResp
	return _result, _err
}

func _postOSSObject_opResponse(response_ *dara.Response) (_result map[string]interface{}, _err error) {
	var respMap map[string]interface{}
	bodyStr, _err := dara.ReadAsString(response_.Body)
	if _err != nil {
		return _result, _err
	}

	if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 600) {
		respMap = dara.ParseXml(bodyStr, nil)
		err := dara.ToMap(respMap["Error"])
		_err = &openapi.ClientError{
			Code:    dara.String(dara.ToString(err["Code"])),
			Message: dara.String(dara.ToString(err["Message"])),
			Data: map[string]interface{}{
				"httpCode":  dara.IntValue(response_.StatusCode),
				"requestId": dara.ToString(err["RequestId"]),
				"hostId":    dara.ToString(err["HostId"]),
			},
		}
		return _result, _err
	}

	respMap = dara.ParseXml(bodyStr, nil)
	_result = make(map[string]interface{})
	_err = dara.Convert(dara.ToMap(respMap), &_result)

	return _result, _err
}
