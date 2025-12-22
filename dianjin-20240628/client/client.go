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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("dianjin"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
		request_.Protocol = dara.String("HTTPS")
		request_.Method = dara.String("POST")
		request_.Pathname = dara.String("/")
		request_.Headers = map[string]*string{
			"host":       dara.String(dara.ToString(form["host"])),
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
// 创建按年文档总结任务
//
// @param request - CreateAnnualDocSummaryTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAnnualDocSummaryTaskResponse
func (client *Client) CreateAnnualDocSummaryTaskWithOptions(workspaceId *string, request *CreateAnnualDocSummaryTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAnnualDocSummaryTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AnaYears) {
		body["anaYears"] = request.AnaYears
	}

	if !dara.IsNil(request.DocInfos) {
		body["docInfos"] = request.DocInfos
	}

	if !dara.IsNil(request.EnableTable) {
		body["enableTable"] = request.EnableTable
	}

	if !dara.IsNil(request.Instruction) {
		body["instruction"] = request.Instruction
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAnnualDocSummaryTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/summary/doc/annual"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAnnualDocSummaryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建按年文档总结任务
//
// @param request - CreateAnnualDocSummaryTaskRequest
//
// @return CreateAnnualDocSummaryTaskResponse
func (client *Client) CreateAnnualDocSummaryTask(workspaceId *string, request *CreateAnnualDocSummaryTaskRequest) (_result *CreateAnnualDocSummaryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAnnualDocSummaryTaskResponse{}
	_body, _err := client.CreateAnnualDocSummaryTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建外呼会话
//
// @param request - CreateDialogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDialogResponse
func (client *Client) CreateDialogWithOptions(workspaceId *string, request *CreateDialogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDialogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Channel) {
		body["channel"] = request.Channel
	}

	if !dara.IsNil(request.EnableLibrary) {
		body["enableLibrary"] = request.EnableLibrary
	}

	if !dara.IsNil(request.MetaData) {
		body["metaData"] = request.MetaData
	}

	if !dara.IsNil(request.PlayCode) {
		body["playCode"] = request.PlayCode
	}

	if !dara.IsNil(request.QaLibraryList) {
		body["qaLibraryList"] = request.QaLibraryList
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.SelfDirected) {
		body["selfDirected"] = request.SelfDirected
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDialog"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/dialog/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDialogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建外呼会话
//
// @param request - CreateDialogRequest
//
// @return CreateDialogResponse
func (client *Client) CreateDialog(workspaceId *string, request *CreateDialogRequest) (_result *CreateDialogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDialogResponse{}
	_body, _err := client.CreateDialogWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建会话分析任务
//
// @param request - CreateDialogAnalysisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDialogAnalysisTaskResponse
func (client *Client) CreateDialogAnalysisTaskWithOptions(workspaceId *string, request *CreateDialogAnalysisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDialogAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AnalysisNodes) {
		body["analysisNodes"] = request.AnalysisNodes
	}

	if !dara.IsNil(request.ConversationList) {
		body["conversationList"] = request.ConversationList
	}

	if !dara.IsNil(request.MetaData) {
		body["metaData"] = request.MetaData
	}

	if !dara.IsNil(request.PlayCode) {
		body["playCode"] = request.PlayCode
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDialogAnalysisTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/dialog/analysis/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDialogAnalysisTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建会话分析任务
//
// @param request - CreateDialogAnalysisTaskRequest
//
// @return CreateDialogAnalysisTaskResponse
func (client *Client) CreateDialogAnalysisTask(workspaceId *string, request *CreateDialogAnalysisTaskRequest) (_result *CreateDialogAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDialogAnalysisTaskResponse{}
	_body, _err := client.CreateDialogAnalysisTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建财报总结任务
//
// @param request - CreateDocsSummaryTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDocsSummaryTaskResponse
func (client *Client) CreateDocsSummaryTaskWithOptions(workspaceId *string, request *CreateDocsSummaryTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDocsSummaryTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocInfos) {
		body["docInfos"] = request.DocInfos
	}

	if !dara.IsNil(request.EnableTable) {
		body["enableTable"] = request.EnableTable
	}

	if !dara.IsNil(request.Instruction) {
		body["instruction"] = request.Instruction
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDocsSummaryTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/summary/docs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDocsSummaryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建财报总结任务
//
// @param request - CreateDocsSummaryTaskRequest
//
// @return CreateDocsSummaryTaskResponse
func (client *Client) CreateDocsSummaryTask(workspaceId *string, request *CreateDocsSummaryTaskRequest) (_result *CreateDocsSummaryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDocsSummaryTaskResponse{}
	_body, _err := client.CreateDocsSummaryTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建财报总结任务
//
// @param request - CreateFinReportSummaryTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFinReportSummaryTaskResponse
func (client *Client) CreateFinReportSummaryTaskWithOptions(workspaceId *string, request *CreateFinReportSummaryTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFinReportSummaryTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["docId"] = request.DocId
	}

	if !dara.IsNil(request.EnableTable) {
		body["enableTable"] = request.EnableTable
	}

	if !dara.IsNil(request.EndPage) {
		body["endPage"] = request.EndPage
	}

	if !dara.IsNil(request.Instruction) {
		body["instruction"] = request.Instruction
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.StartPage) {
		body["startPage"] = request.StartPage
	}

	if !dara.IsNil(request.TaskType) {
		body["taskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFinReportSummaryTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/summary"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFinReportSummaryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建财报总结任务
//
// @param request - CreateFinReportSummaryTaskRequest
//
// @return CreateFinReportSummaryTaskResponse
func (client *Client) CreateFinReportSummaryTask(workspaceId *string, request *CreateFinReportSummaryTaskRequest) (_result *CreateFinReportSummaryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFinReportSummaryTaskResponse{}
	_body, _err := client.CreateFinReportSummaryTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建图片检测任务
//
// @param request - CreateImageDetectionTaskRequest
//
// @param headers - CreateImageDetectionTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageDetectionTaskResponse
func (client *Client) CreateImageDetectionTaskWithOptions(workspaceId *string, request *CreateImageDetectionTaskRequest, headers *CreateImageDetectionTaskHeaders, runtime *dara.RuntimeOptions) (_result *CreateImageDetectionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileInfo) {
		body["fileInfo"] = request.FileInfo
	}

	if !dara.IsNil(request.FileUrl) {
		body["fileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XLoadTest) {
		realHeaders["X-Load-Test"] = dara.String(dara.Stringify(dara.BoolValue(headers.XLoadTest)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageDetectionTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/imageDetect/task/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageDetectionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建图片检测任务
//
// @param request - CreateImageDetectionTaskRequest
//
// @return CreateImageDetectionTaskResponse
func (client *Client) CreateImageDetectionTask(workspaceId *string, request *CreateImageDetectionTaskRequest) (_result *CreateImageDetectionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateImageDetectionTaskHeaders{}
	_result = &CreateImageDetectionTaskResponse{}
	_body, _err := client.CreateImageDetectionTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建文档库
//
// @param request - CreateLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLibraryResponse
func (client *Client) CreateLibraryWithOptions(workspaceId *string, request *CreateLibraryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLibraryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.IndexSetting) {
		body["indexSetting"] = request.IndexSetting
	}

	if !dara.IsNil(request.LibraryName) {
		body["libraryName"] = request.LibraryName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLibrary"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLibraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建文档库
//
// @param request - CreateLibraryRequest
//
// @return CreateLibraryResponse
func (client *Client) CreateLibrary(workspaceId *string, request *CreateLibraryRequest) (_result *CreateLibraryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLibraryResponse{}
	_body, _err := client.CreateLibraryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建PDF翻译任务
//
// @param request - CreatePdfTranslateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePdfTranslateTaskResponse
func (client *Client) CreatePdfTranslateTaskWithOptions(workspaceId *string, request *CreatePdfTranslateTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePdfTranslateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["docId"] = request.DocId
	}

	if !dara.IsNil(request.Knowledge) {
		body["knowledge"] = request.Knowledge
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.TranslateTo) {
		body["translateTo"] = request.TranslateTo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePdfTranslateTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/pdfTranslate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePdfTranslateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建PDF翻译任务
//
// @param request - CreatePdfTranslateTaskRequest
//
// @return CreatePdfTranslateTaskResponse
func (client *Client) CreatePdfTranslateTask(workspaceId *string, request *CreatePdfTranslateTaskRequest) (_result *CreatePdfTranslateTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePdfTranslateTaskResponse{}
	_body, _err := client.CreatePdfTranslateTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建预定义文档
//
// @param request - CreatePredefinedDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePredefinedDocumentResponse
func (client *Client) CreatePredefinedDocumentWithOptions(workspaceId *string, request *CreatePredefinedDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePredefinedDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Chunks) {
		body["chunks"] = request.Chunks
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.Metadata) {
		body["metadata"] = request.Metadata
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePredefinedDocument"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/document/createPredefinedDocument"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePredefinedDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建预定义文档
//
// @param request - CreatePredefinedDocumentRequest
//
// @return CreatePredefinedDocumentResponse
func (client *Client) CreatePredefinedDocument(workspaceId *string, request *CreatePredefinedDocumentRequest) (_result *CreatePredefinedDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePredefinedDocumentResponse{}
	_body, _err := client.CreatePredefinedDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建财报总结的任务
//
// @param request - CreateQualityCheckTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQualityCheckTaskResponse
func (client *Client) CreateQualityCheckTaskWithOptions(workspaceId *string, request *CreateQualityCheckTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateQualityCheckTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConversationList) {
		body["conversationList"] = request.ConversationList
	}

	if !dara.IsNil(request.GmtService) {
		body["gmtService"] = request.GmtService
	}

	if !dara.IsNil(request.MetaData) {
		body["metaData"] = request.MetaData
	}

	if !dara.IsNil(request.QualityGroup) {
		body["qualityGroup"] = request.QualityGroup
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.SceneCode) {
		body["sceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQualityCheckTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/qualitycheck/task/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQualityCheckTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建财报总结的任务
//
// @param request - CreateQualityCheckTaskRequest
//
// @return CreateQualityCheckTaskResponse
func (client *Client) CreateQualityCheckTask(workspaceId *string, request *CreateQualityCheckTaskRequest) (_result *CreateQualityCheckTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateQualityCheckTaskResponse{}
	_body, _err := client.CreateQualityCheckTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建视频生成任务
//
// @param request - CreateVideoCreationTaskRequest
//
// @param headers - CreateVideoCreationTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVideoCreationTaskResponse
func (client *Client) CreateVideoCreationTaskWithOptions(workspaceId *string, request *CreateVideoCreationTaskRequest, headers *CreateVideoCreationTaskHeaders, runtime *dara.RuntimeOptions) (_result *CreateVideoCreationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CreationInstruction) {
		body["creationInstruction"] = request.CreationInstruction
	}

	if !dara.IsNil(request.FileInfo) {
		body["fileInfo"] = request.FileInfo
	}

	if !dara.IsNil(request.ImageDetectionTaskId) {
		body["imageDetectionTaskId"] = request.ImageDetectionTaskId
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XLoadTest) {
		realHeaders["X-Load-Test"] = dara.String(dara.Stringify(dara.BoolValue(headers.XLoadTest)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVideoCreationTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/videoCreation/task/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVideoCreationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建视频生成任务
//
// @param request - CreateVideoCreationTaskRequest
//
// @return CreateVideoCreationTaskResponse
func (client *Client) CreateVideoCreationTask(workspaceId *string, request *CreateVideoCreationTaskRequest) (_result *CreateVideoCreationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateVideoCreationTaskHeaders{}
	_result = &CreateVideoCreationTaskResponse{}
	_body, _err := client.CreateVideoCreationTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除文档
//
// @param request - DeleteDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDocumentResponse
func (client *Client) DeleteDocumentWithOptions(workspaceId *string, request *DeleteDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocIds) {
		body["docIds"] = request.DocIds
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDocument"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/document/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除文档
//
// @param request - DeleteDocumentRequest
//
// @return DeleteDocumentResponse
func (client *Client) DeleteDocument(workspaceId *string, request *DeleteDocumentRequest) (_result *DeleteDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDocumentResponse{}
	_body, _err := client.DeleteDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除文档库
//
// @param request - DeleteLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLibraryResponse
func (client *Client) DeleteLibraryWithOptions(workspaceId *string, request *DeleteLibraryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLibraryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LibraryId) {
		query["libraryId"] = request.LibraryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLibrary"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/delete"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLibraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除文档库
//
// @param request - DeleteLibraryRequest
//
// @return DeleteLibraryResponse
func (client *Client) DeleteLibrary(workspaceId *string, request *DeleteLibraryRequest) (_result *DeleteLibraryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLibraryResponse{}
	_body, _err := client.DeleteLibraryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 端到端实时对话
//
// @param request - EndToEndRealTimeDialogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EndToEndRealTimeDialogResponse
func (client *Client) EndToEndRealTimeDialogWithOptions(workspaceId *string, request *EndToEndRealTimeDialogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EndToEndRealTimeDialogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AsrModelId) {
		query["asrModelId"] = request.AsrModelId
	}

	if !dara.IsNil(request.InputFormat) {
		query["inputFormat"] = request.InputFormat
	}

	if !dara.IsNil(request.OutputFormat) {
		query["outputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.PitchRate) {
		query["pitchRate"] = request.PitchRate
	}

	if !dara.IsNil(request.SampleRate) {
		query["sampleRate"] = request.SampleRate
	}

	if !dara.IsNil(request.SpeechRate) {
		query["speechRate"] = request.SpeechRate
	}

	if !dara.IsNil(request.TtsModelId) {
		query["ttsModelId"] = request.TtsModelId
	}

	if !dara.IsNil(request.VoiceCode) {
		query["voiceCode"] = request.VoiceCode
	}

	if !dara.IsNil(request.Volume) {
		query["volume"] = request.Volume
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EndToEndRealTimeDialog"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/ws/realtime/dialog"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EndToEndRealTimeDialogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 端到端实时对话
//
// @param request - EndToEndRealTimeDialogRequest
//
// @return EndToEndRealTimeDialogResponse
func (client *Client) EndToEndRealTimeDialog(workspaceId *string, request *EndToEndRealTimeDialogRequest) (_result *EndToEndRealTimeDialogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EndToEndRealTimeDialogResponse{}
	_body, _err := client.EndToEndRealTimeDialogWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 中断任务
//
// @param request - EvictTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EvictTaskResponse
func (client *Client) EvictTaskWithOptions(workspaceId *string, request *EvictTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EvictTaskResponse, _err error) {
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
		Action:      dara.String("EvictTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/evict"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EvictTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 中断任务
//
// @param request - EvictTaskRequest
//
// @return EvictTaskResponse
func (client *Client) EvictTask(workspaceId *string, request *EvictTaskRequest) (_result *EvictTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EvictTaskResponse{}
	_body, _err := client.EvictTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据文档解析问答QA
//
// @param request - GenDocQaResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenDocQaResultResponse
func (client *Client) GenDocQaResultWithOptions(workspaceId *string, request *GenDocQaResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenDocQaResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["docId"] = request.DocId
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenDocQaResult"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/qa/parse"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenDocQaResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据文档解析问答QA
//
// @param request - GenDocQaResultRequest
//
// @return GenDocQaResultResponse
func (client *Client) GenDocQaResult(workspaceId *string, request *GenDocQaResultRequest) (_result *GenDocQaResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenDocQaResultResponse{}
	_body, _err := client.GenDocQaResultWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取app配置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppConfigResponse
func (client *Client) GetAppConfigWithOptions(workspaceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAppConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppConfig"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/app/config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取app配置
//
// @return GetAppConfigResponse
func (client *Client) GetAppConfig(workspaceId *string) (_result *GetAppConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAppConfigResponse{}
	_body, _err := client.GetAppConfigWithOptions(workspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取问答结果
//
// @param request - GetChatQuestionRespRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatQuestionRespResponse
func (client *Client) GetChatQuestionRespWithOptions(workspaceId *string, request *GetChatQuestionRespRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetChatQuestionRespResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BatchId) {
		body["batchId"] = request.BatchId
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatQuestionResp"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/chat/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatQuestionRespResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取问答结果
//
// @param request - GetChatQuestionRespRequest
//
// @return GetChatQuestionRespResponse
func (client *Client) GetChatQuestionResp(workspaceId *string, request *GetChatQuestionRespRequest) (_result *GetChatQuestionRespResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetChatQuestionRespResponse{}
	_body, _err := client.GetChatQuestionRespWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取外呼会话分析结果
//
// @param request - GetDialogAnalysisResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDialogAnalysisResultResponse
func (client *Client) GetDialogAnalysisResultWithOptions(workspaceId *string, request *GetDialogAnalysisResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDialogAnalysisResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Asc) {
		body["asc"] = request.Asc
	}

	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.SessionIds) {
		body["sessionIds"] = request.SessionIds
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.UseUrl) {
		body["useUrl"] = request.UseUrl
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDialogAnalysisResult"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/dialog/analysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDialogAnalysisResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取外呼会话分析结果
//
// @param request - GetDialogAnalysisResultRequest
//
// @return GetDialogAnalysisResultResponse
func (client *Client) GetDialogAnalysisResult(workspaceId *string, request *GetDialogAnalysisResultRequest) (_result *GetDialogAnalysisResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDialogAnalysisResultResponse{}
	_body, _err := client.GetDialogAnalysisResultWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异步任务的结果
//
// @param request - GetDialogDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDialogDetailResponse
func (client *Client) GetDialogDetailWithOptions(workspaceId *string, request *GetDialogDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDialogDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDialogDetail"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/dialog/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDialogDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异步任务的结果
//
// @param request - GetDialogDetailRequest
//
// @return GetDialogDetailResponse
func (client *Client) GetDialogDetail(workspaceId *string, request *GetDialogDetailRequest) (_result *GetDialogDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDialogDetailResponse{}
	_body, _err := client.GetDialogDetailWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询会话日志
//
// @param request - GetDialogLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDialogLogResponse
func (client *Client) GetDialogLogWithOptions(workspaceId *string, request *GetDialogLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDialogLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDialogLog"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/dialog/log"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDialogLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会话日志
//
// @param request - GetDialogLogRequest
//
// @return GetDialogLogResponse
func (client *Client) GetDialogLog(workspaceId *string, request *GetDialogLogRequest) (_result *GetDialogLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDialogLogResponse{}
	_body, _err := client.GetDialogLogWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档的chunk列表
//
// @param request - GetDocumentChunkListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentChunkListResponse
func (client *Client) GetDocumentChunkListWithOptions(workspaceId *string, request *GetDocumentChunkListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocumentChunkListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChunkIdList) {
		body["chunkIdList"] = request.ChunkIdList
	}

	if !dara.IsNil(request.DocId) {
		body["docId"] = request.DocId
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.Order) {
		body["order"] = request.Order
	}

	if !dara.IsNil(request.OrderBy) {
		body["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchQuery) {
		body["searchQuery"] = request.SearchQuery
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentChunkList"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/getDocumentChunk"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentChunkListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档的chunk列表
//
// @param request - GetDocumentChunkListRequest
//
// @return GetDocumentChunkListResponse
func (client *Client) GetDocumentChunkList(workspaceId *string, request *GetDocumentChunkListRequest) (_result *GetDocumentChunkListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocumentChunkListResponse{}
	_body, _err := client.GetDocumentChunkListWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询文档库的文档列表
//
// @param request - GetDocumentListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentListResponse
func (client *Client) GetDocumentListWithOptions(workspaceId *string, request *GetDocumentListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocumentListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LibraryId) {
		query["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentList"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/listDocument"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询文档库的文档列表
//
// @param request - GetDocumentListRequest
//
// @return GetDocumentListResponse
func (client *Client) GetDocumentList(workspaceId *string, request *GetDocumentListRequest) (_result *GetDocumentListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocumentListResponse{}
	_body, _err := client.GetDocumentListWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档URL
//
// @param request - GetDocumentUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentUrlResponse
func (client *Client) GetDocumentUrlWithOptions(workspaceId *string, request *GetDocumentUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocumentUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DocumentId) {
		query["documentId"] = request.DocumentId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentUrl"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/document/url"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档URL
//
// @param request - GetDocumentUrlRequest
//
// @return GetDocumentUrlResponse
func (client *Client) GetDocumentUrl(workspaceId *string, request *GetDocumentUrlRequest) (_result *GetDocumentUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocumentUrlResponse{}
	_body, _err := client.GetDocumentUrlWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 带条件的分页查询文档库的文档列表
//
// @param request - GetFilterDocumentListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFilterDocumentListResponse
func (client *Client) GetFilterDocumentListWithOptions(workspaceId *string, request *GetFilterDocumentListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFilterDocumentListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.And) {
		body["and"] = request.And
	}

	if !dara.IsNil(request.DocIdList) {
		body["docIdList"] = request.DocIdList
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.Or) {
		body["or"] = request.Or
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFilterDocumentList"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/filterDocument"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFilterDocumentListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 带条件的分页查询文档库的文档列表
//
// @param request - GetFilterDocumentListRequest
//
// @return GetFilterDocumentListResponse
func (client *Client) GetFilterDocumentList(workspaceId *string, request *GetFilterDocumentListRequest) (_result *GetFilterDocumentListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFilterDocumentListResponse{}
	_body, _err := client.GetFilterDocumentListWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询文档库列表
//
// @param request - GetHistoryListByBizTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHistoryListByBizTypeResponse
func (client *Client) GetHistoryListByBizTypeWithOptions(workspaceId *string, request *GetHistoryListByBizTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHistoryListByBizTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["bizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["bizType"] = request.BizType
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHistoryListByBizType"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/history/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHistoryListByBizTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询文档库列表
//
// @param request - GetHistoryListByBizTypeRequest
//
// @return GetHistoryListByBizTypeResponse
func (client *Client) GetHistoryListByBizType(workspaceId *string, request *GetHistoryListByBizTypeRequest) (_result *GetHistoryListByBizTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHistoryListByBizTypeResponse{}
	_body, _err := client.GetHistoryListByBizTypeWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取检测结果
//
// @param request - GetImageDetectionTaskResultRequest
//
// @param headers - GetImageDetectionTaskResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageDetectionTaskResultResponse
func (client *Client) GetImageDetectionTaskResultWithOptions(workspaceId *string, request *GetImageDetectionTaskResultRequest, headers *GetImageDetectionTaskResultHeaders, runtime *dara.RuntimeOptions) (_result *GetImageDetectionTaskResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XLoadTest) {
		realHeaders["X-Load-Test"] = dara.String(dara.Stringify(dara.BoolValue(headers.XLoadTest)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageDetectionTaskResult"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/imageDetect/task/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageDetectionTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取检测结果
//
// @param request - GetImageDetectionTaskResultRequest
//
// @return GetImageDetectionTaskResultResponse
func (client *Client) GetImageDetectionTaskResult(workspaceId *string, request *GetImageDetectionTaskResultRequest) (_result *GetImageDetectionTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetImageDetectionTaskResultHeaders{}
	_result = &GetImageDetectionTaskResultResponse{}
	_body, _err := client.GetImageDetectionTaskResultWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档库配置详情
//
// @param request - GetLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLibraryResponse
func (client *Client) GetLibraryWithOptions(workspaceId *string, request *GetLibraryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLibraryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LibraryId) {
		query["libraryId"] = request.LibraryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLibrary"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/get"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLibraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档库配置详情
//
// @param request - GetLibraryRequest
//
// @return GetLibraryResponse
func (client *Client) GetLibrary(workspaceId *string, request *GetLibraryRequest) (_result *GetLibraryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLibraryResponse{}
	_body, _err := client.GetLibraryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询文档库列表
//
// @param request - GetLibraryListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLibraryListResponse
func (client *Client) GetLibraryListWithOptions(workspaceId *string, request *GetLibraryListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLibraryListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLibraryList"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLibraryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询文档库列表
//
// @param request - GetLibraryListRequest
//
// @return GetLibraryListResponse
func (client *Client) GetLibraryList(workspaceId *string, request *GetLibraryListRequest) (_result *GetLibraryListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLibraryListResponse{}
	_body, _err := client.GetLibraryListWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取解析结果
//
// @param request - GetParseResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetParseResultResponse
func (client *Client) GetParseResultWithOptions(workspaceId *string, request *GetParseResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetParseResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["docId"] = request.DocId
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.UseUrlResult) {
		body["useUrlResult"] = request.UseUrlResult
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetParseResult"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/document/getParseResult"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetParseResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取解析结果
//
// @param request - GetParseResultRequest
//
// @return GetParseResultResponse
func (client *Client) GetParseResult(workspaceId *string, request *GetParseResultRequest) (_result *GetParseResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetParseResultResponse{}
	_body, _err := client.GetParseResultWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异步任务的结果
//
// @param request - GetQualityCheckTaskResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityCheckTaskResultResponse
func (client *Client) GetQualityCheckTaskResultWithOptions(workspaceId *string, request *GetQualityCheckTaskResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetQualityCheckTaskResultResponse, _err error) {
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
		Action:      dara.String("GetQualityCheckTaskResult"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/qualitycheck/task/query"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityCheckTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异步任务的结果
//
// @param request - GetQualityCheckTaskResultRequest
//
// @return GetQualityCheckTaskResultResponse
func (client *Client) GetQualityCheckTaskResult(workspaceId *string, request *GetQualityCheckTaskResultRequest) (_result *GetQualityCheckTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQualityCheckTaskResultResponse{}
	_body, _err := client.GetQualityCheckTaskResultWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取财报总结任务结果
//
// @param request - GetSummaryTaskResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSummaryTaskResultResponse
func (client *Client) GetSummaryTaskResultWithOptions(workspaceId *string, request *GetSummaryTaskResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSummaryTaskResultResponse, _err error) {
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
		Action:      dara.String("GetSummaryTaskResult"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/summary/result"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSummaryTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取财报总结任务结果
//
// @param request - GetSummaryTaskResultRequest
//
// @return GetSummaryTaskResultResponse
func (client *Client) GetSummaryTaskResult(workspaceId *string, request *GetSummaryTaskResultRequest) (_result *GetSummaryTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSummaryTaskResultResponse{}
	_body, _err := client.GetSummaryTaskResultWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异步任务结果
//
// @param request - GetTaskResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResultResponse
func (client *Client) GetTaskResultWithOptions(workspaceId *string, request *GetTaskResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskResultResponse, _err error) {
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
		Action:      dara.String("GetTaskResult"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/result"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异步任务结果
//
// @param request - GetTaskResultRequest
//
// @return GetTaskResultResponse
func (client *Client) GetTaskResult(workspaceId *string, request *GetTaskResultRequest) (_result *GetTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskResultResponse{}
	_body, _err := client.GetTaskResultWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取财报总结任务结果
//
// @param request - GetTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskStatusResponse
func (client *Client) GetTaskStatusWithOptions(workspaceId *string, request *GetTaskStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskStatusResponse, _err error) {
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
		Action:      dara.String("GetTaskStatus"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取财报总结任务结果
//
// @param request - GetTaskStatusRequest
//
// @return GetTaskStatusResponse
func (client *Client) GetTaskStatus(workspaceId *string, request *GetTaskStatusRequest) (_result *GetTaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskStatusResponse{}
	_body, _err := client.GetTaskStatusWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取视频生成任务结果
//
// @param request - GetVideoCreationTaskResultRequest
//
// @param headers - GetVideoCreationTaskResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoCreationTaskResultResponse
func (client *Client) GetVideoCreationTaskResultWithOptions(workspaceId *string, request *GetVideoCreationTaskResultRequest, headers *GetVideoCreationTaskResultHeaders, runtime *dara.RuntimeOptions) (_result *GetVideoCreationTaskResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XLoadTest) {
		realHeaders["X-Load-Test"] = dara.String(dara.Stringify(dara.BoolValue(headers.XLoadTest)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoCreationTaskResult"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/videoCreation/task/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoCreationTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取视频生成任务结果
//
// @param request - GetVideoCreationTaskResultRequest
//
// @return GetVideoCreationTaskResultResponse
func (client *Client) GetVideoCreationTaskResult(workspaceId *string, request *GetVideoCreationTaskResultRequest) (_result *GetVideoCreationTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetVideoCreationTaskResultHeaders{}
	_result = &GetVideoCreationTaskResultResponse{}
	_body, _err := client.GetVideoCreationTaskResultWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 插件调试接口
//
// @param request - InvokePluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokePluginResponse
func (client *Client) InvokePluginWithOptions(workspaceId *string, request *InvokePluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InvokePluginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	if !dara.IsNil(request.PluginId) {
		body["pluginId"] = request.PluginId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokePlugin"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/plugin/invoke"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InvokePluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 插件调试接口
//
// @param request - InvokePluginRequest
//
// @return InvokePluginResponse
func (client *Client) InvokePlugin(workspaceId *string, request *InvokePluginRequest) (_result *InvokePluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InvokePluginResponse{}
	_body, _err := client.InvokePluginWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档预览
//
// @param request - PreviewDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreviewDocumentResponse
func (client *Client) PreviewDocumentWithOptions(workspaceId *string, request *PreviewDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PreviewDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DocumentId) {
		query["documentId"] = request.DocumentId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreviewDocument"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/document/preview"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PreviewDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档预览
//
// @param request - PreviewDocumentRequest
//
// @return PreviewDocumentResponse
func (client *Client) PreviewDocument(workspaceId *string, request *PreviewDocumentRequest) (_result *PreviewDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PreviewDocumentResponse{}
	_body, _err := client.PreviewDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重新索引
//
// @param request - ReIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReIndexResponse
func (client *Client) ReIndexWithOptions(workspaceId *string, request *ReIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReIndexResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DocumentId) {
		query["documentId"] = request.DocumentId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReIndex"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/document/reIndex"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重新索引
//
// @param request - ReIndexRequest
//
// @return ReIndexResponse
func (client *Client) ReIndex(workspaceId *string, request *ReIndexRequest) (_result *ReIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReIndexResponse{}
	_body, _err := client.ReIndexWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实时对话
//
// @param request - RealTimeDialogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RealTimeDialogResponse
func (client *Client) RealTimeDialogWithSSE(workspaceId *string, request *RealTimeDialogRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RealTimeDialogResponse, _yieldErr chan error) {
	defer close(_yield)
	client.realTimeDialogWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, request, headers, runtime)
	return
}

// Summary:
//
// 实时对话
//
// @param request - RealTimeDialogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RealTimeDialogResponse
func (client *Client) RealTimeDialogWithOptions(workspaceId *string, request *RealTimeDialogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RealTimeDialogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Analysis) {
		body["analysis"] = request.Analysis
	}

	if !dara.IsNil(request.BizType) {
		body["bizType"] = request.BizType
	}

	if !dara.IsNil(request.ConversationModel) {
		body["conversationModel"] = request.ConversationModel
	}

	if !dara.IsNil(request.DialogMemoryTurns) {
		body["dialogMemoryTurns"] = request.DialogMemoryTurns
	}

	if !dara.IsNil(request.MetaData) {
		body["metaData"] = request.MetaData
	}

	if !dara.IsNil(request.OpType) {
		body["opType"] = request.OpType
	}

	if !dara.IsNil(request.Recommend) {
		body["recommend"] = request.Recommend
	}

	if !dara.IsNil(request.ScriptContentPlayed) {
		body["scriptContentPlayed"] = request.ScriptContentPlayed
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.UserVad) {
		body["userVad"] = request.UserVad
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RealTimeDialog"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/realtime/dialog/chat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RealTimeDialogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实时对话
//
// @param request - RealTimeDialogRequest
//
// @return RealTimeDialogResponse
func (client *Client) RealTimeDialog(workspaceId *string, request *RealTimeDialogRequest) (_result *RealTimeDialogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RealTimeDialogResponse{}
	_body, _err := client.RealTimeDialogWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实时会话辅助
//
// @param request - RealtimeDialogAssistRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RealtimeDialogAssistResponse
func (client *Client) RealtimeDialogAssistWithOptions(workspaceId *string, request *RealtimeDialogAssistRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RealtimeDialogAssistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Analysis) {
		body["analysis"] = request.Analysis
	}

	if !dara.IsNil(request.BizType) {
		body["bizType"] = request.BizType
	}

	if !dara.IsNil(request.ConversationModel) {
		body["conversationModel"] = request.ConversationModel
	}

	if !dara.IsNil(request.DialogMemoryTurns) {
		body["dialogMemoryTurns"] = request.DialogMemoryTurns
	}

	if !dara.IsNil(request.HangUpDialog) {
		body["hangUpDialog"] = request.HangUpDialog
	}

	if !dara.IsNil(request.MetaData) {
		body["metaData"] = request.MetaData
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RealtimeDialogAssist"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/realtime/dialog/assist"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RealtimeDialogAssistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实时会话辅助
//
// @param request - RealtimeDialogAssistRequest
//
// @return RealtimeDialogAssistResponse
func (client *Client) RealtimeDialogAssist(workspaceId *string, request *RealtimeDialogAssistRequest) (_result *RealtimeDialogAssistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RealtimeDialogAssistResponse{}
	_body, _err := client.RealtimeDialogAssistWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重建任务
//
// @param request - RebuildTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebuildTaskResponse
func (client *Client) RebuildTaskWithOptions(workspaceId *string, request *RebuildTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RebuildTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskIds) {
		body["taskIds"] = request.TaskIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebuildTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/rebuild"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RebuildTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重建任务
//
// @param request - RebuildTaskRequest
//
// @return RebuildTaskResponse
func (client *Client) RebuildTask(workspaceId *string, request *RebuildTaskRequest) (_result *RebuildTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RebuildTaskResponse{}
	_body, _err := client.RebuildTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档召回。
//
// @param request - RecallDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecallDocumentResponse
func (client *Client) RecallDocumentWithOptions(workspaceId *string, request *RecallDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RecallDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Filters) {
		body["filters"] = request.Filters
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.Rearrangement) {
		body["rearrangement"] = request.Rearrangement
	}

	if !dara.IsNil(request.TopK) {
		body["topK"] = request.TopK
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecallDocument"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/recallDocument"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RecallDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档召回。
//
// @param request - RecallDocumentRequest
//
// @return RecallDocumentResponse
func (client *Client) RecallDocument(workspaceId *string, request *RecallDocumentRequest) (_result *RecallDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecallDocumentResponse{}
	_body, _err := client.RecallDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图识别
//
// @param request - RecognizeIntentionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeIntentionResponse
func (client *Client) RecognizeIntentionWithOptions(workspaceId *string, request *RecognizeIntentionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RecognizeIntentionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Analysis) {
		body["analysis"] = request.Analysis
	}

	if !dara.IsNil(request.BizType) {
		body["bizType"] = request.BizType
	}

	if !dara.IsNil(request.Conversation) {
		body["conversation"] = request.Conversation
	}

	if !dara.IsNil(request.GlobalIntentionList) {
		body["globalIntentionList"] = request.GlobalIntentionList
	}

	if !dara.IsNil(request.HierarchicalIntentionList) {
		body["hierarchicalIntentionList"] = request.HierarchicalIntentionList
	}

	if !dara.IsNil(request.IntentionDomainCode) {
		body["intentionDomainCode"] = request.IntentionDomainCode
	}

	if !dara.IsNil(request.IntentionList) {
		body["intentionList"] = request.IntentionList
	}

	if !dara.IsNil(request.OpType) {
		body["opType"] = request.OpType
	}

	if !dara.IsNil(request.Recommend) {
		body["recommend"] = request.Recommend
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecognizeIntention"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/recog/intent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RecognizeIntentionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图识别
//
// @param request - RecognizeIntentionRequest
//
// @return RecognizeIntentionResponse
func (client *Client) RecognizeIntention(workspaceId *string, request *RecognizeIntentionRequest) (_result *RecognizeIntentionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecognizeIntentionResponse{}
	_body, _err := client.RecognizeIntentionWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 运行智能体
//
// @param request - RunAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunAgentResponse
func (client *Client) RunAgentWithSSE(workspaceId *string, request *RunAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunAgentResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runAgentWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, request, headers, runtime)
	return
}

// Summary:
//
// 运行智能体
//
// @param request - RunAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunAgentResponse
func (client *Client) RunAgentWithOptions(workspaceId *string, request *RunAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BotId) {
		body["botId"] = request.BotId
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.ThreadId) {
		body["threadId"] = request.ThreadId
	}

	if !dara.IsNil(request.UseDraft) {
		body["useDraft"] = request.UseDraft
	}

	if !dara.IsNil(request.UserContent) {
		body["userContent"] = request.UserContent
	}

	if !dara.IsNil(request.UserInputs) {
		body["userInputs"] = request.UserInputs
	}

	if !dara.IsNil(request.VersionId) {
		body["versionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunAgent"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/bot/thread/run"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RunAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运行智能体
//
// @param request - RunAgentRequest
//
// @return RunAgentResponse
func (client *Client) RunAgent(workspaceId *string, request *RunAgentRequest) (_result *RunAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunAgentResponse{}
	_body, _err := client.RunAgentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunChatResultGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunChatResultGenerationResponse
func (client *Client) RunChatResultGenerationWithSSE(workspaceId *string, request *RunChatResultGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunChatResultGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runChatResultGenerationWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, request, headers, runtime)
	return
}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunChatResultGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunChatResultGenerationResponse
func (client *Client) RunChatResultGenerationWithOptions(workspaceId *string, request *RunChatResultGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunChatResultGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InferenceParameters) {
		body["inferenceParameters"] = request.InferenceParameters
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.Tools) {
		body["tools"] = request.Tools
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunChatResultGeneration"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/run/chat/generation"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RunChatResultGenerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunChatResultGenerationRequest
//
// @return RunChatResultGenerationResponse
func (client *Client) RunChatResultGeneration(workspaceId *string, request *RunChatResultGenerationRequest) (_result *RunChatResultGenerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunChatResultGenerationResponse{}
	_body, _err := client.RunChatResultGenerationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 流式获取外呼会话分析结果
//
// @param request - RunDialogAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDialogAnalysisResponse
func (client *Client) RunDialogAnalysisWithSSE(workspaceId *string, request *RunDialogAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunDialogAnalysisResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDialogAnalysisWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, request, headers, runtime)
	return
}

// Summary:
//
// 流式获取外呼会话分析结果
//
// @param request - RunDialogAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDialogAnalysisResponse
func (client *Client) RunDialogAnalysisWithOptions(workspaceId *string, request *RunDialogAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunDialogAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDialogAnalysis"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/dialog/stream/analysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RunDialogAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 流式获取外呼会话分析结果
//
// @param request - RunDialogAnalysisRequest
//
// @return RunDialogAnalysisResponse
func (client *Client) RunDialogAnalysis(workspaceId *string, request *RunDialogAnalysisRequest) (_result *RunDialogAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunDialogAnalysisResponse{}
	_body, _err := client.RunDialogAnalysisWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunLibraryChatGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunLibraryChatGenerationResponse
func (client *Client) RunLibraryChatGenerationWithSSE(workspaceId *string, request *RunLibraryChatGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunLibraryChatGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runLibraryChatGenerationWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, request, headers, runtime)
	return
}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunLibraryChatGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunLibraryChatGenerationResponse
func (client *Client) RunLibraryChatGenerationWithOptions(workspaceId *string, request *RunLibraryChatGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunLibraryChatGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocIdList) {
		body["docIdList"] = request.DocIdList
	}

	if !dara.IsNil(request.EnableFollowUp) {
		body["enableFollowUp"] = request.EnableFollowUp
	}

	if !dara.IsNil(request.EnableMultiQuery) {
		body["enableMultiQuery"] = request.EnableMultiQuery
	}

	if !dara.IsNil(request.EnableOpenQa) {
		body["enableOpenQa"] = request.EnableOpenQa
	}

	if !dara.IsNil(request.FollowUpLlm) {
		body["followUpLlm"] = request.FollowUpLlm
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.LlmType) {
		body["llmType"] = request.LlmType
	}

	if !dara.IsNil(request.MultiQueryLlm) {
		body["multiQueryLlm"] = request.MultiQueryLlm
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.QueryCriteria) {
		body["queryCriteria"] = request.QueryCriteria
	}

	if !dara.IsNil(request.RerankType) {
		body["rerankType"] = request.RerankType
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.SubQueryList) {
		body["subQueryList"] = request.SubQueryList
	}

	if !dara.IsNil(request.TextSearchParameter) {
		body["textSearchParameter"] = request.TextSearchParameter
	}

	if !dara.IsNil(request.TopK) {
		body["topK"] = request.TopK
	}

	if !dara.IsNil(request.VectorSearchParameter) {
		body["vectorSearchParameter"] = request.VectorSearchParameter
	}

	if !dara.IsNil(request.WithDocumentReference) {
		body["withDocumentReference"] = request.WithDocumentReference
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunLibraryChatGeneration"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/run/library/chat/generation"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RunLibraryChatGenerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunLibraryChatGenerationRequest
//
// @return RunLibraryChatGenerationResponse
func (client *Client) RunLibraryChatGeneration(workspaceId *string, request *RunLibraryChatGenerationRequest) (_result *RunLibraryChatGenerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunLibraryChatGenerationResponse{}
	_body, _err := client.RunLibraryChatGenerationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交问题列表
//
// @param request - SubmitChatQuestionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitChatQuestionResponse
func (client *Client) SubmitChatQuestionWithOptions(workspaceId *string, request *SubmitChatQuestionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitChatQuestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GmtService) {
		body["gmtService"] = request.GmtService
	}

	if !dara.IsNil(request.LiveScriptContent) {
		body["liveScriptContent"] = request.LiveScriptContent
	}

	if !dara.IsNil(request.OpenSmallTalk) {
		body["openSmallTalk"] = request.OpenSmallTalk
	}

	if !dara.IsNil(request.QuestionList) {
		body["questionList"] = request.QuestionList
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitChatQuestion"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/chat/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitChatQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交问题列表
//
// @param request - SubmitChatQuestionRequest
//
// @return SubmitChatQuestionResponse
func (client *Client) SubmitChatQuestion(workspaceId *string, request *SubmitChatQuestionRequest) (_result *SubmitChatQuestionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitChatQuestionResponse{}
	_body, _err := client.SubmitChatQuestionWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新文档
//
// @param request - UpdateDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDocumentResponse
func (client *Client) UpdateDocumentWithOptions(workspaceId *string, request *UpdateDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["docId"] = request.DocId
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.Meta) {
		body["meta"] = request.Meta
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDocument"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/document/updateDocument"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新文档
//
// @param request - UpdateDocumentRequest
//
// @return UpdateDocumentResponse
func (client *Client) UpdateDocument(workspaceId *string, request *UpdateDocumentRequest) (_result *UpdateDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDocumentResponse{}
	_body, _err := client.UpdateDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新文档的chunk
//
// @param request - UpdateDocumentChunkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDocumentChunkResponse
func (client *Client) UpdateDocumentChunkWithOptions(workspaceId *string, request *UpdateDocumentChunkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDocumentChunkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Chunks) {
		body["chunks"] = request.Chunks
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDocumentChunk"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/updateDocumentChunk"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDocumentChunkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新文档的chunk
//
// @param request - UpdateDocumentChunkRequest
//
// @return UpdateDocumentChunkResponse
func (client *Client) UpdateDocumentChunk(workspaceId *string, request *UpdateDocumentChunkRequest) (_result *UpdateDocumentChunkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDocumentChunkResponse{}
	_body, _err := client.UpdateDocumentChunkWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新文档库配置
//
// @param request - UpdateLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLibraryResponse
func (client *Client) UpdateLibraryWithOptions(workspaceId *string, request *UpdateLibraryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLibraryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.IndexSetting) {
		body["indexSetting"] = request.IndexSetting
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.LibraryName) {
		body["libraryName"] = request.LibraryName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLibrary"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/update"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLibraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新文档库配置
//
// @param request - UpdateLibraryRequest
//
// @return UpdateLibraryResponse
func (client *Client) UpdateLibrary(workspaceId *string, request *UpdateLibraryRequest) (_result *UpdateLibraryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLibraryResponse{}
	_body, _err := client.UpdateLibraryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新QA问答库
//
// @param request - UpdateQaLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQaLibraryResponse
func (client *Client) UpdateQaLibraryWithOptions(workspaceId *string, request *UpdateQaLibraryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateQaLibraryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ParseQaResults) {
		body["parseQaResults"] = request.ParseQaResults
	}

	if !dara.IsNil(request.QaLibraryId) {
		body["qaLibraryId"] = request.QaLibraryId
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateQaLibrary"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/qa/upload"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQaLibraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新QA问答库
//
// @param request - UpdateQaLibraryRequest
//
// @return UpdateQaLibraryResponse
func (client *Client) UpdateQaLibrary(workspaceId *string, request *UpdateQaLibraryRequest) (_result *UpdateQaLibraryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateQaLibraryResponse{}
	_body, _err := client.UpdateQaLibraryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传文档到文档库
//
// @param request - UploadDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDocumentResponse
func (client *Client) UploadDocumentWithOptions(workspaceId *string, request *UploadDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UploadDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		body["data"] = request.Data
	}

	if !dara.IsNil(request.FileName) {
		body["fileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		body["fileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadDocument"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/document/upload"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传文档到文档库
//
// @param request - UploadDocumentRequest
//
// @return UploadDocumentResponse
func (client *Client) UploadDocument(workspaceId *string, request *UploadDocumentRequest) (_result *UploadDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadDocumentResponse{}
	_body, _err := client.UploadDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadDocumentAdvance(workspaceId *string, request *UploadDocumentAdvanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UploadDocumentResponse, _err error) {
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
		"Product":  dara.String("DianJin"),
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
	uploadDocumentReq := &UploadDocumentRequest{}
	openapiutil.Convert(request, uploadDocumentReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
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
		uploadDocumentReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	uploadDocumentResp, _err := client.UploadDocumentWithOptions(workspaceId, uploadDocumentReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = uploadDocumentResp
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

func (client *Client) realTimeDialogWithSSE_opYieldFunc(_yield chan *RealTimeDialogResponse, _yieldErr chan error, workspaceId *string, request *RealTimeDialogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Analysis) {
		body["analysis"] = request.Analysis
	}

	if !dara.IsNil(request.BizType) {
		body["bizType"] = request.BizType
	}

	if !dara.IsNil(request.ConversationModel) {
		body["conversationModel"] = request.ConversationModel
	}

	if !dara.IsNil(request.DialogMemoryTurns) {
		body["dialogMemoryTurns"] = request.DialogMemoryTurns
	}

	if !dara.IsNil(request.MetaData) {
		body["metaData"] = request.MetaData
	}

	if !dara.IsNil(request.OpType) {
		body["opType"] = request.OpType
	}

	if !dara.IsNil(request.Recommend) {
		body["recommend"] = request.Recommend
	}

	if !dara.IsNil(request.ScriptContentPlayed) {
		body["scriptContentPlayed"] = request.ScriptContentPlayed
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.UserVad) {
		body["userVad"] = request.UserVad
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RealTimeDialog"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/realtime/dialog/chat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runAgentWithSSE_opYieldFunc(_yield chan *RunAgentResponse, _yieldErr chan error, workspaceId *string, request *RunAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BotId) {
		body["botId"] = request.BotId
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.ThreadId) {
		body["threadId"] = request.ThreadId
	}

	if !dara.IsNil(request.UseDraft) {
		body["useDraft"] = request.UseDraft
	}

	if !dara.IsNil(request.UserContent) {
		body["userContent"] = request.UserContent
	}

	if !dara.IsNil(request.UserInputs) {
		body["userInputs"] = request.UserInputs
	}

	if !dara.IsNil(request.VersionId) {
		body["versionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunAgent"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/bot/thread/run"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runChatResultGenerationWithSSE_opYieldFunc(_yield chan *RunChatResultGenerationResponse, _yieldErr chan error, workspaceId *string, request *RunChatResultGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InferenceParameters) {
		body["inferenceParameters"] = request.InferenceParameters
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.Tools) {
		body["tools"] = request.Tools
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunChatResultGeneration"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/run/chat/generation"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runDialogAnalysisWithSSE_opYieldFunc(_yield chan *RunDialogAnalysisResponse, _yieldErr chan error, workspaceId *string, request *RunDialogAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDialogAnalysis"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/dialog/stream/analysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runLibraryChatGenerationWithSSE_opYieldFunc(_yield chan *RunLibraryChatGenerationResponse, _yieldErr chan error, workspaceId *string, request *RunLibraryChatGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocIdList) {
		body["docIdList"] = request.DocIdList
	}

	if !dara.IsNil(request.EnableFollowUp) {
		body["enableFollowUp"] = request.EnableFollowUp
	}

	if !dara.IsNil(request.EnableMultiQuery) {
		body["enableMultiQuery"] = request.EnableMultiQuery
	}

	if !dara.IsNil(request.EnableOpenQa) {
		body["enableOpenQa"] = request.EnableOpenQa
	}

	if !dara.IsNil(request.FollowUpLlm) {
		body["followUpLlm"] = request.FollowUpLlm
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.LlmType) {
		body["llmType"] = request.LlmType
	}

	if !dara.IsNil(request.MultiQueryLlm) {
		body["multiQueryLlm"] = request.MultiQueryLlm
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.QueryCriteria) {
		body["queryCriteria"] = request.QueryCriteria
	}

	if !dara.IsNil(request.RerankType) {
		body["rerankType"] = request.RerankType
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.SubQueryList) {
		body["subQueryList"] = request.SubQueryList
	}

	if !dara.IsNil(request.TextSearchParameter) {
		body["textSearchParameter"] = request.TextSearchParameter
	}

	if !dara.IsNil(request.TopK) {
		body["topK"] = request.TopK
	}

	if !dara.IsNil(request.VectorSearchParameter) {
		body["vectorSearchParameter"] = request.VectorSearchParameter
	}

	if !dara.IsNil(request.WithDocumentReference) {
		body["withDocumentReference"] = request.WithDocumentReference
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunLibraryChatGeneration"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/run/library/chat/generation"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}
