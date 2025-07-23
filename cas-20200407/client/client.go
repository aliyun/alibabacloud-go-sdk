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
// Revokes an issued certificate and cancels the application order of the certificate.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CancelCertificateForPackageRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelCertificateForPackageRequestResponse
func (client *Client) CancelCertificateForPackageRequestWithOptions(request *CancelCertificateForPackageRequestRequest, runtime *dara.RuntimeOptions) (_result *CancelCertificateForPackageRequestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Revokes an issued certificate and cancels the application order of the certificate.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Cancels a certificate application order that is in the pending validation or being reviewed state.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CancelOrderRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelOrderRequestResponse
func (client *Client) CancelOrderRequestWithOptions(request *CancelOrderRequestRequest, runtime *dara.RuntimeOptions) (_result *CancelOrderRequestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Cancels a certificate application order that is in the pending validation or being reviewed state.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Submits a certificate application.
//
// Description:
//
//	  Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455800.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
//		- After you call this operation to submit a certificate application and the certificate is issued, the certificate quota provided by the resource plan that you purchased is consumed. When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate that you want to apply for.
//
//		- After you call this operation to submit a certificate application, you also need to call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required for domain name ownership verification and manually complete the verification. Then, your certificate application is reviewed by the certificate authority (CA). If you use the Domain Name System (DNS) verification method, you must complete the verification on your DNS service provider system. If you use the file verification method, you must complete the verification on the DNS server.
//
// @param request - CreateCertificateForPackageRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCertificateForPackageRequestResponse
func (client *Client) CreateCertificateForPackageRequestWithOptions(request *CreateCertificateForPackageRequestRequest, runtime *dara.RuntimeOptions) (_result *CreateCertificateForPackageRequestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Submits a certificate application.
//
// Description:
//
//	  Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455800.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
//		- After you call this operation to submit a certificate application and the certificate is issued, the certificate quota provided by the resource plan that you purchased is consumed. When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate that you want to apply for.
//
//		- After you call this operation to submit a certificate application, you also need to call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required for domain name ownership verification and manually complete the verification. Then, your certificate application is reviewed by the certificate authority (CA). If you use the Domain Name System (DNS) verification method, you must complete the verification on your DNS service provider system. If you use the file verification method, you must complete the verification on the DNS server.
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
//	  You can call this operation to apply for only DV certificates. If you want to apply for an organization validated (OV) or extended validation (EV) certificate, we recommend that you call the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/455296.html) operation. This operation allows you to apply for certificates of all specifications and specify the method to generate a certificate signing request (CSR) file.
//
//		- Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455803.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
//		- When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate.
//
//		- After you call this operation to submit a certificate application, Certificate Management Service automatically creates a CSR file for your application and consumes the certificate quota in the certificate resource plans of the specified specifications that you purchased. After you call this operation, you also need to call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required to complete domain name verification, and manually complete the verification. If you use the DNS verification method, you must complete the verification on the management platform of the domain name. If you use the file verification method, you must complete the verification on your DNS server. Then, the certificate authority (CA) will review your certificate application.
//
// @param request - CreateCertificateRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCertificateRequestResponse
func (client *Client) CreateCertificateRequestWithOptions(request *CreateCertificateRequestRequest, runtime *dara.RuntimeOptions) (_result *CreateCertificateRequestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
//	  You can call this operation to apply for only DV certificates. If you want to apply for an organization validated (OV) or extended validation (EV) certificate, we recommend that you call the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/455296.html) operation. This operation allows you to apply for certificates of all specifications and specify the method to generate a certificate signing request (CSR) file.
//
//		- Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455803.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
//		- When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate.
//
//		- After you call this operation to submit a certificate application, Certificate Management Service automatically creates a CSR file for your application and consumes the certificate quota in the certificate resource plans of the specified specifications that you purchased. After you call this operation, you also need to call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required to complete domain name verification, and manually complete the verification. If you use the DNS verification method, you must complete the verification on the management platform of the domain name. If you use the file verification method, you must complete the verification on your DNS server. Then, the certificate authority (CA) will review your certificate application.
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
// Purchases, applies for, and issues a domain validated (DV) certificate by using a custom certificate signing request (CSR) file. You can use extended certificate services to purchase and apply for a DV certificate with a few clicks.
//
// Description:
//
//	  You can use this operation to apply for only a domain validated (DV) certificate. You cannot use this operation to apply for an organization validated (OV) certificate. We recommend that you use the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/455296.html) operation to apply for a certificate. You can use the CreateCertificateForPackageRequest operation to apply for certificates of all types and specify the CSR generation method.
//
//		- Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455803.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
//		- When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate that you want to apply for.
//
//		- After you call this operation to submit a certificate application, the certificate quota of the required specifications that you purchased is consumed. After you call this operation, you must call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required for domain name ownership verification and manually complete the verification. Then, your certificate application is reviewed by the certificate authority (CA). If you use the Domain Name System (DNS) verification method, you must complete the verification on your DNS service provider system. If you use the file verification method, you must complete the verification on the DNS server.
//
// @param request - CreateCertificateWithCsrRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCertificateWithCsrRequestResponse
func (client *Client) CreateCertificateWithCsrRequestWithOptions(request *CreateCertificateWithCsrRequestRequest, runtime *dara.RuntimeOptions) (_result *CreateCertificateWithCsrRequestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Purchases, applies for, and issues a domain validated (DV) certificate by using a custom certificate signing request (CSR) file. You can use extended certificate services to purchase and apply for a DV certificate with a few clicks.
//
// Description:
//
//	  You can use this operation to apply for only a domain validated (DV) certificate. You cannot use this operation to apply for an organization validated (OV) certificate. We recommend that you use the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/455296.html) operation to apply for a certificate. You can use the CreateCertificateForPackageRequest operation to apply for certificates of all types and specify the CSR generation method.
//
//		- Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455803.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
//		- When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate that you want to apply for.
//
//		- After you call this operation to submit a certificate application, the certificate quota of the required specifications that you purchased is consumed. After you call this operation, you must call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required for domain name ownership verification and manually complete the verification. Then, your certificate application is reviewed by the certificate authority (CA). If you use the Domain Name System (DNS) verification method, you must complete the verification on your DNS service provider system. If you use the file verification method, you must complete the verification on the DNS server.
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
// Creates a certificate signing request (CSR). A CSR file contains the information about an SSL certificate that you want to apply for. The information includes the domain names that you want to bind to the certificate and the name and the geographical location of the certificate holder. When you submit a certificate application to a certificate authority (CA), you must provide a CSR. After the CA approves your certificate application, the CA uses the private key of the root CA to sign your CSR and generates a public key file. The public key file is the SSL certificate that the CA issues to you. The private key of the SSL certificate is generated when you create the CSR.
//
// @param request - CreateCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCsrResponse
func (client *Client) CreateCsrWithOptions(request *CreateCsrRequest, runtime *dara.RuntimeOptions) (_result *CreateCsrResponse, _err error) {
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
// Creates a certificate signing request (CSR). A CSR file contains the information about an SSL certificate that you want to apply for. The information includes the domain names that you want to bind to the certificate and the name and the geographical location of the certificate holder. When you submit a certificate application to a certificate authority (CA), you must provide a CSR. After the CA approves your certificate application, the CA uses the private key of the root CA to sign your CSR and generates a public key file. The public key file is the SSL certificate that the CA issues to you. The private key of the SSL certificate is generated when you create the CSR.
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
// Creates a certificate deployment task. After an SSL certificate is issued, you can create a certificate deployment task to immediately deploy the certificate to an Alibaba Cloud service or deploy the certificate to the service at a specific point in time. Then, the certificate can implement trusted identity authentication and ensure the security of data transmission for your website hosted on the service.
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Creates a certificate deployment task. After an SSL certificate is issued, you can create a certificate deployment task to immediately deploy the certificate to an Alibaba Cloud service or deploy the certificate to the service at a specific point in time. Then, the certificate can implement trusted identity authentication and ensure the security of data transmission for your website hosted on the service.
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
// Decrypts a certificate in a certificate repository.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DecryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DecryptResponse
func (client *Client) DecryptWithOptions(request *DecryptRequest, runtime *dara.RuntimeOptions) (_result *DecryptResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

	if !dara.IsNil(request.MessageType) {
		query["MessageType"] = request.MessageType
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
// Decrypts a certificate in a certificate repository.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Deletes an order in which the application for a domain validated (DV) certificate failed.
//
// Description:
//
// You can call this operation to delete a certificate application order only in the following scenarios:
//
//   - The status of the order is **review failed**. You have called the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to query the status of the certificate application order and the value of the **Type*	- parameter is **verify_fail**.
//
//   - The status of the order is **pending application**. You have called the [CancelOrderRequest](https://help.aliyun.com/document_detail/455299.html) operation to cancel a certificate application order whose status is pending review or being reviewed. The status of the certificate application order that is canceled in this case changes to **pending application**.
//
// @param request - DeleteCertificateRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCertificateRequestResponse
func (client *Client) DeleteCertificateRequestWithOptions(request *DeleteCertificateRequestRequest, runtime *dara.RuntimeOptions) (_result *DeleteCertificateRequestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Deletes an order in which the application for a domain validated (DV) certificate failed.
//
// Description:
//
// You can call this operation to delete a certificate application order only in the following scenarios:
//
//   - The status of the order is **review failed**. You have called the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to query the status of the certificate application order and the value of the **Type*	- parameter is **verify_fail**.
//
//   - The status of the order is **pending application**. You have called the [CancelOrderRequest](https://help.aliyun.com/document_detail/455299.html) operation to cancel a certificate application order whose status is pending review or being reviewed. The status of the certificate application order that is canceled in this case changes to **pending application**.
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
// Deletes a certificate signing request (CSR) file.
//
// @param request - DeleteCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCsrResponse
func (client *Client) DeleteCsrWithOptions(request *DeleteCsrRequest, runtime *dara.RuntimeOptions) (_result *DeleteCsrResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Deletes a certificate signing request (CSR) file.
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
// Deletes a deployment task.
//
// @param request - DeleteDeploymentJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeploymentJobResponse
func (client *Client) DeleteDeploymentJobWithOptions(request *DeleteDeploymentJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteDeploymentJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Deletes a deployment task.
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
// Deletes an expired or uploaded certificate.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteUserCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserCertificateResponse
func (client *Client) DeleteUserCertificateWithOptions(request *DeleteUserCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Deletes an expired or uploaded certificate.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Deletes the worker of a deployment task.
//
// @param request - DeleteWorkerResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkerResourceResponse
func (client *Client) DeleteWorkerResourceWithOptions(request *DeleteWorkerResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteWorkerResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Deletes the worker of a deployment task.
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
// Queries the status of a specified certificate application order.
//
// Description:
//
// If you do not complete the verification of the domain name ownership after you submit a certificate application, you can call this operation to obtain the information that is required to complete the verification. You can complete the verification of the domain name ownership based on the data returned. If you use the DNS verification method, you must complete the verification on the management platform of the domain name. If you use the file verification method, you must complete the verification on the DNS server.
//
// The certificate authority (CA) reviews your certificate application only after you complete the verification of the domain name ownership. After the CA approves your certificate application, the CA issues the certificate. If a certificate is issued, you can call this operation to obtain the CA certificate and private key of the certificate.
//
// @param request - DescribeCertificateStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCertificateStateResponse
func (client *Client) DescribeCertificateStateWithOptions(request *DescribeCertificateStateRequest, runtime *dara.RuntimeOptions) (_result *DescribeCertificateStateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the status of a specified certificate application order.
//
// Description:
//
// If you do not complete the verification of the domain name ownership after you submit a certificate application, you can call this operation to obtain the information that is required to complete the verification. You can complete the verification of the domain name ownership based on the data returned. If you use the DNS verification method, you must complete the verification on the management platform of the domain name. If you use the file verification method, you must complete the verification on the DNS server.
//
// The certificate authority (CA) reviews your certificate application only after you complete the verification of the domain name ownership. After the CA approves your certificate application, the CA issues the certificate. If a certificate is issued, you can call this operation to obtain the CA certificate and private key of the certificate.
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
// Queries the number of third-party cloud resources on which you deployed certificates by using a multi-cloud deployment task.
//
// @param request - DescribeCloudResourceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudResourceStatusResponse
func (client *Client) DescribeCloudResourceStatusWithOptions(request *DescribeCloudResourceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudResourceStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the number of third-party cloud resources on which you deployed certificates by using a multi-cloud deployment task.
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
// Queries the details of a deployment task. You can call the CreateDeploymentJob operation to create a deployment task and obtain the ID of the task.
//
// @param request - DescribeDeploymentJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeploymentJobResponse
func (client *Client) DescribeDeploymentJobWithOptions(request *DescribeDeploymentJobRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeploymentJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the details of a deployment task. You can call the CreateDeploymentJob operation to create a deployment task and obtain the ID of the task.
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
// Queries the number of worker tasks in a deployment task.
//
// @param request - DescribeDeploymentJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeploymentJobStatusResponse
func (client *Client) DescribeDeploymentJobStatusWithOptions(request *DescribeDeploymentJobStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeploymentJobStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the number of worker tasks in a deployment task.
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
// Queries the quota for domain validated (DV) certificates that you purchase and the quota usage.
//
// @param request - DescribePackageStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePackageStateResponse
func (client *Client) DescribePackageStateWithOptions(request *DescribePackageStateRequest, runtime *dara.RuntimeOptions) (_result *DescribePackageStateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the quota for domain validated (DV) certificates that you purchase and the quota usage.
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
// Encrypts a certificate in a certificate repository.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - EncryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EncryptResponse
func (client *Client) EncryptWithOptions(request *EncryptRequest, runtime *dara.RuntimeOptions) (_result *EncryptResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.MessageType) {
		query["MessageType"] = request.MessageType
	}

	if !dara.IsNil(request.Plaintext) {
		query["Plaintext"] = request.Plaintext
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
// Encrypts a certificate in a certificate repository.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Queries the quota for certificate repositories.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetCertWarehouseQuotaRequest
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
// Queries the quota for certificate repositories.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Obtains the content of a certificate signing request (CSR) file.
//
// @param request - GetCsrDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCsrDetailResponse
func (client *Client) GetCsrDetailWithOptions(request *GetCsrDetailRequest, runtime *dara.RuntimeOptions) (_result *GetCsrDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Obtains the content of a certificate signing request (CSR) file.
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
// Queries the details of a certificate.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetUserCertificateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserCertificateDetailResponse
func (client *Client) GetUserCertificateDetailWithOptions(request *GetUserCertificateDetailRequest, runtime *dara.RuntimeOptions) (_result *GetUserCertificateDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the details of a certificate.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Queries the certificates in a certificate repository.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
	if !dara.IsNil(request.CertType) {
		query["CertType"] = request.CertType
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
// Queries the certificates in a certificate repository.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Queries certificate repositories.
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries certificate repositories.
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
// Queries a list of AccessKey pairs for multi-cloud deployment.
//
// @param request - ListCloudAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudAccessResponse
func (client *Client) ListCloudAccessWithOptions(request *ListCloudAccessRequest, runtime *dara.RuntimeOptions) (_result *ListCloudAccessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries a list of AccessKey pairs for multi-cloud deployment.
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
// Queries the certificate resources of a cloud service provider and cloud services.
//
// @param tmpReq - ListCloudResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudResourcesResponse
func (client *Client) ListCloudResourcesWithOptions(tmpReq *ListCloudResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListCloudResourcesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the certificate resources of a cloud service provider and cloud services.
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
// Queries a list of contacts.
//
// @param request - ListContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContactResponse
func (client *Client) ListContactWithOptions(request *ListContactRequest, runtime *dara.RuntimeOptions) (_result *ListContactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries a list of contacts.
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
// Queries certificate signing requests (CSRs).
//
// @param request - ListCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCsrResponse
func (client *Client) ListCsrWithOptions(request *ListCsrRequest, runtime *dara.RuntimeOptions) (_result *ListCsrResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries certificate signing requests (CSRs).
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
// Queries a list of deployment tasks that are created.
//
// @param request - ListDeploymentJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentJobResponse
func (client *Client) ListDeploymentJobWithOptions(request *ListDeploymentJobRequest, runtime *dara.RuntimeOptions) (_result *ListDeploymentJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries a list of deployment tasks that are created.
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
// Queries the basic information about a deployment task. After you create a deployment task, you can call this operation to obtain the basic information about the deployment task, including the instance ID, type, and name of the certificate.
//
// @param request - ListDeploymentJobCertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentJobCertResponse
func (client *Client) ListDeploymentJobCertWithOptions(request *ListDeploymentJobCertRequest, runtime *dara.RuntimeOptions) (_result *ListDeploymentJobCertResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the basic information about a deployment task. After you create a deployment task, you can call this operation to obtain the basic information about the deployment task, including the instance ID, type, and name of the certificate.
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
// Queries the cloud resources of cloud services in a deployment task.
//
// @param request - ListDeploymentJobResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentJobResourceResponse
func (client *Client) ListDeploymentJobResourceWithOptions(request *ListDeploymentJobResourceRequest, runtime *dara.RuntimeOptions) (_result *ListDeploymentJobResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the cloud resources of cloud services in a deployment task.
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
// Queries the certificates or certificate orders of users.
//
// Description:
//
// You can call the ListUserCertificateOrder operation to query the certificates or certificate orders of users. If you set OrderType to CERT or UPLOAD, certificates are returned. If you set OrderType to CPACK or BUY, certificate orders are returned.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListUserCertificateOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserCertificateOrderResponse
func (client *Client) ListUserCertificateOrderWithOptions(request *ListUserCertificateOrderRequest, runtime *dara.RuntimeOptions) (_result *ListUserCertificateOrderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the certificates or certificate orders of users.
//
// Description:
//
// You can call the ListUserCertificateOrder operation to query the certificates or certificate orders of users. If you set OrderType to CERT or UPLOAD, certificates are returned. If you set OrderType to CPACK or BUY, certificate orders are returned.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Queries the details about the worker tasks of a deployment task. Alibaba Cloud allows you to deploy multiple certificates at a time. Therefore, a deployment task may include multiple worker tasks in multiple cloud services. A worker task refers to a task that deploys a certificate to a cloud resource in a cloud service.
//
// @param request - ListWorkerResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkerResourceResponse
func (client *Client) ListWorkerResourceWithOptions(request *ListWorkerResourceRequest, runtime *dara.RuntimeOptions) (_result *ListWorkerResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the details about the worker tasks of a deployment task. Alibaba Cloud allows you to deploy multiple certificates at a time. Therefore, a deployment task may include multiple worker tasks in multiple cloud services. A worker task refers to a task that deploys a certificate to a cloud resource in a cloud service.
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Submits a renewal application for an issued certificate.
//
// Description:
//
// You can call the RenewCertificateOrderForPackageRequest operation to submit a renewal application for a certificate only when the order of the certificate is in the expiring state. After the renewal is complete, a new certificate order whose status is pending application is generated. You must submit a certificate application for the new certificate order and install the new certificate after the new certificate is issued.
//
// >  You can call the [DescribeCertificateState](https://help.aliyun.com/document_detail/164111.html) operation to query the status of a certificate application order. If the value of the **Type*	- response parameter is **certificate**, the certificate is issued.
//
// @param request - RenewCertificateOrderForPackageRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewCertificateOrderForPackageRequestResponse
func (client *Client) RenewCertificateOrderForPackageRequestWithOptions(request *RenewCertificateOrderForPackageRequestRequest, runtime *dara.RuntimeOptions) (_result *RenewCertificateOrderForPackageRequestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Submits a renewal application for an issued certificate.
//
// Description:
//
// You can call the RenewCertificateOrderForPackageRequest operation to submit a renewal application for a certificate only when the order of the certificate is in the expiring state. After the renewal is complete, a new certificate order whose status is pending application is generated. You must submit a certificate application for the new certificate order and install the new certificate after the new certificate is issued.
//
// >  You can call the [DescribeCertificateState](https://help.aliyun.com/document_detail/164111.html) operation to query the status of a certificate application order. If the value of the **Type*	- response parameter is **certificate**, the certificate is issued.
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
// Signs a private certificate in a certificate application repository.
//
// Description:
//
// You can call the Sign operation to sign a private certificate in a certificate application repository.
//
// ### Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SignResponse
func (client *Client) SignWithOptions(request *SignRequest, runtime *dara.RuntimeOptions) (_result *SignResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
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
// Signs a private certificate in a certificate application repository.
//
// Description:
//
// You can call the Sign operation to sign a private certificate in a certificate application repository.
//
// ### Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Updates the private key of a certificate signing request (CSR).
//
// @param request - UpdateCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCsrResponse
func (client *Client) UpdateCsrWithOptions(request *UpdateCsrRequest, runtime *dara.RuntimeOptions) (_result *UpdateCsrResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Updates the private key of a certificate signing request (CSR).
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
// Updates a deployment task.
//
// @param request - UpdateDeploymentJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeploymentJobResponse
func (client *Client) UpdateDeploymentJobWithOptions(request *UpdateDeploymentJobRequest, runtime *dara.RuntimeOptions) (_result *UpdateDeploymentJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Updates a deployment task.
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
// Updates the status of a deployment task.
//
// @param request - UpdateDeploymentJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeploymentJobStatusResponse
func (client *Client) UpdateDeploymentJobStatusWithOptions(request *UpdateDeploymentJobStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateDeploymentJobStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Updates the status of a deployment task.
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
// Rolls back or executes a worker task in a deployment task.
//
// @param request - UpdateWorkerResourceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkerResourceStatusResponse
func (client *Client) UpdateWorkerResourceStatusWithOptions(request *UpdateWorkerResourceStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkerResourceStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Rolls back or executes a worker task in a deployment task.
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
// # Uploads a certificate signing request (CSR) file
//
// @param request - UploadCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadCsrResponse
func (client *Client) UploadCsrWithOptions(request *UploadCsrRequest, runtime *dara.RuntimeOptions) (_result *UploadCsrResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # Uploads a certificate signing request (CSR) file
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
// Uploads a certificate.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UploadUserCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadUserCertificateResponse
func (client *Client) UploadUserCertificateWithOptions(request *UploadUserCertificateRequest, runtime *dara.RuntimeOptions) (_result *UploadUserCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Uploads a certificate.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Verifies the signature of a private certificate in a certificate application repository.
//
// Description:
//
// You can call the Verify operation to verify the signature of a private certificate in a certificate application repository.
//
// ### Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - VerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyResponse
func (client *Client) VerifyWithOptions(request *VerifyRequest, runtime *dara.RuntimeOptions) (_result *VerifyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
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
// Verifies the signature of a private certificate in a certificate application repository.
//
// Description:
//
// You can call the Verify operation to verify the signature of a private certificate in a certificate application repository.
//
// ### Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
