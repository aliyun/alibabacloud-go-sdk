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
		"cn-hangzhou":                 dara.String("mt.cn-hangzhou.aliyuncs.com"),
		"ap-northeast-1":              dara.String("mt.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("mt.aliyuncs.com"),
		"ap-south-1":                  dara.String("mt.aliyuncs.com"),
		"ap-southeast-1":              dara.String("mt.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":              dara.String("mt.aliyuncs.com"),
		"ap-southeast-3":              dara.String("mt.aliyuncs.com"),
		"ap-southeast-5":              dara.String("mt.aliyuncs.com"),
		"cn-beijing":                  dara.String("mt.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("mt.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("mt.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("mt.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("mt.aliyuncs.com"),
		"cn-chengdu":                  dara.String("mt.aliyuncs.com"),
		"cn-edge-1":                   dara.String("mt.aliyuncs.com"),
		"cn-fujian":                   dara.String("mt.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("mt.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("mt.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("mt.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("mt.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("mt.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("mt.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("mt.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("mt.aliyuncs.com"),
		"cn-hongkong":                 dara.String("mt.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("mt.aliyuncs.com"),
		"cn-huhehaote":                dara.String("mt.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("mt.aliyuncs.com"),
		"cn-qingdao":                  dara.String("mt.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("mt.aliyuncs.com"),
		"cn-shanghai":                 dara.String("mt.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("mt.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("mt.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("mt.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("mt.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("mt.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("mt.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("mt.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("mt.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("mt.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("mt.aliyuncs.com"),
		"cn-wuhan":                    dara.String("mt.aliyuncs.com"),
		"cn-yushanfang":               dara.String("mt.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("mt.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("mt.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("mt.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("mt.aliyuncs.com"),
		"eu-central-1":                dara.String("mt.aliyuncs.com"),
		"eu-west-1":                   dara.String("mt.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("mt.aliyuncs.com"),
		"me-east-1":                   dara.String("mt.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("mt.aliyuncs.com"),
		"us-east-1":                   dara.String("mt.aliyuncs.com"),
		"us-west-1":                   dara.String("mt.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("alimt"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 大文本异步翻译，支持5000-50000字翻译
//
// @param request - CreateAsyncTranslateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAsyncTranslateResponse
func (client *Client) CreateAsyncTranslateWithOptions(request *CreateAsyncTranslateRequest, runtime *dara.RuntimeOptions) (_result *CreateAsyncTranslateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiType) {
		body["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.FormatType) {
		body["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAsyncTranslate"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAsyncTranslateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 大文本异步翻译，支持5000-50000字翻译
//
// @param request - CreateAsyncTranslateRequest
//
// @return CreateAsyncTranslateResponse
func (client *Client) CreateAsyncTranslate(request *CreateAsyncTranslateRequest) (_result *CreateAsyncTranslateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAsyncTranslateResponse{}
	_body, _err := client.CreateAsyncTranslateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateDocTranslateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDocTranslateTaskResponse
func (client *Client) CreateDocTranslateTaskWithOptions(request *CreateDocTranslateTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDocTranslateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CallbackUrl) {
		body["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDocTranslateTask"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDocTranslateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateDocTranslateTaskRequest
//
// @return CreateDocTranslateTaskResponse
func (client *Client) CreateDocTranslateTask(request *CreateDocTranslateTaskRequest) (_result *CreateDocTranslateTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDocTranslateTaskResponse{}
	_body, _err := client.CreateDocTranslateTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDocTranslateTaskAdvance(request *CreateDocTranslateTaskAdvanceRequest, runtime *dara.RuntimeOptions) (_result *CreateDocTranslateTaskResponse, _err error) {
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
		"Product":  dara.String("alimt"),
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
	createDocTranslateTaskReq := &CreateDocTranslateTaskRequest{}
	openapiutil.Convert(request, createDocTranslateTaskReq)
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
		createDocTranslateTaskReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	createDocTranslateTaskResp, _err := client.CreateDocTranslateTaskWithOptions(createDocTranslateTaskReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = createDocTranslateTaskResp
	return _result, _err
}

// Summary:
//
// 创建图片翻译任务
//
// @param request - CreateImageTranslateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageTranslateTaskResponse
func (client *Client) CreateImageTranslateTaskWithOptions(request *CreateImageTranslateTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateImageTranslateTaskResponse, _err error) {
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

	if !dara.IsNil(request.Extra) {
		body["Extra"] = request.Extra
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	if !dara.IsNil(request.UrlList) {
		body["UrlList"] = request.UrlList
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageTranslateTask"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageTranslateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建图片翻译任务
//
// @param request - CreateImageTranslateTaskRequest
//
// @return CreateImageTranslateTaskResponse
func (client *Client) CreateImageTranslateTask(request *CreateImageTranslateTaskRequest) (_result *CreateImageTranslateTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateImageTranslateTaskResponse{}
	_body, _err := client.CreateImageTranslateTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 大文本异步翻译，支持5000-50000字翻译
//
// @param request - GetAsyncTranslateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncTranslateResponse
func (client *Client) GetAsyncTranslateWithOptions(request *GetAsyncTranslateRequest, runtime *dara.RuntimeOptions) (_result *GetAsyncTranslateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAsyncTranslate"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAsyncTranslateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 大文本异步翻译，支持5000-50000字翻译
//
// @param request - GetAsyncTranslateRequest
//
// @return GetAsyncTranslateResponse
func (client *Client) GetAsyncTranslate(request *GetAsyncTranslateRequest) (_result *GetAsyncTranslateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAsyncTranslateResponse{}
	_body, _err := client.GetAsyncTranslateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量文本翻译
//
// @param request - GetBatchTranslateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTranslateResponse
func (client *Client) GetBatchTranslateWithOptions(request *GetBatchTranslateRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTranslateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiType) {
		body["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.FormatType) {
		body["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBatchTranslate"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBatchTranslateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量文本翻译
//
// @param request - GetBatchTranslateRequest
//
// @return GetBatchTranslateResponse
func (client *Client) GetBatchTranslate(request *GetBatchTranslateRequest) (_result *GetBatchTranslateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBatchTranslateResponse{}
	_body, _err := client.GetBatchTranslateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetBatchTranslateByVPC
//
// @param request - GetBatchTranslateByVPCRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTranslateByVPCResponse
func (client *Client) GetBatchTranslateByVPCWithOptions(request *GetBatchTranslateByVPCRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTranslateByVPCResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiType) {
		body["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.FormatType) {
		body["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBatchTranslateByVPC"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBatchTranslateByVPCResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetBatchTranslateByVPC
//
// @param request - GetBatchTranslateByVPCRequest
//
// @return GetBatchTranslateByVPCResponse
func (client *Client) GetBatchTranslateByVPC(request *GetBatchTranslateByVPCRequest) (_result *GetBatchTranslateByVPCResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBatchTranslateByVPCResponse{}
	_body, _err := client.GetBatchTranslateByVPCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 语种识别
//
// @param request - GetDetectLanguageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDetectLanguageResponse
func (client *Client) GetDetectLanguageWithOptions(request *GetDetectLanguageRequest, runtime *dara.RuntimeOptions) (_result *GetDetectLanguageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDetectLanguage"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDetectLanguageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语种识别
//
// @param request - GetDetectLanguageRequest
//
// @return GetDetectLanguageResponse
func (client *Client) GetDetectLanguage(request *GetDetectLanguageRequest) (_result *GetDetectLanguageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDetectLanguageResponse{}
	_body, _err := client.GetDetectLanguageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 语种识别
//
// @param request - GetDetectLanguageVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDetectLanguageVpcResponse
func (client *Client) GetDetectLanguageVpcWithOptions(request *GetDetectLanguageVpcRequest, runtime *dara.RuntimeOptions) (_result *GetDetectLanguageVpcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDetectLanguageVpc"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDetectLanguageVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语种识别
//
// @param request - GetDetectLanguageVpcRequest
//
// @return GetDetectLanguageVpcResponse
func (client *Client) GetDetectLanguageVpc(request *GetDetectLanguageVpcRequest) (_result *GetDetectLanguageVpcResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDetectLanguageVpcResponse{}
	_body, _err := client.GetDetectLanguageVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档翻译任务
//
// @param request - GetDocTranslateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocTranslateTaskResponse
func (client *Client) GetDocTranslateTaskWithOptions(request *GetDocTranslateTaskRequest, runtime *dara.RuntimeOptions) (_result *GetDocTranslateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocTranslateTask"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocTranslateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档翻译任务
//
// @param request - GetDocTranslateTaskRequest
//
// @return GetDocTranslateTaskResponse
func (client *Client) GetDocTranslateTask(request *GetDocTranslateTaskRequest) (_result *GetDocTranslateTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDocTranslateTaskResponse{}
	_body, _err := client.GetDocTranslateTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetImageDiagnoseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageDiagnoseResponse
func (client *Client) GetImageDiagnoseWithOptions(request *GetImageDiagnoseRequest, runtime *dara.RuntimeOptions) (_result *GetImageDiagnoseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Extra) {
		body["Extra"] = request.Extra
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageDiagnose"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageDiagnoseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetImageDiagnoseRequest
//
// @return GetImageDiagnoseResponse
func (client *Client) GetImageDiagnose(request *GetImageDiagnoseRequest) (_result *GetImageDiagnoseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImageDiagnoseResponse{}
	_body, _err := client.GetImageDiagnoseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取图片翻译结果
//
// @param request - GetImageTranslateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageTranslateResponse
func (client *Client) GetImageTranslateWithOptions(request *GetImageTranslateRequest, runtime *dara.RuntimeOptions) (_result *GetImageTranslateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Extra) {
		body["Extra"] = request.Extra
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageTranslate"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageTranslateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取图片翻译结果
//
// @param request - GetImageTranslateRequest
//
// @return GetImageTranslateResponse
func (client *Client) GetImageTranslate(request *GetImageTranslateRequest) (_result *GetImageTranslateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImageTranslateResponse{}
	_body, _err := client.GetImageTranslateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取图片翻译任务
//
// @param request - GetImageTranslateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageTranslateTaskResponse
func (client *Client) GetImageTranslateTaskWithOptions(request *GetImageTranslateTaskRequest, runtime *dara.RuntimeOptions) (_result *GetImageTranslateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageTranslateTask"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageTranslateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取图片翻译任务
//
// @param request - GetImageTranslateTaskRequest
//
// @return GetImageTranslateTaskResponse
func (client *Client) GetImageTranslateTask(request *GetImageTranslateTaskRequest) (_result *GetImageTranslateTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImageTranslateTaskResponse{}
	_body, _err := client.GetImageTranslateTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetTitleDiagnose
//
// @param request - GetTitleDiagnoseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTitleDiagnoseResponse
func (client *Client) GetTitleDiagnoseWithOptions(request *GetTitleDiagnoseRequest, runtime *dara.RuntimeOptions) (_result *GetTitleDiagnoseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Extra) {
		body["Extra"] = request.Extra
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTitleDiagnose"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTitleDiagnoseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetTitleDiagnose
//
// @param request - GetTitleDiagnoseRequest
//
// @return GetTitleDiagnoseResponse
func (client *Client) GetTitleDiagnose(request *GetTitleDiagnoseRequest) (_result *GetTitleDiagnoseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTitleDiagnoseResponse{}
	_body, _err := client.GetTitleDiagnoseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetTitleGenerate
//
// @param request - GetTitleGenerateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTitleGenerateResponse
func (client *Client) GetTitleGenerateWithOptions(request *GetTitleGenerateRequest, runtime *dara.RuntimeOptions) (_result *GetTitleGenerateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Attributes) {
		body["Attributes"] = request.Attributes
	}

	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Extra) {
		body["Extra"] = request.Extra
	}

	if !dara.IsNil(request.HotWords) {
		body["HotWords"] = request.HotWords
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTitleGenerate"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTitleGenerateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetTitleGenerate
//
// @param request - GetTitleGenerateRequest
//
// @return GetTitleGenerateResponse
func (client *Client) GetTitleGenerate(request *GetTitleGenerateRequest) (_result *GetTitleGenerateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTitleGenerateResponse{}
	_body, _err := client.GetTitleGenerateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetTitleIntelligence
//
// @param request - GetTitleIntelligenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTitleIntelligenceResponse
func (client *Client) GetTitleIntelligenceWithOptions(request *GetTitleIntelligenceRequest, runtime *dara.RuntimeOptions) (_result *GetTitleIntelligenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CatLevelThreeId) {
		body["CatLevelThreeId"] = request.CatLevelThreeId
	}

	if !dara.IsNil(request.CatLevelTwoId) {
		body["CatLevelTwoId"] = request.CatLevelTwoId
	}

	if !dara.IsNil(request.Extra) {
		body["Extra"] = request.Extra
	}

	if !dara.IsNil(request.Keywords) {
		body["Keywords"] = request.Keywords
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTitleIntelligence"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTitleIntelligenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetTitleIntelligence
//
// @param request - GetTitleIntelligenceRequest
//
// @return GetTitleIntelligenceResponse
func (client *Client) GetTitleIntelligence(request *GetTitleIntelligenceRequest) (_result *GetTitleIntelligenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTitleIntelligenceResponse{}
	_body, _err := client.GetTitleIntelligenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取图片批量翻译结果
//
// @param request - GetTranslateImageBatchResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranslateImageBatchResultResponse
func (client *Client) GetTranslateImageBatchResultWithOptions(request *GetTranslateImageBatchResultRequest, runtime *dara.RuntimeOptions) (_result *GetTranslateImageBatchResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTranslateImageBatchResult"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTranslateImageBatchResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取图片批量翻译结果
//
// @param request - GetTranslateImageBatchResultRequest
//
// @return GetTranslateImageBatchResultResponse
func (client *Client) GetTranslateImageBatchResult(request *GetTranslateImageBatchResultRequest) (_result *GetTranslateImageBatchResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTranslateImageBatchResultResponse{}
	_body, _err := client.GetTranslateImageBatchResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetTranslateReport
//
// @param request - GetTranslateReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranslateReportResponse
func (client *Client) GetTranslateReportWithOptions(request *GetTranslateReportRequest, runtime *dara.RuntimeOptions) (_result *GetTranslateReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTranslateReport"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTranslateReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetTranslateReport
//
// @param request - GetTranslateReportRequest
//
// @return GetTranslateReportResponse
func (client *Client) GetTranslateReport(request *GetTranslateReportRequest) (_result *GetTranslateReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTranslateReportResponse{}
	_body, _err := client.GetTranslateReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开通服务
//
// @param request - OpenAlimtServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenAlimtServiceResponse
func (client *Client) OpenAlimtServiceWithOptions(request *OpenAlimtServiceRequest, runtime *dara.RuntimeOptions) (_result *OpenAlimtServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenAlimtService"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenAlimtServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开通服务
//
// @param request - OpenAlimtServiceRequest
//
// @return OpenAlimtServiceResponse
func (client *Client) OpenAlimtService(request *OpenAlimtServiceRequest) (_result *OpenAlimtServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenAlimtServiceResponse{}
	_body, _err := client.OpenAlimtServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 专业文本翻译
//
// @param request - TranslateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateResponse
func (client *Client) TranslateWithOptions(request *TranslateRequest, runtime *dara.RuntimeOptions) (_result *TranslateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Context) {
		query["Context"] = request.Context
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FormatType) {
		body["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Translate"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TranslateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 专业文本翻译
//
// @param request - TranslateRequest
//
// @return TranslateResponse
func (client *Client) Translate(request *TranslateRequest) (_result *TranslateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TranslateResponse{}
	_body, _err := client.TranslateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # TranslateCertificate
//
// @param request - TranslateCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateCertificateResponse
func (client *Client) TranslateCertificateWithOptions(request *TranslateCertificateRequest, runtime *dara.RuntimeOptions) (_result *TranslateCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CertificateType) {
		body["CertificateType"] = request.CertificateType
	}

	if !dara.IsNil(request.ImageUrl) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.ResultType) {
		body["ResultType"] = request.ResultType
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TranslateCertificate"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TranslateCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # TranslateCertificate
//
// @param request - TranslateCertificateRequest
//
// @return TranslateCertificateResponse
func (client *Client) TranslateCertificate(request *TranslateCertificateRequest) (_result *TranslateCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TranslateCertificateResponse{}
	_body, _err := client.TranslateCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TranslateCertificateAdvance(request *TranslateCertificateAdvanceRequest, runtime *dara.RuntimeOptions) (_result *TranslateCertificateResponse, _err error) {
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
		"Product":  dara.String("alimt"),
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
	translateCertificateReq := &TranslateCertificateRequest{}
	openapiutil.Convert(request, translateCertificateReq)
	if !dara.IsNil(request.ImageUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.ImageUrlObject,
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
		translateCertificateReq.ImageUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	translateCertificateResp, _err := client.TranslateCertificateWithOptions(translateCertificateReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = translateCertificateResp
	return _result, _err
}

// Deprecated: OpenAPI TranslateECommerce is deprecated, please use alimt::2018-10-12::Translate instead.
//
// Summary:
//
// # TranslateECommerce
//
// @param request - TranslateECommerceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateECommerceResponse
func (client *Client) TranslateECommerceWithOptions(request *TranslateECommerceRequest, runtime *dara.RuntimeOptions) (_result *TranslateECommerceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Context) {
		query["Context"] = request.Context
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FormatType) {
		body["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TranslateECommerce"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TranslateECommerceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI TranslateECommerce is deprecated, please use alimt::2018-10-12::Translate instead.
//
// Summary:
//
// # TranslateECommerce
//
// @param request - TranslateECommerceRequest
//
// @return TranslateECommerceResponse
// Deprecated
func (client *Client) TranslateECommerce(request *TranslateECommerceRequest) (_result *TranslateECommerceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TranslateECommerceResponse{}
	_body, _err := client.TranslateECommerceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文本通用翻译
//
// @param request - TranslateGeneralRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateGeneralResponse
func (client *Client) TranslateGeneralWithOptions(request *TranslateGeneralRequest, runtime *dara.RuntimeOptions) (_result *TranslateGeneralResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Context) {
		query["Context"] = request.Context
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FormatType) {
		body["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TranslateGeneral"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TranslateGeneralResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文本通用翻译
//
// @param request - TranslateGeneralRequest
//
// @return TranslateGeneralResponse
func (client *Client) TranslateGeneral(request *TranslateGeneralRequest) (_result *TranslateGeneralResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TranslateGeneralResponse{}
	_body, _err := client.TranslateGeneralWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # TranslateGeneralVpc
//
// @param request - TranslateGeneralVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateGeneralVpcResponse
func (client *Client) TranslateGeneralVpcWithOptions(request *TranslateGeneralVpcRequest, runtime *dara.RuntimeOptions) (_result *TranslateGeneralVpcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Context) {
		query["Context"] = request.Context
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FormatType) {
		body["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TranslateGeneralVpc"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TranslateGeneralVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # TranslateGeneralVpc
//
// @param request - TranslateGeneralVpcRequest
//
// @return TranslateGeneralVpcResponse
func (client *Client) TranslateGeneralVpc(request *TranslateGeneralVpcRequest) (_result *TranslateGeneralVpcResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TranslateGeneralVpcResponse{}
	_body, _err := client.TranslateGeneralVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 公有云图片翻译产品API
//
// @param request - TranslateImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateImageResponse
func (client *Client) TranslateImageWithOptions(request *TranslateImageRequest, runtime *dara.RuntimeOptions) (_result *TranslateImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Ext) {
		body["Ext"] = request.Ext
	}

	if !dara.IsNil(request.Field) {
		body["Field"] = request.Field
	}

	if !dara.IsNil(request.ImageBase64) {
		body["ImageBase64"] = request.ImageBase64
	}

	if !dara.IsNil(request.ImageUrl) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TranslateImage"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TranslateImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 公有云图片翻译产品API
//
// @param request - TranslateImageRequest
//
// @return TranslateImageResponse
func (client *Client) TranslateImage(request *TranslateImageRequest) (_result *TranslateImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TranslateImageResponse{}
	_body, _err := client.TranslateImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量图片翻译接口
//
// @param request - TranslateImageBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateImageBatchResponse
func (client *Client) TranslateImageBatchWithOptions(request *TranslateImageBatchRequest, runtime *dara.RuntimeOptions) (_result *TranslateImageBatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomTaskId) {
		body["CustomTaskId"] = request.CustomTaskId
	}

	if !dara.IsNil(request.Ext) {
		body["Ext"] = request.Ext
	}

	if !dara.IsNil(request.Field) {
		body["Field"] = request.Field
	}

	if !dara.IsNil(request.ImageUrls) {
		body["ImageUrls"] = request.ImageUrls
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TranslateImageBatch"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TranslateImageBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量图片翻译接口
//
// @param request - TranslateImageBatchRequest
//
// @return TranslateImageBatchResponse
func (client *Client) TranslateImageBatch(request *TranslateImageBatchRequest) (_result *TranslateImageBatchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TranslateImageBatchResponse{}
	_body, _err := client.TranslateImageBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 搜索翻译
//
// @param request - TranslateSearchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateSearchResponse
func (client *Client) TranslateSearchWithOptions(request *TranslateSearchRequest, runtime *dara.RuntimeOptions) (_result *TranslateSearchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FormatType) {
		body["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TranslateSearch"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TranslateSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索翻译
//
// @param request - TranslateSearchRequest
//
// @return TranslateSearchResponse
func (client *Client) TranslateSearch(request *TranslateSearchRequest) (_result *TranslateSearchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TranslateSearchResponse{}
	_body, _err := client.TranslateSearchWithOptions(request, runtime)
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
