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
		"ap-southeast-2":              dara.String("cas.ap-southeast-2.aliyuncs.com"),
		"ap-northeast-1":              dara.String("cas.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-1":              dara.String("cas.ap-southeast-1.aliyuncs.com"),
		"eu-central-1":                dara.String("cas.eu-central-1.aliyuncs.com"),
		"me-central-1":                dara.String("cas.me-central-1.aliyuncs.com"),
		"ap-south-1":                  dara.String("cas.ap-south-1.aliyuncs.com"),
		"me-east-1":                   dara.String("cas.me-east-1.aliyuncs.com"),
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
// Assigns the certificate quota to a subordinate certificate authority (CA) instance.
//
// Description:
//
// ## QPS limit
//
// This API operation is limited to 10 queries per second (QPS) per user. If you exceed this limit, the system throttles your API calls, which can affect your business. Plan your calls accordingly.
//
// @param request - AssignCertificateCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignCertificateCountResponse
func (client *Client) AssignCertificateCountWithOptions(request *AssignCertificateCountRequest, runtime *dara.RuntimeOptions) (_result *AssignCertificateCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertTotalCount) {
		query["CertTotalCount"] = request.CertTotalCount
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssignCertificateCount"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssignCertificateCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns the certificate quota to a subordinate certificate authority (CA) instance.
//
// Description:
//
// ## QPS limit
//
// This API operation is limited to 10 queries per second (QPS) per user. If you exceed this limit, the system throttles your API calls, which can affect your business. Plan your calls accordingly.
//
// @param request - AssignCertificateCountRequest
//
// @return AssignCertificateCountResponse
func (client *Client) AssignCertificateCount(request *AssignCertificateCountRequest) (_result *AssignCertificateCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssignCertificateCountResponse{}
	_body, _err := client.AssignCertificateCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Issues a single client certificate based on a system-generated certificate signing request (CSR).
//
// Description:
//
// Before you call this operation, you must have called [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) to create a root CA certificate and called [CreateSubCACertificate](https://help.aliyun.com/document_detail/465959.html) to create a subordinate CA certificate. Only subordinate CA certificates can issue client certificates.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
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
	if !dara.IsNil(request.AfterTime) {
		query["AfterTime"] = request.AfterTime
	}

	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.AliasName) {
		query["AliasName"] = request.AliasName
	}

	if !dara.IsNil(request.AsynchronousFlag) {
		query["AsynchronousFlag"] = request.AsynchronousFlag
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

	if !dara.IsNil(request.CustomIdentifier) {
		query["CustomIdentifier"] = request.CustomIdentifier
	}

	if !dara.IsNil(request.Days) {
		query["Days"] = request.Days
	}

	if !dara.IsNil(request.EnableCrl) {
		query["EnableCrl"] = request.EnableCrl
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Years) {
		query["Years"] = request.Years
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateClientCertificate"),
		Version:     dara.String("2020-06-30"),
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
// Issues a single client certificate based on a system-generated certificate signing request (CSR).
//
// Description:
//
// Before you call this operation, you must have called [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) to create a root CA certificate and called [CreateSubCACertificate](https://help.aliyun.com/document_detail/465959.html) to create a subordinate CA certificate. Only subordinate CA certificates can issue client certificates.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
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
// Issues a single client certificate based on a custom certificate signing request (CSR).
//
// Description:
//
// Before you call this operation, you must have called [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) to create a root CA certificate and called [CreateSubCACertificate](https://help.aliyun.com/document_detail/465959.html) to create a sub-CA certificate. Only sub-CA certificates can issue client certificates.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation as appropriate.
//
// @param request - CreateClientCertificateWithCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClientCertificateWithCsrResponse
func (client *Client) CreateClientCertificateWithCsrWithOptions(request *CreateClientCertificateWithCsrRequest, runtime *dara.RuntimeOptions) (_result *CreateClientCertificateWithCsrResponse, _err error) {
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

	if !dara.IsNil(request.AsynchronousFlag) {
		query["AsynchronousFlag"] = request.AsynchronousFlag
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

	if !dara.IsNil(request.CustomIdentifier) {
		query["CustomIdentifier"] = request.CustomIdentifier
	}

	if !dara.IsNil(request.Days) {
		query["Days"] = request.Days
	}

	if !dara.IsNil(request.EnableCrl) {
		query["EnableCrl"] = request.EnableCrl
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Years) {
		query["Years"] = request.Years
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateClientCertificateWithCsr"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClientCertificateWithCsrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Issues a single client certificate based on a custom certificate signing request (CSR).
//
// Description:
//
// Before you call this operation, you must have called [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) to create a root CA certificate and called [CreateSubCACertificate](https://help.aliyun.com/document_detail/465959.html) to create a sub-CA certificate. Only sub-CA certificates can issue client certificates.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation as appropriate.
//
// @param request - CreateClientCertificateWithCsrRequest
//
// @return CreateClientCertificateWithCsrResponse
func (client *Client) CreateClientCertificateWithCsr(request *CreateClientCertificateWithCsrRequest) (_result *CreateClientCertificateWithCsrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateClientCertificateWithCsrResponse{}
	_body, _err := client.CreateClientCertificateWithCsrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Issues a custom digital certificate with specified subject, subject alternative names (SANs), key usage, and extended key usage attributes.
//
// Description:
//
// By default, the certificate subject is retrieved from the Certificate Signing Request (CSR). If you specify a certificate subject, the subject from the CSR is ignored and the specified subject is used to issue the certificate.
//
// You must specify the key usage or extended key usage based on your scenario. The following examples show common scenarios:
//
// - Server-side authentication certificate
//
// Key usage: digitalSignature, keyEncipherment
//
// Extended key usage: serverAuth
//
// - Client authentication certificate
//
// Key usage: digitalSignature, keyEncipherment
//
// Extended key usage: clientAuth
//
// - mTLS mutual authentication certificate
//
// Key usage: digitalSignature, keyEncipherment
//
// Extended key usage: serverAuth, clientAuth
//
// - Email signing certificate
//
// Key usage: digitalSignature, contentCommitment
//
// Extended key usage: emailProtection
//
// Note: Compliance CAs are managed by third-party authorities and do not support this operation.
//
// @param request - CreateCustomCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomCertificateResponse
func (client *Client) CreateCustomCertificateWithOptions(request *CreateCustomCertificateRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiPassthrough) {
		query["ApiPassthrough"] = request.ApiPassthrough
	}

	if !dara.IsNil(request.Csr) {
		query["Csr"] = request.Csr
	}

	if !dara.IsNil(request.EnableCrl) {
		query["EnableCrl"] = request.EnableCrl
	}

	if !dara.IsNil(request.Immediately) {
		query["Immediately"] = request.Immediately
	}

	if !dara.IsNil(request.ParentIdentifier) {
		query["ParentIdentifier"] = request.ParentIdentifier
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Validity) {
		query["Validity"] = request.Validity
	}

	if !dara.IsNil(request.CustomIdentifier) {
		query["customIdentifier"] = request.CustomIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomCertificate"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Issues a custom digital certificate with specified subject, subject alternative names (SANs), key usage, and extended key usage attributes.
//
// Description:
//
// By default, the certificate subject is retrieved from the Certificate Signing Request (CSR). If you specify a certificate subject, the subject from the CSR is ignored and the specified subject is used to issue the certificate.
//
// You must specify the key usage or extended key usage based on your scenario. The following examples show common scenarios:
//
// - Server-side authentication certificate
//
// Key usage: digitalSignature, keyEncipherment
//
// Extended key usage: serverAuth
//
// - Client authentication certificate
//
// Key usage: digitalSignature, keyEncipherment
//
// Extended key usage: clientAuth
//
// - mTLS mutual authentication certificate
//
// Key usage: digitalSignature, keyEncipherment
//
// Extended key usage: serverAuth, clientAuth
//
// - Email signing certificate
//
// Key usage: digitalSignature, contentCommitment
//
// Extended key usage: emailProtection
//
// Note: Compliance CAs are managed by third-party authorities and do not support this operation.
//
// @param request - CreateCustomCertificateRequest
//
// @return CreateCustomCertificateResponse
func (client *Client) CreateCustomCertificate(request *CreateCustomCertificateRequest) (_result *CreateCustomCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCustomCertificateResponse{}
	_body, _err := client.CreateCustomCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create and issue an external subordinate CA certificate using a CSR and API parameters.
//
// Description:
//
// ## Request
//
// - Creates an external subordinate CA certificate from a certificate signing request (CSR) and optional API pass-through parameters.
//
// - The required `InstanceId` parameter specifies the instance ID of the external subordinate CA.
//
// - The `Csr` parameter must contain a valid certificate signing request.
//
// - The `Validity` parameter specifies the certificate\\"s validity period and accepts values in either relative or absolute time formats.
//
// - The `ApiPassthrough` parameter lets you override information in the CSR, such as subject information, or add certificate extensions.
//
// - Note: For end-entity CA certificates, set the `pathLenConstraint` parameter to 0.
//
// @param tmpReq - CreateExternalCACertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExternalCACertificateResponse
func (client *Client) CreateExternalCACertificateWithOptions(tmpReq *CreateExternalCACertificateRequest, runtime *dara.RuntimeOptions) (_result *CreateExternalCACertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateExternalCACertificateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApiPassthrough) {
		request.ApiPassthroughShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApiPassthrough, dara.String("ApiPassthrough"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiPassthroughShrink) {
		query["ApiPassthrough"] = request.ApiPassthroughShrink
	}

	if !dara.IsNil(request.CertMaxTime) {
		query["CertMaxTime"] = request.CertMaxTime
	}

	if !dara.IsNil(request.Csr) {
		query["Csr"] = request.Csr
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Validity) {
		query["Validity"] = request.Validity
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExternalCACertificate"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExternalCACertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create and issue an external subordinate CA certificate using a CSR and API parameters.
//
// Description:
//
// ## Request
//
// - Creates an external subordinate CA certificate from a certificate signing request (CSR) and optional API pass-through parameters.
//
// - The required `InstanceId` parameter specifies the instance ID of the external subordinate CA.
//
// - The `Csr` parameter must contain a valid certificate signing request.
//
// - The `Validity` parameter specifies the certificate\\"s validity period and accepts values in either relative or absolute time formats.
//
// - The `ApiPassthrough` parameter lets you override information in the CSR, such as subject information, or add certificate extensions.
//
// - Note: For end-entity CA certificates, set the `pathLenConstraint` parameter to 0.
//
// @param request - CreateExternalCACertificateRequest
//
// @return CreateExternalCACertificateResponse
func (client *Client) CreateExternalCACertificate(request *CreateExternalCACertificateRequest) (_result *CreateExternalCACertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateExternalCACertificateResponse{}
	_body, _err := client.CreateExternalCACertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes a client certificate or a server certificate issued by a private certificate authority (CA).
//
// Description:
//
// After a client or server certificate is revoked, the client or server where the certificate is installed cannot establish HTTPS connections with other devices.
//
// After a client or server certificate is revoked, you can call [DeleteClientCertificate](https://help.aliyun.com/document_detail/465981.html) to permanently delete the certificate.
//
// ## QPS limit
//
// The limit on queries per second (QPS) for this operation is 10 per user. If you exceed this limit, API calls are throttled, which can affect your business. Plan your API calls accordingly.
//
// @param request - CreateRevokeClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRevokeClientCertificateResponse
func (client *Client) CreateRevokeClientCertificateWithOptions(request *CreateRevokeClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *CreateRevokeClientCertificateResponse, _err error) {
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
		Action:      dara.String("CreateRevokeClientCertificate"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRevokeClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes a client certificate or a server certificate issued by a private certificate authority (CA).
//
// Description:
//
// After a client or server certificate is revoked, the client or server where the certificate is installed cannot establish HTTPS connections with other devices.
//
// After a client or server certificate is revoked, you can call [DeleteClientCertificate](https://help.aliyun.com/document_detail/465981.html) to permanently delete the certificate.
//
// ## QPS limit
//
// The limit on queries per second (QPS) for this operation is 10 per user. If you exceed this limit, API calls are throttled, which can affect your business. Plan your API calls accordingly.
//
// @param request - CreateRevokeClientCertificateRequest
//
// @return CreateRevokeClientCertificateResponse
func (client *Client) CreateRevokeClientCertificate(request *CreateRevokeClientCertificateRequest) (_result *CreateRevokeClientCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRevokeClientCertificateResponse{}
	_body, _err := client.CreateRevokeClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a root certificate authority (CA) certificate.
//
// Description:
//
// This operation creates a self-signed root CA certificate. A root CA certificate is the starting point of a private trust chain within an enterprise. After you create a root CA certificate, you can use it to issue intermediate CA certificates. You can then use the intermediate CA certificates to issue client and server-side certificates.
//
// Before calling this operation, purchase a private root CA in the [SSL Certificate Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). Otherwise, the call fails. For more information, see [Purchase a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second per user. If you exceed the limit, API calls are throttled, which may affect your business. Call the API at a reasonable rate.
//
// @param request - CreateRootCACertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRootCACertificateResponse
func (client *Client) CreateRootCACertificateWithOptions(request *CreateRootCACertificateRequest, runtime *dara.RuntimeOptions) (_result *CreateRootCACertificateResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CommonName) {
		query["CommonName"] = request.CommonName
	}

	if !dara.IsNil(request.CountryCode) {
		query["CountryCode"] = request.CountryCode
	}

	if !dara.IsNil(request.Locality) {
		query["Locality"] = request.Locality
	}

	if !dara.IsNil(request.Organization) {
		query["Organization"] = request.Organization
	}

	if !dara.IsNil(request.OrganizationUnit) {
		query["OrganizationUnit"] = request.OrganizationUnit
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Years) {
		query["Years"] = request.Years
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRootCACertificate"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRootCACertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a root certificate authority (CA) certificate.
//
// Description:
//
// This operation creates a self-signed root CA certificate. A root CA certificate is the starting point of a private trust chain within an enterprise. After you create a root CA certificate, you can use it to issue intermediate CA certificates. You can then use the intermediate CA certificates to issue client and server-side certificates.
//
// Before calling this operation, purchase a private root CA in the [SSL Certificate Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). Otherwise, the call fails. For more information, see [Purchase a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second per user. If you exceed the limit, API calls are throttled, which may affect your business. Call the API at a reasonable rate.
//
// @param request - CreateRootCACertificateRequest
//
// @return CreateRootCACertificateResponse
func (client *Client) CreateRootCACertificate(request *CreateRootCACertificateRequest) (_result *CreateRootCACertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRootCACertificateResponse{}
	_body, _err := client.CreateRootCACertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Issues a single server certificate based on a system-generated certificate signing request (CSR).
//
// Description:
//
// Before you call this operation, you must have called [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) to create a root CA certificate and called [CreateSubCACertificate](https://help.aliyun.com/document_detail/465975.html) to create a subordinate CA certificate. Only subordinate CA certificates can issue server certificates.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - CreateServerCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServerCertificateResponse
func (client *Client) CreateServerCertificateWithOptions(request *CreateServerCertificateRequest, runtime *dara.RuntimeOptions) (_result *CreateServerCertificateResponse, _err error) {
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

	if !dara.IsNil(request.AsynchronousFlag) {
		query["AsynchronousFlag"] = request.AsynchronousFlag
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

	if !dara.IsNil(request.CustomIdentifier) {
		query["CustomIdentifier"] = request.CustomIdentifier
	}

	if !dara.IsNil(request.Days) {
		query["Days"] = request.Days
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.EnableCrl) {
		query["EnableCrl"] = request.EnableCrl
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Years) {
		query["Years"] = request.Years
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServerCertificate"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServerCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Issues a single server certificate based on a system-generated certificate signing request (CSR).
//
// Description:
//
// Before you call this operation, you must have called [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) to create a root CA certificate and called [CreateSubCACertificate](https://help.aliyun.com/document_detail/465975.html) to create a subordinate CA certificate. Only subordinate CA certificates can issue server certificates.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - CreateServerCertificateRequest
//
// @return CreateServerCertificateResponse
func (client *Client) CreateServerCertificate(request *CreateServerCertificateRequest) (_result *CreateServerCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateServerCertificateResponse{}
	_body, _err := client.CreateServerCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Issues a single server certificate based on a custom CSR.
//
// Description:
//
// Before you call this operation, you must have called [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) to create a root CA certificate and called [CreateSubCACertificate](https://help.aliyun.com/document_detail/465959.html) to create a subordinate CA certificate. Only subordinate CA certificates can issue server certificates.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation as appropriate.
//
// @param request - CreateServerCertificateWithCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServerCertificateWithCsrResponse
func (client *Client) CreateServerCertificateWithCsrWithOptions(request *CreateServerCertificateWithCsrRequest, runtime *dara.RuntimeOptions) (_result *CreateServerCertificateWithCsrResponse, _err error) {
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

	if !dara.IsNil(request.AsynchronousFlag) {
		query["AsynchronousFlag"] = request.AsynchronousFlag
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

	if !dara.IsNil(request.CustomIdentifier) {
		query["CustomIdentifier"] = request.CustomIdentifier
	}

	if !dara.IsNil(request.Days) {
		query["Days"] = request.Days
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.EnableCrl) {
		query["EnableCrl"] = request.EnableCrl
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Years) {
		query["Years"] = request.Years
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServerCertificateWithCsr"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServerCertificateWithCsrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Issues a single server certificate based on a custom CSR.
//
// Description:
//
// Before you call this operation, you must have called [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) to create a root CA certificate and called [CreateSubCACertificate](https://help.aliyun.com/document_detail/465959.html) to create a subordinate CA certificate. Only subordinate CA certificates can issue server certificates.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation as appropriate.
//
// @param request - CreateServerCertificateWithCsrRequest
//
// @return CreateServerCertificateWithCsrResponse
func (client *Client) CreateServerCertificateWithCsr(request *CreateServerCertificateWithCsrRequest) (_result *CreateServerCertificateWithCsrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateServerCertificateWithCsrResponse{}
	_body, _err := client.CreateServerCertificateWithCsrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a subordinate certificate authority (CA) certificate under an existing root CA.
//
// Description:
//
// This operation issues an intermediate CA certificate from an existing root CA certificate. You can use the intermediate CA certificate to issue client and server certificates.
//
// Before calling this operation, you must call the [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) operation to create a root CA certificate.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per user. Calls that exceed this limit are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - CreateSubCACertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSubCACertificateResponse
func (client *Client) CreateSubCACertificateWithOptions(request *CreateSubCACertificateRequest, runtime *dara.RuntimeOptions) (_result *CreateSubCACertificateResponse, _err error) {
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

	if !dara.IsNil(request.CertMaxTime) {
		query["CertMaxTime"] = request.CertMaxTime
	}

	if !dara.IsNil(request.CommonName) {
		query["CommonName"] = request.CommonName
	}

	if !dara.IsNil(request.CountryCode) {
		query["CountryCode"] = request.CountryCode
	}

	if !dara.IsNil(request.CrlDay) {
		query["CrlDay"] = request.CrlDay
	}

	if !dara.IsNil(request.EnableCrl) {
		query["EnableCrl"] = request.EnableCrl
	}

	if !dara.IsNil(request.ExtendedKeyUsages) {
		query["ExtendedKeyUsages"] = request.ExtendedKeyUsages
	}

	if !dara.IsNil(request.Locality) {
		query["Locality"] = request.Locality
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

	if !dara.IsNil(request.PathLenConstraint) {
		query["PathLenConstraint"] = request.PathLenConstraint
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Years) {
		query["Years"] = request.Years
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSubCACertificate"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSubCACertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a subordinate certificate authority (CA) certificate under an existing root CA.
//
// Description:
//
// This operation issues an intermediate CA certificate from an existing root CA certificate. You can use the intermediate CA certificate to issue client and server certificates.
//
// Before calling this operation, you must call the [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) operation to create a root CA certificate.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per user. Calls that exceed this limit are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - CreateSubCACertificateRequest
//
// @return CreateSubCACertificateResponse
func (client *Client) CreateSubCACertificate(request *CreateSubCACertificateRequest) (_result *CreateSubCACertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSubCACertificateResponse{}
	_body, _err := client.CreateSubCACertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a revoked client certificate or server certificate. Only revoked certificates can be deleted.
//
// Description:
//
// Before you call this operation, you must call [CreateRevokeClientCertificate](https://help.aliyun.com/document_detail/465972.html) to revoke the client or server-side certificate.
//
// ## QPS limit
//
// This operation supports up to 10 queries per second (QPS) for each user. If you exceed the limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.Identifier) {
		query["Identifier"] = request.Identifier
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteClientCertificate"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
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
// Deletes a revoked client certificate or server certificate. Only revoked certificates can be deleted.
//
// Description:
//
// Before you call this operation, you must call [CreateRevokeClientCertificate](https://help.aliyun.com/document_detail/465972.html) to revoke the client or server-side certificate.
//
// ## QPS limit
//
// This operation supports up to 10 queries per second (QPS) for each user. If you exceed the limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
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
// Queries the details of a CA certificate.
//
// Description:
//
// This operation queries the details of a root CA certificate or sub-CA certificate by using the unique identifier of the certificate. The details include the serial number, subject information, and certificate content of the CA certificate.
//
// Before you call this operation, you must have called [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) to create a root CA certificate and [CreateSubCACertificate](https://help.aliyun.com/document_detail/465959.html) to create a sub-CA certificate.
//
// ## QPS limit
//
// The China site (aliyun.com) allows up to 10 queries per second (QPS) for a single user. If the number of calls per second exceeds the limit, throttling is triggered. Throttling may affect your business. Call this operation as appropriate.
//
// @param request - DescribeCACertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCACertificateResponse
func (client *Client) DescribeCACertificateWithOptions(request *DescribeCACertificateRequest, runtime *dara.RuntimeOptions) (_result *DescribeCACertificateResponse, _err error) {
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
		Action:      dara.String("DescribeCACertificate"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCACertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a CA certificate.
//
// Description:
//
// This operation queries the details of a root CA certificate or sub-CA certificate by using the unique identifier of the certificate. The details include the serial number, subject information, and certificate content of the CA certificate.
//
// Before you call this operation, you must have called [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) to create a root CA certificate and [CreateSubCACertificate](https://help.aliyun.com/document_detail/465959.html) to create a sub-CA certificate.
//
// ## QPS limit
//
// The China site (aliyun.com) allows up to 10 queries per second (QPS) for a single user. If the number of calls per second exceeds the limit, throttling is triggered. Throttling may affect your business. Call this operation as appropriate.
//
// @param request - DescribeCACertificateRequest
//
// @return DescribeCACertificateResponse
func (client *Client) DescribeCACertificate(request *DescribeCACertificateRequest) (_result *DescribeCACertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCACertificateResponse{}
	_body, _err := client.DescribeCACertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of certificate authority (CA) certificates that you created.
//
// Description:
//
// This operation queries the number of CA certificates that you have created, including root CA certificates and subordinate CA certificates.
//
// ## QPS limit
//
// Each user is limited to 10 queries per second (QPS) for this API operation. If you exceed the limit, your API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCACertificateCountResponse
func (client *Client) DescribeCACertificateCountWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeCACertificateCountResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCACertificateCount"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCACertificateCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of certificate authority (CA) certificates that you created.
//
// Description:
//
// This operation queries the number of CA certificates that you have created, including root CA certificates and subordinate CA certificates.
//
// ## QPS limit
//
// Each user is limited to 10 queries per second (QPS) for this API operation. If you exceed the limit, your API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @return DescribeCACertificateCountResponse
func (client *Client) DescribeCACertificateCount() (_result *DescribeCACertificateCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCACertificateCountResponse{}
	_body, _err := client.DescribeCACertificateCountWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all root and subordinate certificate authority (CA) certificates.
//
// Description:
//
// You can call this operation to query the details of all your root and intermediate CA certificates by page. The details include the unique identifier, serial number, subject information, and content of each certificate.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 for each user. If you exceed the limit, API calls are throttled. This may affect your business. We recommend that you call this operation a reasonable number of times.
//
// @param request - DescribeCACertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCACertificateListResponse
func (client *Client) DescribeCACertificateListWithOptions(request *DescribeCACertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCACertificateListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaStatus) {
		query["CaStatus"] = request.CaStatus
	}

	if !dara.IsNil(request.CertType) {
		query["CertType"] = request.CertType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Identifier) {
		query["Identifier"] = request.Identifier
	}

	if !dara.IsNil(request.IssuerType) {
		query["IssuerType"] = request.IssuerType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ShowSize) {
		query["ShowSize"] = request.ShowSize
	}

	if !dara.IsNil(request.ValidStatus) {
		query["ValidStatus"] = request.ValidStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCACertificateList"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCACertificateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all root and subordinate certificate authority (CA) certificates.
//
// Description:
//
// You can call this operation to query the details of all your root and intermediate CA certificates by page. The details include the unique identifier, serial number, subject information, and content of each certificate.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 for each user. If you exceed the limit, API calls are throttled. This may affect your business. We recommend that you call this operation a reasonable number of times.
//
// @param request - DescribeCACertificateListRequest
//
// @return DescribeCACertificateListResponse
func (client *Client) DescribeCACertificateList(request *DescribeCACertificateListRequest) (_result *DescribeCACertificateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCACertificateListResponse{}
	_body, _err := client.DescribeCACertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the encrypted private key of a client certificate or server certificate.
//
// Description:
//
// This API applies only to certificates that are issued from a system-generated Certificate Signing Request (CSR). You can use this API to retrieve the encrypted private key of a client certificate or a server-side certificate. Before you call this API, you must have issued a client or server-side certificate by calling one of the following APIs:
//
// - [CreateClientCertificate](https://help.aliyun.com/document_detail/465967.html)
//
// - [CreateServerCertificate](https://help.aliyun.com/document_detail/465975.html)
//
// To keep the private key secure during transmission, this API uses a password that you set to encrypt the private key. The API then returns the encrypted private key. After you retrieve the encrypted private key, you can decrypt it using one of the following methods:
//
// - If the certificate uses the RSA encryption algorithm, run the `openssl rsa -in <encrypted_private_key_file> -passin pass:<private_key_password> -out <decrypted_private_key_file>` command to decrypt the private key. You must run this command on a computer that has [OpenSSL](https://www.openssl.org/source/) or [BabaSSL](https://github.com/BabaSSL/BabaSSL) installed.
//
// - If the certificate uses the ECC encryption algorithm, run the `openssl ec -in <encrypted_private_key_file> -passin pass:<private_key_password> -out <decrypted_private_key_file>` command to decrypt the private key. You must run this command on a computer that has [OpenSSL](https://www.openssl.org/source/) or [BabaSSL](https://github.com/BabaSSL/BabaSSL) installed.
//
// - If the certificate uses the SM2 encryption algorithm, run the `openssl ec -in <encrypted_private_key_file> -passin pass:<private_key_password> -out <decrypted_private_key_file>` command to decrypt the private key. You must run this command on a computer that has [BabaSSL](https://github.com/BabaSSL/BabaSSL) installed.
//
// > You can call [DescribeClientCertificate](https://help.aliyun.com/document_detail/465985.html) to query the encryption algorithm of the client or server-side certificate.
//
// ## QPS limit
//
// This API has a queries per second (QPS) limit of 10 for each user. If you exceed this limit, your API calls are throttled. Throttling can affect your business. Plan your calls accordingly.
//
// @param request - DescribeCertificatePrivateKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCertificatePrivateKeyResponse
func (client *Client) DescribeCertificatePrivateKeyWithOptions(request *DescribeCertificatePrivateKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeCertificatePrivateKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncryptedCode) {
		query["EncryptedCode"] = request.EncryptedCode
	}

	if !dara.IsNil(request.Identifier) {
		query["Identifier"] = request.Identifier
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCertificatePrivateKey"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCertificatePrivateKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the encrypted private key of a client certificate or server certificate.
//
// Description:
//
// This API applies only to certificates that are issued from a system-generated Certificate Signing Request (CSR). You can use this API to retrieve the encrypted private key of a client certificate or a server-side certificate. Before you call this API, you must have issued a client or server-side certificate by calling one of the following APIs:
//
// - [CreateClientCertificate](https://help.aliyun.com/document_detail/465967.html)
//
// - [CreateServerCertificate](https://help.aliyun.com/document_detail/465975.html)
//
// To keep the private key secure during transmission, this API uses a password that you set to encrypt the private key. The API then returns the encrypted private key. After you retrieve the encrypted private key, you can decrypt it using one of the following methods:
//
// - If the certificate uses the RSA encryption algorithm, run the `openssl rsa -in <encrypted_private_key_file> -passin pass:<private_key_password> -out <decrypted_private_key_file>` command to decrypt the private key. You must run this command on a computer that has [OpenSSL](https://www.openssl.org/source/) or [BabaSSL](https://github.com/BabaSSL/BabaSSL) installed.
//
// - If the certificate uses the ECC encryption algorithm, run the `openssl ec -in <encrypted_private_key_file> -passin pass:<private_key_password> -out <decrypted_private_key_file>` command to decrypt the private key. You must run this command on a computer that has [OpenSSL](https://www.openssl.org/source/) or [BabaSSL](https://github.com/BabaSSL/BabaSSL) installed.
//
// - If the certificate uses the SM2 encryption algorithm, run the `openssl ec -in <encrypted_private_key_file> -passin pass:<private_key_password> -out <decrypted_private_key_file>` command to decrypt the private key. You must run this command on a computer that has [BabaSSL](https://github.com/BabaSSL/BabaSSL) installed.
//
// > You can call [DescribeClientCertificate](https://help.aliyun.com/document_detail/465985.html) to query the encryption algorithm of the client or server-side certificate.
//
// ## QPS limit
//
// This API has a queries per second (QPS) limit of 10 for each user. If you exceed this limit, your API calls are throttled. Throttling can affect your business. Plan your calls accordingly.
//
// @param request - DescribeCertificatePrivateKeyRequest
//
// @return DescribeCertificatePrivateKeyResponse
func (client *Client) DescribeCertificatePrivateKey(request *DescribeCertificatePrivateKeyRequest) (_result *DescribeCertificatePrivateKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCertificatePrivateKeyResponse{}
	_body, _err := client.DescribeCertificatePrivateKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a client certificate or server certificate by its unique identifier.
//
// Description:
//
// You can call this operation to query the details of a client certificate or a server-side certificate by its unique identifier. The details include the serial number, subject, content, and status of the certificate.
//
// Before you call this operation, you must create a client certificate or a server-side certificate.
//
// To create a client certificate by calling an API, see the following topics:
//
// - [CreateClientCertificate](https://help.aliyun.com/document_detail/465967.html)
//
// - [CreateClientCertificateWithCsr](https://help.aliyun.com/document_detail/465970.html)
//
// To create a server-side certificate by calling an API, see the following topics:
//
// - [CreateServerCertificate](https://help.aliyun.com/document_detail/465975.html)
//
// - [CreateServerCertificateWithCsr](https://help.aliyun.com/document_detail/465979.html)
//
// ## Limits
//
// The queries per second (QPS) limit for this API call is 10 per user. If you exceed this limit, throttling is triggered, which may affect your business. Plan your calls accordingly.
//
// @param request - DescribeClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClientCertificateResponse
func (client *Client) DescribeClientCertificateWithOptions(request *DescribeClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *DescribeClientCertificateResponse, _err error) {
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
		Action:      dara.String("DescribeClientCertificate"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a client certificate or server certificate by its unique identifier.
//
// Description:
//
// You can call this operation to query the details of a client certificate or a server-side certificate by its unique identifier. The details include the serial number, subject, content, and status of the certificate.
//
// Before you call this operation, you must create a client certificate or a server-side certificate.
//
// To create a client certificate by calling an API, see the following topics:
//
// - [CreateClientCertificate](https://help.aliyun.com/document_detail/465967.html)
//
// - [CreateClientCertificateWithCsr](https://help.aliyun.com/document_detail/465970.html)
//
// To create a server-side certificate by calling an API, see the following topics:
//
// - [CreateServerCertificate](https://help.aliyun.com/document_detail/465975.html)
//
// - [CreateServerCertificateWithCsr](https://help.aliyun.com/document_detail/465979.html)
//
// ## Limits
//
// The queries per second (QPS) limit for this API call is 10 per user. If you exceed this limit, throttling is triggered, which may affect your business. Plan your calls accordingly.
//
// @param request - DescribeClientCertificateRequest
//
// @return DescribeClientCertificateResponse
func (client *Client) DescribeClientCertificate(request *DescribeClientCertificateRequest) (_result *DescribeClientCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClientCertificateResponse{}
	_body, _err := client.DescribeClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeClientCertificateForSerialNumber is deprecated, please use cas::2020-06-30::DescribeClientCertificate instead.
//
// Summary:
//
// Retrieves the details of multiple client or server certificates by serial number.
//
// Description:
//
// # Usage
//
// Retrieves the details of multiple client or server certificates by serial number. The response includes each certificate\\"s serial number, subject information, content, and status.
//
// Before calling this operation, ensure you have created a client certificate or a server certificate.
//
// To create a client certificate, see:
//
// - [CreateClientCertificate](https://help.aliyun.com/document_detail/330873.html)
//
// - [CreateClientCertificateWithCsr](https://help.aliyun.com/document_detail/330875.html)
//
// To create a server certificate, see:
//
// - [CreateServerCertificate](https://help.aliyun.com/document_detail/330877.html)
//
// - [CreateServerCertificateWithCsr](https://help.aliyun.com/document_detail/330878.html)
//
// # QPS limit
//
// The QPS limit for this operation is 10 calls per second per account. Exceeding this limit triggers throttling, which can impact your business. Plan your calls accordingly.
//
// @param request - DescribeClientCertificateForSerialNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClientCertificateForSerialNumberResponse
func (client *Client) DescribeClientCertificateForSerialNumberWithOptions(request *DescribeClientCertificateForSerialNumberRequest, runtime *dara.RuntimeOptions) (_result *DescribeClientCertificateForSerialNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClientCertificateForSerialNumber"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClientCertificateForSerialNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeClientCertificateForSerialNumber is deprecated, please use cas::2020-06-30::DescribeClientCertificate instead.
//
// Summary:
//
// Retrieves the details of multiple client or server certificates by serial number.
//
// Description:
//
// # Usage
//
// Retrieves the details of multiple client or server certificates by serial number. The response includes each certificate\\"s serial number, subject information, content, and status.
//
// Before calling this operation, ensure you have created a client certificate or a server certificate.
//
// To create a client certificate, see:
//
// - [CreateClientCertificate](https://help.aliyun.com/document_detail/330873.html)
//
// - [CreateClientCertificateWithCsr](https://help.aliyun.com/document_detail/330875.html)
//
// To create a server certificate, see:
//
// - [CreateServerCertificate](https://help.aliyun.com/document_detail/330877.html)
//
// - [CreateServerCertificateWithCsr](https://help.aliyun.com/document_detail/330878.html)
//
// # QPS limit
//
// The QPS limit for this operation is 10 calls per second per account. Exceeding this limit triggers throttling, which can impact your business. Plan your calls accordingly.
//
// @param request - DescribeClientCertificateForSerialNumberRequest
//
// @return DescribeClientCertificateForSerialNumberResponse
// Deprecated
func (client *Client) DescribeClientCertificateForSerialNumber(request *DescribeClientCertificateForSerialNumberRequest) (_result *DescribeClientCertificateForSerialNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClientCertificateForSerialNumberResponse{}
	_body, _err := client.DescribeClientCertificateForSerialNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a client certificate or server certificate by its unique identifier.
//
// Description:
//
// This operation queries the status of multiple client or server-side certificates in a batch using their unique identifiers. For example, you can check whether a certificate is revoked.
//
// ## QPS limits
//
// This operation is limited to 10 queries per second (QPS) for each user. API calls that exceed this limit are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DescribeClientCertificateStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClientCertificateStatusResponse
func (client *Client) DescribeClientCertificateStatusWithOptions(request *DescribeClientCertificateStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeClientCertificateStatusResponse, _err error) {
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
		Action:      dara.String("DescribeClientCertificateStatus"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClientCertificateStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a client certificate or server certificate by its unique identifier.
//
// Description:
//
// This operation queries the status of multiple client or server-side certificates in a batch using their unique identifiers. For example, you can check whether a certificate is revoked.
//
// ## QPS limits
//
// This operation is limited to 10 queries per second (QPS) for each user. API calls that exceed this limit are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DescribeClientCertificateStatusRequest
//
// @return DescribeClientCertificateStatusResponse
func (client *Client) DescribeClientCertificateStatus(request *DescribeClientCertificateStatusRequest) (_result *DescribeClientCertificateStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClientCertificateStatusResponse{}
	_body, _err := client.DescribeClientCertificateStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of client and server-side certificates using their serial numbers.
//
// Description:
//
// This operation queries the status of one or more client or server certificates by their serial numbers. For example, you can check whether a certificate is revoked.
//
// ## QPS limit
//
// You can make up to 10 API calls per second per account. Exceeding this limit triggers throttling, which can impact your business. Plan your calls accordingly.
//
// @param request - DescribeClientCertificateStatusForSerialNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClientCertificateStatusForSerialNumberResponse
func (client *Client) DescribeClientCertificateStatusForSerialNumberWithOptions(request *DescribeClientCertificateStatusForSerialNumberRequest, runtime *dara.RuntimeOptions) (_result *DescribeClientCertificateStatusForSerialNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClientCertificateStatusForSerialNumber"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClientCertificateStatusForSerialNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of client and server-side certificates using their serial numbers.
//
// Description:
//
// This operation queries the status of one or more client or server certificates by their serial numbers. For example, you can check whether a certificate is revoked.
//
// ## QPS limit
//
// You can make up to 10 API calls per second per account. Exceeding this limit triggers throttling, which can impact your business. Plan your calls accordingly.
//
// @param request - DescribeClientCertificateStatusForSerialNumberRequest
//
// @return DescribeClientCertificateStatusForSerialNumberResponse
func (client *Client) DescribeClientCertificateStatusForSerialNumber(request *DescribeClientCertificateStatusForSerialNumberRequest) (_result *DescribeClientCertificateStatusForSerialNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClientCertificateStatusForSerialNumberResponse{}
	_body, _err := client.DescribeClientCertificateStatusForSerialNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Returns all your certification authority (CA) certificates. These include certificates generated by Private CA and imported external certificates.
//
// Description:
//
// This operation performs a paged query to retrieve the details of all your CA certificates, including root and subordinate CA certificates. These details include the unique identifier, serial number, subject information, and certificate content.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10. The system throttles API calls that exceed this limit, which may affect your business. To prevent throttling, call this operation within the stated limit.
//
// @param request - DescribePcaAndExternalCACertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePcaAndExternalCACertificateListResponse
func (client *Client) DescribePcaAndExternalCACertificateListWithOptions(request *DescribePcaAndExternalCACertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribePcaAndExternalCACertificateListResponse, _err error) {
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

	if !dara.IsNil(request.Identifiers) {
		query["Identifiers"] = request.Identifiers
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
		Action:      dara.String("DescribePcaAndExternalCACertificateList"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePcaAndExternalCACertificateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns all your certification authority (CA) certificates. These include certificates generated by Private CA and imported external certificates.
//
// Description:
//
// This operation performs a paged query to retrieve the details of all your CA certificates, including root and subordinate CA certificates. These details include the unique identifier, serial number, subject information, and certificate content.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10. The system throttles API calls that exceed this limit, which may affect your business. To prevent throttling, call this operation within the stated limit.
//
// @param request - DescribePcaAndExternalCACertificateListRequest
//
// @return DescribePcaAndExternalCACertificateListResponse
func (client *Client) DescribePcaAndExternalCACertificateList(request *DescribePcaAndExternalCACertificateListRequest) (_result *DescribePcaAndExternalCACertificateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePcaAndExternalCACertificateListResponse{}
	_body, _err := client.DescribePcaAndExternalCACertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a private certificate authority (CA) instance, including the remaining certificate quota.
//
// Description:
//
// Use this operation to query the status of a private CA instance by its ID. The status information includes the instance\\"s status, the total number of certificates it can issue, and the number of certificates already issued.
//
// Before you call this operation, purchase a private CA in the [CAS console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Purchase a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## QPS limits
//
// This operation has a queries per second (QPS) limit of 10 for each user. If you exceed this limit, API calls are throttled, which can affect your business. Call this operation at a reasonable rate.
//
// @param request - GetCAInstanceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCAInstanceStatusResponse
func (client *Client) GetCAInstanceStatusWithOptions(request *GetCAInstanceStatusRequest, runtime *dara.RuntimeOptions) (_result *GetCAInstanceStatusResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCAInstanceStatus"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCAInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a private certificate authority (CA) instance, including the remaining certificate quota.
//
// Description:
//
// Use this operation to query the status of a private CA instance by its ID. The status information includes the instance\\"s status, the total number of certificates it can issue, and the number of certificates already issued.
//
// Before you call this operation, purchase a private CA in the [CAS console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Purchase a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## QPS limits
//
// This operation has a queries per second (QPS) limit of 10 for each user. If you exceed this limit, API calls are throttled, which can affect your business. Call this operation at a reasonable rate.
//
// @param request - GetCAInstanceStatusRequest
//
// @return GetCAInstanceStatusResponse
func (client *Client) GetCAInstanceStatus(request *GetCAInstanceStatusRequest) (_result *GetCAInstanceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCAInstanceStatusResponse{}
	_body, _err := client.GetCAInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the certificate revocation list (CRL) URL of a certificate authority (CA).
//
// Description:
//
// Queries the status information of a private CA instance that you purchased in the SSL Certificate console by using the ID of the private CA instance. For example, you can query the status of the CA instance, the number of digital certificates included, and the number of digital certificates issued.
//
// Before you invoke this operation, you must have purchased a private CA in the [Certificate Management Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Purchase a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, throttling is triggered. This may affect your business. Invoke this operation as needed.
//
// @param request - GetCaInstanceCrlAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCaInstanceCrlAddressResponse
func (client *Client) GetCaInstanceCrlAddressWithOptions(request *GetCaInstanceCrlAddressRequest, runtime *dara.RuntimeOptions) (_result *GetCaInstanceCrlAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaIdentifier) {
		query["CaIdentifier"] = request.CaIdentifier
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCaInstanceCrlAddress"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCaInstanceCrlAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the certificate revocation list (CRL) URL of a certificate authority (CA).
//
// Description:
//
// Queries the status information of a private CA instance that you purchased in the SSL Certificate console by using the ID of the private CA instance. For example, you can query the status of the CA instance, the number of digital certificates included, and the number of digital certificates issued.
//
// Before you invoke this operation, you must have purchased a private CA in the [Certificate Management Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Purchase a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, throttling is triggered. This may affect your business. Invoke this operation as needed.
//
// @param request - GetCaInstanceCrlAddressRequest
//
// @return GetCaInstanceCrlAddressResponse
func (client *Client) GetCaInstanceCrlAddress(request *GetCaInstanceCrlAddressRequest) (_result *GetCaInstanceCrlAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCaInstanceCrlAddressResponse{}
	_body, _err := client.GetCaInstanceCrlAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the end-entity instances under a private certificate authority (CA).
//
// Description:
//
// This operation queries the status of a private Certificate Authority (CA) instance by its ID. It returns details for a private CA instance that you purchased in the Certificate Management Service (CAS) console. These details include the instance status, the number of certificates it contains, and the number of issued certificates.
//
// Before calling this operation, purchase a private CA from the [CAS console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Purchase a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. If you exceed the limit, API calls are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - ListAllEndEntityInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllEndEntityInstanceResponse
func (client *Client) ListAllEndEntityInstanceWithOptions(request *ListAllEndEntityInstanceRequest, runtime *dara.RuntimeOptions) (_result *ListAllEndEntityInstanceResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.RecursiveChildren) {
		query["RecursiveChildren"] = request.RecursiveChildren
	}

	if !dara.IsNil(request.ShowSize) {
		query["ShowSize"] = request.ShowSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAllEndEntityInstance"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllEndEntityInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the end-entity instances under a private certificate authority (CA).
//
// Description:
//
// This operation queries the status of a private Certificate Authority (CA) instance by its ID. It returns details for a private CA instance that you purchased in the Certificate Management Service (CAS) console. These details include the instance status, the number of certificates it contains, and the number of issued certificates.
//
// Before calling this operation, purchase a private CA from the [CAS console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Purchase a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. If you exceed the limit, API calls are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - ListAllEndEntityInstanceRequest
//
// @return ListAllEndEntityInstanceResponse
func (client *Client) ListAllEndEntityInstance(request *ListAllEndEntityInstanceRequest) (_result *ListAllEndEntityInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAllEndEntityInstanceResponse{}
	_body, _err := client.ListAllEndEntityInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operation logs of a certificate authority (CA) certificate, such as issuance and revocation events.
//
// Description:
//
// You can use this API to query the operation logs for a Certificate Authority (CA) certificate. These logs record operations, such as certificate creation and status changes, for both root and subordinate CA certificates.
//
// This API is limited to 10 queries per second (QPS) per user. API calls that exceed this limit are throttled. This can impact your business. Ensure that you call the API within this limit.
//
// @param request - ListCACertificateLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCACertificateLogResponse
func (client *Client) ListCACertificateLogWithOptions(request *ListCACertificateLogRequest, runtime *dara.RuntimeOptions) (_result *ListCACertificateLogResponse, _err error) {
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
		Action:      dara.String("ListCACertificateLog"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCACertificateLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operation logs of a certificate authority (CA) certificate, such as issuance and revocation events.
//
// Description:
//
// You can use this API to query the operation logs for a Certificate Authority (CA) certificate. These logs record operations, such as certificate creation and status changes, for both root and subordinate CA certificates.
//
// This API is limited to 10 queries per second (QPS) per user. API calls that exceed this limit are throttled. This can impact your business. Ensure that you call the API within this limit.
//
// @param request - ListCACertificateLogRequest
//
// @return ListCACertificateLogResponse
func (client *Client) ListCACertificateLog(request *ListCACertificateLogRequest) (_result *ListCACertificateLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCACertificateLogResponse{}
	_body, _err := client.ListCACertificateLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of certificates.
//
// Description:
//
// The QPS limit for this API is 10 per user. If you exceed this limit, your API calls will be throttled. To avoid impacting your business, please plan your calls accordingly.
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
	if !dara.IsNil(request.AfterDate) {
		query["AfterDate"] = request.AfterDate
	}

	if !dara.IsNil(request.BeforeDate) {
		query["BeforeDate"] = request.BeforeDate
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.InstanceUuid) {
		query["InstanceUuid"] = request.InstanceUuid
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ParentIdentifier) {
		query["ParentIdentifier"] = request.ParentIdentifier
	}

	if !dara.IsNil(request.ShowSize) {
		query["ShowSize"] = request.ShowSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCert"),
		Version:     dara.String("2020-06-30"),
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
// Retrieves a list of certificates.
//
// Description:
//
// The QPS limit for this API is 10 per user. If you exceed this limit, your API calls will be throttled. To avoid impacting your business, please plan your calls accordingly.
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
// Lists all client certificates and server certificates issued by a private certificate authority (CA).
//
// Description:
//
// This API performs a paged query to retrieve the details of all client and server-side certificates that you have created. These details include the unique identifier, serial number, subject information, content, and status of each certificate.
//
// ## QPS limit
//
// The QPS limit for a single user is 10 calls per second. If you exceed this limit, your API calls are throttled, which may affect your business. Call this API at a reasonable rate.
//
// @param request - ListClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClientCertificateResponse
func (client *Client) ListClientCertificateWithOptions(request *ListClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *ListClientCertificateResponse, _err error) {
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

	if !dara.IsNil(request.Identifier) {
		query["Identifier"] = request.Identifier
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
		Action:      dara.String("ListClientCertificate"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all client certificates and server certificates issued by a private certificate authority (CA).
//
// Description:
//
// This API performs a paged query to retrieve the details of all client and server-side certificates that you have created. These details include the unique identifier, serial number, subject information, content, and status of each certificate.
//
// ## QPS limit
//
// The QPS limit for a single user is 10 calls per second. If you exceed this limit, your API calls are throttled, which may affect your business. Call this API at a reasonable rate.
//
// @param request - ListClientCertificateRequest
//
// @return ListClientCertificateResponse
func (client *Client) ListClientCertificate(request *ListClientCertificateRequest) (_result *ListClientCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClientCertificateResponse{}
	_body, _err := client.ListClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists private certificate authority (CA) certificates.
//
// Description:
//
// This operation lists CA certificates, including root and subordinate CA certificates.
//
// This operation has a limit of 10 queries per second (QPS) for each user. If you exceed the limit, API calls are throttled, which may affect your business. Call this operation at a reasonable rate.
//
// @param request - ListPcaCaCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPcaCaCertificateResponse
func (client *Client) ListPcaCaCertificateWithOptions(request *ListPcaCaCertificateRequest, runtime *dara.RuntimeOptions) (_result *ListPcaCaCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPcaCaCertificate"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPcaCaCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists private certificate authority (CA) certificates.
//
// Description:
//
// This operation lists CA certificates, including root and subordinate CA certificates.
//
// This operation has a limit of 10 queries per second (QPS) for each user. If you exceed the limit, API calls are throttled, which may affect your business. Call this operation at a reasonable rate.
//
// @param request - ListPcaCaCertificateRequest
//
// @return ListPcaCaCertificateResponse
func (client *Client) ListPcaCaCertificate(request *ListPcaCaCertificateRequest) (_result *ListPcaCaCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPcaCaCertificateResponse{}
	_body, _err := client.ListPcaCaCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all revoked client certificates and server certificates.
//
// Description:
//
// Performs a paged query to retrieve the details of all revoked client and server-side certificates, such as the unique identifier, serial number, and revocation date.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this API is 10 for each user. If you exceed this limit, API calls are throttled, which may affect your business. We recommend that you call this API at a reasonable rate.
//
// @param request - ListRevokeCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRevokeCertificateResponse
func (client *Client) ListRevokeCertificateWithOptions(request *ListRevokeCertificateRequest, runtime *dara.RuntimeOptions) (_result *ListRevokeCertificateResponse, _err error) {
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

	if !dara.IsNil(request.ShowSize) {
		query["ShowSize"] = request.ShowSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRevokeCertificate"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRevokeCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all revoked client certificates and server certificates.
//
// Description:
//
// Performs a paged query to retrieve the details of all revoked client and server-side certificates, such as the unique identifier, serial number, and revocation date.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this API is 10 for each user. If you exceed this limit, API calls are throttled, which may affect your business. We recommend that you call this API at a reasonable rate.
//
// @param request - ListRevokeCertificateRequest
//
// @return ListRevokeCertificateResponse
func (client *Client) ListRevokeCertificate(request *ListRevokeCertificateRequest) (_result *ListRevokeCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRevokeCertificateResponse{}
	_body, _err := client.ListRevokeCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are attached to Certificate Management Service resources.
//
// Description:
//
// Before you call this operation, you must purchase a private CA in the [Certificate Management Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Purchase a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// You can call this operation up to 10 times per second per Alibaba Cloud account. If the number of calls per second exceeds this limit, throttling is triggered. This may affect your business. We recommend that you plan your calls accordingly.
//
// ## QPS limit
//
// You can call this operation up to 10 times per second for each Alibaba Cloud account. If the number of calls per second exceeds this limit, throttling is triggered, which may affect your business. We recommend that you plan your calls accordingly.
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
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2020-06-30"),
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
// Queries the tags that are attached to Certificate Management Service resources.
//
// Description:
//
// Before you call this operation, you must purchase a private CA in the [Certificate Management Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Purchase a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// You can call this operation up to 10 times per second per Alibaba Cloud account. If the number of calls per second exceeds this limit, throttling is triggered. This may affect your business. We recommend that you plan your calls accordingly.
//
// ## QPS limit
//
// You can call this operation up to 10 times per second for each Alibaba Cloud account. If the number of calls per second exceeds this limit, throttling is triggered, which may affect your business. We recommend that you plan your calls accordingly.
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
// Moves a Certificate Management Service resource to a different resource group.
//
// Description:
//
// This operation queries the status information of a private CA instance that you purchased in the Certificate Management Service (CAS) console. You can query by the private CA instance ID to retrieve information such as the status of the CA instance, the number of certificates it contains, and the number of issued certificates.
//
// Before you call this operation, you must purchase a private CA in the [CAS console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Purchase a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## QPS limit
//
// This operation is limited to 10 queries per second (QPS) for each user. If you exceed this limit, API calls are throttled, which can affect your business. Call this operation within the specified limit.
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
		Version:     dara.String("2020-06-30"),
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
// Moves a Certificate Management Service resource to a different resource group.
//
// Description:
//
// This operation queries the status information of a private CA instance that you purchased in the Certificate Management Service (CAS) console. You can query by the private CA instance ID to retrieve information such as the status of the CA instance, the number of certificates it contains, and the number of issued certificates.
//
// Before you call this operation, you must purchase a private CA in the [CAS console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Purchase a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## QPS limit
//
// This operation is limited to 10 queries per second (QPS) for each user. If you exceed this limit, API calls are throttled, which can affect your business. Call this operation within the specified limit.
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
// Attaches tags to one or more Certificate Management Service resources.
//
// Description:
//
// This operation adds tags to one or more resources. You can add tags to private CA instances that you purchased in the Certificate Management Service (CAS) console.
//
// Before calling this operation, purchase a private CA in the [CAS console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Purchase a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## QPS limit
//
// This operation is limited to 10 queries per second (QPS) per user. If you exceed this limit, API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
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
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2020-06-30"),
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
// Attaches tags to one or more Certificate Management Service resources.
//
// Description:
//
// This operation adds tags to one or more resources. You can add tags to private CA instances that you purchased in the Certificate Management Service (CAS) console.
//
// Before calling this operation, purchase a private CA in the [CAS console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Purchase a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## QPS limit
//
// This operation is limited to 10 queries per second (QPS) per user. If you exceed this limit, API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
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
// Removes tags from one or more Certificate Management Service resources.
//
// Description:
//
// This operation queries status information for a private Certificate Authority (CA) instance that you purchased in the Certificate Management Service (CAS) console. You can use the private CA instance ID to retrieve information such as the instance status, the number of certificates it contains, and the number of certificates issued.
//
// Before you call this operation, you must purchase a private CA from the [CAS console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Purchase a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second per user. If you exceed this limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2020-06-30"),
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
// Removes tags from one or more Certificate Management Service resources.
//
// Description:
//
// This operation queries status information for a private Certificate Authority (CA) instance that you purchased in the Certificate Management Service (CAS) console. You can use the private CA instance ID to retrieve information such as the instance status, the number of certificates it contains, and the number of certificates issued.
//
// Before you call this operation, you must purchase a private CA from the [CAS console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Purchase a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second per user. If you exceed this limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
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
// Changes the status of a root or subordinate certificate authority (CA) certificate, such as revoking or disabling an active CA.
//
// Description:
//
// When you create a CA certificate, its status is ISSUE by default. You can call this API operation to change the status of a CA certificate from ISSUE to REVOKE. A CA certificate in the ISSUE state can be used to issue certificates. A CA certificate in the REVOKE state cannot be used to issue certificates, and all certificates issued by this CA certificate become invalid.
//
// Before you call this API operation, create a root CA certificate by calling [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) and an intermediate CA certificate by calling [CreateSubCACertificate](https://help.aliyun.com/document_detail/465959.html).
//
// ## QPS limit
//
// This operation is limited to 10 queries per second (QPS) for each user. If you exceed this limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - UpdateCACertificateStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCACertificateStatusResponse
func (client *Client) UpdateCACertificateStatusWithOptions(request *UpdateCACertificateStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateCACertificateStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Identifier) {
		query["Identifier"] = request.Identifier
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCACertificateStatus"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCACertificateStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the status of a root or subordinate certificate authority (CA) certificate, such as revoking or disabling an active CA.
//
// Description:
//
// When you create a CA certificate, its status is ISSUE by default. You can call this API operation to change the status of a CA certificate from ISSUE to REVOKE. A CA certificate in the ISSUE state can be used to issue certificates. A CA certificate in the REVOKE state cannot be used to issue certificates, and all certificates issued by this CA certificate become invalid.
//
// Before you call this API operation, create a root CA certificate by calling [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) and an intermediate CA certificate by calling [CreateSubCACertificate](https://help.aliyun.com/document_detail/465959.html).
//
// ## QPS limit
//
// This operation is limited to 10 queries per second (QPS) for each user. If you exceed this limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - UpdateCACertificateStatusRequest
//
// @return UpdateCACertificateStatusResponse
func (client *Client) UpdateCACertificateStatus(request *UpdateCACertificateStatusRequest) (_result *UpdateCACertificateStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCACertificateStatusResponse{}
	_body, _err := client.UpdateCACertificateStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the properties of a private CA certificate, such as the export status.
//
// Description:
//
// When a Certificate Authority (CA) certificate is created, its status is Normal by default. You can call this API operation to change the status of a CA certificate to Revoked. A CA certificate in the Normal status can be used to issue certificates. A revoked CA certificate cannot be used to issue certificates, and all certificates previously issued by it become invalid.
//
// Before you call this API operation, you must create a root CA certificate by calling [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) and a subordinate CA certificate by calling [CreateSubCACertificate](https://help.aliyun.com/document_detail/465959.html).
//
// ## QPS limits
//
// This API operation is limited to 10 queries per second (QPS) per user. If you exceed this limit, API calls are throttled, which may affect your business. Call this API operation at a reasonable rate.
//
// @param request - UpdatePcaCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePcaCertificateResponse
func (client *Client) UpdatePcaCertificateWithOptions(request *UpdatePcaCertificateRequest, runtime *dara.RuntimeOptions) (_result *UpdatePcaCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliasName) {
		query["AliasName"] = request.AliasName
	}

	if !dara.IsNil(request.Identifier) {
		query["Identifier"] = request.Identifier
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePcaCertificate"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePcaCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the properties of a private CA certificate, such as the export status.
//
// Description:
//
// When a Certificate Authority (CA) certificate is created, its status is Normal by default. You can call this API operation to change the status of a CA certificate to Revoked. A CA certificate in the Normal status can be used to issue certificates. A revoked CA certificate cannot be used to issue certificates, and all certificates previously issued by it become invalid.
//
// Before you call this API operation, you must create a root CA certificate by calling [CreateRootCACertificate](https://help.aliyun.com/document_detail/465962.html) and a subordinate CA certificate by calling [CreateSubCACertificate](https://help.aliyun.com/document_detail/465959.html).
//
// ## QPS limits
//
// This API operation is limited to 10 queries per second (QPS) per user. If you exceed this limit, API calls are throttled, which may affect your business. Call this API operation at a reasonable rate.
//
// @param request - UpdatePcaCertificateRequest
//
// @return UpdatePcaCertificateResponse
func (client *Client) UpdatePcaCertificate(request *UpdatePcaCertificateRequest) (_result *UpdatePcaCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePcaCertificateResponse{}
	_body, _err := client.UpdatePcaCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads a private CA certificate to Certificate Management Service for centralized management.
//
// Description:
//
// This API operation uploads a PCA certificate to a certificate repository.
//
// ## QPS limit
//
// This operation has a queries per second (QPS) limit of 10 calls per second for each user. If you exceed the limit, your API calls are throttled. Throttling may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - UploadPcaCertToCasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadPcaCertToCasResponse
func (client *Client) UploadPcaCertToCasWithOptions(request *UploadPcaCertToCasRequest, runtime *dara.RuntimeOptions) (_result *UploadPcaCertToCasResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadPcaCertToCas"),
		Version:     dara.String("2020-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadPcaCertToCasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a private CA certificate to Certificate Management Service for centralized management.
//
// Description:
//
// This API operation uploads a PCA certificate to a certificate repository.
//
// ## QPS limit
//
// This operation has a queries per second (QPS) limit of 10 calls per second for each user. If you exceed the limit, your API calls are throttled. Throttling may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - UploadPcaCertToCasRequest
//
// @return UploadPcaCertToCasResponse
func (client *Client) UploadPcaCertToCas(request *UploadPcaCertToCasRequest) (_result *UploadPcaCertToCasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadPcaCertToCasResponse{}
	_body, _err := client.UploadPcaCertToCasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
