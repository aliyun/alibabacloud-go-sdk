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
		"cn-hangzhou":                 dara.String("cas.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("cas.aliyuncs.com"),
		"ap-southeast-3":              dara.String("cas.aliyuncs.com"),
		"ap-southeast-5":              dara.String("cas.aliyuncs.com"),
		"cn-beijing":                  dara.String("cas.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("cas.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("cas.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("cas.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("cas.aliyuncs.com"),
		"cn-chengdu":                  dara.String("cas.aliyuncs.com"),
		"cn-edge-1":                   dara.String("cas.aliyuncs.com"),
		"cn-fujian":                   dara.String("cas.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("cas.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("cas.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("cas.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("cas.aliyuncs.com"),
		"cn-hongkong":                 dara.String("cas.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("cas.aliyuncs.com"),
		"cn-huhehaote":                dara.String("cas.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("cas.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("cas.aliyuncs.com"),
		"cn-qingdao":                  dara.String("cas.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("cas.aliyuncs.com"),
		"cn-shanghai":                 dara.String("cas.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("cas.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("cas.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("cas.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("cas.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("cas.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("cas.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("cas.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("cas.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("cas.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("cas.aliyuncs.com"),
		"cn-wuhan":                    dara.String("cas.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("cas.aliyuncs.com"),
		"cn-yushanfang":               dara.String("cas.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("cas.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("cas.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("cas.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("cas.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("cas.aliyuncs.com"),
		"eu-west-1":                   dara.String("cas.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("cas.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("cas.aliyuncs.com"),
		"us-east-1":                   dara.String("cas.aliyuncs.com"),
		"us-west-1":                   dara.String("cas.aliyuncs.com"),
		"me-east-1":                   dara.String("cas.me-east-1.aliyuncs.com"),
		"eu-central-1":                dara.String("cas.eu-central-1.aliyuncs.com"),
		"ap-southeast-2":              dara.String("cas.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-1":              dara.String("cas.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":                  dara.String("cas.ap-south-1.aliyuncs.com"),
		"ap-northeast-1":              dara.String("cas.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("cas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
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
// Adds an AccessKey for authorization.
//
// Description:
//
// The single-user QPS limit for this API is 100 queries per second (QPS). Calls that exceed this limit are throttled, which can affect your business operations. Call this API at a reasonable rate to avoid throttling.
//
// @param request - AddCloudAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCloudAccessResponse
func (client *Client) AddCloudAccessWithOptions(request *AddCloudAccessRequest, runtime *dara.RuntimeOptions) (_result *AddCloudAccessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloudName) {
		query["CloudName"] = request.CloudName
	}

	if !dara.IsNil(request.SecretId) {
		query["SecretId"] = request.SecretId
	}

	if !dara.IsNil(request.SecretKey) {
		query["SecretKey"] = request.SecretKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCloudAccess"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCloudAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an AccessKey for authorization.
//
// Description:
//
// The single-user QPS limit for this API is 100 queries per second (QPS). Calls that exceed this limit are throttled, which can affect your business operations. Call this API at a reasonable rate to avoid throttling.
//
// @param request - AddCloudAccessRequest
//
// @return AddCloudAccessResponse
func (client *Client) AddCloudAccess(request *AddCloudAccessRequest) (_result *AddCloudAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddCloudAccessResponse{}
	_body, _err := client.AddCloudAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a certificate application for a Certificate Management Service instance.
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyCertificate"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
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
// Submits a certificate application for a Certificate Management Service instance.
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
// # Updates the notification status in batches
//
// Description:
//
// After a CA certificate is created, it is in the normal issuance state by default. You can call this operation to change the status of a CA certificate from normal issuance to revoked. In the normal issuance state, the CA certificate can be used to issue certificates. In the revoked state, the CA certificate cannot be used to issue certificates, and the certificates that have been issued by the CA certificate also become invalid accordingly.
//
// Before you call this operation, you must have called [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) to create a root CA certificate and called [CreateSubCACertificate](https://help.aliyun.com/document_detail/465959.html) to create a sub CA certificate.
//
// ## QPS limit
//
// The QPS limit per user for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation properly.
//
// @param request - BatchUpdateNoticeStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateNoticeStatusResponse
func (client *Client) BatchUpdateNoticeStatusWithOptions(request *BatchUpdateNoticeStatusRequest, runtime *dara.RuntimeOptions) (_result *BatchUpdateNoticeStatusResponse, _err error) {
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

	if !dara.IsNil(request.NoticeBiz) {
		query["NoticeBiz"] = request.NoticeBiz
	}

	if !dara.IsNil(request.NoticeStatus) {
		query["NoticeStatus"] = request.NoticeStatus
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUpdateNoticeStatus"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUpdateNoticeStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Updates the notification status in batches
//
// Description:
//
// After a CA certificate is created, it is in the normal issuance state by default. You can call this operation to change the status of a CA certificate from normal issuance to revoked. In the normal issuance state, the CA certificate can be used to issue certificates. In the revoked state, the CA certificate cannot be used to issue certificates, and the certificates that have been issued by the CA certificate also become invalid accordingly.
//
// Before you call this operation, you must have called [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) to create a root CA certificate and called [CreateSubCACertificate](https://help.aliyun.com/document_detail/465959.html) to create a sub CA certificate.
//
// ## QPS limit
//
// The QPS limit per user for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation properly.
//
// @param request - BatchUpdateNoticeStatusRequest
//
// @return BatchUpdateNoticeStatusResponse
func (client *Client) BatchUpdateNoticeStatus(request *BatchUpdateNoticeStatusRequest) (_result *BatchUpdateNoticeStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchUpdateNoticeStatusResponse{}
	_body, _err := client.BatchUpdateNoticeStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes an issued certificate or cancels a pending certificate order and restores the quota.
//
// Description:
//
// This API has a limit of 10 queries per second (QPS) for each user. If you exceed this limit, API calls are throttled. This can affect your business. Call the API at a reasonable rate.
//
// @param request - CancelCertificateForPackageRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelCertificateForPackageRequestResponse
func (client *Client) CancelCertificateForPackageRequestWithOptions(request *CancelCertificateForPackageRequestRequest, runtime *dara.RuntimeOptions) (_result *CancelCertificateForPackageRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelCertificateForPackageRequest"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelCertificateForPackageRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes an issued certificate or cancels a pending certificate order and restores the quota.
//
// Description:
//
// This API has a limit of 10 queries per second (QPS) for each user. If you exceed this limit, API calls are throttled. This can affect your business. Call the API at a reasonable rate.
//
// @param request - CancelCertificateForPackageRequestRequest
//
// @return CancelCertificateForPackageRequestResponse
func (client *Client) CancelCertificateForPackageRequest(request *CancelCertificateForPackageRequestRequest) (_result *CancelCertificateForPackageRequestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelCertificateForPackageRequestResponse{}
	_body, _err := client.CancelCertificateForPackageRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a certificate application order that is pending domain verification or under review.
//
// Description:
//
// This API is limited to 100 queries per second (QPS) for each user. API calls that exceed this limit are throttled. Because this can impact your business, you should call this API at a reasonable rate.
//
// @param request - CancelOrderRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelOrderRequestResponse
func (client *Client) CancelOrderRequestWithOptions(request *CancelOrderRequestRequest, runtime *dara.RuntimeOptions) (_result *CancelOrderRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelOrderRequest"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelOrderRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a certificate application order that is pending domain verification or under review.
//
// Description:
//
// This API is limited to 100 queries per second (QPS) for each user. API calls that exceed this limit are throttled. Because this can impact your business, you should call this API at a reasonable rate.
//
// @param request - CancelOrderRequestRequest
//
// @return CancelOrderRequestResponse
func (client *Client) CancelOrderRequest(request *CancelOrderRequestRequest) (_result *CancelOrderRequestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelOrderRequestResponse{}
	_body, _err := client.CancelOrderRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a pending certificate application that has not been issued.
//
// @param request - CancelPendingCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelPendingCertificateResponse
func (client *Client) CancelPendingCertificateWithOptions(request *CancelPendingCertificateRequest, runtime *dara.RuntimeOptions) (_result *CancelPendingCertificateResponse, _err error) {
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
		Action:      dara.String("CancelPendingCertificate"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelPendingCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a pending certificate application that has not been issued.
//
// @param request - CancelPendingCertificateRequest
//
// @return CancelPendingCertificateResponse
func (client *Client) CancelPendingCertificate(request *CancelPendingCertificateRequest) (_result *CancelPendingCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelPendingCertificateResponse{}
	_body, _err := client.CancelPendingCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a certificate application by using a purchased certificate package quota.
//
// Description:
//
// - Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455800.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
// - After you call this operation to submit a certificate application and the certificate is issued, the certificate quota provided by the resource plan that you purchased is consumed. When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate that you want to apply for.
//
// - After you call this operation to submit a certificate application, you also need to call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required for domain name ownership verification and manually complete the verification. Then, your certificate application is reviewed by the certificate authority (CA). If you use the Domain Name System (DNS) verification method, you must complete the verification on your DNS service provider system. If you use the file verification method, you must complete the verification on the DNS server.
//
// @param request - CreateCertificateForPackageRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCertificateForPackageRequestResponse
func (client *Client) CreateCertificateForPackageRequestWithOptions(request *CreateCertificateForPackageRequestRequest, runtime *dara.RuntimeOptions) (_result *CreateCertificateForPackageRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyName) {
		query["CompanyName"] = request.CompanyName
	}

	if !dara.IsNil(request.Csr) {
		query["Csr"] = request.Csr
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	if !dara.IsNil(request.ValidateType) {
		query["ValidateType"] = request.ValidateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCertificateForPackageRequest"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCertificateForPackageRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a certificate application by using a purchased certificate package quota.
//
// Description:
//
// - Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455800.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
// - After you call this operation to submit a certificate application and the certificate is issued, the certificate quota provided by the resource plan that you purchased is consumed. When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate that you want to apply for.
//
// - After you call this operation to submit a certificate application, you also need to call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required for domain name ownership verification and manually complete the verification. Then, your certificate application is reviewed by the certificate authority (CA). If you use the Domain Name System (DNS) verification method, you must complete the verification on your DNS service provider system. If you use the file verification method, you must complete the verification on the DNS server.
//
// @param request - CreateCertificateForPackageRequestRequest
//
// @return CreateCertificateForPackageRequestResponse
func (client *Client) CreateCertificateForPackageRequest(request *CreateCertificateForPackageRequestRequest) (_result *CreateCertificateForPackageRequestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCertificateForPackageRequestResponse{}
	_body, _err := client.CreateCertificateForPackageRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Purchases, applies for, and issues a domain validated (DV) certificate by using extended certificate services.
//
// Description:
//
// - You can call this operation to apply for only DV certificates. If you want to apply for an organization validated (OV) or extended validation (EV) certificate, we recommend that you call the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/455296.html) operation. This operation allows you to apply for certificates of all specifications and specify the method to generate a certificate signing request (CSR) file.
//
// - Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455803.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
// - When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate.
//
// - After you call this operation to submit a certificate application, Certificate Management Service automatically creates a CSR file for your application and consumes the certificate quota in the certificate resource plans of the specified specifications that you purchased. After you call this operation, you also need to call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required to complete domain name verification, and manually complete the verification. If you use the DNS verification method, you must complete the verification on the management platform of the domain name. If you use the file verification method, you must complete the verification on your DNS server. Then, the certificate authority (CA) will review your certificate application.
//
// @param request - CreateCertificateRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCertificateRequestResponse
func (client *Client) CreateCertificateRequestWithOptions(request *CreateCertificateRequestRequest, runtime *dara.RuntimeOptions) (_result *CreateCertificateRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	if !dara.IsNil(request.ValidateType) {
		query["ValidateType"] = request.ValidateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCertificateRequest"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCertificateRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Purchases, applies for, and issues a domain validated (DV) certificate by using extended certificate services.
//
// Description:
//
// - You can call this operation to apply for only DV certificates. If you want to apply for an organization validated (OV) or extended validation (EV) certificate, we recommend that you call the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/455296.html) operation. This operation allows you to apply for certificates of all specifications and specify the method to generate a certificate signing request (CSR) file.
//
// - Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455803.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
// - When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate.
//
// - After you call this operation to submit a certificate application, Certificate Management Service automatically creates a CSR file for your application and consumes the certificate quota in the certificate resource plans of the specified specifications that you purchased. After you call this operation, you also need to call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required to complete domain name verification, and manually complete the verification. If you use the DNS verification method, you must complete the verification on the management platform of the domain name. If you use the file verification method, you must complete the verification on your DNS server. Then, the certificate authority (CA) will review your certificate application.
//
// @param request - CreateCertificateRequestRequest
//
// @return CreateCertificateRequestResponse
func (client *Client) CreateCertificateRequest(request *CreateCertificateRequestRequest) (_result *CreateCertificateRequestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCertificateRequestResponse{}
	_body, _err := client.CreateCertificateRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Purchases, applies for, and issues a domain validated (DV) certificate by using a custom certificate signing request (CSR) file.
//
// Description:
//
// - You can use this operation to apply for only a domain validated (DV) certificate. You cannot use this operation to apply for an organization validated (OV) certificate. We recommend that you use the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/455296.html) operation to apply for a certificate. You can use the CreateCertificateForPackageRequest operation to apply for certificates of all types and specify the CSR generation method.
//
// - Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455803.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
// - When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate that you want to apply for.
//
// - After you call this operation to submit a certificate application, the certificate quota of the required specifications that you purchased is consumed. After you call this operation, you must call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required for domain name ownership verification and manually complete the verification. Then, your certificate application is reviewed by the certificate authority (CA). If you use the Domain Name System (DNS) verification method, you must complete the verification on your DNS service provider system. If you use the file verification method, you must complete the verification on the DNS server.
//
// @param request - CreateCertificateWithCsrRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCertificateWithCsrRequestResponse
func (client *Client) CreateCertificateWithCsrRequestWithOptions(request *CreateCertificateWithCsrRequestRequest, runtime *dara.RuntimeOptions) (_result *CreateCertificateWithCsrRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Csr) {
		query["Csr"] = request.Csr
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	if !dara.IsNil(request.ValidateType) {
		query["ValidateType"] = request.ValidateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCertificateWithCsrRequest"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCertificateWithCsrRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Purchases, applies for, and issues a domain validated (DV) certificate by using a custom certificate signing request (CSR) file.
//
// Description:
//
// - You can use this operation to apply for only a domain validated (DV) certificate. You cannot use this operation to apply for an organization validated (OV) certificate. We recommend that you use the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/455296.html) operation to apply for a certificate. You can use the CreateCertificateForPackageRequest operation to apply for certificates of all types and specify the CSR generation method.
//
// - Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455803.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
// - When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate that you want to apply for.
//
// - After you call this operation to submit a certificate application, the certificate quota of the required specifications that you purchased is consumed. After you call this operation, you must call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required for domain name ownership verification and manually complete the verification. Then, your certificate application is reviewed by the certificate authority (CA). If you use the Domain Name System (DNS) verification method, you must complete the verification on your DNS service provider system. If you use the file verification method, you must complete the verification on the DNS server.
//
// @param request - CreateCertificateWithCsrRequestRequest
//
// @return CreateCertificateWithCsrRequestResponse
func (client *Client) CreateCertificateWithCsrRequest(request *CreateCertificateWithCsrRequestRequest) (_result *CreateCertificateWithCsrRequestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCertificateWithCsrRequestResponse{}
	_body, _err := client.CreateCertificateWithCsrRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a certificate signing request (CSR) that contains information about an SSL certificate to apply for, such as the domain names and the certificate holder. You must provide a CSR when you submit a certificate application to a certificate authority (CA).
//
// @param request - CreateCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCsrResponse
func (client *Client) CreateCsrWithOptions(request *CreateCsrRequest, runtime *dara.RuntimeOptions) (_result *CreateCsrResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CommonName) {
		query["CommonName"] = request.CommonName
	}

	if !dara.IsNil(request.CorpName) {
		query["CorpName"] = request.CorpName
	}

	if !dara.IsNil(request.CountryCode) {
		query["CountryCode"] = request.CountryCode
	}

	if !dara.IsNil(request.Department) {
		query["Department"] = request.Department
	}

	if !dara.IsNil(request.KeySize) {
		query["KeySize"] = request.KeySize
	}

	if !dara.IsNil(request.Locality) {
		query["Locality"] = request.Locality
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Province) {
		query["Province"] = request.Province
	}

	if !dara.IsNil(request.Sans) {
		query["Sans"] = request.Sans
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCsr"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCsrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a certificate signing request (CSR) that contains information about an SSL certificate to apply for, such as the domain names and the certificate holder. You must provide a CSR when you submit a certificate application to a certificate authority (CA).
//
// @param request - CreateCsrRequest
//
// @return CreateCsrResponse
func (client *Client) CreateCsr(request *CreateCsrRequest) (_result *CreateCsrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCsrResponse{}
	_body, _err := client.CreateCsrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a certificate deployment task to deploy an SSL certificate to one or more Alibaba Cloud services immediately or at a scheduled time.
//
// Description:
//
// After the task creation is completed, the task will be in the editing state. You need to call the UpdateDeploymentJobStatus interface to change the status to the pending state, otherwise the task will not be executed.
//
// @param request - CreateDeploymentJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeploymentJobResponse
func (client *Client) CreateDeploymentJobWithOptions(request *CreateDeploymentJobRequest, runtime *dara.RuntimeOptions) (_result *CreateDeploymentJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIds) {
		query["CertIds"] = request.CertIds
	}

	if !dara.IsNil(request.ContactIds) {
		query["ContactIds"] = request.ContactIds
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ScheduleTime) {
		query["ScheduleTime"] = request.ScheduleTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeploymentJob"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeploymentJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a certificate deployment task to deploy an SSL certificate to one or more Alibaba Cloud services immediately or at a scheduled time.
//
// Description:
//
// After the task creation is completed, the task will be in the editing state. You need to call the UpdateDeploymentJobStatus interface to change the status to the pending state, otherwise the task will not be executed.
//
// @param request - CreateDeploymentJobRequest
//
// @return CreateDeploymentJobResponse
func (client *Client) CreateDeploymentJob(request *CreateDeploymentJobRequest) (_result *CreateDeploymentJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDeploymentJobResponse{}
	_body, _err := client.CreateDeploymentJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Issues a single client certificate from the general user certificate repository.
//
// Description:
//
// This API is limited to 10 QPS per user. Exceeding this limit triggers throttling, which can affect your business. Call this API at a reasonable rate to avoid disruption.
//
// @param request - CreateWHClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWHClientCertificateResponse
func (client *Client) CreateWHClientCertificateWithOptions(request *CreateWHClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *CreateWHClientCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AfterTime) {
		query["AfterTime"] = request.AfterTime
	}

	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.BeforeTime) {
		query["BeforeTime"] = request.BeforeTime
	}

	if !dara.IsNil(request.CommonName) {
		query["CommonName"] = request.CommonName
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.Csr) {
		query["Csr"] = request.Csr
	}

	if !dara.IsNil(request.Days) {
		query["Days"] = request.Days
	}

	if !dara.IsNil(request.Immediately) {
		query["Immediately"] = request.Immediately
	}

	if !dara.IsNil(request.Locality) {
		query["Locality"] = request.Locality
	}

	if !dara.IsNil(request.Months) {
		query["Months"] = request.Months
	}

	if !dara.IsNil(request.Organization) {
		query["Organization"] = request.Organization
	}

	if !dara.IsNil(request.OrganizationUnit) {
		query["OrganizationUnit"] = request.OrganizationUnit
	}

	if !dara.IsNil(request.ParentIdentifier) {
		query["ParentIdentifier"] = request.ParentIdentifier
	}

	if !dara.IsNil(request.SanType) {
		query["SanType"] = request.SanType
	}

	if !dara.IsNil(request.SanValue) {
		query["SanValue"] = request.SanValue
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Years) {
		query["Years"] = request.Years
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWHClientCertificate"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWHClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Issues a single client certificate from the general user certificate repository.
//
// Description:
//
// This API is limited to 10 QPS per user. Exceeding this limit triggers throttling, which can affect your business. Call this API at a reasonable rate to avoid disruption.
//
// @param request - CreateWHClientCertificateRequest
//
// @return CreateWHClientCertificateResponse
func (client *Client) CreateWHClientCertificate(request *CreateWHClientCertificateRequest) (_result *CreateWHClientCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWHClientCertificateResponse{}
	_body, _err := client.CreateWHClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a certificate warehouse.
//
// @param request - CreateWarehouseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWarehouseResponse
func (client *Client) CreateWarehouseWithOptions(request *CreateWarehouseRequest, runtime *dara.RuntimeOptions) (_result *CreateWarehouseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Biz) {
		query["Biz"] = request.Biz
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWarehouse"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWarehouseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a certificate warehouse.
//
// @param request - CreateWarehouseRequest
//
// @return CreateWarehouseResponse
func (client *Client) CreateWarehouse(request *CreateWarehouseRequest) (_result *CreateWarehouseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWarehouseResponse{}
	_body, _err := client.CreateWarehouseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Decrypts data that was encrypted by using a certificate in a certificate application repository.
//
// Description:
//
// The queries per second (QPS) limit for this API operation is 10 per user. If you exceed the limit, API calls are throttled, which may affect your business. Call this operation at a reasonable rate.
//
// @param request - DecryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DecryptResponse
func (client *Client) DecryptWithOptions(request *DecryptRequest, runtime *dara.RuntimeOptions) (_result *DecryptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.CiphertextBlob) {
		query["CiphertextBlob"] = request.CiphertextBlob
	}

	if !dara.IsNil(request.CustomIdentifier) {
		query["CustomIdentifier"] = request.CustomIdentifier
	}

	if !dara.IsNil(request.MessageType) {
		query["MessageType"] = request.MessageType
	}

	if !dara.IsNil(request.WarehouseId) {
		query["WarehouseId"] = request.WarehouseId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Decrypt"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DecryptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Decrypts data that was encrypted by using a certificate in a certificate application repository.
//
// Description:
//
// The queries per second (QPS) limit for this API operation is 10 per user. If you exceed the limit, API calls are throttled, which may affect your business. Call this operation at a reasonable rate.
//
// @param request - DecryptRequest
//
// @return DecryptResponse
func (client *Client) Decrypt(request *DecryptRequest) (_result *DecryptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DecryptResponse{}
	_body, _err := client.DecryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a failed domain validated (DV) certificate application order.
//
// Description:
//
// You can call this operation to delete a certificate application order only in the following scenarios:
//
// - The status of the order is **review failed**. You have called the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to query the status of the certificate application order and the value of the **Type*	- parameter is **verify_fail**.
//
// - The status of the order is **pending application**. You have called the [CancelOrderRequest](https://help.aliyun.com/document_detail/455299.html) operation to cancel a certificate application order whose status is pending review or being reviewed. The status of the certificate application order that is canceled in this case changes to **pending application**.
//
// @param request - DeleteCertificateRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCertificateRequestResponse
func (client *Client) DeleteCertificateRequestWithOptions(request *DeleteCertificateRequestRequest, runtime *dara.RuntimeOptions) (_result *DeleteCertificateRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCertificateRequest"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCertificateRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a failed domain validated (DV) certificate application order.
//
// Description:
//
// You can call this operation to delete a certificate application order only in the following scenarios:
//
// - The status of the order is **review failed**. You have called the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to query the status of the certificate application order and the value of the **Type*	- parameter is **verify_fail**.
//
// - The status of the order is **pending application**. You have called the [CancelOrderRequest](https://help.aliyun.com/document_detail/455299.html) operation to cancel a certificate application order whose status is pending review or being reviewed. The status of the certificate application order that is canceled in this case changes to **pending application**.
//
// @param request - DeleteCertificateRequestRequest
//
// @return DeleteCertificateRequestResponse
func (client *Client) DeleteCertificateRequest(request *DeleteCertificateRequestRequest) (_result *DeleteCertificateRequestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCertificateRequestResponse{}
	_body, _err := client.DeleteCertificateRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access key.
//
// Description:
//
// This operation is limited to 100 queries per second (QPS) per user. API calls that exceed this limit are throttled, which can impact your business.
//
// @param request - DeleteCloudAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudAccessResponse
func (client *Client) DeleteCloudAccessWithOptions(request *DeleteCloudAccessRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudAccessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessId) {
		query["AccessId"] = request.AccessId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudAccess"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access key.
//
// Description:
//
// This operation is limited to 100 queries per second (QPS) per user. API calls that exceed this limit are throttled, which can impact your business.
//
// @param request - DeleteCloudAccessRequest
//
// @return DeleteCloudAccessResponse
func (client *Client) DeleteCloudAccess(request *DeleteCloudAccessRequest) (_result *DeleteCloudAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloudAccessResponse{}
	_body, _err := client.DeleteCloudAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a certificate signing request (CSR).
//
// @param request - DeleteCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCsrResponse
func (client *Client) DeleteCsrWithOptions(request *DeleteCsrRequest, runtime *dara.RuntimeOptions) (_result *DeleteCsrResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CsrId) {
		query["CsrId"] = request.CsrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCsr"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCsrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a certificate signing request (CSR).
//
// @param request - DeleteCsrRequest
//
// @return DeleteCsrResponse
func (client *Client) DeleteCsr(request *DeleteCsrRequest) (_result *DeleteCsrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCsrResponse{}
	_body, _err := client.DeleteCsrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a certificate deployment task.
//
// @param request - DeleteDeploymentJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeploymentJobResponse
func (client *Client) DeleteDeploymentJobWithOptions(request *DeleteDeploymentJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteDeploymentJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDeploymentJob"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDeploymentJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a certificate deployment task.
//
// @param request - DeleteDeploymentJobRequest
//
// @return DeleteDeploymentJobResponse
func (client *Client) DeleteDeploymentJob(request *DeleteDeploymentJobRequest) (_result *DeleteDeploymentJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDeploymentJobResponse{}
	_body, _err := client.DeleteDeploymentJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Certificate Management Service instance.
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
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
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Certificate Management Service instance.
//
// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a private certificate from a certificate application repository.
//
// Description:
//
// You can call the DeletePCACert operation to delete a private certificate from a certificate application repository.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeletePCACertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePCACertResponse
func (client *Client) DeletePCACertWithOptions(request *DeletePCACertRequest, runtime *dara.RuntimeOptions) (_result *DeletePCACertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Identifier) {
		query["Identifier"] = request.Identifier
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePCACert"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePCACertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a private certificate from a certificate application repository.
//
// Description:
//
// You can call the DeletePCACert operation to delete a private certificate from a certificate application repository.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeletePCACertRequest
//
// @return DeletePCACertResponse
func (client *Client) DeletePCACert(request *DeletePCACertRequest) (_result *DeletePCACertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePCACertResponse{}
	_body, _err := client.DeletePCACertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an expired, revoked, or manually uploaded certificate from Certificate Management Service.
//
// Description:
//
// This operation is limited to 100 queries per second (QPS) per user. API calls exceeding this limit are throttled, which can impact your business. We recommend calling this operation at a reasonable rate to avoid this.
//
// @param request - DeleteUserCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserCertificateResponse
func (client *Client) DeleteUserCertificateWithOptions(request *DeleteUserCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertId) {
		query["CertId"] = request.CertId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserCertificate"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an expired, revoked, or manually uploaded certificate from Certificate Management Service.
//
// Description:
//
// This operation is limited to 100 queries per second (QPS) per user. API calls exceeding this limit are throttled, which can impact your business. We recommend calling this operation at a reasonable rate to avoid this.
//
// @param request - DeleteUserCertificateRequest
//
// @return DeleteUserCertificateResponse
func (client *Client) DeleteUserCertificate(request *DeleteUserCertificateRequest) (_result *DeleteUserCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserCertificateResponse{}
	_body, _err := client.DeleteUserCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a certificate warehouse.
//
// Description:
//
// This operation deletes a certificate warehouse.
//
// ### QPS limit
//
// This operation has a QPS limit of 10 requests per second per user. Exceeding this limit causes subsequent API calls to be throttled, which can impact your services. To ensure service availability, call this operation at a reasonable rate.
//
// @param request - DeleteWarehouseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWarehouseResponse
func (client *Client) DeleteWarehouseWithOptions(request *DeleteWarehouseRequest, runtime *dara.RuntimeOptions) (_result *DeleteWarehouseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WarehouseInstanceId) {
		query["WarehouseInstanceId"] = request.WarehouseInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWarehouse"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWarehouseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a certificate warehouse.
//
// Description:
//
// This operation deletes a certificate warehouse.
//
// ### QPS limit
//
// This operation has a QPS limit of 10 requests per second per user. Exceeding this limit causes subsequent API calls to be throttled, which can impact your services. To ensure service availability, call this operation at a reasonable rate.
//
// @param request - DeleteWarehouseRequest
//
// @return DeleteWarehouseResponse
func (client *Client) DeleteWarehouse(request *DeleteWarehouseRequest) (_result *DeleteWarehouseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWarehouseResponse{}
	_body, _err := client.DeleteWarehouseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a worker task from a certificate deployment task.
//
// @param request - DeleteWorkerResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkerResourceResponse
func (client *Client) DeleteWorkerResourceWithOptions(request *DeleteWorkerResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteWorkerResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.WorkerId) {
		query["WorkerId"] = request.WorkerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkerResource"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkerResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a worker task from a certificate deployment task.
//
// @param request - DeleteWorkerResourceRequest
//
// @return DeleteWorkerResourceResponse
func (client *Client) DeleteWorkerResource(request *DeleteWorkerResourceRequest) (_result *DeleteWorkerResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWorkerResourceResponse{}
	_body, _err := client.DeleteWorkerResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a certificate application order, such as domain validation progress.
//
// Description:
//
// If you have not completed domain ownership validation after submitting a certificate request, you can call this operation to obtain the information required to complete domain validation. Using the returned domain validation information, you can complete domain validation on the DNS management platform (DNS validation method) or on the domain server (file validation method).
//
// Your certificate request will enter the CA center review stage only after you complete domain validation. After the CA center approves your certificate request, a certificate will be issued to you. If the certificate has been issued, you can call this operation to obtain the issued certificate file and private key content.
//
// <props="china">
//
// For the complete process of requesting a certificate using the resource plan API, see [Process of requesting a certificate using API operations](https://help.aliyun.com/document_detail/204741.html).
//
// @param request - DescribeCertificateStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCertificateStateResponse
func (client *Client) DescribeCertificateStateWithOptions(request *DescribeCertificateStateRequest, runtime *dara.RuntimeOptions) (_result *DescribeCertificateStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCertificateState"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCertificateStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a certificate application order, such as domain validation progress.
//
// Description:
//
// If you have not completed domain ownership validation after submitting a certificate request, you can call this operation to obtain the information required to complete domain validation. Using the returned domain validation information, you can complete domain validation on the DNS management platform (DNS validation method) or on the domain server (file validation method).
//
// Your certificate request will enter the CA center review stage only after you complete domain validation. After the CA center approves your certificate request, a certificate will be issued to you. If the certificate has been issued, you can call this operation to obtain the issued certificate file and private key content.
//
// <props="china">
//
// For the complete process of requesting a certificate using the resource plan API, see [Process of requesting a certificate using API operations](https://help.aliyun.com/document_detail/204741.html).
//
// @param request - DescribeCertificateStateRequest
//
// @return DescribeCertificateStateResponse
func (client *Client) DescribeCertificateState(request *DescribeCertificateStateRequest) (_result *DescribeCertificateStateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCertificateStateResponse{}
	_body, _err := client.DescribeCertificateStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of cloud resources on which certificates were deployed by using a multi-cloud deployment task.
//
// @param request - DescribeCloudResourceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudResourceStatusResponse
func (client *Client) DescribeCloudResourceStatusWithOptions(request *DescribeCloudResourceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudResourceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecretId) {
		query["SecretId"] = request.SecretId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudResourceStatus"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudResourceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of cloud resources on which certificates were deployed by using a multi-cloud deployment task.
//
// @param request - DescribeCloudResourceStatusRequest
//
// @return DescribeCloudResourceStatusResponse
func (client *Client) DescribeCloudResourceStatus(request *DescribeCloudResourceStatusRequest) (_result *DescribeCloudResourceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudResourceStatusResponse{}
	_body, _err := client.DescribeCloudResourceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves information about a certificate deployment task, including the task status, target resources, and certificates.
//
// @param request - DescribeDeploymentJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeploymentJobResponse
func (client *Client) DescribeDeploymentJobWithOptions(request *DescribeDeploymentJobRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeploymentJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeploymentJob"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeploymentJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about a certificate deployment task, including the task status, target resources, and certificates.
//
// @param request - DescribeDeploymentJobRequest
//
// @return DescribeDeploymentJobResponse
func (client *Client) DescribeDeploymentJob(request *DescribeDeploymentJobRequest) (_result *DescribeDeploymentJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDeploymentJobResponse{}
	_body, _err := client.DescribeDeploymentJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution status summary of a certificate deployment task, including the number of succeeded and failed workers.
//
// @param request - DescribeDeploymentJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeploymentJobStatusResponse
func (client *Client) DescribeDeploymentJobStatusWithOptions(request *DescribeDeploymentJobStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeploymentJobStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeploymentJobStatus"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeploymentJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution status summary of a certificate deployment task, including the number of succeeded and failed workers.
//
// @param request - DescribeDeploymentJobStatusRequest
//
// @return DescribeDeploymentJobStatusResponse
func (client *Client) DescribeDeploymentJobStatus(request *DescribeDeploymentJobStatusRequest) (_result *DescribeDeploymentJobStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDeploymentJobStatusResponse{}
	_body, _err := client.DescribeDeploymentJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the quota and usage of domain validated (DV) certificate packages.
//
// @param request - DescribePackageStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePackageStateResponse
func (client *Client) DescribePackageStateWithOptions(request *DescribePackageStateRequest, runtime *dara.RuntimeOptions) (_result *DescribePackageStateResponse, _err error) {
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
		Action:      dara.String("DescribePackageState"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePackageStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quota and usage of domain validated (DV) certificate packages.
//
// @param request - DescribePackageStateRequest
//
// @return DescribePackageStateResponse
func (client *Client) DescribePackageState(request *DescribePackageStateRequest) (_result *DescribePackageStateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePackageStateResponse{}
	_body, _err := client.DescribePackageStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a certificate stored in a certificate warehouse.
//
// @param request - DescribeWarehouseCertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWarehouseCertResponse
func (client *Client) DescribeWarehouseCertWithOptions(request *DescribeWarehouseCertRequest, runtime *dara.RuntimeOptions) (_result *DescribeWarehouseCertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWarehouseCert"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWarehouseCertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a certificate stored in a certificate warehouse.
//
// @param request - DescribeWarehouseCertRequest
//
// @return DescribeWarehouseCertResponse
func (client *Client) DescribeWarehouseCert(request *DescribeWarehouseCertRequest) (_result *DescribeWarehouseCertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWarehouseCertResponse{}
	_body, _err := client.DescribeWarehouseCertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Encrypts data by using a certificate in a certificate application repository.
//
// Description:
//
// The queries per second (QPS) limit for a single user is 10. If you exceed this limit, API calls are throttled, which may affect your business. To prevent this, call this operation at a reasonable rate.
//
// @param request - EncryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EncryptResponse
func (client *Client) EncryptWithOptions(request *EncryptRequest, runtime *dara.RuntimeOptions) (_result *EncryptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.CustomIdentifier) {
		query["CustomIdentifier"] = request.CustomIdentifier
	}

	if !dara.IsNil(request.MessageType) {
		query["MessageType"] = request.MessageType
	}

	if !dara.IsNil(request.Plaintext) {
		query["Plaintext"] = request.Plaintext
	}

	if !dara.IsNil(request.WarehouseId) {
		query["WarehouseId"] = request.WarehouseId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Encrypt"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EncryptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Encrypts data by using a certificate in a certificate application repository.
//
// Description:
//
// The queries per second (QPS) limit for a single user is 10. If you exceed this limit, API calls are throttled, which may affect your business. To prevent this, call this operation at a reasonable rate.
//
// @param request - EncryptRequest
//
// @return EncryptResponse
func (client *Client) Encrypt(request *EncryptRequest) (_result *EncryptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EncryptResponse{}
	_body, _err := client.EncryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the total number of certificate-related assets, such as websites and cloud resources.
//
// Description:
//
// This API call queries the number of CA certificates that you have created, including root CA certificates and sub-CA certificates.
//
// ## QPS Limit
//
// This API call has a single-user limit of 10 queries per second (QPS). If you exceed this limit, API calls are rate-limited. This may affect your business. We recommend that you call this API operation at a reasonable rate.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAssetCountResponse
func (client *Client) GetAssetCountWithOptions(runtime *dara.RuntimeOptions) (_result *GetAssetCountResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetAssetCount"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAssetCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total number of certificate-related assets, such as websites and cloud resources.
//
// Description:
//
// This API call queries the number of CA certificates that you have created, including root CA certificates and sub-CA certificates.
//
// ## QPS Limit
//
// This API call has a single-user limit of 10 queries per second (QPS). If you exceed this limit, API calls are rate-limited. This may affect your business. We recommend that you call this API operation at a reasonable rate.
//
// @return GetAssetCountResponse
func (client *Client) GetAssetCount() (_result *GetAssetCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAssetCountResponse{}
	_body, _err := client.GetAssetCountWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the remaining quota for certificate application repository operations.
//
// Description:
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. If you exceed the limit, your API calls are throttled. This may impact your business. Call this operation at a reasonable rate.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCertWarehouseQuotaResponse
func (client *Client) GetCertWarehouseQuotaWithOptions(runtime *dara.RuntimeOptions) (_result *GetCertWarehouseQuotaResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetCertWarehouseQuota"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCertWarehouseQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the remaining quota for certificate application repository operations.
//
// Description:
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. If you exceed the limit, your API calls are throttled. This may impact your business. Call this operation at a reasonable rate.
//
// @return GetCertWarehouseQuotaResponse
func (client *Client) GetCertWarehouseQuota() (_result *GetCertWarehouseQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCertWarehouseQuotaResponse{}
	_body, _err := client.GetCertWarehouseQuotaWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves certificate details, excluding the certificate body and private key.
//
// @param request - GetCertificateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCertificateDetailResponse
func (client *Client) GetCertificateDetailWithOptions(request *GetCertificateDetailRequest, runtime *dara.RuntimeOptions) (_result *GetCertificateDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertificateId) {
		query["CertificateId"] = request.CertificateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCertificateDetail"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCertificateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves certificate details, excluding the certificate body and private key.
//
// @param request - GetCertificateDetailRequest
//
// @return GetCertificateDetailResponse
func (client *Client) GetCertificateDetail(request *GetCertificateDetailRequest) (_result *GetCertificateDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCertificateDetailResponse{}
	_body, _err := client.GetCertificateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the content of a certificate signing request (CSR).
//
// @param request - GetCsrDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCsrDetailResponse
func (client *Client) GetCsrDetailWithOptions(request *GetCsrDetailRequest, runtime *dara.RuntimeOptions) (_result *GetCsrDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CsrId) {
		query["CsrId"] = request.CsrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCsrDetail"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCsrDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the content of a certificate signing request (CSR).
//
// @param request - GetCsrDetailRequest
//
// @return GetCsrDetailResponse
func (client *Client) GetCsrDetail(request *GetCsrDetailRequest) (_result *GetCsrDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCsrDetailResponse{}
	_body, _err := client.GetCsrDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an instance.
//
// @param request - GetInstanceDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceDetailResponse
func (client *Client) GetInstanceDetailWithOptions(request *GetInstanceDetailRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceDetailResponse, _err error) {
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
		Action:      dara.String("GetInstanceDetail"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an instance.
//
// @param request - GetInstanceDetailRequest
//
// @return GetInstanceDetailResponse
func (client *Client) GetInstanceDetail(request *GetInstanceDetailRequest) (_result *GetInstanceDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceDetailResponse{}
	_body, _err := client.GetInstanceDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the summary statistics of Certificate Management Service instances, such as certificate counts by status.
//
// @param request - GetInstanceSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceSummaryResponse
func (client *Client) GetInstanceSummaryWithOptions(request *GetInstanceSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceSummary"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the summary statistics of Certificate Management Service instances, such as certificate counts by status.
//
// @param request - GetInstanceSummaryRequest
//
// @return GetInstanceSummaryResponse
func (client *Client) GetInstanceSummary(request *GetInstanceSummaryRequest) (_result *GetInstanceSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceSummaryResponse{}
	_body, _err := client.GetInstanceSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the resources that match a certificate.
//
// Description:
//
// 本接口用于通过私有 CA 实例的 ID，查询您通过 SSL 证书服务控制台购买的私有 CA 实例的状态信息，例如，CA 实例的状态、包含的证书数量、已签发的证书数量等。
//
// 调用本接口前，您必须已经通过[数字证书管理服务控制台](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist)购买了私有 CA。具体操作，请参见[购买私有 CA](https://help.aliyun.com/document_detail/208553.html)。
//
// ## QPS 限制
//
// 本接口的单用户 QPS 限制为 10 次/秒。超过限制，API 调用将会被限流，这可能影响您的业务，请合理调用。
//
// @param request - GetMatchedResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMatchedResourcesResponse
func (client *Client) GetMatchedResourcesWithOptions(request *GetMatchedResourcesRequest, runtime *dara.RuntimeOptions) (_result *GetMatchedResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIds) {
		query["CertIds"] = request.CertIds
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceScope) {
		query["ResourceScope"] = request.ResourceScope
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMatchedResources"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMatchedResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the resources that match a certificate.
//
// Description:
//
// 本接口用于通过私有 CA 实例的 ID，查询您通过 SSL 证书服务控制台购买的私有 CA 实例的状态信息，例如，CA 实例的状态、包含的证书数量、已签发的证书数量等。
//
// 调用本接口前，您必须已经通过[数字证书管理服务控制台](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist)购买了私有 CA。具体操作，请参见[购买私有 CA](https://help.aliyun.com/document_detail/208553.html)。
//
// ## QPS 限制
//
// 本接口的单用户 QPS 限制为 10 次/秒。超过限制，API 调用将会被限流，这可能影响您的业务，请合理调用。
//
// @param request - GetMatchedResourcesRequest
//
// @return GetMatchedResourcesResponse
func (client *Client) GetMatchedResources(request *GetMatchedResourcesRequest) (_result *GetMatchedResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMatchedResourcesResponse{}
	_body, _err := client.GetMatchedResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of assets with certificate-related risks, such as expired or soon-to-expire certificates.
//
// Description:
//
// This operation queries the number of created Certificate Authority (CA) certificates, including root and subordinate CA certificates.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed this limit, API calls are rate-limited, which may affect your business. We recommend that you call this operation at a reasonable frequency.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRiskCountResponse
func (client *Client) GetRiskCountWithOptions(runtime *dara.RuntimeOptions) (_result *GetRiskCountResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetRiskCount"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRiskCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of assets with certificate-related risks, such as expired or soon-to-expire certificates.
//
// Description:
//
// This operation queries the number of created Certificate Authority (CA) certificates, including root and subordinate CA certificates.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed this limit, API calls are rate-limited, which may affect your business. We recommend that you call this operation at a reasonable frequency.
//
// @return GetRiskCountResponse
func (client *Client) GetRiskCount() (_result *GetRiskCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRiskCountResponse{}
	_body, _err := client.GetRiskCountWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the processing result and status of a submitted certificate application.
//
// @param request - GetTaskAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskAttributeResponse
func (client *Client) GetTaskAttributeWithOptions(request *GetTaskAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetTaskAttributeResponse, _err error) {
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

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskAttribute"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the processing result and status of a submitted certificate application.
//
// @param request - GetTaskAttributeRequest
//
// @return GetTaskAttributeResponse
func (client *Client) GetTaskAttribute(request *GetTaskAttributeRequest) (_result *GetTaskAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTaskAttributeResponse{}
	_body, _err := client.GetTaskAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves certificate details, including the basic information, certificate body, and private key. You can also use this operation to download the certificate content and private key.
//
// Description:
//
// The queries per second (QPS) limit for each user is 100. If you exceed this limit, the system throttles your API calls, which may affect your business. We recommend that you call this operation within this limit.
//
// @param request - GetUserCertificateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserCertificateDetailResponse
func (client *Client) GetUserCertificateDetailWithOptions(request *GetUserCertificateDetailRequest, runtime *dara.RuntimeOptions) (_result *GetUserCertificateDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertFilter) {
		query["CertFilter"] = request.CertFilter
	}

	if !dara.IsNil(request.CertId) {
		query["CertId"] = request.CertId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserCertificateDetail"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserCertificateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves certificate details, including the basic information, certificate body, and private key. You can also use this operation to download the certificate content and private key.
//
// Description:
//
// The queries per second (QPS) limit for each user is 100. If you exceed this limit, the system throttles your API calls, which may affect your business. We recommend that you call this operation within this limit.
//
// @param request - GetUserCertificateDetailRequest
//
// @return GetUserCertificateDetailResponse
func (client *Client) GetUserCertificateDetail(request *GetUserCertificateDetailRequest) (_result *GetUserCertificateDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserCertificateDetailResponse{}
	_body, _err := client.GetUserCertificateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the certificate deployment statistics by cloud service type.
//
// Description:
//
// Queries the number of created Certificate Authority (CA) certificates, including root and subordinate CA certificates.
//
// ## QPS limit
//
// Each user can make up to 10 queries per second (QPS). If you exceed this limit, the system applies rate limiting to your API calls. This may affect your business. Make API calls at a reasonable rate.
//
// @param request - ListAssetCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAssetCountResponse
func (client *Client) ListAssetCountWithOptions(request *ListAssetCountRequest, runtime *dara.RuntimeOptions) (_result *ListAssetCountResponse, _err error) {
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

	if !dara.IsNil(request.ShowSize) {
		query["ShowSize"] = request.ShowSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAssetCount"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAssetCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificate deployment statistics by cloud service type.
//
// Description:
//
// Queries the number of created Certificate Authority (CA) certificates, including root and subordinate CA certificates.
//
// ## QPS limit
//
// Each user can make up to 10 queries per second (QPS). If you exceed this limit, the system applies rate limiting to your API calls. This may affect your business. Make API calls at a reasonable rate.
//
// @param request - ListAssetCountRequest
//
// @return ListAssetCountResponse
func (client *Client) ListAssetCount(request *ListAssetCountRequest) (_result *ListAssetCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAssetCountResponse{}
	_body, _err := client.ListAssetCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This API queries certificates in the certificate store.
//
// Description:
//
// The single-user QPS limit for this API is 10. Calls exceeding this limit are throttled, which may impact your business. Plan your API calls accordingly.
//
// @param request - ListCertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCertResponse
func (client *Client) ListCertWithOptions(request *ListCertRequest, runtime *dara.RuntimeOptions) (_result *ListCertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertType) {
		query["CertType"] = request.CertType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Identifiers) {
		query["Identifiers"] = request.Identifiers
	}

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.ShowSize) {
		query["ShowSize"] = request.ShowSize
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WarehouseId) {
		query["WarehouseId"] = request.WarehouseId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCert"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API queries certificates in the certificate store.
//
// Description:
//
// The single-user QPS limit for this API is 10. Calls exceeding this limit are throttled, which may impact your business. Plan your API calls accordingly.
//
// @param request - ListCertRequest
//
// @return ListCertResponse
func (client *Client) ListCert(request *ListCertRequest) (_result *ListCertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCertResponse{}
	_body, _err := client.ListCertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the certificate application repositories in your account.
//
// Description:
//
// You can call the ListCertWarehouse operation to query certificate repositories.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListCertWarehouseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCertWarehouseResponse
func (client *Client) ListCertWarehouseWithOptions(request *ListCertWarehouseRequest, runtime *dara.RuntimeOptions) (_result *ListCertWarehouseResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ShowSize) {
		query["ShowSize"] = request.ShowSize
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCertWarehouse"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCertWarehouseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificate application repositories in your account.
//
// Description:
//
// You can call the ListCertWarehouse operation to query certificate repositories.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListCertWarehouseRequest
//
// @return ListCertWarehouseResponse
func (client *Client) ListCertWarehouse(request *ListCertWarehouseRequest) (_result *ListCertWarehouseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCertWarehouseResponse{}
	_body, _err := client.ListCertWarehouseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the certificates managed by Certificate Management Service.
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertificateSource) {
		query["CertificateSource"] = request.CertificateSource
	}

	if !dara.IsNil(request.CertificateStatus) {
		query["CertificateStatus"] = request.CertificateStatus
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ShowSize) {
		query["ShowSize"] = request.ShowSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCertificates"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
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
// Queries the certificates managed by Certificate Management Service.
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
// Queries the AccessKey pairs that are configured for multi-cloud certificate deployment.
//
// @param request - ListCloudAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudAccessResponse
func (client *Client) ListCloudAccessWithOptions(request *ListCloudAccessRequest, runtime *dara.RuntimeOptions) (_result *ListCloudAccessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloudName) {
		query["CloudName"] = request.CloudName
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.SecretId) {
		query["SecretId"] = request.SecretId
	}

	if !dara.IsNil(request.ShowSize) {
		query["ShowSize"] = request.ShowSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudAccess"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the AccessKey pairs that are configured for multi-cloud certificate deployment.
//
// @param request - ListCloudAccessRequest
//
// @return ListCloudAccessResponse
func (client *Client) ListCloudAccess(request *ListCloudAccessRequest) (_result *ListCloudAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudAccessResponse{}
	_body, _err := client.ListCloudAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the cloud resources on which certificates are deployed, such as Server Load Balancer (SLB) instances and CDN domains.
//
// @param tmpReq - ListCloudResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudResourcesResponse
func (client *Client) ListCloudResourcesWithOptions(tmpReq *ListCloudResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListCloudResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListCloudResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CertIds) {
		request.CertIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CertIds, dara.String("CertIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIdsShrink) {
		query["CertIds"] = request.CertIdsShrink
	}

	if !dara.IsNil(request.CloudName) {
		query["CloudName"] = request.CloudName
	}

	if !dara.IsNil(request.CloudProduct) {
		query["CloudProduct"] = request.CloudProduct
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.SecretId) {
		query["SecretId"] = request.SecretId
	}

	if !dara.IsNil(request.ShowSize) {
		query["ShowSize"] = request.ShowSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudResources"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cloud resources on which certificates are deployed, such as Server Load Balancer (SLB) instances and CDN domains.
//
// @param request - ListCloudResourcesRequest
//
// @return ListCloudResourcesResponse
func (client *Client) ListCloudResources(request *ListCloudResourcesRequest) (_result *ListCloudResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudResourcesResponse{}
	_body, _err := client.ListCloudResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the contacts that receive certificate deployment notifications.
//
// @param request - ListContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContactResponse
func (client *Client) ListContactWithOptions(request *ListContactRequest, runtime *dara.RuntimeOptions) (_result *ListContactResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.ShowSize) {
		query["ShowSize"] = request.ShowSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContact"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the contacts that receive certificate deployment notifications.
//
// @param request - ListContactRequest
//
// @return ListContactResponse
func (client *Client) ListContact(request *ListContactRequest) (_result *ListContactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListContactResponse{}
	_body, _err := client.ListContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the certificate signing requests (CSRs) in your account.
//
// @param request - ListCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCsrResponse
func (client *Client) ListCsrWithOptions(request *ListCsrRequest, runtime *dara.RuntimeOptions) (_result *ListCsrResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.ShowSize) {
		query["ShowSize"] = request.ShowSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCsr"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCsrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificate signing requests (CSRs) in your account.
//
// @param request - ListCsrRequest
//
// @return ListCsrResponse
func (client *Client) ListCsr(request *ListCsrRequest) (_result *ListCsrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCsrResponse{}
	_body, _err := client.ListCsrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the certificate deployment tasks that are created in your account.
//
// @param request - ListDeploymentJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentJobResponse
func (client *Client) ListDeploymentJobWithOptions(request *ListDeploymentJobRequest, runtime *dara.RuntimeOptions) (_result *ListDeploymentJobResponse, _err error) {
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

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.ShowSize) {
		query["ShowSize"] = request.ShowSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeploymentJob"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeploymentJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificate deployment tasks that are created in your account.
//
// @param request - ListDeploymentJobRequest
//
// @return ListDeploymentJobResponse
func (client *Client) ListDeploymentJob(request *ListDeploymentJobRequest) (_result *ListDeploymentJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDeploymentJobResponse{}
	_body, _err := client.ListDeploymentJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the certificates associated with a deployment task, such as the certificate instance ID, type, and name.
//
// @param request - ListDeploymentJobCertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentJobCertResponse
func (client *Client) ListDeploymentJobCertWithOptions(request *ListDeploymentJobCertRequest, runtime *dara.RuntimeOptions) (_result *ListDeploymentJobCertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeploymentJobCert"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeploymentJobCertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificates associated with a deployment task, such as the certificate instance ID, type, and name.
//
// @param request - ListDeploymentJobCertRequest
//
// @return ListDeploymentJobCertResponse
func (client *Client) ListDeploymentJobCert(request *ListDeploymentJobCertRequest) (_result *ListDeploymentJobCertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDeploymentJobCertResponse{}
	_body, _err := client.ListDeploymentJobCertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the cloud resources associated with a deployment task. An empty list indicates that the resources are invalid and must be re-associated.
//
// @param request - ListDeploymentJobResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentJobResourceResponse
func (client *Client) ListDeploymentJobResourceWithOptions(request *ListDeploymentJobResourceRequest, runtime *dara.RuntimeOptions) (_result *ListDeploymentJobResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeploymentJobResource"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeploymentJobResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cloud resources associated with a deployment task. An empty list indicates that the resources are invalid and must be re-associated.
//
// @param request - ListDeploymentJobResourceRequest
//
// @return ListDeploymentJobResourceResponse
func (client *Client) ListDeploymentJobResource(request *ListDeploymentJobResourceRequest) (_result *ListDeploymentJobResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDeploymentJobResourceResponse{}
	_body, _err := client.ListDeploymentJobResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of instances.
//
// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Brand) {
		query["Brand"] = request.Brand
	}

	if !dara.IsNil(request.CertificateStatus) {
		query["CertificateStatus"] = request.CertificateStatus
	}

	if !dara.IsNil(request.CertificateType) {
		query["CertificateType"] = request.CertificateType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ShowSize) {
		query["ShowSize"] = request.ShowSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of instances.
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the SSL certificates and certificate orders in your account.
//
// Description:
//
// This operation queries a list of your certificates or orders. Set OrderType to CERT or UPLOAD to query certificates. Set OrderType to CPACK or BUY to query orders.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed this limit, API calls are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - ListUserCertificateOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserCertificateOrderResponse
func (client *Client) ListUserCertificateOrderWithOptions(request *ListUserCertificateOrderRequest, runtime *dara.RuntimeOptions) (_result *ListUserCertificateOrderResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ShowSize) {
		query["ShowSize"] = request.ShowSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserCertificateOrder"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserCertificateOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the SSL certificates and certificate orders in your account.
//
// Description:
//
// This operation queries a list of your certificates or orders. Set OrderType to CERT or UPLOAD to query certificates. Set OrderType to CPACK or BUY to query orders.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed this limit, API calls are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - ListUserCertificateOrderRequest
//
// @return ListUserCertificateOrderResponse
func (client *Client) ListUserCertificateOrder(request *ListUserCertificateOrderRequest) (_result *ListUserCertificateOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserCertificateOrderResponse{}
	_body, _err := client.ListUserCertificateOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists warehouses.
//
// Description:
//
// This operation lists your warehouses.
//
// ### QPS limit
//
// This operation has a per-user QPS limit of 10 requests per second. Calls exceeding this limit are throttled, which can affect your business.
//
// @param tmpReq - ListWarehouseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWarehouseResponse
func (client *Client) ListWarehouseWithOptions(tmpReq *ListWarehouseRequest, runtime *dara.RuntimeOptions) (_result *ListWarehouseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListWarehouseShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WarehouseInstanceIds) {
		request.WarehouseInstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WarehouseInstanceIds, dara.String("WarehouseInstanceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.WarehouseTypes) {
		request.WarehouseTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WarehouseTypes, dara.String("WarehouseTypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.WarehouseInstanceIdsShrink) {
		query["WarehouseInstanceIds"] = request.WarehouseInstanceIdsShrink
	}

	if !dara.IsNil(request.WarehouseTypesShrink) {
		query["WarehouseTypes"] = request.WarehouseTypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWarehouse"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWarehouseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists warehouses.
//
// Description:
//
// This operation lists your warehouses.
//
// ### QPS limit
//
// This operation has a per-user QPS limit of 10 requests per second. Calls exceeding this limit are throttled, which can affect your business.
//
// @param request - ListWarehouseRequest
//
// @return ListWarehouseResponse
func (client *Client) ListWarehouse(request *ListWarehouseRequest) (_result *ListWarehouseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWarehouseResponse{}
	_body, _err := client.ListWarehouseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the worker tasks of a deployment task. Each worker task deploys a certificate to a specific cloud resource in a cloud service.
//
// @param request - ListWorkerResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkerResourceResponse
func (client *Client) ListWorkerResourceWithOptions(request *ListWorkerResourceRequest, runtime *dara.RuntimeOptions) (_result *ListWorkerResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloudProduct) {
		query["CloudProduct"] = request.CloudProduct
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.ShowSize) {
		query["ShowSize"] = request.ShowSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkerResource"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkerResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the worker tasks of a deployment task. Each worker task deploys a certificate to a specific cloud resource in a cloud service.
//
// @param request - ListWorkerResourceRequest
//
// @return ListWorkerResourceResponse
func (client *Client) ListWorkerResource(request *ListWorkerResourceRequest) (_result *ListWorkerResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWorkerResourceResponse{}
	_body, _err := client.ListWorkerResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a certificate or certificate order belongs.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveResourceGroup"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a certificate or certificate order belongs.
//
// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Refunds a Certificate Management Service instance if the refund is requested within seven days of purchase.
//
// @param request - RefundInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefundInstanceResponse
func (client *Client) RefundInstanceWithOptions(request *RefundInstanceRequest, runtime *dara.RuntimeOptions) (_result *RefundInstanceResponse, _err error) {
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
		Action:      dara.String("RefundInstance"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefundInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Refunds a Certificate Management Service instance if the refund is requested within seven days of purchase.
//
// @param request - RefundInstanceRequest
//
// @return RefundInstanceResponse
func (client *Client) RefundInstance(request *RefundInstanceRequest) (_result *RefundInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefundInstanceResponse{}
	_body, _err := client.RefundInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a renewal application for an issued SSL certificate.
//
// Description:
//
// You can call the RenewCertificateOrderForPackageRequest operation to submit a renewal application for a certificate only when the order of the certificate is in the expiring state. After the renewal is complete, a new certificate order whose status is pending application is generated. You must submit a certificate application for the new certificate order and install the new certificate after the new certificate is issued.
//
// > You can call the [DescribeCertificateState](https://help.aliyun.com/document_detail/164111.html) operation to query the status of a certificate application order. If the value of the **Type*	- response parameter is **certificate**, the certificate is issued.
//
// @param request - RenewCertificateOrderForPackageRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewCertificateOrderForPackageRequestResponse
func (client *Client) RenewCertificateOrderForPackageRequestWithOptions(request *RenewCertificateOrderForPackageRequestRequest, runtime *dara.RuntimeOptions) (_result *RenewCertificateOrderForPackageRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Csr) {
		query["Csr"] = request.Csr
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewCertificateOrderForPackageRequest"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewCertificateOrderForPackageRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a renewal application for an issued SSL certificate.
//
// Description:
//
// You can call the RenewCertificateOrderForPackageRequest operation to submit a renewal application for a certificate only when the order of the certificate is in the expiring state. After the renewal is complete, a new certificate order whose status is pending application is generated. You must submit a certificate application for the new certificate order and install the new certificate after the new certificate is issued.
//
// > You can call the [DescribeCertificateState](https://help.aliyun.com/document_detail/164111.html) operation to query the status of a certificate application order. If the value of the **Type*	- response parameter is **certificate**, the certificate is issued.
//
// @param request - RenewCertificateOrderForPackageRequestRequest
//
// @return RenewCertificateOrderForPackageRequestResponse
func (client *Client) RenewCertificateOrderForPackageRequest(request *RenewCertificateOrderForPackageRequestRequest) (_result *RenewCertificateOrderForPackageRequestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewCertificateOrderForPackageRequestResponse{}
	_body, _err := client.RenewCertificateOrderForPackageRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes a certificate.
//
// @param request - RevokeCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeCertificateResponse
func (client *Client) RevokeCertificateWithOptions(request *RevokeCertificateRequest, runtime *dara.RuntimeOptions) (_result *RevokeCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertificateId) {
		query["CertificateId"] = request.CertificateId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeCertificate"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes a certificate.
//
// @param request - RevokeCertificateRequest
//
// @return RevokeCertificateResponse
func (client *Client) RevokeCertificate(request *RevokeCertificateRequest) (_result *RevokeCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeCertificateResponse{}
	_body, _err := client.RevokeCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes a client certificate from the certificate repository.
//
// Description:
//
// The rate limit for this API is 10 queries per second (QPS) per user. If you exceed this limit, subsequent API calls will be throttled, which can disrupt your services. We recommend that you call this API at a reasonable rate.
//
// @param request - RevokeWHClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeWHClientCertificateResponse
func (client *Client) RevokeWHClientCertificateWithOptions(request *RevokeWHClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *RevokeWHClientCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Identifier) {
		query["Identifier"] = request.Identifier
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeWHClientCertificate"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeWHClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes a client certificate from the certificate repository.
//
// Description:
//
// The rate limit for this API is 10 queries per second (QPS) per user. If you exceed this limit, subsequent API calls will be throttled, which can disrupt your services. We recommend that you call this API at a reasonable rate.
//
// @param request - RevokeWHClientCertificateRequest
//
// @return RevokeWHClientCertificateResponse
func (client *Client) RevokeWHClientCertificate(request *RevokeWHClientCertificateRequest) (_result *RevokeWHClientCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeWHClientCertificateResponse{}
	_body, _err := client.RevokeWHClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation creates a digital signature with a PCA certificate from a certificate repository.
//
// Description:
//
// This operation creates a digital signature with a PCA certificate from a certificate repository.
//
// ### QPS limit
//
// This operation supports up to 1,000 queries per second (QPS) for a single user. If you exceed this limit, the system throttles your API calls, which can impact your business. Plan your API calls accordingly.
//
// @param request - SignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SignResponse
func (client *Client) SignWithOptions(request *SignRequest, runtime *dara.RuntimeOptions) (_result *SignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.CustomIdentifier) {
		query["CustomIdentifier"] = request.CustomIdentifier
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.MessageType) {
		query["MessageType"] = request.MessageType
	}

	if !dara.IsNil(request.SigningAlgorithm) {
		query["SigningAlgorithm"] = request.SigningAlgorithm
	}

	if !dara.IsNil(request.WarehouseId) {
		query["WarehouseId"] = request.WarehouseId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Sign"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation creates a digital signature with a PCA certificate from a certificate repository.
//
// Description:
//
// This operation creates a digital signature with a PCA certificate from a certificate repository.
//
// ### QPS limit
//
// This operation supports up to 1,000 queries per second (QPS) for a single user. If you exceed this limit, the system throttles your API calls, which can impact your business. Plan your API calls accordingly.
//
// @param request - SignRequest
//
// @return SignResponse
func (client *Client) Sign(request *SignRequest) (_result *SignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SignResponse{}
	_body, _err := client.SignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the private key associated with a certificate signing request (CSR).
//
// @param request - UpdateCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCsrResponse
func (client *Client) UpdateCsrWithOptions(request *UpdateCsrRequest, runtime *dara.RuntimeOptions) (_result *UpdateCsrResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CsrId) {
		query["CsrId"] = request.CsrId
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCsr"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCsrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the private key associated with a certificate signing request (CSR).
//
// @param request - UpdateCsrRequest
//
// @return UpdateCsrResponse
func (client *Client) UpdateCsr(request *UpdateCsrRequest) (_result *UpdateCsrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCsrResponse{}
	_body, _err := client.UpdateCsrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a certificate deployment task, such as the certificates or target resources.
//
// @param request - UpdateDeploymentJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeploymentJobResponse
func (client *Client) UpdateDeploymentJobWithOptions(request *UpdateDeploymentJobRequest, runtime *dara.RuntimeOptions) (_result *UpdateDeploymentJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIds) {
		query["CertIds"] = request.CertIds
	}

	if !dara.IsNil(request.ContactIds) {
		query["ContactIds"] = request.ContactIds
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ScheduleTime) {
		query["ScheduleTime"] = request.ScheduleTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDeploymentJob"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDeploymentJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a certificate deployment task, such as the certificates or target resources.
//
// @param request - UpdateDeploymentJobRequest
//
// @return UpdateDeploymentJobResponse
func (client *Client) UpdateDeploymentJob(request *UpdateDeploymentJobRequest) (_result *UpdateDeploymentJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDeploymentJobResponse{}
	_body, _err := client.UpdateDeploymentJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the status of a certificate deployment task, such as changing from editing to pending execution.
//
// @param request - UpdateDeploymentJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeploymentJobStatusResponse
func (client *Client) UpdateDeploymentJobStatusWithOptions(request *UpdateDeploymentJobStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateDeploymentJobStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDeploymentJobStatus"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDeploymentJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of a certificate deployment task, such as changing from editing to pending execution.
//
// @param request - UpdateDeploymentJobStatusRequest
//
// @return UpdateDeploymentJobStatusResponse
func (client *Client) UpdateDeploymentJobStatus(request *UpdateDeploymentJobStatusRequest) (_result *UpdateDeploymentJobStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDeploymentJobStatusResponse{}
	_body, _err := client.UpdateDeploymentJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a Certificate Management Service instance.
//
// @param request - UpdateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithOptions(request *UpdateInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoReissue) {
		query["AutoReissue"] = request.AutoReissue
	}

	if !dara.IsNil(request.CertificateName) {
		query["CertificateName"] = request.CertificateName
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.CompanyId) {
		query["CompanyId"] = request.CompanyId
	}

	if !dara.IsNil(request.ContactIdList) {
		query["ContactIdList"] = request.ContactIdList
	}

	if !dara.IsNil(request.CountryCode) {
		query["CountryCode"] = request.CountryCode
	}

	if !dara.IsNil(request.Csr) {
		query["Csr"] = request.Csr
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.GenerateCsrMethod) {
		query["GenerateCsrMethod"] = request.GenerateCsrMethod
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.KeyAlgorithm) {
		query["KeyAlgorithm"] = request.KeyAlgorithm
	}

	if !dara.IsNil(request.Province) {
		query["Province"] = request.Province
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.ValidationMethod) {
		query["ValidationMethod"] = request.ValidationMethod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstance"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a Certificate Management Service instance.
//
// @param request - UpdateInstanceRequest
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstance(request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rolls back or re-executes a worker task in a certificate deployment task.
//
// @param request - UpdateWorkerResourceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkerResourceStatusResponse
func (client *Client) UpdateWorkerResourceStatusWithOptions(request *UpdateWorkerResourceStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkerResourceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WorkerId) {
		query["WorkerId"] = request.WorkerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkerResourceStatus"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkerResourceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back or re-executes a worker task in a certificate deployment task.
//
// @param request - UpdateWorkerResourceStatusRequest
//
// @return UpdateWorkerResourceStatusResponse
func (client *Client) UpdateWorkerResourceStatus(request *UpdateWorkerResourceStatusRequest) (_result *UpdateWorkerResourceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWorkerResourceStatusResponse{}
	_body, _err := client.UpdateWorkerResourceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads an existing certificate signing request (CSR) to Certificate Management Service. After the upload, you can use the CSR to apply for certificates.
//
// @param request - UploadCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadCsrResponse
func (client *Client) UploadCsrWithOptions(request *UploadCsrRequest, runtime *dara.RuntimeOptions) (_result *UploadCsrResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Csr) {
		query["Csr"] = request.Csr
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadCsr"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadCsrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads an existing certificate signing request (CSR) to Certificate Management Service. After the upload, you can use the CSR to apply for certificates.
//
// @param request - UploadCsrRequest
//
// @return UploadCsrResponse
func (client *Client) UploadCsr(request *UploadCsrRequest) (_result *UploadCsrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadCsrResponse{}
	_body, _err := client.UploadCsrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads a PCA certificate to a certificate warehouse.
//
// Description:
//
// Use this operation to upload a PCA certificate to a certificate warehouse.
//
// ## QPS limit
//
// The QPS limit for this operation is 10 requests per second per user. Exceeding this limit triggers throttling, which can affect your business.
//
// @param request - UploadPCACertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadPCACertResponse
func (client *Client) UploadPCACertWithOptions(request *UploadPCACertRequest, runtime *dara.RuntimeOptions) (_result *UploadPCACertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cert) {
		query["Cert"] = request.Cert
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PrivateKey) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !dara.IsNil(request.WarehouseId) {
		query["WarehouseId"] = request.WarehouseId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadPCACert"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadPCACertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a PCA certificate to a certificate warehouse.
//
// Description:
//
// Use this operation to upload a PCA certificate to a certificate warehouse.
//
// ## QPS limit
//
// The QPS limit for this operation is 10 requests per second per user. Exceeding this limit triggers throttling, which can affect your business.
//
// @param request - UploadPCACertRequest
//
// @return UploadPCACertResponse
func (client *Client) UploadPCACert(request *UploadPCACertRequest) (_result *UploadPCACertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadPCACertResponse{}
	_body, _err := client.UploadPCACertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads a certificate and its private key to Certificate Management Service. Both SM and non-SM certificates are supported.
//
// Description:
//
// The queries per second (QPS) limit for this operation is 100 for each user. If you exceed this limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - UploadUserCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadUserCertificateResponse
func (client *Client) UploadUserCertificateWithOptions(request *UploadUserCertificateRequest, runtime *dara.RuntimeOptions) (_result *UploadUserCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cert) {
		query["Cert"] = request.Cert
	}

	if !dara.IsNil(request.EncryptCert) {
		query["EncryptCert"] = request.EncryptCert
	}

	if !dara.IsNil(request.EncryptPrivateKey) {
		query["EncryptPrivateKey"] = request.EncryptPrivateKey
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SignCert) {
		query["SignCert"] = request.SignCert
	}

	if !dara.IsNil(request.SignPrivateKey) {
		query["SignPrivateKey"] = request.SignPrivateKey
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadUserCertificate"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadUserCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a certificate and its private key to Certificate Management Service. Both SM and non-SM certificates are supported.
//
// Description:
//
// The queries per second (QPS) limit for this operation is 100 for each user. If you exceed this limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - UploadUserCertificateRequest
//
// @return UploadUserCertificateResponse
func (client *Client) UploadUserCertificate(request *UploadUserCertificateRequest) (_result *UploadUserCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadUserCertificateResponse{}
	_body, _err := client.UploadUserCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies a data signature by using a private certificate in a certificate application repository.
//
// Description:
//
// This API verifies the signatures of PCA certificates and SSL certificates in the certificate repository.
//
// ### QPS limits
//
// The queries per second (QPS) limit for this API is 1,000 for a single user. For your specific QPS limit, refer to the certificate repository. If you exceed this limit, API calls are throttled, which may affect your business. Plan your API calls accordingly.
//
// @param request - VerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyResponse
func (client *Client) VerifyWithOptions(request *VerifyRequest, runtime *dara.RuntimeOptions) (_result *VerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.CustomIdentifier) {
		query["CustomIdentifier"] = request.CustomIdentifier
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.MessageType) {
		query["MessageType"] = request.MessageType
	}

	if !dara.IsNil(request.SignatureValue) {
		query["SignatureValue"] = request.SignatureValue
	}

	if !dara.IsNil(request.SigningAlgorithm) {
		query["SigningAlgorithm"] = request.SigningAlgorithm
	}

	if !dara.IsNil(request.WarehouseId) {
		query["WarehouseId"] = request.WarehouseId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Verify"),
		Version:     dara.String("2020-04-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies a data signature by using a private certificate in a certificate application repository.
//
// Description:
//
// This API verifies the signatures of PCA certificates and SSL certificates in the certificate repository.
//
// ### QPS limits
//
// The queries per second (QPS) limit for this API is 1,000 for a single user. For your specific QPS limit, refer to the certificate repository. If you exceed this limit, API calls are throttled, which may affect your business. Plan your API calls accordingly.
//
// @param request - VerifyRequest
//
// @return VerifyResponse
func (client *Client) Verify(request *VerifyRequest) (_result *VerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyResponse{}
	_body, _err := client.VerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
