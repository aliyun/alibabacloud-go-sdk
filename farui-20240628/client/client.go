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
	client.Endpoint, _err = client.GetEndpoint(dara.String("farui"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 上传合同文件
//
// @param request - CreateTextFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTextFileResponse
func (client *Client) CreateTextFileWithOptions(WorkspaceId *string, request *CreateTextFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTextFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ContractId) {
		body["ContractId"] = request.ContractId
	}

	if !dara.IsNil(request.CreateTime) {
		body["CreateTime"] = request.CreateTime
	}

	if !dara.IsNil(request.TextFileName) {
		body["TextFileName"] = request.TextFileName
	}

	if !dara.IsNil(request.TextFileUrl) {
		body["TextFileUrl"] = request.TextFileUrl
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTextFile"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/data/textFile"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTextFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传合同文件
//
// @param request - CreateTextFileRequest
//
// @return CreateTextFileResponse
func (client *Client) CreateTextFile(WorkspaceId *string, request *CreateTextFileRequest) (_result *CreateTextFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTextFileResponse{}
	_body, _err := client.CreateTextFileWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTextFileAdvance(WorkspaceId *string, request *CreateTextFileAdvanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTextFileResponse, _err error) {
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
		"Product":  dara.String("FaRui"),
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
	createTextFileReq := &CreateTextFileRequest{}
	openapiutil.Convert(request, createTextFileReq)
	if !dara.IsNil(request.TextFileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.TextFileUrlObject,
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
		createTextFileReq.TextFileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	createTextFileResp, _err := client.CreateTextFileWithOptions(WorkspaceId, createTextFileReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = createTextFileResp
	return _result, _err
}

// Summary:
//
// 合同抽取
//
// @param tmpReq - RunContractExtractRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunContractExtractResponse
func (client *Client) RunContractExtractWithOptions(workspaceId *string, tmpReq *RunContractExtractRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunContractExtractResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunContractExtractShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FieldsToExtract) {
		request.FieldsToExtractShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FieldsToExtract, dara.String("fieldsToExtract"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.FieldsToExtractShrink) {
		body["fieldsToExtract"] = request.FieldsToExtractShrink
	}

	if !dara.IsNil(request.FileOssUrl) {
		body["fileOssUrl"] = request.FileOssUrl
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunContractExtract"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/pop/contract/extraction"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunContractExtractResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合同抽取
//
// @param request - RunContractExtractRequest
//
// @return RunContractExtractResponse
func (client *Client) RunContractExtract(workspaceId *string, request *RunContractExtractRequest) (_result *RunContractExtractResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunContractExtractResponse{}
	_body, _err := client.RunContractExtractWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成合同审查结果
//
// @param tmpReq - RunContractResultGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunContractResultGenerationResponse
func (client *Client) RunContractResultGenerationWithSSE(workspaceId *string, tmpReq *RunContractResultGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunContractResultGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runContractResultGenerationWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, tmpReq, headers, runtime)
	return
}

// Summary:
//
// 生成合同审查结果
//
// @param tmpReq - RunContractResultGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunContractResultGenerationResponse
func (client *Client) RunContractResultGenerationWithOptions(workspaceId *string, tmpReq *RunContractResultGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunContractResultGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunContractResultGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Assistant) {
		request.AssistantShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Assistant, dara.String("assistant"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.AssistantShrink) {
		body["assistant"] = request.AssistantShrink
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunContractResultGeneration"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/farui/contract/result/genarate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunContractResultGenerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成合同审查结果
//
// @param request - RunContractResultGenerationRequest
//
// @return RunContractResultGenerationResponse
func (client *Client) RunContractResultGeneration(workspaceId *string, request *RunContractResultGenerationRequest) (_result *RunContractResultGenerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunContractResultGenerationResponse{}
	_body, _err := client.RunContractResultGenerationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成合同审查规则
//
// @param tmpReq - RunContractRuleGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunContractRuleGenerationResponse
func (client *Client) RunContractRuleGenerationWithSSE(workspaceId *string, tmpReq *RunContractRuleGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunContractRuleGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runContractRuleGenerationWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, tmpReq, headers, runtime)
	return
}

// Summary:
//
// 生成合同审查规则
//
// @param tmpReq - RunContractRuleGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunContractRuleGenerationResponse
func (client *Client) RunContractRuleGenerationWithOptions(workspaceId *string, tmpReq *RunContractRuleGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunContractRuleGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunContractRuleGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Assistant) {
		request.AssistantShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Assistant, dara.String("assistant"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.AssistantShrink) {
		body["assistant"] = request.AssistantShrink
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunContractRuleGeneration"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/farui/contract/rule/genarate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunContractRuleGenerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成合同审查规则
//
// @param request - RunContractRuleGenerationRequest
//
// @return RunContractRuleGenerationResponse
func (client *Client) RunContractRuleGeneration(workspaceId *string, request *RunContractRuleGenerationRequest) (_result *RunContractRuleGenerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunContractRuleGenerationResponse{}
	_body, _err := client.RunContractRuleGenerationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 法律咨询
//
// @param tmpReq - RunLegalAdviceConsultationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunLegalAdviceConsultationResponse
func (client *Client) RunLegalAdviceConsultationWithSSE(workspaceId *string, tmpReq *RunLegalAdviceConsultationRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunLegalAdviceConsultationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runLegalAdviceConsultationWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, tmpReq, headers, runtime)
	return
}

// Summary:
//
// 法律咨询
//
// @param tmpReq - RunLegalAdviceConsultationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunLegalAdviceConsultationResponse
func (client *Client) RunLegalAdviceConsultationWithOptions(workspaceId *string, tmpReq *RunLegalAdviceConsultationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunLegalAdviceConsultationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunLegalAdviceConsultationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Assistant) {
		request.AssistantShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Assistant, dara.String("assistant"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("extra"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Thread) {
		request.ThreadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Thread, dara.String("thread"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.AssistantShrink) {
		body["assistant"] = request.AssistantShrink
	}

	if !dara.IsNil(request.ExtraShrink) {
		body["extra"] = request.ExtraShrink
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.ThreadShrink) {
		body["thread"] = request.ThreadShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunLegalAdviceConsultation"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/farui/legalAdvice/consult"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunLegalAdviceConsultationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 法律咨询
//
// @param request - RunLegalAdviceConsultationRequest
//
// @return RunLegalAdviceConsultationResponse
func (client *Client) RunLegalAdviceConsultation(workspaceId *string, request *RunLegalAdviceConsultationRequest) (_result *RunLegalAdviceConsultationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunLegalAdviceConsultationResponse{}
	_body, _err := client.RunLegalAdviceConsultationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 案例检索
//
// @param tmpReq - RunSearchCaseFullTextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSearchCaseFullTextResponse
func (client *Client) RunSearchCaseFullTextWithOptions(workspaceId *string, tmpReq *RunSearchCaseFullTextRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunSearchCaseFullTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunSearchCaseFullTextShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterCondition) {
		request.FilterConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterCondition, dara.String("filterCondition"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PageParam) {
		request.PageParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PageParam, dara.String("pageParam"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.QueryKeywords) {
		request.QueryKeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryKeywords, dara.String("queryKeywords"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SortKeyAndDirection) {
		request.SortKeyAndDirectionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SortKeyAndDirection, dara.String("sortKeyAndDirection"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Thread) {
		request.ThreadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Thread, dara.String("thread"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.FilterConditionShrink) {
		body["filterCondition"] = request.FilterConditionShrink
	}

	if !dara.IsNil(request.PageParamShrink) {
		body["pageParam"] = request.PageParamShrink
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.QueryKeywordsShrink) {
		body["queryKeywords"] = request.QueryKeywordsShrink
	}

	if !dara.IsNil(request.ReferLevel) {
		body["referLevel"] = request.ReferLevel
	}

	if !dara.IsNil(request.SortKeyAndDirectionShrink) {
		body["sortKeyAndDirection"] = request.SortKeyAndDirectionShrink
	}

	if !dara.IsNil(request.ThreadShrink) {
		body["thread"] = request.ThreadShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunSearchCaseFullText"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/farui/search/case/fulltext"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunSearchCaseFullTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 案例检索
//
// @param request - RunSearchCaseFullTextRequest
//
// @return RunSearchCaseFullTextResponse
func (client *Client) RunSearchCaseFullText(workspaceId *string, request *RunSearchCaseFullTextRequest) (_result *RunSearchCaseFullTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunSearchCaseFullTextResponse{}
	_body, _err := client.RunSearchCaseFullTextWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 法规搜索
//
// @param tmpReq - RunSearchLawQueryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSearchLawQueryResponse
func (client *Client) RunSearchLawQueryWithOptions(workspaceId *string, tmpReq *RunSearchLawQueryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunSearchLawQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunSearchLawQueryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterCondition) {
		request.FilterConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterCondition, dara.String("filterCondition"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PageParam) {
		request.PageParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PageParam, dara.String("pageParam"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.QueryKeywords) {
		request.QueryKeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryKeywords, dara.String("queryKeywords"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Thread) {
		request.ThreadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Thread, dara.String("thread"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.FilterConditionShrink) {
		body["filterCondition"] = request.FilterConditionShrink
	}

	if !dara.IsNil(request.PageParamShrink) {
		body["pageParam"] = request.PageParamShrink
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.QueryKeywordsShrink) {
		body["queryKeywords"] = request.QueryKeywordsShrink
	}

	if !dara.IsNil(request.ThreadShrink) {
		body["thread"] = request.ThreadShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunSearchLawQuery"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/farui/search/law/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunSearchLawQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 法规搜索
//
// @param request - RunSearchLawQueryRequest
//
// @return RunSearchLawQueryResponse
func (client *Client) RunSearchLawQuery(workspaceId *string, request *RunSearchLawQueryRequest) (_result *RunSearchLawQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunSearchLawQueryResponse{}
	_body, _err := client.RunSearchLawQueryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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

func (client *Client) runContractResultGenerationWithSSE_opYieldFunc(_yield chan *RunContractResultGenerationResponse, _yieldErr chan error, workspaceId *string, tmpReq *RunContractResultGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunContractResultGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Assistant) {
		request.AssistantShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Assistant, dara.String("assistant"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.AssistantShrink) {
		body["assistant"] = request.AssistantShrink
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunContractResultGeneration"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/farui/contract/result/genarate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
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

func (client *Client) runContractRuleGenerationWithSSE_opYieldFunc(_yield chan *RunContractRuleGenerationResponse, _yieldErr chan error, workspaceId *string, tmpReq *RunContractRuleGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunContractRuleGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Assistant) {
		request.AssistantShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Assistant, dara.String("assistant"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.AssistantShrink) {
		body["assistant"] = request.AssistantShrink
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunContractRuleGeneration"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/farui/contract/rule/genarate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
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

func (client *Client) runLegalAdviceConsultationWithSSE_opYieldFunc(_yield chan *RunLegalAdviceConsultationResponse, _yieldErr chan error, workspaceId *string, tmpReq *RunLegalAdviceConsultationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunLegalAdviceConsultationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Assistant) {
		request.AssistantShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Assistant, dara.String("assistant"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("extra"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Thread) {
		request.ThreadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Thread, dara.String("thread"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.AssistantShrink) {
		body["assistant"] = request.AssistantShrink
	}

	if !dara.IsNil(request.ExtraShrink) {
		body["extra"] = request.ExtraShrink
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.ThreadShrink) {
		body["thread"] = request.ThreadShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunLegalAdviceConsultation"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/farui/legalAdvice/consult"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
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
