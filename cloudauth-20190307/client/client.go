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
		"us-west-1":             dara.String("cloudauth.aliyuncs.com"),
		"us-east-1":             dara.String("cloudauth.aliyuncs.com"),
		"me-east-1":             dara.String("cloudauth.aliyuncs.com"),
		"eu-west-1":             dara.String("cloudauth.aliyuncs.com"),
		"eu-central-1":          dara.String("cloudauth.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("cloudauth.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("cloudauth.aliyuncs.com"),
		"cn-shenzhen":           dara.String("cloudauth.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("cloudauth.aliyuncs.com"),
		"cn-shanghai":           dara.String("cloudauth.aliyuncs.com"),
		"cn-qingdao":            dara.String("cloudauth.cn-qingdao.aliyuncs.com"),
		"cn-north-2-gov-1":      dara.String("cloudauth.aliyuncs.com"),
		"cn-huhehaote":          dara.String("cloudauth.aliyuncs.com"),
		"cn-hongkong":           dara.String("cloudauth.aliyuncs.com"),
		"cn-hangzhou-finance":   dara.String("cloudauth.aliyuncs.com"),
		"cn-hangzhou":           dara.String("cloudauth.aliyuncs.com"),
		"cn-chengdu":            dara.String("cloudauth.aliyuncs.com"),
		"cn-beijing":            dara.String("cloudauth.cn-beijing.aliyuncs.com"),
		"ap-southeast-5":        dara.String("cloudauth.aliyuncs.com"),
		"ap-southeast-3":        dara.String("cloudauth.aliyuncs.com"),
		"ap-southeast-2":        dara.String("cloudauth.aliyuncs.com"),
		"ap-southeast-1":        dara.String("cloudauth.aliyuncs.com"),
		"ap-south-1":            dara.String("cloudauth.aliyuncs.com"),
		"ap-northeast-1":        dara.String("cloudauth.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("cloudauth"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Detects whether an image is generated by AIGC and returns the detection result.
//
// @param request - AIGCFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AIGCFaceVerifyResponse
func (client *Client) AIGCFaceVerifyWithOptions(request *AIGCFaceVerifyRequest, runtime *dara.RuntimeOptions) (_result *AIGCFaceVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FaceContrastPictureUrl) {
		query["FaceContrastPictureUrl"] = request.FaceContrastPictureUrl
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssObjectName) {
		query["OssObjectName"] = request.OssObjectName
	}

	if !dara.IsNil(request.OuterOrderNo) {
		query["OuterOrderNo"] = request.OuterOrderNo
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FaceContrastPicture) {
		body["FaceContrastPicture"] = request.FaceContrastPicture
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AIGCFaceVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AIGCFaceVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects whether an image is generated by AIGC and returns the detection result.
//
// @param request - AIGCFaceVerifyRequest
//
// @return AIGCFaceVerifyResponse
func (client *Client) AIGCFaceVerify(request *AIGCFaceVerifyRequest) (_result *AIGCFaceVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AIGCFaceVerifyResponse{}
	_body, _err := client.AIGCFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Provides API operations for bank card element verification, including two-element, three-element, and four-element verification.
//
// Description:
//
// Verifies bank card information consistency, including two-element verification (name + bank card number), three-element verification (name + ID card number + bank card number), and four-element verification (name + ID card number + phone number + bank card number).
//
// - Service endpoint:
//
//   - Singapore region: cloudauth.ap-southeast-1.aliyuncs.com (IPv4) or cloudauth-dualstack.ap-southeast-1.aliyuncs.com (IPv6).
//
//   - Malaysia region: cloudauth.ap-southeast-3.aliyuncs.com (IPv4) or cloudauth-dualstack.ap-southeast-3.aliyuncs.com (IPv6).
//
// - Request method: POST and GET.
//
// - Transfer protocol: HTTPS.
//
// @param request - BankMetaVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BankMetaVerifyResponse
func (client *Client) BankMetaVerifyWithOptions(request *BankMetaVerifyRequest, runtime *dara.RuntimeOptions) (_result *BankMetaVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BankCard) {
		query["BankCard"] = request.BankCard
	}

	if !dara.IsNil(request.IdentifyNum) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.IdentityType) {
		query["IdentityType"] = request.IdentityType
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		query["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.VerifyMode) {
		query["VerifyMode"] = request.VerifyMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BankMetaVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BankMetaVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides API operations for bank card element verification, including two-element, three-element, and four-element verification.
//
// Description:
//
// Verifies bank card information consistency, including two-element verification (name + bank card number), three-element verification (name + ID card number + bank card number), and four-element verification (name + ID card number + phone number + bank card number).
//
// - Service endpoint:
//
//   - Singapore region: cloudauth.ap-southeast-1.aliyuncs.com (IPv4) or cloudauth-dualstack.ap-southeast-1.aliyuncs.com (IPv6).
//
//   - Malaysia region: cloudauth.ap-southeast-3.aliyuncs.com (IPv4) or cloudauth-dualstack.ap-southeast-3.aliyuncs.com (IPv6).
//
// - Request method: POST and GET.
//
// - Transfer protocol: HTTPS.
//
// @param request - BankMetaVerifyRequest
//
// @return BankMetaVerifyResponse
func (client *Client) BankMetaVerify(request *BankMetaVerifyRequest) (_result *BankMetaVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BankMetaVerifyResponse{}
	_body, _err := client.BankMetaVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Provides a financial-grade server-side API for face comparison.
//
// Description:
//
// - API operation: CompareFaceVerify.
//
// - Endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// - Operation description: Implements ID Verification through server-side integration.
//
// #### Photo format requirements
//
// When performing face comparison, submit two face photos that meet all of the following conditions:
//
// - Recent photos or recent reference photos with a complete, clear, and unobstructed face, a natural expression, and the subject facing the camera directly.
//
// - Clear photos with normal exposure. The face must not be too dark, too bright, or have lens flare, and the angle must not deviate significantly.
//
// - Resolution must not exceed 1920×1080 and must be at least 640×480. The short side is recommended to be scaled to 720 pixels with a compression ratio greater than 0.9.
//
// - Photo size: < 1 MB.
//
// - Photos rotated 90, 180, and 270 degrees are supported. If multiple faces are detected, the largest face is selected.
//
// @param request - CompareFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompareFaceVerifyResponse
func (client *Client) CompareFaceVerifyWithOptions(request *CompareFaceVerifyRequest, runtime *dara.RuntimeOptions) (_result *CompareFaceVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Crop) {
		body["Crop"] = request.Crop
	}

	if !dara.IsNil(request.OuterOrderNo) {
		body["OuterOrderNo"] = request.OuterOrderNo
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.SourceCertifyId) {
		body["SourceCertifyId"] = request.SourceCertifyId
	}

	if !dara.IsNil(request.SourceFaceContrastPicture) {
		body["SourceFaceContrastPicture"] = request.SourceFaceContrastPicture
	}

	if !dara.IsNil(request.SourceFaceContrastPictureUrl) {
		body["SourceFaceContrastPictureUrl"] = request.SourceFaceContrastPictureUrl
	}

	if !dara.IsNil(request.SourceOssBucketName) {
		body["SourceOssBucketName"] = request.SourceOssBucketName
	}

	if !dara.IsNil(request.SourceOssObjectName) {
		body["SourceOssObjectName"] = request.SourceOssObjectName
	}

	if !dara.IsNil(request.TargetCertifyId) {
		body["TargetCertifyId"] = request.TargetCertifyId
	}

	if !dara.IsNil(request.TargetFaceContrastPicture) {
		body["TargetFaceContrastPicture"] = request.TargetFaceContrastPicture
	}

	if !dara.IsNil(request.TargetFaceContrastPictureUrl) {
		body["TargetFaceContrastPictureUrl"] = request.TargetFaceContrastPictureUrl
	}

	if !dara.IsNil(request.TargetOssBucketName) {
		body["TargetOssBucketName"] = request.TargetOssBucketName
	}

	if !dara.IsNil(request.TargetOssObjectName) {
		body["TargetOssObjectName"] = request.TargetOssObjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompareFaceVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompareFaceVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides a financial-grade server-side API for face comparison.
//
// Description:
//
// - API operation: CompareFaceVerify.
//
// - Endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// - Operation description: Implements ID Verification through server-side integration.
//
// #### Photo format requirements
//
// When performing face comparison, submit two face photos that meet all of the following conditions:
//
// - Recent photos or recent reference photos with a complete, clear, and unobstructed face, a natural expression, and the subject facing the camera directly.
//
// - Clear photos with normal exposure. The face must not be too dark, too bright, or have lens flare, and the angle must not deviate significantly.
//
// - Resolution must not exceed 1920×1080 and must be at least 640×480. The short side is recommended to be scaled to 720 pixels with a compression ratio greater than 0.9.
//
// - Photo size: < 1 MB.
//
// - Photos rotated 90, 180, and 270 degrees are supported. If multiple faces are detected, the largest face is selected.
//
// @param request - CompareFaceVerifyRequest
//
// @return CompareFaceVerifyResponse
func (client *Client) CompareFaceVerify(request *CompareFaceVerifyRequest) (_result *CompareFaceVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CompareFaceVerifyResponse{}
	_body, _err := client.CompareFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs face comparison.
//
// Description:
//
// Request method: Only HTTPS POST requests are supported.
//
// Operation description: Specifies two face images for comparison and returns a similarity score between the faces in the two images.
//
// - At least one of the specified comparison images must be of the face photo type (FacePic).
//
// - If an image contains multiple faces, the algorithm automatically selects the face that occupies the largest area in the image.
//
// - If no face is detected in one of the two comparison images, the system returns a "No face detected" error.
//
// When you submit images, you must provide the corresponding HTTP URL or Base64 encoding of each image.
//
// - HTTP URL: A publicly accessible HTTP URL. For example, `http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg`.
//
// - Base64 encoding: A Base64-encoded image in the format `base64://<Base64-encoded image string>`.
//
// Image limits.
//
// - Relative paths or absolute paths of local images are not supported.
//
// - Keep the size of each image within 2 MB to avoid algorithm retrieval timeout.
//
// - The body of a single request has a size limit of 8 MB. Make sure that the total size of all images and other information in the request does not exceed this limit.
//
// - When you use Base64 to transmit images, set the request method to POST. Remove the header description from the Base64 character string, such as `data:image/png;base64,`.
//
// @param request - CompareFacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompareFacesResponse
func (client *Client) CompareFacesWithOptions(request *CompareFacesRequest, runtime *dara.RuntimeOptions) (_result *CompareFacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceImageType) {
		body["SourceImageType"] = request.SourceImageType
	}

	if !dara.IsNil(request.SourceImageValue) {
		body["SourceImageValue"] = request.SourceImageValue
	}

	if !dara.IsNil(request.TargetImageType) {
		body["TargetImageType"] = request.TargetImageType
	}

	if !dara.IsNil(request.TargetImageValue) {
		body["TargetImageValue"] = request.TargetImageValue
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompareFaces"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompareFacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs face comparison.
//
// Description:
//
// Request method: Only HTTPS POST requests are supported.
//
// Operation description: Specifies two face images for comparison and returns a similarity score between the faces in the two images.
//
// - At least one of the specified comparison images must be of the face photo type (FacePic).
//
// - If an image contains multiple faces, the algorithm automatically selects the face that occupies the largest area in the image.
//
// - If no face is detected in one of the two comparison images, the system returns a "No face detected" error.
//
// When you submit images, you must provide the corresponding HTTP URL or Base64 encoding of each image.
//
// - HTTP URL: A publicly accessible HTTP URL. For example, `http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg`.
//
// - Base64 encoding: A Base64-encoded image in the format `base64://<Base64-encoded image string>`.
//
// Image limits.
//
// - Relative paths or absolute paths of local images are not supported.
//
// - Keep the size of each image within 2 MB to avoid algorithm retrieval timeout.
//
// - The body of a single request has a size limit of 8 MB. Make sure that the total size of all images and other information in the request does not exceed this limit.
//
// - When you use Base64 to transmit images, set the request method to POST. Remove the header description from the Base64 character string, such as `data:image/png;base64,`.
//
// @param request - CompareFacesRequest
//
// @return CompareFacesResponse
func (client *Client) CompareFaces(request *CompareFacesRequest) (_result *CompareFacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CompareFacesResponse{}
	_body, _err := client.CompareFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits verification materials for identity comparison and synchronously returns the verification result.
//
// Description:
//
// - API operation: ContrastFaceVerify.
//
// - Endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// - Operation description: An API operation that implements ID Verification through server-side integration.
//
// #### Image format requirements
//
// When performing ID Verification, submit images that meet all of the following conditions:
//
// - A recent photo with a complete, clear, and unobstructed face, a natural expression, and the subject facing the camera directly.
//
// - A clear photo with normal exposure. The face must not be too dark, too bright, or have glare, and the angle must not deviate significantly.
//
// - The resolution must not exceed 1920 × 1080 and must be at least 640 × 480. We recommend scaling the short side to 720 pixels with a compression ratio greater than 0.9.
//
// - Photo size: < 1 MB.
//
// - Photos rotated 90, 180, and 270 degrees are supported. If multiple faces are detected, the largest face is selected.
//
// @param request - ContrastFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContrastFaceVerifyResponse
func (client *Client) ContrastFaceVerifyWithOptions(request *ContrastFaceVerifyRequest, runtime *dara.RuntimeOptions) (_result *ContrastFaceVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CertName) {
		body["CertName"] = request.CertName
	}

	if !dara.IsNil(request.CertNo) {
		body["CertNo"] = request.CertNo
	}

	if !dara.IsNil(request.CertType) {
		body["CertType"] = request.CertType
	}

	if !dara.IsNil(request.CertifyId) {
		body["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.Crop) {
		body["Crop"] = request.Crop
	}

	if !dara.IsNil(request.DeviceToken) {
		body["DeviceToken"] = request.DeviceToken
	}

	if !dara.IsNil(request.EncryptType) {
		body["EncryptType"] = request.EncryptType
	}

	if !dara.IsNil(request.FaceContrastFile) {
		body["FaceContrastFile"] = request.FaceContrastFile
	}

	if !dara.IsNil(request.FaceContrastPicture) {
		body["FaceContrastPicture"] = request.FaceContrastPicture
	}

	if !dara.IsNil(request.FaceContrastPictureUrl) {
		body["FaceContrastPictureUrl"] = request.FaceContrastPictureUrl
	}

	if !dara.IsNil(request.Ip) {
		body["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.OssBucketName) {
		body["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssObjectName) {
		body["OssObjectName"] = request.OssObjectName
	}

	if !dara.IsNil(request.OuterOrderNo) {
		body["OuterOrderNo"] = request.OuterOrderNo
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContrastFaceVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ContrastFaceVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits verification materials for identity comparison and synchronously returns the verification result.
//
// Description:
//
// - API operation: ContrastFaceVerify.
//
// - Endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// - Operation description: An API operation that implements ID Verification through server-side integration.
//
// #### Image format requirements
//
// When performing ID Verification, submit images that meet all of the following conditions:
//
// - A recent photo with a complete, clear, and unobstructed face, a natural expression, and the subject facing the camera directly.
//
// - A clear photo with normal exposure. The face must not be too dark, too bright, or have glare, and the angle must not deviate significantly.
//
// - The resolution must not exceed 1920 × 1080 and must be at least 640 × 480. We recommend scaling the short side to 720 pixels with a compression ratio greater than 0.9.
//
// - Photo size: < 1 MB.
//
// - Photos rotated 90, 180, and 270 degrees are supported. If multiple faces are detected, the largest face is selected.
//
// @param request - ContrastFaceVerifyRequest
//
// @return ContrastFaceVerifyResponse
func (client *Client) ContrastFaceVerify(request *ContrastFaceVerifyRequest) (_result *ContrastFaceVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ContrastFaceVerifyResponse{}
	_body, _err := client.ContrastFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ContrastFaceVerifyAdvance(request *ContrastFaceVerifyAdvanceRequest, runtime *dara.RuntimeOptions) (_result *ContrastFaceVerifyResponse, _err error) {
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
		"Product":  dara.String("Cloudauth"),
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
	contrastFaceVerifyReq := &ContrastFaceVerifyRequest{}
	openapiutil.Convert(request, contrastFaceVerifyReq)
	if !dara.IsNil(request.FaceContrastFileObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FaceContrastFileObject,
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
		contrastFaceVerifyReq.FaceContrastFile = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	contrastFaceVerifyResp, _err := client.ContrastFaceVerifyWithOptions(contrastFaceVerifyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = contrastFaceVerifyResp
	return _result, _err
}

// Summary:
//
// Creates a network access credential.
//
// Description:
//
// Request method: Supports sending requests by using HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. Obtain a new key before each activation.
//
// @param request - CreateAntCloudAuthSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAntCloudAuthSceneResponse
func (client *Client) CreateAntCloudAuthSceneWithOptions(request *CreateAntCloudAuthSceneRequest, runtime *dara.RuntimeOptions) (_result *CreateAntCloudAuthSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BindMiniProgram) {
		query["BindMiniProgram"] = request.BindMiniProgram
	}

	if !dara.IsNil(request.CheckFileBody) {
		query["CheckFileBody"] = request.CheckFileBody
	}

	if !dara.IsNil(request.CheckFileName) {
		query["CheckFileName"] = request.CheckFileName
	}

	if !dara.IsNil(request.DeviceRiskPlus) {
		query["DeviceRiskPlus"] = request.DeviceRiskPlus
	}

	if !dara.IsNil(request.MiniProgramName) {
		query["MiniProgramName"] = request.MiniProgramName
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.ReturnPicCount) {
		query["ReturnPicCount"] = request.ReturnPicCount
	}

	if !dara.IsNil(request.ReturnVideoLength) {
		query["ReturnVideoLength"] = request.ReturnVideoLength
	}

	if !dara.IsNil(request.SceneName) {
		query["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.StoreImage) {
		query["StoreImage"] = request.StoreImage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAntCloudAuthScene"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAntCloudAuthSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a network access credential.
//
// Description:
//
// Request method: Supports sending requests by using HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. Obtain a new key before each activation.
//
// @param request - CreateAntCloudAuthSceneRequest
//
// @return CreateAntCloudAuthSceneResponse
func (client *Client) CreateAntCloudAuthScene(request *CreateAntCloudAuthSceneRequest) (_result *CreateAntCloudAuthSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAntCloudAuthSceneResponse{}
	_body, _err := client.CreateAntCloudAuthSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains an authorization key for activating the offline facial recognition SDK.
//
// Description:
//
// Request method: Supports sending requests by using HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. Obtain a new key before each activation.
//
// @param request - CreateAuthKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAuthKeyResponse
func (client *Client) CreateAuthKeyWithOptions(request *CreateAuthKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateAuthKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthYears) {
		query["AuthYears"] = request.AuthYears
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.Test) {
		query["Test"] = request.Test
	}

	if !dara.IsNil(request.UserDeviceId) {
		query["UserDeviceId"] = request.UserDeviceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAuthKey"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAuthKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains an authorization key for activating the offline facial recognition SDK.
//
// Description:
//
// Request method: Supports sending requests by using HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. Obtain a new key before each activation.
//
// @param request - CreateAuthKeyRequest
//
// @return CreateAuthKeyResponse
func (client *Client) CreateAuthKey(request *CreateAuthKeyRequest) (_result *CreateAuthKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAuthKeyResponse{}
	_body, _err := client.CreateAuthKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an authentication scenario for enhanced ID Verification.
//
// Description:
//
// Request method: Supports sending requests by using HTTPS POST and GET methods.
//
// > The authorization code is valid for 30 minutes and cannot be reused. Obtain a new authorization code before each activation.
//
// @param request - CreateCloudauthstSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudauthstSceneResponse
func (client *Client) CreateCloudauthstSceneWithOptions(request *CreateCloudauthstSceneRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudauthstSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneName) {
		query["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.StoreImage) {
		query["StoreImage"] = request.StoreImage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudauthstScene"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudauthstSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an authentication scenario for enhanced ID Verification.
//
// Description:
//
// Request method: Supports sending requests by using HTTPS POST and GET methods.
//
// > The authorization code is valid for 30 minutes and cannot be reused. Obtain a new authorization code before each activation.
//
// @param request - CreateCloudauthstSceneRequest
//
// @return CreateCloudauthstSceneResponse
func (client *Client) CreateCloudauthstScene(request *CreateCloudauthstSceneRequest) (_result *CreateCloudauthstSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudauthstSceneResponse{}
	_body, _err := client.CreateCloudauthstSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a willingness authentication configuration.
//
// Description:
//
// Request method: Send requests by using the HTTPS POST method.
//
// Request URL: cloudauth.aliyuncs.com.
//
// > The authorization key is valid for 30 minutes and cannot be reused. Obtain a new key before each activation.
//
// @param request - CreateSceneConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSceneConfigResponse
func (client *Client) CreateSceneConfigWithOptions(request *CreateSceneConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateSceneConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.SceneId) {
		body["sceneId"] = request.SceneId
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSceneConfig"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSceneConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a willingness authentication configuration.
//
// Description:
//
// Request method: Send requests by using the HTTPS POST method.
//
// Request URL: cloudauth.aliyuncs.com.
//
// > The authorization key is valid for 30 minutes and cannot be reused. Obtain a new key before each activation.
//
// @param request - CreateSceneConfigRequest
//
// @return CreateSceneConfigResponse
func (client *Client) CreateSceneConfig(request *CreateSceneConfigRequest) (_result *CreateSceneConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSceneConfigResponse{}
	_body, _err := client.CreateSceneConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an ID Verification scenario configuration. This is equivalent to creating a verification scenario in the console.
//
// Description:
//
// Request method: Only HTTPS POST requests are supported.
//
// @param request - CreateVerifySettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVerifySettingResponse
func (client *Client) CreateVerifySettingWithOptions(request *CreateVerifySettingRequest, runtime *dara.RuntimeOptions) (_result *CreateVerifySettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizName) {
		query["BizName"] = request.BizName
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.GuideStep) {
		query["GuideStep"] = request.GuideStep
	}

	if !dara.IsNil(request.PrivacyStep) {
		query["PrivacyStep"] = request.PrivacyStep
	}

	if !dara.IsNil(request.ResultStep) {
		query["ResultStep"] = request.ResultStep
	}

	if !dara.IsNil(request.Solution) {
		query["Solution"] = request.Solution
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVerifySetting"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVerifySettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an ID Verification scenario configuration. This is equivalent to creating a verification scenario in the console.
//
// Description:
//
// Request method: Only HTTPS POST requests are supported.
//
// @param request - CreateVerifySettingRequest
//
// @return CreateVerifySettingResponse
func (client *Client) CreateVerifySetting(request *CreateVerifySettingRequest) (_result *CreateVerifySettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVerifySettingResponse{}
	_body, _err := client.CreateVerifySettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an authentication whitelist.
//
// Description:
//
// Request method: Only HTTPS POST requests are supported.
//
// @param request - CreateWhitelistSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWhitelistSettingResponse
func (client *Client) CreateWhitelistSettingWithOptions(request *CreateWhitelistSettingRequest, runtime *dara.RuntimeOptions) (_result *CreateWhitelistSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertNo) {
		query["CertNo"] = request.CertNo
	}

	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.ValidDay) {
		query["ValidDay"] = request.ValidDay
	}

	if !dara.IsNil(request.WhitelistType) {
		query["WhitelistType"] = request.WhitelistType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWhitelistSetting"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWhitelistSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an authentication whitelist.
//
// Description:
//
// Request method: Only HTTPS POST requests are supported.
//
// @param request - CreateWhitelistSettingRequest
//
// @return CreateWhitelistSettingResponse
func (client *Client) CreateWhitelistSetting(request *CreateWhitelistSettingRequest) (_result *CreateWhitelistSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWhitelistSettingResponse{}
	_body, _err := client.CreateWhitelistSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// API operation for the product image tampering detection service. Detects image tampering, quality (clarity), and similar images.
//
// Description:
//
// Submits an e-commerce product image to perform tampering detection, quality (clarity) assessment, and similar image detection. Returns risk labels and risk scores.
//
// @param request - CredentialProductVerifyV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CredentialProductVerifyV2Response
func (client *Client) CredentialProductVerifyV2WithOptions(request *CredentialProductVerifyV2Request, runtime *dara.RuntimeOptions) (_result *CredentialProductVerifyV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredName) {
		query["CredName"] = request.CredName
	}

	if !dara.IsNil(request.CredType) {
		query["CredType"] = request.CredType
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.MerchantId) {
		query["MerchantId"] = request.MerchantId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageFile) {
		body["ImageFile"] = request.ImageFile
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CredentialProductVerifyV2"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CredentialProductVerifyV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// API operation for the product image tampering detection service. Detects image tampering, quality (clarity), and similar images.
//
// Description:
//
// Submits an e-commerce product image to perform tampering detection, quality (clarity) assessment, and similar image detection. Returns risk labels and risk scores.
//
// @param request - CredentialProductVerifyV2Request
//
// @return CredentialProductVerifyV2Response
func (client *Client) CredentialProductVerifyV2(request *CredentialProductVerifyV2Request) (_result *CredentialProductVerifyV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CredentialProductVerifyV2Response{}
	_body, _err := client.CredentialProductVerifyV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CredentialProductVerifyV2Advance(request *CredentialProductVerifyV2AdvanceRequest, runtime *dara.RuntimeOptions) (_result *CredentialProductVerifyV2Response, _err error) {
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
		"Product":  dara.String("Cloudauth"),
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
	credentialProductVerifyV2Req := &CredentialProductVerifyV2Request{}
	openapiutil.Convert(request, credentialProductVerifyV2Req)
	if !dara.IsNil(request.ImageFileObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.ImageFileObject,
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
		credentialProductVerifyV2Req.ImageFile = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	credentialProductVerifyV2Resp, _err := client.CredentialProductVerifyV2WithOptions(credentialProductVerifyV2Req, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = credentialProductVerifyV2Resp
	return _result, _err
}

// Summary:
//
// Credential verification.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com (IPv4) or cloudauth-dualstack.aliyuncs.com (IPv6).
//
// - Request method: POST and GET.
//
// - Transfer protocol: HTTPS.
//
// @param tmpReq - CredentialVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CredentialVerifyResponse
func (client *Client) CredentialVerifyWithOptions(tmpReq *CredentialVerifyRequest, runtime *dara.RuntimeOptions) (_result *CredentialVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CredentialVerifyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MerchantDetail) {
		request.MerchantDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MerchantDetail, dara.String("MerchantDetail"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CertNum) {
		query["CertNum"] = request.CertNum
	}

	if !dara.IsNil(request.CredName) {
		query["CredName"] = request.CredName
	}

	if !dara.IsNil(request.CredType) {
		query["CredType"] = request.CredType
	}

	if !dara.IsNil(request.IdentifyNum) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.IsCheck) {
		query["IsCheck"] = request.IsCheck
	}

	if !dara.IsNil(request.IsOCR) {
		query["IsOCR"] = request.IsOCR
	}

	if !dara.IsNil(request.MerchantDetailShrink) {
		query["MerchantDetail"] = request.MerchantDetailShrink
	}

	if !dara.IsNil(request.MerchantId) {
		query["MerchantId"] = request.MerchantId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.PromptModel) {
		query["PromptModel"] = request.PromptModel
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageContext) {
		body["ImageContext"] = request.ImageContext
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CredentialVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CredentialVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Credential verification.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com (IPv4) or cloudauth-dualstack.aliyuncs.com (IPv6).
//
// - Request method: POST and GET.
//
// - Transfer protocol: HTTPS.
//
// @param request - CredentialVerifyRequest
//
// @return CredentialVerifyResponse
func (client *Client) CredentialVerify(request *CredentialVerifyRequest) (_result *CredentialVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CredentialVerifyResponse{}
	_body, _err := client.CredentialVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Credential verification.
//
// Description:
//
// Submits a credential image to perform image tampering and forgery detection and image semantic understanding. Returns risk detection results.
//
// @param tmpReq - CredentialVerifyV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CredentialVerifyV2Response
func (client *Client) CredentialVerifyV2WithOptions(tmpReq *CredentialVerifyV2Request, runtime *dara.RuntimeOptions) (_result *CredentialVerifyV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CredentialVerifyV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MerchantDetail) {
		request.MerchantDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MerchantDetail, dara.String("MerchantDetail"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CertNum) {
		query["CertNum"] = request.CertNum
	}

	if !dara.IsNil(request.CredName) {
		query["CredName"] = request.CredName
	}

	if !dara.IsNil(request.CredType) {
		query["CredType"] = request.CredType
	}

	if !dara.IsNil(request.IdentifyNum) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.IsCheck) {
		query["IsCheck"] = request.IsCheck
	}

	if !dara.IsNil(request.IsOcr) {
		query["IsOcr"] = request.IsOcr
	}

	if !dara.IsNil(request.MerchantDetailShrink) {
		query["MerchantDetail"] = request.MerchantDetailShrink
	}

	if !dara.IsNil(request.MerchantId) {
		query["MerchantId"] = request.MerchantId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.PromptModel) {
		query["PromptModel"] = request.PromptModel
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageContext) {
		body["ImageContext"] = request.ImageContext
	}

	if !dara.IsNil(request.ImageFile) {
		body["ImageFile"] = request.ImageFile
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CredentialVerifyV2"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CredentialVerifyV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Credential verification.
//
// Description:
//
// Submits a credential image to perform image tampering and forgery detection and image semantic understanding. Returns risk detection results.
//
// @param request - CredentialVerifyV2Request
//
// @return CredentialVerifyV2Response
func (client *Client) CredentialVerifyV2(request *CredentialVerifyV2Request) (_result *CredentialVerifyV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CredentialVerifyV2Response{}
	_body, _err := client.CredentialVerifyV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CredentialVerifyV2Advance(request *CredentialVerifyV2AdvanceRequest, runtime *dara.RuntimeOptions) (_result *CredentialVerifyV2Response, _err error) {
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
		"Product":  dara.String("Cloudauth"),
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
	credentialVerifyV2Req := &CredentialVerifyV2Request{}
	openapiutil.Convert(request, credentialVerifyV2Req)
	if !dara.IsNil(request.ImageFileObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.ImageFileObject,
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
		credentialVerifyV2Req.ImageFile = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	credentialVerifyV2Resp, _err := client.CredentialVerifyV2WithOptions(credentialVerifyV2Req, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = credentialVerifyV2Resp
	return _result, _err
}

// Summary:
//
// Accepts a face image and uses algorithms to detect deepfake risks. Covers risk scenarios such as AIGC-generated faces, deepfake face swaps, template faces, and recaptured faces. Returns risk labels and confidence levels.
//
// Description:
//
// > The face deepfake detection operation is currently in free public preview. The free public preview ends at 23:59:59 on August 30, 2024. During the public preview, the QPS cannot exceed 3 queries per second.
//
// - Service endpoint: cloudauth.aliyuncs.com (IPv4) or cloudauth-dualstack.aliyuncs.com (IPv6).
//
// - Request method: POST and GET.
//
// - Transfer protocol: HTTPS.
//
// @param request - DeepfakeDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeepfakeDetectResponse
func (client *Client) DeepfakeDetectWithOptions(request *DeepfakeDetectRequest, runtime *dara.RuntimeOptions) (_result *DeepfakeDetectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FaceInputType) {
		query["FaceInputType"] = request.FaceInputType
	}

	if !dara.IsNil(request.FaceUrl) {
		query["FaceUrl"] = request.FaceUrl
	}

	if !dara.IsNil(request.OuterOrderNo) {
		query["OuterOrderNo"] = request.OuterOrderNo
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FaceBase64) {
		body["FaceBase64"] = request.FaceBase64
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeepfakeDetect"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeepfakeDetectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Accepts a face image and uses algorithms to detect deepfake risks. Covers risk scenarios such as AIGC-generated faces, deepfake face swaps, template faces, and recaptured faces. Returns risk labels and confidence levels.
//
// Description:
//
// > The face deepfake detection operation is currently in free public preview. The free public preview ends at 23:59:59 on August 30, 2024. During the public preview, the QPS cannot exceed 3 queries per second.
//
// - Service endpoint: cloudauth.aliyuncs.com (IPv4) or cloudauth-dualstack.aliyuncs.com (IPv6).
//
// - Request method: POST and GET.
//
// - Transfer protocol: HTTPS.
//
// @param request - DeepfakeDetectRequest
//
// @return DeepfakeDetectResponse
func (client *Client) DeepfakeDetect(request *DeepfakeDetectRequest) (_result *DeepfakeDetectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeepfakeDetectResponse{}
	_body, _err := client.DeepfakeDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes all custom rate limiting policies.
//
// Description:
//
// Request method: Supports sending requests by using the HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. Obtain a new key before each activation.
//
// @param request - DeleteAllCustomizeFlowStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAllCustomizeFlowStrategyResponse
func (client *Client) DeleteAllCustomizeFlowStrategyWithOptions(request *DeleteAllCustomizeFlowStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteAllCustomizeFlowStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAllCustomizeFlowStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAllCustomizeFlowStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes all custom rate limiting policies.
//
// Description:
//
// Request method: Supports sending requests by using the HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. Obtain a new key before each activation.
//
// @param request - DeleteAllCustomizeFlowStrategyRequest
//
// @return DeleteAllCustomizeFlowStrategyResponse
func (client *Client) DeleteAllCustomizeFlowStrategy(request *DeleteAllCustomizeFlowStrategyRequest) (_result *DeleteAllCustomizeFlowStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAllCustomizeFlowStrategyResponse{}
	_body, _err := client.DeleteAllCustomizeFlowStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a China Finance Verification scene.
//
// Description:
//
// - Endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DeleteAntCloudAuthSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAntCloudAuthSceneResponse
func (client *Client) DeleteAntCloudAuthSceneWithOptions(request *DeleteAntCloudAuthSceneRequest, runtime *dara.RuntimeOptions) (_result *DeleteAntCloudAuthSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAntCloudAuthScene"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAntCloudAuthSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a China Finance Verification scene.
//
// Description:
//
// - Endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DeleteAntCloudAuthSceneRequest
//
// @return DeleteAntCloudAuthSceneResponse
func (client *Client) DeleteAntCloudAuthScene(request *DeleteAntCloudAuthSceneRequest) (_result *DeleteAntCloudAuthSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAntCloudAuthSceneResponse{}
	_body, _err := client.DeleteAntCloudAuthSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a blacklist rule.
//
// Description:
//
// Request method: Only HTTPS POST requests are supported.
//
// @param request - DeleteBlackListStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBlackListStrategyResponse
func (client *Client) DeleteBlackListStrategyWithOptions(request *DeleteBlackListStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteBlackListStrategyResponse, _err error) {
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

	if !dara.IsNil(request.ProductName) {
		query["ProductName"] = request.ProductName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBlackListStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBlackListStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a blacklist rule.
//
// Description:
//
// Request method: Only HTTPS POST requests are supported.
//
// @param request - DeleteBlackListStrategyRequest
//
// @return DeleteBlackListStrategyResponse
func (client *Client) DeleteBlackListStrategy(request *DeleteBlackListStrategyRequest) (_result *DeleteBlackListStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBlackListStrategyResponse{}
	_body, _err := client.DeleteBlackListStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete a specified authentication scene for Enhanced Real-person Identity Verification
//
// Description:
//
// Request method: Supports sending requests using HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. We recommend that you obtain a new key before each activation.
//
// @param request - DeleteCloudauthstSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudauthstSceneResponse
func (client *Client) DeleteCloudauthstSceneWithOptions(request *DeleteCloudauthstSceneRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudauthstSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudauthstScene"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudauthstSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete a specified authentication scene for Enhanced Real-person Identity Verification
//
// Description:
//
// Request method: Supports sending requests using HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. We recommend that you obtain a new key before each activation.
//
// @param request - DeleteCloudauthstSceneRequest
//
// @return DeleteCloudauthstSceneResponse
func (client *Client) DeleteCloudauthstScene(request *DeleteCloudauthstSceneRequest) (_result *DeleteCloudauthstSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloudauthstSceneResponse{}
	_body, _err := client.DeleteCloudauthstSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete a stability alert rule.
//
// Description:
//
// Request method: Supports sending requests using the HTTPS POST method.
//
// Request URL: cloudauth.aliyuncs.com.
//
// @param request - DeleteControlStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteControlStrategyResponse
func (client *Client) DeleteControlStrategyWithOptions(request *DeleteControlStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteControlStrategyResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteControlStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteControlStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a stability alert rule.
//
// Description:
//
// Request method: Supports sending requests using the HTTPS POST method.
//
// Request URL: cloudauth.aliyuncs.com.
//
// @param request - DeleteControlStrategyRequest
//
// @return DeleteControlStrategyResponse
func (client *Client) DeleteControlStrategy(request *DeleteControlStrategyRequest) (_result *DeleteControlStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteControlStrategyResponse{}
	_body, _err := client.DeleteControlStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom rate limiting policy.
//
// Description:
//
// Request method: Supports sending requests by using HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. Obtain a new key before each activation.
//
// @param request - DeleteCustomizeFlowStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomizeFlowStrategyResponse
func (client *Client) DeleteCustomizeFlowStrategyWithOptions(request *DeleteCustomizeFlowStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomizeFlowStrategyResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomizeFlowStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomizeFlowStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom rate limiting policy.
//
// Description:
//
// Request method: Supports sending requests by using HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. Obtain a new key before each activation.
//
// @param request - DeleteCustomizeFlowStrategyRequest
//
// @return DeleteCustomizeFlowStrategyResponse
func (client *Client) DeleteCustomizeFlowStrategy(request *DeleteCustomizeFlowStrategyRequest) (_result *DeleteCustomizeFlowStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomizeFlowStrategyResponse{}
	_body, _err := client.DeleteCustomizeFlowStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// API for deleting sensitive data in financial-grade services.
//
// Description:
//
// Deletes all personal information fields from the request, including name, ID card number, phone number, IP address, images, videos, device information, etc.
//
// @param request - DeleteFaceVerifyResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaceVerifyResultResponse
func (client *Client) DeleteFaceVerifyResultWithOptions(request *DeleteFaceVerifyResultRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceVerifyResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.DeleteAfterQuery) {
		query["DeleteAfterQuery"] = request.DeleteAfterQuery
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFaceVerifyResult"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFaceVerifyResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// API for deleting sensitive data in financial-grade services.
//
// Description:
//
// Deletes all personal information fields from the request, including name, ID card number, phone number, IP address, images, videos, device information, etc.
//
// @param request - DeleteFaceVerifyResultRequest
//
// @return DeleteFaceVerifyResultResponse
func (client *Client) DeleteFaceVerifyResult(request *DeleteFaceVerifyResultRequest) (_result *DeleteFaceVerifyResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFaceVerifyResultResponse{}
	_body, _err := client.DeleteFaceVerifyResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a scenario configuration.
//
// Description:
//
// - Request method: HTTPS POST and GET methods are supported.
//
// - Request URL: cloudauth.aliyuncs.com.
//
// > The authorization code is valid for 30 minutes and cannot be reused. Obtain a new authorization code before each activation.
//
// @param request - DeleteSceneConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSceneConfigResponse
func (client *Client) DeleteSceneConfigWithOptions(request *DeleteSceneConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteSceneConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SceneConfigId) {
		body["sceneConfigId"] = request.SceneConfigId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSceneConfig"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSceneConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scenario configuration.
//
// Description:
//
// - Request method: HTTPS POST and GET methods are supported.
//
// - Request URL: cloudauth.aliyuncs.com.
//
// > The authorization code is valid for 30 minutes and cannot be reused. Obtain a new authorization code before each activation.
//
// @param request - DeleteSceneConfigRequest
//
// @return DeleteSceneConfigResponse
func (client *Client) DeleteSceneConfig(request *DeleteSceneConfigRequest) (_result *DeleteSceneConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSceneConfigResponse{}
	_body, _err := client.DeleteSceneConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the version of a specified cluster.
//
// Description:
//
// Request method: Only HTTPS POST requests are supported.
//
// @param request - DeleteWhitelistSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWhitelistSettingResponse
func (client *Client) DeleteWhitelistSettingWithOptions(request *DeleteWhitelistSettingRequest, runtime *dara.RuntimeOptions) (_result *DeleteWhitelistSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWhitelistSetting"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWhitelistSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the version of a specified cluster.
//
// Description:
//
// Request method: Only HTTPS POST requests are supported.
//
// @param request - DeleteWhitelistSettingRequest
//
// @return DeleteWhitelistSettingResponse
func (client *Client) DeleteWhitelistSetting(request *DeleteWhitelistSettingRequest) (_result *DeleteWhitelistSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWhitelistSettingResponse{}
	_body, _err := client.DeleteWhitelistSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the activation status of different ID Verification product plans.
//
// Description:
//
// Request method: Supports sending requests by using HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. Obtain a new key before each activation.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntAndCloudAuthUserStatusResponse
func (client *Client) DescribeAntAndCloudAuthUserStatusWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeAntAndCloudAuthUserStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAntAndCloudAuthUserStatus"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAntAndCloudAuthUserStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the activation status of different ID Verification product plans.
//
// Description:
//
// Request method: Supports sending requests by using HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. Obtain a new key before each activation.
//
// @return DescribeAntAndCloudAuthUserStatusResponse
func (client *Client) DescribeAntAndCloudAuthUserStatus() (_result *DescribeAntAndCloudAuthUserStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAntAndCloudAuthUserStatusResponse{}
	_body, _err := client.DescribeAntAndCloudAuthUserStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves OCR results.
//
// @param request - DescribeAuthVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuthVerifyResponse
func (client *Client) DescribeAuthVerifyWithOptions(request *DescribeAuthVerifyRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuthVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CertifyId) {
		body["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuthVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuthVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves OCR results.
//
// @param request - DescribeAuthVerifyRequest
//
// @return DescribeAuthVerifyResponse
func (client *Client) DescribeAuthVerify(request *DescribeAuthVerifyRequest) (_result *DescribeAuthVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAuthVerifyResponse{}
	_body, _err := client.DescribeAuthVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the verification result for image element verification.
//
// Description:
//
// After you receive a callback notification, you can call this operation on the server side to obtain the corresponding verification status and verification materials.
//
// @param request - DescribeCardVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCardVerifyResponse
func (client *Client) DescribeCardVerifyWithOptions(request *DescribeCardVerifyRequest, runtime *dara.RuntimeOptions) (_result *DescribeCardVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCardVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCardVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the verification result for image element verification.
//
// Description:
//
// After you receive a callback notification, you can call this operation on the server side to obtain the corresponding verification status and verification materials.
//
// @param request - DescribeCardVerifyRequest
//
// @return DescribeCardVerifyResponse
func (client *Client) DescribeCardVerify(request *DescribeCardVerifyRequest) (_result *DescribeCardVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCardVerifyResponse{}
	_body, _err := client.DescribeCardVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the authentication scenarios for enhanced ID Verification.
//
// Description:
//
// Request method: Supports sending requests by using HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. Obtain a new key before each activation.
//
// @param request - DescribeCloudauthstSceneListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudauthstSceneListResponse
func (client *Client) DescribeCloudauthstSceneListWithOptions(request *DescribeCloudauthstSceneListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudauthstSceneListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudauthstSceneList"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudauthstSceneListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authentication scenarios for enhanced ID Verification.
//
// Description:
//
// Request method: Supports sending requests by using HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. Obtain a new key before each activation.
//
// @param request - DescribeCloudauthstSceneListRequest
//
// @return DescribeCloudauthstSceneListResponse
func (client *Client) DescribeCloudauthstSceneList(request *DescribeCloudauthstSceneListRequest) (_result *DescribeCloudauthstSceneListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudauthstSceneListResponse{}
	_body, _err := client.DescribeCloudauthstSceneListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Such as authorization validity period, custom business identifiers defined by the integrating party, and device IDs.
//
// Description:
//
// Request method: supports sending requests using HTTPS POST and GET methods.
//
// @param request - DescribeDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceInfoResponse
func (client *Client) DescribeDeviceInfoWithOptions(request *DescribeDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeviceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.ExpiredEndDay) {
		query["ExpiredEndDay"] = request.ExpiredEndDay
	}

	if !dara.IsNil(request.ExpiredStartDay) {
		query["ExpiredStartDay"] = request.ExpiredStartDay
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserDeviceId) {
		query["UserDeviceId"] = request.UserDeviceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeviceInfo"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Such as authorization validity period, custom business identifiers defined by the integrating party, and device IDs.
//
// Description:
//
// Request method: supports sending requests using HTTPS POST and GET methods.
//
// @param request - DescribeDeviceInfoRequest
//
// @return DescribeDeviceInfoResponse
func (client *Client) DescribeDeviceInfo(request *DescribeDeviceInfoRequest) (_result *DescribeDeviceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDeviceInfoResponse{}
	_body, _err := client.DescribeDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Financial-grade Face Guard service.
//
// @param request - DescribeFaceGuardRiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaceGuardRiskResponse
func (client *Client) DescribeFaceGuardRiskWithOptions(request *DescribeFaceGuardRiskRequest, runtime *dara.RuntimeOptions) (_result *DescribeFaceGuardRiskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DeviceToken) {
		query["DeviceToken"] = request.DeviceToken
	}

	if !dara.IsNil(request.OuterOrderNo) {
		query["OuterOrderNo"] = request.OuterOrderNo
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFaceGuardRisk"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFaceGuardRiskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Financial-grade Face Guard service.
//
// @param request - DescribeFaceGuardRiskRequest
//
// @return DescribeFaceGuardRiskResponse
func (client *Client) DescribeFaceGuardRisk(request *DescribeFaceGuardRiskRequest) (_result *DescribeFaceGuardRiskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFaceGuardRiskResponse{}
	_body, _err := client.DescribeFaceGuardRiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// After the China site (Chinese mainland) mobile client receives a callback, the China site (Chinese mainland) server can call this operation to obtain the corresponding verification status and verification materials.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaceVerifyResponse
func (client *Client) DescribeFaceVerifyWithOptions(request *DescribeFaceVerifyRequest, runtime *dara.RuntimeOptions) (_result *DescribeFaceVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.PictureReturnType) {
		query["PictureReturnType"] = request.PictureReturnType
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFaceVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFaceVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// After the China site (Chinese mainland) mobile client receives a callback, the China site (Chinese mainland) server can call this operation to obtain the corresponding verification status and verification materials.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeFaceVerifyRequest
//
// @return DescribeFaceVerifyResponse
func (client *Client) DescribeFaceVerify(request *DescribeFaceVerifyRequest) (_result *DescribeFaceVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFaceVerifyResponse{}
	_body, _err := client.DescribeFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information verification export tasks by page.
//
// Description:
//
// Request method: Supports sending requests by using the HTTPS POST and GET methods.
//
// @param request - DescribeInfoCheckExportRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInfoCheckExportRecordResponse
func (client *Client) DescribeInfoCheckExportRecordWithOptions(request *DescribeInfoCheckExportRecordRequest, runtime *dara.RuntimeOptions) (_result *DescribeInfoCheckExportRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInfoCheckExportRecord"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInfoCheckExportRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information verification export tasks by page.
//
// Description:
//
// Request method: Supports sending requests by using the HTTPS POST and GET methods.
//
// @param request - DescribeInfoCheckExportRecordRequest
//
// @return DescribeInfoCheckExportRecordResponse
func (client *Client) DescribeInfoCheckExportRecord(request *DescribeInfoCheckExportRecordRequest) (_result *DescribeInfoCheckExportRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInfoCheckExportRecordResponse{}
	_body, _err := client.DescribeInfoCheckExportRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a specified.
//
// Description:
//
// Request method: Supports HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. Obtain a new key before each activation.
//
// @param request - DescribeListAntCloudAuthScenesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeListAntCloudAuthScenesResponse
func (client *Client) DescribeListAntCloudAuthScenesWithOptions(request *DescribeListAntCloudAuthScenesRequest, runtime *dara.RuntimeOptions) (_result *DescribeListAntCloudAuthScenesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeListAntCloudAuthScenes"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeListAntCloudAuthScenesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a specified.
//
// Description:
//
// Request method: Supports HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. Obtain a new key before each activation.
//
// @param request - DescribeListAntCloudAuthScenesRequest
//
// @return DescribeListAntCloudAuthScenesResponse
func (client *Client) DescribeListAntCloudAuthScenes(request *DescribeListAntCloudAuthScenesRequest) (_result *DescribeListAntCloudAuthScenesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeListAntCloudAuthScenesResponse{}
	_body, _err := client.DescribeListAntCloudAuthScenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of facial recognition data.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeListFaceVerifyDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeListFaceVerifyDataResponse
func (client *Client) DescribeListFaceVerifyDataWithOptions(request *DescribeListFaceVerifyDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeListFaceVerifyDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GmtEnd) {
		query["GmtEnd"] = request.GmtEnd
	}

	if !dara.IsNil(request.GmtStart) {
		query["GmtStart"] = request.GmtStart
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeListFaceVerifyData"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeListFaceVerifyDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of facial recognition data.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeListFaceVerifyDataRequest
//
// @return DescribeListFaceVerifyDataResponse
func (client *Client) DescribeListFaceVerifyData(request *DescribeListFaceVerifyDataRequest) (_result *DescribeListFaceVerifyDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeListFaceVerifyDataResponse{}
	_body, _err := client.DescribeListFaceVerifyDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves facial recognition authentication records (legacy).
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeListFaceVerifyInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeListFaceVerifyInfosResponse
func (client *Client) DescribeListFaceVerifyInfosWithOptions(request *DescribeListFaceVerifyInfosRequest, runtime *dara.RuntimeOptions) (_result *DescribeListFaceVerifyInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.GmtEnd) {
		query["GmtEnd"] = request.GmtEnd
	}

	if !dara.IsNil(request.GmtStart) {
		query["GmtStart"] = request.GmtStart
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeListFaceVerifyInfos"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeListFaceVerifyInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves facial recognition authentication records (legacy).
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeListFaceVerifyInfosRequest
//
// @return DescribeListFaceVerifyInfosResponse
func (client *Client) DescribeListFaceVerifyInfos(request *DescribeListFaceVerifyInfosRequest) (_result *DescribeListFaceVerifyInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeListFaceVerifyInfosResponse{}
	_body, _err := client.DescribeListFaceVerifyInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information verification details by paging.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeMetaSearchPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetaSearchPageListResponse
func (client *Client) DescribeMetaSearchPageListWithOptions(request *DescribeMetaSearchPageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetaSearchPageListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Api) {
		query["Api"] = request.Api
	}

	if !dara.IsNil(request.BankCard) {
		query["BankCard"] = request.BankCard
	}

	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.IdentifyNum) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.IspName) {
		query["IspName"] = request.IspName
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReqId) {
		query["ReqId"] = request.ReqId
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SubCode) {
		query["SubCode"] = request.SubCode
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.VehicleNum) {
		query["VehicleNum"] = request.VehicleNum
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetaSearchPageList"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetaSearchPageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information verification details by paging.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeMetaSearchPageListRequest
//
// @return DescribeMetaSearchPageListResponse
func (client *Client) DescribeMetaSearchPageList(request *DescribeMetaSearchPageListRequest) (_result *DescribeMetaSearchPageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetaSearchPageListResponse{}
	_body, _err := client.DescribeMetaSearchPageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries statistics information for information verification and authentication.
//
// Description:
//
// - Request method: HTTPS POST and GET methods are supported.
//
// - Service address: cloudauth.aliyuncs.com.
//
// @param request - DescribeMetaStatisticsListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetaStatisticsListResponse
func (client *Client) DescribeMetaStatisticsListWithOptions(request *DescribeMetaStatisticsListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetaStatisticsListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Api) {
		query["Api"] = request.Api
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetaStatisticsList"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetaStatisticsListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries statistics information for information verification and authentication.
//
// Description:
//
// - Request method: HTTPS POST and GET methods are supported.
//
// - Service address: cloudauth.aliyuncs.com.
//
// @param request - DescribeMetaStatisticsListRequest
//
// @return DescribeMetaStatisticsListResponse
func (client *Client) DescribeMetaStatisticsList(request *DescribeMetaStatisticsListRequest) (_result *DescribeMetaStatisticsListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetaStatisticsListResponse{}
	_body, _err := client.DescribeMetaStatisticsListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information verification and authentication data with pagination.
//
// Description:
//
// - Request method: Supports sending requests using HTTPS POST and GET methods.
//
// - Service address: cloudauth.aliyuncs.com.
//
// @param request - DescribeMetaStatisticsPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetaStatisticsPageListResponse
func (client *Client) DescribeMetaStatisticsPageListWithOptions(request *DescribeMetaStatisticsPageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetaStatisticsPageListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Api) {
		query["Api"] = request.Api
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetaStatisticsPageList"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetaStatisticsPageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information verification and authentication data with pagination.
//
// Description:
//
// - Request method: Supports sending requests using HTTPS POST and GET methods.
//
// - Service address: cloudauth.aliyuncs.com.
//
// @param request - DescribeMetaStatisticsPageListRequest
//
// @return DescribeMetaStatisticsPageListResponse
func (client *Client) DescribeMetaStatisticsPageList(request *DescribeMetaStatisticsPageListRequest) (_result *DescribeMetaStatisticsPageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetaStatisticsPageListResponse{}
	_body, _err := client.DescribeMetaStatisticsPageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of OSS.
//
// Description:
//
// - Request method: HTTPS POST and GET methods are supported.
//
// - Endpoint: cloudauth.aliyuncs.com.
//
// @param request - DescribeOssStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssStatusResponse
func (client *Client) DescribeOssStatusWithOptions(request *DescribeOssStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeOssStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOssStatus"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOssStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of OSS.
//
// Description:
//
// - Request method: HTTPS POST and GET methods are supported.
//
// - Endpoint: cloudauth.aliyuncs.com.
//
// @param request - DescribeOssStatusRequest
//
// @return DescribeOssStatusResponse
func (client *Client) DescribeOssStatus(request *DescribeOssStatusRequest) (_result *DescribeOssStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOssStatusResponse{}
	_body, _err := client.DescribeOssStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the OSS status of a user (V2).
//
// Description:
//
// - Request method: HTTPS POST and GET methods are supported.
//
// - Service address: cloudauth.aliyuncs.com.
//
// @param request - DescribeOssStatusV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssStatusV2Response
func (client *Client) DescribeOssStatusV2WithOptions(request *DescribeOssStatusV2Request, runtime *dara.RuntimeOptions) (_result *DescribeOssStatusV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOssStatusV2"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOssStatusV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the OSS status of a user (V2).
//
// Description:
//
// - Request method: HTTPS POST and GET methods are supported.
//
// - Service address: cloudauth.aliyuncs.com.
//
// @param request - DescribeOssStatusV2Request
//
// @return DescribeOssStatusV2Response
func (client *Client) DescribeOssStatusV2(request *DescribeOssStatusV2Request) (_result *DescribeOssStatusV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOssStatusV2Response{}
	_body, _err := client.DescribeOssStatusV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the token required for uploading photos to OSS.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssUploadTokenResponse
func (client *Client) DescribeOssUploadTokenWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeOssUploadTokenResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOssUploadToken"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOssUploadTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the token required for uploading photos to OSS.
//
// @return DescribeOssUploadTokenResponse
func (client *Client) DescribeOssUploadToken() (_result *DescribeOssUploadTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOssUploadTokenResponse{}
	_body, _err := client.DescribeOssUploadTokenWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries financial-grade ID Verification call statistics by using a paging query operation.
//
// @param request - DescribePageFaceVerifyDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePageFaceVerifyDataResponse
func (client *Client) DescribePageFaceVerifyDataWithOptions(request *DescribePageFaceVerifyDataRequest, runtime *dara.RuntimeOptions) (_result *DescribePageFaceVerifyDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePageFaceVerifyData"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePageFaceVerifyDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries financial-grade ID Verification call statistics by using a paging query operation.
//
// @param request - DescribePageFaceVerifyDataRequest
//
// @return DescribePageFaceVerifyDataResponse
func (client *Client) DescribePageFaceVerifyData(request *DescribePageFaceVerifyDataRequest) (_result *DescribePageFaceVerifyDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePageFaceVerifyDataResponse{}
	_body, _err := client.DescribePageFaceVerifyDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries page settings and returns the reasons for authentication failure.
//
// Description:
//
// Request method: Only HTTPS POST requests are supported.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePageSettingResponse
func (client *Client) DescribePageSettingWithOptions(runtime *dara.RuntimeOptions) (_result *DescribePageSettingResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePageSetting"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePageSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries page settings and returns the reasons for authentication failure.
//
// Description:
//
// Request method: Only HTTPS POST requests are supported.
//
// @return DescribePageSettingResponse
func (client *Client) DescribePageSetting() (_result *DescribePageSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePageSettingResponse{}
	_body, _err := client.DescribePageSettingWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the product codes for financial-grade ID Verification.
//
// Description:
//
// Request method: Send requests by using the HTTPS GET or POST method.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProductCodeResponse
func (client *Client) DescribeProductCodeWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeProductCodeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProductCode"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProductCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the product codes for financial-grade ID Verification.
//
// Description:
//
// Request method: Send requests by using the HTTPS GET or POST method.
//
// @return DescribeProductCodeResponse
func (client *Client) DescribeProductCode() (_result *DescribeProductCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProductCodeResponse{}
	_body, _err := client.DescribeProductCodeWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the invoke statistics of enhanced ID Verification by using a paged query.
//
// @param request - DescribeSmartStatisticsPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSmartStatisticsPageListResponse
func (client *Client) DescribeSmartStatisticsPageListWithOptions(request *DescribeSmartStatisticsPageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSmartStatisticsPageListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSmartStatisticsPageList"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSmartStatisticsPageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the invoke statistics of enhanced ID Verification by using a paged query.
//
// @param request - DescribeSmartStatisticsPageListRequest
//
// @return DescribeSmartStatisticsPageListResponse
func (client *Client) DescribeSmartStatisticsPageList(request *DescribeSmartStatisticsPageListRequest) (_result *DescribeSmartStatisticsPageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSmartStatisticsPageListResponse{}
	_body, _err := client.DescribeSmartStatisticsPageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves statistics of verification devices.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeVerifyDeviceRiskStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyDeviceRiskStatisticsResponse
func (client *Client) DescribeVerifyDeviceRiskStatisticsWithOptions(request *DescribeVerifyDeviceRiskStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyDeviceRiskStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyDeviceRiskStatistics"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyDeviceRiskStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves statistics of verification devices.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeVerifyDeviceRiskStatisticsRequest
//
// @return DescribeVerifyDeviceRiskStatisticsResponse
func (client *Client) DescribeVerifyDeviceRiskStatistics(request *DescribeVerifyDeviceRiskStatisticsRequest) (_result *DescribeVerifyDeviceRiskStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVerifyDeviceRiskStatisticsResponse{}
	_body, _err := client.DescribeVerifyDeviceRiskStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the total number of failed authentication requests.
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeVerifyFailStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyFailStatisticsResponse
func (client *Client) DescribeVerifyFailStatisticsWithOptions(request *DescribeVerifyFailStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyFailStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgeGt) {
		query["AgeGt"] = request.AgeGt
	}

	if !dara.IsNil(request.Api) {
		query["Api"] = request.Api
	}

	if !dara.IsNil(request.DeviceType) {
		query["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyFailStatistics"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyFailStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total number of failed authentication requests.
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeVerifyFailStatisticsRequest
//
// @return DescribeVerifyFailStatisticsResponse
func (client *Client) DescribeVerifyFailStatistics(request *DescribeVerifyFailStatisticsRequest) (_result *DescribeVerifyFailStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVerifyFailStatisticsResponse{}
	_body, _err := client.DescribeVerifyFailStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the distribution data of phone models used by authenticated users.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeVerifyPersonasDeviceModelStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyPersonasDeviceModelStatisticsResponse
func (client *Client) DescribeVerifyPersonasDeviceModelStatisticsWithOptions(request *DescribeVerifyPersonasDeviceModelStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyPersonasDeviceModelStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.TimeRange) {
		query["TimeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyPersonasDeviceModelStatistics"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyPersonasDeviceModelStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the distribution data of phone models used by authenticated users.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeVerifyPersonasDeviceModelStatisticsRequest
//
// @return DescribeVerifyPersonasDeviceModelStatisticsResponse
func (client *Client) DescribeVerifyPersonasDeviceModelStatistics(request *DescribeVerifyPersonasDeviceModelStatisticsRequest) (_result *DescribeVerifyPersonasDeviceModelStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVerifyPersonasDeviceModelStatisticsResponse{}
	_body, _err := client.DescribeVerifyPersonasDeviceModelStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the distribution data of ID Verification devices.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeVerifyPersonasOsStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyPersonasOsStatisticsResponse
func (client *Client) DescribeVerifyPersonasOsStatisticsWithOptions(request *DescribeVerifyPersonasOsStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyPersonasOsStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.TimeRange) {
		query["TimeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyPersonasOsStatistics"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyPersonasOsStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the distribution data of ID Verification devices.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeVerifyPersonasOsStatisticsRequest
//
// @return DescribeVerifyPersonasOsStatisticsResponse
func (client *Client) DescribeVerifyPersonasOsStatistics(request *DescribeVerifyPersonasOsStatisticsRequest) (_result *DescribeVerifyPersonasOsStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVerifyPersonasOsStatisticsResponse{}
	_body, _err := client.DescribeVerifyPersonasOsStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries authentication statistics by province of the individual.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeVerifyPersonasProvinceStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyPersonasProvinceStatisticsResponse
func (client *Client) DescribeVerifyPersonasProvinceStatisticsWithOptions(request *DescribeVerifyPersonasProvinceStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyPersonasProvinceStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.TimeRange) {
		query["TimeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyPersonasProvinceStatistics"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyPersonasProvinceStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries authentication statistics by province of the individual.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeVerifyPersonasProvinceStatisticsRequest
//
// @return DescribeVerifyPersonasProvinceStatisticsResponse
func (client *Client) DescribeVerifyPersonasProvinceStatistics(request *DescribeVerifyPersonasProvinceStatisticsRequest) (_result *DescribeVerifyPersonasProvinceStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVerifyPersonasProvinceStatisticsResponse{}
	_body, _err := client.DescribeVerifyPersonasProvinceStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries ID Verification statistics by gender.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeVerifyPersonasSexStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyPersonasSexStatisticsResponse
func (client *Client) DescribeVerifyPersonasSexStatisticsWithOptions(request *DescribeVerifyPersonasSexStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyPersonasSexStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.TimeRange) {
		query["TimeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyPersonasSexStatistics"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyPersonasSexStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries ID Verification statistics by gender.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeVerifyPersonasSexStatisticsRequest
//
// @return DescribeVerifyPersonasSexStatisticsResponse
func (client *Client) DescribeVerifyPersonasSexStatistics(request *DescribeVerifyPersonasSexStatisticsRequest) (_result *DescribeVerifyPersonasSexStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVerifyPersonasSexStatisticsResponse{}
	_body, _err := client.DescribeVerifyPersonasSexStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the result of an ID Verification task.
//
// Description:
//
// Before you begin: Before calling this API, make sure that you have completed the required preparations. For more information, see [ID Verification server-side integration preparations](https://help.aliyun.com/document_detail/127471.html) and [Face liveness verification server-side integration preparations](https://help.aliyun.com/document_detail/127717.html).
//
// > Alibaba Cloud ID Verification retains verification data for only the last 180 days. To use verification data for subsequent business purposes, call this operation promptly to retrieve and store the data to avoid data loss.
//
// Request method: HTTPS POST and GET.
//
// Operation description: After the caller\\"s mobile client receives a callback, the server can call this operation to obtain the corresponding verification status and verification materials.
//
// Applicable scope: This operation is applicable to the SDK + server-side integration verification solution.
//
// @param request - DescribeVerifyResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyResultResponse
func (client *Client) DescribeVerifyResultWithOptions(request *DescribeVerifyResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyResult"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of an ID Verification task.
//
// Description:
//
// Before you begin: Before calling this API, make sure that you have completed the required preparations. For more information, see [ID Verification server-side integration preparations](https://help.aliyun.com/document_detail/127471.html) and [Face liveness verification server-side integration preparations](https://help.aliyun.com/document_detail/127717.html).
//
// > Alibaba Cloud ID Verification retains verification data for only the last 180 days. To use verification data for subsequent business purposes, call this operation promptly to retrieve and store the data to avoid data loss.
//
// Request method: HTTPS POST and GET.
//
// Operation description: After the caller\\"s mobile client receives a callback, the server can call this operation to obtain the corresponding verification status and verification materials.
//
// Applicable scope: This operation is applicable to the SDK + server-side integration verification solution.
//
// @param request - DescribeVerifyResultRequest
//
// @return DescribeVerifyResultResponse
func (client *Client) DescribeVerifyResult(request *DescribeVerifyResultRequest) (_result *DescribeVerifyResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVerifyResultResponse{}
	_body, _err := client.DescribeVerifyResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the download URL of an offline SDK.
//
// Description:
//
// Request method: Supports HTTPS POST and GET methods.
//
// Operation description: Retrieves the result of an offline facial recognition SDK generation task based on the task ID.
//
// @param request - DescribeVerifySDKRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifySDKResponse
func (client *Client) DescribeVerifySDKWithOptions(request *DescribeVerifySDKRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifySDKResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifySDK"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifySDKResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the download URL of an offline SDK.
//
// Description:
//
// Request method: Supports HTTPS POST and GET methods.
//
// Operation description: Retrieves the result of an offline facial recognition SDK generation task based on the task ID.
//
// @param request - DescribeVerifySDKRequest
//
// @return DescribeVerifySDKResponse
func (client *Client) DescribeVerifySDK(request *DescribeVerifySDKRequest) (_result *DescribeVerifySDKResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVerifySDKResponse{}
	_body, _err := client.DescribeVerifySDKWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query authentication details by page with conditions.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request methods: HTTPS POST and GET.
//
// @param request - DescribeVerifySearchPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifySearchPageListResponse
func (client *Client) DescribeVerifySearchPageListWithOptions(request *DescribeVerifySearchPageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifySearchPageListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertNo) {
		query["CertNo"] = request.CertNo
	}

	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.HasDeviceRisk) {
		query["HasDeviceRisk"] = request.HasDeviceRisk
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.OuterOrderNo) {
		query["OuterOrderNo"] = request.OuterOrderNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Passed) {
		query["Passed"] = request.Passed
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.RiskBizScenario) {
		query["RiskBizScenario"] = request.RiskBizScenario
	}

	if !dara.IsNil(request.RiskDevice) {
		query["RiskDevice"] = request.RiskDevice
	}

	if !dara.IsNil(request.RiskDeviceToken) {
		query["RiskDeviceToken"] = request.RiskDeviceToken
	}

	if !dara.IsNil(request.RiskGeneric) {
		query["RiskGeneric"] = request.RiskGeneric
	}

	if !dara.IsNil(request.RiskModelMining) {
		query["RiskModelMining"] = request.RiskModelMining
	}

	if !dara.IsNil(request.Root) {
		query["Root"] = request.Root
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.Simulator) {
		query["Simulator"] = request.Simulator
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SubCode) {
		query["SubCode"] = request.SubCode
	}

	if !dara.IsNil(request.SubCodes) {
		query["SubCodes"] = request.SubCodes
	}

	if !dara.IsNil(request.VirtualVideo) {
		query["VirtualVideo"] = request.VirtualVideo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifySearchPageList"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifySearchPageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query authentication details by page with conditions.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request methods: HTTPS POST and GET.
//
// @param request - DescribeVerifySearchPageListRequest
//
// @return DescribeVerifySearchPageListResponse
func (client *Client) DescribeVerifySearchPageList(request *DescribeVerifySearchPageListRequest) (_result *DescribeVerifySearchPageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVerifySearchPageListResponse{}
	_body, _err := client.DescribeVerifySearchPageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics information of authentication requests.
//
// Description:
//
// - Request method: HTTPS POST and GET methods are supported.
//
// - Service address: cloudauth.aliyuncs.com.
//
// @param request - DescribeVerifyStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyStatisticsResponse
func (client *Client) DescribeVerifyStatisticsWithOptions(request *DescribeVerifyStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgeGt) {
		query["AgeGt"] = request.AgeGt
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyStatistics"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics information of authentication requests.
//
// Description:
//
// - Request method: HTTPS POST and GET methods are supported.
//
// - Service address: cloudauth.aliyuncs.com.
//
// @param request - DescribeVerifyStatisticsRequest
//
// @return DescribeVerifyStatisticsResponse
func (client *Client) DescribeVerifyStatistics(request *DescribeVerifyStatisticsRequest) (_result *DescribeVerifyStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVerifyStatisticsResponse{}
	_body, _err := client.DescribeVerifyStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates an authentication request and obtains an authentication token. This operation is applicable to authentication schemes that use SDK and server-side integration.
//
// Description:
//
// Preparations: Before calling this operation, make sure that you have completed the required preparations. For more information, see [Overview of the ID Verification scheme integration process](https://help.aliyun.com/document_detail/127536.html) and [Overview of the face liveness verification scheme (liveness detection scheme) integration process](https://help.aliyun.com/document_detail/127687.html).
//
// Request method: HTTPS POST and GET.
//
// Operation description: Before each authentication, call this operation to obtain an authentication token (VerifyToken), which is used to connect the various operations in the authentication request.
//
// Scope of application: This operation is applicable to mobile SDK integration.
//
// Image URL: Use a publicly accessible HTTP or HTTPS URL. Example: `http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg`.
//
// Image limits:
//
// - The relative path or absolute path of local images is not supported.
//
// - Keep the size of a single image within 2 MB to avoid algorithm fetch timeout.
//
// - The face area in the image must be at least 64 × 64 pixels (px).
//
// - The body of a single request has a size limit of 8 MB. Make sure that the total size of all images and other information in the request does not exceed this limit.
//
// @param request - DescribeVerifyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyTokenResponse
func (client *Client) DescribeVerifyTokenWithOptions(request *DescribeVerifyTokenRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.CallbackSeed) {
		query["CallbackSeed"] = request.CallbackSeed
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.FaceRetainedImageUrl) {
		query["FaceRetainedImageUrl"] = request.FaceRetainedImageUrl
	}

	if !dara.IsNil(request.FailedRedirectUrl) {
		query["FailedRedirectUrl"] = request.FailedRedirectUrl
	}

	if !dara.IsNil(request.IdCardBackImageUrl) {
		query["IdCardBackImageUrl"] = request.IdCardBackImageUrl
	}

	if !dara.IsNil(request.IdCardFrontImageUrl) {
		query["IdCardFrontImageUrl"] = request.IdCardFrontImageUrl
	}

	if !dara.IsNil(request.IdCardNumber) {
		query["IdCardNumber"] = request.IdCardNumber
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PassedRedirectUrl) {
		query["PassedRedirectUrl"] = request.PassedRedirectUrl
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserIp) {
		query["UserIp"] = request.UserIp
	}

	if !dara.IsNil(request.UserPhoneNumber) {
		query["UserPhoneNumber"] = request.UserPhoneNumber
	}

	if !dara.IsNil(request.UserRegistTime) {
		query["UserRegistTime"] = request.UserRegistTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyToken"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates an authentication request and obtains an authentication token. This operation is applicable to authentication schemes that use SDK and server-side integration.
//
// Description:
//
// Preparations: Before calling this operation, make sure that you have completed the required preparations. For more information, see [Overview of the ID Verification scheme integration process](https://help.aliyun.com/document_detail/127536.html) and [Overview of the face liveness verification scheme (liveness detection scheme) integration process](https://help.aliyun.com/document_detail/127687.html).
//
// Request method: HTTPS POST and GET.
//
// Operation description: Before each authentication, call this operation to obtain an authentication token (VerifyToken), which is used to connect the various operations in the authentication request.
//
// Scope of application: This operation is applicable to mobile SDK integration.
//
// Image URL: Use a publicly accessible HTTP or HTTPS URL. Example: `http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg`.
//
// Image limits:
//
// - The relative path or absolute path of local images is not supported.
//
// - Keep the size of a single image within 2 MB to avoid algorithm fetch timeout.
//
// - The face area in the image must be at least 64 × 64 pixels (px).
//
// - The body of a single request has a size limit of 8 MB. Make sure that the total size of all images and other information in the request does not exceed this limit.
//
// @param request - DescribeVerifyTokenRequest
//
// @return DescribeVerifyTokenResponse
func (client *Client) DescribeVerifyToken(request *DescribeVerifyTokenRequest) (_result *DescribeVerifyTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVerifyTokenResponse{}
	_body, _err := client.DescribeVerifyTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the whitelist of a scenario.
//
// Description:
//
// Request method: Only HTTPS POST requests are supported.
//
// @param request - DescribeWhitelistSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWhitelistSettingResponse
func (client *Client) DescribeWhitelistSettingWithOptions(request *DescribeWhitelistSettingRequest, runtime *dara.RuntimeOptions) (_result *DescribeWhitelistSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertNo) {
		query["CertNo"] = request.CertNo
	}

	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.ValidEndDate) {
		query["ValidEndDate"] = request.ValidEndDate
	}

	if !dara.IsNil(request.ValidStartDate) {
		query["ValidStartDate"] = request.ValidStartDate
	}

	if !dara.IsNil(request.WhitelistType) {
		query["WhitelistType"] = request.WhitelistType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWhitelistSetting"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWhitelistSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the whitelist of a scenario.
//
// Description:
//
// Request method: Only HTTPS POST requests are supported.
//
// @param request - DescribeWhitelistSettingRequest
//
// @return DescribeWhitelistSettingResponse
func (client *Client) DescribeWhitelistSetting(request *DescribeWhitelistSettingRequest) (_result *DescribeWhitelistSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWhitelistSettingResponse{}
	_body, _err := client.DescribeWhitelistSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detects validity attributes of faces in a photo.
//
// Description:
//
// Request method: Only HTTPS POST requests are supported.
//
// Operation description: Detects validity-related attributes of faces in an input photo, helping you determine whether the photo meets your business requirements for retention or comparison. Currently supported face validity attributes include: whether a face is present, whether the face is blurry, whether glasses are worn, facial pose, and whether the face is smiling.
//
// Notes on uploading image addresses: When submitting an image, provide its HTTP URL, OSS address, or Base64 encoding.
//
// - HTTP address: A publicly accessible HTTP URL. For example, `http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg`.
//
// - Base64 encoding: A Base64-encoded image in the format `base64://<Base64-encoded image string>`.
//
// Image limits:
//
// - Relative paths or absolute paths of local images are not supported.
//
// - Keep the size of a single image within 2 MB to avoid algorithm fetch timeouts.
//
// - The request body has a size limit of 8 MB. Make sure the total size of all images and other information in the request does not exceed this limit.
//
// - When using Base64 to transmit images, set the request method to POST. Remove the header description from the Base64 character string, such as `data:image/png,base64`.
//
// @param request - DetectFaceAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectFaceAttributesResponse
func (client *Client) DetectFaceAttributesWithOptions(request *DetectFaceAttributesRequest, runtime *dara.RuntimeOptions) (_result *DetectFaceAttributesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.MaterialValue) {
		body["MaterialValue"] = request.MaterialValue
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectFaceAttributes"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectFaceAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects validity attributes of faces in a photo.
//
// Description:
//
// Request method: Only HTTPS POST requests are supported.
//
// Operation description: Detects validity-related attributes of faces in an input photo, helping you determine whether the photo meets your business requirements for retention or comparison. Currently supported face validity attributes include: whether a face is present, whether the face is blurry, whether glasses are worn, facial pose, and whether the face is smiling.
//
// Notes on uploading image addresses: When submitting an image, provide its HTTP URL, OSS address, or Base64 encoding.
//
// - HTTP address: A publicly accessible HTTP URL. For example, `http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg`.
//
// - Base64 encoding: A Base64-encoded image in the format `base64://<Base64-encoded image string>`.
//
// Image limits:
//
// - Relative paths or absolute paths of local images are not supported.
//
// - Keep the size of a single image within 2 MB to avoid algorithm fetch timeouts.
//
// - The request body has a size limit of 8 MB. Make sure the total size of all images and other information in the request does not exceed this limit.
//
// - When using Base64 to transmit images, set the request method to POST. Remove the header description from the Base64 character string, such as `data:image/png,base64`.
//
// @param request - DetectFaceAttributesRequest
//
// @return DetectFaceAttributesResponse
func (client *Client) DetectFaceAttributes(request *DetectFaceAttributesRequest) (_result *DetectFaceAttributesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectFaceAttributesResponse{}
	_body, _err := client.DetectFaceAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves statistical call volume data.
//
// Description:
//
// Retrieves the download link for the statistical call data file under a product plan based on the specified query conditions.
//
// - Method: HTTPS POST
//
// - Endpoint: cloudauth.aliyuncs.com
//
// > ID Verification counts call volume by CertifyId. To facilitate reconciliation, retain the CertifyId field in your system.
//
// @param request - DownloadVerifyRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadVerifyRecordsResponse
func (client *Client) DownloadVerifyRecordsWithOptions(request *DownloadVerifyRecordsRequest, runtime *dara.RuntimeOptions) (_result *DownloadVerifyRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizParam) {
		query["BizParam"] = request.BizParam
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadVerifyRecords"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadVerifyRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves statistical call volume data.
//
// Description:
//
// Retrieves the download link for the statistical call data file under a product plan based on the specified query conditions.
//
// - Method: HTTPS POST
//
// - Endpoint: cloudauth.aliyuncs.com
//
// > ID Verification counts call volume by CertifyId. To facilitate reconciliation, retain the CertifyId field in your system.
//
// @param request - DownloadVerifyRecordsRequest
//
// @return DownloadVerifyRecordsResponse
func (client *Client) DownloadVerifyRecords(request *DownloadVerifyRecordsRequest) (_result *DownloadVerifyRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DownloadVerifyRecordsResponse{}
	_body, _err := client.DownloadVerifyRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the validity period of a two-factor identity document.
//
// @param request - Id2MetaPeriodVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id2MetaPeriodVerifyResponse
func (client *Client) Id2MetaPeriodVerifyWithOptions(request *Id2MetaPeriodVerifyRequest, runtime *dara.RuntimeOptions) (_result *Id2MetaPeriodVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	if !dara.IsNil(request.ValidityEndDate) {
		body["ValidityEndDate"] = request.ValidityEndDate
	}

	if !dara.IsNil(request.ValidityStartDate) {
		body["ValidityStartDate"] = request.ValidityStartDate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Id2MetaPeriodVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Id2MetaPeriodVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the validity period of a two-factor identity document.
//
// @param request - Id2MetaPeriodVerifyRequest
//
// @return Id2MetaPeriodVerifyResponse
func (client *Client) Id2MetaPeriodVerify(request *Id2MetaPeriodVerifyRequest) (_result *Id2MetaPeriodVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &Id2MetaPeriodVerifyResponse{}
	_body, _err := client.Id2MetaPeriodVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Two-factor identity verification Standard Edition.
//
// @param request - Id2MetaStandardVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id2MetaStandardVerifyResponse
func (client *Client) Id2MetaStandardVerifyWithOptions(request *Id2MetaStandardVerifyRequest, runtime *dara.RuntimeOptions) (_result *Id2MetaStandardVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Id2MetaStandardVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Id2MetaStandardVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Two-factor identity verification Standard Edition.
//
// @param request - Id2MetaStandardVerifyRequest
//
// @return Id2MetaStandardVerifyResponse
func (client *Client) Id2MetaStandardVerify(request *Id2MetaStandardVerifyRequest) (_result *Id2MetaStandardVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &Id2MetaStandardVerifyResponse{}
	_body, _err := client.Id2MetaStandardVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the authenticity and consistency of a name and ID card number against an authoritative data source.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com (IPv4) or cloudauth-dualstack.aliyuncs.com (IPv6).
//
// - Request method: POST and GET.
//
// - Transfer protocol: HTTPS.
//
// @param request - Id2MetaVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id2MetaVerifyResponse
func (client *Client) Id2MetaVerifyWithOptions(request *Id2MetaVerifyRequest, runtime *dara.RuntimeOptions) (_result *Id2MetaVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Id2MetaVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Id2MetaVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the authenticity and consistency of a name and ID card number against an authoritative data source.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com (IPv4) or cloudauth-dualstack.aliyuncs.com (IPv6).
//
// - Request method: POST and GET.
//
// - Transfer protocol: HTTPS.
//
// @param request - Id2MetaVerifyRequest
//
// @return Id2MetaVerifyResponse
func (client *Client) Id2MetaVerify(request *Id2MetaVerifyRequest) (_result *Id2MetaVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &Id2MetaVerifyResponse{}
	_body, _err := client.Id2MetaVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the authenticity and consistency of an identity by taking images of both sides of an ID card, extracting the name and ID number via OCR, and checking them against an authoritative source.
//
// Description:
//
// Takes images of both sides of an ID card and returns the verification result of the two factors from an authoritative data source.
//
// @param request - Id2MetaVerifyWithOCRRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id2MetaVerifyWithOCRResponse
func (client *Client) Id2MetaVerifyWithOCRWithOptions(request *Id2MetaVerifyWithOCRRequest, runtime *dara.RuntimeOptions) (_result *Id2MetaVerifyWithOCRResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CertFile) {
		body["CertFile"] = request.CertFile
	}

	if !dara.IsNil(request.CertNationalFile) {
		body["CertNationalFile"] = request.CertNationalFile
	}

	if !dara.IsNil(request.CertNationalUrl) {
		body["CertNationalUrl"] = request.CertNationalUrl
	}

	if !dara.IsNil(request.CertUrl) {
		body["CertUrl"] = request.CertUrl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Id2MetaVerifyWithOCR"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Id2MetaVerifyWithOCRResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the authenticity and consistency of an identity by taking images of both sides of an ID card, extracting the name and ID number via OCR, and checking them against an authoritative source.
//
// Description:
//
// Takes images of both sides of an ID card and returns the verification result of the two factors from an authoritative data source.
//
// @param request - Id2MetaVerifyWithOCRRequest
//
// @return Id2MetaVerifyWithOCRResponse
func (client *Client) Id2MetaVerifyWithOCR(request *Id2MetaVerifyWithOCRRequest) (_result *Id2MetaVerifyWithOCRResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &Id2MetaVerifyWithOCRResponse{}
	_body, _err := client.Id2MetaVerifyWithOCRWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) Id2MetaVerifyWithOCRAdvance(request *Id2MetaVerifyWithOCRAdvanceRequest, runtime *dara.RuntimeOptions) (_result *Id2MetaVerifyWithOCRResponse, _err error) {
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
		"Product":  dara.String("Cloudauth"),
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
	id2MetaVerifyWithOCRReq := &Id2MetaVerifyWithOCRRequest{}
	openapiutil.Convert(request, id2MetaVerifyWithOCRReq)
	if !dara.IsNil(request.CertFileObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.CertFileObject,
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
		id2MetaVerifyWithOCRReq.CertFile = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	if !dara.IsNil(request.CertNationalFileObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.CertNationalFileObject,
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
		id2MetaVerifyWithOCRReq.CertNationalFile = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	id2MetaVerifyWithOCRResp, _err := client.Id2MetaVerifyWithOCRWithOptions(id2MetaVerifyWithOCRReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = id2MetaVerifyWithOCRResp
	return _result, _err
}

// Summary:
//
// Verifies the authenticity and consistency of a name, ID card number, and facial photo against an authoritative data source.
//
// Description:
//
// Verifies the authenticity and consistency of a name, ID card number, and facial photo against an authoritative data source.
//
// @param request - Id3MetaVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id3MetaVerifyResponse
func (client *Client) Id3MetaVerifyWithOptions(request *Id3MetaVerifyRequest, runtime *dara.RuntimeOptions) (_result *Id3MetaVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Crop) {
		body["Crop"] = request.Crop
	}

	if !dara.IsNil(request.FaceFile) {
		body["FaceFile"] = request.FaceFile
	}

	if !dara.IsNil(request.FacePicture) {
		body["FacePicture"] = request.FacePicture
	}

	if !dara.IsNil(request.FaceUrl) {
		body["FaceUrl"] = request.FaceUrl
	}

	if !dara.IsNil(request.IdentifyNum) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Id3MetaVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Id3MetaVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the authenticity and consistency of a name, ID card number, and facial photo against an authoritative data source.
//
// Description:
//
// Verifies the authenticity and consistency of a name, ID card number, and facial photo against an authoritative data source.
//
// @param request - Id3MetaVerifyRequest
//
// @return Id3MetaVerifyResponse
func (client *Client) Id3MetaVerify(request *Id3MetaVerifyRequest) (_result *Id3MetaVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &Id3MetaVerifyResponse{}
	_body, _err := client.Id3MetaVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) Id3MetaVerifyAdvance(request *Id3MetaVerifyAdvanceRequest, runtime *dara.RuntimeOptions) (_result *Id3MetaVerifyResponse, _err error) {
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
		"Product":  dara.String("Cloudauth"),
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
	id3MetaVerifyReq := &Id3MetaVerifyRequest{}
	openapiutil.Convert(request, id3MetaVerifyReq)
	if !dara.IsNil(request.FaceFileObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FaceFileObject,
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
		id3MetaVerifyReq.FaceFile = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	id3MetaVerifyResp, _err := client.Id3MetaVerifyWithOptions(id3MetaVerifyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = id3MetaVerifyResp
	return _result, _err
}

// Summary:
//
// Accepts images of the front and back of an ID card, extracts the name, ID number, and facial photo by using OCR, and verifies the authenticity and consistency of the three facial elements against an authoritative source.
//
// Description:
//
// Submits images of the front and back of an ID card and returns the verification result of the three facial elements from an authoritative data source.
//
// @param request - Id3MetaVerifyWithOCRRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id3MetaVerifyWithOCRResponse
func (client *Client) Id3MetaVerifyWithOCRWithOptions(request *Id3MetaVerifyWithOCRRequest, runtime *dara.RuntimeOptions) (_result *Id3MetaVerifyWithOCRResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CertFile) {
		body["CertFile"] = request.CertFile
	}

	if !dara.IsNil(request.CertNationalFile) {
		body["CertNationalFile"] = request.CertNationalFile
	}

	if !dara.IsNil(request.CertNationalUrl) {
		body["CertNationalUrl"] = request.CertNationalUrl
	}

	if !dara.IsNil(request.CertUrl) {
		body["CertUrl"] = request.CertUrl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Id3MetaVerifyWithOCR"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Id3MetaVerifyWithOCRResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Accepts images of the front and back of an ID card, extracts the name, ID number, and facial photo by using OCR, and verifies the authenticity and consistency of the three facial elements against an authoritative source.
//
// Description:
//
// Submits images of the front and back of an ID card and returns the verification result of the three facial elements from an authoritative data source.
//
// @param request - Id3MetaVerifyWithOCRRequest
//
// @return Id3MetaVerifyWithOCRResponse
func (client *Client) Id3MetaVerifyWithOCR(request *Id3MetaVerifyWithOCRRequest) (_result *Id3MetaVerifyWithOCRResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &Id3MetaVerifyWithOCRResponse{}
	_body, _err := client.Id3MetaVerifyWithOCRWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) Id3MetaVerifyWithOCRAdvance(request *Id3MetaVerifyWithOCRAdvanceRequest, runtime *dara.RuntimeOptions) (_result *Id3MetaVerifyWithOCRResponse, _err error) {
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
		"Product":  dara.String("Cloudauth"),
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
	id3MetaVerifyWithOCRReq := &Id3MetaVerifyWithOCRRequest{}
	openapiutil.Convert(request, id3MetaVerifyWithOCRReq)
	if !dara.IsNil(request.CertFileObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.CertFileObject,
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
		id3MetaVerifyWithOCRReq.CertFile = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	if !dara.IsNil(request.CertNationalFileObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.CertNationalFileObject,
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
		id3MetaVerifyWithOCRReq.CertNationalFile = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	id3MetaVerifyWithOCRResp, _err := client.Id3MetaVerifyWithOCRWithOptions(id3MetaVerifyWithOCRReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = id3MetaVerifyWithOCRResp
	return _result, _err
}

// Summary:
//
// Initiates an OCR request.
//
// @param request - InitAuthVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitAuthVerifyResponse
func (client *Client) InitAuthVerifyWithOptions(request *InitAuthVerifyRequest, runtime *dara.RuntimeOptions) (_result *InitAuthVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CallbackToken) {
		body["CallbackToken"] = request.CallbackToken
	}

	if !dara.IsNil(request.CallbackUrl) {
		body["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.CardPageNumber) {
		body["CardPageNumber"] = request.CardPageNumber
	}

	if !dara.IsNil(request.CardType) {
		body["CardType"] = request.CardType
	}

	if !dara.IsNil(request.DocScanMode) {
		body["DocScanMode"] = request.DocScanMode
	}

	if !dara.IsNil(request.IdSpoof) {
		body["IdSpoof"] = request.IdSpoof
	}

	if !dara.IsNil(request.MetaInfo) {
		body["MetaInfo"] = request.MetaInfo
	}

	if !dara.IsNil(request.OuterOrderNo) {
		body["OuterOrderNo"] = request.OuterOrderNo
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitAuthVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitAuthVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates an OCR request.
//
// @param request - InitAuthVerifyRequest
//
// @return InitAuthVerifyResponse
func (client *Client) InitAuthVerify(request *InitAuthVerifyRequest) (_result *InitAuthVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitAuthVerifyResponse{}
	_body, _err := client.InitAuthVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates an image verification authentication request.
//
// Description:
//
// Retrieves a CertifyId before each authentication session. The CertifyId links all API operations within the authentication request.
//
// @param request - InitCardVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitCardVerifyResponse
func (client *Client) InitCardVerifyWithOptions(request *InitCardVerifyRequest, runtime *dara.RuntimeOptions) (_result *InitCardVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallbackToken) {
		query["CallbackToken"] = request.CallbackToken
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.CardPageNumber) {
		query["CardPageNumber"] = request.CardPageNumber
	}

	if !dara.IsNil(request.CardType) {
		query["CardType"] = request.CardType
	}

	if !dara.IsNil(request.DocScanMode) {
		query["DocScanMode"] = request.DocScanMode
	}

	if !dara.IsNil(request.MerchantBizId) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.MetaInfo) {
		query["MetaInfo"] = request.MetaInfo
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.PictureSave) {
		query["PictureSave"] = request.PictureSave
	}

	if !dara.IsNil(request.VerifyMeta) {
		query["VerifyMeta"] = request.VerifyMeta
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitCardVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitCardVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates an image verification authentication request.
//
// Description:
//
// Retrieves a CertifyId before each authentication session. The CertifyId links all API operations within the authentication request.
//
// @param request - InitCardVerifyRequest
//
// @return InitCardVerifyResponse
func (client *Client) InitCardVerify(request *InitCardVerifyRequest) (_result *InitCardVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitCardVerifyResponse{}
	_body, _err := client.InitCardVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a CertifyId before each authentication to link the interfaces in the authentication request.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com
//
// - Request method: HTTPS POST and GET.
//
// - This operation uses different parameters for different product plans. For more information, refer to the [official documentation](https://www.alibabacloud.com/help/en/id-verification/financial-grade-id-verification/product-overview/introduction/).
//
// #### Image format requirements
//
// When performing ID Verification, submit images that meet all of the following conditions:
//
// - A recent photo with a complete, clear, and unobstructed face, a natural expression, and the subject facing the camera directly.
//
// - A clear photo with normal exposure. The face must not be too dark, too bright, or have glare, and the angle must not deviate significantly.
//
// - Resolution must not exceed 1920×1080 and must be at least 640×480. Scale the short side to 720 pixels and use a compression ratio greater than 0.9.
//
// - Photo size: < 1 MB.
//
// - Photos rotated 90, 180, and 270 degrees are supported. For photos with multiple faces, the largest face is selected.
//
// @param request - InitFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitFaceVerifyResponse
func (client *Client) InitFaceVerifyWithOptions(request *InitFaceVerifyRequest, runtime *dara.RuntimeOptions) (_result *InitFaceVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppQualityCheck) {
		query["AppQualityCheck"] = request.AppQualityCheck
	}

	if !dara.IsNil(request.Birthday) {
		query["Birthday"] = request.Birthday
	}

	if !dara.IsNil(request.CallbackToken) {
		query["CallbackToken"] = request.CallbackToken
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.CameraSelection) {
		query["CameraSelection"] = request.CameraSelection
	}

	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

	if !dara.IsNil(request.CertNo) {
		query["CertNo"] = request.CertNo
	}

	if !dara.IsNil(request.CertType) {
		query["CertType"] = request.CertType
	}

	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.CertifyUrlStyle) {
		query["CertifyUrlStyle"] = request.CertifyUrlStyle
	}

	if !dara.IsNil(request.CertifyUrlType) {
		query["CertifyUrlType"] = request.CertifyUrlType
	}

	if !dara.IsNil(request.EnableBeauty) {
		query["EnableBeauty"] = request.EnableBeauty
	}

	if !dara.IsNil(request.EncryptType) {
		query["EncryptType"] = request.EncryptType
	}

	if !dara.IsNil(request.FaceContrastPictureUrl) {
		query["FaceContrastPictureUrl"] = request.FaceContrastPictureUrl
	}

	if !dara.IsNil(request.FaceGuardOutput) {
		query["FaceGuardOutput"] = request.FaceGuardOutput
	}

	if !dara.IsNil(request.H5DegradeConfirmBtn) {
		query["H5DegradeConfirmBtn"] = request.H5DegradeConfirmBtn
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.MetaInfo) {
		query["MetaInfo"] = request.MetaInfo
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.NeedMultiFaceCheck) {
		query["NeedMultiFaceCheck"] = request.NeedMultiFaceCheck
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssObjectName) {
		query["OssObjectName"] = request.OssObjectName
	}

	if !dara.IsNil(request.OuterOrderNo) {
		query["OuterOrderNo"] = request.OuterOrderNo
	}

	if !dara.IsNil(request.ProcedurePriority) {
		query["ProcedurePriority"] = request.ProcedurePriority
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.RarelyCharacters) {
		query["RarelyCharacters"] = request.RarelyCharacters
	}

	if !dara.IsNil(request.ReadImg) {
		query["ReadImg"] = request.ReadImg
	}

	if !dara.IsNil(request.ReturnUrl) {
		query["ReturnUrl"] = request.ReturnUrl
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.SuitableType) {
		query["SuitableType"] = request.SuitableType
	}

	if !dara.IsNil(request.UiCustomUrl) {
		query["UiCustomUrl"] = request.UiCustomUrl
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.ValidityDate) {
		query["ValidityDate"] = request.ValidityDate
	}

	if !dara.IsNil(request.VideoEvidence) {
		query["VideoEvidence"] = request.VideoEvidence
	}

	if !dara.IsNil(request.VoluntaryCustomizedContent) {
		query["VoluntaryCustomizedContent"] = request.VoluntaryCustomizedContent
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthId) {
		body["AuthId"] = request.AuthId
	}

	if !dara.IsNil(request.Crop) {
		body["Crop"] = request.Crop
	}

	if !dara.IsNil(request.FaceContrastPicture) {
		body["FaceContrastPicture"] = request.FaceContrastPicture
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitFaceVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitFaceVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a CertifyId before each authentication to link the interfaces in the authentication request.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com
//
// - Request method: HTTPS POST and GET.
//
// - This operation uses different parameters for different product plans. For more information, refer to the [official documentation](https://www.alibabacloud.com/help/en/id-verification/financial-grade-id-verification/product-overview/introduction/).
//
// #### Image format requirements
//
// When performing ID Verification, submit images that meet all of the following conditions:
//
// - A recent photo with a complete, clear, and unobstructed face, a natural expression, and the subject facing the camera directly.
//
// - A clear photo with normal exposure. The face must not be too dark, too bright, or have glare, and the angle must not deviate significantly.
//
// - Resolution must not exceed 1920×1080 and must be at least 640×480. Scale the short side to 720 pixels and use a compression ratio greater than 0.9.
//
// - Photo size: < 1 MB.
//
// - Photos rotated 90, 180, and 270 degrees are supported. For photos with multiple faces, the largest face is selected.
//
// @param request - InitFaceVerifyRequest
//
// @return InitFaceVerifyResponse
func (client *Client) InitFaceVerify(request *InitFaceVerifyRequest) (_result *InitFaceVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitFaceVerifyResponse{}
	_body, _err := client.InitFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an ID Verification whitelist entry.
//
// @param request - InsertWhiteListSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertWhiteListSettingResponse
func (client *Client) InsertWhiteListSettingWithOptions(request *InsertWhiteListSettingRequest, runtime *dara.RuntimeOptions) (_result *InsertWhiteListSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertNo) {
		query["CertNo"] = request.CertNo
	}

	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.ValidDay) {
		query["ValidDay"] = request.ValidDay
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertWhiteListSetting"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertWhiteListSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an ID Verification whitelist entry.
//
// @param request - InsertWhiteListSettingRequest
//
// @return InsertWhiteListSettingResponse
func (client *Client) InsertWhiteListSetting(request *InsertWhiteListSettingRequest) (_result *InsertWhiteListSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InsertWhiteListSettingResponse{}
	_body, _err := client.InsertWhiteListSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// LivenessFaceVerify is a service that performs passive liveness detection on face images submitted through an API operation. The algorithm primarily identifies liveness attack types such as screen replay and printed photo attacks. This service is suitable for low-risk business scenarios or for use in conjunction with an offline facial recognition SDK. If your business requires a higher level of security for real face verification, integrate the App or WebSDK pattern, or integrate the face Deepfake detection service to identify face forgery risks across more dimensions.
//
// Description:
//
// Calls the LivenessFaceVerify operation to perform liveness detection on a face image.
//
// @param request - LivenessFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LivenessFaceVerifyResponse
func (client *Client) LivenessFaceVerifyWithOptions(request *LivenessFaceVerifyRequest, runtime *dara.RuntimeOptions) (_result *LivenessFaceVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CertifyId) {
		body["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.Crop) {
		body["Crop"] = request.Crop
	}

	if !dara.IsNil(request.DeviceToken) {
		body["DeviceToken"] = request.DeviceToken
	}

	if !dara.IsNil(request.FaceContrastPicture) {
		body["FaceContrastPicture"] = request.FaceContrastPicture
	}

	if !dara.IsNil(request.FaceContrastPictureUrl) {
		body["FaceContrastPictureUrl"] = request.FaceContrastPictureUrl
	}

	if !dara.IsNil(request.Ip) {
		body["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.OssBucketName) {
		body["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssObjectName) {
		body["OssObjectName"] = request.OssObjectName
	}

	if !dara.IsNil(request.OuterOrderNo) {
		body["OuterOrderNo"] = request.OuterOrderNo
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LivenessFaceVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LivenessFaceVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// LivenessFaceVerify is a service that performs passive liveness detection on face images submitted through an API operation. The algorithm primarily identifies liveness attack types such as screen replay and printed photo attacks. This service is suitable for low-risk business scenarios or for use in conjunction with an offline facial recognition SDK. If your business requires a higher level of security for real face verification, integrate the App or WebSDK pattern, or integrate the face Deepfake detection service to identify face forgery risks across more dimensions.
//
// Description:
//
// Calls the LivenessFaceVerify operation to perform liveness detection on a face image.
//
// @param request - LivenessFaceVerifyRequest
//
// @return LivenessFaceVerifyResponse
func (client *Client) LivenessFaceVerify(request *LivenessFaceVerifyRequest) (_result *LivenessFaceVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LivenessFaceVerifyResponse{}
	_body, _err := client.LivenessFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the consistency of a mobile phone number and the owner\\"s name.
//
// Description:
//
// Passes in a mobile phone number and a name, and verifies their authenticity and consistency through an authoritative data source.
//
// @param request - Mobile2MetaVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Mobile2MetaVerifyResponse
func (client *Client) Mobile2MetaVerifyWithOptions(request *Mobile2MetaVerifyRequest, runtime *dara.RuntimeOptions) (_result *Mobile2MetaVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Mobile2MetaVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Mobile2MetaVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the consistency of a mobile phone number and the owner\\"s name.
//
// Description:
//
// Passes in a mobile phone number and a name, and verifies their authenticity and consistency through an authoritative data source.
//
// @param request - Mobile2MetaVerifyRequest
//
// @return Mobile2MetaVerifyResponse
func (client *Client) Mobile2MetaVerify(request *Mobile2MetaVerifyRequest) (_result *Mobile2MetaVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &Mobile2MetaVerifyResponse{}
	_body, _err := client.Mobile2MetaVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the three-element identity of a phone number (detailed version - standard edition).
//
// Description:
//
// Passes in a phone number, name, and ID card number, and verifies their authenticity and consistency through an authoritative data source. If the information is inconsistent, the reason for the inconsistency is returned.
//
// @param request - Mobile3MetaDetailStandardVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Mobile3MetaDetailStandardVerifyResponse
func (client *Client) Mobile3MetaDetailStandardVerifyWithOptions(request *Mobile3MetaDetailStandardVerifyRequest, runtime *dara.RuntimeOptions) (_result *Mobile3MetaDetailStandardVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Mobile3MetaDetailStandardVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Mobile3MetaDetailStandardVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the three-element identity of a phone number (detailed version - standard edition).
//
// Description:
//
// Passes in a phone number, name, and ID card number, and verifies their authenticity and consistency through an authoritative data source. If the information is inconsistent, the reason for the inconsistency is returned.
//
// @param request - Mobile3MetaDetailStandardVerifyRequest
//
// @return Mobile3MetaDetailStandardVerifyResponse
func (client *Client) Mobile3MetaDetailStandardVerify(request *Mobile3MetaDetailStandardVerifyRequest) (_result *Mobile3MetaDetailStandardVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &Mobile3MetaDetailStandardVerifyResponse{}
	_body, _err := client.Mobile3MetaDetailStandardVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the authenticity and consistency of a phone number, name, and ID card number against authoritative data sources, and returns the reason for any inconsistency.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com (IPv4) or cloudauth-dualstack.aliyuncs.com (IPv6).
//
// - Request method: POST and GET.
//
// - Transfer protocol: HTTPS.
//
// @param request - Mobile3MetaDetailVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Mobile3MetaDetailVerifyResponse
func (client *Client) Mobile3MetaDetailVerifyWithOptions(request *Mobile3MetaDetailVerifyRequest, runtime *dara.RuntimeOptions) (_result *Mobile3MetaDetailVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Mobile3MetaDetailVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Mobile3MetaDetailVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the authenticity and consistency of a phone number, name, and ID card number against authoritative data sources, and returns the reason for any inconsistency.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com (IPv4) or cloudauth-dualstack.aliyuncs.com (IPv6).
//
// - Request method: POST and GET.
//
// - Transfer protocol: HTTPS.
//
// @param request - Mobile3MetaDetailVerifyRequest
//
// @return Mobile3MetaDetailVerifyResponse
func (client *Client) Mobile3MetaDetailVerify(request *Mobile3MetaDetailVerifyRequest) (_result *Mobile3MetaDetailVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &Mobile3MetaDetailVerifyResponse{}
	_body, _err := client.Mobile3MetaDetailVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the three-element identity of a phone number (Standard Edition).
//
// Description:
//
// Passes in a phone number, name, and ID card number, and verifies their authenticity and consistency through an authoritative data source.
//
// @param request - Mobile3MetaSimpleStandardVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Mobile3MetaSimpleStandardVerifyResponse
func (client *Client) Mobile3MetaSimpleStandardVerifyWithOptions(request *Mobile3MetaSimpleStandardVerifyRequest, runtime *dara.RuntimeOptions) (_result *Mobile3MetaSimpleStandardVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Mobile3MetaSimpleStandardVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Mobile3MetaSimpleStandardVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the three-element identity of a phone number (Standard Edition).
//
// Description:
//
// Passes in a phone number, name, and ID card number, and verifies their authenticity and consistency through an authoritative data source.
//
// @param request - Mobile3MetaSimpleStandardVerifyRequest
//
// @return Mobile3MetaSimpleStandardVerifyResponse
func (client *Client) Mobile3MetaSimpleStandardVerify(request *Mobile3MetaSimpleStandardVerifyRequest) (_result *Mobile3MetaSimpleStandardVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &Mobile3MetaSimpleStandardVerifyResponse{}
	_body, _err := client.Mobile3MetaSimpleStandardVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pass in the name and ID card number to verify their authenticity and consistency through authoritative data sources.
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com (IPv4) or cloudauth-dualstack.aliyuncs.com (IPv6).
//
// - Request method: POST and GET.
//
// - Transfer protocol: HTTPS.
//
// @param request - Mobile3MetaSimpleVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Mobile3MetaSimpleVerifyResponse
func (client *Client) Mobile3MetaSimpleVerifyWithOptions(request *Mobile3MetaSimpleVerifyRequest, runtime *dara.RuntimeOptions) (_result *Mobile3MetaSimpleVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Mobile3MetaSimpleVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Mobile3MetaSimpleVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pass in the name and ID card number to verify their authenticity and consistency through authoritative data sources.
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com (IPv4) or cloudauth-dualstack.aliyuncs.com (IPv6).
//
// - Request method: POST and GET.
//
// - Transfer protocol: HTTPS.
//
// @param request - Mobile3MetaSimpleVerifyRequest
//
// @return Mobile3MetaSimpleVerifyResponse
func (client *Client) Mobile3MetaSimpleVerify(request *Mobile3MetaSimpleVerifyRequest) (_result *Mobile3MetaSimpleVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &Mobile3MetaSimpleVerifyResponse{}
	_body, _err := client.Mobile3MetaSimpleVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detects phone numbers.
//
// @param request - MobileDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MobileDetectResponse
func (client *Client) MobileDetectWithOptions(request *MobileDetectRequest, runtime *dara.RuntimeOptions) (_result *MobileDetectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Mobiles) {
		body["Mobiles"] = request.Mobiles
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MobileDetect"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MobileDetectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects phone numbers.
//
// @param request - MobileDetectRequest
//
// @return MobileDetectResponse
func (client *Client) MobileDetect(request *MobileDetectRequest) (_result *MobileDetectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MobileDetectResponse{}
	_body, _err := client.MobileDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the network availability status of a phone number.
//
// @param request - MobileOnlineStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MobileOnlineStatusResponse
func (client *Client) MobileOnlineStatusWithOptions(request *MobileOnlineStatusRequest, runtime *dara.RuntimeOptions) (_result *MobileOnlineStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MobileOnlineStatus"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MobileOnlineStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the network availability status of a phone number.
//
// @param request - MobileOnlineStatusRequest
//
// @return MobileOnlineStatusResponse
func (client *Client) MobileOnlineStatus(request *MobileOnlineStatusRequest) (_result *MobileOnlineStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MobileOnlineStatusResponse{}
	_body, _err := client.MobileOnlineStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the length of time a phone number has been active on a network.
//
// @param request - MobileOnlineTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MobileOnlineTimeResponse
func (client *Client) MobileOnlineTimeWithOptions(request *MobileOnlineTimeRequest, runtime *dara.RuntimeOptions) (_result *MobileOnlineTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MobileOnlineTime"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MobileOnlineTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the length of time a phone number has been active on a network.
//
// @param request - MobileOnlineTimeRequest
//
// @return MobileOnlineTimeResponse
func (client *Client) MobileOnlineTime(request *MobileOnlineTimeRequest) (_result *MobileOnlineTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MobileOnlineTimeResponse{}
	_body, _err := client.MobileOnlineTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies whether a phone number has been reassigned to a new subscriber.
//
// Description:
//
// Passes in a phone number and its registration date, and verifies the authenticity and consistency of the information through an authoritative data source.
//
// @param request - MobileRecycledMetaVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MobileRecycledMetaVerifyResponse
func (client *Client) MobileRecycledMetaVerifyWithOptions(request *MobileRecycledMetaVerifyRequest, runtime *dara.RuntimeOptions) (_result *MobileRecycledMetaVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		query["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.RegisterDate) {
		query["RegisterDate"] = request.RegisterDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MobileRecycledMetaVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MobileRecycledMetaVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies whether a phone number has been reassigned to a new subscriber.
//
// Description:
//
// Passes in a phone number and its registration date, and verifies the authenticity and consistency of the information through an authoritative data source.
//
// @param request - MobileRecycledMetaVerifyRequest
//
// @return MobileRecycledMetaVerifyResponse
func (client *Client) MobileRecycledMetaVerify(request *MobileRecycledMetaVerifyRequest) (_result *MobileRecycledMetaVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MobileRecycledMetaVerifyResponse{}
	_body, _err := client.MobileRecycledMetaVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or modifies a blacklist rule.
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// - Operation description: Creates or modifies a blacklist rule.
//
// @param tmpReq - ModifyBlackListStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBlackListStrategyResponse
func (client *Client) ModifyBlackListStrategyWithOptions(tmpReq *ModifyBlackListStrategyRequest, runtime *dara.RuntimeOptions) (_result *ModifyBlackListStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyBlackListStrategyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BlackListStrategy) {
		request.BlackListStrategyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BlackListStrategy, dara.String("BlackListStrategy"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BlackListStrategyShrink) {
		query["BlackListStrategy"] = request.BlackListStrategyShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBlackListStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBlackListStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies a blacklist rule.
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// - Operation description: Creates or modifies a blacklist rule.
//
// @param request - ModifyBlackListStrategyRequest
//
// @return ModifyBlackListStrategyResponse
func (client *Client) ModifyBlackListStrategy(request *ModifyBlackListStrategyRequest) (_result *ModifyBlackListStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBlackListStrategyResponse{}
	_body, _err := client.ModifyBlackListStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adjusts stability alert rules.
//
// Description:
//
// - Request method: HTTPS POST.
//
// - Request URL: cloudauth.aliyuncs.com.
//
// @param tmpReq - ModifyControlStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyControlStrategyResponse
func (client *Client) ModifyControlStrategyWithOptions(tmpReq *ModifyControlStrategyRequest, runtime *dara.RuntimeOptions) (_result *ModifyControlStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyControlStrategyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ControlStrategyList) {
		request.ControlStrategyListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ControlStrategyList, dara.String("ControlStrategyList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ControlStrategyListShrink) {
		query["ControlStrategyList"] = request.ControlStrategyListShrink
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyControlStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyControlStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adjusts stability alert rules.
//
// Description:
//
// - Request method: HTTPS POST.
//
// - Request URL: cloudauth.aliyuncs.com.
//
// @param request - ModifyControlStrategyRequest
//
// @return ModifyControlStrategyResponse
func (client *Client) ModifyControlStrategy(request *ModifyControlStrategyRequest) (_result *ModifyControlStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyControlStrategyResponse{}
	_body, _err := client.ModifyControlStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adjusts the list of custom rate limiting policies.
//
// Description:
//
// - Request method: HTTPS POST and GET methods are supported.
//
// - Endpoint: cloudauth.aliyuncs.com.
//
// @param tmpReq - ModifyCustomizeFlowStrategyListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCustomizeFlowStrategyListResponse
func (client *Client) ModifyCustomizeFlowStrategyListWithOptions(tmpReq *ModifyCustomizeFlowStrategyListRequest, runtime *dara.RuntimeOptions) (_result *ModifyCustomizeFlowStrategyListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyCustomizeFlowStrategyListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StrategyObject) {
		request.StrategyObjectShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StrategyObject, dara.String("StrategyObject"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.StrategyObjectShrink) {
		query["StrategyObject"] = request.StrategyObjectShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCustomizeFlowStrategyList"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCustomizeFlowStrategyListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adjusts the list of custom rate limiting policies.
//
// Description:
//
// - Request method: HTTPS POST and GET methods are supported.
//
// - Endpoint: cloudauth.aliyuncs.com.
//
// @param request - ModifyCustomizeFlowStrategyListRequest
//
// @return ModifyCustomizeFlowStrategyListResponse
func (client *Client) ModifyCustomizeFlowStrategyList(request *ModifyCustomizeFlowStrategyListRequest) (_result *ModifyCustomizeFlowStrategyListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCustomizeFlowStrategyListResponse{}
	_body, _err := client.ModifyCustomizeFlowStrategyListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates device-related information, such as extending the device authorization validity period, updating custom business identifiers, and updating device IDs. This operation is applicable to scenarios such as device validity period renewal.
//
// Description:
//
// Request method: You can send requests by using the HTTPS POST and GET methods.
//
// @param request - ModifyDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDeviceInfoResponse
func (client *Client) ModifyDeviceInfoWithOptions(request *ModifyDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *ModifyDeviceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.ExpiredDay) {
		query["ExpiredDay"] = request.ExpiredDay
	}

	if !dara.IsNil(request.UserDeviceId) {
		query["UserDeviceId"] = request.UserDeviceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDeviceInfo"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates device-related information, such as extending the device authorization validity period, updating custom business identifiers, and updating device IDs. This operation is applicable to scenarios such as device validity period renewal.
//
// Description:
//
// Request method: You can send requests by using the HTTPS POST and GET methods.
//
// @param request - ModifyDeviceInfoRequest
//
// @return ModifyDeviceInfoResponse
func (client *Client) ModifyDeviceInfo(request *ModifyDeviceInfoRequest) (_result *ModifyDeviceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDeviceInfoResponse{}
	_body, _err := client.ModifyDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries ID Verification whitelist configurations by using paging.
//
// @param request - PageQueryWhiteListSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageQueryWhiteListSettingResponse
func (client *Client) PageQueryWhiteListSettingWithOptions(request *PageQueryWhiteListSettingRequest, runtime *dara.RuntimeOptions) (_result *PageQueryWhiteListSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertNo) {
		query["CertNo"] = request.CertNo
	}

	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.ValidEndDate) {
		query["ValidEndDate"] = request.ValidEndDate
	}

	if !dara.IsNil(request.ValidStartDate) {
		query["ValidStartDate"] = request.ValidStartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PageQueryWhiteListSetting"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PageQueryWhiteListSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries ID Verification whitelist configurations by using paging.
//
// @param request - PageQueryWhiteListSettingRequest
//
// @return PageQueryWhiteListSettingResponse
func (client *Client) PageQueryWhiteListSetting(request *PageQueryWhiteListSettingRequest) (_result *PageQueryWhiteListSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PageQueryWhiteListSettingResponse{}
	_body, _err := client.PageQueryWhiteListSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the blacklist configuration list.
//
// Description:
//
// - Request endpoint: cloudauth.aliyuncs.com
//
// - Request method: HTTPS POST and GET.
//
// > You can configure blacklists for IP addresses, ID card numbers, phone numbers, and bank card numbers. When a request matches a blacklist entry, the system rejects the request and returns a fixed error code.
//
// You can configure blacklists for IP addresses, ID card numbers, phone numbers, and bank card numbers. When a request matches a blacklist entry, the system rejects the request and returns a fixed error code.
//
// @param request - QueryBlackListStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBlackListStrategyResponse
func (client *Client) QueryBlackListStrategyWithOptions(request *QueryBlackListStrategyRequest, runtime *dara.RuntimeOptions) (_result *QueryBlackListStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryBlackListStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryBlackListStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the blacklist configuration list.
//
// Description:
//
// - Request endpoint: cloudauth.aliyuncs.com
//
// - Request method: HTTPS POST and GET.
//
// > You can configure blacklists for IP addresses, ID card numbers, phone numbers, and bank card numbers. When a request matches a blacklist entry, the system rejects the request and returns a fixed error code.
//
// You can configure blacklists for IP addresses, ID card numbers, phone numbers, and bank card numbers. When a request matches a blacklist entry, the system rejects the request and returns a fixed error code.
//
// @param request - QueryBlackListStrategyRequest
//
// @return QueryBlackListStrategyResponse
func (client *Client) QueryBlackListStrategy(request *QueryBlackListStrategyRequest) (_result *QueryBlackListStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryBlackListStrategyResponse{}
	_body, _err := client.QueryBlackListStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of stability rules. The system monitors the stability of the server-side InitFaceVerify operation and API integration operations. If a threshold is triggered, the system sends an alert.
//
// Description:
//
// - Request method: HTTPS POST and GET.
//
// - Request URL: cloudauth.aliyuncs.com.
//
// @param request - QueryControlStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryControlStrategyResponse
func (client *Client) QueryControlStrategyWithOptions(request *QueryControlStrategyRequest, runtime *dara.RuntimeOptions) (_result *QueryControlStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryControlStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryControlStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of stability rules. The system monitors the stability of the server-side InitFaceVerify operation and API integration operations. If a threshold is triggered, the system sends an alert.
//
// Description:
//
// - Request method: HTTPS POST and GET.
//
// - Request URL: cloudauth.aliyuncs.com.
//
// @param request - QueryControlStrategyRequest
//
// @return QueryControlStrategyResponse
func (client *Client) QueryControlStrategy(request *QueryControlStrategyRequest) (_result *QueryControlStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryControlStrategyResponse{}
	_body, _err := client.QueryControlStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries security rules.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com
//
// - Request method: HTTPS POST and GET.
//
// - Security rules: monitoring rules that ensure system security, such as API abuse and abnormal account theft. When a threshold is triggered, the system sends an alert.
//
// @param request - QueryCustomizeFlowStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomizeFlowStrategyResponse
func (client *Client) QueryCustomizeFlowStrategyWithOptions(request *QueryCustomizeFlowStrategyRequest, runtime *dara.RuntimeOptions) (_result *QueryCustomizeFlowStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCustomizeFlowStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCustomizeFlowStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries security rules.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com
//
// - Request method: HTTPS POST and GET.
//
// - Security rules: monitoring rules that ensure system security, such as API abuse and abnormal account theft. When a threshold is triggered, the system sends an alert.
//
// @param request - QueryCustomizeFlowStrategyRequest
//
// @return QueryCustomizeFlowStrategyResponse
func (client *Client) QueryCustomizeFlowStrategy(request *QueryCustomizeFlowStrategyRequest) (_result *QueryCustomizeFlowStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCustomizeFlowStrategyResponse{}
	_body, _err := client.QueryCustomizeFlowStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of intent verification settings.
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - QuerySceneConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySceneConfigsResponse
func (client *Client) QuerySceneConfigsWithOptions(request *QuerySceneConfigsRequest, runtime *dara.RuntimeOptions) (_result *QuerySceneConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySceneConfigs"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySceneConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of intent verification settings.
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - QuerySceneConfigsRequest
//
// @return QuerySceneConfigsResponse
func (client *Client) QuerySceneConfigs(request *QuerySceneConfigsRequest) (_result *QuerySceneConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySceneConfigsResponse{}
	_body, _err := client.QuerySceneConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries an ID Verification download task.
//
// Description:
//
// Retrieves the download link of a statistical call data file under a product plan based on query conditions.
//
// - Method: HTTPS POST
//
// - Endpoint: cloudauth.aliyuncs.com
//
// > ID Verification uses CertifyId to calculate the call volume. To facilitate reconciliation, retain the CertifyId field in your system.
//
// @param request - QueryVerifyDownloadTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVerifyDownloadTaskResponse
func (client *Client) QueryVerifyDownloadTaskWithOptions(request *QueryVerifyDownloadTaskRequest, runtime *dara.RuntimeOptions) (_result *QueryVerifyDownloadTaskResponse, _err error) {
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
		Action:      dara.String("QueryVerifyDownloadTask"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryVerifyDownloadTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an ID Verification download task.
//
// Description:
//
// Retrieves the download link of a statistical call data file under a product plan based on query conditions.
//
// - Method: HTTPS POST
//
// - Endpoint: cloudauth.aliyuncs.com
//
// > ID Verification uses CertifyId to calculate the call volume. To facilitate reconciliation, retain the CertifyId field in your system.
//
// @param request - QueryVerifyDownloadTaskRequest
//
// @return QueryVerifyDownloadTaskResponse
func (client *Client) QueryVerifyDownloadTask(request *QueryVerifyDownloadTaskRequest) (_result *QueryVerifyDownloadTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryVerifyDownloadTaskResponse{}
	_body, _err := client.QueryVerifyDownloadTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries data transfer plans.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com
//
// - Request method: HTTPS POST and GET.
//
// - This operation uses different parameters for different product plans. For more information, see [official documentation](https://www.alibabacloud.com/help/en/id-verification/financial-grade-id-verification/product-overview/introduction/).
//
// @param request - QueryVerifyFlowPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVerifyFlowPackageResponse
func (client *Client) QueryVerifyFlowPackageWithOptions(request *QueryVerifyFlowPackageRequest, runtime *dara.RuntimeOptions) (_result *QueryVerifyFlowPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryVerifyFlowPackage"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryVerifyFlowPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries data transfer plans.
//
// Description:
//
// - Service endpoint: cloudauth.aliyuncs.com
//
// - Request method: HTTPS POST and GET.
//
// - This operation uses different parameters for different product plans. For more information, see [official documentation](https://www.alibabacloud.com/help/en/id-verification/financial-grade-id-verification/product-overview/introduction/).
//
// @param request - QueryVerifyFlowPackageRequest
//
// @return QueryVerifyFlowPackageResponse
func (client *Client) QueryVerifyFlowPackage(request *QueryVerifyFlowPackageRequest) (_result *QueryVerifyFlowPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryVerifyFlowPackageResponse{}
	_body, _err := client.QueryVerifyFlowPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the call volume of a product code based on different product plans.
//
// Description:
//
// - Request endpoint: cloudauth.aliyuncs.com
//
// - Request method: HTTPS POST and GET.
//
// > ID Verification counts call volume by CertifyId. To facilitate reconciliation, retain the CertifyId field in your system.
//
// @param request - QueryVerifyInvokeSatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVerifyInvokeSatisticResponse
func (client *Client) QueryVerifyInvokeSatisticWithOptions(request *QueryVerifyInvokeSatisticRequest, runtime *dara.RuntimeOptions) (_result *QueryVerifyInvokeSatisticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductProgramList) {
		query["ProductProgramList"] = request.ProductProgramList
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.SceneIdList) {
		query["SceneIdList"] = request.SceneIdList
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.StatisticsType) {
		query["StatisticsType"] = request.StatisticsType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryVerifyInvokeSatistic"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryVerifyInvokeSatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the call volume of a product code based on different product plans.
//
// Description:
//
// - Request endpoint: cloudauth.aliyuncs.com
//
// - Request method: HTTPS POST and GET.
//
// > ID Verification counts call volume by CertifyId. To facilitate reconciliation, retain the CertifyId field in your system.
//
// @param request - QueryVerifyInvokeSatisticRequest
//
// @return QueryVerifyInvokeSatisticResponse
func (client *Client) QueryVerifyInvokeSatistic(request *QueryVerifyInvokeSatisticRequest) (_result *QueryVerifyInvokeSatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryVerifyInvokeSatisticResponse{}
	_body, _err := client.QueryVerifyInvokeSatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an ID Verification whitelist.
//
// @param tmpReq - RemoveWhiteListSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveWhiteListSettingResponse
func (client *Client) RemoveWhiteListSettingWithOptions(tmpReq *RemoveWhiteListSettingRequest, runtime *dara.RuntimeOptions) (_result *RemoveWhiteListSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveWhiteListSettingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IdsShrink) {
		query["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveWhiteListSetting"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveWhiteListSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an ID Verification whitelist.
//
// @param request - RemoveWhiteListSettingRequest
//
// @return RemoveWhiteListSettingResponse
func (client *Client) RemoveWhiteListSetting(request *RemoveWhiteListSettingRequest) (_result *RemoveWhiteListSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveWhiteListSettingResponse{}
	_body, _err := client.RemoveWhiteListSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a China Finance Certification Initiative (CFCI) scenario.
//
// Description:
//
// Updates the information of a China Finance Certification Initiative (CFCI) scenario based on the scenario ID.
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST.
//
// @param request - UpdateAntCloudAuthSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAntCloudAuthSceneResponse
func (client *Client) UpdateAntCloudAuthSceneWithOptions(request *UpdateAntCloudAuthSceneRequest, runtime *dara.RuntimeOptions) (_result *UpdateAntCloudAuthSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BindMiniProgram) {
		query["BindMiniProgram"] = request.BindMiniProgram
	}

	if !dara.IsNil(request.CheckFileBody) {
		query["CheckFileBody"] = request.CheckFileBody
	}

	if !dara.IsNil(request.CheckFileName) {
		query["CheckFileName"] = request.CheckFileName
	}

	if !dara.IsNil(request.DeviceRiskPlus) {
		query["DeviceRiskPlus"] = request.DeviceRiskPlus
	}

	if !dara.IsNil(request.MiniProgramName) {
		query["MiniProgramName"] = request.MiniProgramName
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.ReturnPicCount) {
		query["ReturnPicCount"] = request.ReturnPicCount
	}

	if !dara.IsNil(request.ReturnVideoLength) {
		query["ReturnVideoLength"] = request.ReturnVideoLength
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.SceneName) {
		query["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StoreImage) {
		query["StoreImage"] = request.StoreImage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAntCloudAuthScene"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAntCloudAuthSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a China Finance Certification Initiative (CFCI) scenario.
//
// Description:
//
// Updates the information of a China Finance Certification Initiative (CFCI) scenario based on the scenario ID.
//
// - Service endpoint: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST.
//
// @param request - UpdateAntCloudAuthSceneRequest
//
// @return UpdateAntCloudAuthSceneResponse
func (client *Client) UpdateAntCloudAuthScene(request *UpdateAntCloudAuthSceneRequest) (_result *UpdateAntCloudAuthSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAntCloudAuthSceneResponse{}
	_body, _err := client.UpdateAntCloudAuthSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the intent scenario configuration.
//
// Description:
//
// - Request method: HTTPS POST.
//
// - Request URL: cloudauth.aliyuncs.com.
//
// @param request - UpdateSceneConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSceneConfigResponse
func (client *Client) UpdateSceneConfigWithOptions(request *UpdateSceneConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateSceneConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	if !dara.IsNil(request.SceneId) {
		body["sceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSceneConfig"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSceneConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the intent scenario configuration.
//
// Description:
//
// - Request method: HTTPS POST.
//
// - Request URL: cloudauth.aliyuncs.com.
//
// @param request - UpdateSceneConfigRequest
//
// @return UpdateSceneConfigResponse
func (client *Client) UpdateSceneConfig(request *UpdateSceneConfigRequest) (_result *UpdateSceneConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSceneConfigResponse{}
	_body, _err := client.UpdateSceneConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Identifies the five key attributes of a vehicle.
//
// Description:
//
// Queries basic vehicle information by license plate number and vehicle type.
//
// @param request - Vehicle5ItemQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Vehicle5ItemQueryResponse
func (client *Client) Vehicle5ItemQueryWithOptions(request *Vehicle5ItemQueryRequest, runtime *dara.RuntimeOptions) (_result *Vehicle5ItemQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ParamType) {
		query["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.VehicleNum) {
		query["VehicleNum"] = request.VehicleNum
	}

	if !dara.IsNil(request.VehicleType) {
		query["VehicleType"] = request.VehicleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Vehicle5ItemQuery"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Vehicle5ItemQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Identifies the five key attributes of a vehicle.
//
// Description:
//
// Queries basic vehicle information by license plate number and vehicle type.
//
// @param request - Vehicle5ItemQueryRequest
//
// @return Vehicle5ItemQueryResponse
func (client *Client) Vehicle5ItemQuery(request *Vehicle5ItemQueryRequest) (_result *Vehicle5ItemQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &Vehicle5ItemQueryResponse{}
	_body, _err := client.Vehicle5ItemQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the insurance date of a vehicle.
//
// Description:
//
// Queries the insurance date of a vehicle by license plate number, vehicle type, and vehicle identification number (VIN).
//
// @param request - VehicleInsureQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VehicleInsureQueryResponse
func (client *Client) VehicleInsureQueryWithOptions(request *VehicleInsureQueryRequest, runtime *dara.RuntimeOptions) (_result *VehicleInsureQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ParamType) {
		query["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.VehicleNum) {
		query["VehicleNum"] = request.VehicleNum
	}

	if !dara.IsNil(request.VehicleType) {
		query["VehicleType"] = request.VehicleType
	}

	if !dara.IsNil(request.Vin) {
		query["Vin"] = request.Vin
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VehicleInsureQuery"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VehicleInsureQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the insurance date of a vehicle.
//
// Description:
//
// Queries the insurance date of a vehicle by license plate number, vehicle type, and vehicle identification number (VIN).
//
// @param request - VehicleInsureQueryRequest
//
// @return VehicleInsureQueryResponse
func (client *Client) VehicleInsureQuery(request *VehicleInsureQueryRequest) (_result *VehicleInsureQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VehicleInsureQueryResponse{}
	_body, _err := client.VehicleInsureQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies vehicle element consistency.
//
// Description:
//
// Verifies the consistency of the name, ID card number, license plate number, and vehicle type.
//
// @param request - VehicleMetaVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VehicleMetaVerifyResponse
func (client *Client) VehicleMetaVerifyWithOptions(request *VehicleMetaVerifyRequest, runtime *dara.RuntimeOptions) (_result *VehicleMetaVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.ParamType) {
		query["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.VehicleNum) {
		query["VehicleNum"] = request.VehicleNum
	}

	if !dara.IsNil(request.VehicleType) {
		query["VehicleType"] = request.VehicleType
	}

	if !dara.IsNil(request.VerifyMetaType) {
		query["VerifyMetaType"] = request.VerifyMetaType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VehicleMetaVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VehicleMetaVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies vehicle element consistency.
//
// Description:
//
// Verifies the consistency of the name, ID card number, license plate number, and vehicle type.
//
// @param request - VehicleMetaVerifyRequest
//
// @return VehicleMetaVerifyResponse
func (client *Client) VehicleMetaVerify(request *VehicleMetaVerifyRequest) (_result *VehicleMetaVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VehicleMetaVerifyResponse{}
	_body, _err := client.VehicleMetaVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies vehicle element information (enhanced edition).
//
// Description:
//
// Verifies the consistency of the name, ID card number, license plate number, and vehicle type, and supports returning vehicle details.
//
// @param request - VehicleMetaVerifyV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VehicleMetaVerifyV2Response
func (client *Client) VehicleMetaVerifyV2WithOptions(request *VehicleMetaVerifyV2Request, runtime *dara.RuntimeOptions) (_result *VehicleMetaVerifyV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.ParamType) {
		query["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.VehicleNum) {
		query["VehicleNum"] = request.VehicleNum
	}

	if !dara.IsNil(request.VehicleType) {
		query["VehicleType"] = request.VehicleType
	}

	if !dara.IsNil(request.VerifyMetaType) {
		query["VerifyMetaType"] = request.VerifyMetaType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VehicleMetaVerifyV2"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VehicleMetaVerifyV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies vehicle element information (enhanced edition).
//
// Description:
//
// Verifies the consistency of the name, ID card number, license plate number, and vehicle type, and supports returning vehicle details.
//
// @param request - VehicleMetaVerifyV2Request
//
// @return VehicleMetaVerifyV2Response
func (client *Client) VehicleMetaVerifyV2(request *VehicleMetaVerifyV2Request) (_result *VehicleMetaVerifyV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VehicleMetaVerifyV2Response{}
	_body, _err := client.VehicleMetaVerifyV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Identifies vehicle information.
//
// Description:
//
// Queries detailed vehicle information by license plate number and vehicle type.
//
// @param request - VehicleQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VehicleQueryResponse
func (client *Client) VehicleQueryWithOptions(request *VehicleQueryRequest, runtime *dara.RuntimeOptions) (_result *VehicleQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ParamType) {
		query["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.VehicleNum) {
		query["VehicleNum"] = request.VehicleNum
	}

	if !dara.IsNil(request.VehicleType) {
		query["VehicleType"] = request.VehicleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VehicleQuery"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VehicleQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Identifies vehicle information.
//
// Description:
//
// Queries detailed vehicle information by license plate number and vehicle type.
//
// @param request - VehicleQueryRequest
//
// @return VehicleQueryResponse
func (client *Client) VehicleQuery(request *VehicleQueryRequest) (_result *VehicleQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VehicleQueryResponse{}
	_body, _err := client.VehicleQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a name, ID card number, facial photo, and optional front and back photos of the ID card for ID Verification, and synchronously returns the result.
//
// Description:
//
// Before you begin: Before calling this API, ensure that you have completed the required preparations. For more information, see [Server-side integration preparations](https://help.aliyun.com/document_detail/127471.html).
//
// Request method: HTTPS POST and GET.
//
// Operation description: The caller\\"s server submits verification materials to the ID Verification service for comparison and validation. The result is returned synchronously.
//
// Applicable scope: This operation is applicable only to server-side-only verification solutions.
//
// Image URL description:
//
// - HTTP or HTTPS URL: Publicly accessible HTTP or HTTPS URLs are supported. For example, `http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg`.
//
// - OSS URL: If the caller\\"s images are local files, Alibaba Cloud provides an upload SDK that allows you to upload images to the OSS bucket designated by the ID Verification service and obtain the corresponding OSS URL to use as the image URL parameter. If your business requires the upload SDK, [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/cloudauth/today) to contact us.
//
// Image restrictions:
//
// - Relative paths or absolute paths of local images are not supported.
//
// - Keep each image within 2 MB to avoid algorithm fetch timeouts.
//
// - The face area in the image must be at least 64 × 64 pixels.
//
// - The request body has an 8 MB size limit. Ensure that the total size of all images and other information in the request does not exceed this limit.
//
// @param request - VerifyMaterialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyMaterialResponse
func (client *Client) VerifyMaterialWithOptions(request *VerifyMaterialRequest, runtime *dara.RuntimeOptions) (_result *VerifyMaterialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.FaceImageUrl) {
		query["FaceImageUrl"] = request.FaceImageUrl
	}

	if !dara.IsNil(request.IdCardBackImageUrl) {
		query["IdCardBackImageUrl"] = request.IdCardBackImageUrl
	}

	if !dara.IsNil(request.IdCardFrontImageUrl) {
		query["IdCardFrontImageUrl"] = request.IdCardFrontImageUrl
	}

	if !dara.IsNil(request.IdCardNumber) {
		query["IdCardNumber"] = request.IdCardNumber
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyMaterial"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyMaterialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a name, ID card number, facial photo, and optional front and back photos of the ID card for ID Verification, and synchronously returns the result.
//
// Description:
//
// Before you begin: Before calling this API, ensure that you have completed the required preparations. For more information, see [Server-side integration preparations](https://help.aliyun.com/document_detail/127471.html).
//
// Request method: HTTPS POST and GET.
//
// Operation description: The caller\\"s server submits verification materials to the ID Verification service for comparison and validation. The result is returned synchronously.
//
// Applicable scope: This operation is applicable only to server-side-only verification solutions.
//
// Image URL description:
//
// - HTTP or HTTPS URL: Publicly accessible HTTP or HTTPS URLs are supported. For example, `http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg`.
//
// - OSS URL: If the caller\\"s images are local files, Alibaba Cloud provides an upload SDK that allows you to upload images to the OSS bucket designated by the ID Verification service and obtain the corresponding OSS URL to use as the image URL parameter. If your business requires the upload SDK, [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/cloudauth/today) to contact us.
//
// Image restrictions:
//
// - Relative paths or absolute paths of local images are not supported.
//
// - Keep each image within 2 MB to avoid algorithm fetch timeouts.
//
// - The face area in the image must be at least 64 × 64 pixels.
//
// - The request body has an 8 MB size limit. Ensure that the total size of all images and other information in the request does not exceed this limit.
//
// @param request - VerifyMaterialRequest
//
// @return VerifyMaterialResponse
func (client *Client) VerifyMaterial(request *VerifyMaterialRequest) (_result *VerifyMaterialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyMaterialResponse{}
	_body, _err := client.VerifyMaterialWithOptions(request, runtime)
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
