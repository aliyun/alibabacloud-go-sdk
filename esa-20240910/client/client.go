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
	client.Endpoint, _err = client.GetEndpoint(dara.String("esa"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Activates the client based on the certificate ID.
//
// @param request - ActivateClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateClientCertificateResponse
func (client *Client) ActivateClientCertificateWithOptions(request *ActivateClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *ActivateClientCertificateResponse, _err error) {
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
		Action:      dara.String("ActivateClientCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates the client based on the certificate ID.
//
// @param request - ActivateClientCertificateRequest
//
// @return ActivateClientCertificateResponse
func (client *Client) ActivateClientCertificate(request *ActivateClientCertificateRequest) (_result *ActivateClientCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ActivateClientCertificateResponse{}
	_body, _err := client.ActivateClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Enable Version Management
//
// @param request - ActivateVersionManagementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateVersionManagementResponse
func (client *Client) ActivateVersionManagementWithOptions(request *ActivateVersionManagementRequest, runtime *dara.RuntimeOptions) (_result *ActivateVersionManagementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActivateVersionManagement"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateVersionManagementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Enable Version Management
//
// @param request - ActivateVersionManagementRequest
//
// @return ActivateVersionManagementResponse
func (client *Client) ActivateVersionManagement(request *ActivateVersionManagementRequest) (_result *ActivateVersionManagementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ActivateVersionManagementResponse{}
	_body, _err := client.ActivateVersionManagementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Applies for a free SSL certificate.
//
// @param request - ApplyCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyCertificateResponse
func (client *Client) ApplyCertificateWithOptions(request *ApplyCertificateRequest, runtime *dara.RuntimeOptions) (_result *ApplyCertificateResponse, _err error) {
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
		Action:      dara.String("ApplyCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies for a free SSL certificate.
//
// @param request - ApplyCertificateRequest
//
// @return ApplyCertificateResponse
func (client *Client) ApplyCertificate(request *ApplyCertificateRequest) (_result *ApplyCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyCertificateResponse{}
	_body, _err := client.ApplyCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds DNS records of different record types at a time..
//
// Description:
//
// This operation allows you to create or update multiple DNS records at a time. It is suitable for managing a large number of DNS configurations. Supported record types include but are not limited to A/AAAA, CNAME, NS, MX, TXT, CAA, SRV, and URI. The operation allows you to configure the priority, flag, tag, and weight for DNS records. In addition, for specific types of records, such as CERT, SSHFP, SMIMEA, and TLSA, advanced settings such as certificate information and encryption algorithms are also supported.
//
// Successful and failed records along with error messages are listed in the response.
//
// @param tmpReq - BatchCreateRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateRecordsResponse
func (client *Client) BatchCreateRecordsWithOptions(tmpReq *BatchCreateRecordsRequest, runtime *dara.RuntimeOptions) (_result *BatchCreateRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchCreateRecordsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RecordList) {
		request.RecordListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecordList, dara.String("RecordList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RecordListShrink) {
		query["RecordList"] = request.RecordListShrink
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCreateRecords"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCreateRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds DNS records of different record types at a time..
//
// Description:
//
// This operation allows you to create or update multiple DNS records at a time. It is suitable for managing a large number of DNS configurations. Supported record types include but are not limited to A/AAAA, CNAME, NS, MX, TXT, CAA, SRV, and URI. The operation allows you to configure the priority, flag, tag, and weight for DNS records. In addition, for specific types of records, such as CERT, SSHFP, SMIMEA, and TLSA, advanced settings such as certificate information and encryption algorithms are also supported.
//
// Successful and failed records along with error messages are listed in the response.
//
// @param request - BatchCreateRecordsRequest
//
// @return BatchCreateRecordsResponse
func (client *Client) BatchCreateRecords(request *BatchCreateRecordsRequest) (_result *BatchCreateRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchCreateRecordsResponse{}
	_body, _err := client.BatchCreateRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Batch Create WAF Rules
//
// @param tmpReq - BatchCreateWafRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateWafRulesResponse
func (client *Client) BatchCreateWafRulesWithOptions(tmpReq *BatchCreateWafRulesRequest, runtime *dara.RuntimeOptions) (_result *BatchCreateWafRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchCreateWafRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Configs) {
		request.ConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Configs, dara.String("Configs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Shared) {
		request.SharedShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Shared, dara.String("Shared"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigsShrink) {
		body["Configs"] = request.ConfigsShrink
	}

	if !dara.IsNil(request.Phase) {
		body["Phase"] = request.Phase
	}

	if !dara.IsNil(request.RulesetId) {
		body["RulesetId"] = request.RulesetId
	}

	if !dara.IsNil(request.SharedShrink) {
		body["Shared"] = request.SharedShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCreateWafRules"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCreateWafRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Batch Create WAF Rules
//
// @param request - BatchCreateWafRulesRequest
//
// @return BatchCreateWafRulesResponse
func (client *Client) BatchCreateWafRules(request *BatchCreateWafRulesRequest) (_result *BatchCreateWafRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchCreateWafRulesResponse{}
	_body, _err := client.BatchCreateWafRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes key-value pairs from a namespace at a time based on keys.
//
// @param tmpReq - BatchDeleteKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteKvResponse
func (client *Client) BatchDeleteKvWithOptions(tmpReq *BatchDeleteKvRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteKvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchDeleteKvShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Keys) {
		request.KeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keys, dara.String("Keys"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.KeysShrink) {
		body["Keys"] = request.KeysShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteKv"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteKvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes key-value pairs from a namespace at a time based on keys.
//
// @param request - BatchDeleteKvRequest
//
// @return BatchDeleteKvResponse
func (client *Client) BatchDeleteKv(request *BatchDeleteKvRequest) (_result *BatchDeleteKvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteKvResponse{}
	_body, _err := client.BatchDeleteKvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple key-value pairs from a namespace at a time based on specified keys. The request body can be up to 100 MB.
//
// Description:
//
// This operation allows you to upload a larger request body than by using [BatchDeleteKv](https://help.aliyun.com/document_detail/2850204.html). For small request bodies, we recommend that you use [BatchDeleteKv](https://help.aliyun.com/document_detail/2850204.html) to minimize the server processing time. This operation must be called by using SDKs. The following sample code uses the Golang SDK and BatchDeleteKvWithHighCapacityAdvance to call the operation.
//
//	func TestBatchDeleteWithHighCapacity() error {
//
//		// Initialize the configurations.
//
//		cfg := new(openapi.Config)
//
//		cfg.SetAccessKeyId("xxxxxxxxx")
//
//		cfg.SetAccessKeySecret("xxxxxxxxxx")
//
//		cli, err := NewClient(cfg)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		runtime := &util.RuntimeOptions{}
//
//		// Construct a request for deleting key-value pairs at a time.
//
//		namespace := "test_batch_put"
//
//		rawReq := BatchDeleteKvRequest{
//
//			Namespace: &namespace,
//
//		}
//
//		for i := 0; i < 10000; i++ {
//
//			key := fmt.Sprintf("test_key_%d", i)
//
//			rawReq.Keys = append(rawReq.Keys, &key)
//
//		}
//
//		payload, err := json.Marshal(rawReq)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		// If the payload is greater than 2 MB, call the BatchDeleteKvWithHighCapacity operation for deletion.
//
//		reqHighCapacity := BatchDeleteKvWithHighCapacityAdvanceRequest{
//
//			Namespace: &namespace,
//
//			UrlObject: bytes.NewReader(payload),
//
//		}
//
//		resp, err := cli.BatchDeleteKvWithHighCapacityAdvance(&reqHighCapacity, runtime)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		return nil
//
//	}
//
// @param request - BatchDeleteKvWithHighCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteKvWithHighCapacityResponse
func (client *Client) BatchDeleteKvWithHighCapacityWithOptions(request *BatchDeleteKvWithHighCapacityRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteKvWithHighCapacityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteKvWithHighCapacity"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteKvWithHighCapacityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple key-value pairs from a namespace at a time based on specified keys. The request body can be up to 100 MB.
//
// Description:
//
// This operation allows you to upload a larger request body than by using [BatchDeleteKv](https://help.aliyun.com/document_detail/2850204.html). For small request bodies, we recommend that you use [BatchDeleteKv](https://help.aliyun.com/document_detail/2850204.html) to minimize the server processing time. This operation must be called by using SDKs. The following sample code uses the Golang SDK and BatchDeleteKvWithHighCapacityAdvance to call the operation.
//
//	func TestBatchDeleteWithHighCapacity() error {
//
//		// Initialize the configurations.
//
//		cfg := new(openapi.Config)
//
//		cfg.SetAccessKeyId("xxxxxxxxx")
//
//		cfg.SetAccessKeySecret("xxxxxxxxxx")
//
//		cli, err := NewClient(cfg)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		runtime := &util.RuntimeOptions{}
//
//		// Construct a request for deleting key-value pairs at a time.
//
//		namespace := "test_batch_put"
//
//		rawReq := BatchDeleteKvRequest{
//
//			Namespace: &namespace,
//
//		}
//
//		for i := 0; i < 10000; i++ {
//
//			key := fmt.Sprintf("test_key_%d", i)
//
//			rawReq.Keys = append(rawReq.Keys, &key)
//
//		}
//
//		payload, err := json.Marshal(rawReq)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		// If the payload is greater than 2 MB, call the BatchDeleteKvWithHighCapacity operation for deletion.
//
//		reqHighCapacity := BatchDeleteKvWithHighCapacityAdvanceRequest{
//
//			Namespace: &namespace,
//
//			UrlObject: bytes.NewReader(payload),
//
//		}
//
//		resp, err := cli.BatchDeleteKvWithHighCapacityAdvance(&reqHighCapacity, runtime)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		return nil
//
//	}
//
// @param request - BatchDeleteKvWithHighCapacityRequest
//
// @return BatchDeleteKvWithHighCapacityResponse
func (client *Client) BatchDeleteKvWithHighCapacity(request *BatchDeleteKvWithHighCapacityRequest) (_result *BatchDeleteKvWithHighCapacityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteKvWithHighCapacityResponse{}
	_body, _err := client.BatchDeleteKvWithHighCapacityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteKvWithHighCapacityAdvance(request *BatchDeleteKvWithHighCapacityAdvanceRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteKvWithHighCapacityResponse, _err error) {
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
		"Product":  dara.String("ESA"),
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
	batchDeleteKvWithHighCapacityReq := &BatchDeleteKvWithHighCapacityRequest{}
	openapiutil.Convert(request, batchDeleteKvWithHighCapacityReq)
	if !dara.IsNil(request.UrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.UrlObject,
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
		batchDeleteKvWithHighCapacityReq.Url = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	batchDeleteKvWithHighCapacityResp, _err := client.BatchDeleteKvWithHighCapacityWithOptions(batchDeleteKvWithHighCapacityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = batchDeleteKvWithHighCapacityResp
	return _result, _err
}

// Summary:
//
// # Batch Get Expression Matches
//
// @param tmpReq - BatchGetExpressionFieldsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetExpressionFieldsResponse
func (client *Client) BatchGetExpressionFieldsWithOptions(tmpReq *BatchGetExpressionFieldsRequest, runtime *dara.RuntimeOptions) (_result *BatchGetExpressionFieldsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchGetExpressionFieldsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Expressions) {
		request.ExpressionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Expressions, dara.String("Expressions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PlanNameEn) {
		query["PlanNameEn"] = request.PlanNameEn
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExpressionsShrink) {
		body["Expressions"] = request.ExpressionsShrink
	}

	if !dara.IsNil(request.Kind) {
		body["Kind"] = request.Kind
	}

	if !dara.IsNil(request.Phase) {
		body["Phase"] = request.Phase
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGetExpressionFields"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGetExpressionFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Batch Get Expression Matches
//
// @param request - BatchGetExpressionFieldsRequest
//
// @return BatchGetExpressionFieldsResponse
func (client *Client) BatchGetExpressionFields(request *BatchGetExpressionFieldsRequest) (_result *BatchGetExpressionFieldsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchGetExpressionFieldsResponse{}
	_body, _err := client.BatchGetExpressionFieldsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures key-value pairs for a namespace at a time based on specified keys.
//
// @param tmpReq - BatchPutKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchPutKvResponse
func (client *Client) BatchPutKvWithOptions(tmpReq *BatchPutKvRequest, runtime *dara.RuntimeOptions) (_result *BatchPutKvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchPutKvShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.KvList) {
		request.KvListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KvList, dara.String("KvList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.KvListShrink) {
		body["KvList"] = request.KvListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchPutKv"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchPutKvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures key-value pairs for a namespace at a time based on specified keys.
//
// @param request - BatchPutKvRequest
//
// @return BatchPutKvResponse
func (client *Client) BatchPutKv(request *BatchPutKvRequest) (_result *BatchPutKvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchPutKvResponse{}
	_body, _err := client.BatchPutKvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures key-value pairs for a namespace at a time based on specified keys. The request body can be up to 100 MB.
//
// Description:
//
// This operation allows you to upload a larger request body than by using [BatchPutKv](https://help.aliyun.com/document_detail/2850203.html). For small request bodies, we recommend that you use [BatchPutKv](https://help.aliyun.com/document_detail/2850203.html) to minimize the server processing time. This operation must be called by using SDKs. The following sample code uses the Golang SDK and BatchPutKvWithHighCapacityAdvance to call the operation.
//
//	func TestBatchPutKvWithHighCapacity() error {
//
//		// Initialize the configurations.
//
//		cfg := new(openapi.Config)
//
//		cfg.SetAccessKeyId("xxxxxxxxx")
//
//		cfg.SetAccessKeySecret("xxxxxxxxxx")
//
//		cli, err := NewClient(cfg)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		runtime := &util.RuntimeOptions{}
//
//		// Construct a request for uploading key-value pairs at a time.
//
//		namespace := "test_batch_put"
//
//		numKv := 10000
//
//		kvList := make([]*BatchPutKvRequestKvList, numKv)
//
//		test_value := strings.Repeat("a", 10*1024)
//
//		for i := 0; i < numKv; i++ {
//
//			key := fmt.Sprintf("test_key_%d", i)
//
//			value := test_value
//
//			kvList[i] = &BatchPutKvRequestKvList{
//
//				Key:   &key,
//
//				Value: &value,
//
//			}
//
//		}
//
//		rawReq := BatchPutKvRequest{
//
//			Namespace: &namespace,
//
//			KvList:    kvList,
//
//		}
//
//		payload, err := json.Marshal(rawReq)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		// If the payload is greater than 2 MB, call the BatchPutKvWithHighCapacity operation for upload.
//
//		reqHighCapacity := BatchPutKvWithHighCapacityAdvanceRequest{
//
//			Namespace: &namespace,
//
//			UrlObject: bytes.NewReader(payload),
//
//		}
//
//		resp, err := cli.BatchPutKvWithHighCapacityAdvance(&reqHighCapacity, runtime)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		return nil
//
//	}
//
// @param request - BatchPutKvWithHighCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchPutKvWithHighCapacityResponse
func (client *Client) BatchPutKvWithHighCapacityWithOptions(request *BatchPutKvWithHighCapacityRequest, runtime *dara.RuntimeOptions) (_result *BatchPutKvWithHighCapacityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchPutKvWithHighCapacity"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchPutKvWithHighCapacityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures key-value pairs for a namespace at a time based on specified keys. The request body can be up to 100 MB.
//
// Description:
//
// This operation allows you to upload a larger request body than by using [BatchPutKv](https://help.aliyun.com/document_detail/2850203.html). For small request bodies, we recommend that you use [BatchPutKv](https://help.aliyun.com/document_detail/2850203.html) to minimize the server processing time. This operation must be called by using SDKs. The following sample code uses the Golang SDK and BatchPutKvWithHighCapacityAdvance to call the operation.
//
//	func TestBatchPutKvWithHighCapacity() error {
//
//		// Initialize the configurations.
//
//		cfg := new(openapi.Config)
//
//		cfg.SetAccessKeyId("xxxxxxxxx")
//
//		cfg.SetAccessKeySecret("xxxxxxxxxx")
//
//		cli, err := NewClient(cfg)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		runtime := &util.RuntimeOptions{}
//
//		// Construct a request for uploading key-value pairs at a time.
//
//		namespace := "test_batch_put"
//
//		numKv := 10000
//
//		kvList := make([]*BatchPutKvRequestKvList, numKv)
//
//		test_value := strings.Repeat("a", 10*1024)
//
//		for i := 0; i < numKv; i++ {
//
//			key := fmt.Sprintf("test_key_%d", i)
//
//			value := test_value
//
//			kvList[i] = &BatchPutKvRequestKvList{
//
//				Key:   &key,
//
//				Value: &value,
//
//			}
//
//		}
//
//		rawReq := BatchPutKvRequest{
//
//			Namespace: &namespace,
//
//			KvList:    kvList,
//
//		}
//
//		payload, err := json.Marshal(rawReq)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		// If the payload is greater than 2 MB, call the BatchPutKvWithHighCapacity operation for upload.
//
//		reqHighCapacity := BatchPutKvWithHighCapacityAdvanceRequest{
//
//			Namespace: &namespace,
//
//			UrlObject: bytes.NewReader(payload),
//
//		}
//
//		resp, err := cli.BatchPutKvWithHighCapacityAdvance(&reqHighCapacity, runtime)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		return nil
//
//	}
//
// @param request - BatchPutKvWithHighCapacityRequest
//
// @return BatchPutKvWithHighCapacityResponse
func (client *Client) BatchPutKvWithHighCapacity(request *BatchPutKvWithHighCapacityRequest) (_result *BatchPutKvWithHighCapacityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchPutKvWithHighCapacityResponse{}
	_body, _err := client.BatchPutKvWithHighCapacityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchPutKvWithHighCapacityAdvance(request *BatchPutKvWithHighCapacityAdvanceRequest, runtime *dara.RuntimeOptions) (_result *BatchPutKvWithHighCapacityResponse, _err error) {
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
		"Product":  dara.String("ESA"),
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
	batchPutKvWithHighCapacityReq := &BatchPutKvWithHighCapacityRequest{}
	openapiutil.Convert(request, batchPutKvWithHighCapacityReq)
	if !dara.IsNil(request.UrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.UrlObject,
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
		batchPutKvWithHighCapacityReq.Url = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	batchPutKvWithHighCapacityResp, _err := client.BatchPutKvWithHighCapacityWithOptions(batchPutKvWithHighCapacityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = batchPutKvWithHighCapacityResp
	return _result, _err
}

// Summary:
//
// Modifies multiple rules in a specific Web Application Firewall (WAF) ruleset at a time.
//
// @param tmpReq - BatchUpdateWafRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateWafRulesResponse
func (client *Client) BatchUpdateWafRulesWithOptions(tmpReq *BatchUpdateWafRulesRequest, runtime *dara.RuntimeOptions) (_result *BatchUpdateWafRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchUpdateWafRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Configs) {
		request.ConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Configs, dara.String("Configs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Shared) {
		request.SharedShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Shared, dara.String("Shared"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigsShrink) {
		body["Configs"] = request.ConfigsShrink
	}

	if !dara.IsNil(request.Phase) {
		body["Phase"] = request.Phase
	}

	if !dara.IsNil(request.RulesetId) {
		body["RulesetId"] = request.RulesetId
	}

	if !dara.IsNil(request.SharedShrink) {
		body["Shared"] = request.SharedShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUpdateWafRules"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUpdateWafRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies multiple rules in a specific Web Application Firewall (WAF) ruleset at a time.
//
// @param request - BatchUpdateWafRulesRequest
//
// @return BatchUpdateWafRulesResponse
func (client *Client) BatchUpdateWafRules(request *BatchUpdateWafRulesRequest) (_result *BatchUpdateWafRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchUpdateWafRulesResponse{}
	_body, _err := client.BatchUpdateWafRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Blocks URLs.
//
// @param tmpReq - BlockObjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BlockObjectResponse
func (client *Client) BlockObjectWithOptions(tmpReq *BlockObjectRequest, runtime *dara.RuntimeOptions) (_result *BlockObjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BlockObjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Content) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, dara.String("Content"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ContentShrink) {
		query["Content"] = request.ContentShrink
	}

	if !dara.IsNil(request.Extension) {
		query["Extension"] = request.Extension
	}

	if !dara.IsNil(request.Maxage) {
		query["Maxage"] = request.Maxage
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BlockObject"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BlockObjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Blocks URLs.
//
// @param request - BlockObjectRequest
//
// @return BlockObjectResponse
func (client *Client) BlockObject(request *BlockObjectRequest) (_result *BlockObjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BlockObjectResponse{}
	_body, _err := client.BlockObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检查实时日志slr角色是否已创建
//
// @param request - CheckAssumeSlrRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAssumeSlrRoleResponse
func (client *Client) CheckAssumeSlrRoleWithOptions(runtime *dara.RuntimeOptions) (_result *CheckAssumeSlrRoleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("CheckAssumeSlrRole"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckAssumeSlrRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查实时日志slr角色是否已创建
//
// @return CheckAssumeSlrRoleResponse
func (client *Client) CheckAssumeSlrRole() (_result *CheckAssumeSlrRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckAssumeSlrRoleResponse{}
	_body, _err := client.CheckAssumeSlrRoleWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether a specified website name is available.
//
// @param request - CheckSiteNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckSiteNameResponse
func (client *Client) CheckSiteNameWithOptions(request *CheckSiteNameRequest, runtime *dara.RuntimeOptions) (_result *CheckSiteNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteName) {
		query["SiteName"] = request.SiteName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckSiteName"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckSiteNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a specified website name is available.
//
// @param request - CheckSiteNameRequest
//
// @return CheckSiteNameResponse
func (client *Client) CheckSiteName(request *CheckSiteNameRequest) (_result *CheckSiteNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckSiteNameResponse{}
	_body, _err := client.CheckSiteNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks the name of a real-time log delivery task.
//
// @param request - CheckSiteProjectNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckSiteProjectNameResponse
func (client *Client) CheckSiteProjectNameWithOptions(request *CheckSiteProjectNameRequest, runtime *dara.RuntimeOptions) (_result *CheckSiteProjectNameResponse, _err error) {
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
		Action:      dara.String("CheckSiteProjectName"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckSiteProjectNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the name of a real-time log delivery task.
//
// @param request - CheckSiteProjectNameRequest
//
// @return CheckSiteProjectNameResponse
func (client *Client) CheckSiteProjectName(request *CheckSiteProjectNameRequest) (_result *CheckSiteProjectNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckSiteProjectNameResponse{}
	_body, _err := client.CheckSiteProjectNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks the name of a real-time log delivery task by account.
//
// @param request - CheckUserProjectNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckUserProjectNameResponse
func (client *Client) CheckUserProjectNameWithOptions(request *CheckUserProjectNameRequest, runtime *dara.RuntimeOptions) (_result *CheckUserProjectNameResponse, _err error) {
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
		Action:      dara.String("CheckUserProjectName"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckUserProjectNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the name of a real-time log delivery task by account.
//
// @param request - CheckUserProjectNameRequest
//
// @return CheckUserProjectNameResponse
func (client *Client) CheckUserProjectName(request *CheckUserProjectNameRequest) (_result *CheckUserProjectNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckUserProjectNameResponse{}
	_body, _err := client.CheckUserProjectNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Commits the unstable code in the staging environment to generate an official code version.
//
// @param request - CommitRoutineStagingCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CommitRoutineStagingCodeResponse
func (client *Client) CommitRoutineStagingCodeWithOptions(request *CommitRoutineStagingCodeRequest, runtime *dara.RuntimeOptions) (_result *CommitRoutineStagingCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CodeDescription) {
		body["CodeDescription"] = request.CodeDescription
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CommitRoutineStagingCode"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CommitRoutineStagingCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Commits the unstable code in the staging environment to generate an official code version.
//
// @param request - CommitRoutineStagingCodeRequest
//
// @return CommitRoutineStagingCodeResponse
func (client *Client) CommitRoutineStagingCode(request *CommitRoutineStagingCodeRequest) (_result *CommitRoutineStagingCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CommitRoutineStagingCodeResponse{}
	_body, _err := client.CommitRoutineStagingCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create a new site cache configuration
//
// @param request - CreateCacheRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCacheRuleResponse
func (client *Client) CreateCacheRuleWithOptions(request *CreateCacheRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateCacheRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdditionalCacheablePorts) {
		query["AdditionalCacheablePorts"] = request.AdditionalCacheablePorts
	}

	if !dara.IsNil(request.BrowserCacheMode) {
		query["BrowserCacheMode"] = request.BrowserCacheMode
	}

	if !dara.IsNil(request.BrowserCacheTtl) {
		query["BrowserCacheTtl"] = request.BrowserCacheTtl
	}

	if !dara.IsNil(request.BypassCache) {
		query["BypassCache"] = request.BypassCache
	}

	if !dara.IsNil(request.CacheDeceptionArmor) {
		query["CacheDeceptionArmor"] = request.CacheDeceptionArmor
	}

	if !dara.IsNil(request.CacheReserveEligibility) {
		query["CacheReserveEligibility"] = request.CacheReserveEligibility
	}

	if !dara.IsNil(request.CheckPresenceCookie) {
		query["CheckPresenceCookie"] = request.CheckPresenceCookie
	}

	if !dara.IsNil(request.CheckPresenceHeader) {
		query["CheckPresenceHeader"] = request.CheckPresenceHeader
	}

	if !dara.IsNil(request.EdgeCacheMode) {
		query["EdgeCacheMode"] = request.EdgeCacheMode
	}

	if !dara.IsNil(request.EdgeCacheTtl) {
		query["EdgeCacheTtl"] = request.EdgeCacheTtl
	}

	if !dara.IsNil(request.EdgeStatusCodeCacheTtl) {
		query["EdgeStatusCodeCacheTtl"] = request.EdgeStatusCodeCacheTtl
	}

	if !dara.IsNil(request.IncludeCookie) {
		query["IncludeCookie"] = request.IncludeCookie
	}

	if !dara.IsNil(request.IncludeHeader) {
		query["IncludeHeader"] = request.IncludeHeader
	}

	if !dara.IsNil(request.PostBodyCacheKey) {
		query["PostBodyCacheKey"] = request.PostBodyCacheKey
	}

	if !dara.IsNil(request.PostBodySizeLimit) {
		query["PostBodySizeLimit"] = request.PostBodySizeLimit
	}

	if !dara.IsNil(request.PostCache) {
		query["PostCache"] = request.PostCache
	}

	if !dara.IsNil(request.QueryString) {
		query["QueryString"] = request.QueryString
	}

	if !dara.IsNil(request.QueryStringMode) {
		query["QueryStringMode"] = request.QueryStringMode
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.ServeStale) {
		query["ServeStale"] = request.ServeStale
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	if !dara.IsNil(request.SortQueryStringForCache) {
		query["SortQueryStringForCache"] = request.SortQueryStringForCache
	}

	if !dara.IsNil(request.UserDeviceType) {
		query["UserDeviceType"] = request.UserDeviceType
	}

	if !dara.IsNil(request.UserGeo) {
		query["UserGeo"] = request.UserGeo
	}

	if !dara.IsNil(request.UserLanguage) {
		query["UserLanguage"] = request.UserLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCacheRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCacheRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a new site cache configuration
//
// @param request - CreateCacheRuleRequest
//
// @return CreateCacheRuleResponse
func (client *Client) CreateCacheRule(request *CreateCacheRuleRequest) (_result *CreateCacheRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCacheRuleResponse{}
	_body, _err := client.CreateCacheRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uses the ESA-managed certificate authority (CA) to issue client certificates.
//
// @param request - CreateClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClientCertificateResponse
func (client *Client) CreateClientCertificateWithOptions(request *CreateClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *CreateClientCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CSR) {
		body["CSR"] = request.CSR
	}

	if !dara.IsNil(request.PkeyType) {
		body["PkeyType"] = request.PkeyType
	}

	if !dara.IsNil(request.ValidityDays) {
		body["ValidityDays"] = request.ValidityDays
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateClientCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uses the ESA-managed certificate authority (CA) to issue client certificates.
//
// @param request - CreateClientCertificateRequest
//
// @return CreateClientCertificateResponse
func (client *Client) CreateClientCertificate(request *CreateClientCertificateRequest) (_result *CreateClientCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateClientCertificateResponse{}
	_body, _err := client.CreateClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Add a compression rule
//
// @param request - CreateCompressionRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCompressionRuleResponse
func (client *Client) CreateCompressionRuleWithOptions(request *CreateCompressionRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateCompressionRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Brotli) {
		query["Brotli"] = request.Brotli
	}

	if !dara.IsNil(request.Gzip) {
		query["Gzip"] = request.Gzip
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	if !dara.IsNil(request.Zstd) {
		query["Zstd"] = request.Zstd
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCompressionRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCompressionRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add a compression rule
//
// @param request - CreateCompressionRuleRequest
//
// @return CreateCompressionRuleResponse
func (client *Client) CreateCompressionRule(request *CreateCompressionRuleRequest) (_result *CreateCompressionRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCompressionRuleResponse{}
	_body, _err := client.CreateCompressionRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an account-level custom scenario policy. You can execute a policy after you associate the policy with a website.
//
// @param request - CreateCustomScenePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomScenePolicyResponse
func (client *Client) CreateCustomScenePolicyWithOptions(request *CreateCustomScenePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomScenePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Objects) {
		query["Objects"] = request.Objects
	}

	if !dara.IsNil(request.SiteIds) {
		query["SiteIds"] = request.SiteIds
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomScenePolicy"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomScenePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an account-level custom scenario policy. You can execute a policy after you associate the policy with a website.
//
// @param request - CreateCustomScenePolicyRequest
//
// @return CreateCustomScenePolicyResponse
func (client *Client) CreateCustomScenePolicy(request *CreateCustomScenePolicyRequest) (_result *CreateCustomScenePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCustomScenePolicyResponse{}
	_body, _err := client.CreateCustomScenePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a containerized application. You can deploy and release a version of the application across points of presence (POPs).
//
// @param request - CreateEdgeContainerAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEdgeContainerAppResponse
func (client *Client) CreateEdgeContainerAppWithOptions(request *CreateEdgeContainerAppRequest, runtime *dara.RuntimeOptions) (_result *CreateEdgeContainerAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HealthCheckFailTimes) {
		body["HealthCheckFailTimes"] = request.HealthCheckFailTimes
	}

	if !dara.IsNil(request.HealthCheckHost) {
		body["HealthCheckHost"] = request.HealthCheckHost
	}

	if !dara.IsNil(request.HealthCheckHttpCode) {
		body["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		body["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckMethod) {
		body["HealthCheckMethod"] = request.HealthCheckMethod
	}

	if !dara.IsNil(request.HealthCheckPort) {
		body["HealthCheckPort"] = request.HealthCheckPort
	}

	if !dara.IsNil(request.HealthCheckSuccTimes) {
		body["HealthCheckSuccTimes"] = request.HealthCheckSuccTimes
	}

	if !dara.IsNil(request.HealthCheckTimeout) {
		body["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !dara.IsNil(request.HealthCheckType) {
		body["HealthCheckType"] = request.HealthCheckType
	}

	if !dara.IsNil(request.HealthCheckURI) {
		body["HealthCheckURI"] = request.HealthCheckURI
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Remarks) {
		body["Remarks"] = request.Remarks
	}

	if !dara.IsNil(request.ServicePort) {
		body["ServicePort"] = request.ServicePort
	}

	if !dara.IsNil(request.TargetPort) {
		body["TargetPort"] = request.TargetPort
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEdgeContainerApp"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEdgeContainerAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a containerized application. You can deploy and release a version of the application across points of presence (POPs).
//
// @param request - CreateEdgeContainerAppRequest
//
// @return CreateEdgeContainerAppResponse
func (client *Client) CreateEdgeContainerApp(request *CreateEdgeContainerAppRequest) (_result *CreateEdgeContainerAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEdgeContainerAppResponse{}
	_body, _err := client.CreateEdgeContainerAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a domain name with a containerized application. This way, requests destined for the associated domain name are forwarded to the application.
//
// @param request - CreateEdgeContainerAppRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEdgeContainerAppRecordResponse
func (client *Client) CreateEdgeContainerAppRecordWithOptions(request *CreateEdgeContainerAppRecordRequest, runtime *dara.RuntimeOptions) (_result *CreateEdgeContainerAppRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.RecordName) {
		body["RecordName"] = request.RecordName
	}

	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEdgeContainerAppRecord"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEdgeContainerAppRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a domain name with a containerized application. This way, requests destined for the associated domain name are forwarded to the application.
//
// @param request - CreateEdgeContainerAppRecordRequest
//
// @return CreateEdgeContainerAppRecordResponse
func (client *Client) CreateEdgeContainerAppRecord(request *CreateEdgeContainerAppRecordRequest) (_result *CreateEdgeContainerAppRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEdgeContainerAppRecordResponse{}
	_body, _err := client.CreateEdgeContainerAppRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a version for a containerized application. You can iterate the application based on the version.
//
// @param tmpReq - CreateEdgeContainerAppVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEdgeContainerAppVersionResponse
func (client *Client) CreateEdgeContainerAppVersionWithOptions(tmpReq *CreateEdgeContainerAppVersionRequest, runtime *dara.RuntimeOptions) (_result *CreateEdgeContainerAppVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateEdgeContainerAppVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Containers) {
		request.ContainersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Containers, dara.String("Containers"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ContainersShrink) {
		body["Containers"] = request.ContainersShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Remarks) {
		body["Remarks"] = request.Remarks
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEdgeContainerAppVersion"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEdgeContainerAppVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a version for a containerized application. You can iterate the application based on the version.
//
// @param request - CreateEdgeContainerAppVersionRequest
//
// @return CreateEdgeContainerAppVersionResponse
func (client *Client) CreateEdgeContainerAppVersion(request *CreateEdgeContainerAppVersionRequest) (_result *CreateEdgeContainerAppVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEdgeContainerAppVersionResponse{}
	_body, _err := client.CreateEdgeContainerAppVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds the configuration of modifying HTTP request headers for a website.
//
// @param tmpReq - CreateHttpIncomingRequestHeaderModificationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHttpIncomingRequestHeaderModificationRuleResponse
func (client *Client) CreateHttpIncomingRequestHeaderModificationRuleWithOptions(tmpReq *CreateHttpIncomingRequestHeaderModificationRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateHttpIncomingRequestHeaderModificationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RequestHeaderModification) {
		request.RequestHeaderModificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RequestHeaderModification, dara.String("RequestHeaderModification"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RequestHeaderModificationShrink) {
		query["RequestHeaderModification"] = request.RequestHeaderModificationShrink
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHttpIncomingRequestHeaderModificationRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHttpIncomingRequestHeaderModificationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds the configuration of modifying HTTP request headers for a website.
//
// @param request - CreateHttpIncomingRequestHeaderModificationRuleRequest
//
// @return CreateHttpIncomingRequestHeaderModificationRuleResponse
func (client *Client) CreateHttpIncomingRequestHeaderModificationRule(request *CreateHttpIncomingRequestHeaderModificationRuleRequest) (_result *CreateHttpIncomingRequestHeaderModificationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHttpIncomingRequestHeaderModificationRuleResponse{}
	_body, _err := client.CreateHttpIncomingRequestHeaderModificationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds the configuration of modifying HTTP response headers for a website.
//
// @param tmpReq - CreateHttpIncomingResponseHeaderModificationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHttpIncomingResponseHeaderModificationRuleResponse
func (client *Client) CreateHttpIncomingResponseHeaderModificationRuleWithOptions(tmpReq *CreateHttpIncomingResponseHeaderModificationRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateHttpIncomingResponseHeaderModificationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResponseHeaderModification) {
		request.ResponseHeaderModificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResponseHeaderModification, dara.String("ResponseHeaderModification"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ResponseHeaderModificationShrink) {
		query["ResponseHeaderModification"] = request.ResponseHeaderModificationShrink
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHttpIncomingResponseHeaderModificationRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHttpIncomingResponseHeaderModificationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds the configuration of modifying HTTP response headers for a website.
//
// @param request - CreateHttpIncomingResponseHeaderModificationRuleRequest
//
// @return CreateHttpIncomingResponseHeaderModificationRuleResponse
func (client *Client) CreateHttpIncomingResponseHeaderModificationRule(request *CreateHttpIncomingResponseHeaderModificationRuleRequest) (_result *CreateHttpIncomingResponseHeaderModificationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHttpIncomingResponseHeaderModificationRuleResponse{}
	_body, _err := client.CreateHttpIncomingResponseHeaderModificationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Add HTTP Request Header Rule
//
// @param tmpReq - CreateHttpRequestHeaderModificationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHttpRequestHeaderModificationRuleResponse
func (client *Client) CreateHttpRequestHeaderModificationRuleWithOptions(tmpReq *CreateHttpRequestHeaderModificationRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateHttpRequestHeaderModificationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateHttpRequestHeaderModificationRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RequestHeaderModification) {
		request.RequestHeaderModificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RequestHeaderModification, dara.String("RequestHeaderModification"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RequestHeaderModificationShrink) {
		query["RequestHeaderModification"] = request.RequestHeaderModificationShrink
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHttpRequestHeaderModificationRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHttpRequestHeaderModificationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add HTTP Request Header Rule
//
// @param request - CreateHttpRequestHeaderModificationRuleRequest
//
// @return CreateHttpRequestHeaderModificationRuleResponse
func (client *Client) CreateHttpRequestHeaderModificationRule(request *CreateHttpRequestHeaderModificationRuleRequest) (_result *CreateHttpRequestHeaderModificationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHttpRequestHeaderModificationRuleResponse{}
	_body, _err := client.CreateHttpRequestHeaderModificationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Add HTTP Response Header Rule
//
// @param tmpReq - CreateHttpResponseHeaderModificationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHttpResponseHeaderModificationRuleResponse
func (client *Client) CreateHttpResponseHeaderModificationRuleWithOptions(tmpReq *CreateHttpResponseHeaderModificationRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateHttpResponseHeaderModificationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateHttpResponseHeaderModificationRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResponseHeaderModification) {
		request.ResponseHeaderModificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResponseHeaderModification, dara.String("ResponseHeaderModification"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ResponseHeaderModificationShrink) {
		query["ResponseHeaderModification"] = request.ResponseHeaderModificationShrink
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHttpResponseHeaderModificationRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHttpResponseHeaderModificationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add HTTP Response Header Rule
//
// @param request - CreateHttpResponseHeaderModificationRuleRequest
//
// @return CreateHttpResponseHeaderModificationRuleResponse
func (client *Client) CreateHttpResponseHeaderModificationRule(request *CreateHttpResponseHeaderModificationRuleRequest) (_result *CreateHttpResponseHeaderModificationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHttpResponseHeaderModificationRuleResponse{}
	_body, _err := client.CreateHttpResponseHeaderModificationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create a new site HTTPS application configuration
//
// @param request - CreateHttpsApplicationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHttpsApplicationConfigurationResponse
func (client *Client) CreateHttpsApplicationConfigurationWithOptions(request *CreateHttpsApplicationConfigurationRequest, runtime *dara.RuntimeOptions) (_result *CreateHttpsApplicationConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AltSvc) {
		query["AltSvc"] = request.AltSvc
	}

	if !dara.IsNil(request.AltSvcClear) {
		query["AltSvcClear"] = request.AltSvcClear
	}

	if !dara.IsNil(request.AltSvcMa) {
		query["AltSvcMa"] = request.AltSvcMa
	}

	if !dara.IsNil(request.AltSvcPersist) {
		query["AltSvcPersist"] = request.AltSvcPersist
	}

	if !dara.IsNil(request.Hsts) {
		query["Hsts"] = request.Hsts
	}

	if !dara.IsNil(request.HstsIncludeSubdomains) {
		query["HstsIncludeSubdomains"] = request.HstsIncludeSubdomains
	}

	if !dara.IsNil(request.HstsMaxAge) {
		query["HstsMaxAge"] = request.HstsMaxAge
	}

	if !dara.IsNil(request.HstsPreload) {
		query["HstsPreload"] = request.HstsPreload
	}

	if !dara.IsNil(request.HttpsForce) {
		query["HttpsForce"] = request.HttpsForce
	}

	if !dara.IsNil(request.HttpsForceCode) {
		query["HttpsForceCode"] = request.HttpsForceCode
	}

	if !dara.IsNil(request.HttpsNoSniDeny) {
		query["HttpsNoSniDeny"] = request.HttpsNoSniDeny
	}

	if !dara.IsNil(request.HttpsSniVerify) {
		query["HttpsSniVerify"] = request.HttpsSniVerify
	}

	if !dara.IsNil(request.HttpsSniWhitelist) {
		query["HttpsSniWhitelist"] = request.HttpsSniWhitelist
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHttpsApplicationConfiguration"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHttpsApplicationConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a new site HTTPS application configuration
//
// @param request - CreateHttpsApplicationConfigurationRequest
//
// @return CreateHttpsApplicationConfigurationResponse
func (client *Client) CreateHttpsApplicationConfiguration(request *CreateHttpsApplicationConfigurationRequest) (_result *CreateHttpsApplicationConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHttpsApplicationConfigurationResponse{}
	_body, _err := client.CreateHttpsApplicationConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create a new site HTTPS basic configuration
//
// @param request - CreateHttpsBasicConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHttpsBasicConfigurationResponse
func (client *Client) CreateHttpsBasicConfigurationWithOptions(request *CreateHttpsBasicConfigurationRequest, runtime *dara.RuntimeOptions) (_result *CreateHttpsBasicConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ciphersuite) {
		query["Ciphersuite"] = request.Ciphersuite
	}

	if !dara.IsNil(request.CiphersuiteGroup) {
		query["CiphersuiteGroup"] = request.CiphersuiteGroup
	}

	if !dara.IsNil(request.Http2) {
		query["Http2"] = request.Http2
	}

	if !dara.IsNil(request.Http3) {
		query["Http3"] = request.Http3
	}

	if !dara.IsNil(request.Https) {
		query["Https"] = request.Https
	}

	if !dara.IsNil(request.OcspStapling) {
		query["OcspStapling"] = request.OcspStapling
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.Tls10) {
		query["Tls10"] = request.Tls10
	}

	if !dara.IsNil(request.Tls11) {
		query["Tls11"] = request.Tls11
	}

	if !dara.IsNil(request.Tls12) {
		query["Tls12"] = request.Tls12
	}

	if !dara.IsNil(request.Tls13) {
		query["Tls13"] = request.Tls13
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHttpsBasicConfiguration"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHttpsBasicConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a new site HTTPS basic configuration
//
// @param request - CreateHttpsBasicConfigurationRequest
//
// @return CreateHttpsBasicConfigurationResponse
func (client *Client) CreateHttpsBasicConfiguration(request *CreateHttpsBasicConfigurationRequest) (_result *CreateHttpsBasicConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHttpsBasicConfigurationResponse{}
	_body, _err := client.CreateHttpsBasicConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Add Site Image Transformation Configuration
//
// @param request - CreateImageTransformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageTransformResponse
func (client *Client) CreateImageTransformWithOptions(request *CreateImageTransformRequest, runtime *dara.RuntimeOptions) (_result *CreateImageTransformResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageTransform"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageTransformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add Site Image Transformation Configuration
//
// @param request - CreateImageTransformRequest
//
// @return CreateImageTransformResponse
func (client *Client) CreateImageTransform(request *CreateImageTransformRequest) (_result *CreateImageTransformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateImageTransformResponse{}
	_body, _err := client.CreateImageTransformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a namespace in your Alibaba Cloud account.
//
// @param request - CreateKvNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKvNamespaceResponse
func (client *Client) CreateKvNamespaceWithOptions(request *CreateKvNamespaceRequest, runtime *dara.RuntimeOptions) (_result *CreateKvNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKvNamespace"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKvNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a namespace in your Alibaba Cloud account.
//
// @param request - CreateKvNamespaceRequest
//
// @return CreateKvNamespaceResponse
func (client *Client) CreateKvNamespace(request *CreateKvNamespaceRequest) (_result *CreateKvNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateKvNamespaceResponse{}
	_body, _err := client.CreateKvNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a list. Lists are used for the referencing of values in the rules engine to implement complex logic and control in security policies.
//
// @param tmpReq - CreateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateListResponse
func (client *Client) CreateListWithOptions(tmpReq *CreateListRequest, runtime *dara.RuntimeOptions) (_result *CreateListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Items) {
		request.ItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Items, dara.String("Items"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ItemsShrink) {
		body["Items"] = request.ItemsShrink
	}

	if !dara.IsNil(request.Kind) {
		body["Kind"] = request.Kind
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateList"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a list. Lists are used for the referencing of values in the rules engine to implement complex logic and control in security policies.
//
// @param request - CreateListRequest
//
// @return CreateListResponse
func (client *Client) CreateList(request *CreateListRequest) (_result *CreateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateListResponse{}
	_body, _err := client.CreateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Add a New Load Balancer
//
// Description:
//
// Through this API, users can configure load balancing services according to their business needs, including but not limited to adaptive routing, weighted round-robin, rule matching, health checks, and more, to achieve effective traffic management and optimization.
//
// @param tmpReq - CreateLoadBalancerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoadBalancerResponse
func (client *Client) CreateLoadBalancerWithOptions(tmpReq *CreateLoadBalancerRequest, runtime *dara.RuntimeOptions) (_result *CreateLoadBalancerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateLoadBalancerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AdaptiveRouting) {
		request.AdaptiveRoutingShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdaptiveRouting, dara.String("AdaptiveRouting"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DefaultPools) {
		request.DefaultPoolsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DefaultPools, dara.String("DefaultPools"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Monitor) {
		request.MonitorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Monitor, dara.String("Monitor"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RandomSteering) {
		request.RandomSteeringShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RandomSteering, dara.String("RandomSteering"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Rules) {
		request.RulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rules, dara.String("Rules"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdaptiveRoutingShrink) {
		query["AdaptiveRouting"] = request.AdaptiveRoutingShrink
	}

	if !dara.IsNil(request.DefaultPoolsShrink) {
		query["DefaultPools"] = request.DefaultPoolsShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.FallbackPool) {
		query["FallbackPool"] = request.FallbackPool
	}

	if !dara.IsNil(request.MonitorShrink) {
		query["Monitor"] = request.MonitorShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RandomSteeringShrink) {
		query["RandomSteering"] = request.RandomSteeringShrink
	}

	if !dara.IsNil(request.RegionPools) {
		query["RegionPools"] = request.RegionPools
	}

	if !dara.IsNil(request.RulesShrink) {
		query["Rules"] = request.RulesShrink
	}

	if !dara.IsNil(request.SessionAffinity) {
		query["SessionAffinity"] = request.SessionAffinity
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SteeringPolicy) {
		query["SteeringPolicy"] = request.SteeringPolicy
	}

	if !dara.IsNil(request.SubRegionPools) {
		query["SubRegionPools"] = request.SubRegionPools
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLoadBalancer"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add a New Load Balancer
//
// Description:
//
// Through this API, users can configure load balancing services according to their business needs, including but not limited to adaptive routing, weighted round-robin, rule matching, health checks, and more, to achieve effective traffic management and optimization.
//
// @param request - CreateLoadBalancerRequest
//
// @return CreateLoadBalancerResponse
func (client *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (_result *CreateLoadBalancerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLoadBalancerResponse{}
	_body, _err := client.CreateLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create a new site network optimization configuration
//
// @param request - CreateNetworkOptimizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkOptimizationResponse
func (client *Client) CreateNetworkOptimizationWithOptions(request *CreateNetworkOptimizationRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkOptimizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Grpc) {
		query["Grpc"] = request.Grpc
	}

	if !dara.IsNil(request.Http2Origin) {
		query["Http2Origin"] = request.Http2Origin
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	if !dara.IsNil(request.SmartRouting) {
		query["SmartRouting"] = request.SmartRouting
	}

	if !dara.IsNil(request.UploadMaxFilesize) {
		query["UploadMaxFilesize"] = request.UploadMaxFilesize
	}

	if !dara.IsNil(request.Websocket) {
		query["Websocket"] = request.Websocket
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetworkOptimization"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetworkOptimizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a new site network optimization configuration
//
// @param request - CreateNetworkOptimizationRequest
//
// @return CreateNetworkOptimizationResponse
func (client *Client) CreateNetworkOptimization(request *CreateNetworkOptimizationRequest) (_result *CreateNetworkOptimizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNetworkOptimizationResponse{}
	_body, _err := client.CreateNetworkOptimizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Add a new origin address pool
//
// Description:
//
// Multiple origins can be added under the origin address, supporting domain names, IPs, OSS, S3, and other types of origins. It supports authentication for OSS and S3 type origins.
//
// @param tmpReq - CreateOriginPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOriginPoolResponse
func (client *Client) CreateOriginPoolWithOptions(tmpReq *CreateOriginPoolRequest, runtime *dara.RuntimeOptions) (_result *CreateOriginPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateOriginPoolShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Origins) {
		request.OriginsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Origins, dara.String("Origins"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OriginsShrink) {
		query["Origins"] = request.OriginsShrink
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOriginPool"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOriginPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add a new origin address pool
//
// Description:
//
// Multiple origins can be added under the origin address, supporting domain names, IPs, OSS, S3, and other types of origins. It supports authentication for OSS and S3 type origins.
//
// @param request - CreateOriginPoolRequest
//
// @return CreateOriginPoolResponse
func (client *Client) CreateOriginPool(request *CreateOriginPoolRequest) (_result *CreateOriginPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOriginPoolResponse{}
	_body, _err := client.CreateOriginPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables origin protection.
//
// @param request - CreateOriginProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOriginProtectionResponse
func (client *Client) CreateOriginProtectionWithOptions(request *CreateOriginProtectionRequest, runtime *dara.RuntimeOptions) (_result *CreateOriginProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoConfirmIPList) {
		query["AutoConfirmIPList"] = request.AutoConfirmIPList
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOriginProtection"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOriginProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables origin protection.
//
// @param request - CreateOriginProtectionRequest
//
// @return CreateOriginProtectionResponse
func (client *Client) CreateOriginProtection(request *CreateOriginProtectionRequest) (_result *CreateOriginProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOriginProtectionResponse{}
	_body, _err := client.CreateOriginProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create a new origin rule configuration for the site
//
// @param request - CreateOriginRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOriginRuleResponse
func (client *Client) CreateOriginRuleWithOptions(request *CreateOriginRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateOriginRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DnsRecord) {
		query["DnsRecord"] = request.DnsRecord
	}

	if !dara.IsNil(request.Follow302Enable) {
		query["Follow302Enable"] = request.Follow302Enable
	}

	if !dara.IsNil(request.Follow302MaxTries) {
		query["Follow302MaxTries"] = request.Follow302MaxTries
	}

	if !dara.IsNil(request.Follow302RetainArgs) {
		query["Follow302RetainArgs"] = request.Follow302RetainArgs
	}

	if !dara.IsNil(request.Follow302RetainHeader) {
		query["Follow302RetainHeader"] = request.Follow302RetainHeader
	}

	if !dara.IsNil(request.Follow302TargetHost) {
		query["Follow302TargetHost"] = request.Follow302TargetHost
	}

	if !dara.IsNil(request.OriginHost) {
		query["OriginHost"] = request.OriginHost
	}

	if !dara.IsNil(request.OriginHttpPort) {
		query["OriginHttpPort"] = request.OriginHttpPort
	}

	if !dara.IsNil(request.OriginHttpsPort) {
		query["OriginHttpsPort"] = request.OriginHttpsPort
	}

	if !dara.IsNil(request.OriginMtls) {
		query["OriginMtls"] = request.OriginMtls
	}

	if !dara.IsNil(request.OriginReadTimeout) {
		query["OriginReadTimeout"] = request.OriginReadTimeout
	}

	if !dara.IsNil(request.OriginScheme) {
		query["OriginScheme"] = request.OriginScheme
	}

	if !dara.IsNil(request.OriginSni) {
		query["OriginSni"] = request.OriginSni
	}

	if !dara.IsNil(request.OriginVerify) {
		query["OriginVerify"] = request.OriginVerify
	}

	if !dara.IsNil(request.Range) {
		query["Range"] = request.Range
	}

	if !dara.IsNil(request.RangeChunkSize) {
		query["RangeChunkSize"] = request.RangeChunkSize
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOriginRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOriginRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a new origin rule configuration for the site
//
// @param request - CreateOriginRuleRequest
//
// @return CreateOriginRuleResponse
func (client *Client) CreateOriginRule(request *CreateOriginRuleRequest) (_result *CreateOriginRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOriginRuleResponse{}
	_body, _err := client.CreateOriginRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom error page, which is displayed when a request is blocked by Web Application Firewall (WAF). You can configure the HTML content, page type, and description, and submit the Base64-encoded page content.
//
// @param request - CreatePageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePageResponse
func (client *Client) CreatePageWithOptions(request *CreatePageRequest, runtime *dara.RuntimeOptions) (_result *CreatePageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.ContentType) {
		body["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePage"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom error page, which is displayed when a request is blocked by Web Application Firewall (WAF). You can configure the HTML content, page type, and description, and submit the Base64-encoded page content.
//
// @param request - CreatePageRequest
//
// @return CreatePageResponse
func (client *Client) CreatePage(request *CreatePageRequest) (_result *CreatePageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePageResponse{}
	_body, _err := client.CreatePageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a DNS record for a specific website.
//
// @param tmpReq - CreateRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecordResponse
func (client *Client) CreateRecordWithOptions(tmpReq *CreateRecordRequest, runtime *dara.RuntimeOptions) (_result *CreateRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuthConf) {
		request.AuthConfShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthConf, dara.String("AuthConf"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Data) {
		request.DataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Data, dara.String("Data"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthConfShrink) {
		query["AuthConf"] = request.AuthConfShrink
	}

	if !dara.IsNil(request.BizName) {
		query["BizName"] = request.BizName
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DataShrink) {
		query["Data"] = request.DataShrink
	}

	if !dara.IsNil(request.HostPolicy) {
		query["HostPolicy"] = request.HostPolicy
	}

	if !dara.IsNil(request.Proxied) {
		query["Proxied"] = request.Proxied
	}

	if !dara.IsNil(request.RecordName) {
		query["RecordName"] = request.RecordName
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRecord"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a DNS record for a specific website.
//
// @param request - CreateRecordRequest
//
// @return CreateRecordResponse
func (client *Client) CreateRecord(request *CreateRecordRequest) (_result *CreateRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRecordResponse{}
	_body, _err := client.CreateRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Add a Redirect Rule
//
// @param request - CreateRedirectRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRedirectRuleResponse
func (client *Client) CreateRedirectRuleWithOptions(request *CreateRedirectRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateRedirectRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReserveQueryString) {
		query["ReserveQueryString"] = request.ReserveQueryString
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	if !dara.IsNil(request.StatusCode) {
		query["StatusCode"] = request.StatusCode
	}

	if !dara.IsNil(request.TargetUrl) {
		query["TargetUrl"] = request.TargetUrl
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRedirectRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRedirectRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add a Redirect Rule
//
// @param request - CreateRedirectRuleRequest
//
// @return CreateRedirectRuleResponse
func (client *Client) CreateRedirectRule(request *CreateRedirectRuleRequest) (_result *CreateRedirectRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRedirectRuleResponse{}
	_body, _err := client.CreateRedirectRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Add Rewrite URL Rule
//
// @param request - CreateRewriteUrlRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRewriteUrlRuleResponse
func (client *Client) CreateRewriteUrlRuleWithOptions(request *CreateRewriteUrlRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateRewriteUrlRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QueryString) {
		query["QueryString"] = request.QueryString
	}

	if !dara.IsNil(request.RewriteQueryStringType) {
		query["RewriteQueryStringType"] = request.RewriteQueryStringType
	}

	if !dara.IsNil(request.RewriteUriType) {
		query["RewriteUriType"] = request.RewriteUriType
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	if !dara.IsNil(request.Uri) {
		query["Uri"] = request.Uri
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRewriteUrlRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRewriteUrlRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add Rewrite URL Rule
//
// @param request - CreateRewriteUrlRuleRequest
//
// @return CreateRewriteUrlRuleResponse
func (client *Client) CreateRewriteUrlRule(request *CreateRewriteUrlRuleRequest) (_result *CreateRewriteUrlRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRewriteUrlRuleResponse{}
	_body, _err := client.CreateRewriteUrlRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a routine.
//
// @param request - CreateRoutineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoutineResponse
func (client *Client) CreateRoutineWithOptions(request *CreateRoutineRequest, runtime *dara.RuntimeOptions) (_result *CreateRoutineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.HasAssets) {
		body["HasAssets"] = request.HasAssets
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRoutine"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRoutineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a routine.
//
// @param request - CreateRoutineRequest
//
// @return CreateRoutineResponse
func (client *Client) CreateRoutine(request *CreateRoutineRequest) (_result *CreateRoutineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRoutineResponse{}
	_body, _err := client.CreateRoutineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Release the version of the function code in proportion to the specified environment.
//
// Description:
//
// ## [](#)Request description
//
//   - When you create a version for deployment, you can set the environment name `Env` parameter only to the test environment `staging` or the production environment `production`.
//
//   - `CodeVersions` parameter supports up to two versions of a phased release, and the sum of the proportions of these versions must be equal to 100%.
//
// @param tmpReq - CreateRoutineCodeDeploymentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoutineCodeDeploymentResponse
func (client *Client) CreateRoutineCodeDeploymentWithOptions(tmpReq *CreateRoutineCodeDeploymentRequest, runtime *dara.RuntimeOptions) (_result *CreateRoutineCodeDeploymentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateRoutineCodeDeploymentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CodeVersions) {
		request.CodeVersionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CodeVersions, dara.String("CodeVersions"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CodeVersionsShrink) {
		body["CodeVersions"] = request.CodeVersionsShrink
	}

	if !dara.IsNil(request.Env) {
		body["Env"] = request.Env
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Strategy) {
		body["Strategy"] = request.Strategy
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRoutineCodeDeployment"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRoutineCodeDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Release the version of the function code in proportion to the specified environment.
//
// Description:
//
// ## [](#)Request description
//
//   - When you create a version for deployment, you can set the environment name `Env` parameter only to the test environment `staging` or the production environment `production`.
//
//   - `CodeVersions` parameter supports up to two versions of a phased release, and the sum of the proportions of these versions must be equal to 100%.
//
// @param request - CreateRoutineCodeDeploymentRequest
//
// @return CreateRoutineCodeDeploymentResponse
func (client *Client) CreateRoutineCodeDeployment(request *CreateRoutineCodeDeploymentRequest) (_result *CreateRoutineCodeDeploymentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRoutineCodeDeploymentResponse{}
	_body, _err := client.CreateRoutineCodeDeploymentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a record to map a domain that is associated with a routine. This record is used to trigger the associated routine code.
//
// @param request - CreateRoutineRelatedRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoutineRelatedRecordResponse
func (client *Client) CreateRoutineRelatedRecordWithOptions(request *CreateRoutineRelatedRecordRequest, runtime *dara.RuntimeOptions) (_result *CreateRoutineRelatedRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RecordName) {
		body["RecordName"] = request.RecordName
	}

	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRoutineRelatedRecord"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRoutineRelatedRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a record to map a domain that is associated with a routine. This record is used to trigger the associated routine code.
//
// @param request - CreateRoutineRelatedRecordRequest
//
// @return CreateRoutineRelatedRecordResponse
func (client *Client) CreateRoutineRelatedRecord(request *CreateRoutineRelatedRecordRequest) (_result *CreateRoutineRelatedRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRoutineRelatedRecordResponse{}
	_body, _err := client.CreateRoutineRelatedRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds edge function routing configurations.
//
// @param request - CreateRoutineRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoutineRouteResponse
func (client *Client) CreateRoutineRouteWithOptions(request *CreateRoutineRouteRequest, runtime *dara.RuntimeOptions) (_result *CreateRoutineRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bypass) {
		query["Bypass"] = request.Bypass
	}

	if !dara.IsNil(request.Fallback) {
		query["Fallback"] = request.Fallback
	}

	if !dara.IsNil(request.RouteEnable) {
		query["RouteEnable"] = request.RouteEnable
	}

	if !dara.IsNil(request.RouteName) {
		query["RouteName"] = request.RouteName
	}

	if !dara.IsNil(request.RoutineName) {
		query["RoutineName"] = request.RoutineName
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRoutineRoute"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRoutineRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds edge function routing configurations.
//
// @param request - CreateRoutineRouteRequest
//
// @return CreateRoutineRouteResponse
func (client *Client) CreateRoutineRoute(request *CreateRoutineRouteRequest) (_result *CreateRoutineRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRoutineRouteResponse{}
	_body, _err := client.CreateRoutineRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建带Assets资源的Routine代码版本
//
// @param tmpReq - CreateRoutineWithAssetsCodeVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoutineWithAssetsCodeVersionResponse
func (client *Client) CreateRoutineWithAssetsCodeVersionWithOptions(tmpReq *CreateRoutineWithAssetsCodeVersionRequest, runtime *dara.RuntimeOptions) (_result *CreateRoutineWithAssetsCodeVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateRoutineWithAssetsCodeVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ConfOptions) {
		request.ConfOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfOptions, dara.String("ConfOptions"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BuildId) {
		body["BuildId"] = request.BuildId
	}

	if !dara.IsNil(request.CodeDescription) {
		body["CodeDescription"] = request.CodeDescription
	}

	if !dara.IsNil(request.ConfOptionsShrink) {
		body["ConfOptions"] = request.ConfOptionsShrink
	}

	if !dara.IsNil(request.ExtraInfo) {
		body["ExtraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRoutineWithAssetsCodeVersion"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRoutineWithAssetsCodeVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建带Assets资源的Routine代码版本
//
// @param request - CreateRoutineWithAssetsCodeVersionRequest
//
// @return CreateRoutineWithAssetsCodeVersionResponse
func (client *Client) CreateRoutineWithAssetsCodeVersion(request *CreateRoutineWithAssetsCodeVersionRequest) (_result *CreateRoutineWithAssetsCodeVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRoutineWithAssetsCodeVersionResponse{}
	_body, _err := client.CreateRoutineWithAssetsCodeVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates scheduled prefetch plans.
//
// @param tmpReq - CreateScheduledPreloadExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduledPreloadExecutionsResponse
func (client *Client) CreateScheduledPreloadExecutionsWithOptions(tmpReq *CreateScheduledPreloadExecutionsRequest, runtime *dara.RuntimeOptions) (_result *CreateScheduledPreloadExecutionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateScheduledPreloadExecutionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Executions) {
		request.ExecutionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Executions, dara.String("Executions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExecutionsShrink) {
		body["Executions"] = request.ExecutionsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScheduledPreloadExecutions"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScheduledPreloadExecutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates scheduled prefetch plans.
//
// @param request - CreateScheduledPreloadExecutionsRequest
//
// @return CreateScheduledPreloadExecutionsResponse
func (client *Client) CreateScheduledPreloadExecutions(request *CreateScheduledPreloadExecutionsRequest) (_result *CreateScheduledPreloadExecutionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateScheduledPreloadExecutionsResponse{}
	_body, _err := client.CreateScheduledPreloadExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a scheduled prefetch task.
//
// @param request - CreateScheduledPreloadJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduledPreloadJobResponse
func (client *Client) CreateScheduledPreloadJobWithOptions(request *CreateScheduledPreloadJobRequest, runtime *dara.RuntimeOptions) (_result *CreateScheduledPreloadJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InsertWay) {
		body["InsertWay"] = request.InsertWay
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OssUrl) {
		body["OssUrl"] = request.OssUrl
	}

	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.UrlList) {
		body["UrlList"] = request.UrlList
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScheduledPreloadJob"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScheduledPreloadJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a scheduled prefetch task.
//
// @param request - CreateScheduledPreloadJobRequest
//
// @return CreateScheduledPreloadJobResponse
func (client *Client) CreateScheduledPreloadJob(request *CreateScheduledPreloadJobRequest) (_result *CreateScheduledPreloadJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateScheduledPreloadJobResponse{}
	_body, _err := client.CreateScheduledPreloadJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a website.
//
// Description:
//
//	  Make sure that you have an available plan before you add a website.
//
//		- Make sure that your website domain name has an ICP filing if the location you want to specify covers the Chinese mainland.
//
// @param request - CreateSiteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSiteResponse
func (client *Client) CreateSiteWithOptions(request *CreateSiteRequest, runtime *dara.RuntimeOptions) (_result *CreateSiteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessType) {
		query["AccessType"] = request.AccessType
	}

	if !dara.IsNil(request.Coverage) {
		query["Coverage"] = request.Coverage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SiteName) {
		query["SiteName"] = request.SiteName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSite"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSiteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a website.
//
// Description:
//
//	  Make sure that you have an available plan before you add a website.
//
//		- Make sure that your website domain name has an ICP filing if the location you want to specify covers the Chinese mainland.
//
// @param request - CreateSiteRequest
//
// @return CreateSiteResponse
func (client *Client) CreateSite(request *CreateSiteRequest) (_result *CreateSiteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSiteResponse{}
	_body, _err := client.CreateSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds the configuration of custom request header, response header, and cookie fields that are used to capture logs of a website.
//
// Description:
//
//	  **Custom field limits**: The key name of a custom field can contain only letters, digits, underscores (_), and spaces. The key name cannot contain other characters. Otherwise, errors may occur.
//
//		- **Parameter passing**: Submit `SiteId`, `RequestHeaders`, `ResponseHeaders`, and `Cookies` by using `formData`. Each array element matches a custom field name.
//
//		- **(Required) SiteId**: Although `SiteId` is not marked as required in the Required column, you must specify a website ID by using this parameter when you can call this API operation.
//
// @param tmpReq - CreateSiteCustomLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSiteCustomLogResponse
func (client *Client) CreateSiteCustomLogWithOptions(tmpReq *CreateSiteCustomLogRequest, runtime *dara.RuntimeOptions) (_result *CreateSiteCustomLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSiteCustomLogShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Cookies) {
		request.CookiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Cookies, dara.String("Cookies"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RequestHeaders) {
		request.RequestHeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RequestHeaders, dara.String("RequestHeaders"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResponseHeaders) {
		request.ResponseHeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResponseHeaders, dara.String("ResponseHeaders"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CookiesShrink) {
		body["Cookies"] = request.CookiesShrink
	}

	if !dara.IsNil(request.RequestHeadersShrink) {
		body["RequestHeaders"] = request.RequestHeadersShrink
	}

	if !dara.IsNil(request.ResponseHeadersShrink) {
		body["ResponseHeaders"] = request.ResponseHeadersShrink
	}

	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSiteCustomLog"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSiteCustomLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds the configuration of custom request header, response header, and cookie fields that are used to capture logs of a website.
//
// Description:
//
//	  **Custom field limits**: The key name of a custom field can contain only letters, digits, underscores (_), and spaces. The key name cannot contain other characters. Otherwise, errors may occur.
//
//		- **Parameter passing**: Submit `SiteId`, `RequestHeaders`, `ResponseHeaders`, and `Cookies` by using `formData`. Each array element matches a custom field name.
//
//		- **(Required) SiteId**: Although `SiteId` is not marked as required in the Required column, you must specify a website ID by using this parameter when you can call this API operation.
//
// @param request - CreateSiteCustomLogRequest
//
// @return CreateSiteCustomLogResponse
func (client *Client) CreateSiteCustomLog(request *CreateSiteCustomLogRequest) (_result *CreateSiteCustomLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSiteCustomLogResponse{}
	_body, _err := client.CreateSiteCustomLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a real-time log delivery task.
//
// @param tmpReq - CreateSiteDeliveryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSiteDeliveryTaskResponse
func (client *Client) CreateSiteDeliveryTaskWithOptions(tmpReq *CreateSiteDeliveryTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateSiteDeliveryTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSiteDeliveryTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HttpDelivery) {
		request.HttpDeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HttpDelivery, dara.String("HttpDelivery"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.KafkaDelivery) {
		request.KafkaDeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KafkaDelivery, dara.String("KafkaDelivery"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OssDelivery) {
		request.OssDeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OssDelivery, dara.String("OssDelivery"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.S3Delivery) {
		request.S3DeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.S3Delivery, dara.String("S3Delivery"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SlsDelivery) {
		request.SlsDeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SlsDelivery, dara.String("SlsDelivery"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		body["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.DataCenter) {
		body["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.DeliveryType) {
		body["DeliveryType"] = request.DeliveryType
	}

	if !dara.IsNil(request.DiscardRate) {
		body["DiscardRate"] = request.DiscardRate
	}

	if !dara.IsNil(request.FieldName) {
		body["FieldName"] = request.FieldName
	}

	if !dara.IsNil(request.FilterVer) {
		body["FilterVer"] = request.FilterVer
	}

	if !dara.IsNil(request.HttpDeliveryShrink) {
		body["HttpDelivery"] = request.HttpDeliveryShrink
	}

	if !dara.IsNil(request.KafkaDeliveryShrink) {
		body["KafkaDelivery"] = request.KafkaDeliveryShrink
	}

	if !dara.IsNil(request.OssDeliveryShrink) {
		body["OssDelivery"] = request.OssDeliveryShrink
	}

	if !dara.IsNil(request.S3DeliveryShrink) {
		body["S3Delivery"] = request.S3DeliveryShrink
	}

	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SlsDeliveryShrink) {
		body["SlsDelivery"] = request.SlsDeliveryShrink
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSiteDeliveryTask"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSiteDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a real-time log delivery task.
//
// @param request - CreateSiteDeliveryTaskRequest
//
// @return CreateSiteDeliveryTaskResponse
func (client *Client) CreateSiteDeliveryTask(request *CreateSiteDeliveryTaskRequest) (_result *CreateSiteDeliveryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSiteDeliveryTaskResponse{}
	_body, _err := client.CreateSiteDeliveryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建一个实时日志slr角色
//
// @param request - CreateSlrRoleForRealtimeLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSlrRoleForRealtimeLogResponse
func (client *Client) CreateSlrRoleForRealtimeLogWithOptions(runtime *dara.RuntimeOptions) (_result *CreateSlrRoleForRealtimeLogResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSlrRoleForRealtimeLog"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSlrRoleForRealtimeLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一个实时日志slr角色
//
// @return CreateSlrRoleForRealtimeLogResponse
func (client *Client) CreateSlrRoleForRealtimeLog() (_result *CreateSlrRoleForRealtimeLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSlrRoleForRealtimeLogResponse{}
	_body, _err := client.CreateSlrRoleForRealtimeLogWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Transport Layer Application
//
// @param tmpReq - CreateTransportLayerApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransportLayerApplicationResponse
func (client *Client) CreateTransportLayerApplicationWithOptions(tmpReq *CreateTransportLayerApplicationRequest, runtime *dara.RuntimeOptions) (_result *CreateTransportLayerApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTransportLayerApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Rules) {
		request.RulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rules, dara.String("Rules"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossBorderOptimization) {
		query["CrossBorderOptimization"] = request.CrossBorderOptimization
	}

	if !dara.IsNil(request.IpAccessRule) {
		query["IpAccessRule"] = request.IpAccessRule
	}

	if !dara.IsNil(request.Ipv6) {
		query["Ipv6"] = request.Ipv6
	}

	if !dara.IsNil(request.RecordName) {
		query["RecordName"] = request.RecordName
	}

	if !dara.IsNil(request.RulesShrink) {
		query["Rules"] = request.RulesShrink
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.StaticIp) {
		query["StaticIp"] = request.StaticIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTransportLayerApplication"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTransportLayerApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Transport Layer Application
//
// @param request - CreateTransportLayerApplicationRequest
//
// @return CreateTransportLayerApplicationResponse
func (client *Client) CreateTransportLayerApplication(request *CreateTransportLayerApplicationRequest) (_result *CreateTransportLayerApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTransportLayerApplicationResponse{}
	_body, _err := client.CreateTransportLayerApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a web page monitoring configuration.
//
// @param request - CreateUrlObservationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUrlObservationResponse
func (client *Client) CreateUrlObservationWithOptions(request *CreateUrlObservationRequest, runtime *dara.RuntimeOptions) (_result *CreateUrlObservationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SdkType) {
		query["SdkType"] = request.SdkType
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUrlObservation"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUrlObservationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a web page monitoring configuration.
//
// @param request - CreateUrlObservationRequest
//
// @return CreateUrlObservationResponse
func (client *Client) CreateUrlObservation(request *CreateUrlObservationRequest) (_result *CreateUrlObservationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUrlObservationResponse{}
	_body, _err := client.CreateUrlObservationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a log delivery task to ship logs to the specified destination.
//
// Description:
//
// This API operation allows you to deliver logs to destinations such as Simple Log Service (SLS), HTTP servers, Object Storage Service (OSS), Amazon Simple Storage Service (S3), and Kafka. You can specify the task name, log fields to deliver, data center, discard rate, delivery type, and delivery details.
//
//   - **Field filtering**: Use the `FieldName` parameter to specify log fields to deliver.
//
//   - **Filtering rules**: Use the `FilterRules` parameter to pre-process and filter log data.
//
//   - **Diverse delivery destinations**: Logs can be delivered to different destinations. Configuration parameters vary with delivery destinations.
//
// ## [](#)Precautions
//
//   - Make sure that you have sufficient permissions to perform delivery tasks.
//
//   - If you enable encryption or authentication, properly configure corresponding parameters.
//
//   - Verify the syntax of `FilterRules` to make sure that filtering logic works as expected.
//
//   - Specify advanced settings such as the number of retries and timeout period based on your needs to have optimal delivery efficiency and stability.
//
// @param tmpReq - CreateUserDeliveryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserDeliveryTaskResponse
func (client *Client) CreateUserDeliveryTaskWithOptions(tmpReq *CreateUserDeliveryTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateUserDeliveryTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateUserDeliveryTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HttpDelivery) {
		request.HttpDeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HttpDelivery, dara.String("HttpDelivery"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.KafkaDelivery) {
		request.KafkaDeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KafkaDelivery, dara.String("KafkaDelivery"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OssDelivery) {
		request.OssDeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OssDelivery, dara.String("OssDelivery"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.S3Delivery) {
		request.S3DeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.S3Delivery, dara.String("S3Delivery"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SlsDelivery) {
		request.SlsDeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SlsDelivery, dara.String("SlsDelivery"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		body["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.DataCenter) {
		body["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.DeliveryType) {
		body["DeliveryType"] = request.DeliveryType
	}

	if !dara.IsNil(request.Details) {
		body["Details"] = request.Details
	}

	if !dara.IsNil(request.DiscardRate) {
		body["DiscardRate"] = request.DiscardRate
	}

	if !dara.IsNil(request.FieldName) {
		body["FieldName"] = request.FieldName
	}

	if !dara.IsNil(request.FilterVer) {
		body["FilterVer"] = request.FilterVer
	}

	if !dara.IsNil(request.HttpDeliveryShrink) {
		body["HttpDelivery"] = request.HttpDeliveryShrink
	}

	if !dara.IsNil(request.KafkaDeliveryShrink) {
		body["KafkaDelivery"] = request.KafkaDeliveryShrink
	}

	if !dara.IsNil(request.OssDeliveryShrink) {
		body["OssDelivery"] = request.OssDeliveryShrink
	}

	if !dara.IsNil(request.S3DeliveryShrink) {
		body["S3Delivery"] = request.S3DeliveryShrink
	}

	if !dara.IsNil(request.SlsDeliveryShrink) {
		body["SlsDelivery"] = request.SlsDeliveryShrink
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserDeliveryTask"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a log delivery task to ship logs to the specified destination.
//
// Description:
//
// This API operation allows you to deliver logs to destinations such as Simple Log Service (SLS), HTTP servers, Object Storage Service (OSS), Amazon Simple Storage Service (S3), and Kafka. You can specify the task name, log fields to deliver, data center, discard rate, delivery type, and delivery details.
//
//   - **Field filtering**: Use the `FieldName` parameter to specify log fields to deliver.
//
//   - **Filtering rules**: Use the `FilterRules` parameter to pre-process and filter log data.
//
//   - **Diverse delivery destinations**: Logs can be delivered to different destinations. Configuration parameters vary with delivery destinations.
//
// ## [](#)Precautions
//
//   - Make sure that you have sufficient permissions to perform delivery tasks.
//
//   - If you enable encryption or authentication, properly configure corresponding parameters.
//
//   - Verify the syntax of `FilterRules` to make sure that filtering logic works as expected.
//
//   - Specify advanced settings such as the number of retries and timeout period based on your needs to have optimal delivery efficiency and stability.
//
// @param request - CreateUserDeliveryTaskRequest
//
// @return CreateUserDeliveryTaskResponse
func (client *Client) CreateUserDeliveryTask(request *CreateUserDeliveryTaskRequest) (_result *CreateUserDeliveryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserDeliveryTaskResponse{}
	_body, _err := client.CreateUserDeliveryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add video processing configurations for a website.
//
// @param request - CreateVideoProcessingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVideoProcessingResponse
func (client *Client) CreateVideoProcessingWithOptions(request *CreateVideoProcessingRequest, runtime *dara.RuntimeOptions) (_result *CreateVideoProcessingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlvSeekEnd) {
		query["FlvSeekEnd"] = request.FlvSeekEnd
	}

	if !dara.IsNil(request.FlvSeekStart) {
		query["FlvSeekStart"] = request.FlvSeekStart
	}

	if !dara.IsNil(request.FlvVideoSeekMode) {
		query["FlvVideoSeekMode"] = request.FlvVideoSeekMode
	}

	if !dara.IsNil(request.Mp4SeekEnd) {
		query["Mp4SeekEnd"] = request.Mp4SeekEnd
	}

	if !dara.IsNil(request.Mp4SeekStart) {
		query["Mp4SeekStart"] = request.Mp4SeekStart
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	if !dara.IsNil(request.VideoSeekEnable) {
		query["VideoSeekEnable"] = request.VideoSeekEnable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVideoProcessing"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVideoProcessingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add video processing configurations for a website.
//
// @param request - CreateVideoProcessingRequest
//
// @return CreateVideoProcessingResponse
func (client *Client) CreateVideoProcessing(request *CreateVideoProcessingRequest) (_result *CreateVideoProcessingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVideoProcessingResponse{}
	_body, _err := client.CreateVideoProcessingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create WAF Rule
//
// @param tmpReq - CreateWafRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWafRuleResponse
func (client *Client) CreateWafRuleWithOptions(tmpReq *CreateWafRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateWafRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateWafRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Config) {
		request.ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Config, dara.String("Config"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigShrink) {
		body["Config"] = request.ConfigShrink
	}

	if !dara.IsNil(request.Phase) {
		body["Phase"] = request.Phase
	}

	if !dara.IsNil(request.RulesetId) {
		body["RulesetId"] = request.RulesetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWafRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWafRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create WAF Rule
//
// @param request - CreateWafRuleRequest
//
// @return CreateWafRuleResponse
func (client *Client) CreateWafRule(request *CreateWafRuleRequest) (_result *CreateWafRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWafRuleResponse{}
	_body, _err := client.CreateWafRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create WAF Ruleset
//
// @param request - CreateWafRulesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWafRulesetResponse
func (client *Client) CreateWafRulesetWithOptions(request *CreateWafRulesetRequest, runtime *dara.RuntimeOptions) (_result *CreateWafRulesetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Phase) {
		body["Phase"] = request.Phase
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWafRuleset"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWafRulesetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create WAF Ruleset
//
// @param request - CreateWafRulesetRequest
//
// @return CreateWafRulesetResponse
func (client *Client) CreateWafRuleset(request *CreateWafRulesetRequest) (_result *CreateWafRulesetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWafRulesetResponse{}
	_body, _err := client.CreateWafRulesetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a waiting room for a website.
//
// @param tmpReq - CreateWaitingRoomRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWaitingRoomResponse
func (client *Client) CreateWaitingRoomWithOptions(tmpReq *CreateWaitingRoomRequest, runtime *dara.RuntimeOptions) (_result *CreateWaitingRoomResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateWaitingRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HostNameAndPath) {
		request.HostNameAndPathShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HostNameAndPath, dara.String("HostNameAndPath"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CookieName) {
		query["CookieName"] = request.CookieName
	}

	if !dara.IsNil(request.CustomPageHtml) {
		query["CustomPageHtml"] = request.CustomPageHtml
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisableSessionRenewalEnable) {
		query["DisableSessionRenewalEnable"] = request.DisableSessionRenewalEnable
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.HostNameAndPathShrink) {
		query["HostNameAndPath"] = request.HostNameAndPathShrink
	}

	if !dara.IsNil(request.JsonResponseEnable) {
		query["JsonResponseEnable"] = request.JsonResponseEnable
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NewUsersPerMinute) {
		query["NewUsersPerMinute"] = request.NewUsersPerMinute
	}

	if !dara.IsNil(request.QueueAllEnable) {
		query["QueueAllEnable"] = request.QueueAllEnable
	}

	if !dara.IsNil(request.QueuingMethod) {
		query["QueuingMethod"] = request.QueuingMethod
	}

	if !dara.IsNil(request.QueuingStatusCode) {
		query["QueuingStatusCode"] = request.QueuingStatusCode
	}

	if !dara.IsNil(request.SessionDuration) {
		query["SessionDuration"] = request.SessionDuration
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.TotalActiveUsers) {
		query["TotalActiveUsers"] = request.TotalActiveUsers
	}

	if !dara.IsNil(request.WaitingRoomType) {
		query["WaitingRoomType"] = request.WaitingRoomType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWaitingRoom"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWaitingRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a waiting room for a website.
//
// @param request - CreateWaitingRoomRequest
//
// @return CreateWaitingRoomResponse
func (client *Client) CreateWaitingRoom(request *CreateWaitingRoomRequest) (_result *CreateWaitingRoomResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWaitingRoomResponse{}
	_body, _err := client.CreateWaitingRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a waiting room event.
//
// @param request - CreateWaitingRoomEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWaitingRoomEventResponse
func (client *Client) CreateWaitingRoomEventWithOptions(request *CreateWaitingRoomEventRequest, runtime *dara.RuntimeOptions) (_result *CreateWaitingRoomEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomPageHtml) {
		query["CustomPageHtml"] = request.CustomPageHtml
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisableSessionRenewalEnable) {
		query["DisableSessionRenewalEnable"] = request.DisableSessionRenewalEnable
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.JsonResponseEnable) {
		query["JsonResponseEnable"] = request.JsonResponseEnable
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NewUsersPerMinute) {
		query["NewUsersPerMinute"] = request.NewUsersPerMinute
	}

	if !dara.IsNil(request.PreQueueEnable) {
		query["PreQueueEnable"] = request.PreQueueEnable
	}

	if !dara.IsNil(request.PreQueueStartTime) {
		query["PreQueueStartTime"] = request.PreQueueStartTime
	}

	if !dara.IsNil(request.QueuingMethod) {
		query["QueuingMethod"] = request.QueuingMethod
	}

	if !dara.IsNil(request.QueuingStatusCode) {
		query["QueuingStatusCode"] = request.QueuingStatusCode
	}

	if !dara.IsNil(request.RandomPreQueueEnable) {
		query["RandomPreQueueEnable"] = request.RandomPreQueueEnable
	}

	if !dara.IsNil(request.SessionDuration) {
		query["SessionDuration"] = request.SessionDuration
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TotalActiveUsers) {
		query["TotalActiveUsers"] = request.TotalActiveUsers
	}

	if !dara.IsNil(request.WaitingRoomId) {
		query["WaitingRoomId"] = request.WaitingRoomId
	}

	if !dara.IsNil(request.WaitingRoomType) {
		query["WaitingRoomType"] = request.WaitingRoomType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWaitingRoomEvent"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWaitingRoomEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a waiting room event.
//
// @param request - CreateWaitingRoomEventRequest
//
// @return CreateWaitingRoomEventResponse
func (client *Client) CreateWaitingRoomEvent(request *CreateWaitingRoomEventRequest) (_result *CreateWaitingRoomEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWaitingRoomEventResponse{}
	_body, _err := client.CreateWaitingRoomEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Waiting Room Rule
//
// @param request - CreateWaitingRoomRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWaitingRoomRuleResponse
func (client *Client) CreateWaitingRoomRuleWithOptions(request *CreateWaitingRoomRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateWaitingRoomRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.WaitingRoomId) {
		query["WaitingRoomId"] = request.WaitingRoomId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWaitingRoomRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWaitingRoomRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Waiting Room Rule
//
// @param request - CreateWaitingRoomRuleRequest
//
// @return CreateWaitingRoomRuleResponse
func (client *Client) CreateWaitingRoomRule(request *CreateWaitingRoomRuleRequest) (_result *CreateWaitingRoomRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWaitingRoomRuleResponse{}
	_body, _err := client.CreateWaitingRoomRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables version management for a website.
//
// Description:
//
// You can disable version management only when the default environment and version 0 exist.
//
// @param request - DeactivateVersionManagementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeactivateVersionManagementResponse
func (client *Client) DeactivateVersionManagementWithOptions(request *DeactivateVersionManagementRequest, runtime *dara.RuntimeOptions) (_result *DeactivateVersionManagementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeactivateVersionManagement"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeactivateVersionManagementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables version management for a website.
//
// Description:
//
// You can disable version management only when the default environment and version 0 exist.
//
// @param request - DeactivateVersionManagementRequest
//
// @return DeactivateVersionManagementResponse
func (client *Client) DeactivateVersionManagement(request *DeactivateVersionManagementRequest) (_result *DeactivateVersionManagementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeactivateVersionManagementResponse{}
	_body, _err := client.DeactivateVersionManagementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Cache Configuration
//
// @param request - DeleteCacheRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCacheRuleResponse
func (client *Client) DeleteCacheRuleWithOptions(request *DeleteCacheRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteCacheRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCacheRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCacheRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Cache Configuration
//
// @param request - DeleteCacheRuleRequest
//
// @return DeleteCacheRuleResponse
func (client *Client) DeleteCacheRule(request *DeleteCacheRuleRequest) (_result *DeleteCacheRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCacheRuleResponse{}
	_body, _err := client.DeleteCacheRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a certificate for a website.
//
// @param request - DeleteCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCertificateResponse
func (client *Client) DeleteCertificateWithOptions(request *DeleteCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteCertificateResponse, _err error) {
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
		Action:      dara.String("DeleteCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a certificate for a website.
//
// @param request - DeleteCertificateRequest
//
// @return DeleteCertificateResponse
func (client *Client) DeleteCertificate(request *DeleteCertificateRequest) (_result *DeleteCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCertificateResponse{}
	_body, _err := client.DeleteCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a client CA certificate.
//
// @param request - DeleteClientCaCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClientCaCertificateResponse
func (client *Client) DeleteClientCaCertificateWithOptions(request *DeleteClientCaCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteClientCaCertificateResponse, _err error) {
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
		Action:      dara.String("DeleteClientCaCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClientCaCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a client CA certificate.
//
// @param request - DeleteClientCaCertificateRequest
//
// @return DeleteClientCaCertificateResponse
func (client *Client) DeleteClientCaCertificate(request *DeleteClientCaCertificateRequest) (_result *DeleteClientCaCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteClientCaCertificateResponse{}
	_body, _err := client.DeleteClientCaCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a revoked client certificate.
//
// @param request - DeleteClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClientCertificateResponse
func (client *Client) DeleteClientCertificateWithOptions(request *DeleteClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteClientCertificateResponse, _err error) {
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
		Action:      dara.String("DeleteClientCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a revoked client certificate.
//
// @param request - DeleteClientCertificateRequest
//
// @return DeleteClientCertificateResponse
func (client *Client) DeleteClientCertificate(request *DeleteClientCertificateRequest) (_result *DeleteClientCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteClientCertificateResponse{}
	_body, _err := client.DeleteClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete compression rule
//
// @param request - DeleteCompressionRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCompressionRuleResponse
func (client *Client) DeleteCompressionRuleWithOptions(request *DeleteCompressionRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteCompressionRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCompressionRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCompressionRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete compression rule
//
// @param request - DeleteCompressionRuleRequest
//
// @return DeleteCompressionRuleResponse
func (client *Client) DeleteCompressionRule(request *DeleteCompressionRuleRequest) (_result *DeleteCompressionRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCompressionRuleResponse{}
	_body, _err := client.DeleteCompressionRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a scenario-specific custom policy.
//
// @param request - DeleteCustomScenePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomScenePolicyResponse
func (client *Client) DeleteCustomScenePolicyWithOptions(request *DeleteCustomScenePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomScenePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomScenePolicy"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomScenePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scenario-specific custom policy.
//
// @param request - DeleteCustomScenePolicyRequest
//
// @return DeleteCustomScenePolicyResponse
func (client *Client) DeleteCustomScenePolicy(request *DeleteCustomScenePolicyRequest) (_result *DeleteCustomScenePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomScenePolicyResponse{}
	_body, _err := client.DeleteCustomScenePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a containerized application.
//
// @param request - DeleteEdgeContainerAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEdgeContainerAppResponse
func (client *Client) DeleteEdgeContainerAppWithOptions(request *DeleteEdgeContainerAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteEdgeContainerAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEdgeContainerApp"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEdgeContainerAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a containerized application.
//
// @param request - DeleteEdgeContainerAppRequest
//
// @return DeleteEdgeContainerAppResponse
func (client *Client) DeleteEdgeContainerApp(request *DeleteEdgeContainerAppRequest) (_result *DeleteEdgeContainerAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEdgeContainerAppResponse{}
	_body, _err := client.DeleteEdgeContainerAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates a domain name from a containerized application. After the dissociation, you can no longer use the domain name to access the containerized application.
//
// @param request - DeleteEdgeContainerAppRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEdgeContainerAppRecordResponse
func (client *Client) DeleteEdgeContainerAppRecordWithOptions(request *DeleteEdgeContainerAppRecordRequest, runtime *dara.RuntimeOptions) (_result *DeleteEdgeContainerAppRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.RecordName) {
		body["RecordName"] = request.RecordName
	}

	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEdgeContainerAppRecord"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEdgeContainerAppRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a domain name from a containerized application. After the dissociation, you can no longer use the domain name to access the containerized application.
//
// @param request - DeleteEdgeContainerAppRecordRequest
//
// @return DeleteEdgeContainerAppRecordResponse
func (client *Client) DeleteEdgeContainerAppRecord(request *DeleteEdgeContainerAppRecordRequest) (_result *DeleteEdgeContainerAppRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEdgeContainerAppRecordResponse{}
	_body, _err := client.DeleteEdgeContainerAppRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a version of a containerized application.
//
// @param request - DeleteEdgeContainerAppVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEdgeContainerAppVersionResponse
func (client *Client) DeleteEdgeContainerAppVersionWithOptions(request *DeleteEdgeContainerAppVersionRequest, runtime *dara.RuntimeOptions) (_result *DeleteEdgeContainerAppVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEdgeContainerAppVersion"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEdgeContainerAppVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a version of a containerized application.
//
// @param request - DeleteEdgeContainerAppVersionRequest
//
// @return DeleteEdgeContainerAppVersionResponse
func (client *Client) DeleteEdgeContainerAppVersion(request *DeleteEdgeContainerAppVersionRequest) (_result *DeleteEdgeContainerAppVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEdgeContainerAppVersionResponse{}
	_body, _err := client.DeleteEdgeContainerAppVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the configuration of modifying incoming HTTP request headers for a website.
//
// @param request - DeleteHttpIncomingRequestHeaderModificationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHttpIncomingRequestHeaderModificationRuleResponse
func (client *Client) DeleteHttpIncomingRequestHeaderModificationRuleWithOptions(request *DeleteHttpIncomingRequestHeaderModificationRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteHttpIncomingRequestHeaderModificationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHttpIncomingRequestHeaderModificationRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHttpIncomingRequestHeaderModificationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the configuration of modifying incoming HTTP request headers for a website.
//
// @param request - DeleteHttpIncomingRequestHeaderModificationRuleRequest
//
// @return DeleteHttpIncomingRequestHeaderModificationRuleResponse
func (client *Client) DeleteHttpIncomingRequestHeaderModificationRule(request *DeleteHttpIncomingRequestHeaderModificationRuleRequest) (_result *DeleteHttpIncomingRequestHeaderModificationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHttpIncomingRequestHeaderModificationRuleResponse{}
	_body, _err := client.DeleteHttpIncomingRequestHeaderModificationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the configuration of modifying HTTP response headers for a website.
//
// @param request - DeleteHttpIncomingResponseHeaderModificationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHttpIncomingResponseHeaderModificationRuleResponse
func (client *Client) DeleteHttpIncomingResponseHeaderModificationRuleWithOptions(request *DeleteHttpIncomingResponseHeaderModificationRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteHttpIncomingResponseHeaderModificationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHttpIncomingResponseHeaderModificationRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHttpIncomingResponseHeaderModificationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the configuration of modifying HTTP response headers for a website.
//
// @param request - DeleteHttpIncomingResponseHeaderModificationRuleRequest
//
// @return DeleteHttpIncomingResponseHeaderModificationRuleResponse
func (client *Client) DeleteHttpIncomingResponseHeaderModificationRule(request *DeleteHttpIncomingResponseHeaderModificationRuleRequest) (_result *DeleteHttpIncomingResponseHeaderModificationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHttpIncomingResponseHeaderModificationRuleResponse{}
	_body, _err := client.DeleteHttpIncomingResponseHeaderModificationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the configuration of modifying HTTP request headers for a website.
//
// @param request - DeleteHttpRequestHeaderModificationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHttpRequestHeaderModificationRuleResponse
func (client *Client) DeleteHttpRequestHeaderModificationRuleWithOptions(request *DeleteHttpRequestHeaderModificationRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteHttpRequestHeaderModificationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHttpRequestHeaderModificationRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHttpRequestHeaderModificationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the configuration of modifying HTTP request headers for a website.
//
// @param request - DeleteHttpRequestHeaderModificationRuleRequest
//
// @return DeleteHttpRequestHeaderModificationRuleResponse
func (client *Client) DeleteHttpRequestHeaderModificationRule(request *DeleteHttpRequestHeaderModificationRuleRequest) (_result *DeleteHttpRequestHeaderModificationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHttpRequestHeaderModificationRuleResponse{}
	_body, _err := client.DeleteHttpRequestHeaderModificationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the configuration of modifying HTTP response headers for a website.
//
// @param request - DeleteHttpResponseHeaderModificationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHttpResponseHeaderModificationRuleResponse
func (client *Client) DeleteHttpResponseHeaderModificationRuleWithOptions(request *DeleteHttpResponseHeaderModificationRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteHttpResponseHeaderModificationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHttpResponseHeaderModificationRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHttpResponseHeaderModificationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the configuration of modifying HTTP response headers for a website.
//
// @param request - DeleteHttpResponseHeaderModificationRuleRequest
//
// @return DeleteHttpResponseHeaderModificationRuleResponse
func (client *Client) DeleteHttpResponseHeaderModificationRule(request *DeleteHttpResponseHeaderModificationRuleRequest) (_result *DeleteHttpResponseHeaderModificationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHttpResponseHeaderModificationRuleResponse{}
	_body, _err := client.DeleteHttpResponseHeaderModificationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete HTTPS Application Configuration
//
// @param request - DeleteHttpsApplicationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHttpsApplicationConfigurationResponse
func (client *Client) DeleteHttpsApplicationConfigurationWithOptions(request *DeleteHttpsApplicationConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DeleteHttpsApplicationConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHttpsApplicationConfiguration"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHttpsApplicationConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete HTTPS Application Configuration
//
// @param request - DeleteHttpsApplicationConfigurationRequest
//
// @return DeleteHttpsApplicationConfigurationResponse
func (client *Client) DeleteHttpsApplicationConfiguration(request *DeleteHttpsApplicationConfigurationRequest) (_result *DeleteHttpsApplicationConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHttpsApplicationConfigurationResponse{}
	_body, _err := client.DeleteHttpsApplicationConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete HTTPS Basic Configuration
//
// @param request - DeleteHttpsBasicConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHttpsBasicConfigurationResponse
func (client *Client) DeleteHttpsBasicConfigurationWithOptions(request *DeleteHttpsBasicConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DeleteHttpsBasicConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHttpsBasicConfiguration"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHttpsBasicConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete HTTPS Basic Configuration
//
// @param request - DeleteHttpsBasicConfigurationRequest
//
// @return DeleteHttpsBasicConfigurationResponse
func (client *Client) DeleteHttpsBasicConfiguration(request *DeleteHttpsBasicConfigurationRequest) (_result *DeleteHttpsBasicConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHttpsBasicConfigurationResponse{}
	_body, _err := client.DeleteHttpsBasicConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Site Image Transformation Configuration
//
// @param request - DeleteImageTransformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteImageTransformResponse
func (client *Client) DeleteImageTransformWithOptions(request *DeleteImageTransformRequest, runtime *dara.RuntimeOptions) (_result *DeleteImageTransformResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteImageTransform"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteImageTransformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Site Image Transformation Configuration
//
// @param request - DeleteImageTransformRequest
//
// @return DeleteImageTransformResponse
func (client *Client) DeleteImageTransform(request *DeleteImageTransformRequest) (_result *DeleteImageTransformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteImageTransformResponse{}
	_body, _err := client.DeleteImageTransformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a key-value pair from a namespace.
//
// @param request - DeleteKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKvResponse
func (client *Client) DeleteKvWithOptions(request *DeleteKvRequest, runtime *dara.RuntimeOptions) (_result *DeleteKvResponse, _err error) {
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
		Action:      dara.String("DeleteKv"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteKvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a key-value pair from a namespace.
//
// @param request - DeleteKvRequest
//
// @return DeleteKvResponse
func (client *Client) DeleteKv(request *DeleteKvRequest) (_result *DeleteKvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteKvResponse{}
	_body, _err := client.DeleteKvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a namespace from an Alibaba Cloud account.
//
// @param request - DeleteKvNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKvNamespaceResponse
func (client *Client) DeleteKvNamespaceWithOptions(request *DeleteKvNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteKvNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteKvNamespace"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteKvNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a namespace from an Alibaba Cloud account.
//
// @param request - DeleteKvNamespaceRequest
//
// @return DeleteKvNamespaceResponse
func (client *Client) DeleteKvNamespace(request *DeleteKvNamespaceRequest) (_result *DeleteKvNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteKvNamespaceResponse{}
	_body, _err := client.DeleteKvNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom list that is no longer needed.
//
// @param request - DeleteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteListResponse
func (client *Client) DeleteListWithOptions(request *DeleteListRequest, runtime *dara.RuntimeOptions) (_result *DeleteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteList"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom list that is no longer needed.
//
// @param request - DeleteListRequest
//
// @return DeleteListResponse
func (client *Client) DeleteList(request *DeleteListRequest) (_result *DeleteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteListResponse{}
	_body, _err := client.DeleteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Load Balancer
//
// Description:
//
// Delete a load balancer by its ID, only one can be deleted at a time.
//
// @param request - DeleteLoadBalancerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLoadBalancerResponse
func (client *Client) DeleteLoadBalancerWithOptions(request *DeleteLoadBalancerRequest, runtime *dara.RuntimeOptions) (_result *DeleteLoadBalancerResponse, _err error) {
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

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLoadBalancer"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Load Balancer
//
// Description:
//
// Delete a load balancer by its ID, only one can be deleted at a time.
//
// @param request - DeleteLoadBalancerRequest
//
// @return DeleteLoadBalancerResponse
func (client *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (_result *DeleteLoadBalancerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLoadBalancerResponse{}
	_body, _err := client.DeleteLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Network Optimization Configuration
//
// @param request - DeleteNetworkOptimizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkOptimizationResponse
func (client *Client) DeleteNetworkOptimizationWithOptions(request *DeleteNetworkOptimizationRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkOptimizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNetworkOptimization"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetworkOptimizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Network Optimization Configuration
//
// @param request - DeleteNetworkOptimizationRequest
//
// @return DeleteNetworkOptimizationResponse
func (client *Client) DeleteNetworkOptimization(request *DeleteNetworkOptimizationRequest) (_result *DeleteNetworkOptimizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNetworkOptimizationResponse{}
	_body, _err := client.DeleteNetworkOptimizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除源服务器CA证书
//
// @param request - DeleteOriginCaCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOriginCaCertificateResponse
func (client *Client) DeleteOriginCaCertificateWithOptions(request *DeleteOriginCaCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteOriginCaCertificateResponse, _err error) {
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
		Action:      dara.String("DeleteOriginCaCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOriginCaCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除源服务器CA证书
//
// @param request - DeleteOriginCaCertificateRequest
//
// @return DeleteOriginCaCertificateResponse
func (client *Client) DeleteOriginCaCertificate(request *DeleteOriginCaCertificateRequest) (_result *DeleteOriginCaCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteOriginCaCertificateResponse{}
	_body, _err := client.DeleteOriginCaCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除域名回源客户端证书
//
// @param request - DeleteOriginClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOriginClientCertificateResponse
func (client *Client) DeleteOriginClientCertificateWithOptions(request *DeleteOriginClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteOriginClientCertificateResponse, _err error) {
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
		Action:      dara.String("DeleteOriginClientCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOriginClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除域名回源客户端证书
//
// @param request - DeleteOriginClientCertificateRequest
//
// @return DeleteOriginClientCertificateResponse
func (client *Client) DeleteOriginClientCertificate(request *DeleteOriginClientCertificateRequest) (_result *DeleteOriginClientCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteOriginClientCertificateResponse{}
	_body, _err := client.DeleteOriginClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Origin Address Pool
//
// @param request - DeleteOriginPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOriginPoolResponse
func (client *Client) DeleteOriginPoolWithOptions(request *DeleteOriginPoolRequest, runtime *dara.RuntimeOptions) (_result *DeleteOriginPoolResponse, _err error) {
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

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOriginPool"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOriginPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Origin Address Pool
//
// @param request - DeleteOriginPoolRequest
//
// @return DeleteOriginPoolResponse
func (client *Client) DeleteOriginPool(request *DeleteOriginPoolRequest) (_result *DeleteOriginPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteOriginPoolResponse{}
	_body, _err := client.DeleteOriginPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables origin protection.
//
// @param request - DeleteOriginProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOriginProtectionResponse
func (client *Client) DeleteOriginProtectionWithOptions(request *DeleteOriginProtectionRequest, runtime *dara.RuntimeOptions) (_result *DeleteOriginProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOriginProtection"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOriginProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables origin protection.
//
// @param request - DeleteOriginProtectionRequest
//
// @return DeleteOriginProtectionResponse
func (client *Client) DeleteOriginProtection(request *DeleteOriginProtectionRequest) (_result *DeleteOriginProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteOriginProtectionResponse{}
	_body, _err := client.DeleteOriginProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Origin Rule Configuration
//
// @param request - DeleteOriginRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOriginRuleResponse
func (client *Client) DeleteOriginRuleWithOptions(request *DeleteOriginRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteOriginRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOriginRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOriginRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Origin Rule Configuration
//
// @param request - DeleteOriginRuleRequest
//
// @return DeleteOriginRuleResponse
func (client *Client) DeleteOriginRule(request *DeleteOriginRuleRequest) (_result *DeleteOriginRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteOriginRuleResponse{}
	_body, _err := client.DeleteOriginRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom error page that is no longer needed.
//
// @param request - DeletePageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePageResponse
func (client *Client) DeletePageWithOptions(request *DeletePageRequest, runtime *dara.RuntimeOptions) (_result *DeletePageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePage"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom error page that is no longer needed.
//
// @param request - DeletePageRequest
//
// @return DeletePageResponse
func (client *Client) DeletePage(request *DeletePageRequest) (_result *DeletePageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePageResponse{}
	_body, _err := client.DeletePageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a DNS record of a website based on the specified RecordId.
//
// @param request - DeleteRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecordResponse
func (client *Client) DeleteRecordWithOptions(request *DeleteRecordRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecord"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a DNS record of a website based on the specified RecordId.
//
// @param request - DeleteRecordRequest
//
// @return DeleteRecordResponse
func (client *Client) DeleteRecord(request *DeleteRecordRequest) (_result *DeleteRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRecordResponse{}
	_body, _err := client.DeleteRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a URL redirect rule for a website.
//
// @param request - DeleteRedirectRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRedirectRuleResponse
func (client *Client) DeleteRedirectRuleWithOptions(request *DeleteRedirectRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRedirectRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRedirectRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRedirectRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a URL redirect rule for a website.
//
// @param request - DeleteRedirectRuleRequest
//
// @return DeleteRedirectRuleResponse
func (client *Client) DeleteRedirectRule(request *DeleteRedirectRuleRequest) (_result *DeleteRedirectRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRedirectRuleResponse{}
	_body, _err := client.DeleteRedirectRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a URL rewrite rule for a website.
//
// @param request - DeleteRewriteUrlRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRewriteUrlRuleResponse
func (client *Client) DeleteRewriteUrlRuleWithOptions(request *DeleteRewriteUrlRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRewriteUrlRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRewriteUrlRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRewriteUrlRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a URL rewrite rule for a website.
//
// @param request - DeleteRewriteUrlRuleRequest
//
// @return DeleteRewriteUrlRuleResponse
func (client *Client) DeleteRewriteUrlRule(request *DeleteRewriteUrlRuleRequest) (_result *DeleteRewriteUrlRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRewriteUrlRuleResponse{}
	_body, _err := client.DeleteRewriteUrlRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a routine in Edge Routine.
//
// @param request - DeleteRoutineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoutineResponse
func (client *Client) DeleteRoutineWithOptions(request *DeleteRoutineRequest, runtime *dara.RuntimeOptions) (_result *DeleteRoutineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRoutine"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRoutineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a routine in Edge Routine.
//
// @param request - DeleteRoutineRequest
//
// @return DeleteRoutineResponse
func (client *Client) DeleteRoutine(request *DeleteRoutineRequest) (_result *DeleteRoutineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRoutineResponse{}
	_body, _err := client.DeleteRoutineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a code version of a routine.
//
// @param request - DeleteRoutineCodeVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoutineCodeVersionResponse
func (client *Client) DeleteRoutineCodeVersionWithOptions(request *DeleteRoutineCodeVersionRequest, runtime *dara.RuntimeOptions) (_result *DeleteRoutineCodeVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CodeVersion) {
		body["CodeVersion"] = request.CodeVersion
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRoutineCodeVersion"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRoutineCodeVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a code version of a routine.
//
// @param request - DeleteRoutineCodeVersionRequest
//
// @return DeleteRoutineCodeVersionResponse
func (client *Client) DeleteRoutineCodeVersion(request *DeleteRoutineCodeVersionRequest) (_result *DeleteRoutineCodeVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRoutineCodeVersionResponse{}
	_body, _err := client.DeleteRoutineCodeVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a record that is associated with a routine.
//
// @param request - DeleteRoutineRelatedRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoutineRelatedRecordResponse
func (client *Client) DeleteRoutineRelatedRecordWithOptions(request *DeleteRoutineRelatedRecordRequest, runtime *dara.RuntimeOptions) (_result *DeleteRoutineRelatedRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RecordId) {
		body["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.RecordName) {
		body["RecordName"] = request.RecordName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRoutineRelatedRecord"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRoutineRelatedRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a record that is associated with a routine.
//
// @param request - DeleteRoutineRelatedRecordRequest
//
// @return DeleteRoutineRelatedRecordResponse
func (client *Client) DeleteRoutineRelatedRecord(request *DeleteRoutineRelatedRecordRequest) (_result *DeleteRoutineRelatedRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRoutineRelatedRecordResponse{}
	_body, _err := client.DeleteRoutineRelatedRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the route configuration of an edge function.
//
// @param request - DeleteRoutineRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoutineRouteResponse
func (client *Client) DeleteRoutineRouteWithOptions(request *DeleteRoutineRouteRequest, runtime *dara.RuntimeOptions) (_result *DeleteRoutineRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRoutineRoute"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRoutineRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the route configuration of an edge function.
//
// @param request - DeleteRoutineRouteRequest
//
// @return DeleteRoutineRouteResponse
func (client *Client) DeleteRoutineRoute(request *DeleteRoutineRouteRequest) (_result *DeleteRoutineRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRoutineRouteResponse{}
	_body, _err := client.DeleteRoutineRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a scheduled prefetch plan based on the plan ID.
//
// @param request - DeleteScheduledPreloadExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScheduledPreloadExecutionResponse
func (client *Client) DeleteScheduledPreloadExecutionWithOptions(request *DeleteScheduledPreloadExecutionRequest, runtime *dara.RuntimeOptions) (_result *DeleteScheduledPreloadExecutionResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScheduledPreloadExecution"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScheduledPreloadExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scheduled prefetch plan based on the plan ID.
//
// @param request - DeleteScheduledPreloadExecutionRequest
//
// @return DeleteScheduledPreloadExecutionResponse
func (client *Client) DeleteScheduledPreloadExecution(request *DeleteScheduledPreloadExecutionRequest) (_result *DeleteScheduledPreloadExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteScheduledPreloadExecutionResponse{}
	_body, _err := client.DeleteScheduledPreloadExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified scheduled prefetch task based on the task ID.
//
// @param request - DeleteScheduledPreloadJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScheduledPreloadJobResponse
func (client *Client) DeleteScheduledPreloadJobWithOptions(request *DeleteScheduledPreloadJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteScheduledPreloadJobResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScheduledPreloadJob"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScheduledPreloadJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified scheduled prefetch task based on the task ID.
//
// @param request - DeleteScheduledPreloadJobRequest
//
// @return DeleteScheduledPreloadJobResponse
func (client *Client) DeleteScheduledPreloadJob(request *DeleteScheduledPreloadJobRequest) (_result *DeleteScheduledPreloadJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteScheduledPreloadJobResponse{}
	_body, _err := client.DeleteScheduledPreloadJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a website based on the specified website ID.
//
// @param request - DeleteSiteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSiteResponse
func (client *Client) DeleteSiteWithOptions(request *DeleteSiteRequest, runtime *dara.RuntimeOptions) (_result *DeleteSiteResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSite"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSiteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a website based on the specified website ID.
//
// @param request - DeleteSiteRequest
//
// @return DeleteSiteResponse
func (client *Client) DeleteSite(request *DeleteSiteRequest) (_result *DeleteSiteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSiteResponse{}
	_body, _err := client.DeleteSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a real-time log delivery task.
//
// @param request - DeleteSiteDeliveryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSiteDeliveryTaskResponse
func (client *Client) DeleteSiteDeliveryTaskWithOptions(request *DeleteSiteDeliveryTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteSiteDeliveryTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSiteDeliveryTask"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSiteDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a real-time log delivery task.
//
// @param request - DeleteSiteDeliveryTaskRequest
//
// @return DeleteSiteDeliveryTaskResponse
func (client *Client) DeleteSiteDeliveryTask(request *DeleteSiteDeliveryTaskRequest) (_result *DeleteSiteDeliveryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSiteDeliveryTaskResponse{}
	_body, _err := client.DeleteSiteDeliveryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除站点回源客户端证书
//
// @param request - DeleteSiteOriginClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSiteOriginClientCertificateResponse
func (client *Client) DeleteSiteOriginClientCertificateWithOptions(request *DeleteSiteOriginClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteSiteOriginClientCertificateResponse, _err error) {
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
		Action:      dara.String("DeleteSiteOriginClientCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSiteOriginClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除站点回源客户端证书
//
// @param request - DeleteSiteOriginClientCertificateRequest
//
// @return DeleteSiteOriginClientCertificateResponse
func (client *Client) DeleteSiteOriginClientCertificate(request *DeleteSiteOriginClientCertificateRequest) (_result *DeleteSiteOriginClientCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSiteOriginClientCertificateResponse{}
	_body, _err := client.DeleteSiteOriginClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Transport Layer Application
//
// @param request - DeleteTransportLayerApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransportLayerApplicationResponse
func (client *Client) DeleteTransportLayerApplicationWithOptions(request *DeleteTransportLayerApplicationRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransportLayerApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTransportLayerApplication"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTransportLayerApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Transport Layer Application
//
// @param request - DeleteTransportLayerApplicationRequest
//
// @return DeleteTransportLayerApplicationResponse
func (client *Client) DeleteTransportLayerApplication(request *DeleteTransportLayerApplicationRequest) (_result *DeleteTransportLayerApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTransportLayerApplicationResponse{}
	_body, _err := client.DeleteTransportLayerApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes page monitoring configurations.
//
// @param request - DeleteUrlObservationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUrlObservationResponse
func (client *Client) DeleteUrlObservationWithOptions(request *DeleteUrlObservationRequest, runtime *dara.RuntimeOptions) (_result *DeleteUrlObservationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUrlObservation"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUrlObservationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes page monitoring configurations.
//
// @param request - DeleteUrlObservationRequest
//
// @return DeleteUrlObservationResponse
func (client *Client) DeleteUrlObservation(request *DeleteUrlObservationRequest) (_result *DeleteUrlObservationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUrlObservationResponse{}
	_body, _err := client.DeleteUrlObservationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a log delivery task from your Alibaba Cloud account.
//
// Description:
//
// *****>
//
//   - Deleted tasks cannot be restored. Proceed with caution.
//
//   - To call this operation, you must have an account that has the required permissions.
//
//   - The returned `RequestId` value can be used to track the request processing progress and troubleshoot issues.
//
// @param request - DeleteUserDeliveryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserDeliveryTaskResponse
func (client *Client) DeleteUserDeliveryTaskWithOptions(request *DeleteUserDeliveryTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserDeliveryTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserDeliveryTask"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a log delivery task from your Alibaba Cloud account.
//
// Description:
//
// *****>
//
//   - Deleted tasks cannot be restored. Proceed with caution.
//
//   - To call this operation, you must have an account that has the required permissions.
//
//   - The returned `RequestId` value can be used to track the request processing progress and troubleshoot issues.
//
// @param request - DeleteUserDeliveryTaskRequest
//
// @return DeleteUserDeliveryTaskResponse
func (client *Client) DeleteUserDeliveryTask(request *DeleteUserDeliveryTaskRequest) (_result *DeleteUserDeliveryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserDeliveryTaskResponse{}
	_body, _err := client.DeleteUserDeliveryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a video processing configuration.
//
// @param request - DeleteVideoProcessingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVideoProcessingResponse
func (client *Client) DeleteVideoProcessingWithOptions(request *DeleteVideoProcessingRequest, runtime *dara.RuntimeOptions) (_result *DeleteVideoProcessingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVideoProcessing"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVideoProcessingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a video processing configuration.
//
// @param request - DeleteVideoProcessingRequest
//
// @return DeleteVideoProcessingResponse
func (client *Client) DeleteVideoProcessing(request *DeleteVideoProcessingRequest) (_result *DeleteVideoProcessingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVideoProcessingResponse{}
	_body, _err := client.DeleteVideoProcessingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete WAF Rule
//
// @param request - DeleteWafRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWafRuleResponse
func (client *Client) DeleteWafRuleWithOptions(request *DeleteWafRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteWafRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWafRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWafRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete WAF Rule
//
// @param request - DeleteWafRuleRequest
//
// @return DeleteWafRuleResponse
func (client *Client) DeleteWafRule(request *DeleteWafRuleRequest) (_result *DeleteWafRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWafRuleResponse{}
	_body, _err := client.DeleteWafRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete WAF Ruleset
//
// @param request - DeleteWafRulesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWafRulesetResponse
func (client *Client) DeleteWafRulesetWithOptions(request *DeleteWafRulesetRequest, runtime *dara.RuntimeOptions) (_result *DeleteWafRulesetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWafRuleset"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWafRulesetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete WAF Ruleset
//
// @param request - DeleteWafRulesetRequest
//
// @return DeleteWafRulesetResponse
func (client *Client) DeleteWafRuleset(request *DeleteWafRulesetRequest) (_result *DeleteWafRulesetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWafRulesetResponse{}
	_body, _err := client.DeleteWafRulesetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a waiting room.
//
// @param request - DeleteWaitingRoomRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWaitingRoomResponse
func (client *Client) DeleteWaitingRoomWithOptions(request *DeleteWaitingRoomRequest, runtime *dara.RuntimeOptions) (_result *DeleteWaitingRoomResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.WaitingRoomId) {
		query["WaitingRoomId"] = request.WaitingRoomId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWaitingRoom"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWaitingRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a waiting room.
//
// @param request - DeleteWaitingRoomRequest
//
// @return DeleteWaitingRoomResponse
func (client *Client) DeleteWaitingRoom(request *DeleteWaitingRoomRequest) (_result *DeleteWaitingRoomResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWaitingRoomResponse{}
	_body, _err := client.DeleteWaitingRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a waiting room event.
//
// @param request - DeleteWaitingRoomEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWaitingRoomEventResponse
func (client *Client) DeleteWaitingRoomEventWithOptions(request *DeleteWaitingRoomEventRequest, runtime *dara.RuntimeOptions) (_result *DeleteWaitingRoomEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.WaitingRoomEventId) {
		query["WaitingRoomEventId"] = request.WaitingRoomEventId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWaitingRoomEvent"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWaitingRoomEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a waiting room event.
//
// @param request - DeleteWaitingRoomEventRequest
//
// @return DeleteWaitingRoomEventResponse
func (client *Client) DeleteWaitingRoomEvent(request *DeleteWaitingRoomEventRequest) (_result *DeleteWaitingRoomEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWaitingRoomEventResponse{}
	_body, _err := client.DeleteWaitingRoomEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a waiting room bypass rule.
//
// @param request - DeleteWaitingRoomRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWaitingRoomRuleResponse
func (client *Client) DeleteWaitingRoomRuleWithOptions(request *DeleteWaitingRoomRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteWaitingRoomRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.WaitingRoomRuleId) {
		query["WaitingRoomRuleId"] = request.WaitingRoomRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWaitingRoomRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWaitingRoomRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a waiting room bypass rule.
//
// @param request - DeleteWaitingRoomRuleRequest
//
// @return DeleteWaitingRoomRuleResponse
func (client *Client) DeleteWaitingRoomRule(request *DeleteWaitingRoomRuleRequest) (_result *DeleteWaitingRoomRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWaitingRoomRuleResponse{}
	_body, _err := client.DeleteWaitingRoomRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of a scenario-specific policy.
//
// @param request - DescribeCustomScenePoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomScenePoliciesResponse
func (client *Client) DescribeCustomScenePoliciesWithOptions(request *DescribeCustomScenePoliciesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomScenePoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomScenePolicies"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomScenePoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a scenario-specific policy.
//
// @param request - DescribeCustomScenePoliciesRequest
//
// @return DescribeCustomScenePoliciesResponse
func (client *Client) DescribeCustomScenePolicies(request *DescribeCustomScenePoliciesRequest) (_result *DescribeCustomScenePoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomScenePoliciesResponse{}
	_body, _err := client.DescribeCustomScenePoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries DDoS attack events.
//
// @param request - DescribeDDoSAllEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDDoSAllEventListResponse
func (client *Client) DescribeDDoSAllEventListWithOptions(request *DescribeDDoSAllEventListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDDoSAllEventListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDDoSAllEventList"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDDoSAllEventListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries DDoS attack events.
//
// @param request - DescribeDDoSAllEventListRequest
//
// @return DescribeDDoSAllEventListResponse
func (client *Client) DescribeDDoSAllEventList(request *DescribeDDoSAllEventListRequest) (_result *DescribeDDoSAllEventListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDDoSAllEventListResponse{}
	_body, _err := client.DescribeDDoSAllEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query DCDN DDoS user bps and pps data
//
// @param request - DescribeDDoSBpsListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDDoSBpsListResponse
func (client *Client) DescribeDDoSBpsListWithOptions(request *DescribeDDoSBpsListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDDoSBpsListResponse, _err error) {
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
		Action:      dara.String("DescribeDDoSBpsList"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDDoSBpsListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query DCDN DDoS user bps and pps data
//
// @param request - DescribeDDoSBpsListRequest
//
// @return DescribeDDoSBpsListResponse
func (client *Client) DescribeDDoSBpsList(request *DescribeDDoSBpsListRequest) (_result *DescribeDDoSBpsListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDDoSBpsListResponse{}
	_body, _err := client.DescribeDDoSBpsListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DDoS Analysis Layer 7 QPS Trend Chart API
//
// @param request - DescribeDDoSL7QpsListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDDoSL7QpsListResponse
func (client *Client) DescribeDDoSL7QpsListWithOptions(request *DescribeDDoSL7QpsListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDDoSL7QpsListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDDoSL7QpsList"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDDoSL7QpsListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DDoS Analysis Layer 7 QPS Trend Chart API
//
// @param request - DescribeDDoSL7QpsListRequest
//
// @return DescribeDDoSL7QpsListResponse
func (client *Client) DescribeDDoSL7QpsList(request *DescribeDDoSL7QpsListRequest) (_result *DescribeDDoSL7QpsListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDDoSL7QpsListResponse{}
	_body, _err := client.DescribeDDoSL7QpsListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Provides monitoring data for metrics of ESA edge containers.
//
// @param request - DescribeEdgeContainerAppStatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEdgeContainerAppStatsResponse
func (client *Client) DescribeEdgeContainerAppStatsWithOptions(request *DescribeEdgeContainerAppStatsRequest, runtime *dara.RuntimeOptions) (_result *DescribeEdgeContainerAppStatsResponse, _err error) {
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
		Action:      dara.String("DescribeEdgeContainerAppStats"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEdgeContainerAppStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides monitoring data for metrics of ESA edge containers.
//
// @param request - DescribeEdgeContainerAppStatsRequest
//
// @return DescribeEdgeContainerAppStatsResponse
func (client *Client) DescribeEdgeContainerAppStats(request *DescribeEdgeContainerAppStatsRequest) (_result *DescribeEdgeContainerAppStatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEdgeContainerAppStatsResponse{}
	_body, _err := client.DescribeEdgeContainerAppStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration of smart HTTP DDoS protection for a website.
//
// @param request - DescribeHttpDDoSAttackIntelligentProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHttpDDoSAttackIntelligentProtectionResponse
func (client *Client) DescribeHttpDDoSAttackIntelligentProtectionWithOptions(request *DescribeHttpDDoSAttackIntelligentProtectionRequest, runtime *dara.RuntimeOptions) (_result *DescribeHttpDDoSAttackIntelligentProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHttpDDoSAttackIntelligentProtection"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHttpDDoSAttackIntelligentProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of smart HTTP DDoS protection for a website.
//
// @param request - DescribeHttpDDoSAttackIntelligentProtectionRequest
//
// @return DescribeHttpDDoSAttackIntelligentProtectionResponse
func (client *Client) DescribeHttpDDoSAttackIntelligentProtection(request *DescribeHttpDDoSAttackIntelligentProtectionRequest) (_result *DescribeHttpDDoSAttackIntelligentProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHttpDDoSAttackIntelligentProtectionResponse{}
	_body, _err := client.DescribeHttpDDoSAttackIntelligentProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of HTTP DDoS attack protection.
//
// @param request - DescribeHttpDDoSAttackProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHttpDDoSAttackProtectionResponse
func (client *Client) DescribeHttpDDoSAttackProtectionWithOptions(request *DescribeHttpDDoSAttackProtectionRequest, runtime *dara.RuntimeOptions) (_result *DescribeHttpDDoSAttackProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHttpDDoSAttackProtection"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHttpDDoSAttackProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of HTTP DDoS attack protection.
//
// @param request - DescribeHttpDDoSAttackProtectionRequest
//
// @return DescribeHttpDDoSAttackProtectionResponse
func (client *Client) DescribeHttpDDoSAttackProtection(request *DescribeHttpDDoSAttackProtectionRequest) (_result *DescribeHttpDDoSAttackProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHttpDDoSAttackProtectionResponse{}
	_body, _err := client.DescribeHttpDDoSAttackProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether Edge KV is activated in your Alibaba Cloud account.
//
// @param request - DescribeKvAccountStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKvAccountStatusResponse
func (client *Client) DescribeKvAccountStatusWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeKvAccountStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeKvAccountStatus"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKvAccountStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether Edge KV is activated in your Alibaba Cloud account.
//
// @return DescribeKvAccountStatusResponse
func (client *Client) DescribeKvAccountStatus() (_result *DescribeKvAccountStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeKvAccountStatusResponse{}
	_body, _err := client.DescribeKvAccountStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of prefetch tasks by time, task status, or prefetch URL.
//
// @param request - DescribePreloadTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePreloadTasksResponse
func (client *Client) DescribePreloadTasksWithOptions(request *DescribePreloadTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribePreloadTasksResponse, _err error) {
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
		Action:      dara.String("DescribePreloadTasks"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePreloadTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of prefetch tasks by time, task status, or prefetch URL.
//
// @param request - DescribePreloadTasksRequest
//
// @return DescribePreloadTasksResponse
func (client *Client) DescribePreloadTasks(request *DescribePreloadTasksRequest) (_result *DescribePreloadTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePreloadTasksResponse{}
	_body, _err := client.DescribePreloadTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of purge tasks.
//
// @param request - DescribePurgeTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePurgeTasksResponse
func (client *Client) DescribePurgeTasksWithOptions(request *DescribePurgeTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribePurgeTasksResponse, _err error) {
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
		Action:      dara.String("DescribePurgeTasks"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePurgeTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of purge tasks.
//
// @param request - DescribePurgeTasksRequest
//
// @return DescribePurgeTasksResponse
func (client *Client) DescribePurgeTasks(request *DescribePurgeTasksRequest) (_result *DescribePurgeTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePurgeTasksResponse{}
	_body, _err := client.DescribePurgeTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of an instance that uses a plan.
//
// Description:
//
// You can query the status of an instance after you purchase a plan for the instance.
//
// @param request - DescribeRatePlanInstanceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRatePlanInstanceStatusResponse
func (client *Client) DescribeRatePlanInstanceStatusWithOptions(request *DescribeRatePlanInstanceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeRatePlanInstanceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRatePlanInstanceStatus"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRatePlanInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of an instance that uses a plan.
//
// Description:
//
// You can query the status of an instance after you purchase a plan for the instance.
//
// @param request - DescribeRatePlanInstanceStatusRequest
//
// @return DescribeRatePlanInstanceStatusResponse
func (client *Client) DescribeRatePlanInstanceStatus(request *DescribeRatePlanInstanceStatusRequest) (_result *DescribeRatePlanInstanceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRatePlanInstanceStatusResponse{}
	_body, _err := client.DescribeRatePlanInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the URLs from which you can download the raw access logs of a website.
//
// Description:
//
//	  If you do not specify StartTime or EndTime, the log data generated in the last 24 hours is queried. If you specify StartTime and EndTime, the log data generated within the specified time range is queried.
//
//		- The log data is collected every hour.
//
//		- You can call this operation up to 50 times per second per account.
//
//		- You can query only logs in the last month. The time range cannot exceed 31 days.
//
// @param request - DescribeSiteLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSiteLogsResponse
func (client *Client) DescribeSiteLogsWithOptions(request *DescribeSiteLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSiteLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSiteLogs"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSiteLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the URLs from which you can download the raw access logs of a website.
//
// Description:
//
//	  If you do not specify StartTime or EndTime, the log data generated in the last 24 hours is queried. If you specify StartTime and EndTime, the log data generated within the specified time range is queried.
//
//		- The log data is collected every hour.
//
//		- You can call this operation up to 50 times per second per account.
//
//		- You can query only logs in the last month. The time range cannot exceed 31 days.
//
// @param request - DescribeSiteLogsRequest
//
// @return DescribeSiteLogsResponse
func (client *Client) DescribeSiteLogs(request *DescribeSiteLogsRequest) (_result *DescribeSiteLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSiteLogsResponse{}
	_body, _err := client.DescribeSiteLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query traffic analysis time series data
//
// Description:
//
// - If you do not specify `StartTime` and `EndTime`, the API returns data for the past 24 hours; if you specify `StartTime` and `EndTime`, the API returns data for the specified time period.
//
// - The API returns different time granularities based on the span between `StartTime` and `EndTime`.
//
//   - For a span of 3 hours or less, it returns 1-minute granularity data.
//
//   - For a span greater than 3 hours but no more than 12 hours, it returns 5-minute granularity data.
//
//   - For a span greater than 12 hours but no more than 1 day, it returns 15-minute granularity data.
//
//   - For a span greater than 1 day but no more than 10 days, it returns hourly granularity data.
//
//   - For a span greater than 10 days but no more than 31 days, it returns daily granularity data.
//
// - Due to the high number of accesses during the query period, the data analysis may be sampled.
//
// @param tmpReq - DescribeSiteTimeSeriesDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSiteTimeSeriesDataResponse
func (client *Client) DescribeSiteTimeSeriesDataWithOptions(tmpReq *DescribeSiteTimeSeriesDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeSiteTimeSeriesDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeSiteTimeSeriesDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Fields) {
		request.FieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Fields, dara.String("Fields"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FieldsShrink) {
		query["Fields"] = request.FieldsShrink
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSiteTimeSeriesData"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSiteTimeSeriesDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query traffic analysis time series data
//
// Description:
//
// - If you do not specify `StartTime` and `EndTime`, the API returns data for the past 24 hours; if you specify `StartTime` and `EndTime`, the API returns data for the specified time period.
//
// - The API returns different time granularities based on the span between `StartTime` and `EndTime`.
//
//   - For a span of 3 hours or less, it returns 1-minute granularity data.
//
//   - For a span greater than 3 hours but no more than 12 hours, it returns 5-minute granularity data.
//
//   - For a span greater than 12 hours but no more than 1 day, it returns 15-minute granularity data.
//
//   - For a span greater than 1 day but no more than 10 days, it returns hourly granularity data.
//
//   - For a span greater than 10 days but no more than 31 days, it returns daily granularity data.
//
// - Due to the high number of accesses during the query period, the data analysis may be sampled.
//
// @param request - DescribeSiteTimeSeriesDataRequest
//
// @return DescribeSiteTimeSeriesDataResponse
func (client *Client) DescribeSiteTimeSeriesData(request *DescribeSiteTimeSeriesDataRequest) (_result *DescribeSiteTimeSeriesDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSiteTimeSeriesDataResponse{}
	_body, _err := client.DescribeSiteTimeSeriesDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top-ranking records in a traffic analytics report by website or Alibaba Cloud account.
//
// Description:
//
//	If you do not specify the StartTime or EndTime parameter, the request returns the data collected in the previous 24 hours. If you specify both parameters, the request returns the data collected within the specified time range.
//
// @param tmpReq - DescribeSiteTopDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSiteTopDataResponse
func (client *Client) DescribeSiteTopDataWithOptions(tmpReq *DescribeSiteTopDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeSiteTopDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeSiteTopDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Fields) {
		request.FieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Fields, dara.String("Fields"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FieldsShrink) {
		query["Fields"] = request.FieldsShrink
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSiteTopData"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSiteTopDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top-ranking records in a traffic analytics report by website or Alibaba Cloud account.
//
// Description:
//
//	If you do not specify the StartTime or EndTime parameter, the request returns the data collected in the previous 24 hours. If you specify both parameters, the request returns the data collected within the specified time range.
//
// @param request - DescribeSiteTopDataRequest
//
// @return DescribeSiteTopDataResponse
func (client *Client) DescribeSiteTopData(request *DescribeSiteTopDataRequest) (_result *DescribeSiteTopDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSiteTopDataResponse{}
	_body, _err := client.DescribeSiteTopDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the page monitoring data.
//
// Description:
//
// If you do not specify the StartTime or EndTime parameter, this operation returns the data collected within the last 24 hours. If you specify both parameters, this operation returns the data collected within the specified time range.
//
// @param request - DescribeUrlObservationDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUrlObservationDataResponse
func (client *Client) DescribeUrlObservationDataWithOptions(request *DescribeUrlObservationDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeUrlObservationDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientPlatform) {
		query["ClientPlatform"] = request.ClientPlatform
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Metric) {
		query["Metric"] = request.Metric
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUrlObservationData"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUrlObservationDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the page monitoring data.
//
// Description:
//
// If you do not specify the StartTime or EndTime parameter, this operation returns the data collected within the last 24 hours. If you specify both parameters, this operation returns the data collected within the specified time range.
//
// @param request - DescribeUrlObservationDataRequest
//
// @return DescribeUrlObservationDataResponse
func (client *Client) DescribeUrlObservationData(request *DescribeUrlObservationDataRequest) (_result *DescribeUrlObservationDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUrlObservationDataResponse{}
	_body, _err := client.DescribeUrlObservationDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a scenario-specific policy.
//
// @param request - DisableCustomScenePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableCustomScenePolicyResponse
func (client *Client) DisableCustomScenePolicyWithOptions(request *DisableCustomScenePolicyRequest, runtime *dara.RuntimeOptions) (_result *DisableCustomScenePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableCustomScenePolicy"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableCustomScenePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a scenario-specific policy.
//
// @param request - DisableCustomScenePolicyRequest
//
// @return DisableCustomScenePolicyResponse
func (client *Client) DisableCustomScenePolicy(request *DisableCustomScenePolicyRequest) (_result *DisableCustomScenePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableCustomScenePolicyResponse{}
	_body, _err := client.DisableCustomScenePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Edit WAF Configuration for a Site
//
// @param tmpReq - EditSiteWafSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditSiteWafSettingsResponse
func (client *Client) EditSiteWafSettingsWithOptions(tmpReq *EditSiteWafSettingsRequest, runtime *dara.RuntimeOptions) (_result *EditSiteWafSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &EditSiteWafSettingsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Settings) {
		request.SettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Settings, dara.String("Settings"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SettingsShrink) {
		body["Settings"] = request.SettingsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditSiteWafSettings"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditSiteWafSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Edit WAF Configuration for a Site
//
// @param request - EditSiteWafSettingsRequest
//
// @return EditSiteWafSettingsResponse
func (client *Client) EditSiteWafSettings(request *EditSiteWafSettingsRequest) (_result *EditSiteWafSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EditSiteWafSettingsResponse{}
	_body, _err := client.EditSiteWafSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a scenario-specific policy.
//
// @param request - EnableCustomScenePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableCustomScenePolicyResponse
func (client *Client) EnableCustomScenePolicyWithOptions(request *EnableCustomScenePolicyRequest, runtime *dara.RuntimeOptions) (_result *EnableCustomScenePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableCustomScenePolicy"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableCustomScenePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a scenario-specific policy.
//
// @param request - EnableCustomScenePolicyRequest
//
// @return EnableCustomScenePolicyResponse
func (client *Client) EnableCustomScenePolicy(request *EnableCustomScenePolicyRequest) (_result *EnableCustomScenePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableCustomScenePolicyResponse{}
	_body, _err := client.EnableCustomScenePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Exports all DNS records of a website domain as a TXT file.
//
// @param request - ExportRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportRecordsResponse
func (client *Client) ExportRecordsWithOptions(request *ExportRecordsRequest, runtime *dara.RuntimeOptions) (_result *ExportRecordsResponse, _err error) {
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
		Action:      dara.String("ExportRecords"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports all DNS records of a website domain as a TXT file.
//
// @param request - ExportRecordsRequest
//
// @return ExportRecordsResponse
func (client *Client) ExportRecords(request *ExportRecordsRequest) (_result *ExportRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportRecordsResponse{}
	_body, _err := client.ExportRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the usage of the upload file quota for API security schema verification.
//
// @param request - GetApiSchemaUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApiSchemaUsageResponse
func (client *Client) GetApiSchemaUsageWithOptions(request *GetApiSchemaUsageRequest, runtime *dara.RuntimeOptions) (_result *GetApiSchemaUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApiSchemaUsage"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApiSchemaUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage of the upload file quota for API security schema verification.
//
// @param request - GetApiSchemaUsageRequest
//
// @return GetApiSchemaUsageResponse
func (client *Client) GetApiSchemaUsage(request *GetApiSchemaUsageRequest) (_result *GetApiSchemaUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApiSchemaUsageResponse{}
	_body, _err := client.GetApiSchemaUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available specifications of cache reserve instances.
//
// @param request - GetCacheReserveSpecificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCacheReserveSpecificationResponse
func (client *Client) GetCacheReserveSpecificationWithOptions(runtime *dara.RuntimeOptions) (_result *GetCacheReserveSpecificationResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetCacheReserveSpecification"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCacheReserveSpecificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available specifications of cache reserve instances.
//
// @return GetCacheReserveSpecificationResponse
func (client *Client) GetCacheReserveSpecification() (_result *GetCacheReserveSpecificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCacheReserveSpecificationResponse{}
	_body, _err := client.GetCacheReserveSpecificationWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query a single cache configuration
//
// @param request - GetCacheRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCacheRuleResponse
func (client *Client) GetCacheRuleWithOptions(request *GetCacheRuleRequest, runtime *dara.RuntimeOptions) (_result *GetCacheRuleResponse, _err error) {
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
		Action:      dara.String("GetCacheRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCacheRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query a single cache configuration
//
// @param request - GetCacheRuleRequest
//
// @return GetCacheRuleResponse
func (client *Client) GetCacheRule(request *GetCacheRuleRequest) (_result *GetCacheRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCacheRuleResponse{}
	_body, _err := client.GetCacheRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Site Cache Tag Configuration
//
// @param request - GetCacheTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCacheTagResponse
func (client *Client) GetCacheTagWithOptions(request *GetCacheTagRequest, runtime *dara.RuntimeOptions) (_result *GetCacheTagResponse, _err error) {
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
		Action:      dara.String("GetCacheTag"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCacheTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Site Cache Tag Configuration
//
// @param request - GetCacheTagRequest
//
// @return GetCacheTagResponse
func (client *Client) GetCacheTag(request *GetCacheTagRequest) (_result *GetCacheTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCacheTagResponse{}
	_body, _err := client.GetCacheTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Retrieve the certificate, private key, and certificate information
//
// @param request - GetCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCertificateResponse
func (client *Client) GetCertificateWithOptions(request *GetCertificateRequest, runtime *dara.RuntimeOptions) (_result *GetCertificateResponse, _err error) {
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
		Action:      dara.String("GetCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve the certificate, private key, and certificate information
//
// @param request - GetCertificateRequest
//
// @return GetCertificateResponse
func (client *Client) GetCertificate(request *GetCertificateRequest) (_result *GetCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCertificateResponse{}
	_body, _err := client.GetCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query certificate quota and usage
//
// @param request - GetCertificateQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCertificateQuotaResponse
func (client *Client) GetCertificateQuotaWithOptions(request *GetCertificateQuotaRequest, runtime *dara.RuntimeOptions) (_result *GetCertificateQuotaResponse, _err error) {
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
		Action:      dara.String("GetCertificateQuota"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCertificateQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query certificate quota and usage
//
// @param request - GetCertificateQuotaRequest
//
// @return GetCertificateQuotaResponse
func (client *Client) GetCertificateQuota(request *GetCertificateQuotaRequest) (_result *GetCertificateQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCertificateQuotaResponse{}
	_body, _err := client.GetCertificateQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a client CA certificate.
//
// @param request - GetClientCaCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClientCaCertificateResponse
func (client *Client) GetClientCaCertificateWithOptions(request *GetClientCaCertificateRequest, runtime *dara.RuntimeOptions) (_result *GetClientCaCertificateResponse, _err error) {
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
		Action:      dara.String("GetClientCaCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClientCaCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a client CA certificate.
//
// @param request - GetClientCaCertificateRequest
//
// @return GetClientCaCertificateResponse
func (client *Client) GetClientCaCertificate(request *GetClientCaCertificateRequest) (_result *GetClientCaCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetClientCaCertificateResponse{}
	_body, _err := client.GetClientCaCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a client certificate.
//
// @param request - GetClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClientCertificateResponse
func (client *Client) GetClientCertificateWithOptions(request *GetClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *GetClientCertificateResponse, _err error) {
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
		Action:      dara.String("GetClientCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a client certificate.
//
// @param request - GetClientCertificateRequest
//
// @return GetClientCertificateResponse
func (client *Client) GetClientCertificate(request *GetClientCertificateRequest) (_result *GetClientCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetClientCertificateResponse{}
	_body, _err := client.GetClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries domain names associated with a client CA certificate. If no certificate is specified, domain names associated with an Edge Security Acceleration(ESA)-managed CA certificate are returned.
//
// @param request - GetClientCertificateHostnamesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClientCertificateHostnamesResponse
func (client *Client) GetClientCertificateHostnamesWithOptions(request *GetClientCertificateHostnamesRequest, runtime *dara.RuntimeOptions) (_result *GetClientCertificateHostnamesResponse, _err error) {
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
		Action:      dara.String("GetClientCertificateHostnames"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClientCertificateHostnamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries domain names associated with a client CA certificate. If no certificate is specified, domain names associated with an Edge Security Acceleration(ESA)-managed CA certificate are returned.
//
// @param request - GetClientCertificateHostnamesRequest
//
// @return GetClientCertificateHostnamesResponse
func (client *Client) GetClientCertificateHostnames(request *GetClientCertificateHostnamesRequest) (_result *GetClientCertificateHostnamesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetClientCertificateHostnamesResponse{}
	_body, _err := client.GetClientCertificateHostnamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Queries the CNAME flattening configuration of a website
//
// @param request - GetCnameFlatteningRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCnameFlatteningResponse
func (client *Client) GetCnameFlatteningWithOptions(request *GetCnameFlatteningRequest, runtime *dara.RuntimeOptions) (_result *GetCnameFlatteningResponse, _err error) {
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
		Action:      dara.String("GetCnameFlattening"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCnameFlatteningResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Queries the CNAME flattening configuration of a website
//
// @param request - GetCnameFlatteningRequest
//
// @return GetCnameFlatteningResponse
func (client *Client) GetCnameFlattening(request *GetCnameFlatteningRequest) (_result *GetCnameFlatteningResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCnameFlatteningResponse{}
	_body, _err := client.GetCnameFlatteningWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Compression Rule Details
//
// @param request - GetCompressionRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCompressionRuleResponse
func (client *Client) GetCompressionRuleWithOptions(request *GetCompressionRuleRequest, runtime *dara.RuntimeOptions) (_result *GetCompressionRuleResponse, _err error) {
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
		Action:      dara.String("GetCompressionRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCompressionRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Compression Rule Details
//
// @param request - GetCompressionRuleRequest
//
// @return GetCompressionRuleResponse
func (client *Client) GetCompressionRule(request *GetCompressionRuleRequest) (_result *GetCompressionRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCompressionRuleResponse{}
	_body, _err := client.GetCompressionRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration of Chinese mainland access optimization.
//
// @param request - GetCrossBorderOptimizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCrossBorderOptimizationResponse
func (client *Client) GetCrossBorderOptimizationWithOptions(request *GetCrossBorderOptimizationRequest, runtime *dara.RuntimeOptions) (_result *GetCrossBorderOptimizationResponse, _err error) {
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
		Action:      dara.String("GetCrossBorderOptimization"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCrossBorderOptimizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of Chinese mainland access optimization.
//
// @param request - GetCrossBorderOptimizationRequest
//
// @return GetCrossBorderOptimizationResponse
func (client *Client) GetCrossBorderOptimization(request *GetCrossBorderOptimizationRequest) (_result *GetCrossBorderOptimizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCrossBorderOptimizationResponse{}
	_body, _err := client.GetCrossBorderOptimizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Site Developer Mode Configuration
//
// @param request - GetDevelopmentModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDevelopmentModeResponse
func (client *Client) GetDevelopmentModeWithOptions(request *GetDevelopmentModeRequest, runtime *dara.RuntimeOptions) (_result *GetDevelopmentModeResponse, _err error) {
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
		Action:      dara.String("GetDevelopmentMode"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDevelopmentModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Site Developer Mode Configuration
//
// @param request - GetDevelopmentModeRequest
//
// @return GetDevelopmentModeResponse
func (client *Client) GetDevelopmentMode(request *GetDevelopmentModeRequest) (_result *GetDevelopmentModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDevelopmentModeResponse{}
	_body, _err := client.GetDevelopmentModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a containerized application, including basic application configurations and health check configurations.
//
// @param request - GetEdgeContainerAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEdgeContainerAppResponse
func (client *Client) GetEdgeContainerAppWithOptions(request *GetEdgeContainerAppRequest, runtime *dara.RuntimeOptions) (_result *GetEdgeContainerAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEdgeContainerApp"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEdgeContainerAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a containerized application, including basic application configurations and health check configurations.
//
// @param request - GetEdgeContainerAppRequest
//
// @return GetEdgeContainerAppResponse
func (client *Client) GetEdgeContainerApp(request *GetEdgeContainerAppRequest) (_result *GetEdgeContainerAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEdgeContainerAppResponse{}
	_body, _err := client.GetEdgeContainerAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the log collection configuration of a containerized application.
//
// @param request - GetEdgeContainerAppLogRiverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEdgeContainerAppLogRiverResponse
func (client *Client) GetEdgeContainerAppLogRiverWithOptions(request *GetEdgeContainerAppLogRiverRequest, runtime *dara.RuntimeOptions) (_result *GetEdgeContainerAppLogRiverResponse, _err error) {
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
		Action:      dara.String("GetEdgeContainerAppLogRiver"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEdgeContainerAppLogRiverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log collection configuration of a containerized application.
//
// @param request - GetEdgeContainerAppLogRiverRequest
//
// @return GetEdgeContainerAppLogRiverResponse
func (client *Client) GetEdgeContainerAppLogRiver(request *GetEdgeContainerAppLogRiverRequest) (_result *GetEdgeContainerAppLogRiverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEdgeContainerAppLogRiverResponse{}
	_body, _err := client.GetEdgeContainerAppLogRiverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the resource reservation configuration of the edge container.
//
// @param request - GetEdgeContainerAppResourceReserveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEdgeContainerAppResourceReserveResponse
func (client *Client) GetEdgeContainerAppResourceReserveWithOptions(request *GetEdgeContainerAppResourceReserveRequest, runtime *dara.RuntimeOptions) (_result *GetEdgeContainerAppResourceReserveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEdgeContainerAppResourceReserve"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEdgeContainerAppResourceReserveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the resource reservation configuration of the edge container.
//
// @param request - GetEdgeContainerAppResourceReserveRequest
//
// @return GetEdgeContainerAppResourceReserveResponse
func (client *Client) GetEdgeContainerAppResourceReserve(request *GetEdgeContainerAppResourceReserveRequest) (_result *GetEdgeContainerAppResourceReserveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEdgeContainerAppResourceReserveResponse{}
	_body, _err := client.GetEdgeContainerAppResourceReserveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the distribution of edge container application resources.
//
// @param request - GetEdgeContainerAppResourceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEdgeContainerAppResourceStatusResponse
func (client *Client) GetEdgeContainerAppResourceStatusWithOptions(request *GetEdgeContainerAppResourceStatusRequest, runtime *dara.RuntimeOptions) (_result *GetEdgeContainerAppResourceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEdgeContainerAppResourceStatus"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEdgeContainerAppResourceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the distribution of edge container application resources.
//
// @param request - GetEdgeContainerAppResourceStatusRequest
//
// @return GetEdgeContainerAppResourceStatusResponse
func (client *Client) GetEdgeContainerAppResourceStatus(request *GetEdgeContainerAppResourceStatusRequest) (_result *GetEdgeContainerAppResourceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEdgeContainerAppResourceStatusResponse{}
	_body, _err := client.GetEdgeContainerAppResourceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status information about a containerized application, including the deployment, release, and rollback of the application.
//
// @param request - GetEdgeContainerAppStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEdgeContainerAppStatusResponse
func (client *Client) GetEdgeContainerAppStatusWithOptions(request *GetEdgeContainerAppStatusRequest, runtime *dara.RuntimeOptions) (_result *GetEdgeContainerAppStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PublishEnv) {
		query["PublishEnv"] = request.PublishEnv
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEdgeContainerAppStatus"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEdgeContainerAppStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status information about a containerized application, including the deployment, release, and rollback of the application.
//
// @param request - GetEdgeContainerAppStatusRequest
//
// @return GetEdgeContainerAppStatusResponse
func (client *Client) GetEdgeContainerAppStatus(request *GetEdgeContainerAppStatusRequest) (_result *GetEdgeContainerAppStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEdgeContainerAppStatusResponse{}
	_body, _err := client.GetEdgeContainerAppStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a version of a containerized application. You can select an application version to release based on the version information.
//
// @param request - GetEdgeContainerAppVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEdgeContainerAppVersionResponse
func (client *Client) GetEdgeContainerAppVersionWithOptions(request *GetEdgeContainerAppVersionRequest, runtime *dara.RuntimeOptions) (_result *GetEdgeContainerAppVersionResponse, _err error) {
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
		Action:      dara.String("GetEdgeContainerAppVersion"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEdgeContainerAppVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a version of a containerized application. You can select an application version to release based on the version information.
//
// @param request - GetEdgeContainerAppVersionRequest
//
// @return GetEdgeContainerAppVersionResponse
func (client *Client) GetEdgeContainerAppVersion(request *GetEdgeContainerAppVersionRequest) (_result *GetEdgeContainerAppVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEdgeContainerAppVersionResponse{}
	_body, _err := client.GetEdgeContainerAppVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries regions where a containerized application is deployed based on the application ID.
//
// @param request - GetEdgeContainerDeployRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEdgeContainerDeployRegionsResponse
func (client *Client) GetEdgeContainerDeployRegionsWithOptions(request *GetEdgeContainerDeployRegionsRequest, runtime *dara.RuntimeOptions) (_result *GetEdgeContainerDeployRegionsResponse, _err error) {
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
		Action:      dara.String("GetEdgeContainerDeployRegions"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEdgeContainerDeployRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries regions where a containerized application is deployed based on the application ID.
//
// @param request - GetEdgeContainerDeployRegionsRequest
//
// @return GetEdgeContainerDeployRegionsResponse
func (client *Client) GetEdgeContainerDeployRegions(request *GetEdgeContainerDeployRegionsRequest) (_result *GetEdgeContainerDeployRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEdgeContainerDeployRegionsResponse{}
	_body, _err := client.GetEdgeContainerDeployRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Edge Container logs.
//
// @param request - GetEdgeContainerLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEdgeContainerLogsResponse
func (client *Client) GetEdgeContainerLogsWithOptions(request *GetEdgeContainerLogsRequest, runtime *dara.RuntimeOptions) (_result *GetEdgeContainerLogsResponse, _err error) {
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
		Action:      dara.String("GetEdgeContainerLogs"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEdgeContainerLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Edge Container logs.
//
// @param request - GetEdgeContainerLogsRequest
//
// @return GetEdgeContainerLogsResponse
func (client *Client) GetEdgeContainerLogs(request *GetEdgeContainerLogsRequest) (_result *GetEdgeContainerLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEdgeContainerLogsResponse{}
	_body, _err := client.GetEdgeContainerLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the deployment status of an application in the staging environment by using the application ID.
//
// @param request - GetEdgeContainerStagingDeployStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEdgeContainerStagingDeployStatusResponse
func (client *Client) GetEdgeContainerStagingDeployStatusWithOptions(request *GetEdgeContainerStagingDeployStatusRequest, runtime *dara.RuntimeOptions) (_result *GetEdgeContainerStagingDeployStatusResponse, _err error) {
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
		Action:      dara.String("GetEdgeContainerStagingDeployStatus"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEdgeContainerStagingDeployStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the deployment status of an application in the staging environment by using the application ID.
//
// @param request - GetEdgeContainerStagingDeployStatusRequest
//
// @return GetEdgeContainerStagingDeployStatusResponse
func (client *Client) GetEdgeContainerStagingDeployStatus(request *GetEdgeContainerStagingDeployStatusRequest) (_result *GetEdgeContainerStagingDeployStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEdgeContainerStagingDeployStatusResponse{}
	_body, _err := client.GetEdgeContainerStagingDeployStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the terminal information of a containerized application.
//
// @param request - GetEdgeContainerTerminalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEdgeContainerTerminalResponse
func (client *Client) GetEdgeContainerTerminalWithOptions(request *GetEdgeContainerTerminalRequest, runtime *dara.RuntimeOptions) (_result *GetEdgeContainerTerminalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEdgeContainerTerminal"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEdgeContainerTerminalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the terminal information of a containerized application.
//
// @param request - GetEdgeContainerTerminalRequest
//
// @return GetEdgeContainerTerminalResponse
func (client *Client) GetEdgeContainerTerminal(request *GetEdgeContainerTerminalRequest) (_result *GetEdgeContainerTerminalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEdgeContainerTerminalResponse{}
	_body, _err := client.GetEdgeContainerTerminalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks the status of Edge Routine.
//
// @param request - GetErServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetErServiceResponse
func (client *Client) GetErServiceWithOptions(request *GetErServiceRequest, runtime *dara.RuntimeOptions) (_result *GetErServiceResponse, _err error) {
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
		Action:      dara.String("GetErService"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetErServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the status of Edge Routine.
//
// @param request - GetErServiceRequest
//
// @return GetErServiceResponse
func (client *Client) GetErService(request *GetErServiceRequest) (_result *GetErServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetErServiceResponse{}
	_body, _err := client.GetErServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration details of an HTTP request header modification rule for a website.
//
// @param request - GetHttpIncomingRequestHeaderModificationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHttpIncomingRequestHeaderModificationRuleResponse
func (client *Client) GetHttpIncomingRequestHeaderModificationRuleWithOptions(request *GetHttpIncomingRequestHeaderModificationRuleRequest, runtime *dara.RuntimeOptions) (_result *GetHttpIncomingRequestHeaderModificationRuleResponse, _err error) {
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
		Action:      dara.String("GetHttpIncomingRequestHeaderModificationRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHttpIncomingRequestHeaderModificationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration details of an HTTP request header modification rule for a website.
//
// @param request - GetHttpIncomingRequestHeaderModificationRuleRequest
//
// @return GetHttpIncomingRequestHeaderModificationRuleResponse
func (client *Client) GetHttpIncomingRequestHeaderModificationRule(request *GetHttpIncomingRequestHeaderModificationRuleRequest) (_result *GetHttpIncomingRequestHeaderModificationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHttpIncomingRequestHeaderModificationRuleResponse{}
	_body, _err := client.GetHttpIncomingRequestHeaderModificationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration details of an incoming HTTP response header modification rule for a website.
//
// @param request - GetHttpIncomingResponseHeaderModificationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHttpIncomingResponseHeaderModificationRuleResponse
func (client *Client) GetHttpIncomingResponseHeaderModificationRuleWithOptions(request *GetHttpIncomingResponseHeaderModificationRuleRequest, runtime *dara.RuntimeOptions) (_result *GetHttpIncomingResponseHeaderModificationRuleResponse, _err error) {
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
		Action:      dara.String("GetHttpIncomingResponseHeaderModificationRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHttpIncomingResponseHeaderModificationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration details of an incoming HTTP response header modification rule for a website.
//
// @param request - GetHttpIncomingResponseHeaderModificationRuleRequest
//
// @return GetHttpIncomingResponseHeaderModificationRuleResponse
func (client *Client) GetHttpIncomingResponseHeaderModificationRule(request *GetHttpIncomingResponseHeaderModificationRuleRequest) (_result *GetHttpIncomingResponseHeaderModificationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHttpIncomingResponseHeaderModificationRuleResponse{}
	_body, _err := client.GetHttpIncomingResponseHeaderModificationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query HTTP Request Header Rule Details
//
// @param request - GetHttpRequestHeaderModificationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHttpRequestHeaderModificationRuleResponse
func (client *Client) GetHttpRequestHeaderModificationRuleWithOptions(request *GetHttpRequestHeaderModificationRuleRequest, runtime *dara.RuntimeOptions) (_result *GetHttpRequestHeaderModificationRuleResponse, _err error) {
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
		Action:      dara.String("GetHttpRequestHeaderModificationRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHttpRequestHeaderModificationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query HTTP Request Header Rule Details
//
// @param request - GetHttpRequestHeaderModificationRuleRequest
//
// @return GetHttpRequestHeaderModificationRuleResponse
func (client *Client) GetHttpRequestHeaderModificationRule(request *GetHttpRequestHeaderModificationRuleRequest) (_result *GetHttpRequestHeaderModificationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHttpRequestHeaderModificationRuleResponse{}
	_body, _err := client.GetHttpRequestHeaderModificationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query HTTP Response Header Rules
//
// @param request - GetHttpResponseHeaderModificationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHttpResponseHeaderModificationRuleResponse
func (client *Client) GetHttpResponseHeaderModificationRuleWithOptions(request *GetHttpResponseHeaderModificationRuleRequest, runtime *dara.RuntimeOptions) (_result *GetHttpResponseHeaderModificationRuleResponse, _err error) {
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
		Action:      dara.String("GetHttpResponseHeaderModificationRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHttpResponseHeaderModificationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query HTTP Response Header Rules
//
// @param request - GetHttpResponseHeaderModificationRuleRequest
//
// @return GetHttpResponseHeaderModificationRuleResponse
func (client *Client) GetHttpResponseHeaderModificationRule(request *GetHttpResponseHeaderModificationRuleRequest) (_result *GetHttpResponseHeaderModificationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHttpResponseHeaderModificationRuleResponse{}
	_body, _err := client.GetHttpResponseHeaderModificationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query a Single HTTPS Application Configuration
//
// @param request - GetHttpsApplicationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHttpsApplicationConfigurationResponse
func (client *Client) GetHttpsApplicationConfigurationWithOptions(request *GetHttpsApplicationConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetHttpsApplicationConfigurationResponse, _err error) {
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
		Action:      dara.String("GetHttpsApplicationConfiguration"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHttpsApplicationConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query a Single HTTPS Application Configuration
//
// @param request - GetHttpsApplicationConfigurationRequest
//
// @return GetHttpsApplicationConfigurationResponse
func (client *Client) GetHttpsApplicationConfiguration(request *GetHttpsApplicationConfigurationRequest) (_result *GetHttpsApplicationConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHttpsApplicationConfigurationResponse{}
	_body, _err := client.GetHttpsApplicationConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query a Single HTTPS Basic Configuration
//
// @param request - GetHttpsBasicConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHttpsBasicConfigurationResponse
func (client *Client) GetHttpsBasicConfigurationWithOptions(request *GetHttpsBasicConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetHttpsBasicConfigurationResponse, _err error) {
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
		Action:      dara.String("GetHttpsBasicConfiguration"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHttpsBasicConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query a Single HTTPS Basic Configuration
//
// @param request - GetHttpsBasicConfigurationRequest
//
// @return GetHttpsBasicConfigurationResponse
func (client *Client) GetHttpsBasicConfiguration(request *GetHttpsBasicConfigurationRequest) (_result *GetHttpsBasicConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHttpsBasicConfigurationResponse{}
	_body, _err := client.GetHttpsBasicConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IPv6 configuration of a website.
//
// @param request - GetIPv6Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIPv6Response
func (client *Client) GetIPv6WithOptions(request *GetIPv6Request, runtime *dara.RuntimeOptions) (_result *GetIPv6Response, _err error) {
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
		Action:      dara.String("GetIPv6"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIPv6Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IPv6 configuration of a website.
//
// @param request - GetIPv6Request
//
// @return GetIPv6Response
func (client *Client) GetIPv6(request *GetIPv6Request) (_result *GetIPv6Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIPv6Response{}
	_body, _err := client.GetIPv6WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Single Site Image Transformation Configuration
//
// @param request - GetImageTransformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageTransformResponse
func (client *Client) GetImageTransformWithOptions(request *GetImageTransformRequest, runtime *dara.RuntimeOptions) (_result *GetImageTransformResponse, _err error) {
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
		Action:      dara.String("GetImageTransform"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageTransformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Single Site Image Transformation Configuration
//
// @param request - GetImageTransformRequest
//
// @return GetImageTransformResponse
func (client *Client) GetImageTransform(request *GetImageTransformRequest) (_result *GetImageTransformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImageTransformResponse{}
	_body, _err := client.GetImageTransformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the value of a key in a key-value pair.
//
// @param request - GetKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKvResponse
func (client *Client) GetKvWithOptions(request *GetKvRequest, runtime *dara.RuntimeOptions) (_result *GetKvResponse, _err error) {
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
		Action:      dara.String("GetKv"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the value of a key in a key-value pair.
//
// @param request - GetKvRequest
//
// @return GetKvResponse
func (client *Client) GetKv(request *GetKvRequest) (_result *GetKvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetKvResponse{}
	_body, _err := client.GetKvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Edge KV usage in your Alibaba Cloud account, including the information about all namespaces.
//
// @param request - GetKvAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKvAccountResponse
func (client *Client) GetKvAccountWithOptions(runtime *dara.RuntimeOptions) (_result *GetKvAccountResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetKvAccount"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKvAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Edge KV usage in your Alibaba Cloud account, including the information about all namespaces.
//
// @return GetKvAccountResponse
func (client *Client) GetKvAccount() (_result *GetKvAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetKvAccountResponse{}
	_body, _err := client.GetKvAccountWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a namespace in your Alibaba Cloud account.
//
// @param request - GetKvNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKvNamespaceResponse
func (client *Client) GetKvNamespaceWithOptions(request *GetKvNamespaceRequest, runtime *dara.RuntimeOptions) (_result *GetKvNamespaceResponse, _err error) {
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
		Action:      dara.String("GetKvNamespace"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKvNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a namespace in your Alibaba Cloud account.
//
// @param request - GetKvNamespaceRequest
//
// @return GetKvNamespaceResponse
func (client *Client) GetKvNamespace(request *GetKvNamespaceRequest) (_result *GetKvNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetKvNamespaceResponse{}
	_body, _err := client.GetKvNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a custom list, such as the name, description, type, and content.
//
// @param request - GetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetListResponse
func (client *Client) GetListWithOptions(request *GetListRequest, runtime *dara.RuntimeOptions) (_result *GetListResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetList"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a custom list, such as the name, description, type, and content.
//
// @param request - GetListRequest
//
// @return GetListResponse
func (client *Client) GetList(request *GetListRequest) (_result *GetListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetListResponse{}
	_body, _err := client.GetListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query a Specific Load Balancer
//
// Description:
//
// This API allows users to query the configuration details of a specific load balancer by providing necessary authentication information and resource identifiers, including but not limited to name, session persistence strategy, routing policy, etc.
//
// @param request - GetLoadBalancerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoadBalancerResponse
func (client *Client) GetLoadBalancerWithOptions(request *GetLoadBalancerRequest, runtime *dara.RuntimeOptions) (_result *GetLoadBalancerResponse, _err error) {
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
		Action:      dara.String("GetLoadBalancer"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query a Specific Load Balancer
//
// Description:
//
// This API allows users to query the configuration details of a specific load balancer by providing necessary authentication information and resource identifiers, including but not limited to name, session persistence strategy, routing policy, etc.
//
// @param request - GetLoadBalancerRequest
//
// @return GetLoadBalancerResponse
func (client *Client) GetLoadBalancer(request *GetLoadBalancerRequest) (_result *GetLoadBalancerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLoadBalancerResponse{}
	_body, _err := client.GetLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Managed Transform Configuration
//
// @param request - GetManagedTransformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetManagedTransformResponse
func (client *Client) GetManagedTransformWithOptions(request *GetManagedTransformRequest, runtime *dara.RuntimeOptions) (_result *GetManagedTransformResponse, _err error) {
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
		Action:      dara.String("GetManagedTransform"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetManagedTransformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Managed Transform Configuration
//
// @param request - GetManagedTransformRequest
//
// @return GetManagedTransformResponse
func (client *Client) GetManagedTransform(request *GetManagedTransformRequest) (_result *GetManagedTransformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetManagedTransformResponse{}
	_body, _err := client.GetManagedTransformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query a single network optimization configuration
//
// @param request - GetNetworkOptimizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkOptimizationResponse
func (client *Client) GetNetworkOptimizationWithOptions(request *GetNetworkOptimizationRequest, runtime *dara.RuntimeOptions) (_result *GetNetworkOptimizationResponse, _err error) {
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
		Action:      dara.String("GetNetworkOptimization"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNetworkOptimizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query a single network optimization configuration
//
// @param request - GetNetworkOptimizationRequest
//
// @return GetNetworkOptimizationResponse
func (client *Client) GetNetworkOptimization(request *GetNetworkOptimizationRequest) (_result *GetNetworkOptimizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNetworkOptimizationResponse{}
	_body, _err := client.GetNetworkOptimizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取源服务器CA证书信息
//
// @param request - GetOriginCaCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOriginCaCertificateResponse
func (client *Client) GetOriginCaCertificateWithOptions(request *GetOriginCaCertificateRequest, runtime *dara.RuntimeOptions) (_result *GetOriginCaCertificateResponse, _err error) {
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
		Action:      dara.String("GetOriginCaCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOriginCaCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取源服务器CA证书信息
//
// @param request - GetOriginCaCertificateRequest
//
// @return GetOriginCaCertificateResponse
func (client *Client) GetOriginCaCertificate(request *GetOriginCaCertificateRequest) (_result *GetOriginCaCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOriginCaCertificateResponse{}
	_body, _err := client.GetOriginCaCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取域名回源客户端证书信息
//
// @param request - GetOriginClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOriginClientCertificateResponse
func (client *Client) GetOriginClientCertificateWithOptions(request *GetOriginClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *GetOriginClientCertificateResponse, _err error) {
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
		Action:      dara.String("GetOriginClientCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOriginClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取域名回源客户端证书信息
//
// @param request - GetOriginClientCertificateRequest
//
// @return GetOriginClientCertificateResponse
func (client *Client) GetOriginClientCertificate(request *GetOriginClientCertificateRequest) (_result *GetOriginClientCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOriginClientCertificateResponse{}
	_body, _err := client.GetOriginClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取域名回源客户端证书绑定的域名列表
//
// @param request - GetOriginClientCertificateHostnamesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOriginClientCertificateHostnamesResponse
func (client *Client) GetOriginClientCertificateHostnamesWithOptions(request *GetOriginClientCertificateHostnamesRequest, runtime *dara.RuntimeOptions) (_result *GetOriginClientCertificateHostnamesResponse, _err error) {
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
		Action:      dara.String("GetOriginClientCertificateHostnames"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOriginClientCertificateHostnamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取域名回源客户端证书绑定的域名列表
//
// @param request - GetOriginClientCertificateHostnamesRequest
//
// @return GetOriginClientCertificateHostnamesResponse
func (client *Client) GetOriginClientCertificateHostnames(request *GetOriginClientCertificateHostnamesRequest) (_result *GetOriginClientCertificateHostnamesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOriginClientCertificateHostnamesResponse{}
	_body, _err := client.GetOriginClientCertificateHostnamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query a specific origin pool
//
// @param request - GetOriginPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOriginPoolResponse
func (client *Client) GetOriginPoolWithOptions(request *GetOriginPoolRequest, runtime *dara.RuntimeOptions) (_result *GetOriginPoolResponse, _err error) {
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
		Action:      dara.String("GetOriginPool"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOriginPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query a specific origin pool
//
// @param request - GetOriginPoolRequest
//
// @return GetOriginPoolResponse
func (client *Client) GetOriginPool(request *GetOriginPoolRequest) (_result *GetOriginPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOriginPoolResponse{}
	_body, _err := client.GetOriginPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the origin protection configurations of a website, including the origin protection, IP convergence, and the status and details of the IP whitelist for origin protection. The details includes the IP whitelist used by the website, the latest IP whitelist, and the differences between them.
//
// @param request - GetOriginProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOriginProtectionResponse
func (client *Client) GetOriginProtectionWithOptions(request *GetOriginProtectionRequest, runtime *dara.RuntimeOptions) (_result *GetOriginProtectionResponse, _err error) {
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
		Action:      dara.String("GetOriginProtection"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOriginProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the origin protection configurations of a website, including the origin protection, IP convergence, and the status and details of the IP whitelist for origin protection. The details includes the IP whitelist used by the website, the latest IP whitelist, and the differences between them.
//
// @param request - GetOriginProtectionRequest
//
// @return GetOriginProtectionResponse
func (client *Client) GetOriginProtection(request *GetOriginProtectionRequest) (_result *GetOriginProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOriginProtectionResponse{}
	_body, _err := client.GetOriginProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query a Single Origin Rule Configuration
//
// @param request - GetOriginRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOriginRuleResponse
func (client *Client) GetOriginRuleWithOptions(request *GetOriginRuleRequest, runtime *dara.RuntimeOptions) (_result *GetOriginRuleResponse, _err error) {
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
		Action:      dara.String("GetOriginRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOriginRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query a Single Origin Rule Configuration
//
// @param request - GetOriginRuleRequest
//
// @return GetOriginRuleResponse
func (client *Client) GetOriginRule(request *GetOriginRuleRequest) (_result *GetOriginRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOriginRuleResponse{}
	_body, _err := client.GetOriginRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a custom error page based on the error page ID.
//
// @param request - GetPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPageResponse
func (client *Client) GetPageWithOptions(request *GetPageRequest, runtime *dara.RuntimeOptions) (_result *GetPageResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPage"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a custom error page based on the error page ID.
//
// @param request - GetPageRequest
//
// @return GetPageResponse
func (client *Client) GetPage(request *GetPageRequest) (_result *GetPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPageResponse{}
	_body, _err := client.GetPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the quotas and quota usage for different cache purge options.
//
// @param request - GetPurgeQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPurgeQuotaResponse
func (client *Client) GetPurgeQuotaWithOptions(request *GetPurgeQuotaRequest, runtime *dara.RuntimeOptions) (_result *GetPurgeQuotaResponse, _err error) {
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
		Action:      dara.String("GetPurgeQuota"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPurgeQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quotas and quota usage for different cache purge options.
//
// @param request - GetPurgeQuotaRequest
//
// @return GetPurgeQuotaResponse
func (client *Client) GetPurgeQuota(request *GetPurgeQuotaRequest) (_result *GetPurgeQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPurgeQuotaResponse{}
	_body, _err := client.GetPurgeQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the fields in real-time logs based on the log category.
//
// @param request - GetRealtimeDeliveryFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRealtimeDeliveryFieldResponse
func (client *Client) GetRealtimeDeliveryFieldWithOptions(request *GetRealtimeDeliveryFieldRequest, runtime *dara.RuntimeOptions) (_result *GetRealtimeDeliveryFieldResponse, _err error) {
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
		Action:      dara.String("GetRealtimeDeliveryField"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRealtimeDeliveryFieldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the fields in real-time logs based on the log category.
//
// @param request - GetRealtimeDeliveryFieldRequest
//
// @return GetRealtimeDeliveryFieldResponse
func (client *Client) GetRealtimeDeliveryField(request *GetRealtimeDeliveryFieldRequest) (_result *GetRealtimeDeliveryFieldResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRealtimeDeliveryFieldResponse{}
	_body, _err := client.GetRealtimeDeliveryFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration of a single DNS record, such as the record value, priority, and origin authentication setting (exclusive to CNAME records).
//
// @param request - GetRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecordResponse
func (client *Client) GetRecordWithOptions(request *GetRecordRequest, runtime *dara.RuntimeOptions) (_result *GetRecordResponse, _err error) {
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
		Action:      dara.String("GetRecord"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of a single DNS record, such as the record value, priority, and origin authentication setting (exclusive to CNAME records).
//
// @param request - GetRecordRequest
//
// @return GetRecordResponse
func (client *Client) GetRecord(request *GetRecordRequest) (_result *GetRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRecordResponse{}
	_body, _err := client.GetRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Redirect Rule Details
//
// @param request - GetRedirectRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRedirectRuleResponse
func (client *Client) GetRedirectRuleWithOptions(request *GetRedirectRuleRequest, runtime *dara.RuntimeOptions) (_result *GetRedirectRuleResponse, _err error) {
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
		Action:      dara.String("GetRedirectRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRedirectRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Redirect Rule Details
//
// @param request - GetRedirectRuleRequest
//
// @return GetRedirectRuleResponse
func (client *Client) GetRedirectRule(request *GetRedirectRuleRequest) (_result *GetRedirectRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRedirectRuleResponse{}
	_body, _err := client.GetRedirectRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query details of the rewrite URL rule
//
// @param request - GetRewriteUrlRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRewriteUrlRuleResponse
func (client *Client) GetRewriteUrlRuleWithOptions(request *GetRewriteUrlRuleRequest, runtime *dara.RuntimeOptions) (_result *GetRewriteUrlRuleResponse, _err error) {
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
		Action:      dara.String("GetRewriteUrlRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRewriteUrlRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query details of the rewrite URL rule
//
// @param request - GetRewriteUrlRuleRequest
//
// @return GetRewriteUrlRuleResponse
func (client *Client) GetRewriteUrlRule(request *GetRewriteUrlRuleRequest) (_result *GetRewriteUrlRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRewriteUrlRuleResponse{}
	_body, _err := client.GetRewriteUrlRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of a routine, including the code versions and the configurations of the environments, associated domain names, and associated routes.
//
// @param request - GetRoutineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoutineResponse
func (client *Client) GetRoutineWithOptions(request *GetRoutineRequest, runtime *dara.RuntimeOptions) (_result *GetRoutineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRoutine"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRoutineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a routine, including the code versions and the configurations of the environments, associated domain names, and associated routes.
//
// @param request - GetRoutineRequest
//
// @return GetRoutineResponse
func (client *Client) GetRoutine(request *GetRoutineRequest) (_result *GetRoutineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRoutineResponse{}
	_body, _err := client.GetRoutineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a code version of a routine.
//
// @param request - GetRoutineCodeVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoutineCodeVersionResponse
func (client *Client) GetRoutineCodeVersionWithOptions(request *GetRoutineCodeVersionRequest, runtime *dara.RuntimeOptions) (_result *GetRoutineCodeVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CodeVersion) {
		body["CodeVersion"] = request.CodeVersion
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRoutineCodeVersion"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRoutineCodeVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a code version of a routine.
//
// @param request - GetRoutineCodeVersionRequest
//
// @return GetRoutineCodeVersionResponse
func (client *Client) GetRoutineCodeVersion(request *GetRoutineCodeVersionRequest) (_result *GetRoutineCodeVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRoutineCodeVersionResponse{}
	_body, _err := client.GetRoutineCodeVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the route configurations of a single edge function.
//
// @param request - GetRoutineRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoutineRouteResponse
func (client *Client) GetRoutineRouteWithOptions(request *GetRoutineRouteRequest, runtime *dara.RuntimeOptions) (_result *GetRoutineRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRoutineRoute"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRoutineRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the route configurations of a single edge function.
//
// @param request - GetRoutineRouteRequest
//
// @return GetRoutineRouteResponse
func (client *Client) GetRoutineRoute(request *GetRoutineRouteRequest) (_result *GetRoutineRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRoutineRouteResponse{}
	_body, _err := client.GetRoutineRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the release information about the routine code that is released to the staging environment. This information can be used to upload the test code to Object Storage Service (OSS).
//
// Description:
//
//	  Every time the code of a routine is released to the staging environment, a version number is generated. Such code is for tests only.
//
//		- A routine can retain a maximum of 10 code versions. If the number of versions reaches the limit, you must call the DeleteRoutineCodeRevision operation to delete unwanted versions.
//
// @param request - GetRoutineStagingCodeUploadInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoutineStagingCodeUploadInfoResponse
func (client *Client) GetRoutineStagingCodeUploadInfoWithOptions(request *GetRoutineStagingCodeUploadInfoRequest, runtime *dara.RuntimeOptions) (_result *GetRoutineStagingCodeUploadInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CodeDescription) {
		body["CodeDescription"] = request.CodeDescription
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRoutineStagingCodeUploadInfo"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRoutineStagingCodeUploadInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the release information about the routine code that is released to the staging environment. This information can be used to upload the test code to Object Storage Service (OSS).
//
// Description:
//
//	  Every time the code of a routine is released to the staging environment, a version number is generated. Such code is for tests only.
//
//		- A routine can retain a maximum of 10 code versions. If the number of versions reaches the limit, you must call the DeleteRoutineCodeRevision operation to delete unwanted versions.
//
// @param request - GetRoutineStagingCodeUploadInfoRequest
//
// @return GetRoutineStagingCodeUploadInfoResponse
func (client *Client) GetRoutineStagingCodeUploadInfo(request *GetRoutineStagingCodeUploadInfoRequest) (_result *GetRoutineStagingCodeUploadInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRoutineStagingCodeUploadInfoResponse{}
	_body, _err := client.GetRoutineStagingCodeUploadInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IP addresses of staging environments for Edge Routine.
//
// @param request - GetRoutineStagingEnvIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoutineStagingEnvIpResponse
func (client *Client) GetRoutineStagingEnvIpWithOptions(runtime *dara.RuntimeOptions) (_result *GetRoutineStagingEnvIpResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetRoutineStagingEnvIp"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRoutineStagingEnvIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP addresses of staging environments for Edge Routine.
//
// @return GetRoutineStagingEnvIpResponse
func (client *Client) GetRoutineStagingEnvIp() (_result *GetRoutineStagingEnvIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRoutineStagingEnvIpResponse{}
	_body, _err := client.GetRoutineStagingEnvIpWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Edge Routine information in your Alibaba Cloud account, including the associated subdomain and created routines.
//
// @param request - GetRoutineUserInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoutineUserInfoResponse
func (client *Client) GetRoutineUserInfoWithOptions(runtime *dara.RuntimeOptions) (_result *GetRoutineUserInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetRoutineUserInfo"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRoutineUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Edge Routine information in your Alibaba Cloud account, including the associated subdomain and created routines.
//
// @return GetRoutineUserInfoResponse
func (client *Client) GetRoutineUserInfo() (_result *GetRoutineUserInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRoutineUserInfoResponse{}
	_body, _err := client.GetRoutineUserInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a specified scheduled prefetch task based on the task ID.
//
// @param request - GetScheduledPreloadJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScheduledPreloadJobResponse
func (client *Client) GetScheduledPreloadJobWithOptions(request *GetScheduledPreloadJobRequest, runtime *dara.RuntimeOptions) (_result *GetScheduledPreloadJobResponse, _err error) {
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
		Action:      dara.String("GetScheduledPreloadJob"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScheduledPreloadJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a specified scheduled prefetch task based on the task ID.
//
// @param request - GetScheduledPreloadJobRequest
//
// @return GetScheduledPreloadJobResponse
func (client *Client) GetScheduledPreloadJob(request *GetScheduledPreloadJobRequest) (_result *GetScheduledPreloadJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetScheduledPreloadJobResponse{}
	_body, _err := client.GetScheduledPreloadJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration for search engine crawler of a website.
//
// @param request - GetSeoBypassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSeoBypassResponse
func (client *Client) GetSeoBypassWithOptions(request *GetSeoBypassRequest, runtime *dara.RuntimeOptions) (_result *GetSeoBypassResponse, _err error) {
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
		Action:      dara.String("GetSeoBypass"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSeoBypassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration for search engine crawler of a website.
//
// @param request - GetSeoBypassRequest
//
// @return GetSeoBypassResponse
func (client *Client) GetSeoBypass(request *GetSeoBypassRequest) (_result *GetSeoBypassResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSeoBypassResponse{}
	_body, _err := client.GetSeoBypassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a website based on the website ID.
//
// @param request - GetSiteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSiteResponse
func (client *Client) GetSiteWithOptions(request *GetSiteRequest, runtime *dara.RuntimeOptions) (_result *GetSiteResponse, _err error) {
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
		Action:      dara.String("GetSite"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSiteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a website based on the website ID.
//
// @param request - GetSiteRequest
//
// @return GetSiteResponse
func (client *Client) GetSite(request *GetSiteRequest) (_result *GetSiteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSiteResponse{}
	_body, _err := client.GetSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the nameservers configured for a website.
//
// @param request - GetSiteCurrentNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSiteCurrentNSResponse
func (client *Client) GetSiteCurrentNSWithOptions(request *GetSiteCurrentNSRequest, runtime *dara.RuntimeOptions) (_result *GetSiteCurrentNSResponse, _err error) {
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
		Action:      dara.String("GetSiteCurrentNS"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSiteCurrentNSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the nameservers configured for a website.
//
// @param request - GetSiteCurrentNSRequest
//
// @return GetSiteCurrentNSResponse
func (client *Client) GetSiteCurrentNS(request *GetSiteCurrentNSRequest) (_result *GetSiteCurrentNSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSiteCurrentNSResponse{}
	_body, _err := client.GetSiteCurrentNSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration of custom log fields for a website.
//
// Description:
//
//	  **Description**: You can call this operation to query the configuration of custom log fields for a website, including custom fields in request headers, response headers, and cookies.
//
//		- **Scenarios**: You can call this operation in scenarios where you need to obtain specific HTTP headers or cookie information for log analysis.
//
//		- ****
//
// @param request - GetSiteCustomLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSiteCustomLogResponse
func (client *Client) GetSiteCustomLogWithOptions(request *GetSiteCustomLogRequest, runtime *dara.RuntimeOptions) (_result *GetSiteCustomLogResponse, _err error) {
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
		Action:      dara.String("GetSiteCustomLog"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSiteCustomLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of custom log fields for a website.
//
// Description:
//
//	  **Description**: You can call this operation to query the configuration of custom log fields for a website, including custom fields in request headers, response headers, and cookies.
//
//		- **Scenarios**: You can call this operation in scenarios where you need to obtain specific HTTP headers or cookie information for log analysis.
//
//		- ****
//
// @param request - GetSiteCustomLogRequest
//
// @return GetSiteCustomLogResponse
func (client *Client) GetSiteCustomLog(request *GetSiteCustomLogRequest) (_result *GetSiteCustomLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSiteCustomLogResponse{}
	_body, _err := client.GetSiteCustomLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a real-time log delivery task.
//
// @param request - GetSiteDeliveryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSiteDeliveryTaskResponse
func (client *Client) GetSiteDeliveryTaskWithOptions(request *GetSiteDeliveryTaskRequest, runtime *dara.RuntimeOptions) (_result *GetSiteDeliveryTaskResponse, _err error) {
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
		Action:      dara.String("GetSiteDeliveryTask"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSiteDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a real-time log delivery task.
//
// @param request - GetSiteDeliveryTaskRequest
//
// @return GetSiteDeliveryTaskResponse
func (client *Client) GetSiteDeliveryTask(request *GetSiteDeliveryTaskRequest) (_result *GetSiteDeliveryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSiteDeliveryTaskResponse{}
	_body, _err := client.GetSiteDeliveryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the remaining quota for delivering a specific category of real-time logs in a website.
//
// Description:
//
// You can call this operation to query the remaining quota for delivering a specific category of real-time logs in a website within an Alibaba Cloud account. This is essential for monitoring and managing your log delivery capacity to ensure that logs can be delivered to the destination and prevent data loss or latency caused by insufficient quota.
//
// **Take note of the following parameters:**
//
//   - “
//
//   - `BusinessType` is required. You must specify a log category to obtain the corresponding quota information.
//
//   - `SiteId` specifies the ID of a website, which must be a valid integer that corresponds to a website that you configured on Alibaba Cloud.
//
// **Response:**
//
//   - If a request is successful, the system returns the remaining log delivery quota (`FreeQuota`), request ID (`RequestId`), website ID (`SiteId`), and log category (`BusinessType`). You can confirm and record the returned data.
//
// @param request - GetSiteLogDeliveryQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSiteLogDeliveryQuotaResponse
func (client *Client) GetSiteLogDeliveryQuotaWithOptions(request *GetSiteLogDeliveryQuotaRequest, runtime *dara.RuntimeOptions) (_result *GetSiteLogDeliveryQuotaResponse, _err error) {
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
		Action:      dara.String("GetSiteLogDeliveryQuota"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSiteLogDeliveryQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the remaining quota for delivering a specific category of real-time logs in a website.
//
// Description:
//
// You can call this operation to query the remaining quota for delivering a specific category of real-time logs in a website within an Alibaba Cloud account. This is essential for monitoring and managing your log delivery capacity to ensure that logs can be delivered to the destination and prevent data loss or latency caused by insufficient quota.
//
// **Take note of the following parameters:**
//
//   - “
//
//   - `BusinessType` is required. You must specify a log category to obtain the corresponding quota information.
//
//   - `SiteId` specifies the ID of a website, which must be a valid integer that corresponds to a website that you configured on Alibaba Cloud.
//
// **Response:**
//
//   - If a request is successful, the system returns the remaining log delivery quota (`FreeQuota`), request ID (`RequestId`), website ID (`SiteId`), and log category (`BusinessType`). You can confirm and record the returned data.
//
// @param request - GetSiteLogDeliveryQuotaRequest
//
// @return GetSiteLogDeliveryQuotaResponse
func (client *Client) GetSiteLogDeliveryQuota(request *GetSiteLogDeliveryQuotaRequest) (_result *GetSiteLogDeliveryQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSiteLogDeliveryQuotaResponse{}
	_body, _err := client.GetSiteLogDeliveryQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the site hold configuration of a website. After you enable site hold, other accounts cannot add your website domain or its subdomains to ESA.
//
// @param request - GetSiteNameExclusiveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSiteNameExclusiveResponse
func (client *Client) GetSiteNameExclusiveWithOptions(request *GetSiteNameExclusiveRequest, runtime *dara.RuntimeOptions) (_result *GetSiteNameExclusiveResponse, _err error) {
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
		Action:      dara.String("GetSiteNameExclusive"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSiteNameExclusiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the site hold configuration of a website. After you enable site hold, other accounts cannot add your website domain or its subdomains to ESA.
//
// @param request - GetSiteNameExclusiveRequest
//
// @return GetSiteNameExclusiveResponse
func (client *Client) GetSiteNameExclusive(request *GetSiteNameExclusiveRequest) (_result *GetSiteNameExclusiveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSiteNameExclusiveResponse{}
	_body, _err := client.GetSiteNameExclusiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取站点回源客户端证书信息
//
// @param request - GetSiteOriginClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSiteOriginClientCertificateResponse
func (client *Client) GetSiteOriginClientCertificateWithOptions(request *GetSiteOriginClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *GetSiteOriginClientCertificateResponse, _err error) {
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
		Action:      dara.String("GetSiteOriginClientCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSiteOriginClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取站点回源客户端证书信息
//
// @param request - GetSiteOriginClientCertificateRequest
//
// @return GetSiteOriginClientCertificateResponse
func (client *Client) GetSiteOriginClientCertificate(request *GetSiteOriginClientCertificateRequest) (_result *GetSiteOriginClientCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSiteOriginClientCertificateResponse{}
	_body, _err := client.GetSiteOriginClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ESA proxy configuration of a website.
//
// @param request - GetSitePauseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSitePauseResponse
func (client *Client) GetSitePauseWithOptions(request *GetSitePauseRequest, runtime *dara.RuntimeOptions) (_result *GetSitePauseResponse, _err error) {
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
		Action:      dara.String("GetSitePause"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSitePauseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ESA proxy configuration of a website.
//
// @param request - GetSitePauseRequest
//
// @return GetSitePauseResponse
func (client *Client) GetSitePause(request *GetSitePauseRequest) (_result *GetSitePauseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSitePauseResponse{}
	_body, _err := client.GetSitePauseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get WAF Configuration for a Site
//
// @param request - GetSiteWafSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSiteWafSettingsResponse
func (client *Client) GetSiteWafSettingsWithOptions(request *GetSiteWafSettingsRequest, runtime *dara.RuntimeOptions) (_result *GetSiteWafSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSiteWafSettings"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSiteWafSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get WAF Configuration for a Site
//
// @param request - GetSiteWafSettingsRequest
//
// @return GetSiteWafSettingsResponse
func (client *Client) GetSiteWafSettings(request *GetSiteWafSettingsRequest) (_result *GetSiteWafSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSiteWafSettingsResponse{}
	_body, _err := client.GetSiteWafSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Multi-level Cache Configuration for Site
//
// @param request - GetTieredCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTieredCacheResponse
func (client *Client) GetTieredCacheWithOptions(request *GetTieredCacheRequest, runtime *dara.RuntimeOptions) (_result *GetTieredCacheResponse, _err error) {
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
		Action:      dara.String("GetTieredCache"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTieredCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Multi-level Cache Configuration for Site
//
// @param request - GetTieredCacheRequest
//
// @return GetTieredCacheResponse
func (client *Client) GetTieredCache(request *GetTieredCacheRequest) (_result *GetTieredCacheResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTieredCacheResponse{}
	_body, _err := client.GetTieredCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query details of the transport layer application
//
// @param request - GetTransportLayerApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTransportLayerApplicationResponse
func (client *Client) GetTransportLayerApplicationWithOptions(request *GetTransportLayerApplicationRequest, runtime *dara.RuntimeOptions) (_result *GetTransportLayerApplicationResponse, _err error) {
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
		Action:      dara.String("GetTransportLayerApplication"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTransportLayerApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query details of the transport layer application
//
// @param request - GetTransportLayerApplicationRequest
//
// @return GetTransportLayerApplicationResponse
func (client *Client) GetTransportLayerApplication(request *GetTransportLayerApplicationRequest) (_result *GetTransportLayerApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTransportLayerApplicationResponse{}
	_body, _err := client.GetTransportLayerApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution status and running information of a file upload task based on the task ID.
//
// @param request - GetUploadTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUploadTaskResponse
func (client *Client) GetUploadTaskWithOptions(request *GetUploadTaskRequest, runtime *dara.RuntimeOptions) (_result *GetUploadTaskResponse, _err error) {
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
		Action:      dara.String("GetUploadTask"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUploadTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution status and running information of a file upload task based on the task ID.
//
// @param request - GetUploadTaskRequest
//
// @return GetUploadTaskResponse
func (client *Client) GetUploadTask(request *GetUploadTaskRequest) (_result *GetUploadTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUploadTaskResponse{}
	_body, _err := client.GetUploadTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a log delivery task by account.
//
// Description:
//
//	  This API operation queries the details of a delivery task, including the task name, discard rate, region, log category, status, delivery destination, configuration, and filtering rules.****
//
//		- You can call this operation to query detailed information about a log delivery task to analyze log processing efficiency or troubleshoot delivery problems.****
//
//		- ****````
//
// @param request - GetUserDeliveryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserDeliveryTaskResponse
func (client *Client) GetUserDeliveryTaskWithOptions(request *GetUserDeliveryTaskRequest, runtime *dara.RuntimeOptions) (_result *GetUserDeliveryTaskResponse, _err error) {
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
		Action:      dara.String("GetUserDeliveryTask"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a log delivery task by account.
//
// Description:
//
//	  This API operation queries the details of a delivery task, including the task name, discard rate, region, log category, status, delivery destination, configuration, and filtering rules.****
//
//		- You can call this operation to query detailed information about a log delivery task to analyze log processing efficiency or troubleshoot delivery problems.****
//
//		- ****````
//
// @param request - GetUserDeliveryTaskRequest
//
// @return GetUserDeliveryTaskResponse
func (client *Client) GetUserDeliveryTask(request *GetUserDeliveryTaskRequest) (_result *GetUserDeliveryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserDeliveryTaskResponse{}
	_body, _err := client.GetUserDeliveryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the remaining log delivery quota of each log category in your account.
//
// Description:
//
// This operation allows you to query the remaining real-time log delivery quota of each log category in your Alibaba Cloud account. You must provide your Alibaba Cloud account ID (aliUid) and log category (BusinessType). The system then returns the remaining quota of the log category to help you track the usage.
//
// @param request - GetUserLogDeliveryQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserLogDeliveryQuotaResponse
func (client *Client) GetUserLogDeliveryQuotaWithOptions(request *GetUserLogDeliveryQuotaRequest, runtime *dara.RuntimeOptions) (_result *GetUserLogDeliveryQuotaResponse, _err error) {
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
		Action:      dara.String("GetUserLogDeliveryQuota"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserLogDeliveryQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the remaining log delivery quota of each log category in your account.
//
// Description:
//
// This operation allows you to query the remaining real-time log delivery quota of each log category in your Alibaba Cloud account. You must provide your Alibaba Cloud account ID (aliUid) and log category (BusinessType). The system then returns the remaining quota of the log category to help you track the usage.
//
// @param request - GetUserLogDeliveryQuotaRequest
//
// @return GetUserLogDeliveryQuotaResponse
func (client *Client) GetUserLogDeliveryQuota(request *GetUserLogDeliveryQuotaRequest) (_result *GetUserLogDeliveryQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserLogDeliveryQuotaResponse{}
	_body, _err := client.GetUserLogDeliveryQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the video processing configuration details of a site.
//
// @param request - GetVideoProcessingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoProcessingResponse
func (client *Client) GetVideoProcessingWithOptions(request *GetVideoProcessingRequest, runtime *dara.RuntimeOptions) (_result *GetVideoProcessingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoProcessing"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoProcessingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the video processing configuration details of a site.
//
// @param request - GetVideoProcessingRequest
//
// @return GetVideoProcessingResponse
func (client *Client) GetVideoProcessing(request *GetVideoProcessingRequest) (_result *GetVideoProcessingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVideoProcessingResponse{}
	_body, _err := client.GetVideoProcessingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This interface is used to obtain the application key (AppKey) for the BOT behavior detection feature in the site\\"s Web Application Firewall (WAF). The key is typically used for authentication and data exchange with the WAF service.
//
// @param request - GetWafBotAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWafBotAppKeyResponse
func (client *Client) GetWafBotAppKeyWithOptions(runtime *dara.RuntimeOptions) (_result *GetWafBotAppKeyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetWafBotAppKey"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWafBotAppKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to obtain the application key (AppKey) for the BOT behavior detection feature in the site\\"s Web Application Firewall (WAF). The key is typically used for authentication and data exchange with the WAF service.
//
// @return GetWafBotAppKeyResponse
func (client *Client) GetWafBotAppKey() (_result *GetWafBotAppKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWafBotAppKeyResponse{}
	_body, _err := client.GetWafBotAppKeyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the conditions for matching incoming requests that are configured in a WAF rule category for a website. These conditions define how WAF detects and processes different types of requests.
//
// @param request - GetWafFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWafFilterResponse
func (client *Client) GetWafFilterWithOptions(request *GetWafFilterRequest, runtime *dara.RuntimeOptions) (_result *GetWafFilterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Phase) {
		query["Phase"] = request.Phase
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWafFilter"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWafFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the conditions for matching incoming requests that are configured in a WAF rule category for a website. These conditions define how WAF detects and processes different types of requests.
//
// @param request - GetWafFilterRequest
//
// @return GetWafFilterResponse
func (client *Client) GetWafFilter(request *GetWafFilterRequest) (_result *GetWafFilterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWafFilterResponse{}
	_body, _err := client.GetWafFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get WAF Quota Details
//
// @param request - GetWafQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWafQuotaResponse
func (client *Client) GetWafQuotaWithOptions(request *GetWafQuotaRequest, runtime *dara.RuntimeOptions) (_result *GetWafQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Paths) {
		query["Paths"] = request.Paths
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWafQuota"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWafQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get WAF Quota Details
//
// @param request - GetWafQuotaRequest
//
// @return GetWafQuotaResponse
func (client *Client) GetWafQuota(request *GetWafQuotaRequest) (_result *GetWafQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWafQuotaResponse{}
	_body, _err := client.GetWafQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Details of a Single WAF Rule
//
// @param request - GetWafRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWafRuleResponse
func (client *Client) GetWafRuleWithOptions(request *GetWafRuleRequest, runtime *dara.RuntimeOptions) (_result *GetWafRuleResponse, _err error) {
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

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWafRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWafRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Details of a Single WAF Rule
//
// @param request - GetWafRuleRequest
//
// @return GetWafRuleResponse
func (client *Client) GetWafRule(request *GetWafRuleRequest) (_result *GetWafRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWafRuleResponse{}
	_body, _err := client.GetWafRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get WAF Ruleset Details
//
// @param request - GetWafRulesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWafRulesetResponse
func (client *Client) GetWafRulesetWithOptions(request *GetWafRulesetRequest, runtime *dara.RuntimeOptions) (_result *GetWafRulesetResponse, _err error) {
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

	if !dara.IsNil(request.Phase) {
		query["Phase"] = request.Phase
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWafRuleset"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWafRulesetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get WAF Ruleset Details
//
// @param request - GetWafRulesetRequest
//
// @return GetWafRulesetResponse
func (client *Client) GetWafRuleset(request *GetWafRulesetRequest) (_result *GetWafRulesetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWafRulesetResponse{}
	_body, _err := client.GetWafRulesetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Cache Reserve Instance List
//
// @param request - ListCacheReserveInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCacheReserveInstancesResponse
func (client *Client) ListCacheReserveInstancesWithOptions(request *ListCacheReserveInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListCacheReserveInstancesResponse, _err error) {
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
		Action:      dara.String("ListCacheReserveInstances"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCacheReserveInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Cache Reserve Instance List
//
// @param request - ListCacheReserveInstancesRequest
//
// @return ListCacheReserveInstancesResponse
func (client *Client) ListCacheReserveInstances(request *ListCacheReserveInstancesRequest) (_result *ListCacheReserveInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCacheReserveInstancesResponse{}
	_body, _err := client.ListCacheReserveInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query multiple cache configurations
//
// @param request - ListCacheRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCacheRulesResponse
func (client *Client) ListCacheRulesWithOptions(request *ListCacheRulesRequest, runtime *dara.RuntimeOptions) (_result *ListCacheRulesResponse, _err error) {
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
		Action:      dara.String("ListCacheRules"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCacheRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query multiple cache configurations
//
// @param request - ListCacheRulesRequest
//
// @return ListCacheRulesResponse
func (client *Client) ListCacheRules(request *ListCacheRulesRequest) (_result *ListCacheRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCacheRulesResponse{}
	_body, _err := client.ListCacheRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists certificates of a website.
//
// @param request - ListCertificatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCertificatesResponse
func (client *Client) ListCertificatesWithOptions(request *ListCertificatesRequest, runtime *dara.RuntimeOptions) (_result *ListCertificatesResponse, _err error) {
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
		Action:      dara.String("ListCertificates"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists certificates of a website.
//
// @param request - ListCertificatesRequest
//
// @return ListCertificatesResponse
func (client *Client) ListCertificates(request *ListCertificatesRequest) (_result *ListCertificatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCertificatesResponse{}
	_body, _err := client.ListCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists certificates that match specified records for a website. You can specify multiple records at a time.
//
// @param request - ListCertificatesByRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCertificatesByRecordResponse
func (client *Client) ListCertificatesByRecordWithOptions(request *ListCertificatesByRecordRequest, runtime *dara.RuntimeOptions) (_result *ListCertificatesByRecordResponse, _err error) {
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
		Action:      dara.String("ListCertificatesByRecord"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCertificatesByRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists certificates that match specified records for a website. You can specify multiple records at a time.
//
// @param request - ListCertificatesByRecordRequest
//
// @return ListCertificatesByRecordResponse
func (client *Client) ListCertificatesByRecord(request *ListCertificatesByRecordRequest) (_result *ListCertificatesByRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCertificatesByRecordResponse{}
	_body, _err := client.ListCertificatesByRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query TLS Cipher Suite List
//
// @param request - ListCiphersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCiphersResponse
func (client *Client) ListCiphersWithOptions(request *ListCiphersRequest, runtime *dara.RuntimeOptions) (_result *ListCiphersResponse, _err error) {
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
		Action:      dara.String("ListCiphers"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCiphersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query TLS Cipher Suite List
//
// @param request - ListCiphersRequest
//
// @return ListCiphersResponse
func (client *Client) ListCiphers(request *ListCiphersRequest) (_result *ListCiphersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCiphersResponse{}
	_body, _err := client.ListCiphersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of client certificate authority (CA) certificates for a website.
//
// @param request - ListClientCaCertificatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClientCaCertificatesResponse
func (client *Client) ListClientCaCertificatesWithOptions(request *ListClientCaCertificatesRequest, runtime *dara.RuntimeOptions) (_result *ListClientCaCertificatesResponse, _err error) {
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
		Action:      dara.String("ListClientCaCertificates"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClientCaCertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of client certificate authority (CA) certificates for a website.
//
// @param request - ListClientCaCertificatesRequest
//
// @return ListClientCaCertificatesResponse
func (client *Client) ListClientCaCertificates(request *ListClientCaCertificatesRequest) (_result *ListClientCaCertificatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClientCaCertificatesResponse{}
	_body, _err := client.ListClientCaCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries client certificates configured for a website.
//
// @param request - ListClientCertificatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClientCertificatesResponse
func (client *Client) ListClientCertificatesWithOptions(request *ListClientCertificatesRequest, runtime *dara.RuntimeOptions) (_result *ListClientCertificatesResponse, _err error) {
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
		Action:      dara.String("ListClientCertificates"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClientCertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries client certificates configured for a website.
//
// @param request - ListClientCertificatesRequest
//
// @return ListClientCertificatesResponse
func (client *Client) ListClientCertificates(request *ListClientCertificatesRequest) (_result *ListClientCertificatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClientCertificatesResponse{}
	_body, _err := client.ListClientCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the list of compression rules
//
// @param request - ListCompressionRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCompressionRulesResponse
func (client *Client) ListCompressionRulesWithOptions(request *ListCompressionRulesRequest, runtime *dara.RuntimeOptions) (_result *ListCompressionRulesResponse, _err error) {
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
		Action:      dara.String("ListCompressionRules"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCompressionRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list of compression rules
//
// @param request - ListCompressionRulesRequest
//
// @return ListCompressionRulesResponse
func (client *Client) ListCompressionRules(request *ListCompressionRulesRequest) (_result *ListCompressionRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCompressionRulesResponse{}
	_body, _err := client.ListCompressionRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch query whether the IP address is included in the ESA resolution result.
//
// Description:
//
// This interface is used to check whether the vs_addr parameter in the vipInfo collection is vip.
//
// @param request - ListESAIPInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListESAIPInfoResponse
func (client *Client) ListESAIPInfoWithOptions(request *ListESAIPInfoRequest, runtime *dara.RuntimeOptions) (_result *ListESAIPInfoResponse, _err error) {
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
		Action:      dara.String("ListESAIPInfo"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListESAIPInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch query whether the IP address is included in the ESA resolution result.
//
// Description:
//
// This interface is used to check whether the vs_addr parameter in the vipInfo collection is vip.
//
// @param request - ListESAIPInfoRequest
//
// @return ListESAIPInfoResponse
func (client *Client) ListESAIPInfo(request *ListESAIPInfoRequest) (_result *ListESAIPInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListESAIPInfoResponse{}
	_body, _err := client.ListESAIPInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists domain names that are associated with a containerized application.
//
// @param request - ListEdgeContainerAppRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEdgeContainerAppRecordsResponse
func (client *Client) ListEdgeContainerAppRecordsWithOptions(request *ListEdgeContainerAppRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListEdgeContainerAppRecordsResponse, _err error) {
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
		Action:      dara.String("ListEdgeContainerAppRecords"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEdgeContainerAppRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists domain names that are associated with a containerized application.
//
// @param request - ListEdgeContainerAppRecordsRequest
//
// @return ListEdgeContainerAppRecordsResponse
func (client *Client) ListEdgeContainerAppRecords(request *ListEdgeContainerAppRecordsRequest) (_result *ListEdgeContainerAppRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEdgeContainerAppRecordsResponse{}
	_body, _err := client.ListEdgeContainerAppRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists versions of all containerized applications.
//
// @param request - ListEdgeContainerAppVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEdgeContainerAppVersionsResponse
func (client *Client) ListEdgeContainerAppVersionsWithOptions(request *ListEdgeContainerAppVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListEdgeContainerAppVersionsResponse, _err error) {
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
		Action:      dara.String("ListEdgeContainerAppVersions"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEdgeContainerAppVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists versions of all containerized applications.
//
// @param request - ListEdgeContainerAppVersionsRequest
//
// @return ListEdgeContainerAppVersionsResponse
func (client *Client) ListEdgeContainerAppVersions(request *ListEdgeContainerAppVersionsRequest) (_result *ListEdgeContainerAppVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEdgeContainerAppVersionsResponse{}
	_body, _err := client.ListEdgeContainerAppVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all containerized applications in your Alibaba Cloud account.
//
// @param request - ListEdgeContainerAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEdgeContainerAppsResponse
func (client *Client) ListEdgeContainerAppsWithOptions(request *ListEdgeContainerAppsRequest, runtime *dara.RuntimeOptions) (_result *ListEdgeContainerAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderKey) {
		query["OrderKey"] = request.OrderKey
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.SearchType) {
		query["SearchType"] = request.SearchType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEdgeContainerApps"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEdgeContainerAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all containerized applications in your Alibaba Cloud account.
//
// @param request - ListEdgeContainerAppsRequest
//
// @return ListEdgeContainerAppsResponse
func (client *Client) ListEdgeContainerApps(request *ListEdgeContainerAppsRequest) (_result *ListEdgeContainerAppsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEdgeContainerAppsResponse{}
	_body, _err := client.ListEdgeContainerAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the records that are associated with Edge Container for a website.
//
// @param request - ListEdgeContainerRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEdgeContainerRecordsResponse
func (client *Client) ListEdgeContainerRecordsWithOptions(request *ListEdgeContainerRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListEdgeContainerRecordsResponse, _err error) {
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
		Action:      dara.String("ListEdgeContainerRecords"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEdgeContainerRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the records that are associated with Edge Container for a website.
//
// @param request - ListEdgeContainerRecordsRequest
//
// @return ListEdgeContainerRecordsResponse
func (client *Client) ListEdgeContainerRecords(request *ListEdgeContainerRecordsRequest) (_result *ListEdgeContainerRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEdgeContainerRecordsResponse{}
	_body, _err := client.ListEdgeContainerRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Edge Routine plans.
//
// @param request - ListEdgeRoutinePlansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEdgeRoutinePlansResponse
func (client *Client) ListEdgeRoutinePlansWithOptions(runtime *dara.RuntimeOptions) (_result *ListEdgeRoutinePlansResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListEdgeRoutinePlans"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEdgeRoutinePlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Edge Routine plans.
//
// @return ListEdgeRoutinePlansResponse
func (client *Client) ListEdgeRoutinePlans() (_result *ListEdgeRoutinePlansResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEdgeRoutinePlansResponse{}
	_body, _err := client.ListEdgeRoutinePlansWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the records that are associated with Edge Routine routes for a website.
//
// Description:
//
// >  You can call this operation 100 times per second.
//
// @param request - ListEdgeRoutineRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEdgeRoutineRecordsResponse
func (client *Client) ListEdgeRoutineRecordsWithOptions(request *ListEdgeRoutineRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListEdgeRoutineRecordsResponse, _err error) {
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
		Action:      dara.String("ListEdgeRoutineRecords"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEdgeRoutineRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the records that are associated with Edge Routine routes for a website.
//
// Description:
//
// >  You can call this operation 100 times per second.
//
// @param request - ListEdgeRoutineRecordsRequest
//
// @return ListEdgeRoutineRecordsResponse
func (client *Client) ListEdgeRoutineRecords(request *ListEdgeRoutineRecordsRequest) (_result *ListEdgeRoutineRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEdgeRoutineRecordsResponse{}
	_body, _err := client.ListEdgeRoutineRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration details of an incoming HTTP request header modification rule for a website.
//
// @param request - ListHttpIncomingRequestHeaderModificationRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHttpIncomingRequestHeaderModificationRulesResponse
func (client *Client) ListHttpIncomingRequestHeaderModificationRulesWithOptions(request *ListHttpIncomingRequestHeaderModificationRulesRequest, runtime *dara.RuntimeOptions) (_result *ListHttpIncomingRequestHeaderModificationRulesResponse, _err error) {
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
		Action:      dara.String("ListHttpIncomingRequestHeaderModificationRules"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHttpIncomingRequestHeaderModificationRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration details of an incoming HTTP request header modification rule for a website.
//
// @param request - ListHttpIncomingRequestHeaderModificationRulesRequest
//
// @return ListHttpIncomingRequestHeaderModificationRulesResponse
func (client *Client) ListHttpIncomingRequestHeaderModificationRules(request *ListHttpIncomingRequestHeaderModificationRulesRequest) (_result *ListHttpIncomingRequestHeaderModificationRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHttpIncomingRequestHeaderModificationRulesResponse{}
	_body, _err := client.ListHttpIncomingRequestHeaderModificationRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of an incoming HTTP response header modification rule for a website.
//
// @param request - ListHttpIncomingResponseHeaderModificationRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHttpIncomingResponseHeaderModificationRulesResponse
func (client *Client) ListHttpIncomingResponseHeaderModificationRulesWithOptions(request *ListHttpIncomingResponseHeaderModificationRulesRequest, runtime *dara.RuntimeOptions) (_result *ListHttpIncomingResponseHeaderModificationRulesResponse, _err error) {
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
		Action:      dara.String("ListHttpIncomingResponseHeaderModificationRules"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHttpIncomingResponseHeaderModificationRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of an incoming HTTP response header modification rule for a website.
//
// @param request - ListHttpIncomingResponseHeaderModificationRulesRequest
//
// @return ListHttpIncomingResponseHeaderModificationRulesResponse
func (client *Client) ListHttpIncomingResponseHeaderModificationRules(request *ListHttpIncomingResponseHeaderModificationRulesRequest) (_result *ListHttpIncomingResponseHeaderModificationRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHttpIncomingResponseHeaderModificationRulesResponse{}
	_body, _err := client.ListHttpIncomingResponseHeaderModificationRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List of HTTP Request Header Rules
//
// @param request - ListHttpRequestHeaderModificationRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHttpRequestHeaderModificationRulesResponse
func (client *Client) ListHttpRequestHeaderModificationRulesWithOptions(request *ListHttpRequestHeaderModificationRulesRequest, runtime *dara.RuntimeOptions) (_result *ListHttpRequestHeaderModificationRulesResponse, _err error) {
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
		Action:      dara.String("ListHttpRequestHeaderModificationRules"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHttpRequestHeaderModificationRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List of HTTP Request Header Rules
//
// @param request - ListHttpRequestHeaderModificationRulesRequest
//
// @return ListHttpRequestHeaderModificationRulesResponse
func (client *Client) ListHttpRequestHeaderModificationRules(request *ListHttpRequestHeaderModificationRulesRequest) (_result *ListHttpRequestHeaderModificationRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHttpRequestHeaderModificationRulesResponse{}
	_body, _err := client.ListHttpRequestHeaderModificationRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List of HTTP Response Header Rules
//
// @param request - ListHttpResponseHeaderModificationRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHttpResponseHeaderModificationRulesResponse
func (client *Client) ListHttpResponseHeaderModificationRulesWithOptions(request *ListHttpResponseHeaderModificationRulesRequest, runtime *dara.RuntimeOptions) (_result *ListHttpResponseHeaderModificationRulesResponse, _err error) {
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
		Action:      dara.String("ListHttpResponseHeaderModificationRules"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHttpResponseHeaderModificationRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List of HTTP Response Header Rules
//
// @param request - ListHttpResponseHeaderModificationRulesRequest
//
// @return ListHttpResponseHeaderModificationRulesResponse
func (client *Client) ListHttpResponseHeaderModificationRules(request *ListHttpResponseHeaderModificationRulesRequest) (_result *ListHttpResponseHeaderModificationRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHttpResponseHeaderModificationRulesResponse{}
	_body, _err := client.ListHttpResponseHeaderModificationRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query multiple HTTPS application configurations
//
// @param request - ListHttpsApplicationConfigurationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHttpsApplicationConfigurationsResponse
func (client *Client) ListHttpsApplicationConfigurationsWithOptions(request *ListHttpsApplicationConfigurationsRequest, runtime *dara.RuntimeOptions) (_result *ListHttpsApplicationConfigurationsResponse, _err error) {
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
		Action:      dara.String("ListHttpsApplicationConfigurations"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHttpsApplicationConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query multiple HTTPS application configurations
//
// @param request - ListHttpsApplicationConfigurationsRequest
//
// @return ListHttpsApplicationConfigurationsResponse
func (client *Client) ListHttpsApplicationConfigurations(request *ListHttpsApplicationConfigurationsRequest) (_result *ListHttpsApplicationConfigurationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHttpsApplicationConfigurationsResponse{}
	_body, _err := client.ListHttpsApplicationConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query multiple HTTPS basic configurations
//
// @param request - ListHttpsBasicConfigurationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHttpsBasicConfigurationsResponse
func (client *Client) ListHttpsBasicConfigurationsWithOptions(request *ListHttpsBasicConfigurationsRequest, runtime *dara.RuntimeOptions) (_result *ListHttpsBasicConfigurationsResponse, _err error) {
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
		Action:      dara.String("ListHttpsBasicConfigurations"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHttpsBasicConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query multiple HTTPS basic configurations
//
// @param request - ListHttpsBasicConfigurationsRequest
//
// @return ListHttpsBasicConfigurationsResponse
func (client *Client) ListHttpsBasicConfigurations(request *ListHttpsBasicConfigurationsRequest) (_result *ListHttpsBasicConfigurationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHttpsBasicConfigurationsResponse{}
	_body, _err := client.ListHttpsBasicConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Multiple Site Image Transformation Configurations
//
// @param request - ListImageTransformsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImageTransformsResponse
func (client *Client) ListImageTransformsWithOptions(request *ListImageTransformsRequest, runtime *dara.RuntimeOptions) (_result *ListImageTransformsResponse, _err error) {
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
		Action:      dara.String("ListImageTransforms"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImageTransformsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Multiple Site Image Transformation Configurations
//
// @param request - ListImageTransformsRequest
//
// @return ListImageTransformsResponse
func (client *Client) ListImageTransforms(request *ListImageTransformsRequest) (_result *ListImageTransformsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListImageTransformsResponse{}
	_body, _err := client.ListImageTransformsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the quota details in a subscription plan.
//
// @param request - ListInstanceQuotasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceQuotasResponse
func (client *Client) ListInstanceQuotasWithOptions(request *ListInstanceQuotasRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceQuotasResponse, _err error) {
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
		Action:      dara.String("ListInstanceQuotas"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceQuotasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quota details in a subscription plan.
//
// @param request - ListInstanceQuotasRequest
//
// @return ListInstanceQuotasResponse
func (client *Client) ListInstanceQuotas(request *ListInstanceQuotasRequest) (_result *ListInstanceQuotasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceQuotasResponse{}
	_body, _err := client.ListInstanceQuotasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries quotas and the actual usage in a plan based on the website or plan ID.
//
// @param request - ListInstanceQuotasWithUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceQuotasWithUsageResponse
func (client *Client) ListInstanceQuotasWithUsageWithOptions(request *ListInstanceQuotasWithUsageRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceQuotasWithUsageResponse, _err error) {
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
		Action:      dara.String("ListInstanceQuotasWithUsage"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceQuotasWithUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quotas and the actual usage in a plan based on the website or plan ID.
//
// @param request - ListInstanceQuotasWithUsageRequest
//
// @return ListInstanceQuotasWithUsageResponse
func (client *Client) ListInstanceQuotasWithUsage(request *ListInstanceQuotasWithUsageRequest) (_result *ListInstanceQuotasWithUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceQuotasWithUsageResponse{}
	_body, _err := client.ListInstanceQuotasWithUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all key-value pairs in a namespace in your Alibaba Cloud account.
//
// @param request - ListKvsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKvsResponse
func (client *Client) ListKvsWithOptions(request *ListKvsRequest, runtime *dara.RuntimeOptions) (_result *ListKvsResponse, _err error) {
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
		Action:      dara.String("ListKvs"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKvsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all key-value pairs in a namespace in your Alibaba Cloud account.
//
// @param request - ListKvsRequest
//
// @return ListKvsResponse
func (client *Client) ListKvs(request *ListKvsRequest) (_result *ListKvsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListKvsResponse{}
	_body, _err := client.ListKvsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all custom lists and their details in an Alibaba Cloud account. You can specify query arguments to filter the results and display the returned lists by page.
//
// @param tmpReq - ListListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListListsResponse
func (client *Client) ListListsWithOptions(tmpReq *ListListsRequest, runtime *dara.RuntimeOptions) (_result *ListListsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListListsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.QueryArgs) {
		request.QueryArgsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryArgs, dara.String("QueryArgs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryArgsShrink) {
		query["QueryArgs"] = request.QueryArgsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLists"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListListsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all custom lists and their details in an Alibaba Cloud account. You can specify query arguments to filter the results and display the returned lists by page.
//
// @param request - ListListsRequest
//
// @return ListListsResponse
func (client *Client) ListLists(request *ListListsRequest) (_result *ListListsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListListsResponse{}
	_body, _err := client.ListListsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the status of origins in load balancers
//
// Description:
//
// Query the status of origins under load balancers. You can pass multiple load balancer IDs at once, separated by commas. This is for load balancers that have monitors configured. It will probe the origins in the source address pools used by the load balancers and record the current status of each origin.
//
// - Healthy(healthy): The probe result is available.
//
// - Unhealthy(unhealthy): The probe result is unavailable.
//
// - Unknown(unknown): Unknown, the monitor has not yet probed.
//
// - Undetected(undetected): The load balancer to which the origin belongs is not bound to a monitor.
//
// @param request - ListLoadBalancerOriginStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLoadBalancerOriginStatusResponse
func (client *Client) ListLoadBalancerOriginStatusWithOptions(request *ListLoadBalancerOriginStatusRequest, runtime *dara.RuntimeOptions) (_result *ListLoadBalancerOriginStatusResponse, _err error) {
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
		Action:      dara.String("ListLoadBalancerOriginStatus"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLoadBalancerOriginStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the status of origins in load balancers
//
// Description:
//
// Query the status of origins under load balancers. You can pass multiple load balancer IDs at once, separated by commas. This is for load balancers that have monitors configured. It will probe the origins in the source address pools used by the load balancers and record the current status of each origin.
//
// - Healthy(healthy): The probe result is available.
//
// - Unhealthy(unhealthy): The probe result is unavailable.
//
// - Unknown(unknown): Unknown, the monitor has not yet probed.
//
// - Undetected(undetected): The load balancer to which the origin belongs is not bound to a monitor.
//
// @param request - ListLoadBalancerOriginStatusRequest
//
// @return ListLoadBalancerOriginStatusResponse
func (client *Client) ListLoadBalancerOriginStatus(request *ListLoadBalancerOriginStatusRequest) (_result *ListLoadBalancerOriginStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLoadBalancerOriginStatusResponse{}
	_body, _err := client.ListLoadBalancerOriginStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Load Balancer Region List
//
// Description:
//
// When creating a load balancer \\"based on country/region scheduling\\" strategy through OpenAPI, use the code of primary or secondary regions to represent traffic from this geographical area.
//
// @param request - ListLoadBalancerRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLoadBalancerRegionsResponse
func (client *Client) ListLoadBalancerRegionsWithOptions(request *ListLoadBalancerRegionsRequest, runtime *dara.RuntimeOptions) (_result *ListLoadBalancerRegionsResponse, _err error) {
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
		Action:      dara.String("ListLoadBalancerRegions"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLoadBalancerRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Load Balancer Region List
//
// Description:
//
// When creating a load balancer \\"based on country/region scheduling\\" strategy through OpenAPI, use the code of primary or secondary regions to represent traffic from this geographical area.
//
// @param request - ListLoadBalancerRegionsRequest
//
// @return ListLoadBalancerRegionsResponse
func (client *Client) ListLoadBalancerRegions(request *ListLoadBalancerRegionsRequest) (_result *ListLoadBalancerRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLoadBalancerRegionsResponse{}
	_body, _err := client.ListLoadBalancerRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the list of load balancers
//
// @param request - ListLoadBalancersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLoadBalancersResponse
func (client *Client) ListLoadBalancersWithOptions(request *ListLoadBalancersRequest, runtime *dara.RuntimeOptions) (_result *ListLoadBalancersResponse, _err error) {
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
		Action:      dara.String("ListLoadBalancers"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLoadBalancersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list of load balancers
//
// @param request - ListLoadBalancersRequest
//
// @return ListLoadBalancersResponse
func (client *Client) ListLoadBalancers(request *ListLoadBalancersRequest) (_result *ListLoadBalancersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLoadBalancersResponse{}
	_body, _err := client.ListLoadBalancersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List Custom Managed Rule Groups
//
// @param request - ListManagedRulesGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListManagedRulesGroupsResponse
func (client *Client) ListManagedRulesGroupsWithOptions(request *ListManagedRulesGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListManagedRulesGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("ListManagedRulesGroups"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListManagedRulesGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List Custom Managed Rule Groups
//
// @param request - ListManagedRulesGroupsRequest
//
// @return ListManagedRulesGroupsResponse
func (client *Client) ListManagedRulesGroups(request *ListManagedRulesGroupsRequest) (_result *ListManagedRulesGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListManagedRulesGroupsResponse{}
	_body, _err := client.ListManagedRulesGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query multiple network optimization configurations
//
// @param request - ListNetworkOptimizationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkOptimizationsResponse
func (client *Client) ListNetworkOptimizationsWithOptions(request *ListNetworkOptimizationsRequest, runtime *dara.RuntimeOptions) (_result *ListNetworkOptimizationsResponse, _err error) {
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
		Action:      dara.String("ListNetworkOptimizations"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNetworkOptimizationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query multiple network optimization configurations
//
// @param request - ListNetworkOptimizationsRequest
//
// @return ListNetworkOptimizationsResponse
func (client *Client) ListNetworkOptimizations(request *ListNetworkOptimizationsRequest) (_result *ListNetworkOptimizationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNetworkOptimizationsResponse{}
	_body, _err := client.ListNetworkOptimizationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询源服务器CA证书列表
//
// @param request - ListOriginCaCertificatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOriginCaCertificatesResponse
func (client *Client) ListOriginCaCertificatesWithOptions(request *ListOriginCaCertificatesRequest, runtime *dara.RuntimeOptions) (_result *ListOriginCaCertificatesResponse, _err error) {
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
		Action:      dara.String("ListOriginCaCertificates"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOriginCaCertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询源服务器CA证书列表
//
// @param request - ListOriginCaCertificatesRequest
//
// @return ListOriginCaCertificatesResponse
func (client *Client) ListOriginCaCertificates(request *ListOriginCaCertificatesRequest) (_result *ListOriginCaCertificatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOriginCaCertificatesResponse{}
	_body, _err := client.ListOriginCaCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询域名回源客户端证书列表
//
// @param request - ListOriginClientCertificatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOriginClientCertificatesResponse
func (client *Client) ListOriginClientCertificatesWithOptions(request *ListOriginClientCertificatesRequest, runtime *dara.RuntimeOptions) (_result *ListOriginClientCertificatesResponse, _err error) {
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
		Action:      dara.String("ListOriginClientCertificates"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOriginClientCertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询域名回源客户端证书列表
//
// @param request - ListOriginClientCertificatesRequest
//
// @return ListOriginClientCertificatesResponse
func (client *Client) ListOriginClientCertificates(request *ListOriginClientCertificatesRequest) (_result *ListOriginClientCertificatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOriginClientCertificatesResponse{}
	_body, _err := client.ListOriginClientCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List Origin Pools
//
// @param request - ListOriginPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOriginPoolsResponse
func (client *Client) ListOriginPoolsWithOptions(request *ListOriginPoolsRequest, runtime *dara.RuntimeOptions) (_result *ListOriginPoolsResponse, _err error) {
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
		Action:      dara.String("ListOriginPools"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOriginPoolsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List Origin Pools
//
// @param request - ListOriginPoolsRequest
//
// @return ListOriginPoolsResponse
func (client *Client) ListOriginPools(request *ListOriginPoolsRequest) (_result *ListOriginPoolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOriginPoolsResponse{}
	_body, _err := client.ListOriginPoolsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query multiple origin rule configurations
//
// @param request - ListOriginRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOriginRulesResponse
func (client *Client) ListOriginRulesWithOptions(request *ListOriginRulesRequest, runtime *dara.RuntimeOptions) (_result *ListOriginRulesResponse, _err error) {
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
		Action:      dara.String("ListOriginRules"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOriginRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query multiple origin rule configurations
//
// @param request - ListOriginRulesRequest
//
// @return ListOriginRulesResponse
func (client *Client) ListOriginRules(request *ListOriginRulesRequest) (_result *ListOriginRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOriginRulesResponse{}
	_body, _err := client.ListOriginRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all custom error pages that you created. You can define the page number and the number of entries per page to display the response.
//
// @param tmpReq - ListPagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPagesResponse
func (client *Client) ListPagesWithOptions(tmpReq *ListPagesRequest, runtime *dara.RuntimeOptions) (_result *ListPagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.QueryArgs) {
		request.QueryArgsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryArgs, dara.String("QueryArgs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryArgsShrink) {
		query["QueryArgs"] = request.QueryArgsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPages"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all custom error pages that you created. You can define the page number and the number of entries per page to display the response.
//
// @param request - ListPagesRequest
//
// @return ListPagesResponse
func (client *Client) ListPages(request *ListPagesRequest) (_result *ListPagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPagesResponse{}
	_body, _err := client.ListPagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of Domain Name System (DNS) records of a website, including the record value, priority, and authentication configurations. Supports filtering by specifying parameters such as RecordName and RecordMatchType.
//
// Description:
//
// The DNS records related to Edge Container, Edge Routine, and TCP/UDP proxy are not returned in this operation.
//
// @param request - ListRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecordsResponse
func (client *Client) ListRecordsWithOptions(request *ListRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListRecordsResponse, _err error) {
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
		Action:      dara.String("ListRecords"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Domain Name System (DNS) records of a website, including the record value, priority, and authentication configurations. Supports filtering by specifying parameters such as RecordName and RecordMatchType.
//
// Description:
//
// The DNS records related to Edge Container, Edge Routine, and TCP/UDP proxy are not returned in this operation.
//
// @param request - ListRecordsRequest
//
// @return ListRecordsResponse
func (client *Client) ListRecords(request *ListRecordsRequest) (_result *ListRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRecordsResponse{}
	_body, _err := client.ListRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Redirect Rule List
//
// @param request - ListRedirectRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRedirectRulesResponse
func (client *Client) ListRedirectRulesWithOptions(request *ListRedirectRulesRequest, runtime *dara.RuntimeOptions) (_result *ListRedirectRulesResponse, _err error) {
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
		Action:      dara.String("ListRedirectRules"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRedirectRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Redirect Rule List
//
// @param request - ListRedirectRulesRequest
//
// @return ListRedirectRulesResponse
func (client *Client) ListRedirectRules(request *ListRedirectRulesRequest) (_result *ListRedirectRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRedirectRulesResponse{}
	_body, _err := client.ListRedirectRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List of Rewrite URL Rules
//
// @param request - ListRewriteUrlRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRewriteUrlRulesResponse
func (client *Client) ListRewriteUrlRulesWithOptions(request *ListRewriteUrlRulesRequest, runtime *dara.RuntimeOptions) (_result *ListRewriteUrlRulesResponse, _err error) {
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
		Action:      dara.String("ListRewriteUrlRules"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRewriteUrlRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List of Rewrite URL Rules
//
// @param request - ListRewriteUrlRulesRequest
//
// @return ListRewriteUrlRulesResponse
func (client *Client) ListRewriteUrlRules(request *ListRewriteUrlRulesRequest) (_result *ListRewriteUrlRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRewriteUrlRulesResponse{}
	_body, _err := client.ListRewriteUrlRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the regions to which Edge Routine code can be released for canary deployment.
//
// @param request - ListRoutineCanaryAreasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRoutineCanaryAreasResponse
func (client *Client) ListRoutineCanaryAreasWithOptions(runtime *dara.RuntimeOptions) (_result *ListRoutineCanaryAreasResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListRoutineCanaryAreas"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRoutineCanaryAreasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the regions to which Edge Routine code can be released for canary deployment.
//
// @return ListRoutineCanaryAreasResponse
func (client *Client) ListRoutineCanaryAreas() (_result *ListRoutineCanaryAreasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRoutineCanaryAreasResponse{}
	_body, _err := client.ListRoutineCanaryAreasWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the code versions of a function (routine) by page.
//
// Description:
//
// Call this operation to query the code versions of a specific function. Paged query and fuzzy search are supported. You can configure `Name` to specify the name of a function.
//
// Specify `PageNumber` and `PageSize` to control the number of entries returned in a request, and use `SearchKeyWord` to specify a keyword for fuzzy search.
//
// The response includes the number, description, and creation time of each code version.
//
// @param request - ListRoutineCodeVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRoutineCodeVersionsResponse
func (client *Client) ListRoutineCodeVersionsWithOptions(request *ListRoutineCodeVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListRoutineCodeVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKeyWord) {
		body["SearchKeyWord"] = request.SearchKeyWord
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRoutineCodeVersions"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRoutineCodeVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the code versions of a function (routine) by page.
//
// Description:
//
// Call this operation to query the code versions of a specific function. Paged query and fuzzy search are supported. You can configure `Name` to specify the name of a function.
//
// Specify `PageNumber` and `PageSize` to control the number of entries returned in a request, and use `SearchKeyWord` to specify a keyword for fuzzy search.
//
// The response includes the number, description, and creation time of each code version.
//
// @param request - ListRoutineCodeVersionsRequest
//
// @return ListRoutineCodeVersionsResponse
func (client *Client) ListRoutineCodeVersions(request *ListRoutineCodeVersionsRequest) (_result *ListRoutineCodeVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRoutineCodeVersionsResponse{}
	_body, _err := client.ListRoutineCodeVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The records associated with the function.
//
// Description:
//
// You can call this operation to query the routes associated with a function. You can specify paged query parameters to obtain the specified number of routes or specify a keyword for fuzzy search to filter specific routes.
//
// @param request - ListRoutineRelatedRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRoutineRelatedRecordsResponse
func (client *Client) ListRoutineRelatedRecordsWithOptions(request *ListRoutineRelatedRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListRoutineRelatedRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKeyWord) {
		body["SearchKeyWord"] = request.SearchKeyWord
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRoutineRelatedRecords"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRoutineRelatedRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The records associated with the function.
//
// Description:
//
// You can call this operation to query the routes associated with a function. You can specify paged query parameters to obtain the specified number of routes or specify a keyword for fuzzy search to filter specific routes.
//
// @param request - ListRoutineRelatedRecordsRequest
//
// @return ListRoutineRelatedRecordsResponse
func (client *Client) ListRoutineRelatedRecords(request *ListRoutineRelatedRecordsRequest) (_result *ListRoutineRelatedRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRoutineRelatedRecordsResponse{}
	_body, _err := client.ListRoutineRelatedRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the routes of an edge function.
//
// @param request - ListRoutineRoutesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRoutineRoutesResponse
func (client *Client) ListRoutineRoutesWithOptions(request *ListRoutineRoutesRequest, runtime *dara.RuntimeOptions) (_result *ListRoutineRoutesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RoutineName) {
		query["RoutineName"] = request.RoutineName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRoutineRoutes"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRoutineRoutesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the routes of an edge function.
//
// @param request - ListRoutineRoutesRequest
//
// @return ListRoutineRoutesResponse
func (client *Client) ListRoutineRoutes(request *ListRoutineRoutesRequest) (_result *ListRoutineRoutesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRoutineRoutesResponse{}
	_body, _err := client.ListRoutineRoutesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the plans in a scheduled prefetch task by task ID.
//
// @param request - ListScheduledPreloadExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduledPreloadExecutionsResponse
func (client *Client) ListScheduledPreloadExecutionsWithOptions(request *ListScheduledPreloadExecutionsRequest, runtime *dara.RuntimeOptions) (_result *ListScheduledPreloadExecutionsResponse, _err error) {
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
		Action:      dara.String("ListScheduledPreloadExecutions"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScheduledPreloadExecutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the plans in a scheduled prefetch task by task ID.
//
// @param request - ListScheduledPreloadExecutionsRequest
//
// @return ListScheduledPreloadExecutionsResponse
func (client *Client) ListScheduledPreloadExecutions(request *ListScheduledPreloadExecutionsRequest) (_result *ListScheduledPreloadExecutionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListScheduledPreloadExecutionsResponse{}
	_body, _err := client.ListScheduledPreloadExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the scheduled prefetch tasks for a website.
//
// @param request - ListScheduledPreloadJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduledPreloadJobsResponse
func (client *Client) ListScheduledPreloadJobsWithOptions(request *ListScheduledPreloadJobsRequest, runtime *dara.RuntimeOptions) (_result *ListScheduledPreloadJobsResponse, _err error) {
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
		Action:      dara.String("ListScheduledPreloadJobs"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScheduledPreloadJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the scheduled prefetch tasks for a website.
//
// @param request - ListScheduledPreloadJobsRequest
//
// @return ListScheduledPreloadJobsResponse
func (client *Client) ListScheduledPreloadJobs(request *ListScheduledPreloadJobsRequest) (_result *ListScheduledPreloadJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListScheduledPreloadJobsResponse{}
	_body, _err := client.ListScheduledPreloadJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all log delivery tasks that are in progress.
//
// @param request - ListSiteDeliveryTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSiteDeliveryTasksResponse
func (client *Client) ListSiteDeliveryTasksWithOptions(request *ListSiteDeliveryTasksRequest, runtime *dara.RuntimeOptions) (_result *ListSiteDeliveryTasksResponse, _err error) {
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
		Action:      dara.String("ListSiteDeliveryTasks"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSiteDeliveryTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all log delivery tasks that are in progress.
//
// @param request - ListSiteDeliveryTasksRequest
//
// @return ListSiteDeliveryTasksResponse
func (client *Client) ListSiteDeliveryTasks(request *ListSiteDeliveryTasksRequest) (_result *ListSiteDeliveryTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSiteDeliveryTasksResponse{}
	_body, _err := client.ListSiteDeliveryTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the edge function routes for a website.
//
// @param request - ListSiteRoutesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSiteRoutesResponse
func (client *Client) ListSiteRoutesWithOptions(request *ListSiteRoutesRequest, runtime *dara.RuntimeOptions) (_result *ListSiteRoutesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.ConfigType) {
		query["ConfigType"] = request.ConfigType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RouteName) {
		query["RouteName"] = request.RouteName
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSiteRoutes"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSiteRoutesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the edge function routes for a website.
//
// @param request - ListSiteRoutesRequest
//
// @return ListSiteRoutesResponse
func (client *Client) ListSiteRoutes(request *ListSiteRoutesRequest) (_result *ListSiteRoutesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSiteRoutesResponse{}
	_body, _err := client.ListSiteRoutesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about websites in your account, such as the name, status, and configuration of each website.
//
// @param tmpReq - ListSitesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSitesResponse
func (client *Client) ListSitesWithOptions(tmpReq *ListSitesRequest, runtime *dara.RuntimeOptions) (_result *ListSitesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSitesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TagFilter) {
		request.TagFilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagFilter, dara.String("TagFilter"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSites"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSitesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about websites in your account, such as the name, status, and configuration of each website.
//
// @param request - ListSitesRequest
//
// @return ListSitesResponse
func (client *Client) ListSites(request *ListSitesRequest) (_result *ListSitesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSitesResponse{}
	_body, _err := client.ListSitesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries tags based on the region ID and resource type.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxItem) {
		query["MaxItem"] = request.MaxItem
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tags based on the region ID and resource type.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List of Transport Layer Applications
//
// @param request - ListTransportLayerApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransportLayerApplicationsResponse
func (client *Client) ListTransportLayerApplicationsWithOptions(request *ListTransportLayerApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListTransportLayerApplicationsResponse, _err error) {
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
		Action:      dara.String("ListTransportLayerApplications"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransportLayerApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List of Transport Layer Applications
//
// @param request - ListTransportLayerApplicationsRequest
//
// @return ListTransportLayerApplicationsResponse
func (client *Client) ListTransportLayerApplications(request *ListTransportLayerApplicationsRequest) (_result *ListTransportLayerApplicationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransportLayerApplicationsResponse{}
	_body, _err := client.ListTransportLayerApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution status and running information of file upload tasks based on the task time and type.
//
// @param request - ListUploadTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUploadTasksResponse
func (client *Client) ListUploadTasksWithOptions(request *ListUploadTasksRequest, runtime *dara.RuntimeOptions) (_result *ListUploadTasksResponse, _err error) {
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
		Action:      dara.String("ListUploadTasks"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUploadTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution status and running information of file upload tasks based on the task time and type.
//
// @param request - ListUploadTasksRequest
//
// @return ListUploadTasksResponse
func (client *Client) ListUploadTasks(request *ListUploadTasksRequest) (_result *ListUploadTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUploadTasksResponse{}
	_body, _err := client.ListUploadTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of page monitoring configurations.
//
// @param request - ListUrlObservationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUrlObservationsResponse
func (client *Client) ListUrlObservationsWithOptions(request *ListUrlObservationsRequest, runtime *dara.RuntimeOptions) (_result *ListUrlObservationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUrlObservations"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUrlObservationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of page monitoring configurations.
//
// @param request - ListUrlObservationsRequest
//
// @return ListUrlObservationsResponse
func (client *Client) ListUrlObservations(request *ListUrlObservationsRequest) (_result *ListUrlObservationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUrlObservationsResponse{}
	_body, _err := client.ListUrlObservationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all delivery tasks in your Alibaba Cloud account by page. You can filter the delivery tasks by the category of the delivered real-time logs.
//
// @param request - ListUserDeliveryTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserDeliveryTasksResponse
func (client *Client) ListUserDeliveryTasksWithOptions(request *ListUserDeliveryTasksRequest, runtime *dara.RuntimeOptions) (_result *ListUserDeliveryTasksResponse, _err error) {
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
		Action:      dara.String("ListUserDeliveryTasks"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserDeliveryTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all delivery tasks in your Alibaba Cloud account by page. You can filter the delivery tasks by the category of the delivered real-time logs.
//
// @param request - ListUserDeliveryTasksRequest
//
// @return ListUserDeliveryTasksResponse
func (client *Client) ListUserDeliveryTasks(request *ListUserDeliveryTasksRequest) (_result *ListUserDeliveryTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserDeliveryTasksResponse{}
	_body, _err := client.ListUserDeliveryTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the plans that you purchased and the details of the plans.
//
// @param request - ListUserRatePlanInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserRatePlanInstancesResponse
func (client *Client) ListUserRatePlanInstancesWithOptions(request *ListUserRatePlanInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListUserRatePlanInstancesResponse, _err error) {
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
		Action:      dara.String("ListUserRatePlanInstances"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserRatePlanInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the plans that you purchased and the details of the plans.
//
// @param request - ListUserRatePlanInstancesRequest
//
// @return ListUserRatePlanInstancesResponse
func (client *Client) ListUserRatePlanInstances(request *ListUserRatePlanInstancesRequest) (_result *ListUserRatePlanInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserRatePlanInstancesResponse{}
	_body, _err := client.ListUserRatePlanInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the functions created in your account and the maximum number of functions supported by your plan. You can call this operation to perform a paged query.
//
// Description:
//
// You can call this operation to perform a paged query to query all functions created in your account, the maximum number of functions supported by the billing plan that you use, and the number of functions already created. You can specify `PageNumber` and `PageSize` to control the number of entries to be returned in the response and specify `SearchKeyWord` to perform a fuzzy search to filter specific routine names.
//
// @param request - ListUserRoutinesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserRoutinesResponse
func (client *Client) ListUserRoutinesWithOptions(request *ListUserRoutinesRequest, runtime *dara.RuntimeOptions) (_result *ListUserRoutinesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKeyWord) {
		query["SearchKeyWord"] = request.SearchKeyWord
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserRoutines"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserRoutinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the functions created in your account and the maximum number of functions supported by your plan. You can call this operation to perform a paged query.
//
// Description:
//
// You can call this operation to perform a paged query to query all functions created in your account, the maximum number of functions supported by the billing plan that you use, and the number of functions already created. You can specify `PageNumber` and `PageSize` to control the number of entries to be returned in the response and specify `SearchKeyWord` to perform a fuzzy search to filter specific routine names.
//
// @param request - ListUserRoutinesRequest
//
// @return ListUserRoutinesResponse
func (client *Client) ListUserRoutines(request *ListUserRoutinesRequest) (_result *ListUserRoutinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserRoutinesResponse{}
	_body, _err := client.ListUserRoutinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the video processing configurations of a site.
//
// @param request - ListVideoProcessingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVideoProcessingsResponse
func (client *Client) ListVideoProcessingsWithOptions(request *ListVideoProcessingsRequest, runtime *dara.RuntimeOptions) (_result *ListVideoProcessingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.ConfigType) {
		query["ConfigType"] = request.ConfigType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVideoProcessings"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVideoProcessingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the video processing configurations of a site.
//
// @param request - ListVideoProcessingsRequest
//
// @return ListVideoProcessingsResponse
func (client *Client) ListVideoProcessings(request *ListVideoProcessingsRequest) (_result *ListVideoProcessingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVideoProcessingsResponse{}
	_body, _err := client.ListVideoProcessingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List WAF Managed Rules
//
// @param tmpReq - ListWafManagedRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWafManagedRulesResponse
func (client *Client) ListWafManagedRulesWithOptions(tmpReq *ListWafManagedRulesRequest, runtime *dara.RuntimeOptions) (_result *ListWafManagedRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListWafManagedRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ManagedRuleset) {
		request.ManagedRulesetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ManagedRuleset, dara.String("ManagedRuleset"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.QueryArgs) {
		request.QueryArgsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryArgs, dara.String("QueryArgs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttackType) {
		query["AttackType"] = request.AttackType
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.ManagedRulesetShrink) {
		query["ManagedRuleset"] = request.ManagedRulesetShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProtectionLevel) {
		query["ProtectionLevel"] = request.ProtectionLevel
	}

	if !dara.IsNil(request.QueryArgsShrink) {
		query["QueryArgs"] = request.QueryArgsShrink
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWafManagedRules"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWafManagedRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List WAF Managed Rules
//
// @param request - ListWafManagedRulesRequest
//
// @return ListWafManagedRulesResponse
func (client *Client) ListWafManagedRules(request *ListWafManagedRulesRequest) (_result *ListWafManagedRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWafManagedRulesResponse{}
	_body, _err := client.ListWafManagedRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List WAF Phases
//
// @param request - ListWafPhasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWafPhasesResponse
func (client *Client) ListWafPhasesWithOptions(request *ListWafPhasesRequest, runtime *dara.RuntimeOptions) (_result *ListWafPhasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWafPhases"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWafPhasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List WAF Phases
//
// @param request - ListWafPhasesRequest
//
// @return ListWafPhasesResponse
func (client *Client) ListWafPhases(request *ListWafPhasesRequest) (_result *ListWafPhasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWafPhasesResponse{}
	_body, _err := client.ListWafPhasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List WAF Rules
//
// @param tmpReq - ListWafRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWafRulesResponse
func (client *Client) ListWafRulesWithOptions(tmpReq *ListWafRulesRequest, runtime *dara.RuntimeOptions) (_result *ListWafRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListWafRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.QueryArgs) {
		request.QueryArgsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryArgs, dara.String("QueryArgs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Phase) {
		query["Phase"] = request.Phase
	}

	if !dara.IsNil(request.QueryArgsShrink) {
		query["QueryArgs"] = request.QueryArgsShrink
	}

	if !dara.IsNil(request.RulesetId) {
		query["RulesetId"] = request.RulesetId
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWafRules"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWafRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List WAF Rules
//
// @param request - ListWafRulesRequest
//
// @return ListWafRulesResponse
func (client *Client) ListWafRules(request *ListWafRulesRequest) (_result *ListWafRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWafRulesResponse{}
	_body, _err := client.ListWafRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List WAF Rule Sets
//
// @param tmpReq - ListWafRulesetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWafRulesetsResponse
func (client *Client) ListWafRulesetsWithOptions(tmpReq *ListWafRulesetsRequest, runtime *dara.RuntimeOptions) (_result *ListWafRulesetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListWafRulesetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.QueryArgs) {
		request.QueryArgsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryArgs, dara.String("QueryArgs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Phase) {
		query["Phase"] = request.Phase
	}

	if !dara.IsNil(request.QueryArgsShrink) {
		query["QueryArgs"] = request.QueryArgsShrink
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWafRulesets"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWafRulesetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List WAF Rule Sets
//
// @param request - ListWafRulesetsRequest
//
// @return ListWafRulesetsResponse
func (client *Client) ListWafRulesets(request *ListWafRulesetsRequest) (_result *ListWafRulesetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWafRulesetsResponse{}
	_body, _err := client.ListWafRulesetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List WAF Template Rules
//
// @param tmpReq - ListWafTemplateRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWafTemplateRulesResponse
func (client *Client) ListWafTemplateRulesWithOptions(tmpReq *ListWafTemplateRulesRequest, runtime *dara.RuntimeOptions) (_result *ListWafTemplateRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListWafTemplateRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.QueryArgs) {
		request.QueryArgsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryArgs, dara.String("QueryArgs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Phase) {
		query["Phase"] = request.Phase
	}

	if !dara.IsNil(request.QueryArgsShrink) {
		query["QueryArgs"] = request.QueryArgsShrink
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWafTemplateRules"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWafTemplateRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List WAF Template Rules
//
// @param request - ListWafTemplateRulesRequest
//
// @return ListWafTemplateRulesResponse
func (client *Client) ListWafTemplateRules(request *ListWafTemplateRulesRequest) (_result *ListWafTemplateRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWafTemplateRulesResponse{}
	_body, _err := client.ListWafTemplateRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List WAF Rule Usage
//
// @param request - ListWafUsageOfRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWafUsageOfRulesResponse
func (client *Client) ListWafUsageOfRulesWithOptions(request *ListWafUsageOfRulesRequest, runtime *dara.RuntimeOptions) (_result *ListWafUsageOfRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Phase) {
		query["Phase"] = request.Phase
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWafUsageOfRules"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWafUsageOfRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List WAF Rule Usage
//
// @param request - ListWafUsageOfRulesRequest
//
// @return ListWafUsageOfRulesResponse
func (client *Client) ListWafUsageOfRules(request *ListWafUsageOfRulesRequest) (_result *ListWafUsageOfRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWafUsageOfRulesResponse{}
	_body, _err := client.ListWafUsageOfRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about waiting room events for a waiting room.
//
// Description:
//
// You can call this operation to query details of all waiting room events related to a waiting room in a website.
//
// @param request - ListWaitingRoomEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWaitingRoomEventsResponse
func (client *Client) ListWaitingRoomEventsWithOptions(request *ListWaitingRoomEventsRequest, runtime *dara.RuntimeOptions) (_result *ListWaitingRoomEventsResponse, _err error) {
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
		Action:      dara.String("ListWaitingRoomEvents"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWaitingRoomEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about waiting room events for a waiting room.
//
// Description:
//
// You can call this operation to query details of all waiting room events related to a waiting room in a website.
//
// @param request - ListWaitingRoomEventsRequest
//
// @return ListWaitingRoomEventsResponse
func (client *Client) ListWaitingRoomEvents(request *ListWaitingRoomEventsRequest) (_result *ListWaitingRoomEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWaitingRoomEventsResponse{}
	_body, _err := client.ListWaitingRoomEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Waiting Room Bypass Rules
//
// Description:
//
// This API allows users to query the list of waiting room bypass rules associated with a specific site.
//
// @param request - ListWaitingRoomRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWaitingRoomRulesResponse
func (client *Client) ListWaitingRoomRulesWithOptions(request *ListWaitingRoomRulesRequest, runtime *dara.RuntimeOptions) (_result *ListWaitingRoomRulesResponse, _err error) {
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
		Action:      dara.String("ListWaitingRoomRules"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWaitingRoomRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Waiting Room Bypass Rules
//
// Description:
//
// This API allows users to query the list of waiting room bypass rules associated with a specific site.
//
// @param request - ListWaitingRoomRulesRequest
//
// @return ListWaitingRoomRulesResponse
func (client *Client) ListWaitingRoomRules(request *ListWaitingRoomRulesRequest) (_result *ListWaitingRoomRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWaitingRoomRulesResponse{}
	_body, _err := client.ListWaitingRoomRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about all waiting rooms in a website.
//
// Description:
//
// You can call this operation to query detailed configurations about all waiting rooms in a website, including the status, name, and queuing rules of each waiting room.
//
// @param request - ListWaitingRoomsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWaitingRoomsResponse
func (client *Client) ListWaitingRoomsWithOptions(request *ListWaitingRoomsRequest, runtime *dara.RuntimeOptions) (_result *ListWaitingRoomsResponse, _err error) {
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
		Action:      dara.String("ListWaitingRooms"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWaitingRoomsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all waiting rooms in a website.
//
// Description:
//
// You can call this operation to query detailed configurations about all waiting rooms in a website, including the status, name, and queuing rules of each waiting room.
//
// @param request - ListWaitingRoomsRequest
//
// @return ListWaitingRoomsResponse
func (client *Client) ListWaitingRooms(request *ListWaitingRoomsRequest) (_result *ListWaitingRoomsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWaitingRoomsResponse{}
	_body, _err := client.ListWaitingRoomsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # OpenErService
//
// @param request - OpenErServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenErServiceResponse
func (client *Client) OpenErServiceWithOptions(request *OpenErServiceRequest, runtime *dara.RuntimeOptions) (_result *OpenErServiceResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenErService"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenErServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # OpenErService
//
// @param request - OpenErServiceRequest
//
// @return OpenErServiceResponse
func (client *Client) OpenErService(request *OpenErServiceRequest) (_result *OpenErServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenErServiceResponse{}
	_body, _err := client.OpenErServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Prefetches cache.
//
// @param tmpReq - PreloadCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreloadCachesResponse
func (client *Client) PreloadCachesWithOptions(tmpReq *PreloadCachesRequest, runtime *dara.RuntimeOptions) (_result *PreloadCachesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PreloadCachesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Content) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, dara.String("Content"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Headers) {
		request.HeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Headers, dara.String("Headers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ContentShrink) {
		query["Content"] = request.ContentShrink
	}

	if !dara.IsNil(request.HeadersShrink) {
		query["Headers"] = request.HeadersShrink
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreloadCaches"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreloadCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Prefetches cache.
//
// @param request - PreloadCachesRequest
//
// @return PreloadCachesResponse
func (client *Client) PreloadCaches(request *PreloadCachesRequest) (_result *PreloadCachesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PreloadCachesResponse{}
	_body, _err := client.PreloadCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a specific version of a containerized application. You can call this operation to iterate an application.
//
// @param tmpReq - PublishEdgeContainerAppVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishEdgeContainerAppVersionResponse
func (client *Client) PublishEdgeContainerAppVersionWithOptions(tmpReq *PublishEdgeContainerAppVersionRequest, runtime *dara.RuntimeOptions) (_result *PublishEdgeContainerAppVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PublishEdgeContainerAppVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Regions) {
		request.RegionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Regions, dara.String("Regions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FullRelease) {
		query["FullRelease"] = request.FullRelease
	}

	if !dara.IsNil(request.PublishType) {
		query["PublishType"] = request.PublishType
	}

	if !dara.IsNil(request.RegionsShrink) {
		query["Regions"] = request.RegionsShrink
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Percentage) {
		body["Percentage"] = request.Percentage
	}

	if !dara.IsNil(request.PublishEnv) {
		body["PublishEnv"] = request.PublishEnv
	}

	if !dara.IsNil(request.Remarks) {
		body["Remarks"] = request.Remarks
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishEdgeContainerAppVersion"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishEdgeContainerAppVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a specific version of a containerized application. You can call this operation to iterate an application.
//
// @param request - PublishEdgeContainerAppVersionRequest
//
// @return PublishEdgeContainerAppVersionResponse
func (client *Client) PublishEdgeContainerAppVersion(request *PublishEdgeContainerAppVersionRequest) (_result *PublishEdgeContainerAppVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishEdgeContainerAppVersionResponse{}
	_body, _err := client.PublishEdgeContainerAppVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a code version of a routine to the staging, canary, or production environment. You can specify the regions where the canary environment is deployed to release your code.
//
// @param request - PublishRoutineCodeVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishRoutineCodeVersionResponse
func (client *Client) PublishRoutineCodeVersionWithOptions(request *PublishRoutineCodeVersionRequest, runtime *dara.RuntimeOptions) (_result *PublishRoutineCodeVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CodeVersion) {
		body["CodeVersion"] = request.CodeVersion
	}

	if !dara.IsNil(request.Env) {
		body["Env"] = request.Env
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishRoutineCodeVersion"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishRoutineCodeVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a code version of a routine to the staging, canary, or production environment. You can specify the regions where the canary environment is deployed to release your code.
//
// @param request - PublishRoutineCodeVersionRequest
//
// @return PublishRoutineCodeVersionResponse
func (client *Client) PublishRoutineCodeVersion(request *PublishRoutineCodeVersionRequest) (_result *PublishRoutineCodeVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishRoutineCodeVersionResponse{}
	_body, _err := client.PublishRoutineCodeVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # New Purchase of Cache Retention
//
// @param request - PurchaseCacheReserveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PurchaseCacheReserveResponse
func (client *Client) PurchaseCacheReserveWithOptions(request *PurchaseCacheReserveRequest, runtime *dara.RuntimeOptions) (_result *PurchaseCacheReserveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.CrRegion) {
		query["CrRegion"] = request.CrRegion
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.QuotaGb) {
		query["QuotaGb"] = request.QuotaGb
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PurchaseCacheReserve"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PurchaseCacheReserveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # New Purchase of Cache Retention
//
// @param request - PurchaseCacheReserveRequest
//
// @return PurchaseCacheReserveResponse
func (client *Client) PurchaseCacheReserve(request *PurchaseCacheReserveRequest) (_result *PurchaseCacheReserveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PurchaseCacheReserveResponse{}
	_body, _err := client.PurchaseCacheReserveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Purchase New Package
//
// Description:
//
// 1. The package name and code can be obtained from the DescribeRatePlanPrice interface.
//
// 2. If the acceleration area is not overseas, the site must have successfully completed the filing process.
//
// @param request - PurchaseRatePlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PurchaseRatePlanResponse
func (client *Client) PurchaseRatePlanWithOptions(request *PurchaseRatePlanRequest, runtime *dara.RuntimeOptions) (_result *PurchaseRatePlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.Coverage) {
		query["Coverage"] = request.Coverage
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PlanCode) {
		query["PlanCode"] = request.PlanCode
	}

	if !dara.IsNil(request.PlanName) {
		query["PlanName"] = request.PlanName
	}

	if !dara.IsNil(request.SiteName) {
		query["SiteName"] = request.SiteName
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PurchaseRatePlan"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PurchaseRatePlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Purchase New Package
//
// Description:
//
// 1. The package name and code can be obtained from the DescribeRatePlanPrice interface.
//
// 2. If the acceleration area is not overseas, the site must have successfully completed the filing process.
//
// @param request - PurchaseRatePlanRequest
//
// @return PurchaseRatePlanResponse
func (client *Client) PurchaseRatePlan(request *PurchaseRatePlanRequest) (_result *PurchaseRatePlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PurchaseRatePlanResponse{}
	_body, _err := client.PurchaseRatePlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Cache Refresh
//
// @param tmpReq - PurgeCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PurgeCachesResponse
func (client *Client) PurgeCachesWithOptions(tmpReq *PurgeCachesRequest, runtime *dara.RuntimeOptions) (_result *PurgeCachesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PurgeCachesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Content) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, dara.String("Content"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ContentShrink) {
		query["Content"] = request.ContentShrink
	}

	if !dara.IsNil(request.EdgeComputePurge) {
		query["EdgeComputePurge"] = request.EdgeComputePurge
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PurgeCaches"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PurgeCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Cache Refresh
//
// @param request - PurgeCachesRequest
//
// @return PurgeCachesResponse
func (client *Client) PurgeCaches(request *PurgeCachesRequest) (_result *PurgeCachesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PurgeCachesResponse{}
	_body, _err := client.PurgeCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a key-value pair for a namespace. The request body can be up to 2 MB.
//
// @param request - PutKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutKvResponse
func (client *Client) PutKvWithOptions(request *PutKvRequest, runtime *dara.RuntimeOptions) (_result *PutKvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Base64) {
		query["Base64"] = request.Base64
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.ExpirationTtl) {
		query["ExpirationTtl"] = request.ExpirationTtl
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Value) {
		body["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutKv"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutKvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a key-value pair for a namespace. The request body can be up to 2 MB.
//
// @param request - PutKvRequest
//
// @return PutKvResponse
func (client *Client) PutKv(request *PutKvRequest) (_result *PutKvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutKvResponse{}
	_body, _err := client.PutKvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a large key-value pair for a namespace. The request body can be up to 25 MB.
//
// Description:
//
// This operation allows you to upload a larger request body than by using [PutKv](~~PutKv~~). For small request bodies, we recommend that you use [PutKv](~~PutKv~~) to minimize the server processing time. This operation must be called by using SDKs. The following sample code uses the Golang SDK and PutKvWithHighCapacityAdvance to call the operation.
//
//	func TestPutKvWithHighCapacity() {
//
//		// Initialize the configurations.
//
//		cfg := new(openapi.Config)
//
//		cfg.SetAccessKeyId("xxxxxxxxx")
//
//		cfg.SetAccessKeySecret("xxxxxxxxxx")
//
//		cli, err := NewClient(cfg)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		runtime := &util.RuntimeOptions{}
//
//		// Construct a request for uploading key-value pairs.
//
//		namespace := "test-put-kv"
//
//		key := "test_PutKvWithHighCapacity_0"
//
//		value := strings.Repeat("t", 10*1024*1024)
//
//		rawReq := &PutKvRequest{
//
//			Namespace: &namespace,
//
//			Key:       &key,
//
//			Value:     &value,
//
//		}
//
//		payload, err := json.Marshal(rawReq)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		// If the payload is greater than 2 MB, call the PutKvWithHighCapacity operation for upload.
//
//		reqHighCapacity := &PutKvWithHighCapacityAdvanceRequest{
//
//			Namespace: &namespace,
//
//			Key:       &key,
//
//			UrlObject: bytes.NewReader([]byte(payload)),
//
//		}
//
//		resp, err := cli.PutKvWithHighCapacityAdvance(reqHighCapacity, runtime)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		return nil
//
//	}
//
// @param request - PutKvWithHighCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutKvWithHighCapacityResponse
func (client *Client) PutKvWithHighCapacityWithOptions(request *PutKvWithHighCapacityRequest, runtime *dara.RuntimeOptions) (_result *PutKvWithHighCapacityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutKvWithHighCapacity"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutKvWithHighCapacityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a large key-value pair for a namespace. The request body can be up to 25 MB.
//
// Description:
//
// This operation allows you to upload a larger request body than by using [PutKv](~~PutKv~~). For small request bodies, we recommend that you use [PutKv](~~PutKv~~) to minimize the server processing time. This operation must be called by using SDKs. The following sample code uses the Golang SDK and PutKvWithHighCapacityAdvance to call the operation.
//
//	func TestPutKvWithHighCapacity() {
//
//		// Initialize the configurations.
//
//		cfg := new(openapi.Config)
//
//		cfg.SetAccessKeyId("xxxxxxxxx")
//
//		cfg.SetAccessKeySecret("xxxxxxxxxx")
//
//		cli, err := NewClient(cfg)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		runtime := &util.RuntimeOptions{}
//
//		// Construct a request for uploading key-value pairs.
//
//		namespace := "test-put-kv"
//
//		key := "test_PutKvWithHighCapacity_0"
//
//		value := strings.Repeat("t", 10*1024*1024)
//
//		rawReq := &PutKvRequest{
//
//			Namespace: &namespace,
//
//			Key:       &key,
//
//			Value:     &value,
//
//		}
//
//		payload, err := json.Marshal(rawReq)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		// If the payload is greater than 2 MB, call the PutKvWithHighCapacity operation for upload.
//
//		reqHighCapacity := &PutKvWithHighCapacityAdvanceRequest{
//
//			Namespace: &namespace,
//
//			Key:       &key,
//
//			UrlObject: bytes.NewReader([]byte(payload)),
//
//		}
//
//		resp, err := cli.PutKvWithHighCapacityAdvance(reqHighCapacity, runtime)
//
//		if err != nil {
//
//			return err
//
//		}
//
//		return nil
//
//	}
//
// @param request - PutKvWithHighCapacityRequest
//
// @return PutKvWithHighCapacityResponse
func (client *Client) PutKvWithHighCapacity(request *PutKvWithHighCapacityRequest) (_result *PutKvWithHighCapacityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutKvWithHighCapacityResponse{}
	_body, _err := client.PutKvWithHighCapacityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutKvWithHighCapacityAdvance(request *PutKvWithHighCapacityAdvanceRequest, runtime *dara.RuntimeOptions) (_result *PutKvWithHighCapacityResponse, _err error) {
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
		"Product":  dara.String("ESA"),
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
	putKvWithHighCapacityReq := &PutKvWithHighCapacityRequest{}
	openapiutil.Convert(request, putKvWithHighCapacityReq)
	if !dara.IsNil(request.UrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.UrlObject,
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
		putKvWithHighCapacityReq.Url = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	putKvWithHighCapacityResp, _err := client.PutKvWithHighCapacityWithOptions(putKvWithHighCapacityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = putKvWithHighCapacityResp
	return _result, _err
}

// Summary:
//
// Rebuilds the staging environment for containerized applications.
//
// @param request - RebuildEdgeContainerAppStagingEnvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebuildEdgeContainerAppStagingEnvResponse
func (client *Client) RebuildEdgeContainerAppStagingEnvWithOptions(request *RebuildEdgeContainerAppStagingEnvRequest, runtime *dara.RuntimeOptions) (_result *RebuildEdgeContainerAppStagingEnvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebuildEdgeContainerAppStagingEnv"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebuildEdgeContainerAppStagingEnvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rebuilds the staging environment for containerized applications.
//
// @param request - RebuildEdgeContainerAppStagingEnvRequest
//
// @return RebuildEdgeContainerAppStagingEnvResponse
func (client *Client) RebuildEdgeContainerAppStagingEnv(request *RebuildEdgeContainerAppStagingEnvRequest) (_result *RebuildEdgeContainerAppStagingEnvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RebuildEdgeContainerAppStagingEnvResponse{}
	_body, _err := client.RebuildEdgeContainerAppStagingEnvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the progress of a scheduled prefetch task and starts the prefetch from the beginning.
//
// @param request - ResetScheduledPreloadJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetScheduledPreloadJobResponse
func (client *Client) ResetScheduledPreloadJobWithOptions(request *ResetScheduledPreloadJobRequest, runtime *dara.RuntimeOptions) (_result *ResetScheduledPreloadJobResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetScheduledPreloadJob"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetScheduledPreloadJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the progress of a scheduled prefetch task and starts the prefetch from the beginning.
//
// @param request - ResetScheduledPreloadJobRequest
//
// @return ResetScheduledPreloadJobResponse
func (client *Client) ResetScheduledPreloadJob(request *ResetScheduledPreloadJobRequest) (_result *ResetScheduledPreloadJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetScheduledPreloadJobResponse{}
	_body, _err := client.ResetScheduledPreloadJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes an activated client certificate.
//
// @param request - RevokeClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeClientCertificateResponse
func (client *Client) RevokeClientCertificateWithOptions(request *RevokeClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *RevokeClientCertificateResponse, _err error) {
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
		Action:      dara.String("RevokeClientCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes an activated client certificate.
//
// @param request - RevokeClientCertificateRequest
//
// @return RevokeClientCertificateResponse
func (client *Client) RevokeClientCertificate(request *RevokeClientCertificateRequest) (_result *RevokeClientCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeClientCertificateResponse{}
	_body, _err := client.RevokeClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rolls back a version of a containerized application.
//
// @param request - RollbackEdgeContainerAppVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackEdgeContainerAppVersionResponse
func (client *Client) RollbackEdgeContainerAppVersionWithOptions(request *RollbackEdgeContainerAppVersionRequest, runtime *dara.RuntimeOptions) (_result *RollbackEdgeContainerAppVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Percentage) {
		query["Percentage"] = request.Percentage
	}

	if !dara.IsNil(request.UsedPercent) {
		query["UsedPercent"] = request.UsedPercent
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Remarks) {
		body["Remarks"] = request.Remarks
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackEdgeContainerAppVersion"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackEdgeContainerAppVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back a version of a containerized application.
//
// @param request - RollbackEdgeContainerAppVersionRequest
//
// @return RollbackEdgeContainerAppVersionResponse
func (client *Client) RollbackEdgeContainerAppVersion(request *RollbackEdgeContainerAppVersionRequest) (_result *RollbackEdgeContainerAppVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RollbackEdgeContainerAppVersionResponse{}
	_body, _err := client.RollbackEdgeContainerAppVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures whether to enable certificates and update certificate information for a website.
//
// @param request - SetCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCertificateResponse
func (client *Client) SetCertificateWithOptions(request *SetCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetCertificateResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CasId) {
		body["CasId"] = request.CasId
	}

	if !dara.IsNil(request.Certificate) {
		body["Certificate"] = request.Certificate
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PrivateKey) {
		body["PrivateKey"] = request.PrivateKey
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures whether to enable certificates and update certificate information for a website.
//
// @param request - SetCertificateRequest
//
// @return SetCertificateResponse
func (client *Client) SetCertificate(request *SetCertificateRequest) (_result *SetCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetCertificateResponse{}
	_body, _err := client.SetCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates domain names with a client CA certificate. If no certificate is specified, domain names are associated with an Edge Security Acceleration (ESA)-managed CA certificate.
//
// @param tmpReq - SetClientCertificateHostnamesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetClientCertificateHostnamesResponse
func (client *Client) SetClientCertificateHostnamesWithOptions(tmpReq *SetClientCertificateHostnamesRequest, runtime *dara.RuntimeOptions) (_result *SetClientCertificateHostnamesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetClientCertificateHostnamesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Hostnames) {
		request.HostnamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Hostnames, dara.String("Hostnames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HostnamesShrink) {
		body["Hostnames"] = request.HostnamesShrink
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetClientCertificateHostnames"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetClientCertificateHostnamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates domain names with a client CA certificate. If no certificate is specified, domain names are associated with an Edge Security Acceleration (ESA)-managed CA certificate.
//
// @param request - SetClientCertificateHostnamesRequest
//
// @return SetClientCertificateHostnamesResponse
func (client *Client) SetClientCertificateHostnames(request *SetClientCertificateHostnamesRequest) (_result *SetClientCertificateHostnamesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetClientCertificateHostnamesResponse{}
	_body, _err := client.SetClientCertificateHostnamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures smart HTTP DDoS protection.
//
// @param request - SetHttpDDoSAttackIntelligentProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetHttpDDoSAttackIntelligentProtectionResponse
func (client *Client) SetHttpDDoSAttackIntelligentProtectionWithOptions(request *SetHttpDDoSAttackIntelligentProtectionRequest, runtime *dara.RuntimeOptions) (_result *SetHttpDDoSAttackIntelligentProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AiMode) {
		query["AiMode"] = request.AiMode
	}

	if !dara.IsNil(request.AiTemplate) {
		query["AiTemplate"] = request.AiTemplate
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetHttpDDoSAttackIntelligentProtection"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetHttpDDoSAttackIntelligentProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures smart HTTP DDoS protection.
//
// @param request - SetHttpDDoSAttackIntelligentProtectionRequest
//
// @return SetHttpDDoSAttackIntelligentProtectionResponse
func (client *Client) SetHttpDDoSAttackIntelligentProtection(request *SetHttpDDoSAttackIntelligentProtectionRequest) (_result *SetHttpDDoSAttackIntelligentProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetHttpDDoSAttackIntelligentProtectionResponse{}
	_body, _err := client.SetHttpDDoSAttackIntelligentProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures HTTP DDoS attack protection for a website.
//
// @param request - SetHttpDDoSAttackProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetHttpDDoSAttackProtectionResponse
func (client *Client) SetHttpDDoSAttackProtectionWithOptions(request *SetHttpDDoSAttackProtectionRequest, runtime *dara.RuntimeOptions) (_result *SetHttpDDoSAttackProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GlobalMode) {
		query["GlobalMode"] = request.GlobalMode
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetHttpDDoSAttackProtection"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetHttpDDoSAttackProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures HTTP DDoS attack protection for a website.
//
// @param request - SetHttpDDoSAttackProtectionRequest
//
// @return SetHttpDDoSAttackProtectionResponse
func (client *Client) SetHttpDDoSAttackProtection(request *SetHttpDDoSAttackProtectionRequest) (_result *SetHttpDDoSAttackProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetHttpDDoSAttackProtectionResponse{}
	_body, _err := client.SetHttpDDoSAttackProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为域名回源客户端证书绑定域名
//
// @param tmpReq - SetOriginClientCertificateHostnamesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetOriginClientCertificateHostnamesResponse
func (client *Client) SetOriginClientCertificateHostnamesWithOptions(tmpReq *SetOriginClientCertificateHostnamesRequest, runtime *dara.RuntimeOptions) (_result *SetOriginClientCertificateHostnamesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetOriginClientCertificateHostnamesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Hostnames) {
		request.HostnamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Hostnames, dara.String("Hostnames"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HostnamesShrink) {
		body["Hostnames"] = request.HostnamesShrink
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetOriginClientCertificateHostnames"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetOriginClientCertificateHostnamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为域名回源客户端证书绑定域名
//
// @param request - SetOriginClientCertificateHostnamesRequest
//
// @return SetOriginClientCertificateHostnamesResponse
func (client *Client) SetOriginClientCertificateHostnames(request *SetOriginClientCertificateHostnamesRequest) (_result *SetOriginClientCertificateHostnamesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetOriginClientCertificateHostnamesResponse{}
	_body, _err := client.SetOriginClientCertificateHostnamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a scheduled prefetch plan based on the plan ID.
//
// @param request - StartScheduledPreloadExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartScheduledPreloadExecutionResponse
func (client *Client) StartScheduledPreloadExecutionWithOptions(request *StartScheduledPreloadExecutionRequest, runtime *dara.RuntimeOptions) (_result *StartScheduledPreloadExecutionResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartScheduledPreloadExecution"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartScheduledPreloadExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a scheduled prefetch plan based on the plan ID.
//
// @param request - StartScheduledPreloadExecutionRequest
//
// @return StartScheduledPreloadExecutionResponse
func (client *Client) StartScheduledPreloadExecution(request *StartScheduledPreloadExecutionRequest) (_result *StartScheduledPreloadExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartScheduledPreloadExecutionResponse{}
	_body, _err := client.StartScheduledPreloadExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a scheduled prefetch plan based on the plan ID.
//
// @param request - StopScheduledPreloadExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopScheduledPreloadExecutionResponse
func (client *Client) StopScheduledPreloadExecutionWithOptions(request *StopScheduledPreloadExecutionRequest, runtime *dara.RuntimeOptions) (_result *StopScheduledPreloadExecutionResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopScheduledPreloadExecution"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopScheduledPreloadExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a scheduled prefetch plan based on the plan ID.
//
// @param request - StopScheduledPreloadExecutionRequest
//
// @return StopScheduledPreloadExecutionResponse
func (client *Client) StopScheduledPreloadExecution(request *StopScheduledPreloadExecutionRequest) (_result *StopScheduledPreloadExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopScheduledPreloadExecutionResponse{}
	_body, _err := client.StopScheduledPreloadExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds one or more tags to resources.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds one or more tags to resources.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a resource tag based on a specified resource ID.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a resource tag based on a specified resource ID.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Cache Reserve Specification Change
//
// @param request - UpdateCacheReserveSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCacheReserveSpecResponse
func (client *Client) UpdateCacheReserveSpecWithOptions(request *UpdateCacheReserveSpecRequest, runtime *dara.RuntimeOptions) (_result *UpdateCacheReserveSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TargetQuotaGb) {
		query["TargetQuotaGb"] = request.TargetQuotaGb
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCacheReserveSpec"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCacheReserveSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Cache Reserve Specification Change
//
// @param request - UpdateCacheReserveSpecRequest
//
// @return UpdateCacheReserveSpecResponse
func (client *Client) UpdateCacheReserveSpec(request *UpdateCacheReserveSpecRequest) (_result *UpdateCacheReserveSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCacheReserveSpecResponse{}
	_body, _err := client.UpdateCacheReserveSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify cache configuration
//
// @param request - UpdateCacheRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCacheRuleResponse
func (client *Client) UpdateCacheRuleWithOptions(request *UpdateCacheRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateCacheRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdditionalCacheablePorts) {
		query["AdditionalCacheablePorts"] = request.AdditionalCacheablePorts
	}

	if !dara.IsNil(request.BrowserCacheMode) {
		query["BrowserCacheMode"] = request.BrowserCacheMode
	}

	if !dara.IsNil(request.BrowserCacheTtl) {
		query["BrowserCacheTtl"] = request.BrowserCacheTtl
	}

	if !dara.IsNil(request.BypassCache) {
		query["BypassCache"] = request.BypassCache
	}

	if !dara.IsNil(request.CacheDeceptionArmor) {
		query["CacheDeceptionArmor"] = request.CacheDeceptionArmor
	}

	if !dara.IsNil(request.CacheReserveEligibility) {
		query["CacheReserveEligibility"] = request.CacheReserveEligibility
	}

	if !dara.IsNil(request.CheckPresenceCookie) {
		query["CheckPresenceCookie"] = request.CheckPresenceCookie
	}

	if !dara.IsNil(request.CheckPresenceHeader) {
		query["CheckPresenceHeader"] = request.CheckPresenceHeader
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.EdgeCacheMode) {
		query["EdgeCacheMode"] = request.EdgeCacheMode
	}

	if !dara.IsNil(request.EdgeCacheTtl) {
		query["EdgeCacheTtl"] = request.EdgeCacheTtl
	}

	if !dara.IsNil(request.EdgeStatusCodeCacheTtl) {
		query["EdgeStatusCodeCacheTtl"] = request.EdgeStatusCodeCacheTtl
	}

	if !dara.IsNil(request.IncludeCookie) {
		query["IncludeCookie"] = request.IncludeCookie
	}

	if !dara.IsNil(request.IncludeHeader) {
		query["IncludeHeader"] = request.IncludeHeader
	}

	if !dara.IsNil(request.PostBodyCacheKey) {
		query["PostBodyCacheKey"] = request.PostBodyCacheKey
	}

	if !dara.IsNil(request.PostBodySizeLimit) {
		query["PostBodySizeLimit"] = request.PostBodySizeLimit
	}

	if !dara.IsNil(request.PostCache) {
		query["PostCache"] = request.PostCache
	}

	if !dara.IsNil(request.QueryString) {
		query["QueryString"] = request.QueryString
	}

	if !dara.IsNil(request.QueryStringMode) {
		query["QueryStringMode"] = request.QueryStringMode
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.ServeStale) {
		query["ServeStale"] = request.ServeStale
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SortQueryStringForCache) {
		query["SortQueryStringForCache"] = request.SortQueryStringForCache
	}

	if !dara.IsNil(request.UserDeviceType) {
		query["UserDeviceType"] = request.UserDeviceType
	}

	if !dara.IsNil(request.UserGeo) {
		query["UserGeo"] = request.UserGeo
	}

	if !dara.IsNil(request.UserLanguage) {
		query["UserLanguage"] = request.UserLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCacheRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCacheRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify cache configuration
//
// @param request - UpdateCacheRuleRequest
//
// @return UpdateCacheRuleResponse
func (client *Client) UpdateCacheRule(request *UpdateCacheRuleRequest) (_result *UpdateCacheRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCacheRuleResponse{}
	_body, _err := client.UpdateCacheRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the cache tag configuration of your website. You can call this operation when you need to specify tags in the Cache-Tag response header to use the purge by cache tag feature.
//
// @param request - UpdateCacheTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCacheTagResponse
func (client *Client) UpdateCacheTagWithOptions(request *UpdateCacheTagRequest, runtime *dara.RuntimeOptions) (_result *UpdateCacheTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaseInsensitive) {
		query["CaseInsensitive"] = request.CaseInsensitive
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCacheTag"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCacheTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the cache tag configuration of your website. You can call this operation when you need to specify tags in the Cache-Tag response header to use the purge by cache tag feature.
//
// @param request - UpdateCacheTagRequest
//
// @return UpdateCacheTagResponse
func (client *Client) UpdateCacheTag(request *UpdateCacheTagRequest) (_result *UpdateCacheTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCacheTagResponse{}
	_body, _err := client.UpdateCacheTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the CNAME flattening configuration of a website.
//
// @param request - UpdateCnameFlatteningRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCnameFlatteningResponse
func (client *Client) UpdateCnameFlatteningWithOptions(request *UpdateCnameFlatteningRequest, runtime *dara.RuntimeOptions) (_result *UpdateCnameFlatteningResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlattenMode) {
		query["FlattenMode"] = request.FlattenMode
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCnameFlattening"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCnameFlatteningResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the CNAME flattening configuration of a website.
//
// @param request - UpdateCnameFlatteningRequest
//
// @return UpdateCnameFlatteningResponse
func (client *Client) UpdateCnameFlattening(request *UpdateCnameFlatteningRequest) (_result *UpdateCnameFlatteningResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCnameFlatteningResponse{}
	_body, _err := client.UpdateCnameFlatteningWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify compression rule
//
// @param request - UpdateCompressionRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCompressionRuleResponse
func (client *Client) UpdateCompressionRuleWithOptions(request *UpdateCompressionRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateCompressionRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Brotli) {
		query["Brotli"] = request.Brotli
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.Gzip) {
		query["Gzip"] = request.Gzip
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.Zstd) {
		query["Zstd"] = request.Zstd
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCompressionRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCompressionRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify compression rule
//
// @param request - UpdateCompressionRuleRequest
//
// @return UpdateCompressionRuleResponse
func (client *Client) UpdateCompressionRule(request *UpdateCompressionRuleRequest) (_result *UpdateCompressionRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCompressionRuleResponse{}
	_body, _err := client.UpdateCompressionRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of the Chinese mainland network access optimization.
//
// @param request - UpdateCrossBorderOptimizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCrossBorderOptimizationResponse
func (client *Client) UpdateCrossBorderOptimizationWithOptions(request *UpdateCrossBorderOptimizationRequest, runtime *dara.RuntimeOptions) (_result *UpdateCrossBorderOptimizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCrossBorderOptimization"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCrossBorderOptimizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of the Chinese mainland network access optimization.
//
// @param request - UpdateCrossBorderOptimizationRequest
//
// @return UpdateCrossBorderOptimizationResponse
func (client *Client) UpdateCrossBorderOptimization(request *UpdateCrossBorderOptimizationRequest) (_result *UpdateCrossBorderOptimizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCrossBorderOptimizationResponse{}
	_body, _err := client.UpdateCrossBorderOptimizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a custom scenario-specific policy.
//
// @param request - UpdateCustomScenePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomScenePolicyResponse
func (client *Client) UpdateCustomScenePolicyWithOptions(request *UpdateCustomScenePolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomScenePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Objects) {
		query["Objects"] = request.Objects
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.SiteIds) {
		query["SiteIds"] = request.SiteIds
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomScenePolicy"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomScenePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a custom scenario-specific policy.
//
// @param request - UpdateCustomScenePolicyRequest
//
// @return UpdateCustomScenePolicyResponse
func (client *Client) UpdateCustomScenePolicy(request *UpdateCustomScenePolicyRequest) (_result *UpdateCustomScenePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCustomScenePolicyResponse{}
	_body, _err := client.UpdateCustomScenePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the development mode configuration of your website. If you enable Development Mode, all requests bypass caching components on POPs and are redirected to the origin server. This allows clients to retrieve the most recent resources on the origin server.
//
// @param request - UpdateDevelopmentModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDevelopmentModeResponse
func (client *Client) UpdateDevelopmentModeWithOptions(request *UpdateDevelopmentModeRequest, runtime *dara.RuntimeOptions) (_result *UpdateDevelopmentModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDevelopmentMode"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDevelopmentModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the development mode configuration of your website. If you enable Development Mode, all requests bypass caching components on POPs and are redirected to the origin server. This allows clients to retrieve the most recent resources on the origin server.
//
// @param request - UpdateDevelopmentModeRequest
//
// @return UpdateDevelopmentModeResponse
func (client *Client) UpdateDevelopmentMode(request *UpdateDevelopmentModeRequest) (_result *UpdateDevelopmentModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDevelopmentModeResponse{}
	_body, _err := client.UpdateDevelopmentModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the log collection configuration of a containerized application.
//
// @param request - UpdateEdgeContainerAppLogRiverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEdgeContainerAppLogRiverResponse
func (client *Client) UpdateEdgeContainerAppLogRiverWithOptions(request *UpdateEdgeContainerAppLogRiverRequest, runtime *dara.RuntimeOptions) (_result *UpdateEdgeContainerAppLogRiverResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.Stdout) {
		query["Stdout"] = request.Stdout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEdgeContainerAppLogRiver"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEdgeContainerAppLogRiverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the log collection configuration of a containerized application.
//
// @param request - UpdateEdgeContainerAppLogRiverRequest
//
// @return UpdateEdgeContainerAppLogRiverResponse
func (client *Client) UpdateEdgeContainerAppLogRiver(request *UpdateEdgeContainerAppLogRiverRequest) (_result *UpdateEdgeContainerAppLogRiverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateEdgeContainerAppLogRiverResponse{}
	_body, _err := client.UpdateEdgeContainerAppLogRiverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the resource reservation configuration of an edge container.
//
// @param tmpReq - UpdateEdgeContainerAppResourceReserveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEdgeContainerAppResourceReserveResponse
func (client *Client) UpdateEdgeContainerAppResourceReserveWithOptions(tmpReq *UpdateEdgeContainerAppResourceReserveRequest, runtime *dara.RuntimeOptions) (_result *UpdateEdgeContainerAppResourceReserveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateEdgeContainerAppResourceReserveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReserveSet) {
		request.ReserveSetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReserveSet, dara.String("ReserveSet"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DurationTime) {
		query["DurationTime"] = request.DurationTime
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.Forever) {
		query["Forever"] = request.Forever
	}

	if !dara.IsNil(request.ReserveSetShrink) {
		query["ReserveSet"] = request.ReserveSetShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEdgeContainerAppResourceReserve"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEdgeContainerAppResourceReserveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the resource reservation configuration of an edge container.
//
// @param request - UpdateEdgeContainerAppResourceReserveRequest
//
// @return UpdateEdgeContainerAppResourceReserveResponse
func (client *Client) UpdateEdgeContainerAppResourceReserve(request *UpdateEdgeContainerAppResourceReserveRequest) (_result *UpdateEdgeContainerAppResourceReserveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateEdgeContainerAppResourceReserveResponse{}
	_body, _err := client.UpdateEdgeContainerAppResourceReserveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the HTTP incoming request header modification rule.
//
// @param tmpReq - UpdateHttpIncomingRequestHeaderModificationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHttpIncomingRequestHeaderModificationRuleResponse
func (client *Client) UpdateHttpIncomingRequestHeaderModificationRuleWithOptions(tmpReq *UpdateHttpIncomingRequestHeaderModificationRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateHttpIncomingRequestHeaderModificationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RequestHeaderModification) {
		request.RequestHeaderModificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RequestHeaderModification, dara.String("RequestHeaderModification"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.RequestHeaderModificationShrink) {
		query["RequestHeaderModification"] = request.RequestHeaderModificationShrink
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHttpIncomingRequestHeaderModificationRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHttpIncomingRequestHeaderModificationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the HTTP incoming request header modification rule.
//
// @param request - UpdateHttpIncomingRequestHeaderModificationRuleRequest
//
// @return UpdateHttpIncomingRequestHeaderModificationRuleResponse
func (client *Client) UpdateHttpIncomingRequestHeaderModificationRule(request *UpdateHttpIncomingRequestHeaderModificationRuleRequest) (_result *UpdateHttpIncomingRequestHeaderModificationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateHttpIncomingRequestHeaderModificationRuleResponse{}
	_body, _err := client.UpdateHttpIncomingRequestHeaderModificationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of modifying HTTP response headers for a website.
//
// @param tmpReq - UpdateHttpIncomingResponseHeaderModificationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHttpIncomingResponseHeaderModificationRuleResponse
func (client *Client) UpdateHttpIncomingResponseHeaderModificationRuleWithOptions(tmpReq *UpdateHttpIncomingResponseHeaderModificationRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateHttpIncomingResponseHeaderModificationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResponseHeaderModification) {
		request.ResponseHeaderModificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResponseHeaderModification, dara.String("ResponseHeaderModification"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.ResponseHeaderModificationShrink) {
		query["ResponseHeaderModification"] = request.ResponseHeaderModificationShrink
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHttpIncomingResponseHeaderModificationRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHttpIncomingResponseHeaderModificationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of modifying HTTP response headers for a website.
//
// @param request - UpdateHttpIncomingResponseHeaderModificationRuleRequest
//
// @return UpdateHttpIncomingResponseHeaderModificationRuleResponse
func (client *Client) UpdateHttpIncomingResponseHeaderModificationRule(request *UpdateHttpIncomingResponseHeaderModificationRuleRequest) (_result *UpdateHttpIncomingResponseHeaderModificationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateHttpIncomingResponseHeaderModificationRuleResponse{}
	_body, _err := client.UpdateHttpIncomingResponseHeaderModificationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify HTTP Request Header Rules
//
// @param tmpReq - UpdateHttpRequestHeaderModificationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHttpRequestHeaderModificationRuleResponse
func (client *Client) UpdateHttpRequestHeaderModificationRuleWithOptions(tmpReq *UpdateHttpRequestHeaderModificationRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateHttpRequestHeaderModificationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateHttpRequestHeaderModificationRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RequestHeaderModification) {
		request.RequestHeaderModificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RequestHeaderModification, dara.String("RequestHeaderModification"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.RequestHeaderModificationShrink) {
		query["RequestHeaderModification"] = request.RequestHeaderModificationShrink
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHttpRequestHeaderModificationRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHttpRequestHeaderModificationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify HTTP Request Header Rules
//
// @param request - UpdateHttpRequestHeaderModificationRuleRequest
//
// @return UpdateHttpRequestHeaderModificationRuleResponse
func (client *Client) UpdateHttpRequestHeaderModificationRule(request *UpdateHttpRequestHeaderModificationRuleRequest) (_result *UpdateHttpRequestHeaderModificationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateHttpRequestHeaderModificationRuleResponse{}
	_body, _err := client.UpdateHttpRequestHeaderModificationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify HTTP response header rules
//
// @param tmpReq - UpdateHttpResponseHeaderModificationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHttpResponseHeaderModificationRuleResponse
func (client *Client) UpdateHttpResponseHeaderModificationRuleWithOptions(tmpReq *UpdateHttpResponseHeaderModificationRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateHttpResponseHeaderModificationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateHttpResponseHeaderModificationRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResponseHeaderModification) {
		request.ResponseHeaderModificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResponseHeaderModification, dara.String("ResponseHeaderModification"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.ResponseHeaderModificationShrink) {
		query["ResponseHeaderModification"] = request.ResponseHeaderModificationShrink
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHttpResponseHeaderModificationRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHttpResponseHeaderModificationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify HTTP response header rules
//
// @param request - UpdateHttpResponseHeaderModificationRuleRequest
//
// @return UpdateHttpResponseHeaderModificationRuleResponse
func (client *Client) UpdateHttpResponseHeaderModificationRule(request *UpdateHttpResponseHeaderModificationRuleRequest) (_result *UpdateHttpResponseHeaderModificationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateHttpResponseHeaderModificationRuleResponse{}
	_body, _err := client.UpdateHttpResponseHeaderModificationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify HTTPS Application Configuration
//
// @param request - UpdateHttpsApplicationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHttpsApplicationConfigurationResponse
func (client *Client) UpdateHttpsApplicationConfigurationWithOptions(request *UpdateHttpsApplicationConfigurationRequest, runtime *dara.RuntimeOptions) (_result *UpdateHttpsApplicationConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AltSvc) {
		query["AltSvc"] = request.AltSvc
	}

	if !dara.IsNil(request.AltSvcClear) {
		query["AltSvcClear"] = request.AltSvcClear
	}

	if !dara.IsNil(request.AltSvcMa) {
		query["AltSvcMa"] = request.AltSvcMa
	}

	if !dara.IsNil(request.AltSvcPersist) {
		query["AltSvcPersist"] = request.AltSvcPersist
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.Hsts) {
		query["Hsts"] = request.Hsts
	}

	if !dara.IsNil(request.HstsIncludeSubdomains) {
		query["HstsIncludeSubdomains"] = request.HstsIncludeSubdomains
	}

	if !dara.IsNil(request.HstsMaxAge) {
		query["HstsMaxAge"] = request.HstsMaxAge
	}

	if !dara.IsNil(request.HstsPreload) {
		query["HstsPreload"] = request.HstsPreload
	}

	if !dara.IsNil(request.HttpsForce) {
		query["HttpsForce"] = request.HttpsForce
	}

	if !dara.IsNil(request.HttpsForceCode) {
		query["HttpsForceCode"] = request.HttpsForceCode
	}

	if !dara.IsNil(request.HttpsNoSniDeny) {
		query["HttpsNoSniDeny"] = request.HttpsNoSniDeny
	}

	if !dara.IsNil(request.HttpsSniVerify) {
		query["HttpsSniVerify"] = request.HttpsSniVerify
	}

	if !dara.IsNil(request.HttpsSniWhitelist) {
		query["HttpsSniWhitelist"] = request.HttpsSniWhitelist
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHttpsApplicationConfiguration"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHttpsApplicationConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify HTTPS Application Configuration
//
// @param request - UpdateHttpsApplicationConfigurationRequest
//
// @return UpdateHttpsApplicationConfigurationResponse
func (client *Client) UpdateHttpsApplicationConfiguration(request *UpdateHttpsApplicationConfigurationRequest) (_result *UpdateHttpsApplicationConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateHttpsApplicationConfigurationResponse{}
	_body, _err := client.UpdateHttpsApplicationConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify HTTPS Basic Configuration
//
// @param request - UpdateHttpsBasicConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHttpsBasicConfigurationResponse
func (client *Client) UpdateHttpsBasicConfigurationWithOptions(request *UpdateHttpsBasicConfigurationRequest, runtime *dara.RuntimeOptions) (_result *UpdateHttpsBasicConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ciphersuite) {
		query["Ciphersuite"] = request.Ciphersuite
	}

	if !dara.IsNil(request.CiphersuiteGroup) {
		query["CiphersuiteGroup"] = request.CiphersuiteGroup
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.Http2) {
		query["Http2"] = request.Http2
	}

	if !dara.IsNil(request.Http3) {
		query["Http3"] = request.Http3
	}

	if !dara.IsNil(request.Https) {
		query["Https"] = request.Https
	}

	if !dara.IsNil(request.OcspStapling) {
		query["OcspStapling"] = request.OcspStapling
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.Tls10) {
		query["Tls10"] = request.Tls10
	}

	if !dara.IsNil(request.Tls11) {
		query["Tls11"] = request.Tls11
	}

	if !dara.IsNil(request.Tls12) {
		query["Tls12"] = request.Tls12
	}

	if !dara.IsNil(request.Tls13) {
		query["Tls13"] = request.Tls13
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHttpsBasicConfiguration"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHttpsBasicConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify HTTPS Basic Configuration
//
// @param request - UpdateHttpsBasicConfigurationRequest
//
// @return UpdateHttpsBasicConfigurationResponse
func (client *Client) UpdateHttpsBasicConfiguration(request *UpdateHttpsBasicConfigurationRequest) (_result *UpdateHttpsBasicConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateHttpsBasicConfigurationResponse{}
	_body, _err := client.UpdateHttpsBasicConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the IPv6 configuration of a website.
//
// @param request - UpdateIPv6Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIPv6Response
func (client *Client) UpdateIPv6WithOptions(request *UpdateIPv6Request, runtime *dara.RuntimeOptions) (_result *UpdateIPv6Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIPv6"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIPv6Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the IPv6 configuration of a website.
//
// @param request - UpdateIPv6Request
//
// @return UpdateIPv6Response
func (client *Client) UpdateIPv6(request *UpdateIPv6Request) (_result *UpdateIPv6Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIPv6Response{}
	_body, _err := client.UpdateIPv6WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Site Image Transformation Configuration
//
// @param request - UpdateImageTransformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateImageTransformResponse
func (client *Client) UpdateImageTransformWithOptions(request *UpdateImageTransformRequest, runtime *dara.RuntimeOptions) (_result *UpdateImageTransformResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateImageTransform"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateImageTransformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Site Image Transformation Configuration
//
// @param request - UpdateImageTransformRequest
//
// @return UpdateImageTransformResponse
func (client *Client) UpdateImageTransform(request *UpdateImageTransformRequest) (_result *UpdateImageTransformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateImageTransformResponse{}
	_body, _err := client.UpdateImageTransformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a custom list.
//
// @param tmpReq - UpdateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateListResponse
func (client *Client) UpdateListWithOptions(tmpReq *UpdateListRequest, runtime *dara.RuntimeOptions) (_result *UpdateListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Items) {
		request.ItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Items, dara.String("Items"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ItemsShrink) {
		body["Items"] = request.ItemsShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateList"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a custom list.
//
// @param request - UpdateListRequest
//
// @return UpdateListResponse
func (client *Client) UpdateList(request *UpdateListRequest) (_result *UpdateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateListResponse{}
	_body, _err := client.UpdateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Load Balancer
//
// Description:
//
// Through this interface, you can modify multiple configurations of the load balancer, including but not limited to the name of the load balancer, whether to enable acceleration, session persistence strategy, and various advanced settings related to traffic routing.	Notice: Changes to certain parameters may affect the stability of existing services, please operate with caution.
//
// @param tmpReq - UpdateLoadBalancerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLoadBalancerResponse
func (client *Client) UpdateLoadBalancerWithOptions(tmpReq *UpdateLoadBalancerRequest, runtime *dara.RuntimeOptions) (_result *UpdateLoadBalancerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateLoadBalancerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AdaptiveRouting) {
		request.AdaptiveRoutingShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdaptiveRouting, dara.String("AdaptiveRouting"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DefaultPools) {
		request.DefaultPoolsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DefaultPools, dara.String("DefaultPools"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Monitor) {
		request.MonitorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Monitor, dara.String("Monitor"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RandomSteering) {
		request.RandomSteeringShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RandomSteering, dara.String("RandomSteering"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Rules) {
		request.RulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rules, dara.String("Rules"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdaptiveRoutingShrink) {
		query["AdaptiveRouting"] = request.AdaptiveRoutingShrink
	}

	if !dara.IsNil(request.DefaultPoolsShrink) {
		query["DefaultPools"] = request.DefaultPoolsShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.FallbackPool) {
		query["FallbackPool"] = request.FallbackPool
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.MonitorShrink) {
		query["Monitor"] = request.MonitorShrink
	}

	if !dara.IsNil(request.RandomSteeringShrink) {
		query["RandomSteering"] = request.RandomSteeringShrink
	}

	if !dara.IsNil(request.RegionPools) {
		query["RegionPools"] = request.RegionPools
	}

	if !dara.IsNil(request.RulesShrink) {
		query["Rules"] = request.RulesShrink
	}

	if !dara.IsNil(request.SessionAffinity) {
		query["SessionAffinity"] = request.SessionAffinity
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SteeringPolicy) {
		query["SteeringPolicy"] = request.SteeringPolicy
	}

	if !dara.IsNil(request.SubRegionPools) {
		query["SubRegionPools"] = request.SubRegionPools
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLoadBalancer"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Load Balancer
//
// Description:
//
// Through this interface, you can modify multiple configurations of the load balancer, including but not limited to the name of the load balancer, whether to enable acceleration, session persistence strategy, and various advanced settings related to traffic routing.	Notice: Changes to certain parameters may affect the stability of existing services, please operate with caution.
//
// @param request - UpdateLoadBalancerRequest
//
// @return UpdateLoadBalancerResponse
func (client *Client) UpdateLoadBalancer(request *UpdateLoadBalancerRequest) (_result *UpdateLoadBalancerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLoadBalancerResponse{}
	_body, _err := client.UpdateLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of managed transforms for your website.
//
// @param request - UpdateManagedTransformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateManagedTransformResponse
func (client *Client) UpdateManagedTransformWithOptions(request *UpdateManagedTransformRequest, runtime *dara.RuntimeOptions) (_result *UpdateManagedTransformResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddClientGeolocationHeader) {
		query["AddClientGeolocationHeader"] = request.AddClientGeolocationHeader
	}

	if !dara.IsNil(request.AddRealClientIpHeader) {
		query["AddRealClientIpHeader"] = request.AddRealClientIpHeader
	}

	if !dara.IsNil(request.RealClientIpHeaderName) {
		query["RealClientIpHeaderName"] = request.RealClientIpHeaderName
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateManagedTransform"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateManagedTransformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of managed transforms for your website.
//
// @param request - UpdateManagedTransformRequest
//
// @return UpdateManagedTransformResponse
func (client *Client) UpdateManagedTransform(request *UpdateManagedTransformRequest) (_result *UpdateManagedTransformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateManagedTransformResponse{}
	_body, _err := client.UpdateManagedTransformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify network optimization configuration
//
// @param request - UpdateNetworkOptimizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNetworkOptimizationResponse
func (client *Client) UpdateNetworkOptimizationWithOptions(request *UpdateNetworkOptimizationRequest, runtime *dara.RuntimeOptions) (_result *UpdateNetworkOptimizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.Grpc) {
		query["Grpc"] = request.Grpc
	}

	if !dara.IsNil(request.Http2Origin) {
		query["Http2Origin"] = request.Http2Origin
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SmartRouting) {
		query["SmartRouting"] = request.SmartRouting
	}

	if !dara.IsNil(request.UploadMaxFilesize) {
		query["UploadMaxFilesize"] = request.UploadMaxFilesize
	}

	if !dara.IsNil(request.Websocket) {
		query["Websocket"] = request.Websocket
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNetworkOptimization"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNetworkOptimizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify network optimization configuration
//
// @param request - UpdateNetworkOptimizationRequest
//
// @return UpdateNetworkOptimizationResponse
func (client *Client) UpdateNetworkOptimization(request *UpdateNetworkOptimizationRequest) (_result *UpdateNetworkOptimizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNetworkOptimizationResponse{}
	_body, _err := client.UpdateNetworkOptimizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify the Monitor
//
// @param tmpReq - UpdateOriginPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOriginPoolResponse
func (client *Client) UpdateOriginPoolWithOptions(tmpReq *UpdateOriginPoolRequest, runtime *dara.RuntimeOptions) (_result *UpdateOriginPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateOriginPoolShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Origins) {
		request.OriginsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Origins, dara.String("Origins"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OriginsShrink) {
		query["Origins"] = request.OriginsShrink
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOriginPool"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOriginPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify the Monitor
//
// @param request - UpdateOriginPoolRequest
//
// @return UpdateOriginPoolResponse
func (client *Client) UpdateOriginPool(request *UpdateOriginPoolRequest) (_result *UpdateOriginPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOriginPoolResponse{}
	_body, _err := client.UpdateOriginPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables IP convergence.
//
// @param request - UpdateOriginProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOriginProtectionResponse
func (client *Client) UpdateOriginProtectionWithOptions(request *UpdateOriginProtectionRequest, runtime *dara.RuntimeOptions) (_result *UpdateOriginProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoConfirmIPList) {
		query["AutoConfirmIPList"] = request.AutoConfirmIPList
	}

	if !dara.IsNil(request.OriginConverge) {
		query["OriginConverge"] = request.OriginConverge
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOriginProtection"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOriginProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables IP convergence.
//
// @param request - UpdateOriginProtectionRequest
//
// @return UpdateOriginProtectionResponse
func (client *Client) UpdateOriginProtection(request *UpdateOriginProtectionRequest) (_result *UpdateOriginProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOriginProtectionResponse{}
	_body, _err := client.UpdateOriginProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the IP whitelist for origin protection used by a website to the latest version.
//
// @param request - UpdateOriginProtectionIpWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOriginProtectionIpWhiteListResponse
func (client *Client) UpdateOriginProtectionIpWhiteListWithOptions(request *UpdateOriginProtectionIpWhiteListRequest, runtime *dara.RuntimeOptions) (_result *UpdateOriginProtectionIpWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOriginProtectionIpWhiteList"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOriginProtectionIpWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the IP whitelist for origin protection used by a website to the latest version.
//
// @param request - UpdateOriginProtectionIpWhiteListRequest
//
// @return UpdateOriginProtectionIpWhiteListResponse
func (client *Client) UpdateOriginProtectionIpWhiteList(request *UpdateOriginProtectionIpWhiteListRequest) (_result *UpdateOriginProtectionIpWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOriginProtectionIpWhiteListResponse{}
	_body, _err := client.UpdateOriginProtectionIpWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Origin Rule Configuration for Site
//
// @param request - UpdateOriginRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOriginRuleResponse
func (client *Client) UpdateOriginRuleWithOptions(request *UpdateOriginRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateOriginRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.DnsRecord) {
		query["DnsRecord"] = request.DnsRecord
	}

	if !dara.IsNil(request.Follow302Enable) {
		query["Follow302Enable"] = request.Follow302Enable
	}

	if !dara.IsNil(request.Follow302MaxTries) {
		query["Follow302MaxTries"] = request.Follow302MaxTries
	}

	if !dara.IsNil(request.Follow302RetainArgs) {
		query["Follow302RetainArgs"] = request.Follow302RetainArgs
	}

	if !dara.IsNil(request.Follow302RetainHeader) {
		query["Follow302RetainHeader"] = request.Follow302RetainHeader
	}

	if !dara.IsNil(request.Follow302TargetHost) {
		query["Follow302TargetHost"] = request.Follow302TargetHost
	}

	if !dara.IsNil(request.OriginHost) {
		query["OriginHost"] = request.OriginHost
	}

	if !dara.IsNil(request.OriginHttpPort) {
		query["OriginHttpPort"] = request.OriginHttpPort
	}

	if !dara.IsNil(request.OriginHttpsPort) {
		query["OriginHttpsPort"] = request.OriginHttpsPort
	}

	if !dara.IsNil(request.OriginMtls) {
		query["OriginMtls"] = request.OriginMtls
	}

	if !dara.IsNil(request.OriginReadTimeout) {
		query["OriginReadTimeout"] = request.OriginReadTimeout
	}

	if !dara.IsNil(request.OriginScheme) {
		query["OriginScheme"] = request.OriginScheme
	}

	if !dara.IsNil(request.OriginSni) {
		query["OriginSni"] = request.OriginSni
	}

	if !dara.IsNil(request.OriginVerify) {
		query["OriginVerify"] = request.OriginVerify
	}

	if !dara.IsNil(request.Range) {
		query["Range"] = request.Range
	}

	if !dara.IsNil(request.RangeChunkSize) {
		query["RangeChunkSize"] = request.RangeChunkSize
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOriginRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOriginRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Origin Rule Configuration for Site
//
// @param request - UpdateOriginRuleRequest
//
// @return UpdateOriginRuleResponse
func (client *Client) UpdateOriginRule(request *UpdateOriginRuleRequest) (_result *UpdateOriginRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOriginRuleResponse{}
	_body, _err := client.UpdateOriginRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a custom error page, such as the name, description, content type, and content of the page.
//
// @param request - UpdatePageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePageResponse
func (client *Client) UpdatePageWithOptions(request *UpdatePageRequest, runtime *dara.RuntimeOptions) (_result *UpdatePageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.ContentType) {
		body["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePage"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a custom error page, such as the name, description, content type, and content of the page.
//
// @param request - UpdatePageRequest
//
// @return UpdatePageResponse
func (client *Client) UpdatePage(request *UpdatePageRequest) (_result *UpdatePageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePageResponse{}
	_body, _err := client.UpdatePageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 套餐变配
//
// @param request - UpdateRatePlanSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRatePlanSpecResponse
func (client *Client) UpdateRatePlanSpecWithOptions(request *UpdateRatePlanSpecRequest, runtime *dara.RuntimeOptions) (_result *UpdateRatePlanSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.TargetPlanCode) {
		query["TargetPlanCode"] = request.TargetPlanCode
	}

	if !dara.IsNil(request.TargetPlanName) {
		query["TargetPlanName"] = request.TargetPlanName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRatePlanSpec"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRatePlanSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 套餐变配
//
// @param request - UpdateRatePlanSpecRequest
//
// @return UpdateRatePlanSpecResponse
func (client *Client) UpdateRatePlanSpec(request *UpdateRatePlanSpecRequest) (_result *UpdateRatePlanSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRatePlanSpecResponse{}
	_body, _err := client.UpdateRatePlanSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates multiple types of DNS records and origin authentication configurations.
//
// Description:
//
// This operation allows you to update multiple types of DNS records, including but not limited to A/AAAA, CNAME, NS, MX, TXT, CAA, SRV, and URI. You can modify the record content by providing the necessary fields such as Value, Priority, and Flag. For origins added in CNAME records such as OSS and S3, the API enables you to configure authentication details to ensure secure access.
//
// ### [](#)Usage notes
//
//   - The record value (Value) must match the record type. For example, the CNAME record should correspond to the target domain name.
//
//   - You must specify a priority (Priority) for some record types, such as MX and SRV.
//
//   - You must specify specific fields such as Flag and Tag for CAA records.
//
//   - When you update security records such as CERT and SSHFP, you must accurately set fields such as Type and Algorithm.
//
//   - If your origin type is OSS or S3, configure the authentication details in AuthConf based on the permissions.
//
// @param tmpReq - UpdateRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecordResponse
func (client *Client) UpdateRecordWithOptions(tmpReq *UpdateRecordRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuthConf) {
		request.AuthConfShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthConf, dara.String("AuthConf"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Data) {
		request.DataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Data, dara.String("Data"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthConfShrink) {
		query["AuthConf"] = request.AuthConfShrink
	}

	if !dara.IsNil(request.BizName) {
		query["BizName"] = request.BizName
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DataShrink) {
		query["Data"] = request.DataShrink
	}

	if !dara.IsNil(request.HostPolicy) {
		query["HostPolicy"] = request.HostPolicy
	}

	if !dara.IsNil(request.Proxied) {
		query["Proxied"] = request.Proxied
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecord"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates multiple types of DNS records and origin authentication configurations.
//
// Description:
//
// This operation allows you to update multiple types of DNS records, including but not limited to A/AAAA, CNAME, NS, MX, TXT, CAA, SRV, and URI. You can modify the record content by providing the necessary fields such as Value, Priority, and Flag. For origins added in CNAME records such as OSS and S3, the API enables you to configure authentication details to ensure secure access.
//
// ### [](#)Usage notes
//
//   - The record value (Value) must match the record type. For example, the CNAME record should correspond to the target domain name.
//
//   - You must specify a priority (Priority) for some record types, such as MX and SRV.
//
//   - You must specify specific fields such as Flag and Tag for CAA records.
//
//   - When you update security records such as CERT and SSHFP, you must accurately set fields such as Type and Algorithm.
//
//   - If your origin type is OSS or S3, configure the authentication details in AuthConf based on the permissions.
//
// @param request - UpdateRecordRequest
//
// @return UpdateRecordResponse
func (client *Client) UpdateRecord(request *UpdateRecordRequest) (_result *UpdateRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecordResponse{}
	_body, _err := client.UpdateRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Redirect Rule
//
// @param request - UpdateRedirectRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRedirectRuleResponse
func (client *Client) UpdateRedirectRuleWithOptions(request *UpdateRedirectRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateRedirectRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.ReserveQueryString) {
		query["ReserveQueryString"] = request.ReserveQueryString
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.StatusCode) {
		query["StatusCode"] = request.StatusCode
	}

	if !dara.IsNil(request.TargetUrl) {
		query["TargetUrl"] = request.TargetUrl
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRedirectRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRedirectRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Redirect Rule
//
// @param request - UpdateRedirectRuleRequest
//
// @return UpdateRedirectRuleResponse
func (client *Client) UpdateRedirectRule(request *UpdateRedirectRuleRequest) (_result *UpdateRedirectRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRedirectRuleResponse{}
	_body, _err := client.UpdateRedirectRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Rewrite URL Rule
//
// @param request - UpdateRewriteUrlRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRewriteUrlRuleResponse
func (client *Client) UpdateRewriteUrlRuleWithOptions(request *UpdateRewriteUrlRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateRewriteUrlRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.QueryString) {
		query["QueryString"] = request.QueryString
	}

	if !dara.IsNil(request.RewriteQueryStringType) {
		query["RewriteQueryStringType"] = request.RewriteQueryStringType
	}

	if !dara.IsNil(request.RewriteUriType) {
		query["RewriteUriType"] = request.RewriteUriType
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.Uri) {
		query["Uri"] = request.Uri
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRewriteUrlRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRewriteUrlRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Rewrite URL Rule
//
// @param request - UpdateRewriteUrlRuleRequest
//
// @return UpdateRewriteUrlRuleResponse
func (client *Client) UpdateRewriteUrlRule(request *UpdateRewriteUrlRuleRequest) (_result *UpdateRewriteUrlRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRewriteUrlRuleResponse{}
	_body, _err := client.UpdateRewriteUrlRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of a routine.
//
// @param request - UpdateRoutineConfigDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRoutineConfigDescriptionResponse
func (client *Client) UpdateRoutineConfigDescriptionWithOptions(request *UpdateRoutineConfigDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateRoutineConfigDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRoutineConfigDescription"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRoutineConfigDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a routine.
//
// @param request - UpdateRoutineConfigDescriptionRequest
//
// @return UpdateRoutineConfigDescriptionResponse
func (client *Client) UpdateRoutineConfigDescription(request *UpdateRoutineConfigDescriptionRequest) (_result *UpdateRoutineConfigDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRoutineConfigDescriptionResponse{}
	_body, _err := client.UpdateRoutineConfigDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the route configuration of an edge function.
//
// @param request - UpdateRoutineRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRoutineRouteResponse
func (client *Client) UpdateRoutineRouteWithOptions(request *UpdateRoutineRouteRequest, runtime *dara.RuntimeOptions) (_result *UpdateRoutineRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bypass) {
		query["Bypass"] = request.Bypass
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.Fallback) {
		query["Fallback"] = request.Fallback
	}

	if !dara.IsNil(request.RouteEnable) {
		query["RouteEnable"] = request.RouteEnable
	}

	if !dara.IsNil(request.RouteName) {
		query["RouteName"] = request.RouteName
	}

	if !dara.IsNil(request.RoutineName) {
		query["RoutineName"] = request.RoutineName
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRoutineRoute"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRoutineRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the route configuration of an edge function.
//
// @param request - UpdateRoutineRouteRequest
//
// @return UpdateRoutineRouteResponse
func (client *Client) UpdateRoutineRoute(request *UpdateRoutineRouteRequest) (_result *UpdateRoutineRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRoutineRouteResponse{}
	_body, _err := client.UpdateRoutineRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a scheduled prefetch plan based on the plan ID.
//
// @param request - UpdateScheduledPreloadExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScheduledPreloadExecutionResponse
func (client *Client) UpdateScheduledPreloadExecutionWithOptions(request *UpdateScheduledPreloadExecutionRequest, runtime *dara.RuntimeOptions) (_result *UpdateScheduledPreloadExecutionResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		body["Interval"] = request.Interval
	}

	if !dara.IsNil(request.SliceLen) {
		body["SliceLen"] = request.SliceLen
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateScheduledPreloadExecution"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateScheduledPreloadExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a scheduled prefetch plan based on the plan ID.
//
// @param request - UpdateScheduledPreloadExecutionRequest
//
// @return UpdateScheduledPreloadExecutionResponse
func (client *Client) UpdateScheduledPreloadExecution(request *UpdateScheduledPreloadExecutionRequest) (_result *UpdateScheduledPreloadExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateScheduledPreloadExecutionResponse{}
	_body, _err := client.UpdateScheduledPreloadExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the search engine crawler configuration for a website.
//
// @param request - UpdateSeoBypassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSeoBypassResponse
func (client *Client) UpdateSeoBypassWithOptions(request *UpdateSeoBypassRequest, runtime *dara.RuntimeOptions) (_result *UpdateSeoBypassResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSeoBypass"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSeoBypassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the search engine crawler configuration for a website.
//
// @param request - UpdateSeoBypassRequest
//
// @return UpdateSeoBypassResponse
func (client *Client) UpdateSeoBypass(request *UpdateSeoBypassRequest) (_result *UpdateSeoBypassResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSeoBypassResponse{}
	_body, _err := client.UpdateSeoBypassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Converts the DNS setup option of a website.
//
// Description:
//
// When you change the DNS setup of a website from NS to CNAME, note the following prerequisites:
//
//   - The website only has proxied A/AAAA and CNAME records.
//
//   - The DNS passthrough mode and custom nameserver features are not enabled for the website.
//
// @param request - UpdateSiteAccessTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSiteAccessTypeResponse
func (client *Client) UpdateSiteAccessTypeWithOptions(request *UpdateSiteAccessTypeRequest, runtime *dara.RuntimeOptions) (_result *UpdateSiteAccessTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessType) {
		query["AccessType"] = request.AccessType
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSiteAccessType"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSiteAccessTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Converts the DNS setup option of a website.
//
// Description:
//
// When you change the DNS setup of a website from NS to CNAME, note the following prerequisites:
//
//   - The website only has proxied A/AAAA and CNAME records.
//
//   - The DNS passthrough mode and custom nameserver features are not enabled for the website.
//
// @param request - UpdateSiteAccessTypeRequest
//
// @return UpdateSiteAccessTypeResponse
func (client *Client) UpdateSiteAccessType(request *UpdateSiteAccessTypeRequest) (_result *UpdateSiteAccessTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSiteAccessTypeResponse{}
	_body, _err := client.UpdateSiteAccessTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the service location for a single website. This updates the acceleration configuration of the website to adapt to changes in traffic distribution, and improve user experience in specific regions.
//
// @param request - UpdateSiteCoverageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSiteCoverageResponse
func (client *Client) UpdateSiteCoverageWithOptions(request *UpdateSiteCoverageRequest, runtime *dara.RuntimeOptions) (_result *UpdateSiteCoverageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Coverage) {
		query["Coverage"] = request.Coverage
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSiteCoverage"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSiteCoverageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the service location for a single website. This updates the acceleration configuration of the website to adapt to changes in traffic distribution, and improve user experience in specific regions.
//
// @param request - UpdateSiteCoverageRequest
//
// @return UpdateSiteCoverageResponse
func (client *Client) UpdateSiteCoverage(request *UpdateSiteCoverageRequest) (_result *UpdateSiteCoverageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSiteCoverageResponse{}
	_body, _err := client.UpdateSiteCoverageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of custom request header, response header, and cookie fields that are used to capture logs of a website.
//
// @param tmpReq - UpdateSiteCustomLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSiteCustomLogResponse
func (client *Client) UpdateSiteCustomLogWithOptions(tmpReq *UpdateSiteCustomLogRequest, runtime *dara.RuntimeOptions) (_result *UpdateSiteCustomLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSiteCustomLogShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Cookies) {
		request.CookiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Cookies, dara.String("Cookies"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RequestHeaders) {
		request.RequestHeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RequestHeaders, dara.String("RequestHeaders"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResponseHeaders) {
		request.ResponseHeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResponseHeaders, dara.String("ResponseHeaders"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CookiesShrink) {
		body["Cookies"] = request.CookiesShrink
	}

	if !dara.IsNil(request.RequestHeadersShrink) {
		body["RequestHeaders"] = request.RequestHeadersShrink
	}

	if !dara.IsNil(request.ResponseHeadersShrink) {
		body["ResponseHeaders"] = request.ResponseHeadersShrink
	}

	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSiteCustomLog"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSiteCustomLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of custom request header, response header, and cookie fields that are used to capture logs of a website.
//
// @param request - UpdateSiteCustomLogRequest
//
// @return UpdateSiteCustomLogResponse
func (client *Client) UpdateSiteCustomLog(request *UpdateSiteCustomLogRequest) (_result *UpdateSiteCustomLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSiteCustomLogResponse{}
	_body, _err := client.UpdateSiteCustomLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a real-time log delivery task.
//
// @param request - UpdateSiteDeliveryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSiteDeliveryTaskResponse
func (client *Client) UpdateSiteDeliveryTaskWithOptions(request *UpdateSiteDeliveryTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateSiteDeliveryTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		body["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.DiscardRate) {
		body["DiscardRate"] = request.DiscardRate
	}

	if !dara.IsNil(request.FieldName) {
		body["FieldName"] = request.FieldName
	}

	if !dara.IsNil(request.FilterVer) {
		body["FilterVer"] = request.FilterVer
	}

	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSiteDeliveryTask"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSiteDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a real-time log delivery task.
//
// @param request - UpdateSiteDeliveryTaskRequest
//
// @return UpdateSiteDeliveryTaskResponse
func (client *Client) UpdateSiteDeliveryTask(request *UpdateSiteDeliveryTaskRequest) (_result *UpdateSiteDeliveryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSiteDeliveryTaskResponse{}
	_body, _err := client.UpdateSiteDeliveryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the status of a real-time log delivery task.
//
// @param request - UpdateSiteDeliveryTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSiteDeliveryTaskStatusResponse
func (client *Client) UpdateSiteDeliveryTaskStatusWithOptions(request *UpdateSiteDeliveryTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateSiteDeliveryTaskStatusResponse, _err error) {
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
		Action:      dara.String("UpdateSiteDeliveryTaskStatus"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSiteDeliveryTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the status of a real-time log delivery task.
//
// @param request - UpdateSiteDeliveryTaskStatusRequest
//
// @return UpdateSiteDeliveryTaskStatusResponse
func (client *Client) UpdateSiteDeliveryTaskStatus(request *UpdateSiteDeliveryTaskStatusRequest) (_result *UpdateSiteDeliveryTaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSiteDeliveryTaskStatusResponse{}
	_body, _err := client.UpdateSiteDeliveryTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the site hold configuration of a website. After you enable site hold, other accounts cannot add your website domain or its subdomains to ESA.
//
// @param request - UpdateSiteNameExclusiveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSiteNameExclusiveResponse
func (client *Client) UpdateSiteNameExclusiveWithOptions(request *UpdateSiteNameExclusiveRequest, runtime *dara.RuntimeOptions) (_result *UpdateSiteNameExclusiveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSiteNameExclusive"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSiteNameExclusiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the site hold configuration of a website. After you enable site hold, other accounts cannot add your website domain or its subdomains to ESA.
//
// @param request - UpdateSiteNameExclusiveRequest
//
// @return UpdateSiteNameExclusiveResponse
func (client *Client) UpdateSiteNameExclusive(request *UpdateSiteNameExclusiveRequest) (_result *UpdateSiteNameExclusiveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSiteNameExclusiveResponse{}
	_body, _err := client.UpdateSiteNameExclusiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the ESA proxy configuration of a website.
//
// @param request - UpdateSitePauseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSitePauseResponse
func (client *Client) UpdateSitePauseWithOptions(request *UpdateSitePauseRequest, runtime *dara.RuntimeOptions) (_result *UpdateSitePauseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Paused) {
		query["Paused"] = request.Paused
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSitePause"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSitePauseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the ESA proxy configuration of a website.
//
// @param request - UpdateSitePauseRequest
//
// @return UpdateSitePauseResponse
func (client *Client) UpdateSitePause(request *UpdateSitePauseRequest) (_result *UpdateSitePauseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSitePauseResponse{}
	_body, _err := client.UpdateSitePauseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the custom nameserver names for a single website.
//
// @param request - UpdateSiteVanityNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSiteVanityNSResponse
func (client *Client) UpdateSiteVanityNSWithOptions(request *UpdateSiteVanityNSRequest, runtime *dara.RuntimeOptions) (_result *UpdateSiteVanityNSResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.VanityNSList) {
		query["VanityNSList"] = request.VanityNSList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSiteVanityNS"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSiteVanityNSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the custom nameserver names for a single website.
//
// @param request - UpdateSiteVanityNSRequest
//
// @return UpdateSiteVanityNSResponse
func (client *Client) UpdateSiteVanityNS(request *UpdateSiteVanityNSRequest) (_result *UpdateSiteVanityNSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSiteVanityNSResponse{}
	_body, _err := client.UpdateSiteVanityNSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the tiered cache configuration of your website.
//
// @param request - UpdateTieredCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTieredCacheResponse
func (client *Client) UpdateTieredCacheWithOptions(request *UpdateTieredCacheRequest, runtime *dara.RuntimeOptions) (_result *UpdateTieredCacheResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CacheArchitectureMode) {
		query["CacheArchitectureMode"] = request.CacheArchitectureMode
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTieredCache"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTieredCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the tiered cache configuration of your website.
//
// @param request - UpdateTieredCacheRequest
//
// @return UpdateTieredCacheResponse
func (client *Client) UpdateTieredCache(request *UpdateTieredCacheRequest) (_result *UpdateTieredCacheResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTieredCacheResponse{}
	_body, _err := client.UpdateTieredCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Transport Layer Application
//
// @param tmpReq - UpdateTransportLayerApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransportLayerApplicationResponse
func (client *Client) UpdateTransportLayerApplicationWithOptions(tmpReq *UpdateTransportLayerApplicationRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransportLayerApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTransportLayerApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Rules) {
		request.RulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rules, dara.String("Rules"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.CrossBorderOptimization) {
		query["CrossBorderOptimization"] = request.CrossBorderOptimization
	}

	if !dara.IsNil(request.IpAccessRule) {
		query["IpAccessRule"] = request.IpAccessRule
	}

	if !dara.IsNil(request.Ipv6) {
		query["Ipv6"] = request.Ipv6
	}

	if !dara.IsNil(request.RulesShrink) {
		query["Rules"] = request.RulesShrink
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.StaticIp) {
		query["StaticIp"] = request.StaticIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTransportLayerApplication"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTransportLayerApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Transport Layer Application
//
// @param request - UpdateTransportLayerApplicationRequest
//
// @return UpdateTransportLayerApplicationResponse
func (client *Client) UpdateTransportLayerApplication(request *UpdateTransportLayerApplicationRequest) (_result *UpdateTransportLayerApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTransportLayerApplicationResponse{}
	_body, _err := client.UpdateTransportLayerApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the webpage monitoring configuration.
//
// @param request - UpdateUrlObservationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUrlObservationResponse
func (client *Client) UpdateUrlObservationWithOptions(request *UpdateUrlObservationRequest, runtime *dara.RuntimeOptions) (_result *UpdateUrlObservationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.SdkType) {
		query["SdkType"] = request.SdkType
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUrlObservation"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUrlObservationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the webpage monitoring configuration.
//
// @param request - UpdateUrlObservationRequest
//
// @return UpdateUrlObservationResponse
func (client *Client) UpdateUrlObservation(request *UpdateUrlObservationRequest) (_result *UpdateUrlObservationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUrlObservationResponse{}
	_body, _err := client.UpdateUrlObservationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a delivery task, including the task name, log field, log category, and discard rate.
//
// @param request - UpdateUserDeliveryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserDeliveryTaskResponse
func (client *Client) UpdateUserDeliveryTaskWithOptions(request *UpdateUserDeliveryTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserDeliveryTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		body["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.Details) {
		body["Details"] = request.Details
	}

	if !dara.IsNil(request.DiscardRate) {
		body["DiscardRate"] = request.DiscardRate
	}

	if !dara.IsNil(request.FieldName) {
		body["FieldName"] = request.FieldName
	}

	if !dara.IsNil(request.FilterVer) {
		body["FilterVer"] = request.FilterVer
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserDeliveryTask"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a delivery task, including the task name, log field, log category, and discard rate.
//
// @param request - UpdateUserDeliveryTaskRequest
//
// @return UpdateUserDeliveryTaskResponse
func (client *Client) UpdateUserDeliveryTask(request *UpdateUserDeliveryTaskRequest) (_result *UpdateUserDeliveryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserDeliveryTaskResponse{}
	_body, _err := client.UpdateUserDeliveryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the status of a delivery task in your Alibaba Cloud account.
//
// Description:
//
// ## [](#)
//
// You can call this operation to enable or disable a delivery task by using TaskName and Method. The response includes the most recent status and operation result details of the task.
//
// @param request - UpdateUserDeliveryTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserDeliveryTaskStatusResponse
func (client *Client) UpdateUserDeliveryTaskStatusWithOptions(request *UpdateUserDeliveryTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserDeliveryTaskStatusResponse, _err error) {
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
		Action:      dara.String("UpdateUserDeliveryTaskStatus"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserDeliveryTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the status of a delivery task in your Alibaba Cloud account.
//
// Description:
//
// ## [](#)
//
// You can call this operation to enable or disable a delivery task by using TaskName and Method. The response includes the most recent status and operation result details of the task.
//
// @param request - UpdateUserDeliveryTaskStatusRequest
//
// @return UpdateUserDeliveryTaskStatusResponse
func (client *Client) UpdateUserDeliveryTaskStatus(request *UpdateUserDeliveryTaskStatusRequest) (_result *UpdateUserDeliveryTaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserDeliveryTaskStatusResponse{}
	_body, _err := client.UpdateUserDeliveryTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the video processing configuration of the site.
//
// @param request - UpdateVideoProcessingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVideoProcessingResponse
func (client *Client) UpdateVideoProcessingWithOptions(request *UpdateVideoProcessingRequest, runtime *dara.RuntimeOptions) (_result *UpdateVideoProcessingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.FlvSeekEnd) {
		query["FlvSeekEnd"] = request.FlvSeekEnd
	}

	if !dara.IsNil(request.FlvSeekStart) {
		query["FlvSeekStart"] = request.FlvSeekStart
	}

	if !dara.IsNil(request.FlvVideoSeekMode) {
		query["FlvVideoSeekMode"] = request.FlvVideoSeekMode
	}

	if !dara.IsNil(request.Mp4SeekEnd) {
		query["Mp4SeekEnd"] = request.Mp4SeekEnd
	}

	if !dara.IsNil(request.Mp4SeekStart) {
		query["Mp4SeekStart"] = request.Mp4SeekStart
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.VideoSeekEnable) {
		query["VideoSeekEnable"] = request.VideoSeekEnable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVideoProcessing"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVideoProcessingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the video processing configuration of the site.
//
// @param request - UpdateVideoProcessingRequest
//
// @return UpdateVideoProcessingResponse
func (client *Client) UpdateVideoProcessing(request *UpdateVideoProcessingRequest) (_result *UpdateVideoProcessingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateVideoProcessingResponse{}
	_body, _err := client.UpdateVideoProcessingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update WAF Rule Page
//
// @param tmpReq - UpdateWafRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWafRuleResponse
func (client *Client) UpdateWafRuleWithOptions(tmpReq *UpdateWafRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateWafRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateWafRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Config) {
		request.ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Config, dara.String("Config"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigShrink) {
		body["Config"] = request.ConfigShrink
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Position) {
		body["Position"] = request.Position
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWafRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWafRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update WAF Rule Page
//
// @param request - UpdateWafRuleRequest
//
// @return UpdateWafRuleResponse
func (client *Client) UpdateWafRule(request *UpdateWafRuleRequest) (_result *UpdateWafRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWafRuleResponse{}
	_body, _err := client.UpdateWafRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update WAF Ruleset
//
// @param request - UpdateWafRulesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWafRulesetResponse
func (client *Client) UpdateWafRulesetWithOptions(request *UpdateWafRulesetRequest, runtime *dara.RuntimeOptions) (_result *UpdateWafRulesetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWafRuleset"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWafRulesetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update WAF Ruleset
//
// @param request - UpdateWafRulesetRequest
//
// @return UpdateWafRulesetResponse
func (client *Client) UpdateWafRuleset(request *UpdateWafRulesetRequest) (_result *UpdateWafRulesetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWafRulesetResponse{}
	_body, _err := client.UpdateWafRulesetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a waiting room.
//
// @param tmpReq - UpdateWaitingRoomRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWaitingRoomResponse
func (client *Client) UpdateWaitingRoomWithOptions(tmpReq *UpdateWaitingRoomRequest, runtime *dara.RuntimeOptions) (_result *UpdateWaitingRoomResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateWaitingRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HostNameAndPath) {
		request.HostNameAndPathShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HostNameAndPath, dara.String("HostNameAndPath"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CookieName) {
		query["CookieName"] = request.CookieName
	}

	if !dara.IsNil(request.CustomPageHtml) {
		query["CustomPageHtml"] = request.CustomPageHtml
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisableSessionRenewalEnable) {
		query["DisableSessionRenewalEnable"] = request.DisableSessionRenewalEnable
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.HostNameAndPathShrink) {
		query["HostNameAndPath"] = request.HostNameAndPathShrink
	}

	if !dara.IsNil(request.JsonResponseEnable) {
		query["JsonResponseEnable"] = request.JsonResponseEnable
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NewUsersPerMinute) {
		query["NewUsersPerMinute"] = request.NewUsersPerMinute
	}

	if !dara.IsNil(request.QueueAllEnable) {
		query["QueueAllEnable"] = request.QueueAllEnable
	}

	if !dara.IsNil(request.QueuingMethod) {
		query["QueuingMethod"] = request.QueuingMethod
	}

	if !dara.IsNil(request.QueuingStatusCode) {
		query["QueuingStatusCode"] = request.QueuingStatusCode
	}

	if !dara.IsNil(request.SessionDuration) {
		query["SessionDuration"] = request.SessionDuration
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.TotalActiveUsers) {
		query["TotalActiveUsers"] = request.TotalActiveUsers
	}

	if !dara.IsNil(request.WaitingRoomId) {
		query["WaitingRoomId"] = request.WaitingRoomId
	}

	if !dara.IsNil(request.WaitingRoomType) {
		query["WaitingRoomType"] = request.WaitingRoomType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWaitingRoom"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWaitingRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a waiting room.
//
// @param request - UpdateWaitingRoomRequest
//
// @return UpdateWaitingRoomResponse
func (client *Client) UpdateWaitingRoom(request *UpdateWaitingRoomRequest) (_result *UpdateWaitingRoomResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWaitingRoomResponse{}
	_body, _err := client.UpdateWaitingRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a waiting room event.
//
// @param request - UpdateWaitingRoomEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWaitingRoomEventResponse
func (client *Client) UpdateWaitingRoomEventWithOptions(request *UpdateWaitingRoomEventRequest, runtime *dara.RuntimeOptions) (_result *UpdateWaitingRoomEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomPageHtml) {
		query["CustomPageHtml"] = request.CustomPageHtml
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisableSessionRenewalEnable) {
		query["DisableSessionRenewalEnable"] = request.DisableSessionRenewalEnable
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.JsonResponseEnable) {
		query["JsonResponseEnable"] = request.JsonResponseEnable
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NewUsersPerMinute) {
		query["NewUsersPerMinute"] = request.NewUsersPerMinute
	}

	if !dara.IsNil(request.PreQueueEnable) {
		query["PreQueueEnable"] = request.PreQueueEnable
	}

	if !dara.IsNil(request.PreQueueStartTime) {
		query["PreQueueStartTime"] = request.PreQueueStartTime
	}

	if !dara.IsNil(request.QueuingMethod) {
		query["QueuingMethod"] = request.QueuingMethod
	}

	if !dara.IsNil(request.QueuingStatusCode) {
		query["QueuingStatusCode"] = request.QueuingStatusCode
	}

	if !dara.IsNil(request.RandomPreQueueEnable) {
		query["RandomPreQueueEnable"] = request.RandomPreQueueEnable
	}

	if !dara.IsNil(request.SessionDuration) {
		query["SessionDuration"] = request.SessionDuration
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TotalActiveUsers) {
		query["TotalActiveUsers"] = request.TotalActiveUsers
	}

	if !dara.IsNil(request.WaitingRoomEventId) {
		query["WaitingRoomEventId"] = request.WaitingRoomEventId
	}

	if !dara.IsNil(request.WaitingRoomType) {
		query["WaitingRoomType"] = request.WaitingRoomType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWaitingRoomEvent"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWaitingRoomEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a waiting room event.
//
// @param request - UpdateWaitingRoomEventRequest
//
// @return UpdateWaitingRoomEventResponse
func (client *Client) UpdateWaitingRoomEvent(request *UpdateWaitingRoomEventRequest) (_result *UpdateWaitingRoomEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWaitingRoomEventResponse{}
	_body, _err := client.UpdateWaitingRoomEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Waiting Room Rule
//
// Description:
//
// This interface allows you to modify the rule settings of a specific waiting room in a site, including the rule name, enable status, and rule content, etc.
//
// @param request - UpdateWaitingRoomRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWaitingRoomRuleResponse
func (client *Client) UpdateWaitingRoomRuleWithOptions(request *UpdateWaitingRoomRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateWaitingRoomRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.RuleEnable) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.WaitingRoomRuleId) {
		query["WaitingRoomRuleId"] = request.WaitingRoomRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWaitingRoomRule"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWaitingRoomRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Waiting Room Rule
//
// Description:
//
// This interface allows you to modify the rule settings of a specific waiting room in a site, including the rule name, enable status, and rule content, etc.
//
// @param request - UpdateWaitingRoomRuleRequest
//
// @return UpdateWaitingRoomRuleResponse
func (client *Client) UpdateWaitingRoomRule(request *UpdateWaitingRoomRuleRequest) (_result *UpdateWaitingRoomRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWaitingRoomRuleResponse{}
	_body, _err := client.UpdateWaitingRoomRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads a client certificate authority (CA) certificate.
//
// @param request - UploadClientCaCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadClientCaCertificateResponse
func (client *Client) UploadClientCaCertificateWithOptions(request *UploadClientCaCertificateRequest, runtime *dara.RuntimeOptions) (_result *UploadClientCaCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Certificate) {
		body["Certificate"] = request.Certificate
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadClientCaCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadClientCaCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a client certificate authority (CA) certificate.
//
// @param request - UploadClientCaCertificateRequest
//
// @return UploadClientCaCertificateResponse
func (client *Client) UploadClientCaCertificate(request *UploadClientCaCertificateRequest) (_result *UploadClientCaCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadClientCaCertificateResponse{}
	_body, _err := client.UploadClientCaCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads the file that contains resources to be purged or prefetched.
//
// Description:
//
// >
//
//   - The file can be up to 10 MB in size.
//
// @param request - UploadFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadFileResponse
func (client *Client) UploadFileWithOptions(request *UploadFileRequest, runtime *dara.RuntimeOptions) (_result *UploadFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UploadTaskName) {
		query["UploadTaskName"] = request.UploadTaskName
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadFile"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads the file that contains resources to be purged or prefetched.
//
// Description:
//
// >
//
//   - The file can be up to 10 MB in size.
//
// @param request - UploadFileRequest
//
// @return UploadFileResponse
func (client *Client) UploadFile(request *UploadFileRequest) (_result *UploadFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadFileResponse{}
	_body, _err := client.UploadFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadFileAdvance(request *UploadFileAdvanceRequest, runtime *dara.RuntimeOptions) (_result *UploadFileResponse, _err error) {
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
		"Product":  dara.String("ESA"),
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
	uploadFileReq := &UploadFileRequest{}
	openapiutil.Convert(request, uploadFileReq)
	if !dara.IsNil(request.UrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.UrlObject,
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
		uploadFileReq.Url = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	uploadFileResp, _err := client.UploadFileWithOptions(uploadFileReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = uploadFileResp
	return _result, _err
}

// Summary:
//
// 上传源服务器CA证书
//
// @param request - UploadOriginCaCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadOriginCaCertificateResponse
func (client *Client) UploadOriginCaCertificateWithOptions(request *UploadOriginCaCertificateRequest, runtime *dara.RuntimeOptions) (_result *UploadOriginCaCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Certificate) {
		body["Certificate"] = request.Certificate
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadOriginCaCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadOriginCaCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传源服务器CA证书
//
// @param request - UploadOriginCaCertificateRequest
//
// @return UploadOriginCaCertificateResponse
func (client *Client) UploadOriginCaCertificate(request *UploadOriginCaCertificateRequest) (_result *UploadOriginCaCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadOriginCaCertificateResponse{}
	_body, _err := client.UploadOriginCaCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传域名回源客户端证书
//
// @param request - UploadOriginClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadOriginClientCertificateResponse
func (client *Client) UploadOriginClientCertificateWithOptions(request *UploadOriginClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *UploadOriginClientCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Certificate) {
		body["Certificate"] = request.Certificate
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PrivateKey) {
		body["PrivateKey"] = request.PrivateKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadOriginClientCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadOriginClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传域名回源客户端证书
//
// @param request - UploadOriginClientCertificateRequest
//
// @return UploadOriginClientCertificateResponse
func (client *Client) UploadOriginClientCertificate(request *UploadOriginClientCertificateRequest) (_result *UploadOriginClientCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadOriginClientCertificateResponse{}
	_body, _err := client.UploadOriginClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Upload site origin client certificate
//
// @param request - UploadSiteOriginClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadSiteOriginClientCertificateResponse
func (client *Client) UploadSiteOriginClientCertificateWithOptions(request *UploadSiteOriginClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *UploadSiteOriginClientCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Certificate) {
		body["Certificate"] = request.Certificate
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PrivateKey) {
		body["PrivateKey"] = request.PrivateKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadSiteOriginClientCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadSiteOriginClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Upload site origin client certificate
//
// @param request - UploadSiteOriginClientCertificateRequest
//
// @return UploadSiteOriginClientCertificateResponse
func (client *Client) UploadSiteOriginClientCertificate(request *UploadSiteOriginClientCertificateRequest) (_result *UploadSiteOriginClientCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadSiteOriginClientCertificateResponse{}
	_body, _err := client.UploadSiteOriginClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the ownership of a website domain. Websites that pass the verification are automatically activated.
//
// Description:
//
// 1.  For a website connected by using NS setup, this operation verifies whether the nameservers of the website are the nameservers assigned by Alibaba Cloud.
//
// 2.  For a website connected by using CNAME setup, this operation verifies whether the website has a TXT record whose hostname is  _esaauth.[websiteDomainName] and record value is the value of VerifyCode to the DNS records of your domain. You can see the VerifyCode field in the site information.
//
// @param request - VerifySiteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifySiteResponse
func (client *Client) VerifySiteWithOptions(request *VerifySiteRequest, runtime *dara.RuntimeOptions) (_result *VerifySiteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifySite"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifySiteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the ownership of a website domain. Websites that pass the verification are automatically activated.
//
// Description:
//
// 1.  For a website connected by using NS setup, this operation verifies whether the nameservers of the website are the nameservers assigned by Alibaba Cloud.
//
// 2.  For a website connected by using CNAME setup, this operation verifies whether the website has a TXT record whose hostname is  _esaauth.[websiteDomainName] and record value is the value of VerifyCode to the DNS records of your domain. You can see the VerifyCode field in the site information.
//
// @param request - VerifySiteRequest
//
// @return VerifySiteResponse
func (client *Client) VerifySite(request *VerifySiteRequest) (_result *VerifySiteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifySiteResponse{}
	_body, _err := client.VerifySiteWithOptions(request, runtime)
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
