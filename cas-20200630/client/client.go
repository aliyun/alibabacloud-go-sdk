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
// Issues a client certificate by using a system-generated certificate signing request (CSR) file.
//
// Description:
//
// Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](~~CreateRootCACertificate~~) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](~~CreateRootCACertificate~~) operation. Only intermediate CA certificates can issue client certificates.
//
// ## QPS limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClientCertificateResponse
func (client *Client) CreateClientCertificateWithOptions(request *CreateClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *CreateClientCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Issues a client certificate by using a system-generated certificate signing request (CSR) file.
//
// Description:
//
// Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](~~CreateRootCACertificate~~) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](~~CreateRootCACertificate~~) operation. Only intermediate CA certificates can issue client certificates.
//
// ## QPS limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Issues a client certificate by using a custom certificate signing request (CSR) file.
//
// Description:
//
// Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](https://help.aliyun.com/document_detail/328093.html) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](https://help.aliyun.com/document_detail/328094.html) operation. Only intermediate CA certificates can be used to issue client certificates.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateClientCertificateWithCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClientCertificateWithCsrResponse
func (client *Client) CreateClientCertificateWithCsrWithOptions(request *CreateClientCertificateWithCsrRequest, runtime *dara.RuntimeOptions) (_result *CreateClientCertificateWithCsrResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Issues a client certificate by using a custom certificate signing request (CSR) file.
//
// Description:
//
// Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](https://help.aliyun.com/document_detail/328093.html) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](https://help.aliyun.com/document_detail/328094.html) operation. Only intermediate CA certificates can be used to issue client certificates.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Issues a certificate based on the specified key usage, extended key usage, and name and alias of the entity that uses the certificate.
//
// Description:
//
// By default, the name of the entity is obtained from the certificate signing request (CSR) of the certificate that you want to issue. If you specify a different name for the entity, the name of the entity in the CSR becomes invalid. The specified name is used to issue the certificate.
//
// You must specify the key usage and extended key usage based on the certificate type. The following list describes common certificate types:
//
//   - Server certificate
//
// Key usage: digitalSignature or keyEncipherment
//
// Extended key usage: serverAuth
//
//   - Client certificate
//
// Key usage: digitalSignature or keyEncipherment
//
// Extended key usage: clientAuth
//
//   - Mutual Transport Layer Security (TLS) authentication certificate
//
// Key usage: digitalSignature or keyEncipherment
//
// Extended key usage: serverAuth or clientAuth
//
//   - Email certificate
//
// Key usage: digitalSignature or contentCommitment
//
// Extended key usage: emailProtection
//
// Note: Compliant certificate authorities (CAs) are managed by third-party authorities. This operation is not supported for compliant CAs.
//
// @param request - CreateCustomCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomCertificateResponse
func (client *Client) CreateCustomCertificateWithOptions(request *CreateCustomCertificateRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

	if !dara.IsNil(request.Validity) {
		query["Validity"] = request.Validity
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
// Issues a certificate based on the specified key usage, extended key usage, and name and alias of the entity that uses the certificate.
//
// Description:
//
// By default, the name of the entity is obtained from the certificate signing request (CSR) of the certificate that you want to issue. If you specify a different name for the entity, the name of the entity in the CSR becomes invalid. The specified name is used to issue the certificate.
//
// You must specify the key usage and extended key usage based on the certificate type. The following list describes common certificate types:
//
//   - Server certificate
//
// Key usage: digitalSignature or keyEncipherment
//
// Extended key usage: serverAuth
//
//   - Client certificate
//
// Key usage: digitalSignature or keyEncipherment
//
// Extended key usage: clientAuth
//
//   - Mutual Transport Layer Security (TLS) authentication certificate
//
// Key usage: digitalSignature or keyEncipherment
//
// Extended key usage: serverAuth or clientAuth
//
//   - Email certificate
//
// Key usage: digitalSignature or contentCommitment
//
// Extended key usage: emailProtection
//
// Note: Compliant certificate authorities (CAs) are managed by third-party authorities. This operation is not supported for compliant CAs.
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
// 创建外部子CA证书
//
// @param tmpReq - CreateExternalCACertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExternalCACertificateResponse
func (client *Client) CreateExternalCACertificateWithOptions(tmpReq *CreateExternalCACertificateRequest, runtime *dara.RuntimeOptions) (_result *CreateExternalCACertificateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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

	if !dara.IsNil(request.Csr) {
		query["Csr"] = request.Csr
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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
// 创建外部子CA证书
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
// Revokes a client certificate or a server certificate.
//
// Description:
//
// After a client certificate or a server certificate is revoked, the client or the server on which the certificate is installed cannot establish HTTPS connections with other devices.
//
// After a client certificate or a server certificate is revoked, you can call the [DeleteClientCertificate](https://help.aliyun.com/document_detail/330880.html) operation to permanently delete the certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateRevokeClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRevokeClientCertificateResponse
func (client *Client) CreateRevokeClientCertificateWithOptions(request *CreateRevokeClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *CreateRevokeClientCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Revokes a client certificate or a server certificate.
//
// Description:
//
// After a client certificate or a server certificate is revoked, the client or the server on which the certificate is installed cannot establish HTTPS connections with other devices.
//
// After a client certificate or a server certificate is revoked, you can call the [DeleteClientCertificate](https://help.aliyun.com/document_detail/330880.html) operation to permanently delete the certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// You can call the CreateRootCACertificate operation to create a self-signed root CA certificate. A root CA certificate is the trust anchor in a chain of trust for private certificates that are used within an enterprise. You must create a root CA certificate before you can use the root CA certificate to issue intermediate CA certificates. Then, you can use the intermediate CA certificates to issue client certificates and server certificates.
//
// Before you call this operation, make sure that you have purchased a private root CA instance by using the [Certificate Management Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Create a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateRootCACertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRootCACertificateResponse
func (client *Client) CreateRootCACertificateWithOptions(request *CreateRootCACertificateRequest, runtime *dara.RuntimeOptions) (_result *CreateRootCACertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// You can call the CreateRootCACertificate operation to create a self-signed root CA certificate. A root CA certificate is the trust anchor in a chain of trust for private certificates that are used within an enterprise. You must create a root CA certificate before you can use the root CA certificate to issue intermediate CA certificates. Then, you can use the intermediate CA certificates to issue client certificates and server certificates.
//
// Before you call this operation, make sure that you have purchased a private root CA instance by using the [Certificate Management Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Create a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Issues a server certificate by using a system-generated certificate signing request (CSR) file.
//
// Description:
//
// Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](https://help.aliyun.com/document_detail/328093.html) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](https://help.aliyun.com/document_detail/328094.html) operation. Only intermediate CA certificates can be used to issue server certificates.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateServerCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServerCertificateResponse
func (client *Client) CreateServerCertificateWithOptions(request *CreateServerCertificateRequest, runtime *dara.RuntimeOptions) (_result *CreateServerCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Issues a server certificate by using a system-generated certificate signing request (CSR) file.
//
// Description:
//
// Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](https://help.aliyun.com/document_detail/328093.html) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](https://help.aliyun.com/document_detail/328094.html) operation. Only intermediate CA certificates can be used to issue server certificates.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Issues a server certificate by using a custom certificate signing request (CSR) file.
//
// Description:
//
// ## Usage notes
//
// Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](https://help.aliyun.com/document_detail/328093.html) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](https://help.aliyun.com/document_detail/328094.html) operation. Only intermediate CA certificates can be used to issue server certificates.
//
// @param request - CreateServerCertificateWithCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServerCertificateWithCsrResponse
func (client *Client) CreateServerCertificateWithCsrWithOptions(request *CreateServerCertificateWithCsrRequest, runtime *dara.RuntimeOptions) (_result *CreateServerCertificateWithCsrResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Issues a server certificate by using a custom certificate signing request (CSR) file.
//
// Description:
//
// ## Usage notes
//
// Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](https://help.aliyun.com/document_detail/328093.html) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](https://help.aliyun.com/document_detail/328094.html) operation. Only intermediate CA certificates can be used to issue server certificates.
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
// Creates an intermediate certificate authority (CA) certificate.
//
// Description:
//
// You can call this operation to issue an intermediate certificate authority (CA) certificate by using an existing root CA certificate. Intermediate CA certificates can be used to issue client certificates and server certificates.
//
// Before you call this operation, make sure that you have issued a root CA certificate by calling the [CreateRootCACertificate] operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateSubCACertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSubCACertificateResponse
func (client *Client) CreateSubCACertificateWithOptions(request *CreateSubCACertificateRequest, runtime *dara.RuntimeOptions) (_result *CreateSubCACertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
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
// Creates an intermediate certificate authority (CA) certificate.
//
// Description:
//
// You can call this operation to issue an intermediate certificate authority (CA) certificate by using an existing root CA certificate. Intermediate CA certificates can be used to issue client certificates and server certificates.
//
// Before you call this operation, make sure that you have issued a root CA certificate by calling the [CreateRootCACertificate] operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Deletes a client certificate or a server certificate that is revoked.
//
// Description:
//
// Before you call this operation, you must call the [CreateRevokeClientCertificate](https://help.aliyun.com/document_detail/330876.html) operation to revoke a client certificate or a server certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClientCertificateResponse
func (client *Client) DeleteClientCertificateWithOptions(request *DeleteClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteClientCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Deletes a client certificate or a server certificate that is revoked.
//
// Description:
//
// Before you call this operation, you must call the [CreateRevokeClientCertificate](https://help.aliyun.com/document_detail/330876.html) operation to revoke a client certificate or a server certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Queries the details about a root certificate authority (CA) certificate or an intermediate CA certificate.
//
// Description:
//
// You can call the DescribeCACertificate operation to query the details about a root CA certificate or an intermediate CA certificate by using the unique identifier of the root CA certificate or intermediate CA certificate. The details include the serial number, user information, and content of a CA certificate.
//
// Before you call this operation, make sure that you have created a root CA by calling the [CreateRootCACertificate] operation or an intermediate CA certificate by calling the [CreateSubCACertificate] operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeCACertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCACertificateResponse
func (client *Client) DescribeCACertificateWithOptions(request *DescribeCACertificateRequest, runtime *dara.RuntimeOptions) (_result *DescribeCACertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the details about a root certificate authority (CA) certificate or an intermediate CA certificate.
//
// Description:
//
// You can call the DescribeCACertificate operation to query the details about a root CA certificate or an intermediate CA certificate by using the unique identifier of the root CA certificate or intermediate CA certificate. The details include the serial number, user information, and content of a CA certificate.
//
// Before you call this operation, make sure that you have created a root CA by calling the [CreateRootCACertificate] operation or an intermediate CA certificate by calling the [CreateSubCACertificate] operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Queries the number of certificate authority (CA) certificates that you create.
//
// Description:
//
// You can call the DescribeCACertificateCount operation to query the number of created CA certificates, which includes root CA certificates and intermediate CA certificates.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeCACertificateCountRequest
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
// Queries the number of certificate authority (CA) certificates that you create.
//
// Description:
//
// You can call the DescribeCACertificateCount operation to query the number of created CA certificates, which includes root CA certificates and intermediate CA certificates.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Queries the details about all root certificate authority (CA) certificates and intermediate CA certificates.
//
// Description:
//
// You can call the DescribeCACertificateList operation to perform a paged query of the details about all CA certificates that you create. The details include the unique identifier, serial number, user information, and content of each root CA certificate or intermediate CA certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeCACertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCACertificateListResponse
func (client *Client) DescribeCACertificateListWithOptions(request *DescribeCACertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCACertificateListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the details about all root certificate authority (CA) certificates and intermediate CA certificates.
//
// Description:
//
// You can call the DescribeCACertificateList operation to perform a paged query of the details about all CA certificates that you create. The details include the unique identifier, serial number, user information, and content of each root CA certificate or intermediate CA certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Queries the encrypted private key of a client certificate or a server certificate.
//
// Description:
//
// ## Usage notes
//
// You can call the DescribeCertificatePrivateKey operation to obtain the encrypted private key of a client certificate or a server certificate. The certificate is issued based on a system-generated certificate signing request (CSR). Before you call this operation, make sure that you have issued a client certificate or a server certificate by calling the following operation:
//
//   - [CreateClientCertificate](https://help.aliyun.com/document_detail/330873.html)
//
//   - [CreateServerCertificate](https://help.aliyun.com/document_detail/330877.html)
//
// To ensure the security of private key transmission, the DescribeCertificatePrivateKey operation encrypts the private key by using the private key password that you specify and returns the encrypted private key. The private key password is a string that is used to encrypt the private key. After you obtain the encrypted private key of the certificate, you can use the following methods to decrypt the private key:
//
//   - If the encryption algorithm of the certificate is RSA, you must run the `openssl rsa -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [OpenSSL](https://www.openssl.org/source/) or [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
//
//   - If the encryption algorithm of the certificate is ECC, you must run the `openssl ec -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [OpenSSL](https://www.openssl.org/source/) or [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
//
//   - If the encryption algorithm of the certificate is SM2, you must run the `openssl ec -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
//
// >  You can call the [DescribeClientCertificate] operation to query the encryption algorithm type of a client certificate or a server certificate.
//
// ## Limits
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeCertificatePrivateKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCertificatePrivateKeyResponse
func (client *Client) DescribeCertificatePrivateKeyWithOptions(request *DescribeCertificatePrivateKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeCertificatePrivateKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncryptedCode) {
		query["EncryptedCode"] = request.EncryptedCode
	}

	if !dara.IsNil(request.Identifier) {
		query["Identifier"] = request.Identifier
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
// Queries the encrypted private key of a client certificate or a server certificate.
//
// Description:
//
// ## Usage notes
//
// You can call the DescribeCertificatePrivateKey operation to obtain the encrypted private key of a client certificate or a server certificate. The certificate is issued based on a system-generated certificate signing request (CSR). Before you call this operation, make sure that you have issued a client certificate or a server certificate by calling the following operation:
//
//   - [CreateClientCertificate](https://help.aliyun.com/document_detail/330873.html)
//
//   - [CreateServerCertificate](https://help.aliyun.com/document_detail/330877.html)
//
// To ensure the security of private key transmission, the DescribeCertificatePrivateKey operation encrypts the private key by using the private key password that you specify and returns the encrypted private key. The private key password is a string that is used to encrypt the private key. After you obtain the encrypted private key of the certificate, you can use the following methods to decrypt the private key:
//
//   - If the encryption algorithm of the certificate is RSA, you must run the `openssl rsa -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [OpenSSL](https://www.openssl.org/source/) or [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
//
//   - If the encryption algorithm of the certificate is ECC, you must run the `openssl ec -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [OpenSSL](https://www.openssl.org/source/) or [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
//
//   - If the encryption algorithm of the certificate is SM2, you must run the `openssl ec -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
//
// >  You can call the [DescribeClientCertificate] operation to query the encryption algorithm type of a client certificate or a server certificate.
//
// ## Limits
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Queries the details about a client certificate or a server certificate by using the unique identifier of the certificate.
//
// Description:
//
// You can call the DescribeClientCertificate operation to query the details about a client certificate or a server certificate by using the unique identifier of the certificate. The details include the serial number, user information, content, and status of each certificate.
//
// Before you call this operation, make sure that you have created a client certificate or a server certificate.
//
// For more information about how to call an operation to create a client certificate, see the following topics:
//
//   - [CreateClientCertificate](https://help.aliyun.com/document_detail/330873.html)
//
//   - [CreateClientCertificateWithCsr](https://help.aliyun.com/document_detail/330875.html)
//
// For more information about how to call an operation to create a server certificate, see the following topics:
//
//   - [CreateServerCertificate](https://help.aliyun.com/document_detail/330877.html)
//
//   - [CreateServerCertificateWithCsr](https://help.aliyun.com/document_detail/330878.html)
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClientCertificateResponse
func (client *Client) DescribeClientCertificateWithOptions(request *DescribeClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *DescribeClientCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the details about a client certificate or a server certificate by using the unique identifier of the certificate.
//
// Description:
//
// You can call the DescribeClientCertificate operation to query the details about a client certificate or a server certificate by using the unique identifier of the certificate. The details include the serial number, user information, content, and status of each certificate.
//
// Before you call this operation, make sure that you have created a client certificate or a server certificate.
//
// For more information about how to call an operation to create a client certificate, see the following topics:
//
//   - [CreateClientCertificate](https://help.aliyun.com/document_detail/330873.html)
//
//   - [CreateClientCertificateWithCsr](https://help.aliyun.com/document_detail/330875.html)
//
// For more information about how to call an operation to create a server certificate, see the following topics:
//
//   - [CreateServerCertificate](https://help.aliyun.com/document_detail/330877.html)
//
//   - [CreateServerCertificateWithCsr](https://help.aliyun.com/document_detail/330878.html)
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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

// Summary:
//
// Queries the status information about client certificates and server certificates by using the unique identifiers of the certificates.
//
// Description:
//
// You can call the DescribeClientCertificateStatus operation to query the status information about multiple client certificates or server certificates at a time by using the unique identifiers of the certificates. For example, you can check whether a certificate is revoked.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeClientCertificateStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClientCertificateStatusResponse
func (client *Client) DescribeClientCertificateStatusWithOptions(request *DescribeClientCertificateStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeClientCertificateStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the status information about client certificates and server certificates by using the unique identifiers of the certificates.
//
// Description:
//
// You can call the DescribeClientCertificateStatus operation to query the status information about multiple client certificates or server certificates at a time by using the unique identifiers of the certificates. For example, you can check whether a certificate is revoked.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Queries the status information about a private root certificate authority (CA) instance or a private intermediate CA instance that you purchase by using the Certificate Management Service console.
//
// Description:
//
// ## Usage notes
//
// You can call the GetCAInstanceStatus operation to query the status information of a private CA instance by using the ID of the instance. The instance is purchased by using the SSL Certificates Service console. The status information includes the status of the private CA instance, the number of certificates that can be issued by using the private CA instance, and the number of issued certificates.
//
// Before you call this operation, make sure that you have purchased a private CA by using the [SSL Certificates Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Create a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// @param request - GetCAInstanceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCAInstanceStatusResponse
func (client *Client) GetCAInstanceStatusWithOptions(request *GetCAInstanceStatusRequest, runtime *dara.RuntimeOptions) (_result *GetCAInstanceStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the status information about a private root certificate authority (CA) instance or a private intermediate CA instance that you purchase by using the Certificate Management Service console.
//
// Description:
//
// ## Usage notes
//
// You can call the GetCAInstanceStatus operation to query the status information of a private CA instance by using the ID of the instance. The instance is purchased by using the SSL Certificates Service console. The status information includes the status of the private CA instance, the number of certificates that can be issued by using the private CA instance, and the number of issued certificates.
//
// Before you call this operation, make sure that you have purchased a private CA by using the [SSL Certificates Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Create a private CA](https://help.aliyun.com/document_detail/208553.html).
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
// 获取证书列表
//
// @param request - ListCertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCertResponse
func (client *Client) ListCertWithOptions(request *ListCertRequest, runtime *dara.RuntimeOptions) (_result *ListCertResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取证书列表
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
// Queries the details about all client certificates and server certificates.
//
// Description:
//
// You can call the ListClientCertificate operation to perform a paged query of the details about all client certificates and server certificates that you create. The details include the unique identifier, serial number, user information, content, and status of each certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClientCertificateResponse
func (client *Client) ListClientCertificateWithOptions(request *ListClientCertificateRequest, runtime *dara.RuntimeOptions) (_result *ListClientCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Identifier) {
		query["Identifier"] = request.Identifier
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
// Queries the details about all client certificates and server certificates.
//
// Description:
//
// You can call the ListClientCertificate operation to perform a paged query of the details about all client certificates and server certificates that you create. The details include the unique identifier, serial number, user information, content, and status of each certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// 查询私有CA机构证书
//
// @param request - ListPcaCaCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPcaCaCertificateResponse
func (client *Client) ListPcaCaCertificateWithOptions(request *ListPcaCaCertificateRequest, runtime *dara.RuntimeOptions) (_result *ListPcaCaCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询私有CA机构证书
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
// Queries the details about all client certificates and server certificates that are revoked.
//
// Description:
//
// You can call the ListRevokeCertificate operation to perform a paged query of the details about all revoked client certificates and server certificates. The details include the unique identifier, serial number, and revocation date of each certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListRevokeCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRevokeCertificateResponse
func (client *Client) ListRevokeCertificateWithOptions(request *ListRevokeCertificateRequest, runtime *dara.RuntimeOptions) (_result *ListRevokeCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the details about all client certificates and server certificates that are revoked.
//
// Description:
//
// You can call the ListRevokeCertificate operation to perform a paged query of the details about all revoked client certificates and server certificates. The details include the unique identifier, serial number, and revocation date of each certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Changes the status of a root certificate authority (CA) certificate or an intermediate CA certificate from ISSUE to REVOKE.
//
// Description:
//
// After a CA certificate is created, the CA certificate is in the ISSUE state by default. You can call the UpdateCACertificateStatus operation to change the status of a CA certificate from ISSUE to REVOKE. If a CA certificate is in the ISSUE state, the CA certificate can be used to issue certificates. If a CA certificate is in the REVOKE state, the CA certificate cannot be used to issue certificates, and the certificates that are issued from the CA certificate become invalid.
//
// Before you call this operation, make sure that you have created a root CA by calling the [CreateRootCACertificate] operation or an intermediate CA certificate by calling the [CreateSubCACertificate] operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UpdateCACertificateStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCACertificateStatusResponse
func (client *Client) UpdateCACertificateStatusWithOptions(request *UpdateCACertificateStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateCACertificateStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Changes the status of a root certificate authority (CA) certificate or an intermediate CA certificate from ISSUE to REVOKE.
//
// Description:
//
// After a CA certificate is created, the CA certificate is in the ISSUE state by default. You can call the UpdateCACertificateStatus operation to change the status of a CA certificate from ISSUE to REVOKE. If a CA certificate is in the ISSUE state, the CA certificate can be used to issue certificates. If a CA certificate is in the REVOKE state, the CA certificate cannot be used to issue certificates, and the certificates that are issued from the CA certificate become invalid.
//
// Before you call this operation, make sure that you have created a root CA by calling the [CreateRootCACertificate] operation or an intermediate CA certificate by calling the [CreateSubCACertificate] operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// 上传pca证书到SSL上传证书
//
// @param request - UploadPcaCertToCasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadPcaCertToCasResponse
func (client *Client) UploadPcaCertToCasWithOptions(request *UploadPcaCertToCasRequest, runtime *dara.RuntimeOptions) (_result *UploadPcaCertToCasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 上传pca证书到SSL上传证书
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
