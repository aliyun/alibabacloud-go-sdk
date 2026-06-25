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
		"eu-central-1":   dara.String("imagesearch.eu-central-1.aliyuncs.com"),
		"cn-shenzhen":    dara.String("imagesearch.cn-shenzhen.aliyuncs.com"),
		"cn-shanghai":    dara.String("imagesearch.cn-shanghai.aliyuncs.com"),
		"cn-hongkong":    dara.String("imagesearch.cn-hongkong.aliyuncs.com"),
		"cn-hangzhou":    dara.String("imagesearch.cn-hangzhou.aliyuncs.com"),
		"cn-beijing":     dara.String("imagesearch.cn-beijing.aliyuncs.com"),
		"ap-southeast-2": dara.String("imagesearch.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-1": dara.String("imagesearch.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":     dara.String("imagesearch.ap-south-1.aliyuncs.com"),
		"ap-northeast-1": dara.String("imagesearch.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("imagesearch"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Describes the syntax and provides examples of the AddImage operation, which adds image information to an Image Search instance.
//
// Description:
//
// ## Description
//
// This operation adds image information to an Image Search instance.
//
// ## QPS limit
//
// An instance with a maximum image capacity of 100,000 has a default concurrency of 1, which means that a maximum of 1 image addition request can be processed per second.
//
// Instances with other image capacities have a default concurrency of 5, which means that a maximum of 5 image addition requests can be processed per second.
//
// @param request - AddImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddImageResponse
func (client *Client) AddImageWithOptions(request *AddImageRequest, runtime *dara.RuntimeOptions) (_result *AddImageResponse, _err error) {
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

	if !dara.IsNil(request.Crop) {
		body["Crop"] = request.Crop
	}

	if !dara.IsNil(request.CustomContent) {
		body["CustomContent"] = request.CustomContent
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.IntAttr) {
		body["IntAttr"] = request.IntAttr
	}

	if !dara.IsNil(request.IntAttr2) {
		body["IntAttr2"] = request.IntAttr2
	}

	if !dara.IsNil(request.IntAttr3) {
		body["IntAttr3"] = request.IntAttr3
	}

	if !dara.IsNil(request.IntAttr4) {
		body["IntAttr4"] = request.IntAttr4
	}

	if !dara.IsNil(request.PicContent) {
		body["PicContent"] = request.PicContent
	}

	if !dara.IsNil(request.PicName) {
		body["PicName"] = request.PicName
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.StrAttr) {
		body["StrAttr"] = request.StrAttr
	}

	if !dara.IsNil(request.StrAttr2) {
		body["StrAttr2"] = request.StrAttr2
	}

	if !dara.IsNil(request.StrAttr3) {
		body["StrAttr3"] = request.StrAttr3
	}

	if !dara.IsNil(request.StrAttr4) {
		body["StrAttr4"] = request.StrAttr4
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddImage"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes the syntax and provides examples of the AddImage operation, which adds image information to an Image Search instance.
//
// Description:
//
// ## Description
//
// This operation adds image information to an Image Search instance.
//
// ## QPS limit
//
// An instance with a maximum image capacity of 100,000 has a default concurrency of 1, which means that a maximum of 1 image addition request can be processed per second.
//
// Instances with other image capacities have a default concurrency of 5, which means that a maximum of 5 image addition requests can be processed per second.
//
// @param request - AddImageRequest
//
// @return AddImageResponse
func (client *Client) AddImage(request *AddImageRequest) (_result *AddImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddImageResponse{}
	_body, _err := client.AddImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddImageAdvance(request *AddImageAdvanceRequest, runtime *dara.RuntimeOptions) (_result *AddImageResponse, _err error) {
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
		"Product":  dara.String("ImageSearch"),
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
	addImageReq := &AddImageRequest{}
	openapiutil.Convert(request, addImageReq)
	if !dara.IsNil(request.PicContentObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.PicContentObject,
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
		addImageReq.PicContent = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	addImageResp, _err := client.AddImageWithOptions(addImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = addImageResp
	return _result, _err
}

// Summary:
//
// # CheckImageExists
//
// Description:
//
// ## How-To
//
// This API is used to query image information in an Image Search instance based on an image.
//
// ## QPS Limit
//
// The default maximum queries per second (QPS) for query operations can be viewed in the console. It corresponds to the Visit Frequency (QPS) you selected when purchasing the instance. Supported QPS values are 1, 5, and 10.
//
// ### SDK Version Guide
//
// Upgrade the Image SDK to version V3.1.1 to use the "subject identification" and "similarity score" features. For more information, see [Java SDK](https://help.aliyun.com/document_detail/179188.html).
//
// @param request - CheckImageExistsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckImageExistsResponse
func (client *Client) CheckImageExistsWithOptions(request *CheckImageExistsRequest, runtime *dara.RuntimeOptions) (_result *CheckImageExistsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PicName) {
		body["PicName"] = request.PicName
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckImageExists"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckImageExistsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CheckImageExists
//
// Description:
//
// ## How-To
//
// This API is used to query image information in an Image Search instance based on an image.
//
// ## QPS Limit
//
// The default maximum queries per second (QPS) for query operations can be viewed in the console. It corresponds to the Visit Frequency (QPS) you selected when purchasing the instance. Supported QPS values are 1, 5, and 10.
//
// ### SDK Version Guide
//
// Upgrade the Image SDK to version V3.1.1 to use the "subject identification" and "similarity score" features. For more information, see [Java SDK](https://help.aliyun.com/document_detail/179188.html).
//
// @param request - CheckImageExistsRequest
//
// @return CheckImageExistsResponse
func (client *Client) CheckImageExists(request *CheckImageExistsRequest) (_result *CheckImageExistsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckImageExistsResponse{}
	_body, _err := client.CheckImageExistsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Compares two images and returns a similarity score.
//
// @param request - CompareSimilarByImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompareSimilarByImageResponse
func (client *Client) CompareSimilarByImageWithOptions(request *CompareSimilarByImageRequest, runtime *dara.RuntimeOptions) (_result *CompareSimilarByImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PrimaryPicContent) {
		body["PrimaryPicContent"] = request.PrimaryPicContent
	}

	if !dara.IsNil(request.SecondaryPicContent) {
		body["SecondaryPicContent"] = request.SecondaryPicContent
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompareSimilarByImage"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompareSimilarByImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Compares two images and returns a similarity score.
//
// @param request - CompareSimilarByImageRequest
//
// @return CompareSimilarByImageResponse
func (client *Client) CompareSimilarByImage(request *CompareSimilarByImageRequest) (_result *CompareSimilarByImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CompareSimilarByImageResponse{}
	_body, _err := client.CompareSimilarByImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompareSimilarByImageAdvance(request *CompareSimilarByImageAdvanceRequest, runtime *dara.RuntimeOptions) (_result *CompareSimilarByImageResponse, _err error) {
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
		"Product":  dara.String("ImageSearch"),
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
	compareSimilarByImageReq := &CompareSimilarByImageRequest{}
	openapiutil.Convert(request, compareSimilarByImageReq)
	if !dara.IsNil(request.PrimaryPicContentObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.PrimaryPicContentObject,
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
		compareSimilarByImageReq.PrimaryPicContent = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	if !dara.IsNil(request.SecondaryPicContentObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.SecondaryPicContentObject,
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
		compareSimilarByImageReq.SecondaryPicContent = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	compareSimilarByImageResp, _err := client.CompareSimilarByImageWithOptions(compareSimilarByImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = compareSimilarByImageResp
	return _result, _err
}

// Summary:
//
// This topic describes the syntax and examples of the DeleteImage operation, which is used to delete image information from an Image Search instance.
//
// Description:
//
// ## Operation description
//
// This operation is used to delete image information from an Image Search instance.
//
// >- If the specified image does not exist in the Image Search instance, this operation still returns a success response. Do not use the response to determine whether the image exists.
//
// ## QPS limit
//
// The default concurrency for delete operations is 20, which means a maximum of 20 delete requests can be processed per second.
//
// @param request - DeleteImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteImageResponse
func (client *Client) DeleteImageWithOptions(request *DeleteImageRequest, runtime *dara.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.IsDeleteByFilter) {
		body["IsDeleteByFilter"] = request.IsDeleteByFilter
	}

	if !dara.IsNil(request.PicName) {
		body["PicName"] = request.PicName
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteImage"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the syntax and examples of the DeleteImage operation, which is used to delete image information from an Image Search instance.
//
// Description:
//
// ## Operation description
//
// This operation is used to delete image information from an Image Search instance.
//
// >- If the specified image does not exist in the Image Search instance, this operation still returns a success response. Do not use the response to determine whether the image exists.
//
// ## QPS limit
//
// The default concurrency for delete operations is 20, which means a maximum of 20 delete requests can be processed per second.
//
// @param request - DeleteImageRequest
//
// @return DeleteImageResponse
func (client *Client) DeleteImage(request *DeleteImageRequest) (_result *DeleteImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteImageResponse{}
	_body, _err := client.DeleteImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes the syntax and examples of the Detail operation, which queries information about an Image Search instance by name.
//
// Description:
//
// ## Operation description
//
// This operation queries instance information from an Image Search instance.
//
// > For more product details or technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ## QPS limit
//
// The default concurrency for query operations is 1, which means a maximum of 1 request is processed per second.
//
// @param request - DetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetailResponse
func (client *Client) DetailWithOptions(request *DetailRequest, runtime *dara.RuntimeOptions) (_result *DetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Detail"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the syntax and examples of the Detail operation, which queries information about an Image Search instance by name.
//
// Description:
//
// ## Operation description
//
// This operation queries instance information from an Image Search instance.
//
// > For more product details or technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ## QPS limit
//
// The default concurrency for query operations is 1, which means a maximum of 1 request is processed per second.
//
// @param request - DetailRequest
//
// @return DetailResponse
func (client *Client) Detail(request *DetailRequest) (_result *DetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetailResponse{}
	_body, _err := client.DetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes the syntax and examples of the DumpMeta operation, which creates a metadata export task for Image Search by name.
//
// Description:
//
// ## Operation description
//
// This operation submits a metadata export task to an Image Search instance.
//
// > For more product details and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ## QPS limit
//
// The default concurrency for submit operations is 1, which means a maximum of 1 request is processed per second.
//
// > You cannot submit a new metadata export task while the previous metadata export task is still in progress.
//
// @param request - DumpMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DumpMetaResponse
func (client *Client) DumpMetaWithOptions(request *DumpMetaRequest, runtime *dara.RuntimeOptions) (_result *DumpMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DumpMeta"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DumpMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the syntax and examples of the DumpMeta operation, which creates a metadata export task for Image Search by name.
//
// Description:
//
// ## Operation description
//
// This operation submits a metadata export task to an Image Search instance.
//
// > For more product details and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ## QPS limit
//
// The default concurrency for submit operations is 1, which means a maximum of 1 request is processed per second.
//
// > You cannot submit a new metadata export task while the previous metadata export task is still in progress.
//
// @param request - DumpMetaRequest
//
// @return DumpMetaResponse
func (client *Client) DumpMeta(request *DumpMetaRequest) (_result *DumpMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DumpMetaResponse{}
	_body, _err := client.DumpMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Describes the syntax and provides examples of the DumpMetaList operation, which queries the list of metadata export tasks in an Image Search instance.
//
// Description:
//
// ## Operation description
//
// This operation queries metadata export tasks in an Image Search instance.
//
// > For more product details and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ## QPS limit
//
// The default concurrency for query operations is 1, which means a maximum of 1 request is processed per second.
//
// @param request - DumpMetaListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DumpMetaListResponse
func (client *Client) DumpMetaListWithOptions(request *DumpMetaListRequest, runtime *dara.RuntimeOptions) (_result *DumpMetaListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DumpMetaList"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DumpMetaListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes the syntax and provides examples of the DumpMetaList operation, which queries the list of metadata export tasks in an Image Search instance.
//
// Description:
//
// ## Operation description
//
// This operation queries metadata export tasks in an Image Search instance.
//
// > For more product details and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ## QPS limit
//
// The default concurrency for query operations is 1, which means a maximum of 1 request is processed per second.
//
// @param request - DumpMetaListRequest
//
// @return DumpMetaListResponse
func (client *Client) DumpMetaList(request *DumpMetaListRequest) (_result *DumpMetaListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DumpMetaListResponse{}
	_body, _err := client.DumpMetaListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Describes the syntax and provides examples of the IncreaseInstance operation, which is used to create a batch task for an Image Search instance by name.
//
// Description:
//
// ## Operation description
//
// This operation is used to submit a batch task to an Image Search instance.
//
// > <props="china">For more information about the product or technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us by using DingTalk group 35035130.
//
// ## QPS limit
//
// Only one batch task can run at a time.
//
// > You cannot submit a new batch task until the previous batch task is complete.
//
// @param request - IncreaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IncreaseInstanceResponse
func (client *Client) IncreaseInstanceWithOptions(request *IncreaseInstanceRequest, runtime *dara.RuntimeOptions) (_result *IncreaseInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.CallbackAddress) {
		query["CallbackAddress"] = request.CallbackAddress
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IncreaseInstance"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IncreaseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes the syntax and provides examples of the IncreaseInstance operation, which is used to create a batch task for an Image Search instance by name.
//
// Description:
//
// ## Operation description
//
// This operation is used to submit a batch task to an Image Search instance.
//
// > <props="china">For more information about the product or technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us by using DingTalk group 35035130.
//
// ## QPS limit
//
// Only one batch task can run at a time.
//
// > You cannot submit a new batch task until the previous batch task is complete.
//
// @param request - IncreaseInstanceRequest
//
// @return IncreaseInstanceResponse
func (client *Client) IncreaseInstance(request *IncreaseInstanceRequest) (_result *IncreaseInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &IncreaseInstanceResponse{}
	_body, _err := client.IncreaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of batch tasks in an Image Search instance by calling the IncreaseList operation. This topic describes the syntax and provides examples.
//
// Description:
//
// ## Operation description
//
// This operation is used to query batch tasks in an Image Search instance.
//
// > For more product details or technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ## QPS limit
//
// The default concurrency for query operations is 1, which means a maximum of 1 request is processed per second.
//
// @param request - IncreaseListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IncreaseListResponse
func (client *Client) IncreaseListWithOptions(request *IncreaseListRequest, runtime *dara.RuntimeOptions) (_result *IncreaseListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IncreaseList"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IncreaseListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of batch tasks in an Image Search instance by calling the IncreaseList operation. This topic describes the syntax and provides examples.
//
// Description:
//
// ## Operation description
//
// This operation is used to query batch tasks in an Image Search instance.
//
// > For more product details or technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ## QPS limit
//
// The default concurrency for query operations is 1, which means a maximum of 1 request is processed per second.
//
// @param request - IncreaseListRequest
//
// @return IncreaseListResponse
func (client *Client) IncreaseList(request *IncreaseListRequest) (_result *IncreaseListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &IncreaseListResponse{}
	_body, _err := client.IncreaseListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes the syntax and examples of SearchImageByFilter, which is used to query image information in an Image Search instance based on filter conditions.
//
// @param request - SearchImageByFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchImageByFilterResponse
func (client *Client) SearchImageByFilterWithOptions(request *SearchImageByFilterRequest, runtime *dara.RuntimeOptions) (_result *SearchImageByFilterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Num) {
		body["Num"] = request.Num
	}

	if !dara.IsNil(request.Start) {
		body["Start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchImageByFilter"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchImageByFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the syntax and examples of SearchImageByFilter, which is used to query image information in an Image Search instance based on filter conditions.
//
// @param request - SearchImageByFilterRequest
//
// @return SearchImageByFilterResponse
func (client *Client) SearchImageByFilter(request *SearchImageByFilterRequest) (_result *SearchImageByFilterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchImageByFilterResponse{}
	_body, _err := client.SearchImageByFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes the syntax and examples of the SearchByName operation, which is used to query image information in an Image Search instance by name.
//
// Description:
//
// ### Operation description
//
// This operation queries image information in an Image Search instance by name (ProductId and PicName).
//
// > For more product details and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ### QPS limit
//
// The default maximum query rate can be viewed in the console. It is the QPS value you selected at the time of purchase. Currently supported values are 1 QPS, 5 QPS, and 10 QPS.
//
// @param request - SearchImageByNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchImageByNameResponse
func (client *Client) SearchImageByNameWithOptions(request *SearchImageByNameRequest, runtime *dara.RuntimeOptions) (_result *SearchImageByNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ScoreThreshold) {
		query["ScoreThreshold"] = request.ScoreThreshold
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.DistinctProductId) {
		body["DistinctProductId"] = request.DistinctProductId
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Num) {
		body["Num"] = request.Num
	}

	if !dara.IsNil(request.PicName) {
		body["PicName"] = request.PicName
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.Start) {
		body["Start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchImageByName"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchImageByNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the syntax and examples of the SearchByName operation, which is used to query image information in an Image Search instance by name.
//
// Description:
//
// ### Operation description
//
// This operation queries image information in an Image Search instance by name (ProductId and PicName).
//
// > For more product details and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ### QPS limit
//
// The default maximum query rate can be viewed in the console. It is the QPS value you selected at the time of purchase. Currently supported values are 1 QPS, 5 QPS, and 10 QPS.
//
// @param request - SearchImageByNameRequest
//
// @return SearchImageByNameResponse
func (client *Client) SearchImageByName(request *SearchImageByNameRequest) (_result *SearchImageByNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchImageByNameResponse{}
	_body, _err := client.SearchImageByNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes the syntax and examples of SearchByPic, which is used to search for image information in an Image Search instance by image.
//
// Description:
//
// ## Operation description
//
// This operation is used to search for image information in an Image Search instance by image.
//
// > <props="china">For more product details and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ## QPS limit
//
// You can view the default maximum access frequency for query operations in the console. The frequency is the QPS value that you selected when you made the purchase. The supported values are 1 QPS, 5 QPS, and 10 QPS.
//
// ### SDK version description
//
// Upgrade the Image Search SDK to V3.1.1 to use the multi-subject identification and similarity score features. For more information, see [Java SDK](https://help.aliyun.com/document_detail/179188.html).
//
// @param request - SearchImageByPicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchImageByPicResponse
func (client *Client) SearchImageByPicWithOptions(request *SearchImageByPicRequest, runtime *dara.RuntimeOptions) (_result *SearchImageByPicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ScoreThreshold) {
		query["ScoreThreshold"] = request.ScoreThreshold
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Crop) {
		body["Crop"] = request.Crop
	}

	if !dara.IsNil(request.DistinctProductId) {
		body["DistinctProductId"] = request.DistinctProductId
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Num) {
		body["Num"] = request.Num
	}

	if !dara.IsNil(request.PicContent) {
		body["PicContent"] = request.PicContent
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.Start) {
		body["Start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchImageByPic"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchImageByPicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the syntax and examples of SearchByPic, which is used to search for image information in an Image Search instance by image.
//
// Description:
//
// ## Operation description
//
// This operation is used to search for image information in an Image Search instance by image.
//
// > <props="china">For more product details and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ## QPS limit
//
// You can view the default maximum access frequency for query operations in the console. The frequency is the QPS value that you selected when you made the purchase. The supported values are 1 QPS, 5 QPS, and 10 QPS.
//
// ### SDK version description
//
// Upgrade the Image Search SDK to V3.1.1 to use the multi-subject identification and similarity score features. For more information, see [Java SDK](https://help.aliyun.com/document_detail/179188.html).
//
// @param request - SearchImageByPicRequest
//
// @return SearchImageByPicResponse
func (client *Client) SearchImageByPic(request *SearchImageByPicRequest) (_result *SearchImageByPicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchImageByPicResponse{}
	_body, _err := client.SearchImageByPicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchImageByPicAdvance(request *SearchImageByPicAdvanceRequest, runtime *dara.RuntimeOptions) (_result *SearchImageByPicResponse, _err error) {
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
		"Product":  dara.String("ImageSearch"),
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
	searchImageByPicReq := &SearchImageByPicRequest{}
	openapiutil.Convert(request, searchImageByPicReq)
	if !dara.IsNil(request.PicContentObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.PicContentObject,
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
		searchImageByPicReq.PicContent = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	searchImageByPicResp, _err := client.SearchImageByPicWithOptions(searchImageByPicReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = searchImageByPicResp
	return _result, _err
}

// Summary:
//
// This topic describes the syntax and examples of SearchImageByText, which is used to search for image information in an Image Search instance based on text.
//
// Description:
//
// ## Operation description
//
// This operation is used to search for image information in an Image Search instance based on text. This operation is available only for instances whose service type is product multimodal search.
//
// > <props="china">For more product details and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through the DingTalk group (35035130).
//
// ## QPS limit
//
// You can view the default maximum access frequency for query operations in the console. The frequency is the QPS value you selected at the time of purchase. Currently supported values are 1 QPS, 5 QPS, and 10 QPS.
//
// ### SDK version description
//
// Upgrade the Image Search SDK to V3.1.1 to use the multi-subject identification and similarity score features. For more information, see [Java SDK](https://help.aliyun.com/document_detail/179188.html).
//
// @param request - SearchImageByTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchImageByTextResponse
func (client *Client) SearchImageByTextWithOptions(request *SearchImageByTextRequest, runtime *dara.RuntimeOptions) (_result *SearchImageByTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ScoreThreshold) {
		query["ScoreThreshold"] = request.ScoreThreshold
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DistinctProductId) {
		body["DistinctProductId"] = request.DistinctProductId
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Num) {
		body["Num"] = request.Num
	}

	if !dara.IsNil(request.Start) {
		body["Start"] = request.Start
	}

	if !dara.IsNil(request.Text) {
		body["Text"] = request.Text
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchImageByText"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchImageByTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the syntax and examples of SearchImageByText, which is used to search for image information in an Image Search instance based on text.
//
// Description:
//
// ## Operation description
//
// This operation is used to search for image information in an Image Search instance based on text. This operation is available only for instances whose service type is product multimodal search.
//
// > <props="china">For more product details and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through the DingTalk group (35035130).
//
// ## QPS limit
//
// You can view the default maximum access frequency for query operations in the console. The frequency is the QPS value you selected at the time of purchase. Currently supported values are 1 QPS, 5 QPS, and 10 QPS.
//
// ### SDK version description
//
// Upgrade the Image Search SDK to V3.1.1 to use the multi-subject identification and similarity score features. For more information, see [Java SDK](https://help.aliyun.com/document_detail/179188.html).
//
// @param request - SearchImageByTextRequest
//
// @return SearchImageByTextResponse
func (client *Client) SearchImageByText(request *SearchImageByTextRequest) (_result *SearchImageByTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchImageByTextResponse{}
	_body, _err := client.SearchImageByTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the image information in an Image Search instance.
//
// Description:
//
// ## Usage notes
//
// This operation updates the image information in an Image Search instance based on the product ID and image name.
//
// > - The instance must meet the creation date requirements.
//
// <props="china">
//
// - Instances created after June 2021 in the Shanghai and Hangzhou regions are supported. Instances in other regions can be used normally.
//
// <props="intl">
//
// - Instances created after December 2021 in the Singapore region are supported. Instances in other regions are currently unavailable.
//
// - For more information about the product and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through the DingTalk group (35035130).
//
// ## QPS limit
//
// The default concurrency for update operations is 20, which means that a maximum of 20 requests can be processed per second.
//
// @param request - UpdateImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateImageResponse
func (client *Client) UpdateImageWithOptions(request *UpdateImageRequest, runtime *dara.RuntimeOptions) (_result *UpdateImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IntAttr3) {
		query["IntAttr3"] = request.IntAttr3
	}

	if !dara.IsNil(request.IntAttr4) {
		query["IntAttr4"] = request.IntAttr4
	}

	if !dara.IsNil(request.StrAttr3) {
		query["StrAttr3"] = request.StrAttr3
	}

	if !dara.IsNil(request.StrAttr4) {
		query["StrAttr4"] = request.StrAttr4
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomContent) {
		body["CustomContent"] = request.CustomContent
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.IntAttr) {
		body["IntAttr"] = request.IntAttr
	}

	if !dara.IsNil(request.IntAttr2) {
		body["IntAttr2"] = request.IntAttr2
	}

	if !dara.IsNil(request.PicName) {
		body["PicName"] = request.PicName
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.StrAttr) {
		body["StrAttr"] = request.StrAttr
	}

	if !dara.IsNil(request.StrAttr2) {
		body["StrAttr2"] = request.StrAttr2
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateImage"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the image information in an Image Search instance.
//
// Description:
//
// ## Usage notes
//
// This operation updates the image information in an Image Search instance based on the product ID and image name.
//
// > - The instance must meet the creation date requirements.
//
// <props="china">
//
// - Instances created after June 2021 in the Shanghai and Hangzhou regions are supported. Instances in other regions can be used normally.
//
// <props="intl">
//
// - Instances created after December 2021 in the Singapore region are supported. Instances in other regions are currently unavailable.
//
// - For more information about the product and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through the DingTalk group (35035130).
//
// ## QPS limit
//
// The default concurrency for update operations is 20, which means that a maximum of 20 requests can be processed per second.
//
// @param request - UpdateImageRequest
//
// @return UpdateImageResponse
func (client *Client) UpdateImage(request *UpdateImageRequest) (_result *UpdateImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateImageResponse{}
	_body, _err := client.UpdateImageWithOptions(request, runtime)
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
